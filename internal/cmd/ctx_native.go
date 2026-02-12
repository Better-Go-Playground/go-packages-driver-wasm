//go:build !wasm

package cmd

import (
	"context"
	"os"
	"os/signal"
)

func newApplicationContext() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt)
}
