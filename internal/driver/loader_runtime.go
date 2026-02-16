package driver

import (
	"go/build"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/tools/go/packages"
)

// loader_runtime.go contains runtime/env initialization and mode helpers.

type moduleInfo struct {
	Path      string
	Dir       string
	GoModPath string
	GoVersion string
	Main      bool
}

func (mi *moduleInfo) toPackageModule() *packages.Module {
	if mi == nil {
		return nil
	}

	return &packages.Module{
		Path:      mi.Path,
		Main:      mi.Main,
		Dir:       mi.Dir,
		GoMod:     mi.GoModPath,
		GoVersion: mi.GoVersion,
	}
}

type resolvedImport struct {
	dir   string
	pkgID string
	mod   *packages.Module
}

type goModSummary struct {
	modulePath string
	goVersion  string
	requires   map[string]string
}

type loaderRuntime struct {
	cfg Config

	mode     packages.LoadMode
	buildCtx build.Context

	goroot string
	gopath []string

	overlay    map[string][]byte
	overlayDir map[string][]string

	moduleCache map[string]*moduleInfo
	moduleReqs  map[string]map[string]string
	moduleByRef map[string]*moduleInfo
	goModCache  []string
	mainModule  *moduleInfo
}

// newLoaderRuntime builds immutable runtime inputs used by the loader state
// machine (toolchain env, overlay index, and module caches).
func newLoaderRuntime(cfg Config) (*loaderRuntime, error) {
	mode := cfg.Mode
	if mode == 0 {
		mode = packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles
	}

	goos := cfg.Env["GOOS"]
	if goos == "" {
		goos = runtime.GOOS
	}

	goarch := cfg.Env["GOARCH"]
	if goarch == "" {
		goarch = runtime.GOARCH
	}

	goroot := cfg.Env["GOROOT"]
	if goroot == "" {
		goroot = runtime.GOROOT()
	}
	goroot = filepath.Clean(goroot)

	gopath := splitPathList(cfg.Env["GOPATH"])
	if len(gopath) == 0 {
		gopath = splitPathList(os.Getenv("GOPATH"))
	}

	buildCtx := build.Default
	buildCtx.GOOS = goos
	buildCtx.GOARCH = goarch
	buildCtx.GOROOT = goroot
	if len(gopath) > 0 {
		buildCtx.GOPATH = strings.Join(gopath, string(os.PathListSeparator))
	}
	if cgoEnabled, ok := cfg.Env["CGO_ENABLED"]; ok {
		buildCtx.CgoEnabled = cgoEnabled != "0"
	}

	rt := &loaderRuntime{
		cfg:         cfg,
		mode:        mode,
		buildCtx:    buildCtx,
		goroot:      goroot,
		gopath:      gopath,
		overlay:     make(map[string][]byte, len(cfg.Overlay)),
		overlayDir:  make(map[string][]string, len(cfg.Overlay)),
		moduleCache: make(map[string]*moduleInfo),
		moduleReqs:  make(map[string]map[string]string),
		moduleByRef: make(map[string]*moduleInfo),
		goModCache:  resolveGoModCacheDirs(cfg.Env, gopath),
	}

	for inPath, data := range cfg.Overlay {
		abs := normalizeAbsolutePath(cfg.Dir, inPath)
		rt.overlay[abs] = data
		dir := filepath.Clean(filepath.Dir(abs))
		rt.overlayDir[dir] = append(rt.overlayDir[dir], filepath.Base(abs))
	}

	for dir, names := range rt.overlayDir {
		rt.overlayDir[dir] = uniqueStrings(names)
	}

	mainModule, err := rt.moduleForDir(cfg.Dir)
	if err != nil {
		log.Printf("driver: failed to parse go.mod from %q: %s", cfg.Dir, err)
	}
	rt.mainModule = mainModule
	if rt.mainModule != nil {
		rt.mainModule.Main = true
	}

	return rt, nil
}

func (rt *loaderRuntime) wantsName() bool {
	return rt.mode&packages.NeedName != 0
}

func (rt *loaderRuntime) wantsFiles() bool {
	return rt.mode&packages.NeedFiles != 0
}

func (rt *loaderRuntime) wantsCompiledGoFiles() bool {
	return rt.mode&packages.NeedCompiledGoFiles != 0
}

func (rt *loaderRuntime) wantsImports() bool {
	return rt.mode&packages.NeedImports != 0
}

func (rt *loaderRuntime) wantsDeps() bool {
	return rt.mode&packages.NeedDeps != 0
}

func (rt *loaderRuntime) wantsModule() bool {
	return rt.mode&packages.NeedModule != 0
}
