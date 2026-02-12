# Project Implementation Plan

## Goal

Deliver a Go packages driver compatible with gopls, producing outputs that match the "gold" request/response pairs, without spawning external processes (e.g., the `go` toolchain) due to environment constraints.

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

## Already Implemented (No Action)

- JSON marshal/unmarshal, transport, and routing in `internal/jsonrpc`.
- Server/stdio wiring in `internal/cmd/root.go`.
- Request cancellation in `internal/jsonrpc/handler.go`.

## Plan

### 1) Golden Contract and Acceptance Criteria

- [ ] Treat `docs/reference/driver-requests-responses.md` as the source of truth for inputs and required outputs.
- [ ] Align the request envelope with `internal/models/request.go` (`DriverServerRequest`).
- [ ] Document acceptance criteria: "gold" output parity, `builtin` always present, absolute paths anchored to `PWD`, and `NotHandled` semantics.

### 2) Driver Core Behavior (internal/driver)

- [ ] Implement business logic in `internal/driver/loader.go`.
- [ ] Resolve env, build flags, tests, and overlay handling for each request.
- [ ] Implement pattern chunking and response merge semantics (including `NotHandled` propagation).
- [ ] Build package graph assembly: `Roots`, `Packages`, `Imports`, `ForTest`, `Module`, `GoVersion`, `DepsErrors`.
- [ ] Ensure path normalization and `PWD` anchoring for workspace files.
- [ ] Always include `builtin` in responses; handle stdlib and no-match cases.
- [ ] Use `/usr/lib/go/src/cmd/go/internal/list/list.go` as behavioral reference, but keep implementation compact (avoid copy-paste if possible).

**Algorithm Outline (Loader.Load)**

1. Build a request-scoped runtime view from `Config`: `Dir`, env map, build flags, tests, overlay, `Mode`, and `GoVersion`.
2. Normalize incoming patterns: resolve relative patterns against `Config.Dir` and preserve `builtin`.
3. Chunk patterns to stay within a safe command-line size limit, mirroring gopls behavior.
4. For each chunk, execute the loader pipeline: resolve module/workspace context and go env values, enumerate packages and file lists, apply overlays, and populate fields required by `Mode` and the gold outputs.
5. Merge chunk responses: if any chunk is `NotHandled`, return `NotHandled` overall; reconcile imports and roots.
6. Finalize response: always include `builtin` and ensure absolute paths anchored to `Config.Dir`.
7. Return a `packages.DriverResponse` that matches the golden outputs.

### 3) Test/Validation and Trace Parity

- [ ] Create fixture tests using "gold" request/response pairs from `docs/reference/driver-requests-responses.md`.
- [ ] Add integration tests for overlays, tests=true/false, and workspace patterns (e.g., `./...`, `builtin`).
- [ ] Compare driver outputs with trace analysis summaries (`docs/research/trace-analysis.md`).

### 4) Documentation and Integration (Non-WASM)

- [ ] Update README with driver setup steps and required env variables.
- [ ] Document troubleshooting and expected error modes.

## Research Gaps to Fill

- [ ] Overlay edge cases: large overlays, binary data, and mixed absolute/relative paths.
- [ ] Modes coverage: confirm field population across different `packages.LoadMode` combinations.
- [ ] Go version selection: when `GoVersion` should be the runtime minor vs `0`.
- [ ] Module/workspace variants: `go.work`, `vendor`, `GOPATH`-mode, and replace directives.
- [ ] Windows path behavior and path separator expectations.
- [ ] Error and partial-result semantics: when to return `NotHandled` vs. errors.
- [ ] Performance constraints: caching strategy, request coalescing, and parallel chunk behavior.
