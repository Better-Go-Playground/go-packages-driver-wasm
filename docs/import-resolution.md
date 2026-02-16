# Import Resolution Guide

This document describes how import resolution works in the custom packages driver.

Primary implementation: `internal/driver/loader.go`.

## Goals

- Resolve package imports without spawning external processes.
- Match gopls built-in driver behavior closely in module-mode workspaces.
- Keep behavior deterministic and testable.

## Resolution Order

For `resolveImport(importPath, srcDir)` the resolver uses this order:

1. `builtin`
2. Main module path match
3. Source module path match (`moduleForDir(srcDir)`)
4. Source-dir-aware vendor mapping (GOROOT `vendor/`)
5. `GOPATH/src`
6. `GOROOT/src`
7. Module cache lookup (`GOMODCACHE`, then `$GOPATH/pkg/mod`)

If none match, resolution fails with `cannot resolve import path`.

## Module Cache Resolution

Module cache resolution uses request/runtime env and parsed `go.mod` files:

- Cache roots:
  - `GOMODCACHE` from request env
  - process `GOMODCACHE` fallback
  - `$GOPATH/pkg/mod` fallback(s)
- Requirement sources:
  - main module requirements first
  - current source module requirements next (if different)
- Path matching:
  - derive module path candidates from import path (longest-prefix strategy)
- Version selection:
  - prefer exact required version
  - if exact cache dir is missing, use best available cached version (semver compare when possible)
- Parsed module requirements are cached by `go.mod` path.

## Vendored Imports

When `srcDir` is under `GOROOT/src`, vendor resolution is source-dir-aware:

- walk up from `srcDir` to nearest directory containing `vendor/`
- resolve import from that vendor root when present
- canonical package IDs become `vendor/<import-path>`

This is required for parity on stdlib vendored edges (for example `golang.org/x/...` referenced by stdlib packages).

## `tests=true` Variants

When `Config.Tests=true` and `NeedForTest` is requested:

- root package `<pkg>` can produce:
  - `<pkg> [<pkg>.test]` (test variant package)
  - `<pkg>.test` (synthetic test main)
- test variant includes same-package `_test.go` files
- synthetic test main imports:
  - root package key mapped to test variant ID
  - `os`, `reflect`, `testing`, `testing/internal/testdeps`

## cgo Handling

- `CGO_ENABLED` is applied to `build.Context` from request env.
- Parsed import `"C"` maps to `runtime/cgo` when cgo is enabled.
- `"C"` imports are skipped when cgo is disabled.
- Self-import guard prevents `runtime/cgo => runtime/cgo` edges.

## Notes For Future Changes

- Preserve resolution order to avoid parity regressions.
- Avoid caching solely by raw import string when `srcDir` can change vendoring outcomes.
- Keep trace-based validation in loop for behavior changes.
