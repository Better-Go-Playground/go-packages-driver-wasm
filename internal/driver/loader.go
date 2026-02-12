// Package driver implements package driver business logic.
package driver

import (
	"context"
	"errors"

	"golang.org/x/tools/go/packages"
)

type Loader struct {
	cfg Config
}

func NewLoader(cfg Config) *Loader {
	return &Loader{
		cfg: cfg.WithDefaults(),
	}
}

func (ldr *Loader) Load(ctx context.Context, patterns []string) (*packages.DriverResponse, error) {
	// TODO: implement
	return nil, errors.New("TODO: not implemented")
}
