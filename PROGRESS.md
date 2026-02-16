# Progress

## Snapshot (2026-02-15)

### Status

- MVP loader behavior is still in place and covered by current unit tests.
- Live gopls validation uncovered a high-priority parity bug: module-mode external imports are not resolved.

### New Bug: External Imports Unresolved In gopls

- Repro environment:
  - custom gopls built with `scripts/build-gopls.sh` (`-tags pipetransport`)
  - baseline run via `scripts/lsp-drv-standard.sh`
  - custom driver run via `scripts/lsp-drv-custom.sh`
  - sample workspace: `/home/x1unix/tmp/book`
- Trace artifacts collected in `logs/`:
  - baseline: `rpc-expected.trace.jsonl`
  - custom: `rpc-got.trace.jsonl`
- Editor diagnostics observed with custom driver:
  - `could not import github.com/jackc/pgx/v5/stdlib (go/packages driver could not load "github.com/jackc/pgx/v5/stdlib")`
  - `could not import github.com/pressly/goose/v3 (go/packages driver could not load "github.com/pressly/goose/v3")`

NOTE: do not run `scripts/ldp-drv*` scripts as they start an interactive neovim session.
Instead, ask the operator (human) to collect new samples.

### Investigation Findings (Code-Level)

- `internal/driver/loader.go` currently resolves imports from:
  - main module path match
  - current package module path match (`moduleForDir(srcDir)`)
  - `GOPATH/src`
  - `GOROOT/src`
- It does not yet resolve non-stdlib module dependencies from module cache locations (`GOMODCACHE` or `$GOPATH/pkg/mod`).
- As a result, imports like `github.com/jackc/pgx/v5/stdlib` fail when not in the main module and not present in `GOPATH/src`, producing package load errors and BrokenImport diagnostics.

### Investigation And Bugfix Plan

1. Add trace diff analysis focused on failed package IDs/imports between `rpc-expected.trace.jsonl` and `rpc-got.trace.jsonl`.
2. Extend runtime env handling for module cache resolution (`GOMODCACHE`, fallback `$GOPATH/pkg/mod`).
3. Implement module-aware import resolution for external dependencies:
   - derive candidate module path prefixes from import path
   - map module path to version from nearest `go.mod` requirements
   - resolve module root directory in module cache
4. Add on-demand transitive module requirement loading (parse dependency `go.mod` when needed, with caching).
5. Add deterministic tests:
   - resolver unit tests for external module imports
   - integration test that reproduces BrokenImport scenario
   - fixture parity checks against reference traces
6. Re-run validation:
   - `go test ./...`
   - side-by-side gopls smoke test with standard vs custom driver scripts

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
