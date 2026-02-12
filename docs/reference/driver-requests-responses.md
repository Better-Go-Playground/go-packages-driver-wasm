# Driver Trace Reference (Requests & Responses)

This document extracts request/response references from `docs/research/traces/*.jsonl`.
Only `drv` objects are used for request/response material. `cmd` objects are summarized as a reference for the underlying `go list` behavior.

Note: `Packages` arrays in responses are truncated to the first 5 entries for readability. See the trace files for full payloads.

## Sources

- `docs/research/traces/1.jsonl` (`drv`: 18, `cmd`: 17)
- `docs/research/traces/2.jsonl` (`drv`: 13, `cmd`: 13)

## `cmd` Reference (go list invocations)

The following unique argument sets were observed in `cmd` events. These provide context for how `go list` was invoked.

### 1.jsonl

- Count: 7
- Args: `-e`, `-json=Name,ImportPath,Error,Dir,GoFiles,IgnoredGoFiles,IgnoredOtherFiles,CFiles,CgoFiles,CXXFiles,MFiles,HFiles,FFiles,SFiles,SwigFiles,SwigCXXFiles,SysoFiles,TestGoFiles,XTestGoFiles,CompiledGoFiles,Export,DepOnly,Imports,ImportMap,TestImports,XTestImports,ForTest,DepsErrors,Module,EmbedFiles`, `-compiled=true`, `-test=true`, `-export=false`, `-deps=true`, `-find=false`, `-pgo=off`, `-tags`, `integration`, `--`, `github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver`
- Count: 6
- Args: `-e`, `-json=Name,ImportPath,Error,Dir,GoFiles,IgnoredGoFiles,IgnoredOtherFiles,CFiles,CgoFiles,CXXFiles,MFiles,HFiles,FFiles,SFiles,SwigFiles,SwigCXXFiles,SysoFiles,TestGoFiles,XTestGoFiles,CompiledGoFiles,Export,DepOnly,Imports,ImportMap,TestImports,XTestImports,ForTest,DepsErrors,Module,EmbedFiles`, `-compiled=true`, `-test=true`, `-export=false`, `-deps=true`, `-find=false`, `-pgo=off`, `-tags`, `integration`, `--`, `/home/username/prj/go-packages-driver-wasm/internal/driver`
- Count: 2
- Args: `-e`, `-json=Name,ImportPath,Error,Dir,GoFiles,IgnoredGoFiles,IgnoredOtherFiles,CFiles,CgoFiles,CXXFiles,MFiles,HFiles,FFiles,SFiles,SwigFiles,SwigCXXFiles,SysoFiles,TestGoFiles,XTestGoFiles,CompiledGoFiles,Export,DepOnly,Imports,ImportMap,TestImports,XTestImports,ForTest,DepsErrors,Module,EmbedFiles`, `-compiled=true`, `-test=true`, `-export=false`, `-deps=true`, `-find=false`, `-pgo=off`, `-tags`, `integration`, `--`, `/home/username/prj/go-packages-driver-wasm/...`, `builtin`
- Count: 2
- Args: `-e`, `-json=Name,ImportPath,Error,Dir,GoFiles,IgnoredGoFiles,IgnoredOtherFiles,CFiles,CgoFiles,CXXFiles,MFiles,HFiles,FFiles,SFiles,SwigFiles,SwigCXXFiles,SysoFiles,TestGoFiles,XTestGoFiles,CompiledGoFiles,Export,DepOnly,Imports,ImportMap,TestImports,XTestImports,ForTest,DepsErrors,Module,EmbedFiles`, `-compiled=true`, `-test=true`, `-export=false`, `-deps=true`, `-find=false`, `-pgo=off`, `-tags`, `integration`, `--`, `github.com/Better-Go-Playground/go-packages-driver-wasm/cmd/gopackagesdriver`, `github.com/Better-Go-Playground/go-packages-driver-wasm/internal/server`

### 2.jsonl

- Count: 8
- Args: `-e`, `-json=Name,ImportPath,Error,Dir,GoFiles,IgnoredGoFiles,IgnoredOtherFiles,CFiles,CgoFiles,CXXFiles,MFiles,HFiles,FFiles,SFiles,SwigFiles,SwigCXXFiles,SysoFiles,TestGoFiles,XTestGoFiles,CompiledGoFiles,Export,DepOnly,Imports,ImportMap,TestImports,XTestImports,ForTest,DepsErrors,Module,EmbedFiles`, `-compiled=true`, `-test=true`, `-export=false`, `-deps=true`, `-find=false`, `-pgo=off`, `-tags`, `integration`, `--`, `github.com/grafana/alloy`, `github.com/grafana/alloy/internal/alloycli`, `github.com/grafana/alloy/internal/cmd/listcomponents`, `github.com/grafana/alloy/internal/component/all`, `github.com/grafana/alloy/internal/component/metadata`, `github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus`, `github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test`, `github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus`, `github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test`, `github.com/grafana/alloy/internal/component/prometheus`, `github.com/grafana/alloy/internal/component/prometheus/enrich`, `github.com/grafana/alloy/internal/component/prometheus/exporter/tests`, `github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test`, `github.com/grafana/alloy/internal/component/prometheus/operator`, `github.com/grafana/alloy/internal/component/prometheus/operator/common`, `github.com/grafana/alloy/internal/component/prometheus/operator/configgen`, `github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors`, `github.com/grafana/alloy/internal/component/prometheus/operator/probes`, `github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs`, `github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors`, `github.com/grafana/alloy/internal/component/prometheus/receive_http`, `github.com/grafana/alloy/internal/component/prometheus/relabel`, `github.com/grafana/alloy/internal/component/prometheus/remotewrite`, `github.com/grafana/alloy/internal/component/prometheus/remotewrite_test`, `github.com/grafana/alloy/internal/component/prometheus/scrape`, `github.com/grafana/alloy/internal/component/prometheus_test`, `github.com/grafana/alloy/internal/component/pyroscope/scrape`, `github.com/grafana/alloy/internal/converter`, `github.com/grafana/alloy/internal/converter/internal/otelcolconvert`, `github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test`, `github.com/grafana/alloy/internal/converter/internal/prometheusconvert`, `github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component`, `github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test`, `github.com/grafana/alloy/internal/converter/internal/promtailconvert`, `github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build`, `github.com/grafana/alloy/internal/converter/internal/promtailconvert_test`, `github.com/grafana/alloy/internal/converter/internal/staticconvert`, `github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build`, `github.com/grafana/alloy/internal/converter/internal/staticconvert_test`, `github.com/grafana/alloy/internal/service/cluster`, `github.com/grafana/alloy/internal/service/cluster_test`, `github.com/grafana/alloy/internal/tools/docs_generator`, `github.com/grafana/alloy/internal/tools/docs_generator_test`, `github.com/grafana/alloy/internal/validator`
- Count: 2
- Args: `-e`, `-json=Name,ImportPath,Error,Dir,GoFiles,IgnoredGoFiles,IgnoredOtherFiles,CFiles,CgoFiles,CXXFiles,MFiles,HFiles,FFiles,SFiles,SwigFiles,SwigCXXFiles,SysoFiles,TestGoFiles,XTestGoFiles,CompiledGoFiles,Export,DepOnly,Imports,ImportMap,TestImports,XTestImports,ForTest,DepsErrors,Module,EmbedFiles`, `-compiled=true`, `-test=true`, `-export=false`, `-deps=true`, `-find=false`, `-pgo=off`, `-tags`, `integration`, `--`, `/home/username/work/grafana/alloy/internal/component/prometheus/appenders`
- Count: 2
- Args: `-e`, `-json=Name,ImportPath,Error,Dir,GoFiles,IgnoredGoFiles,IgnoredOtherFiles,CFiles,CgoFiles,CXXFiles,MFiles,HFiles,FFiles,SFiles,SwigFiles,SwigCXXFiles,SysoFiles,TestGoFiles,XTestGoFiles,CompiledGoFiles,Export,DepOnly,Imports,ImportMap,TestImports,XTestImports,ForTest,DepsErrors,Module,EmbedFiles`, `-compiled=true`, `-test=true`, `-export=false`, `-deps=true`, `-find=false`, `-pgo=off`, `-tags`, `integration`, `--`, `github.com/grafana/alloy`, `github.com/grafana/alloy/internal/alloycli`, `github.com/grafana/alloy/internal/cmd/listcomponents`, `github.com/grafana/alloy/internal/component/all`, `github.com/grafana/alloy/internal/component/metadata`, `github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus`, `github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test`, `github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus`, `github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test`, `github.com/grafana/alloy/internal/component/prometheus`, `github.com/grafana/alloy/internal/component/prometheus/appenders`, `github.com/grafana/alloy/internal/component/prometheus/enrich`, `github.com/grafana/alloy/internal/component/prometheus/exporter/tests`, `github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test`, `github.com/grafana/alloy/internal/component/prometheus/operator`, `github.com/grafana/alloy/internal/component/prometheus/operator/common`, `github.com/grafana/alloy/internal/component/prometheus/operator/configgen`, `github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors`, `github.com/grafana/alloy/internal/component/prometheus/operator/probes`, `github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs`, `github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors`, `github.com/grafana/alloy/internal/component/prometheus/receive_http`, `github.com/grafana/alloy/internal/component/prometheus/relabel`, `github.com/grafana/alloy/internal/component/prometheus/remotewrite`, `github.com/grafana/alloy/internal/component/prometheus/remotewrite_test`, `github.com/grafana/alloy/internal/component/prometheus/scrape`, `github.com/grafana/alloy/internal/component/prometheus_test`, `github.com/grafana/alloy/internal/component/pyroscope/scrape`, `github.com/grafana/alloy/internal/converter`, `github.com/grafana/alloy/internal/converter/internal/otelcolconvert`, `github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test`, `github.com/grafana/alloy/internal/converter/internal/prometheusconvert`, `github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component`, `github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test`, `github.com/grafana/alloy/internal/converter/internal/promtailconvert`, `github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build`, `github.com/grafana/alloy/internal/converter/internal/promtailconvert_test`, `github.com/grafana/alloy/internal/converter/internal/staticconvert`, `github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build`, `github.com/grafana/alloy/internal/converter/internal/staticconvert_test`, `github.com/grafana/alloy/internal/service/cluster`, `github.com/grafana/alloy/internal/service/cluster_test`, `github.com/grafana/alloy/internal/tools/docs_generator`, `github.com/grafana/alloy/internal/tools/docs_generator_test`, `github.com/grafana/alloy/internal/validator`
- Count: 1
- Args: `-e`, `-json=Name,ImportPath,Error,Dir,GoFiles,IgnoredGoFiles,IgnoredOtherFiles,CFiles,CgoFiles,CXXFiles,MFiles,HFiles,FFiles,SFiles,SwigFiles,SwigCXXFiles,SysoFiles,TestGoFiles,XTestGoFiles,CompiledGoFiles,Export,DepOnly,Imports,ImportMap,TestImports,XTestImports,ForTest,DepsErrors,Module,EmbedFiles`, `-compiled=true`, `-test=true`, `-export=false`, `-deps=true`, `-find=false`, `-pgo=off`, `-tags`, `integration`, `--`, `/home/username/work/grafana/alloy/...`, `builtin`

## `drv` Requests & Responses

Requests are shown as `DriverRequestEnvelope` values: `{ workDir, patterns, driverRequest }`.
Responses are shown as `packages.DriverResponse` (see `internal/server/server.go`).

### 1.jsonl

Total drv events: 18

#### drv #1

Trace meta: spanId=1, ts=1770696368800, ts_iso=2026-02-10T04:06:08.800000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": ["/home/username/prj/go-packages-driver-wasm/...", "builtin"],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "GO111MODULE=auto"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {}
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/fakenet",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/jsonrpc",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/server",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/cmd/gopackagesdriver",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver",
    "builtin"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "builtin",
      "Name": "builtin",
      "PkgPath": "builtin",
      "GoFiles": ["/usr/lib/go/src/builtin/builtin.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/builtin/builtin.go"],
      "Imports": {
        "cmp": "cmp"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    }
  ],
  "GoVersion": 25
}
```

Notes:

- `Packages` truncated from 157 to 5 entries.
- `builtin` package has always to be present in the output.

#### drv #2

Trace meta: spanId=4, ts=1770696377207, ts_iso=2026-02-10T04:06:17.207000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "file=/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJnb2xhbmcub3JnL3gvdG9vbHMvZ28vcGFja2FnZXMiCikKCi8vIENvbmZpZyBpcyBwYWNrYWdlIGxvYWQgY29uZmlndXJhdGlvbi4KLy8KLy8gSW5zcGlyZWQgYnkgW2dvbGFuZy5vcmcveC90b29scy9wYWNrYWdlcy5Db25maWddIHN0cnVjdC4KdHlwZSBDb25maWcgc3RydWN0IHsKCQp9CgpmdW5jIExvYWQoY2ZnICpDb25maWcsIHBhdHRlcm5zIFtdc3RyaW5nKSAoKnBhY2thZ2VzLkRyaXZlclJlc3BvbnNlLCBlcnJvcikgewoJcmV0dXJuIG5pbCwgbmlsCn0K"
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 0
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #3

Trace meta: spanId=3, ts=1770696377206, ts_iso=2026-02-10T04:06:17.206000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJnb2xhbmcub3JnL3gvdG9vbHMvZ28vcGFja2FnZXMiCikKCi8vIENvbmZpZyBpcyBwYWNrYWdlIGxvYWQgY29uZmlndXJhdGlvbi4KLy8KLy8gSW5zcGlyZWQgYnkgW2dvbGFuZy5vcmcveC90b29scy9wYWNrYWdlcy5Db25maWddIHN0cnVjdC4KdHlwZSBDb25maWcgc3RydWN0IHsKCQp9CgpmdW5jIExvYWQoY2ZnICpDb25maWcsIHBhdHRlcm5zIFtdc3RyaW5nKSAoKnBhY2thZ2VzLkRyaXZlclJlc3BvbnNlLCBlcnJvcikgewoJcmV0dXJuIG5pbCwgbmlsCn0K"
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #4

Trace meta: spanId=7, ts=1770696378369, ts_iso=2026-02-10T04:06:18.369000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiZ29sYW5nLm9yZy94L3Rvb2xzL2dvL3BhY2thZ2VzIgopCgovLyBDb25maWcgaXMgcGFja2FnZSBsb2FkIGNvbmZpZ3VyYXRpb24uCi8vCi8vIEluc3BpcmVkIGJ5IFtnb2xhbmcub3JnL3gvdG9vbHMvcGFja2FnZXMuQ29uZmlnXSBzdHJ1Y3QuCnR5cGUgQ29uZmlnIHN0cnVjdHt9CgpmdW5jIExvYWQoY2ZnICpDb25maWcsIHBhdHRlcm5zIFtdc3RyaW5nKSAoKnBhY2thZ2VzLkRyaXZlclJlc3BvbnNlLCBlcnJvcikgewoJcmV0dXJuIG5pbCwgbmlsCn0K"
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #5

Trace meta: spanId=8, ts=1770696378369, ts_iso=2026-02-10T04:06:18.369000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "file=/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiZ29sYW5nLm9yZy94L3Rvb2xzL2dvL3BhY2thZ2VzIgopCgovLyBDb25maWcgaXMgcGFja2FnZSBsb2FkIGNvbmZpZ3VyYXRpb24uCi8vCi8vIEluc3BpcmVkIGJ5IFtnb2xhbmcub3JnL3gvdG9vbHMvcGFja2FnZXMuQ29uZmlnXSBzdHJ1Y3QuCnR5cGUgQ29uZmlnIHN0cnVjdHt9CgpmdW5jIExvYWQoY2ZnICpDb25maWcsIHBhdHRlcm5zIFtdc3RyaW5nKSAoKnBhY2thZ2VzLkRyaXZlclJlc3BvbnNlLCBlcnJvcikgewoJcmV0dXJuIG5pbCwgbmlsCn0K"
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 0
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #6

Trace meta: spanId=13, ts=1770696379337, ts_iso=2026-02-10T04:06:19.337000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "file=/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCgkiZ29sYW5nLm9yZy94L3Rvb2xzL2dvL3BhY2thZ2VzIgopCgovLyBDb25maWcgaXMgcGFja2FnZSBsb2FkIGNvbmZpZ3VyYXRpb24uCi8vCi8vIEluc3BpcmVkIGJ5IFtnb2xhbmcub3JnL3gvdG9vbHMvcGFja2FnZXMuQ29uZmlnXSBzdHJ1Y3QuCnR5cGUgQ29uZmlnIHN0cnVjdHt9CgpmdW5jIExvYWQoY2ZnICpDb25maWcsIHBhdHRlcm5zIFtdc3RyaW5nKSAoKnBhY2thZ2VzLkRyaXZlclJlc3BvbnNlLCBlcnJvcikgewoJcmV0dXJuIG5pbCwgbmlsCn0K"
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 0
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #7

Trace meta: spanId=11, ts=1770696379333, ts_iso=2026-02-10T04:06:19.333000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCgkiZ29sYW5nLm9yZy94L3Rvb2xzL2dvL3BhY2thZ2VzIgopCgovLyBDb25maWcgaXMgcGFja2FnZSBsb2FkIGNvbmZpZ3VyYXRpb24uCi8vCi8vIEluc3BpcmVkIGJ5IFtnb2xhbmcub3JnL3gvdG9vbHMvcGFja2FnZXMuQ29uZmlnXSBzdHJ1Y3QuCnR5cGUgQ29uZmlnIHN0cnVjdHt9CgpmdW5jIExvYWQoY2ZnICpDb25maWcsIHBhdHRlcm5zIFtdc3RyaW5nKSAoKnBhY2thZ2VzLkRyaXZlclJlc3BvbnNlLCBlcnJvcikgewoJcmV0dXJuIG5pbCwgbmlsCn0K"
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #8

Trace meta: spanId=15, ts=1770696380688, ts_iso=2026-02-10T04:06:20.688000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJnby9hc3QiCgkiZ28vdG9rZW4iCgoJImdvbGFuZy5vcmcveC90b29scy9nby9wYWNrYWdlcyIKKQoKLy8gQ29uZmlnIGlzIHBhY2thZ2UgbG9hZCBjb25maWd1cmF0aW9uLgovLwovLyBJbnNwaXJlZCBieSBbZ29sYW5nLm9yZy94L3Rvb2xzL3BhY2thZ2VzLkNvbmZpZ10gc3RydWN0Lgp0eXBlIENvbmZpZyBzdHJ1Y3QgewoJCn0KCmZ1bmMgTG9hZChjZmcgKkNvbmZpZywgcGF0dGVybnMgW11zdHJpbmcpICgqcGFja2FnZXMuRHJpdmVyUmVzcG9uc2UsIGVycm9yKSB7CglyZXR1cm4gbmlsLCBuaWwKfQo="
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #9

Trace meta: spanId=16, ts=1770696380689, ts_iso=2026-02-10T04:06:20.689000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "file=/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJnby9hc3QiCgkiZ28vdG9rZW4iCgoJImdvbGFuZy5vcmcveC90b29scy9nby9wYWNrYWdlcyIKKQoKLy8gQ29uZmlnIGlzIHBhY2thZ2UgbG9hZCBjb25maWd1cmF0aW9uLgovLwovLyBJbnNwaXJlZCBieSBbZ29sYW5nLm9yZy94L3Rvb2xzL3BhY2thZ2VzLkNvbmZpZ10gc3RydWN0Lgp0eXBlIENvbmZpZyBzdHJ1Y3QgewoJCn0KCmZ1bmMgTG9hZChjZmcgKkNvbmZpZywgcGF0dGVybnMgW11zdHJpbmcpICgqcGFja2FnZXMuRHJpdmVyUmVzcG9uc2UsIGVycm9yKSB7CglyZXR1cm4gbmlsLCBuaWwKfQo="
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 0
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #10

Trace meta: spanId=20, ts=1770696387724, ts_iso=2026-02-10T04:06:27.724000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "file=/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJmbXQiCgkiZ28vYXN0IgoJImdvL3Rva2VuIgoKCSJnb2xhbmcub3JnL3gvdG9vbHMvZ28vcGFja2FnZXMiCikKCi8vIENvbmZpZyBpcyBwYWNrYWdlIGxvYWQgY29uZmlndXJhdGlvbi4KLy8KLy8gSW5zcGlyZWQgYnkgW2dvbGFuZy5vcmcveC90b29scy9wYWNrYWdlcy5Db25maWddIHN0cnVjdC4KdHlwZSBDb25maWcgc3RydWN0IHsKCS8vIE1vZGUgY29udHJvbHMgdGhlIGxldmVsIG9mIGluZm9ybWF0aW9uIHJldHVybmVkIGZvciBlYWNoIHBhY2thZ2UuCglNb2RlIHBhY2thZ2VzLkxvYWRNb2RlCgoJLy8gQ29udGV4dCBzcGVjaWZpZXMgdGhlIGNvbnRleHQgZm9yIHRoZSBsb2FkIG9wZXJhdGlvbi4KCS8vIENhbmNlbGxpbmcgdGhlIGNvbnRleHQgbWF5IGNhdXNlIFtMb2FkXSB0byBhYm9ydCBhbmQKCS8vIHJldHVybiBhbiBlcnJvci4KCUNvbnRleHQgY29udGV4dC5Db250ZXh0CgoJLy8gRGlyIGlzIHRoZSBkaXJlY3RvcnkgaW4gd2hpY2ggdG8gcnVuIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sCgkvLyB0aGF0IHByb3ZpZGVzIGluZm9ybWF0aW9uIGFib3V0IHRoZSBwYWNrYWdlcy4KCS8vIElmIERpciBpcyBlbXB0eSwgdGhlIHRvb2wgaXMgcnVuIGluIHRoZSBjdXJyZW50IGRpcmVjdG9yeS4KCURpciBzdHJpbmcKCgkvLyBFbnYgaXMgdGhlIGVudmlyb25tZW50IHRvIHVzZSB3aGVuIGludm9raW5nIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sLgoJRW52IG1hcFtzdHJpbmddc3RyaW5nCgoJLy8gQnVpbGRGbGFncyBpcyBhIGxpc3Qgb2YgY29tbWFuZC1saW5lIGZsYWdzIHRvIGJlIHBhc3NlZCB0aHJvdWdoIHRvCgkvLyB0aGUgYnVpbGQgc3lzdGVtJ3MgcXVlcnkgdG9vbC4KCUJ1aWxkRmxhZ3MgW11zdHJpbmcKCgkvLyBGc2V0IHByb3ZpZGVzIHNvdXJjZSBwb3NpdGlvbiBpbmZvcm1hdGlvbiBmb3Igc3ludGF4IHRyZWVzIGFuZCB0eXBlcy4KCS8vIElmIEZzZXQgaXMgbmlsLCBMb2FkIHdpbGwgdXNlIGEgbmV3IGZpbGVzZXQsIGJ1dCBwcmVzZXJ2ZSBGc2V0J3MgdmFsdWUuCglGc2V0ICp0b2tlbi5GaWxlU2V0CgoJLy8gUGFyc2VGaWxlIGlzIGNhbGxlZCB0byByZWFkIGFuZCBwYXJzZSBlYWNoIGZpbGUKCS8vIHdoZW4gcHJlcGFyaW5nIGEgcGFja2FnZSdzIHR5cGUtY2hlY2tlZCBzeW50YXggdHJlZS4KCS8vIEl0IG11c3QgYmUgc2FmZSB0byBjYWxsIFBhcnNlRmlsZSBzaW11bHRhbmVvdXNseSBmcm9tIG11bHRpcGxlIGdvcm91dGluZXMuCgkvLyBJZiBQYXJzZUZpbGUgaXMgbmlsLCB0aGUgbG9hZGVyIHdpbGwgdXNlcyBwYXJzZXIuUGFyc2VGaWxlLgoJLy8KCS8vIFBhcnNlRmlsZSBzaG91bGQgcGFyc2UgdGhlIHNvdXJjZSBmcm9tIHNyYyBhbmQgdXNlIGZpbGVuYW1lIG9ubHkgZm9yCgkvLyByZWNvcmRpbmcgcG9zaXRpb24gaW5mb3JtYXRpb24uCgkvLwoJLy8gQW4gYXBwbGljYXRpb24gbWF5IHN1cHBseSBhIGN1c3RvbSBpbXBsZW1lbnRhdGlvbiBvZiBQYXJzZUZpbGUKCS8vIHRvIGNoYW5nZSB0aGUgZWZmZWN0aXZlIGZpbGUgY29udGVudHMgb3IgdGhlIGJlaGF2aW9yIG9mIHRoZSBwYXJzZXIsCgkvLyBvciB0byBtb2RpZnkgdGhlIHN5bnRheCB0cmVlLiBGb3IgZXhhbXBsZSwgc2VsZWN0aXZlbHkgZWxpbWluYXRpbmcKCS8vIHVud2FudGVkIGZ1bmN0aW9uIGJvZGllcyBjYW4gc2lnbmlmaWNhbnRseSBhY2NlbGVyYXRlIHR5cGUgY2hlY2tpbmcuCglQYXJzZUZpbGUgZnVuYyhmc2V0ICp0b2tlbi5GaWxlU2V0LCBmaWxlbmFtZSBzdHJpbmcsIHNyYyBbXWJ5dGUpICgqYXN0LkZpbGUsIGVycm9yKQoKCS8vIElmIFRlc3RzIGlzIHNldCwgdGhlIGxvYWRlciBpbmNsdWRlcyBub3QganVzdCB0aGUgcGFja2FnZXMKCS8vIG1hdGNoaW5nIGEgcGFydGljdWxhciBwYXR0ZXJuIGJ1dCBhbHNvIGFueSByZWxhdGVkIHRlc3QgcGFja2FnZXMsCgkvLyBpbmNsdWRpbmcgdGVzdC1vbmx5IHZhcmlhbnRzIG9mIHRoZSBwYWNrYWdlIGFuZCB0aGUgdGVzdCBleGVjdXRhYmxlLgoJLy8KCS8vIEZvciBleGFtcGxlLCB3aGVuIHVzaW5nIHRoZSBnbyBjb21tYW5kLCBsb2FkaW5nICJmbXQiIHdpdGggVGVzdHM9dHJ1ZQoJLy8gcmV0dXJucyBmb3VyIHBhY2thZ2VzLCB3aXRoIElEcyAiZm10IiAodGhlIHN0YW5kYXJkIHBhY2thZ2UpLAoJLy8gImZtdCBbZm10LnRlc3RdIiAodGhlIHBhY2thZ2UgYXMgY29tcGlsZWQgZm9yIHRoZSB0ZXN0KSwKCS8vICJmbXRfdGVzdCIgKHRoZSB0ZXN0IGZ1bmN0aW9ucyBmcm9tIHNvdXJjZSBmaWxlcyBpbiBwYWNrYWdlIGZtdF90ZXN0KSwKCS8vIGFuZCAiZm10LnRlc3QiICh0aGUgdGVzdCBiaW5hcnkpLgoJLy8KCS8vIEluIGJ1aWxkIHN5c3RlbXMgd2l0aCBleHBsaWNpdCBuYW1lcyBmb3IgdGVzdHMsCgkvLyBzZXR0aW5nIFRlc3RzIG1heSBoYXZlIG5vIGVmZmVjdC4KCVRlc3RzIGJvb2wKCgkvLyBPdmVybGF5IGlzIGEgbWFwcGluZyBmcm9tIGFic29sdXRlIGZpbGUgcGF0aHMgdG8gZmlsZSBjb250ZW50cy4KCS8vCgkvLyBGb3IgZWFjaCBtYXAgZW50cnksIFtMb2FkXSB1c2VzIHRoZSBhbHRlcm5hdGl2ZSBmaWxlCgkvLyBjb250ZW50cyBwcm92aWRlZCBieSB0aGUgb3ZlcmxheSBtYXBwaW5nIGluc3RlYWQgb2YgcmVhZGluZwoJLy8gZnJvbSB0aGUgZmlsZSBzeXN0ZW0uIFRoaXMgbWVjaGFuaXNtIGNhbiBiZSB1c2VkIHRvIGVuYWJsZQoJLy8gZWRpdG9yLWludGVncmF0ZWQgdG9vbHMgdG8gY29ycmVjdGx5IGFuYWx5emUgdGhlIGNvbnRlbnRzCgkvLyBvZiBtb2RpZmllZCBidXQgdW5zYXZlZCBidWZmZXJzLCBmb3IgZXhhbXBsZS4KCS8vCgkvLyBUaGUgb3ZlcmxheSBtYXBwaW5nIGlzIHBhc3NlZCB0byB0aGUgYnVpbGQgc3lzdGVtJ3MgZHJpdmVyCgkvLyAoc2VlICJUaGUgZHJpdmVyIHByb3RvY29sIikgc28gdGhhdCBpdCB0b28gY2FuIHJlcG9ydAoJLy8gY29uc2lzdGVudCBwYWNrYWdlIG1ldGFkYXRhIGFib3V0IHVuc2F2ZWQgZmlsZXMuIEhvd2V2ZXIsCgkvLyBkcml2ZXJzIG1heSB2YXJ5IGluIHRoZWlyIGxldmVsIG9mIHN1cHBvcnQgZm9yIG92ZXJsYXlzLgoJT3ZlcmxheSBtYXBbc3RyaW5nXVtdYnl0ZQp9CgpmdW5jIExvYWQoY2ZnICpDb25maWcsIHBhdHRlcm5zIFtdc3RyaW5nKSAoKnBhY2thZ2VzLkRyaXZlclJlc3BvbnNlLCBlcnJvcikgewoJZm10LlByaW50bG4oKQoJcmV0dXJuIG5pbCwgbmlsCn0K"
    }
  }
}
```

Response (packages.DriverResponse):
Error: `err: context canceled: stderr: `

#### drv #11

Trace meta: spanId=19, ts=1770696387724, ts_iso=2026-02-10T04:06:27.724000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJmbXQiCgkiZ28vYXN0IgoJImdvL3Rva2VuIgoKCSJnb2xhbmcub3JnL3gvdG9vbHMvZ28vcGFja2FnZXMiCikKCi8vIENvbmZpZyBpcyBwYWNrYWdlIGxvYWQgY29uZmlndXJhdGlvbi4KLy8KLy8gSW5zcGlyZWQgYnkgW2dvbGFuZy5vcmcveC90b29scy9wYWNrYWdlcy5Db25maWddIHN0cnVjdC4KdHlwZSBDb25maWcgc3RydWN0IHsKCS8vIE1vZGUgY29udHJvbHMgdGhlIGxldmVsIG9mIGluZm9ybWF0aW9uIHJldHVybmVkIGZvciBlYWNoIHBhY2thZ2UuCglNb2RlIHBhY2thZ2VzLkxvYWRNb2RlCgoJLy8gQ29udGV4dCBzcGVjaWZpZXMgdGhlIGNvbnRleHQgZm9yIHRoZSBsb2FkIG9wZXJhdGlvbi4KCS8vIENhbmNlbGxpbmcgdGhlIGNvbnRleHQgbWF5IGNhdXNlIFtMb2FkXSB0byBhYm9ydCBhbmQKCS8vIHJldHVybiBhbiBlcnJvci4KCUNvbnRleHQgY29udGV4dC5Db250ZXh0CgoJLy8gRGlyIGlzIHRoZSBkaXJlY3RvcnkgaW4gd2hpY2ggdG8gcnVuIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sCgkvLyB0aGF0IHByb3ZpZGVzIGluZm9ybWF0aW9uIGFib3V0IHRoZSBwYWNrYWdlcy4KCS8vIElmIERpciBpcyBlbXB0eSwgdGhlIHRvb2wgaXMgcnVuIGluIHRoZSBjdXJyZW50IGRpcmVjdG9yeS4KCURpciBzdHJpbmcKCgkvLyBFbnYgaXMgdGhlIGVudmlyb25tZW50IHRvIHVzZSB3aGVuIGludm9raW5nIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sLgoJRW52IG1hcFtzdHJpbmddc3RyaW5nCgoJLy8gQnVpbGRGbGFncyBpcyBhIGxpc3Qgb2YgY29tbWFuZC1saW5lIGZsYWdzIHRvIGJlIHBhc3NlZCB0aHJvdWdoIHRvCgkvLyB0aGUgYnVpbGQgc3lzdGVtJ3MgcXVlcnkgdG9vbC4KCUJ1aWxkRmxhZ3MgW11zdHJpbmcKCgkvLyBGc2V0IHByb3ZpZGVzIHNvdXJjZSBwb3NpdGlvbiBpbmZvcm1hdGlvbiBmb3Igc3ludGF4IHRyZWVzIGFuZCB0eXBlcy4KCS8vIElmIEZzZXQgaXMgbmlsLCBMb2FkIHdpbGwgdXNlIGEgbmV3IGZpbGVzZXQsIGJ1dCBwcmVzZXJ2ZSBGc2V0J3MgdmFsdWUuCglGc2V0ICp0b2tlbi5GaWxlU2V0CgoJLy8gUGFyc2VGaWxlIGlzIGNhbGxlZCB0byByZWFkIGFuZCBwYXJzZSBlYWNoIGZpbGUKCS8vIHdoZW4gcHJlcGFyaW5nIGEgcGFja2FnZSdzIHR5cGUtY2hlY2tlZCBzeW50YXggdHJlZS4KCS8vIEl0IG11c3QgYmUgc2FmZSB0byBjYWxsIFBhcnNlRmlsZSBzaW11bHRhbmVvdXNseSBmcm9tIG11bHRpcGxlIGdvcm91dGluZXMuCgkvLyBJZiBQYXJzZUZpbGUgaXMgbmlsLCB0aGUgbG9hZGVyIHdpbGwgdXNlcyBwYXJzZXIuUGFyc2VGaWxlLgoJLy8KCS8vIFBhcnNlRmlsZSBzaG91bGQgcGFyc2UgdGhlIHNvdXJjZSBmcm9tIHNyYyBhbmQgdXNlIGZpbGVuYW1lIG9ubHkgZm9yCgkvLyByZWNvcmRpbmcgcG9zaXRpb24gaW5mb3JtYXRpb24uCgkvLwoJLy8gQW4gYXBwbGljYXRpb24gbWF5IHN1cHBseSBhIGN1c3RvbSBpbXBsZW1lbnRhdGlvbiBvZiBQYXJzZUZpbGUKCS8vIHRvIGNoYW5nZSB0aGUgZWZmZWN0aXZlIGZpbGUgY29udGVudHMgb3IgdGhlIGJlaGF2aW9yIG9mIHRoZSBwYXJzZXIsCgkvLyBvciB0byBtb2RpZnkgdGhlIHN5bnRheCB0cmVlLiBGb3IgZXhhbXBsZSwgc2VsZWN0aXZlbHkgZWxpbWluYXRpbmcKCS8vIHVud2FudGVkIGZ1bmN0aW9uIGJvZGllcyBjYW4gc2lnbmlmaWNhbnRseSBhY2NlbGVyYXRlIHR5cGUgY2hlY2tpbmcuCglQYXJzZUZpbGUgZnVuYyhmc2V0ICp0b2tlbi5GaWxlU2V0LCBmaWxlbmFtZSBzdHJpbmcsIHNyYyBbXWJ5dGUpICgqYXN0LkZpbGUsIGVycm9yKQoKCS8vIElmIFRlc3RzIGlzIHNldCwgdGhlIGxvYWRlciBpbmNsdWRlcyBub3QganVzdCB0aGUgcGFja2FnZXMKCS8vIG1hdGNoaW5nIGEgcGFydGljdWxhciBwYXR0ZXJuIGJ1dCBhbHNvIGFueSByZWxhdGVkIHRlc3QgcGFja2FnZXMsCgkvLyBpbmNsdWRpbmcgdGVzdC1vbmx5IHZhcmlhbnRzIG9mIHRoZSBwYWNrYWdlIGFuZCB0aGUgdGVzdCBleGVjdXRhYmxlLgoJLy8KCS8vIEZvciBleGFtcGxlLCB3aGVuIHVzaW5nIHRoZSBnbyBjb21tYW5kLCBsb2FkaW5nICJmbXQiIHdpdGggVGVzdHM9dHJ1ZQoJLy8gcmV0dXJucyBmb3VyIHBhY2thZ2VzLCB3aXRoIElEcyAiZm10IiAodGhlIHN0YW5kYXJkIHBhY2thZ2UpLAoJLy8gImZtdCBbZm10LnRlc3RdIiAodGhlIHBhY2thZ2UgYXMgY29tcGlsZWQgZm9yIHRoZSB0ZXN0KSwKCS8vICJmbXRfdGVzdCIgKHRoZSB0ZXN0IGZ1bmN0aW9ucyBmcm9tIHNvdXJjZSBmaWxlcyBpbiBwYWNrYWdlIGZtdF90ZXN0KSwKCS8vIGFuZCAiZm10LnRlc3QiICh0aGUgdGVzdCBiaW5hcnkpLgoJLy8KCS8vIEluIGJ1aWxkIHN5c3RlbXMgd2l0aCBleHBsaWNpdCBuYW1lcyBmb3IgdGVzdHMsCgkvLyBzZXR0aW5nIFRlc3RzIG1heSBoYXZlIG5vIGVmZmVjdC4KCVRlc3RzIGJvb2wKCgkvLyBPdmVybGF5IGlzIGEgbWFwcGluZyBmcm9tIGFic29sdXRlIGZpbGUgcGF0aHMgdG8gZmlsZSBjb250ZW50cy4KCS8vCgkvLyBGb3IgZWFjaCBtYXAgZW50cnksIFtMb2FkXSB1c2VzIHRoZSBhbHRlcm5hdGl2ZSBmaWxlCgkvLyBjb250ZW50cyBwcm92aWRlZCBieSB0aGUgb3ZlcmxheSBtYXBwaW5nIGluc3RlYWQgb2YgcmVhZGluZwoJLy8gZnJvbSB0aGUgZmlsZSBzeXN0ZW0uIFRoaXMgbWVjaGFuaXNtIGNhbiBiZSB1c2VkIHRvIGVuYWJsZQoJLy8gZWRpdG9yLWludGVncmF0ZWQgdG9vbHMgdG8gY29ycmVjdGx5IGFuYWx5emUgdGhlIGNvbnRlbnRzCgkvLyBvZiBtb2RpZmllZCBidXQgdW5zYXZlZCBidWZmZXJzLCBmb3IgZXhhbXBsZS4KCS8vCgkvLyBUaGUgb3ZlcmxheSBtYXBwaW5nIGlzIHBhc3NlZCB0byB0aGUgYnVpbGQgc3lzdGVtJ3MgZHJpdmVyCgkvLyAoc2VlICJUaGUgZHJpdmVyIHByb3RvY29sIikgc28gdGhhdCBpdCB0b28gY2FuIHJlcG9ydAoJLy8gY29uc2lzdGVudCBwYWNrYWdlIG1ldGFkYXRhIGFib3V0IHVuc2F2ZWQgZmlsZXMuIEhvd2V2ZXIsCgkvLyBkcml2ZXJzIG1heSB2YXJ5IGluIHRoZWlyIGxldmVsIG9mIHN1cHBvcnQgZm9yIG92ZXJsYXlzLgoJT3ZlcmxheSBtYXBbc3RyaW5nXVtdYnl0ZQp9CgpmdW5jIExvYWQoY2ZnICpDb25maWcsIHBhdHRlcm5zIFtdc3RyaW5nKSAoKnBhY2thZ2VzLkRyaXZlclJlc3BvbnNlLCBlcnJvcikgewoJZm10LlByaW50bG4oKQoJcmV0dXJuIG5pbCwgbmlsCn0K"
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #12

Trace meta: spanId=21, ts=1770696387725, ts_iso=2026-02-10T04:06:27.725000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "file=/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJmbXQiCgkiZ28vYXN0IgoJImdvL3Rva2VuIgoKCSJnb2xhbmcub3JnL3gvdG9vbHMvZ28vcGFja2FnZXMiCikKCi8vIENvbmZpZyBpcyBwYWNrYWdlIGxvYWQgY29uZmlndXJhdGlvbi4KLy8KLy8gSW5zcGlyZWQgYnkgW2dvbGFuZy5vcmcveC90b29scy9wYWNrYWdlcy5Db25maWddIHN0cnVjdC4KdHlwZSBDb25maWcgc3RydWN0IHsKCS8vIE1vZGUgY29udHJvbHMgdGhlIGxldmVsIG9mIGluZm9ybWF0aW9uIHJldHVybmVkIGZvciBlYWNoIHBhY2thZ2UuCglNb2RlIHBhY2thZ2VzLkxvYWRNb2RlCgoJLy8gQ29udGV4dCBzcGVjaWZpZXMgdGhlIGNvbnRleHQgZm9yIHRoZSBsb2FkIG9wZXJhdGlvbi4KCS8vIENhbmNlbGxpbmcgdGhlIGNvbnRleHQgbWF5IGNhdXNlIFtMb2FkXSB0byBhYm9ydCBhbmQKCS8vIHJldHVybiBhbiBlcnJvci4KCUNvbnRleHQgY29udGV4dC5Db250ZXh0CgoJLy8gRGlyIGlzIHRoZSBkaXJlY3RvcnkgaW4gd2hpY2ggdG8gcnVuIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sCgkvLyB0aGF0IHByb3ZpZGVzIGluZm9ybWF0aW9uIGFib3V0IHRoZSBwYWNrYWdlcy4KCS8vIElmIERpciBpcyBlbXB0eSwgdGhlIHRvb2wgaXMgcnVuIGluIHRoZSBjdXJyZW50IGRpcmVjdG9yeS4KCURpciBzdHJpbmcKCgkvLyBFbnYgaXMgdGhlIGVudmlyb25tZW50IHRvIHVzZSB3aGVuIGludm9raW5nIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sLgoJRW52IG1hcFtzdHJpbmddc3RyaW5nCgoJLy8gQnVpbGRGbGFncyBpcyBhIGxpc3Qgb2YgY29tbWFuZC1saW5lIGZsYWdzIHRvIGJlIHBhc3NlZCB0aHJvdWdoIHRvCgkvLyB0aGUgYnVpbGQgc3lzdGVtJ3MgcXVlcnkgdG9vbC4KCUJ1aWxkRmxhZ3MgW11zdHJpbmcKCgkvLyBGc2V0IHByb3ZpZGVzIHNvdXJjZSBwb3NpdGlvbiBpbmZvcm1hdGlvbiBmb3Igc3ludGF4IHRyZWVzIGFuZCB0eXBlcy4KCS8vIElmIEZzZXQgaXMgbmlsLCBMb2FkIHdpbGwgdXNlIGEgbmV3IGZpbGVzZXQsIGJ1dCBwcmVzZXJ2ZSBGc2V0J3MgdmFsdWUuCglGc2V0ICp0b2tlbi5GaWxlU2V0CgoJLy8gUGFyc2VGaWxlIGlzIGNhbGxlZCB0byByZWFkIGFuZCBwYXJzZSBlYWNoIGZpbGUKCS8vIHdoZW4gcHJlcGFyaW5nIGEgcGFja2FnZSdzIHR5cGUtY2hlY2tlZCBzeW50YXggdHJlZS4KCS8vIEl0IG11c3QgYmUgc2FmZSB0byBjYWxsIFBhcnNlRmlsZSBzaW11bHRhbmVvdXNseSBmcm9tIG11bHRpcGxlIGdvcm91dGluZXMuCgkvLyBJZiBQYXJzZUZpbGUgaXMgbmlsLCB0aGUgbG9hZGVyIHdpbGwgdXNlcyBwYXJzZXIuUGFyc2VGaWxlLgoJLy8KCS8vIFBhcnNlRmlsZSBzaG91bGQgcGFyc2UgdGhlIHNvdXJjZSBmcm9tIHNyYyBhbmQgdXNlIGZpbGVuYW1lIG9ubHkgZm9yCgkvLyByZWNvcmRpbmcgcG9zaXRpb24gaW5mb3JtYXRpb24uCgkvLwoJLy8gQW4gYXBwbGljYXRpb24gbWF5IHN1cHBseSBhIGN1c3RvbSBpbXBsZW1lbnRhdGlvbiBvZiBQYXJzZUZpbGUKCS8vIHRvIGNoYW5nZSB0aGUgZWZmZWN0aXZlIGZpbGUgY29udGVudHMgb3IgdGhlIGJlaGF2aW9yIG9mIHRoZSBwYXJzZXIsCgkvLyBvciB0byBtb2RpZnkgdGhlIHN5bnRheCB0cmVlLiBGb3IgZXhhbXBsZSwgc2VsZWN0aXZlbHkgZWxpbWluYXRpbmcKCS8vIHVud2FudGVkIGZ1bmN0aW9uIGJvZGllcyBjYW4gc2lnbmlmaWNhbnRseSBhY2NlbGVyYXRlIHR5cGUgY2hlY2tpbmcuCglQYXJzZUZpbGUgZnVuYyhmc2V0ICp0b2tlbi5GaWxlU2V0LCBmaWxlbmFtZSBzdHJpbmcsIHNyYyBbXWJ5dGUpICgqYXN0LkZpbGUsIGVycm9yKQoKCS8vIElmIFRlc3RzIGlzIHNldCwgdGhlIGxvYWRlciBpbmNsdWRlcyBub3QganVzdCB0aGUgcGFja2FnZXMKCS8vIG1hdGNoaW5nIGEgcGFydGljdWxhciBwYXR0ZXJuIGJ1dCBhbHNvIGFueSByZWxhdGVkIHRlc3QgcGFja2FnZXMsCgkvLyBpbmNsdWRpbmcgdGVzdC1vbmx5IHZhcmlhbnRzIG9mIHRoZSBwYWNrYWdlIGFuZCB0aGUgdGVzdCBleGVjdXRhYmxlLgoJLy8KCS8vIEZvciBleGFtcGxlLCB3aGVuIHVzaW5nIHRoZSBnbyBjb21tYW5kLCBsb2FkaW5nICJmbXQiIHdpdGggVGVzdHM9dHJ1ZQoJLy8gcmV0dXJucyBmb3VyIHBhY2thZ2VzLCB3aXRoIElEcyAiZm10IiAodGhlIHN0YW5kYXJkIHBhY2thZ2UpLAoJLy8gImZtdCBbZm10LnRlc3RdIiAodGhlIHBhY2thZ2UgYXMgY29tcGlsZWQgZm9yIHRoZSB0ZXN0KSwKCS8vICJmbXRfdGVzdCIgKHRoZSB0ZXN0IGZ1bmN0aW9ucyBmcm9tIHNvdXJjZSBmaWxlcyBpbiBwYWNrYWdlIGZtdF90ZXN0KSwKCS8vIGFuZCAiZm10LnRlc3QiICh0aGUgdGVzdCBiaW5hcnkpLgoJLy8KCS8vIEluIGJ1aWxkIHN5c3RlbXMgd2l0aCBleHBsaWNpdCBuYW1lcyBmb3IgdGVzdHMsCgkvLyBzZXR0aW5nIFRlc3RzIG1heSBoYXZlIG5vIGVmZmVjdC4KCVRlc3RzIGJvb2wKCgkvLyBPdmVybGF5IGlzIGEgbWFwcGluZyBmcm9tIGFic29sdXRlIGZpbGUgcGF0aHMgdG8gZmlsZSBjb250ZW50cy4KCS8vCgkvLyBGb3IgZWFjaCBtYXAgZW50cnksIFtMb2FkXSB1c2VzIHRoZSBhbHRlcm5hdGl2ZSBmaWxlCgkvLyBjb250ZW50cyBwcm92aWRlZCBieSB0aGUgb3ZlcmxheSBtYXBwaW5nIGluc3RlYWQgb2YgcmVhZGluZwoJLy8gZnJvbSB0aGUgZmlsZSBzeXN0ZW0uIFRoaXMgbWVjaGFuaXNtIGNhbiBiZSB1c2VkIHRvIGVuYWJsZQoJLy8gZWRpdG9yLWludGVncmF0ZWQgdG9vbHMgdG8gY29ycmVjdGx5IGFuYWx5emUgdGhlIGNvbnRlbnRzCgkvLyBvZiBtb2RpZmllZCBidXQgdW5zYXZlZCBidWZmZXJzLCBmb3IgZXhhbXBsZS4KCS8vCgkvLyBUaGUgb3ZlcmxheSBtYXBwaW5nIGlzIHBhc3NlZCB0byB0aGUgYnVpbGQgc3lzdGVtJ3MgZHJpdmVyCgkvLyAoc2VlICJUaGUgZHJpdmVyIHByb3RvY29sIikgc28gdGhhdCBpdCB0b28gY2FuIHJlcG9ydAoJLy8gY29uc2lzdGVudCBwYWNrYWdlIG1ldGFkYXRhIGFib3V0IHVuc2F2ZWQgZmlsZXMuIEhvd2V2ZXIsCgkvLyBkcml2ZXJzIG1heSB2YXJ5IGluIHRoZWlyIGxldmVsIG9mIHN1cHBvcnQgZm9yIG92ZXJsYXlzLgoJT3ZlcmxheSBtYXBbc3RyaW5nXVtdYnl0ZQp9CgpmdW5jIExvYWQoY2ZnICpDb25maWcsIHBhdHRlcm5zIFtdc3RyaW5nKSAoKnBhY2thZ2VzLkRyaXZlclJlc3BvbnNlLCBlcnJvcikgewoJZm10LlByaW50bG4oKQoJcmV0dXJuIG5pbCwgbmlsCn0K"
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 0
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #13

Trace meta: spanId=24, ts=1770696391529, ts_iso=2026-02-10T04:06:31.529000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": ["/home/username/prj/go-packages-driver-wasm/...", "builtin"],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GO111MODULE=auto",
      "GOARCH=wasm",
      "GOOS=js"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJmbXQiCgkiZ28vYXN0IgoJImdvL3Rva2VuIgoKCSJnb2xhbmcub3JnL3gvdG9vbHMvZ28vcGFja2FnZXMiCikKCi8vIENvbmZpZyBpcyBwYWNrYWdlIGxvYWQgY29uZmlndXJhdGlvbi4KLy8KLy8gSW5zcGlyZWQgYnkgW2dvbGFuZy5vcmcveC90b29scy9wYWNrYWdlcy5Db25maWddIHN0cnVjdC4KdHlwZSBDb25maWcgc3RydWN0IHsKCS8vIE1vZGUgY29udHJvbHMgdGhlIGxldmVsIG9mIGluZm9ybWF0aW9uIHJldHVybmVkIGZvciBlYWNoIHBhY2thZ2UuCglNb2RlIHBhY2thZ2VzLkxvYWRNb2RlCgoJLy8gQ29udGV4dCBzcGVjaWZpZXMgdGhlIGNvbnRleHQgZm9yIHRoZSBsb2FkIG9wZXJhdGlvbi4KCS8vIENhbmNlbGxpbmcgdGhlIGNvbnRleHQgbWF5IGNhdXNlIFtMb2FkXSB0byBhYm9ydCBhbmQKCS8vIHJldHVybiBhbiBlcnJvci4KCUNvbnRleHQgY29udGV4dC5Db250ZXh0CgoJLy8gRGlyIGlzIHRoZSBkaXJlY3RvcnkgaW4gd2hpY2ggdG8gcnVuIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sCgkvLyB0aGF0IHByb3ZpZGVzIGluZm9ybWF0aW9uIGFib3V0IHRoZSBwYWNrYWdlcy4KCS8vIElmIERpciBpcyBlbXB0eSwgdGhlIHRvb2wgaXMgcnVuIGluIHRoZSBjdXJyZW50IGRpcmVjdG9yeS4KCURpciBzdHJpbmcKCgkvLyBFbnYgaXMgdGhlIGVudmlyb25tZW50IHRvIHVzZSB3aGVuIGludm9raW5nIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sLgoJRW52IG1hcFtzdHJpbmddc3RyaW5nCgoJLy8gQnVpbGRGbGFncyBpcyBhIGxpc3Qgb2YgY29tbWFuZC1saW5lIGZsYWdzIHRvIGJlIHBhc3NlZCB0aHJvdWdoIHRvCgkvLyB0aGUgYnVpbGQgc3lzdGVtJ3MgcXVlcnkgdG9vbC4KCUJ1aWxkRmxhZ3MgW11zdHJpbmcKCgkvLyBGc2V0IHByb3ZpZGVzIHNvdXJjZSBwb3NpdGlvbiBpbmZvcm1hdGlvbiBmb3Igc3ludGF4IHRyZWVzIGFuZCB0eXBlcy4KCS8vIElmIEZzZXQgaXMgbmlsLCBMb2FkIHdpbGwgdXNlIGEgbmV3IGZpbGVzZXQsIGJ1dCBwcmVzZXJ2ZSBGc2V0J3MgdmFsdWUuCglGc2V0ICp0b2tlbi5GaWxlU2V0CgoJLy8gUGFyc2VGaWxlIGlzIGNhbGxlZCB0byByZWFkIGFuZCBwYXJzZSBlYWNoIGZpbGUKCS8vIHdoZW4gcHJlcGFyaW5nIGEgcGFja2FnZSdzIHR5cGUtY2hlY2tlZCBzeW50YXggdHJlZS4KCS8vIEl0IG11c3QgYmUgc2FmZSB0byBjYWxsIFBhcnNlRmlsZSBzaW11bHRhbmVvdXNseSBmcm9tIG11bHRpcGxlIGdvcm91dGluZXMuCgkvLyBJZiBQYXJzZUZpbGUgaXMgbmlsLCB0aGUgbG9hZGVyIHdpbGwgdXNlcyBwYXJzZXIuUGFyc2VGaWxlLgoJLy8KCS8vIFBhcnNlRmlsZSBzaG91bGQgcGFyc2UgdGhlIHNvdXJjZSBmcm9tIHNyYyBhbmQgdXNlIGZpbGVuYW1lIG9ubHkgZm9yCgkvLyByZWNvcmRpbmcgcG9zaXRpb24gaW5mb3JtYXRpb24uCgkvLwoJLy8gQW4gYXBwbGljYXRpb24gbWF5IHN1cHBseSBhIGN1c3RvbSBpbXBsZW1lbnRhdGlvbiBvZiBQYXJzZUZpbGUKCS8vIHRvIGNoYW5nZSB0aGUgZWZmZWN0aXZlIGZpbGUgY29udGVudHMgb3IgdGhlIGJlaGF2aW9yIG9mIHRoZSBwYXJzZXIsCgkvLyBvciB0byBtb2RpZnkgdGhlIHN5bnRheCB0cmVlLiBGb3IgZXhhbXBsZSwgc2VsZWN0aXZlbHkgZWxpbWluYXRpbmcKCS8vIHVud2FudGVkIGZ1bmN0aW9uIGJvZGllcyBjYW4gc2lnbmlmaWNhbnRseSBhY2NlbGVyYXRlIHR5cGUgY2hlY2tpbmcuCglQYXJzZUZpbGUgZnVuYyhmc2V0ICp0b2tlbi5GaWxlU2V0LCBmaWxlbmFtZSBzdHJpbmcsIHNyYyBbXWJ5dGUpICgqYXN0LkZpbGUsIGVycm9yKQoKCS8vIElmIFRlc3RzIGlzIHNldCwgdGhlIGxvYWRlciBpbmNsdWRlcyBub3QganVzdCB0aGUgcGFja2FnZXMKCS8vIG1hdGNoaW5nIGEgcGFydGljdWxhciBwYXR0ZXJuIGJ1dCBhbHNvIGFueSByZWxhdGVkIHRlc3QgcGFja2FnZXMsCgkvLyBpbmNsdWRpbmcgdGVzdC1vbmx5IHZhcmlhbnRzIG9mIHRoZSBwYWNrYWdlIGFuZCB0aGUgdGVzdCBleGVjdXRhYmxlLgoJLy8KCS8vIEZvciBleGFtcGxlLCB3aGVuIHVzaW5nIHRoZSBnbyBjb21tYW5kLCBsb2FkaW5nICJmbXQiIHdpdGggVGVzdHM9dHJ1ZQoJLy8gcmV0dXJucyBmb3VyIHBhY2thZ2VzLCB3aXRoIElEcyAiZm10IiAodGhlIHN0YW5kYXJkIHBhY2thZ2UpLAoJLy8gImZtdCBbZm10LnRlc3RdIiAodGhlIHBhY2thZ2UgYXMgY29tcGlsZWQgZm9yIHRoZSB0ZXN0KSwKCS8vICJmbXRfdGVzdCIgKHRoZSB0ZXN0IGZ1bmN0aW9ucyBmcm9tIHNvdXJjZSBmaWxlcyBpbiBwYWNrYWdlIGZtdF90ZXN0KSwKCS8vIGFuZCAiZm10LnRlc3QiICh0aGUgdGVzdCBiaW5hcnkpLgoJLy8KCS8vIEluIGJ1aWxkIHN5c3RlbXMgd2l0aCBleHBsaWNpdCBuYW1lcyBmb3IgdGVzdHMsCgkvLyBzZXR0aW5nIFRlc3RzIG1heSBoYXZlIG5vIGVmZmVjdC4KCVRlc3RzIGJvb2wKCgkvLyBPdmVybGF5IGlzIGEgbWFwcGluZyBmcm9tIGFic29sdXRlIGZpbGUgcGF0aHMgdG8gZmlsZSBjb250ZW50cy4KCS8vCgkvLyBGb3IgZWFjaCBtYXAgZW50cnksIFtMb2FkXSB1c2VzIHRoZSBhbHRlcm5hdGl2ZSBmaWxlCgkvLyBjb250ZW50cyBwcm92aWRlZCBieSB0aGUgb3ZlcmxheSBtYXBwaW5nIGluc3RlYWQgb2YgcmVhZGluZwoJLy8gZnJvbSB0aGUgZmlsZSBzeXN0ZW0uIFRoaXMgbWVjaGFuaXNtIGNhbiBiZSB1c2VkIHRvIGVuYWJsZQoJLy8gZWRpdG9yLWludGVncmF0ZWQgdG9vbHMgdG8gY29ycmVjdGx5IGFuYWx5emUgdGhlIGNvbnRlbnRzCgkvLyBvZiBtb2RpZmllZCBidXQgdW5zYXZlZCBidWZmZXJzLCBmb3IgZXhhbXBsZS4KCS8vCgkvLyBUaGUgb3ZlcmxheSBtYXBwaW5nIGlzIHBhc3NlZCB0byB0aGUgYnVpbGQgc3lzdGVtJ3MgZHJpdmVyCgkvLyAoc2VlICJUaGUgZHJpdmVyIHByb3RvY29sIikgc28gdGhhdCBpdCB0b28gY2FuIHJlcG9ydAoJLy8gY29uc2lzdGVudCBwYWNrYWdlIG1ldGFkYXRhIGFib3V0IHVuc2F2ZWQgZmlsZXMuIEhvd2V2ZXIsCgkvLyBkcml2ZXJzIG1heSB2YXJ5IGluIHRoZWlyIGxldmVsIG9mIHN1cHBvcnQgZm9yIG92ZXJsYXlzLgoJT3ZlcmxheSBtYXBbc3RyaW5nXVtdYnl0ZQp9CgpmdW5jIExvYWQoY2ZnICpDb25maWcsIHBhdHRlcm5zIFtdc3RyaW5nKSAoKnBhY2thZ2VzLkRyaXZlclJlc3BvbnNlLCBlcnJvcikgewoJZm10LlByaW50bG4oImZhIikKCXJldHVybiBuaWwsIG5pbAp9Cg=="
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "wasm",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/fakenet",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/jsonrpc",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/server",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/cmd/gopackagesdriver",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver",
    "builtin"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bufio/net_test.go"],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "builtin",
      "Name": "builtin",
      "PkgPath": "builtin",
      "GoFiles": ["/usr/lib/go/src/builtin/builtin.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/builtin/builtin.go"],
      "Imports": {
        "cmp": "cmp"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/boundary_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 153 to 5 entries.

#### drv #14

Trace meta: spanId=26, ts=1770696392287, ts_iso=2026-02-10T04:06:32.287000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/cmd/gopackagesdriver",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/server"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJmbXQiCgkiZ28vYXN0IgoJImdvL3Rva2VuIgoKCSJnb2xhbmcub3JnL3gvdG9vbHMvZ28vcGFja2FnZXMiCikKCi8vIENvbmZpZyBpcyBwYWNrYWdlIGxvYWQgY29uZmlndXJhdGlvbi4KLy8KLy8gSW5zcGlyZWQgYnkgW2dvbGFuZy5vcmcveC90b29scy9wYWNrYWdlcy5Db25maWddIHN0cnVjdC4KdHlwZSBDb25maWcgc3RydWN0IHsKCS8vIE1vZGUgY29udHJvbHMgdGhlIGxldmVsIG9mIGluZm9ybWF0aW9uIHJldHVybmVkIGZvciBlYWNoIHBhY2thZ2UuCglNb2RlIHBhY2thZ2VzLkxvYWRNb2RlCgoJLy8gQ29udGV4dCBzcGVjaWZpZXMgdGhlIGNvbnRleHQgZm9yIHRoZSBsb2FkIG9wZXJhdGlvbi4KCS8vIENhbmNlbGxpbmcgdGhlIGNvbnRleHQgbWF5IGNhdXNlIFtMb2FkXSB0byBhYm9ydCBhbmQKCS8vIHJldHVybiBhbiBlcnJvci4KCUNvbnRleHQgY29udGV4dC5Db250ZXh0CgoJLy8gRGlyIGlzIHRoZSBkaXJlY3RvcnkgaW4gd2hpY2ggdG8gcnVuIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sCgkvLyB0aGF0IHByb3ZpZGVzIGluZm9ybWF0aW9uIGFib3V0IHRoZSBwYWNrYWdlcy4KCS8vIElmIERpciBpcyBlbXB0eSwgdGhlIHRvb2wgaXMgcnVuIGluIHRoZSBjdXJyZW50IGRpcmVjdG9yeS4KCURpciBzdHJpbmcKCgkvLyBFbnYgaXMgdGhlIGVudmlyb25tZW50IHRvIHVzZSB3aGVuIGludm9raW5nIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sLgoJRW52IG1hcFtzdHJpbmddc3RyaW5nCgoJLy8gQnVpbGRGbGFncyBpcyBhIGxpc3Qgb2YgY29tbWFuZC1saW5lIGZsYWdzIHRvIGJlIHBhc3NlZCB0aHJvdWdoIHRvCgkvLyB0aGUgYnVpbGQgc3lzdGVtJ3MgcXVlcnkgdG9vbC4KCUJ1aWxkRmxhZ3MgW11zdHJpbmcKCgkvLyBGc2V0IHByb3ZpZGVzIHNvdXJjZSBwb3NpdGlvbiBpbmZvcm1hdGlvbiBmb3Igc3ludGF4IHRyZWVzIGFuZCB0eXBlcy4KCS8vIElmIEZzZXQgaXMgbmlsLCBMb2FkIHdpbGwgdXNlIGEgbmV3IGZpbGVzZXQsIGJ1dCBwcmVzZXJ2ZSBGc2V0J3MgdmFsdWUuCglGc2V0ICp0b2tlbi5GaWxlU2V0CgoJLy8gUGFyc2VGaWxlIGlzIGNhbGxlZCB0byByZWFkIGFuZCBwYXJzZSBlYWNoIGZpbGUKCS8vIHdoZW4gcHJlcGFyaW5nIGEgcGFja2FnZSdzIHR5cGUtY2hlY2tlZCBzeW50YXggdHJlZS4KCS8vIEl0IG11c3QgYmUgc2FmZSB0byBjYWxsIFBhcnNlRmlsZSBzaW11bHRhbmVvdXNseSBmcm9tIG11bHRpcGxlIGdvcm91dGluZXMuCgkvLyBJZiBQYXJzZUZpbGUgaXMgbmlsLCB0aGUgbG9hZGVyIHdpbGwgdXNlcyBwYXJzZXIuUGFyc2VGaWxlLgoJLy8KCS8vIFBhcnNlRmlsZSBzaG91bGQgcGFyc2UgdGhlIHNvdXJjZSBmcm9tIHNyYyBhbmQgdXNlIGZpbGVuYW1lIG9ubHkgZm9yCgkvLyByZWNvcmRpbmcgcG9zaXRpb24gaW5mb3JtYXRpb24uCgkvLwoJLy8gQW4gYXBwbGljYXRpb24gbWF5IHN1cHBseSBhIGN1c3RvbSBpbXBsZW1lbnRhdGlvbiBvZiBQYXJzZUZpbGUKCS8vIHRvIGNoYW5nZSB0aGUgZWZmZWN0aXZlIGZpbGUgY29udGVudHMgb3IgdGhlIGJlaGF2aW9yIG9mIHRoZSBwYXJzZXIsCgkvLyBvciB0byBtb2RpZnkgdGhlIHN5bnRheCB0cmVlLiBGb3IgZXhhbXBsZSwgc2VsZWN0aXZlbHkgZWxpbWluYXRpbmcKCS8vIHVud2FudGVkIGZ1bmN0aW9uIGJvZGllcyBjYW4gc2lnbmlmaWNhbnRseSBhY2NlbGVyYXRlIHR5cGUgY2hlY2tpbmcuCglQYXJzZUZpbGUgZnVuYyhmc2V0ICp0b2tlbi5GaWxlU2V0LCBmaWxlbmFtZSBzdHJpbmcsIHNyYyBbXWJ5dGUpICgqYXN0LkZpbGUsIGVycm9yKQoKCS8vIElmIFRlc3RzIGlzIHNldCwgdGhlIGxvYWRlciBpbmNsdWRlcyBub3QganVzdCB0aGUgcGFja2FnZXMKCS8vIG1hdGNoaW5nIGEgcGFydGljdWxhciBwYXR0ZXJuIGJ1dCBhbHNvIGFueSByZWxhdGVkIHRlc3QgcGFja2FnZXMsCgkvLyBpbmNsdWRpbmcgdGVzdC1vbmx5IHZhcmlhbnRzIG9mIHRoZSBwYWNrYWdlIGFuZCB0aGUgdGVzdCBleGVjdXRhYmxlLgoJLy8KCS8vIEZvciBleGFtcGxlLCB3aGVuIHVzaW5nIHRoZSBnbyBjb21tYW5kLCBsb2FkaW5nICJmbXQiIHdpdGggVGVzdHM9dHJ1ZQoJLy8gcmV0dXJucyBmb3VyIHBhY2thZ2VzLCB3aXRoIElEcyAiZm10IiAodGhlIHN0YW5kYXJkIHBhY2thZ2UpLAoJLy8gImZtdCBbZm10LnRlc3RdIiAodGhlIHBhY2thZ2UgYXMgY29tcGlsZWQgZm9yIHRoZSB0ZXN0KSwKCS8vICJmbXRfdGVzdCIgKHRoZSB0ZXN0IGZ1bmN0aW9ucyBmcm9tIHNvdXJjZSBmaWxlcyBpbiBwYWNrYWdlIGZtdF90ZXN0KSwKCS8vIGFuZCAiZm10LnRlc3QiICh0aGUgdGVzdCBiaW5hcnkpLgoJLy8KCS8vIEluIGJ1aWxkIHN5c3RlbXMgd2l0aCBleHBsaWNpdCBuYW1lcyBmb3IgdGVzdHMsCgkvLyBzZXR0aW5nIFRlc3RzIG1heSBoYXZlIG5vIGVmZmVjdC4KCVRlc3RzIGJvb2wKCgkvLyBPdmVybGF5IGlzIGEgbWFwcGluZyBmcm9tIGFic29sdXRlIGZpbGUgcGF0aHMgdG8gZmlsZSBjb250ZW50cy4KCS8vCgkvLyBGb3IgZWFjaCBtYXAgZW50cnksIFtMb2FkXSB1c2VzIHRoZSBhbHRlcm5hdGl2ZSBmaWxlCgkvLyBjb250ZW50cyBwcm92aWRlZCBieSB0aGUgb3ZlcmxheSBtYXBwaW5nIGluc3RlYWQgb2YgcmVhZGluZwoJLy8gZnJvbSB0aGUgZmlsZSBzeXN0ZW0uIFRoaXMgbWVjaGFuaXNtIGNhbiBiZSB1c2VkIHRvIGVuYWJsZQoJLy8gZWRpdG9yLWludGVncmF0ZWQgdG9vbHMgdG8gY29ycmVjdGx5IGFuYWx5emUgdGhlIGNvbnRlbnRzCgkvLyBvZiBtb2RpZmllZCBidXQgdW5zYXZlZCBidWZmZXJzLCBmb3IgZXhhbXBsZS4KCS8vCgkvLyBUaGUgb3ZlcmxheSBtYXBwaW5nIGlzIHBhc3NlZCB0byB0aGUgYnVpbGQgc3lzdGVtJ3MgZHJpdmVyCgkvLyAoc2VlICJUaGUgZHJpdmVyIHByb3RvY29sIikgc28gdGhhdCBpdCB0b28gY2FuIHJlcG9ydAoJLy8gY29uc2lzdGVudCBwYWNrYWdlIG1ldGFkYXRhIGFib3V0IHVuc2F2ZWQgZmlsZXMuIEhvd2V2ZXIsCgkvLyBkcml2ZXJzIG1heSB2YXJ5IGluIHRoZWlyIGxldmVsIG9mIHN1cHBvcnQgZm9yIG92ZXJsYXlzLgoJT3ZlcmxheSBtYXBbc3RyaW5nXVtdYnl0ZQp9CgpmdW5jIExvYWQoY2ZnICpDb25maWcsIHBhdHRlcm5zIFtdc3RyaW5nKSAoKnBhY2thZ2VzLkRyaXZlclJlc3BvbnNlLCBlcnJvcikgewoJZm10LlByaW50bG4oImZhIikKCXJldHVybiBuaWwsIG5pbAp9Cg=="
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/server",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/cmd/gopackagesdriver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 155 to 5 entries.

#### drv #15

Trace meta: spanId=29, ts=1770696395405, ts_iso=2026-02-10T04:06:35.405000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "file=/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJnby9hc3QiCgkiZ28vdG9rZW4iCgoJImdvbGFuZy5vcmcveC90b29scy9nby9wYWNrYWdlcyIKKQoKLy8gQ29uZmlnIGlzIHBhY2thZ2UgbG9hZCBjb25maWd1cmF0aW9uLgovLwovLyBJbnNwaXJlZCBieSBbZ29sYW5nLm9yZy94L3Rvb2xzL3BhY2thZ2VzLkNvbmZpZ10gc3RydWN0Lgp0eXBlIENvbmZpZyBzdHJ1Y3QgewoJLy8gTW9kZSBjb250cm9scyB0aGUgbGV2ZWwgb2YgaW5mb3JtYXRpb24gcmV0dXJuZWQgZm9yIGVhY2ggcGFja2FnZS4KCU1vZGUgcGFja2FnZXMuTG9hZE1vZGUKCgkvLyBDb250ZXh0IHNwZWNpZmllcyB0aGUgY29udGV4dCBmb3IgdGhlIGxvYWQgb3BlcmF0aW9uLgoJLy8gQ2FuY2VsbGluZyB0aGUgY29udGV4dCBtYXkgY2F1c2UgW0xvYWRdIHRvIGFib3J0IGFuZAoJLy8gcmV0dXJuIGFuIGVycm9yLgoJQ29udGV4dCBjb250ZXh0LkNvbnRleHQKCgkvLyBEaXIgaXMgdGhlIGRpcmVjdG9yeSBpbiB3aGljaCB0byBydW4gdGhlIGJ1aWxkIHN5c3RlbSdzIHF1ZXJ5IHRvb2wKCS8vIHRoYXQgcHJvdmlkZXMgaW5mb3JtYXRpb24gYWJvdXQgdGhlIHBhY2thZ2VzLgoJLy8gSWYgRGlyIGlzIGVtcHR5LCB0aGUgdG9vbCBpcyBydW4gaW4gdGhlIGN1cnJlbnQgZGlyZWN0b3J5LgoJRGlyIHN0cmluZwoKCS8vIEVudiBpcyB0aGUgZW52aXJvbm1lbnQgdG8gdXNlIHdoZW4gaW52b2tpbmcgdGhlIGJ1aWxkIHN5c3RlbSdzIHF1ZXJ5IHRvb2wuCglFbnYgbWFwW3N0cmluZ11zdHJpbmcKCgkvLyBCdWlsZEZsYWdzIGlzIGEgbGlzdCBvZiBjb21tYW5kLWxpbmUgZmxhZ3MgdG8gYmUgcGFzc2VkIHRocm91Z2ggdG8KCS8vIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sLgoJQnVpbGRGbGFncyBbXXN0cmluZwoKCS8vIEZzZXQgcHJvdmlkZXMgc291cmNlIHBvc2l0aW9uIGluZm9ybWF0aW9uIGZvciBzeW50YXggdHJlZXMgYW5kIHR5cGVzLgoJLy8gSWYgRnNldCBpcyBuaWwsIExvYWQgd2lsbCB1c2UgYSBuZXcgZmlsZXNldCwgYnV0IHByZXNlcnZlIEZzZXQncyB2YWx1ZS4KCUZzZXQgKnRva2VuLkZpbGVTZXQKCgkvLyBQYXJzZUZpbGUgaXMgY2FsbGVkIHRvIHJlYWQgYW5kIHBhcnNlIGVhY2ggZmlsZQoJLy8gd2hlbiBwcmVwYXJpbmcgYSBwYWNrYWdlJ3MgdHlwZS1jaGVja2VkIHN5bnRheCB0cmVlLgoJLy8gSXQgbXVzdCBiZSBzYWZlIHRvIGNhbGwgUGFyc2VGaWxlIHNpbXVsdGFuZW91c2x5IGZyb20gbXVsdGlwbGUgZ29yb3V0aW5lcy4KCS8vIElmIFBhcnNlRmlsZSBpcyBuaWwsIHRoZSBsb2FkZXIgd2lsbCB1c2VzIHBhcnNlci5QYXJzZUZpbGUuCgkvLwoJLy8gUGFyc2VGaWxlIHNob3VsZCBwYXJzZSB0aGUgc291cmNlIGZyb20gc3JjIGFuZCB1c2UgZmlsZW5hbWUgb25seSBmb3IKCS8vIHJlY29yZGluZyBwb3NpdGlvbiBpbmZvcm1hdGlvbi4KCS8vCgkvLyBBbiBhcHBsaWNhdGlvbiBtYXkgc3VwcGx5IGEgY3VzdG9tIGltcGxlbWVudGF0aW9uIG9mIFBhcnNlRmlsZQoJLy8gdG8gY2hhbmdlIHRoZSBlZmZlY3RpdmUgZmlsZSBjb250ZW50cyBvciB0aGUgYmVoYXZpb3Igb2YgdGhlIHBhcnNlciwKCS8vIG9yIHRvIG1vZGlmeSB0aGUgc3ludGF4IHRyZWUuIEZvciBleGFtcGxlLCBzZWxlY3RpdmVseSBlbGltaW5hdGluZwoJLy8gdW53YW50ZWQgZnVuY3Rpb24gYm9kaWVzIGNhbiBzaWduaWZpY2FudGx5IGFjY2VsZXJhdGUgdHlwZSBjaGVja2luZy4KCVBhcnNlRmlsZSBmdW5jKGZzZXQgKnRva2VuLkZpbGVTZXQsIGZpbGVuYW1lIHN0cmluZywgc3JjIFtdYnl0ZSkgKCphc3QuRmlsZSwgZXJyb3IpCgoJLy8gSWYgVGVzdHMgaXMgc2V0LCB0aGUgbG9hZGVyIGluY2x1ZGVzIG5vdCBqdXN0IHRoZSBwYWNrYWdlcwoJLy8gbWF0Y2hpbmcgYSBwYXJ0aWN1bGFyIHBhdHRlcm4gYnV0IGFsc28gYW55IHJlbGF0ZWQgdGVzdCBwYWNrYWdlcywKCS8vIGluY2x1ZGluZyB0ZXN0LW9ubHkgdmFyaWFudHMgb2YgdGhlIHBhY2thZ2UgYW5kIHRoZSB0ZXN0IGV4ZWN1dGFibGUuCgkvLwoJLy8gRm9yIGV4YW1wbGUsIHdoZW4gdXNpbmcgdGhlIGdvIGNvbW1hbmQsIGxvYWRpbmcgImZtdCIgd2l0aCBUZXN0cz10cnVlCgkvLyByZXR1cm5zIGZvdXIgcGFja2FnZXMsIHdpdGggSURzICJmbXQiICh0aGUgc3RhbmRhcmQgcGFja2FnZSksCgkvLyAiZm10IFtmbXQudGVzdF0iICh0aGUgcGFja2FnZSBhcyBjb21waWxlZCBmb3IgdGhlIHRlc3QpLAoJLy8gImZtdF90ZXN0IiAodGhlIHRlc3QgZnVuY3Rpb25zIGZyb20gc291cmNlIGZpbGVzIGluIHBhY2thZ2UgZm10X3Rlc3QpLAoJLy8gYW5kICJmbXQudGVzdCIgKHRoZSB0ZXN0IGJpbmFyeSkuCgkvLwoJLy8gSW4gYnVpbGQgc3lzdGVtcyB3aXRoIGV4cGxpY2l0IG5hbWVzIGZvciB0ZXN0cywKCS8vIHNldHRpbmcgVGVzdHMgbWF5IGhhdmUgbm8gZWZmZWN0LgoJVGVzdHMgYm9vbAoKCS8vIE92ZXJsYXkgaXMgYSBtYXBwaW5nIGZyb20gYWJzb2x1dGUgZmlsZSBwYXRocyB0byBmaWxlIGNvbnRlbnRzLgoJLy8KCS8vIEZvciBlYWNoIG1hcCBlbnRyeSwgW0xvYWRdIHVzZXMgdGhlIGFsdGVybmF0aXZlIGZpbGUKCS8vIGNvbnRlbnRzIHByb3ZpZGVkIGJ5IHRoZSBvdmVybGF5IG1hcHBpbmcgaW5zdGVhZCBvZiByZWFkaW5nCgkvLyBmcm9tIHRoZSBmaWxlIHN5c3RlbS4gVGhpcyBtZWNoYW5pc20gY2FuIGJlIHVzZWQgdG8gZW5hYmxlCgkvLyBlZGl0b3ItaW50ZWdyYXRlZCB0b29scyB0byBjb3JyZWN0bHkgYW5hbHl6ZSB0aGUgY29udGVudHMKCS8vIG9mIG1vZGlmaWVkIGJ1dCB1bnNhdmVkIGJ1ZmZlcnMsIGZvciBleGFtcGxlLgoJLy8KCS8vIFRoZSBvdmVybGF5IG1hcHBpbmcgaXMgcGFzc2VkIHRvIHRoZSBidWlsZCBzeXN0ZW0ncyBkcml2ZXIKCS8vIChzZWUgIlRoZSBkcml2ZXIgcHJvdG9jb2wiKSBzbyB0aGF0IGl0IHRvbyBjYW4gcmVwb3J0CgkvLyBjb25zaXN0ZW50IHBhY2thZ2UgbWV0YWRhdGEgYWJvdXQgdW5zYXZlZCBmaWxlcy4gSG93ZXZlciwKCS8vIGRyaXZlcnMgbWF5IHZhcnkgaW4gdGhlaXIgbGV2ZWwgb2Ygc3VwcG9ydCBmb3Igb3ZlcmxheXMuCglPdmVybGF5IG1hcFtzdHJpbmddW11ieXRlCn0KCmZ1bmMgTG9hZChjZmcgKkNvbmZpZywgcGF0dGVybnMgW11zdHJpbmcpICgqcGFja2FnZXMuRHJpdmVyUmVzcG9uc2UsIGVycm9yKSB7CglyZXR1cm4gbmlsLCBuaWwKfQo="
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 0
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #16

Trace meta: spanId=28, ts=1770696395404, ts_iso=2026-02-10T04:06:35.404000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJnby9hc3QiCgkiZ28vdG9rZW4iCgoJImdvbGFuZy5vcmcveC90b29scy9nby9wYWNrYWdlcyIKKQoKLy8gQ29uZmlnIGlzIHBhY2thZ2UgbG9hZCBjb25maWd1cmF0aW9uLgovLwovLyBJbnNwaXJlZCBieSBbZ29sYW5nLm9yZy94L3Rvb2xzL3BhY2thZ2VzLkNvbmZpZ10gc3RydWN0Lgp0eXBlIENvbmZpZyBzdHJ1Y3QgewoJLy8gTW9kZSBjb250cm9scyB0aGUgbGV2ZWwgb2YgaW5mb3JtYXRpb24gcmV0dXJuZWQgZm9yIGVhY2ggcGFja2FnZS4KCU1vZGUgcGFja2FnZXMuTG9hZE1vZGUKCgkvLyBDb250ZXh0IHNwZWNpZmllcyB0aGUgY29udGV4dCBmb3IgdGhlIGxvYWQgb3BlcmF0aW9uLgoJLy8gQ2FuY2VsbGluZyB0aGUgY29udGV4dCBtYXkgY2F1c2UgW0xvYWRdIHRvIGFib3J0IGFuZAoJLy8gcmV0dXJuIGFuIGVycm9yLgoJQ29udGV4dCBjb250ZXh0LkNvbnRleHQKCgkvLyBEaXIgaXMgdGhlIGRpcmVjdG9yeSBpbiB3aGljaCB0byBydW4gdGhlIGJ1aWxkIHN5c3RlbSdzIHF1ZXJ5IHRvb2wKCS8vIHRoYXQgcHJvdmlkZXMgaW5mb3JtYXRpb24gYWJvdXQgdGhlIHBhY2thZ2VzLgoJLy8gSWYgRGlyIGlzIGVtcHR5LCB0aGUgdG9vbCBpcyBydW4gaW4gdGhlIGN1cnJlbnQgZGlyZWN0b3J5LgoJRGlyIHN0cmluZwoKCS8vIEVudiBpcyB0aGUgZW52aXJvbm1lbnQgdG8gdXNlIHdoZW4gaW52b2tpbmcgdGhlIGJ1aWxkIHN5c3RlbSdzIHF1ZXJ5IHRvb2wuCglFbnYgbWFwW3N0cmluZ11zdHJpbmcKCgkvLyBCdWlsZEZsYWdzIGlzIGEgbGlzdCBvZiBjb21tYW5kLWxpbmUgZmxhZ3MgdG8gYmUgcGFzc2VkIHRocm91Z2ggdG8KCS8vIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sLgoJQnVpbGRGbGFncyBbXXN0cmluZwoKCS8vIEZzZXQgcHJvdmlkZXMgc291cmNlIHBvc2l0aW9uIGluZm9ybWF0aW9uIGZvciBzeW50YXggdHJlZXMgYW5kIHR5cGVzLgoJLy8gSWYgRnNldCBpcyBuaWwsIExvYWQgd2lsbCB1c2UgYSBuZXcgZmlsZXNldCwgYnV0IHByZXNlcnZlIEZzZXQncyB2YWx1ZS4KCUZzZXQgKnRva2VuLkZpbGVTZXQKCgkvLyBQYXJzZUZpbGUgaXMgY2FsbGVkIHRvIHJlYWQgYW5kIHBhcnNlIGVhY2ggZmlsZQoJLy8gd2hlbiBwcmVwYXJpbmcgYSBwYWNrYWdlJ3MgdHlwZS1jaGVja2VkIHN5bnRheCB0cmVlLgoJLy8gSXQgbXVzdCBiZSBzYWZlIHRvIGNhbGwgUGFyc2VGaWxlIHNpbXVsdGFuZW91c2x5IGZyb20gbXVsdGlwbGUgZ29yb3V0aW5lcy4KCS8vIElmIFBhcnNlRmlsZSBpcyBuaWwsIHRoZSBsb2FkZXIgd2lsbCB1c2VzIHBhcnNlci5QYXJzZUZpbGUuCgkvLwoJLy8gUGFyc2VGaWxlIHNob3VsZCBwYXJzZSB0aGUgc291cmNlIGZyb20gc3JjIGFuZCB1c2UgZmlsZW5hbWUgb25seSBmb3IKCS8vIHJlY29yZGluZyBwb3NpdGlvbiBpbmZvcm1hdGlvbi4KCS8vCgkvLyBBbiBhcHBsaWNhdGlvbiBtYXkgc3VwcGx5IGEgY3VzdG9tIGltcGxlbWVudGF0aW9uIG9mIFBhcnNlRmlsZQoJLy8gdG8gY2hhbmdlIHRoZSBlZmZlY3RpdmUgZmlsZSBjb250ZW50cyBvciB0aGUgYmVoYXZpb3Igb2YgdGhlIHBhcnNlciwKCS8vIG9yIHRvIG1vZGlmeSB0aGUgc3ludGF4IHRyZWUuIEZvciBleGFtcGxlLCBzZWxlY3RpdmVseSBlbGltaW5hdGluZwoJLy8gdW53YW50ZWQgZnVuY3Rpb24gYm9kaWVzIGNhbiBzaWduaWZpY2FudGx5IGFjY2VsZXJhdGUgdHlwZSBjaGVja2luZy4KCVBhcnNlRmlsZSBmdW5jKGZzZXQgKnRva2VuLkZpbGVTZXQsIGZpbGVuYW1lIHN0cmluZywgc3JjIFtdYnl0ZSkgKCphc3QuRmlsZSwgZXJyb3IpCgoJLy8gSWYgVGVzdHMgaXMgc2V0LCB0aGUgbG9hZGVyIGluY2x1ZGVzIG5vdCBqdXN0IHRoZSBwYWNrYWdlcwoJLy8gbWF0Y2hpbmcgYSBwYXJ0aWN1bGFyIHBhdHRlcm4gYnV0IGFsc28gYW55IHJlbGF0ZWQgdGVzdCBwYWNrYWdlcywKCS8vIGluY2x1ZGluZyB0ZXN0LW9ubHkgdmFyaWFudHMgb2YgdGhlIHBhY2thZ2UgYW5kIHRoZSB0ZXN0IGV4ZWN1dGFibGUuCgkvLwoJLy8gRm9yIGV4YW1wbGUsIHdoZW4gdXNpbmcgdGhlIGdvIGNvbW1hbmQsIGxvYWRpbmcgImZtdCIgd2l0aCBUZXN0cz10cnVlCgkvLyByZXR1cm5zIGZvdXIgcGFja2FnZXMsIHdpdGggSURzICJmbXQiICh0aGUgc3RhbmRhcmQgcGFja2FnZSksCgkvLyAiZm10IFtmbXQudGVzdF0iICh0aGUgcGFja2FnZSBhcyBjb21waWxlZCBmb3IgdGhlIHRlc3QpLAoJLy8gImZtdF90ZXN0IiAodGhlIHRlc3QgZnVuY3Rpb25zIGZyb20gc291cmNlIGZpbGVzIGluIHBhY2thZ2UgZm10X3Rlc3QpLAoJLy8gYW5kICJmbXQudGVzdCIgKHRoZSB0ZXN0IGJpbmFyeSkuCgkvLwoJLy8gSW4gYnVpbGQgc3lzdGVtcyB3aXRoIGV4cGxpY2l0IG5hbWVzIGZvciB0ZXN0cywKCS8vIHNldHRpbmcgVGVzdHMgbWF5IGhhdmUgbm8gZWZmZWN0LgoJVGVzdHMgYm9vbAoKCS8vIE92ZXJsYXkgaXMgYSBtYXBwaW5nIGZyb20gYWJzb2x1dGUgZmlsZSBwYXRocyB0byBmaWxlIGNvbnRlbnRzLgoJLy8KCS8vIEZvciBlYWNoIG1hcCBlbnRyeSwgW0xvYWRdIHVzZXMgdGhlIGFsdGVybmF0aXZlIGZpbGUKCS8vIGNvbnRlbnRzIHByb3ZpZGVkIGJ5IHRoZSBvdmVybGF5IG1hcHBpbmcgaW5zdGVhZCBvZiByZWFkaW5nCgkvLyBmcm9tIHRoZSBmaWxlIHN5c3RlbS4gVGhpcyBtZWNoYW5pc20gY2FuIGJlIHVzZWQgdG8gZW5hYmxlCgkvLyBlZGl0b3ItaW50ZWdyYXRlZCB0b29scyB0byBjb3JyZWN0bHkgYW5hbHl6ZSB0aGUgY29udGVudHMKCS8vIG9mIG1vZGlmaWVkIGJ1dCB1bnNhdmVkIGJ1ZmZlcnMsIGZvciBleGFtcGxlLgoJLy8KCS8vIFRoZSBvdmVybGF5IG1hcHBpbmcgaXMgcGFzc2VkIHRvIHRoZSBidWlsZCBzeXN0ZW0ncyBkcml2ZXIKCS8vIChzZWUgIlRoZSBkcml2ZXIgcHJvdG9jb2wiKSBzbyB0aGF0IGl0IHRvbyBjYW4gcmVwb3J0CgkvLyBjb25zaXN0ZW50IHBhY2thZ2UgbWV0YWRhdGEgYWJvdXQgdW5zYXZlZCBmaWxlcy4gSG93ZXZlciwKCS8vIGRyaXZlcnMgbWF5IHZhcnkgaW4gdGhlaXIgbGV2ZWwgb2Ygc3VwcG9ydCBmb3Igb3ZlcmxheXMuCglPdmVybGF5IG1hcFtzdHJpbmddW11ieXRlCn0KCmZ1bmMgTG9hZChjZmcgKkNvbmZpZywgcGF0dGVybnMgW11zdHJpbmcpICgqcGFja2FnZXMuRHJpdmVyUmVzcG9uc2UsIGVycm9yKSB7CglyZXR1cm4gbmlsLCBuaWwKfQo="
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 143 to 5 entries.

#### drv #17

Trace meta: spanId=32, ts=1770696395447, ts_iso=2026-02-10T04:06:35.447000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GO111MODULE=auto",
      "GOARCH=wasm",
      "GOOS=js",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJnby9hc3QiCgkiZ28vdG9rZW4iCgoJImdvbGFuZy5vcmcveC90b29scy9nby9wYWNrYWdlcyIKKQoKLy8gQ29uZmlnIGlzIHBhY2thZ2UgbG9hZCBjb25maWd1cmF0aW9uLgovLwovLyBJbnNwaXJlZCBieSBbZ29sYW5nLm9yZy94L3Rvb2xzL3BhY2thZ2VzLkNvbmZpZ10gc3RydWN0Lgp0eXBlIENvbmZpZyBzdHJ1Y3QgewoJLy8gTW9kZSBjb250cm9scyB0aGUgbGV2ZWwgb2YgaW5mb3JtYXRpb24gcmV0dXJuZWQgZm9yIGVhY2ggcGFja2FnZS4KCU1vZGUgcGFja2FnZXMuTG9hZE1vZGUKCgkvLyBDb250ZXh0IHNwZWNpZmllcyB0aGUgY29udGV4dCBmb3IgdGhlIGxvYWQgb3BlcmF0aW9uLgoJLy8gQ2FuY2VsbGluZyB0aGUgY29udGV4dCBtYXkgY2F1c2UgW0xvYWRdIHRvIGFib3J0IGFuZAoJLy8gcmV0dXJuIGFuIGVycm9yLgoJQ29udGV4dCBjb250ZXh0LkNvbnRleHQKCgkvLyBEaXIgaXMgdGhlIGRpcmVjdG9yeSBpbiB3aGljaCB0byBydW4gdGhlIGJ1aWxkIHN5c3RlbSdzIHF1ZXJ5IHRvb2wKCS8vIHRoYXQgcHJvdmlkZXMgaW5mb3JtYXRpb24gYWJvdXQgdGhlIHBhY2thZ2VzLgoJLy8gSWYgRGlyIGlzIGVtcHR5LCB0aGUgdG9vbCBpcyBydW4gaW4gdGhlIGN1cnJlbnQgZGlyZWN0b3J5LgoJRGlyIHN0cmluZwoKCS8vIEVudiBpcyB0aGUgZW52aXJvbm1lbnQgdG8gdXNlIHdoZW4gaW52b2tpbmcgdGhlIGJ1aWxkIHN5c3RlbSdzIHF1ZXJ5IHRvb2wuCglFbnYgbWFwW3N0cmluZ11zdHJpbmcKCgkvLyBCdWlsZEZsYWdzIGlzIGEgbGlzdCBvZiBjb21tYW5kLWxpbmUgZmxhZ3MgdG8gYmUgcGFzc2VkIHRocm91Z2ggdG8KCS8vIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sLgoJQnVpbGRGbGFncyBbXXN0cmluZwoKCS8vIEZzZXQgcHJvdmlkZXMgc291cmNlIHBvc2l0aW9uIGluZm9ybWF0aW9uIGZvciBzeW50YXggdHJlZXMgYW5kIHR5cGVzLgoJLy8gSWYgRnNldCBpcyBuaWwsIExvYWQgd2lsbCB1c2UgYSBuZXcgZmlsZXNldCwgYnV0IHByZXNlcnZlIEZzZXQncyB2YWx1ZS4KCUZzZXQgKnRva2VuLkZpbGVTZXQKCgkvLyBQYXJzZUZpbGUgaXMgY2FsbGVkIHRvIHJlYWQgYW5kIHBhcnNlIGVhY2ggZmlsZQoJLy8gd2hlbiBwcmVwYXJpbmcgYSBwYWNrYWdlJ3MgdHlwZS1jaGVja2VkIHN5bnRheCB0cmVlLgoJLy8gSXQgbXVzdCBiZSBzYWZlIHRvIGNhbGwgUGFyc2VGaWxlIHNpbXVsdGFuZW91c2x5IGZyb20gbXVsdGlwbGUgZ29yb3V0aW5lcy4KCS8vIElmIFBhcnNlRmlsZSBpcyBuaWwsIHRoZSBsb2FkZXIgd2lsbCB1c2VzIHBhcnNlci5QYXJzZUZpbGUuCgkvLwoJLy8gUGFyc2VGaWxlIHNob3VsZCBwYXJzZSB0aGUgc291cmNlIGZyb20gc3JjIGFuZCB1c2UgZmlsZW5hbWUgb25seSBmb3IKCS8vIHJlY29yZGluZyBwb3NpdGlvbiBpbmZvcm1hdGlvbi4KCS8vCgkvLyBBbiBhcHBsaWNhdGlvbiBtYXkgc3VwcGx5IGEgY3VzdG9tIGltcGxlbWVudGF0aW9uIG9mIFBhcnNlRmlsZQoJLy8gdG8gY2hhbmdlIHRoZSBlZmZlY3RpdmUgZmlsZSBjb250ZW50cyBvciB0aGUgYmVoYXZpb3Igb2YgdGhlIHBhcnNlciwKCS8vIG9yIHRvIG1vZGlmeSB0aGUgc3ludGF4IHRyZWUuIEZvciBleGFtcGxlLCBzZWxlY3RpdmVseSBlbGltaW5hdGluZwoJLy8gdW53YW50ZWQgZnVuY3Rpb24gYm9kaWVzIGNhbiBzaWduaWZpY2FudGx5IGFjY2VsZXJhdGUgdHlwZSBjaGVja2luZy4KCVBhcnNlRmlsZSBmdW5jKGZzZXQgKnRva2VuLkZpbGVTZXQsIGZpbGVuYW1lIHN0cmluZywgc3JjIFtdYnl0ZSkgKCphc3QuRmlsZSwgZXJyb3IpCgoJLy8gSWYgVGVzdHMgaXMgc2V0LCB0aGUgbG9hZGVyIGluY2x1ZGVzIG5vdCBqdXN0IHRoZSBwYWNrYWdlcwoJLy8gbWF0Y2hpbmcgYSBwYXJ0aWN1bGFyIHBhdHRlcm4gYnV0IGFsc28gYW55IHJlbGF0ZWQgdGVzdCBwYWNrYWdlcywKCS8vIGluY2x1ZGluZyB0ZXN0LW9ubHkgdmFyaWFudHMgb2YgdGhlIHBhY2thZ2UgYW5kIHRoZSB0ZXN0IGV4ZWN1dGFibGUuCgkvLwoJLy8gRm9yIGV4YW1wbGUsIHdoZW4gdXNpbmcgdGhlIGdvIGNvbW1hbmQsIGxvYWRpbmcgImZtdCIgd2l0aCBUZXN0cz10cnVlCgkvLyByZXR1cm5zIGZvdXIgcGFja2FnZXMsIHdpdGggSURzICJmbXQiICh0aGUgc3RhbmRhcmQgcGFja2FnZSksCgkvLyAiZm10IFtmbXQudGVzdF0iICh0aGUgcGFja2FnZSBhcyBjb21waWxlZCBmb3IgdGhlIHRlc3QpLAoJLy8gImZtdF90ZXN0IiAodGhlIHRlc3QgZnVuY3Rpb25zIGZyb20gc291cmNlIGZpbGVzIGluIHBhY2thZ2UgZm10X3Rlc3QpLAoJLy8gYW5kICJmbXQudGVzdCIgKHRoZSB0ZXN0IGJpbmFyeSkuCgkvLwoJLy8gSW4gYnVpbGQgc3lzdGVtcyB3aXRoIGV4cGxpY2l0IG5hbWVzIGZvciB0ZXN0cywKCS8vIHNldHRpbmcgVGVzdHMgbWF5IGhhdmUgbm8gZWZmZWN0LgoJVGVzdHMgYm9vbAoKCS8vIE92ZXJsYXkgaXMgYSBtYXBwaW5nIGZyb20gYWJzb2x1dGUgZmlsZSBwYXRocyB0byBmaWxlIGNvbnRlbnRzLgoJLy8KCS8vIEZvciBlYWNoIG1hcCBlbnRyeSwgW0xvYWRdIHVzZXMgdGhlIGFsdGVybmF0aXZlIGZpbGUKCS8vIGNvbnRlbnRzIHByb3ZpZGVkIGJ5IHRoZSBvdmVybGF5IG1hcHBpbmcgaW5zdGVhZCBvZiByZWFkaW5nCgkvLyBmcm9tIHRoZSBmaWxlIHN5c3RlbS4gVGhpcyBtZWNoYW5pc20gY2FuIGJlIHVzZWQgdG8gZW5hYmxlCgkvLyBlZGl0b3ItaW50ZWdyYXRlZCB0b29scyB0byBjb3JyZWN0bHkgYW5hbHl6ZSB0aGUgY29udGVudHMKCS8vIG9mIG1vZGlmaWVkIGJ1dCB1bnNhdmVkIGJ1ZmZlcnMsIGZvciBleGFtcGxlLgoJLy8KCS8vIFRoZSBvdmVybGF5IG1hcHBpbmcgaXMgcGFzc2VkIHRvIHRoZSBidWlsZCBzeXN0ZW0ncyBkcml2ZXIKCS8vIChzZWUgIlRoZSBkcml2ZXIgcHJvdG9jb2wiKSBzbyB0aGF0IGl0IHRvbyBjYW4gcmVwb3J0CgkvLyBjb25zaXN0ZW50IHBhY2thZ2UgbWV0YWRhdGEgYWJvdXQgdW5zYXZlZCBmaWxlcy4gSG93ZXZlciwKCS8vIGRyaXZlcnMgbWF5IHZhcnkgaW4gdGhlaXIgbGV2ZWwgb2Ygc3VwcG9ydCBmb3Igb3ZlcmxheXMuCglPdmVybGF5IG1hcFtzdHJpbmddW11ieXRlCn0KCmZ1bmMgTG9hZChjZmcgKkNvbmZpZywgcGF0dGVybnMgW11zdHJpbmcpICgqcGFja2FnZXMuRHJpdmVyUmVzcG9uc2UsIGVycm9yKSB7CglyZXR1cm4gbmlsLCBuaWwKfQo="
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "wasm",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/driver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bufio/net_test.go"],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/boundary_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 141 to 5 entries.

#### drv #18

Trace meta: spanId=34, ts=1770696406199, ts_iso=2026-02-10T04:06:46.199000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/prj/go-packages-driver-wasm",
  "patterns": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/cmd/gopackagesdriver",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/server"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "CGO_FFLAGS=-O2 -g",
      "GONOSUMDB=",
      "AR=ar",
      "GOFLAGS=",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1927722828=/tmp/go-build -gno-record-gcc-switches",
      "GODEBUG=",
      "GONOPROXY=",
      "GOTELEMETRY=off",
      "GOPRIVATE=",
      "GOAMD64=v1",
      "GOPROXY=https://proxy.golang.org,direct",
      "CGO_ENABLED=1",
      "GOTMPDIR=",
      "CGO_CPPFLAGS=",
      "GOHOSTARCH=amd64",
      "CGO_CXXFLAGS=-O2 -g",
      "GOEXPERIMENT=nodwarf5",
      "CC=gcc",
      "CGO_LDFLAGS=-O2 -g",
      "GOFIPS140=off",
      "GOENV=/home/username/.config/go/env",
      "PKG_CONFIG=pkg-config",
      "GO111MODULE=",
      "GOOS=linux",
      "GOHOSTOS=linux",
      "GOPATH=/home/username/go",
      "GOBIN=",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "CXX=g++",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOAUTH=netrc",
      "GOTOOLCHAIN=auto",
      "GOROOT=/usr/lib/go",
      "GOVERSION=go1.25.6 X:nodwarf5",
      "GOMOD=/home/username/prj/go-packages-driver-wasm/go.mod",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOVCS=",
      "GOSUMDB=sum.golang.org",
      "GOARCH=amd64",
      "GOCACHE=/home/username/.cache/go-build",
      "GOEXE=",
      "GO111MODULE=auto",
      "GOARCH=wasm",
      "GOOS=js",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/prj/go-packages-driver-wasm/internal/driver/loader.go": "cGFja2FnZSBkcml2ZXIKCmltcG9ydCAoCgkiY29udGV4dCIKCSJnby9hc3QiCgkiZ28vdG9rZW4iCgoJImdvbGFuZy5vcmcveC90b29scy9nby9wYWNrYWdlcyIKKQoKLy8gQ29uZmlnIGlzIHBhY2thZ2UgbG9hZCBjb25maWd1cmF0aW9uLgovLwovLyBJbnNwaXJlZCBieSBbZ29sYW5nLm9yZy94L3Rvb2xzL3BhY2thZ2VzLkNvbmZpZ10gc3RydWN0Lgp0eXBlIENvbmZpZyBzdHJ1Y3QgewoJLy8gTW9kZSBjb250cm9scyB0aGUgbGV2ZWwgb2YgaW5mb3JtYXRpb24gcmV0dXJuZWQgZm9yIGVhY2ggcGFja2FnZS4KCU1vZGUgcGFja2FnZXMuTG9hZE1vZGUKCgkvLyBDb250ZXh0IHNwZWNpZmllcyB0aGUgY29udGV4dCBmb3IgdGhlIGxvYWQgb3BlcmF0aW9uLgoJLy8gQ2FuY2VsbGluZyB0aGUgY29udGV4dCBtYXkgY2F1c2UgW0xvYWRdIHRvIGFib3J0IGFuZAoJLy8gcmV0dXJuIGFuIGVycm9yLgoJQ29udGV4dCBjb250ZXh0LkNvbnRleHQKCgkvLyBEaXIgaXMgdGhlIGRpcmVjdG9yeSBpbiB3aGljaCB0byBydW4gdGhlIGJ1aWxkIHN5c3RlbSdzIHF1ZXJ5IHRvb2wKCS8vIHRoYXQgcHJvdmlkZXMgaW5mb3JtYXRpb24gYWJvdXQgdGhlIHBhY2thZ2VzLgoJLy8gSWYgRGlyIGlzIGVtcHR5LCB0aGUgdG9vbCBpcyBydW4gaW4gdGhlIGN1cnJlbnQgZGlyZWN0b3J5LgoJRGlyIHN0cmluZwoKCS8vIEVudiBpcyB0aGUgZW52aXJvbm1lbnQgdG8gdXNlIHdoZW4gaW52b2tpbmcgdGhlIGJ1aWxkIHN5c3RlbSdzIHF1ZXJ5IHRvb2wuCglFbnYgbWFwW3N0cmluZ11zdHJpbmcKCgkvLyBCdWlsZEZsYWdzIGlzIGEgbGlzdCBvZiBjb21tYW5kLWxpbmUgZmxhZ3MgdG8gYmUgcGFzc2VkIHRocm91Z2ggdG8KCS8vIHRoZSBidWlsZCBzeXN0ZW0ncyBxdWVyeSB0b29sLgoJQnVpbGRGbGFncyBbXXN0cmluZwoKCS8vIEZzZXQgcHJvdmlkZXMgc291cmNlIHBvc2l0aW9uIGluZm9ybWF0aW9uIGZvciBzeW50YXggdHJlZXMgYW5kIHR5cGVzLgoJLy8gSWYgRnNldCBpcyBuaWwsIExvYWQgd2lsbCB1c2UgYSBuZXcgZmlsZXNldCwgYnV0IHByZXNlcnZlIEZzZXQncyB2YWx1ZS4KCUZzZXQgKnRva2VuLkZpbGVTZXQKCgkvLyBQYXJzZUZpbGUgaXMgY2FsbGVkIHRvIHJlYWQgYW5kIHBhcnNlIGVhY2ggZmlsZQoJLy8gd2hlbiBwcmVwYXJpbmcgYSBwYWNrYWdlJ3MgdHlwZS1jaGVja2VkIHN5bnRheCB0cmVlLgoJLy8gSXQgbXVzdCBiZSBzYWZlIHRvIGNhbGwgUGFyc2VGaWxlIHNpbXVsdGFuZW91c2x5IGZyb20gbXVsdGlwbGUgZ29yb3V0aW5lcy4KCS8vIElmIFBhcnNlRmlsZSBpcyBuaWwsIHRoZSBsb2FkZXIgd2lsbCB1c2VzIHBhcnNlci5QYXJzZUZpbGUuCgkvLwoJLy8gUGFyc2VGaWxlIHNob3VsZCBwYXJzZSB0aGUgc291cmNlIGZyb20gc3JjIGFuZCB1c2UgZmlsZW5hbWUgb25seSBmb3IKCS8vIHJlY29yZGluZyBwb3NpdGlvbiBpbmZvcm1hdGlvbi4KCS8vCgkvLyBBbiBhcHBsaWNhdGlvbiBtYXkgc3VwcGx5IGEgY3VzdG9tIGltcGxlbWVudGF0aW9uIG9mIFBhcnNlRmlsZQoJLy8gdG8gY2hhbmdlIHRoZSBlZmZlY3RpdmUgZmlsZSBjb250ZW50cyBvciB0aGUgYmVoYXZpb3Igb2YgdGhlIHBhcnNlciwKCS8vIG9yIHRvIG1vZGlmeSB0aGUgc3ludGF4IHRyZWUuIEZvciBleGFtcGxlLCBzZWxlY3RpdmVseSBlbGltaW5hdGluZwoJLy8gdW53YW50ZWQgZnVuY3Rpb24gYm9kaWVzIGNhbiBzaWduaWZpY2FudGx5IGFjY2VsZXJhdGUgdHlwZSBjaGVja2luZy4KCVBhcnNlRmlsZSBmdW5jKGZzZXQgKnRva2VuLkZpbGVTZXQsIGZpbGVuYW1lIHN0cmluZywgc3JjIFtdYnl0ZSkgKCphc3QuRmlsZSwgZXJyb3IpCgoJLy8gSWYgVGVzdHMgaXMgc2V0LCB0aGUgbG9hZGVyIGluY2x1ZGVzIG5vdCBqdXN0IHRoZSBwYWNrYWdlcwoJLy8gbWF0Y2hpbmcgYSBwYXJ0aWN1bGFyIHBhdHRlcm4gYnV0IGFsc28gYW55IHJlbGF0ZWQgdGVzdCBwYWNrYWdlcywKCS8vIGluY2x1ZGluZyB0ZXN0LW9ubHkgdmFyaWFudHMgb2YgdGhlIHBhY2thZ2UgYW5kIHRoZSB0ZXN0IGV4ZWN1dGFibGUuCgkvLwoJLy8gRm9yIGV4YW1wbGUsIHdoZW4gdXNpbmcgdGhlIGdvIGNvbW1hbmQsIGxvYWRpbmcgImZtdCIgd2l0aCBUZXN0cz10cnVlCgkvLyByZXR1cm5zIGZvdXIgcGFja2FnZXMsIHdpdGggSURzICJmbXQiICh0aGUgc3RhbmRhcmQgcGFja2FnZSksCgkvLyAiZm10IFtmbXQudGVzdF0iICh0aGUgcGFja2FnZSBhcyBjb21waWxlZCBmb3IgdGhlIHRlc3QpLAoJLy8gImZtdF90ZXN0IiAodGhlIHRlc3QgZnVuY3Rpb25zIGZyb20gc291cmNlIGZpbGVzIGluIHBhY2thZ2UgZm10X3Rlc3QpLAoJLy8gYW5kICJmbXQudGVzdCIgKHRoZSB0ZXN0IGJpbmFyeSkuCgkvLwoJLy8gSW4gYnVpbGQgc3lzdGVtcyB3aXRoIGV4cGxpY2l0IG5hbWVzIGZvciB0ZXN0cywKCS8vIHNldHRpbmcgVGVzdHMgbWF5IGhhdmUgbm8gZWZmZWN0LgoJVGVzdHMgYm9vbAoKCS8vIE92ZXJsYXkgaXMgYSBtYXBwaW5nIGZyb20gYWJzb2x1dGUgZmlsZSBwYXRocyB0byBmaWxlIGNvbnRlbnRzLgoJLy8KCS8vIEZvciBlYWNoIG1hcCBlbnRyeSwgW0xvYWRdIHVzZXMgdGhlIGFsdGVybmF0aXZlIGZpbGUKCS8vIGNvbnRlbnRzIHByb3ZpZGVkIGJ5IHRoZSBvdmVybGF5IG1hcHBpbmcgaW5zdGVhZCBvZiByZWFkaW5nCgkvLyBmcm9tIHRoZSBmaWxlIHN5c3RlbS4gVGhpcyBtZWNoYW5pc20gY2FuIGJlIHVzZWQgdG8gZW5hYmxlCgkvLyBlZGl0b3ItaW50ZWdyYXRlZCB0b29scyB0byBjb3JyZWN0bHkgYW5hbHl6ZSB0aGUgY29udGVudHMKCS8vIG9mIG1vZGlmaWVkIGJ1dCB1bnNhdmVkIGJ1ZmZlcnMsIGZvciBleGFtcGxlLgoJLy8KCS8vIFRoZSBvdmVybGF5IG1hcHBpbmcgaXMgcGFzc2VkIHRvIHRoZSBidWlsZCBzeXN0ZW0ncyBkcml2ZXIKCS8vIChzZWUgIlRoZSBkcml2ZXIgcHJvdG9jb2wiKSBzbyB0aGF0IGl0IHRvbyBjYW4gcmVwb3J0CgkvLyBjb25zaXN0ZW50IHBhY2thZ2UgbWV0YWRhdGEgYWJvdXQgdW5zYXZlZCBmaWxlcy4gSG93ZXZlciwKCS8vIGRyaXZlcnMgbWF5IHZhcnkgaW4gdGhlaXIgbGV2ZWwgb2Ygc3VwcG9ydCBmb3Igb3ZlcmxheXMuCglPdmVybGF5IG1hcFtzdHJpbmddW11ieXRlCn0KCmZ1bmMgTG9hZChjZmcgKkNvbmZpZywgcGF0dGVybnMgW11zdHJpbmcpICgqcGFja2FnZXMuRHJpdmVyUmVzcG9uc2UsIGVycm9yKSB7CglyZXR1cm4gbmlsLCBuaWwKfQo="
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "wasm",
  "Roots": [
    "github.com/Better-Go-Playground/go-packages-driver-wasm/internal/server",
    "github.com/Better-Go-Playground/go-packages-driver-wasm/cmd/gopackagesdriver"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bufio/net_test.go"],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/boundary_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cmp",
      "Name": "cmp",
      "PkgPath": "cmp",
      "GoFiles": ["/usr/lib/go/src/cmp/cmp.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/cmp/cmp.go"]
    },
    {
      "ID": "container/heap",
      "Name": "heap",
      "PkgPath": "container/heap",
      "GoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/container/heap/heap.go"],
      "Imports": {
        "sort": "sort"
      }
    },
    {
      "ID": "context",
      "Name": "context",
      "PkgPath": "context",
      "GoFiles": ["/usr/lib/go/src/context/context.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/context/context.go"],
      "Imports": {
        "errors": "errors",
        "internal/reflectlite": "internal/reflectlite",
        "sync": "sync",
        "sync/atomic": "sync/atomic",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 151 to 5 entries.

### 2.jsonl

Total drv events: 13

#### drv #1

Trace meta: spanId=1, ts=1770837112298, ts_iso=2026-02-11T19:11:52.298000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": ["/home/username/work/grafana/alloy/...", "builtin"],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GO111MODULE=auto"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {}
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/grafana/alloy/internal/component/common/loki",
    "github.com/grafana/alloy/internal/slogadapter",
    "github.com/grafana/alloy/internal/runtime/logging",
    "github.com/grafana/alloy/internal/runtime/logging/level",
    "github.com/grafana/alloy/internal/alloyseed",
    "github.com/grafana/alloy/internal/boringcrypto",
    "github.com/grafana/alloy/internal/build",
    "github.com/grafana/alloy/internal/featuregate",
    "github.com/grafana/alloy/internal/component",
    "github.com/grafana/alloy/internal/component/common/relabel",
    "github.com/grafana/alloy/internal/runtime/equality",
    "github.com/grafana/alloy/internal/service",
    "github.com/grafana/alloy/internal/service/cluster/discovery",
    "github.com/grafana/alloy/internal/component/common/config",
    "github.com/grafana/alloy/internal/dag",
    "github.com/grafana/alloy/internal/nodeconf/argument",
    "github.com/grafana/alloy/internal/nodeconf/export",
    "github.com/grafana/alloy/internal/nodeconf/foreach",
    "github.com/grafana/alloy/internal/useragent",
    "github.com/grafana/alloy/internal/component/remote/http",
    "github.com/grafana/alloy/internal/filedetector/internal/inotifyinfo",
    "github.com/grafana/alloy/internal/filedetector",
    "github.com/grafana/alloy/internal/util",
    "github.com/grafana/alloy/internal/vcs",
    "github.com/grafana/alloy/internal/nodeconf/importsource",
    "github.com/grafana/alloy/internal/component/otelcol/config",
    "github.com/grafana/alloy/internal/component/otelcol/internal/lazycollector",
    "github.com/grafana/alloy/internal/component/otelcol/internal/scheduler",
    "github.com/grafana/alloy/internal/component/otelcol/util",
    "github.com/grafana/alloy/internal/util/zapadapter",
    "github.com/grafana/alloy/internal/component/otelcol/auth",
    "github.com/grafana/alloy/internal/component/otelcol/extension",
    "github.com/grafana/alloy/internal/component/otelcol",
    "github.com/grafana/alloy/internal/component/pyroscope",
    "github.com/grafana/alloy/internal/runner",
    "github.com/grafana/alloy/internal/runtime/internal/worker",
    "github.com/grafana/alloy/internal/runtime/tracing/internal/jaegerremote/utils",
    "github.com/grafana/alloy/internal/runtime/tracing/internal/jaegerremote",
    "github.com/grafana/alloy/internal/runtime/tracing",
    "github.com/grafana/alloy/internal/util/ast",
    "github.com/grafana/alloy/internal/runtime/internal/controller",
    "github.com/grafana/alloy/internal/static/config/encoder",
    "github.com/grafana/alloy/internal/runtime",
    "github.com/grafana/alloy/internal/util/jitter",
    "github.com/grafana/alloy/internal/service/remotecfg",
    "github.com/grafana/alloy/internal/static/server",
    "github.com/grafana/alloy/internal/service/http",
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/service/livedebugging",
    "github.com/grafana/alloy/internal/component/discovery",
    "github.com/grafana/alloy/internal/component/beyla/ebpf",
    "github.com/grafana/alloy/internal/component/database_observability",
    "github.com/grafana/alloy/internal/component/database_observability/mysql/collector/parser",
    "github.com/grafana/alloy/internal/component/database_observability/mysql/collector",
    "github.com/grafana/alloy/internal/component/database_observability/mysql",
    "github.com/grafana/alloy/internal/component/database_observability/postgres/collector",
    "github.com/grafana/alloy/internal/component/database_observability/postgres",
    "github.com/grafana/alloy/internal/component/discovery/aws",
    "github.com/grafana/alloy/internal/component/discovery/azure",
    "github.com/grafana/alloy/internal/component/discovery/consul",
    "github.com/grafana/alloy/internal/component/discovery/consulagent",
    "github.com/grafana/alloy/internal/component/discovery/digitalocean",
    "github.com/grafana/alloy/internal/component/discovery/dns",
    "github.com/grafana/alloy/internal/component/discovery/docker",
    "github.com/grafana/alloy/internal/component/discovery/dockerswarm",
    "github.com/grafana/alloy/internal/component/discovery/eureka",
    "github.com/grafana/alloy/internal/component/discovery/file",
    "github.com/grafana/alloy/internal/component/discovery/gce",
    "github.com/grafana/alloy/internal/component/discovery/hetzner",
    "github.com/grafana/alloy/internal/component/discovery/http",
    "github.com/grafana/alloy/internal/component/discovery/ionos",
    "github.com/grafana/alloy/internal/component/discovery/kubelet",
    "github.com/grafana/alloy/internal/component/discovery/kubernetes",
    "github.com/grafana/alloy/internal/component/discovery/kuma",
    "github.com/grafana/alloy/internal/component/discovery/linode",
    "github.com/grafana/alloy/internal/component/discovery/marathon",
    "github.com/grafana/alloy/internal/component/discovery/nerve",
    "github.com/grafana/alloy/internal/component/discovery/nomad",
    "github.com/grafana/alloy/internal/component/discovery/openstack",
    "github.com/grafana/alloy/internal/component/discovery/ovhcloud",
    "github.com/grafana/alloy/internal/component/discovery/process",
    "github.com/grafana/alloy/internal/component/discovery/puppetdb",
    "github.com/grafana/alloy/internal/component/discovery/relabel",
    "github.com/grafana/alloy/internal/component/discovery/scaleway",
    "github.com/grafana/alloy/internal/component/discovery/serverset",
    "github.com/grafana/alloy/internal/component/discovery/triton",
    "github.com/grafana/alloy/internal/component/discovery/uyuni",
    "github.com/grafana/alloy/internal/component/faro/receiver/internal/payload",
    "github.com/grafana/alloy/internal/util/wildcard",
    "github.com/grafana/alloy/internal/component/faro/receiver",
    "github.com/grafana/alloy/internal/component/local/file",
    "github.com/grafana/alloy/internal/component/local/file_match",
    "github.com/grafana/alloy/internal/component/loki/echo",
    "github.com/grafana/alloy/internal/component/loki/enrich",
    "github.com/grafana/alloy/internal/loki/util",
    "github.com/grafana/alloy/internal/component/loki/process/metric",
    "github.com/grafana/alloy/internal/loki/logql",
    "github.com/grafana/alloy/internal/component/loki/process/stages",
    "github.com/grafana/alloy/internal/component/loki/process",
    "github.com/grafana/alloy/internal/component/loki/relabel",
    "github.com/grafana/alloy/internal/mimir/alertmanager",
    "github.com/grafana/alloy/internal/mimir/client/internal",
    "github.com/grafana/alloy/internal/mimir/client",
    "github.com/grafana/alloy/internal/component/common/kubernetes",
    "github.com/grafana/alloy/internal/loki/client/internal",
    "github.com/grafana/alloy/internal/loki/client",
    "github.com/grafana/alloy/internal/component/loki/rules/kubernetes",
    "github.com/grafana/alloy/internal/component/loki/secretfilter",
    "github.com/grafana/alloy/internal/component/common/net",
    "github.com/grafana/alloy/internal/component/loki/source",
    "github.com/grafana/alloy/internal/component/common/loki/wal/internal",
    "github.com/grafana/alloy/internal/component/common/loki/wal",
    "github.com/grafana/alloy/internal/component/common/loki/client/internal",
    "github.com/grafana/alloy/internal/component/common/loki/client",
    "github.com/grafana/alloy/internal/component/loki/source/api/internal/loghttp",
    "github.com/grafana/alloy/internal/component/loki/source/api/internal/lokipush",
    "github.com/grafana/alloy/internal/component/loki/source/api",
    "github.com/grafana/alloy/internal/component/loki/source/aws_firehose/internal",
    "github.com/grafana/alloy/internal/component/loki/source/aws_firehose",
    "github.com/grafana/alloy/internal/component/loki/source/azure_event_hubs/internal/parser",
    "github.com/grafana/alloy/internal/component/loki/source/internal/kafkatarget",
    "github.com/grafana/alloy/internal/component/loki/source/azure_event_hubs",
    "github.com/grafana/alloy/internal/component/loki/source/internal/positions",
    "github.com/grafana/alloy/internal/component/loki/source/cloudflare",
    "github.com/grafana/alloy/internal/component/loki/source/docker",
    "github.com/grafana/alloy/internal/component/loki/source/file/internal/tail/fileext",
    "github.com/grafana/alloy/internal/component/loki/source/file/internal/tail",
    "github.com/grafana/alloy/internal/component/loki/source/file",
    "github.com/grafana/alloy/internal/component/loki/source/gcplog/gcptypes",
    "github.com/grafana/alloy/internal/component/loki/source/gcplog/internal/gcplogtarget",
    "github.com/grafana/alloy/internal/component/loki/source/gcplog",
    "github.com/grafana/alloy/internal/loki/promtail/discovery/consulagent",
    "github.com/grafana/alloy/internal/loki/util/flagext",
    "github.com/grafana/alloy/internal/loki/promtail/stages",
    "github.com/grafana/alloy/internal/loki/promtail/scrapeconfig",
    "github.com/grafana/alloy/internal/component/loki/source/gelf/internal/target",
    "github.com/grafana/alloy/internal/component/loki/source/gelf",
    "github.com/grafana/alloy/internal/component/loki/source/heroku/internal/herokutarget",
    "github.com/grafana/alloy/internal/component/loki/source/heroku",
    "github.com/grafana/alloy/internal/component/loki/source/journal",
    "github.com/grafana/alloy/internal/component/loki/source/kafka",
    "github.com/grafana/alloy/internal/component/loki/source/kubernetes/kubetail",
    "github.com/grafana/alloy/internal/component/loki/source/kubernetes",
    "github.com/grafana/alloy/internal/component/loki/source/kubernetes_events",
    "github.com/grafana/alloy/internal/component/loki/source/podlogs/internal/apis/monitoring/v1alpha2",
    "github.com/grafana/alloy/internal/component/loki/source/podlogs",
    "github.com/grafana/alloy/internal/component/loki/source/syslog/config",
    "github.com/grafana/alloy/internal/component/loki/source/syslog/internal/syslogtarget/syslogparser",
    "github.com/grafana/alloy/internal/component/loki/source/syslog/internal/syslogtarget",
    "github.com/grafana/alloy/internal/component/loki/source/syslog",
    "github.com/grafana/alloy/internal/component/loki/source/windowsevent",
    "github.com/grafana/alloy/internal/component/loki/write",
    "github.com/grafana/alloy/internal/component/mimir/util",
    "github.com/grafana/alloy/internal/component/mimir/alerts/kubernetes",
    "github.com/grafana/alloy/internal/component/mimir/rules/kubernetes",
    "github.com/grafana/alloy/internal/component/otelcol/auth/basic",
    "github.com/grafana/alloy/internal/component/otelcol/auth/bearer",
    "github.com/grafana/alloy/internal/component/otelcol/auth/headers",
    "github.com/grafana/alloy/internal/component/otelcol/auth/oauth2",
    "github.com/grafana/alloy/internal/component/otelcol/auth/sigv4",
    "github.com/grafana/alloy/internal/component/otelcol/internal/fanoutconsumer",
    "github.com/grafana/alloy/internal/component/otelcol/internal/interceptconsumer",
    "github.com/grafana/alloy/internal/component/otelcol/internal/lazyconsumer",
    "github.com/grafana/alloy/internal/component/otelcol/internal/textmarshaler",
    "github.com/grafana/alloy/internal/component/otelcol/internal/livedebuggingpublisher",
    "github.com/grafana/alloy/internal/component/otelcol/connector",
    "github.com/grafana/alloy/internal/component/otelcol/connector/count",
    "github.com/grafana/alloy/internal/component/otelcol/connector/host_info",
    "github.com/grafana/alloy/internal/component/otelcol/connector/servicegraph",
    "github.com/grafana/alloy/internal/component/otelcol/connector/spanlogs",
    "github.com/grafana/alloy/internal/component/otelcol/connector/spanmetrics",
    "github.com/grafana/alloy/internal/component/otelcol/internal/views",
    "github.com/grafana/alloy/internal/component/otelcol/exporter",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/awss3",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/datadog/config",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/datadog",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/debug",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/faro",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/file",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/googlecloud/config",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/googlecloud",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/googlecloudpubsub/config",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/googlecloudpubsub",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/kafka",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/loadbalancing",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/loki/internal/convert",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/loki",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/otlp",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/otlphttp",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus/internal/convert",
    "github.com/grafana/alloy/internal/component/prometheus/appenders",
    "github.com/grafana/alloy/internal/service/labelstore",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/splunkhec/config",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/splunkhec",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/syslog",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/server/grpc",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/server/http",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/filesource",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/remotesource",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/strategy_store",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling",
    "github.com/grafana/alloy/internal/component/otelcol/processor",
    "github.com/grafana/alloy/internal/component/otelcol/processor/attributes",
    "github.com/grafana/alloy/internal/component/otelcol/processor/batch",
    "github.com/grafana/alloy/internal/component/otelcol/processor/cumulativetodelta",
    "github.com/grafana/alloy/internal/component/otelcol/processor/deltatocumulative",
    "github.com/grafana/alloy/internal/static/traces/promsdprocessor/consumer",
    "github.com/grafana/alloy/internal/component/otelcol/processor/discovery",
    "github.com/grafana/alloy/internal/component/otelcol/processor/filter",
    "github.com/grafana/alloy/internal/component/otelcol/processor/groupbyattrs",
    "github.com/grafana/alloy/internal/component/otelcol/processor/interval",
    "github.com/grafana/alloy/internal/component/otelcol/processor/k8sattributes",
    "github.com/grafana/alloy/internal/component/otelcol/processor/memorylimiter",
    "github.com/grafana/alloy/internal/component/otelcol/processor/probabilistic_sampler",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/resource_attribute_config",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/akamai",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/aws/ec2",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/aws/ecs",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/aws/eks",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/aws/elasticbeanstalk",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/aws/lambda",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/azure",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/azure/aks",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/consul",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/digitalocean",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/docker",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/dynatrace",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/gcp",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/heroku",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/hetzner",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/k8snode",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/kubeadm",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/openshift",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/openstacknova",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/oraclecloud",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/scaleway",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/system",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/upcloud",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection/internal/vultr",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection",
    "github.com/grafana/alloy/internal/component/otelcol/processor/span",
    "github.com/grafana/alloy/internal/component/otelcol/processor/tail_sampling",
    "github.com/grafana/alloy/internal/component/otelcol/processor/transform",
    "github.com/grafana/alloy/internal/component/otelcol/receiver",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/awscloudwatch",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/awsecscontainermetrics",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/awss3",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/cloudflare",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/datadog",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/faro",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/file_stats",
    "github.com/grafana/alloy/internal/component/otelcol/internal/textutils",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/filelog",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/fluentforward",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/googlecloudpubsub",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/influxdb",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/jaeger",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/kafka",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/loki",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/opencensus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/otlp",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus/internal",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/solace",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/splunkhec",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/syslog",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/tcplog",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/vcenter",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/zipkin",
    "github.com/grafana/alloy/internal/component/otelcol/storage/file",
    "github.com/grafana/alloy/internal/component/prometheus/echo",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/static/integrations/config",
    "github.com/grafana/alloy/internal/static/metrics/cluster/client",
    "github.com/grafana/alloy/internal/util/log",
    "github.com/grafana/alloy/internal/static/metrics/cluster",
    "github.com/grafana/alloy/internal/static/metrics/instance",
    "github.com/grafana/alloy/internal/static/metrics",
    "github.com/grafana/alloy/internal/static/integrations",
    "github.com/grafana/alloy/internal/component/prometheus/exporter",
    "github.com/grafana/alloy/internal/static/integrations/apache_http",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/apache",
    "github.com/grafana/alloy/internal/static/integrations/v2/autoscrape",
    "github.com/grafana/alloy/internal/static/integrations/v2/common",
    "github.com/grafana/alloy/internal/util/otelfeaturegatefix",
    "github.com/grafana/alloy/internal/static/integrations/v2",
    "github.com/grafana/alloy/internal/static/integrations/v2/metricsutils",
    "github.com/grafana/alloy/internal/static/integrations/azure_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/azure",
    "github.com/grafana/alloy/internal/static/integrations/blackbox_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/blackbox",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/common",
    "github.com/grafana/alloy/internal/static/integrations/cadvisor",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/cadvisor",
    "github.com/grafana/alloy/internal/static/integrations/catchpoint_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/catchpoint",
    "github.com/grafana/alloy/internal/static/integrations/cloudwatch_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/cloudwatch",
    "github.com/grafana/alloy/internal/static/integrations/consul_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/consul",
    "github.com/grafana/alloy/internal/static/integrations/dnsmasq_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/dnsmasq",
    "github.com/grafana/alloy/internal/static/integrations/elasticsearch_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/elasticsearch",
    "github.com/grafana/alloy/internal/static/integrations/gcp_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/gcp",
    "github.com/grafana/alloy/internal/static/integrations/github_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/github",
    "github.com/grafana/alloy/internal/static/integrations/kafka_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/kafka",
    "github.com/grafana/alloy/internal/static/integrations/memcached_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/memcached",
    "github.com/grafana/alloy/internal/static/integrations/mongodb_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/mongodb",
    "github.com/grafana/alloy/internal/static/integrations/mssql",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/mssql",
    "github.com/grafana/alloy/internal/static/integrations/mysqld_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/mysql",
    "github.com/grafana/alloy/internal/static/integrations/oracledb_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/oracledb",
    "github.com/grafana/alloy/internal/static/integrations/postgres_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/postgres",
    "github.com/grafana/alloy/internal/static/integrations/process_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/process",
    "github.com/grafana/alloy/internal/static/integrations/redis_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/redis",
    "github.com/grafana/alloy/internal/static/integrations/agent",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/self",
    "github.com/grafana/alloy/internal/static/integrations/snmp_exporter/common",
    "github.com/grafana/alloy/internal/static/integrations/snmp_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/snmp",
    "github.com/grafana/alloy/internal/static/integrations/snowflake_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/snowflake",
    "github.com/grafana/alloy/internal/static/integrations/squid_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/squid",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/static",
    "github.com/grafana/alloy/internal/static/integrations/statsd_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/statsd",
    "github.com/grafana/alloy/internal/static/integrations/node_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/unix",
    "github.com/grafana/alloy/internal/static/integrations/windows_exporter",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/windows",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/static/agentctl/waltools",
    "github.com/grafana/alloy/internal/static/metrics/wal",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/prometheus/write/queue",
    "github.com/grafana/alloy/internal/component/pyroscope/ebpf/reporter",
    "github.com/grafana/alloy/internal/component/pyroscope/ebpf",
    "github.com/grafana/alloy/internal/component/pyroscope/enrich",
    "github.com/grafana/alloy/internal/component/pyroscope/java/asprof",
    "github.com/grafana/alloy/internal/component/pyroscope/java",
    "github.com/grafana/alloy/internal/component/pyroscope/util",
    "github.com/grafana/alloy/internal/component/pyroscope/util/metrics",
    "github.com/grafana/alloy/internal/component/pyroscope/write",
    "github.com/grafana/alloy/internal/component/pyroscope/receive_http",
    "github.com/grafana/alloy/internal/component/pyroscope/relabel",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape/internal/pproflite",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape/internal/fastdelta",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/component/pyroscope/util/glue",
    "github.com/grafana/alloy/internal/component/pyroscope/write/glue",
    "github.com/grafana/alloy/internal/component/remote/kubernetes",
    "github.com/grafana/alloy/internal/component/remote/kubernetes/configmap",
    "github.com/grafana/alloy/internal/component/remote/kubernetes/secret",
    "github.com/grafana/alloy/internal/component/remote/s3",
    "github.com/grafana/alloy/internal/component/remote/vault",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/converter/diag",
    "github.com/grafana/alloy/internal/converter/internal/common",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert/envprovider",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/build",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/loki/promtail/client",
    "github.com/grafana/alloy/internal/loki/promtail/file",
    "github.com/grafana/alloy/internal/loki/promtail/metric",
    "github.com/grafana/alloy/internal/loki/promtail/positions",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/loki/promtail/limit",
    "github.com/grafana/alloy/internal/loki/promtail/server",
    "github.com/grafana/alloy/internal/loki/promtail/tracing",
    "github.com/grafana/alloy/internal/loki/promtail/wal",
    "github.com/grafana/alloy/internal/loki/promtail/config",
    "github.com/grafana/alloy/internal/loki/util/cfg",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/static/config/features",
    "github.com/grafana/alloy/internal/static/logs",
    "github.com/grafana/alloy/internal/static/traces/automaticloggingprocessor",
    "github.com/grafana/alloy/internal/static/traces/noopreceiver",
    "github.com/grafana/alloy/internal/static/traces/promsdprocessor",
    "github.com/grafana/alloy/internal/static/traces/pushreceiver",
    "github.com/grafana/alloy/internal/static/traces/remotewriteexporter",
    "github.com/grafana/alloy/internal/static/traces/servicegraphprocessor",
    "github.com/grafana/alloy/internal/static/traces/spanmetricsprocessor/internal/cache",
    "github.com/grafana/alloy/internal/static/traces/spanmetricsprocessor/internal/metadata",
    "github.com/grafana/alloy/internal/static/traces/spanmetricsprocessor",
    "github.com/grafana/alloy/internal/static/traces",
    "github.com/grafana/alloy/internal/static/config",
    "github.com/grafana/alloy/internal/static/integrations/v2/agent",
    "github.com/grafana/alloy/internal/static/integrations/v2/apache_http",
    "github.com/grafana/alloy/internal/static/integrations/v2/app_agent_receiver",
    "github.com/grafana/alloy/internal/static/integrations/v2/blackbox_exporter",
    "github.com/grafana/alloy/internal/static/integrations/v2/eventhandler",
    "github.com/grafana/alloy/internal/static/integrations/v2/snmp_exporter",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/static/integrations/v2/vmware_exporter",
    "github.com/grafana/alloy/internal/static/integrations/vmware_exporter",
    "github.com/grafana/alloy/internal/static/integrations/install",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/service/otel",
    "github.com/grafana/alloy/internal/web/api",
    "github.com/grafana/alloy/internal/web/ui",
    "github.com/grafana/alloy/internal/service/ui",
    "github.com/grafana/alloy/internal/static/config/instrumentation",
    "github.com/grafana/alloy/internal/usagestats",
    "github.com/grafana/alloy/internal/util/windowspriority",
    "github.com/grafana/alloy/internal/validator",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy/internal/winmanifest",
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/docs/developer/key-deps-update/tools/dependency-updates/latest-release",
    "github.com/grafana/alloy/docs/developer/key-deps-update/tools/tool-example-template",
    "github.com/grafana/alloy/internal/cmd/alloy-service",
    "github.com/grafana/alloy/internal/cmd/alloylint/internal/findcomponents",
    "github.com/grafana/alloy/internal/cmd/alloylint/internal/syntaxtags",
    "github.com/grafana/alloy/internal/cmd/alloylint",
    "github.com/grafana/alloy/internal/cmd/integration-tests/common",
    "github.com/grafana/alloy/internal/cmd/integration-tests",
    "github.com/grafana/alloy/internal/cmd/integration-tests/configs/kafka",
    "github.com/grafana/alloy/internal/cmd/integration-tests/configs/otel-gen",
    "github.com/grafana/alloy/internal/cmd/integration-tests/configs/prom-gen",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/beyla",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/blackbox",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/kafka",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/loki-enrich",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/otlp-metadata",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/otlp-metrics",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/prom-enrich",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/prom-metadata",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/read-log-file",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/redis",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/scrape-prom-metrics",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/snmp",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/static",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/unix",
    "github.com/grafana/alloy/internal/cmd/integration-tests-k8s/tests/mimir-alerts-kubernetes",
    "github.com/grafana/alloy/internal/cmd/integration-tests-k8s/util",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/loki/source/journal/internal/target",
    "github.com/grafana/alloy/internal/component/loki/source/windowsevent/win_eventlog",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/mocks",
    "github.com/grafana/alloy/internal/component/otelcol/internal/fakeconsumer",
    "github.com/grafana/alloy/internal/component/otelcol/internal/testutils",
    "github.com/grafana/alloy/internal/runtime/internal/testcomponents/module",
    "github.com/grafana/alloy/internal/runtime/componenttest",
    "github.com/grafana/alloy/internal/component/otelcol/processor/processortest",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/pyroscope/java/integration",
    "github.com/grafana/alloy/internal/component/pyroscope/testutil",
    "github.com/grafana/alloy/internal/component/pyroscope/util/internal/cmd/playground",
    "github.com/grafana/alloy/internal/component/pyroscope/util/test",
    "github.com/grafana/alloy/internal/component/pyroscope/util/test/container",
    "github.com/grafana/alloy/internal/component/pyroscope/util/testlog",
    "github.com/grafana/alloy/internal/converter/internal/test_common",
    "github.com/grafana/alloy/internal/runtime/internal/testcomponents",
    "github.com/grafana/alloy/internal/runtime/internal/testcomponents/module/file",
    "github.com/grafana/alloy/internal/runtime/internal/testcomponents/module/git",
    "github.com/grafana/alloy/internal/runtime/internal/testcomponents/module/http",
    "github.com/grafana/alloy/internal/runtime/internal/testcomponents/module/string",
    "github.com/grafana/alloy/internal/runtime/internal/testcomponents/targets",
    "github.com/grafana/alloy/internal/runtime/internal/testservices",
    "github.com/grafana/alloy/internal/service/cluster/internal/testcomponent",
    "github.com/grafana/alloy/internal/static/integrations/cloudwatch_exporter/docs",
    "github.com/grafana/alloy/internal/static/traces/contextkeys",
    "github.com/grafana/alloy/internal/static/traces/spanmetricsprocessor/mocks",
    "github.com/grafana/alloy/internal/static/traces/traceutils",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/util/assertmetrics",
    "github.com/grafana/alloy/internal/util/syncbuffer",
    "github.com/grafana/alloy/internal/util/testappender/internal/dtobuilder",
    "github.com/grafana/alloy/internal/util/testappender",
    "github.com/grafana/alloy/internal/util/testlivedebugging",
    "github.com/grafana/alloy/internal/util/testtarget",
    "builtin",
    "github.com/grafana/alloy [github.com/grafana/alloy.test]",
    "github.com/grafana/alloy.test",
    "github.com/grafana/alloy/docs/developer/key-deps-update/tools/dependency-updates/latest-release [github.com/grafana/alloy/docs/developer/key-deps-update/tools/dependency-updates/latest-release.test]",
    "github.com/grafana/alloy/docs/developer/key-deps-update/tools/dependency-updates/latest-release.test",
    "github.com/grafana/alloy/internal/alloycli [github.com/grafana/alloy/internal/alloycli.test]",
    "github.com/grafana/alloy/internal/alloycli.test",
    "github.com/grafana/alloy/internal/alloyseed [github.com/grafana/alloy/internal/alloyseed.test]",
    "github.com/grafana/alloy/internal/alloyseed.test",
    "github.com/grafana/alloy/internal/build [github.com/grafana/alloy/internal/build.test]",
    "github.com/grafana/alloy/internal/build.test",
    "github.com/grafana/alloy/internal/cmd/alloy-service [github.com/grafana/alloy/internal/cmd/alloy-service.test]",
    "github.com/grafana/alloy/internal/cmd/alloy-service.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/beyla [github.com/grafana/alloy/internal/cmd/integration-tests/tests/beyla.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/beyla.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/blackbox [github.com/grafana/alloy/internal/cmd/integration-tests/tests/blackbox.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/blackbox.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/kafka [github.com/grafana/alloy/internal/cmd/integration-tests/tests/kafka.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/kafka.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/loki-enrich [github.com/grafana/alloy/internal/cmd/integration-tests/tests/loki-enrich.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/loki-enrich.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/otlp-metadata [github.com/grafana/alloy/internal/cmd/integration-tests/tests/otlp-metadata.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/otlp-metadata.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/otlp-metrics [github.com/grafana/alloy/internal/cmd/integration-tests/tests/otlp-metrics.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/otlp-metrics.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/prom-enrich [github.com/grafana/alloy/internal/cmd/integration-tests/tests/prom-enrich.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/prom-enrich.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/prom-metadata [github.com/grafana/alloy/internal/cmd/integration-tests/tests/prom-metadata.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/prom-metadata.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/read-log-file [github.com/grafana/alloy/internal/cmd/integration-tests/tests/read-log-file.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/read-log-file.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/redis [github.com/grafana/alloy/internal/cmd/integration-tests/tests/redis.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/redis.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/scrape-prom-metrics [github.com/grafana/alloy/internal/cmd/integration-tests/tests/scrape-prom-metrics.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/scrape-prom-metrics.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/snmp [github.com/grafana/alloy/internal/cmd/integration-tests/tests/snmp.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/snmp.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/static [github.com/grafana/alloy/internal/cmd/integration-tests/tests/static.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/static.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/unix [github.com/grafana/alloy/internal/cmd/integration-tests/tests/unix.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests/tests/unix.test",
    "github.com/grafana/alloy/internal/cmd/integration-tests-k8s/tests/mimir-alerts-kubernetes [github.com/grafana/alloy/internal/cmd/integration-tests-k8s/tests/mimir-alerts-kubernetes.test]",
    "github.com/grafana/alloy/internal/cmd/integration-tests-k8s/tests/mimir-alerts-kubernetes.test",
    "github.com/grafana/alloy/internal/component [github.com/grafana/alloy/internal/component.test]",
    "github.com/grafana/alloy/internal/component_test [github.com/grafana/alloy/internal/component.test]",
    "github.com/grafana/alloy/internal/component.test",
    "github.com/grafana/alloy/internal/component/all [github.com/grafana/alloy/internal/component/all.test]",
    "github.com/grafana/alloy/internal/component/all.test",
    "github.com/grafana/alloy/internal/component/beyla/ebpf [github.com/grafana/alloy/internal/component/beyla/ebpf.test]",
    "github.com/grafana/alloy/internal/component/beyla/ebpf.test",
    "github.com/grafana/alloy/internal/component/common/config [github.com/grafana/alloy/internal/component/common/config.test]",
    "github.com/grafana/alloy/internal/component/common/config.test",
    "github.com/grafana/alloy/internal/component/common/kubernetes [github.com/grafana/alloy/internal/component/common/kubernetes.test]",
    "github.com/grafana/alloy/internal/component/common/kubernetes.test",
    "github.com/grafana/alloy/internal/component/common/loki [github.com/grafana/alloy/internal/component/common/loki.test]",
    "github.com/grafana/alloy/internal/component/common/loki.test",
    "github.com/grafana/alloy/internal/component/common/loki/client [github.com/grafana/alloy/internal/component/common/loki/client.test]",
    "github.com/grafana/alloy/internal/component/common/loki/client.test",
    "github.com/grafana/alloy/internal/component/common/loki/client/internal [github.com/grafana/alloy/internal/component/common/loki/client/internal.test]",
    "github.com/grafana/alloy/internal/component/common/loki/client/internal.test",
    "github.com/grafana/alloy/internal/component/common/loki/wal [github.com/grafana/alloy/internal/component/common/loki/wal.test]",
    "github.com/grafana/alloy/internal/component/common/loki/wal.test",
    "github.com/grafana/alloy/internal/component/common/net [github.com/grafana/alloy/internal/component/common/net.test]",
    "github.com/grafana/alloy/internal/component/common/net.test",
    "github.com/grafana/alloy/internal/component/common/relabel [github.com/grafana/alloy/internal/component/common/relabel.test]",
    "github.com/grafana/alloy/internal/component/common/relabel.test",
    "github.com/grafana/alloy/internal/component/database_observability [github.com/grafana/alloy/internal/component/database_observability.test]",
    "github.com/grafana/alloy/internal/component/database_observability.test",
    "github.com/grafana/alloy/internal/component/database_observability/mysql [github.com/grafana/alloy/internal/component/database_observability/mysql.test]",
    "github.com/grafana/alloy/internal/component/database_observability/mysql.test",
    "github.com/grafana/alloy/internal/component/database_observability/mysql/collector [github.com/grafana/alloy/internal/component/database_observability/mysql/collector.test]",
    "github.com/grafana/alloy/internal/component/database_observability/mysql/collector.test",
    "github.com/grafana/alloy/internal/component/database_observability/mysql/collector/parser_test [github.com/grafana/alloy/internal/component/database_observability/mysql/collector/parser.test]",
    "github.com/grafana/alloy/internal/component/database_observability/mysql/collector/parser.test",
    "github.com/grafana/alloy/internal/component/database_observability/postgres [github.com/grafana/alloy/internal/component/database_observability/postgres.test]",
    "github.com/grafana/alloy/internal/component/database_observability/postgres.test",
    "github.com/grafana/alloy/internal/component/database_observability/postgres/collector [github.com/grafana/alloy/internal/component/database_observability/postgres/collector.test]",
    "github.com/grafana/alloy/internal/component/database_observability/postgres/collector.test",
    "github.com/grafana/alloy/internal/component/discovery [github.com/grafana/alloy/internal/component/discovery.test]",
    "github.com/grafana/alloy/internal/component/discovery.test",
    "github.com/grafana/alloy/internal/component/discovery/aws [github.com/grafana/alloy/internal/component/discovery/aws.test]",
    "github.com/grafana/alloy/internal/component/discovery/aws.test",
    "github.com/grafana/alloy/internal/component/discovery/azure [github.com/grafana/alloy/internal/component/discovery/azure.test]",
    "github.com/grafana/alloy/internal/component/discovery/azure.test",
    "github.com/grafana/alloy/internal/component/discovery/consul [github.com/grafana/alloy/internal/component/discovery/consul.test]",
    "github.com/grafana/alloy/internal/component/discovery/consul.test",
    "github.com/grafana/alloy/internal/component/discovery/consulagent [github.com/grafana/alloy/internal/component/discovery/consulagent.test]",
    "github.com/grafana/alloy/internal/component/discovery/consulagent.test",
    "github.com/grafana/alloy/internal/component/discovery/digitalocean [github.com/grafana/alloy/internal/component/discovery/digitalocean.test]",
    "github.com/grafana/alloy/internal/component/discovery/digitalocean.test",
    "github.com/grafana/alloy/internal/component/discovery/dns [github.com/grafana/alloy/internal/component/discovery/dns.test]",
    "github.com/grafana/alloy/internal/component/discovery/dns.test",
    "github.com/grafana/alloy/internal/component/discovery/docker [github.com/grafana/alloy/internal/component/discovery/docker.test]",
    "github.com/grafana/alloy/internal/component/discovery/docker.test",
    "github.com/grafana/alloy/internal/component/discovery/dockerswarm [github.com/grafana/alloy/internal/component/discovery/dockerswarm.test]",
    "github.com/grafana/alloy/internal/component/discovery/dockerswarm.test",
    "github.com/grafana/alloy/internal/component/discovery/eureka [github.com/grafana/alloy/internal/component/discovery/eureka.test]",
    "github.com/grafana/alloy/internal/component/discovery/eureka.test",
    "github.com/grafana/alloy/internal/component/discovery/file [github.com/grafana/alloy/internal/component/discovery/file.test]",
    "github.com/grafana/alloy/internal/component/discovery/file.test",
    "github.com/grafana/alloy/internal/component/discovery/gce [github.com/grafana/alloy/internal/component/discovery/gce.test]",
    "github.com/grafana/alloy/internal/component/discovery/gce.test",
    "github.com/grafana/alloy/internal/component/discovery/hetzner [github.com/grafana/alloy/internal/component/discovery/hetzner.test]",
    "github.com/grafana/alloy/internal/component/discovery/hetzner.test",
    "github.com/grafana/alloy/internal/component/discovery/http [github.com/grafana/alloy/internal/component/discovery/http.test]",
    "github.com/grafana/alloy/internal/component/discovery/http.test",
    "github.com/grafana/alloy/internal/component/discovery/ionos [github.com/grafana/alloy/internal/component/discovery/ionos.test]",
    "github.com/grafana/alloy/internal/component/discovery/ionos.test",
    "github.com/grafana/alloy/internal/component/discovery/kubelet [github.com/grafana/alloy/internal/component/discovery/kubelet.test]",
    "github.com/grafana/alloy/internal/component/discovery/kubelet.test",
    "github.com/grafana/alloy/internal/component/discovery/kubernetes [github.com/grafana/alloy/internal/component/discovery/kubernetes.test]",
    "github.com/grafana/alloy/internal/component/discovery/kubernetes.test",
    "github.com/grafana/alloy/internal/component/discovery/kuma [github.com/grafana/alloy/internal/component/discovery/kuma.test]",
    "github.com/grafana/alloy/internal/component/discovery/kuma.test",
    "github.com/grafana/alloy/internal/component/discovery/linode [github.com/grafana/alloy/internal/component/discovery/linode.test]",
    "github.com/grafana/alloy/internal/component/discovery/linode.test",
    "github.com/grafana/alloy/internal/component/discovery/marathon [github.com/grafana/alloy/internal/component/discovery/marathon.test]",
    "github.com/grafana/alloy/internal/component/discovery/marathon.test",
    "github.com/grafana/alloy/internal/component/discovery/nerve [github.com/grafana/alloy/internal/component/discovery/nerve.test]",
    "github.com/grafana/alloy/internal/component/discovery/nerve.test",
    "github.com/grafana/alloy/internal/component/discovery/nomad [github.com/grafana/alloy/internal/component/discovery/nomad.test]",
    "github.com/grafana/alloy/internal/component/discovery/nomad.test",
    "github.com/grafana/alloy/internal/component/discovery/openstack [github.com/grafana/alloy/internal/component/discovery/openstack.test]",
    "github.com/grafana/alloy/internal/component/discovery/openstack.test",
    "github.com/grafana/alloy/internal/component/discovery/ovhcloud_test [github.com/grafana/alloy/internal/component/discovery/ovhcloud.test]",
    "github.com/grafana/alloy/internal/component/discovery/ovhcloud.test",
    "github.com/grafana/alloy/internal/component/discovery/process [github.com/grafana/alloy/internal/component/discovery/process.test]",
    "github.com/grafana/alloy/internal/component/discovery/process.test",
    "github.com/grafana/alloy/internal/component/discovery/puppetdb [github.com/grafana/alloy/internal/component/discovery/puppetdb.test]",
    "github.com/grafana/alloy/internal/component/discovery/puppetdb.test",
    "github.com/grafana/alloy/internal/component/discovery/relabel_test [github.com/grafana/alloy/internal/component/discovery/relabel.test]",
    "github.com/grafana/alloy/internal/component/discovery/relabel.test",
    "github.com/grafana/alloy/internal/component/discovery/scaleway [github.com/grafana/alloy/internal/component/discovery/scaleway.test]",
    "github.com/grafana/alloy/internal/component/discovery/scaleway.test",
    "github.com/grafana/alloy/internal/component/discovery/serverset [github.com/grafana/alloy/internal/component/discovery/serverset.test]",
    "github.com/grafana/alloy/internal/component/discovery/serverset.test",
    "github.com/grafana/alloy/internal/component/discovery/triton [github.com/grafana/alloy/internal/component/discovery/triton.test]",
    "github.com/grafana/alloy/internal/component/discovery/triton.test",
    "github.com/grafana/alloy/internal/component/discovery/uyuni [github.com/grafana/alloy/internal/component/discovery/uyuni.test]",
    "github.com/grafana/alloy/internal/component/discovery/uyuni.test",
    "github.com/grafana/alloy/internal/component/faro/receiver [github.com/grafana/alloy/internal/component/faro/receiver.test]",
    "github.com/grafana/alloy/internal/component/faro/receiver.test",
    "github.com/grafana/alloy/internal/component/faro/receiver/internal/payload [github.com/grafana/alloy/internal/component/faro/receiver/internal/payload.test]",
    "github.com/grafana/alloy/internal/component/faro/receiver/internal/payload.test",
    "github.com/grafana/alloy/internal/component/local/file_test [github.com/grafana/alloy/internal/component/local/file.test]",
    "github.com/grafana/alloy/internal/component/local/file.test",
    "github.com/grafana/alloy/internal/component/local/file_match [github.com/grafana/alloy/internal/component/local/file_match.test]",
    "github.com/grafana/alloy/internal/component/local/file_match.test",
    "github.com/grafana/alloy/internal/component/loki/enrich [github.com/grafana/alloy/internal/component/loki/enrich.test]",
    "github.com/grafana/alloy/internal/component/loki/enrich.test",
    "github.com/grafana/alloy/internal/component/loki/process [github.com/grafana/alloy/internal/component/loki/process.test]",
    "github.com/grafana/alloy/internal/component/loki/process.test",
    "github.com/grafana/alloy/internal/component/loki/process/metric [github.com/grafana/alloy/internal/component/loki/process/metric.test]",
    "github.com/grafana/alloy/internal/component/loki/process/metric.test",
    "github.com/grafana/alloy/internal/component/loki/process/stages [github.com/grafana/alloy/internal/component/loki/process/stages.test]",
    "github.com/grafana/alloy/internal/component/loki/process/stages.test",
    "github.com/grafana/alloy/internal/component/loki/relabel [github.com/grafana/alloy/internal/component/loki/relabel.test]",
    "github.com/grafana/alloy/internal/component/loki/relabel.test",
    "github.com/grafana/alloy/internal/component/loki/rules/kubernetes [github.com/grafana/alloy/internal/component/loki/rules/kubernetes.test]",
    "github.com/grafana/alloy/internal/component/loki/rules/kubernetes.test",
    "github.com/grafana/alloy/internal/component/loki/secretfilter [github.com/grafana/alloy/internal/component/loki/secretfilter.test]",
    "github.com/grafana/alloy/internal/component/loki/secretfilter.test",
    "github.com/grafana/alloy/internal/component/loki/source [github.com/grafana/alloy/internal/component/loki/source.test]",
    "github.com/grafana/alloy/internal/component/loki/source.test",
    "github.com/grafana/alloy/internal/component/loki/source/api [github.com/grafana/alloy/internal/component/loki/source/api.test]",
    "github.com/grafana/alloy/internal/component/loki/source/api.test",
    "github.com/grafana/alloy/internal/component/loki/source/api/internal/lokipush [github.com/grafana/alloy/internal/component/loki/source/api/internal/lokipush.test]",
    "github.com/grafana/alloy/internal/component/loki/source/api/internal/lokipush.test",
    "github.com/grafana/alloy/internal/component/loki/source/aws_firehose [github.com/grafana/alloy/internal/component/loki/source/aws_firehose.test]",
    "github.com/grafana/alloy/internal/component/loki/source/aws_firehose.test",
    "github.com/grafana/alloy/internal/component/loki/source/aws_firehose/internal [github.com/grafana/alloy/internal/component/loki/source/aws_firehose/internal.test]",
    "github.com/grafana/alloy/internal/component/loki/source/aws_firehose/internal.test",
    "github.com/grafana/alloy/internal/component/loki/source/azure_event_hubs [github.com/grafana/alloy/internal/component/loki/source/azure_event_hubs.test]",
    "github.com/grafana/alloy/internal/component/loki/source/azure_event_hubs.test",
    "github.com/grafana/alloy/internal/component/loki/source/azure_event_hubs/internal/parser [github.com/grafana/alloy/internal/component/loki/source/azure_event_hubs/internal/parser.test]",
    "github.com/grafana/alloy/internal/component/loki/source/azure_event_hubs/internal/parser.test",
    "github.com/grafana/alloy/internal/component/loki/source/cloudflare [github.com/grafana/alloy/internal/component/loki/source/cloudflare.test]",
    "github.com/grafana/alloy/internal/component/loki/source/cloudflare.test",
    "github.com/grafana/alloy/internal/component/loki/source/docker [github.com/grafana/alloy/internal/component/loki/source/docker.test]",
    "github.com/grafana/alloy/internal/component/loki/source/docker.test",
    "github.com/grafana/alloy/internal/component/loki/source/file [github.com/grafana/alloy/internal/component/loki/source/file.test]",
    "github.com/grafana/alloy/internal/component/loki/source/file.test",
    "github.com/grafana/alloy/internal/component/loki/source/file/internal/tail [github.com/grafana/alloy/internal/component/loki/source/file/internal/tail.test]",
    "github.com/grafana/alloy/internal/component/loki/source/file/internal/tail.test",
    "github.com/grafana/alloy/internal/component/loki/source/gcplog [github.com/grafana/alloy/internal/component/loki/source/gcplog.test]",
    "github.com/grafana/alloy/internal/component/loki/source/gcplog.test",
    "github.com/grafana/alloy/internal/component/loki/source/gcplog/internal/gcplogtarget [github.com/grafana/alloy/internal/component/loki/source/gcplog/internal/gcplogtarget.test]",
    "github.com/grafana/alloy/internal/component/loki/source/gcplog/internal/gcplogtarget.test",
    "github.com/grafana/alloy/internal/component/loki/source/gelf [github.com/grafana/alloy/internal/component/loki/source/gelf.test]",
    "github.com/grafana/alloy/internal/component/loki/source/gelf.test",
    "github.com/grafana/alloy/internal/component/loki/source/heroku [github.com/grafana/alloy/internal/component/loki/source/heroku.test]",
    "github.com/grafana/alloy/internal/component/loki/source/heroku.test",
    "github.com/grafana/alloy/internal/component/loki/source/heroku/internal/herokutarget [github.com/grafana/alloy/internal/component/loki/source/heroku/internal/herokutarget.test]",
    "github.com/grafana/alloy/internal/component/loki/source/heroku/internal/herokutarget.test",
    "github.com/grafana/alloy/internal/component/loki/source/internal/kafkatarget [github.com/grafana/alloy/internal/component/loki/source/internal/kafkatarget.test]",
    "github.com/grafana/alloy/internal/component/loki/source/internal/kafkatarget.test",
    "github.com/grafana/alloy/internal/component/loki/source/internal/positions [github.com/grafana/alloy/internal/component/loki/source/internal/positions.test]",
    "github.com/grafana/alloy/internal/component/loki/source/internal/positions.test",
    "github.com/grafana/alloy/internal/component/loki/source/kafka [github.com/grafana/alloy/internal/component/loki/source/kafka.test]",
    "github.com/grafana/alloy/internal/component/loki/source/kafka.test",
    "github.com/grafana/alloy/internal/component/loki/source/kubernetes [github.com/grafana/alloy/internal/component/loki/source/kubernetes.test]",
    "github.com/grafana/alloy/internal/component/loki/source/kubernetes.test",
    "github.com/grafana/alloy/internal/component/loki/source/kubernetes/kubetail [github.com/grafana/alloy/internal/component/loki/source/kubernetes/kubetail.test]",
    "github.com/grafana/alloy/internal/component/loki/source/kubernetes/kubetail.test",
    "github.com/grafana/alloy/internal/component/loki/source/podlogs [github.com/grafana/alloy/internal/component/loki/source/podlogs.test]",
    "github.com/grafana/alloy/internal/component/loki/source/podlogs.test",
    "github.com/grafana/alloy/internal/component/loki/source/syslog [github.com/grafana/alloy/internal/component/loki/source/syslog.test]",
    "github.com/grafana/alloy/internal/component/loki/source/syslog.test",
    "github.com/grafana/alloy/internal/component/loki/source/syslog/internal/syslogtarget [github.com/grafana/alloy/internal/component/loki/source/syslog/internal/syslogtarget.test]",
    "github.com/grafana/alloy/internal/component/loki/source/syslog/internal/syslogtarget.test",
    "github.com/grafana/alloy/internal/component/loki/source/syslog/internal/syslogtarget/syslogparser [github.com/grafana/alloy/internal/component/loki/source/syslog/internal/syslogtarget/syslogparser.test]",
    "github.com/grafana/alloy/internal/component/loki/source/syslog/internal/syslogtarget/syslogparser_test [github.com/grafana/alloy/internal/component/loki/source/syslog/internal/syslogtarget/syslogparser.test]",
    "github.com/grafana/alloy/internal/component/loki/source/syslog/internal/syslogtarget/syslogparser.test",
    "github.com/grafana/alloy/internal/component/loki/write [github.com/grafana/alloy/internal/component/loki/write.test]",
    "github.com/grafana/alloy/internal/component/loki/write.test",
    "github.com/grafana/alloy/internal/component/metadata [github.com/grafana/alloy/internal/component/metadata.test]",
    "github.com/grafana/alloy/internal/component/metadata.test",
    "github.com/grafana/alloy/internal/component/mimir/alerts/kubernetes [github.com/grafana/alloy/internal/component/mimir/alerts/kubernetes.test]",
    "github.com/grafana/alloy/internal/component/mimir/alerts/kubernetes.test",
    "github.com/grafana/alloy/internal/component/mimir/rules/kubernetes [github.com/grafana/alloy/internal/component/mimir/rules/kubernetes.test]",
    "github.com/grafana/alloy/internal/component/mimir/rules/kubernetes.test",
    "github.com/grafana/alloy/internal/component/otelcol_test [github.com/grafana/alloy/internal/component/otelcol.test]",
    "github.com/grafana/alloy/internal/component/otelcol.test",
    "github.com/grafana/alloy/internal/component/otelcol/auth_test [github.com/grafana/alloy/internal/component/otelcol/auth.test]",
    "github.com/grafana/alloy/internal/component/otelcol/auth.test",
    "github.com/grafana/alloy/internal/component/otelcol/auth/basic_test [github.com/grafana/alloy/internal/component/otelcol/auth/basic.test]",
    "github.com/grafana/alloy/internal/component/otelcol/auth/basic.test",
    "github.com/grafana/alloy/internal/component/otelcol/auth/bearer_test [github.com/grafana/alloy/internal/component/otelcol/auth/bearer.test]",
    "github.com/grafana/alloy/internal/component/otelcol/auth/bearer.test",
    "github.com/grafana/alloy/internal/component/otelcol/auth/headers_test [github.com/grafana/alloy/internal/component/otelcol/auth/headers.test]",
    "github.com/grafana/alloy/internal/component/otelcol/auth/headers.test",
    "github.com/grafana/alloy/internal/component/otelcol/auth/oauth2_test [github.com/grafana/alloy/internal/component/otelcol/auth/oauth2.test]",
    "github.com/grafana/alloy/internal/component/otelcol/auth/oauth2.test",
    "github.com/grafana/alloy/internal/component/otelcol/auth/sigv4_test [github.com/grafana/alloy/internal/component/otelcol/auth/sigv4.test]",
    "github.com/grafana/alloy/internal/component/otelcol/auth/sigv4.test",
    "github.com/grafana/alloy/internal/component/otelcol/connector/count_test [github.com/grafana/alloy/internal/component/otelcol/connector/count.test]",
    "github.com/grafana/alloy/internal/component/otelcol/connector/count.test",
    "github.com/grafana/alloy/internal/component/otelcol/connector/host_info [github.com/grafana/alloy/internal/component/otelcol/connector/host_info.test]",
    "github.com/grafana/alloy/internal/component/otelcol/connector/host_info.test",
    "github.com/grafana/alloy/internal/component/otelcol/connector/servicegraph_test [github.com/grafana/alloy/internal/component/otelcol/connector/servicegraph.test]",
    "github.com/grafana/alloy/internal/component/otelcol/connector/servicegraph.test",
    "github.com/grafana/alloy/internal/component/otelcol/connector/spanlogs_test [github.com/grafana/alloy/internal/component/otelcol/connector/spanlogs.test]",
    "github.com/grafana/alloy/internal/component/otelcol/connector/spanlogs.test",
    "github.com/grafana/alloy/internal/component/otelcol/connector/spanmetrics_test [github.com/grafana/alloy/internal/component/otelcol/connector/spanmetrics.test]",
    "github.com/grafana/alloy/internal/component/otelcol/connector/spanmetrics.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter_test [github.com/grafana/alloy/internal/component/otelcol/exporter.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/awss3_test [github.com/grafana/alloy/internal/component/otelcol/exporter/awss3.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/awss3.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/datadog_test [github.com/grafana/alloy/internal/component/otelcol/exporter/datadog.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/datadog.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/datadog/config_test [github.com/grafana/alloy/internal/component/otelcol/exporter/datadog/config.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/datadog/config.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/debug_test [github.com/grafana/alloy/internal/component/otelcol/exporter/debug.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/debug.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/faro_test [github.com/grafana/alloy/internal/component/otelcol/exporter/faro.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/faro.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/file [github.com/grafana/alloy/internal/component/otelcol/exporter/file.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/file.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/googlecloud_test [github.com/grafana/alloy/internal/component/otelcol/exporter/googlecloud.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/googlecloud.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/googlecloudpubsub_test [github.com/grafana/alloy/internal/component/otelcol/exporter/googlecloudpubsub.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/googlecloudpubsub.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/kafka_test [github.com/grafana/alloy/internal/component/otelcol/exporter/kafka.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/kafka.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/loadbalancing_test [github.com/grafana/alloy/internal/component/otelcol/exporter/loadbalancing.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/loadbalancing.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/loki/internal/convert_test [github.com/grafana/alloy/internal/component/otelcol/exporter/loki/internal/convert.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/loki/internal/convert.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/otlp_test [github.com/grafana/alloy/internal/component/otelcol/exporter/otlp.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/otlp.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/otlphttp_test [github.com/grafana/alloy/internal/component/otelcol/exporter/otlphttp.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/otlphttp.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test [github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus/internal/convert_test [github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus/internal/convert.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus/internal/convert.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/splunkhec_test [github.com/grafana/alloy/internal/component/otelcol/exporter/splunkhec.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/splunkhec.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/splunkhec/config_test [github.com/grafana/alloy/internal/component/otelcol/exporter/splunkhec/config.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/splunkhec/config.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/syslog_test [github.com/grafana/alloy/internal/component/otelcol/exporter/syslog.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/syslog.test",
    "github.com/grafana/alloy/internal/component/otelcol/extension_test [github.com/grafana/alloy/internal/component/otelcol/extension.test]",
    "github.com/grafana/alloy/internal/component/otelcol/extension.test",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling_test [github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling.test]",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling.test",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling [github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling.test]",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling.test",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal [github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal.test]",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal.test",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/server/grpc [github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/server/grpc.test]",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/server/grpc.test",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/server/http [github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/server/http.test]",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/server/http.test",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/filesource [github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/filesource.test]",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/filesource.test",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/remotesource [github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/remotesource.test]",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/remotesource.test",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/strategy_store [github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/strategy_store.test]",
    "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling/internal/jaegerremotesampling/internal/source/strategy_store.test",
    "github.com/grafana/alloy/internal/component/otelcol/internal/lazyconsumer [github.com/grafana/alloy/internal/component/otelcol/internal/lazyconsumer.test]",
    "github.com/grafana/alloy/internal/component/otelcol/internal/lazyconsumer.test",
    "github.com/grafana/alloy/internal/component/otelcol/internal/scheduler_test [github.com/grafana/alloy/internal/component/otelcol/internal/scheduler.test]",
    "github.com/grafana/alloy/internal/component/otelcol/internal/scheduler.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor_test [github.com/grafana/alloy/internal/component/otelcol/processor.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/attributes_test [github.com/grafana/alloy/internal/component/otelcol/processor/attributes.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/attributes.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/batch_test [github.com/grafana/alloy/internal/component/otelcol/processor/batch.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/batch.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/cumulativetodelta_test [github.com/grafana/alloy/internal/component/otelcol/processor/cumulativetodelta.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/cumulativetodelta.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/discovery_test [github.com/grafana/alloy/internal/component/otelcol/processor/discovery.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/discovery.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/filter_test [github.com/grafana/alloy/internal/component/otelcol/processor/filter.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/filter.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/groupbyattrs_test [github.com/grafana/alloy/internal/component/otelcol/processor/groupbyattrs.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/groupbyattrs.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/k8sattributes_test [github.com/grafana/alloy/internal/component/otelcol/processor/k8sattributes.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/k8sattributes.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/memorylimiter_test [github.com/grafana/alloy/internal/component/otelcol/processor/memorylimiter.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/memorylimiter.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/probabilistic_sampler_test [github.com/grafana/alloy/internal/component/otelcol/processor/probabilistic_sampler.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/probabilistic_sampler.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/processortest [github.com/grafana/alloy/internal/component/otelcol/processor/processortest.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/processortest.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection_test [github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/span_test [github.com/grafana/alloy/internal/component/otelcol/processor/span.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/span.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/tail_sampling [github.com/grafana/alloy/internal/component/otelcol/processor/tail_sampling.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/tail_sampling.test",
    "github.com/grafana/alloy/internal/component/otelcol/processor/transform_test [github.com/grafana/alloy/internal/component/otelcol/processor/transform.test]",
    "github.com/grafana/alloy/internal/component/otelcol/processor/transform.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver_test [github.com/grafana/alloy/internal/component/otelcol/receiver.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/awscloudwatch_test [github.com/grafana/alloy/internal/component/otelcol/receiver/awscloudwatch.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/awscloudwatch.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/awsecscontainermetrics_test [github.com/grafana/alloy/internal/component/otelcol/receiver/awsecscontainermetrics.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/awsecscontainermetrics.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/awss3_test [github.com/grafana/alloy/internal/component/otelcol/receiver/awss3.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/awss3.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/cloudflare_test [github.com/grafana/alloy/internal/component/otelcol/receiver/cloudflare.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/cloudflare.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/datadog_test [github.com/grafana/alloy/internal/component/otelcol/receiver/datadog.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/datadog.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/faro_test [github.com/grafana/alloy/internal/component/otelcol/receiver/faro.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/faro.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/file_stats_test [github.com/grafana/alloy/internal/component/otelcol/receiver/file_stats.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/file_stats.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/filelog_test [github.com/grafana/alloy/internal/component/otelcol/receiver/filelog.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/filelog.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/fluentforward_test [github.com/grafana/alloy/internal/component/otelcol/receiver/fluentforward.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/fluentforward.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/googlecloudpubsub [github.com/grafana/alloy/internal/component/otelcol/receiver/googlecloudpubsub.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/googlecloudpubsub.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/influxdb_test [github.com/grafana/alloy/internal/component/otelcol/receiver/influxdb.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/influxdb.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/jaeger_test [github.com/grafana/alloy/internal/component/otelcol/receiver/jaeger.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/jaeger.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/kafka_test [github.com/grafana/alloy/internal/component/otelcol/receiver/kafka.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/kafka.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/loki [github.com/grafana/alloy/internal/component/otelcol/receiver/loki.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/loki.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/opencensus_test [github.com/grafana/alloy/internal/component/otelcol/receiver/opencensus.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/opencensus.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/otlp_test [github.com/grafana/alloy/internal/component/otelcol/receiver/otlp.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/otlp.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test [github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus/internal [github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus/internal.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus/internal_test [github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus/internal.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus/internal.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/solace_test [github.com/grafana/alloy/internal/component/otelcol/receiver/solace.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/solace.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/splunkhec [github.com/grafana/alloy/internal/component/otelcol/receiver/splunkhec.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/splunkhec.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/syslog_test [github.com/grafana/alloy/internal/component/otelcol/receiver/syslog.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/syslog.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/tcplog_test [github.com/grafana/alloy/internal/component/otelcol/receiver/tcplog.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/tcplog.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/vcenter_test [github.com/grafana/alloy/internal/component/otelcol/receiver/vcenter.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/vcenter.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/zipkin_test [github.com/grafana/alloy/internal/component/otelcol/receiver/zipkin.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/zipkin.test",
    "github.com/grafana/alloy/internal/component/otelcol/storage/file_test [github.com/grafana/alloy/internal/component/otelcol/storage/file.test]",
    "github.com/grafana/alloy/internal/component/otelcol/storage/file.test",
    "github.com/grafana/alloy/internal/component/prometheus [github.com/grafana/alloy/internal/component/prometheus.test]",
    "github.com/grafana/alloy/internal/component/prometheus_test [github.com/grafana/alloy/internal/component/prometheus.test]",
    "github.com/grafana/alloy/internal/component/prometheus.test",
    "github.com/grafana/alloy/internal/component/prometheus/appenders [github.com/grafana/alloy/internal/component/prometheus/appenders.test]",
    "github.com/grafana/alloy/internal/component/prometheus/appenders.test",
    "github.com/grafana/alloy/internal/component/prometheus/echo [github.com/grafana/alloy/internal/component/prometheus/echo.test]",
    "github.com/grafana/alloy/internal/component/prometheus/echo.test",
    "github.com/grafana/alloy/internal/component/prometheus/enrich [github.com/grafana/alloy/internal/component/prometheus/enrich.test]",
    "github.com/grafana/alloy/internal/component/prometheus/enrich.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/blackbox [github.com/grafana/alloy/internal/component/prometheus/exporter/blackbox.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/blackbox.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/cadvisor [github.com/grafana/alloy/internal/component/prometheus/exporter/cadvisor.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/cadvisor.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/catchpoint [github.com/grafana/alloy/internal/component/prometheus/exporter/catchpoint.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/catchpoint.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/cloudwatch [github.com/grafana/alloy/internal/component/prometheus/exporter/cloudwatch.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/cloudwatch.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/dnsmasq [github.com/grafana/alloy/internal/component/prometheus/exporter/dnsmasq.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/dnsmasq.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/elasticsearch [github.com/grafana/alloy/internal/component/prometheus/exporter/elasticsearch.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/elasticsearch.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/gcp [github.com/grafana/alloy/internal/component/prometheus/exporter/gcp.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/gcp.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/github [github.com/grafana/alloy/internal/component/prometheus/exporter/github.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/github.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/kafka [github.com/grafana/alloy/internal/component/prometheus/exporter/kafka.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/kafka.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/memcached [github.com/grafana/alloy/internal/component/prometheus/exporter/memcached.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/memcached.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/mongodb [github.com/grafana/alloy/internal/component/prometheus/exporter/mongodb.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/mongodb.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/mssql [github.com/grafana/alloy/internal/component/prometheus/exporter/mssql.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/mssql.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/mysql [github.com/grafana/alloy/internal/component/prometheus/exporter/mysql.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/mysql.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/oracledb [github.com/grafana/alloy/internal/component/prometheus/exporter/oracledb.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/oracledb.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/postgres [github.com/grafana/alloy/internal/component/prometheus/exporter/postgres.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/postgres.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/process [github.com/grafana/alloy/internal/component/prometheus/exporter/process.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/process.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/redis [github.com/grafana/alloy/internal/component/prometheus/exporter/redis.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/redis.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/snmp [github.com/grafana/alloy/internal/component/prometheus/exporter/snmp.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/snmp.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/snowflake [github.com/grafana/alloy/internal/component/prometheus/exporter/snowflake.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/snowflake.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/squid [github.com/grafana/alloy/internal/component/prometheus/exporter/squid.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/squid.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/static [github.com/grafana/alloy/internal/component/prometheus/exporter/static.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/static.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/statsd [github.com/grafana/alloy/internal/component/prometheus/exporter/statsd.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/statsd.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test [github.com/grafana/alloy/internal/component/prometheus/exporter/tests.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/windows [github.com/grafana/alloy/internal/component/prometheus/exporter/windows.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/windows.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator [github.com/grafana/alloy/internal/component/prometheus/operator.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common [github.com/grafana/alloy/internal/component/prometheus/operator/common.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen [github.com/grafana/alloy/internal/component/prometheus/operator/configgen.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen.test",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http [github.com/grafana/alloy/internal/component/prometheus/receive_http.test]",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http.test",
    "github.com/grafana/alloy/internal/component/prometheus/relabel [github.com/grafana/alloy/internal/component/prometheus/relabel.test]",
    "github.com/grafana/alloy/internal/component/prometheus/relabel.test",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite [github.com/grafana/alloy/internal/component/prometheus/remotewrite.test]",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test [github.com/grafana/alloy/internal/component/prometheus/remotewrite.test]",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite.test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape [github.com/grafana/alloy/internal/component/prometheus/scrape.test]",
    "github.com/grafana/alloy/internal/component/prometheus/scrape.test",
    "github.com/grafana/alloy/internal/component/prometheus/write/queue [github.com/grafana/alloy/internal/component/prometheus/write/queue.test]",
    "github.com/grafana/alloy/internal/component/prometheus/write/queue.test",
    "github.com/grafana/alloy/internal/component/pyroscope [github.com/grafana/alloy/internal/component/pyroscope.test]",
    "github.com/grafana/alloy/internal/component/pyroscope.test",
    "github.com/grafana/alloy/internal/component/pyroscope/ebpf [github.com/grafana/alloy/internal/component/pyroscope/ebpf.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/ebpf.test",
    "github.com/grafana/alloy/internal/component/pyroscope/ebpf/reporter [github.com/grafana/alloy/internal/component/pyroscope/ebpf/reporter.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/ebpf/reporter.test",
    "github.com/grafana/alloy/internal/component/pyroscope/enrich [github.com/grafana/alloy/internal/component/pyroscope/enrich.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/enrich.test",
    "github.com/grafana/alloy/internal/component/pyroscope/java [github.com/grafana/alloy/internal/component/pyroscope/java.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/java.test",
    "github.com/grafana/alloy/internal/component/pyroscope/java/asprof [github.com/grafana/alloy/internal/component/pyroscope/java/asprof.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/java/asprof.test",
    "github.com/grafana/alloy/internal/component/pyroscope/java/integration [github.com/grafana/alloy/internal/component/pyroscope/java/integration.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/java/integration.test",
    "github.com/grafana/alloy/internal/component/pyroscope/receive_http [github.com/grafana/alloy/internal/component/pyroscope/receive_http.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/receive_http.test",
    "github.com/grafana/alloy/internal/component/pyroscope/relabel [github.com/grafana/alloy/internal/component/pyroscope/relabel.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/relabel.test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape [github.com/grafana/alloy/internal/component/pyroscope/scrape.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape.test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape/internal/fastdelta [github.com/grafana/alloy/internal/component/pyroscope/scrape/internal/fastdelta.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape/internal/fastdelta_test [github.com/grafana/alloy/internal/component/pyroscope/scrape/internal/fastdelta.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape/internal/fastdelta.test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape/internal/pproflite_test [github.com/grafana/alloy/internal/component/pyroscope/scrape/internal/pproflite.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape/internal/pproflite.test",
    "github.com/grafana/alloy/internal/component/pyroscope/write [github.com/grafana/alloy/internal/component/pyroscope/write.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/write.test",
    "github.com/grafana/alloy/internal/component/remote/http_test [github.com/grafana/alloy/internal/component/remote/http.test]",
    "github.com/grafana/alloy/internal/component/remote/http.test",
    "github.com/grafana/alloy/internal/component/remote/kubernetes [github.com/grafana/alloy/internal/component/remote/kubernetes.test]",
    "github.com/grafana/alloy/internal/component/remote/kubernetes.test",
    "github.com/grafana/alloy/internal/component/remote/s3 [github.com/grafana/alloy/internal/component/remote/s3.test]",
    "github.com/grafana/alloy/internal/component/remote/s3.test",
    "github.com/grafana/alloy/internal/component/remote/vault [github.com/grafana/alloy/internal/component/remote/vault.test]",
    "github.com/grafana/alloy/internal/component/remote/vault.test",
    "github.com/grafana/alloy/internal/converter/diag [github.com/grafana/alloy/internal/converter/diag.test]",
    "github.com/grafana/alloy/internal/converter/diag.test",
    "github.com/grafana/alloy/internal/converter/internal/common_test [github.com/grafana/alloy/internal/converter/internal/common.test]",
    "github.com/grafana/alloy/internal/converter/internal/common.test",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test [github.com/grafana/alloy/internal/converter/internal/otelcolconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test [github.com/grafana/alloy/internal/converter/internal/prometheusconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test [github.com/grafana/alloy/internal/converter/internal/promtailconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test [github.com/grafana/alloy/internal/converter/internal/staticconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build [github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build.test]",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build.test",
    "github.com/grafana/alloy/internal/dag [github.com/grafana/alloy/internal/dag.test]",
    "github.com/grafana/alloy/internal/dag.test",
    "github.com/grafana/alloy/internal/featuregate [github.com/grafana/alloy/internal/featuregate.test]",
    "github.com/grafana/alloy/internal/featuregate.test",
    "github.com/grafana/alloy/internal/loki/client [github.com/grafana/alloy/internal/loki/client.test]",
    "github.com/grafana/alloy/internal/loki/client.test",
    "github.com/grafana/alloy/internal/loki/client/internal [github.com/grafana/alloy/internal/loki/client/internal.test]",
    "github.com/grafana/alloy/internal/loki/client/internal.test",
    "github.com/grafana/alloy/internal/loki/logql [github.com/grafana/alloy/internal/loki/logql.test]",
    "github.com/grafana/alloy/internal/loki/logql.test",
    "github.com/grafana/alloy/internal/loki/util [github.com/grafana/alloy/internal/loki/util.test]",
    "github.com/grafana/alloy/internal/loki/util_test [github.com/grafana/alloy/internal/loki/util.test]",
    "github.com/grafana/alloy/internal/loki/util.test",
    "github.com/grafana/alloy/internal/loki/util/cfg [github.com/grafana/alloy/internal/loki/util/cfg.test]",
    "github.com/grafana/alloy/internal/loki/util/cfg.test",
    "github.com/grafana/alloy/internal/loki/util/flagext [github.com/grafana/alloy/internal/loki/util/flagext.test]",
    "github.com/grafana/alloy/internal/loki/util/flagext.test",
    "github.com/grafana/alloy/internal/mimir/client [github.com/grafana/alloy/internal/mimir/client.test]",
    "github.com/grafana/alloy/internal/mimir/client.test",
    "github.com/grafana/alloy/internal/mimir/client/internal [github.com/grafana/alloy/internal/mimir/client/internal.test]",
    "github.com/grafana/alloy/internal/mimir/client/internal.test",
    "github.com/grafana/alloy/internal/runner [github.com/grafana/alloy/internal/runner.test]",
    "github.com/grafana/alloy/internal/runner_test [github.com/grafana/alloy/internal/runner.test]",
    "github.com/grafana/alloy/internal/runner.test",
    "github.com/grafana/alloy/internal/runtime [github.com/grafana/alloy/internal/runtime.test]",
    "github.com/grafana/alloy/internal/runtime_test [github.com/grafana/alloy/internal/runtime.test]",
    "github.com/grafana/alloy/internal/runtime.test",
    "github.com/grafana/alloy/internal/runtime/equality [github.com/grafana/alloy/internal/runtime/equality.test]",
    "github.com/grafana/alloy/internal/runtime/equality.test",
    "github.com/grafana/alloy/internal/runtime/internal/controller [github.com/grafana/alloy/internal/runtime/internal/controller.test]",
    "github.com/grafana/alloy/internal/runtime/internal/controller_test [github.com/grafana/alloy/internal/runtime/internal/controller.test]",
    "github.com/grafana/alloy/internal/runtime/internal/controller.test",
    "github.com/grafana/alloy/internal/runtime/internal/worker [github.com/grafana/alloy/internal/runtime/internal/worker.test]",
    "github.com/grafana/alloy/internal/runtime/internal/worker.test",
    "github.com/grafana/alloy/internal/runtime/logging [github.com/grafana/alloy/internal/runtime/logging.test]",
    "github.com/grafana/alloy/internal/runtime/logging_test [github.com/grafana/alloy/internal/runtime/logging.test]",
    "github.com/grafana/alloy/internal/runtime/logging.test",
    "github.com/grafana/alloy/internal/service/cluster [github.com/grafana/alloy/internal/service/cluster.test]",
    "github.com/grafana/alloy/internal/service/cluster_test [github.com/grafana/alloy/internal/service/cluster.test]",
    "github.com/grafana/alloy/internal/service/cluster.test",
    "github.com/grafana/alloy/internal/service/cluster/discovery [github.com/grafana/alloy/internal/service/cluster/discovery.test]",
    "github.com/grafana/alloy/internal/service/cluster/discovery.test",
    "github.com/grafana/alloy/internal/service/http [github.com/grafana/alloy/internal/service/http.test]",
    "github.com/grafana/alloy/internal/service/http.test",
    "github.com/grafana/alloy/internal/service/labelstore [github.com/grafana/alloy/internal/service/labelstore.test]",
    "github.com/grafana/alloy/internal/service/labelstore.test",
    "github.com/grafana/alloy/internal/service/livedebugging [github.com/grafana/alloy/internal/service/livedebugging.test]",
    "github.com/grafana/alloy/internal/service/livedebugging.test",
    "github.com/grafana/alloy/internal/service/remotecfg [github.com/grafana/alloy/internal/service/remotecfg.test]",
    "github.com/grafana/alloy/internal/service/remotecfg.test",
    "github.com/grafana/alloy/internal/slogadapter [github.com/grafana/alloy/internal/slogadapter.test]",
    "github.com/grafana/alloy/internal/slogadapter.test",
    "github.com/grafana/alloy/internal/static/agentctl/waltools [github.com/grafana/alloy/internal/static/agentctl/waltools.test]",
    "github.com/grafana/alloy/internal/static/agentctl/waltools.test",
    "github.com/grafana/alloy/internal/static/config [github.com/grafana/alloy/internal/static/config.test]",
    "github.com/grafana/alloy/internal/static/config.test",
    "github.com/grafana/alloy/internal/static/config/features [github.com/grafana/alloy/internal/static/config/features.test]",
    "github.com/grafana/alloy/internal/static/config/features.test",
    "github.com/grafana/alloy/internal/static/integrations [github.com/grafana/alloy/internal/static/integrations.test]",
    "github.com/grafana/alloy/internal/static/integrations.test",
    "github.com/grafana/alloy/internal/static/integrations/azure_exporter_test [github.com/grafana/alloy/internal/static/integrations/azure_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/azure_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/blackbox_exporter [github.com/grafana/alloy/internal/static/integrations/blackbox_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/blackbox_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/cadvisor [github.com/grafana/alloy/internal/static/integrations/cadvisor.test]",
    "github.com/grafana/alloy/internal/static/integrations/cadvisor.test",
    "github.com/grafana/alloy/internal/static/integrations/catchpoint_exporter [github.com/grafana/alloy/internal/static/integrations/catchpoint_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/catchpoint_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/cloudwatch_exporter [github.com/grafana/alloy/internal/static/integrations/cloudwatch_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/cloudwatch_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/gcp_exporter_test [github.com/grafana/alloy/internal/static/integrations/gcp_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/gcp_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/github_exporter [github.com/grafana/alloy/internal/static/integrations/github_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/github_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/install [github.com/grafana/alloy/internal/static/integrations/install.test]",
    "github.com/grafana/alloy/internal/static/integrations/install.test",
    "github.com/grafana/alloy/internal/static/integrations/kafka_exporter [github.com/grafana/alloy/internal/static/integrations/kafka_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/kafka_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/mongodb_exporter [github.com/grafana/alloy/internal/static/integrations/mongodb_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/mongodb_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/mssql [github.com/grafana/alloy/internal/static/integrations/mssql.test]",
    "github.com/grafana/alloy/internal/static/integrations/mssql.test",
    "github.com/grafana/alloy/internal/static/integrations/mysqld_exporter [github.com/grafana/alloy/internal/static/integrations/mysqld_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/mysqld_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/node_exporter [github.com/grafana/alloy/internal/static/integrations/node_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/node_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/oracledb_exporter [github.com/grafana/alloy/internal/static/integrations/oracledb_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/oracledb_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/postgres_exporter [github.com/grafana/alloy/internal/static/integrations/postgres_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/postgres_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/redis_exporter [github.com/grafana/alloy/internal/static/integrations/redis_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/redis_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/snmp_exporter [github.com/grafana/alloy/internal/static/integrations/snmp_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/snmp_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/snowflake_exporter [github.com/grafana/alloy/internal/static/integrations/snowflake_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/snowflake_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/squid_exporter [github.com/grafana/alloy/internal/static/integrations/squid_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/squid_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/v2 [github.com/grafana/alloy/internal/static/integrations/v2.test]",
    "github.com/grafana/alloy/internal/static/integrations/v2.test",
    "github.com/grafana/alloy/internal/static/integrations/v2/apache_http [github.com/grafana/alloy/internal/static/integrations/v2/apache_http.test]",
    "github.com/grafana/alloy/internal/static/integrations/v2/apache_http.test",
    "github.com/grafana/alloy/internal/static/integrations/v2/app_agent_receiver [github.com/grafana/alloy/internal/static/integrations/v2/app_agent_receiver.test]",
    "github.com/grafana/alloy/internal/static/integrations/v2/app_agent_receiver.test",
    "github.com/grafana/alloy/internal/static/integrations/v2/blackbox_exporter [github.com/grafana/alloy/internal/static/integrations/v2/blackbox_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/v2/blackbox_exporter.test",
    "github.com/grafana/alloy/internal/static/integrations/v2/metricsutils [github.com/grafana/alloy/internal/static/integrations/v2/metricsutils.test]",
    "github.com/grafana/alloy/internal/static/integrations/v2/metricsutils.test",
    "github.com/grafana/alloy/internal/static/integrations/v2/snmp_exporter [github.com/grafana/alloy/internal/static/integrations/v2/snmp_exporter.test]",
    "github.com/grafana/alloy/internal/static/integrations/v2/snmp_exporter.test",
    "github.com/grafana/alloy/internal/static/logs [github.com/grafana/alloy/internal/static/logs.test]",
    "github.com/grafana/alloy/internal/static/logs.test",
    "github.com/grafana/alloy/internal/static/metrics [github.com/grafana/alloy/internal/static/metrics.test]",
    "github.com/grafana/alloy/internal/static/metrics.test",
    "github.com/grafana/alloy/internal/static/metrics/cluster [github.com/grafana/alloy/internal/static/metrics/cluster.test]",
    "github.com/grafana/alloy/internal/static/metrics/cluster.test",
    "github.com/grafana/alloy/internal/static/metrics/instance [github.com/grafana/alloy/internal/static/metrics/instance.test]",
    "github.com/grafana/alloy/internal/static/metrics/instance.test",
    "github.com/grafana/alloy/internal/static/metrics/wal [github.com/grafana/alloy/internal/static/metrics/wal.test]",
    "github.com/grafana/alloy/internal/static/metrics/wal.test",
    "github.com/grafana/alloy/internal/static/server [github.com/grafana/alloy/internal/static/server.test]",
    "github.com/grafana/alloy/internal/static/server.test",
    "github.com/grafana/alloy/internal/static/traces [github.com/grafana/alloy/internal/static/traces.test]",
    "github.com/grafana/alloy/internal/static/traces.test",
    "github.com/grafana/alloy/internal/static/traces/automaticloggingprocessor [github.com/grafana/alloy/internal/static/traces/automaticloggingprocessor.test]",
    "github.com/grafana/alloy/internal/static/traces/automaticloggingprocessor.test",
    "github.com/grafana/alloy/internal/static/traces/promsdprocessor [github.com/grafana/alloy/internal/static/traces/promsdprocessor.test]",
    "github.com/grafana/alloy/internal/static/traces/promsdprocessor.test",
    "github.com/grafana/alloy/internal/static/traces/promsdprocessor/consumer [github.com/grafana/alloy/internal/static/traces/promsdprocessor/consumer.test]",
    "github.com/grafana/alloy/internal/static/traces/promsdprocessor/consumer.test",
    "github.com/grafana/alloy/internal/static/traces/servicegraphprocessor [github.com/grafana/alloy/internal/static/traces/servicegraphprocessor.test]",
    "github.com/grafana/alloy/internal/static/traces/servicegraphprocessor.test",
    "github.com/grafana/alloy/internal/static/traces/spanmetricsprocessor/internal/cache [github.com/grafana/alloy/internal/static/traces/spanmetricsprocessor/internal/cache.test]",
    "github.com/grafana/alloy/internal/static/traces/spanmetricsprocessor/internal/cache.test",
    "github.com/grafana/alloy/internal/tools/docs_generator_test [github.com/grafana/alloy/internal/tools/docs_generator.test]",
    "github.com/grafana/alloy/internal/tools/docs_generator.test",
    "github.com/grafana/alloy/internal/usagestats [github.com/grafana/alloy/internal/usagestats.test]",
    "github.com/grafana/alloy/internal/usagestats.test",
    "github.com/grafana/alloy/internal/useragent [github.com/grafana/alloy/internal/useragent.test]",
    "github.com/grafana/alloy/internal/useragent.test",
    "github.com/grafana/alloy/internal/util [github.com/grafana/alloy/internal/util.test]",
    "github.com/grafana/alloy/internal/util.test",
    "github.com/grafana/alloy/internal/util/assertmetrics [github.com/grafana/alloy/internal/util/assertmetrics.test]",
    "github.com/grafana/alloy/internal/util/assertmetrics.test",
    "github.com/grafana/alloy/internal/util/jitter [github.com/grafana/alloy/internal/util/jitter.test]",
    "github.com/grafana/alloy/internal/util/jitter.test",
    "github.com/grafana/alloy/internal/util/testappender_test [github.com/grafana/alloy/internal/util/testappender.test]",
    "github.com/grafana/alloy/internal/util/testappender.test",
    "github.com/grafana/alloy/internal/util/testlivedebugging_test [github.com/grafana/alloy/internal/util/testlivedebugging.test]",
    "github.com/grafana/alloy/internal/util/testlivedebugging.test",
    "github.com/grafana/alloy/internal/util/testtarget [github.com/grafana/alloy/internal/util/testtarget.test]",
    "github.com/grafana/alloy/internal/util/testtarget.test",
    "github.com/grafana/alloy/internal/util/wildcard_test [github.com/grafana/alloy/internal/util/wildcard.test]",
    "github.com/grafana/alloy/internal/util/wildcard.test",
    "github.com/grafana/alloy/internal/util/zapadapter_test [github.com/grafana/alloy/internal/util/zapadapter.test]",
    "github.com/grafana/alloy/internal/util/zapadapter.test",
    "github.com/grafana/alloy/internal/validator [github.com/grafana/alloy/internal/validator.test]",
    "github.com/grafana/alloy/internal/validator.test",
    "github.com/grafana/alloy/internal/vcs_test [github.com/grafana/alloy/internal/vcs.test]",
    "github.com/grafana/alloy/internal/vcs.test",
    "github.com/grafana/alloy/internal/web/ui [github.com/grafana/alloy/internal/web/ui.test]",
    "github.com/grafana/alloy/internal/web/ui.test"
  ],
  "Packages": [
    {
      "ID": "archive/tar",
      "Name": "tar",
      "PkgPath": "archive/tar",
      "GoFiles": [
        "/usr/lib/go/src/archive/tar/common.go",
        "/usr/lib/go/src/archive/tar/format.go",
        "/usr/lib/go/src/archive/tar/reader.go",
        "/usr/lib/go/src/archive/tar/stat_actime1.go",
        "/usr/lib/go/src/archive/tar/stat_unix.go",
        "/usr/lib/go/src/archive/tar/strconv.go",
        "/usr/lib/go/src/archive/tar/writer.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/archive/tar/common.go",
        "/usr/lib/go/src/archive/tar/format.go",
        "/usr/lib/go/src/archive/tar/reader.go",
        "/usr/lib/go/src/archive/tar/stat_actime1.go",
        "/usr/lib/go/src/archive/tar/stat_unix.go",
        "/usr/lib/go/src/archive/tar/strconv.go",
        "/usr/lib/go/src/archive/tar/writer.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/archive/tar/stat_actime2.go"],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "fmt": "fmt",
        "internal/godebug": "internal/godebug",
        "io": "io",
        "io/fs": "io/fs",
        "maps": "maps",
        "math": "math",
        "os/user": "os/user",
        "path": "path",
        "path/filepath": "path/filepath",
        "reflect": "reflect",
        "runtime": "runtime",
        "slices": "slices",
        "strconv": "strconv",
        "strings": "strings",
        "sync": "sync",
        "syscall": "syscall",
        "time": "time"
      }
    },
    {
      "ID": "archive/zip",
      "Name": "zip",
      "PkgPath": "archive/zip",
      "GoFiles": [
        "/usr/lib/go/src/archive/zip/reader.go",
        "/usr/lib/go/src/archive/zip/register.go",
        "/usr/lib/go/src/archive/zip/struct.go",
        "/usr/lib/go/src/archive/zip/writer.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/archive/zip/reader.go",
        "/usr/lib/go/src/archive/zip/register.go",
        "/usr/lib/go/src/archive/zip/struct.go",
        "/usr/lib/go/src/archive/zip/writer.go"
      ],
      "Imports": {
        "bufio": "bufio",
        "compress/flate": "compress/flate",
        "encoding/binary": "encoding/binary",
        "errors": "errors",
        "fmt": "fmt",
        "hash": "hash",
        "hash/crc32": "hash/crc32",
        "internal/godebug": "internal/godebug",
        "io": "io",
        "io/fs": "io/fs",
        "os": "os",
        "path": "path",
        "path/filepath": "path/filepath",
        "slices": "slices",
        "strings": "strings",
        "sync": "sync",
        "time": "time",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "builtin",
      "Name": "builtin",
      "PkgPath": "builtin",
      "GoFiles": ["/usr/lib/go/src/builtin/builtin.go"],
      "CompiledGoFiles": ["/usr/lib/go/src/builtin/builtin.go"],
      "Imports": {
        "cmp": "cmp"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 5631 to 5 entries.

#### drv #2

Trace meta: spanId=3, ts=1770837170484, ts_iso=2026-02-11T19:12:50.484000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "file=/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkiZm10IgoJInN5bmMiCgkic3luYy9hdG9taWMiCgkidGltZSIKCgkiZ2l0aHViLmNvbS9oYXNoaWNvcnAvZ28tbXVsdGllcnJvciIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvY2xpZW50X2dvbGFuZy9wcm9tZXRoZXVzIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2V4ZW1wbGFyIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2hpc3RvZ3JhbSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9sYWJlbHMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvbWV0YWRhdGEiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvc3RvcmFnZSIKKQoKdHlwZSBNYXBwaW5nU3RvcmUgaW50ZXJmYWNlIHsKCUdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYKCUNyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmCglVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpCglUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKQoJR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwKfQoKdHlwZSBzZXJpZXNSZWZNYXBwaW5nIHN0cnVjdCB7CglzdGFydCAgICB0aW1lLlRpbWUKCWNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlcgoJc3RvcmUgICAgTWFwcGluZ1N0b3JlCgoJdW5pcXVlUmVmQ2VsbCAqQ2VsbAoKCS8vIGNoaWxkUmVmcyBpcyByZXVzZWQgZm9yIGVhY2ggYXBwZW5kIGNhbGwgdG8gYXZvaWQgYWxsb2NhdGlvbnMuIFRoaXMgaXMgc2FmZSBiZWNhdXNlIHN0b3JhZ2UuQXBwZW5kZXIgc2hvdWxkIG5ldmVyCgkvLyBoYXZlIGNvbmN1cnJlbnQgY2FsbHMgdG8gQXBwZW5kIG1ldGhvZHMuCgljaGlsZFJlZnMgICAgICAgIFtdc3RvcmFnZS5TZXJpZXNSZWYKCXdyaXRlTGF0ZW5jeSAgICAgcHJvbWV0aGV1cy5IaXN0b2dyYW0KCXNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZyhjaGlsZHJlbiBbXXN0b3JhZ2UuQXBwZW5kZXIsIHN0b3JlIE1hcHBpbmdTdG9yZSwgd3JpdGVMYXRlbmN5IHByb21ldGhldXMuSGlzdG9ncmFtLCBzYW1wbGVzRm9yd2FyZGVkIHByb21ldGhldXMuQ291bnRlcikgc3RvcmFnZS5BcHBlbmRlciB7Cgl1bmlxdWVSZWZDZWxsIDo9IHN0b3JlLkdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpCgoJcmV0dXJuICZzZXJpZXNSZWZNYXBwaW5newoJCWNoaWxkcmVuOiAgICAgICAgIGNoaWxkcmVuLAoJCXN0b3JlOiAgICAgICAgICAgIHN0b3JlLAoJCXdyaXRlTGF0ZW5jeTogICAgIHdyaXRlTGF0ZW5jeSwKCQlzYW1wbGVzRm9yd2FyZGVkOiBzYW1wbGVzRm9yd2FyZGVkLAoKCQl1bmlxdWVSZWZDZWxsOiB1bmlxdWVSZWZDZWxsLAoJCWNoaWxkUmVmczogICAgIG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgbGVuKGNoaWxkcmVuKSksCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFNldE9wdGlvbnMob3B0cyAqc3RvcmFnZS5BcHBlbmRPcHRpb25zKSB7CglmbXQuUHJpbnRsbigpCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQljLlNldE9wdGlvbnMob3B0cykKCX0KfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQ29tbWl0KCkgZXJyb3IgewoJcy5zdG9yZS5UcmFja0FwcGVuZGVkU2VyaWVzKHRpbWUuTm93KCkuVW5peCgpLCBzLnVuaXF1ZVJlZkNlbGwpCgoJdmFyIG11bHRpRXJyIGVycm9yCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQllcnIgOj0gYy5Db21taXQoKQoJCWlmIGVyciAhPSBuaWwgewoJCQltdWx0aUVyciA9IG11bHRpZXJyb3IuQXBwZW5kKG11bHRpRXJyLCBlcnIpCgkJfQoJfQoJcmV0dXJuIG11bHRpRXJyCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFJvbGxiYWNrKCkgZXJyb3IgewoJLy8gV2Ugc3RpbGwgdHJhY2sgcm9sbGVkIGJhY2sgc2VyaWVzIHNvIHdlIGNhbiBwcm9wZXJseQoJLy8gY2xlYW4gdXAgYW55IHNlcmllcyB0aGF0IHdhcyBhcHBlbmRlZAoJcy5zdG9yZS5UcmFja0FwcGVuZGVkU2VyaWVzKHRpbWUuTm93KCkuVW5peCgpLCBzLnVuaXF1ZVJlZkNlbGwpCgoJdmFyIG11bHRpRXJyIGVycm9yCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQllcnIgOj0gYy5Sb2xsYmFjaygpCgkJaWYgZXJyICE9IG5pbCB7CgkJCW11bHRpRXJyID0gbXVsdGllcnJvci5BcHBlbmQobXVsdGlFcnIsIGVycikKCQl9Cgl9CglyZXR1cm4gbXVsdGlFcnIKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgcmVjb3JkTGF0ZW5jeSgpIHsKCWlmIHMuc3RhcnQuSXNaZXJvKCkgewoJCXJldHVybgoJfQoJZHVyYXRpb24gOj0gdGltZS5TaW5jZShzLnN0YXJ0KQoJcy53cml0ZUxhdGVuY3kuT2JzZXJ2ZShkdXJhdGlvbi5TZWNvbmRzKCkpCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIHJlc2V0RmllbGRzKCkgewoJLy8gUmVzZXQgY2hpbGRSZWZzIHNsaWNlIGxlbmd0aCB0byAwIGZvciByZXVzZQoJcy5jaGlsZFJlZnMgPSBzLmNoaWxkUmVmc1s6MF0KfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0IGludDY0LCB2IGZsb2F0NjQpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCW5ld1JlZiwgZXJyIDo9IGFwcGVuZGVyLkFwcGVuZChyZWYsIGwsIHQsIHYpCgkJaWYgZXJyID09IG5pbCB7CgkJCXMuc2FtcGxlc0ZvcndhcmRlZC5JbmMoKQoJCX0KCQlyZXR1cm4gbmV3UmVmLCBlcnIKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEV4ZW1wbGFyKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCBlIGV4ZW1wbGFyLkV4ZW1wbGFyKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kRXhlbXBsYXIocmVmLCBsLCBlKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kSGlzdG9ncmFtKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0IGludDY0LCBoICpoaXN0b2dyYW0uSGlzdG9ncmFtLCBmaCAqaGlzdG9ncmFtLkZsb2F0SGlzdG9ncmFtKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kSGlzdG9ncmFtKHJlZiwgbCwgdCwgaCwgZmgpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRIaXN0b2dyYW1DVFplcm9TYW1wbGUocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQsIGN0IGludDY0LCBoICpoaXN0b2dyYW0uSGlzdG9ncmFtLCBmaCAqaGlzdG9ncmFtLkZsb2F0SGlzdG9ncmFtKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kSGlzdG9ncmFtQ1RaZXJvU2FtcGxlKHJlZiwgbCwgdCwgY3QsIGgsIGZoKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgVXBkYXRlTWV0YWRhdGEocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIG0gbWV0YWRhdGEuTWV0YWRhdGEpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5VcGRhdGVNZXRhZGF0YShyZWYsIGwsIG0pCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRDVFplcm9TYW1wbGUocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQsIGN0IGludDY0KSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kQ1RaZXJvU2FtcGxlKHJlZiwgbCwgdCwgY3QpCgl9KQp9Cgp0eXBlIGFwcGVuZGVyRnVuYyBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikKCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIGFwcGVuZFRvQ2hpbGRyZW4ocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMsIGFmIGFwcGVuZGVyRnVuYykgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJZGVmZXIgcy5yZXNldEZpZWxkcygpCgoJaWYgcy5zdGFydC5Jc1plcm8oKSB7CgkJcy5zdGFydCA9IHRpbWUuTm93KCkKCX0KCgkvLyBDaGVjayBpZiB0aGUgaW5jb21pbmcgcmVmIGhhcyByZWYgbWFwcGluZ3MKCWV4aXN0aW5nQ2hpbGRSZWZzIDo9IHMuc3RvcmUuR2V0TWFwcGluZyhyZWYsIGxibHMpCgoJdmFyIGFwcGVuZEVyciBlcnJvcgoKCS8vIFNhbml0eSBjaGVjazogaWYgd2UgaGF2ZSBleGlzdGluZyBjaGlsZCByZWZzLCB0aGV5IG11c3QgbWF0Y2ggdGhlIG51bWJlciBvZiBjaGlsZHJlbgoJaWYgZXhpc3RpbmdDaGlsZFJlZnMgIT0gbmlsICYmIGxlbihleGlzdGluZ0NoaWxkUmVmcykgPT0gbGVuKHMuY2hpbGRyZW4pIHsKCQlzLnVuaXF1ZVJlZkNlbGwuUmVmcyA9IGFwcGVuZChzLnVuaXF1ZVJlZkNlbGwuUmVmcywgcmVmKQoKCQlyZWZVcGRhdGVSZXF1aXJlZCA6PSBmYWxzZQoJCWZvciBjaGlsZEluZGV4LCBjaGlsZFJlZiA6PSByYW5nZSBleGlzdGluZ0NoaWxkUmVmcyB7CgkJCW5ld0NoaWxkUmVmLCBlcnIgOj0gYWYocy5jaGlsZHJlbltjaGlsZEluZGV4XSwgY2hpbGRSZWYpCgkJCWlmIGVyciAhPSBuaWwgewoJCQkJYXBwZW5kRXJyID0gbXVsdGllcnJvci5BcHBlbmQoYXBwZW5kRXJyLCBlcnIpCgkJCX0KCgkJCWlmIG5ld0NoaWxkUmVmICE9IGNoaWxkUmVmIHsKCQkJCS8vIENoaWxkIHJlZiBjaGFuZ2VkLCBuZWVkIHRvIHVwZGF0ZSBtYXBwaW5nCgkJCQlleGlzdGluZ0NoaWxkUmVmc1tjaGlsZEluZGV4XSA9IG5ld0NoaWxkUmVmCgkJCQlyZWZVcGRhdGVSZXF1aXJlZCA9IHRydWUKCQkJfQoJCX0KCgkJaWYgYXBwZW5kRXJyICE9IG5pbCB7CgkJCXJldHVybiAwLCBhcHBlbmRFcnIKCQl9CgoJCWlmIHJlZlVwZGF0ZVJlcXVpcmVkIHsKCQkJcy5zdG9yZS5VcGRhdGVNYXBwaW5nKHJlZiwgZXhpc3RpbmdDaGlsZFJlZnMsIGxibHMpCgkJfQoKCQlyZXR1cm4gcmVmLCBuaWwKCX0KCgkvLyBObyBleGlzdGluZyBtYXBwaW5nLCBwcm9jZWVkIHdpdGggbm9ybWFsIGFwcGVuZCB0byBhbGwgY2hpbGRyZW4KCXZhciBmaXJzdE5vblplcm9SZWYgc3RvcmFnZS5TZXJpZXNSZWYKCXZhciBub25aZXJvQ291bnQgaW50CgoJLy8gTm90ZTogdGhlcmUncyBhbm90aGVyIG9wdGltaXphdGlvbiB3aGVyZSB3ZSBjb3VsZCB1c2UgdGhlIHJldHVybmVkIHJlZiBpZiBhbGwgdGhlIG5vbiB6ZXJvIHJlZnMKCS8vICBhcmUgdGhlIHNhbWUgdmFsdWUuIFRoaXMgaXNuJ3Qgc2FmZSBhcyB3ZSB3aWxsIG1peCBkb3duc3RyZWFtIHJlZnMgd2l0aCB1bmlxdWUgcmVmcyB3aGljaCBjb3VsZAoJLy8gIGNvbGxpZGUuIFdlIGNvdWxkIHN0YXJ0IGF0IG1heCB1bml0NjQgZm9yIHVuaXF1ZSByZWZzIGFuZCBnbyBiYWNrd2FyZHMgbGVzc2VuaW5nIHRoZSBjaGFuY2Ugb2YKCS8vIAljb2xsaXNpb25zIGJ1dCBpdCdzIHJhdGhlciBkYW5nZXJvdXMgZm9yIGFuIHVubGlrZWx5IGVkZ2UgY2FzZS4gSWYgdHdvIGNvbXBvbmVudHMgYXJlIHJldHVybmluZwoJLy8gCXRoZSBzYW1lIHJlZiBpdCdzIHR3byByZW1vdGVfd3JpdGUgY29tcG9uZW50cyB3aGljaCBzaG91bGQgcHJvYmFibHkgYmUgbWVyZ2VkIGluIHRvIG9uZS4KCWZvciBfLCBjaGlsZCA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQljaGlsZFJlZiwgZXJyIDo9IGFmKGNoaWxkLCByZWYpCgkJaWYgZXJyICE9IG5pbCB7CgkJCWFwcGVuZEVyciA9IG11bHRpZXJyb3IuQXBwZW5kKGFwcGVuZEVyciwgZXJyKQoKCQkJLy8gVE9ETyBzaG91bGQgSSBtYWtlIHRoZSBjaGlsZFJlZiB6ZXJvIGhlcmU/CgkJfQoKCQlzLmNoaWxkUmVmcyA9IGFwcGVuZChzLmNoaWxkUmVmcywgY2hpbGRSZWYpCgkJaWYgY2hpbGRSZWYgIT0gMCB7CgkJCWlmIGZpcnN0Tm9uWmVyb1JlZiA9PSAwIHsKCQkJCWZpcnN0Tm9uWmVyb1JlZiA9IGNoaWxkUmVmCgkJCX0KCQkJbm9uWmVyb0NvdW50KysKCQl9Cgl9CgoJaWYgYXBwZW5kRXJyICE9IG5pbCB7CgkJcmV0dXJuIDAsIGFwcGVuZEVycgoJfQoKCWlmIG5vblplcm9Db3VudCA9PSAwIHsKCQkvLyBBbGwgY2hpbGRyZW4gcmV0dXJuZWQgcmVmIDAsIHNvIHJldHVybiB0aGUgaW5wdXQgcmVmCgkJcmV0dXJuIHJlZiwgbmlsCgl9CgoJLy8gT25seSBvbmUgY2hpbGQgcmV0dXJuZWQgYSBub24temVybyByZWYsIHVzZSB0aGF0CglpZiBub25aZXJvQ291bnQgPT0gMSB7CgkJcmV0dXJuIGZpcnN0Tm9uWmVyb1JlZiwgbmlsCgl9CgoJLy8gV2UgZ290IGRpZmZlcmVudCByZWZzIGJhY2sgYW5kIG5lZWQgdG8gY3JlYXRlIGEgbmV3IG1hcHBpbmcKCXVuaXF1ZVJlZiA6PSBzLnN0b3JlLkNyZWF0ZU1hcHBpbmcocy5jaGlsZFJlZnMsIGxibHMpCglzLnVuaXF1ZVJlZkNlbGwuUmVmcyA9IGFwcGVuZChzLnVuaXF1ZVJlZkNlbGwuUmVmcywgdW5pcXVlUmVmKQoJcmV0dXJuIHVuaXF1ZVJlZiwgbmlsCn0KCnR5cGUgdW5pcVJlZkNoaWxkcmVuIHN0cnVjdCB7CgljaGlsZFJlZnMgKltdc3RvcmFnZS5TZXJpZXNSZWYKCWxhYmVsSGFzaCB1aW50NjQKfQoKdHlwZSBTZXJpZXNSZWZNYXBwaW5nU3RvcmUgc3RydWN0IHsKCS8vIHJlZk1hcHBpbmdNdSBwcm90ZWN0cyB1bmlxdWVSZWZUb0NoaWxkUmVmcywgbGFiZWxIYXNoVG9VbmlxdWVSZWYgYW5kIG5leHRVbmlxdWVSZWYKCXJlZk1hcHBpbmdNdSBzeW5jLlJXTXV0ZXgKCS8vIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzIG1hcHMgdGhlIHVuaXF1ZSByZWYgdG8gdGhlIGV4cGVjdGVkIGNoaWxkIHJlZiBpbiBvcmRlcgoJdW5pcXVlUmVmVG9DaGlsZFJlZnMgbWFwW3N0b3JhZ2UuU2VyaWVzUmVmXXVuaXFSZWZDaGlsZHJlbgoJLy8gbGFiZWxIYXNoVG9VbmlxdWVSZWYgbWFwcyB0aGUgbGFiZWwgaGFzaCB0byB1bmlxdWUgcmVmLgoJbGFiZWxIYXNoVG9VbmlxdWVSZWYgbWFwW3VpbnQ2NF1zdG9yYWdlLlNlcmllc1JlZgoKCS8vIG5leHRVbmlxdWVSZWYgaXMgdGhlIG5leHQgcmVmIElEIHdlIHdpbGwgaGFuZCBvdXQKCW5leHRVbmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYKCgkvLyB0aW1lc3RhbXBUcmFja2luZ011IHByb3RlY3RzIHVuaXF1ZVJlZlRpbWVzdGFtcHMgYW5kIGNlbGxQb29sCgl0aW1lc3RhbXBUcmFja2luZ011IHN5bmMuTXV0ZXgKCS8vIHVuaXF1ZVJlZlRpbWVzdGFtcHMgbWFwcyB1bmlxdWUgcmVmcyB0byB0aGVpciBsYXN0IGFwcGVuZCB0aW1lc3RhbXAKCXVuaXF1ZVJlZlRpbWVzdGFtcHMgbWFwW3N0b3JhZ2UuU2VyaWVzUmVmXWludDY0CgkvLyBjZWxsUG9vbCBpcyB1c2VkIHRvIHBvb2wgc2xpY2VzIG9mIFNlcmllc1JlZnMgdXNlZCBmb3IgdHJhY2tpbmcgdW5pcXVlIHJlZnMgaW4gVHJhY2tBcHBlbmRlZFNlcmllcy4KCWNlbGxQb29sIHN5bmMuUG9vbAoKCS8vIENsZWFudXAgZ29yb3V0aW5lIGNvb3JkaW5hdGlvbiAobm8gbG9jayByZXF1aXJlZCkKCXN0YXJ0UmVmQ2xlYW51cCBzeW5jLk9uY2UKCWNsZWFudXBTdGFydGVkICBhdG9taWMuQm9vbAoJc3RvcENsZWFudXAgICAgIGNoYW4gc3RydWN0e30KCWNsZWFudXBTdG9wcGVkICBjaGFuIHN0cnVjdHt9CgoJLy8gTWV0cmljcyAoc2FmZSBmb3IgY29uY3VycmVudCBhY2Nlc3MsIG5vIGxvY2sgcmVxdWlyZWQpCglhY3RpdmVNYXBwaW5ncyAgcHJvbWV0aGV1cy5HYXVnZQoJdHJhY2tlZFJlZnMgICAgIHByb21ldGhldXMuR2F1Z2UKCXJlZnNDbGVhbmVkICAgICBwcm9tZXRoZXVzLkNvdW50ZXIKCXVuaXF1ZVJlZnNUb3RhbCBwcm9tZXRoZXVzLkNvdW50ZXIKfQoKZnVuYyBOZXdTZXJpZXNSZWZNYXBwaW5nU3RvcmUocmVnIHByb21ldGhldXMuUmVnaXN0ZXJlcikgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSB7CglhY3RpdmVNYXBwaW5ncyA6PSBwcm9tZXRoZXVzLk5ld0dhdWdlKHByb21ldGhldXMuR2F1Z2VPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV9tYXBwaW5nc190b3RhbCIsCgkJSGVscDogIk51bWJlciBvZiBhY3RpdmUgdW5pcXVlIHJlZiBtYXBwaW5ncyBpbiB0aGUgc3RvcmUuIiwKCX0pCgl0cmFja2VkUmVmcyA6PSBwcm9tZXRoZXVzLk5ld0dhdWdlKHByb21ldGhldXMuR2F1Z2VPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV90cmFja2VkX3JlZnNfdG90YWwiLAoJCUhlbHA6ICJOdW1iZXIgb2YgcmVmcyBiZWluZyB0cmFja2VkIGZvciB0aW1lc3RhbXAtYmFzZWQgY2xlYW51cC4iLAoJfSkKCXJlZnNDbGVhbmVkIDo9IHByb21ldGhldXMuTmV3Q291bnRlcihwcm9tZXRoZXVzLkNvdW50ZXJPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV9yZWZzX2NsZWFuZWRfdG90YWwiLAoJCUhlbHA6ICJUb3RhbCBudW1iZXIgb2Ygc3RhbGUgcmVmcyBjbGVhbmVkIHVwIG92ZXIgdGltZS4iLAoJfSkKCXVuaXF1ZVJlZnNUb3RhbCA6PSBwcm9tZXRoZXVzLk5ld0NvdW50ZXIocHJvbWV0aGV1cy5Db3VudGVyT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfdW5pcXVlX3JlZnNfY3JlYXRlZF90b3RhbCIsCgkJSGVscDogIlRvdGFsIG51bWJlciBvZiB1bmlxdWUgcmVmcyBjcmVhdGVkLiIsCgl9KQoKCWlmIHJlZyAhPSBuaWwgewoJCXJlZy5SZWdpc3RlcihhY3RpdmVNYXBwaW5ncykKCQlyZWcuUmVnaXN0ZXIodHJhY2tlZFJlZnMpCgkJcmVnLlJlZ2lzdGVyKHJlZnNDbGVhbmVkKQoJCXJlZy5SZWdpc3Rlcih1bmlxdWVSZWZzVG90YWwpCgl9CgoJcmV0dXJuICZTZXJpZXNSZWZNYXBwaW5nU3RvcmV7CgkJdW5pcXVlUmVmVG9DaGlsZFJlZnM6IG1ha2UobWFwW3N0b3JhZ2UuU2VyaWVzUmVmXXVuaXFSZWZDaGlsZHJlbiksCgkJbmV4dFVuaXF1ZVJlZjogICAgICAgIDEsCgkJdW5pcXVlUmVmVGltZXN0YW1wczogIG1ha2UobWFwW3N0b3JhZ2UuU2VyaWVzUmVmXWludDY0KSwKCQlsYWJlbEhhc2hUb1VuaXF1ZVJlZjogbWFrZShtYXBbdWludDY0XXN0b3JhZ2UuU2VyaWVzUmVmKSwKCQljZWxsUG9vbDogc3luYy5Qb29sewoJCQlOZXc6IGZ1bmMoKSBhbnkgewoJCQkJcmV0dXJuICZDZWxse1JlZnM6IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMTAwKX0KCQkJfSwKCQl9LAoJCXN0b3BDbGVhbnVwOiAgICAgbWFrZShjaGFuIHN0cnVjdHt9KSwKCQljbGVhbnVwU3RvcHBlZDogIG1ha2UoY2hhbiBzdHJ1Y3R7fSksCgkJYWN0aXZlTWFwcGluZ3M6ICBhY3RpdmVNYXBwaW5ncywKCQl0cmFja2VkUmVmczogICAgIHRyYWNrZWRSZWZzLAoJCXJlZnNDbGVhbmVkOiAgICAgcmVmc0NsZWFuZWQsCgkJdW5pcXVlUmVmc1RvdGFsOiB1bmlxdWVSZWZzVG90YWwsCgl9Cn0KCnR5cGUgQ2VsbCBzdHJ1Y3QgewoJUmVmcyBbXXN0b3JhZ2UuU2VyaWVzUmVmCn0KCi8vIEdldE1hcHBpbmcgcmV0dXJucyBleGlzdGluZyBjaGlsZCByZWYgcmVzdWx0cyBmb3IgdGhlIGdpdmVuIHVuaXF1ZSByZWYgaWYgb25lIGV4aXN0cy4KLy8KLy8gSWYgdGhlIHBhc3NlZCB1bmlxdWVSZWYgaXMgemVybywgdGhlIG1ldGhvZCB3aWxsIGF0dGVtcHQgdG8gZmluZCBhIG1hcHBpbmcgdXNpbmcgcGFzc2VkIGxhYmVscy4KLy8gUmV0dXJucyBuaWwgaWYgbm8gbWFwcGluZyBleGlzdHMuCi8vCi8vIFRoZSByZXR1cm5lZCBzbGljZSBtYXkgYmUgbW9kaWZpZWQgYnkgdGhlIGNhbGxlciwgYnV0IFVwZGF0ZU1hcHBpbmcgbXVzdCBiZSBjYWxsZWQKLy8gYWZ0ZXJ3YXJkcyB0byBwZXJzaXN0IGNoYW5nZXMuIE5vdGUgdGhhdCBjb25jdXJyZW50IGFwcGVuZGVycyBtYXkgcmFjZSB0byB1cGRhdGUgdGhlCi8vIHNhbWUgbWFwcGluZyB3aXRoIGRpZmZlcmVudCB2YWx1ZXMsIHdoaWNoIGlzIHNhZmUgYmVjYXVzZSBzdGFsZSBtYXBwaW5ncyBhcmUgc2VsZi1jb3JyZWN0aW5nIC0KLy8gdXNpbmcgYSBzdGFsZSByZWYgd2lsbCBjYXVzZSB0aGUgY2hpbGQgYXBwZW5kZXIgdG8gcmV0dXJuIGEgbmV3IHJlZiBvbiB0aGUgbmV4dCBhcHBlbmQuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgR2V0TWFwcGluZyh1bmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgW11zdG9yYWdlLlNlcmllc1JlZiB7CglzLnJlZk1hcHBpbmdNdS5STG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5SVW5sb2NrKCkKCglpZiB1bmlxdWVSZWYgPT0gMCB7CgkJLy8gU29tZSBjb25zdW1lcnMgZG9uJ3QgbWVtbyB0aGUgZ2xvYmFsIHJlZi4gVHJ5IHRvIGxvb2t1cCBhIHJlZiBieSBsYWJlbCBoYXNoLgoJCWxhYmVsSGFzaCA6PSBsYmxzLkhhc2goKQoJCWdvdFJlZiwgb2sgOj0gcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltsYWJlbEhhc2hdCgkJaWYgIW9rIHsKCQkJcmV0dXJuIG5pbAoJCX0KCgkJdW5pcXVlUmVmID0gZ290UmVmCgl9CgoJaWYgbWFwcGluZywgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdOyBvayB7CgkJcmV0dXJuICptYXBwaW5nLmNoaWxkUmVmcwoJfQoJcmV0dXJuIG5pbAp9CgovLyBDcmVhdGVNYXBwaW5nIGNyZWF0ZXMgYSBuZXcgdW5pcXVlIHJlZiBtYXBwaW5nIGZvciB0aGUgZ2l2ZW4gY2hpbGQgcmVmIHJlc3VsdHMuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgQ3JlYXRlTWFwcGluZyhyZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgc3RvcmFnZS5TZXJpZXNSZWYgewoJLy8gU3RhcnQgY2xlYW51cCBnb3JvdXRpbmUgb24gZmlyc3QgbWFwcGluZwoJcy5zdGFydFJlZkNsZWFudXAuRG8oZnVuYygpIHsKCQlzLmNsZWFudXBTdGFydGVkLlN0b3JlKHRydWUpCgkJZ28gcy5jbGVhbnVwU3RhbGVSZWZzKCkKCX0pCgoJLy8gU3RvcmUgYSBjb3B5IG9mIHRoZSBjaGlsZCByZWYgcmVzdWx0cyBkaXJlY3RseQoJY2hpbGRSZWZTbGljZSA6PSBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxlbihyZWZSZXN1bHRzKSkKCWNvcHkoY2hpbGRSZWZTbGljZSwgcmVmUmVzdWx0cykKCgkvLyBIYXNoIGxhYmVscyB0byBmb3IgdGhlIGZhbGxiYWNrIGxvb2t1cCB0YWJsZQoJbGFiZWxIYXNoIDo9IGxibHMuSGFzaCgpCgoJcy5yZWZNYXBwaW5nTXUuTG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoKCS8vIENyZWF0ZSBhIG5ldyB1bmlxdWUgcmVmCgl1bmlxdWVSZWYgOj0gcy5uZXh0VW5pcXVlUmVmCglzLm5leHRVbmlxdWVSZWYrKwoKCXMubGFiZWxIYXNoVG9VbmlxdWVSZWZbbGFiZWxIYXNoXSA9IHVuaXF1ZVJlZgoJcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdID0gdW5pcVJlZkNoaWxkcmVuewoJCWNoaWxkUmVmczogJmNoaWxkUmVmU2xpY2UsCgkJbGFiZWxIYXNoOiBsYWJlbEhhc2gsCgl9CgoJcy5hY3RpdmVNYXBwaW5ncy5JbmMoKQoJcy51bmlxdWVSZWZzVG90YWwuSW5jKCkKCglyZXR1cm4gdW5pcXVlUmVmCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgVXBkYXRlTWFwcGluZyh1bmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYsIHJlZlJlc3VsdHMgW11zdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSB7CglpZiB1bmlxdWVSZWYgPT0gMCB7CgkJcmV0dXJuCgl9CgoJY2hpbGRSZWZTbGljZSA6PSBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxlbihyZWZSZXN1bHRzKSkKCWNvcHkoY2hpbGRSZWZTbGljZSwgcmVmUmVzdWx0cykKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJLy8gRW5zdXJlIHRoYXQgbGFiZWwgaGFzaCBpbmRleCBpcyB1cCB0byBkYXRlIHRvIGhhbmRsZSBwb3NzaWJsZSBoYXNoIGNvbGxpc2lvbnMuCgkvLyBUT0RPOiBpcyB0aGlzIG5lY2Vzc2FyeT8KCW5ld0hhc2ggOj0gbGJscy5IYXNoKCkKCXByZXYsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXQoJaWYgb2sgJiYgcHJldi5sYWJlbEhhc2ggIT0gbmV3SGFzaCB7CgkJZGVsZXRlKHMubGFiZWxIYXNoVG9VbmlxdWVSZWYsIHByZXYubGFiZWxIYXNoKQoJCXMubGFiZWxIYXNoVG9VbmlxdWVSZWZbbmV3SGFzaF0gPSB1bmlxdWVSZWYKCX0KCglzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0gPSB1bmlxUmVmQ2hpbGRyZW57CgkJY2hpbGRSZWZzOiAmY2hpbGRSZWZTbGljZSwKCQlsYWJlbEhhc2g6IGxibHMuSGFzaCgpLAoJfQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIFRyYWNrQXBwZW5kZWRTZXJpZXModHMgaW50NjQsIGNlbGwgKkNlbGwpIHsKCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCWRlZmVyIHMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCWZvciBfLCByIDo9IHJhbmdlIGNlbGwuUmVmcyB7CgkJcy51bmlxdWVSZWZUaW1lc3RhbXBzW3JdID0gdHMKCX0KCglzLnRyYWNrZWRSZWZzLlNldChmbG9hdDY0KGxlbihzLnVuaXF1ZVJlZlRpbWVzdGFtcHMpKSkKCgljZWxsLlJlZnMgPSBjZWxsLlJlZnNbOjBdCglzLmNlbGxQb29sLlB1dChjZWxsKQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIEdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpICpDZWxsIHsKCXJldHVybiBzLmNlbGxQb29sLkdldCgpLigqQ2VsbCkKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBjbGVhbnVwU3RhbGVSZWZzKCkgewoJZGVmZXIgY2xvc2Uocy5jbGVhbnVwU3RvcHBlZCkKCgl0aWNrZXIgOj0gdGltZS5OZXdUaWNrZXIoMTUgKiB0aW1lLk1pbnV0ZSkKCWRlZmVyIHRpY2tlci5TdG9wKCkKCglmb3IgewoJCXNlbGVjdCB7CgkJY2FzZSA8LXRpY2tlci5DOgoJCQljdXRvZmZUaW1lIDo9IHRpbWUuTm93KCkuQWRkKC0xNSAqIHRpbWUuTWludXRlKS5Vbml4KCkKCgkJCS8vIEhvbGQgYm90aCBsb2NrcyB0byBwcmV2ZW50IHJhY2UgY29uZGl0aW9uIHdoZXJlIGEgcmVmIGNvdWxkIGJlCgkJCS8vIGFwcGVuZGVkIGFmdGVyIHdlIGRlbGV0ZSBpdCBmcm9tIHVuaXF1ZVJlZkNlbGwgYnV0IGJlZm9yZQoJCQkvLyB3ZSBkZWxldGUgaXQgZnJvbSB1bmlxdWVSZWZUb0NoaWxkUmVmcwoJCQlzLnRpbWVzdGFtcFRyYWNraW5nTXUuTG9jaygpCgkJCXMucmVmTWFwcGluZ011LkxvY2soKQoKCQkJc3RhbGVSZWZDb3VudCA6PSAwCgkJCWZvciByZWYsIHRzIDo9IHJhbmdlIHMudW5pcXVlUmVmVGltZXN0YW1wcyB7CgkJCQlpZiB0cyA8IGN1dG9mZlRpbWUgewoJCQkJCXN0YWxlUmVmQ291bnQrKwoKCQkJCQl2LCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3JlZl0KCQkJCQlpZiBvayB7CgkJCQkJCWRlbGV0ZShzLmxhYmVsSGFzaFRvVW5pcXVlUmVmLCB2LmxhYmVsSGFzaCkKCQkJCQl9CgoJCQkJCWRlbGV0ZShzLnVuaXF1ZVJlZlRpbWVzdGFtcHMsIHJlZikKCQkJCQlkZWxldGUocy51bmlxdWVSZWZUb0NoaWxkUmVmcywgcmVmKQoJCQkJfQoJCQl9CgoJCQkvLyBVcGRhdGUgbWV0cmljcwoJCQlpZiBzdGFsZVJlZkNvdW50ID4gMCB7CgkJCQlzLnJlZnNDbGVhbmVkLkFkZChmbG9hdDY0KHN0YWxlUmVmQ291bnQpKQoJCQkJcy5hY3RpdmVNYXBwaW5ncy5TdWIoZmxvYXQ2NChzdGFsZVJlZkNvdW50KSkKCQkJCXMudHJhY2tlZFJlZnMuU2V0KGZsb2F0NjQobGVuKHMudW5pcXVlUmVmVGltZXN0YW1wcykpKQoJCQl9CgoJCQlzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoJCQlzLnRpbWVzdGFtcFRyYWNraW5nTXUuVW5sb2NrKCkKCgkJY2FzZSA8LXMuc3RvcENsZWFudXA6CgkJCXJldHVybgoJCX0KCX0KfQoKLy8gQ2xlYXIgd2lsbCBjbGVhciBhbGwgaW50ZXJuYWwgbWFwcGluZ3MgYW5kIHN0b3AgdGhlIGNsZWFuZXIgZ29yb3V0aW5lIGlmIGl0IGlzIHJ1bm5pbmcuCi8vIEl0IGlzIHNhZmUgdG8gcmUtdXNlIHRoZSBzYW1lIGluc3RhbmNlIGFmdGVyIGNhbGxpbmcgQ2xlYXIuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgQ2xlYXIoKSB7CgkvLyBTdG9wIHRoZSBjbGVhbnVwIGdvcm91dGluZSBhbmQgd2FpdCBmb3IgaXQgdG8gYmUgc3RvcHBlZCBzbyB3ZSBjYW4KCS8vIGF2b2lkIGEgcG9zc2libGUgZGVhZGxvY2sgd2l0aCBjbGVhbnVwIHRoYXQgYWxzbyBob2xkcyBib3RoIGxvY2tzCglpZiBzLmNsZWFudXBTdGFydGVkLkxvYWQoKSB7CgkJc2VsZWN0IHsKCQljYXNlIDwtcy5zdG9wQ2xlYW51cDoKCQkJLy8gQWxyZWFkeSBjbG9zZWQKCQlkZWZhdWx0OgoJCQljbG9zZShzLnN0b3BDbGVhbnVwKQoJCQk8LXMuY2xlYW51cFN0b3BwZWQKCQl9Cgl9CgoJLy8gV2UgbmVlZCB0byBob2xkIGJvdGggbG9ja3MgdG8gZG8gdGhpcyBzYWZlbHkgYW5kIHdlIGRvIGl0IGluIHRoZSBzYW1lIG9yZGVyIGFzCgkvLyBjbGVhbnVwU3RhbGVSZWZzLiBXZSBzdG9wcGVkIGFuZCB3YWl0ZWQgZm9yIHRoZSBiYWNrZ3JvdW5kIHdvcmtlciB0aGF0IGNhbGxzIGl0CgkvLyB0byBmaW5pc2ggYnV0IHNvbWUgZXh0cmEgc2FmZXR5IHdvbid0IGh1cnQuCglzLnRpbWVzdGFtcFRyYWNraW5nTXUuTG9jaygpCglkZWZlciBzLnRpbWVzdGFtcFRyYWNraW5nTXUuVW5sb2NrKCkKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJY2xlYXIocy51bmlxdWVSZWZUb0NoaWxkUmVmcykKCWNsZWFyKHMudW5pcXVlUmVmVGltZXN0YW1wcykKCgkvLyByZXNldCB0aGUgcG9vbAoJcy5jZWxsUG9vbCA9IHN5bmMuUG9vbHsKCQlOZXc6IGZ1bmMoKSBhbnkgewoJCQlyZXR1cm4gJkNlbGx7UmVmczogbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAwLCAxMDApfQoJCX0sCgl9CgoJLy8gTk9URTogV2UgZG8gTk9UIHJlc2V0IG5leHRVbmlxdWVSZWYgaGVyZS4gUmVzZXR0aW5nIGl0IHdvdWxkIGNhdXNlIHJlZiBjb2xsaXNpb25zCgkvLyB3aXRoIGNvbXBvbmVudHMgbGlrZSBwcm9tZXRoZXVzLnNjcmFwZSB3aGljaCB3aWxsIGtlZXAgcmUtc2VuZGluZyB0aGUgc2FtZSBjYWNoZWQgcmVmcy4KCS8vIFdlIGNvbnRpbnVlIGluY3JlbWVudGluZyB0byBlbnN1cmUgYWxsIHJlZnMgcmVtYWluIHVuaXF1ZSBhY3Jvc3MgdGhlIGxpZmV0aW1lIG9mIHRoZSBwcm9jZXNzLgoKCS8vIFJlc2V0IG1ldHJpY3MKCXMuYWN0aXZlTWFwcGluZ3MuU2V0KDApCglzLnRyYWNrZWRSZWZzLlNldCgwKQoKCS8vIFJlc2V0IGNoYW5uZWxzIGFuZCBmbGFncwoJcy5zdG9wQ2xlYW51cCA9IG1ha2UoY2hhbiBzdHJ1Y3R7fSkKCXMuY2xlYW51cFN0b3BwZWQgPSBtYWtlKGNoYW4gc3RydWN0e30pCglzLnN0YXJ0UmVmQ2xlYW51cCA9IHN5bmMuT25jZXt9CglzLmNsZWFudXBTdGFydGVkLlN0b3JlKGZhbHNlKQp9Cg=="
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/grafana/alloy/internal/component/prometheus/appenders",
    "github.com/grafana/alloy/internal/component/prometheus/appenders [github.com/grafana/alloy/internal/component/prometheus/appenders.test]"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cloud.google.com/go/auth",
      "Name": "auth",
      "PkgPath": "cloud.google.com/go/auth",
      "GoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/auth.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/threelegged.go"
      ],
      "CompiledGoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/auth.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/threelegged.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "cloud.google.com/go/auth/internal": "cloud.google.com/go/auth/internal",
        "cloud.google.com/go/auth/internal/jwt": "cloud.google.com/go/auth/internal/jwt",
        "context": "context",
        "encoding/json": "encoding/json",
        "errors": "errors",
        "fmt": "fmt",
        "github.com/googleapis/gax-go/v2/internallog": "github.com/googleapis/gax-go/v2/internallog",
        "log/slog": "log/slog",
        "mime": "mime",
        "net/http": "net/http",
        "net/url": "net/url",
        "strconv": "strconv",
        "strings": "strings",
        "sync": "sync",
        "time": "time"
      }
    },
    {
      "ID": "cloud.google.com/go/auth/credentials",
      "Name": "credentials",
      "PkgPath": "cloud.google.com/go/auth/credentials",
      "GoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/compute.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/detect.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/doc.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/filetypes.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/selfsignedjwt.go"
      ],
      "CompiledGoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/compute.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/detect.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/doc.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/filetypes.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/selfsignedjwt.go"
      ],
      "Imports": {
        "cloud.google.com/go/auth": "cloud.google.com/go/auth",
        "cloud.google.com/go/auth/credentials/internal/externalaccount": "cloud.google.com/go/auth/credentials/internal/externalaccount",
        "cloud.google.com/go/auth/credentials/internal/externalaccountuser": "cloud.google.com/go/auth/credentials/internal/externalaccountuser",
        "cloud.google.com/go/auth/credentials/internal/gdch": "cloud.google.com/go/auth/credentials/internal/gdch",
        "cloud.google.com/go/auth/credentials/internal/impersonate": "cloud.google.com/go/auth/credentials/internal/impersonate",
        "cloud.google.com/go/auth/internal": "cloud.google.com/go/auth/internal",
        "cloud.google.com/go/auth/internal/credsfile": "cloud.google.com/go/auth/internal/credsfile",
        "cloud.google.com/go/auth/internal/jwt": "cloud.google.com/go/auth/internal/jwt",
        "cloud.google.com/go/auth/internal/trustboundary": "cloud.google.com/go/auth/internal/trustboundary",
        "cloud.google.com/go/compute/metadata": "cloud.google.com/go/compute/metadata",
        "context": "context",
        "crypto": "crypto",
        "encoding/json": "encoding/json",
        "errors": "errors",
        "fmt": "fmt",
        "github.com/googleapis/gax-go/v2/internallog": "github.com/googleapis/gax-go/v2/internallog",
        "log/slog": "log/slog",
        "net/http": "net/http",
        "net/url": "net/url",
        "os": "os",
        "strings": "strings",
        "time": "time"
      }
    },
    {
      "ID": "cloud.google.com/go/auth/credentials/internal/externalaccount",
      "Name": "externalaccount",
      "PkgPath": "cloud.google.com/go/auth/credentials/internal/externalaccount",
      "GoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/aws_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/executable_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/externalaccount.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/file_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/info.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/programmatic_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/url_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/x509_provider.go"
      ],
      "CompiledGoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/aws_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/executable_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/externalaccount.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/file_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/info.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/programmatic_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/url_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/x509_provider.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "cloud.google.com/go/auth": "cloud.google.com/go/auth",
        "cloud.google.com/go/auth/credentials/internal/impersonate": "cloud.google.com/go/auth/credentials/internal/impersonate",
        "cloud.google.com/go/auth/credentials/internal/stsexchange": "cloud.google.com/go/auth/credentials/internal/stsexchange",
        "cloud.google.com/go/auth/internal": "cloud.google.com/go/auth/internal",
        "cloud.google.com/go/auth/internal/credsfile": "cloud.google.com/go/auth/internal/credsfile",
        "cloud.google.com/go/auth/internal/transport/cert": "cloud.google.com/go/auth/internal/transport/cert",
        "context": "context",
        "crypto/hmac": "crypto/hmac",
        "crypto/sha256": "crypto/sha256",
        "crypto/tls": "crypto/tls",
        "crypto/x509": "crypto/x509",
        "encoding/base64": "encoding/base64",
        "encoding/hex": "encoding/hex",
        "encoding/json": "encoding/json",
        "encoding/pem": "encoding/pem",
        "errors": "errors",
        "fmt": "fmt",
        "github.com/googleapis/gax-go/v2/internallog": "github.com/googleapis/gax-go/v2/internallog",
        "io/fs": "io/fs",
        "log/slog": "log/slog",
        "net/http": "net/http",
        "net/url": "net/url",
        "os": "os",
        "os/exec": "os/exec",
        "path": "path",
        "regexp": "regexp",
        "runtime": "runtime",
        "sort": "sort",
        "strconv": "strconv",
        "strings": "strings",
        "time": "time",
        "unicode": "unicode"
      }
    }
  ],
  "GoVersion": 0
}
```

Note: `Packages` truncated from 661 to 5 entries.

#### drv #3

Trace meta: spanId=4, ts=1770837170484, ts_iso=2026-02-11T19:12:50.484000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/appenders",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy/internal/validator"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkiZm10IgoJInN5bmMiCgkic3luYy9hdG9taWMiCgkidGltZSIKCgkiZ2l0aHViLmNvbS9oYXNoaWNvcnAvZ28tbXVsdGllcnJvciIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvY2xpZW50X2dvbGFuZy9wcm9tZXRoZXVzIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2V4ZW1wbGFyIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2hpc3RvZ3JhbSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9sYWJlbHMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvbWV0YWRhdGEiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvc3RvcmFnZSIKKQoKdHlwZSBNYXBwaW5nU3RvcmUgaW50ZXJmYWNlIHsKCUdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYKCUNyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmCglVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpCglUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKQoJR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwKfQoKdHlwZSBzZXJpZXNSZWZNYXBwaW5nIHN0cnVjdCB7CglzdGFydCAgICB0aW1lLlRpbWUKCWNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlcgoJc3RvcmUgICAgTWFwcGluZ1N0b3JlCgoJdW5pcXVlUmVmQ2VsbCAqQ2VsbAoKCS8vIGNoaWxkUmVmcyBpcyByZXVzZWQgZm9yIGVhY2ggYXBwZW5kIGNhbGwgdG8gYXZvaWQgYWxsb2NhdGlvbnMuIFRoaXMgaXMgc2FmZSBiZWNhdXNlIHN0b3JhZ2UuQXBwZW5kZXIgc2hvdWxkIG5ldmVyCgkvLyBoYXZlIGNvbmN1cnJlbnQgY2FsbHMgdG8gQXBwZW5kIG1ldGhvZHMuCgljaGlsZFJlZnMgICAgICAgIFtdc3RvcmFnZS5TZXJpZXNSZWYKCXdyaXRlTGF0ZW5jeSAgICAgcHJvbWV0aGV1cy5IaXN0b2dyYW0KCXNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZyhjaGlsZHJlbiBbXXN0b3JhZ2UuQXBwZW5kZXIsIHN0b3JlIE1hcHBpbmdTdG9yZSwgd3JpdGVMYXRlbmN5IHByb21ldGhldXMuSGlzdG9ncmFtLCBzYW1wbGVzRm9yd2FyZGVkIHByb21ldGhldXMuQ291bnRlcikgc3RvcmFnZS5BcHBlbmRlciB7Cgl1bmlxdWVSZWZDZWxsIDo9IHN0b3JlLkdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpCgoJcmV0dXJuICZzZXJpZXNSZWZNYXBwaW5newoJCWNoaWxkcmVuOiAgICAgICAgIGNoaWxkcmVuLAoJCXN0b3JlOiAgICAgICAgICAgIHN0b3JlLAoJCXdyaXRlTGF0ZW5jeTogICAgIHdyaXRlTGF0ZW5jeSwKCQlzYW1wbGVzRm9yd2FyZGVkOiBzYW1wbGVzRm9yd2FyZGVkLAoKCQl1bmlxdWVSZWZDZWxsOiB1bmlxdWVSZWZDZWxsLAoJCWNoaWxkUmVmczogICAgIG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgbGVuKGNoaWxkcmVuKSksCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFNldE9wdGlvbnMob3B0cyAqc3RvcmFnZS5BcHBlbmRPcHRpb25zKSB7CglmbXQuUHJpbnRsbigpCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQljLlNldE9wdGlvbnMob3B0cykKCX0KfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQ29tbWl0KCkgZXJyb3IgewoJcy5zdG9yZS5UcmFja0FwcGVuZGVkU2VyaWVzKHRpbWUuTm93KCkuVW5peCgpLCBzLnVuaXF1ZVJlZkNlbGwpCgoJdmFyIG11bHRpRXJyIGVycm9yCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQllcnIgOj0gYy5Db21taXQoKQoJCWlmIGVyciAhPSBuaWwgewoJCQltdWx0aUVyciA9IG11bHRpZXJyb3IuQXBwZW5kKG11bHRpRXJyLCBlcnIpCgkJfQoJfQoJcmV0dXJuIG11bHRpRXJyCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFJvbGxiYWNrKCkgZXJyb3IgewoJLy8gV2Ugc3RpbGwgdHJhY2sgcm9sbGVkIGJhY2sgc2VyaWVzIHNvIHdlIGNhbiBwcm9wZXJseQoJLy8gY2xlYW4gdXAgYW55IHNlcmllcyB0aGF0IHdhcyBhcHBlbmRlZAoJcy5zdG9yZS5UcmFja0FwcGVuZGVkU2VyaWVzKHRpbWUuTm93KCkuVW5peCgpLCBzLnVuaXF1ZVJlZkNlbGwpCgoJdmFyIG11bHRpRXJyIGVycm9yCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQllcnIgOj0gYy5Sb2xsYmFjaygpCgkJaWYgZXJyICE9IG5pbCB7CgkJCW11bHRpRXJyID0gbXVsdGllcnJvci5BcHBlbmQobXVsdGlFcnIsIGVycikKCQl9Cgl9CglyZXR1cm4gbXVsdGlFcnIKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgcmVjb3JkTGF0ZW5jeSgpIHsKCWlmIHMuc3RhcnQuSXNaZXJvKCkgewoJCXJldHVybgoJfQoJZHVyYXRpb24gOj0gdGltZS5TaW5jZShzLnN0YXJ0KQoJcy53cml0ZUxhdGVuY3kuT2JzZXJ2ZShkdXJhdGlvbi5TZWNvbmRzKCkpCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIHJlc2V0RmllbGRzKCkgewoJLy8gUmVzZXQgY2hpbGRSZWZzIHNsaWNlIGxlbmd0aCB0byAwIGZvciByZXVzZQoJcy5jaGlsZFJlZnMgPSBzLmNoaWxkUmVmc1s6MF0KfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0IGludDY0LCB2IGZsb2F0NjQpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCW5ld1JlZiwgZXJyIDo9IGFwcGVuZGVyLkFwcGVuZChyZWYsIGwsIHQsIHYpCgkJaWYgZXJyID09IG5pbCB7CgkJCXMuc2FtcGxlc0ZvcndhcmRlZC5JbmMoKQoJCX0KCQlyZXR1cm4gbmV3UmVmLCBlcnIKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEV4ZW1wbGFyKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCBlIGV4ZW1wbGFyLkV4ZW1wbGFyKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kRXhlbXBsYXIocmVmLCBsLCBlKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kSGlzdG9ncmFtKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0IGludDY0LCBoICpoaXN0b2dyYW0uSGlzdG9ncmFtLCBmaCAqaGlzdG9ncmFtLkZsb2F0SGlzdG9ncmFtKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kSGlzdG9ncmFtKHJlZiwgbCwgdCwgaCwgZmgpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRIaXN0b2dyYW1DVFplcm9TYW1wbGUocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQsIGN0IGludDY0LCBoICpoaXN0b2dyYW0uSGlzdG9ncmFtLCBmaCAqaGlzdG9ncmFtLkZsb2F0SGlzdG9ncmFtKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kSGlzdG9ncmFtQ1RaZXJvU2FtcGxlKHJlZiwgbCwgdCwgY3QsIGgsIGZoKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgVXBkYXRlTWV0YWRhdGEocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIG0gbWV0YWRhdGEuTWV0YWRhdGEpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5VcGRhdGVNZXRhZGF0YShyZWYsIGwsIG0pCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRDVFplcm9TYW1wbGUocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQsIGN0IGludDY0KSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kQ1RaZXJvU2FtcGxlKHJlZiwgbCwgdCwgY3QpCgl9KQp9Cgp0eXBlIGFwcGVuZGVyRnVuYyBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikKCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIGFwcGVuZFRvQ2hpbGRyZW4ocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMsIGFmIGFwcGVuZGVyRnVuYykgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJZGVmZXIgcy5yZXNldEZpZWxkcygpCgoJaWYgcy5zdGFydC5Jc1plcm8oKSB7CgkJcy5zdGFydCA9IHRpbWUuTm93KCkKCX0KCgkvLyBDaGVjayBpZiB0aGUgaW5jb21pbmcgcmVmIGhhcyByZWYgbWFwcGluZ3MKCWV4aXN0aW5nQ2hpbGRSZWZzIDo9IHMuc3RvcmUuR2V0TWFwcGluZyhyZWYsIGxibHMpCgoJdmFyIGFwcGVuZEVyciBlcnJvcgoKCS8vIFNhbml0eSBjaGVjazogaWYgd2UgaGF2ZSBleGlzdGluZyBjaGlsZCByZWZzLCB0aGV5IG11c3QgbWF0Y2ggdGhlIG51bWJlciBvZiBjaGlsZHJlbgoJaWYgZXhpc3RpbmdDaGlsZFJlZnMgIT0gbmlsICYmIGxlbihleGlzdGluZ0NoaWxkUmVmcykgPT0gbGVuKHMuY2hpbGRyZW4pIHsKCQlzLnVuaXF1ZVJlZkNlbGwuUmVmcyA9IGFwcGVuZChzLnVuaXF1ZVJlZkNlbGwuUmVmcywgcmVmKQoKCQlyZWZVcGRhdGVSZXF1aXJlZCA6PSBmYWxzZQoJCWZvciBjaGlsZEluZGV4LCBjaGlsZFJlZiA6PSByYW5nZSBleGlzdGluZ0NoaWxkUmVmcyB7CgkJCW5ld0NoaWxkUmVmLCBlcnIgOj0gYWYocy5jaGlsZHJlbltjaGlsZEluZGV4XSwgY2hpbGRSZWYpCgkJCWlmIGVyciAhPSBuaWwgewoJCQkJYXBwZW5kRXJyID0gbXVsdGllcnJvci5BcHBlbmQoYXBwZW5kRXJyLCBlcnIpCgkJCX0KCgkJCWlmIG5ld0NoaWxkUmVmICE9IGNoaWxkUmVmIHsKCQkJCS8vIENoaWxkIHJlZiBjaGFuZ2VkLCBuZWVkIHRvIHVwZGF0ZSBtYXBwaW5nCgkJCQlleGlzdGluZ0NoaWxkUmVmc1tjaGlsZEluZGV4XSA9IG5ld0NoaWxkUmVmCgkJCQlyZWZVcGRhdGVSZXF1aXJlZCA9IHRydWUKCQkJfQoJCX0KCgkJaWYgYXBwZW5kRXJyICE9IG5pbCB7CgkJCXJldHVybiAwLCBhcHBlbmRFcnIKCQl9CgoJCWlmIHJlZlVwZGF0ZVJlcXVpcmVkIHsKCQkJcy5zdG9yZS5VcGRhdGVNYXBwaW5nKHJlZiwgZXhpc3RpbmdDaGlsZFJlZnMsIGxibHMpCgkJfQoKCQlyZXR1cm4gcmVmLCBuaWwKCX0KCgkvLyBObyBleGlzdGluZyBtYXBwaW5nLCBwcm9jZWVkIHdpdGggbm9ybWFsIGFwcGVuZCB0byBhbGwgY2hpbGRyZW4KCXZhciBmaXJzdE5vblplcm9SZWYgc3RvcmFnZS5TZXJpZXNSZWYKCXZhciBub25aZXJvQ291bnQgaW50CgoJLy8gTm90ZTogdGhlcmUncyBhbm90aGVyIG9wdGltaXphdGlvbiB3aGVyZSB3ZSBjb3VsZCB1c2UgdGhlIHJldHVybmVkIHJlZiBpZiBhbGwgdGhlIG5vbiB6ZXJvIHJlZnMKCS8vICBhcmUgdGhlIHNhbWUgdmFsdWUuIFRoaXMgaXNuJ3Qgc2FmZSBhcyB3ZSB3aWxsIG1peCBkb3duc3RyZWFtIHJlZnMgd2l0aCB1bmlxdWUgcmVmcyB3aGljaCBjb3VsZAoJLy8gIGNvbGxpZGUuIFdlIGNvdWxkIHN0YXJ0IGF0IG1heCB1bml0NjQgZm9yIHVuaXF1ZSByZWZzIGFuZCBnbyBiYWNrd2FyZHMgbGVzc2VuaW5nIHRoZSBjaGFuY2Ugb2YKCS8vIAljb2xsaXNpb25zIGJ1dCBpdCdzIHJhdGhlciBkYW5nZXJvdXMgZm9yIGFuIHVubGlrZWx5IGVkZ2UgY2FzZS4gSWYgdHdvIGNvbXBvbmVudHMgYXJlIHJldHVybmluZwoJLy8gCXRoZSBzYW1lIHJlZiBpdCdzIHR3byByZW1vdGVfd3JpdGUgY29tcG9uZW50cyB3aGljaCBzaG91bGQgcHJvYmFibHkgYmUgbWVyZ2VkIGluIHRvIG9uZS4KCWZvciBfLCBjaGlsZCA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQljaGlsZFJlZiwgZXJyIDo9IGFmKGNoaWxkLCByZWYpCgkJaWYgZXJyICE9IG5pbCB7CgkJCWFwcGVuZEVyciA9IG11bHRpZXJyb3IuQXBwZW5kKGFwcGVuZEVyciwgZXJyKQoKCQkJLy8gVE9ETyBzaG91bGQgSSBtYWtlIHRoZSBjaGlsZFJlZiB6ZXJvIGhlcmU/CgkJfQoKCQlzLmNoaWxkUmVmcyA9IGFwcGVuZChzLmNoaWxkUmVmcywgY2hpbGRSZWYpCgkJaWYgY2hpbGRSZWYgIT0gMCB7CgkJCWlmIGZpcnN0Tm9uWmVyb1JlZiA9PSAwIHsKCQkJCWZpcnN0Tm9uWmVyb1JlZiA9IGNoaWxkUmVmCgkJCX0KCQkJbm9uWmVyb0NvdW50KysKCQl9Cgl9CgoJaWYgYXBwZW5kRXJyICE9IG5pbCB7CgkJcmV0dXJuIDAsIGFwcGVuZEVycgoJfQoKCWlmIG5vblplcm9Db3VudCA9PSAwIHsKCQkvLyBBbGwgY2hpbGRyZW4gcmV0dXJuZWQgcmVmIDAsIHNvIHJldHVybiB0aGUgaW5wdXQgcmVmCgkJcmV0dXJuIHJlZiwgbmlsCgl9CgoJLy8gT25seSBvbmUgY2hpbGQgcmV0dXJuZWQgYSBub24temVybyByZWYsIHVzZSB0aGF0CglpZiBub25aZXJvQ291bnQgPT0gMSB7CgkJcmV0dXJuIGZpcnN0Tm9uWmVyb1JlZiwgbmlsCgl9CgoJLy8gV2UgZ290IGRpZmZlcmVudCByZWZzIGJhY2sgYW5kIG5lZWQgdG8gY3JlYXRlIGEgbmV3IG1hcHBpbmcKCXVuaXF1ZVJlZiA6PSBzLnN0b3JlLkNyZWF0ZU1hcHBpbmcocy5jaGlsZFJlZnMsIGxibHMpCglzLnVuaXF1ZVJlZkNlbGwuUmVmcyA9IGFwcGVuZChzLnVuaXF1ZVJlZkNlbGwuUmVmcywgdW5pcXVlUmVmKQoJcmV0dXJuIHVuaXF1ZVJlZiwgbmlsCn0KCnR5cGUgdW5pcVJlZkNoaWxkcmVuIHN0cnVjdCB7CgljaGlsZFJlZnMgKltdc3RvcmFnZS5TZXJpZXNSZWYKCWxhYmVsSGFzaCB1aW50NjQKfQoKdHlwZSBTZXJpZXNSZWZNYXBwaW5nU3RvcmUgc3RydWN0IHsKCS8vIHJlZk1hcHBpbmdNdSBwcm90ZWN0cyB1bmlxdWVSZWZUb0NoaWxkUmVmcywgbGFiZWxIYXNoVG9VbmlxdWVSZWYgYW5kIG5leHRVbmlxdWVSZWYKCXJlZk1hcHBpbmdNdSBzeW5jLlJXTXV0ZXgKCS8vIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzIG1hcHMgdGhlIHVuaXF1ZSByZWYgdG8gdGhlIGV4cGVjdGVkIGNoaWxkIHJlZiBpbiBvcmRlcgoJdW5pcXVlUmVmVG9DaGlsZFJlZnMgbWFwW3N0b3JhZ2UuU2VyaWVzUmVmXXVuaXFSZWZDaGlsZHJlbgoJLy8gbGFiZWxIYXNoVG9VbmlxdWVSZWYgbWFwcyB0aGUgbGFiZWwgaGFzaCB0byB1bmlxdWUgcmVmLgoJbGFiZWxIYXNoVG9VbmlxdWVSZWYgbWFwW3VpbnQ2NF1zdG9yYWdlLlNlcmllc1JlZgoKCS8vIG5leHRVbmlxdWVSZWYgaXMgdGhlIG5leHQgcmVmIElEIHdlIHdpbGwgaGFuZCBvdXQKCW5leHRVbmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYKCgkvLyB0aW1lc3RhbXBUcmFja2luZ011IHByb3RlY3RzIHVuaXF1ZVJlZlRpbWVzdGFtcHMgYW5kIGNlbGxQb29sCgl0aW1lc3RhbXBUcmFja2luZ011IHN5bmMuTXV0ZXgKCS8vIHVuaXF1ZVJlZlRpbWVzdGFtcHMgbWFwcyB1bmlxdWUgcmVmcyB0byB0aGVpciBsYXN0IGFwcGVuZCB0aW1lc3RhbXAKCXVuaXF1ZVJlZlRpbWVzdGFtcHMgbWFwW3N0b3JhZ2UuU2VyaWVzUmVmXWludDY0CgkvLyBjZWxsUG9vbCBpcyB1c2VkIHRvIHBvb2wgc2xpY2VzIG9mIFNlcmllc1JlZnMgdXNlZCBmb3IgdHJhY2tpbmcgdW5pcXVlIHJlZnMgaW4gVHJhY2tBcHBlbmRlZFNlcmllcy4KCWNlbGxQb29sIHN5bmMuUG9vbAoKCS8vIENsZWFudXAgZ29yb3V0aW5lIGNvb3JkaW5hdGlvbiAobm8gbG9jayByZXF1aXJlZCkKCXN0YXJ0UmVmQ2xlYW51cCBzeW5jLk9uY2UKCWNsZWFudXBTdGFydGVkICBhdG9taWMuQm9vbAoJc3RvcENsZWFudXAgICAgIGNoYW4gc3RydWN0e30KCWNsZWFudXBTdG9wcGVkICBjaGFuIHN0cnVjdHt9CgoJLy8gTWV0cmljcyAoc2FmZSBmb3IgY29uY3VycmVudCBhY2Nlc3MsIG5vIGxvY2sgcmVxdWlyZWQpCglhY3RpdmVNYXBwaW5ncyAgcHJvbWV0aGV1cy5HYXVnZQoJdHJhY2tlZFJlZnMgICAgIHByb21ldGhldXMuR2F1Z2UKCXJlZnNDbGVhbmVkICAgICBwcm9tZXRoZXVzLkNvdW50ZXIKCXVuaXF1ZVJlZnNUb3RhbCBwcm9tZXRoZXVzLkNvdW50ZXIKfQoKZnVuYyBOZXdTZXJpZXNSZWZNYXBwaW5nU3RvcmUocmVnIHByb21ldGhldXMuUmVnaXN0ZXJlcikgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSB7CglhY3RpdmVNYXBwaW5ncyA6PSBwcm9tZXRoZXVzLk5ld0dhdWdlKHByb21ldGhldXMuR2F1Z2VPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV9tYXBwaW5nc190b3RhbCIsCgkJSGVscDogIk51bWJlciBvZiBhY3RpdmUgdW5pcXVlIHJlZiBtYXBwaW5ncyBpbiB0aGUgc3RvcmUuIiwKCX0pCgl0cmFja2VkUmVmcyA6PSBwcm9tZXRoZXVzLk5ld0dhdWdlKHByb21ldGhldXMuR2F1Z2VPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV90cmFja2VkX3JlZnNfdG90YWwiLAoJCUhlbHA6ICJOdW1iZXIgb2YgcmVmcyBiZWluZyB0cmFja2VkIGZvciB0aW1lc3RhbXAtYmFzZWQgY2xlYW51cC4iLAoJfSkKCXJlZnNDbGVhbmVkIDo9IHByb21ldGhldXMuTmV3Q291bnRlcihwcm9tZXRoZXVzLkNvdW50ZXJPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV9yZWZzX2NsZWFuZWRfdG90YWwiLAoJCUhlbHA6ICJUb3RhbCBudW1iZXIgb2Ygc3RhbGUgcmVmcyBjbGVhbmVkIHVwIG92ZXIgdGltZS4iLAoJfSkKCXVuaXF1ZVJlZnNUb3RhbCA6PSBwcm9tZXRoZXVzLk5ld0NvdW50ZXIocHJvbWV0aGV1cy5Db3VudGVyT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfdW5pcXVlX3JlZnNfY3JlYXRlZF90b3RhbCIsCgkJSGVscDogIlRvdGFsIG51bWJlciBvZiB1bmlxdWUgcmVmcyBjcmVhdGVkLiIsCgl9KQoKCWlmIHJlZyAhPSBuaWwgewoJCXJlZy5SZWdpc3RlcihhY3RpdmVNYXBwaW5ncykKCQlyZWcuUmVnaXN0ZXIodHJhY2tlZFJlZnMpCgkJcmVnLlJlZ2lzdGVyKHJlZnNDbGVhbmVkKQoJCXJlZy5SZWdpc3Rlcih1bmlxdWVSZWZzVG90YWwpCgl9CgoJcmV0dXJuICZTZXJpZXNSZWZNYXBwaW5nU3RvcmV7CgkJdW5pcXVlUmVmVG9DaGlsZFJlZnM6IG1ha2UobWFwW3N0b3JhZ2UuU2VyaWVzUmVmXXVuaXFSZWZDaGlsZHJlbiksCgkJbmV4dFVuaXF1ZVJlZjogICAgICAgIDEsCgkJdW5pcXVlUmVmVGltZXN0YW1wczogIG1ha2UobWFwW3N0b3JhZ2UuU2VyaWVzUmVmXWludDY0KSwKCQlsYWJlbEhhc2hUb1VuaXF1ZVJlZjogbWFrZShtYXBbdWludDY0XXN0b3JhZ2UuU2VyaWVzUmVmKSwKCQljZWxsUG9vbDogc3luYy5Qb29sewoJCQlOZXc6IGZ1bmMoKSBhbnkgewoJCQkJcmV0dXJuICZDZWxse1JlZnM6IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMTAwKX0KCQkJfSwKCQl9LAoJCXN0b3BDbGVhbnVwOiAgICAgbWFrZShjaGFuIHN0cnVjdHt9KSwKCQljbGVhbnVwU3RvcHBlZDogIG1ha2UoY2hhbiBzdHJ1Y3R7fSksCgkJYWN0aXZlTWFwcGluZ3M6ICBhY3RpdmVNYXBwaW5ncywKCQl0cmFja2VkUmVmczogICAgIHRyYWNrZWRSZWZzLAoJCXJlZnNDbGVhbmVkOiAgICAgcmVmc0NsZWFuZWQsCgkJdW5pcXVlUmVmc1RvdGFsOiB1bmlxdWVSZWZzVG90YWwsCgl9Cn0KCnR5cGUgQ2VsbCBzdHJ1Y3QgewoJUmVmcyBbXXN0b3JhZ2UuU2VyaWVzUmVmCn0KCi8vIEdldE1hcHBpbmcgcmV0dXJucyBleGlzdGluZyBjaGlsZCByZWYgcmVzdWx0cyBmb3IgdGhlIGdpdmVuIHVuaXF1ZSByZWYgaWYgb25lIGV4aXN0cy4KLy8KLy8gSWYgdGhlIHBhc3NlZCB1bmlxdWVSZWYgaXMgemVybywgdGhlIG1ldGhvZCB3aWxsIGF0dGVtcHQgdG8gZmluZCBhIG1hcHBpbmcgdXNpbmcgcGFzc2VkIGxhYmVscy4KLy8gUmV0dXJucyBuaWwgaWYgbm8gbWFwcGluZyBleGlzdHMuCi8vCi8vIFRoZSByZXR1cm5lZCBzbGljZSBtYXkgYmUgbW9kaWZpZWQgYnkgdGhlIGNhbGxlciwgYnV0IFVwZGF0ZU1hcHBpbmcgbXVzdCBiZSBjYWxsZWQKLy8gYWZ0ZXJ3YXJkcyB0byBwZXJzaXN0IGNoYW5nZXMuIE5vdGUgdGhhdCBjb25jdXJyZW50IGFwcGVuZGVycyBtYXkgcmFjZSB0byB1cGRhdGUgdGhlCi8vIHNhbWUgbWFwcGluZyB3aXRoIGRpZmZlcmVudCB2YWx1ZXMsIHdoaWNoIGlzIHNhZmUgYmVjYXVzZSBzdGFsZSBtYXBwaW5ncyBhcmUgc2VsZi1jb3JyZWN0aW5nIC0KLy8gdXNpbmcgYSBzdGFsZSByZWYgd2lsbCBjYXVzZSB0aGUgY2hpbGQgYXBwZW5kZXIgdG8gcmV0dXJuIGEgbmV3IHJlZiBvbiB0aGUgbmV4dCBhcHBlbmQuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgR2V0TWFwcGluZyh1bmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgW11zdG9yYWdlLlNlcmllc1JlZiB7CglzLnJlZk1hcHBpbmdNdS5STG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5SVW5sb2NrKCkKCglpZiB1bmlxdWVSZWYgPT0gMCB7CgkJLy8gU29tZSBjb25zdW1lcnMgZG9uJ3QgbWVtbyB0aGUgZ2xvYmFsIHJlZi4gVHJ5IHRvIGxvb2t1cCBhIHJlZiBieSBsYWJlbCBoYXNoLgoJCWxhYmVsSGFzaCA6PSBsYmxzLkhhc2goKQoJCWdvdFJlZiwgb2sgOj0gcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltsYWJlbEhhc2hdCgkJaWYgIW9rIHsKCQkJcmV0dXJuIG5pbAoJCX0KCgkJdW5pcXVlUmVmID0gZ290UmVmCgl9CgoJaWYgbWFwcGluZywgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdOyBvayB7CgkJcmV0dXJuICptYXBwaW5nLmNoaWxkUmVmcwoJfQoJcmV0dXJuIG5pbAp9CgovLyBDcmVhdGVNYXBwaW5nIGNyZWF0ZXMgYSBuZXcgdW5pcXVlIHJlZiBtYXBwaW5nIGZvciB0aGUgZ2l2ZW4gY2hpbGQgcmVmIHJlc3VsdHMuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgQ3JlYXRlTWFwcGluZyhyZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgc3RvcmFnZS5TZXJpZXNSZWYgewoJLy8gU3RhcnQgY2xlYW51cCBnb3JvdXRpbmUgb24gZmlyc3QgbWFwcGluZwoJcy5zdGFydFJlZkNsZWFudXAuRG8oZnVuYygpIHsKCQlzLmNsZWFudXBTdGFydGVkLlN0b3JlKHRydWUpCgkJZ28gcy5jbGVhbnVwU3RhbGVSZWZzKCkKCX0pCgoJLy8gU3RvcmUgYSBjb3B5IG9mIHRoZSBjaGlsZCByZWYgcmVzdWx0cyBkaXJlY3RseQoJY2hpbGRSZWZTbGljZSA6PSBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxlbihyZWZSZXN1bHRzKSkKCWNvcHkoY2hpbGRSZWZTbGljZSwgcmVmUmVzdWx0cykKCgkvLyBIYXNoIGxhYmVscyB0byBmb3IgdGhlIGZhbGxiYWNrIGxvb2t1cCB0YWJsZQoJbGFiZWxIYXNoIDo9IGxibHMuSGFzaCgpCgoJcy5yZWZNYXBwaW5nTXUuTG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoKCS8vIENyZWF0ZSBhIG5ldyB1bmlxdWUgcmVmCgl1bmlxdWVSZWYgOj0gcy5uZXh0VW5pcXVlUmVmCglzLm5leHRVbmlxdWVSZWYrKwoKCXMubGFiZWxIYXNoVG9VbmlxdWVSZWZbbGFiZWxIYXNoXSA9IHVuaXF1ZVJlZgoJcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdID0gdW5pcVJlZkNoaWxkcmVuewoJCWNoaWxkUmVmczogJmNoaWxkUmVmU2xpY2UsCgkJbGFiZWxIYXNoOiBsYWJlbEhhc2gsCgl9CgoJcy5hY3RpdmVNYXBwaW5ncy5JbmMoKQoJcy51bmlxdWVSZWZzVG90YWwuSW5jKCkKCglyZXR1cm4gdW5pcXVlUmVmCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgVXBkYXRlTWFwcGluZyh1bmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYsIHJlZlJlc3VsdHMgW11zdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSB7CglpZiB1bmlxdWVSZWYgPT0gMCB7CgkJcmV0dXJuCgl9CgoJY2hpbGRSZWZTbGljZSA6PSBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxlbihyZWZSZXN1bHRzKSkKCWNvcHkoY2hpbGRSZWZTbGljZSwgcmVmUmVzdWx0cykKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJLy8gRW5zdXJlIHRoYXQgbGFiZWwgaGFzaCBpbmRleCBpcyB1cCB0byBkYXRlIHRvIGhhbmRsZSBwb3NzaWJsZSBoYXNoIGNvbGxpc2lvbnMuCgkvLyBUT0RPOiBpcyB0aGlzIG5lY2Vzc2FyeT8KCW5ld0hhc2ggOj0gbGJscy5IYXNoKCkKCXByZXYsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXQoJaWYgb2sgJiYgcHJldi5sYWJlbEhhc2ggIT0gbmV3SGFzaCB7CgkJZGVsZXRlKHMubGFiZWxIYXNoVG9VbmlxdWVSZWYsIHByZXYubGFiZWxIYXNoKQoJCXMubGFiZWxIYXNoVG9VbmlxdWVSZWZbbmV3SGFzaF0gPSB1bmlxdWVSZWYKCX0KCglzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0gPSB1bmlxUmVmQ2hpbGRyZW57CgkJY2hpbGRSZWZzOiAmY2hpbGRSZWZTbGljZSwKCQlsYWJlbEhhc2g6IGxibHMuSGFzaCgpLAoJfQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIFRyYWNrQXBwZW5kZWRTZXJpZXModHMgaW50NjQsIGNlbGwgKkNlbGwpIHsKCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCWRlZmVyIHMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCWZvciBfLCByIDo9IHJhbmdlIGNlbGwuUmVmcyB7CgkJcy51bmlxdWVSZWZUaW1lc3RhbXBzW3JdID0gdHMKCX0KCglzLnRyYWNrZWRSZWZzLlNldChmbG9hdDY0KGxlbihzLnVuaXF1ZVJlZlRpbWVzdGFtcHMpKSkKCgljZWxsLlJlZnMgPSBjZWxsLlJlZnNbOjBdCglzLmNlbGxQb29sLlB1dChjZWxsKQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIEdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpICpDZWxsIHsKCXJldHVybiBzLmNlbGxQb29sLkdldCgpLigqQ2VsbCkKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBjbGVhbnVwU3RhbGVSZWZzKCkgewoJZGVmZXIgY2xvc2Uocy5jbGVhbnVwU3RvcHBlZCkKCgl0aWNrZXIgOj0gdGltZS5OZXdUaWNrZXIoMTUgKiB0aW1lLk1pbnV0ZSkKCWRlZmVyIHRpY2tlci5TdG9wKCkKCglmb3IgewoJCXNlbGVjdCB7CgkJY2FzZSA8LXRpY2tlci5DOgoJCQljdXRvZmZUaW1lIDo9IHRpbWUuTm93KCkuQWRkKC0xNSAqIHRpbWUuTWludXRlKS5Vbml4KCkKCgkJCS8vIEhvbGQgYm90aCBsb2NrcyB0byBwcmV2ZW50IHJhY2UgY29uZGl0aW9uIHdoZXJlIGEgcmVmIGNvdWxkIGJlCgkJCS8vIGFwcGVuZGVkIGFmdGVyIHdlIGRlbGV0ZSBpdCBmcm9tIHVuaXF1ZVJlZkNlbGwgYnV0IGJlZm9yZQoJCQkvLyB3ZSBkZWxldGUgaXQgZnJvbSB1bmlxdWVSZWZUb0NoaWxkUmVmcwoJCQlzLnRpbWVzdGFtcFRyYWNraW5nTXUuTG9jaygpCgkJCXMucmVmTWFwcGluZ011LkxvY2soKQoKCQkJc3RhbGVSZWZDb3VudCA6PSAwCgkJCWZvciByZWYsIHRzIDo9IHJhbmdlIHMudW5pcXVlUmVmVGltZXN0YW1wcyB7CgkJCQlpZiB0cyA8IGN1dG9mZlRpbWUgewoJCQkJCXN0YWxlUmVmQ291bnQrKwoKCQkJCQl2LCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3JlZl0KCQkJCQlpZiBvayB7CgkJCQkJCWRlbGV0ZShzLmxhYmVsSGFzaFRvVW5pcXVlUmVmLCB2LmxhYmVsSGFzaCkKCQkJCQl9CgoJCQkJCWRlbGV0ZShzLnVuaXF1ZVJlZlRpbWVzdGFtcHMsIHJlZikKCQkJCQlkZWxldGUocy51bmlxdWVSZWZUb0NoaWxkUmVmcywgcmVmKQoJCQkJfQoJCQl9CgoJCQkvLyBVcGRhdGUgbWV0cmljcwoJCQlpZiBzdGFsZVJlZkNvdW50ID4gMCB7CgkJCQlzLnJlZnNDbGVhbmVkLkFkZChmbG9hdDY0KHN0YWxlUmVmQ291bnQpKQoJCQkJcy5hY3RpdmVNYXBwaW5ncy5TdWIoZmxvYXQ2NChzdGFsZVJlZkNvdW50KSkKCQkJCXMudHJhY2tlZFJlZnMuU2V0KGZsb2F0NjQobGVuKHMudW5pcXVlUmVmVGltZXN0YW1wcykpKQoJCQl9CgoJCQlzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoJCQlzLnRpbWVzdGFtcFRyYWNraW5nTXUuVW5sb2NrKCkKCgkJY2FzZSA8LXMuc3RvcENsZWFudXA6CgkJCXJldHVybgoJCX0KCX0KfQoKLy8gQ2xlYXIgd2lsbCBjbGVhciBhbGwgaW50ZXJuYWwgbWFwcGluZ3MgYW5kIHN0b3AgdGhlIGNsZWFuZXIgZ29yb3V0aW5lIGlmIGl0IGlzIHJ1bm5pbmcuCi8vIEl0IGlzIHNhZmUgdG8gcmUtdXNlIHRoZSBzYW1lIGluc3RhbmNlIGFmdGVyIGNhbGxpbmcgQ2xlYXIuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgQ2xlYXIoKSB7CgkvLyBTdG9wIHRoZSBjbGVhbnVwIGdvcm91dGluZSBhbmQgd2FpdCBmb3IgaXQgdG8gYmUgc3RvcHBlZCBzbyB3ZSBjYW4KCS8vIGF2b2lkIGEgcG9zc2libGUgZGVhZGxvY2sgd2l0aCBjbGVhbnVwIHRoYXQgYWxzbyBob2xkcyBib3RoIGxvY2tzCglpZiBzLmNsZWFudXBTdGFydGVkLkxvYWQoKSB7CgkJc2VsZWN0IHsKCQljYXNlIDwtcy5zdG9wQ2xlYW51cDoKCQkJLy8gQWxyZWFkeSBjbG9zZWQKCQlkZWZhdWx0OgoJCQljbG9zZShzLnN0b3BDbGVhbnVwKQoJCQk8LXMuY2xlYW51cFN0b3BwZWQKCQl9Cgl9CgoJLy8gV2UgbmVlZCB0byBob2xkIGJvdGggbG9ja3MgdG8gZG8gdGhpcyBzYWZlbHkgYW5kIHdlIGRvIGl0IGluIHRoZSBzYW1lIG9yZGVyIGFzCgkvLyBjbGVhbnVwU3RhbGVSZWZzLiBXZSBzdG9wcGVkIGFuZCB3YWl0ZWQgZm9yIHRoZSBiYWNrZ3JvdW5kIHdvcmtlciB0aGF0IGNhbGxzIGl0CgkvLyB0byBmaW5pc2ggYnV0IHNvbWUgZXh0cmEgc2FmZXR5IHdvbid0IGh1cnQuCglzLnRpbWVzdGFtcFRyYWNraW5nTXUuTG9jaygpCglkZWZlciBzLnRpbWVzdGFtcFRyYWNraW5nTXUuVW5sb2NrKCkKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJY2xlYXIocy51bmlxdWVSZWZUb0NoaWxkUmVmcykKCWNsZWFyKHMudW5pcXVlUmVmVGltZXN0YW1wcykKCgkvLyByZXNldCB0aGUgcG9vbAoJcy5jZWxsUG9vbCA9IHN5bmMuUG9vbHsKCQlOZXc6IGZ1bmMoKSBhbnkgewoJCQlyZXR1cm4gJkNlbGx7UmVmczogbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAwLCAxMDApfQoJCX0sCgl9CgoJLy8gTk9URTogV2UgZG8gTk9UIHJlc2V0IG5leHRVbmlxdWVSZWYgaGVyZS4gUmVzZXR0aW5nIGl0IHdvdWxkIGNhdXNlIHJlZiBjb2xsaXNpb25zCgkvLyB3aXRoIGNvbXBvbmVudHMgbGlrZSBwcm9tZXRoZXVzLnNjcmFwZSB3aGljaCB3aWxsIGtlZXAgcmUtc2VuZGluZyB0aGUgc2FtZSBjYWNoZWQgcmVmcy4KCS8vIFdlIGNvbnRpbnVlIGluY3JlbWVudGluZyB0byBlbnN1cmUgYWxsIHJlZnMgcmVtYWluIHVuaXF1ZSBhY3Jvc3MgdGhlIGxpZmV0aW1lIG9mIHRoZSBwcm9jZXNzLgoKCS8vIFJlc2V0IG1ldHJpY3MKCXMuYWN0aXZlTWFwcGluZ3MuU2V0KDApCglzLnRyYWNrZWRSZWZzLlNldCgwKQoKCS8vIFJlc2V0IGNoYW5uZWxzIGFuZCBmbGFncwoJcy5zdG9wQ2xlYW51cCA9IG1ha2UoY2hhbiBzdHJ1Y3R7fSkKCXMuY2xlYW51cFN0b3BwZWQgPSBtYWtlKGNoYW4gc3RydWN0e30pCglzLnN0YXJ0UmVmQ2xlYW51cCA9IHN5bmMuT25jZXt9CglzLmNsZWFudXBTdGFydGVkLlN0b3JlKGZhbHNlKQp9Cg=="
    }
  }
}
```

Response (packages.DriverResponse):
Error: `err: context canceled: stderr: `

#### drv #4

Trace meta: spanId=7, ts=1770837170931, ts_iso=2026-02-11T19:12:50.931000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy/internal/validator"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkiZm10IgoJInN5bmMiCgkic3luYy9hdG9taWMiCgkidGltZSIKCgkiZ2l0aHViLmNvbS9oYXNoaWNvcnAvZ28tbXVsdGllcnJvciIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvY2xpZW50X2dvbGFuZy9wcm9tZXRoZXVzIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2V4ZW1wbGFyIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2hpc3RvZ3JhbSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9sYWJlbHMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvbWV0YWRhdGEiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvc3RvcmFnZSIKKQoKdHlwZSBNYXBwaW5nU3RvcmUgaW50ZXJmYWNlIHsKCUdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYKCUNyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmCglVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpCglUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKQoJR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwKfQoKdHlwZSBzZXJpZXNSZWZNYXBwaW5nIHN0cnVjdCB7CglzdGFydCAgICB0aW1lLlRpbWUKCWNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlcgoJc3RvcmUgICAgTWFwcGluZ1N0b3JlCgoJdW5pcXVlUmVmQ2VsbCAqQ2VsbAoKCS8vIGNoaWxkUmVmcyBpcyByZXVzZWQgZm9yIGVhY2ggYXBwZW5kIGNhbGwgdG8gYXZvaWQgYWxsb2NhdGlvbnMuIFRoaXMgaXMgc2FmZSBiZWNhdXNlIHN0b3JhZ2UuQXBwZW5kZXIgc2hvdWxkIG5ldmVyCgkvLyBoYXZlIGNvbmN1cnJlbnQgY2FsbHMgdG8gQXBwZW5kIG1ldGhvZHMuCgljaGlsZFJlZnMgICAgICAgIFtdc3RvcmFnZS5TZXJpZXNSZWYKCXdyaXRlTGF0ZW5jeSAgICAgcHJvbWV0aGV1cy5IaXN0b2dyYW0KCXNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZyhjaGlsZHJlbiBbXXN0b3JhZ2UuQXBwZW5kZXIsIHN0b3JlIE1hcHBpbmdTdG9yZSwgd3JpdGVMYXRlbmN5IHByb21ldGhldXMuSGlzdG9ncmFtLCBzYW1wbGVzRm9yd2FyZGVkIHByb21ldGhldXMuQ291bnRlcikgc3RvcmFnZS5BcHBlbmRlciB7Cgl1bmlxdWVSZWZDZWxsIDo9IHN0b3JlLkdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpCgoJcmV0dXJuICZzZXJpZXNSZWZNYXBwaW5newoJCWNoaWxkcmVuOiAgICAgICAgIGNoaWxkcmVuLAoJCXN0b3JlOiAgICAgICAgICAgIHN0b3JlLAoJCXdyaXRlTGF0ZW5jeTogICAgIHdyaXRlTGF0ZW5jeSwKCQlzYW1wbGVzRm9yd2FyZGVkOiBzYW1wbGVzRm9yd2FyZGVkLAoKCQl1bmlxdWVSZWZDZWxsOiB1bmlxdWVSZWZDZWxsLAoJCWNoaWxkUmVmczogICAgIG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgbGVuKGNoaWxkcmVuKSksCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFNldE9wdGlvbnMob3B0cyAqc3RvcmFnZS5BcHBlbmRPcHRpb25zKSB7CglmbXQuUHJpbnRsbigiKQoJZm9yIF8sIGMgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJYy5TZXRPcHRpb25zKG9wdHMpCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIENvbW1pdCgpIGVycm9yIHsKCXMuc3RvcmUuVHJhY2tBcHBlbmRlZFNlcmllcyh0aW1lLk5vdygpLlVuaXgoKSwgcy51bmlxdWVSZWZDZWxsKQoKCXZhciBtdWx0aUVyciBlcnJvcgoJZm9yIF8sIGMgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJZXJyIDo9IGMuQ29tbWl0KCkKCQlpZiBlcnIgIT0gbmlsIHsKCQkJbXVsdGlFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChtdWx0aUVyciwgZXJyKQoJCX0KCX0KCXJldHVybiBtdWx0aUVycgp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBSb2xsYmFjaygpIGVycm9yIHsKCS8vIFdlIHN0aWxsIHRyYWNrIHJvbGxlZCBiYWNrIHNlcmllcyBzbyB3ZSBjYW4gcHJvcGVybHkKCS8vIGNsZWFuIHVwIGFueSBzZXJpZXMgdGhhdCB3YXMgYXBwZW5kZWQKCXMuc3RvcmUuVHJhY2tBcHBlbmRlZFNlcmllcyh0aW1lLk5vdygpLlVuaXgoKSwgcy51bmlxdWVSZWZDZWxsKQoKCXZhciBtdWx0aUVyciBlcnJvcgoJZm9yIF8sIGMgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJZXJyIDo9IGMuUm9sbGJhY2soKQoJCWlmIGVyciAhPSBuaWwgewoJCQltdWx0aUVyciA9IG11bHRpZXJyb3IuQXBwZW5kKG11bHRpRXJyLCBlcnIpCgkJfQoJfQoJcmV0dXJuIG11bHRpRXJyCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIHJlY29yZExhdGVuY3koKSB7CglpZiBzLnN0YXJ0LklzWmVybygpIHsKCQlyZXR1cm4KCX0KCWR1cmF0aW9uIDo9IHRpbWUuU2luY2Uocy5zdGFydCkKCXMud3JpdGVMYXRlbmN5Lk9ic2VydmUoZHVyYXRpb24uU2Vjb25kcygpKQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSByZXNldEZpZWxkcygpIHsKCS8vIFJlc2V0IGNoaWxkUmVmcyBzbGljZSBsZW5ndGggdG8gMCBmb3IgcmV1c2UKCXMuY2hpbGRSZWZzID0gcy5jaGlsZFJlZnNbOjBdCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZChyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCBpbnQ2NCwgdiBmbG9hdDY0KSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQluZXdSZWYsIGVyciA6PSBhcHBlbmRlci5BcHBlbmQocmVmLCBsLCB0LCB2KQoJCWlmIGVyciA9PSBuaWwgewoJCQlzLnNhbXBsZXNGb3J3YXJkZWQuSW5jKCkKCQl9CgkJcmV0dXJuIG5ld1JlZiwgZXJyCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRFeGVtcGxhcihyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgZSBleGVtcGxhci5FeGVtcGxhcikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZEV4ZW1wbGFyKHJlZiwgbCwgZSkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEhpc3RvZ3JhbShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCBpbnQ2NCwgaCAqaGlzdG9ncmFtLkhpc3RvZ3JhbSwgZmggKmhpc3RvZ3JhbS5GbG9hdEhpc3RvZ3JhbSkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZEhpc3RvZ3JhbShyZWYsIGwsIHQsIGgsIGZoKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kSGlzdG9ncmFtQ1RaZXJvU2FtcGxlKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0LCBjdCBpbnQ2NCwgaCAqaGlzdG9ncmFtLkhpc3RvZ3JhbSwgZmggKmhpc3RvZ3JhbS5GbG9hdEhpc3RvZ3JhbSkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZEhpc3RvZ3JhbUNUWmVyb1NhbXBsZShyZWYsIGwsIHQsIGN0LCBoLCBmaCkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFVwZGF0ZU1ldGFkYXRhKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCBtIG1ldGFkYXRhLk1ldGFkYXRhKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuVXBkYXRlTWV0YWRhdGEocmVmLCBsLCBtKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kQ1RaZXJvU2FtcGxlKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0LCBjdCBpbnQ2NCkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZENUWmVyb1NhbXBsZShyZWYsIGwsIHQsIGN0KQoJfSkKfQoKdHlwZSBhcHBlbmRlckZ1bmMgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpCgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBhcHBlbmRUb0NoaWxkcmVuKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzLCBhZiBhcHBlbmRlckZ1bmMpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCWRlZmVyIHMucmVzZXRGaWVsZHMoKQoKCWlmIHMuc3RhcnQuSXNaZXJvKCkgewoJCXMuc3RhcnQgPSB0aW1lLk5vdygpCgl9CgoJLy8gQ2hlY2sgaWYgdGhlIGluY29taW5nIHJlZiBoYXMgcmVmIG1hcHBpbmdzCglleGlzdGluZ0NoaWxkUmVmcyA6PSBzLnN0b3JlLkdldE1hcHBpbmcocmVmLCBsYmxzKQoKCXZhciBhcHBlbmRFcnIgZXJyb3IKCgkvLyBTYW5pdHkgY2hlY2s6IGlmIHdlIGhhdmUgZXhpc3RpbmcgY2hpbGQgcmVmcywgdGhleSBtdXN0IG1hdGNoIHRoZSBudW1iZXIgb2YgY2hpbGRyZW4KCWlmIGV4aXN0aW5nQ2hpbGRSZWZzICE9IG5pbCAmJiBsZW4oZXhpc3RpbmdDaGlsZFJlZnMpID09IGxlbihzLmNoaWxkcmVuKSB7CgkJcy51bmlxdWVSZWZDZWxsLlJlZnMgPSBhcHBlbmQocy51bmlxdWVSZWZDZWxsLlJlZnMsIHJlZikKCgkJcmVmVXBkYXRlUmVxdWlyZWQgOj0gZmFsc2UKCQlmb3IgY2hpbGRJbmRleCwgY2hpbGRSZWYgOj0gcmFuZ2UgZXhpc3RpbmdDaGlsZFJlZnMgewoJCQluZXdDaGlsZFJlZiwgZXJyIDo9IGFmKHMuY2hpbGRyZW5bY2hpbGRJbmRleF0sIGNoaWxkUmVmKQoJCQlpZiBlcnIgIT0gbmlsIHsKCQkJCWFwcGVuZEVyciA9IG11bHRpZXJyb3IuQXBwZW5kKGFwcGVuZEVyciwgZXJyKQoJCQl9CgoJCQlpZiBuZXdDaGlsZFJlZiAhPSBjaGlsZFJlZiB7CgkJCQkvLyBDaGlsZCByZWYgY2hhbmdlZCwgbmVlZCB0byB1cGRhdGUgbWFwcGluZwoJCQkJZXhpc3RpbmdDaGlsZFJlZnNbY2hpbGRJbmRleF0gPSBuZXdDaGlsZFJlZgoJCQkJcmVmVXBkYXRlUmVxdWlyZWQgPSB0cnVlCgkJCX0KCQl9CgoJCWlmIGFwcGVuZEVyciAhPSBuaWwgewoJCQlyZXR1cm4gMCwgYXBwZW5kRXJyCgkJfQoKCQlpZiByZWZVcGRhdGVSZXF1aXJlZCB7CgkJCXMuc3RvcmUuVXBkYXRlTWFwcGluZyhyZWYsIGV4aXN0aW5nQ2hpbGRSZWZzLCBsYmxzKQoJCX0KCgkJcmV0dXJuIHJlZiwgbmlsCgl9CgoJLy8gTm8gZXhpc3RpbmcgbWFwcGluZywgcHJvY2VlZCB3aXRoIG5vcm1hbCBhcHBlbmQgdG8gYWxsIGNoaWxkcmVuCgl2YXIgZmlyc3ROb25aZXJvUmVmIHN0b3JhZ2UuU2VyaWVzUmVmCgl2YXIgbm9uWmVyb0NvdW50IGludAoKCS8vIE5vdGU6IHRoZXJlJ3MgYW5vdGhlciBvcHRpbWl6YXRpb24gd2hlcmUgd2UgY291bGQgdXNlIHRoZSByZXR1cm5lZCByZWYgaWYgYWxsIHRoZSBub24gemVybyByZWZzCgkvLyAgYXJlIHRoZSBzYW1lIHZhbHVlLiBUaGlzIGlzbid0IHNhZmUgYXMgd2Ugd2lsbCBtaXggZG93bnN0cmVhbSByZWZzIHdpdGggdW5pcXVlIHJlZnMgd2hpY2ggY291bGQKCS8vICBjb2xsaWRlLiBXZSBjb3VsZCBzdGFydCBhdCBtYXggdW5pdDY0IGZvciB1bmlxdWUgcmVmcyBhbmQgZ28gYmFja3dhcmRzIGxlc3NlbmluZyB0aGUgY2hhbmNlIG9mCgkvLyAJY29sbGlzaW9ucyBidXQgaXQncyByYXRoZXIgZGFuZ2Vyb3VzIGZvciBhbiB1bmxpa2VseSBlZGdlIGNhc2UuIElmIHR3byBjb21wb25lbnRzIGFyZSByZXR1cm5pbmcKCS8vIAl0aGUgc2FtZSByZWYgaXQncyB0d28gcmVtb3RlX3dyaXRlIGNvbXBvbmVudHMgd2hpY2ggc2hvdWxkIHByb2JhYmx5IGJlIG1lcmdlZCBpbiB0byBvbmUuCglmb3IgXywgY2hpbGQgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJY2hpbGRSZWYsIGVyciA6PSBhZihjaGlsZCwgcmVmKQoJCWlmIGVyciAhPSBuaWwgewoJCQlhcHBlbmRFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChhcHBlbmRFcnIsIGVycikKCgkJCS8vIFRPRE8gc2hvdWxkIEkgbWFrZSB0aGUgY2hpbGRSZWYgemVybyBoZXJlPwoJCX0KCgkJcy5jaGlsZFJlZnMgPSBhcHBlbmQocy5jaGlsZFJlZnMsIGNoaWxkUmVmKQoJCWlmIGNoaWxkUmVmICE9IDAgewoJCQlpZiBmaXJzdE5vblplcm9SZWYgPT0gMCB7CgkJCQlmaXJzdE5vblplcm9SZWYgPSBjaGlsZFJlZgoJCQl9CgkJCW5vblplcm9Db3VudCsrCgkJfQoJfQoKCWlmIGFwcGVuZEVyciAhPSBuaWwgewoJCXJldHVybiAwLCBhcHBlbmRFcnIKCX0KCglpZiBub25aZXJvQ291bnQgPT0gMCB7CgkJLy8gQWxsIGNoaWxkcmVuIHJldHVybmVkIHJlZiAwLCBzbyByZXR1cm4gdGhlIGlucHV0IHJlZgoJCXJldHVybiByZWYsIG5pbAoJfQoKCS8vIE9ubHkgb25lIGNoaWxkIHJldHVybmVkIGEgbm9uLXplcm8gcmVmLCB1c2UgdGhhdAoJaWYgbm9uWmVyb0NvdW50ID09IDEgewoJCXJldHVybiBmaXJzdE5vblplcm9SZWYsIG5pbAoJfQoKCS8vIFdlIGdvdCBkaWZmZXJlbnQgcmVmcyBiYWNrIGFuZCBuZWVkIHRvIGNyZWF0ZSBhIG5ldyBtYXBwaW5nCgl1bmlxdWVSZWYgOj0gcy5zdG9yZS5DcmVhdGVNYXBwaW5nKHMuY2hpbGRSZWZzLCBsYmxzKQoJcy51bmlxdWVSZWZDZWxsLlJlZnMgPSBhcHBlbmQocy51bmlxdWVSZWZDZWxsLlJlZnMsIHVuaXF1ZVJlZikKCXJldHVybiB1bmlxdWVSZWYsIG5pbAp9Cgp0eXBlIHVuaXFSZWZDaGlsZHJlbiBzdHJ1Y3QgewoJY2hpbGRSZWZzICpbXXN0b3JhZ2UuU2VyaWVzUmVmCglsYWJlbEhhc2ggdWludDY0Cn0KCnR5cGUgU2VyaWVzUmVmTWFwcGluZ1N0b3JlIHN0cnVjdCB7CgkvLyByZWZNYXBwaW5nTXUgcHJvdGVjdHMgdW5pcXVlUmVmVG9DaGlsZFJlZnMsIGxhYmVsSGFzaFRvVW5pcXVlUmVmIGFuZCBuZXh0VW5pcXVlUmVmCglyZWZNYXBwaW5nTXUgc3luYy5SV011dGV4CgkvLyB1bmlxdWVSZWZUb0NoaWxkUmVmcyBtYXBzIHRoZSB1bmlxdWUgcmVmIHRvIHRoZSBleHBlY3RlZCBjaGlsZCByZWYgaW4gb3JkZXIKCXVuaXF1ZVJlZlRvQ2hpbGRSZWZzIG1hcFtzdG9yYWdlLlNlcmllc1JlZl11bmlxUmVmQ2hpbGRyZW4KCS8vIGxhYmVsSGFzaFRvVW5pcXVlUmVmIG1hcHMgdGhlIGxhYmVsIGhhc2ggdG8gdW5pcXVlIHJlZi4KCWxhYmVsSGFzaFRvVW5pcXVlUmVmIG1hcFt1aW50NjRdc3RvcmFnZS5TZXJpZXNSZWYKCgkvLyBuZXh0VW5pcXVlUmVmIGlzIHRoZSBuZXh0IHJlZiBJRCB3ZSB3aWxsIGhhbmQgb3V0CgluZXh0VW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmCgoJLy8gdGltZXN0YW1wVHJhY2tpbmdNdSBwcm90ZWN0cyB1bmlxdWVSZWZUaW1lc3RhbXBzIGFuZCBjZWxsUG9vbAoJdGltZXN0YW1wVHJhY2tpbmdNdSBzeW5jLk11dGV4CgkvLyB1bmlxdWVSZWZUaW1lc3RhbXBzIG1hcHMgdW5pcXVlIHJlZnMgdG8gdGhlaXIgbGFzdCBhcHBlbmQgdGltZXN0YW1wCgl1bmlxdWVSZWZUaW1lc3RhbXBzIG1hcFtzdG9yYWdlLlNlcmllc1JlZl1pbnQ2NAoJLy8gY2VsbFBvb2wgaXMgdXNlZCB0byBwb29sIHNsaWNlcyBvZiBTZXJpZXNSZWZzIHVzZWQgZm9yIHRyYWNraW5nIHVuaXF1ZSByZWZzIGluIFRyYWNrQXBwZW5kZWRTZXJpZXMuCgljZWxsUG9vbCBzeW5jLlBvb2wKCgkvLyBDbGVhbnVwIGdvcm91dGluZSBjb29yZGluYXRpb24gKG5vIGxvY2sgcmVxdWlyZWQpCglzdGFydFJlZkNsZWFudXAgc3luYy5PbmNlCgljbGVhbnVwU3RhcnRlZCAgYXRvbWljLkJvb2wKCXN0b3BDbGVhbnVwICAgICBjaGFuIHN0cnVjdHt9CgljbGVhbnVwU3RvcHBlZCAgY2hhbiBzdHJ1Y3R7fQoKCS8vIE1ldHJpY3MgKHNhZmUgZm9yIGNvbmN1cnJlbnQgYWNjZXNzLCBubyBsb2NrIHJlcXVpcmVkKQoJYWN0aXZlTWFwcGluZ3MgIHByb21ldGhldXMuR2F1Z2UKCXRyYWNrZWRSZWZzICAgICBwcm9tZXRoZXVzLkdhdWdlCglyZWZzQ2xlYW5lZCAgICAgcHJvbWV0aGV1cy5Db3VudGVyCgl1bmlxdWVSZWZzVG90YWwgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZ1N0b3JlKHJlZyBwcm9tZXRoZXVzLlJlZ2lzdGVyZXIpICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUgewoJYWN0aXZlTWFwcGluZ3MgOj0gcHJvbWV0aGV1cy5OZXdHYXVnZShwcm9tZXRoZXVzLkdhdWdlT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfbWFwcGluZ3NfdG90YWwiLAoJCUhlbHA6ICJOdW1iZXIgb2YgYWN0aXZlIHVuaXF1ZSByZWYgbWFwcGluZ3MgaW4gdGhlIHN0b3JlLiIsCgl9KQoJdHJhY2tlZFJlZnMgOj0gcHJvbWV0aGV1cy5OZXdHYXVnZShwcm9tZXRoZXVzLkdhdWdlT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfdHJhY2tlZF9yZWZzX3RvdGFsIiwKCQlIZWxwOiAiTnVtYmVyIG9mIHJlZnMgYmVpbmcgdHJhY2tlZCBmb3IgdGltZXN0YW1wLWJhc2VkIGNsZWFudXAuIiwKCX0pCglyZWZzQ2xlYW5lZCA6PSBwcm9tZXRoZXVzLk5ld0NvdW50ZXIocHJvbWV0aGV1cy5Db3VudGVyT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfcmVmc19jbGVhbmVkX3RvdGFsIiwKCQlIZWxwOiAiVG90YWwgbnVtYmVyIG9mIHN0YWxlIHJlZnMgY2xlYW5lZCB1cCBvdmVyIHRpbWUuIiwKCX0pCgl1bmlxdWVSZWZzVG90YWwgOj0gcHJvbWV0aGV1cy5OZXdDb3VudGVyKHByb21ldGhldXMuQ291bnRlck9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3VuaXF1ZV9yZWZzX2NyZWF0ZWRfdG90YWwiLAoJCUhlbHA6ICJUb3RhbCBudW1iZXIgb2YgdW5pcXVlIHJlZnMgY3JlYXRlZC4iLAoJfSkKCglpZiByZWcgIT0gbmlsIHsKCQlyZWcuUmVnaXN0ZXIoYWN0aXZlTWFwcGluZ3MpCgkJcmVnLlJlZ2lzdGVyKHRyYWNrZWRSZWZzKQoJCXJlZy5SZWdpc3RlcihyZWZzQ2xlYW5lZCkKCQlyZWcuUmVnaXN0ZXIodW5pcXVlUmVmc1RvdGFsKQoJfQoKCXJldHVybiAmU2VyaWVzUmVmTWFwcGluZ1N0b3JlewoJCXVuaXF1ZVJlZlRvQ2hpbGRSZWZzOiBtYWtlKG1hcFtzdG9yYWdlLlNlcmllc1JlZl11bmlxUmVmQ2hpbGRyZW4pLAoJCW5leHRVbmlxdWVSZWY6ICAgICAgICAxLAoJCXVuaXF1ZVJlZlRpbWVzdGFtcHM6ICBtYWtlKG1hcFtzdG9yYWdlLlNlcmllc1JlZl1pbnQ2NCksCgkJbGFiZWxIYXNoVG9VbmlxdWVSZWY6IG1ha2UobWFwW3VpbnQ2NF1zdG9yYWdlLlNlcmllc1JlZiksCgkJY2VsbFBvb2w6IHN5bmMuUG9vbHsKCQkJTmV3OiBmdW5jKCkgYW55IHsKCQkJCXJldHVybiAmQ2VsbHtSZWZzOiBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIDEwMCl9CgkJCX0sCgkJfSwKCQlzdG9wQ2xlYW51cDogICAgIG1ha2UoY2hhbiBzdHJ1Y3R7fSksCgkJY2xlYW51cFN0b3BwZWQ6ICBtYWtlKGNoYW4gc3RydWN0e30pLAoJCWFjdGl2ZU1hcHBpbmdzOiAgYWN0aXZlTWFwcGluZ3MsCgkJdHJhY2tlZFJlZnM6ICAgICB0cmFja2VkUmVmcywKCQlyZWZzQ2xlYW5lZDogICAgIHJlZnNDbGVhbmVkLAoJCXVuaXF1ZVJlZnNUb3RhbDogdW5pcXVlUmVmc1RvdGFsLAoJfQp9Cgp0eXBlIENlbGwgc3RydWN0IHsKCVJlZnMgW11zdG9yYWdlLlNlcmllc1JlZgp9CgovLyBHZXRNYXBwaW5nIHJldHVybnMgZXhpc3RpbmcgY2hpbGQgcmVmIHJlc3VsdHMgZm9yIHRoZSBnaXZlbiB1bmlxdWUgcmVmIGlmIG9uZSBleGlzdHMuCi8vCi8vIElmIHRoZSBwYXNzZWQgdW5pcXVlUmVmIGlzIHplcm8sIHRoZSBtZXRob2Qgd2lsbCBhdHRlbXB0IHRvIGZpbmQgYSBtYXBwaW5nIHVzaW5nIHBhc3NlZCBsYWJlbHMuCi8vIFJldHVybnMgbmlsIGlmIG5vIG1hcHBpbmcgZXhpc3RzLgovLwovLyBUaGUgcmV0dXJuZWQgc2xpY2UgbWF5IGJlIG1vZGlmaWVkIGJ5IHRoZSBjYWxsZXIsIGJ1dCBVcGRhdGVNYXBwaW5nIG11c3QgYmUgY2FsbGVkCi8vIGFmdGVyd2FyZHMgdG8gcGVyc2lzdCBjaGFuZ2VzLiBOb3RlIHRoYXQgY29uY3VycmVudCBhcHBlbmRlcnMgbWF5IHJhY2UgdG8gdXBkYXRlIHRoZQovLyBzYW1lIG1hcHBpbmcgd2l0aCBkaWZmZXJlbnQgdmFsdWVzLCB3aGljaCBpcyBzYWZlIGJlY2F1c2Ugc3RhbGUgbWFwcGluZ3MgYXJlIHNlbGYtY29ycmVjdGluZyAtCi8vIHVzaW5nIGEgc3RhbGUgcmVmIHdpbGwgY2F1c2UgdGhlIGNoaWxkIGFwcGVuZGVyIHRvIHJldHVybiBhIG5ldyByZWYgb24gdGhlIG5leHQgYXBwZW5kLgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIEdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYgewoJcy5yZWZNYXBwaW5nTXUuUkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuUlVubG9jaygpCgoJaWYgdW5pcXVlUmVmID09IDAgewoJCS8vIFNvbWUgY29uc3VtZXJzIGRvbid0IG1lbW8gdGhlIGdsb2JhbCByZWYuIFRyeSB0byBsb29rdXAgYSByZWYgYnkgbGFiZWwgaGFzaC4KCQlsYWJlbEhhc2ggOj0gbGJscy5IYXNoKCkKCQlnb3RSZWYsIG9rIDo9IHMubGFiZWxIYXNoVG9VbmlxdWVSZWZbbGFiZWxIYXNoXQoJCWlmICFvayB7CgkJCXJldHVybiBuaWwKCQl9CgoJCXVuaXF1ZVJlZiA9IGdvdFJlZgoJfQoKCWlmIG1hcHBpbmcsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXTsgb2sgewoJCXJldHVybiAqbWFwcGluZy5jaGlsZFJlZnMKCX0KCXJldHVybiBuaWwKfQoKLy8gQ3JlYXRlTWFwcGluZyBjcmVhdGVzIGEgbmV3IHVuaXF1ZSByZWYgbWFwcGluZyBmb3IgdGhlIGdpdmVuIGNoaWxkIHJlZiByZXN1bHRzLgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIENyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmIHsKCS8vIFN0YXJ0IGNsZWFudXAgZ29yb3V0aW5lIG9uIGZpcnN0IG1hcHBpbmcKCXMuc3RhcnRSZWZDbGVhbnVwLkRvKGZ1bmMoKSB7CgkJcy5jbGVhbnVwU3RhcnRlZC5TdG9yZSh0cnVlKQoJCWdvIHMuY2xlYW51cFN0YWxlUmVmcygpCgl9KQoKCS8vIFN0b3JlIGEgY29weSBvZiB0aGUgY2hpbGQgcmVmIHJlc3VsdHMgZGlyZWN0bHkKCWNoaWxkUmVmU2xpY2UgOj0gbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsZW4ocmVmUmVzdWx0cykpCgljb3B5KGNoaWxkUmVmU2xpY2UsIHJlZlJlc3VsdHMpCgoJLy8gSGFzaCBsYWJlbHMgdG8gZm9yIHRoZSBmYWxsYmFjayBsb29rdXAgdGFibGUKCWxhYmVsSGFzaCA6PSBsYmxzLkhhc2goKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgkvLyBDcmVhdGUgYSBuZXcgdW5pcXVlIHJlZgoJdW5pcXVlUmVmIDo9IHMubmV4dFVuaXF1ZVJlZgoJcy5uZXh0VW5pcXVlUmVmKysKCglzLmxhYmVsSGFzaFRvVW5pcXVlUmVmW2xhYmVsSGFzaF0gPSB1bmlxdWVSZWYKCXMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXSA9IHVuaXFSZWZDaGlsZHJlbnsKCQljaGlsZFJlZnM6ICZjaGlsZFJlZlNsaWNlLAoJCWxhYmVsSGFzaDogbGFiZWxIYXNoLAoJfQoKCXMuYWN0aXZlTWFwcGluZ3MuSW5jKCkKCXMudW5pcXVlUmVmc1RvdGFsLkluYygpCgoJcmV0dXJuIHVuaXF1ZVJlZgp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIFVwZGF0ZU1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCByZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgewoJaWYgdW5pcXVlUmVmID09IDAgewoJCXJldHVybgoJfQoKCWNoaWxkUmVmU2xpY2UgOj0gbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsZW4ocmVmUmVzdWx0cykpCgljb3B5KGNoaWxkUmVmU2xpY2UsIHJlZlJlc3VsdHMpCgoJcy5yZWZNYXBwaW5nTXUuTG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoKCS8vIEVuc3VyZSB0aGF0IGxhYmVsIGhhc2ggaW5kZXggaXMgdXAgdG8gZGF0ZSB0byBoYW5kbGUgcG9zc2libGUgaGFzaCBjb2xsaXNpb25zLgoJLy8gVE9ETzogaXMgdGhpcyBuZWNlc3Nhcnk/CgluZXdIYXNoIDo9IGxibHMuSGFzaCgpCglwcmV2LCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0KCWlmIG9rICYmIHByZXYubGFiZWxIYXNoICE9IG5ld0hhc2ggewoJCWRlbGV0ZShzLmxhYmVsSGFzaFRvVW5pcXVlUmVmLCBwcmV2LmxhYmVsSGFzaCkKCQlzLmxhYmVsSGFzaFRvVW5pcXVlUmVmW25ld0hhc2hdID0gdW5pcXVlUmVmCgl9CgoJcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdID0gdW5pcVJlZkNoaWxkcmVuewoJCWNoaWxkUmVmczogJmNoaWxkUmVmU2xpY2UsCgkJbGFiZWxIYXNoOiBsYmxzLkhhc2goKSwKCX0KfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKSB7CglzLnRpbWVzdGFtcFRyYWNraW5nTXUuTG9jaygpCglkZWZlciBzLnRpbWVzdGFtcFRyYWNraW5nTXUuVW5sb2NrKCkKCglmb3IgXywgciA6PSByYW5nZSBjZWxsLlJlZnMgewoJCXMudW5pcXVlUmVmVGltZXN0YW1wc1tyXSA9IHRzCgl9CgoJcy50cmFja2VkUmVmcy5TZXQoZmxvYXQ2NChsZW4ocy51bmlxdWVSZWZUaW1lc3RhbXBzKSkpCgoJY2VsbC5SZWZzID0gY2VsbC5SZWZzWzowXQoJcy5jZWxsUG9vbC5QdXQoY2VsbCkKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBHZXRDZWxsRm9yQXBwZW5kZWRTZXJpZXMoKSAqQ2VsbCB7CglyZXR1cm4gcy5jZWxsUG9vbC5HZXQoKS4oKkNlbGwpCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgY2xlYW51cFN0YWxlUmVmcygpIHsKCWRlZmVyIGNsb3NlKHMuY2xlYW51cFN0b3BwZWQpCgoJdGlja2VyIDo9IHRpbWUuTmV3VGlja2VyKDE1ICogdGltZS5NaW51dGUpCglkZWZlciB0aWNrZXIuU3RvcCgpCgoJZm9yIHsKCQlzZWxlY3QgewoJCWNhc2UgPC10aWNrZXIuQzoKCQkJY3V0b2ZmVGltZSA6PSB0aW1lLk5vdygpLkFkZCgtMTUgKiB0aW1lLk1pbnV0ZSkuVW5peCgpCgoJCQkvLyBIb2xkIGJvdGggbG9ja3MgdG8gcHJldmVudCByYWNlIGNvbmRpdGlvbiB3aGVyZSBhIHJlZiBjb3VsZCBiZQoJCQkvLyBhcHBlbmRlZCBhZnRlciB3ZSBkZWxldGUgaXQgZnJvbSB1bmlxdWVSZWZDZWxsIGJ1dCBiZWZvcmUKCQkJLy8gd2UgZGVsZXRlIGl0IGZyb20gdW5pcXVlUmVmVG9DaGlsZFJlZnMKCQkJcy50aW1lc3RhbXBUcmFja2luZ011LkxvY2soKQoJCQlzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCgkJCXN0YWxlUmVmQ291bnQgOj0gMAoJCQlmb3IgcmVmLCB0cyA6PSByYW5nZSBzLnVuaXF1ZVJlZlRpbWVzdGFtcHMgewoJCQkJaWYgdHMgPCBjdXRvZmZUaW1lIHsKCQkJCQlzdGFsZVJlZkNvdW50KysKCgkJCQkJdiwgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1tyZWZdCgkJCQkJaWYgb2sgewoJCQkJCQlkZWxldGUocy5sYWJlbEhhc2hUb1VuaXF1ZVJlZiwgdi5sYWJlbEhhc2gpCgkJCQkJfQoKCQkJCQlkZWxldGUocy51bmlxdWVSZWZUaW1lc3RhbXBzLCByZWYpCgkJCQkJZGVsZXRlKHMudW5pcXVlUmVmVG9DaGlsZFJlZnMsIHJlZikKCQkJCX0KCQkJfQoKCQkJLy8gVXBkYXRlIG1ldHJpY3MKCQkJaWYgc3RhbGVSZWZDb3VudCA+IDAgewoJCQkJcy5yZWZzQ2xlYW5lZC5BZGQoZmxvYXQ2NChzdGFsZVJlZkNvdW50KSkKCQkJCXMuYWN0aXZlTWFwcGluZ3MuU3ViKGZsb2F0NjQoc3RhbGVSZWZDb3VudCkpCgkJCQlzLnRyYWNrZWRSZWZzLlNldChmbG9hdDY0KGxlbihzLnVuaXF1ZVJlZlRpbWVzdGFtcHMpKSkKCQkJfQoKCQkJcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCQkJcy50aW1lc3RhbXBUcmFja2luZ011LlVubG9jaygpCgoJCWNhc2UgPC1zLnN0b3BDbGVhbnVwOgoJCQlyZXR1cm4KCQl9Cgl9Cn0KCi8vIENsZWFyIHdpbGwgY2xlYXIgYWxsIGludGVybmFsIG1hcHBpbmdzIGFuZCBzdG9wIHRoZSBjbGVhbmVyIGdvcm91dGluZSBpZiBpdCBpcyBydW5uaW5nLgovLyBJdCBpcyBzYWZlIHRvIHJlLXVzZSB0aGUgc2FtZSBpbnN0YW5jZSBhZnRlciBjYWxsaW5nIENsZWFyLgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIENsZWFyKCkgewoJLy8gU3RvcCB0aGUgY2xlYW51cCBnb3JvdXRpbmUgYW5kIHdhaXQgZm9yIGl0IHRvIGJlIHN0b3BwZWQgc28gd2UgY2FuCgkvLyBhdm9pZCBhIHBvc3NpYmxlIGRlYWRsb2NrIHdpdGggY2xlYW51cCB0aGF0IGFsc28gaG9sZHMgYm90aCBsb2NrcwoJaWYgcy5jbGVhbnVwU3RhcnRlZC5Mb2FkKCkgewoJCXNlbGVjdCB7CgkJY2FzZSA8LXMuc3RvcENsZWFudXA6CgkJCS8vIEFscmVhZHkgY2xvc2VkCgkJZGVmYXVsdDoKCQkJY2xvc2Uocy5zdG9wQ2xlYW51cCkKCQkJPC1zLmNsZWFudXBTdG9wcGVkCgkJfQoJfQoKCS8vIFdlIG5lZWQgdG8gaG9sZCBib3RoIGxvY2tzIHRvIGRvIHRoaXMgc2FmZWx5IGFuZCB3ZSBkbyBpdCBpbiB0aGUgc2FtZSBvcmRlciBhcwoJLy8gY2xlYW51cFN0YWxlUmVmcy4gV2Ugc3RvcHBlZCBhbmQgd2FpdGVkIGZvciB0aGUgYmFja2dyb3VuZCB3b3JrZXIgdGhhdCBjYWxscyBpdAoJLy8gdG8gZmluaXNoIGJ1dCBzb21lIGV4dHJhIHNhZmV0eSB3b24ndCBodXJ0LgoJcy50aW1lc3RhbXBUcmFja2luZ011LkxvY2soKQoJZGVmZXIgcy50aW1lc3RhbXBUcmFja2luZ011LlVubG9jaygpCgoJcy5yZWZNYXBwaW5nTXUuTG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoKCWNsZWFyKHMudW5pcXVlUmVmVG9DaGlsZFJlZnMpCgljbGVhcihzLnVuaXF1ZVJlZlRpbWVzdGFtcHMpCgoJLy8gcmVzZXQgdGhlIHBvb2wKCXMuY2VsbFBvb2wgPSBzeW5jLlBvb2x7CgkJTmV3OiBmdW5jKCkgYW55IHsKCQkJcmV0dXJuICZDZWxse1JlZnM6IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgMTAwKX0KCQl9LAoJfQoKCS8vIE5PVEU6IFdlIGRvIE5PVCByZXNldCBuZXh0VW5pcXVlUmVmIGhlcmUuIFJlc2V0dGluZyBpdCB3b3VsZCBjYXVzZSByZWYgY29sbGlzaW9ucwoJLy8gd2l0aCBjb21wb25lbnRzIGxpa2UgcHJvbWV0aGV1cy5zY3JhcGUgd2hpY2ggd2lsbCBrZWVwIHJlLXNlbmRpbmcgdGhlIHNhbWUgY2FjaGVkIHJlZnMuCgkvLyBXZSBjb250aW51ZSBpbmNyZW1lbnRpbmcgdG8gZW5zdXJlIGFsbCByZWZzIHJlbWFpbiB1bmlxdWUgYWNyb3NzIHRoZSBsaWZldGltZSBvZiB0aGUgcHJvY2Vzcy4KCgkvLyBSZXNldCBtZXRyaWNzCglzLmFjdGl2ZU1hcHBpbmdzLlNldCgwKQoJcy50cmFja2VkUmVmcy5TZXQoMCkKCgkvLyBSZXNldCBjaGFubmVscyBhbmQgZmxhZ3MKCXMuc3RvcENsZWFudXAgPSBtYWtlKGNoYW4gc3RydWN0e30pCglzLmNsZWFudXBTdG9wcGVkID0gbWFrZShjaGFuIHN0cnVjdHt9KQoJcy5zdGFydFJlZkNsZWFudXAgPSBzeW5jLk9uY2V7fQoJcy5jbGVhbnVwU3RhcnRlZC5TdG9yZShmYWxzZSkKfQo="
    }
  }
}
```

Response (packages.DriverResponse):
Error: `err: context canceled: stderr: `

#### drv #5

Trace meta: spanId=9, ts=1770837171179, ts_iso=2026-02-11T19:12:51.179000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy/internal/validator"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkiZm10IgoJInN5bmMiCgkic3luYy9hdG9taWMiCgkidGltZSIKCgkiZ2l0aHViLmNvbS9oYXNoaWNvcnAvZ28tbXVsdGllcnJvciIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvY2xpZW50X2dvbGFuZy9wcm9tZXRoZXVzIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2V4ZW1wbGFyIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2hpc3RvZ3JhbSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9sYWJlbHMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvbWV0YWRhdGEiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvc3RvcmFnZSIKKQoKdHlwZSBNYXBwaW5nU3RvcmUgaW50ZXJmYWNlIHsKCUdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYKCUNyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmCglVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpCglUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKQoJR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwKfQoKdHlwZSBzZXJpZXNSZWZNYXBwaW5nIHN0cnVjdCB7CglzdGFydCAgICB0aW1lLlRpbWUKCWNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlcgoJc3RvcmUgICAgTWFwcGluZ1N0b3JlCgoJdW5pcXVlUmVmQ2VsbCAqQ2VsbAoKCS8vIGNoaWxkUmVmcyBpcyByZXVzZWQgZm9yIGVhY2ggYXBwZW5kIGNhbGwgdG8gYXZvaWQgYWxsb2NhdGlvbnMuIFRoaXMgaXMgc2FmZSBiZWNhdXNlIHN0b3JhZ2UuQXBwZW5kZXIgc2hvdWxkIG5ldmVyCgkvLyBoYXZlIGNvbmN1cnJlbnQgY2FsbHMgdG8gQXBwZW5kIG1ldGhvZHMuCgljaGlsZFJlZnMgICAgICAgIFtdc3RvcmFnZS5TZXJpZXNSZWYKCXdyaXRlTGF0ZW5jeSAgICAgcHJvbWV0aGV1cy5IaXN0b2dyYW0KCXNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZyhjaGlsZHJlbiBbXXN0b3JhZ2UuQXBwZW5kZXIsIHN0b3JlIE1hcHBpbmdTdG9yZSwgd3JpdGVMYXRlbmN5IHByb21ldGhldXMuSGlzdG9ncmFtLCBzYW1wbGVzRm9yd2FyZGVkIHByb21ldGhldXMuQ291bnRlcikgc3RvcmFnZS5BcHBlbmRlciB7Cgl1bmlxdWVSZWZDZWxsIDo9IHN0b3JlLkdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpCgoJcmV0dXJuICZzZXJpZXNSZWZNYXBwaW5newoJCWNoaWxkcmVuOiAgICAgICAgIGNoaWxkcmVuLAoJCXN0b3JlOiAgICAgICAgICAgIHN0b3JlLAoJCXdyaXRlTGF0ZW5jeTogICAgIHdyaXRlTGF0ZW5jeSwKCQlzYW1wbGVzRm9yd2FyZGVkOiBzYW1wbGVzRm9yd2FyZGVkLAoKCQl1bmlxdWVSZWZDZWxsOiB1bmlxdWVSZWZDZWxsLAoJCWNoaWxkUmVmczogICAgIG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgbGVuKGNoaWxkcmVuKSksCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFNldE9wdGlvbnMob3B0cyAqc3RvcmFnZS5BcHBlbmRPcHRpb25zKSB7CglmbXQuUHJpbnRsbigiYSkKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWMuU2V0T3B0aW9ucyhvcHRzKQoJfQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBDb21taXQoKSBlcnJvciB7CglzLnN0b3JlLlRyYWNrQXBwZW5kZWRTZXJpZXModGltZS5Ob3coKS5Vbml4KCksIHMudW5pcXVlUmVmQ2VsbCkKCgl2YXIgbXVsdGlFcnIgZXJyb3IKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWVyciA6PSBjLkNvbW1pdCgpCgkJaWYgZXJyICE9IG5pbCB7CgkJCW11bHRpRXJyID0gbXVsdGllcnJvci5BcHBlbmQobXVsdGlFcnIsIGVycikKCQl9Cgl9CglyZXR1cm4gbXVsdGlFcnIKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgUm9sbGJhY2soKSBlcnJvciB7CgkvLyBXZSBzdGlsbCB0cmFjayByb2xsZWQgYmFjayBzZXJpZXMgc28gd2UgY2FuIHByb3Blcmx5CgkvLyBjbGVhbiB1cCBhbnkgc2VyaWVzIHRoYXQgd2FzIGFwcGVuZGVkCglzLnN0b3JlLlRyYWNrQXBwZW5kZWRTZXJpZXModGltZS5Ob3coKS5Vbml4KCksIHMudW5pcXVlUmVmQ2VsbCkKCgl2YXIgbXVsdGlFcnIgZXJyb3IKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWVyciA6PSBjLlJvbGxiYWNrKCkKCQlpZiBlcnIgIT0gbmlsIHsKCQkJbXVsdGlFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChtdWx0aUVyciwgZXJyKQoJCX0KCX0KCXJldHVybiBtdWx0aUVycgp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSByZWNvcmRMYXRlbmN5KCkgewoJaWYgcy5zdGFydC5Jc1plcm8oKSB7CgkJcmV0dXJuCgl9CglkdXJhdGlvbiA6PSB0aW1lLlNpbmNlKHMuc3RhcnQpCglzLndyaXRlTGF0ZW5jeS5PYnNlcnZlKGR1cmF0aW9uLlNlY29uZHMoKSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgcmVzZXRGaWVsZHMoKSB7CgkvLyBSZXNldCBjaGlsZFJlZnMgc2xpY2UgbGVuZ3RoIHRvIDAgZm9yIHJldXNlCglzLmNoaWxkUmVmcyA9IHMuY2hpbGRSZWZzWzowXQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmQocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQgaW50NjQsIHYgZmxvYXQ2NCkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJbmV3UmVmLCBlcnIgOj0gYXBwZW5kZXIuQXBwZW5kKHJlZiwgbCwgdCwgdikKCQlpZiBlcnIgPT0gbmlsIHsKCQkJcy5zYW1wbGVzRm9yd2FyZGVkLkluYygpCgkJfQoJCXJldHVybiBuZXdSZWYsIGVycgoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kRXhlbXBsYXIocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIGUgZXhlbXBsYXIuRXhlbXBsYXIpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRFeGVtcGxhcihyZWYsIGwsIGUpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRIaXN0b2dyYW0ocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQgaW50NjQsIGggKmhpc3RvZ3JhbS5IaXN0b2dyYW0sIGZoICpoaXN0b2dyYW0uRmxvYXRIaXN0b2dyYW0pIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRIaXN0b2dyYW0ocmVmLCBsLCB0LCBoLCBmaCkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEhpc3RvZ3JhbUNUWmVyb1NhbXBsZShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCwgY3QgaW50NjQsIGggKmhpc3RvZ3JhbS5IaXN0b2dyYW0sIGZoICpoaXN0b2dyYW0uRmxvYXRIaXN0b2dyYW0pIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRIaXN0b2dyYW1DVFplcm9TYW1wbGUocmVmLCBsLCB0LCBjdCwgaCwgZmgpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBVcGRhdGVNZXRhZGF0YShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgbSBtZXRhZGF0YS5NZXRhZGF0YSkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLlVwZGF0ZU1ldGFkYXRhKHJlZiwgbCwgbSkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZENUWmVyb1NhbXBsZShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCwgY3QgaW50NjQpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRDVFplcm9TYW1wbGUocmVmLCBsLCB0LCBjdCkKCX0pCn0KCnR5cGUgYXBwZW5kZXJGdW5jIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgYXBwZW5kVG9DaGlsZHJlbihyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscywgYWYgYXBwZW5kZXJGdW5jKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglkZWZlciBzLnJlc2V0RmllbGRzKCkKCglpZiBzLnN0YXJ0LklzWmVybygpIHsKCQlzLnN0YXJ0ID0gdGltZS5Ob3coKQoJfQoKCS8vIENoZWNrIGlmIHRoZSBpbmNvbWluZyByZWYgaGFzIHJlZiBtYXBwaW5ncwoJZXhpc3RpbmdDaGlsZFJlZnMgOj0gcy5zdG9yZS5HZXRNYXBwaW5nKHJlZiwgbGJscykKCgl2YXIgYXBwZW5kRXJyIGVycm9yCgoJLy8gU2FuaXR5IGNoZWNrOiBpZiB3ZSBoYXZlIGV4aXN0aW5nIGNoaWxkIHJlZnMsIHRoZXkgbXVzdCBtYXRjaCB0aGUgbnVtYmVyIG9mIGNoaWxkcmVuCglpZiBleGlzdGluZ0NoaWxkUmVmcyAhPSBuaWwgJiYgbGVuKGV4aXN0aW5nQ2hpbGRSZWZzKSA9PSBsZW4ocy5jaGlsZHJlbikgewoJCXMudW5pcXVlUmVmQ2VsbC5SZWZzID0gYXBwZW5kKHMudW5pcXVlUmVmQ2VsbC5SZWZzLCByZWYpCgoJCXJlZlVwZGF0ZVJlcXVpcmVkIDo9IGZhbHNlCgkJZm9yIGNoaWxkSW5kZXgsIGNoaWxkUmVmIDo9IHJhbmdlIGV4aXN0aW5nQ2hpbGRSZWZzIHsKCQkJbmV3Q2hpbGRSZWYsIGVyciA6PSBhZihzLmNoaWxkcmVuW2NoaWxkSW5kZXhdLCBjaGlsZFJlZikKCQkJaWYgZXJyICE9IG5pbCB7CgkJCQlhcHBlbmRFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChhcHBlbmRFcnIsIGVycikKCQkJfQoKCQkJaWYgbmV3Q2hpbGRSZWYgIT0gY2hpbGRSZWYgewoJCQkJLy8gQ2hpbGQgcmVmIGNoYW5nZWQsIG5lZWQgdG8gdXBkYXRlIG1hcHBpbmcKCQkJCWV4aXN0aW5nQ2hpbGRSZWZzW2NoaWxkSW5kZXhdID0gbmV3Q2hpbGRSZWYKCQkJCXJlZlVwZGF0ZVJlcXVpcmVkID0gdHJ1ZQoJCQl9CgkJfQoKCQlpZiBhcHBlbmRFcnIgIT0gbmlsIHsKCQkJcmV0dXJuIDAsIGFwcGVuZEVycgoJCX0KCgkJaWYgcmVmVXBkYXRlUmVxdWlyZWQgewoJCQlzLnN0b3JlLlVwZGF0ZU1hcHBpbmcocmVmLCBleGlzdGluZ0NoaWxkUmVmcywgbGJscykKCQl9CgoJCXJldHVybiByZWYsIG5pbAoJfQoKCS8vIE5vIGV4aXN0aW5nIG1hcHBpbmcsIHByb2NlZWQgd2l0aCBub3JtYWwgYXBwZW5kIHRvIGFsbCBjaGlsZHJlbgoJdmFyIGZpcnN0Tm9uWmVyb1JlZiBzdG9yYWdlLlNlcmllc1JlZgoJdmFyIG5vblplcm9Db3VudCBpbnQKCgkvLyBOb3RlOiB0aGVyZSdzIGFub3RoZXIgb3B0aW1pemF0aW9uIHdoZXJlIHdlIGNvdWxkIHVzZSB0aGUgcmV0dXJuZWQgcmVmIGlmIGFsbCB0aGUgbm9uIHplcm8gcmVmcwoJLy8gIGFyZSB0aGUgc2FtZSB2YWx1ZS4gVGhpcyBpc24ndCBzYWZlIGFzIHdlIHdpbGwgbWl4IGRvd25zdHJlYW0gcmVmcyB3aXRoIHVuaXF1ZSByZWZzIHdoaWNoIGNvdWxkCgkvLyAgY29sbGlkZS4gV2UgY291bGQgc3RhcnQgYXQgbWF4IHVuaXQ2NCBmb3IgdW5pcXVlIHJlZnMgYW5kIGdvIGJhY2t3YXJkcyBsZXNzZW5pbmcgdGhlIGNoYW5jZSBvZgoJLy8gCWNvbGxpc2lvbnMgYnV0IGl0J3MgcmF0aGVyIGRhbmdlcm91cyBmb3IgYW4gdW5saWtlbHkgZWRnZSBjYXNlLiBJZiB0d28gY29tcG9uZW50cyBhcmUgcmV0dXJuaW5nCgkvLyAJdGhlIHNhbWUgcmVmIGl0J3MgdHdvIHJlbW90ZV93cml0ZSBjb21wb25lbnRzIHdoaWNoIHNob3VsZCBwcm9iYWJseSBiZSBtZXJnZWQgaW4gdG8gb25lLgoJZm9yIF8sIGNoaWxkIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWNoaWxkUmVmLCBlcnIgOj0gYWYoY2hpbGQsIHJlZikKCQlpZiBlcnIgIT0gbmlsIHsKCQkJYXBwZW5kRXJyID0gbXVsdGllcnJvci5BcHBlbmQoYXBwZW5kRXJyLCBlcnIpCgoJCQkvLyBUT0RPIHNob3VsZCBJIG1ha2UgdGhlIGNoaWxkUmVmIHplcm8gaGVyZT8KCQl9CgoJCXMuY2hpbGRSZWZzID0gYXBwZW5kKHMuY2hpbGRSZWZzLCBjaGlsZFJlZikKCQlpZiBjaGlsZFJlZiAhPSAwIHsKCQkJaWYgZmlyc3ROb25aZXJvUmVmID09IDAgewoJCQkJZmlyc3ROb25aZXJvUmVmID0gY2hpbGRSZWYKCQkJfQoJCQlub25aZXJvQ291bnQrKwoJCX0KCX0KCglpZiBhcHBlbmRFcnIgIT0gbmlsIHsKCQlyZXR1cm4gMCwgYXBwZW5kRXJyCgl9CgoJaWYgbm9uWmVyb0NvdW50ID09IDAgewoJCS8vIEFsbCBjaGlsZHJlbiByZXR1cm5lZCByZWYgMCwgc28gcmV0dXJuIHRoZSBpbnB1dCByZWYKCQlyZXR1cm4gcmVmLCBuaWwKCX0KCgkvLyBPbmx5IG9uZSBjaGlsZCByZXR1cm5lZCBhIG5vbi16ZXJvIHJlZiwgdXNlIHRoYXQKCWlmIG5vblplcm9Db3VudCA9PSAxIHsKCQlyZXR1cm4gZmlyc3ROb25aZXJvUmVmLCBuaWwKCX0KCgkvLyBXZSBnb3QgZGlmZmVyZW50IHJlZnMgYmFjayBhbmQgbmVlZCB0byBjcmVhdGUgYSBuZXcgbWFwcGluZwoJdW5pcXVlUmVmIDo9IHMuc3RvcmUuQ3JlYXRlTWFwcGluZyhzLmNoaWxkUmVmcywgbGJscykKCXMudW5pcXVlUmVmQ2VsbC5SZWZzID0gYXBwZW5kKHMudW5pcXVlUmVmQ2VsbC5SZWZzLCB1bmlxdWVSZWYpCglyZXR1cm4gdW5pcXVlUmVmLCBuaWwKfQoKdHlwZSB1bmlxUmVmQ2hpbGRyZW4gc3RydWN0IHsKCWNoaWxkUmVmcyAqW11zdG9yYWdlLlNlcmllc1JlZgoJbGFiZWxIYXNoIHVpbnQ2NAp9Cgp0eXBlIFNlcmllc1JlZk1hcHBpbmdTdG9yZSBzdHJ1Y3QgewoJLy8gcmVmTWFwcGluZ011IHByb3RlY3RzIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzLCBsYWJlbEhhc2hUb1VuaXF1ZVJlZiBhbmQgbmV4dFVuaXF1ZVJlZgoJcmVmTWFwcGluZ011IHN5bmMuUldNdXRleAoJLy8gdW5pcXVlUmVmVG9DaGlsZFJlZnMgbWFwcyB0aGUgdW5pcXVlIHJlZiB0byB0aGUgZXhwZWN0ZWQgY2hpbGQgcmVmIGluIG9yZGVyCgl1bmlxdWVSZWZUb0NoaWxkUmVmcyBtYXBbc3RvcmFnZS5TZXJpZXNSZWZddW5pcVJlZkNoaWxkcmVuCgkvLyBsYWJlbEhhc2hUb1VuaXF1ZVJlZiBtYXBzIHRoZSBsYWJlbCBoYXNoIHRvIHVuaXF1ZSByZWYuCglsYWJlbEhhc2hUb1VuaXF1ZVJlZiBtYXBbdWludDY0XXN0b3JhZ2UuU2VyaWVzUmVmCgoJLy8gbmV4dFVuaXF1ZVJlZiBpcyB0aGUgbmV4dCByZWYgSUQgd2Ugd2lsbCBoYW5kIG91dAoJbmV4dFVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZgoKCS8vIHRpbWVzdGFtcFRyYWNraW5nTXUgcHJvdGVjdHMgdW5pcXVlUmVmVGltZXN0YW1wcyBhbmQgY2VsbFBvb2wKCXRpbWVzdGFtcFRyYWNraW5nTXUgc3luYy5NdXRleAoJLy8gdW5pcXVlUmVmVGltZXN0YW1wcyBtYXBzIHVuaXF1ZSByZWZzIHRvIHRoZWlyIGxhc3QgYXBwZW5kIHRpbWVzdGFtcAoJdW5pcXVlUmVmVGltZXN0YW1wcyBtYXBbc3RvcmFnZS5TZXJpZXNSZWZdaW50NjQKCS8vIGNlbGxQb29sIGlzIHVzZWQgdG8gcG9vbCBzbGljZXMgb2YgU2VyaWVzUmVmcyB1c2VkIGZvciB0cmFja2luZyB1bmlxdWUgcmVmcyBpbiBUcmFja0FwcGVuZGVkU2VyaWVzLgoJY2VsbFBvb2wgc3luYy5Qb29sCgoJLy8gQ2xlYW51cCBnb3JvdXRpbmUgY29vcmRpbmF0aW9uIChubyBsb2NrIHJlcXVpcmVkKQoJc3RhcnRSZWZDbGVhbnVwIHN5bmMuT25jZQoJY2xlYW51cFN0YXJ0ZWQgIGF0b21pYy5Cb29sCglzdG9wQ2xlYW51cCAgICAgY2hhbiBzdHJ1Y3R7fQoJY2xlYW51cFN0b3BwZWQgIGNoYW4gc3RydWN0e30KCgkvLyBNZXRyaWNzIChzYWZlIGZvciBjb25jdXJyZW50IGFjY2Vzcywgbm8gbG9jayByZXF1aXJlZCkKCWFjdGl2ZU1hcHBpbmdzICBwcm9tZXRoZXVzLkdhdWdlCgl0cmFja2VkUmVmcyAgICAgcHJvbWV0aGV1cy5HYXVnZQoJcmVmc0NsZWFuZWQgICAgIHByb21ldGhldXMuQ291bnRlcgoJdW5pcXVlUmVmc1RvdGFsIHByb21ldGhldXMuQ291bnRlcgp9CgpmdW5jIE5ld1Nlcmllc1JlZk1hcHBpbmdTdG9yZShyZWcgcHJvbWV0aGV1cy5SZWdpc3RlcmVyKSAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlIHsKCWFjdGl2ZU1hcHBpbmdzIDo9IHByb21ldGhldXMuTmV3R2F1Z2UocHJvbWV0aGV1cy5HYXVnZU9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX21hcHBpbmdzX3RvdGFsIiwKCQlIZWxwOiAiTnVtYmVyIG9mIGFjdGl2ZSB1bmlxdWUgcmVmIG1hcHBpbmdzIGluIHRoZSBzdG9yZS4iLAoJfSkKCXRyYWNrZWRSZWZzIDo9IHByb21ldGhldXMuTmV3R2F1Z2UocHJvbWV0aGV1cy5HYXVnZU9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3RyYWNrZWRfcmVmc190b3RhbCIsCgkJSGVscDogIk51bWJlciBvZiByZWZzIGJlaW5nIHRyYWNrZWQgZm9yIHRpbWVzdGFtcC1iYXNlZCBjbGVhbnVwLiIsCgl9KQoJcmVmc0NsZWFuZWQgOj0gcHJvbWV0aGV1cy5OZXdDb3VudGVyKHByb21ldGhldXMuQ291bnRlck9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3JlZnNfY2xlYW5lZF90b3RhbCIsCgkJSGVscDogIlRvdGFsIG51bWJlciBvZiBzdGFsZSByZWZzIGNsZWFuZWQgdXAgb3ZlciB0aW1lLiIsCgl9KQoJdW5pcXVlUmVmc1RvdGFsIDo9IHByb21ldGhldXMuTmV3Q291bnRlcihwcm9tZXRoZXVzLkNvdW50ZXJPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV91bmlxdWVfcmVmc19jcmVhdGVkX3RvdGFsIiwKCQlIZWxwOiAiVG90YWwgbnVtYmVyIG9mIHVuaXF1ZSByZWZzIGNyZWF0ZWQuIiwKCX0pCgoJaWYgcmVnICE9IG5pbCB7CgkJcmVnLlJlZ2lzdGVyKGFjdGl2ZU1hcHBpbmdzKQoJCXJlZy5SZWdpc3Rlcih0cmFja2VkUmVmcykKCQlyZWcuUmVnaXN0ZXIocmVmc0NsZWFuZWQpCgkJcmVnLlJlZ2lzdGVyKHVuaXF1ZVJlZnNUb3RhbCkKCX0KCglyZXR1cm4gJlNlcmllc1JlZk1hcHBpbmdTdG9yZXsKCQl1bmlxdWVSZWZUb0NoaWxkUmVmczogbWFrZShtYXBbc3RvcmFnZS5TZXJpZXNSZWZddW5pcVJlZkNoaWxkcmVuKSwKCQluZXh0VW5pcXVlUmVmOiAgICAgICAgMSwKCQl1bmlxdWVSZWZUaW1lc3RhbXBzOiAgbWFrZShtYXBbc3RvcmFnZS5TZXJpZXNSZWZdaW50NjQpLAoJCWxhYmVsSGFzaFRvVW5pcXVlUmVmOiBtYWtlKG1hcFt1aW50NjRdc3RvcmFnZS5TZXJpZXNSZWYpLAoJCWNlbGxQb29sOiBzeW5jLlBvb2x7CgkJCU5ldzogZnVuYygpIGFueSB7CgkJCQlyZXR1cm4gJkNlbGx7UmVmczogbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAxMDApfQoJCQl9LAoJCX0sCgkJc3RvcENsZWFudXA6ICAgICBtYWtlKGNoYW4gc3RydWN0e30pLAoJCWNsZWFudXBTdG9wcGVkOiAgbWFrZShjaGFuIHN0cnVjdHt9KSwKCQlhY3RpdmVNYXBwaW5nczogIGFjdGl2ZU1hcHBpbmdzLAoJCXRyYWNrZWRSZWZzOiAgICAgdHJhY2tlZFJlZnMsCgkJcmVmc0NsZWFuZWQ6ICAgICByZWZzQ2xlYW5lZCwKCQl1bmlxdWVSZWZzVG90YWw6IHVuaXF1ZVJlZnNUb3RhbCwKCX0KfQoKdHlwZSBDZWxsIHN0cnVjdCB7CglSZWZzIFtdc3RvcmFnZS5TZXJpZXNSZWYKfQoKLy8gR2V0TWFwcGluZyByZXR1cm5zIGV4aXN0aW5nIGNoaWxkIHJlZiByZXN1bHRzIGZvciB0aGUgZ2l2ZW4gdW5pcXVlIHJlZiBpZiBvbmUgZXhpc3RzLgovLwovLyBJZiB0aGUgcGFzc2VkIHVuaXF1ZVJlZiBpcyB6ZXJvLCB0aGUgbWV0aG9kIHdpbGwgYXR0ZW1wdCB0byBmaW5kIGEgbWFwcGluZyB1c2luZyBwYXNzZWQgbGFiZWxzLgovLyBSZXR1cm5zIG5pbCBpZiBubyBtYXBwaW5nIGV4aXN0cy4KLy8KLy8gVGhlIHJldHVybmVkIHNsaWNlIG1heSBiZSBtb2RpZmllZCBieSB0aGUgY2FsbGVyLCBidXQgVXBkYXRlTWFwcGluZyBtdXN0IGJlIGNhbGxlZAovLyBhZnRlcndhcmRzIHRvIHBlcnNpc3QgY2hhbmdlcy4gTm90ZSB0aGF0IGNvbmN1cnJlbnQgYXBwZW5kZXJzIG1heSByYWNlIHRvIHVwZGF0ZSB0aGUKLy8gc2FtZSBtYXBwaW5nIHdpdGggZGlmZmVyZW50IHZhbHVlcywgd2hpY2ggaXMgc2FmZSBiZWNhdXNlIHN0YWxlIG1hcHBpbmdzIGFyZSBzZWxmLWNvcnJlY3RpbmcgLQovLyB1c2luZyBhIHN0YWxlIHJlZiB3aWxsIGNhdXNlIHRoZSBjaGlsZCBhcHBlbmRlciB0byByZXR1cm4gYSBuZXcgcmVmIG9uIHRoZSBuZXh0IGFwcGVuZC4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBHZXRNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSBbXXN0b3JhZ2UuU2VyaWVzUmVmIHsKCXMucmVmTWFwcGluZ011LlJMb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlJVbmxvY2soKQoKCWlmIHVuaXF1ZVJlZiA9PSAwIHsKCQkvLyBTb21lIGNvbnN1bWVycyBkb24ndCBtZW1vIHRoZSBnbG9iYWwgcmVmLiBUcnkgdG8gbG9va3VwIGEgcmVmIGJ5IGxhYmVsIGhhc2guCgkJbGFiZWxIYXNoIDo9IGxibHMuSGFzaCgpCgkJZ290UmVmLCBvayA6PSBzLmxhYmVsSGFzaFRvVW5pcXVlUmVmW2xhYmVsSGFzaF0KCQlpZiAhb2sgewoJCQlyZXR1cm4gbmlsCgkJfQoKCQl1bmlxdWVSZWYgPSBnb3RSZWYKCX0KCglpZiBtYXBwaW5nLCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl07IG9rIHsKCQlyZXR1cm4gKm1hcHBpbmcuY2hpbGRSZWZzCgl9CglyZXR1cm4gbmlsCn0KCi8vIENyZWF0ZU1hcHBpbmcgY3JlYXRlcyBhIG5ldyB1bmlxdWUgcmVmIG1hcHBpbmcgZm9yIHRoZSBnaXZlbiBjaGlsZCByZWYgcmVzdWx0cy4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBDcmVhdGVNYXBwaW5nKHJlZlJlc3VsdHMgW11zdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSBzdG9yYWdlLlNlcmllc1JlZiB7CgkvLyBTdGFydCBjbGVhbnVwIGdvcm91dGluZSBvbiBmaXJzdCBtYXBwaW5nCglzLnN0YXJ0UmVmQ2xlYW51cC5EbyhmdW5jKCkgewoJCXMuY2xlYW51cFN0YXJ0ZWQuU3RvcmUodHJ1ZSkKCQlnbyBzLmNsZWFudXBTdGFsZVJlZnMoKQoJfSkKCgkvLyBTdG9yZSBhIGNvcHkgb2YgdGhlIGNoaWxkIHJlZiByZXN1bHRzIGRpcmVjdGx5CgljaGlsZFJlZlNsaWNlIDo9IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgbGVuKHJlZlJlc3VsdHMpKQoJY29weShjaGlsZFJlZlNsaWNlLCByZWZSZXN1bHRzKQoKCS8vIEhhc2ggbGFiZWxzIHRvIGZvciB0aGUgZmFsbGJhY2sgbG9va3VwIHRhYmxlCglsYWJlbEhhc2ggOj0gbGJscy5IYXNoKCkKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJLy8gQ3JlYXRlIGEgbmV3IHVuaXF1ZSByZWYKCXVuaXF1ZVJlZiA6PSBzLm5leHRVbmlxdWVSZWYKCXMubmV4dFVuaXF1ZVJlZisrCgoJcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltsYWJlbEhhc2hdID0gdW5pcXVlUmVmCglzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0gPSB1bmlxUmVmQ2hpbGRyZW57CgkJY2hpbGRSZWZzOiAmY2hpbGRSZWZTbGljZSwKCQlsYWJlbEhhc2g6IGxhYmVsSGFzaCwKCX0KCglzLmFjdGl2ZU1hcHBpbmdzLkluYygpCglzLnVuaXF1ZVJlZnNUb3RhbC5JbmMoKQoKCXJldHVybiB1bmlxdWVSZWYKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHsKCWlmIHVuaXF1ZVJlZiA9PSAwIHsKCQlyZXR1cm4KCX0KCgljaGlsZFJlZlNsaWNlIDo9IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgbGVuKHJlZlJlc3VsdHMpKQoJY29weShjaGlsZFJlZlNsaWNlLCByZWZSZXN1bHRzKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgkvLyBFbnN1cmUgdGhhdCBsYWJlbCBoYXNoIGluZGV4IGlzIHVwIHRvIGRhdGUgdG8gaGFuZGxlIHBvc3NpYmxlIGhhc2ggY29sbGlzaW9ucy4KCS8vIFRPRE86IGlzIHRoaXMgbmVjZXNzYXJ5PwoJbmV3SGFzaCA6PSBsYmxzLkhhc2goKQoJcHJldiwgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdCglpZiBvayAmJiBwcmV2LmxhYmVsSGFzaCAhPSBuZXdIYXNoIHsKCQlkZWxldGUocy5sYWJlbEhhc2hUb1VuaXF1ZVJlZiwgcHJldi5sYWJlbEhhc2gpCgkJcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltuZXdIYXNoXSA9IHVuaXF1ZVJlZgoJfQoKCXMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXSA9IHVuaXFSZWZDaGlsZHJlbnsKCQljaGlsZFJlZnM6ICZjaGlsZFJlZlNsaWNlLAoJCWxhYmVsSGFzaDogbGJscy5IYXNoKCksCgl9Cn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgVHJhY2tBcHBlbmRlZFNlcmllcyh0cyBpbnQ2NCwgY2VsbCAqQ2VsbCkgewoJcy50aW1lc3RhbXBUcmFja2luZ011LkxvY2soKQoJZGVmZXIgcy50aW1lc3RhbXBUcmFja2luZ011LlVubG9jaygpCgoJZm9yIF8sIHIgOj0gcmFuZ2UgY2VsbC5SZWZzIHsKCQlzLnVuaXF1ZVJlZlRpbWVzdGFtcHNbcl0gPSB0cwoJfQoKCXMudHJhY2tlZFJlZnMuU2V0KGZsb2F0NjQobGVuKHMudW5pcXVlUmVmVGltZXN0YW1wcykpKQoKCWNlbGwuUmVmcyA9IGNlbGwuUmVmc1s6MF0KCXMuY2VsbFBvb2wuUHV0KGNlbGwpCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwgewoJcmV0dXJuIHMuY2VsbFBvb2wuR2V0KCkuKCpDZWxsKQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIGNsZWFudXBTdGFsZVJlZnMoKSB7CglkZWZlciBjbG9zZShzLmNsZWFudXBTdG9wcGVkKQoKCXRpY2tlciA6PSB0aW1lLk5ld1RpY2tlcigxNSAqIHRpbWUuTWludXRlKQoJZGVmZXIgdGlja2VyLlN0b3AoKQoKCWZvciB7CgkJc2VsZWN0IHsKCQljYXNlIDwtdGlja2VyLkM6CgkJCWN1dG9mZlRpbWUgOj0gdGltZS5Ob3coKS5BZGQoLTE1ICogdGltZS5NaW51dGUpLlVuaXgoKQoKCQkJLy8gSG9sZCBib3RoIGxvY2tzIHRvIHByZXZlbnQgcmFjZSBjb25kaXRpb24gd2hlcmUgYSByZWYgY291bGQgYmUKCQkJLy8gYXBwZW5kZWQgYWZ0ZXIgd2UgZGVsZXRlIGl0IGZyb20gdW5pcXVlUmVmQ2VsbCBidXQgYmVmb3JlCgkJCS8vIHdlIGRlbGV0ZSBpdCBmcm9tIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzCgkJCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCQkJcy5yZWZNYXBwaW5nTXUuTG9jaygpCgoJCQlzdGFsZVJlZkNvdW50IDo9IDAKCQkJZm9yIHJlZiwgdHMgOj0gcmFuZ2Ugcy51bmlxdWVSZWZUaW1lc3RhbXBzIHsKCQkJCWlmIHRzIDwgY3V0b2ZmVGltZSB7CgkJCQkJc3RhbGVSZWZDb3VudCsrCgoJCQkJCXYsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbcmVmXQoJCQkJCWlmIG9rIHsKCQkJCQkJZGVsZXRlKHMubGFiZWxIYXNoVG9VbmlxdWVSZWYsIHYubGFiZWxIYXNoKQoJCQkJCX0KCgkJCQkJZGVsZXRlKHMudW5pcXVlUmVmVGltZXN0YW1wcywgcmVmKQoJCQkJCWRlbGV0ZShzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzLCByZWYpCgkJCQl9CgkJCX0KCgkJCS8vIFVwZGF0ZSBtZXRyaWNzCgkJCWlmIHN0YWxlUmVmQ291bnQgPiAwIHsKCQkJCXMucmVmc0NsZWFuZWQuQWRkKGZsb2F0NjQoc3RhbGVSZWZDb3VudCkpCgkJCQlzLmFjdGl2ZU1hcHBpbmdzLlN1YihmbG9hdDY0KHN0YWxlUmVmQ291bnQpKQoJCQkJcy50cmFja2VkUmVmcy5TZXQoZmxvYXQ2NChsZW4ocy51bmlxdWVSZWZUaW1lc3RhbXBzKSkpCgkJCX0KCgkJCXMucmVmTWFwcGluZ011LlVubG9jaygpCgkJCXMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCQljYXNlIDwtcy5zdG9wQ2xlYW51cDoKCQkJcmV0dXJuCgkJfQoJfQp9CgovLyBDbGVhciB3aWxsIGNsZWFyIGFsbCBpbnRlcm5hbCBtYXBwaW5ncyBhbmQgc3RvcCB0aGUgY2xlYW5lciBnb3JvdXRpbmUgaWYgaXQgaXMgcnVubmluZy4KLy8gSXQgaXMgc2FmZSB0byByZS11c2UgdGhlIHNhbWUgaW5zdGFuY2UgYWZ0ZXIgY2FsbGluZyBDbGVhci4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBDbGVhcigpIHsKCS8vIFN0b3AgdGhlIGNsZWFudXAgZ29yb3V0aW5lIGFuZCB3YWl0IGZvciBpdCB0byBiZSBzdG9wcGVkIHNvIHdlIGNhbgoJLy8gYXZvaWQgYSBwb3NzaWJsZSBkZWFkbG9jayB3aXRoIGNsZWFudXAgdGhhdCBhbHNvIGhvbGRzIGJvdGggbG9ja3MKCWlmIHMuY2xlYW51cFN0YXJ0ZWQuTG9hZCgpIHsKCQlzZWxlY3QgewoJCWNhc2UgPC1zLnN0b3BDbGVhbnVwOgoJCQkvLyBBbHJlYWR5IGNsb3NlZAoJCWRlZmF1bHQ6CgkJCWNsb3NlKHMuc3RvcENsZWFudXApCgkJCTwtcy5jbGVhbnVwU3RvcHBlZAoJCX0KCX0KCgkvLyBXZSBuZWVkIHRvIGhvbGQgYm90aCBsb2NrcyB0byBkbyB0aGlzIHNhZmVseSBhbmQgd2UgZG8gaXQgaW4gdGhlIHNhbWUgb3JkZXIgYXMKCS8vIGNsZWFudXBTdGFsZVJlZnMuIFdlIHN0b3BwZWQgYW5kIHdhaXRlZCBmb3IgdGhlIGJhY2tncm91bmQgd29ya2VyIHRoYXQgY2FsbHMgaXQKCS8vIHRvIGZpbmlzaCBidXQgc29tZSBleHRyYSBzYWZldHkgd29uJ3QgaHVydC4KCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCWRlZmVyIHMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgljbGVhcihzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzKQoJY2xlYXIocy51bmlxdWVSZWZUaW1lc3RhbXBzKQoKCS8vIHJlc2V0IHRoZSBwb29sCglzLmNlbGxQb29sID0gc3luYy5Qb29sewoJCU5ldzogZnVuYygpIGFueSB7CgkJCXJldHVybiAmQ2VsbHtSZWZzOiBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIDAsIDEwMCl9CgkJfSwKCX0KCgkvLyBOT1RFOiBXZSBkbyBOT1QgcmVzZXQgbmV4dFVuaXF1ZVJlZiBoZXJlLiBSZXNldHRpbmcgaXQgd291bGQgY2F1c2UgcmVmIGNvbGxpc2lvbnMKCS8vIHdpdGggY29tcG9uZW50cyBsaWtlIHByb21ldGhldXMuc2NyYXBlIHdoaWNoIHdpbGwga2VlcCByZS1zZW5kaW5nIHRoZSBzYW1lIGNhY2hlZCByZWZzLgoJLy8gV2UgY29udGludWUgaW5jcmVtZW50aW5nIHRvIGVuc3VyZSBhbGwgcmVmcyByZW1haW4gdW5pcXVlIGFjcm9zcyB0aGUgbGlmZXRpbWUgb2YgdGhlIHByb2Nlc3MuCgoJLy8gUmVzZXQgbWV0cmljcwoJcy5hY3RpdmVNYXBwaW5ncy5TZXQoMCkKCXMudHJhY2tlZFJlZnMuU2V0KDApCgoJLy8gUmVzZXQgY2hhbm5lbHMgYW5kIGZsYWdzCglzLnN0b3BDbGVhbnVwID0gbWFrZShjaGFuIHN0cnVjdHt9KQoJcy5jbGVhbnVwU3RvcHBlZCA9IG1ha2UoY2hhbiBzdHJ1Y3R7fSkKCXMuc3RhcnRSZWZDbGVhbnVwID0gc3luYy5PbmNle30KCXMuY2xlYW51cFN0YXJ0ZWQuU3RvcmUoZmFsc2UpCn0K"
    }
  }
}
```

Response (packages.DriverResponse):
Error: `err: context canceled: stderr: `

#### drv #6

Trace meta: spanId=11, ts=1770837171301, ts_iso=2026-02-11T19:12:51.301000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy/internal/validator"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkiZm10IgoJInN5bmMiCgkic3luYy9hdG9taWMiCgkidGltZSIKCgkiZ2l0aHViLmNvbS9oYXNoaWNvcnAvZ28tbXVsdGllcnJvciIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvY2xpZW50X2dvbGFuZy9wcm9tZXRoZXVzIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2V4ZW1wbGFyIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2hpc3RvZ3JhbSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9sYWJlbHMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvbWV0YWRhdGEiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvc3RvcmFnZSIKKQoKdHlwZSBNYXBwaW5nU3RvcmUgaW50ZXJmYWNlIHsKCUdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYKCUNyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmCglVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpCglUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKQoJR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwKfQoKdHlwZSBzZXJpZXNSZWZNYXBwaW5nIHN0cnVjdCB7CglzdGFydCAgICB0aW1lLlRpbWUKCWNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlcgoJc3RvcmUgICAgTWFwcGluZ1N0b3JlCgoJdW5pcXVlUmVmQ2VsbCAqQ2VsbAoKCS8vIGNoaWxkUmVmcyBpcyByZXVzZWQgZm9yIGVhY2ggYXBwZW5kIGNhbGwgdG8gYXZvaWQgYWxsb2NhdGlvbnMuIFRoaXMgaXMgc2FmZSBiZWNhdXNlIHN0b3JhZ2UuQXBwZW5kZXIgc2hvdWxkIG5ldmVyCgkvLyBoYXZlIGNvbmN1cnJlbnQgY2FsbHMgdG8gQXBwZW5kIG1ldGhvZHMuCgljaGlsZFJlZnMgICAgICAgIFtdc3RvcmFnZS5TZXJpZXNSZWYKCXdyaXRlTGF0ZW5jeSAgICAgcHJvbWV0aGV1cy5IaXN0b2dyYW0KCXNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZyhjaGlsZHJlbiBbXXN0b3JhZ2UuQXBwZW5kZXIsIHN0b3JlIE1hcHBpbmdTdG9yZSwgd3JpdGVMYXRlbmN5IHByb21ldGhldXMuSGlzdG9ncmFtLCBzYW1wbGVzRm9yd2FyZGVkIHByb21ldGhldXMuQ291bnRlcikgc3RvcmFnZS5BcHBlbmRlciB7Cgl1bmlxdWVSZWZDZWxsIDo9IHN0b3JlLkdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpCgoJcmV0dXJuICZzZXJpZXNSZWZNYXBwaW5newoJCWNoaWxkcmVuOiAgICAgICAgIGNoaWxkcmVuLAoJCXN0b3JlOiAgICAgICAgICAgIHN0b3JlLAoJCXdyaXRlTGF0ZW5jeTogICAgIHdyaXRlTGF0ZW5jeSwKCQlzYW1wbGVzRm9yd2FyZGVkOiBzYW1wbGVzRm9yd2FyZGVkLAoKCQl1bmlxdWVSZWZDZWxsOiB1bmlxdWVSZWZDZWxsLAoJCWNoaWxkUmVmczogICAgIG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgbGVuKGNoaWxkcmVuKSksCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFNldE9wdGlvbnMob3B0cyAqc3RvcmFnZS5BcHBlbmRPcHRpb25zKSB7CglmbXQuUHJpbnRsbigiYXMpCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQljLlNldE9wdGlvbnMob3B0cykKCX0KfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQ29tbWl0KCkgZXJyb3IgewoJcy5zdG9yZS5UcmFja0FwcGVuZGVkU2VyaWVzKHRpbWUuTm93KCkuVW5peCgpLCBzLnVuaXF1ZVJlZkNlbGwpCgoJdmFyIG11bHRpRXJyIGVycm9yCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQllcnIgOj0gYy5Db21taXQoKQoJCWlmIGVyciAhPSBuaWwgewoJCQltdWx0aUVyciA9IG11bHRpZXJyb3IuQXBwZW5kKG11bHRpRXJyLCBlcnIpCgkJfQoJfQoJcmV0dXJuIG11bHRpRXJyCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFJvbGxiYWNrKCkgZXJyb3IgewoJLy8gV2Ugc3RpbGwgdHJhY2sgcm9sbGVkIGJhY2sgc2VyaWVzIHNvIHdlIGNhbiBwcm9wZXJseQoJLy8gY2xlYW4gdXAgYW55IHNlcmllcyB0aGF0IHdhcyBhcHBlbmRlZAoJcy5zdG9yZS5UcmFja0FwcGVuZGVkU2VyaWVzKHRpbWUuTm93KCkuVW5peCgpLCBzLnVuaXF1ZVJlZkNlbGwpCgoJdmFyIG11bHRpRXJyIGVycm9yCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQllcnIgOj0gYy5Sb2xsYmFjaygpCgkJaWYgZXJyICE9IG5pbCB7CgkJCW11bHRpRXJyID0gbXVsdGllcnJvci5BcHBlbmQobXVsdGlFcnIsIGVycikKCQl9Cgl9CglyZXR1cm4gbXVsdGlFcnIKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgcmVjb3JkTGF0ZW5jeSgpIHsKCWlmIHMuc3RhcnQuSXNaZXJvKCkgewoJCXJldHVybgoJfQoJZHVyYXRpb24gOj0gdGltZS5TaW5jZShzLnN0YXJ0KQoJcy53cml0ZUxhdGVuY3kuT2JzZXJ2ZShkdXJhdGlvbi5TZWNvbmRzKCkpCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIHJlc2V0RmllbGRzKCkgewoJLy8gUmVzZXQgY2hpbGRSZWZzIHNsaWNlIGxlbmd0aCB0byAwIGZvciByZXVzZQoJcy5jaGlsZFJlZnMgPSBzLmNoaWxkUmVmc1s6MF0KfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0IGludDY0LCB2IGZsb2F0NjQpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCW5ld1JlZiwgZXJyIDo9IGFwcGVuZGVyLkFwcGVuZChyZWYsIGwsIHQsIHYpCgkJaWYgZXJyID09IG5pbCB7CgkJCXMuc2FtcGxlc0ZvcndhcmRlZC5JbmMoKQoJCX0KCQlyZXR1cm4gbmV3UmVmLCBlcnIKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEV4ZW1wbGFyKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCBlIGV4ZW1wbGFyLkV4ZW1wbGFyKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kRXhlbXBsYXIocmVmLCBsLCBlKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kSGlzdG9ncmFtKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0IGludDY0LCBoICpoaXN0b2dyYW0uSGlzdG9ncmFtLCBmaCAqaGlzdG9ncmFtLkZsb2F0SGlzdG9ncmFtKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kSGlzdG9ncmFtKHJlZiwgbCwgdCwgaCwgZmgpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRIaXN0b2dyYW1DVFplcm9TYW1wbGUocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQsIGN0IGludDY0LCBoICpoaXN0b2dyYW0uSGlzdG9ncmFtLCBmaCAqaGlzdG9ncmFtLkZsb2F0SGlzdG9ncmFtKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kSGlzdG9ncmFtQ1RaZXJvU2FtcGxlKHJlZiwgbCwgdCwgY3QsIGgsIGZoKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgVXBkYXRlTWV0YWRhdGEocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIG0gbWV0YWRhdGEuTWV0YWRhdGEpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5VcGRhdGVNZXRhZGF0YShyZWYsIGwsIG0pCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRDVFplcm9TYW1wbGUocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQsIGN0IGludDY0KSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kQ1RaZXJvU2FtcGxlKHJlZiwgbCwgdCwgY3QpCgl9KQp9Cgp0eXBlIGFwcGVuZGVyRnVuYyBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikKCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIGFwcGVuZFRvQ2hpbGRyZW4ocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMsIGFmIGFwcGVuZGVyRnVuYykgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJZGVmZXIgcy5yZXNldEZpZWxkcygpCgoJaWYgcy5zdGFydC5Jc1plcm8oKSB7CgkJcy5zdGFydCA9IHRpbWUuTm93KCkKCX0KCgkvLyBDaGVjayBpZiB0aGUgaW5jb21pbmcgcmVmIGhhcyByZWYgbWFwcGluZ3MKCWV4aXN0aW5nQ2hpbGRSZWZzIDo9IHMuc3RvcmUuR2V0TWFwcGluZyhyZWYsIGxibHMpCgoJdmFyIGFwcGVuZEVyciBlcnJvcgoKCS8vIFNhbml0eSBjaGVjazogaWYgd2UgaGF2ZSBleGlzdGluZyBjaGlsZCByZWZzLCB0aGV5IG11c3QgbWF0Y2ggdGhlIG51bWJlciBvZiBjaGlsZHJlbgoJaWYgZXhpc3RpbmdDaGlsZFJlZnMgIT0gbmlsICYmIGxlbihleGlzdGluZ0NoaWxkUmVmcykgPT0gbGVuKHMuY2hpbGRyZW4pIHsKCQlzLnVuaXF1ZVJlZkNlbGwuUmVmcyA9IGFwcGVuZChzLnVuaXF1ZVJlZkNlbGwuUmVmcywgcmVmKQoKCQlyZWZVcGRhdGVSZXF1aXJlZCA6PSBmYWxzZQoJCWZvciBjaGlsZEluZGV4LCBjaGlsZFJlZiA6PSByYW5nZSBleGlzdGluZ0NoaWxkUmVmcyB7CgkJCW5ld0NoaWxkUmVmLCBlcnIgOj0gYWYocy5jaGlsZHJlbltjaGlsZEluZGV4XSwgY2hpbGRSZWYpCgkJCWlmIGVyciAhPSBuaWwgewoJCQkJYXBwZW5kRXJyID0gbXVsdGllcnJvci5BcHBlbmQoYXBwZW5kRXJyLCBlcnIpCgkJCX0KCgkJCWlmIG5ld0NoaWxkUmVmICE9IGNoaWxkUmVmIHsKCQkJCS8vIENoaWxkIHJlZiBjaGFuZ2VkLCBuZWVkIHRvIHVwZGF0ZSBtYXBwaW5nCgkJCQlleGlzdGluZ0NoaWxkUmVmc1tjaGlsZEluZGV4XSA9IG5ld0NoaWxkUmVmCgkJCQlyZWZVcGRhdGVSZXF1aXJlZCA9IHRydWUKCQkJfQoJCX0KCgkJaWYgYXBwZW5kRXJyICE9IG5pbCB7CgkJCXJldHVybiAwLCBhcHBlbmRFcnIKCQl9CgoJCWlmIHJlZlVwZGF0ZVJlcXVpcmVkIHsKCQkJcy5zdG9yZS5VcGRhdGVNYXBwaW5nKHJlZiwgZXhpc3RpbmdDaGlsZFJlZnMsIGxibHMpCgkJfQoKCQlyZXR1cm4gcmVmLCBuaWwKCX0KCgkvLyBObyBleGlzdGluZyBtYXBwaW5nLCBwcm9jZWVkIHdpdGggbm9ybWFsIGFwcGVuZCB0byBhbGwgY2hpbGRyZW4KCXZhciBmaXJzdE5vblplcm9SZWYgc3RvcmFnZS5TZXJpZXNSZWYKCXZhciBub25aZXJvQ291bnQgaW50CgoJLy8gTm90ZTogdGhlcmUncyBhbm90aGVyIG9wdGltaXphdGlvbiB3aGVyZSB3ZSBjb3VsZCB1c2UgdGhlIHJldHVybmVkIHJlZiBpZiBhbGwgdGhlIG5vbiB6ZXJvIHJlZnMKCS8vICBhcmUgdGhlIHNhbWUgdmFsdWUuIFRoaXMgaXNuJ3Qgc2FmZSBhcyB3ZSB3aWxsIG1peCBkb3duc3RyZWFtIHJlZnMgd2l0aCB1bmlxdWUgcmVmcyB3aGljaCBjb3VsZAoJLy8gIGNvbGxpZGUuIFdlIGNvdWxkIHN0YXJ0IGF0IG1heCB1bml0NjQgZm9yIHVuaXF1ZSByZWZzIGFuZCBnbyBiYWNrd2FyZHMgbGVzc2VuaW5nIHRoZSBjaGFuY2Ugb2YKCS8vIAljb2xsaXNpb25zIGJ1dCBpdCdzIHJhdGhlciBkYW5nZXJvdXMgZm9yIGFuIHVubGlrZWx5IGVkZ2UgY2FzZS4gSWYgdHdvIGNvbXBvbmVudHMgYXJlIHJldHVybmluZwoJLy8gCXRoZSBzYW1lIHJlZiBpdCdzIHR3byByZW1vdGVfd3JpdGUgY29tcG9uZW50cyB3aGljaCBzaG91bGQgcHJvYmFibHkgYmUgbWVyZ2VkIGluIHRvIG9uZS4KCWZvciBfLCBjaGlsZCA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQljaGlsZFJlZiwgZXJyIDo9IGFmKGNoaWxkLCByZWYpCgkJaWYgZXJyICE9IG5pbCB7CgkJCWFwcGVuZEVyciA9IG11bHRpZXJyb3IuQXBwZW5kKGFwcGVuZEVyciwgZXJyKQoKCQkJLy8gVE9ETyBzaG91bGQgSSBtYWtlIHRoZSBjaGlsZFJlZiB6ZXJvIGhlcmU/CgkJfQoKCQlzLmNoaWxkUmVmcyA9IGFwcGVuZChzLmNoaWxkUmVmcywgY2hpbGRSZWYpCgkJaWYgY2hpbGRSZWYgIT0gMCB7CgkJCWlmIGZpcnN0Tm9uWmVyb1JlZiA9PSAwIHsKCQkJCWZpcnN0Tm9uWmVyb1JlZiA9IGNoaWxkUmVmCgkJCX0KCQkJbm9uWmVyb0NvdW50KysKCQl9Cgl9CgoJaWYgYXBwZW5kRXJyICE9IG5pbCB7CgkJcmV0dXJuIDAsIGFwcGVuZEVycgoJfQoKCWlmIG5vblplcm9Db3VudCA9PSAwIHsKCQkvLyBBbGwgY2hpbGRyZW4gcmV0dXJuZWQgcmVmIDAsIHNvIHJldHVybiB0aGUgaW5wdXQgcmVmCgkJcmV0dXJuIHJlZiwgbmlsCgl9CgoJLy8gT25seSBvbmUgY2hpbGQgcmV0dXJuZWQgYSBub24temVybyByZWYsIHVzZSB0aGF0CglpZiBub25aZXJvQ291bnQgPT0gMSB7CgkJcmV0dXJuIGZpcnN0Tm9uWmVyb1JlZiwgbmlsCgl9CgoJLy8gV2UgZ290IGRpZmZlcmVudCByZWZzIGJhY2sgYW5kIG5lZWQgdG8gY3JlYXRlIGEgbmV3IG1hcHBpbmcKCXVuaXF1ZVJlZiA6PSBzLnN0b3JlLkNyZWF0ZU1hcHBpbmcocy5jaGlsZFJlZnMsIGxibHMpCglzLnVuaXF1ZVJlZkNlbGwuUmVmcyA9IGFwcGVuZChzLnVuaXF1ZVJlZkNlbGwuUmVmcywgdW5pcXVlUmVmKQoJcmV0dXJuIHVuaXF1ZVJlZiwgbmlsCn0KCnR5cGUgdW5pcVJlZkNoaWxkcmVuIHN0cnVjdCB7CgljaGlsZFJlZnMgKltdc3RvcmFnZS5TZXJpZXNSZWYKCWxhYmVsSGFzaCB1aW50NjQKfQoKdHlwZSBTZXJpZXNSZWZNYXBwaW5nU3RvcmUgc3RydWN0IHsKCS8vIHJlZk1hcHBpbmdNdSBwcm90ZWN0cyB1bmlxdWVSZWZUb0NoaWxkUmVmcywgbGFiZWxIYXNoVG9VbmlxdWVSZWYgYW5kIG5leHRVbmlxdWVSZWYKCXJlZk1hcHBpbmdNdSBzeW5jLlJXTXV0ZXgKCS8vIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzIG1hcHMgdGhlIHVuaXF1ZSByZWYgdG8gdGhlIGV4cGVjdGVkIGNoaWxkIHJlZiBpbiBvcmRlcgoJdW5pcXVlUmVmVG9DaGlsZFJlZnMgbWFwW3N0b3JhZ2UuU2VyaWVzUmVmXXVuaXFSZWZDaGlsZHJlbgoJLy8gbGFiZWxIYXNoVG9VbmlxdWVSZWYgbWFwcyB0aGUgbGFiZWwgaGFzaCB0byB1bmlxdWUgcmVmLgoJbGFiZWxIYXNoVG9VbmlxdWVSZWYgbWFwW3VpbnQ2NF1zdG9yYWdlLlNlcmllc1JlZgoKCS8vIG5leHRVbmlxdWVSZWYgaXMgdGhlIG5leHQgcmVmIElEIHdlIHdpbGwgaGFuZCBvdXQKCW5leHRVbmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYKCgkvLyB0aW1lc3RhbXBUcmFja2luZ011IHByb3RlY3RzIHVuaXF1ZVJlZlRpbWVzdGFtcHMgYW5kIGNlbGxQb29sCgl0aW1lc3RhbXBUcmFja2luZ011IHN5bmMuTXV0ZXgKCS8vIHVuaXF1ZVJlZlRpbWVzdGFtcHMgbWFwcyB1bmlxdWUgcmVmcyB0byB0aGVpciBsYXN0IGFwcGVuZCB0aW1lc3RhbXAKCXVuaXF1ZVJlZlRpbWVzdGFtcHMgbWFwW3N0b3JhZ2UuU2VyaWVzUmVmXWludDY0CgkvLyBjZWxsUG9vbCBpcyB1c2VkIHRvIHBvb2wgc2xpY2VzIG9mIFNlcmllc1JlZnMgdXNlZCBmb3IgdHJhY2tpbmcgdW5pcXVlIHJlZnMgaW4gVHJhY2tBcHBlbmRlZFNlcmllcy4KCWNlbGxQb29sIHN5bmMuUG9vbAoKCS8vIENsZWFudXAgZ29yb3V0aW5lIGNvb3JkaW5hdGlvbiAobm8gbG9jayByZXF1aXJlZCkKCXN0YXJ0UmVmQ2xlYW51cCBzeW5jLk9uY2UKCWNsZWFudXBTdGFydGVkICBhdG9taWMuQm9vbAoJc3RvcENsZWFudXAgICAgIGNoYW4gc3RydWN0e30KCWNsZWFudXBTdG9wcGVkICBjaGFuIHN0cnVjdHt9CgoJLy8gTWV0cmljcyAoc2FmZSBmb3IgY29uY3VycmVudCBhY2Nlc3MsIG5vIGxvY2sgcmVxdWlyZWQpCglhY3RpdmVNYXBwaW5ncyAgcHJvbWV0aGV1cy5HYXVnZQoJdHJhY2tlZFJlZnMgICAgIHByb21ldGhldXMuR2F1Z2UKCXJlZnNDbGVhbmVkICAgICBwcm9tZXRoZXVzLkNvdW50ZXIKCXVuaXF1ZVJlZnNUb3RhbCBwcm9tZXRoZXVzLkNvdW50ZXIKfQoKZnVuYyBOZXdTZXJpZXNSZWZNYXBwaW5nU3RvcmUocmVnIHByb21ldGhldXMuUmVnaXN0ZXJlcikgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSB7CglhY3RpdmVNYXBwaW5ncyA6PSBwcm9tZXRoZXVzLk5ld0dhdWdlKHByb21ldGhldXMuR2F1Z2VPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV9tYXBwaW5nc190b3RhbCIsCgkJSGVscDogIk51bWJlciBvZiBhY3RpdmUgdW5pcXVlIHJlZiBtYXBwaW5ncyBpbiB0aGUgc3RvcmUuIiwKCX0pCgl0cmFja2VkUmVmcyA6PSBwcm9tZXRoZXVzLk5ld0dhdWdlKHByb21ldGhldXMuR2F1Z2VPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV90cmFja2VkX3JlZnNfdG90YWwiLAoJCUhlbHA6ICJOdW1iZXIgb2YgcmVmcyBiZWluZyB0cmFja2VkIGZvciB0aW1lc3RhbXAtYmFzZWQgY2xlYW51cC4iLAoJfSkKCXJlZnNDbGVhbmVkIDo9IHByb21ldGhldXMuTmV3Q291bnRlcihwcm9tZXRoZXVzLkNvdW50ZXJPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV9yZWZzX2NsZWFuZWRfdG90YWwiLAoJCUhlbHA6ICJUb3RhbCBudW1iZXIgb2Ygc3RhbGUgcmVmcyBjbGVhbmVkIHVwIG92ZXIgdGltZS4iLAoJfSkKCXVuaXF1ZVJlZnNUb3RhbCA6PSBwcm9tZXRoZXVzLk5ld0NvdW50ZXIocHJvbWV0aGV1cy5Db3VudGVyT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfdW5pcXVlX3JlZnNfY3JlYXRlZF90b3RhbCIsCgkJSGVscDogIlRvdGFsIG51bWJlciBvZiB1bmlxdWUgcmVmcyBjcmVhdGVkLiIsCgl9KQoKCWlmIHJlZyAhPSBuaWwgewoJCXJlZy5SZWdpc3RlcihhY3RpdmVNYXBwaW5ncykKCQlyZWcuUmVnaXN0ZXIodHJhY2tlZFJlZnMpCgkJcmVnLlJlZ2lzdGVyKHJlZnNDbGVhbmVkKQoJCXJlZy5SZWdpc3Rlcih1bmlxdWVSZWZzVG90YWwpCgl9CgoJcmV0dXJuICZTZXJpZXNSZWZNYXBwaW5nU3RvcmV7CgkJdW5pcXVlUmVmVG9DaGlsZFJlZnM6IG1ha2UobWFwW3N0b3JhZ2UuU2VyaWVzUmVmXXVuaXFSZWZDaGlsZHJlbiksCgkJbmV4dFVuaXF1ZVJlZjogICAgICAgIDEsCgkJdW5pcXVlUmVmVGltZXN0YW1wczogIG1ha2UobWFwW3N0b3JhZ2UuU2VyaWVzUmVmXWludDY0KSwKCQlsYWJlbEhhc2hUb1VuaXF1ZVJlZjogbWFrZShtYXBbdWludDY0XXN0b3JhZ2UuU2VyaWVzUmVmKSwKCQljZWxsUG9vbDogc3luYy5Qb29sewoJCQlOZXc6IGZ1bmMoKSBhbnkgewoJCQkJcmV0dXJuICZDZWxse1JlZnM6IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMTAwKX0KCQkJfSwKCQl9LAoJCXN0b3BDbGVhbnVwOiAgICAgbWFrZShjaGFuIHN0cnVjdHt9KSwKCQljbGVhbnVwU3RvcHBlZDogIG1ha2UoY2hhbiBzdHJ1Y3R7fSksCgkJYWN0aXZlTWFwcGluZ3M6ICBhY3RpdmVNYXBwaW5ncywKCQl0cmFja2VkUmVmczogICAgIHRyYWNrZWRSZWZzLAoJCXJlZnNDbGVhbmVkOiAgICAgcmVmc0NsZWFuZWQsCgkJdW5pcXVlUmVmc1RvdGFsOiB1bmlxdWVSZWZzVG90YWwsCgl9Cn0KCnR5cGUgQ2VsbCBzdHJ1Y3QgewoJUmVmcyBbXXN0b3JhZ2UuU2VyaWVzUmVmCn0KCi8vIEdldE1hcHBpbmcgcmV0dXJucyBleGlzdGluZyBjaGlsZCByZWYgcmVzdWx0cyBmb3IgdGhlIGdpdmVuIHVuaXF1ZSByZWYgaWYgb25lIGV4aXN0cy4KLy8KLy8gSWYgdGhlIHBhc3NlZCB1bmlxdWVSZWYgaXMgemVybywgdGhlIG1ldGhvZCB3aWxsIGF0dGVtcHQgdG8gZmluZCBhIG1hcHBpbmcgdXNpbmcgcGFzc2VkIGxhYmVscy4KLy8gUmV0dXJucyBuaWwgaWYgbm8gbWFwcGluZyBleGlzdHMuCi8vCi8vIFRoZSByZXR1cm5lZCBzbGljZSBtYXkgYmUgbW9kaWZpZWQgYnkgdGhlIGNhbGxlciwgYnV0IFVwZGF0ZU1hcHBpbmcgbXVzdCBiZSBjYWxsZWQKLy8gYWZ0ZXJ3YXJkcyB0byBwZXJzaXN0IGNoYW5nZXMuIE5vdGUgdGhhdCBjb25jdXJyZW50IGFwcGVuZGVycyBtYXkgcmFjZSB0byB1cGRhdGUgdGhlCi8vIHNhbWUgbWFwcGluZyB3aXRoIGRpZmZlcmVudCB2YWx1ZXMsIHdoaWNoIGlzIHNhZmUgYmVjYXVzZSBzdGFsZSBtYXBwaW5ncyBhcmUgc2VsZi1jb3JyZWN0aW5nIC0KLy8gdXNpbmcgYSBzdGFsZSByZWYgd2lsbCBjYXVzZSB0aGUgY2hpbGQgYXBwZW5kZXIgdG8gcmV0dXJuIGEgbmV3IHJlZiBvbiB0aGUgbmV4dCBhcHBlbmQuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgR2V0TWFwcGluZyh1bmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgW11zdG9yYWdlLlNlcmllc1JlZiB7CglzLnJlZk1hcHBpbmdNdS5STG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5SVW5sb2NrKCkKCglpZiB1bmlxdWVSZWYgPT0gMCB7CgkJLy8gU29tZSBjb25zdW1lcnMgZG9uJ3QgbWVtbyB0aGUgZ2xvYmFsIHJlZi4gVHJ5IHRvIGxvb2t1cCBhIHJlZiBieSBsYWJlbCBoYXNoLgoJCWxhYmVsSGFzaCA6PSBsYmxzLkhhc2goKQoJCWdvdFJlZiwgb2sgOj0gcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltsYWJlbEhhc2hdCgkJaWYgIW9rIHsKCQkJcmV0dXJuIG5pbAoJCX0KCgkJdW5pcXVlUmVmID0gZ290UmVmCgl9CgoJaWYgbWFwcGluZywgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdOyBvayB7CgkJcmV0dXJuICptYXBwaW5nLmNoaWxkUmVmcwoJfQoJcmV0dXJuIG5pbAp9CgovLyBDcmVhdGVNYXBwaW5nIGNyZWF0ZXMgYSBuZXcgdW5pcXVlIHJlZiBtYXBwaW5nIGZvciB0aGUgZ2l2ZW4gY2hpbGQgcmVmIHJlc3VsdHMuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgQ3JlYXRlTWFwcGluZyhyZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgc3RvcmFnZS5TZXJpZXNSZWYgewoJLy8gU3RhcnQgY2xlYW51cCBnb3JvdXRpbmUgb24gZmlyc3QgbWFwcGluZwoJcy5zdGFydFJlZkNsZWFudXAuRG8oZnVuYygpIHsKCQlzLmNsZWFudXBTdGFydGVkLlN0b3JlKHRydWUpCgkJZ28gcy5jbGVhbnVwU3RhbGVSZWZzKCkKCX0pCgoJLy8gU3RvcmUgYSBjb3B5IG9mIHRoZSBjaGlsZCByZWYgcmVzdWx0cyBkaXJlY3RseQoJY2hpbGRSZWZTbGljZSA6PSBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxlbihyZWZSZXN1bHRzKSkKCWNvcHkoY2hpbGRSZWZTbGljZSwgcmVmUmVzdWx0cykKCgkvLyBIYXNoIGxhYmVscyB0byBmb3IgdGhlIGZhbGxiYWNrIGxvb2t1cCB0YWJsZQoJbGFiZWxIYXNoIDo9IGxibHMuSGFzaCgpCgoJcy5yZWZNYXBwaW5nTXUuTG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoKCS8vIENyZWF0ZSBhIG5ldyB1bmlxdWUgcmVmCgl1bmlxdWVSZWYgOj0gcy5uZXh0VW5pcXVlUmVmCglzLm5leHRVbmlxdWVSZWYrKwoKCXMubGFiZWxIYXNoVG9VbmlxdWVSZWZbbGFiZWxIYXNoXSA9IHVuaXF1ZVJlZgoJcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdID0gdW5pcVJlZkNoaWxkcmVuewoJCWNoaWxkUmVmczogJmNoaWxkUmVmU2xpY2UsCgkJbGFiZWxIYXNoOiBsYWJlbEhhc2gsCgl9CgoJcy5hY3RpdmVNYXBwaW5ncy5JbmMoKQoJcy51bmlxdWVSZWZzVG90YWwuSW5jKCkKCglyZXR1cm4gdW5pcXVlUmVmCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgVXBkYXRlTWFwcGluZyh1bmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYsIHJlZlJlc3VsdHMgW11zdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSB7CglpZiB1bmlxdWVSZWYgPT0gMCB7CgkJcmV0dXJuCgl9CgoJY2hpbGRSZWZTbGljZSA6PSBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxlbihyZWZSZXN1bHRzKSkKCWNvcHkoY2hpbGRSZWZTbGljZSwgcmVmUmVzdWx0cykKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJLy8gRW5zdXJlIHRoYXQgbGFiZWwgaGFzaCBpbmRleCBpcyB1cCB0byBkYXRlIHRvIGhhbmRsZSBwb3NzaWJsZSBoYXNoIGNvbGxpc2lvbnMuCgkvLyBUT0RPOiBpcyB0aGlzIG5lY2Vzc2FyeT8KCW5ld0hhc2ggOj0gbGJscy5IYXNoKCkKCXByZXYsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXQoJaWYgb2sgJiYgcHJldi5sYWJlbEhhc2ggIT0gbmV3SGFzaCB7CgkJZGVsZXRlKHMubGFiZWxIYXNoVG9VbmlxdWVSZWYsIHByZXYubGFiZWxIYXNoKQoJCXMubGFiZWxIYXNoVG9VbmlxdWVSZWZbbmV3SGFzaF0gPSB1bmlxdWVSZWYKCX0KCglzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0gPSB1bmlxUmVmQ2hpbGRyZW57CgkJY2hpbGRSZWZzOiAmY2hpbGRSZWZTbGljZSwKCQlsYWJlbEhhc2g6IGxibHMuSGFzaCgpLAoJfQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIFRyYWNrQXBwZW5kZWRTZXJpZXModHMgaW50NjQsIGNlbGwgKkNlbGwpIHsKCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCWRlZmVyIHMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCWZvciBfLCByIDo9IHJhbmdlIGNlbGwuUmVmcyB7CgkJcy51bmlxdWVSZWZUaW1lc3RhbXBzW3JdID0gdHMKCX0KCglzLnRyYWNrZWRSZWZzLlNldChmbG9hdDY0KGxlbihzLnVuaXF1ZVJlZlRpbWVzdGFtcHMpKSkKCgljZWxsLlJlZnMgPSBjZWxsLlJlZnNbOjBdCglzLmNlbGxQb29sLlB1dChjZWxsKQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIEdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpICpDZWxsIHsKCXJldHVybiBzLmNlbGxQb29sLkdldCgpLigqQ2VsbCkKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBjbGVhbnVwU3RhbGVSZWZzKCkgewoJZGVmZXIgY2xvc2Uocy5jbGVhbnVwU3RvcHBlZCkKCgl0aWNrZXIgOj0gdGltZS5OZXdUaWNrZXIoMTUgKiB0aW1lLk1pbnV0ZSkKCWRlZmVyIHRpY2tlci5TdG9wKCkKCglmb3IgewoJCXNlbGVjdCB7CgkJY2FzZSA8LXRpY2tlci5DOgoJCQljdXRvZmZUaW1lIDo9IHRpbWUuTm93KCkuQWRkKC0xNSAqIHRpbWUuTWludXRlKS5Vbml4KCkKCgkJCS8vIEhvbGQgYm90aCBsb2NrcyB0byBwcmV2ZW50IHJhY2UgY29uZGl0aW9uIHdoZXJlIGEgcmVmIGNvdWxkIGJlCgkJCS8vIGFwcGVuZGVkIGFmdGVyIHdlIGRlbGV0ZSBpdCBmcm9tIHVuaXF1ZVJlZkNlbGwgYnV0IGJlZm9yZQoJCQkvLyB3ZSBkZWxldGUgaXQgZnJvbSB1bmlxdWVSZWZUb0NoaWxkUmVmcwoJCQlzLnRpbWVzdGFtcFRyYWNraW5nTXUuTG9jaygpCgkJCXMucmVmTWFwcGluZ011LkxvY2soKQoKCQkJc3RhbGVSZWZDb3VudCA6PSAwCgkJCWZvciByZWYsIHRzIDo9IHJhbmdlIHMudW5pcXVlUmVmVGltZXN0YW1wcyB7CgkJCQlpZiB0cyA8IGN1dG9mZlRpbWUgewoJCQkJCXN0YWxlUmVmQ291bnQrKwoKCQkJCQl2LCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3JlZl0KCQkJCQlpZiBvayB7CgkJCQkJCWRlbGV0ZShzLmxhYmVsSGFzaFRvVW5pcXVlUmVmLCB2LmxhYmVsSGFzaCkKCQkJCQl9CgoJCQkJCWRlbGV0ZShzLnVuaXF1ZVJlZlRpbWVzdGFtcHMsIHJlZikKCQkJCQlkZWxldGUocy51bmlxdWVSZWZUb0NoaWxkUmVmcywgcmVmKQoJCQkJfQoJCQl9CgoJCQkvLyBVcGRhdGUgbWV0cmljcwoJCQlpZiBzdGFsZVJlZkNvdW50ID4gMCB7CgkJCQlzLnJlZnNDbGVhbmVkLkFkZChmbG9hdDY0KHN0YWxlUmVmQ291bnQpKQoJCQkJcy5hY3RpdmVNYXBwaW5ncy5TdWIoZmxvYXQ2NChzdGFsZVJlZkNvdW50KSkKCQkJCXMudHJhY2tlZFJlZnMuU2V0KGZsb2F0NjQobGVuKHMudW5pcXVlUmVmVGltZXN0YW1wcykpKQoJCQl9CgoJCQlzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoJCQlzLnRpbWVzdGFtcFRyYWNraW5nTXUuVW5sb2NrKCkKCgkJY2FzZSA8LXMuc3RvcENsZWFudXA6CgkJCXJldHVybgoJCX0KCX0KfQoKLy8gQ2xlYXIgd2lsbCBjbGVhciBhbGwgaW50ZXJuYWwgbWFwcGluZ3MgYW5kIHN0b3AgdGhlIGNsZWFuZXIgZ29yb3V0aW5lIGlmIGl0IGlzIHJ1bm5pbmcuCi8vIEl0IGlzIHNhZmUgdG8gcmUtdXNlIHRoZSBzYW1lIGluc3RhbmNlIGFmdGVyIGNhbGxpbmcgQ2xlYXIuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgQ2xlYXIoKSB7CgkvLyBTdG9wIHRoZSBjbGVhbnVwIGdvcm91dGluZSBhbmQgd2FpdCBmb3IgaXQgdG8gYmUgc3RvcHBlZCBzbyB3ZSBjYW4KCS8vIGF2b2lkIGEgcG9zc2libGUgZGVhZGxvY2sgd2l0aCBjbGVhbnVwIHRoYXQgYWxzbyBob2xkcyBib3RoIGxvY2tzCglpZiBzLmNsZWFudXBTdGFydGVkLkxvYWQoKSB7CgkJc2VsZWN0IHsKCQljYXNlIDwtcy5zdG9wQ2xlYW51cDoKCQkJLy8gQWxyZWFkeSBjbG9zZWQKCQlkZWZhdWx0OgoJCQljbG9zZShzLnN0b3BDbGVhbnVwKQoJCQk8LXMuY2xlYW51cFN0b3BwZWQKCQl9Cgl9CgoJLy8gV2UgbmVlZCB0byBob2xkIGJvdGggbG9ja3MgdG8gZG8gdGhpcyBzYWZlbHkgYW5kIHdlIGRvIGl0IGluIHRoZSBzYW1lIG9yZGVyIGFzCgkvLyBjbGVhbnVwU3RhbGVSZWZzLiBXZSBzdG9wcGVkIGFuZCB3YWl0ZWQgZm9yIHRoZSBiYWNrZ3JvdW5kIHdvcmtlciB0aGF0IGNhbGxzIGl0CgkvLyB0byBmaW5pc2ggYnV0IHNvbWUgZXh0cmEgc2FmZXR5IHdvbid0IGh1cnQuCglzLnRpbWVzdGFtcFRyYWNraW5nTXUuTG9jaygpCglkZWZlciBzLnRpbWVzdGFtcFRyYWNraW5nTXUuVW5sb2NrKCkKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJY2xlYXIocy51bmlxdWVSZWZUb0NoaWxkUmVmcykKCWNsZWFyKHMudW5pcXVlUmVmVGltZXN0YW1wcykKCgkvLyByZXNldCB0aGUgcG9vbAoJcy5jZWxsUG9vbCA9IHN5bmMuUG9vbHsKCQlOZXc6IGZ1bmMoKSBhbnkgewoJCQlyZXR1cm4gJkNlbGx7UmVmczogbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAwLCAxMDApfQoJCX0sCgl9CgoJLy8gTk9URTogV2UgZG8gTk9UIHJlc2V0IG5leHRVbmlxdWVSZWYgaGVyZS4gUmVzZXR0aW5nIGl0IHdvdWxkIGNhdXNlIHJlZiBjb2xsaXNpb25zCgkvLyB3aXRoIGNvbXBvbmVudHMgbGlrZSBwcm9tZXRoZXVzLnNjcmFwZSB3aGljaCB3aWxsIGtlZXAgcmUtc2VuZGluZyB0aGUgc2FtZSBjYWNoZWQgcmVmcy4KCS8vIFdlIGNvbnRpbnVlIGluY3JlbWVudGluZyB0byBlbnN1cmUgYWxsIHJlZnMgcmVtYWluIHVuaXF1ZSBhY3Jvc3MgdGhlIGxpZmV0aW1lIG9mIHRoZSBwcm9jZXNzLgoKCS8vIFJlc2V0IG1ldHJpY3MKCXMuYWN0aXZlTWFwcGluZ3MuU2V0KDApCglzLnRyYWNrZWRSZWZzLlNldCgwKQoKCS8vIFJlc2V0IGNoYW5uZWxzIGFuZCBmbGFncwoJcy5zdG9wQ2xlYW51cCA9IG1ha2UoY2hhbiBzdHJ1Y3R7fSkKCXMuY2xlYW51cFN0b3BwZWQgPSBtYWtlKGNoYW4gc3RydWN0e30pCglzLnN0YXJ0UmVmQ2xlYW51cCA9IHN5bmMuT25jZXt9CglzLmNsZWFudXBTdGFydGVkLlN0b3JlKGZhbHNlKQp9Cg=="
    }
  }
}
```

Response (packages.DriverResponse):
Error: `err: context canceled: stderr: `

#### drv #7

Trace meta: spanId=13, ts=1770837171385, ts_iso=2026-02-11T19:12:51.385000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy/internal/validator"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkiZm10IgoJInN5bmMiCgkic3luYy9hdG9taWMiCgkidGltZSIKCgkiZ2l0aHViLmNvbS9oYXNoaWNvcnAvZ28tbXVsdGllcnJvciIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvY2xpZW50X2dvbGFuZy9wcm9tZXRoZXVzIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2V4ZW1wbGFyIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2hpc3RvZ3JhbSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9sYWJlbHMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvbWV0YWRhdGEiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvc3RvcmFnZSIKKQoKdHlwZSBNYXBwaW5nU3RvcmUgaW50ZXJmYWNlIHsKCUdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYKCUNyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmCglVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpCglUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKQoJR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwKfQoKdHlwZSBzZXJpZXNSZWZNYXBwaW5nIHN0cnVjdCB7CglzdGFydCAgICB0aW1lLlRpbWUKCWNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlcgoJc3RvcmUgICAgTWFwcGluZ1N0b3JlCgoJdW5pcXVlUmVmQ2VsbCAqQ2VsbAoKCS8vIGNoaWxkUmVmcyBpcyByZXVzZWQgZm9yIGVhY2ggYXBwZW5kIGNhbGwgdG8gYXZvaWQgYWxsb2NhdGlvbnMuIFRoaXMgaXMgc2FmZSBiZWNhdXNlIHN0b3JhZ2UuQXBwZW5kZXIgc2hvdWxkIG5ldmVyCgkvLyBoYXZlIGNvbmN1cnJlbnQgY2FsbHMgdG8gQXBwZW5kIG1ldGhvZHMuCgljaGlsZFJlZnMgICAgICAgIFtdc3RvcmFnZS5TZXJpZXNSZWYKCXdyaXRlTGF0ZW5jeSAgICAgcHJvbWV0aGV1cy5IaXN0b2dyYW0KCXNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZyhjaGlsZHJlbiBbXXN0b3JhZ2UuQXBwZW5kZXIsIHN0b3JlIE1hcHBpbmdTdG9yZSwgd3JpdGVMYXRlbmN5IHByb21ldGhldXMuSGlzdG9ncmFtLCBzYW1wbGVzRm9yd2FyZGVkIHByb21ldGhldXMuQ291bnRlcikgc3RvcmFnZS5BcHBlbmRlciB7Cgl1bmlxdWVSZWZDZWxsIDo9IHN0b3JlLkdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpCgoJcmV0dXJuICZzZXJpZXNSZWZNYXBwaW5newoJCWNoaWxkcmVuOiAgICAgICAgIGNoaWxkcmVuLAoJCXN0b3JlOiAgICAgICAgICAgIHN0b3JlLAoJCXdyaXRlTGF0ZW5jeTogICAgIHdyaXRlTGF0ZW5jeSwKCQlzYW1wbGVzRm9yd2FyZGVkOiBzYW1wbGVzRm9yd2FyZGVkLAoKCQl1bmlxdWVSZWZDZWxsOiB1bmlxdWVSZWZDZWxsLAoJCWNoaWxkUmVmczogICAgIG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgbGVuKGNoaWxkcmVuKSksCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFNldE9wdGlvbnMob3B0cyAqc3RvcmFnZS5BcHBlbmRPcHRpb25zKSB7CglmbXQuUHJpbnRsbigiYXNkKQoJZm9yIF8sIGMgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJYy5TZXRPcHRpb25zKG9wdHMpCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIENvbW1pdCgpIGVycm9yIHsKCXMuc3RvcmUuVHJhY2tBcHBlbmRlZFNlcmllcyh0aW1lLk5vdygpLlVuaXgoKSwgcy51bmlxdWVSZWZDZWxsKQoKCXZhciBtdWx0aUVyciBlcnJvcgoJZm9yIF8sIGMgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJZXJyIDo9IGMuQ29tbWl0KCkKCQlpZiBlcnIgIT0gbmlsIHsKCQkJbXVsdGlFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChtdWx0aUVyciwgZXJyKQoJCX0KCX0KCXJldHVybiBtdWx0aUVycgp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBSb2xsYmFjaygpIGVycm9yIHsKCS8vIFdlIHN0aWxsIHRyYWNrIHJvbGxlZCBiYWNrIHNlcmllcyBzbyB3ZSBjYW4gcHJvcGVybHkKCS8vIGNsZWFuIHVwIGFueSBzZXJpZXMgdGhhdCB3YXMgYXBwZW5kZWQKCXMuc3RvcmUuVHJhY2tBcHBlbmRlZFNlcmllcyh0aW1lLk5vdygpLlVuaXgoKSwgcy51bmlxdWVSZWZDZWxsKQoKCXZhciBtdWx0aUVyciBlcnJvcgoJZm9yIF8sIGMgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJZXJyIDo9IGMuUm9sbGJhY2soKQoJCWlmIGVyciAhPSBuaWwgewoJCQltdWx0aUVyciA9IG11bHRpZXJyb3IuQXBwZW5kKG11bHRpRXJyLCBlcnIpCgkJfQoJfQoJcmV0dXJuIG11bHRpRXJyCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIHJlY29yZExhdGVuY3koKSB7CglpZiBzLnN0YXJ0LklzWmVybygpIHsKCQlyZXR1cm4KCX0KCWR1cmF0aW9uIDo9IHRpbWUuU2luY2Uocy5zdGFydCkKCXMud3JpdGVMYXRlbmN5Lk9ic2VydmUoZHVyYXRpb24uU2Vjb25kcygpKQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSByZXNldEZpZWxkcygpIHsKCS8vIFJlc2V0IGNoaWxkUmVmcyBzbGljZSBsZW5ndGggdG8gMCBmb3IgcmV1c2UKCXMuY2hpbGRSZWZzID0gcy5jaGlsZFJlZnNbOjBdCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZChyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCBpbnQ2NCwgdiBmbG9hdDY0KSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQluZXdSZWYsIGVyciA6PSBhcHBlbmRlci5BcHBlbmQocmVmLCBsLCB0LCB2KQoJCWlmIGVyciA9PSBuaWwgewoJCQlzLnNhbXBsZXNGb3J3YXJkZWQuSW5jKCkKCQl9CgkJcmV0dXJuIG5ld1JlZiwgZXJyCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRFeGVtcGxhcihyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgZSBleGVtcGxhci5FeGVtcGxhcikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZEV4ZW1wbGFyKHJlZiwgbCwgZSkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEhpc3RvZ3JhbShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCBpbnQ2NCwgaCAqaGlzdG9ncmFtLkhpc3RvZ3JhbSwgZmggKmhpc3RvZ3JhbS5GbG9hdEhpc3RvZ3JhbSkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZEhpc3RvZ3JhbShyZWYsIGwsIHQsIGgsIGZoKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kSGlzdG9ncmFtQ1RaZXJvU2FtcGxlKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0LCBjdCBpbnQ2NCwgaCAqaGlzdG9ncmFtLkhpc3RvZ3JhbSwgZmggKmhpc3RvZ3JhbS5GbG9hdEhpc3RvZ3JhbSkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZEhpc3RvZ3JhbUNUWmVyb1NhbXBsZShyZWYsIGwsIHQsIGN0LCBoLCBmaCkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFVwZGF0ZU1ldGFkYXRhKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCBtIG1ldGFkYXRhLk1ldGFkYXRhKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuVXBkYXRlTWV0YWRhdGEocmVmLCBsLCBtKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kQ1RaZXJvU2FtcGxlKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0LCBjdCBpbnQ2NCkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZENUWmVyb1NhbXBsZShyZWYsIGwsIHQsIGN0KQoJfSkKfQoKdHlwZSBhcHBlbmRlckZ1bmMgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpCgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBhcHBlbmRUb0NoaWxkcmVuKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzLCBhZiBhcHBlbmRlckZ1bmMpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCWRlZmVyIHMucmVzZXRGaWVsZHMoKQoKCWlmIHMuc3RhcnQuSXNaZXJvKCkgewoJCXMuc3RhcnQgPSB0aW1lLk5vdygpCgl9CgoJLy8gQ2hlY2sgaWYgdGhlIGluY29taW5nIHJlZiBoYXMgcmVmIG1hcHBpbmdzCglleGlzdGluZ0NoaWxkUmVmcyA6PSBzLnN0b3JlLkdldE1hcHBpbmcocmVmLCBsYmxzKQoKCXZhciBhcHBlbmRFcnIgZXJyb3IKCgkvLyBTYW5pdHkgY2hlY2s6IGlmIHdlIGhhdmUgZXhpc3RpbmcgY2hpbGQgcmVmcywgdGhleSBtdXN0IG1hdGNoIHRoZSBudW1iZXIgb2YgY2hpbGRyZW4KCWlmIGV4aXN0aW5nQ2hpbGRSZWZzICE9IG5pbCAmJiBsZW4oZXhpc3RpbmdDaGlsZFJlZnMpID09IGxlbihzLmNoaWxkcmVuKSB7CgkJcy51bmlxdWVSZWZDZWxsLlJlZnMgPSBhcHBlbmQocy51bmlxdWVSZWZDZWxsLlJlZnMsIHJlZikKCgkJcmVmVXBkYXRlUmVxdWlyZWQgOj0gZmFsc2UKCQlmb3IgY2hpbGRJbmRleCwgY2hpbGRSZWYgOj0gcmFuZ2UgZXhpc3RpbmdDaGlsZFJlZnMgewoJCQluZXdDaGlsZFJlZiwgZXJyIDo9IGFmKHMuY2hpbGRyZW5bY2hpbGRJbmRleF0sIGNoaWxkUmVmKQoJCQlpZiBlcnIgIT0gbmlsIHsKCQkJCWFwcGVuZEVyciA9IG11bHRpZXJyb3IuQXBwZW5kKGFwcGVuZEVyciwgZXJyKQoJCQl9CgoJCQlpZiBuZXdDaGlsZFJlZiAhPSBjaGlsZFJlZiB7CgkJCQkvLyBDaGlsZCByZWYgY2hhbmdlZCwgbmVlZCB0byB1cGRhdGUgbWFwcGluZwoJCQkJZXhpc3RpbmdDaGlsZFJlZnNbY2hpbGRJbmRleF0gPSBuZXdDaGlsZFJlZgoJCQkJcmVmVXBkYXRlUmVxdWlyZWQgPSB0cnVlCgkJCX0KCQl9CgoJCWlmIGFwcGVuZEVyciAhPSBuaWwgewoJCQlyZXR1cm4gMCwgYXBwZW5kRXJyCgkJfQoKCQlpZiByZWZVcGRhdGVSZXF1aXJlZCB7CgkJCXMuc3RvcmUuVXBkYXRlTWFwcGluZyhyZWYsIGV4aXN0aW5nQ2hpbGRSZWZzLCBsYmxzKQoJCX0KCgkJcmV0dXJuIHJlZiwgbmlsCgl9CgoJLy8gTm8gZXhpc3RpbmcgbWFwcGluZywgcHJvY2VlZCB3aXRoIG5vcm1hbCBhcHBlbmQgdG8gYWxsIGNoaWxkcmVuCgl2YXIgZmlyc3ROb25aZXJvUmVmIHN0b3JhZ2UuU2VyaWVzUmVmCgl2YXIgbm9uWmVyb0NvdW50IGludAoKCS8vIE5vdGU6IHRoZXJlJ3MgYW5vdGhlciBvcHRpbWl6YXRpb24gd2hlcmUgd2UgY291bGQgdXNlIHRoZSByZXR1cm5lZCByZWYgaWYgYWxsIHRoZSBub24gemVybyByZWZzCgkvLyAgYXJlIHRoZSBzYW1lIHZhbHVlLiBUaGlzIGlzbid0IHNhZmUgYXMgd2Ugd2lsbCBtaXggZG93bnN0cmVhbSByZWZzIHdpdGggdW5pcXVlIHJlZnMgd2hpY2ggY291bGQKCS8vICBjb2xsaWRlLiBXZSBjb3VsZCBzdGFydCBhdCBtYXggdW5pdDY0IGZvciB1bmlxdWUgcmVmcyBhbmQgZ28gYmFja3dhcmRzIGxlc3NlbmluZyB0aGUgY2hhbmNlIG9mCgkvLyAJY29sbGlzaW9ucyBidXQgaXQncyByYXRoZXIgZGFuZ2Vyb3VzIGZvciBhbiB1bmxpa2VseSBlZGdlIGNhc2UuIElmIHR3byBjb21wb25lbnRzIGFyZSByZXR1cm5pbmcKCS8vIAl0aGUgc2FtZSByZWYgaXQncyB0d28gcmVtb3RlX3dyaXRlIGNvbXBvbmVudHMgd2hpY2ggc2hvdWxkIHByb2JhYmx5IGJlIG1lcmdlZCBpbiB0byBvbmUuCglmb3IgXywgY2hpbGQgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJY2hpbGRSZWYsIGVyciA6PSBhZihjaGlsZCwgcmVmKQoJCWlmIGVyciAhPSBuaWwgewoJCQlhcHBlbmRFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChhcHBlbmRFcnIsIGVycikKCgkJCS8vIFRPRE8gc2hvdWxkIEkgbWFrZSB0aGUgY2hpbGRSZWYgemVybyBoZXJlPwoJCX0KCgkJcy5jaGlsZFJlZnMgPSBhcHBlbmQocy5jaGlsZFJlZnMsIGNoaWxkUmVmKQoJCWlmIGNoaWxkUmVmICE9IDAgewoJCQlpZiBmaXJzdE5vblplcm9SZWYgPT0gMCB7CgkJCQlmaXJzdE5vblplcm9SZWYgPSBjaGlsZFJlZgoJCQl9CgkJCW5vblplcm9Db3VudCsrCgkJfQoJfQoKCWlmIGFwcGVuZEVyciAhPSBuaWwgewoJCXJldHVybiAwLCBhcHBlbmRFcnIKCX0KCglpZiBub25aZXJvQ291bnQgPT0gMCB7CgkJLy8gQWxsIGNoaWxkcmVuIHJldHVybmVkIHJlZiAwLCBzbyByZXR1cm4gdGhlIGlucHV0IHJlZgoJCXJldHVybiByZWYsIG5pbAoJfQoKCS8vIE9ubHkgb25lIGNoaWxkIHJldHVybmVkIGEgbm9uLXplcm8gcmVmLCB1c2UgdGhhdAoJaWYgbm9uWmVyb0NvdW50ID09IDEgewoJCXJldHVybiBmaXJzdE5vblplcm9SZWYsIG5pbAoJfQoKCS8vIFdlIGdvdCBkaWZmZXJlbnQgcmVmcyBiYWNrIGFuZCBuZWVkIHRvIGNyZWF0ZSBhIG5ldyBtYXBwaW5nCgl1bmlxdWVSZWYgOj0gcy5zdG9yZS5DcmVhdGVNYXBwaW5nKHMuY2hpbGRSZWZzLCBsYmxzKQoJcy51bmlxdWVSZWZDZWxsLlJlZnMgPSBhcHBlbmQocy51bmlxdWVSZWZDZWxsLlJlZnMsIHVuaXF1ZVJlZikKCXJldHVybiB1bmlxdWVSZWYsIG5pbAp9Cgp0eXBlIHVuaXFSZWZDaGlsZHJlbiBzdHJ1Y3QgewoJY2hpbGRSZWZzICpbXXN0b3JhZ2UuU2VyaWVzUmVmCglsYWJlbEhhc2ggdWludDY0Cn0KCnR5cGUgU2VyaWVzUmVmTWFwcGluZ1N0b3JlIHN0cnVjdCB7CgkvLyByZWZNYXBwaW5nTXUgcHJvdGVjdHMgdW5pcXVlUmVmVG9DaGlsZFJlZnMsIGxhYmVsSGFzaFRvVW5pcXVlUmVmIGFuZCBuZXh0VW5pcXVlUmVmCglyZWZNYXBwaW5nTXUgc3luYy5SV011dGV4CgkvLyB1bmlxdWVSZWZUb0NoaWxkUmVmcyBtYXBzIHRoZSB1bmlxdWUgcmVmIHRvIHRoZSBleHBlY3RlZCBjaGlsZCByZWYgaW4gb3JkZXIKCXVuaXF1ZVJlZlRvQ2hpbGRSZWZzIG1hcFtzdG9yYWdlLlNlcmllc1JlZl11bmlxUmVmQ2hpbGRyZW4KCS8vIGxhYmVsSGFzaFRvVW5pcXVlUmVmIG1hcHMgdGhlIGxhYmVsIGhhc2ggdG8gdW5pcXVlIHJlZi4KCWxhYmVsSGFzaFRvVW5pcXVlUmVmIG1hcFt1aW50NjRdc3RvcmFnZS5TZXJpZXNSZWYKCgkvLyBuZXh0VW5pcXVlUmVmIGlzIHRoZSBuZXh0IHJlZiBJRCB3ZSB3aWxsIGhhbmQgb3V0CgluZXh0VW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmCgoJLy8gdGltZXN0YW1wVHJhY2tpbmdNdSBwcm90ZWN0cyB1bmlxdWVSZWZUaW1lc3RhbXBzIGFuZCBjZWxsUG9vbAoJdGltZXN0YW1wVHJhY2tpbmdNdSBzeW5jLk11dGV4CgkvLyB1bmlxdWVSZWZUaW1lc3RhbXBzIG1hcHMgdW5pcXVlIHJlZnMgdG8gdGhlaXIgbGFzdCBhcHBlbmQgdGltZXN0YW1wCgl1bmlxdWVSZWZUaW1lc3RhbXBzIG1hcFtzdG9yYWdlLlNlcmllc1JlZl1pbnQ2NAoJLy8gY2VsbFBvb2wgaXMgdXNlZCB0byBwb29sIHNsaWNlcyBvZiBTZXJpZXNSZWZzIHVzZWQgZm9yIHRyYWNraW5nIHVuaXF1ZSByZWZzIGluIFRyYWNrQXBwZW5kZWRTZXJpZXMuCgljZWxsUG9vbCBzeW5jLlBvb2wKCgkvLyBDbGVhbnVwIGdvcm91dGluZSBjb29yZGluYXRpb24gKG5vIGxvY2sgcmVxdWlyZWQpCglzdGFydFJlZkNsZWFudXAgc3luYy5PbmNlCgljbGVhbnVwU3RhcnRlZCAgYXRvbWljLkJvb2wKCXN0b3BDbGVhbnVwICAgICBjaGFuIHN0cnVjdHt9CgljbGVhbnVwU3RvcHBlZCAgY2hhbiBzdHJ1Y3R7fQoKCS8vIE1ldHJpY3MgKHNhZmUgZm9yIGNvbmN1cnJlbnQgYWNjZXNzLCBubyBsb2NrIHJlcXVpcmVkKQoJYWN0aXZlTWFwcGluZ3MgIHByb21ldGhldXMuR2F1Z2UKCXRyYWNrZWRSZWZzICAgICBwcm9tZXRoZXVzLkdhdWdlCglyZWZzQ2xlYW5lZCAgICAgcHJvbWV0aGV1cy5Db3VudGVyCgl1bmlxdWVSZWZzVG90YWwgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZ1N0b3JlKHJlZyBwcm9tZXRoZXVzLlJlZ2lzdGVyZXIpICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUgewoJYWN0aXZlTWFwcGluZ3MgOj0gcHJvbWV0aGV1cy5OZXdHYXVnZShwcm9tZXRoZXVzLkdhdWdlT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfbWFwcGluZ3NfdG90YWwiLAoJCUhlbHA6ICJOdW1iZXIgb2YgYWN0aXZlIHVuaXF1ZSByZWYgbWFwcGluZ3MgaW4gdGhlIHN0b3JlLiIsCgl9KQoJdHJhY2tlZFJlZnMgOj0gcHJvbWV0aGV1cy5OZXdHYXVnZShwcm9tZXRoZXVzLkdhdWdlT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfdHJhY2tlZF9yZWZzX3RvdGFsIiwKCQlIZWxwOiAiTnVtYmVyIG9mIHJlZnMgYmVpbmcgdHJhY2tlZCBmb3IgdGltZXN0YW1wLWJhc2VkIGNsZWFudXAuIiwKCX0pCglyZWZzQ2xlYW5lZCA6PSBwcm9tZXRoZXVzLk5ld0NvdW50ZXIocHJvbWV0aGV1cy5Db3VudGVyT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfcmVmc19jbGVhbmVkX3RvdGFsIiwKCQlIZWxwOiAiVG90YWwgbnVtYmVyIG9mIHN0YWxlIHJlZnMgY2xlYW5lZCB1cCBvdmVyIHRpbWUuIiwKCX0pCgl1bmlxdWVSZWZzVG90YWwgOj0gcHJvbWV0aGV1cy5OZXdDb3VudGVyKHByb21ldGhldXMuQ291bnRlck9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3VuaXF1ZV9yZWZzX2NyZWF0ZWRfdG90YWwiLAoJCUhlbHA6ICJUb3RhbCBudW1iZXIgb2YgdW5pcXVlIHJlZnMgY3JlYXRlZC4iLAoJfSkKCglpZiByZWcgIT0gbmlsIHsKCQlyZWcuUmVnaXN0ZXIoYWN0aXZlTWFwcGluZ3MpCgkJcmVnLlJlZ2lzdGVyKHRyYWNrZWRSZWZzKQoJCXJlZy5SZWdpc3RlcihyZWZzQ2xlYW5lZCkKCQlyZWcuUmVnaXN0ZXIodW5pcXVlUmVmc1RvdGFsKQoJfQoKCXJldHVybiAmU2VyaWVzUmVmTWFwcGluZ1N0b3JlewoJCXVuaXF1ZVJlZlRvQ2hpbGRSZWZzOiBtYWtlKG1hcFtzdG9yYWdlLlNlcmllc1JlZl11bmlxUmVmQ2hpbGRyZW4pLAoJCW5leHRVbmlxdWVSZWY6ICAgICAgICAxLAoJCXVuaXF1ZVJlZlRpbWVzdGFtcHM6ICBtYWtlKG1hcFtzdG9yYWdlLlNlcmllc1JlZl1pbnQ2NCksCgkJbGFiZWxIYXNoVG9VbmlxdWVSZWY6IG1ha2UobWFwW3VpbnQ2NF1zdG9yYWdlLlNlcmllc1JlZiksCgkJY2VsbFBvb2w6IHN5bmMuUG9vbHsKCQkJTmV3OiBmdW5jKCkgYW55IHsKCQkJCXJldHVybiAmQ2VsbHtSZWZzOiBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIDEwMCl9CgkJCX0sCgkJfSwKCQlzdG9wQ2xlYW51cDogICAgIG1ha2UoY2hhbiBzdHJ1Y3R7fSksCgkJY2xlYW51cFN0b3BwZWQ6ICBtYWtlKGNoYW4gc3RydWN0e30pLAoJCWFjdGl2ZU1hcHBpbmdzOiAgYWN0aXZlTWFwcGluZ3MsCgkJdHJhY2tlZFJlZnM6ICAgICB0cmFja2VkUmVmcywKCQlyZWZzQ2xlYW5lZDogICAgIHJlZnNDbGVhbmVkLAoJCXVuaXF1ZVJlZnNUb3RhbDogdW5pcXVlUmVmc1RvdGFsLAoJfQp9Cgp0eXBlIENlbGwgc3RydWN0IHsKCVJlZnMgW11zdG9yYWdlLlNlcmllc1JlZgp9CgovLyBHZXRNYXBwaW5nIHJldHVybnMgZXhpc3RpbmcgY2hpbGQgcmVmIHJlc3VsdHMgZm9yIHRoZSBnaXZlbiB1bmlxdWUgcmVmIGlmIG9uZSBleGlzdHMuCi8vCi8vIElmIHRoZSBwYXNzZWQgdW5pcXVlUmVmIGlzIHplcm8sIHRoZSBtZXRob2Qgd2lsbCBhdHRlbXB0IHRvIGZpbmQgYSBtYXBwaW5nIHVzaW5nIHBhc3NlZCBsYWJlbHMuCi8vIFJldHVybnMgbmlsIGlmIG5vIG1hcHBpbmcgZXhpc3RzLgovLwovLyBUaGUgcmV0dXJuZWQgc2xpY2UgbWF5IGJlIG1vZGlmaWVkIGJ5IHRoZSBjYWxsZXIsIGJ1dCBVcGRhdGVNYXBwaW5nIG11c3QgYmUgY2FsbGVkCi8vIGFmdGVyd2FyZHMgdG8gcGVyc2lzdCBjaGFuZ2VzLiBOb3RlIHRoYXQgY29uY3VycmVudCBhcHBlbmRlcnMgbWF5IHJhY2UgdG8gdXBkYXRlIHRoZQovLyBzYW1lIG1hcHBpbmcgd2l0aCBkaWZmZXJlbnQgdmFsdWVzLCB3aGljaCBpcyBzYWZlIGJlY2F1c2Ugc3RhbGUgbWFwcGluZ3MgYXJlIHNlbGYtY29ycmVjdGluZyAtCi8vIHVzaW5nIGEgc3RhbGUgcmVmIHdpbGwgY2F1c2UgdGhlIGNoaWxkIGFwcGVuZGVyIHRvIHJldHVybiBhIG5ldyByZWYgb24gdGhlIG5leHQgYXBwZW5kLgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIEdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYgewoJcy5yZWZNYXBwaW5nTXUuUkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuUlVubG9jaygpCgoJaWYgdW5pcXVlUmVmID09IDAgewoJCS8vIFNvbWUgY29uc3VtZXJzIGRvbid0IG1lbW8gdGhlIGdsb2JhbCByZWYuIFRyeSB0byBsb29rdXAgYSByZWYgYnkgbGFiZWwgaGFzaC4KCQlsYWJlbEhhc2ggOj0gbGJscy5IYXNoKCkKCQlnb3RSZWYsIG9rIDo9IHMubGFiZWxIYXNoVG9VbmlxdWVSZWZbbGFiZWxIYXNoXQoJCWlmICFvayB7CgkJCXJldHVybiBuaWwKCQl9CgoJCXVuaXF1ZVJlZiA9IGdvdFJlZgoJfQoKCWlmIG1hcHBpbmcsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXTsgb2sgewoJCXJldHVybiAqbWFwcGluZy5jaGlsZFJlZnMKCX0KCXJldHVybiBuaWwKfQoKLy8gQ3JlYXRlTWFwcGluZyBjcmVhdGVzIGEgbmV3IHVuaXF1ZSByZWYgbWFwcGluZyBmb3IgdGhlIGdpdmVuIGNoaWxkIHJlZiByZXN1bHRzLgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIENyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmIHsKCS8vIFN0YXJ0IGNsZWFudXAgZ29yb3V0aW5lIG9uIGZpcnN0IG1hcHBpbmcKCXMuc3RhcnRSZWZDbGVhbnVwLkRvKGZ1bmMoKSB7CgkJcy5jbGVhbnVwU3RhcnRlZC5TdG9yZSh0cnVlKQoJCWdvIHMuY2xlYW51cFN0YWxlUmVmcygpCgl9KQoKCS8vIFN0b3JlIGEgY29weSBvZiB0aGUgY2hpbGQgcmVmIHJlc3VsdHMgZGlyZWN0bHkKCWNoaWxkUmVmU2xpY2UgOj0gbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsZW4ocmVmUmVzdWx0cykpCgljb3B5KGNoaWxkUmVmU2xpY2UsIHJlZlJlc3VsdHMpCgoJLy8gSGFzaCBsYWJlbHMgdG8gZm9yIHRoZSBmYWxsYmFjayBsb29rdXAgdGFibGUKCWxhYmVsSGFzaCA6PSBsYmxzLkhhc2goKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgkvLyBDcmVhdGUgYSBuZXcgdW5pcXVlIHJlZgoJdW5pcXVlUmVmIDo9IHMubmV4dFVuaXF1ZVJlZgoJcy5uZXh0VW5pcXVlUmVmKysKCglzLmxhYmVsSGFzaFRvVW5pcXVlUmVmW2xhYmVsSGFzaF0gPSB1bmlxdWVSZWYKCXMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXSA9IHVuaXFSZWZDaGlsZHJlbnsKCQljaGlsZFJlZnM6ICZjaGlsZFJlZlNsaWNlLAoJCWxhYmVsSGFzaDogbGFiZWxIYXNoLAoJfQoKCXMuYWN0aXZlTWFwcGluZ3MuSW5jKCkKCXMudW5pcXVlUmVmc1RvdGFsLkluYygpCgoJcmV0dXJuIHVuaXF1ZVJlZgp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIFVwZGF0ZU1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCByZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgewoJaWYgdW5pcXVlUmVmID09IDAgewoJCXJldHVybgoJfQoKCWNoaWxkUmVmU2xpY2UgOj0gbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsZW4ocmVmUmVzdWx0cykpCgljb3B5KGNoaWxkUmVmU2xpY2UsIHJlZlJlc3VsdHMpCgoJcy5yZWZNYXBwaW5nTXUuTG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoKCS8vIEVuc3VyZSB0aGF0IGxhYmVsIGhhc2ggaW5kZXggaXMgdXAgdG8gZGF0ZSB0byBoYW5kbGUgcG9zc2libGUgaGFzaCBjb2xsaXNpb25zLgoJLy8gVE9ETzogaXMgdGhpcyBuZWNlc3Nhcnk/CgluZXdIYXNoIDo9IGxibHMuSGFzaCgpCglwcmV2LCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0KCWlmIG9rICYmIHByZXYubGFiZWxIYXNoICE9IG5ld0hhc2ggewoJCWRlbGV0ZShzLmxhYmVsSGFzaFRvVW5pcXVlUmVmLCBwcmV2LmxhYmVsSGFzaCkKCQlzLmxhYmVsSGFzaFRvVW5pcXVlUmVmW25ld0hhc2hdID0gdW5pcXVlUmVmCgl9CgoJcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdID0gdW5pcVJlZkNoaWxkcmVuewoJCWNoaWxkUmVmczogJmNoaWxkUmVmU2xpY2UsCgkJbGFiZWxIYXNoOiBsYmxzLkhhc2goKSwKCX0KfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKSB7CglzLnRpbWVzdGFtcFRyYWNraW5nTXUuTG9jaygpCglkZWZlciBzLnRpbWVzdGFtcFRyYWNraW5nTXUuVW5sb2NrKCkKCglmb3IgXywgciA6PSByYW5nZSBjZWxsLlJlZnMgewoJCXMudW5pcXVlUmVmVGltZXN0YW1wc1tyXSA9IHRzCgl9CgoJcy50cmFja2VkUmVmcy5TZXQoZmxvYXQ2NChsZW4ocy51bmlxdWVSZWZUaW1lc3RhbXBzKSkpCgoJY2VsbC5SZWZzID0gY2VsbC5SZWZzWzowXQoJcy5jZWxsUG9vbC5QdXQoY2VsbCkKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBHZXRDZWxsRm9yQXBwZW5kZWRTZXJpZXMoKSAqQ2VsbCB7CglyZXR1cm4gcy5jZWxsUG9vbC5HZXQoKS4oKkNlbGwpCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgY2xlYW51cFN0YWxlUmVmcygpIHsKCWRlZmVyIGNsb3NlKHMuY2xlYW51cFN0b3BwZWQpCgoJdGlja2VyIDo9IHRpbWUuTmV3VGlja2VyKDE1ICogdGltZS5NaW51dGUpCglkZWZlciB0aWNrZXIuU3RvcCgpCgoJZm9yIHsKCQlzZWxlY3QgewoJCWNhc2UgPC10aWNrZXIuQzoKCQkJY3V0b2ZmVGltZSA6PSB0aW1lLk5vdygpLkFkZCgtMTUgKiB0aW1lLk1pbnV0ZSkuVW5peCgpCgoJCQkvLyBIb2xkIGJvdGggbG9ja3MgdG8gcHJldmVudCByYWNlIGNvbmRpdGlvbiB3aGVyZSBhIHJlZiBjb3VsZCBiZQoJCQkvLyBhcHBlbmRlZCBhZnRlciB3ZSBkZWxldGUgaXQgZnJvbSB1bmlxdWVSZWZDZWxsIGJ1dCBiZWZvcmUKCQkJLy8gd2UgZGVsZXRlIGl0IGZyb20gdW5pcXVlUmVmVG9DaGlsZFJlZnMKCQkJcy50aW1lc3RhbXBUcmFja2luZ011LkxvY2soKQoJCQlzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCgkJCXN0YWxlUmVmQ291bnQgOj0gMAoJCQlmb3IgcmVmLCB0cyA6PSByYW5nZSBzLnVuaXF1ZVJlZlRpbWVzdGFtcHMgewoJCQkJaWYgdHMgPCBjdXRvZmZUaW1lIHsKCQkJCQlzdGFsZVJlZkNvdW50KysKCgkJCQkJdiwgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1tyZWZdCgkJCQkJaWYgb2sgewoJCQkJCQlkZWxldGUocy5sYWJlbEhhc2hUb1VuaXF1ZVJlZiwgdi5sYWJlbEhhc2gpCgkJCQkJfQoKCQkJCQlkZWxldGUocy51bmlxdWVSZWZUaW1lc3RhbXBzLCByZWYpCgkJCQkJZGVsZXRlKHMudW5pcXVlUmVmVG9DaGlsZFJlZnMsIHJlZikKCQkJCX0KCQkJfQoKCQkJLy8gVXBkYXRlIG1ldHJpY3MKCQkJaWYgc3RhbGVSZWZDb3VudCA+IDAgewoJCQkJcy5yZWZzQ2xlYW5lZC5BZGQoZmxvYXQ2NChzdGFsZVJlZkNvdW50KSkKCQkJCXMuYWN0aXZlTWFwcGluZ3MuU3ViKGZsb2F0NjQoc3RhbGVSZWZDb3VudCkpCgkJCQlzLnRyYWNrZWRSZWZzLlNldChmbG9hdDY0KGxlbihzLnVuaXF1ZVJlZlRpbWVzdGFtcHMpKSkKCQkJfQoKCQkJcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCQkJcy50aW1lc3RhbXBUcmFja2luZ011LlVubG9jaygpCgoJCWNhc2UgPC1zLnN0b3BDbGVhbnVwOgoJCQlyZXR1cm4KCQl9Cgl9Cn0KCi8vIENsZWFyIHdpbGwgY2xlYXIgYWxsIGludGVybmFsIG1hcHBpbmdzIGFuZCBzdG9wIHRoZSBjbGVhbmVyIGdvcm91dGluZSBpZiBpdCBpcyBydW5uaW5nLgovLyBJdCBpcyBzYWZlIHRvIHJlLXVzZSB0aGUgc2FtZSBpbnN0YW5jZSBhZnRlciBjYWxsaW5nIENsZWFyLgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIENsZWFyKCkgewoJLy8gU3RvcCB0aGUgY2xlYW51cCBnb3JvdXRpbmUgYW5kIHdhaXQgZm9yIGl0IHRvIGJlIHN0b3BwZWQgc28gd2UgY2FuCgkvLyBhdm9pZCBhIHBvc3NpYmxlIGRlYWRsb2NrIHdpdGggY2xlYW51cCB0aGF0IGFsc28gaG9sZHMgYm90aCBsb2NrcwoJaWYgcy5jbGVhbnVwU3RhcnRlZC5Mb2FkKCkgewoJCXNlbGVjdCB7CgkJY2FzZSA8LXMuc3RvcENsZWFudXA6CgkJCS8vIEFscmVhZHkgY2xvc2VkCgkJZGVmYXVsdDoKCQkJY2xvc2Uocy5zdG9wQ2xlYW51cCkKCQkJPC1zLmNsZWFudXBTdG9wcGVkCgkJfQoJfQoKCS8vIFdlIG5lZWQgdG8gaG9sZCBib3RoIGxvY2tzIHRvIGRvIHRoaXMgc2FmZWx5IGFuZCB3ZSBkbyBpdCBpbiB0aGUgc2FtZSBvcmRlciBhcwoJLy8gY2xlYW51cFN0YWxlUmVmcy4gV2Ugc3RvcHBlZCBhbmQgd2FpdGVkIGZvciB0aGUgYmFja2dyb3VuZCB3b3JrZXIgdGhhdCBjYWxscyBpdAoJLy8gdG8gZmluaXNoIGJ1dCBzb21lIGV4dHJhIHNhZmV0eSB3b24ndCBodXJ0LgoJcy50aW1lc3RhbXBUcmFja2luZ011LkxvY2soKQoJZGVmZXIgcy50aW1lc3RhbXBUcmFja2luZ011LlVubG9jaygpCgoJcy5yZWZNYXBwaW5nTXUuTG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoKCWNsZWFyKHMudW5pcXVlUmVmVG9DaGlsZFJlZnMpCgljbGVhcihzLnVuaXF1ZVJlZlRpbWVzdGFtcHMpCgoJLy8gcmVzZXQgdGhlIHBvb2wKCXMuY2VsbFBvb2wgPSBzeW5jLlBvb2x7CgkJTmV3OiBmdW5jKCkgYW55IHsKCQkJcmV0dXJuICZDZWxse1JlZnM6IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgMTAwKX0KCQl9LAoJfQoKCS8vIE5PVEU6IFdlIGRvIE5PVCByZXNldCBuZXh0VW5pcXVlUmVmIGhlcmUuIFJlc2V0dGluZyBpdCB3b3VsZCBjYXVzZSByZWYgY29sbGlzaW9ucwoJLy8gd2l0aCBjb21wb25lbnRzIGxpa2UgcHJvbWV0aGV1cy5zY3JhcGUgd2hpY2ggd2lsbCBrZWVwIHJlLXNlbmRpbmcgdGhlIHNhbWUgY2FjaGVkIHJlZnMuCgkvLyBXZSBjb250aW51ZSBpbmNyZW1lbnRpbmcgdG8gZW5zdXJlIGFsbCByZWZzIHJlbWFpbiB1bmlxdWUgYWNyb3NzIHRoZSBsaWZldGltZSBvZiB0aGUgcHJvY2Vzcy4KCgkvLyBSZXNldCBtZXRyaWNzCglzLmFjdGl2ZU1hcHBpbmdzLlNldCgwKQoJcy50cmFja2VkUmVmcy5TZXQoMCkKCgkvLyBSZXNldCBjaGFubmVscyBhbmQgZmxhZ3MKCXMuc3RvcENsZWFudXAgPSBtYWtlKGNoYW4gc3RydWN0e30pCglzLmNsZWFudXBTdG9wcGVkID0gbWFrZShjaGFuIHN0cnVjdHt9KQoJcy5zdGFydFJlZkNsZWFudXAgPSBzeW5jLk9uY2V7fQoJcy5jbGVhbnVwU3RhcnRlZC5TdG9yZShmYWxzZSkKfQo="
    }
  }
}
```

Response (packages.DriverResponse):
Error: `err: context canceled: stderr: `

#### drv #8

Trace meta: spanId=15, ts=1770837171546, ts_iso=2026-02-11T19:12:51.546000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy/internal/validator"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkiZm10IgoJInN5bmMiCgkic3luYy9hdG9taWMiCgkidGltZSIKCgkiZ2l0aHViLmNvbS9oYXNoaWNvcnAvZ28tbXVsdGllcnJvciIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvY2xpZW50X2dvbGFuZy9wcm9tZXRoZXVzIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2V4ZW1wbGFyIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2hpc3RvZ3JhbSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9sYWJlbHMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvbWV0YWRhdGEiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvc3RvcmFnZSIKKQoKdHlwZSBNYXBwaW5nU3RvcmUgaW50ZXJmYWNlIHsKCUdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYKCUNyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmCglVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpCglUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKQoJR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwKfQoKdHlwZSBzZXJpZXNSZWZNYXBwaW5nIHN0cnVjdCB7CglzdGFydCAgICB0aW1lLlRpbWUKCWNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlcgoJc3RvcmUgICAgTWFwcGluZ1N0b3JlCgoJdW5pcXVlUmVmQ2VsbCAqQ2VsbAoKCS8vIGNoaWxkUmVmcyBpcyByZXVzZWQgZm9yIGVhY2ggYXBwZW5kIGNhbGwgdG8gYXZvaWQgYWxsb2NhdGlvbnMuIFRoaXMgaXMgc2FmZSBiZWNhdXNlIHN0b3JhZ2UuQXBwZW5kZXIgc2hvdWxkIG5ldmVyCgkvLyBoYXZlIGNvbmN1cnJlbnQgY2FsbHMgdG8gQXBwZW5kIG1ldGhvZHMuCgljaGlsZFJlZnMgICAgICAgIFtdc3RvcmFnZS5TZXJpZXNSZWYKCXdyaXRlTGF0ZW5jeSAgICAgcHJvbWV0aGV1cy5IaXN0b2dyYW0KCXNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZyhjaGlsZHJlbiBbXXN0b3JhZ2UuQXBwZW5kZXIsIHN0b3JlIE1hcHBpbmdTdG9yZSwgd3JpdGVMYXRlbmN5IHByb21ldGhldXMuSGlzdG9ncmFtLCBzYW1wbGVzRm9yd2FyZGVkIHByb21ldGhldXMuQ291bnRlcikgc3RvcmFnZS5BcHBlbmRlciB7Cgl1bmlxdWVSZWZDZWxsIDo9IHN0b3JlLkdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpCgoJcmV0dXJuICZzZXJpZXNSZWZNYXBwaW5newoJCWNoaWxkcmVuOiAgICAgICAgIGNoaWxkcmVuLAoJCXN0b3JlOiAgICAgICAgICAgIHN0b3JlLAoJCXdyaXRlTGF0ZW5jeTogICAgIHdyaXRlTGF0ZW5jeSwKCQlzYW1wbGVzRm9yd2FyZGVkOiBzYW1wbGVzRm9yd2FyZGVkLAoKCQl1bmlxdWVSZWZDZWxsOiB1bmlxdWVSZWZDZWxsLAoJCWNoaWxkUmVmczogICAgIG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgbGVuKGNoaWxkcmVuKSksCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFNldE9wdGlvbnMob3B0cyAqc3RvcmFnZS5BcHBlbmRPcHRpb25zKSB7CglmbXQuUHJpbnRsbigiYXNkYSkKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWMuU2V0T3B0aW9ucyhvcHRzKQoJfQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBDb21taXQoKSBlcnJvciB7CglzLnN0b3JlLlRyYWNrQXBwZW5kZWRTZXJpZXModGltZS5Ob3coKS5Vbml4KCksIHMudW5pcXVlUmVmQ2VsbCkKCgl2YXIgbXVsdGlFcnIgZXJyb3IKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWVyciA6PSBjLkNvbW1pdCgpCgkJaWYgZXJyICE9IG5pbCB7CgkJCW11bHRpRXJyID0gbXVsdGllcnJvci5BcHBlbmQobXVsdGlFcnIsIGVycikKCQl9Cgl9CglyZXR1cm4gbXVsdGlFcnIKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgUm9sbGJhY2soKSBlcnJvciB7CgkvLyBXZSBzdGlsbCB0cmFjayByb2xsZWQgYmFjayBzZXJpZXMgc28gd2UgY2FuIHByb3Blcmx5CgkvLyBjbGVhbiB1cCBhbnkgc2VyaWVzIHRoYXQgd2FzIGFwcGVuZGVkCglzLnN0b3JlLlRyYWNrQXBwZW5kZWRTZXJpZXModGltZS5Ob3coKS5Vbml4KCksIHMudW5pcXVlUmVmQ2VsbCkKCgl2YXIgbXVsdGlFcnIgZXJyb3IKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWVyciA6PSBjLlJvbGxiYWNrKCkKCQlpZiBlcnIgIT0gbmlsIHsKCQkJbXVsdGlFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChtdWx0aUVyciwgZXJyKQoJCX0KCX0KCXJldHVybiBtdWx0aUVycgp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSByZWNvcmRMYXRlbmN5KCkgewoJaWYgcy5zdGFydC5Jc1plcm8oKSB7CgkJcmV0dXJuCgl9CglkdXJhdGlvbiA6PSB0aW1lLlNpbmNlKHMuc3RhcnQpCglzLndyaXRlTGF0ZW5jeS5PYnNlcnZlKGR1cmF0aW9uLlNlY29uZHMoKSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgcmVzZXRGaWVsZHMoKSB7CgkvLyBSZXNldCBjaGlsZFJlZnMgc2xpY2UgbGVuZ3RoIHRvIDAgZm9yIHJldXNlCglzLmNoaWxkUmVmcyA9IHMuY2hpbGRSZWZzWzowXQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmQocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQgaW50NjQsIHYgZmxvYXQ2NCkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJbmV3UmVmLCBlcnIgOj0gYXBwZW5kZXIuQXBwZW5kKHJlZiwgbCwgdCwgdikKCQlpZiBlcnIgPT0gbmlsIHsKCQkJcy5zYW1wbGVzRm9yd2FyZGVkLkluYygpCgkJfQoJCXJldHVybiBuZXdSZWYsIGVycgoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kRXhlbXBsYXIocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIGUgZXhlbXBsYXIuRXhlbXBsYXIpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRFeGVtcGxhcihyZWYsIGwsIGUpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRIaXN0b2dyYW0ocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQgaW50NjQsIGggKmhpc3RvZ3JhbS5IaXN0b2dyYW0sIGZoICpoaXN0b2dyYW0uRmxvYXRIaXN0b2dyYW0pIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRIaXN0b2dyYW0ocmVmLCBsLCB0LCBoLCBmaCkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEhpc3RvZ3JhbUNUWmVyb1NhbXBsZShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCwgY3QgaW50NjQsIGggKmhpc3RvZ3JhbS5IaXN0b2dyYW0sIGZoICpoaXN0b2dyYW0uRmxvYXRIaXN0b2dyYW0pIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRIaXN0b2dyYW1DVFplcm9TYW1wbGUocmVmLCBsLCB0LCBjdCwgaCwgZmgpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBVcGRhdGVNZXRhZGF0YShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgbSBtZXRhZGF0YS5NZXRhZGF0YSkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLlVwZGF0ZU1ldGFkYXRhKHJlZiwgbCwgbSkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZENUWmVyb1NhbXBsZShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCwgY3QgaW50NjQpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRDVFplcm9TYW1wbGUocmVmLCBsLCB0LCBjdCkKCX0pCn0KCnR5cGUgYXBwZW5kZXJGdW5jIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgYXBwZW5kVG9DaGlsZHJlbihyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscywgYWYgYXBwZW5kZXJGdW5jKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglkZWZlciBzLnJlc2V0RmllbGRzKCkKCglpZiBzLnN0YXJ0LklzWmVybygpIHsKCQlzLnN0YXJ0ID0gdGltZS5Ob3coKQoJfQoKCS8vIENoZWNrIGlmIHRoZSBpbmNvbWluZyByZWYgaGFzIHJlZiBtYXBwaW5ncwoJZXhpc3RpbmdDaGlsZFJlZnMgOj0gcy5zdG9yZS5HZXRNYXBwaW5nKHJlZiwgbGJscykKCgl2YXIgYXBwZW5kRXJyIGVycm9yCgoJLy8gU2FuaXR5IGNoZWNrOiBpZiB3ZSBoYXZlIGV4aXN0aW5nIGNoaWxkIHJlZnMsIHRoZXkgbXVzdCBtYXRjaCB0aGUgbnVtYmVyIG9mIGNoaWxkcmVuCglpZiBleGlzdGluZ0NoaWxkUmVmcyAhPSBuaWwgJiYgbGVuKGV4aXN0aW5nQ2hpbGRSZWZzKSA9PSBsZW4ocy5jaGlsZHJlbikgewoJCXMudW5pcXVlUmVmQ2VsbC5SZWZzID0gYXBwZW5kKHMudW5pcXVlUmVmQ2VsbC5SZWZzLCByZWYpCgoJCXJlZlVwZGF0ZVJlcXVpcmVkIDo9IGZhbHNlCgkJZm9yIGNoaWxkSW5kZXgsIGNoaWxkUmVmIDo9IHJhbmdlIGV4aXN0aW5nQ2hpbGRSZWZzIHsKCQkJbmV3Q2hpbGRSZWYsIGVyciA6PSBhZihzLmNoaWxkcmVuW2NoaWxkSW5kZXhdLCBjaGlsZFJlZikKCQkJaWYgZXJyICE9IG5pbCB7CgkJCQlhcHBlbmRFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChhcHBlbmRFcnIsIGVycikKCQkJfQoKCQkJaWYgbmV3Q2hpbGRSZWYgIT0gY2hpbGRSZWYgewoJCQkJLy8gQ2hpbGQgcmVmIGNoYW5nZWQsIG5lZWQgdG8gdXBkYXRlIG1hcHBpbmcKCQkJCWV4aXN0aW5nQ2hpbGRSZWZzW2NoaWxkSW5kZXhdID0gbmV3Q2hpbGRSZWYKCQkJCXJlZlVwZGF0ZVJlcXVpcmVkID0gdHJ1ZQoJCQl9CgkJfQoKCQlpZiBhcHBlbmRFcnIgIT0gbmlsIHsKCQkJcmV0dXJuIDAsIGFwcGVuZEVycgoJCX0KCgkJaWYgcmVmVXBkYXRlUmVxdWlyZWQgewoJCQlzLnN0b3JlLlVwZGF0ZU1hcHBpbmcocmVmLCBleGlzdGluZ0NoaWxkUmVmcywgbGJscykKCQl9CgoJCXJldHVybiByZWYsIG5pbAoJfQoKCS8vIE5vIGV4aXN0aW5nIG1hcHBpbmcsIHByb2NlZWQgd2l0aCBub3JtYWwgYXBwZW5kIHRvIGFsbCBjaGlsZHJlbgoJdmFyIGZpcnN0Tm9uWmVyb1JlZiBzdG9yYWdlLlNlcmllc1JlZgoJdmFyIG5vblplcm9Db3VudCBpbnQKCgkvLyBOb3RlOiB0aGVyZSdzIGFub3RoZXIgb3B0aW1pemF0aW9uIHdoZXJlIHdlIGNvdWxkIHVzZSB0aGUgcmV0dXJuZWQgcmVmIGlmIGFsbCB0aGUgbm9uIHplcm8gcmVmcwoJLy8gIGFyZSB0aGUgc2FtZSB2YWx1ZS4gVGhpcyBpc24ndCBzYWZlIGFzIHdlIHdpbGwgbWl4IGRvd25zdHJlYW0gcmVmcyB3aXRoIHVuaXF1ZSByZWZzIHdoaWNoIGNvdWxkCgkvLyAgY29sbGlkZS4gV2UgY291bGQgc3RhcnQgYXQgbWF4IHVuaXQ2NCBmb3IgdW5pcXVlIHJlZnMgYW5kIGdvIGJhY2t3YXJkcyBsZXNzZW5pbmcgdGhlIGNoYW5jZSBvZgoJLy8gCWNvbGxpc2lvbnMgYnV0IGl0J3MgcmF0aGVyIGRhbmdlcm91cyBmb3IgYW4gdW5saWtlbHkgZWRnZSBjYXNlLiBJZiB0d28gY29tcG9uZW50cyBhcmUgcmV0dXJuaW5nCgkvLyAJdGhlIHNhbWUgcmVmIGl0J3MgdHdvIHJlbW90ZV93cml0ZSBjb21wb25lbnRzIHdoaWNoIHNob3VsZCBwcm9iYWJseSBiZSBtZXJnZWQgaW4gdG8gb25lLgoJZm9yIF8sIGNoaWxkIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWNoaWxkUmVmLCBlcnIgOj0gYWYoY2hpbGQsIHJlZikKCQlpZiBlcnIgIT0gbmlsIHsKCQkJYXBwZW5kRXJyID0gbXVsdGllcnJvci5BcHBlbmQoYXBwZW5kRXJyLCBlcnIpCgoJCQkvLyBUT0RPIHNob3VsZCBJIG1ha2UgdGhlIGNoaWxkUmVmIHplcm8gaGVyZT8KCQl9CgoJCXMuY2hpbGRSZWZzID0gYXBwZW5kKHMuY2hpbGRSZWZzLCBjaGlsZFJlZikKCQlpZiBjaGlsZFJlZiAhPSAwIHsKCQkJaWYgZmlyc3ROb25aZXJvUmVmID09IDAgewoJCQkJZmlyc3ROb25aZXJvUmVmID0gY2hpbGRSZWYKCQkJfQoJCQlub25aZXJvQ291bnQrKwoJCX0KCX0KCglpZiBhcHBlbmRFcnIgIT0gbmlsIHsKCQlyZXR1cm4gMCwgYXBwZW5kRXJyCgl9CgoJaWYgbm9uWmVyb0NvdW50ID09IDAgewoJCS8vIEFsbCBjaGlsZHJlbiByZXR1cm5lZCByZWYgMCwgc28gcmV0dXJuIHRoZSBpbnB1dCByZWYKCQlyZXR1cm4gcmVmLCBuaWwKCX0KCgkvLyBPbmx5IG9uZSBjaGlsZCByZXR1cm5lZCBhIG5vbi16ZXJvIHJlZiwgdXNlIHRoYXQKCWlmIG5vblplcm9Db3VudCA9PSAxIHsKCQlyZXR1cm4gZmlyc3ROb25aZXJvUmVmLCBuaWwKCX0KCgkvLyBXZSBnb3QgZGlmZmVyZW50IHJlZnMgYmFjayBhbmQgbmVlZCB0byBjcmVhdGUgYSBuZXcgbWFwcGluZwoJdW5pcXVlUmVmIDo9IHMuc3RvcmUuQ3JlYXRlTWFwcGluZyhzLmNoaWxkUmVmcywgbGJscykKCXMudW5pcXVlUmVmQ2VsbC5SZWZzID0gYXBwZW5kKHMudW5pcXVlUmVmQ2VsbC5SZWZzLCB1bmlxdWVSZWYpCglyZXR1cm4gdW5pcXVlUmVmLCBuaWwKfQoKdHlwZSB1bmlxUmVmQ2hpbGRyZW4gc3RydWN0IHsKCWNoaWxkUmVmcyAqW11zdG9yYWdlLlNlcmllc1JlZgoJbGFiZWxIYXNoIHVpbnQ2NAp9Cgp0eXBlIFNlcmllc1JlZk1hcHBpbmdTdG9yZSBzdHJ1Y3QgewoJLy8gcmVmTWFwcGluZ011IHByb3RlY3RzIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzLCBsYWJlbEhhc2hUb1VuaXF1ZVJlZiBhbmQgbmV4dFVuaXF1ZVJlZgoJcmVmTWFwcGluZ011IHN5bmMuUldNdXRleAoJLy8gdW5pcXVlUmVmVG9DaGlsZFJlZnMgbWFwcyB0aGUgdW5pcXVlIHJlZiB0byB0aGUgZXhwZWN0ZWQgY2hpbGQgcmVmIGluIG9yZGVyCgl1bmlxdWVSZWZUb0NoaWxkUmVmcyBtYXBbc3RvcmFnZS5TZXJpZXNSZWZddW5pcVJlZkNoaWxkcmVuCgkvLyBsYWJlbEhhc2hUb1VuaXF1ZVJlZiBtYXBzIHRoZSBsYWJlbCBoYXNoIHRvIHVuaXF1ZSByZWYuCglsYWJlbEhhc2hUb1VuaXF1ZVJlZiBtYXBbdWludDY0XXN0b3JhZ2UuU2VyaWVzUmVmCgoJLy8gbmV4dFVuaXF1ZVJlZiBpcyB0aGUgbmV4dCByZWYgSUQgd2Ugd2lsbCBoYW5kIG91dAoJbmV4dFVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZgoKCS8vIHRpbWVzdGFtcFRyYWNraW5nTXUgcHJvdGVjdHMgdW5pcXVlUmVmVGltZXN0YW1wcyBhbmQgY2VsbFBvb2wKCXRpbWVzdGFtcFRyYWNraW5nTXUgc3luYy5NdXRleAoJLy8gdW5pcXVlUmVmVGltZXN0YW1wcyBtYXBzIHVuaXF1ZSByZWZzIHRvIHRoZWlyIGxhc3QgYXBwZW5kIHRpbWVzdGFtcAoJdW5pcXVlUmVmVGltZXN0YW1wcyBtYXBbc3RvcmFnZS5TZXJpZXNSZWZdaW50NjQKCS8vIGNlbGxQb29sIGlzIHVzZWQgdG8gcG9vbCBzbGljZXMgb2YgU2VyaWVzUmVmcyB1c2VkIGZvciB0cmFja2luZyB1bmlxdWUgcmVmcyBpbiBUcmFja0FwcGVuZGVkU2VyaWVzLgoJY2VsbFBvb2wgc3luYy5Qb29sCgoJLy8gQ2xlYW51cCBnb3JvdXRpbmUgY29vcmRpbmF0aW9uIChubyBsb2NrIHJlcXVpcmVkKQoJc3RhcnRSZWZDbGVhbnVwIHN5bmMuT25jZQoJY2xlYW51cFN0YXJ0ZWQgIGF0b21pYy5Cb29sCglzdG9wQ2xlYW51cCAgICAgY2hhbiBzdHJ1Y3R7fQoJY2xlYW51cFN0b3BwZWQgIGNoYW4gc3RydWN0e30KCgkvLyBNZXRyaWNzIChzYWZlIGZvciBjb25jdXJyZW50IGFjY2Vzcywgbm8gbG9jayByZXF1aXJlZCkKCWFjdGl2ZU1hcHBpbmdzICBwcm9tZXRoZXVzLkdhdWdlCgl0cmFja2VkUmVmcyAgICAgcHJvbWV0aGV1cy5HYXVnZQoJcmVmc0NsZWFuZWQgICAgIHByb21ldGhldXMuQ291bnRlcgoJdW5pcXVlUmVmc1RvdGFsIHByb21ldGhldXMuQ291bnRlcgp9CgpmdW5jIE5ld1Nlcmllc1JlZk1hcHBpbmdTdG9yZShyZWcgcHJvbWV0aGV1cy5SZWdpc3RlcmVyKSAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlIHsKCWFjdGl2ZU1hcHBpbmdzIDo9IHByb21ldGhldXMuTmV3R2F1Z2UocHJvbWV0aGV1cy5HYXVnZU9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX21hcHBpbmdzX3RvdGFsIiwKCQlIZWxwOiAiTnVtYmVyIG9mIGFjdGl2ZSB1bmlxdWUgcmVmIG1hcHBpbmdzIGluIHRoZSBzdG9yZS4iLAoJfSkKCXRyYWNrZWRSZWZzIDo9IHByb21ldGhldXMuTmV3R2F1Z2UocHJvbWV0aGV1cy5HYXVnZU9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3RyYWNrZWRfcmVmc190b3RhbCIsCgkJSGVscDogIk51bWJlciBvZiByZWZzIGJlaW5nIHRyYWNrZWQgZm9yIHRpbWVzdGFtcC1iYXNlZCBjbGVhbnVwLiIsCgl9KQoJcmVmc0NsZWFuZWQgOj0gcHJvbWV0aGV1cy5OZXdDb3VudGVyKHByb21ldGhldXMuQ291bnRlck9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3JlZnNfY2xlYW5lZF90b3RhbCIsCgkJSGVscDogIlRvdGFsIG51bWJlciBvZiBzdGFsZSByZWZzIGNsZWFuZWQgdXAgb3ZlciB0aW1lLiIsCgl9KQoJdW5pcXVlUmVmc1RvdGFsIDo9IHByb21ldGhldXMuTmV3Q291bnRlcihwcm9tZXRoZXVzLkNvdW50ZXJPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV91bmlxdWVfcmVmc19jcmVhdGVkX3RvdGFsIiwKCQlIZWxwOiAiVG90YWwgbnVtYmVyIG9mIHVuaXF1ZSByZWZzIGNyZWF0ZWQuIiwKCX0pCgoJaWYgcmVnICE9IG5pbCB7CgkJcmVnLlJlZ2lzdGVyKGFjdGl2ZU1hcHBpbmdzKQoJCXJlZy5SZWdpc3Rlcih0cmFja2VkUmVmcykKCQlyZWcuUmVnaXN0ZXIocmVmc0NsZWFuZWQpCgkJcmVnLlJlZ2lzdGVyKHVuaXF1ZVJlZnNUb3RhbCkKCX0KCglyZXR1cm4gJlNlcmllc1JlZk1hcHBpbmdTdG9yZXsKCQl1bmlxdWVSZWZUb0NoaWxkUmVmczogbWFrZShtYXBbc3RvcmFnZS5TZXJpZXNSZWZddW5pcVJlZkNoaWxkcmVuKSwKCQluZXh0VW5pcXVlUmVmOiAgICAgICAgMSwKCQl1bmlxdWVSZWZUaW1lc3RhbXBzOiAgbWFrZShtYXBbc3RvcmFnZS5TZXJpZXNSZWZdaW50NjQpLAoJCWxhYmVsSGFzaFRvVW5pcXVlUmVmOiBtYWtlKG1hcFt1aW50NjRdc3RvcmFnZS5TZXJpZXNSZWYpLAoJCWNlbGxQb29sOiBzeW5jLlBvb2x7CgkJCU5ldzogZnVuYygpIGFueSB7CgkJCQlyZXR1cm4gJkNlbGx7UmVmczogbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAxMDApfQoJCQl9LAoJCX0sCgkJc3RvcENsZWFudXA6ICAgICBtYWtlKGNoYW4gc3RydWN0e30pLAoJCWNsZWFudXBTdG9wcGVkOiAgbWFrZShjaGFuIHN0cnVjdHt9KSwKCQlhY3RpdmVNYXBwaW5nczogIGFjdGl2ZU1hcHBpbmdzLAoJCXRyYWNrZWRSZWZzOiAgICAgdHJhY2tlZFJlZnMsCgkJcmVmc0NsZWFuZWQ6ICAgICByZWZzQ2xlYW5lZCwKCQl1bmlxdWVSZWZzVG90YWw6IHVuaXF1ZVJlZnNUb3RhbCwKCX0KfQoKdHlwZSBDZWxsIHN0cnVjdCB7CglSZWZzIFtdc3RvcmFnZS5TZXJpZXNSZWYKfQoKLy8gR2V0TWFwcGluZyByZXR1cm5zIGV4aXN0aW5nIGNoaWxkIHJlZiByZXN1bHRzIGZvciB0aGUgZ2l2ZW4gdW5pcXVlIHJlZiBpZiBvbmUgZXhpc3RzLgovLwovLyBJZiB0aGUgcGFzc2VkIHVuaXF1ZVJlZiBpcyB6ZXJvLCB0aGUgbWV0aG9kIHdpbGwgYXR0ZW1wdCB0byBmaW5kIGEgbWFwcGluZyB1c2luZyBwYXNzZWQgbGFiZWxzLgovLyBSZXR1cm5zIG5pbCBpZiBubyBtYXBwaW5nIGV4aXN0cy4KLy8KLy8gVGhlIHJldHVybmVkIHNsaWNlIG1heSBiZSBtb2RpZmllZCBieSB0aGUgY2FsbGVyLCBidXQgVXBkYXRlTWFwcGluZyBtdXN0IGJlIGNhbGxlZAovLyBhZnRlcndhcmRzIHRvIHBlcnNpc3QgY2hhbmdlcy4gTm90ZSB0aGF0IGNvbmN1cnJlbnQgYXBwZW5kZXJzIG1heSByYWNlIHRvIHVwZGF0ZSB0aGUKLy8gc2FtZSBtYXBwaW5nIHdpdGggZGlmZmVyZW50IHZhbHVlcywgd2hpY2ggaXMgc2FmZSBiZWNhdXNlIHN0YWxlIG1hcHBpbmdzIGFyZSBzZWxmLWNvcnJlY3RpbmcgLQovLyB1c2luZyBhIHN0YWxlIHJlZiB3aWxsIGNhdXNlIHRoZSBjaGlsZCBhcHBlbmRlciB0byByZXR1cm4gYSBuZXcgcmVmIG9uIHRoZSBuZXh0IGFwcGVuZC4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBHZXRNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSBbXXN0b3JhZ2UuU2VyaWVzUmVmIHsKCXMucmVmTWFwcGluZ011LlJMb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlJVbmxvY2soKQoKCWlmIHVuaXF1ZVJlZiA9PSAwIHsKCQkvLyBTb21lIGNvbnN1bWVycyBkb24ndCBtZW1vIHRoZSBnbG9iYWwgcmVmLiBUcnkgdG8gbG9va3VwIGEgcmVmIGJ5IGxhYmVsIGhhc2guCgkJbGFiZWxIYXNoIDo9IGxibHMuSGFzaCgpCgkJZ290UmVmLCBvayA6PSBzLmxhYmVsSGFzaFRvVW5pcXVlUmVmW2xhYmVsSGFzaF0KCQlpZiAhb2sgewoJCQlyZXR1cm4gbmlsCgkJfQoKCQl1bmlxdWVSZWYgPSBnb3RSZWYKCX0KCglpZiBtYXBwaW5nLCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl07IG9rIHsKCQlyZXR1cm4gKm1hcHBpbmcuY2hpbGRSZWZzCgl9CglyZXR1cm4gbmlsCn0KCi8vIENyZWF0ZU1hcHBpbmcgY3JlYXRlcyBhIG5ldyB1bmlxdWUgcmVmIG1hcHBpbmcgZm9yIHRoZSBnaXZlbiBjaGlsZCByZWYgcmVzdWx0cy4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBDcmVhdGVNYXBwaW5nKHJlZlJlc3VsdHMgW11zdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSBzdG9yYWdlLlNlcmllc1JlZiB7CgkvLyBTdGFydCBjbGVhbnVwIGdvcm91dGluZSBvbiBmaXJzdCBtYXBwaW5nCglzLnN0YXJ0UmVmQ2xlYW51cC5EbyhmdW5jKCkgewoJCXMuY2xlYW51cFN0YXJ0ZWQuU3RvcmUodHJ1ZSkKCQlnbyBzLmNsZWFudXBTdGFsZVJlZnMoKQoJfSkKCgkvLyBTdG9yZSBhIGNvcHkgb2YgdGhlIGNoaWxkIHJlZiByZXN1bHRzIGRpcmVjdGx5CgljaGlsZFJlZlNsaWNlIDo9IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgbGVuKHJlZlJlc3VsdHMpKQoJY29weShjaGlsZFJlZlNsaWNlLCByZWZSZXN1bHRzKQoKCS8vIEhhc2ggbGFiZWxzIHRvIGZvciB0aGUgZmFsbGJhY2sgbG9va3VwIHRhYmxlCglsYWJlbEhhc2ggOj0gbGJscy5IYXNoKCkKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJLy8gQ3JlYXRlIGEgbmV3IHVuaXF1ZSByZWYKCXVuaXF1ZVJlZiA6PSBzLm5leHRVbmlxdWVSZWYKCXMubmV4dFVuaXF1ZVJlZisrCgoJcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltsYWJlbEhhc2hdID0gdW5pcXVlUmVmCglzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0gPSB1bmlxUmVmQ2hpbGRyZW57CgkJY2hpbGRSZWZzOiAmY2hpbGRSZWZTbGljZSwKCQlsYWJlbEhhc2g6IGxhYmVsSGFzaCwKCX0KCglzLmFjdGl2ZU1hcHBpbmdzLkluYygpCglzLnVuaXF1ZVJlZnNUb3RhbC5JbmMoKQoKCXJldHVybiB1bmlxdWVSZWYKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHsKCWlmIHVuaXF1ZVJlZiA9PSAwIHsKCQlyZXR1cm4KCX0KCgljaGlsZFJlZlNsaWNlIDo9IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgbGVuKHJlZlJlc3VsdHMpKQoJY29weShjaGlsZFJlZlNsaWNlLCByZWZSZXN1bHRzKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgkvLyBFbnN1cmUgdGhhdCBsYWJlbCBoYXNoIGluZGV4IGlzIHVwIHRvIGRhdGUgdG8gaGFuZGxlIHBvc3NpYmxlIGhhc2ggY29sbGlzaW9ucy4KCS8vIFRPRE86IGlzIHRoaXMgbmVjZXNzYXJ5PwoJbmV3SGFzaCA6PSBsYmxzLkhhc2goKQoJcHJldiwgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdCglpZiBvayAmJiBwcmV2LmxhYmVsSGFzaCAhPSBuZXdIYXNoIHsKCQlkZWxldGUocy5sYWJlbEhhc2hUb1VuaXF1ZVJlZiwgcHJldi5sYWJlbEhhc2gpCgkJcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltuZXdIYXNoXSA9IHVuaXF1ZVJlZgoJfQoKCXMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXSA9IHVuaXFSZWZDaGlsZHJlbnsKCQljaGlsZFJlZnM6ICZjaGlsZFJlZlNsaWNlLAoJCWxhYmVsSGFzaDogbGJscy5IYXNoKCksCgl9Cn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgVHJhY2tBcHBlbmRlZFNlcmllcyh0cyBpbnQ2NCwgY2VsbCAqQ2VsbCkgewoJcy50aW1lc3RhbXBUcmFja2luZ011LkxvY2soKQoJZGVmZXIgcy50aW1lc3RhbXBUcmFja2luZ011LlVubG9jaygpCgoJZm9yIF8sIHIgOj0gcmFuZ2UgY2VsbC5SZWZzIHsKCQlzLnVuaXF1ZVJlZlRpbWVzdGFtcHNbcl0gPSB0cwoJfQoKCXMudHJhY2tlZFJlZnMuU2V0KGZsb2F0NjQobGVuKHMudW5pcXVlUmVmVGltZXN0YW1wcykpKQoKCWNlbGwuUmVmcyA9IGNlbGwuUmVmc1s6MF0KCXMuY2VsbFBvb2wuUHV0KGNlbGwpCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwgewoJcmV0dXJuIHMuY2VsbFBvb2wuR2V0KCkuKCpDZWxsKQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIGNsZWFudXBTdGFsZVJlZnMoKSB7CglkZWZlciBjbG9zZShzLmNsZWFudXBTdG9wcGVkKQoKCXRpY2tlciA6PSB0aW1lLk5ld1RpY2tlcigxNSAqIHRpbWUuTWludXRlKQoJZGVmZXIgdGlja2VyLlN0b3AoKQoKCWZvciB7CgkJc2VsZWN0IHsKCQljYXNlIDwtdGlja2VyLkM6CgkJCWN1dG9mZlRpbWUgOj0gdGltZS5Ob3coKS5BZGQoLTE1ICogdGltZS5NaW51dGUpLlVuaXgoKQoKCQkJLy8gSG9sZCBib3RoIGxvY2tzIHRvIHByZXZlbnQgcmFjZSBjb25kaXRpb24gd2hlcmUgYSByZWYgY291bGQgYmUKCQkJLy8gYXBwZW5kZWQgYWZ0ZXIgd2UgZGVsZXRlIGl0IGZyb20gdW5pcXVlUmVmQ2VsbCBidXQgYmVmb3JlCgkJCS8vIHdlIGRlbGV0ZSBpdCBmcm9tIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzCgkJCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCQkJcy5yZWZNYXBwaW5nTXUuTG9jaygpCgoJCQlzdGFsZVJlZkNvdW50IDo9IDAKCQkJZm9yIHJlZiwgdHMgOj0gcmFuZ2Ugcy51bmlxdWVSZWZUaW1lc3RhbXBzIHsKCQkJCWlmIHRzIDwgY3V0b2ZmVGltZSB7CgkJCQkJc3RhbGVSZWZDb3VudCsrCgoJCQkJCXYsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbcmVmXQoJCQkJCWlmIG9rIHsKCQkJCQkJZGVsZXRlKHMubGFiZWxIYXNoVG9VbmlxdWVSZWYsIHYubGFiZWxIYXNoKQoJCQkJCX0KCgkJCQkJZGVsZXRlKHMudW5pcXVlUmVmVGltZXN0YW1wcywgcmVmKQoJCQkJCWRlbGV0ZShzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzLCByZWYpCgkJCQl9CgkJCX0KCgkJCS8vIFVwZGF0ZSBtZXRyaWNzCgkJCWlmIHN0YWxlUmVmQ291bnQgPiAwIHsKCQkJCXMucmVmc0NsZWFuZWQuQWRkKGZsb2F0NjQoc3RhbGVSZWZDb3VudCkpCgkJCQlzLmFjdGl2ZU1hcHBpbmdzLlN1YihmbG9hdDY0KHN0YWxlUmVmQ291bnQpKQoJCQkJcy50cmFja2VkUmVmcy5TZXQoZmxvYXQ2NChsZW4ocy51bmlxdWVSZWZUaW1lc3RhbXBzKSkpCgkJCX0KCgkJCXMucmVmTWFwcGluZ011LlVubG9jaygpCgkJCXMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCQljYXNlIDwtcy5zdG9wQ2xlYW51cDoKCQkJcmV0dXJuCgkJfQoJfQp9CgovLyBDbGVhciB3aWxsIGNsZWFyIGFsbCBpbnRlcm5hbCBtYXBwaW5ncyBhbmQgc3RvcCB0aGUgY2xlYW5lciBnb3JvdXRpbmUgaWYgaXQgaXMgcnVubmluZy4KLy8gSXQgaXMgc2FmZSB0byByZS11c2UgdGhlIHNhbWUgaW5zdGFuY2UgYWZ0ZXIgY2FsbGluZyBDbGVhci4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBDbGVhcigpIHsKCS8vIFN0b3AgdGhlIGNsZWFudXAgZ29yb3V0aW5lIGFuZCB3YWl0IGZvciBpdCB0byBiZSBzdG9wcGVkIHNvIHdlIGNhbgoJLy8gYXZvaWQgYSBwb3NzaWJsZSBkZWFkbG9jayB3aXRoIGNsZWFudXAgdGhhdCBhbHNvIGhvbGRzIGJvdGggbG9ja3MKCWlmIHMuY2xlYW51cFN0YXJ0ZWQuTG9hZCgpIHsKCQlzZWxlY3QgewoJCWNhc2UgPC1zLnN0b3BDbGVhbnVwOgoJCQkvLyBBbHJlYWR5IGNsb3NlZAoJCWRlZmF1bHQ6CgkJCWNsb3NlKHMuc3RvcENsZWFudXApCgkJCTwtcy5jbGVhbnVwU3RvcHBlZAoJCX0KCX0KCgkvLyBXZSBuZWVkIHRvIGhvbGQgYm90aCBsb2NrcyB0byBkbyB0aGlzIHNhZmVseSBhbmQgd2UgZG8gaXQgaW4gdGhlIHNhbWUgb3JkZXIgYXMKCS8vIGNsZWFudXBTdGFsZVJlZnMuIFdlIHN0b3BwZWQgYW5kIHdhaXRlZCBmb3IgdGhlIGJhY2tncm91bmQgd29ya2VyIHRoYXQgY2FsbHMgaXQKCS8vIHRvIGZpbmlzaCBidXQgc29tZSBleHRyYSBzYWZldHkgd29uJ3QgaHVydC4KCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCWRlZmVyIHMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgljbGVhcihzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzKQoJY2xlYXIocy51bmlxdWVSZWZUaW1lc3RhbXBzKQoKCS8vIHJlc2V0IHRoZSBwb29sCglzLmNlbGxQb29sID0gc3luYy5Qb29sewoJCU5ldzogZnVuYygpIGFueSB7CgkJCXJldHVybiAmQ2VsbHtSZWZzOiBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIDAsIDEwMCl9CgkJfSwKCX0KCgkvLyBOT1RFOiBXZSBkbyBOT1QgcmVzZXQgbmV4dFVuaXF1ZVJlZiBoZXJlLiBSZXNldHRpbmcgaXQgd291bGQgY2F1c2UgcmVmIGNvbGxpc2lvbnMKCS8vIHdpdGggY29tcG9uZW50cyBsaWtlIHByb21ldGhldXMuc2NyYXBlIHdoaWNoIHdpbGwga2VlcCByZS1zZW5kaW5nIHRoZSBzYW1lIGNhY2hlZCByZWZzLgoJLy8gV2UgY29udGludWUgaW5jcmVtZW50aW5nIHRvIGVuc3VyZSBhbGwgcmVmcyByZW1haW4gdW5pcXVlIGFjcm9zcyB0aGUgbGlmZXRpbWUgb2YgdGhlIHByb2Nlc3MuCgoJLy8gUmVzZXQgbWV0cmljcwoJcy5hY3RpdmVNYXBwaW5ncy5TZXQoMCkKCXMudHJhY2tlZFJlZnMuU2V0KDApCgoJLy8gUmVzZXQgY2hhbm5lbHMgYW5kIGZsYWdzCglzLnN0b3BDbGVhbnVwID0gbWFrZShjaGFuIHN0cnVjdHt9KQoJcy5jbGVhbnVwU3RvcHBlZCA9IG1ha2UoY2hhbiBzdHJ1Y3R7fSkKCXMuc3RhcnRSZWZDbGVhbnVwID0gc3luYy5PbmNle30KCXMuY2xlYW51cFN0YXJ0ZWQuU3RvcmUoZmFsc2UpCn0K"
    }
  }
}
```

Response (packages.DriverResponse):
Error: `err: context canceled: stderr: `

#### drv #9

Trace meta: spanId=17, ts=1770837171634, ts_iso=2026-02-11T19:12:51.634000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy/internal/validator"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkiZm10IgoJInN5bmMiCgkic3luYy9hdG9taWMiCgkidGltZSIKCgkiZ2l0aHViLmNvbS9oYXNoaWNvcnAvZ28tbXVsdGllcnJvciIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvY2xpZW50X2dvbGFuZy9wcm9tZXRoZXVzIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2V4ZW1wbGFyIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2hpc3RvZ3JhbSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9sYWJlbHMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvbWV0YWRhdGEiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvc3RvcmFnZSIKKQoKdHlwZSBNYXBwaW5nU3RvcmUgaW50ZXJmYWNlIHsKCUdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYKCUNyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmCglVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpCglUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKQoJR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwKfQoKdHlwZSBzZXJpZXNSZWZNYXBwaW5nIHN0cnVjdCB7CglzdGFydCAgICB0aW1lLlRpbWUKCWNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlcgoJc3RvcmUgICAgTWFwcGluZ1N0b3JlCgoJdW5pcXVlUmVmQ2VsbCAqQ2VsbAoKCS8vIGNoaWxkUmVmcyBpcyByZXVzZWQgZm9yIGVhY2ggYXBwZW5kIGNhbGwgdG8gYXZvaWQgYWxsb2NhdGlvbnMuIFRoaXMgaXMgc2FmZSBiZWNhdXNlIHN0b3JhZ2UuQXBwZW5kZXIgc2hvdWxkIG5ldmVyCgkvLyBoYXZlIGNvbmN1cnJlbnQgY2FsbHMgdG8gQXBwZW5kIG1ldGhvZHMuCgljaGlsZFJlZnMgICAgICAgIFtdc3RvcmFnZS5TZXJpZXNSZWYKCXdyaXRlTGF0ZW5jeSAgICAgcHJvbWV0aGV1cy5IaXN0b2dyYW0KCXNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZyhjaGlsZHJlbiBbXXN0b3JhZ2UuQXBwZW5kZXIsIHN0b3JlIE1hcHBpbmdTdG9yZSwgd3JpdGVMYXRlbmN5IHByb21ldGhldXMuSGlzdG9ncmFtLCBzYW1wbGVzRm9yd2FyZGVkIHByb21ldGhldXMuQ291bnRlcikgc3RvcmFnZS5BcHBlbmRlciB7Cgl1bmlxdWVSZWZDZWxsIDo9IHN0b3JlLkdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpCgoJcmV0dXJuICZzZXJpZXNSZWZNYXBwaW5newoJCWNoaWxkcmVuOiAgICAgICAgIGNoaWxkcmVuLAoJCXN0b3JlOiAgICAgICAgICAgIHN0b3JlLAoJCXdyaXRlTGF0ZW5jeTogICAgIHdyaXRlTGF0ZW5jeSwKCQlzYW1wbGVzRm9yd2FyZGVkOiBzYW1wbGVzRm9yd2FyZGVkLAoKCQl1bmlxdWVSZWZDZWxsOiB1bmlxdWVSZWZDZWxsLAoJCWNoaWxkUmVmczogICAgIG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgbGVuKGNoaWxkcmVuKSksCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFNldE9wdGlvbnMob3B0cyAqc3RvcmFnZS5BcHBlbmRPcHRpb25zKSB7CglmbXQuUHJpbnRsbigiYXNkYXMpCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQljLlNldE9wdGlvbnMob3B0cykKCX0KfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQ29tbWl0KCkgZXJyb3IgewoJcy5zdG9yZS5UcmFja0FwcGVuZGVkU2VyaWVzKHRpbWUuTm93KCkuVW5peCgpLCBzLnVuaXF1ZVJlZkNlbGwpCgoJdmFyIG11bHRpRXJyIGVycm9yCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQllcnIgOj0gYy5Db21taXQoKQoJCWlmIGVyciAhPSBuaWwgewoJCQltdWx0aUVyciA9IG11bHRpZXJyb3IuQXBwZW5kKG11bHRpRXJyLCBlcnIpCgkJfQoJfQoJcmV0dXJuIG11bHRpRXJyCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFJvbGxiYWNrKCkgZXJyb3IgewoJLy8gV2Ugc3RpbGwgdHJhY2sgcm9sbGVkIGJhY2sgc2VyaWVzIHNvIHdlIGNhbiBwcm9wZXJseQoJLy8gY2xlYW4gdXAgYW55IHNlcmllcyB0aGF0IHdhcyBhcHBlbmRlZAoJcy5zdG9yZS5UcmFja0FwcGVuZGVkU2VyaWVzKHRpbWUuTm93KCkuVW5peCgpLCBzLnVuaXF1ZVJlZkNlbGwpCgoJdmFyIG11bHRpRXJyIGVycm9yCglmb3IgXywgYyA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQllcnIgOj0gYy5Sb2xsYmFjaygpCgkJaWYgZXJyICE9IG5pbCB7CgkJCW11bHRpRXJyID0gbXVsdGllcnJvci5BcHBlbmQobXVsdGlFcnIsIGVycikKCQl9Cgl9CglyZXR1cm4gbXVsdGlFcnIKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgcmVjb3JkTGF0ZW5jeSgpIHsKCWlmIHMuc3RhcnQuSXNaZXJvKCkgewoJCXJldHVybgoJfQoJZHVyYXRpb24gOj0gdGltZS5TaW5jZShzLnN0YXJ0KQoJcy53cml0ZUxhdGVuY3kuT2JzZXJ2ZShkdXJhdGlvbi5TZWNvbmRzKCkpCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIHJlc2V0RmllbGRzKCkgewoJLy8gUmVzZXQgY2hpbGRSZWZzIHNsaWNlIGxlbmd0aCB0byAwIGZvciByZXVzZQoJcy5jaGlsZFJlZnMgPSBzLmNoaWxkUmVmc1s6MF0KfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0IGludDY0LCB2IGZsb2F0NjQpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCW5ld1JlZiwgZXJyIDo9IGFwcGVuZGVyLkFwcGVuZChyZWYsIGwsIHQsIHYpCgkJaWYgZXJyID09IG5pbCB7CgkJCXMuc2FtcGxlc0ZvcndhcmRlZC5JbmMoKQoJCX0KCQlyZXR1cm4gbmV3UmVmLCBlcnIKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEV4ZW1wbGFyKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCBlIGV4ZW1wbGFyLkV4ZW1wbGFyKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kRXhlbXBsYXIocmVmLCBsLCBlKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kSGlzdG9ncmFtKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0IGludDY0LCBoICpoaXN0b2dyYW0uSGlzdG9ncmFtLCBmaCAqaGlzdG9ncmFtLkZsb2F0SGlzdG9ncmFtKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kSGlzdG9ncmFtKHJlZiwgbCwgdCwgaCwgZmgpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRIaXN0b2dyYW1DVFplcm9TYW1wbGUocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQsIGN0IGludDY0LCBoICpoaXN0b2dyYW0uSGlzdG9ncmFtLCBmaCAqaGlzdG9ncmFtLkZsb2F0SGlzdG9ncmFtKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kSGlzdG9ncmFtQ1RaZXJvU2FtcGxlKHJlZiwgbCwgdCwgY3QsIGgsIGZoKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgVXBkYXRlTWV0YWRhdGEocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIG0gbWV0YWRhdGEuTWV0YWRhdGEpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5VcGRhdGVNZXRhZGF0YShyZWYsIGwsIG0pCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRDVFplcm9TYW1wbGUocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQsIGN0IGludDY0KSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuQXBwZW5kQ1RaZXJvU2FtcGxlKHJlZiwgbCwgdCwgY3QpCgl9KQp9Cgp0eXBlIGFwcGVuZGVyRnVuYyBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikKCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIGFwcGVuZFRvQ2hpbGRyZW4ocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMsIGFmIGFwcGVuZGVyRnVuYykgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJZGVmZXIgcy5yZXNldEZpZWxkcygpCgoJaWYgcy5zdGFydC5Jc1plcm8oKSB7CgkJcy5zdGFydCA9IHRpbWUuTm93KCkKCX0KCgkvLyBDaGVjayBpZiB0aGUgaW5jb21pbmcgcmVmIGhhcyByZWYgbWFwcGluZ3MKCWV4aXN0aW5nQ2hpbGRSZWZzIDo9IHMuc3RvcmUuR2V0TWFwcGluZyhyZWYsIGxibHMpCgoJdmFyIGFwcGVuZEVyciBlcnJvcgoKCS8vIFNhbml0eSBjaGVjazogaWYgd2UgaGF2ZSBleGlzdGluZyBjaGlsZCByZWZzLCB0aGV5IG11c3QgbWF0Y2ggdGhlIG51bWJlciBvZiBjaGlsZHJlbgoJaWYgZXhpc3RpbmdDaGlsZFJlZnMgIT0gbmlsICYmIGxlbihleGlzdGluZ0NoaWxkUmVmcykgPT0gbGVuKHMuY2hpbGRyZW4pIHsKCQlzLnVuaXF1ZVJlZkNlbGwuUmVmcyA9IGFwcGVuZChzLnVuaXF1ZVJlZkNlbGwuUmVmcywgcmVmKQoKCQlyZWZVcGRhdGVSZXF1aXJlZCA6PSBmYWxzZQoJCWZvciBjaGlsZEluZGV4LCBjaGlsZFJlZiA6PSByYW5nZSBleGlzdGluZ0NoaWxkUmVmcyB7CgkJCW5ld0NoaWxkUmVmLCBlcnIgOj0gYWYocy5jaGlsZHJlbltjaGlsZEluZGV4XSwgY2hpbGRSZWYpCgkJCWlmIGVyciAhPSBuaWwgewoJCQkJYXBwZW5kRXJyID0gbXVsdGllcnJvci5BcHBlbmQoYXBwZW5kRXJyLCBlcnIpCgkJCX0KCgkJCWlmIG5ld0NoaWxkUmVmICE9IGNoaWxkUmVmIHsKCQkJCS8vIENoaWxkIHJlZiBjaGFuZ2VkLCBuZWVkIHRvIHVwZGF0ZSBtYXBwaW5nCgkJCQlleGlzdGluZ0NoaWxkUmVmc1tjaGlsZEluZGV4XSA9IG5ld0NoaWxkUmVmCgkJCQlyZWZVcGRhdGVSZXF1aXJlZCA9IHRydWUKCQkJfQoJCX0KCgkJaWYgYXBwZW5kRXJyICE9IG5pbCB7CgkJCXJldHVybiAwLCBhcHBlbmRFcnIKCQl9CgoJCWlmIHJlZlVwZGF0ZVJlcXVpcmVkIHsKCQkJcy5zdG9yZS5VcGRhdGVNYXBwaW5nKHJlZiwgZXhpc3RpbmdDaGlsZFJlZnMsIGxibHMpCgkJfQoKCQlyZXR1cm4gcmVmLCBuaWwKCX0KCgkvLyBObyBleGlzdGluZyBtYXBwaW5nLCBwcm9jZWVkIHdpdGggbm9ybWFsIGFwcGVuZCB0byBhbGwgY2hpbGRyZW4KCXZhciBmaXJzdE5vblplcm9SZWYgc3RvcmFnZS5TZXJpZXNSZWYKCXZhciBub25aZXJvQ291bnQgaW50CgoJLy8gTm90ZTogdGhlcmUncyBhbm90aGVyIG9wdGltaXphdGlvbiB3aGVyZSB3ZSBjb3VsZCB1c2UgdGhlIHJldHVybmVkIHJlZiBpZiBhbGwgdGhlIG5vbiB6ZXJvIHJlZnMKCS8vICBhcmUgdGhlIHNhbWUgdmFsdWUuIFRoaXMgaXNuJ3Qgc2FmZSBhcyB3ZSB3aWxsIG1peCBkb3duc3RyZWFtIHJlZnMgd2l0aCB1bmlxdWUgcmVmcyB3aGljaCBjb3VsZAoJLy8gIGNvbGxpZGUuIFdlIGNvdWxkIHN0YXJ0IGF0IG1heCB1bml0NjQgZm9yIHVuaXF1ZSByZWZzIGFuZCBnbyBiYWNrd2FyZHMgbGVzc2VuaW5nIHRoZSBjaGFuY2Ugb2YKCS8vIAljb2xsaXNpb25zIGJ1dCBpdCdzIHJhdGhlciBkYW5nZXJvdXMgZm9yIGFuIHVubGlrZWx5IGVkZ2UgY2FzZS4gSWYgdHdvIGNvbXBvbmVudHMgYXJlIHJldHVybmluZwoJLy8gCXRoZSBzYW1lIHJlZiBpdCdzIHR3byByZW1vdGVfd3JpdGUgY29tcG9uZW50cyB3aGljaCBzaG91bGQgcHJvYmFibHkgYmUgbWVyZ2VkIGluIHRvIG9uZS4KCWZvciBfLCBjaGlsZCA6PSByYW5nZSBzLmNoaWxkcmVuIHsKCQljaGlsZFJlZiwgZXJyIDo9IGFmKGNoaWxkLCByZWYpCgkJaWYgZXJyICE9IG5pbCB7CgkJCWFwcGVuZEVyciA9IG11bHRpZXJyb3IuQXBwZW5kKGFwcGVuZEVyciwgZXJyKQoKCQkJLy8gVE9ETyBzaG91bGQgSSBtYWtlIHRoZSBjaGlsZFJlZiB6ZXJvIGhlcmU/CgkJfQoKCQlzLmNoaWxkUmVmcyA9IGFwcGVuZChzLmNoaWxkUmVmcywgY2hpbGRSZWYpCgkJaWYgY2hpbGRSZWYgIT0gMCB7CgkJCWlmIGZpcnN0Tm9uWmVyb1JlZiA9PSAwIHsKCQkJCWZpcnN0Tm9uWmVyb1JlZiA9IGNoaWxkUmVmCgkJCX0KCQkJbm9uWmVyb0NvdW50KysKCQl9Cgl9CgoJaWYgYXBwZW5kRXJyICE9IG5pbCB7CgkJcmV0dXJuIDAsIGFwcGVuZEVycgoJfQoKCWlmIG5vblplcm9Db3VudCA9PSAwIHsKCQkvLyBBbGwgY2hpbGRyZW4gcmV0dXJuZWQgcmVmIDAsIHNvIHJldHVybiB0aGUgaW5wdXQgcmVmCgkJcmV0dXJuIHJlZiwgbmlsCgl9CgoJLy8gT25seSBvbmUgY2hpbGQgcmV0dXJuZWQgYSBub24temVybyByZWYsIHVzZSB0aGF0CglpZiBub25aZXJvQ291bnQgPT0gMSB7CgkJcmV0dXJuIGZpcnN0Tm9uWmVyb1JlZiwgbmlsCgl9CgoJLy8gV2UgZ290IGRpZmZlcmVudCByZWZzIGJhY2sgYW5kIG5lZWQgdG8gY3JlYXRlIGEgbmV3IG1hcHBpbmcKCXVuaXF1ZVJlZiA6PSBzLnN0b3JlLkNyZWF0ZU1hcHBpbmcocy5jaGlsZFJlZnMsIGxibHMpCglzLnVuaXF1ZVJlZkNlbGwuUmVmcyA9IGFwcGVuZChzLnVuaXF1ZVJlZkNlbGwuUmVmcywgdW5pcXVlUmVmKQoJcmV0dXJuIHVuaXF1ZVJlZiwgbmlsCn0KCnR5cGUgdW5pcVJlZkNoaWxkcmVuIHN0cnVjdCB7CgljaGlsZFJlZnMgKltdc3RvcmFnZS5TZXJpZXNSZWYKCWxhYmVsSGFzaCB1aW50NjQKfQoKdHlwZSBTZXJpZXNSZWZNYXBwaW5nU3RvcmUgc3RydWN0IHsKCS8vIHJlZk1hcHBpbmdNdSBwcm90ZWN0cyB1bmlxdWVSZWZUb0NoaWxkUmVmcywgbGFiZWxIYXNoVG9VbmlxdWVSZWYgYW5kIG5leHRVbmlxdWVSZWYKCXJlZk1hcHBpbmdNdSBzeW5jLlJXTXV0ZXgKCS8vIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzIG1hcHMgdGhlIHVuaXF1ZSByZWYgdG8gdGhlIGV4cGVjdGVkIGNoaWxkIHJlZiBpbiBvcmRlcgoJdW5pcXVlUmVmVG9DaGlsZFJlZnMgbWFwW3N0b3JhZ2UuU2VyaWVzUmVmXXVuaXFSZWZDaGlsZHJlbgoJLy8gbGFiZWxIYXNoVG9VbmlxdWVSZWYgbWFwcyB0aGUgbGFiZWwgaGFzaCB0byB1bmlxdWUgcmVmLgoJbGFiZWxIYXNoVG9VbmlxdWVSZWYgbWFwW3VpbnQ2NF1zdG9yYWdlLlNlcmllc1JlZgoKCS8vIG5leHRVbmlxdWVSZWYgaXMgdGhlIG5leHQgcmVmIElEIHdlIHdpbGwgaGFuZCBvdXQKCW5leHRVbmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYKCgkvLyB0aW1lc3RhbXBUcmFja2luZ011IHByb3RlY3RzIHVuaXF1ZVJlZlRpbWVzdGFtcHMgYW5kIGNlbGxQb29sCgl0aW1lc3RhbXBUcmFja2luZ011IHN5bmMuTXV0ZXgKCS8vIHVuaXF1ZVJlZlRpbWVzdGFtcHMgbWFwcyB1bmlxdWUgcmVmcyB0byB0aGVpciBsYXN0IGFwcGVuZCB0aW1lc3RhbXAKCXVuaXF1ZVJlZlRpbWVzdGFtcHMgbWFwW3N0b3JhZ2UuU2VyaWVzUmVmXWludDY0CgkvLyBjZWxsUG9vbCBpcyB1c2VkIHRvIHBvb2wgc2xpY2VzIG9mIFNlcmllc1JlZnMgdXNlZCBmb3IgdHJhY2tpbmcgdW5pcXVlIHJlZnMgaW4gVHJhY2tBcHBlbmRlZFNlcmllcy4KCWNlbGxQb29sIHN5bmMuUG9vbAoKCS8vIENsZWFudXAgZ29yb3V0aW5lIGNvb3JkaW5hdGlvbiAobm8gbG9jayByZXF1aXJlZCkKCXN0YXJ0UmVmQ2xlYW51cCBzeW5jLk9uY2UKCWNsZWFudXBTdGFydGVkICBhdG9taWMuQm9vbAoJc3RvcENsZWFudXAgICAgIGNoYW4gc3RydWN0e30KCWNsZWFudXBTdG9wcGVkICBjaGFuIHN0cnVjdHt9CgoJLy8gTWV0cmljcyAoc2FmZSBmb3IgY29uY3VycmVudCBhY2Nlc3MsIG5vIGxvY2sgcmVxdWlyZWQpCglhY3RpdmVNYXBwaW5ncyAgcHJvbWV0aGV1cy5HYXVnZQoJdHJhY2tlZFJlZnMgICAgIHByb21ldGhldXMuR2F1Z2UKCXJlZnNDbGVhbmVkICAgICBwcm9tZXRoZXVzLkNvdW50ZXIKCXVuaXF1ZVJlZnNUb3RhbCBwcm9tZXRoZXVzLkNvdW50ZXIKfQoKZnVuYyBOZXdTZXJpZXNSZWZNYXBwaW5nU3RvcmUocmVnIHByb21ldGhldXMuUmVnaXN0ZXJlcikgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSB7CglhY3RpdmVNYXBwaW5ncyA6PSBwcm9tZXRoZXVzLk5ld0dhdWdlKHByb21ldGhldXMuR2F1Z2VPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV9tYXBwaW5nc190b3RhbCIsCgkJSGVscDogIk51bWJlciBvZiBhY3RpdmUgdW5pcXVlIHJlZiBtYXBwaW5ncyBpbiB0aGUgc3RvcmUuIiwKCX0pCgl0cmFja2VkUmVmcyA6PSBwcm9tZXRoZXVzLk5ld0dhdWdlKHByb21ldGhldXMuR2F1Z2VPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV90cmFja2VkX3JlZnNfdG90YWwiLAoJCUhlbHA6ICJOdW1iZXIgb2YgcmVmcyBiZWluZyB0cmFja2VkIGZvciB0aW1lc3RhbXAtYmFzZWQgY2xlYW51cC4iLAoJfSkKCXJlZnNDbGVhbmVkIDo9IHByb21ldGhldXMuTmV3Q291bnRlcihwcm9tZXRoZXVzLkNvdW50ZXJPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV9yZWZzX2NsZWFuZWRfdG90YWwiLAoJCUhlbHA6ICJUb3RhbCBudW1iZXIgb2Ygc3RhbGUgcmVmcyBjbGVhbmVkIHVwIG92ZXIgdGltZS4iLAoJfSkKCXVuaXF1ZVJlZnNUb3RhbCA6PSBwcm9tZXRoZXVzLk5ld0NvdW50ZXIocHJvbWV0aGV1cy5Db3VudGVyT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfdW5pcXVlX3JlZnNfY3JlYXRlZF90b3RhbCIsCgkJSGVscDogIlRvdGFsIG51bWJlciBvZiB1bmlxdWUgcmVmcyBjcmVhdGVkLiIsCgl9KQoKCWlmIHJlZyAhPSBuaWwgewoJCXJlZy5SZWdpc3RlcihhY3RpdmVNYXBwaW5ncykKCQlyZWcuUmVnaXN0ZXIodHJhY2tlZFJlZnMpCgkJcmVnLlJlZ2lzdGVyKHJlZnNDbGVhbmVkKQoJCXJlZy5SZWdpc3Rlcih1bmlxdWVSZWZzVG90YWwpCgl9CgoJcmV0dXJuICZTZXJpZXNSZWZNYXBwaW5nU3RvcmV7CgkJdW5pcXVlUmVmVG9DaGlsZFJlZnM6IG1ha2UobWFwW3N0b3JhZ2UuU2VyaWVzUmVmXXVuaXFSZWZDaGlsZHJlbiksCgkJbmV4dFVuaXF1ZVJlZjogICAgICAgIDEsCgkJdW5pcXVlUmVmVGltZXN0YW1wczogIG1ha2UobWFwW3N0b3JhZ2UuU2VyaWVzUmVmXWludDY0KSwKCQlsYWJlbEhhc2hUb1VuaXF1ZVJlZjogbWFrZShtYXBbdWludDY0XXN0b3JhZ2UuU2VyaWVzUmVmKSwKCQljZWxsUG9vbDogc3luYy5Qb29sewoJCQlOZXc6IGZ1bmMoKSBhbnkgewoJCQkJcmV0dXJuICZDZWxse1JlZnM6IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMTAwKX0KCQkJfSwKCQl9LAoJCXN0b3BDbGVhbnVwOiAgICAgbWFrZShjaGFuIHN0cnVjdHt9KSwKCQljbGVhbnVwU3RvcHBlZDogIG1ha2UoY2hhbiBzdHJ1Y3R7fSksCgkJYWN0aXZlTWFwcGluZ3M6ICBhY3RpdmVNYXBwaW5ncywKCQl0cmFja2VkUmVmczogICAgIHRyYWNrZWRSZWZzLAoJCXJlZnNDbGVhbmVkOiAgICAgcmVmc0NsZWFuZWQsCgkJdW5pcXVlUmVmc1RvdGFsOiB1bmlxdWVSZWZzVG90YWwsCgl9Cn0KCnR5cGUgQ2VsbCBzdHJ1Y3QgewoJUmVmcyBbXXN0b3JhZ2UuU2VyaWVzUmVmCn0KCi8vIEdldE1hcHBpbmcgcmV0dXJucyBleGlzdGluZyBjaGlsZCByZWYgcmVzdWx0cyBmb3IgdGhlIGdpdmVuIHVuaXF1ZSByZWYgaWYgb25lIGV4aXN0cy4KLy8KLy8gSWYgdGhlIHBhc3NlZCB1bmlxdWVSZWYgaXMgemVybywgdGhlIG1ldGhvZCB3aWxsIGF0dGVtcHQgdG8gZmluZCBhIG1hcHBpbmcgdXNpbmcgcGFzc2VkIGxhYmVscy4KLy8gUmV0dXJucyBuaWwgaWYgbm8gbWFwcGluZyBleGlzdHMuCi8vCi8vIFRoZSByZXR1cm5lZCBzbGljZSBtYXkgYmUgbW9kaWZpZWQgYnkgdGhlIGNhbGxlciwgYnV0IFVwZGF0ZU1hcHBpbmcgbXVzdCBiZSBjYWxsZWQKLy8gYWZ0ZXJ3YXJkcyB0byBwZXJzaXN0IGNoYW5nZXMuIE5vdGUgdGhhdCBjb25jdXJyZW50IGFwcGVuZGVycyBtYXkgcmFjZSB0byB1cGRhdGUgdGhlCi8vIHNhbWUgbWFwcGluZyB3aXRoIGRpZmZlcmVudCB2YWx1ZXMsIHdoaWNoIGlzIHNhZmUgYmVjYXVzZSBzdGFsZSBtYXBwaW5ncyBhcmUgc2VsZi1jb3JyZWN0aW5nIC0KLy8gdXNpbmcgYSBzdGFsZSByZWYgd2lsbCBjYXVzZSB0aGUgY2hpbGQgYXBwZW5kZXIgdG8gcmV0dXJuIGEgbmV3IHJlZiBvbiB0aGUgbmV4dCBhcHBlbmQuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgR2V0TWFwcGluZyh1bmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgW11zdG9yYWdlLlNlcmllc1JlZiB7CglzLnJlZk1hcHBpbmdNdS5STG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5SVW5sb2NrKCkKCglpZiB1bmlxdWVSZWYgPT0gMCB7CgkJLy8gU29tZSBjb25zdW1lcnMgZG9uJ3QgbWVtbyB0aGUgZ2xvYmFsIHJlZi4gVHJ5IHRvIGxvb2t1cCBhIHJlZiBieSBsYWJlbCBoYXNoLgoJCWxhYmVsSGFzaCA6PSBsYmxzLkhhc2goKQoJCWdvdFJlZiwgb2sgOj0gcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltsYWJlbEhhc2hdCgkJaWYgIW9rIHsKCQkJcmV0dXJuIG5pbAoJCX0KCgkJdW5pcXVlUmVmID0gZ290UmVmCgl9CgoJaWYgbWFwcGluZywgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdOyBvayB7CgkJcmV0dXJuICptYXBwaW5nLmNoaWxkUmVmcwoJfQoJcmV0dXJuIG5pbAp9CgovLyBDcmVhdGVNYXBwaW5nIGNyZWF0ZXMgYSBuZXcgdW5pcXVlIHJlZiBtYXBwaW5nIGZvciB0aGUgZ2l2ZW4gY2hpbGQgcmVmIHJlc3VsdHMuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgQ3JlYXRlTWFwcGluZyhyZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgc3RvcmFnZS5TZXJpZXNSZWYgewoJLy8gU3RhcnQgY2xlYW51cCBnb3JvdXRpbmUgb24gZmlyc3QgbWFwcGluZwoJcy5zdGFydFJlZkNsZWFudXAuRG8oZnVuYygpIHsKCQlzLmNsZWFudXBTdGFydGVkLlN0b3JlKHRydWUpCgkJZ28gcy5jbGVhbnVwU3RhbGVSZWZzKCkKCX0pCgoJLy8gU3RvcmUgYSBjb3B5IG9mIHRoZSBjaGlsZCByZWYgcmVzdWx0cyBkaXJlY3RseQoJY2hpbGRSZWZTbGljZSA6PSBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxlbihyZWZSZXN1bHRzKSkKCWNvcHkoY2hpbGRSZWZTbGljZSwgcmVmUmVzdWx0cykKCgkvLyBIYXNoIGxhYmVscyB0byBmb3IgdGhlIGZhbGxiYWNrIGxvb2t1cCB0YWJsZQoJbGFiZWxIYXNoIDo9IGxibHMuSGFzaCgpCgoJcy5yZWZNYXBwaW5nTXUuTG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoKCS8vIENyZWF0ZSBhIG5ldyB1bmlxdWUgcmVmCgl1bmlxdWVSZWYgOj0gcy5uZXh0VW5pcXVlUmVmCglzLm5leHRVbmlxdWVSZWYrKwoKCXMubGFiZWxIYXNoVG9VbmlxdWVSZWZbbGFiZWxIYXNoXSA9IHVuaXF1ZVJlZgoJcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdID0gdW5pcVJlZkNoaWxkcmVuewoJCWNoaWxkUmVmczogJmNoaWxkUmVmU2xpY2UsCgkJbGFiZWxIYXNoOiBsYWJlbEhhc2gsCgl9CgoJcy5hY3RpdmVNYXBwaW5ncy5JbmMoKQoJcy51bmlxdWVSZWZzVG90YWwuSW5jKCkKCglyZXR1cm4gdW5pcXVlUmVmCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgVXBkYXRlTWFwcGluZyh1bmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYsIHJlZlJlc3VsdHMgW11zdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSB7CglpZiB1bmlxdWVSZWYgPT0gMCB7CgkJcmV0dXJuCgl9CgoJY2hpbGRSZWZTbGljZSA6PSBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxlbihyZWZSZXN1bHRzKSkKCWNvcHkoY2hpbGRSZWZTbGljZSwgcmVmUmVzdWx0cykKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJLy8gRW5zdXJlIHRoYXQgbGFiZWwgaGFzaCBpbmRleCBpcyB1cCB0byBkYXRlIHRvIGhhbmRsZSBwb3NzaWJsZSBoYXNoIGNvbGxpc2lvbnMuCgkvLyBUT0RPOiBpcyB0aGlzIG5lY2Vzc2FyeT8KCW5ld0hhc2ggOj0gbGJscy5IYXNoKCkKCXByZXYsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXQoJaWYgb2sgJiYgcHJldi5sYWJlbEhhc2ggIT0gbmV3SGFzaCB7CgkJZGVsZXRlKHMubGFiZWxIYXNoVG9VbmlxdWVSZWYsIHByZXYubGFiZWxIYXNoKQoJCXMubGFiZWxIYXNoVG9VbmlxdWVSZWZbbmV3SGFzaF0gPSB1bmlxdWVSZWYKCX0KCglzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0gPSB1bmlxUmVmQ2hpbGRyZW57CgkJY2hpbGRSZWZzOiAmY2hpbGRSZWZTbGljZSwKCQlsYWJlbEhhc2g6IGxibHMuSGFzaCgpLAoJfQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIFRyYWNrQXBwZW5kZWRTZXJpZXModHMgaW50NjQsIGNlbGwgKkNlbGwpIHsKCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCWRlZmVyIHMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCWZvciBfLCByIDo9IHJhbmdlIGNlbGwuUmVmcyB7CgkJcy51bmlxdWVSZWZUaW1lc3RhbXBzW3JdID0gdHMKCX0KCglzLnRyYWNrZWRSZWZzLlNldChmbG9hdDY0KGxlbihzLnVuaXF1ZVJlZlRpbWVzdGFtcHMpKSkKCgljZWxsLlJlZnMgPSBjZWxsLlJlZnNbOjBdCglzLmNlbGxQb29sLlB1dChjZWxsKQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIEdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpICpDZWxsIHsKCXJldHVybiBzLmNlbGxQb29sLkdldCgpLigqQ2VsbCkKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBjbGVhbnVwU3RhbGVSZWZzKCkgewoJZGVmZXIgY2xvc2Uocy5jbGVhbnVwU3RvcHBlZCkKCgl0aWNrZXIgOj0gdGltZS5OZXdUaWNrZXIoMTUgKiB0aW1lLk1pbnV0ZSkKCWRlZmVyIHRpY2tlci5TdG9wKCkKCglmb3IgewoJCXNlbGVjdCB7CgkJY2FzZSA8LXRpY2tlci5DOgoJCQljdXRvZmZUaW1lIDo9IHRpbWUuTm93KCkuQWRkKC0xNSAqIHRpbWUuTWludXRlKS5Vbml4KCkKCgkJCS8vIEhvbGQgYm90aCBsb2NrcyB0byBwcmV2ZW50IHJhY2UgY29uZGl0aW9uIHdoZXJlIGEgcmVmIGNvdWxkIGJlCgkJCS8vIGFwcGVuZGVkIGFmdGVyIHdlIGRlbGV0ZSBpdCBmcm9tIHVuaXF1ZVJlZkNlbGwgYnV0IGJlZm9yZQoJCQkvLyB3ZSBkZWxldGUgaXQgZnJvbSB1bmlxdWVSZWZUb0NoaWxkUmVmcwoJCQlzLnRpbWVzdGFtcFRyYWNraW5nTXUuTG9jaygpCgkJCXMucmVmTWFwcGluZ011LkxvY2soKQoKCQkJc3RhbGVSZWZDb3VudCA6PSAwCgkJCWZvciByZWYsIHRzIDo9IHJhbmdlIHMudW5pcXVlUmVmVGltZXN0YW1wcyB7CgkJCQlpZiB0cyA8IGN1dG9mZlRpbWUgewoJCQkJCXN0YWxlUmVmQ291bnQrKwoKCQkJCQl2LCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3JlZl0KCQkJCQlpZiBvayB7CgkJCQkJCWRlbGV0ZShzLmxhYmVsSGFzaFRvVW5pcXVlUmVmLCB2LmxhYmVsSGFzaCkKCQkJCQl9CgoJCQkJCWRlbGV0ZShzLnVuaXF1ZVJlZlRpbWVzdGFtcHMsIHJlZikKCQkJCQlkZWxldGUocy51bmlxdWVSZWZUb0NoaWxkUmVmcywgcmVmKQoJCQkJfQoJCQl9CgoJCQkvLyBVcGRhdGUgbWV0cmljcwoJCQlpZiBzdGFsZVJlZkNvdW50ID4gMCB7CgkJCQlzLnJlZnNDbGVhbmVkLkFkZChmbG9hdDY0KHN0YWxlUmVmQ291bnQpKQoJCQkJcy5hY3RpdmVNYXBwaW5ncy5TdWIoZmxvYXQ2NChzdGFsZVJlZkNvdW50KSkKCQkJCXMudHJhY2tlZFJlZnMuU2V0KGZsb2F0NjQobGVuKHMudW5pcXVlUmVmVGltZXN0YW1wcykpKQoJCQl9CgoJCQlzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoJCQlzLnRpbWVzdGFtcFRyYWNraW5nTXUuVW5sb2NrKCkKCgkJY2FzZSA8LXMuc3RvcENsZWFudXA6CgkJCXJldHVybgoJCX0KCX0KfQoKLy8gQ2xlYXIgd2lsbCBjbGVhciBhbGwgaW50ZXJuYWwgbWFwcGluZ3MgYW5kIHN0b3AgdGhlIGNsZWFuZXIgZ29yb3V0aW5lIGlmIGl0IGlzIHJ1bm5pbmcuCi8vIEl0IGlzIHNhZmUgdG8gcmUtdXNlIHRoZSBzYW1lIGluc3RhbmNlIGFmdGVyIGNhbGxpbmcgQ2xlYXIuCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgQ2xlYXIoKSB7CgkvLyBTdG9wIHRoZSBjbGVhbnVwIGdvcm91dGluZSBhbmQgd2FpdCBmb3IgaXQgdG8gYmUgc3RvcHBlZCBzbyB3ZSBjYW4KCS8vIGF2b2lkIGEgcG9zc2libGUgZGVhZGxvY2sgd2l0aCBjbGVhbnVwIHRoYXQgYWxzbyBob2xkcyBib3RoIGxvY2tzCglpZiBzLmNsZWFudXBTdGFydGVkLkxvYWQoKSB7CgkJc2VsZWN0IHsKCQljYXNlIDwtcy5zdG9wQ2xlYW51cDoKCQkJLy8gQWxyZWFkeSBjbG9zZWQKCQlkZWZhdWx0OgoJCQljbG9zZShzLnN0b3BDbGVhbnVwKQoJCQk8LXMuY2xlYW51cFN0b3BwZWQKCQl9Cgl9CgoJLy8gV2UgbmVlZCB0byBob2xkIGJvdGggbG9ja3MgdG8gZG8gdGhpcyBzYWZlbHkgYW5kIHdlIGRvIGl0IGluIHRoZSBzYW1lIG9yZGVyIGFzCgkvLyBjbGVhbnVwU3RhbGVSZWZzLiBXZSBzdG9wcGVkIGFuZCB3YWl0ZWQgZm9yIHRoZSBiYWNrZ3JvdW5kIHdvcmtlciB0aGF0IGNhbGxzIGl0CgkvLyB0byBmaW5pc2ggYnV0IHNvbWUgZXh0cmEgc2FmZXR5IHdvbid0IGh1cnQuCglzLnRpbWVzdGFtcFRyYWNraW5nTXUuTG9jaygpCglkZWZlciBzLnRpbWVzdGFtcFRyYWNraW5nTXUuVW5sb2NrKCkKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJY2xlYXIocy51bmlxdWVSZWZUb0NoaWxkUmVmcykKCWNsZWFyKHMudW5pcXVlUmVmVGltZXN0YW1wcykKCgkvLyByZXNldCB0aGUgcG9vbAoJcy5jZWxsUG9vbCA9IHN5bmMuUG9vbHsKCQlOZXc6IGZ1bmMoKSBhbnkgewoJCQlyZXR1cm4gJkNlbGx7UmVmczogbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAwLCAxMDApfQoJCX0sCgl9CgoJLy8gTk9URTogV2UgZG8gTk9UIHJlc2V0IG5leHRVbmlxdWVSZWYgaGVyZS4gUmVzZXR0aW5nIGl0IHdvdWxkIGNhdXNlIHJlZiBjb2xsaXNpb25zCgkvLyB3aXRoIGNvbXBvbmVudHMgbGlrZSBwcm9tZXRoZXVzLnNjcmFwZSB3aGljaCB3aWxsIGtlZXAgcmUtc2VuZGluZyB0aGUgc2FtZSBjYWNoZWQgcmVmcy4KCS8vIFdlIGNvbnRpbnVlIGluY3JlbWVudGluZyB0byBlbnN1cmUgYWxsIHJlZnMgcmVtYWluIHVuaXF1ZSBhY3Jvc3MgdGhlIGxpZmV0aW1lIG9mIHRoZSBwcm9jZXNzLgoKCS8vIFJlc2V0IG1ldHJpY3MKCXMuYWN0aXZlTWFwcGluZ3MuU2V0KDApCglzLnRyYWNrZWRSZWZzLlNldCgwKQoKCS8vIFJlc2V0IGNoYW5uZWxzIGFuZCBmbGFncwoJcy5zdG9wQ2xlYW51cCA9IG1ha2UoY2hhbiBzdHJ1Y3R7fSkKCXMuY2xlYW51cFN0b3BwZWQgPSBtYWtlKGNoYW4gc3RydWN0e30pCglzLnN0YXJ0UmVmQ2xlYW51cCA9IHN5bmMuT25jZXt9CglzLmNsZWFudXBTdGFydGVkLlN0b3JlKGZhbHNlKQp9Cg=="
    }
  }
}
```

Response (packages.DriverResponse):
Error: `err: context canceled: stderr: `

#### drv #10

Trace meta: spanId=19, ts=1770837171658, ts_iso=2026-02-11T19:12:51.658000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy/internal/validator"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkiZm10IgoJInN5bmMiCgkic3luYy9hdG9taWMiCgkidGltZSIKCgkiZ2l0aHViLmNvbS9oYXNoaWNvcnAvZ28tbXVsdGllcnJvciIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvY2xpZW50X2dvbGFuZy9wcm9tZXRoZXVzIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2V4ZW1wbGFyIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2hpc3RvZ3JhbSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9sYWJlbHMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvbWV0YWRhdGEiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvc3RvcmFnZSIKKQoKdHlwZSBNYXBwaW5nU3RvcmUgaW50ZXJmYWNlIHsKCUdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYKCUNyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmCglVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpCglUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKQoJR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwKfQoKdHlwZSBzZXJpZXNSZWZNYXBwaW5nIHN0cnVjdCB7CglzdGFydCAgICB0aW1lLlRpbWUKCWNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlcgoJc3RvcmUgICAgTWFwcGluZ1N0b3JlCgoJdW5pcXVlUmVmQ2VsbCAqQ2VsbAoKCS8vIGNoaWxkUmVmcyBpcyByZXVzZWQgZm9yIGVhY2ggYXBwZW5kIGNhbGwgdG8gYXZvaWQgYWxsb2NhdGlvbnMuIFRoaXMgaXMgc2FmZSBiZWNhdXNlIHN0b3JhZ2UuQXBwZW5kZXIgc2hvdWxkIG5ldmVyCgkvLyBoYXZlIGNvbmN1cnJlbnQgY2FsbHMgdG8gQXBwZW5kIG1ldGhvZHMuCgljaGlsZFJlZnMgICAgICAgIFtdc3RvcmFnZS5TZXJpZXNSZWYKCXdyaXRlTGF0ZW5jeSAgICAgcHJvbWV0aGV1cy5IaXN0b2dyYW0KCXNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZyhjaGlsZHJlbiBbXXN0b3JhZ2UuQXBwZW5kZXIsIHN0b3JlIE1hcHBpbmdTdG9yZSwgd3JpdGVMYXRlbmN5IHByb21ldGhldXMuSGlzdG9ncmFtLCBzYW1wbGVzRm9yd2FyZGVkIHByb21ldGhldXMuQ291bnRlcikgc3RvcmFnZS5BcHBlbmRlciB7Cgl1bmlxdWVSZWZDZWxsIDo9IHN0b3JlLkdldENlbGxGb3JBcHBlbmRlZFNlcmllcygpCgoJcmV0dXJuICZzZXJpZXNSZWZNYXBwaW5newoJCWNoaWxkcmVuOiAgICAgICAgIGNoaWxkcmVuLAoJCXN0b3JlOiAgICAgICAgICAgIHN0b3JlLAoJCXdyaXRlTGF0ZW5jeTogICAgIHdyaXRlTGF0ZW5jeSwKCQlzYW1wbGVzRm9yd2FyZGVkOiBzYW1wbGVzRm9yd2FyZGVkLAoKCQl1bmlxdWVSZWZDZWxsOiB1bmlxdWVSZWZDZWxsLAoJCWNoaWxkUmVmczogICAgIG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgbGVuKGNoaWxkcmVuKSksCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFNldE9wdGlvbnMob3B0cyAqc3RvcmFnZS5BcHBlbmRPcHRpb25zKSB7CglmbXQuUHJpbnRsbigiYXNkYXNkKQoJZm9yIF8sIGMgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJYy5TZXRPcHRpb25zKG9wdHMpCgl9Cn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIENvbW1pdCgpIGVycm9yIHsKCXMuc3RvcmUuVHJhY2tBcHBlbmRlZFNlcmllcyh0aW1lLk5vdygpLlVuaXgoKSwgcy51bmlxdWVSZWZDZWxsKQoKCXZhciBtdWx0aUVyciBlcnJvcgoJZm9yIF8sIGMgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJZXJyIDo9IGMuQ29tbWl0KCkKCQlpZiBlcnIgIT0gbmlsIHsKCQkJbXVsdGlFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChtdWx0aUVyciwgZXJyKQoJCX0KCX0KCXJldHVybiBtdWx0aUVycgp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBSb2xsYmFjaygpIGVycm9yIHsKCS8vIFdlIHN0aWxsIHRyYWNrIHJvbGxlZCBiYWNrIHNlcmllcyBzbyB3ZSBjYW4gcHJvcGVybHkKCS8vIGNsZWFuIHVwIGFueSBzZXJpZXMgdGhhdCB3YXMgYXBwZW5kZWQKCXMuc3RvcmUuVHJhY2tBcHBlbmRlZFNlcmllcyh0aW1lLk5vdygpLlVuaXgoKSwgcy51bmlxdWVSZWZDZWxsKQoKCXZhciBtdWx0aUVyciBlcnJvcgoJZm9yIF8sIGMgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJZXJyIDo9IGMuUm9sbGJhY2soKQoJCWlmIGVyciAhPSBuaWwgewoJCQltdWx0aUVyciA9IG11bHRpZXJyb3IuQXBwZW5kKG11bHRpRXJyLCBlcnIpCgkJfQoJfQoJcmV0dXJuIG11bHRpRXJyCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIHJlY29yZExhdGVuY3koKSB7CglpZiBzLnN0YXJ0LklzWmVybygpIHsKCQlyZXR1cm4KCX0KCWR1cmF0aW9uIDo9IHRpbWUuU2luY2Uocy5zdGFydCkKCXMud3JpdGVMYXRlbmN5Lk9ic2VydmUoZHVyYXRpb24uU2Vjb25kcygpKQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSByZXNldEZpZWxkcygpIHsKCS8vIFJlc2V0IGNoaWxkUmVmcyBzbGljZSBsZW5ndGggdG8gMCBmb3IgcmV1c2UKCXMuY2hpbGRSZWZzID0gcy5jaGlsZFJlZnNbOjBdCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZChyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCBpbnQ2NCwgdiBmbG9hdDY0KSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQluZXdSZWYsIGVyciA6PSBhcHBlbmRlci5BcHBlbmQocmVmLCBsLCB0LCB2KQoJCWlmIGVyciA9PSBuaWwgewoJCQlzLnNhbXBsZXNGb3J3YXJkZWQuSW5jKCkKCQl9CgkJcmV0dXJuIG5ld1JlZiwgZXJyCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRFeGVtcGxhcihyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgZSBleGVtcGxhci5FeGVtcGxhcikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZEV4ZW1wbGFyKHJlZiwgbCwgZSkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEhpc3RvZ3JhbShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCBpbnQ2NCwgaCAqaGlzdG9ncmFtLkhpc3RvZ3JhbSwgZmggKmhpc3RvZ3JhbS5GbG9hdEhpc3RvZ3JhbSkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZEhpc3RvZ3JhbShyZWYsIGwsIHQsIGgsIGZoKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kSGlzdG9ncmFtQ1RaZXJvU2FtcGxlKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0LCBjdCBpbnQ2NCwgaCAqaGlzdG9ncmFtLkhpc3RvZ3JhbSwgZmggKmhpc3RvZ3JhbS5GbG9hdEhpc3RvZ3JhbSkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZEhpc3RvZ3JhbUNUWmVyb1NhbXBsZShyZWYsIGwsIHQsIGN0LCBoLCBmaCkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIFVwZGF0ZU1ldGFkYXRhKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCBtIG1ldGFkYXRhLk1ldGFkYXRhKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglyZXR1cm4gcy5hcHBlbmRUb0NoaWxkcmVuKHJlZiwgbCwgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCQlyZXR1cm4gYXBwZW5kZXIuVXBkYXRlTWV0YWRhdGEocmVmLCBsLCBtKQoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kQ1RaZXJvU2FtcGxlKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbCBsYWJlbHMuTGFiZWxzLCB0LCBjdCBpbnQ2NCkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLkFwcGVuZENUWmVyb1NhbXBsZShyZWYsIGwsIHQsIGN0KQoJfSkKfQoKdHlwZSBhcHBlbmRlckZ1bmMgZnVuYyhhcHBlbmRlciBzdG9yYWdlLkFwcGVuZGVyLCByZWYgc3RvcmFnZS5TZXJpZXNSZWYpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpCgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBhcHBlbmRUb0NoaWxkcmVuKHJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzLCBhZiBhcHBlbmRlckZ1bmMpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCWRlZmVyIHMucmVzZXRGaWVsZHMoKQoKCWlmIHMuc3RhcnQuSXNaZXJvKCkgewoJCXMuc3RhcnQgPSB0aW1lLk5vdygpCgl9CgoJLy8gQ2hlY2sgaWYgdGhlIGluY29taW5nIHJlZiBoYXMgcmVmIG1hcHBpbmdzCglleGlzdGluZ0NoaWxkUmVmcyA6PSBzLnN0b3JlLkdldE1hcHBpbmcocmVmLCBsYmxzKQoKCXZhciBhcHBlbmRFcnIgZXJyb3IKCgkvLyBTYW5pdHkgY2hlY2s6IGlmIHdlIGhhdmUgZXhpc3RpbmcgY2hpbGQgcmVmcywgdGhleSBtdXN0IG1hdGNoIHRoZSBudW1iZXIgb2YgY2hpbGRyZW4KCWlmIGV4aXN0aW5nQ2hpbGRSZWZzICE9IG5pbCAmJiBsZW4oZXhpc3RpbmdDaGlsZFJlZnMpID09IGxlbihzLmNoaWxkcmVuKSB7CgkJcy51bmlxdWVSZWZDZWxsLlJlZnMgPSBhcHBlbmQocy51bmlxdWVSZWZDZWxsLlJlZnMsIHJlZikKCgkJcmVmVXBkYXRlUmVxdWlyZWQgOj0gZmFsc2UKCQlmb3IgY2hpbGRJbmRleCwgY2hpbGRSZWYgOj0gcmFuZ2UgZXhpc3RpbmdDaGlsZFJlZnMgewoJCQluZXdDaGlsZFJlZiwgZXJyIDo9IGFmKHMuY2hpbGRyZW5bY2hpbGRJbmRleF0sIGNoaWxkUmVmKQoJCQlpZiBlcnIgIT0gbmlsIHsKCQkJCWFwcGVuZEVyciA9IG11bHRpZXJyb3IuQXBwZW5kKGFwcGVuZEVyciwgZXJyKQoJCQl9CgoJCQlpZiBuZXdDaGlsZFJlZiAhPSBjaGlsZFJlZiB7CgkJCQkvLyBDaGlsZCByZWYgY2hhbmdlZCwgbmVlZCB0byB1cGRhdGUgbWFwcGluZwoJCQkJZXhpc3RpbmdDaGlsZFJlZnNbY2hpbGRJbmRleF0gPSBuZXdDaGlsZFJlZgoJCQkJcmVmVXBkYXRlUmVxdWlyZWQgPSB0cnVlCgkJCX0KCQl9CgoJCWlmIGFwcGVuZEVyciAhPSBuaWwgewoJCQlyZXR1cm4gMCwgYXBwZW5kRXJyCgkJfQoKCQlpZiByZWZVcGRhdGVSZXF1aXJlZCB7CgkJCXMuc3RvcmUuVXBkYXRlTWFwcGluZyhyZWYsIGV4aXN0aW5nQ2hpbGRSZWZzLCBsYmxzKQoJCX0KCgkJcmV0dXJuIHJlZiwgbmlsCgl9CgoJLy8gTm8gZXhpc3RpbmcgbWFwcGluZywgcHJvY2VlZCB3aXRoIG5vcm1hbCBhcHBlbmQgdG8gYWxsIGNoaWxkcmVuCgl2YXIgZmlyc3ROb25aZXJvUmVmIHN0b3JhZ2UuU2VyaWVzUmVmCgl2YXIgbm9uWmVyb0NvdW50IGludAoKCS8vIE5vdGU6IHRoZXJlJ3MgYW5vdGhlciBvcHRpbWl6YXRpb24gd2hlcmUgd2UgY291bGQgdXNlIHRoZSByZXR1cm5lZCByZWYgaWYgYWxsIHRoZSBub24gemVybyByZWZzCgkvLyAgYXJlIHRoZSBzYW1lIHZhbHVlLiBUaGlzIGlzbid0IHNhZmUgYXMgd2Ugd2lsbCBtaXggZG93bnN0cmVhbSByZWZzIHdpdGggdW5pcXVlIHJlZnMgd2hpY2ggY291bGQKCS8vICBjb2xsaWRlLiBXZSBjb3VsZCBzdGFydCBhdCBtYXggdW5pdDY0IGZvciB1bmlxdWUgcmVmcyBhbmQgZ28gYmFja3dhcmRzIGxlc3NlbmluZyB0aGUgY2hhbmNlIG9mCgkvLyAJY29sbGlzaW9ucyBidXQgaXQncyByYXRoZXIgZGFuZ2Vyb3VzIGZvciBhbiB1bmxpa2VseSBlZGdlIGNhc2UuIElmIHR3byBjb21wb25lbnRzIGFyZSByZXR1cm5pbmcKCS8vIAl0aGUgc2FtZSByZWYgaXQncyB0d28gcmVtb3RlX3dyaXRlIGNvbXBvbmVudHMgd2hpY2ggc2hvdWxkIHByb2JhYmx5IGJlIG1lcmdlZCBpbiB0byBvbmUuCglmb3IgXywgY2hpbGQgOj0gcmFuZ2Ugcy5jaGlsZHJlbiB7CgkJY2hpbGRSZWYsIGVyciA6PSBhZihjaGlsZCwgcmVmKQoJCWlmIGVyciAhPSBuaWwgewoJCQlhcHBlbmRFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChhcHBlbmRFcnIsIGVycikKCgkJCS8vIFRPRE8gc2hvdWxkIEkgbWFrZSB0aGUgY2hpbGRSZWYgemVybyBoZXJlPwoJCX0KCgkJcy5jaGlsZFJlZnMgPSBhcHBlbmQocy5jaGlsZFJlZnMsIGNoaWxkUmVmKQoJCWlmIGNoaWxkUmVmICE9IDAgewoJCQlpZiBmaXJzdE5vblplcm9SZWYgPT0gMCB7CgkJCQlmaXJzdE5vblplcm9SZWYgPSBjaGlsZFJlZgoJCQl9CgkJCW5vblplcm9Db3VudCsrCgkJfQoJfQoKCWlmIGFwcGVuZEVyciAhPSBuaWwgewoJCXJldHVybiAwLCBhcHBlbmRFcnIKCX0KCglpZiBub25aZXJvQ291bnQgPT0gMCB7CgkJLy8gQWxsIGNoaWxkcmVuIHJldHVybmVkIHJlZiAwLCBzbyByZXR1cm4gdGhlIGlucHV0IHJlZgoJCXJldHVybiByZWYsIG5pbAoJfQoKCS8vIE9ubHkgb25lIGNoaWxkIHJldHVybmVkIGEgbm9uLXplcm8gcmVmLCB1c2UgdGhhdAoJaWYgbm9uWmVyb0NvdW50ID09IDEgewoJCXJldHVybiBmaXJzdE5vblplcm9SZWYsIG5pbAoJfQoKCS8vIFdlIGdvdCBkaWZmZXJlbnQgcmVmcyBiYWNrIGFuZCBuZWVkIHRvIGNyZWF0ZSBhIG5ldyBtYXBwaW5nCgl1bmlxdWVSZWYgOj0gcy5zdG9yZS5DcmVhdGVNYXBwaW5nKHMuY2hpbGRSZWZzLCBsYmxzKQoJcy51bmlxdWVSZWZDZWxsLlJlZnMgPSBhcHBlbmQocy51bmlxdWVSZWZDZWxsLlJlZnMsIHVuaXF1ZVJlZikKCXJldHVybiB1bmlxdWVSZWYsIG5pbAp9Cgp0eXBlIHVuaXFSZWZDaGlsZHJlbiBzdHJ1Y3QgewoJY2hpbGRSZWZzICpbXXN0b3JhZ2UuU2VyaWVzUmVmCglsYWJlbEhhc2ggdWludDY0Cn0KCnR5cGUgU2VyaWVzUmVmTWFwcGluZ1N0b3JlIHN0cnVjdCB7CgkvLyByZWZNYXBwaW5nTXUgcHJvdGVjdHMgdW5pcXVlUmVmVG9DaGlsZFJlZnMsIGxhYmVsSGFzaFRvVW5pcXVlUmVmIGFuZCBuZXh0VW5pcXVlUmVmCglyZWZNYXBwaW5nTXUgc3luYy5SV011dGV4CgkvLyB1bmlxdWVSZWZUb0NoaWxkUmVmcyBtYXBzIHRoZSB1bmlxdWUgcmVmIHRvIHRoZSBleHBlY3RlZCBjaGlsZCByZWYgaW4gb3JkZXIKCXVuaXF1ZVJlZlRvQ2hpbGRSZWZzIG1hcFtzdG9yYWdlLlNlcmllc1JlZl11bmlxUmVmQ2hpbGRyZW4KCS8vIGxhYmVsSGFzaFRvVW5pcXVlUmVmIG1hcHMgdGhlIGxhYmVsIGhhc2ggdG8gdW5pcXVlIHJlZi4KCWxhYmVsSGFzaFRvVW5pcXVlUmVmIG1hcFt1aW50NjRdc3RvcmFnZS5TZXJpZXNSZWYKCgkvLyBuZXh0VW5pcXVlUmVmIGlzIHRoZSBuZXh0IHJlZiBJRCB3ZSB3aWxsIGhhbmQgb3V0CgluZXh0VW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmCgoJLy8gdGltZXN0YW1wVHJhY2tpbmdNdSBwcm90ZWN0cyB1bmlxdWVSZWZUaW1lc3RhbXBzIGFuZCBjZWxsUG9vbAoJdGltZXN0YW1wVHJhY2tpbmdNdSBzeW5jLk11dGV4CgkvLyB1bmlxdWVSZWZUaW1lc3RhbXBzIG1hcHMgdW5pcXVlIHJlZnMgdG8gdGhlaXIgbGFzdCBhcHBlbmQgdGltZXN0YW1wCgl1bmlxdWVSZWZUaW1lc3RhbXBzIG1hcFtzdG9yYWdlLlNlcmllc1JlZl1pbnQ2NAoJLy8gY2VsbFBvb2wgaXMgdXNlZCB0byBwb29sIHNsaWNlcyBvZiBTZXJpZXNSZWZzIHVzZWQgZm9yIHRyYWNraW5nIHVuaXF1ZSByZWZzIGluIFRyYWNrQXBwZW5kZWRTZXJpZXMuCgljZWxsUG9vbCBzeW5jLlBvb2wKCgkvLyBDbGVhbnVwIGdvcm91dGluZSBjb29yZGluYXRpb24gKG5vIGxvY2sgcmVxdWlyZWQpCglzdGFydFJlZkNsZWFudXAgc3luYy5PbmNlCgljbGVhbnVwU3RhcnRlZCAgYXRvbWljLkJvb2wKCXN0b3BDbGVhbnVwICAgICBjaGFuIHN0cnVjdHt9CgljbGVhbnVwU3RvcHBlZCAgY2hhbiBzdHJ1Y3R7fQoKCS8vIE1ldHJpY3MgKHNhZmUgZm9yIGNvbmN1cnJlbnQgYWNjZXNzLCBubyBsb2NrIHJlcXVpcmVkKQoJYWN0aXZlTWFwcGluZ3MgIHByb21ldGhldXMuR2F1Z2UKCXRyYWNrZWRSZWZzICAgICBwcm9tZXRoZXVzLkdhdWdlCglyZWZzQ2xlYW5lZCAgICAgcHJvbWV0aGV1cy5Db3VudGVyCgl1bmlxdWVSZWZzVG90YWwgcHJvbWV0aGV1cy5Db3VudGVyCn0KCmZ1bmMgTmV3U2VyaWVzUmVmTWFwcGluZ1N0b3JlKHJlZyBwcm9tZXRoZXVzLlJlZ2lzdGVyZXIpICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUgewoJYWN0aXZlTWFwcGluZ3MgOj0gcHJvbWV0aGV1cy5OZXdHYXVnZShwcm9tZXRoZXVzLkdhdWdlT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfbWFwcGluZ3NfdG90YWwiLAoJCUhlbHA6ICJOdW1iZXIgb2YgYWN0aXZlIHVuaXF1ZSByZWYgbWFwcGluZ3MgaW4gdGhlIHN0b3JlLiIsCgl9KQoJdHJhY2tlZFJlZnMgOj0gcHJvbWV0aGV1cy5OZXdHYXVnZShwcm9tZXRoZXVzLkdhdWdlT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfdHJhY2tlZF9yZWZzX3RvdGFsIiwKCQlIZWxwOiAiTnVtYmVyIG9mIHJlZnMgYmVpbmcgdHJhY2tlZCBmb3IgdGltZXN0YW1wLWJhc2VkIGNsZWFudXAuIiwKCX0pCglyZWZzQ2xlYW5lZCA6PSBwcm9tZXRoZXVzLk5ld0NvdW50ZXIocHJvbWV0aGV1cy5Db3VudGVyT3B0c3sKCQlOYW1lOiAiYWxsb3lfZmFub3V0X21hcHBpbmdfc3RvcmVfcmVmc19jbGVhbmVkX3RvdGFsIiwKCQlIZWxwOiAiVG90YWwgbnVtYmVyIG9mIHN0YWxlIHJlZnMgY2xlYW5lZCB1cCBvdmVyIHRpbWUuIiwKCX0pCgl1bmlxdWVSZWZzVG90YWwgOj0gcHJvbWV0aGV1cy5OZXdDb3VudGVyKHByb21ldGhldXMuQ291bnRlck9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3VuaXF1ZV9yZWZzX2NyZWF0ZWRfdG90YWwiLAoJCUhlbHA6ICJUb3RhbCBudW1iZXIgb2YgdW5pcXVlIHJlZnMgY3JlYXRlZC4iLAoJfSkKCglpZiByZWcgIT0gbmlsIHsKCQlyZWcuUmVnaXN0ZXIoYWN0aXZlTWFwcGluZ3MpCgkJcmVnLlJlZ2lzdGVyKHRyYWNrZWRSZWZzKQoJCXJlZy5SZWdpc3RlcihyZWZzQ2xlYW5lZCkKCQlyZWcuUmVnaXN0ZXIodW5pcXVlUmVmc1RvdGFsKQoJfQoKCXJldHVybiAmU2VyaWVzUmVmTWFwcGluZ1N0b3JlewoJCXVuaXF1ZVJlZlRvQ2hpbGRSZWZzOiBtYWtlKG1hcFtzdG9yYWdlLlNlcmllc1JlZl11bmlxUmVmQ2hpbGRyZW4pLAoJCW5leHRVbmlxdWVSZWY6ICAgICAgICAxLAoJCXVuaXF1ZVJlZlRpbWVzdGFtcHM6ICBtYWtlKG1hcFtzdG9yYWdlLlNlcmllc1JlZl1pbnQ2NCksCgkJbGFiZWxIYXNoVG9VbmlxdWVSZWY6IG1ha2UobWFwW3VpbnQ2NF1zdG9yYWdlLlNlcmllc1JlZiksCgkJY2VsbFBvb2w6IHN5bmMuUG9vbHsKCQkJTmV3OiBmdW5jKCkgYW55IHsKCQkJCXJldHVybiAmQ2VsbHtSZWZzOiBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIDEwMCl9CgkJCX0sCgkJfSwKCQlzdG9wQ2xlYW51cDogICAgIG1ha2UoY2hhbiBzdHJ1Y3R7fSksCgkJY2xlYW51cFN0b3BwZWQ6ICBtYWtlKGNoYW4gc3RydWN0e30pLAoJCWFjdGl2ZU1hcHBpbmdzOiAgYWN0aXZlTWFwcGluZ3MsCgkJdHJhY2tlZFJlZnM6ICAgICB0cmFja2VkUmVmcywKCQlyZWZzQ2xlYW5lZDogICAgIHJlZnNDbGVhbmVkLAoJCXVuaXF1ZVJlZnNUb3RhbDogdW5pcXVlUmVmc1RvdGFsLAoJfQp9Cgp0eXBlIENlbGwgc3RydWN0IHsKCVJlZnMgW11zdG9yYWdlLlNlcmllc1JlZgp9CgovLyBHZXRNYXBwaW5nIHJldHVybnMgZXhpc3RpbmcgY2hpbGQgcmVmIHJlc3VsdHMgZm9yIHRoZSBnaXZlbiB1bmlxdWUgcmVmIGlmIG9uZSBleGlzdHMuCi8vCi8vIElmIHRoZSBwYXNzZWQgdW5pcXVlUmVmIGlzIHplcm8sIHRoZSBtZXRob2Qgd2lsbCBhdHRlbXB0IHRvIGZpbmQgYSBtYXBwaW5nIHVzaW5nIHBhc3NlZCBsYWJlbHMuCi8vIFJldHVybnMgbmlsIGlmIG5vIG1hcHBpbmcgZXhpc3RzLgovLwovLyBUaGUgcmV0dXJuZWQgc2xpY2UgbWF5IGJlIG1vZGlmaWVkIGJ5IHRoZSBjYWxsZXIsIGJ1dCBVcGRhdGVNYXBwaW5nIG11c3QgYmUgY2FsbGVkCi8vIGFmdGVyd2FyZHMgdG8gcGVyc2lzdCBjaGFuZ2VzLiBOb3RlIHRoYXQgY29uY3VycmVudCBhcHBlbmRlcnMgbWF5IHJhY2UgdG8gdXBkYXRlIHRoZQovLyBzYW1lIG1hcHBpbmcgd2l0aCBkaWZmZXJlbnQgdmFsdWVzLCB3aGljaCBpcyBzYWZlIGJlY2F1c2Ugc3RhbGUgbWFwcGluZ3MgYXJlIHNlbGYtY29ycmVjdGluZyAtCi8vIHVzaW5nIGEgc3RhbGUgcmVmIHdpbGwgY2F1c2UgdGhlIGNoaWxkIGFwcGVuZGVyIHRvIHJldHVybiBhIG5ldyByZWYgb24gdGhlIG5leHQgYXBwZW5kLgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIEdldE1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIFtdc3RvcmFnZS5TZXJpZXNSZWYgewoJcy5yZWZNYXBwaW5nTXUuUkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuUlVubG9jaygpCgoJaWYgdW5pcXVlUmVmID09IDAgewoJCS8vIFNvbWUgY29uc3VtZXJzIGRvbid0IG1lbW8gdGhlIGdsb2JhbCByZWYuIFRyeSB0byBsb29rdXAgYSByZWYgYnkgbGFiZWwgaGFzaC4KCQlsYWJlbEhhc2ggOj0gbGJscy5IYXNoKCkKCQlnb3RSZWYsIG9rIDo9IHMubGFiZWxIYXNoVG9VbmlxdWVSZWZbbGFiZWxIYXNoXQoJCWlmICFvayB7CgkJCXJldHVybiBuaWwKCQl9CgoJCXVuaXF1ZVJlZiA9IGdvdFJlZgoJfQoKCWlmIG1hcHBpbmcsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXTsgb2sgewoJCXJldHVybiAqbWFwcGluZy5jaGlsZFJlZnMKCX0KCXJldHVybiBuaWwKfQoKLy8gQ3JlYXRlTWFwcGluZyBjcmVhdGVzIGEgbmV3IHVuaXF1ZSByZWYgbWFwcGluZyBmb3IgdGhlIGdpdmVuIGNoaWxkIHJlZiByZXN1bHRzLgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIENyZWF0ZU1hcHBpbmcocmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHN0b3JhZ2UuU2VyaWVzUmVmIHsKCS8vIFN0YXJ0IGNsZWFudXAgZ29yb3V0aW5lIG9uIGZpcnN0IG1hcHBpbmcKCXMuc3RhcnRSZWZDbGVhbnVwLkRvKGZ1bmMoKSB7CgkJcy5jbGVhbnVwU3RhcnRlZC5TdG9yZSh0cnVlKQoJCWdvIHMuY2xlYW51cFN0YWxlUmVmcygpCgl9KQoKCS8vIFN0b3JlIGEgY29weSBvZiB0aGUgY2hpbGQgcmVmIHJlc3VsdHMgZGlyZWN0bHkKCWNoaWxkUmVmU2xpY2UgOj0gbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsZW4ocmVmUmVzdWx0cykpCgljb3B5KGNoaWxkUmVmU2xpY2UsIHJlZlJlc3VsdHMpCgoJLy8gSGFzaCBsYWJlbHMgdG8gZm9yIHRoZSBmYWxsYmFjayBsb29rdXAgdGFibGUKCWxhYmVsSGFzaCA6PSBsYmxzLkhhc2goKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgkvLyBDcmVhdGUgYSBuZXcgdW5pcXVlIHJlZgoJdW5pcXVlUmVmIDo9IHMubmV4dFVuaXF1ZVJlZgoJcy5uZXh0VW5pcXVlUmVmKysKCglzLmxhYmVsSGFzaFRvVW5pcXVlUmVmW2xhYmVsSGFzaF0gPSB1bmlxdWVSZWYKCXMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXSA9IHVuaXFSZWZDaGlsZHJlbnsKCQljaGlsZFJlZnM6ICZjaGlsZFJlZlNsaWNlLAoJCWxhYmVsSGFzaDogbGFiZWxIYXNoLAoJfQoKCXMuYWN0aXZlTWFwcGluZ3MuSW5jKCkKCXMudW5pcXVlUmVmc1RvdGFsLkluYygpCgoJcmV0dXJuIHVuaXF1ZVJlZgp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIFVwZGF0ZU1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCByZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgewoJaWYgdW5pcXVlUmVmID09IDAgewoJCXJldHVybgoJfQoKCWNoaWxkUmVmU2xpY2UgOj0gbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsZW4ocmVmUmVzdWx0cykpCgljb3B5KGNoaWxkUmVmU2xpY2UsIHJlZlJlc3VsdHMpCgoJcy5yZWZNYXBwaW5nTXUuTG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoKCS8vIEVuc3VyZSB0aGF0IGxhYmVsIGhhc2ggaW5kZXggaXMgdXAgdG8gZGF0ZSB0byBoYW5kbGUgcG9zc2libGUgaGFzaCBjb2xsaXNpb25zLgoJLy8gVE9ETzogaXMgdGhpcyBuZWNlc3Nhcnk/CgluZXdIYXNoIDo9IGxibHMuSGFzaCgpCglwcmV2LCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0KCWlmIG9rICYmIHByZXYubGFiZWxIYXNoICE9IG5ld0hhc2ggewoJCWRlbGV0ZShzLmxhYmVsSGFzaFRvVW5pcXVlUmVmLCBwcmV2LmxhYmVsSGFzaCkKCQlzLmxhYmVsSGFzaFRvVW5pcXVlUmVmW25ld0hhc2hdID0gdW5pcXVlUmVmCgl9CgoJcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdID0gdW5pcVJlZkNoaWxkcmVuewoJCWNoaWxkUmVmczogJmNoaWxkUmVmU2xpY2UsCgkJbGFiZWxIYXNoOiBsYmxzLkhhc2goKSwKCX0KfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBUcmFja0FwcGVuZGVkU2VyaWVzKHRzIGludDY0LCBjZWxsICpDZWxsKSB7CglzLnRpbWVzdGFtcFRyYWNraW5nTXUuTG9jaygpCglkZWZlciBzLnRpbWVzdGFtcFRyYWNraW5nTXUuVW5sb2NrKCkKCglmb3IgXywgciA6PSByYW5nZSBjZWxsLlJlZnMgewoJCXMudW5pcXVlUmVmVGltZXN0YW1wc1tyXSA9IHRzCgl9CgoJcy50cmFja2VkUmVmcy5TZXQoZmxvYXQ2NChsZW4ocy51bmlxdWVSZWZUaW1lc3RhbXBzKSkpCgoJY2VsbC5SZWZzID0gY2VsbC5SZWZzWzowXQoJcy5jZWxsUG9vbC5QdXQoY2VsbCkKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBHZXRDZWxsRm9yQXBwZW5kZWRTZXJpZXMoKSAqQ2VsbCB7CglyZXR1cm4gcy5jZWxsUG9vbC5HZXQoKS4oKkNlbGwpCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgY2xlYW51cFN0YWxlUmVmcygpIHsKCWRlZmVyIGNsb3NlKHMuY2xlYW51cFN0b3BwZWQpCgoJdGlja2VyIDo9IHRpbWUuTmV3VGlja2VyKDE1ICogdGltZS5NaW51dGUpCglkZWZlciB0aWNrZXIuU3RvcCgpCgoJZm9yIHsKCQlzZWxlY3QgewoJCWNhc2UgPC10aWNrZXIuQzoKCQkJY3V0b2ZmVGltZSA6PSB0aW1lLk5vdygpLkFkZCgtMTUgKiB0aW1lLk1pbnV0ZSkuVW5peCgpCgoJCQkvLyBIb2xkIGJvdGggbG9ja3MgdG8gcHJldmVudCByYWNlIGNvbmRpdGlvbiB3aGVyZSBhIHJlZiBjb3VsZCBiZQoJCQkvLyBhcHBlbmRlZCBhZnRlciB3ZSBkZWxldGUgaXQgZnJvbSB1bmlxdWVSZWZDZWxsIGJ1dCBiZWZvcmUKCQkJLy8gd2UgZGVsZXRlIGl0IGZyb20gdW5pcXVlUmVmVG9DaGlsZFJlZnMKCQkJcy50aW1lc3RhbXBUcmFja2luZ011LkxvY2soKQoJCQlzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCgkJCXN0YWxlUmVmQ291bnQgOj0gMAoJCQlmb3IgcmVmLCB0cyA6PSByYW5nZSBzLnVuaXF1ZVJlZlRpbWVzdGFtcHMgewoJCQkJaWYgdHMgPCBjdXRvZmZUaW1lIHsKCQkJCQlzdGFsZVJlZkNvdW50KysKCgkJCQkJdiwgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1tyZWZdCgkJCQkJaWYgb2sgewoJCQkJCQlkZWxldGUocy5sYWJlbEhhc2hUb1VuaXF1ZVJlZiwgdi5sYWJlbEhhc2gpCgkJCQkJfQoKCQkJCQlkZWxldGUocy51bmlxdWVSZWZUaW1lc3RhbXBzLCByZWYpCgkJCQkJZGVsZXRlKHMudW5pcXVlUmVmVG9DaGlsZFJlZnMsIHJlZikKCQkJCX0KCQkJfQoKCQkJLy8gVXBkYXRlIG1ldHJpY3MKCQkJaWYgc3RhbGVSZWZDb3VudCA+IDAgewoJCQkJcy5yZWZzQ2xlYW5lZC5BZGQoZmxvYXQ2NChzdGFsZVJlZkNvdW50KSkKCQkJCXMuYWN0aXZlTWFwcGluZ3MuU3ViKGZsb2F0NjQoc3RhbGVSZWZDb3VudCkpCgkJCQlzLnRyYWNrZWRSZWZzLlNldChmbG9hdDY0KGxlbihzLnVuaXF1ZVJlZlRpbWVzdGFtcHMpKSkKCQkJfQoKCQkJcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCQkJcy50aW1lc3RhbXBUcmFja2luZ011LlVubG9jaygpCgoJCWNhc2UgPC1zLnN0b3BDbGVhbnVwOgoJCQlyZXR1cm4KCQl9Cgl9Cn0KCi8vIENsZWFyIHdpbGwgY2xlYXIgYWxsIGludGVybmFsIG1hcHBpbmdzIGFuZCBzdG9wIHRoZSBjbGVhbmVyIGdvcm91dGluZSBpZiBpdCBpcyBydW5uaW5nLgovLyBJdCBpcyBzYWZlIHRvIHJlLXVzZSB0aGUgc2FtZSBpbnN0YW5jZSBhZnRlciBjYWxsaW5nIENsZWFyLgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIENsZWFyKCkgewoJLy8gU3RvcCB0aGUgY2xlYW51cCBnb3JvdXRpbmUgYW5kIHdhaXQgZm9yIGl0IHRvIGJlIHN0b3BwZWQgc28gd2UgY2FuCgkvLyBhdm9pZCBhIHBvc3NpYmxlIGRlYWRsb2NrIHdpdGggY2xlYW51cCB0aGF0IGFsc28gaG9sZHMgYm90aCBsb2NrcwoJaWYgcy5jbGVhbnVwU3RhcnRlZC5Mb2FkKCkgewoJCXNlbGVjdCB7CgkJY2FzZSA8LXMuc3RvcENsZWFudXA6CgkJCS8vIEFscmVhZHkgY2xvc2VkCgkJZGVmYXVsdDoKCQkJY2xvc2Uocy5zdG9wQ2xlYW51cCkKCQkJPC1zLmNsZWFudXBTdG9wcGVkCgkJfQoJfQoKCS8vIFdlIG5lZWQgdG8gaG9sZCBib3RoIGxvY2tzIHRvIGRvIHRoaXMgc2FmZWx5IGFuZCB3ZSBkbyBpdCBpbiB0aGUgc2FtZSBvcmRlciBhcwoJLy8gY2xlYW51cFN0YWxlUmVmcy4gV2Ugc3RvcHBlZCBhbmQgd2FpdGVkIGZvciB0aGUgYmFja2dyb3VuZCB3b3JrZXIgdGhhdCBjYWxscyBpdAoJLy8gdG8gZmluaXNoIGJ1dCBzb21lIGV4dHJhIHNhZmV0eSB3b24ndCBodXJ0LgoJcy50aW1lc3RhbXBUcmFja2luZ011LkxvY2soKQoJZGVmZXIgcy50aW1lc3RhbXBUcmFja2luZ011LlVubG9jaygpCgoJcy5yZWZNYXBwaW5nTXUuTG9jaygpCglkZWZlciBzLnJlZk1hcHBpbmdNdS5VbmxvY2soKQoKCWNsZWFyKHMudW5pcXVlUmVmVG9DaGlsZFJlZnMpCgljbGVhcihzLnVuaXF1ZVJlZlRpbWVzdGFtcHMpCgoJLy8gcmVzZXQgdGhlIHBvb2wKCXMuY2VsbFBvb2wgPSBzeW5jLlBvb2x7CgkJTmV3OiBmdW5jKCkgYW55IHsKCQkJcmV0dXJuICZDZWxse1JlZnM6IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgMCwgMTAwKX0KCQl9LAoJfQoKCS8vIE5PVEU6IFdlIGRvIE5PVCByZXNldCBuZXh0VW5pcXVlUmVmIGhlcmUuIFJlc2V0dGluZyBpdCB3b3VsZCBjYXVzZSByZWYgY29sbGlzaW9ucwoJLy8gd2l0aCBjb21wb25lbnRzIGxpa2UgcHJvbWV0aGV1cy5zY3JhcGUgd2hpY2ggd2lsbCBrZWVwIHJlLXNlbmRpbmcgdGhlIHNhbWUgY2FjaGVkIHJlZnMuCgkvLyBXZSBjb250aW51ZSBpbmNyZW1lbnRpbmcgdG8gZW5zdXJlIGFsbCByZWZzIHJlbWFpbiB1bmlxdWUgYWNyb3NzIHRoZSBsaWZldGltZSBvZiB0aGUgcHJvY2Vzcy4KCgkvLyBSZXNldCBtZXRyaWNzCglzLmFjdGl2ZU1hcHBpbmdzLlNldCgwKQoJcy50cmFja2VkUmVmcy5TZXQoMCkKCgkvLyBSZXNldCBjaGFubmVscyBhbmQgZmxhZ3MKCXMuc3RvcENsZWFudXAgPSBtYWtlKGNoYW4gc3RydWN0e30pCglzLmNsZWFudXBTdG9wcGVkID0gbWFrZShjaGFuIHN0cnVjdHt9KQoJcy5zdGFydFJlZkNsZWFudXAgPSBzeW5jLk9uY2V7fQoJcy5jbGVhbnVwU3RhcnRlZC5TdG9yZShmYWxzZSkKfQo="
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/validator",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy [github.com/grafana/alloy.test]",
    "github.com/grafana/alloy.test",
    "github.com/grafana/alloy/internal/alloycli [github.com/grafana/alloy/internal/alloycli.test]",
    "github.com/grafana/alloy/internal/alloycli.test",
    "github.com/grafana/alloy/internal/component/all [github.com/grafana/alloy/internal/component/all.test]",
    "github.com/grafana/alloy/internal/component/all.test",
    "github.com/grafana/alloy/internal/component/metadata [github.com/grafana/alloy/internal/component/metadata.test]",
    "github.com/grafana/alloy/internal/component/metadata.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test [github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test [github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus.test",
    "github.com/grafana/alloy/internal/component/prometheus [github.com/grafana/alloy/internal/component/prometheus.test]",
    "github.com/grafana/alloy/internal/component/prometheus_test [github.com/grafana/alloy/internal/component/prometheus.test]",
    "github.com/grafana/alloy/internal/component/prometheus.test",
    "github.com/grafana/alloy/internal/component/prometheus/enrich [github.com/grafana/alloy/internal/component/prometheus/enrich.test]",
    "github.com/grafana/alloy/internal/component/prometheus/enrich.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test [github.com/grafana/alloy/internal/component/prometheus/exporter/tests.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator [github.com/grafana/alloy/internal/component/prometheus/operator.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common [github.com/grafana/alloy/internal/component/prometheus/operator/common.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen [github.com/grafana/alloy/internal/component/prometheus/operator/configgen.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen.test",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http [github.com/grafana/alloy/internal/component/prometheus/receive_http.test]",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http.test",
    "github.com/grafana/alloy/internal/component/prometheus/relabel [github.com/grafana/alloy/internal/component/prometheus/relabel.test]",
    "github.com/grafana/alloy/internal/component/prometheus/relabel.test",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite [github.com/grafana/alloy/internal/component/prometheus/remotewrite.test]",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test [github.com/grafana/alloy/internal/component/prometheus/remotewrite.test]",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite.test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape [github.com/grafana/alloy/internal/component/prometheus/scrape.test]",
    "github.com/grafana/alloy/internal/component/prometheus/scrape.test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape [github.com/grafana/alloy/internal/component/pyroscope/scrape.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape.test",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test [github.com/grafana/alloy/internal/converter/internal/otelcolconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test [github.com/grafana/alloy/internal/converter/internal/prometheusconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test [github.com/grafana/alloy/internal/converter/internal/promtailconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test [github.com/grafana/alloy/internal/converter/internal/staticconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build [github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build.test]",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build.test",
    "github.com/grafana/alloy/internal/service/cluster [github.com/grafana/alloy/internal/service/cluster.test]",
    "github.com/grafana/alloy/internal/service/cluster_test [github.com/grafana/alloy/internal/service/cluster.test]",
    "github.com/grafana/alloy/internal/service/cluster.test",
    "github.com/grafana/alloy/internal/tools/docs_generator_test [github.com/grafana/alloy/internal/tools/docs_generator.test]",
    "github.com/grafana/alloy/internal/tools/docs_generator.test",
    "github.com/grafana/alloy/internal/validator [github.com/grafana/alloy/internal/validator.test]",
    "github.com/grafana/alloy/internal/validator.test"
  ],
  "Packages": [
    {
      "ID": "archive/tar",
      "Name": "tar",
      "PkgPath": "archive/tar",
      "GoFiles": [
        "/usr/lib/go/src/archive/tar/common.go",
        "/usr/lib/go/src/archive/tar/format.go",
        "/usr/lib/go/src/archive/tar/reader.go",
        "/usr/lib/go/src/archive/tar/stat_actime1.go",
        "/usr/lib/go/src/archive/tar/stat_unix.go",
        "/usr/lib/go/src/archive/tar/strconv.go",
        "/usr/lib/go/src/archive/tar/writer.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/archive/tar/common.go",
        "/usr/lib/go/src/archive/tar/format.go",
        "/usr/lib/go/src/archive/tar/reader.go",
        "/usr/lib/go/src/archive/tar/stat_actime1.go",
        "/usr/lib/go/src/archive/tar/stat_unix.go",
        "/usr/lib/go/src/archive/tar/strconv.go",
        "/usr/lib/go/src/archive/tar/writer.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/archive/tar/stat_actime2.go"],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "fmt": "fmt",
        "internal/godebug": "internal/godebug",
        "io": "io",
        "io/fs": "io/fs",
        "maps": "maps",
        "math": "math",
        "os/user": "os/user",
        "path": "path",
        "path/filepath": "path/filepath",
        "reflect": "reflect",
        "runtime": "runtime",
        "slices": "slices",
        "strconv": "strconv",
        "strings": "strings",
        "sync": "sync",
        "syscall": "syscall",
        "time": "time"
      }
    },
    {
      "ID": "archive/zip",
      "Name": "zip",
      "PkgPath": "archive/zip",
      "GoFiles": [
        "/usr/lib/go/src/archive/zip/reader.go",
        "/usr/lib/go/src/archive/zip/register.go",
        "/usr/lib/go/src/archive/zip/struct.go",
        "/usr/lib/go/src/archive/zip/writer.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/archive/zip/reader.go",
        "/usr/lib/go/src/archive/zip/register.go",
        "/usr/lib/go/src/archive/zip/struct.go",
        "/usr/lib/go/src/archive/zip/writer.go"
      ],
      "Imports": {
        "bufio": "bufio",
        "compress/flate": "compress/flate",
        "encoding/binary": "encoding/binary",
        "errors": "errors",
        "fmt": "fmt",
        "hash": "hash",
        "hash/crc32": "hash/crc32",
        "internal/godebug": "internal/godebug",
        "io": "io",
        "io/fs": "io/fs",
        "os": "os",
        "path": "path",
        "path/filepath": "path/filepath",
        "slices": "slices",
        "strings": "strings",
        "sync": "sync",
        "time": "time",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cloud.google.com/go/auth",
      "Name": "auth",
      "PkgPath": "cloud.google.com/go/auth",
      "GoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/auth.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/threelegged.go"
      ],
      "CompiledGoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/auth.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/threelegged.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "cloud.google.com/go/auth/internal": "cloud.google.com/go/auth/internal",
        "cloud.google.com/go/auth/internal/jwt": "cloud.google.com/go/auth/internal/jwt",
        "context": "context",
        "encoding/json": "encoding/json",
        "errors": "errors",
        "fmt": "fmt",
        "github.com/googleapis/gax-go/v2/internallog": "github.com/googleapis/gax-go/v2/internallog",
        "log/slog": "log/slog",
        "mime": "mime",
        "net/http": "net/http",
        "net/url": "net/url",
        "strconv": "strconv",
        "strings": "strings",
        "sync": "sync",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 4797 to 5 entries.

#### drv #11

Trace meta: spanId=22, ts=1770837177333, ts_iso=2026-02-11T19:12:57.333000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "file=/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkic3luYyIKCSJzeW5jL2F0b21pYyIKCSJ0aW1lIgoKCSJnaXRodWIuY29tL2hhc2hpY29ycC9nby1tdWx0aWVycm9yIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9jbGllbnRfZ29sYW5nL3Byb21ldGhldXMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvZXhlbXBsYXIiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvaGlzdG9ncmFtIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2xhYmVscyIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9tZXRhZGF0YSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9zdG9yYWdlIgopCgp0eXBlIE1hcHBpbmdTdG9yZSBpbnRlcmZhY2UgewoJR2V0TWFwcGluZyh1bmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgW11zdG9yYWdlLlNlcmllc1JlZgoJQ3JlYXRlTWFwcGluZyhyZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgc3RvcmFnZS5TZXJpZXNSZWYKCVVwZGF0ZU1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCByZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykKCVRyYWNrQXBwZW5kZWRTZXJpZXModHMgaW50NjQsIGNlbGwgKkNlbGwpCglHZXRDZWxsRm9yQXBwZW5kZWRTZXJpZXMoKSAqQ2VsbAp9Cgp0eXBlIHNlcmllc1JlZk1hcHBpbmcgc3RydWN0IHsKCXN0YXJ0ICAgIHRpbWUuVGltZQoJY2hpbGRyZW4gW11zdG9yYWdlLkFwcGVuZGVyCglzdG9yZSAgICBNYXBwaW5nU3RvcmUKCgl1bmlxdWVSZWZDZWxsICpDZWxsCgoJLy8gY2hpbGRSZWZzIGlzIHJldXNlZCBmb3IgZWFjaCBhcHBlbmQgY2FsbCB0byBhdm9pZCBhbGxvY2F0aW9ucy4gVGhpcyBpcyBzYWZlIGJlY2F1c2Ugc3RvcmFnZS5BcHBlbmRlciBzaG91bGQgbmV2ZXIKCS8vIGhhdmUgY29uY3VycmVudCBjYWxscyB0byBBcHBlbmQgbWV0aG9kcy4KCWNoaWxkUmVmcyAgICAgICAgW11zdG9yYWdlLlNlcmllc1JlZgoJd3JpdGVMYXRlbmN5ICAgICBwcm9tZXRoZXVzLkhpc3RvZ3JhbQoJc2FtcGxlc0ZvcndhcmRlZCBwcm9tZXRoZXVzLkNvdW50ZXIKfQoKZnVuYyBOZXdTZXJpZXNSZWZNYXBwaW5nKGNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlciwgc3RvcmUgTWFwcGluZ1N0b3JlLCB3cml0ZUxhdGVuY3kgcHJvbWV0aGV1cy5IaXN0b2dyYW0sIHNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyKSBzdG9yYWdlLkFwcGVuZGVyIHsKCXVuaXF1ZVJlZkNlbGwgOj0gc3RvcmUuR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkKCglyZXR1cm4gJnNlcmllc1JlZk1hcHBpbmd7CgkJY2hpbGRyZW46ICAgICAgICAgY2hpbGRyZW4sCgkJc3RvcmU6ICAgICAgICAgICAgc3RvcmUsCgkJd3JpdGVMYXRlbmN5OiAgICAgd3JpdGVMYXRlbmN5LAoJCXNhbXBsZXNGb3J3YXJkZWQ6IHNhbXBsZXNGb3J3YXJkZWQsCgoJCXVuaXF1ZVJlZkNlbGw6IHVuaXF1ZVJlZkNlbGwsCgkJY2hpbGRSZWZzOiAgICAgbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAwLCBsZW4oY2hpbGRyZW4pKSwKCX0KfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgU2V0T3B0aW9ucyhvcHRzICpzdG9yYWdlLkFwcGVuZE9wdGlvbnMpIHsKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWMuU2V0T3B0aW9ucyhvcHRzKQoJfQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBDb21taXQoKSBlcnJvciB7CglzLnN0b3JlLlRyYWNrQXBwZW5kZWRTZXJpZXModGltZS5Ob3coKS5Vbml4KCksIHMudW5pcXVlUmVmQ2VsbCkKCgl2YXIgbXVsdGlFcnIgZXJyb3IKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWVyciA6PSBjLkNvbW1pdCgpCgkJaWYgZXJyICE9IG5pbCB7CgkJCW11bHRpRXJyID0gbXVsdGllcnJvci5BcHBlbmQobXVsdGlFcnIsIGVycikKCQl9Cgl9CglyZXR1cm4gbXVsdGlFcnIKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgUm9sbGJhY2soKSBlcnJvciB7CgkvLyBXZSBzdGlsbCB0cmFjayByb2xsZWQgYmFjayBzZXJpZXMgc28gd2UgY2FuIHByb3Blcmx5CgkvLyBjbGVhbiB1cCBhbnkgc2VyaWVzIHRoYXQgd2FzIGFwcGVuZGVkCglzLnN0b3JlLlRyYWNrQXBwZW5kZWRTZXJpZXModGltZS5Ob3coKS5Vbml4KCksIHMudW5pcXVlUmVmQ2VsbCkKCgl2YXIgbXVsdGlFcnIgZXJyb3IKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWVyciA6PSBjLlJvbGxiYWNrKCkKCQlpZiBlcnIgIT0gbmlsIHsKCQkJbXVsdGlFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChtdWx0aUVyciwgZXJyKQoJCX0KCX0KCXJldHVybiBtdWx0aUVycgp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSByZWNvcmRMYXRlbmN5KCkgewoJaWYgcy5zdGFydC5Jc1plcm8oKSB7CgkJcmV0dXJuCgl9CglkdXJhdGlvbiA6PSB0aW1lLlNpbmNlKHMuc3RhcnQpCglzLndyaXRlTGF0ZW5jeS5PYnNlcnZlKGR1cmF0aW9uLlNlY29uZHMoKSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgcmVzZXRGaWVsZHMoKSB7CgkvLyBSZXNldCBjaGlsZFJlZnMgc2xpY2UgbGVuZ3RoIHRvIDAgZm9yIHJldXNlCglzLmNoaWxkUmVmcyA9IHMuY2hpbGRSZWZzWzowXQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmQocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQgaW50NjQsIHYgZmxvYXQ2NCkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJbmV3UmVmLCBlcnIgOj0gYXBwZW5kZXIuQXBwZW5kKHJlZiwgbCwgdCwgdikKCQlpZiBlcnIgPT0gbmlsIHsKCQkJcy5zYW1wbGVzRm9yd2FyZGVkLkluYygpCgkJfQoJCXJldHVybiBuZXdSZWYsIGVycgoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kRXhlbXBsYXIocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIGUgZXhlbXBsYXIuRXhlbXBsYXIpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRFeGVtcGxhcihyZWYsIGwsIGUpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRIaXN0b2dyYW0ocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQgaW50NjQsIGggKmhpc3RvZ3JhbS5IaXN0b2dyYW0sIGZoICpoaXN0b2dyYW0uRmxvYXRIaXN0b2dyYW0pIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRIaXN0b2dyYW0ocmVmLCBsLCB0LCBoLCBmaCkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEhpc3RvZ3JhbUNUWmVyb1NhbXBsZShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCwgY3QgaW50NjQsIGggKmhpc3RvZ3JhbS5IaXN0b2dyYW0sIGZoICpoaXN0b2dyYW0uRmxvYXRIaXN0b2dyYW0pIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRIaXN0b2dyYW1DVFplcm9TYW1wbGUocmVmLCBsLCB0LCBjdCwgaCwgZmgpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBVcGRhdGVNZXRhZGF0YShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgbSBtZXRhZGF0YS5NZXRhZGF0YSkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLlVwZGF0ZU1ldGFkYXRhKHJlZiwgbCwgbSkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZENUWmVyb1NhbXBsZShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCwgY3QgaW50NjQpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRDVFplcm9TYW1wbGUocmVmLCBsLCB0LCBjdCkKCX0pCn0KCnR5cGUgYXBwZW5kZXJGdW5jIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgYXBwZW5kVG9DaGlsZHJlbihyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscywgYWYgYXBwZW5kZXJGdW5jKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglkZWZlciBzLnJlc2V0RmllbGRzKCkKCglpZiBzLnN0YXJ0LklzWmVybygpIHsKCQlzLnN0YXJ0ID0gdGltZS5Ob3coKQoJfQoKCS8vIENoZWNrIGlmIHRoZSBpbmNvbWluZyByZWYgaGFzIHJlZiBtYXBwaW5ncwoJZXhpc3RpbmdDaGlsZFJlZnMgOj0gcy5zdG9yZS5HZXRNYXBwaW5nKHJlZiwgbGJscykKCgl2YXIgYXBwZW5kRXJyIGVycm9yCgoJLy8gU2FuaXR5IGNoZWNrOiBpZiB3ZSBoYXZlIGV4aXN0aW5nIGNoaWxkIHJlZnMsIHRoZXkgbXVzdCBtYXRjaCB0aGUgbnVtYmVyIG9mIGNoaWxkcmVuCglpZiBleGlzdGluZ0NoaWxkUmVmcyAhPSBuaWwgJiYgbGVuKGV4aXN0aW5nQ2hpbGRSZWZzKSA9PSBsZW4ocy5jaGlsZHJlbikgewoJCXMudW5pcXVlUmVmQ2VsbC5SZWZzID0gYXBwZW5kKHMudW5pcXVlUmVmQ2VsbC5SZWZzLCByZWYpCgoJCXJlZlVwZGF0ZVJlcXVpcmVkIDo9IGZhbHNlCgkJZm9yIGNoaWxkSW5kZXgsIGNoaWxkUmVmIDo9IHJhbmdlIGV4aXN0aW5nQ2hpbGRSZWZzIHsKCQkJbmV3Q2hpbGRSZWYsIGVyciA6PSBhZihzLmNoaWxkcmVuW2NoaWxkSW5kZXhdLCBjaGlsZFJlZikKCQkJaWYgZXJyICE9IG5pbCB7CgkJCQlhcHBlbmRFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChhcHBlbmRFcnIsIGVycikKCQkJfQoKCQkJaWYgbmV3Q2hpbGRSZWYgIT0gY2hpbGRSZWYgewoJCQkJLy8gQ2hpbGQgcmVmIGNoYW5nZWQsIG5lZWQgdG8gdXBkYXRlIG1hcHBpbmcKCQkJCWV4aXN0aW5nQ2hpbGRSZWZzW2NoaWxkSW5kZXhdID0gbmV3Q2hpbGRSZWYKCQkJCXJlZlVwZGF0ZVJlcXVpcmVkID0gdHJ1ZQoJCQl9CgkJfQoKCQlpZiBhcHBlbmRFcnIgIT0gbmlsIHsKCQkJcmV0dXJuIDAsIGFwcGVuZEVycgoJCX0KCgkJaWYgcmVmVXBkYXRlUmVxdWlyZWQgewoJCQlzLnN0b3JlLlVwZGF0ZU1hcHBpbmcocmVmLCBleGlzdGluZ0NoaWxkUmVmcywgbGJscykKCQl9CgoJCXJldHVybiByZWYsIG5pbAoJfQoKCS8vIE5vIGV4aXN0aW5nIG1hcHBpbmcsIHByb2NlZWQgd2l0aCBub3JtYWwgYXBwZW5kIHRvIGFsbCBjaGlsZHJlbgoJdmFyIGZpcnN0Tm9uWmVyb1JlZiBzdG9yYWdlLlNlcmllc1JlZgoJdmFyIG5vblplcm9Db3VudCBpbnQKCgkvLyBOb3RlOiB0aGVyZSdzIGFub3RoZXIgb3B0aW1pemF0aW9uIHdoZXJlIHdlIGNvdWxkIHVzZSB0aGUgcmV0dXJuZWQgcmVmIGlmIGFsbCB0aGUgbm9uIHplcm8gcmVmcwoJLy8gIGFyZSB0aGUgc2FtZSB2YWx1ZS4gVGhpcyBpc24ndCBzYWZlIGFzIHdlIHdpbGwgbWl4IGRvd25zdHJlYW0gcmVmcyB3aXRoIHVuaXF1ZSByZWZzIHdoaWNoIGNvdWxkCgkvLyAgY29sbGlkZS4gV2UgY291bGQgc3RhcnQgYXQgbWF4IHVuaXQ2NCBmb3IgdW5pcXVlIHJlZnMgYW5kIGdvIGJhY2t3YXJkcyBsZXNzZW5pbmcgdGhlIGNoYW5jZSBvZgoJLy8gCWNvbGxpc2lvbnMgYnV0IGl0J3MgcmF0aGVyIGRhbmdlcm91cyBmb3IgYW4gdW5saWtlbHkgZWRnZSBjYXNlLiBJZiB0d28gY29tcG9uZW50cyBhcmUgcmV0dXJuaW5nCgkvLyAJdGhlIHNhbWUgcmVmIGl0J3MgdHdvIHJlbW90ZV93cml0ZSBjb21wb25lbnRzIHdoaWNoIHNob3VsZCBwcm9iYWJseSBiZSBtZXJnZWQgaW4gdG8gb25lLgoJZm9yIF8sIGNoaWxkIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWNoaWxkUmVmLCBlcnIgOj0gYWYoY2hpbGQsIHJlZikKCQlpZiBlcnIgIT0gbmlsIHsKCQkJYXBwZW5kRXJyID0gbXVsdGllcnJvci5BcHBlbmQoYXBwZW5kRXJyLCBlcnIpCgoJCQkvLyBUT0RPIHNob3VsZCBJIG1ha2UgdGhlIGNoaWxkUmVmIHplcm8gaGVyZT8KCQl9CgoJCXMuY2hpbGRSZWZzID0gYXBwZW5kKHMuY2hpbGRSZWZzLCBjaGlsZFJlZikKCQlpZiBjaGlsZFJlZiAhPSAwIHsKCQkJaWYgZmlyc3ROb25aZXJvUmVmID09IDAgewoJCQkJZmlyc3ROb25aZXJvUmVmID0gY2hpbGRSZWYKCQkJfQoJCQlub25aZXJvQ291bnQrKwoJCX0KCX0KCglpZiBhcHBlbmRFcnIgIT0gbmlsIHsKCQlyZXR1cm4gMCwgYXBwZW5kRXJyCgl9CgoJaWYgbm9uWmVyb0NvdW50ID09IDAgewoJCS8vIEFsbCBjaGlsZHJlbiByZXR1cm5lZCByZWYgMCwgc28gcmV0dXJuIHRoZSBpbnB1dCByZWYKCQlyZXR1cm4gcmVmLCBuaWwKCX0KCgkvLyBPbmx5IG9uZSBjaGlsZCByZXR1cm5lZCBhIG5vbi16ZXJvIHJlZiwgdXNlIHRoYXQKCWlmIG5vblplcm9Db3VudCA9PSAxIHsKCQlyZXR1cm4gZmlyc3ROb25aZXJvUmVmLCBuaWwKCX0KCgkvLyBXZSBnb3QgZGlmZmVyZW50IHJlZnMgYmFjayBhbmQgbmVlZCB0byBjcmVhdGUgYSBuZXcgbWFwcGluZwoJdW5pcXVlUmVmIDo9IHMuc3RvcmUuQ3JlYXRlTWFwcGluZyhzLmNoaWxkUmVmcywgbGJscykKCXMudW5pcXVlUmVmQ2VsbC5SZWZzID0gYXBwZW5kKHMudW5pcXVlUmVmQ2VsbC5SZWZzLCB1bmlxdWVSZWYpCglyZXR1cm4gdW5pcXVlUmVmLCBuaWwKfQoKdHlwZSB1bmlxUmVmQ2hpbGRyZW4gc3RydWN0IHsKCWNoaWxkUmVmcyAqW11zdG9yYWdlLlNlcmllc1JlZgoJbGFiZWxIYXNoIHVpbnQ2NAp9Cgp0eXBlIFNlcmllc1JlZk1hcHBpbmdTdG9yZSBzdHJ1Y3QgewoJLy8gcmVmTWFwcGluZ011IHByb3RlY3RzIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzLCBsYWJlbEhhc2hUb1VuaXF1ZVJlZiBhbmQgbmV4dFVuaXF1ZVJlZgoJcmVmTWFwcGluZ011IHN5bmMuUldNdXRleAoJLy8gdW5pcXVlUmVmVG9DaGlsZFJlZnMgbWFwcyB0aGUgdW5pcXVlIHJlZiB0byB0aGUgZXhwZWN0ZWQgY2hpbGQgcmVmIGluIG9yZGVyCgl1bmlxdWVSZWZUb0NoaWxkUmVmcyBtYXBbc3RvcmFnZS5TZXJpZXNSZWZddW5pcVJlZkNoaWxkcmVuCgkvLyBsYWJlbEhhc2hUb1VuaXF1ZVJlZiBtYXBzIHRoZSBsYWJlbCBoYXNoIHRvIHVuaXF1ZSByZWYuCglsYWJlbEhhc2hUb1VuaXF1ZVJlZiBtYXBbdWludDY0XXN0b3JhZ2UuU2VyaWVzUmVmCgoJLy8gbmV4dFVuaXF1ZVJlZiBpcyB0aGUgbmV4dCByZWYgSUQgd2Ugd2lsbCBoYW5kIG91dAoJbmV4dFVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZgoKCS8vIHRpbWVzdGFtcFRyYWNraW5nTXUgcHJvdGVjdHMgdW5pcXVlUmVmVGltZXN0YW1wcyBhbmQgY2VsbFBvb2wKCXRpbWVzdGFtcFRyYWNraW5nTXUgc3luYy5NdXRleAoJLy8gdW5pcXVlUmVmVGltZXN0YW1wcyBtYXBzIHVuaXF1ZSByZWZzIHRvIHRoZWlyIGxhc3QgYXBwZW5kIHRpbWVzdGFtcAoJdW5pcXVlUmVmVGltZXN0YW1wcyBtYXBbc3RvcmFnZS5TZXJpZXNSZWZdaW50NjQKCS8vIGNlbGxQb29sIGlzIHVzZWQgdG8gcG9vbCBzbGljZXMgb2YgU2VyaWVzUmVmcyB1c2VkIGZvciB0cmFja2luZyB1bmlxdWUgcmVmcyBpbiBUcmFja0FwcGVuZGVkU2VyaWVzLgoJY2VsbFBvb2wgc3luYy5Qb29sCgoJLy8gQ2xlYW51cCBnb3JvdXRpbmUgY29vcmRpbmF0aW9uIChubyBsb2NrIHJlcXVpcmVkKQoJc3RhcnRSZWZDbGVhbnVwIHN5bmMuT25jZQoJY2xlYW51cFN0YXJ0ZWQgIGF0b21pYy5Cb29sCglzdG9wQ2xlYW51cCAgICAgY2hhbiBzdHJ1Y3R7fQoJY2xlYW51cFN0b3BwZWQgIGNoYW4gc3RydWN0e30KCgkvLyBNZXRyaWNzIChzYWZlIGZvciBjb25jdXJyZW50IGFjY2Vzcywgbm8gbG9jayByZXF1aXJlZCkKCWFjdGl2ZU1hcHBpbmdzICBwcm9tZXRoZXVzLkdhdWdlCgl0cmFja2VkUmVmcyAgICAgcHJvbWV0aGV1cy5HYXVnZQoJcmVmc0NsZWFuZWQgICAgIHByb21ldGhldXMuQ291bnRlcgoJdW5pcXVlUmVmc1RvdGFsIHByb21ldGhldXMuQ291bnRlcgp9CgpmdW5jIE5ld1Nlcmllc1JlZk1hcHBpbmdTdG9yZShyZWcgcHJvbWV0aGV1cy5SZWdpc3RlcmVyKSAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlIHsKCWFjdGl2ZU1hcHBpbmdzIDo9IHByb21ldGhldXMuTmV3R2F1Z2UocHJvbWV0aGV1cy5HYXVnZU9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX21hcHBpbmdzX3RvdGFsIiwKCQlIZWxwOiAiTnVtYmVyIG9mIGFjdGl2ZSB1bmlxdWUgcmVmIG1hcHBpbmdzIGluIHRoZSBzdG9yZS4iLAoJfSkKCXRyYWNrZWRSZWZzIDo9IHByb21ldGhldXMuTmV3R2F1Z2UocHJvbWV0aGV1cy5HYXVnZU9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3RyYWNrZWRfcmVmc190b3RhbCIsCgkJSGVscDogIk51bWJlciBvZiByZWZzIGJlaW5nIHRyYWNrZWQgZm9yIHRpbWVzdGFtcC1iYXNlZCBjbGVhbnVwLiIsCgl9KQoJcmVmc0NsZWFuZWQgOj0gcHJvbWV0aGV1cy5OZXdDb3VudGVyKHByb21ldGhldXMuQ291bnRlck9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3JlZnNfY2xlYW5lZF90b3RhbCIsCgkJSGVscDogIlRvdGFsIG51bWJlciBvZiBzdGFsZSByZWZzIGNsZWFuZWQgdXAgb3ZlciB0aW1lLiIsCgl9KQoJdW5pcXVlUmVmc1RvdGFsIDo9IHByb21ldGhldXMuTmV3Q291bnRlcihwcm9tZXRoZXVzLkNvdW50ZXJPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV91bmlxdWVfcmVmc19jcmVhdGVkX3RvdGFsIiwKCQlIZWxwOiAiVG90YWwgbnVtYmVyIG9mIHVuaXF1ZSByZWZzIGNyZWF0ZWQuIiwKCX0pCgoJaWYgcmVnICE9IG5pbCB7CgkJcmVnLlJlZ2lzdGVyKGFjdGl2ZU1hcHBpbmdzKQoJCXJlZy5SZWdpc3Rlcih0cmFja2VkUmVmcykKCQlyZWcuUmVnaXN0ZXIocmVmc0NsZWFuZWQpCgkJcmVnLlJlZ2lzdGVyKHVuaXF1ZVJlZnNUb3RhbCkKCX0KCglyZXR1cm4gJlNlcmllc1JlZk1hcHBpbmdTdG9yZXsKCQl1bmlxdWVSZWZUb0NoaWxkUmVmczogbWFrZShtYXBbc3RvcmFnZS5TZXJpZXNSZWZddW5pcVJlZkNoaWxkcmVuKSwKCQluZXh0VW5pcXVlUmVmOiAgICAgICAgMSwKCQl1bmlxdWVSZWZUaW1lc3RhbXBzOiAgbWFrZShtYXBbc3RvcmFnZS5TZXJpZXNSZWZdaW50NjQpLAoJCWxhYmVsSGFzaFRvVW5pcXVlUmVmOiBtYWtlKG1hcFt1aW50NjRdc3RvcmFnZS5TZXJpZXNSZWYpLAoJCWNlbGxQb29sOiBzeW5jLlBvb2x7CgkJCU5ldzogZnVuYygpIGFueSB7CgkJCQlyZXR1cm4gJkNlbGx7UmVmczogbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAxMDApfQoJCQl9LAoJCX0sCgkJc3RvcENsZWFudXA6ICAgICBtYWtlKGNoYW4gc3RydWN0e30pLAoJCWNsZWFudXBTdG9wcGVkOiAgbWFrZShjaGFuIHN0cnVjdHt9KSwKCQlhY3RpdmVNYXBwaW5nczogIGFjdGl2ZU1hcHBpbmdzLAoJCXRyYWNrZWRSZWZzOiAgICAgdHJhY2tlZFJlZnMsCgkJcmVmc0NsZWFuZWQ6ICAgICByZWZzQ2xlYW5lZCwKCQl1bmlxdWVSZWZzVG90YWw6IHVuaXF1ZVJlZnNUb3RhbCwKCX0KfQoKdHlwZSBDZWxsIHN0cnVjdCB7CglSZWZzIFtdc3RvcmFnZS5TZXJpZXNSZWYKfQoKLy8gR2V0TWFwcGluZyByZXR1cm5zIGV4aXN0aW5nIGNoaWxkIHJlZiByZXN1bHRzIGZvciB0aGUgZ2l2ZW4gdW5pcXVlIHJlZiBpZiBvbmUgZXhpc3RzLgovLwovLyBJZiB0aGUgcGFzc2VkIHVuaXF1ZVJlZiBpcyB6ZXJvLCB0aGUgbWV0aG9kIHdpbGwgYXR0ZW1wdCB0byBmaW5kIGEgbWFwcGluZyB1c2luZyBwYXNzZWQgbGFiZWxzLgovLyBSZXR1cm5zIG5pbCBpZiBubyBtYXBwaW5nIGV4aXN0cy4KLy8KLy8gVGhlIHJldHVybmVkIHNsaWNlIG1heSBiZSBtb2RpZmllZCBieSB0aGUgY2FsbGVyLCBidXQgVXBkYXRlTWFwcGluZyBtdXN0IGJlIGNhbGxlZAovLyBhZnRlcndhcmRzIHRvIHBlcnNpc3QgY2hhbmdlcy4gTm90ZSB0aGF0IGNvbmN1cnJlbnQgYXBwZW5kZXJzIG1heSByYWNlIHRvIHVwZGF0ZSB0aGUKLy8gc2FtZSBtYXBwaW5nIHdpdGggZGlmZmVyZW50IHZhbHVlcywgd2hpY2ggaXMgc2FmZSBiZWNhdXNlIHN0YWxlIG1hcHBpbmdzIGFyZSBzZWxmLWNvcnJlY3RpbmcgLQovLyB1c2luZyBhIHN0YWxlIHJlZiB3aWxsIGNhdXNlIHRoZSBjaGlsZCBhcHBlbmRlciB0byByZXR1cm4gYSBuZXcgcmVmIG9uIHRoZSBuZXh0IGFwcGVuZC4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBHZXRNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSBbXXN0b3JhZ2UuU2VyaWVzUmVmIHsKCXMucmVmTWFwcGluZ011LlJMb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlJVbmxvY2soKQoKCWlmIHVuaXF1ZVJlZiA9PSAwIHsKCQkvLyBTb21lIGNvbnN1bWVycyBkb24ndCBtZW1vIHRoZSBnbG9iYWwgcmVmLiBUcnkgdG8gbG9va3VwIGEgcmVmIGJ5IGxhYmVsIGhhc2guCgkJbGFiZWxIYXNoIDo9IGxibHMuSGFzaCgpCgkJZ290UmVmLCBvayA6PSBzLmxhYmVsSGFzaFRvVW5pcXVlUmVmW2xhYmVsSGFzaF0KCQlpZiAhb2sgewoJCQlyZXR1cm4gbmlsCgkJfQoKCQl1bmlxdWVSZWYgPSBnb3RSZWYKCX0KCglpZiBtYXBwaW5nLCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl07IG9rIHsKCQlyZXR1cm4gKm1hcHBpbmcuY2hpbGRSZWZzCgl9CglyZXR1cm4gbmlsCn0KCi8vIENyZWF0ZU1hcHBpbmcgY3JlYXRlcyBhIG5ldyB1bmlxdWUgcmVmIG1hcHBpbmcgZm9yIHRoZSBnaXZlbiBjaGlsZCByZWYgcmVzdWx0cy4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBDcmVhdGVNYXBwaW5nKHJlZlJlc3VsdHMgW11zdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSBzdG9yYWdlLlNlcmllc1JlZiB7CgkvLyBTdGFydCBjbGVhbnVwIGdvcm91dGluZSBvbiBmaXJzdCBtYXBwaW5nCglzLnN0YXJ0UmVmQ2xlYW51cC5EbyhmdW5jKCkgewoJCXMuY2xlYW51cFN0YXJ0ZWQuU3RvcmUodHJ1ZSkKCQlnbyBzLmNsZWFudXBTdGFsZVJlZnMoKQoJfSkKCgkvLyBTdG9yZSBhIGNvcHkgb2YgdGhlIGNoaWxkIHJlZiByZXN1bHRzIGRpcmVjdGx5CgljaGlsZFJlZlNsaWNlIDo9IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgbGVuKHJlZlJlc3VsdHMpKQoJY29weShjaGlsZFJlZlNsaWNlLCByZWZSZXN1bHRzKQoKCS8vIEhhc2ggbGFiZWxzIHRvIGZvciB0aGUgZmFsbGJhY2sgbG9va3VwIHRhYmxlCglsYWJlbEhhc2ggOj0gbGJscy5IYXNoKCkKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJLy8gQ3JlYXRlIGEgbmV3IHVuaXF1ZSByZWYKCXVuaXF1ZVJlZiA6PSBzLm5leHRVbmlxdWVSZWYKCXMubmV4dFVuaXF1ZVJlZisrCgoJcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltsYWJlbEhhc2hdID0gdW5pcXVlUmVmCglzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0gPSB1bmlxUmVmQ2hpbGRyZW57CgkJY2hpbGRSZWZzOiAmY2hpbGRSZWZTbGljZSwKCQlsYWJlbEhhc2g6IGxhYmVsSGFzaCwKCX0KCglzLmFjdGl2ZU1hcHBpbmdzLkluYygpCglzLnVuaXF1ZVJlZnNUb3RhbC5JbmMoKQoKCXJldHVybiB1bmlxdWVSZWYKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHsKCWlmIHVuaXF1ZVJlZiA9PSAwIHsKCQlyZXR1cm4KCX0KCgljaGlsZFJlZlNsaWNlIDo9IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgbGVuKHJlZlJlc3VsdHMpKQoJY29weShjaGlsZFJlZlNsaWNlLCByZWZSZXN1bHRzKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgkvLyBFbnN1cmUgdGhhdCBsYWJlbCBoYXNoIGluZGV4IGlzIHVwIHRvIGRhdGUgdG8gaGFuZGxlIHBvc3NpYmxlIGhhc2ggY29sbGlzaW9ucy4KCS8vIFRPRE86IGlzIHRoaXMgbmVjZXNzYXJ5PwoJbmV3SGFzaCA6PSBsYmxzLkhhc2goKQoJcHJldiwgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdCglpZiBvayAmJiBwcmV2LmxhYmVsSGFzaCAhPSBuZXdIYXNoIHsKCQlkZWxldGUocy5sYWJlbEhhc2hUb1VuaXF1ZVJlZiwgcHJldi5sYWJlbEhhc2gpCgkJcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltuZXdIYXNoXSA9IHVuaXF1ZVJlZgoJfQoKCXMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXSA9IHVuaXFSZWZDaGlsZHJlbnsKCQljaGlsZFJlZnM6ICZjaGlsZFJlZlNsaWNlLAoJCWxhYmVsSGFzaDogbGJscy5IYXNoKCksCgl9Cn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgVHJhY2tBcHBlbmRlZFNlcmllcyh0cyBpbnQ2NCwgY2VsbCAqQ2VsbCkgewoJcy50aW1lc3RhbXBUcmFja2luZ011LkxvY2soKQoJZGVmZXIgcy50aW1lc3RhbXBUcmFja2luZ011LlVubG9jaygpCgoJZm9yIF8sIHIgOj0gcmFuZ2UgY2VsbC5SZWZzIHsKCQlzLnVuaXF1ZVJlZlRpbWVzdGFtcHNbcl0gPSB0cwoJfQoKCXMudHJhY2tlZFJlZnMuU2V0KGZsb2F0NjQobGVuKHMudW5pcXVlUmVmVGltZXN0YW1wcykpKQoKCWNlbGwuUmVmcyA9IGNlbGwuUmVmc1s6MF0KCXMuY2VsbFBvb2wuUHV0KGNlbGwpCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwgewoJcmV0dXJuIHMuY2VsbFBvb2wuR2V0KCkuKCpDZWxsKQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIGNsZWFudXBTdGFsZVJlZnMoKSB7CglkZWZlciBjbG9zZShzLmNsZWFudXBTdG9wcGVkKQoKCXRpY2tlciA6PSB0aW1lLk5ld1RpY2tlcigxNSAqIHRpbWUuTWludXRlKQoJZGVmZXIgdGlja2VyLlN0b3AoKQoKCWZvciB7CgkJc2VsZWN0IHsKCQljYXNlIDwtdGlja2VyLkM6CgkJCWN1dG9mZlRpbWUgOj0gdGltZS5Ob3coKS5BZGQoLTE1ICogdGltZS5NaW51dGUpLlVuaXgoKQoKCQkJLy8gSG9sZCBib3RoIGxvY2tzIHRvIHByZXZlbnQgcmFjZSBjb25kaXRpb24gd2hlcmUgYSByZWYgY291bGQgYmUKCQkJLy8gYXBwZW5kZWQgYWZ0ZXIgd2UgZGVsZXRlIGl0IGZyb20gdW5pcXVlUmVmQ2VsbCBidXQgYmVmb3JlCgkJCS8vIHdlIGRlbGV0ZSBpdCBmcm9tIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzCgkJCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCQkJcy5yZWZNYXBwaW5nTXUuTG9jaygpCgoJCQlzdGFsZVJlZkNvdW50IDo9IDAKCQkJZm9yIHJlZiwgdHMgOj0gcmFuZ2Ugcy51bmlxdWVSZWZUaW1lc3RhbXBzIHsKCQkJCWlmIHRzIDwgY3V0b2ZmVGltZSB7CgkJCQkJc3RhbGVSZWZDb3VudCsrCgoJCQkJCXYsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbcmVmXQoJCQkJCWlmIG9rIHsKCQkJCQkJZGVsZXRlKHMubGFiZWxIYXNoVG9VbmlxdWVSZWYsIHYubGFiZWxIYXNoKQoJCQkJCX0KCgkJCQkJZGVsZXRlKHMudW5pcXVlUmVmVGltZXN0YW1wcywgcmVmKQoJCQkJCWRlbGV0ZShzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzLCByZWYpCgkJCQl9CgkJCX0KCgkJCS8vIFVwZGF0ZSBtZXRyaWNzCgkJCWlmIHN0YWxlUmVmQ291bnQgPiAwIHsKCQkJCXMucmVmc0NsZWFuZWQuQWRkKGZsb2F0NjQoc3RhbGVSZWZDb3VudCkpCgkJCQlzLmFjdGl2ZU1hcHBpbmdzLlN1YihmbG9hdDY0KHN0YWxlUmVmQ291bnQpKQoJCQkJcy50cmFja2VkUmVmcy5TZXQoZmxvYXQ2NChsZW4ocy51bmlxdWVSZWZUaW1lc3RhbXBzKSkpCgkJCX0KCgkJCXMucmVmTWFwcGluZ011LlVubG9jaygpCgkJCXMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCQljYXNlIDwtcy5zdG9wQ2xlYW51cDoKCQkJcmV0dXJuCgkJfQoJfQp9CgovLyBDbGVhciB3aWxsIGNsZWFyIGFsbCBpbnRlcm5hbCBtYXBwaW5ncyBhbmQgc3RvcCB0aGUgY2xlYW5lciBnb3JvdXRpbmUgaWYgaXQgaXMgcnVubmluZy4KLy8gSXQgaXMgc2FmZSB0byByZS11c2UgdGhlIHNhbWUgaW5zdGFuY2UgYWZ0ZXIgY2FsbGluZyBDbGVhci4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBDbGVhcigpIHsKCS8vIFN0b3AgdGhlIGNsZWFudXAgZ29yb3V0aW5lIGFuZCB3YWl0IGZvciBpdCB0byBiZSBzdG9wcGVkIHNvIHdlIGNhbgoJLy8gYXZvaWQgYSBwb3NzaWJsZSBkZWFkbG9jayB3aXRoIGNsZWFudXAgdGhhdCBhbHNvIGhvbGRzIGJvdGggbG9ja3MKCWlmIHMuY2xlYW51cFN0YXJ0ZWQuTG9hZCgpIHsKCQlzZWxlY3QgewoJCWNhc2UgPC1zLnN0b3BDbGVhbnVwOgoJCQkvLyBBbHJlYWR5IGNsb3NlZAoJCWRlZmF1bHQ6CgkJCWNsb3NlKHMuc3RvcENsZWFudXApCgkJCTwtcy5jbGVhbnVwU3RvcHBlZAoJCX0KCX0KCgkvLyBXZSBuZWVkIHRvIGhvbGQgYm90aCBsb2NrcyB0byBkbyB0aGlzIHNhZmVseSBhbmQgd2UgZG8gaXQgaW4gdGhlIHNhbWUgb3JkZXIgYXMKCS8vIGNsZWFudXBTdGFsZVJlZnMuIFdlIHN0b3BwZWQgYW5kIHdhaXRlZCBmb3IgdGhlIGJhY2tncm91bmQgd29ya2VyIHRoYXQgY2FsbHMgaXQKCS8vIHRvIGZpbmlzaCBidXQgc29tZSBleHRyYSBzYWZldHkgd29uJ3QgaHVydC4KCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCWRlZmVyIHMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgljbGVhcihzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzKQoJY2xlYXIocy51bmlxdWVSZWZUaW1lc3RhbXBzKQoKCS8vIHJlc2V0IHRoZSBwb29sCglzLmNlbGxQb29sID0gc3luYy5Qb29sewoJCU5ldzogZnVuYygpIGFueSB7CgkJCXJldHVybiAmQ2VsbHtSZWZzOiBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIDAsIDEwMCl9CgkJfSwKCX0KCgkvLyBOT1RFOiBXZSBkbyBOT1QgcmVzZXQgbmV4dFVuaXF1ZVJlZiBoZXJlLiBSZXNldHRpbmcgaXQgd291bGQgY2F1c2UgcmVmIGNvbGxpc2lvbnMKCS8vIHdpdGggY29tcG9uZW50cyBsaWtlIHByb21ldGhldXMuc2NyYXBlIHdoaWNoIHdpbGwga2VlcCByZS1zZW5kaW5nIHRoZSBzYW1lIGNhY2hlZCByZWZzLgoJLy8gV2UgY29udGludWUgaW5jcmVtZW50aW5nIHRvIGVuc3VyZSBhbGwgcmVmcyByZW1haW4gdW5pcXVlIGFjcm9zcyB0aGUgbGlmZXRpbWUgb2YgdGhlIHByb2Nlc3MuCgoJLy8gUmVzZXQgbWV0cmljcwoJcy5hY3RpdmVNYXBwaW5ncy5TZXQoMCkKCXMudHJhY2tlZFJlZnMuU2V0KDApCgoJLy8gUmVzZXQgY2hhbm5lbHMgYW5kIGZsYWdzCglzLnN0b3BDbGVhbnVwID0gbWFrZShjaGFuIHN0cnVjdHt9KQoJcy5jbGVhbnVwU3RvcHBlZCA9IG1ha2UoY2hhbiBzdHJ1Y3R7fSkKCXMuc3RhcnRSZWZDbGVhbnVwID0gc3luYy5PbmNle30KCXMuY2xlYW51cFN0YXJ0ZWQuU3RvcmUoZmFsc2UpCn0K"
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/grafana/alloy/internal/component/prometheus/appenders",
    "github.com/grafana/alloy/internal/component/prometheus/appenders [github.com/grafana/alloy/internal/component/prometheus/appenders.test]"
  ],
  "Packages": [
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cloud.google.com/go/auth",
      "Name": "auth",
      "PkgPath": "cloud.google.com/go/auth",
      "GoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/auth.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/threelegged.go"
      ],
      "CompiledGoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/auth.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/threelegged.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "cloud.google.com/go/auth/internal": "cloud.google.com/go/auth/internal",
        "cloud.google.com/go/auth/internal/jwt": "cloud.google.com/go/auth/internal/jwt",
        "context": "context",
        "encoding/json": "encoding/json",
        "errors": "errors",
        "fmt": "fmt",
        "github.com/googleapis/gax-go/v2/internallog": "github.com/googleapis/gax-go/v2/internallog",
        "log/slog": "log/slog",
        "mime": "mime",
        "net/http": "net/http",
        "net/url": "net/url",
        "strconv": "strconv",
        "strings": "strings",
        "sync": "sync",
        "time": "time"
      }
    },
    {
      "ID": "cloud.google.com/go/auth/credentials",
      "Name": "credentials",
      "PkgPath": "cloud.google.com/go/auth/credentials",
      "GoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/compute.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/detect.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/doc.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/filetypes.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/selfsignedjwt.go"
      ],
      "CompiledGoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/compute.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/detect.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/doc.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/filetypes.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/selfsignedjwt.go"
      ],
      "Imports": {
        "cloud.google.com/go/auth": "cloud.google.com/go/auth",
        "cloud.google.com/go/auth/credentials/internal/externalaccount": "cloud.google.com/go/auth/credentials/internal/externalaccount",
        "cloud.google.com/go/auth/credentials/internal/externalaccountuser": "cloud.google.com/go/auth/credentials/internal/externalaccountuser",
        "cloud.google.com/go/auth/credentials/internal/gdch": "cloud.google.com/go/auth/credentials/internal/gdch",
        "cloud.google.com/go/auth/credentials/internal/impersonate": "cloud.google.com/go/auth/credentials/internal/impersonate",
        "cloud.google.com/go/auth/internal": "cloud.google.com/go/auth/internal",
        "cloud.google.com/go/auth/internal/credsfile": "cloud.google.com/go/auth/internal/credsfile",
        "cloud.google.com/go/auth/internal/jwt": "cloud.google.com/go/auth/internal/jwt",
        "cloud.google.com/go/auth/internal/trustboundary": "cloud.google.com/go/auth/internal/trustboundary",
        "cloud.google.com/go/compute/metadata": "cloud.google.com/go/compute/metadata",
        "context": "context",
        "crypto": "crypto",
        "encoding/json": "encoding/json",
        "errors": "errors",
        "fmt": "fmt",
        "github.com/googleapis/gax-go/v2/internallog": "github.com/googleapis/gax-go/v2/internallog",
        "log/slog": "log/slog",
        "net/http": "net/http",
        "net/url": "net/url",
        "os": "os",
        "strings": "strings",
        "time": "time"
      }
    },
    {
      "ID": "cloud.google.com/go/auth/credentials/internal/externalaccount",
      "Name": "externalaccount",
      "PkgPath": "cloud.google.com/go/auth/credentials/internal/externalaccount",
      "GoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/aws_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/executable_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/externalaccount.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/file_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/info.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/programmatic_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/url_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/x509_provider.go"
      ],
      "CompiledGoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/aws_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/executable_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/externalaccount.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/file_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/info.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/programmatic_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/url_provider.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/credentials/internal/externalaccount/x509_provider.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "cloud.google.com/go/auth": "cloud.google.com/go/auth",
        "cloud.google.com/go/auth/credentials/internal/impersonate": "cloud.google.com/go/auth/credentials/internal/impersonate",
        "cloud.google.com/go/auth/credentials/internal/stsexchange": "cloud.google.com/go/auth/credentials/internal/stsexchange",
        "cloud.google.com/go/auth/internal": "cloud.google.com/go/auth/internal",
        "cloud.google.com/go/auth/internal/credsfile": "cloud.google.com/go/auth/internal/credsfile",
        "cloud.google.com/go/auth/internal/transport/cert": "cloud.google.com/go/auth/internal/transport/cert",
        "context": "context",
        "crypto/hmac": "crypto/hmac",
        "crypto/sha256": "crypto/sha256",
        "crypto/tls": "crypto/tls",
        "crypto/x509": "crypto/x509",
        "encoding/base64": "encoding/base64",
        "encoding/hex": "encoding/hex",
        "encoding/json": "encoding/json",
        "encoding/pem": "encoding/pem",
        "errors": "errors",
        "fmt": "fmt",
        "github.com/googleapis/gax-go/v2/internallog": "github.com/googleapis/gax-go/v2/internallog",
        "io/fs": "io/fs",
        "log/slog": "log/slog",
        "net/http": "net/http",
        "net/url": "net/url",
        "os": "os",
        "os/exec": "os/exec",
        "path": "path",
        "regexp": "regexp",
        "runtime": "runtime",
        "sort": "sort",
        "strconv": "strconv",
        "strings": "strings",
        "time": "time",
        "unicode": "unicode"
      }
    }
  ],
  "GoVersion": 0
}
```

Note: `Packages` truncated from 661 to 5 entries.

#### drv #12

Trace meta: spanId=21, ts=1770837177333, ts_iso=2026-02-11T19:12:57.333000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/appenders",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy/internal/validator"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkic3luYyIKCSJzeW5jL2F0b21pYyIKCSJ0aW1lIgoKCSJnaXRodWIuY29tL2hhc2hpY29ycC9nby1tdWx0aWVycm9yIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9jbGllbnRfZ29sYW5nL3Byb21ldGhldXMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvZXhlbXBsYXIiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvaGlzdG9ncmFtIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2xhYmVscyIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9tZXRhZGF0YSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9zdG9yYWdlIgopCgp0eXBlIE1hcHBpbmdTdG9yZSBpbnRlcmZhY2UgewoJR2V0TWFwcGluZyh1bmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgW11zdG9yYWdlLlNlcmllc1JlZgoJQ3JlYXRlTWFwcGluZyhyZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgc3RvcmFnZS5TZXJpZXNSZWYKCVVwZGF0ZU1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCByZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykKCVRyYWNrQXBwZW5kZWRTZXJpZXModHMgaW50NjQsIGNlbGwgKkNlbGwpCglHZXRDZWxsRm9yQXBwZW5kZWRTZXJpZXMoKSAqQ2VsbAp9Cgp0eXBlIHNlcmllc1JlZk1hcHBpbmcgc3RydWN0IHsKCXN0YXJ0ICAgIHRpbWUuVGltZQoJY2hpbGRyZW4gW11zdG9yYWdlLkFwcGVuZGVyCglzdG9yZSAgICBNYXBwaW5nU3RvcmUKCgl1bmlxdWVSZWZDZWxsICpDZWxsCgoJLy8gY2hpbGRSZWZzIGlzIHJldXNlZCBmb3IgZWFjaCBhcHBlbmQgY2FsbCB0byBhdm9pZCBhbGxvY2F0aW9ucy4gVGhpcyBpcyBzYWZlIGJlY2F1c2Ugc3RvcmFnZS5BcHBlbmRlciBzaG91bGQgbmV2ZXIKCS8vIGhhdmUgY29uY3VycmVudCBjYWxscyB0byBBcHBlbmQgbWV0aG9kcy4KCWNoaWxkUmVmcyAgICAgICAgW11zdG9yYWdlLlNlcmllc1JlZgoJd3JpdGVMYXRlbmN5ICAgICBwcm9tZXRoZXVzLkhpc3RvZ3JhbQoJc2FtcGxlc0ZvcndhcmRlZCBwcm9tZXRoZXVzLkNvdW50ZXIKfQoKZnVuYyBOZXdTZXJpZXNSZWZNYXBwaW5nKGNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlciwgc3RvcmUgTWFwcGluZ1N0b3JlLCB3cml0ZUxhdGVuY3kgcHJvbWV0aGV1cy5IaXN0b2dyYW0sIHNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyKSBzdG9yYWdlLkFwcGVuZGVyIHsKCXVuaXF1ZVJlZkNlbGwgOj0gc3RvcmUuR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkKCglyZXR1cm4gJnNlcmllc1JlZk1hcHBpbmd7CgkJY2hpbGRyZW46ICAgICAgICAgY2hpbGRyZW4sCgkJc3RvcmU6ICAgICAgICAgICAgc3RvcmUsCgkJd3JpdGVMYXRlbmN5OiAgICAgd3JpdGVMYXRlbmN5LAoJCXNhbXBsZXNGb3J3YXJkZWQ6IHNhbXBsZXNGb3J3YXJkZWQsCgoJCXVuaXF1ZVJlZkNlbGw6IHVuaXF1ZVJlZkNlbGwsCgkJY2hpbGRSZWZzOiAgICAgbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAwLCBsZW4oY2hpbGRyZW4pKSwKCX0KfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgU2V0T3B0aW9ucyhvcHRzICpzdG9yYWdlLkFwcGVuZE9wdGlvbnMpIHsKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWMuU2V0T3B0aW9ucyhvcHRzKQoJfQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBDb21taXQoKSBlcnJvciB7CglzLnN0b3JlLlRyYWNrQXBwZW5kZWRTZXJpZXModGltZS5Ob3coKS5Vbml4KCksIHMudW5pcXVlUmVmQ2VsbCkKCgl2YXIgbXVsdGlFcnIgZXJyb3IKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWVyciA6PSBjLkNvbW1pdCgpCgkJaWYgZXJyICE9IG5pbCB7CgkJCW11bHRpRXJyID0gbXVsdGllcnJvci5BcHBlbmQobXVsdGlFcnIsIGVycikKCQl9Cgl9CglyZXR1cm4gbXVsdGlFcnIKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgUm9sbGJhY2soKSBlcnJvciB7CgkvLyBXZSBzdGlsbCB0cmFjayByb2xsZWQgYmFjayBzZXJpZXMgc28gd2UgY2FuIHByb3Blcmx5CgkvLyBjbGVhbiB1cCBhbnkgc2VyaWVzIHRoYXQgd2FzIGFwcGVuZGVkCglzLnN0b3JlLlRyYWNrQXBwZW5kZWRTZXJpZXModGltZS5Ob3coKS5Vbml4KCksIHMudW5pcXVlUmVmQ2VsbCkKCgl2YXIgbXVsdGlFcnIgZXJyb3IKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWVyciA6PSBjLlJvbGxiYWNrKCkKCQlpZiBlcnIgIT0gbmlsIHsKCQkJbXVsdGlFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChtdWx0aUVyciwgZXJyKQoJCX0KCX0KCXJldHVybiBtdWx0aUVycgp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSByZWNvcmRMYXRlbmN5KCkgewoJaWYgcy5zdGFydC5Jc1plcm8oKSB7CgkJcmV0dXJuCgl9CglkdXJhdGlvbiA6PSB0aW1lLlNpbmNlKHMuc3RhcnQpCglzLndyaXRlTGF0ZW5jeS5PYnNlcnZlKGR1cmF0aW9uLlNlY29uZHMoKSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgcmVzZXRGaWVsZHMoKSB7CgkvLyBSZXNldCBjaGlsZFJlZnMgc2xpY2UgbGVuZ3RoIHRvIDAgZm9yIHJldXNlCglzLmNoaWxkUmVmcyA9IHMuY2hpbGRSZWZzWzowXQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmQocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQgaW50NjQsIHYgZmxvYXQ2NCkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJbmV3UmVmLCBlcnIgOj0gYXBwZW5kZXIuQXBwZW5kKHJlZiwgbCwgdCwgdikKCQlpZiBlcnIgPT0gbmlsIHsKCQkJcy5zYW1wbGVzRm9yd2FyZGVkLkluYygpCgkJfQoJCXJldHVybiBuZXdSZWYsIGVycgoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kRXhlbXBsYXIocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIGUgZXhlbXBsYXIuRXhlbXBsYXIpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRFeGVtcGxhcihyZWYsIGwsIGUpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRIaXN0b2dyYW0ocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQgaW50NjQsIGggKmhpc3RvZ3JhbS5IaXN0b2dyYW0sIGZoICpoaXN0b2dyYW0uRmxvYXRIaXN0b2dyYW0pIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRIaXN0b2dyYW0ocmVmLCBsLCB0LCBoLCBmaCkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEhpc3RvZ3JhbUNUWmVyb1NhbXBsZShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCwgY3QgaW50NjQsIGggKmhpc3RvZ3JhbS5IaXN0b2dyYW0sIGZoICpoaXN0b2dyYW0uRmxvYXRIaXN0b2dyYW0pIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRIaXN0b2dyYW1DVFplcm9TYW1wbGUocmVmLCBsLCB0LCBjdCwgaCwgZmgpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBVcGRhdGVNZXRhZGF0YShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgbSBtZXRhZGF0YS5NZXRhZGF0YSkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLlVwZGF0ZU1ldGFkYXRhKHJlZiwgbCwgbSkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZENUWmVyb1NhbXBsZShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCwgY3QgaW50NjQpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRDVFplcm9TYW1wbGUocmVmLCBsLCB0LCBjdCkKCX0pCn0KCnR5cGUgYXBwZW5kZXJGdW5jIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgYXBwZW5kVG9DaGlsZHJlbihyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscywgYWYgYXBwZW5kZXJGdW5jKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglkZWZlciBzLnJlc2V0RmllbGRzKCkKCglpZiBzLnN0YXJ0LklzWmVybygpIHsKCQlzLnN0YXJ0ID0gdGltZS5Ob3coKQoJfQoKCS8vIENoZWNrIGlmIHRoZSBpbmNvbWluZyByZWYgaGFzIHJlZiBtYXBwaW5ncwoJZXhpc3RpbmdDaGlsZFJlZnMgOj0gcy5zdG9yZS5HZXRNYXBwaW5nKHJlZiwgbGJscykKCgl2YXIgYXBwZW5kRXJyIGVycm9yCgoJLy8gU2FuaXR5IGNoZWNrOiBpZiB3ZSBoYXZlIGV4aXN0aW5nIGNoaWxkIHJlZnMsIHRoZXkgbXVzdCBtYXRjaCB0aGUgbnVtYmVyIG9mIGNoaWxkcmVuCglpZiBleGlzdGluZ0NoaWxkUmVmcyAhPSBuaWwgJiYgbGVuKGV4aXN0aW5nQ2hpbGRSZWZzKSA9PSBsZW4ocy5jaGlsZHJlbikgewoJCXMudW5pcXVlUmVmQ2VsbC5SZWZzID0gYXBwZW5kKHMudW5pcXVlUmVmQ2VsbC5SZWZzLCByZWYpCgoJCXJlZlVwZGF0ZVJlcXVpcmVkIDo9IGZhbHNlCgkJZm9yIGNoaWxkSW5kZXgsIGNoaWxkUmVmIDo9IHJhbmdlIGV4aXN0aW5nQ2hpbGRSZWZzIHsKCQkJbmV3Q2hpbGRSZWYsIGVyciA6PSBhZihzLmNoaWxkcmVuW2NoaWxkSW5kZXhdLCBjaGlsZFJlZikKCQkJaWYgZXJyICE9IG5pbCB7CgkJCQlhcHBlbmRFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChhcHBlbmRFcnIsIGVycikKCQkJfQoKCQkJaWYgbmV3Q2hpbGRSZWYgIT0gY2hpbGRSZWYgewoJCQkJLy8gQ2hpbGQgcmVmIGNoYW5nZWQsIG5lZWQgdG8gdXBkYXRlIG1hcHBpbmcKCQkJCWV4aXN0aW5nQ2hpbGRSZWZzW2NoaWxkSW5kZXhdID0gbmV3Q2hpbGRSZWYKCQkJCXJlZlVwZGF0ZVJlcXVpcmVkID0gdHJ1ZQoJCQl9CgkJfQoKCQlpZiBhcHBlbmRFcnIgIT0gbmlsIHsKCQkJcmV0dXJuIDAsIGFwcGVuZEVycgoJCX0KCgkJaWYgcmVmVXBkYXRlUmVxdWlyZWQgewoJCQlzLnN0b3JlLlVwZGF0ZU1hcHBpbmcocmVmLCBleGlzdGluZ0NoaWxkUmVmcywgbGJscykKCQl9CgoJCXJldHVybiByZWYsIG5pbAoJfQoKCS8vIE5vIGV4aXN0aW5nIG1hcHBpbmcsIHByb2NlZWQgd2l0aCBub3JtYWwgYXBwZW5kIHRvIGFsbCBjaGlsZHJlbgoJdmFyIGZpcnN0Tm9uWmVyb1JlZiBzdG9yYWdlLlNlcmllc1JlZgoJdmFyIG5vblplcm9Db3VudCBpbnQKCgkvLyBOb3RlOiB0aGVyZSdzIGFub3RoZXIgb3B0aW1pemF0aW9uIHdoZXJlIHdlIGNvdWxkIHVzZSB0aGUgcmV0dXJuZWQgcmVmIGlmIGFsbCB0aGUgbm9uIHplcm8gcmVmcwoJLy8gIGFyZSB0aGUgc2FtZSB2YWx1ZS4gVGhpcyBpc24ndCBzYWZlIGFzIHdlIHdpbGwgbWl4IGRvd25zdHJlYW0gcmVmcyB3aXRoIHVuaXF1ZSByZWZzIHdoaWNoIGNvdWxkCgkvLyAgY29sbGlkZS4gV2UgY291bGQgc3RhcnQgYXQgbWF4IHVuaXQ2NCBmb3IgdW5pcXVlIHJlZnMgYW5kIGdvIGJhY2t3YXJkcyBsZXNzZW5pbmcgdGhlIGNoYW5jZSBvZgoJLy8gCWNvbGxpc2lvbnMgYnV0IGl0J3MgcmF0aGVyIGRhbmdlcm91cyBmb3IgYW4gdW5saWtlbHkgZWRnZSBjYXNlLiBJZiB0d28gY29tcG9uZW50cyBhcmUgcmV0dXJuaW5nCgkvLyAJdGhlIHNhbWUgcmVmIGl0J3MgdHdvIHJlbW90ZV93cml0ZSBjb21wb25lbnRzIHdoaWNoIHNob3VsZCBwcm9iYWJseSBiZSBtZXJnZWQgaW4gdG8gb25lLgoJZm9yIF8sIGNoaWxkIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWNoaWxkUmVmLCBlcnIgOj0gYWYoY2hpbGQsIHJlZikKCQlpZiBlcnIgIT0gbmlsIHsKCQkJYXBwZW5kRXJyID0gbXVsdGllcnJvci5BcHBlbmQoYXBwZW5kRXJyLCBlcnIpCgoJCQkvLyBUT0RPIHNob3VsZCBJIG1ha2UgdGhlIGNoaWxkUmVmIHplcm8gaGVyZT8KCQl9CgoJCXMuY2hpbGRSZWZzID0gYXBwZW5kKHMuY2hpbGRSZWZzLCBjaGlsZFJlZikKCQlpZiBjaGlsZFJlZiAhPSAwIHsKCQkJaWYgZmlyc3ROb25aZXJvUmVmID09IDAgewoJCQkJZmlyc3ROb25aZXJvUmVmID0gY2hpbGRSZWYKCQkJfQoJCQlub25aZXJvQ291bnQrKwoJCX0KCX0KCglpZiBhcHBlbmRFcnIgIT0gbmlsIHsKCQlyZXR1cm4gMCwgYXBwZW5kRXJyCgl9CgoJaWYgbm9uWmVyb0NvdW50ID09IDAgewoJCS8vIEFsbCBjaGlsZHJlbiByZXR1cm5lZCByZWYgMCwgc28gcmV0dXJuIHRoZSBpbnB1dCByZWYKCQlyZXR1cm4gcmVmLCBuaWwKCX0KCgkvLyBPbmx5IG9uZSBjaGlsZCByZXR1cm5lZCBhIG5vbi16ZXJvIHJlZiwgdXNlIHRoYXQKCWlmIG5vblplcm9Db3VudCA9PSAxIHsKCQlyZXR1cm4gZmlyc3ROb25aZXJvUmVmLCBuaWwKCX0KCgkvLyBXZSBnb3QgZGlmZmVyZW50IHJlZnMgYmFjayBhbmQgbmVlZCB0byBjcmVhdGUgYSBuZXcgbWFwcGluZwoJdW5pcXVlUmVmIDo9IHMuc3RvcmUuQ3JlYXRlTWFwcGluZyhzLmNoaWxkUmVmcywgbGJscykKCXMudW5pcXVlUmVmQ2VsbC5SZWZzID0gYXBwZW5kKHMudW5pcXVlUmVmQ2VsbC5SZWZzLCB1bmlxdWVSZWYpCglyZXR1cm4gdW5pcXVlUmVmLCBuaWwKfQoKdHlwZSB1bmlxUmVmQ2hpbGRyZW4gc3RydWN0IHsKCWNoaWxkUmVmcyAqW11zdG9yYWdlLlNlcmllc1JlZgoJbGFiZWxIYXNoIHVpbnQ2NAp9Cgp0eXBlIFNlcmllc1JlZk1hcHBpbmdTdG9yZSBzdHJ1Y3QgewoJLy8gcmVmTWFwcGluZ011IHByb3RlY3RzIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzLCBsYWJlbEhhc2hUb1VuaXF1ZVJlZiBhbmQgbmV4dFVuaXF1ZVJlZgoJcmVmTWFwcGluZ011IHN5bmMuUldNdXRleAoJLy8gdW5pcXVlUmVmVG9DaGlsZFJlZnMgbWFwcyB0aGUgdW5pcXVlIHJlZiB0byB0aGUgZXhwZWN0ZWQgY2hpbGQgcmVmIGluIG9yZGVyCgl1bmlxdWVSZWZUb0NoaWxkUmVmcyBtYXBbc3RvcmFnZS5TZXJpZXNSZWZddW5pcVJlZkNoaWxkcmVuCgkvLyBsYWJlbEhhc2hUb1VuaXF1ZVJlZiBtYXBzIHRoZSBsYWJlbCBoYXNoIHRvIHVuaXF1ZSByZWYuCglsYWJlbEhhc2hUb1VuaXF1ZVJlZiBtYXBbdWludDY0XXN0b3JhZ2UuU2VyaWVzUmVmCgoJLy8gbmV4dFVuaXF1ZVJlZiBpcyB0aGUgbmV4dCByZWYgSUQgd2Ugd2lsbCBoYW5kIG91dAoJbmV4dFVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZgoKCS8vIHRpbWVzdGFtcFRyYWNraW5nTXUgcHJvdGVjdHMgdW5pcXVlUmVmVGltZXN0YW1wcyBhbmQgY2VsbFBvb2wKCXRpbWVzdGFtcFRyYWNraW5nTXUgc3luYy5NdXRleAoJLy8gdW5pcXVlUmVmVGltZXN0YW1wcyBtYXBzIHVuaXF1ZSByZWZzIHRvIHRoZWlyIGxhc3QgYXBwZW5kIHRpbWVzdGFtcAoJdW5pcXVlUmVmVGltZXN0YW1wcyBtYXBbc3RvcmFnZS5TZXJpZXNSZWZdaW50NjQKCS8vIGNlbGxQb29sIGlzIHVzZWQgdG8gcG9vbCBzbGljZXMgb2YgU2VyaWVzUmVmcyB1c2VkIGZvciB0cmFja2luZyB1bmlxdWUgcmVmcyBpbiBUcmFja0FwcGVuZGVkU2VyaWVzLgoJY2VsbFBvb2wgc3luYy5Qb29sCgoJLy8gQ2xlYW51cCBnb3JvdXRpbmUgY29vcmRpbmF0aW9uIChubyBsb2NrIHJlcXVpcmVkKQoJc3RhcnRSZWZDbGVhbnVwIHN5bmMuT25jZQoJY2xlYW51cFN0YXJ0ZWQgIGF0b21pYy5Cb29sCglzdG9wQ2xlYW51cCAgICAgY2hhbiBzdHJ1Y3R7fQoJY2xlYW51cFN0b3BwZWQgIGNoYW4gc3RydWN0e30KCgkvLyBNZXRyaWNzIChzYWZlIGZvciBjb25jdXJyZW50IGFjY2Vzcywgbm8gbG9jayByZXF1aXJlZCkKCWFjdGl2ZU1hcHBpbmdzICBwcm9tZXRoZXVzLkdhdWdlCgl0cmFja2VkUmVmcyAgICAgcHJvbWV0aGV1cy5HYXVnZQoJcmVmc0NsZWFuZWQgICAgIHByb21ldGhldXMuQ291bnRlcgoJdW5pcXVlUmVmc1RvdGFsIHByb21ldGhldXMuQ291bnRlcgp9CgpmdW5jIE5ld1Nlcmllc1JlZk1hcHBpbmdTdG9yZShyZWcgcHJvbWV0aGV1cy5SZWdpc3RlcmVyKSAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlIHsKCWFjdGl2ZU1hcHBpbmdzIDo9IHByb21ldGhldXMuTmV3R2F1Z2UocHJvbWV0aGV1cy5HYXVnZU9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX21hcHBpbmdzX3RvdGFsIiwKCQlIZWxwOiAiTnVtYmVyIG9mIGFjdGl2ZSB1bmlxdWUgcmVmIG1hcHBpbmdzIGluIHRoZSBzdG9yZS4iLAoJfSkKCXRyYWNrZWRSZWZzIDo9IHByb21ldGhldXMuTmV3R2F1Z2UocHJvbWV0aGV1cy5HYXVnZU9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3RyYWNrZWRfcmVmc190b3RhbCIsCgkJSGVscDogIk51bWJlciBvZiByZWZzIGJlaW5nIHRyYWNrZWQgZm9yIHRpbWVzdGFtcC1iYXNlZCBjbGVhbnVwLiIsCgl9KQoJcmVmc0NsZWFuZWQgOj0gcHJvbWV0aGV1cy5OZXdDb3VudGVyKHByb21ldGhldXMuQ291bnRlck9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3JlZnNfY2xlYW5lZF90b3RhbCIsCgkJSGVscDogIlRvdGFsIG51bWJlciBvZiBzdGFsZSByZWZzIGNsZWFuZWQgdXAgb3ZlciB0aW1lLiIsCgl9KQoJdW5pcXVlUmVmc1RvdGFsIDo9IHByb21ldGhldXMuTmV3Q291bnRlcihwcm9tZXRoZXVzLkNvdW50ZXJPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV91bmlxdWVfcmVmc19jcmVhdGVkX3RvdGFsIiwKCQlIZWxwOiAiVG90YWwgbnVtYmVyIG9mIHVuaXF1ZSByZWZzIGNyZWF0ZWQuIiwKCX0pCgoJaWYgcmVnICE9IG5pbCB7CgkJcmVnLlJlZ2lzdGVyKGFjdGl2ZU1hcHBpbmdzKQoJCXJlZy5SZWdpc3Rlcih0cmFja2VkUmVmcykKCQlyZWcuUmVnaXN0ZXIocmVmc0NsZWFuZWQpCgkJcmVnLlJlZ2lzdGVyKHVuaXF1ZVJlZnNUb3RhbCkKCX0KCglyZXR1cm4gJlNlcmllc1JlZk1hcHBpbmdTdG9yZXsKCQl1bmlxdWVSZWZUb0NoaWxkUmVmczogbWFrZShtYXBbc3RvcmFnZS5TZXJpZXNSZWZddW5pcVJlZkNoaWxkcmVuKSwKCQluZXh0VW5pcXVlUmVmOiAgICAgICAgMSwKCQl1bmlxdWVSZWZUaW1lc3RhbXBzOiAgbWFrZShtYXBbc3RvcmFnZS5TZXJpZXNSZWZdaW50NjQpLAoJCWxhYmVsSGFzaFRvVW5pcXVlUmVmOiBtYWtlKG1hcFt1aW50NjRdc3RvcmFnZS5TZXJpZXNSZWYpLAoJCWNlbGxQb29sOiBzeW5jLlBvb2x7CgkJCU5ldzogZnVuYygpIGFueSB7CgkJCQlyZXR1cm4gJkNlbGx7UmVmczogbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAxMDApfQoJCQl9LAoJCX0sCgkJc3RvcENsZWFudXA6ICAgICBtYWtlKGNoYW4gc3RydWN0e30pLAoJCWNsZWFudXBTdG9wcGVkOiAgbWFrZShjaGFuIHN0cnVjdHt9KSwKCQlhY3RpdmVNYXBwaW5nczogIGFjdGl2ZU1hcHBpbmdzLAoJCXRyYWNrZWRSZWZzOiAgICAgdHJhY2tlZFJlZnMsCgkJcmVmc0NsZWFuZWQ6ICAgICByZWZzQ2xlYW5lZCwKCQl1bmlxdWVSZWZzVG90YWw6IHVuaXF1ZVJlZnNUb3RhbCwKCX0KfQoKdHlwZSBDZWxsIHN0cnVjdCB7CglSZWZzIFtdc3RvcmFnZS5TZXJpZXNSZWYKfQoKLy8gR2V0TWFwcGluZyByZXR1cm5zIGV4aXN0aW5nIGNoaWxkIHJlZiByZXN1bHRzIGZvciB0aGUgZ2l2ZW4gdW5pcXVlIHJlZiBpZiBvbmUgZXhpc3RzLgovLwovLyBJZiB0aGUgcGFzc2VkIHVuaXF1ZVJlZiBpcyB6ZXJvLCB0aGUgbWV0aG9kIHdpbGwgYXR0ZW1wdCB0byBmaW5kIGEgbWFwcGluZyB1c2luZyBwYXNzZWQgbGFiZWxzLgovLyBSZXR1cm5zIG5pbCBpZiBubyBtYXBwaW5nIGV4aXN0cy4KLy8KLy8gVGhlIHJldHVybmVkIHNsaWNlIG1heSBiZSBtb2RpZmllZCBieSB0aGUgY2FsbGVyLCBidXQgVXBkYXRlTWFwcGluZyBtdXN0IGJlIGNhbGxlZAovLyBhZnRlcndhcmRzIHRvIHBlcnNpc3QgY2hhbmdlcy4gTm90ZSB0aGF0IGNvbmN1cnJlbnQgYXBwZW5kZXJzIG1heSByYWNlIHRvIHVwZGF0ZSB0aGUKLy8gc2FtZSBtYXBwaW5nIHdpdGggZGlmZmVyZW50IHZhbHVlcywgd2hpY2ggaXMgc2FmZSBiZWNhdXNlIHN0YWxlIG1hcHBpbmdzIGFyZSBzZWxmLWNvcnJlY3RpbmcgLQovLyB1c2luZyBhIHN0YWxlIHJlZiB3aWxsIGNhdXNlIHRoZSBjaGlsZCBhcHBlbmRlciB0byByZXR1cm4gYSBuZXcgcmVmIG9uIHRoZSBuZXh0IGFwcGVuZC4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBHZXRNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSBbXXN0b3JhZ2UuU2VyaWVzUmVmIHsKCXMucmVmTWFwcGluZ011LlJMb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlJVbmxvY2soKQoKCWlmIHVuaXF1ZVJlZiA9PSAwIHsKCQkvLyBTb21lIGNvbnN1bWVycyBkb24ndCBtZW1vIHRoZSBnbG9iYWwgcmVmLiBUcnkgdG8gbG9va3VwIGEgcmVmIGJ5IGxhYmVsIGhhc2guCgkJbGFiZWxIYXNoIDo9IGxibHMuSGFzaCgpCgkJZ290UmVmLCBvayA6PSBzLmxhYmVsSGFzaFRvVW5pcXVlUmVmW2xhYmVsSGFzaF0KCQlpZiAhb2sgewoJCQlyZXR1cm4gbmlsCgkJfQoKCQl1bmlxdWVSZWYgPSBnb3RSZWYKCX0KCglpZiBtYXBwaW5nLCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl07IG9rIHsKCQlyZXR1cm4gKm1hcHBpbmcuY2hpbGRSZWZzCgl9CglyZXR1cm4gbmlsCn0KCi8vIENyZWF0ZU1hcHBpbmcgY3JlYXRlcyBhIG5ldyB1bmlxdWUgcmVmIG1hcHBpbmcgZm9yIHRoZSBnaXZlbiBjaGlsZCByZWYgcmVzdWx0cy4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBDcmVhdGVNYXBwaW5nKHJlZlJlc3VsdHMgW11zdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSBzdG9yYWdlLlNlcmllc1JlZiB7CgkvLyBTdGFydCBjbGVhbnVwIGdvcm91dGluZSBvbiBmaXJzdCBtYXBwaW5nCglzLnN0YXJ0UmVmQ2xlYW51cC5EbyhmdW5jKCkgewoJCXMuY2xlYW51cFN0YXJ0ZWQuU3RvcmUodHJ1ZSkKCQlnbyBzLmNsZWFudXBTdGFsZVJlZnMoKQoJfSkKCgkvLyBTdG9yZSBhIGNvcHkgb2YgdGhlIGNoaWxkIHJlZiByZXN1bHRzIGRpcmVjdGx5CgljaGlsZFJlZlNsaWNlIDo9IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgbGVuKHJlZlJlc3VsdHMpKQoJY29weShjaGlsZFJlZlNsaWNlLCByZWZSZXN1bHRzKQoKCS8vIEhhc2ggbGFiZWxzIHRvIGZvciB0aGUgZmFsbGJhY2sgbG9va3VwIHRhYmxlCglsYWJlbEhhc2ggOj0gbGJscy5IYXNoKCkKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJLy8gQ3JlYXRlIGEgbmV3IHVuaXF1ZSByZWYKCXVuaXF1ZVJlZiA6PSBzLm5leHRVbmlxdWVSZWYKCXMubmV4dFVuaXF1ZVJlZisrCgoJcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltsYWJlbEhhc2hdID0gdW5pcXVlUmVmCglzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0gPSB1bmlxUmVmQ2hpbGRyZW57CgkJY2hpbGRSZWZzOiAmY2hpbGRSZWZTbGljZSwKCQlsYWJlbEhhc2g6IGxhYmVsSGFzaCwKCX0KCglzLmFjdGl2ZU1hcHBpbmdzLkluYygpCglzLnVuaXF1ZVJlZnNUb3RhbC5JbmMoKQoKCXJldHVybiB1bmlxdWVSZWYKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHsKCWlmIHVuaXF1ZVJlZiA9PSAwIHsKCQlyZXR1cm4KCX0KCgljaGlsZFJlZlNsaWNlIDo9IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgbGVuKHJlZlJlc3VsdHMpKQoJY29weShjaGlsZFJlZlNsaWNlLCByZWZSZXN1bHRzKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgkvLyBFbnN1cmUgdGhhdCBsYWJlbCBoYXNoIGluZGV4IGlzIHVwIHRvIGRhdGUgdG8gaGFuZGxlIHBvc3NpYmxlIGhhc2ggY29sbGlzaW9ucy4KCS8vIFRPRE86IGlzIHRoaXMgbmVjZXNzYXJ5PwoJbmV3SGFzaCA6PSBsYmxzLkhhc2goKQoJcHJldiwgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdCglpZiBvayAmJiBwcmV2LmxhYmVsSGFzaCAhPSBuZXdIYXNoIHsKCQlkZWxldGUocy5sYWJlbEhhc2hUb1VuaXF1ZVJlZiwgcHJldi5sYWJlbEhhc2gpCgkJcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltuZXdIYXNoXSA9IHVuaXF1ZVJlZgoJfQoKCXMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXSA9IHVuaXFSZWZDaGlsZHJlbnsKCQljaGlsZFJlZnM6ICZjaGlsZFJlZlNsaWNlLAoJCWxhYmVsSGFzaDogbGJscy5IYXNoKCksCgl9Cn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgVHJhY2tBcHBlbmRlZFNlcmllcyh0cyBpbnQ2NCwgY2VsbCAqQ2VsbCkgewoJcy50aW1lc3RhbXBUcmFja2luZ011LkxvY2soKQoJZGVmZXIgcy50aW1lc3RhbXBUcmFja2luZ011LlVubG9jaygpCgoJZm9yIF8sIHIgOj0gcmFuZ2UgY2VsbC5SZWZzIHsKCQlzLnVuaXF1ZVJlZlRpbWVzdGFtcHNbcl0gPSB0cwoJfQoKCXMudHJhY2tlZFJlZnMuU2V0KGZsb2F0NjQobGVuKHMudW5pcXVlUmVmVGltZXN0YW1wcykpKQoKCWNlbGwuUmVmcyA9IGNlbGwuUmVmc1s6MF0KCXMuY2VsbFBvb2wuUHV0KGNlbGwpCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwgewoJcmV0dXJuIHMuY2VsbFBvb2wuR2V0KCkuKCpDZWxsKQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIGNsZWFudXBTdGFsZVJlZnMoKSB7CglkZWZlciBjbG9zZShzLmNsZWFudXBTdG9wcGVkKQoKCXRpY2tlciA6PSB0aW1lLk5ld1RpY2tlcigxNSAqIHRpbWUuTWludXRlKQoJZGVmZXIgdGlja2VyLlN0b3AoKQoKCWZvciB7CgkJc2VsZWN0IHsKCQljYXNlIDwtdGlja2VyLkM6CgkJCWN1dG9mZlRpbWUgOj0gdGltZS5Ob3coKS5BZGQoLTE1ICogdGltZS5NaW51dGUpLlVuaXgoKQoKCQkJLy8gSG9sZCBib3RoIGxvY2tzIHRvIHByZXZlbnQgcmFjZSBjb25kaXRpb24gd2hlcmUgYSByZWYgY291bGQgYmUKCQkJLy8gYXBwZW5kZWQgYWZ0ZXIgd2UgZGVsZXRlIGl0IGZyb20gdW5pcXVlUmVmQ2VsbCBidXQgYmVmb3JlCgkJCS8vIHdlIGRlbGV0ZSBpdCBmcm9tIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzCgkJCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCQkJcy5yZWZNYXBwaW5nTXUuTG9jaygpCgoJCQlzdGFsZVJlZkNvdW50IDo9IDAKCQkJZm9yIHJlZiwgdHMgOj0gcmFuZ2Ugcy51bmlxdWVSZWZUaW1lc3RhbXBzIHsKCQkJCWlmIHRzIDwgY3V0b2ZmVGltZSB7CgkJCQkJc3RhbGVSZWZDb3VudCsrCgoJCQkJCXYsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbcmVmXQoJCQkJCWlmIG9rIHsKCQkJCQkJZGVsZXRlKHMubGFiZWxIYXNoVG9VbmlxdWVSZWYsIHYubGFiZWxIYXNoKQoJCQkJCX0KCgkJCQkJZGVsZXRlKHMudW5pcXVlUmVmVGltZXN0YW1wcywgcmVmKQoJCQkJCWRlbGV0ZShzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzLCByZWYpCgkJCQl9CgkJCX0KCgkJCS8vIFVwZGF0ZSBtZXRyaWNzCgkJCWlmIHN0YWxlUmVmQ291bnQgPiAwIHsKCQkJCXMucmVmc0NsZWFuZWQuQWRkKGZsb2F0NjQoc3RhbGVSZWZDb3VudCkpCgkJCQlzLmFjdGl2ZU1hcHBpbmdzLlN1YihmbG9hdDY0KHN0YWxlUmVmQ291bnQpKQoJCQkJcy50cmFja2VkUmVmcy5TZXQoZmxvYXQ2NChsZW4ocy51bmlxdWVSZWZUaW1lc3RhbXBzKSkpCgkJCX0KCgkJCXMucmVmTWFwcGluZ011LlVubG9jaygpCgkJCXMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCQljYXNlIDwtcy5zdG9wQ2xlYW51cDoKCQkJcmV0dXJuCgkJfQoJfQp9CgovLyBDbGVhciB3aWxsIGNsZWFyIGFsbCBpbnRlcm5hbCBtYXBwaW5ncyBhbmQgc3RvcCB0aGUgY2xlYW5lciBnb3JvdXRpbmUgaWYgaXQgaXMgcnVubmluZy4KLy8gSXQgaXMgc2FmZSB0byByZS11c2UgdGhlIHNhbWUgaW5zdGFuY2UgYWZ0ZXIgY2FsbGluZyBDbGVhci4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBDbGVhcigpIHsKCS8vIFN0b3AgdGhlIGNsZWFudXAgZ29yb3V0aW5lIGFuZCB3YWl0IGZvciBpdCB0byBiZSBzdG9wcGVkIHNvIHdlIGNhbgoJLy8gYXZvaWQgYSBwb3NzaWJsZSBkZWFkbG9jayB3aXRoIGNsZWFudXAgdGhhdCBhbHNvIGhvbGRzIGJvdGggbG9ja3MKCWlmIHMuY2xlYW51cFN0YXJ0ZWQuTG9hZCgpIHsKCQlzZWxlY3QgewoJCWNhc2UgPC1zLnN0b3BDbGVhbnVwOgoJCQkvLyBBbHJlYWR5IGNsb3NlZAoJCWRlZmF1bHQ6CgkJCWNsb3NlKHMuc3RvcENsZWFudXApCgkJCTwtcy5jbGVhbnVwU3RvcHBlZAoJCX0KCX0KCgkvLyBXZSBuZWVkIHRvIGhvbGQgYm90aCBsb2NrcyB0byBkbyB0aGlzIHNhZmVseSBhbmQgd2UgZG8gaXQgaW4gdGhlIHNhbWUgb3JkZXIgYXMKCS8vIGNsZWFudXBTdGFsZVJlZnMuIFdlIHN0b3BwZWQgYW5kIHdhaXRlZCBmb3IgdGhlIGJhY2tncm91bmQgd29ya2VyIHRoYXQgY2FsbHMgaXQKCS8vIHRvIGZpbmlzaCBidXQgc29tZSBleHRyYSBzYWZldHkgd29uJ3QgaHVydC4KCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCWRlZmVyIHMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgljbGVhcihzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzKQoJY2xlYXIocy51bmlxdWVSZWZUaW1lc3RhbXBzKQoKCS8vIHJlc2V0IHRoZSBwb29sCglzLmNlbGxQb29sID0gc3luYy5Qb29sewoJCU5ldzogZnVuYygpIGFueSB7CgkJCXJldHVybiAmQ2VsbHtSZWZzOiBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIDAsIDEwMCl9CgkJfSwKCX0KCgkvLyBOT1RFOiBXZSBkbyBOT1QgcmVzZXQgbmV4dFVuaXF1ZVJlZiBoZXJlLiBSZXNldHRpbmcgaXQgd291bGQgY2F1c2UgcmVmIGNvbGxpc2lvbnMKCS8vIHdpdGggY29tcG9uZW50cyBsaWtlIHByb21ldGhldXMuc2NyYXBlIHdoaWNoIHdpbGwga2VlcCByZS1zZW5kaW5nIHRoZSBzYW1lIGNhY2hlZCByZWZzLgoJLy8gV2UgY29udGludWUgaW5jcmVtZW50aW5nIHRvIGVuc3VyZSBhbGwgcmVmcyByZW1haW4gdW5pcXVlIGFjcm9zcyB0aGUgbGlmZXRpbWUgb2YgdGhlIHByb2Nlc3MuCgoJLy8gUmVzZXQgbWV0cmljcwoJcy5hY3RpdmVNYXBwaW5ncy5TZXQoMCkKCXMudHJhY2tlZFJlZnMuU2V0KDApCgoJLy8gUmVzZXQgY2hhbm5lbHMgYW5kIGZsYWdzCglzLnN0b3BDbGVhbnVwID0gbWFrZShjaGFuIHN0cnVjdHt9KQoJcy5jbGVhbnVwU3RvcHBlZCA9IG1ha2UoY2hhbiBzdHJ1Y3R7fSkKCXMuc3RhcnRSZWZDbGVhbnVwID0gc3luYy5PbmNle30KCXMuY2xlYW51cFN0YXJ0ZWQuU3RvcmUoZmFsc2UpCn0K"
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/component/prometheus/appenders",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/validator",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy [github.com/grafana/alloy.test]",
    "github.com/grafana/alloy.test",
    "github.com/grafana/alloy/internal/alloycli [github.com/grafana/alloy/internal/alloycli.test]",
    "github.com/grafana/alloy/internal/alloycli.test",
    "github.com/grafana/alloy/internal/component/all [github.com/grafana/alloy/internal/component/all.test]",
    "github.com/grafana/alloy/internal/component/all.test",
    "github.com/grafana/alloy/internal/component/metadata [github.com/grafana/alloy/internal/component/metadata.test]",
    "github.com/grafana/alloy/internal/component/metadata.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test [github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test [github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus.test",
    "github.com/grafana/alloy/internal/component/prometheus [github.com/grafana/alloy/internal/component/prometheus.test]",
    "github.com/grafana/alloy/internal/component/prometheus_test [github.com/grafana/alloy/internal/component/prometheus.test]",
    "github.com/grafana/alloy/internal/component/prometheus.test",
    "github.com/grafana/alloy/internal/component/prometheus/appenders [github.com/grafana/alloy/internal/component/prometheus/appenders.test]",
    "github.com/grafana/alloy/internal/component/prometheus/appenders.test",
    "github.com/grafana/alloy/internal/component/prometheus/enrich [github.com/grafana/alloy/internal/component/prometheus/enrich.test]",
    "github.com/grafana/alloy/internal/component/prometheus/enrich.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test [github.com/grafana/alloy/internal/component/prometheus/exporter/tests.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator [github.com/grafana/alloy/internal/component/prometheus/operator.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common [github.com/grafana/alloy/internal/component/prometheus/operator/common.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen [github.com/grafana/alloy/internal/component/prometheus/operator/configgen.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen.test",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http [github.com/grafana/alloy/internal/component/prometheus/receive_http.test]",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http.test",
    "github.com/grafana/alloy/internal/component/prometheus/relabel [github.com/grafana/alloy/internal/component/prometheus/relabel.test]",
    "github.com/grafana/alloy/internal/component/prometheus/relabel.test",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite [github.com/grafana/alloy/internal/component/prometheus/remotewrite.test]",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test [github.com/grafana/alloy/internal/component/prometheus/remotewrite.test]",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite.test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape [github.com/grafana/alloy/internal/component/prometheus/scrape.test]",
    "github.com/grafana/alloy/internal/component/prometheus/scrape.test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape [github.com/grafana/alloy/internal/component/pyroscope/scrape.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape.test",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test [github.com/grafana/alloy/internal/converter/internal/otelcolconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test [github.com/grafana/alloy/internal/converter/internal/prometheusconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test [github.com/grafana/alloy/internal/converter/internal/promtailconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test [github.com/grafana/alloy/internal/converter/internal/staticconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build [github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build.test]",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build.test",
    "github.com/grafana/alloy/internal/service/cluster [github.com/grafana/alloy/internal/service/cluster.test]",
    "github.com/grafana/alloy/internal/service/cluster_test [github.com/grafana/alloy/internal/service/cluster.test]",
    "github.com/grafana/alloy/internal/service/cluster.test",
    "github.com/grafana/alloy/internal/tools/docs_generator_test [github.com/grafana/alloy/internal/tools/docs_generator.test]",
    "github.com/grafana/alloy/internal/tools/docs_generator.test",
    "github.com/grafana/alloy/internal/validator [github.com/grafana/alloy/internal/validator.test]",
    "github.com/grafana/alloy/internal/validator.test"
  ],
  "Packages": [
    {
      "ID": "archive/tar",
      "Name": "tar",
      "PkgPath": "archive/tar",
      "GoFiles": [
        "/usr/lib/go/src/archive/tar/common.go",
        "/usr/lib/go/src/archive/tar/format.go",
        "/usr/lib/go/src/archive/tar/reader.go",
        "/usr/lib/go/src/archive/tar/stat_actime1.go",
        "/usr/lib/go/src/archive/tar/stat_unix.go",
        "/usr/lib/go/src/archive/tar/strconv.go",
        "/usr/lib/go/src/archive/tar/writer.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/archive/tar/common.go",
        "/usr/lib/go/src/archive/tar/format.go",
        "/usr/lib/go/src/archive/tar/reader.go",
        "/usr/lib/go/src/archive/tar/stat_actime1.go",
        "/usr/lib/go/src/archive/tar/stat_unix.go",
        "/usr/lib/go/src/archive/tar/strconv.go",
        "/usr/lib/go/src/archive/tar/writer.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/archive/tar/stat_actime2.go"],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "fmt": "fmt",
        "internal/godebug": "internal/godebug",
        "io": "io",
        "io/fs": "io/fs",
        "maps": "maps",
        "math": "math",
        "os/user": "os/user",
        "path": "path",
        "path/filepath": "path/filepath",
        "reflect": "reflect",
        "runtime": "runtime",
        "slices": "slices",
        "strconv": "strconv",
        "strings": "strings",
        "sync": "sync",
        "syscall": "syscall",
        "time": "time"
      }
    },
    {
      "ID": "archive/zip",
      "Name": "zip",
      "PkgPath": "archive/zip",
      "GoFiles": [
        "/usr/lib/go/src/archive/zip/reader.go",
        "/usr/lib/go/src/archive/zip/register.go",
        "/usr/lib/go/src/archive/zip/struct.go",
        "/usr/lib/go/src/archive/zip/writer.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/archive/zip/reader.go",
        "/usr/lib/go/src/archive/zip/register.go",
        "/usr/lib/go/src/archive/zip/struct.go",
        "/usr/lib/go/src/archive/zip/writer.go"
      ],
      "Imports": {
        "bufio": "bufio",
        "compress/flate": "compress/flate",
        "encoding/binary": "encoding/binary",
        "errors": "errors",
        "fmt": "fmt",
        "hash": "hash",
        "hash/crc32": "hash/crc32",
        "internal/godebug": "internal/godebug",
        "io": "io",
        "io/fs": "io/fs",
        "os": "os",
        "path": "path",
        "path/filepath": "path/filepath",
        "slices": "slices",
        "strings": "strings",
        "sync": "sync",
        "time": "time",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cloud.google.com/go/auth",
      "Name": "auth",
      "PkgPath": "cloud.google.com/go/auth",
      "GoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/auth.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/threelegged.go"
      ],
      "CompiledGoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/auth.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/threelegged.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "cloud.google.com/go/auth/internal": "cloud.google.com/go/auth/internal",
        "cloud.google.com/go/auth/internal/jwt": "cloud.google.com/go/auth/internal/jwt",
        "context": "context",
        "encoding/json": "encoding/json",
        "errors": "errors",
        "fmt": "fmt",
        "github.com/googleapis/gax-go/v2/internallog": "github.com/googleapis/gax-go/v2/internallog",
        "log/slog": "log/slog",
        "mime": "mime",
        "net/http": "net/http",
        "net/url": "net/url",
        "strconv": "strconv",
        "strings": "strings",
        "sync": "sync",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 4800 to 5 entries.

#### drv #13

Trace meta: spanId=25, ts=1770837177798, ts_iso=2026-02-11T19:12:57.798000Z

Request (DriverRequestEnvelope):

```json
{
  "workDir": "/home/username/work/grafana/alloy",
  "patterns": [
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy/internal/validator"
  ],
  "driverRequest": {
    "mode": 32287,
    "env": [
      "GOMODCACHE=/home/username/go/pkg/mod",
      "GOGCCFLAGS=-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4141799624=/tmp/go-build -gno-record-gcc-switches",
      "GONOSUMDB=",
      "GOCACHE=/home/username/.cache/go-build",
      "GOTMPDIR=",
      "CGO_CFLAGS=-O2 -g",
      "GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64",
      "GOOS=linux",
      "GOFLAGS=",
      "GOPATH=/home/username/go",
      "GOAMD64=v1",
      "GODEBUG=",
      "GO111MODULE=",
      "PKG_CONFIG=pkg-config",
      "CGO_CPPFLAGS=",
      "GOFIPS140=off",
      "AR=ar",
      "GOMOD=/home/username/work/grafana/alloy/go.mod",
      "CGO_CXXFLAGS=-O2 -g",
      "GOSUMDB=sum.golang.org",
      "GOEXE=",
      "CC=gcc",
      "GOTOOLCHAIN=auto",
      "GOENV=/home/username/.config/go/env",
      "GOHOSTOS=linux",
      "CGO_ENABLED=1",
      "GONOPROXY=",
      "GOVERSION=go1.25.7 X:nodwarf5",
      "GOTELEMETRY=off",
      "GOHOSTARCH=amd64",
      "CGO_FFLAGS=-O2 -g",
      "GOVCS=",
      "GOPRIVATE=",
      "CXX=g++",
      "GOTELEMETRYDIR=/home/username/.config/go/telemetry",
      "GOEXPERIMENT=nodwarf5",
      "GOROOT=/usr/lib/go",
      "CGO_LDFLAGS=-O2 -g",
      "GOBIN=",
      "GOPROXY=https://proxy.golang.org,direct",
      "GCCGO=gccgo",
      "GOCACHEPROG=",
      "GOARCH=amd64",
      "GOAUTH=netrc",
      "GO111MODULE=auto",
      "GOPROXY=off"
    ],
    "build_flags": ["-tags", "integration"],
    "tests": true,
    "overlay": {
      "/home/username/work/grafana/alloy/internal/component/prometheus/appenders/seriesrefmapping.go": "cGFja2FnZSBhcHBlbmRlcnMKCmltcG9ydCAoCgkic3luYyIKCSJzeW5jL2F0b21pYyIKCSJ0aW1lIgoKCSJnaXRodWIuY29tL2hhc2hpY29ycC9nby1tdWx0aWVycm9yIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9jbGllbnRfZ29sYW5nL3Byb21ldGhldXMiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvZXhlbXBsYXIiCgkiZ2l0aHViLmNvbS9wcm9tZXRoZXVzL3Byb21ldGhldXMvbW9kZWwvaGlzdG9ncmFtIgoJImdpdGh1Yi5jb20vcHJvbWV0aGV1cy9wcm9tZXRoZXVzL21vZGVsL2xhYmVscyIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9tb2RlbC9tZXRhZGF0YSIKCSJnaXRodWIuY29tL3Byb21ldGhldXMvcHJvbWV0aGV1cy9zdG9yYWdlIgopCgp0eXBlIE1hcHBpbmdTdG9yZSBpbnRlcmZhY2UgewoJR2V0TWFwcGluZyh1bmlxdWVSZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgW11zdG9yYWdlLlNlcmllc1JlZgoJQ3JlYXRlTWFwcGluZyhyZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykgc3RvcmFnZS5TZXJpZXNSZWYKCVVwZGF0ZU1hcHBpbmcodW5pcXVlUmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCByZWZSZXN1bHRzIFtdc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscykKCVRyYWNrQXBwZW5kZWRTZXJpZXModHMgaW50NjQsIGNlbGwgKkNlbGwpCglHZXRDZWxsRm9yQXBwZW5kZWRTZXJpZXMoKSAqQ2VsbAp9Cgp0eXBlIHNlcmllc1JlZk1hcHBpbmcgc3RydWN0IHsKCXN0YXJ0ICAgIHRpbWUuVGltZQoJY2hpbGRyZW4gW11zdG9yYWdlLkFwcGVuZGVyCglzdG9yZSAgICBNYXBwaW5nU3RvcmUKCgl1bmlxdWVSZWZDZWxsICpDZWxsCgoJLy8gY2hpbGRSZWZzIGlzIHJldXNlZCBmb3IgZWFjaCBhcHBlbmQgY2FsbCB0byBhdm9pZCBhbGxvY2F0aW9ucy4gVGhpcyBpcyBzYWZlIGJlY2F1c2Ugc3RvcmFnZS5BcHBlbmRlciBzaG91bGQgbmV2ZXIKCS8vIGhhdmUgY29uY3VycmVudCBjYWxscyB0byBBcHBlbmQgbWV0aG9kcy4KCWNoaWxkUmVmcyAgICAgICAgW11zdG9yYWdlLlNlcmllc1JlZgoJd3JpdGVMYXRlbmN5ICAgICBwcm9tZXRoZXVzLkhpc3RvZ3JhbQoJc2FtcGxlc0ZvcndhcmRlZCBwcm9tZXRoZXVzLkNvdW50ZXIKfQoKZnVuYyBOZXdTZXJpZXNSZWZNYXBwaW5nKGNoaWxkcmVuIFtdc3RvcmFnZS5BcHBlbmRlciwgc3RvcmUgTWFwcGluZ1N0b3JlLCB3cml0ZUxhdGVuY3kgcHJvbWV0aGV1cy5IaXN0b2dyYW0sIHNhbXBsZXNGb3J3YXJkZWQgcHJvbWV0aGV1cy5Db3VudGVyKSBzdG9yYWdlLkFwcGVuZGVyIHsKCXVuaXF1ZVJlZkNlbGwgOj0gc3RvcmUuR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkKCglyZXR1cm4gJnNlcmllc1JlZk1hcHBpbmd7CgkJY2hpbGRyZW46ICAgICAgICAgY2hpbGRyZW4sCgkJc3RvcmU6ICAgICAgICAgICAgc3RvcmUsCgkJd3JpdGVMYXRlbmN5OiAgICAgd3JpdGVMYXRlbmN5LAoJCXNhbXBsZXNGb3J3YXJkZWQ6IHNhbXBsZXNGb3J3YXJkZWQsCgoJCXVuaXF1ZVJlZkNlbGw6IHVuaXF1ZVJlZkNlbGwsCgkJY2hpbGRSZWZzOiAgICAgbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAwLCBsZW4oY2hpbGRyZW4pKSwKCX0KfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgU2V0T3B0aW9ucyhvcHRzICpzdG9yYWdlLkFwcGVuZE9wdGlvbnMpIHsKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWMuU2V0T3B0aW9ucyhvcHRzKQoJfQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBDb21taXQoKSBlcnJvciB7CglzLnN0b3JlLlRyYWNrQXBwZW5kZWRTZXJpZXModGltZS5Ob3coKS5Vbml4KCksIHMudW5pcXVlUmVmQ2VsbCkKCgl2YXIgbXVsdGlFcnIgZXJyb3IKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWVyciA6PSBjLkNvbW1pdCgpCgkJaWYgZXJyICE9IG5pbCB7CgkJCW11bHRpRXJyID0gbXVsdGllcnJvci5BcHBlbmQobXVsdGlFcnIsIGVycikKCQl9Cgl9CglyZXR1cm4gbXVsdGlFcnIKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgUm9sbGJhY2soKSBlcnJvciB7CgkvLyBXZSBzdGlsbCB0cmFjayByb2xsZWQgYmFjayBzZXJpZXMgc28gd2UgY2FuIHByb3Blcmx5CgkvLyBjbGVhbiB1cCBhbnkgc2VyaWVzIHRoYXQgd2FzIGFwcGVuZGVkCglzLnN0b3JlLlRyYWNrQXBwZW5kZWRTZXJpZXModGltZS5Ob3coKS5Vbml4KCksIHMudW5pcXVlUmVmQ2VsbCkKCgl2YXIgbXVsdGlFcnIgZXJyb3IKCWZvciBfLCBjIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWVyciA6PSBjLlJvbGxiYWNrKCkKCQlpZiBlcnIgIT0gbmlsIHsKCQkJbXVsdGlFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChtdWx0aUVyciwgZXJyKQoJCX0KCX0KCXJldHVybiBtdWx0aUVycgp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSByZWNvcmRMYXRlbmN5KCkgewoJaWYgcy5zdGFydC5Jc1plcm8oKSB7CgkJcmV0dXJuCgl9CglkdXJhdGlvbiA6PSB0aW1lLlNpbmNlKHMuc3RhcnQpCglzLndyaXRlTGF0ZW5jeS5PYnNlcnZlKGR1cmF0aW9uLlNlY29uZHMoKSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgcmVzZXRGaWVsZHMoKSB7CgkvLyBSZXNldCBjaGlsZFJlZnMgc2xpY2UgbGVuZ3RoIHRvIDAgZm9yIHJldXNlCglzLmNoaWxkUmVmcyA9IHMuY2hpbGRSZWZzWzowXQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmQocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQgaW50NjQsIHYgZmxvYXQ2NCkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJbmV3UmVmLCBlcnIgOj0gYXBwZW5kZXIuQXBwZW5kKHJlZiwgbCwgdCwgdikKCQlpZiBlcnIgPT0gbmlsIHsKCQkJcy5zYW1wbGVzRm9yd2FyZGVkLkluYygpCgkJfQoJCXJldHVybiBuZXdSZWYsIGVycgoJfSkKfQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgQXBwZW5kRXhlbXBsYXIocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIGUgZXhlbXBsYXIuRXhlbXBsYXIpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRFeGVtcGxhcihyZWYsIGwsIGUpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBBcHBlbmRIaXN0b2dyYW0ocmVmIHN0b3JhZ2UuU2VyaWVzUmVmLCBsIGxhYmVscy5MYWJlbHMsIHQgaW50NjQsIGggKmhpc3RvZ3JhbS5IaXN0b2dyYW0sIGZoICpoaXN0b2dyYW0uRmxvYXRIaXN0b2dyYW0pIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRIaXN0b2dyYW0ocmVmLCBsLCB0LCBoLCBmaCkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZEhpc3RvZ3JhbUNUWmVyb1NhbXBsZShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCwgY3QgaW50NjQsIGggKmhpc3RvZ3JhbS5IaXN0b2dyYW0sIGZoICpoaXN0b2dyYW0uRmxvYXRIaXN0b2dyYW0pIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRIaXN0b2dyYW1DVFplcm9TYW1wbGUocmVmLCBsLCB0LCBjdCwgaCwgZmgpCgl9KQp9CgpmdW5jIChzICpzZXJpZXNSZWZNYXBwaW5nKSBVcGRhdGVNZXRhZGF0YShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgbSBtZXRhZGF0YS5NZXRhZGF0YSkgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJcmV0dXJuIHMuYXBwZW5kVG9DaGlsZHJlbihyZWYsIGwsIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CgkJcmV0dXJuIGFwcGVuZGVyLlVwZGF0ZU1ldGFkYXRhKHJlZiwgbCwgbSkKCX0pCn0KCmZ1bmMgKHMgKnNlcmllc1JlZk1hcHBpbmcpIEFwcGVuZENUWmVyb1NhbXBsZShyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGwgbGFiZWxzLkxhYmVscywgdCwgY3QgaW50NjQpIChzdG9yYWdlLlNlcmllc1JlZiwgZXJyb3IpIHsKCXJldHVybiBzLmFwcGVuZFRvQ2hpbGRyZW4ocmVmLCBsLCBmdW5jKGFwcGVuZGVyIHN0b3JhZ2UuQXBwZW5kZXIsIHJlZiBzdG9yYWdlLlNlcmllc1JlZikgKHN0b3JhZ2UuU2VyaWVzUmVmLCBlcnJvcikgewoJCXJldHVybiBhcHBlbmRlci5BcHBlbmRDVFplcm9TYW1wbGUocmVmLCBsLCB0LCBjdCkKCX0pCn0KCnR5cGUgYXBwZW5kZXJGdW5jIGZ1bmMoYXBwZW5kZXIgc3RvcmFnZS5BcHBlbmRlciwgcmVmIHN0b3JhZ2UuU2VyaWVzUmVmKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKQoKZnVuYyAocyAqc2VyaWVzUmVmTWFwcGluZykgYXBwZW5kVG9DaGlsZHJlbihyZWYgc3RvcmFnZS5TZXJpZXNSZWYsIGxibHMgbGFiZWxzLkxhYmVscywgYWYgYXBwZW5kZXJGdW5jKSAoc3RvcmFnZS5TZXJpZXNSZWYsIGVycm9yKSB7CglkZWZlciBzLnJlc2V0RmllbGRzKCkKCglpZiBzLnN0YXJ0LklzWmVybygpIHsKCQlzLnN0YXJ0ID0gdGltZS5Ob3coKQoJfQoKCS8vIENoZWNrIGlmIHRoZSBpbmNvbWluZyByZWYgaGFzIHJlZiBtYXBwaW5ncwoJZXhpc3RpbmdDaGlsZFJlZnMgOj0gcy5zdG9yZS5HZXRNYXBwaW5nKHJlZiwgbGJscykKCgl2YXIgYXBwZW5kRXJyIGVycm9yCgoJLy8gU2FuaXR5IGNoZWNrOiBpZiB3ZSBoYXZlIGV4aXN0aW5nIGNoaWxkIHJlZnMsIHRoZXkgbXVzdCBtYXRjaCB0aGUgbnVtYmVyIG9mIGNoaWxkcmVuCglpZiBleGlzdGluZ0NoaWxkUmVmcyAhPSBuaWwgJiYgbGVuKGV4aXN0aW5nQ2hpbGRSZWZzKSA9PSBsZW4ocy5jaGlsZHJlbikgewoJCXMudW5pcXVlUmVmQ2VsbC5SZWZzID0gYXBwZW5kKHMudW5pcXVlUmVmQ2VsbC5SZWZzLCByZWYpCgoJCXJlZlVwZGF0ZVJlcXVpcmVkIDo9IGZhbHNlCgkJZm9yIGNoaWxkSW5kZXgsIGNoaWxkUmVmIDo9IHJhbmdlIGV4aXN0aW5nQ2hpbGRSZWZzIHsKCQkJbmV3Q2hpbGRSZWYsIGVyciA6PSBhZihzLmNoaWxkcmVuW2NoaWxkSW5kZXhdLCBjaGlsZFJlZikKCQkJaWYgZXJyICE9IG5pbCB7CgkJCQlhcHBlbmRFcnIgPSBtdWx0aWVycm9yLkFwcGVuZChhcHBlbmRFcnIsIGVycikKCQkJfQoKCQkJaWYgbmV3Q2hpbGRSZWYgIT0gY2hpbGRSZWYgewoJCQkJLy8gQ2hpbGQgcmVmIGNoYW5nZWQsIG5lZWQgdG8gdXBkYXRlIG1hcHBpbmcKCQkJCWV4aXN0aW5nQ2hpbGRSZWZzW2NoaWxkSW5kZXhdID0gbmV3Q2hpbGRSZWYKCQkJCXJlZlVwZGF0ZVJlcXVpcmVkID0gdHJ1ZQoJCQl9CgkJfQoKCQlpZiBhcHBlbmRFcnIgIT0gbmlsIHsKCQkJcmV0dXJuIDAsIGFwcGVuZEVycgoJCX0KCgkJaWYgcmVmVXBkYXRlUmVxdWlyZWQgewoJCQlzLnN0b3JlLlVwZGF0ZU1hcHBpbmcocmVmLCBleGlzdGluZ0NoaWxkUmVmcywgbGJscykKCQl9CgoJCXJldHVybiByZWYsIG5pbAoJfQoKCS8vIE5vIGV4aXN0aW5nIG1hcHBpbmcsIHByb2NlZWQgd2l0aCBub3JtYWwgYXBwZW5kIHRvIGFsbCBjaGlsZHJlbgoJdmFyIGZpcnN0Tm9uWmVyb1JlZiBzdG9yYWdlLlNlcmllc1JlZgoJdmFyIG5vblplcm9Db3VudCBpbnQKCgkvLyBOb3RlOiB0aGVyZSdzIGFub3RoZXIgb3B0aW1pemF0aW9uIHdoZXJlIHdlIGNvdWxkIHVzZSB0aGUgcmV0dXJuZWQgcmVmIGlmIGFsbCB0aGUgbm9uIHplcm8gcmVmcwoJLy8gIGFyZSB0aGUgc2FtZSB2YWx1ZS4gVGhpcyBpc24ndCBzYWZlIGFzIHdlIHdpbGwgbWl4IGRvd25zdHJlYW0gcmVmcyB3aXRoIHVuaXF1ZSByZWZzIHdoaWNoIGNvdWxkCgkvLyAgY29sbGlkZS4gV2UgY291bGQgc3RhcnQgYXQgbWF4IHVuaXQ2NCBmb3IgdW5pcXVlIHJlZnMgYW5kIGdvIGJhY2t3YXJkcyBsZXNzZW5pbmcgdGhlIGNoYW5jZSBvZgoJLy8gCWNvbGxpc2lvbnMgYnV0IGl0J3MgcmF0aGVyIGRhbmdlcm91cyBmb3IgYW4gdW5saWtlbHkgZWRnZSBjYXNlLiBJZiB0d28gY29tcG9uZW50cyBhcmUgcmV0dXJuaW5nCgkvLyAJdGhlIHNhbWUgcmVmIGl0J3MgdHdvIHJlbW90ZV93cml0ZSBjb21wb25lbnRzIHdoaWNoIHNob3VsZCBwcm9iYWJseSBiZSBtZXJnZWQgaW4gdG8gb25lLgoJZm9yIF8sIGNoaWxkIDo9IHJhbmdlIHMuY2hpbGRyZW4gewoJCWNoaWxkUmVmLCBlcnIgOj0gYWYoY2hpbGQsIHJlZikKCQlpZiBlcnIgIT0gbmlsIHsKCQkJYXBwZW5kRXJyID0gbXVsdGllcnJvci5BcHBlbmQoYXBwZW5kRXJyLCBlcnIpCgoJCQkvLyBUT0RPIHNob3VsZCBJIG1ha2UgdGhlIGNoaWxkUmVmIHplcm8gaGVyZT8KCQl9CgoJCXMuY2hpbGRSZWZzID0gYXBwZW5kKHMuY2hpbGRSZWZzLCBjaGlsZFJlZikKCQlpZiBjaGlsZFJlZiAhPSAwIHsKCQkJaWYgZmlyc3ROb25aZXJvUmVmID09IDAgewoJCQkJZmlyc3ROb25aZXJvUmVmID0gY2hpbGRSZWYKCQkJfQoJCQlub25aZXJvQ291bnQrKwoJCX0KCX0KCglpZiBhcHBlbmRFcnIgIT0gbmlsIHsKCQlyZXR1cm4gMCwgYXBwZW5kRXJyCgl9CgoJaWYgbm9uWmVyb0NvdW50ID09IDAgewoJCS8vIEFsbCBjaGlsZHJlbiByZXR1cm5lZCByZWYgMCwgc28gcmV0dXJuIHRoZSBpbnB1dCByZWYKCQlyZXR1cm4gcmVmLCBuaWwKCX0KCgkvLyBPbmx5IG9uZSBjaGlsZCByZXR1cm5lZCBhIG5vbi16ZXJvIHJlZiwgdXNlIHRoYXQKCWlmIG5vblplcm9Db3VudCA9PSAxIHsKCQlyZXR1cm4gZmlyc3ROb25aZXJvUmVmLCBuaWwKCX0KCgkvLyBXZSBnb3QgZGlmZmVyZW50IHJlZnMgYmFjayBhbmQgbmVlZCB0byBjcmVhdGUgYSBuZXcgbWFwcGluZwoJdW5pcXVlUmVmIDo9IHMuc3RvcmUuQ3JlYXRlTWFwcGluZyhzLmNoaWxkUmVmcywgbGJscykKCXMudW5pcXVlUmVmQ2VsbC5SZWZzID0gYXBwZW5kKHMudW5pcXVlUmVmQ2VsbC5SZWZzLCB1bmlxdWVSZWYpCglyZXR1cm4gdW5pcXVlUmVmLCBuaWwKfQoKdHlwZSB1bmlxUmVmQ2hpbGRyZW4gc3RydWN0IHsKCWNoaWxkUmVmcyAqW11zdG9yYWdlLlNlcmllc1JlZgoJbGFiZWxIYXNoIHVpbnQ2NAp9Cgp0eXBlIFNlcmllc1JlZk1hcHBpbmdTdG9yZSBzdHJ1Y3QgewoJLy8gcmVmTWFwcGluZ011IHByb3RlY3RzIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzLCBsYWJlbEhhc2hUb1VuaXF1ZVJlZiBhbmQgbmV4dFVuaXF1ZVJlZgoJcmVmTWFwcGluZ011IHN5bmMuUldNdXRleAoJLy8gdW5pcXVlUmVmVG9DaGlsZFJlZnMgbWFwcyB0aGUgdW5pcXVlIHJlZiB0byB0aGUgZXhwZWN0ZWQgY2hpbGQgcmVmIGluIG9yZGVyCgl1bmlxdWVSZWZUb0NoaWxkUmVmcyBtYXBbc3RvcmFnZS5TZXJpZXNSZWZddW5pcVJlZkNoaWxkcmVuCgkvLyBsYWJlbEhhc2hUb1VuaXF1ZVJlZiBtYXBzIHRoZSBsYWJlbCBoYXNoIHRvIHVuaXF1ZSByZWYuCglsYWJlbEhhc2hUb1VuaXF1ZVJlZiBtYXBbdWludDY0XXN0b3JhZ2UuU2VyaWVzUmVmCgoJLy8gbmV4dFVuaXF1ZVJlZiBpcyB0aGUgbmV4dCByZWYgSUQgd2Ugd2lsbCBoYW5kIG91dAoJbmV4dFVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZgoKCS8vIHRpbWVzdGFtcFRyYWNraW5nTXUgcHJvdGVjdHMgdW5pcXVlUmVmVGltZXN0YW1wcyBhbmQgY2VsbFBvb2wKCXRpbWVzdGFtcFRyYWNraW5nTXUgc3luYy5NdXRleAoJLy8gdW5pcXVlUmVmVGltZXN0YW1wcyBtYXBzIHVuaXF1ZSByZWZzIHRvIHRoZWlyIGxhc3QgYXBwZW5kIHRpbWVzdGFtcAoJdW5pcXVlUmVmVGltZXN0YW1wcyBtYXBbc3RvcmFnZS5TZXJpZXNSZWZdaW50NjQKCS8vIGNlbGxQb29sIGlzIHVzZWQgdG8gcG9vbCBzbGljZXMgb2YgU2VyaWVzUmVmcyB1c2VkIGZvciB0cmFja2luZyB1bmlxdWUgcmVmcyBpbiBUcmFja0FwcGVuZGVkU2VyaWVzLgoJY2VsbFBvb2wgc3luYy5Qb29sCgoJLy8gQ2xlYW51cCBnb3JvdXRpbmUgY29vcmRpbmF0aW9uIChubyBsb2NrIHJlcXVpcmVkKQoJc3RhcnRSZWZDbGVhbnVwIHN5bmMuT25jZQoJY2xlYW51cFN0YXJ0ZWQgIGF0b21pYy5Cb29sCglzdG9wQ2xlYW51cCAgICAgY2hhbiBzdHJ1Y3R7fQoJY2xlYW51cFN0b3BwZWQgIGNoYW4gc3RydWN0e30KCgkvLyBNZXRyaWNzIChzYWZlIGZvciBjb25jdXJyZW50IGFjY2Vzcywgbm8gbG9jayByZXF1aXJlZCkKCWFjdGl2ZU1hcHBpbmdzICBwcm9tZXRoZXVzLkdhdWdlCgl0cmFja2VkUmVmcyAgICAgcHJvbWV0aGV1cy5HYXVnZQoJcmVmc0NsZWFuZWQgICAgIHByb21ldGhldXMuQ291bnRlcgoJdW5pcXVlUmVmc1RvdGFsIHByb21ldGhldXMuQ291bnRlcgp9CgpmdW5jIE5ld1Nlcmllc1JlZk1hcHBpbmdTdG9yZShyZWcgcHJvbWV0aGV1cy5SZWdpc3RlcmVyKSAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlIHsKCWFjdGl2ZU1hcHBpbmdzIDo9IHByb21ldGhldXMuTmV3R2F1Z2UocHJvbWV0aGV1cy5HYXVnZU9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX21hcHBpbmdzX3RvdGFsIiwKCQlIZWxwOiAiTnVtYmVyIG9mIGFjdGl2ZSB1bmlxdWUgcmVmIG1hcHBpbmdzIGluIHRoZSBzdG9yZS4iLAoJfSkKCXRyYWNrZWRSZWZzIDo9IHByb21ldGhldXMuTmV3R2F1Z2UocHJvbWV0aGV1cy5HYXVnZU9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3RyYWNrZWRfcmVmc190b3RhbCIsCgkJSGVscDogIk51bWJlciBvZiByZWZzIGJlaW5nIHRyYWNrZWQgZm9yIHRpbWVzdGFtcC1iYXNlZCBjbGVhbnVwLiIsCgl9KQoJcmVmc0NsZWFuZWQgOj0gcHJvbWV0aGV1cy5OZXdDb3VudGVyKHByb21ldGhldXMuQ291bnRlck9wdHN7CgkJTmFtZTogImFsbG95X2Zhbm91dF9tYXBwaW5nX3N0b3JlX3JlZnNfY2xlYW5lZF90b3RhbCIsCgkJSGVscDogIlRvdGFsIG51bWJlciBvZiBzdGFsZSByZWZzIGNsZWFuZWQgdXAgb3ZlciB0aW1lLiIsCgl9KQoJdW5pcXVlUmVmc1RvdGFsIDo9IHByb21ldGhldXMuTmV3Q291bnRlcihwcm9tZXRoZXVzLkNvdW50ZXJPcHRzewoJCU5hbWU6ICJhbGxveV9mYW5vdXRfbWFwcGluZ19zdG9yZV91bmlxdWVfcmVmc19jcmVhdGVkX3RvdGFsIiwKCQlIZWxwOiAiVG90YWwgbnVtYmVyIG9mIHVuaXF1ZSByZWZzIGNyZWF0ZWQuIiwKCX0pCgoJaWYgcmVnICE9IG5pbCB7CgkJcmVnLlJlZ2lzdGVyKGFjdGl2ZU1hcHBpbmdzKQoJCXJlZy5SZWdpc3Rlcih0cmFja2VkUmVmcykKCQlyZWcuUmVnaXN0ZXIocmVmc0NsZWFuZWQpCgkJcmVnLlJlZ2lzdGVyKHVuaXF1ZVJlZnNUb3RhbCkKCX0KCglyZXR1cm4gJlNlcmllc1JlZk1hcHBpbmdTdG9yZXsKCQl1bmlxdWVSZWZUb0NoaWxkUmVmczogbWFrZShtYXBbc3RvcmFnZS5TZXJpZXNSZWZddW5pcVJlZkNoaWxkcmVuKSwKCQluZXh0VW5pcXVlUmVmOiAgICAgICAgMSwKCQl1bmlxdWVSZWZUaW1lc3RhbXBzOiAgbWFrZShtYXBbc3RvcmFnZS5TZXJpZXNSZWZdaW50NjQpLAoJCWxhYmVsSGFzaFRvVW5pcXVlUmVmOiBtYWtlKG1hcFt1aW50NjRdc3RvcmFnZS5TZXJpZXNSZWYpLAoJCWNlbGxQb29sOiBzeW5jLlBvb2x7CgkJCU5ldzogZnVuYygpIGFueSB7CgkJCQlyZXR1cm4gJkNlbGx7UmVmczogbWFrZShbXXN0b3JhZ2UuU2VyaWVzUmVmLCAxMDApfQoJCQl9LAoJCX0sCgkJc3RvcENsZWFudXA6ICAgICBtYWtlKGNoYW4gc3RydWN0e30pLAoJCWNsZWFudXBTdG9wcGVkOiAgbWFrZShjaGFuIHN0cnVjdHt9KSwKCQlhY3RpdmVNYXBwaW5nczogIGFjdGl2ZU1hcHBpbmdzLAoJCXRyYWNrZWRSZWZzOiAgICAgdHJhY2tlZFJlZnMsCgkJcmVmc0NsZWFuZWQ6ICAgICByZWZzQ2xlYW5lZCwKCQl1bmlxdWVSZWZzVG90YWw6IHVuaXF1ZVJlZnNUb3RhbCwKCX0KfQoKdHlwZSBDZWxsIHN0cnVjdCB7CglSZWZzIFtdc3RvcmFnZS5TZXJpZXNSZWYKfQoKLy8gR2V0TWFwcGluZyByZXR1cm5zIGV4aXN0aW5nIGNoaWxkIHJlZiByZXN1bHRzIGZvciB0aGUgZ2l2ZW4gdW5pcXVlIHJlZiBpZiBvbmUgZXhpc3RzLgovLwovLyBJZiB0aGUgcGFzc2VkIHVuaXF1ZVJlZiBpcyB6ZXJvLCB0aGUgbWV0aG9kIHdpbGwgYXR0ZW1wdCB0byBmaW5kIGEgbWFwcGluZyB1c2luZyBwYXNzZWQgbGFiZWxzLgovLyBSZXR1cm5zIG5pbCBpZiBubyBtYXBwaW5nIGV4aXN0cy4KLy8KLy8gVGhlIHJldHVybmVkIHNsaWNlIG1heSBiZSBtb2RpZmllZCBieSB0aGUgY2FsbGVyLCBidXQgVXBkYXRlTWFwcGluZyBtdXN0IGJlIGNhbGxlZAovLyBhZnRlcndhcmRzIHRvIHBlcnNpc3QgY2hhbmdlcy4gTm90ZSB0aGF0IGNvbmN1cnJlbnQgYXBwZW5kZXJzIG1heSByYWNlIHRvIHVwZGF0ZSB0aGUKLy8gc2FtZSBtYXBwaW5nIHdpdGggZGlmZmVyZW50IHZhbHVlcywgd2hpY2ggaXMgc2FmZSBiZWNhdXNlIHN0YWxlIG1hcHBpbmdzIGFyZSBzZWxmLWNvcnJlY3RpbmcgLQovLyB1c2luZyBhIHN0YWxlIHJlZiB3aWxsIGNhdXNlIHRoZSBjaGlsZCBhcHBlbmRlciB0byByZXR1cm4gYSBuZXcgcmVmIG9uIHRoZSBuZXh0IGFwcGVuZC4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBHZXRNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSBbXXN0b3JhZ2UuU2VyaWVzUmVmIHsKCXMucmVmTWFwcGluZ011LlJMb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlJVbmxvY2soKQoKCWlmIHVuaXF1ZVJlZiA9PSAwIHsKCQkvLyBTb21lIGNvbnN1bWVycyBkb24ndCBtZW1vIHRoZSBnbG9iYWwgcmVmLiBUcnkgdG8gbG9va3VwIGEgcmVmIGJ5IGxhYmVsIGhhc2guCgkJbGFiZWxIYXNoIDo9IGxibHMuSGFzaCgpCgkJZ290UmVmLCBvayA6PSBzLmxhYmVsSGFzaFRvVW5pcXVlUmVmW2xhYmVsSGFzaF0KCQlpZiAhb2sgewoJCQlyZXR1cm4gbmlsCgkJfQoKCQl1bmlxdWVSZWYgPSBnb3RSZWYKCX0KCglpZiBtYXBwaW5nLCBvayA6PSBzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl07IG9rIHsKCQlyZXR1cm4gKm1hcHBpbmcuY2hpbGRSZWZzCgl9CglyZXR1cm4gbmlsCn0KCi8vIENyZWF0ZU1hcHBpbmcgY3JlYXRlcyBhIG5ldyB1bmlxdWUgcmVmIG1hcHBpbmcgZm9yIHRoZSBnaXZlbiBjaGlsZCByZWYgcmVzdWx0cy4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBDcmVhdGVNYXBwaW5nKHJlZlJlc3VsdHMgW11zdG9yYWdlLlNlcmllc1JlZiwgbGJscyBsYWJlbHMuTGFiZWxzKSBzdG9yYWdlLlNlcmllc1JlZiB7CgkvLyBTdGFydCBjbGVhbnVwIGdvcm91dGluZSBvbiBmaXJzdCBtYXBwaW5nCglzLnN0YXJ0UmVmQ2xlYW51cC5EbyhmdW5jKCkgewoJCXMuY2xlYW51cFN0YXJ0ZWQuU3RvcmUodHJ1ZSkKCQlnbyBzLmNsZWFudXBTdGFsZVJlZnMoKQoJfSkKCgkvLyBTdG9yZSBhIGNvcHkgb2YgdGhlIGNoaWxkIHJlZiByZXN1bHRzIGRpcmVjdGx5CgljaGlsZFJlZlNsaWNlIDo9IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgbGVuKHJlZlJlc3VsdHMpKQoJY29weShjaGlsZFJlZlNsaWNlLCByZWZSZXN1bHRzKQoKCS8vIEhhc2ggbGFiZWxzIHRvIGZvciB0aGUgZmFsbGJhY2sgbG9va3VwIHRhYmxlCglsYWJlbEhhc2ggOj0gbGJscy5IYXNoKCkKCglzLnJlZk1hcHBpbmdNdS5Mb2NrKCkKCWRlZmVyIHMucmVmTWFwcGluZ011LlVubG9jaygpCgoJLy8gQ3JlYXRlIGEgbmV3IHVuaXF1ZSByZWYKCXVuaXF1ZVJlZiA6PSBzLm5leHRVbmlxdWVSZWYKCXMubmV4dFVuaXF1ZVJlZisrCgoJcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltsYWJlbEhhc2hdID0gdW5pcXVlUmVmCglzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzW3VuaXF1ZVJlZl0gPSB1bmlxUmVmQ2hpbGRyZW57CgkJY2hpbGRSZWZzOiAmY2hpbGRSZWZTbGljZSwKCQlsYWJlbEhhc2g6IGxhYmVsSGFzaCwKCX0KCglzLmFjdGl2ZU1hcHBpbmdzLkluYygpCglzLnVuaXF1ZVJlZnNUb3RhbC5JbmMoKQoKCXJldHVybiB1bmlxdWVSZWYKfQoKZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBVcGRhdGVNYXBwaW5nKHVuaXF1ZVJlZiBzdG9yYWdlLlNlcmllc1JlZiwgcmVmUmVzdWx0cyBbXXN0b3JhZ2UuU2VyaWVzUmVmLCBsYmxzIGxhYmVscy5MYWJlbHMpIHsKCWlmIHVuaXF1ZVJlZiA9PSAwIHsKCQlyZXR1cm4KCX0KCgljaGlsZFJlZlNsaWNlIDo9IG1ha2UoW11zdG9yYWdlLlNlcmllc1JlZiwgbGVuKHJlZlJlc3VsdHMpKQoJY29weShjaGlsZFJlZlNsaWNlLCByZWZSZXN1bHRzKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgkvLyBFbnN1cmUgdGhhdCBsYWJlbCBoYXNoIGluZGV4IGlzIHVwIHRvIGRhdGUgdG8gaGFuZGxlIHBvc3NpYmxlIGhhc2ggY29sbGlzaW9ucy4KCS8vIFRPRE86IGlzIHRoaXMgbmVjZXNzYXJ5PwoJbmV3SGFzaCA6PSBsYmxzLkhhc2goKQoJcHJldiwgb2sgOj0gcy51bmlxdWVSZWZUb0NoaWxkUmVmc1t1bmlxdWVSZWZdCglpZiBvayAmJiBwcmV2LmxhYmVsSGFzaCAhPSBuZXdIYXNoIHsKCQlkZWxldGUocy5sYWJlbEhhc2hUb1VuaXF1ZVJlZiwgcHJldi5sYWJlbEhhc2gpCgkJcy5sYWJlbEhhc2hUb1VuaXF1ZVJlZltuZXdIYXNoXSA9IHVuaXF1ZVJlZgoJfQoKCXMudW5pcXVlUmVmVG9DaGlsZFJlZnNbdW5pcXVlUmVmXSA9IHVuaXFSZWZDaGlsZHJlbnsKCQljaGlsZFJlZnM6ICZjaGlsZFJlZlNsaWNlLAoJCWxhYmVsSGFzaDogbGJscy5IYXNoKCksCgl9Cn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgVHJhY2tBcHBlbmRlZFNlcmllcyh0cyBpbnQ2NCwgY2VsbCAqQ2VsbCkgewoJcy50aW1lc3RhbXBUcmFja2luZ011LkxvY2soKQoJZGVmZXIgcy50aW1lc3RhbXBUcmFja2luZ011LlVubG9jaygpCgoJZm9yIF8sIHIgOj0gcmFuZ2UgY2VsbC5SZWZzIHsKCQlzLnVuaXF1ZVJlZlRpbWVzdGFtcHNbcl0gPSB0cwoJfQoKCXMudHJhY2tlZFJlZnMuU2V0KGZsb2F0NjQobGVuKHMudW5pcXVlUmVmVGltZXN0YW1wcykpKQoKCWNlbGwuUmVmcyA9IGNlbGwuUmVmc1s6MF0KCXMuY2VsbFBvb2wuUHV0KGNlbGwpCn0KCmZ1bmMgKHMgKlNlcmllc1JlZk1hcHBpbmdTdG9yZSkgR2V0Q2VsbEZvckFwcGVuZGVkU2VyaWVzKCkgKkNlbGwgewoJcmV0dXJuIHMuY2VsbFBvb2wuR2V0KCkuKCpDZWxsKQp9CgpmdW5jIChzICpTZXJpZXNSZWZNYXBwaW5nU3RvcmUpIGNsZWFudXBTdGFsZVJlZnMoKSB7CglkZWZlciBjbG9zZShzLmNsZWFudXBTdG9wcGVkKQoKCXRpY2tlciA6PSB0aW1lLk5ld1RpY2tlcigxNSAqIHRpbWUuTWludXRlKQoJZGVmZXIgdGlja2VyLlN0b3AoKQoKCWZvciB7CgkJc2VsZWN0IHsKCQljYXNlIDwtdGlja2VyLkM6CgkJCWN1dG9mZlRpbWUgOj0gdGltZS5Ob3coKS5BZGQoLTE1ICogdGltZS5NaW51dGUpLlVuaXgoKQoKCQkJLy8gSG9sZCBib3RoIGxvY2tzIHRvIHByZXZlbnQgcmFjZSBjb25kaXRpb24gd2hlcmUgYSByZWYgY291bGQgYmUKCQkJLy8gYXBwZW5kZWQgYWZ0ZXIgd2UgZGVsZXRlIGl0IGZyb20gdW5pcXVlUmVmQ2VsbCBidXQgYmVmb3JlCgkJCS8vIHdlIGRlbGV0ZSBpdCBmcm9tIHVuaXF1ZVJlZlRvQ2hpbGRSZWZzCgkJCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCQkJcy5yZWZNYXBwaW5nTXUuTG9jaygpCgoJCQlzdGFsZVJlZkNvdW50IDo9IDAKCQkJZm9yIHJlZiwgdHMgOj0gcmFuZ2Ugcy51bmlxdWVSZWZUaW1lc3RhbXBzIHsKCQkJCWlmIHRzIDwgY3V0b2ZmVGltZSB7CgkJCQkJc3RhbGVSZWZDb3VudCsrCgoJCQkJCXYsIG9rIDo9IHMudW5pcXVlUmVmVG9DaGlsZFJlZnNbcmVmXQoJCQkJCWlmIG9rIHsKCQkJCQkJZGVsZXRlKHMubGFiZWxIYXNoVG9VbmlxdWVSZWYsIHYubGFiZWxIYXNoKQoJCQkJCX0KCgkJCQkJZGVsZXRlKHMudW5pcXVlUmVmVGltZXN0YW1wcywgcmVmKQoJCQkJCWRlbGV0ZShzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzLCByZWYpCgkJCQl9CgkJCX0KCgkJCS8vIFVwZGF0ZSBtZXRyaWNzCgkJCWlmIHN0YWxlUmVmQ291bnQgPiAwIHsKCQkJCXMucmVmc0NsZWFuZWQuQWRkKGZsb2F0NjQoc3RhbGVSZWZDb3VudCkpCgkJCQlzLmFjdGl2ZU1hcHBpbmdzLlN1YihmbG9hdDY0KHN0YWxlUmVmQ291bnQpKQoJCQkJcy50cmFja2VkUmVmcy5TZXQoZmxvYXQ2NChsZW4ocy51bmlxdWVSZWZUaW1lc3RhbXBzKSkpCgkJCX0KCgkJCXMucmVmTWFwcGluZ011LlVubG9jaygpCgkJCXMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCQljYXNlIDwtcy5zdG9wQ2xlYW51cDoKCQkJcmV0dXJuCgkJfQoJfQp9CgovLyBDbGVhciB3aWxsIGNsZWFyIGFsbCBpbnRlcm5hbCBtYXBwaW5ncyBhbmQgc3RvcCB0aGUgY2xlYW5lciBnb3JvdXRpbmUgaWYgaXQgaXMgcnVubmluZy4KLy8gSXQgaXMgc2FmZSB0byByZS11c2UgdGhlIHNhbWUgaW5zdGFuY2UgYWZ0ZXIgY2FsbGluZyBDbGVhci4KZnVuYyAocyAqU2VyaWVzUmVmTWFwcGluZ1N0b3JlKSBDbGVhcigpIHsKCS8vIFN0b3AgdGhlIGNsZWFudXAgZ29yb3V0aW5lIGFuZCB3YWl0IGZvciBpdCB0byBiZSBzdG9wcGVkIHNvIHdlIGNhbgoJLy8gYXZvaWQgYSBwb3NzaWJsZSBkZWFkbG9jayB3aXRoIGNsZWFudXAgdGhhdCBhbHNvIGhvbGRzIGJvdGggbG9ja3MKCWlmIHMuY2xlYW51cFN0YXJ0ZWQuTG9hZCgpIHsKCQlzZWxlY3QgewoJCWNhc2UgPC1zLnN0b3BDbGVhbnVwOgoJCQkvLyBBbHJlYWR5IGNsb3NlZAoJCWRlZmF1bHQ6CgkJCWNsb3NlKHMuc3RvcENsZWFudXApCgkJCTwtcy5jbGVhbnVwU3RvcHBlZAoJCX0KCX0KCgkvLyBXZSBuZWVkIHRvIGhvbGQgYm90aCBsb2NrcyB0byBkbyB0aGlzIHNhZmVseSBhbmQgd2UgZG8gaXQgaW4gdGhlIHNhbWUgb3JkZXIgYXMKCS8vIGNsZWFudXBTdGFsZVJlZnMuIFdlIHN0b3BwZWQgYW5kIHdhaXRlZCBmb3IgdGhlIGJhY2tncm91bmQgd29ya2VyIHRoYXQgY2FsbHMgaXQKCS8vIHRvIGZpbmlzaCBidXQgc29tZSBleHRyYSBzYWZldHkgd29uJ3QgaHVydC4KCXMudGltZXN0YW1wVHJhY2tpbmdNdS5Mb2NrKCkKCWRlZmVyIHMudGltZXN0YW1wVHJhY2tpbmdNdS5VbmxvY2soKQoKCXMucmVmTWFwcGluZ011LkxvY2soKQoJZGVmZXIgcy5yZWZNYXBwaW5nTXUuVW5sb2NrKCkKCgljbGVhcihzLnVuaXF1ZVJlZlRvQ2hpbGRSZWZzKQoJY2xlYXIocy51bmlxdWVSZWZUaW1lc3RhbXBzKQoKCS8vIHJlc2V0IHRoZSBwb29sCglzLmNlbGxQb29sID0gc3luYy5Qb29sewoJCU5ldzogZnVuYygpIGFueSB7CgkJCXJldHVybiAmQ2VsbHtSZWZzOiBtYWtlKFtdc3RvcmFnZS5TZXJpZXNSZWYsIDAsIDEwMCl9CgkJfSwKCX0KCgkvLyBOT1RFOiBXZSBkbyBOT1QgcmVzZXQgbmV4dFVuaXF1ZVJlZiBoZXJlLiBSZXNldHRpbmcgaXQgd291bGQgY2F1c2UgcmVmIGNvbGxpc2lvbnMKCS8vIHdpdGggY29tcG9uZW50cyBsaWtlIHByb21ldGhldXMuc2NyYXBlIHdoaWNoIHdpbGwga2VlcCByZS1zZW5kaW5nIHRoZSBzYW1lIGNhY2hlZCByZWZzLgoJLy8gV2UgY29udGludWUgaW5jcmVtZW50aW5nIHRvIGVuc3VyZSBhbGwgcmVmcyByZW1haW4gdW5pcXVlIGFjcm9zcyB0aGUgbGlmZXRpbWUgb2YgdGhlIHByb2Nlc3MuCgoJLy8gUmVzZXQgbWV0cmljcwoJcy5hY3RpdmVNYXBwaW5ncy5TZXQoMCkKCXMudHJhY2tlZFJlZnMuU2V0KDApCgoJLy8gUmVzZXQgY2hhbm5lbHMgYW5kIGZsYWdzCglzLnN0b3BDbGVhbnVwID0gbWFrZShjaGFuIHN0cnVjdHt9KQoJcy5jbGVhbnVwU3RvcHBlZCA9IG1ha2UoY2hhbiBzdHJ1Y3R7fSkKCXMuc3RhcnRSZWZDbGVhbnVwID0gc3luYy5PbmNle30KCXMuY2xlYW51cFN0YXJ0ZWQuU3RvcmUoZmFsc2UpCn0K"
    }
  }
}
```

Response (packages.DriverResponse):

```json
{
  "NotHandled": false,
  "Compiler": "gc",
  "Arch": "amd64",
  "Roots": [
    "github.com/grafana/alloy/internal/service/cluster",
    "github.com/grafana/alloy/internal/component/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus",
    "github.com/grafana/alloy/internal/component/prometheus/enrich",
    "github.com/grafana/alloy/internal/component/prometheus/scrape",
    "github.com/grafana/alloy/internal/component/prometheus/operator",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common",
    "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors",
    "github.com/grafana/alloy/internal/component/prometheus/operator/probes",
    "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs",
    "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http",
    "github.com/grafana/alloy/internal/component/prometheus/relabel",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape",
    "github.com/grafana/alloy/internal/component/all",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert/component",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert",
    "github.com/grafana/alloy/internal/converter",
    "github.com/grafana/alloy/internal/validator",
    "github.com/grafana/alloy/internal/alloycli",
    "github.com/grafana/alloy",
    "github.com/grafana/alloy/internal/cmd/listcomponents",
    "github.com/grafana/alloy/internal/component/metadata",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test",
    "github.com/grafana/alloy/internal/component/prometheus_test",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test",
    "github.com/grafana/alloy/internal/service/cluster_test",
    "github.com/grafana/alloy/internal/tools/docs_generator",
    "github.com/grafana/alloy/internal/tools/docs_generator_test",
    "github.com/grafana/alloy [github.com/grafana/alloy.test]",
    "github.com/grafana/alloy.test",
    "github.com/grafana/alloy/internal/alloycli [github.com/grafana/alloy/internal/alloycli.test]",
    "github.com/grafana/alloy/internal/alloycli.test",
    "github.com/grafana/alloy/internal/component/all [github.com/grafana/alloy/internal/component/all.test]",
    "github.com/grafana/alloy/internal/component/all.test",
    "github.com/grafana/alloy/internal/component/metadata [github.com/grafana/alloy/internal/component/metadata.test]",
    "github.com/grafana/alloy/internal/component/metadata.test",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus_test [github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus.test]",
    "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus.test",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus_test [github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus.test]",
    "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus.test",
    "github.com/grafana/alloy/internal/component/prometheus [github.com/grafana/alloy/internal/component/prometheus.test]",
    "github.com/grafana/alloy/internal/component/prometheus_test [github.com/grafana/alloy/internal/component/prometheus.test]",
    "github.com/grafana/alloy/internal/component/prometheus.test",
    "github.com/grafana/alloy/internal/component/prometheus/enrich [github.com/grafana/alloy/internal/component/prometheus/enrich.test]",
    "github.com/grafana/alloy/internal/component/prometheus/enrich.test",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests_test [github.com/grafana/alloy/internal/component/prometheus/exporter/tests.test]",
    "github.com/grafana/alloy/internal/component/prometheus/exporter/tests.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator [github.com/grafana/alloy/internal/component/prometheus/operator.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common [github.com/grafana/alloy/internal/component/prometheus/operator/common.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator/common.test",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen [github.com/grafana/alloy/internal/component/prometheus/operator/configgen.test]",
    "github.com/grafana/alloy/internal/component/prometheus/operator/configgen.test",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http [github.com/grafana/alloy/internal/component/prometheus/receive_http.test]",
    "github.com/grafana/alloy/internal/component/prometheus/receive_http.test",
    "github.com/grafana/alloy/internal/component/prometheus/relabel [github.com/grafana/alloy/internal/component/prometheus/relabel.test]",
    "github.com/grafana/alloy/internal/component/prometheus/relabel.test",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite [github.com/grafana/alloy/internal/component/prometheus/remotewrite.test]",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite_test [github.com/grafana/alloy/internal/component/prometheus/remotewrite.test]",
    "github.com/grafana/alloy/internal/component/prometheus/remotewrite.test",
    "github.com/grafana/alloy/internal/component/prometheus/scrape [github.com/grafana/alloy/internal/component/prometheus/scrape.test]",
    "github.com/grafana/alloy/internal/component/prometheus/scrape.test",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape [github.com/grafana/alloy/internal/component/pyroscope/scrape.test]",
    "github.com/grafana/alloy/internal/component/pyroscope/scrape.test",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert_test [github.com/grafana/alloy/internal/converter/internal/otelcolconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/otelcolconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert_test [github.com/grafana/alloy/internal/converter/internal/prometheusconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/prometheusconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert_test [github.com/grafana/alloy/internal/converter/internal/promtailconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/promtailconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert_test [github.com/grafana/alloy/internal/converter/internal/staticconvert.test]",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert.test",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build [github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build.test]",
    "github.com/grafana/alloy/internal/converter/internal/staticconvert/internal/build.test",
    "github.com/grafana/alloy/internal/service/cluster [github.com/grafana/alloy/internal/service/cluster.test]",
    "github.com/grafana/alloy/internal/service/cluster_test [github.com/grafana/alloy/internal/service/cluster.test]",
    "github.com/grafana/alloy/internal/service/cluster.test",
    "github.com/grafana/alloy/internal/tools/docs_generator_test [github.com/grafana/alloy/internal/tools/docs_generator.test]",
    "github.com/grafana/alloy/internal/tools/docs_generator.test",
    "github.com/grafana/alloy/internal/validator [github.com/grafana/alloy/internal/validator.test]",
    "github.com/grafana/alloy/internal/validator.test"
  ],
  "Packages": [
    {
      "ID": "archive/tar",
      "Name": "tar",
      "PkgPath": "archive/tar",
      "GoFiles": [
        "/usr/lib/go/src/archive/tar/common.go",
        "/usr/lib/go/src/archive/tar/format.go",
        "/usr/lib/go/src/archive/tar/reader.go",
        "/usr/lib/go/src/archive/tar/stat_actime1.go",
        "/usr/lib/go/src/archive/tar/stat_unix.go",
        "/usr/lib/go/src/archive/tar/strconv.go",
        "/usr/lib/go/src/archive/tar/writer.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/archive/tar/common.go",
        "/usr/lib/go/src/archive/tar/format.go",
        "/usr/lib/go/src/archive/tar/reader.go",
        "/usr/lib/go/src/archive/tar/stat_actime1.go",
        "/usr/lib/go/src/archive/tar/stat_unix.go",
        "/usr/lib/go/src/archive/tar/strconv.go",
        "/usr/lib/go/src/archive/tar/writer.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/archive/tar/stat_actime2.go"],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "fmt": "fmt",
        "internal/godebug": "internal/godebug",
        "io": "io",
        "io/fs": "io/fs",
        "maps": "maps",
        "math": "math",
        "os/user": "os/user",
        "path": "path",
        "path/filepath": "path/filepath",
        "reflect": "reflect",
        "runtime": "runtime",
        "slices": "slices",
        "strconv": "strconv",
        "strings": "strings",
        "sync": "sync",
        "syscall": "syscall",
        "time": "time"
      }
    },
    {
      "ID": "archive/zip",
      "Name": "zip",
      "PkgPath": "archive/zip",
      "GoFiles": [
        "/usr/lib/go/src/archive/zip/reader.go",
        "/usr/lib/go/src/archive/zip/register.go",
        "/usr/lib/go/src/archive/zip/struct.go",
        "/usr/lib/go/src/archive/zip/writer.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/archive/zip/reader.go",
        "/usr/lib/go/src/archive/zip/register.go",
        "/usr/lib/go/src/archive/zip/struct.go",
        "/usr/lib/go/src/archive/zip/writer.go"
      ],
      "Imports": {
        "bufio": "bufio",
        "compress/flate": "compress/flate",
        "encoding/binary": "encoding/binary",
        "errors": "errors",
        "fmt": "fmt",
        "hash": "hash",
        "hash/crc32": "hash/crc32",
        "internal/godebug": "internal/godebug",
        "io": "io",
        "io/fs": "io/fs",
        "os": "os",
        "path": "path",
        "path/filepath": "path/filepath",
        "slices": "slices",
        "strings": "strings",
        "sync": "sync",
        "time": "time",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bufio",
      "Name": "bufio",
      "PkgPath": "bufio",
      "GoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bufio/bufio.go",
        "/usr/lib/go/src/bufio/scan.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "errors": "errors",
        "io": "io",
        "strings": "strings",
        "unicode/utf8": "unicode/utf8"
      }
    },
    {
      "ID": "bytes",
      "Name": "bytes",
      "PkgPath": "bytes",
      "GoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "CompiledGoFiles": [
        "/usr/lib/go/src/bytes/buffer.go",
        "/usr/lib/go/src/bytes/bytes.go",
        "/usr/lib/go/src/bytes/iter.go",
        "/usr/lib/go/src/bytes/reader.go"
      ],
      "IgnoredFiles": ["/usr/lib/go/src/bytes/bytes_js_wasm_test.go"],
      "Imports": {
        "errors": "errors",
        "internal/bytealg": "internal/bytealg",
        "io": "io",
        "iter": "iter",
        "math/bits": "math/bits",
        "unicode": "unicode",
        "unicode/utf8": "unicode/utf8",
        "unsafe": "unsafe"
      }
    },
    {
      "ID": "cloud.google.com/go/auth",
      "Name": "auth",
      "PkgPath": "cloud.google.com/go/auth",
      "GoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/auth.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/threelegged.go"
      ],
      "CompiledGoFiles": [
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/auth.go",
        "/home/username/go/pkg/mod/cloud.google.com/go/auth@v0.17.0/threelegged.go"
      ],
      "Imports": {
        "bytes": "bytes",
        "cloud.google.com/go/auth/internal": "cloud.google.com/go/auth/internal",
        "cloud.google.com/go/auth/internal/jwt": "cloud.google.com/go/auth/internal/jwt",
        "context": "context",
        "encoding/json": "encoding/json",
        "errors": "errors",
        "fmt": "fmt",
        "github.com/googleapis/gax-go/v2/internallog": "github.com/googleapis/gax-go/v2/internallog",
        "log/slog": "log/slog",
        "mime": "mime",
        "net/http": "net/http",
        "net/url": "net/url",
        "strconv": "strconv",
        "strings": "strings",
        "sync": "sync",
        "time": "time"
      }
    }
  ],
  "GoVersion": 25
}
```

Note: `Packages` truncated from 4797 to 5 entries.
