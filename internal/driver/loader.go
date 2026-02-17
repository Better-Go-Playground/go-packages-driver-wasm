// Package driver implements package driver business logic.
package driver

import (
	"context"
	"log"

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

// Load runs the loader pipeline:
// 1) build runtime/config state
// 2) normalize + load requested patterns
// 3) force builtin package inclusion
// 4) build a deterministic DriverResponse
func (ldr *Loader) Load(ctx context.Context, patterns []string) (*packages.DriverResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	rt, err := newLoaderRuntime(ldr.cfg)
	if err != nil {
		return nil, err
	}

	if rt.cfg.Tests && rt.mode&packages.NeedForTest != 0 {
		log.Println("driver: NeedForTest requested; returning best-effort non-test package metadata")
	}

	if unsupported := rt.mode &^ supportedLoadModeMask; unsupported != 0 {
		log.Printf("driver: unsupported LoadMode bits requested: %d", unsupported)
	}

	if len(patterns) == 0 {
		patterns = []string{"."}
	}

	st := newLoaderState(rt)
	normalizedPatterns := normalizePatterns(rt.cfg.Dir, patterns)
	for _, chunk := range chunkPatterns(normalizedPatterns, safeArgMax) {
		if err := st.loadChunk(ctx, chunk); err != nil {
			return nil, err
		}
	}

	if err := st.ensureBuiltin(ctx); err != nil {
		return nil, err
	}

	return st.buildResponse(), nil
}

const (
	safeArgMax       = 32767 - 16384
	builtinPackageID = "builtin"
)

const supportedLoadModeMask = packages.NeedName |
	packages.NeedFiles |
	packages.NeedCompiledGoFiles |
	packages.NeedImports |
	packages.NeedDeps |
	packages.NeedTypesSizes |
	packages.NeedForTest |
	packages.NeedModule |
	packages.NeedEmbedFiles |
	packages.NeedEmbedPatterns |
	packages.NeedTarget
