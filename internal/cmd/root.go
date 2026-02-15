package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/tools/go/packages"

	"github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
	"github.com/Better-Go-Playground/go-packages-driver-wasm/internal/fakenet"
	"github.com/Better-Go-Playground/go-packages-driver-wasm/internal/jsonrpc"
	"github.com/Better-Go-Playground/go-packages-driver-wasm/internal/models"
)

type Options struct {
	// RunAsServer controls whether to run the app in server mode.
	RunAsServer bool

	// SocketPath path to a UNIX socket to listen for incoming connections.
	//
	// If not set - server listens to stdio.
	// Has effect only if RunAsServer.
	SocketPath string

	// Args is a list of positional arguments.
	//
	// Used for patterns when not in server mode.
	Args []string
}

// Main starts the application.
func Main(opts Options) error {
	ctx, cancelFn := newApplicationContext()
	defer cancelFn()

	ver, err := driver.GoVersionFromRuntime()
	if err != nil {
		return fmt.Errorf("failed to build runtime info")
	}

	if opts.RunAsServer {
		return startServer(ctx, serverOpts{
			goVersion:  ver,
			socketPath: opts.SocketPath,
		})
	}

	// If in regular mode - get patterns from args and write a response into stdout.
	return runAsCommand(ctx, ver, opts.Args)
}

var emptyResponse = &packages.DriverResponse{NotHandled: true}

func runAsCommand(ctx context.Context, ver driver.GoVersion, patterns []string) error {
	var req packages.DriverRequest
	err := json.NewDecoder(os.Stdin).Decode(&req)
	if err != nil {
		dumpResponse(emptyResponse)
		return fmt.Errorf("cannot parse request from stdin: %w", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		dumpResponse(emptyResponse)
		return err
	}

	cfg := driver.ConfigFromDriverRequest(ver, cwd, req)
	loader := driver.NewLoader(cfg)
	rsp, err := loader.Load(ctx, patterns)
	if err != nil {
		dumpResponse(emptyResponse)
		return fmt.Errorf("driver.Load: %w", err)
	}

	dumpResponse(rsp)
	return nil
}

func dumpResponse(rsp *packages.DriverResponse) {
	err := json.NewEncoder(os.Stdout).Encode(rsp)
	if err != nil {
		log.Printf("can't write response to stdout: %s", err)
	}
}

type serverOpts struct {
	goVersion  driver.GoVersion
	socketPath string
}

func startServer(ctx context.Context, opts serverOpts) error {
	ver := opts.goVersion
	log.Printf(
		"starting driver server (compiler=%s arch=%s goVersion=%d)",
		ver.Compiler, ver.Arch, ver.GoMinorVersion,
	)

	listener := jsonrpc.NewListener(map[string]jsonrpc.RequestHandler{
		"goPackageDriver/query": createHandler(ver),
	})

	if opts.socketPath == "" {
		// serve stdio
		conn := fakenet.NewConn("stdio", os.Stdin, os.Stdout)
		defer conn.Close()
		return listener.ServeStream(ctx, conn)
	}

	log.Printf("listening for socket %q", opts.socketPath)
	l, err := net.Listen("unix", opts.socketPath)
	if err != nil {
		return fmt.Errorf("listen error: %w", err)
	}

	go func() {
		<-ctx.Done()
		l.Close()
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			if ctx.Err() != nil {
				break
			}

			return fmt.Errorf("accept error: %w", err)
		}

		go func() {
			if err := listener.ServeStream(ctx, conn); err != nil {
				log.Printf("failed to handle request: %s", err)
			}
		}()
	}

	return nil
}

func createHandler(ver driver.GoVersion) jsonrpc.RequestHandler {
	return jsonrpc.NewHandler(func(ctx context.Context, req *models.DriverServerRequest) (*packages.DriverResponse, error) {
		cfg := driver.ConfigFromDriverRequest(ver, req.WorkDir, req.DriverRequest)
		loader := driver.NewLoader(cfg)
		rsp, err := loader.Load(ctx, req.Patterns)
		if err != nil {
			log.Printf("Error: driver.Load: %s", err)
			return nil, err
		}

		return rsp, nil
	})
}
