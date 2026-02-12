# Project Implementation Plan

## Goals and Assumptions

### Goal

Deliver a Go packages driver compatible with gopls, producing outputs that match the "gold" request/response pairs, without spawning external processes (e.g., the `go` toolchain) due to environment constraints.

### Assumptions

- Runtime is Unix-like only; Windows is explicitly out of scope.
- Overlay paths are absolute.
- Overlay content consists of Go source files.
- `GoVersion` is injected by the server layer (`internal/cmd/root.go`).
- Driver logic should return concrete errors; server layer handles `NotHandled` behavior for clients.

## Key References

- `docs/reference/driver-requests-responses.md` ("gold" request/response examples from traces)
- `docs/research/driver-protocol-notes.md` (protocol and gopls behavior)
- `docs/research/trace-format.md` and `docs/research/trace-analysis.md` (trace shape and analysis)
- `internal/models/request.go` (`DriverServerRequest` envelope used by the server)
- `internal/driver/loader.go` (business logic entry point)

## Milestones

1. Golden contract and acceptance criteria
2. Driver core behavior (business logic)
3. Test/validation and trace parity
4. Documentation and integration guidance

## Implementation Phases

### Phase 1 (MVP Correctness)

- [ ] Support `GOPATH` and `go.mod` workspaces without `replace`.
- [ ] Prioritize correctness over optimization; no caching/coalescing in this phase.
- [ ] Determine and target the most frequent `packages.LoadMode` bitmask combinations from traces.

### Phase 2 (Module Extensions)

- [ ] Add `replace` support in `go.mod`.
- [ ] Add `vendor` mode behavior parity.

### Phase 3 (Workspace Extensions)

- [ ] Add `go.work` support and validate parity for workspace resolution.

## Already Implemented (No Action)

- JSON marshal/unmarshal, transport, and routing in `internal/jsonrpc`.
- Server/stdio wiring in `internal/cmd/root.go`.
- Request cancellation in `internal/jsonrpc/handler.go`.

## Plan

### 1) Golden Contract and Acceptance Criteria

- [ ] [P1] Treat `docs/reference/driver-requests-responses.md` as the source of truth for inputs and required outputs.
- [ ] [P1] Align the request envelope with `internal/models/request.go` (`DriverServerRequest`).
- [ ] [P1] Document acceptance criteria: "gold" output parity, `builtin` always present, absolute paths anchored to `PWD`, and `NotHandled` semantics.

### 2) Driver Core Behavior (internal/driver)

- [ ] [P1] Implement business logic in `internal/driver/loader.go`.
- [ ] [P1] Resolve env, build flags, tests, and overlay handling for each request.
- [ ] [P1] Implement pattern chunking and response merge semantics (including `NotHandled` propagation).
- [ ] [P1] Build package graph assembly: `Roots`, `Packages`, `Imports`, `ForTest`, `Module`, `GoVersion`, `DepsErrors`.
- [ ] [P1] Ensure path normalization and `PWD` anchoring for workspace files.
- [ ] [P1] Always include `builtin` in responses; handle stdlib and no-match cases.
- [ ] [P1] Use `/usr/lib/go/src/cmd/go/internal/list/list.go` as behavioral reference, but keep implementation compact (avoid copy-paste if possible).

**Algorithm Outline (Loader.Load)**

1. Build a request-scoped runtime view from `Config`: `Dir`, env map, build flags, tests, overlay, `Mode`, and `GoVersion`.
2. Normalize incoming patterns: resolve relative patterns against `Config.Dir` and preserve `builtin`.
3. Chunk patterns to stay within a safe command-line size limit, mirroring gopls behavior.
4. For each chunk, execute the loader pipeline: resolve module/workspace context and go env values, enumerate packages and file lists, apply overlays, and populate fields required by `Mode` and the gold outputs.
5. Merge chunk responses: if any chunk is `NotHandled`, return `NotHandled` overall; reconcile imports and roots.
6. Finalize response: always include `builtin` and ensure absolute paths anchored to `Config.Dir`.
7. Return a `packages.DriverResponse` that matches the golden outputs.

### 3) Test/Validation and Trace Parity

- [ ] [P1] Create fixture tests using "gold" request/response pairs from `docs/reference/driver-requests-responses.md`.
- [ ] [P1] Add integration tests for overlays, tests=true/false, and workspace patterns (e.g., `./...`, `builtin`).
- [ ] [P1] Compare driver outputs with trace analysis summaries (`docs/research/trace-analysis.md`).
- [ ] [P2] Add parity tests for `go.mod` `replace` and `vendor`.
- [ ] [P3] Add parity tests for `go.work` multi-module workspace behavior.

### 4) Documentation and Integration (Non-WASM)

- [ ] [P1] Update README with driver setup steps and required env variables.
- [ ] [P1] Document troubleshooting and expected error modes.
- [ ] [P2] Document `replace` and `vendor` support limits/behavior.
- [ ] [P3] Document `go.work` support and constraints.

## Open Questions

- [ ] [P1] Load modes: confirm the most frequent `packages.LoadMode` bitmask combinations from `docs/research/traces/` and define mandatory field population per mode.
