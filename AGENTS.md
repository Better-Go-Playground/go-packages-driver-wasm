# AGENTS

**Project Summary**

- This repo is an alternative Go packages driver for the gopls Go language server, aimed at working in constrained environments, where external process execution is not possible.
- The driver avoids spawning external processes and uses a custom RPC extension to pass patterns and environment data.
- See `internal/models/request.go` for the driver request envelope type.
- See `CODE_STYLE.md` for coding style.

**Project Structure (Internal Focus)**

- `internal/driver/` driver business logic:
  - `loader.go` (public entrypoint)
  - `loader_runtime.go` (runtime/env setup)
  - `loader_state.go` (load state machine and response assembly)
  - `loader_files.go` (resolver, module cache, metadata parsing, helpers)
  - `config.go`, `runtime.go`
- `internal/models/` request envelope types (`request.go`).
- `internal/jsonrpc/` JSON-RPC transport, request handling, and cancellation.
- `internal/cmd/` app entry, stdio/server wiring (`root.go`, `ctx_*`).
- `internal/fakenet/` stdio-backed `net.Conn` shim used by the server mode.
- `cmd/` CLI entry points that invoke `internal/cmd`.
- `docs/reference/` golden request/response pairs.
- `docs/research/` traces and protocol research notes.

**Working Docs**

- `PLAN.md` — canonical implementation plan, milestones, and acceptance criteria.
- `PROGRESS.md` — current snapshot of completed work, validation status, and next targets.
- `docs/import-resolution.md` — canonical import resolver behavior and precedence guide.

**Current Focus**

- External import parity bug investigation is closed (validated in `logs/fix-stage-6`).
- Current work is parity hardening: fixture tests, broader integration tests, and remaining `replace`/full `vendor`/`go.work` behavior.

**Research Docs (start here)**

- `docs/research/README.md` indexes research files and explains the traces directory and schema files.
- `docs/research/drivertypes.ts` and `docs/research/types.ts` describe the JSONL trace schema in TypeScript.

**Driver Protocol Notes**

- Environment variables control driver discovery: `GOPACKAGESDRIVER` (path or `off`) and, for wasm, `GOPACKAGESDRIVER=wasm` plus `GOPACKAGESDRIVERADDR`.
- Requests are JSON on stdin and responses are JSON on stdout when using the external driver path.
- The wasm transport uses an IO pipe and `goPackageDriver/query` with an envelope `{ workDir, patterns, driverRequest }`.
- Paths returned by the driver should be absolute and anchored to the injected `PWD` to preserve logical working directories.
- See `docs/research/driver-protocol-notes.md` for details and related gopls source references.

**Trace Research**

- Trace files are line-delimited JSON (JSONL) under `docs/research/traces`.
- Each trace line contains exactly one `cmd` or `drv` event with top-level headers (`spanId`, `parentSpanId`, `ts`, `stack`).
- Driver overlays are base64-encoded in `DriverRequest.overlay`, and the fallback go command writes a parsed overlay file.
- See `docs/research/trace-format.md` for the exact JSON shape and field naming.
- `docs/research/trace-analysis.md` summarizes a sample trace (`drvtrace.jsonl`).

**Trace Analysis Script**

- `docs/research/scripts/analyze_drvtrace.py` parses a JSONL trace and emits a JSON summary.
- Default input path is `docs/research/drvtrace.jsonl` (override by passing a path argument).

**References**

- README links:
- `docs/research/README.md` for research index.
- gopls packages driver protocol: `golang.org/x/tools/go/packages` driver spec.
- Rules_go driver overview: `bazel-contrib/rules_go` editor/tool integration docs.
