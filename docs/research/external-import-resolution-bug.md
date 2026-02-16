# External Import Resolution Bug (2026-02-15)

## Summary

During live gopls smoke testing, the custom driver fails to resolve external module imports in a module-based workspace. This causes BrokenImport diagnostics in the editor for packages that the built-in gopls driver resolves successfully.

## Reproduction Setup

- Build custom gopls:
  - `scripts/build-gopls.sh`
- Run baseline (built-in driver behavior):
  - `TEST_DIR=/home/x1unix/tmp/book scripts/lsp-drv-standard.sh`
- Run custom driver:
  - `TEST_DIR=/home/x1unix/tmp/book scripts/lsp-drv-custom.sh`
- Compare traces:
  - expected: `logs/rpc-expected.trace.jsonl`
  - got: `logs/rpc-got.trace.jsonl`

## User-Visible Symptoms

Editor diagnostics include:

- `could not import github.com/jackc/pgx/v5/stdlib (go/packages driver could not load "github.com/jackc/pgx/v5/stdlib")`
- `could not import github.com/pressly/goose/v3 (go/packages driver could not load "github.com/pressly/goose/v3")`

## Current Root-Cause Hypothesis

In `internal/driver/loader.go`, `resolveImport` handles:

- main-module package paths
- package paths within the current source module (`moduleForDir(srcDir)`)
- `GOPATH/src`
- `GOROOT/src`

It does not currently resolve module dependencies from module cache paths (`GOMODCACHE` or fallback `$GOPATH/pkg/mod`). Therefore, external imports that are not part of the main module or GOPATH workspace are returned as unresolved.

## Investigation Plan

1. Diff `rpc-expected.trace.jsonl` and `rpc-got.trace.jsonl` to identify exact missing package nodes and import edge differences.
2. Inspect incoming `driverRequest.Env` values in traces for module-related vars (`GOMOD`, `GOMODCACHE`, `GOPATH`, `GOFLAGS`).
3. Confirm failing import paths and first failure point in the recursive dependency loading flow.
4. Add a focused failing test case that reproduces unresolved external module imports.

## Bugfix Plan

1. Add module-cache resolution inputs to runtime:
   - explicit `GOMODCACHE`
   - fallback to `$GOPATH/pkg/mod` when `GOMODCACHE` is missing
2. Implement module-aware external import resolution:
   - derive module path candidates from import path (longest-prefix strategy)
   - map module path to version from nearest parsed `go.mod` requirements
   - locate module root in module cache and append package subpath
3. Add on-demand dependency module requirement parsing (`go.mod`) with caching for transitive imports.
4. Preserve current constraints: no `packages.Load`, no external process calls, single-threaded loading flow.
5. Add regression tests and trace parity checks.

## Acceptance Criteria

- External imports in `/home/x1unix/tmp/book` resolve without BrokenImport diagnostics under the custom driver.
- Driver response includes package entries for previously failing imports.
- New resolver tests pass and prevent regressions.
- `go test ./...` remains green.
