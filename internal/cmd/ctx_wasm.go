package cmd

import (
	"context"
)

func newApplicationContext() (context.Context, context.CancelFunc) {
	// Stub
	return context.WithCancel(context.Background())
}
