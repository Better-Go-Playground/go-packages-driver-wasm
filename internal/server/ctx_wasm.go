package server

import (
	"context"
)

func NewApplicationContext() (context.Context, context.CancelFunc) {
	// Stub
	return context.WithCancel(context.Background())
}
