package cmd

import (
	"context"
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

type serverOpts struct {
	goVersion  driver.GoVersion
	debug      bool
	traceFile  string
	socketPath string
}

func startServer(ctx context.Context, opts serverOpts) error {
	ver := opts.goVersion
	log.Printf(
		"starting driver server (compiler=%s arch=%s goVersion=%d)",
		ver.Compiler, ver.Arch, ver.GoMinorVersion,
	)

	listener := jsonrpc.NewServeMux(map[string]jsonrpc.RequestHandler{
		"goPackageDriver/query": createHandler(ver),
	})

	if opts.traceFile != "" {
		log.Printf("tracing is enabled, file: %q", opts.traceFile)
		tracer, err := newTraceInterceptor(opts.traceFile)
		if err != nil {
			return err
		}

		defer tracer.Close()
		listener.SetInterceptor(tracer)
	}

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
