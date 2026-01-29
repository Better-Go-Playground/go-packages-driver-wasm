// Package jsonrpc provides primitives to serve JSON-RPC requests.
package jsonrpc

import (
	"encoding/json"
	"errors"
	"fmt"
)

const NotificationCancelRequest = "$/cancelRequest"

type Request struct {
	ID     int             `json:"id,omitempty"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params,omitempty"`
}

type Response struct {
	ID     int    `json:"id"`
	Result any    `json:"result,omitempty"`
	Error  *Error `json:"error,omitempty"`
}

type ErrorCode int

const (
	ErrorCodeParseError     ErrorCode = -32700
	ErrorCodeInvalidRequest ErrorCode = -32600
	ErrorCodeMethodNotFound ErrorCode = -32601
	ErrorCodeInvalidParams  ErrorCode = -32602
	ErrorCodeInternalError  ErrorCode = -32603
)

func (code ErrorCode) Errorf(format string, args ...any) *Error {
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	return &Error{
		Code:    code,
		Message: msg,
	}
}

func (code ErrorCode) Error(err error) *Error {
	return NewError(code, err)
}

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Data    any       `json:"data,omitempty"`
}

func NewError(code ErrorCode, err error) *Error {
	return &Error{
		Code:    code,
		Message: err.Error(),
	}
}

func (err *Error) AsResponse(reqID int) *Response {
	return &Response{
		ID:    reqID,
		Error: err,
	}
}

// WrapError wraps error as json-rpc error.
//
// If error is [jsonrpc.Error] - returns original value.
func WrapError(err error) *Error {
	e := new(Error)
	if errors.As(err, &e) {
		return e
	}

	return NewError(ErrorCodeInternalError, err)
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s (code: %d)", err.Message, err.Code)
}

