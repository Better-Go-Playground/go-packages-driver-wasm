# Trace Format Notes

This file documents the observed JSONL structure in `docs/research/drvtrace.jsonl` and how it maps to gopls tracing code in `../gopls/go/packages/debug.go`.

**Envelope**
- Each line is one JSON object.
- Exactly one of `cmd` or `drv` is present.
- Header fields are top-level, not nested under `cmd` or `drv`.

```json
{
  "cmd": { ... } | "drv": { ... },
  "spanId": 123,
  "parentSpanId": 0,
  "ts": 1770696377207,
  "stack": [ { "func": "...", "at": "...:line" } ]
}
```

**Header Fields**
- `spanId` and `parentSpanId` form a parent-child span graph.
- `ts` is a Unix timestamp in milliseconds.
- `stack` is a stack trace captured at the call site.

**`cmd` Body**
- Shape is `{"verb": "list", "args": [...], "result": {"ok": string | "error": string}}`.
- `result.ok` is a string containing `go list -json` output. Multiple JSON objects are concatenated in that string.

**`drv` Body**
- Shape is `{"cwd": string, "patterns": [...], "req": {...}, "result": {"ok": {...} | "error": string}, "overlay": {...}?}`.
- `req` is a `DriverRequest` (see `docs/research/drivertypes.go`).
- `req.overlay` values are base64-encoded file contents (Go `[]byte` JSON encoding).
- `overlay` is the parsed overlay file used by the fallback go command. Example:

```json
"overlay": {
  "path": "/tmp/gocommand-2336204532/overlay.json",
  "Content": {
    "replace": {
      "/home/x1unix/prj/go-packages-driver-wasm/internal/driver/loader.go": "/tmp/gocommand-2336204532/1-loader.go"
    }
  }
}
```

**Notable Field-Name Observations**
- `spanId` uses a lowercase `d` (not `spanID`).
- `parentSpanId` uses a lowercase `d`.
- Overlay payload uses `Content` (uppercase `C`), not `content`.
- `stack` is top-level even for `cmd` and `drv` events.

**Cross-Reference**
- Trace emission is implemented in `../gopls/go/packages/debug.go` (`traceCmd`, `traceDrv`, `traceHeader`).
