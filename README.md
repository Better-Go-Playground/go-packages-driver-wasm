# Go Packages Driver

This repo contains an alternative Go packages driver for [gopls](https://github.com/Better-Go-Playground/tools/tree/gopls-wasm-support) WebAssembly port.

Goal is to provide a portable driver that can work in WASM and browser environments.

## Current status

Project is at early stage of development. Work in progress.

Resolver support status:

- [x] Stdlib
- [x] GOMODCACHE
- [ ] `replace` in `go.mod`
- [ ] `vendor/` directory
- [ ] `go.work` monorepos

## Differences

Classic Go packages uses command-line arguments to pass query patterns and process enviromnent variables. \
As WebAssembly doesn't support spawning external processes, this driver uses a custom RPC extension to provide that context in a query.

See: the [`DriverRequestEnvelope`](./internal/server/server.go) type.

## References

- [gopls traces and research docs](./docs/research/README.md)
- [Packages protocol spec](https://pkg.go.dev/golang.org/x/tools/go/packages#hdr-The_driver_protocol)
- [Go packages driver overview](https://github.com/bazel-contrib/rules_go/wiki/Editor-and-tool-integration#driver-interface)
