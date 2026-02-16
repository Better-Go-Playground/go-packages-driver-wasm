# Research

This folder contains Go package driver wire format and gopls behavior research docs.

Below is description of each document.

## Files

### `external-import-resolution-bug.md`

Live-test bug report for unresolved external module imports when running gopls with the custom driver.
Includes repro scripts, observed diagnostics, root-cause hypothesis, and investigation/bugfix plan.

### `traces`

The @./traces directory contains "trace" results of how gopls calls its built-in package driver which is an adapter for `go` command.
Trace files are used to research and build a compatible third-party Go package driver.

The trace file format is line-delimited JSON (jsonl).

TypeScript schema of jsonl entry is described in @./types.ts and @./drivertypes.ts including field and type documentation.

> [!TIP]
> Use [online jsonl viewer](https://finetunedb.com/tools/jsonl-viewer) to parse and view jsonl format in user-friendly way.

### `scripts`

The @./scripts/ directory stores helper scripts to automate research and data analysis.
