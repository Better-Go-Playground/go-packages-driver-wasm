# Project Implementation Plan

## Goals and Assumptions

### Goal

Deliver a Go packages driver compatible with gopls, producing outputs that match reference behavior without spawning external processes.

### Assumptions

- Runtime is Unix-like only; Windows is out of scope.
- Overlay paths are absolute and overlay content is Go source.
- `GoVersion` is injected by server layer (`internal/cmd/root.go`).
- Request-level failures return concrete errors; per-pattern/package failures return partial results with package-level errors.
- Driver response uses the same `GoVersion` value from `driver.Config`.
- Package/file selection respects request `GOOS`/`GOARCH`, with runtime fallback.

## Constraints

- Do not call `golang.org/x/tools/go/packages.Load` inside driver implementation.
- Keep core loading flow single-threaded.

## Key References

- `docs/reference/driver-requests-responses.md`
- `docs/research/driver-protocol-notes.md`
- `docs/import-resolution.md`
- `internal/models/request.go`
- Driver implementation entry and internals:
  - `internal/driver/loader.go`
  - `internal/driver/loader_runtime.go`
  - `internal/driver/loader_state.go`
  - `internal/driver/loader_files.go`

## Milestones

1. Golden contract and acceptance criteria - in progress
2. Driver core behavior (business logic) - mostly complete
3. Test/validation and trace parity - in progress
4. Documentation and integration guidance - in progress

## Implementation Phases

### Phase 1 (MVP Correctness)

- [x] Support `GOPATH` and `go.mod` workspaces without `replace`.
- [x] Resolve external module imports in module mode via module cache (`GOMODCACHE` or `$GOPATH/pkg/mod`) without invoking `go`.
- [x] Prioritize correctness over optimization.
- [x] Target frequent trace load mode `32287`.

### Phase 2 (Module Extensions)

- [ ] Add `replace` support in `go.mod`.
- [ ] Complete full `vendor` mode parity.
- [x] Canonicalize GOROOT vendored import IDs used by stdlib package graph.

### Phase 3 (Workspace Extensions)

- [ ] Add `go.work` support and validate workspace parity.

## Already Implemented (No Action)

- JSON-RPC transport, marshal/unmarshal, and routing in `internal/jsonrpc`.
- Server/stdio wiring in `internal/cmd/root.go`.
- Request cancellation in `internal/jsonrpc/handler.go`.

## Plan

### 1) Golden Contract and Acceptance Criteria

- [x] Treat `docs/reference/driver-requests-responses.md` as contract input/output reference.
- [x] Align request envelope with `internal/models/request.go` (`DriverServerRequest`).
- [x] Keep `builtin` always present and paths absolute/normalized.
- [ ] Add automated fixture parity assertions against reference request/response pairs.

### 2) Driver Core Behavior (internal/driver)

- [x] Implement business logic entrypoint in `internal/driver/loader.go`.
- [x] Build package graph without `packages.Load`.
- [x] Keep loading sequential (single-threaded).
- [x] Implement pattern normalization, chunking, and merge flow.
- [x] Resolve external imports from module cache with `go.mod` requirement mapping.
- [x] Implement GOROOT vendored import canonicalization.
- [x] Implement `tests=true` variant roots and synthetic test main wiring.
- [x] Implement cgo import mapping (`"C" -> "runtime/cgo"`) with self-import guard.
- [x] Always include `builtin` and handle no-match as package-level errors.
- [ ] Complete richer metadata parity (`DepsErrors`, `TypeErrors`, module/error shaping details).

### 3) Test/Validation and Trace Parity

- [x] Add regression tests for external module import resolution.
- [x] Add regression tests for vendored import ID mapping behavior.
- [x] Add regression tests for `tests=true` variant roots and cgo mapping behavior.
- [x] Validate via operator-collected smoke traces through staged parity fixes (`logs/fix-stage-1` .. `logs/fix-stage-6`).
- [ ] Add fixture parity tests from `docs/reference/driver-requests-responses.md`.
- [ ] Add broader integration tests for overlays, `tests=true/false`, recursive patterns, and no-match semantics.
- [ ] Add parity tests for `replace`, full `vendor` mode, and `go.work`.

### 4) Documentation and Integration

- [x] Document resolver behavior and precedence in `docs/import-resolution.md`.
- [x] Document bug timeline and closure in `PROGRESS.md`.
- [ ] Update README with current setup/troubleshooting guidance.
- [ ] Document unsupported/partial parity areas (`replace`, full `vendor`, `go.work`).

## Resolved Investigation

- [x] External imports unresolved in gopls (module mode) - fixed.
- [x] Tests variant roots missing under `tests=true` - fixed.
- [x] Remaining runtime/cgo metadata mismatch - fixed.
- [x] Final smoke parity reached in `logs/fix-stage-6` (`missing=0`, `extra=0`, `Errors=0`).

## Next Implementation Targets

1. Add golden fixture parity tests from `docs/reference/driver-requests-responses.md`.
2. Add broader integration tests for recursive patterns and error/no-match semantics.
3. Improve metadata parity for `DepsErrors`, `TypeErrors`, and module/error shaping.
4. Implement/document `replace`, full `vendor` mode parity, and `go.work` support.
