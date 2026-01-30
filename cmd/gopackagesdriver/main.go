package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Better-Go-Playground/go-packages-driver-wasm/internal/fakenet"
	"github.com/Better-Go-Playground/go-packages-driver-wasm/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("failed to start package driver: ", err)
	}
}

func run() error {
	ctx, cancelFn := server.NewApplicationContext()
	defer cancelFn()

	ri, err := server.BuildRuntimeInfo()
	if err != nil {
		return fmt.Errorf("failed to build runtime info")
	}

	log.Printf(
		"starting driver server (compiler=%s arch=%s goVersion=%d)",
		ri.Compiler, ri.Arch, ri.GoMinorVersion,
	)
	conn := fakenet.NewConn("stdio", os.Stdin, os.Stdout)
	defer conn.Close()

	srv := server.NewService(ri)
	return srv.Listener().ServeStream(ctx, conn)
}
