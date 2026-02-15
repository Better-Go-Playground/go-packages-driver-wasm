package jsonrpc

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type RequestHandler interface {
	HandleRequest(ctx context.Context, params json.RawMessage) (any, error)
}

type typedHandler[TRequest, TResponse any] struct {
	handleFunc func(ctx context.Context, req TRequest) (*TResponse, error)
}

func (h *typedHandler[TRequest, TResponse]) HandleRequest(ctx context.Context, params json.RawMessage) (any, error) {
	var req TRequest
	if err := json.Unmarshal(params, &req); err != nil {
		return nil, NewError(ErrorCodeParseError, err)
	}

	return h.handleFunc(ctx, req)
}

// NewHandler constructs a new type-safe request handler.
//
// Incoming request is automatically parsed into a [TRequest] type.
func NewHandler[TRequest, TResponse any](fn func(ctx context.Context, req TRequest) (*TResponse, error)) RequestHandler {
	return &typedHandler[TRequest, TResponse]{
		handleFunc: fn,
	}
}

type requestCanceler struct {
	lock sync.Mutex
	reqs map[int]context.CancelFunc
}

func (rc *requestCanceler) cancelRequest(reqID int) {
	rc.lock.Lock()
	defer rc.lock.Unlock()

	cancelFn, ok := rc.reqs[reqID]
	if ok {
		delete(rc.reqs, reqID)
		cancelFn()
	}
}

func (rc *requestCanceler) finishRequest(reqID int) bool {
	rc.lock.Lock()
	defer rc.lock.Unlock()

	_, ok := rc.reqs[reqID]
	if !ok {
		return false
	}

	delete(rc.reqs, reqID)
	return ok
}

func (rc *requestCanceler) cancelAll() {
	rc.lock.Lock()
	defer rc.lock.Unlock()

	for _, cancelFn := range rc.reqs {
		cancelFn()
	}

	rc.reqs = make(map[int]context.CancelFunc)
}

func (rc *requestCanceler) addRequest(reqID int, cancelFn context.CancelFunc) {
	rc.lock.Lock()
	defer rc.lock.Unlock()

	rc.reqs[reqID] = cancelFn
}

type ServeMux struct {
	canceler requestCanceler
	handlers map[string]RequestHandler
}

func NewServeMux(handlers map[string]RequestHandler) *ServeMux {
	return &ServeMux{
		canceler: requestCanceler{
			reqs: make(map[int]context.CancelFunc),
		},
		handlers: handlers,
	}
}

// ServeStream handles incoming requests from a given connection.
func (l *ServeMux) ServeStream(ctx context.Context, conn net.Conn) error {
	connCtx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()
	defer l.canceler.cancelAll()

	go func() {
		<-ctx.Done()
		_ = conn.Close()
	}()

	reader := bufio.NewReader(conn)

	// Requests are delimited by \n character.
	for {
		data, err := reader.ReadBytes('\n')
		if len(data) > 0 {
			trimmed := bytes.TrimSpace(data)
			if len(trimmed) == 0 {
				if err != nil {
					log.Printf("empty request payload %q", data)
				}
				continue
			}
			if err := l.handleRequest(connCtx, conn, trimmed); err != nil {
				log.Printf("failed to handle request: %s", err)
			}
		}

		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed) {
				break
			}
			if connCtx.Err() != nil {
				break
			}

			return fmt.Errorf("connection read failed: %w", err)
		}
	}

	return nil
}

func (l *ServeMux) handleRequest(ctx context.Context, w io.Writer, data []byte) error {
	var req Request
	if err := json.Unmarshal(data, &req); err != nil {
		return l.serveError(w, 0, NewError(ErrorCodeParseError, err))
	}

	if req.ID == 0 {
		err := l.handleNotification(&req)
		if err != nil {
			return l.serveResponse(w, WrapError(err).AsResponse(0))
		}

		return nil
	}

	handler, ok := l.handlers[req.Method]
	if !ok {
		err := ErrorCodeMethodNotFound.Errorf("method not found: %q", req.Method)
		return l.serveResponse(w, err.AsResponse(req.ID))
	}

	reqCtx, cancelFn := context.WithCancel(ctx)
	l.canceler.addRequest(req.ID, cancelFn)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %s", r)
				_ = l.serveError(w, req.ID, ErrorCodeInternalError.Errorf("%s", r))
			}
		}()

		defer cancelFn()
		defer l.canceler.finishRequest(req.ID)

		rsp := &Response{
			ID: req.ID,
		}

		out, err := handler.HandleRequest(reqCtx, req.Params)
		if err != nil {
			rsp.Error = WrapError(err)
		}

		rsp.Result = out
		if ctx.Err() != nil {
			return
		}

		if err := l.serveResponse(w, rsp); err != nil {
			log.Printf(
				"failed to respond: %s (reqID=%v method=%q)",
				err, req.ID, req.Method,
			)
		}
	}()

	return nil
}

func (l *ServeMux) handleNotification(req *Request) error {
	if req.Method != NotificationCancelRequest {
		return ErrorCodeMethodNotFound.Errorf(
			"unsupported notification %q", req.Method,
		)
	}

	var reqID int
	if err := json.Unmarshal(req.Params, &reqID); err != nil {
		return ErrorCodeInvalidParams.Errorf(
			"cannot read params: %s", err,
		)
	}

	if reqID == 0 {
		return ErrorCodeInvalidParams.Errorf(
			"missing request ID",
		)
	}

	l.canceler.cancelRequest(reqID)
	return nil
}

func (l *ServeMux) serveError(dst io.Writer, reqID int, e *Error) error {
	return l.serveResponse(dst, e.AsResponse(reqID))
}

func (l *ServeMux) serveResponse(dst io.Writer, rsp *Response) error {
	buff := bytes.NewBuffer(make([]byte, 1024))

	// NOTE: responses should be delimited by LF (\n).
	if err := json.NewEncoder(buff).Encode(rsp); err != nil {
		return fmt.Errorf("failed to serialize response: %w", err)
	}

	_, err := dst.Write(buff.Bytes())
	return err
}
