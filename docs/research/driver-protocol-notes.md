# Driver Protocol Notes (gopls go/packages)

This note consolidates driver-protocol behavior from the gopls source in `../gopls/go/packages/*`.

**Driver Discovery and Invocation**
- `GOPACKAGESDRIVER` controls the external driver path. If unset, `gopackagesdriver` is searched on `PATH`.
- If `GOPACKAGESDRIVER=off`, no external driver is used.
- Patterns are passed as positional arguments.
- `DriverRequest` is JSON on stdin. `DriverResponse` is JSON on stdout.
- `PWD` is injected into the driver environment to preserve the logical working directory (symlink-safe). The driver must emit absolute paths with the `PWD` prefix when returning files under the working directory.

**DriverRequest Schema (key points)**
- Fields: `mode`, `env`, `build_flags`, `tests`, `overlay`.
- `overlay` maps file path to file contents (encoded as base64 in JSON, due to Go `[]byte` encoding).

**DriverResponse Schema (key points)**
- `NotHandled` tells gopls to fall back to the next driver.
- `Roots` are the root package IDs (needed to reconstruct the graph).
- `Packages` includes the full set of packages, with `Imports` pointing to stub packages by ID.
- `GoVersion` is the minor Go version number used when selecting files. `0` means unknown.

**Pattern Chunking**
- Patterns are split into chunks to avoid exceeding command-line limits.
- `safeArgMax` is `Windows ARG_MAX (32767) - maxEnvSize (16384)`.
- Each chunk is processed concurrently. If any chunk returns `NotHandled`, the whole response is treated as `NotHandled`.
- Responses from multiple chunks are merged and de-duplicated.

**Wasm Driver Transport**
- If `GOPACKAGESDRIVER=wasm`, gopls uses an IO transport defined in `external_pipe.go`.
- `GOPACKAGESDRIVERADDR` must be set. Supported formats:
- `fd:3,4` for file descriptors.
- `file:///path/to/in,file:///path/to/out` for file URLs.
- The driver is invoked over a JSON-RPC-like transport using method `goPackageDriver/query`.
- The request envelope is `{ workDir, patterns, driverRequest }`.

**go list Fallback Behavior**
- Fallback uses `go list -json` with a large fixed field set and flags derived from `Config`.
- `createDriverResponse` converts `go list` JSON into `DriverResponse`.
- `PkgPath` is derived from `ID` (prefix up to the first space).
- `Imports` is constructed from `Imports` and `ImportMap`, with `C` filtered out.
- `Roots` are any packages where `DepOnly` is false.
- If `typecheckCgo` is set, `CompiledGoFiles` may be adjusted or error out when cgo processing fails.

**Relevant Source Files**
- `../gopls/go/packages/doc.go` (driver protocol description)
- `../gopls/go/packages/external.go` (request/response schema and driver discovery)
- `../gopls/go/packages/external_native.go` (native driver invocation)
- `../gopls/go/packages/external_wasm.go` (wasm driver invocation)
- `../gopls/go/packages/external_pipe.go` (IO transport)
- `../gopls/go/packages/golist.go` (go list fallback and response shaping)
- `../gopls/go/packages/debug.go` (trace emission)
