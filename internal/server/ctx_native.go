//go:build !wasm

package server

import (
	"context"
	"os"
	"os/signal"
)

func NewApplicationContext() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt)
}
