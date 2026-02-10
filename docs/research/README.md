# Research

This folder contains Go package driver wire format and gopls behavior research docs.

Below is description of each document.

## Files

### `drvtrace.jsonl`

File contains "trace" results of how gopls calls its built-in package driver which is an adapter for `go` command.

The file format is line-delimited JSON (jsonl).

TypeScript schema of jsonl entry is described in @./types.ts and @./drivertypes.ts including field and type documentation.

> [!TIP]
> Use [online jsonl viewer](https://finetunedb.com/tools/jsonl-viewer) to parse and view jsonl format in user-friendly way.

### `drivertypes.go`

Contains type definitions for `DriverRequest` and `DriverResponse` structs used in Go package driver protocol.
