package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/tools/go/packages"

	"github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
)

type Options struct {
	// RunAsServer controls whether to run the app in server mode.
	RunAsServer bool

	// SocketPath path to a UNIX socket to listen for incoming connections.
	//
	// If not set - server listens to stdio.
	// Has effect only if RunAsServer.
	SocketPath string

	// TraceFile is file to write request traces for debug purposes.
	TraceFile string

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
			traceFile:  opts.TraceFile,
		})
	}

	// If in regular mode - get patterns from args and write a response into stdout.
	return runAsCommand(ctx, ver, opts.Args)
}

var emptyResponse = &packages.DriverResponse{NotHandled: true}

func runAsCommand(ctx context.Context, ver driver.GoVersion, patterns []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		dumpResponse(emptyResponse)
		return err
	}

	type result struct {
		req packages.DriverRequest
		err error
	}

	// Listen for stdin request in background to be able to handle SIGINT
	ch := make(chan result)
	defer close(ch)
	go func() {
		var req packages.DriverRequest
		err = json.NewDecoder(os.Stdin).Decode(&req)
		if ctx.Err() != nil {
			return
		}

		ch <- result{err: err, req: req}
	}()

	var req packages.DriverRequest
	select {
	case <-ctx.Done():
		return nil
	case r := <-ch:
		if r.err != nil {
			dumpResponse(emptyResponse)
			return fmt.Errorf("cannot parse request from stdin: %w", err)
		}

		req = r.req
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
