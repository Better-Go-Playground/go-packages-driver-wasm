package main

import (
	"flag"
	"log"

	"github.com/Better-Go-Playground/go-packages-driver-wasm/internal/cmd"
)

func main() {
	var opts cmd.Options
	flag.BoolVar(&opts.RunAsServer, "serve", false, "Start the driver in a server mode. Used for WebAssembly.")
	flag.Parse()

	opts.Args = flag.Args()
	if err := cmd.Main(opts); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
