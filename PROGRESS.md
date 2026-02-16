# Progress

## Snapshot (2026-02-15)

### Status

- MVP loader behavior is still in place and covered by current unit tests.
- Initial high-priority parity bug (module-mode external imports unresolved) is largely fixed; operator traces show major parity improvement with a small remaining delta.

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

### Trace Diff Findings (Baseline vs Custom)

- Comparison used existing artifacts only:
  - baseline: `logs/rpc-expected.trace.jsonl`
  - custom: `logs/rpc-got.trace.jsonl`
  - NOTE respected: no `scripts/ldp-drv*` execution and no new interactive sample capture.
- Top-level response delta:
  - baseline: `Roots=9`, `Packages=340`, package entries with `Errors`: `0`
  - custom: `Roots=7`, `Packages=202`, package entries with `Errors`: `24`
  - baseline trace has 2 equivalent response entries; custom trace has 1 response entry.
- Package graph parity delta:
  - package IDs missing in custom vs baseline: `146` (`102` external-module IDs, `44` stdlib/transitive IDs)
  - package IDs only in custom vs baseline: `8` (all unresolved `golang.org/x/...` IDs)
  - `16` overlapping package IDs regress to placeholders in custom (`GoFiles=0`, `Imports=0`, `Errors=[cannot resolve import path ...]`)
  - placeholder examples: `github.com/jackc/pgx/v5/stdlib`, `github.com/pressly/goose/v3`, `github.com/jackc/pgx/v5/pgxpool`, `github.com/georgysavva/scany/v2/pgxscan`
- Failure boundary evidence:
  - workspace package `github.com/x1unix/thoughtly-ticket-booking/internal/booking` is identical in both traces (`GoFiles=4`, `Imports=9`, `Errors=0`)
  - divergence starts when resolving external dependencies imported by workspace packages.
- Additional parity gaps observed:
  - `tests=true` root variants missing in custom:
    - `github.com/x1unix/thoughtly-ticket-booking/tests [github.com/x1unix/thoughtly-ticket-booking/tests.test]`
    - `github.com/x1unix/thoughtly-ticket-booking/tests.test`
  - GOROOT vendored import mapping mismatch:
    - baseline maps imports like `golang.org/x/net/http/httpguts` to package ID `vendor/golang.org/x/net/http/httpguts`
    - custom maps same import to ID `golang.org/x/net/http/httpguts`, producing unresolved `golang.org/x/...` placeholders instead of matching `vendor/...` stdlib package IDs.
- Env propagation sanity check:
  - both traces carry `GOMOD`, `GOPATH`, `GOROOT`, and `GOMODCACHE=/home/x1unix/go/pkg/mod`
  - unresolved external imports are therefore consistent with resolver behavior gaps, not missing env values in the request.

### Initial Investigation Findings (Pre-fix)

- `internal/driver/loader.go` currently resolves imports from:
  - main module path match
  - current package module path match (`moduleForDir(srcDir)`)
  - `GOPATH/src`
  - `GOROOT/src`
- It does not yet resolve non-stdlib module dependencies from module cache locations (`GOMODCACHE` or `$GOPATH/pkg/mod`).
- GOROOT vendored imports (`golang.org/x/...` within stdlib packages) are not canonicalized to `vendor/golang.org/x/...` package IDs, causing additional unresolved stdlib dependency nodes.
- As a result, imports like `github.com/jackc/pgx/v5/stdlib` fail when not in the main module and not present in `GOPATH/src`, producing package load errors and BrokenImport diagnostics.

### Investigation And Bugfix Plan

1. Add trace diff analysis focused on failed package IDs/imports between `rpc-expected.trace.jsonl` and `rpc-got.trace.jsonl`.
2. Extend runtime env handling for module cache resolution (`GOMODCACHE`, fallback `$GOPATH/pkg/mod`).
3. Implement module-aware import resolution for external dependencies:
   - derive candidate module path prefixes from import path
   - map module path to version from nearest `go.mod` requirements
   - resolve module root directory in module cache
4. Canonicalize stdlib vendored imports so `golang.org/x/...` edges in GOROOT packages resolve to `vendor/golang.org/x/...` package IDs.
5. Add on-demand transitive module requirement loading (parse dependency `go.mod` when needed, with caching).
6. Add deterministic tests:
   - resolver unit tests for external module imports
   - resolver unit tests for GOROOT vendored `golang.org/x/...` import ID mapping
   - integration test that reproduces BrokenImport scenario
   - fixture parity checks against reference traces
