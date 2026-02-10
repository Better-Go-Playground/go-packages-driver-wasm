# Trace Analysis (drvtrace.jsonl)

This note summarizes the observed behavior in `docs/research/drvtrace.jsonl` based on `docs/research/scripts/analyze_drvtrace.py`.

**Counts**
- Total entries: 35 lines (18 `drv`, 17 `cmd`).
- Parent/child links: 17 `cmd` entries are children of a `drv` span. All `drv` spans are roots (`parentSpanId` is 0 or missing).
- Errors: 1 `drv` call failed with `context canceled` for `file=/home/x1unix/prj/go-packages-driver-wasm/internal/driver/loader.go`. There are no `cmd` errors.

**Patterns Seen**
- `/home/x1unix/prj/go-packages-driver-wasm/...|builtin` (2 calls)
- `file=/home/x1unix/prj/go-packages-driver-wasm/internal/driver/loader.go` (7 calls)
- `github.com/Better-Go-Playground/go-packages-driver-wasm/cmd/gopackagesdriver|github.com/Better-Go-Playground/go-packages-driver-wasm/internal/server` (2 calls)
- `github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver` (7 calls)

**DriverRequest Summary**
- `mode` is always `32287`.
- Bits set in `mode`: `NeedName`, `NeedFiles`, `NeedCompiledGoFiles`, `NeedImports`, `NeedDeps`, `NeedTypesSizes`, `needInternalDepsErrors`, `NeedForTest`, `typecheckCgo`, `NeedModule`, `NeedEmbedFiles`.
- Bits not set in this trace: `NeedExportFile`, `NeedTypes`, `NeedSyntax`, `NeedTypesInfo`, `NeedEmbedPatterns`, `NeedTarget`.
- `tests` is always `true`.
- `build_flags`: always `-tags integration`.

**Environment (DriverRequest.Env)**
- 151 env vars are passed through to the driver.
- Notable Go-related values:
- `GOOS=js`
- `GOARCH=amd64`
- `GOROOT=/usr/lib/go`
- `GOPATH=/home/x1unix/go`
- `GOMOD=/home/x1unix/prj/go-packages-driver-wasm/go.mod`
- `GOMODCACHE=/home/x1unix/go/pkg/mod`
- `GOVERSION=go1.25.6 X:nodwarf5`
- `GOTOOLCHAIN=auto`

**Overlay Behavior**
- `req.overlay` is present in 17 of 18 `drv` calls.
- Each overlay contains exactly one file (`internal/driver/loader.go`).
- Each overlay payload is base64-encoded file contents (Go `[]byte` JSON encoding).
- Total overlay payload size across calls: 41,676 bytes (base64 text length).
- `drv.overlay.Content.replace` maps the real file to a temporary overlay file like `/tmp/gocommand-*/1-loader.go`.

**go list Invocations (`cmd`)**
- All `cmd` entries use `verb: list`.
- 4 unique argument sets, differing only by pattern list.
- Common flags across all calls:
- `-e`
- `-json=Name,ImportPath,Error,Dir,GoFiles,IgnoredGoFiles,IgnoredOtherFiles,CFiles,CgoFiles,CXXFiles,MFiles,HFiles,FFiles,SFiles,SwigFiles,SwigCXXFiles,SysoFiles,TestGoFiles,XTestGoFiles,CompiledGoFiles,Export,DepOnly,Imports,ImportMap,TestImports,XTestImports,ForTest,DepsErrors,Module,EmbedFiles`
- `-compiled=true -test=true -export=false -deps=true -find=false -pgo=off -tags integration -- <patterns>`
- Total JSON objects emitted by `go list` across all `cmd` calls: 2,473.

**DriverResponse Summary**
- `NotHandled` is never true.
- `GoVersion` values: `25` for 11 responses, `0` for 6 responses (all `file=` pattern calls).
- `Compiler|Arch`: `gc|amd64` for 14 responses, `gc|wasm` for 3 responses.
- Package counts by pattern set:
- `/...|builtin`: 153 or 157 packages
- `file=.../loader.go`: 143 packages
- `internal/driver`: 141 or 143 packages
- `cmd/gopackagesdriver|internal/server`: 151 or 155 packages
- Root counts by pattern set:
- `/...|builtin`: 6 roots
- `file=.../loader.go`: 1 root
- `internal/driver`: 1 root
- `cmd/gopackagesdriver|internal/server`: 2 roots

**Observed Package Fields (non-empty)**
- Always populated: `ID`, `Name`, `PkgPath`, `GoFiles`.
- Usually populated: `CompiledGoFiles` (missing for 17 packages), `Imports` (present on 2,163 packages).
- Sometimes populated: `IgnoredFiles` (981 packages), `OtherFiles` (323 packages).
- Not populated in this trace: `Dir`, `ExportFile`, `Target`, `Module`, `EmbedFiles`, `EmbedPatterns`, `TypeErrors`, `Errors`.

**Tooling**
- Analysis script: `docs/research/scripts/analyze_drvtrace.py`.
