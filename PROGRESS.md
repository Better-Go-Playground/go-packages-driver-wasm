# Progress

## Snapshot (2026-02-15)

### Status

- MVP loader behavior is still in place and covered by current unit tests.
- External import resolution parity bug is CLOSED.
- Final operator validation (`logs/fix-stage-6`) confirms baseline/custom parity (`missing=0`, `extra=0`, `Errors=0`) and runtime/cgo metadata alignment.

### Bug: External Imports Unresolved In gopls (Closed)

- Resolution status: CLOSED (validated in `logs/fix-stage-6`).
- NOTE: do not run `scripts/ldp-drv*` from the agent (interactive neovim); operator collects smoke traces.

### Final Outcome

- Final parity (`logs/fix-stage-6`):
  - baseline: `Roots=9`, `Packages=340`, `Errors=0`
  - custom: `Roots=9`, `Packages=340`, `Errors=0`
  - package ID delta: `missing=0`, `extra=0`
- Key closure checks:
  - previously broken external imports resolve (`pgx`, `goose`, `klauspost/compress`, `go-isatty`, `x/sync/semaphore`)
  - `tests=true` roots exist (`<pkg> [<pkg>.test]`, `<pkg>.test`)
  - `runtime/cgo` metadata matches baseline (no custom-only self-edge)

### Timeline (Condensed)

- Initial state (`logs/before-bugfix`): custom had `Packages=202`, `Errors=24`, `missing=146` vs baseline.
- Stage 1-2 (`logs/fix-stage-1`, `logs/fix-stage-2`): module-cache + vendored import fixes reduced unresolved external imports (`Errors 24 -> 0`, `missing 146 -> 15`).
- Stage 3 (`logs/fix-stage-3`): `tests=true` variants/testdeps wiring restored missing test roots and companion graph nodes.
- Stage 4-5 (`logs/fix-stage-4`, `logs/fix-stage-5`): closed vendored text/cgo gaps and reached package ID parity; one runtime/cgo self-edge metadata mismatch remained.
- Stage 6 (`logs/fix-stage-6`): self-edge cleanup validated; baseline/custom parity achieved.

### Implementation Summary

- Core resolver implementation in `internal/driver/loader.go` now includes:
  - module-cache-aware resolution (`GOMODCACHE` and `$GOPATH/pkg/mod` fallback)
  - on-demand `go.mod` requirement parsing and caching for transitive deps
  - best-available-version fallback in cache when exact version dir is missing
  - source-dir-aware GOROOT vendored import canonicalization (`vendor/...` IDs)
  - `tests=true` variant and synthetic test-main package modeling
  - cgo normalization (`"C" -> runtime/cgo`) and self-import guard
  - explicit `CGO_ENABLED` handling in build context
- Regression coverage added in `internal/driver/loader_test.go` across all fix areas (module cache, vendoring, tests variants, cgo).

### Resolver Documentation

- Durable reference for future work: `docs/import-resolution.md`.

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

1. Add golden fixture parity tests from `docs/reference/driver-requests-responses.md`.
2. Add broader integration tests for `tests=true/false`, recursive patterns, and error/no-match semantics.
3. Improve package metadata parity (`DepsErrors`, `TypeErrors`, richer module/error shaping).
4. Extend regression fixtures for vendored/stdlib and cgo edge-cases to lock current parity behavior.