7. Re-run validation:
   - `go test ./...`
   - side-by-side gopls smoke test with standard vs custom driver scripts (can be done only by human).

### Fix Implementation Progress (This Pass)

- Implemented module-cache-aware import resolution in `internal/driver/loader.go`:
  - runtime now discovers module cache roots from `GOMODCACHE` with fallback to `$GOPATH/pkg/mod`
  - import resolver now resolves external module imports via nearest module `go.mod` `require` versions
  - module cache path resolution uses escaped module path/version (`x/mod/module` escaping)
- Implemented on-demand dependency `go.mod` parsing with caching:
  - parsed `require` maps are cached by `go.mod` path
  - dependency module `go.mod` files are parsed as dependencies are traversed, enabling transitive resolution from dependency imports
- Implemented GOROOT vendored import canonicalization:
  - imports like `golang.org/x/net/http/httpguts` from stdlib sources now resolve to `GOROOT/src/vendor/...`
  - package IDs for these imports now canonicalize to `vendor/golang.org/x/...` instead of unresolved `golang.org/x/...` placeholders
- Updated module metadata shaping:
  - module records now preserve `Main` accurately (`true` only for the workspace main module)
- Added regression coverage in `internal/driver/loader_test.go`:
  - `TestLoaderResolvesExternalImportsFromGoModCache`
  - `TestLoaderResolvesTransitiveImportsUsingDependencyGoMod`
  - `TestResolveImportCanonicalizesGorootVendorID`

### Validation (This Pass)

- Ran: `go test ./...`
- Result: PASS
- NOTE respected: no `scripts/ldp-drv*` execution.
- Pending human validation: side-by-side gopls smoke test trace capture (`scripts/lsp-drv-standard.sh` vs `scripts/lsp-drv-custom.sh`).

### Post-Fix Smoke Validation (Operator Traces)

- Operator provided new artifacts in `logs/after-bugfix`:
  - baseline: `logs/after-bugfix/rpc-expected.trace.jsonl`
  - custom: `logs/after-bugfix/rpc-got.trace.jsonl`
  - NOTE respected: no interactive script execution by agent.
- Top-level parity improved significantly:
  - custom package count: `202 -> 314` (baseline still `340`)
  - custom package error entries: `24 -> 5`
  - missing package IDs vs baseline: `146 -> 28`
  - extra package IDs vs baseline: `8 -> 2`
- Confirmed fixed in traces:
  - external imports now resolved for previously failing roots (`github.com/jackc/pgx/v5/stdlib`, `github.com/pressly/goose/v3`, `github.com/jackc/pgx/v5/pgxpool`, `github.com/georgysavva/scany/v2/pgxscan`)
  - GOROOT vendored mappings now match baseline (`golang.org/x/...` imports map to `vendor/golang.org/x/...` IDs)
- Remaining parity issues from `logs/after-bugfix`:
  - unresolved package placeholders remain for:
    - `github.com/klauspost/compress/{flate,gzip,zlib}`
    - `github.com/mattn/go-isatty`
    - `golang.org/x/sync/semaphore`
  - `tests=true` root variants are still missing in custom:
    - `github.com/x1unix/thoughtly-ticket-booking/tests [github.com/x1unix/thoughtly-ticket-booking/tests.test]`
    - `github.com/x1unix/thoughtly-ticket-booking/tests.test`

### Follow-up Fix Progress (This Pass)

- Started second-pass resolver hardening in `internal/driver/loader.go` to address remaining external mismatch:
  - module-cache version selection now prefers main-module requirements first, then source-module requirements
  - when required version is absent in local module cache, resolver falls back to the best available cached version for that module path
  - added version comparison logic for fallback selection (`x/mod/semver`)
- Added targeted regression coverage in `internal/driver/loader_test.go`:
  - `TestLoaderPrefersMainModuleVersionsForDependencies`
- Validation after follow-up changes:
  - Ran: `go test ./...`
  - Result: PASS
- Pending human validation:
  - re-run smoke traces to confirm remaining 5 unresolved imports are eliminated and re-check package graph deltas.

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
