# Progress

## Snapshot (2026-02-12)

### Status
- Implementation started.
- Loader TODO has been replaced with a working MVP implementation in `internal/driver/loader.go`.

### Completed In This Pass
- Implemented request config propagation of `GoVersion`:
  - `internal/driver/config.go`
- Implemented first-pass sequential driver loader without `packages.Load`:
  - `internal/driver/loader.go`
  - Pattern normalization (`./...`, `...`, `file=`, absolute paths, import paths)
  - Pattern chunking (`safeArgMax`) and sequential chunk processing
  - Recursive pattern loading for directory and import-prefix patterns
  - Package resolution from:
    - module workspace (`go.mod`)
    - `GOROOT/src`
    - `GOPATH/src`
  - Recursive dependency loading via parsed imports
  - Overlay-aware import parsing
  - `builtin` package inclusion in every response
  - `NotHandled=false` response path with partial package-level errors
- Added tests for the new baseline behavior:
  - `internal/driver/loader_test.go`
  - `ConfigFromDriverRequest` keeps `GoVersion`
  - Pattern normalization
  - Pattern chunking
  - Module package load + builtin inclusion
  - Overlay import override behavior

### Validation
- Ran: `go test ./...`
- Result: PASS

## PLAN.md Tracking

### 2) Driver Core Behavior (internal/driver)
- [x] [P1] Implement business logic in `internal/driver/loader.go` (initial MVP baseline)
- [x] [P1] Execute loading flow sequentially (single-threaded runtime)
- [x] [P1] Implement pattern chunking and response merge baseline
- [x] [P1] Ensure path normalization baseline for workspace-resolved files
- [x] [P1] Always include `builtin` in responses
- [ ] [P1] Full parity package graph assembly (`ForTest`, `DepsErrors`, `TypeErrors`, richer `Module` behavior)
- [ ] [P1] Complete env/build-flags/tests parity behavior
- [ ] [P1] Gold-output parity with reference traces

### 3) Test/Validation and Trace Parity
- [x] [P1] Added initial loader tests for core behavior
- [ ] [P1] Add fixture parity tests against `docs/reference/driver-requests-responses.md`
- [ ] [P1] Add broader integration tests for `tests=true/false`, `./...`, `builtin`, and error cases

## Next Implementation Targets
1. Implement `tests=true` package variants (`ForTest` and related IDs) for closer `mode=32287` behavior.
2. Improve package metadata parity (`DepsErrors`, `TypeErrors`, richer module/error shaping).
3. Add golden fixture parity tests from `docs/reference/driver-requests-responses.md`.
4. Tighten edge-case handling for import-path-only module dependencies and no-match semantics.
