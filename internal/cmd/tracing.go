package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Better-Go-Playground/go-packages-driver-wasm/internal/jsonrpc"
)

var _ jsonrpc.Interceptor = (*traceInterceptor)(nil)

type traceEvent struct {
	Request  jsonrpc.Request   `json:"request"`
	Response *jsonrpc.Response `json:"response"`
}

type traceInterceptor struct {
	dst *os.File
}

func newTraceInterceptor(dstFile string) (*traceInterceptor, error) {
	// Open a file at dstFile for write.
	f, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o644)
	if err != nil {
		return nil, fmt.Errorf("failed to open tracing file %q for write: %w", dstFile, err)
	}

	return &traceInterceptor{
		dst: f,
	}, nil
}

// InterceptRequest implements [jsonrpc.Interceptor].
func (t *traceInterceptor) InterceptRequest(ctx context.Context, req jsonrpc.Request, next jsonrpc.Interceptor) *jsonrpc.Response {
	log.Printf("trace.request: m=%q id=%d", req.Method, req.ID)
	rsp := next.InterceptRequest(ctx, req, nil)

	if rsp.Error != nil {
		log.Printf("trace.response.error: %q (m=%q id=%d)", rsp.Error.Message, req.Method, req.ID)
	} else {
		log.Printf("trace.response.ok: m=%q id=%d", req.Method, req.ID)
	}

	err := json.NewEncoder(t.dst).Encode(traceEvent{
		Request:  req,
		Response: rsp,
	})
	if err != nil {
		log.Printf("failed to write trace event: %s", err)
	}

	return rsp
}

func (t *traceInterceptor) Close() {
	_ = t.dst.Sync()
	_ = t.dst.Close()
}
