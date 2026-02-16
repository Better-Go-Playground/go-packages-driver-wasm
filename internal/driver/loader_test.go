package driver

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	modmodule "golang.org/x/mod/module"
	"golang.org/x/tools/go/packages"
)

func TestConfigFromDriverRequestKeepsGoVersion(t *testing.T) {
	ver := GoVersion{
		GoMinorVersion: 25,
		Compiler:       "gc",
		Arch:           "amd64",
	}

	cfg := ConfigFromDriverRequest(ver, "/tmp/work", packages.DriverRequest{})
	if cfg.GoVersion != ver {
		t.Fatalf("GoVersion mismatch: got %+v want %+v", cfg.GoVersion, ver)
	}
}

func TestNormalizePatterns(t *testing.T) {
	workDir := filepath.FromSlash("/tmp/workspace")
	got := normalizePatterns(workDir, []string{
		"./...",
		".",
		"file=internal/driver/loader.go",
		"builtin",
		"github.com/acme/project/...",
		"...",
	})

	want := []string{
		filepath.ToSlash(filepath.Clean(workDir)) + "/...",
		filepath.Clean(workDir),
		"file=" + filepath.Clean(filepath.Join(workDir, "internal/driver/loader.go")),
		"builtin",
		"github.com/acme/project/...",
		filepath.ToSlash(filepath.Clean(workDir)) + "/...",
	}

	if len(got) != len(want) {
		t.Fatalf("normalizePatterns length mismatch: got %d want %d (%v)", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("normalizePatterns[%d] mismatch: got %q want %q", i, got[i], want[i])
		}
	}
}

func TestChunkPatterns(t *testing.T) {
	patterns := []string{"aaa", "bbb", "cccc"}
	chunks := chunkPatterns(patterns, 6)

	if len(chunks) != 3 {
		t.Fatalf("chunk count mismatch: got %d want 3", len(chunks))
	}
	for i, chunk := range chunks {
		sum := 0
		for _, p := range chunk {
			sum += len(p) + 1
		}
		if sum > 6 {
			t.Fatalf("chunk %d too large: %d > 6", i, sum)
		}
	}
}

func TestLoaderLoadIncludesBuiltinAndModulePackages(t *testing.T) {
	tmpDir := t.TempDir()
	writeFile(t, filepath.Join(tmpDir, "go.mod"), "module example.com/app\n\ngo 1.25.0\n")
	writeFile(t, filepath.Join(tmpDir, "main.go"), "package app\n\nimport (\n\t\"fmt\"\n\t\"example.com/app/lib\"\n)\n\nfunc Main() {\n\tfmt.Println(lib.Hello())\n}\n")
	writeFile(t, filepath.Join(tmpDir, "lib", "lib.go"), "package lib\n\nfunc Hello() string { return \"ok\" }\n")

	cfg := Config{
		GoVersion: GoVersion{
			GoMinorVersion: 25,
			Compiler:       "gc",
			Arch:           "amd64",
		},
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedModule,
		Dir: tmpDir,
		Env: map[string]string{
			"GOOS":   runtime.GOOS,
			"GOARCH": runtime.GOARCH,
			"GOROOT": runtime.GOROOT(),
		},
	}

	rsp, err := NewLoader(cfg).Load(context.Background(), []string{"./..."})
	if err != nil {
		t.Fatalf("Load failed: %s", err)
	}
	if rsp.NotHandled {
		t.Fatalf("expected NotHandled=false")
	}
	if rsp.GoVersion != 25 {
		t.Fatalf("unexpected GoVersion: got %d want 25", rsp.GoVersion)
	}

	byID := make(map[string]*packages.Package, len(rsp.Packages))
	for _, pkg := range rsp.Packages {
		byID[pkg.ID] = pkg
	}

	rootPkg := byID["example.com/app"]
	if rootPkg == nil {
		t.Fatalf("missing root module package")
	}
	if rootPkg.Module == nil || rootPkg.Module.Path != "example.com/app" {
		t.Fatalf("unexpected module metadata: %+v", rootPkg.Module)
	}
	if !hasKey(rootPkg.Imports, "example.com/app/lib") {
		t.Fatalf("root package imports missing local package: %+v", rootPkg.Imports)
	}
	if !hasKey(byID, "builtin") {
		t.Fatalf("builtin package must always be present")
	}
	if !hasKey(byID, "example.com/app/lib") {
		t.Fatalf("missing nested module package")
	}

	for _, f := range rootPkg.GoFiles {
		if !strings.HasPrefix(f, tmpDir+string(filepath.Separator)) {
			t.Fatalf("root GoFile path is not anchored to workspace dir: %q", f)
		}
	}
}

func TestLoaderOverlayOverridesImports(t *testing.T) {
	tmpDir := t.TempDir()
	writeFile(t, filepath.Join(tmpDir, "go.mod"), "module example.com/overlay\n\ngo 1.25.0\n")
	overlayPath := filepath.Join(tmpDir, "overlay.go")
	writeFile(t, overlayPath, "package overlay\n\nimport \"fmt\"\n\nfunc Name() string { return fmt.Sprint(\"x\") }\n")

	cfg := Config{
		GoVersion: GoVersion{
			GoMinorVersion: 25,
			Compiler:       "gc",
			Arch:           "amd64",
		},
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps,
		Dir: tmpDir,
		Env: map[string]string{
			"GOOS":   runtime.GOOS,
			"GOARCH": runtime.GOARCH,
			"GOROOT": runtime.GOROOT(),
		},
		Overlay: map[string][]byte{
			overlayPath: []byte("package overlay\n\nimport \"strings\"\n\nfunc Name() string { return strings.TrimSpace(\"x\") }\n"),
		},
	}

	rsp, err := NewLoader(cfg).Load(context.Background(), []string{"."})
	if err != nil {
		t.Fatalf("Load failed: %s", err)
	}

	var rootPkg *packages.Package
	for _, pkg := range rsp.Packages {
		if pkg.ID == "example.com/overlay" {
			rootPkg = pkg
			break
		}
	}
	if rootPkg == nil {
		t.Fatalf("overlay package not found in response")
	}

	if !hasKey(rootPkg.Imports, "strings") {
		t.Fatalf("overlay imports were not applied: %+v", rootPkg.Imports)
	}
	if hasKey(rootPkg.Imports, "fmt") {
		t.Fatalf("expected overlay to replace original imports: %+v", rootPkg.Imports)
	}
}

func TestLoaderResolvesExternalImportsFromGoModCache(t *testing.T) {
	tmpDir := t.TempDir()
	workspaceDir := filepath.Join(tmpDir, "workspace")
	modCacheDir := filepath.Join(tmpDir, "modcache")

	writeFile(t, filepath.Join(workspaceDir, "go.mod"), "module example.com/app\n\ngo 1.25.0\n\nrequire example.com/dep/v2 v2.0.1\n")
	writeFile(t, filepath.Join(workspaceDir, "main.go"), "package app\n\nimport _ \"example.com/dep/v2/subpkg\"\n")

	depRoot := filepath.Join(modCacheDir, moduleCachePath(t, "example.com/dep/v2", "v2.0.1"))
	writeFile(t, filepath.Join(depRoot, "go.mod"), "module example.com/dep/v2\n\ngo 1.25.0\n")
	writeFile(t, filepath.Join(depRoot, "subpkg", "subpkg.go"), "package subpkg\n\nfunc Name() string { return \"dep\" }\n")

	cfg := Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps,
		Dir: workspaceDir,
		Env: map[string]string{
			"GOOS":       runtime.GOOS,
			"GOARCH":     runtime.GOARCH,
			"GOROOT":     runtime.GOROOT(),
			"GOMODCACHE": modCacheDir,
		},
	}

	rsp, err := NewLoader(cfg).Load(context.Background(), []string{"."})
	if err != nil {
		t.Fatalf("Load failed: %s", err)
	}

	byID := make(map[string]*packages.Package, len(rsp.Packages))
	for _, pkg := range rsp.Packages {
		byID[pkg.ID] = pkg
	}

	depPkg := byID["example.com/dep/v2/subpkg"]
	if depPkg == nil {
		t.Fatalf("missing package loaded from module cache")
	}
	if len(depPkg.GoFiles) == 0 {
		t.Fatalf("module cache package should have GoFiles")
	}
	if len(depPkg.Errors) > 0 {
		t.Fatalf("unexpected package errors: %+v", depPkg.Errors)
	}
}

func TestLoaderResolvesTransitiveImportsUsingDependencyGoMod(t *testing.T) {
	tmpDir := t.TempDir()
	workspaceDir := filepath.Join(tmpDir, "workspace")
	modCacheDir := filepath.Join(tmpDir, "modcache")

	writeFile(t, filepath.Join(workspaceDir, "go.mod"), "module example.com/app\n\ngo 1.25.0\n\nrequire example.com/dep/v2 v2.0.1\n")
	writeFile(t, filepath.Join(workspaceDir, "main.go"), "package app\n\nimport _ \"example.com/dep/v2/subpkg\"\n")

	depRoot := filepath.Join(modCacheDir, moduleCachePath(t, "example.com/dep/v2", "v2.0.1"))
	writeFile(t, filepath.Join(depRoot, "go.mod"), "module example.com/dep/v2\n\ngo 1.25.0\n\nrequire example.com/leaf v1.4.0\n")
	writeFile(t, filepath.Join(depRoot, "subpkg", "subpkg.go"), "package subpkg\n\nimport _ \"example.com/leaf/pkg\"\n")

	leafRoot := filepath.Join(modCacheDir, moduleCachePath(t, "example.com/leaf", "v1.4.0"))
	writeFile(t, filepath.Join(leafRoot, "go.mod"), "module example.com/leaf\n\ngo 1.25.0\n")
	writeFile(t, filepath.Join(leafRoot, "pkg", "pkg.go"), "package pkg\n\nfunc Value() string { return \"leaf\" }\n")

	cfg := Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps,
		Dir: workspaceDir,
		Env: map[string]string{
			"GOOS":       runtime.GOOS,
			"GOARCH":     runtime.GOARCH,
			"GOROOT":     runtime.GOROOT(),
			"GOMODCACHE": modCacheDir,
		},
	}

	rsp, err := NewLoader(cfg).Load(context.Background(), []string{"."})
	if err != nil {
		t.Fatalf("Load failed: %s", err)
	}

	byID := make(map[string]*packages.Package, len(rsp.Packages))
	for _, pkg := range rsp.Packages {
		byID[pkg.ID] = pkg
	}

	leafPkg := byID["example.com/leaf/pkg"]
	if leafPkg == nil {
		t.Fatalf("missing transitive module package")
	}
	if len(leafPkg.GoFiles) == 0 {
		t.Fatalf("transitive module package should have GoFiles")
	}
	if len(leafPkg.Errors) > 0 {
		t.Fatalf("unexpected transitive package errors: %+v", leafPkg.Errors)
	}
}

func TestLoaderPrefersMainModuleVersionsForDependencies(t *testing.T) {
	tmpDir := t.TempDir()
	workspaceDir := filepath.Join(tmpDir, "workspace")
	modCacheDir := filepath.Join(tmpDir, "modcache")

	writeFile(t, filepath.Join(workspaceDir, "go.mod"), "module example.com/app\n\ngo 1.25.0\n\nrequire (\n\texample.com/common v1.2.0\n\texample.com/dep v1.0.0\n)\n")
	writeFile(t, filepath.Join(workspaceDir, "main.go"), "package app\n\nimport _ \"example.com/dep/subpkg\"\n")

	depRoot := filepath.Join(modCacheDir, moduleCachePath(t, "example.com/dep", "v1.0.0"))
	writeFile(t, filepath.Join(depRoot, "go.mod"), "module example.com/dep\n\ngo 1.25.0\n\nrequire example.com/common v1.0.0\n")
	writeFile(t, filepath.Join(depRoot, "subpkg", "subpkg.go"), "package subpkg\n\nimport _ \"example.com/common/pkg\"\n")

	commonRoot := filepath.Join(modCacheDir, moduleCachePath(t, "example.com/common", "v1.2.0"))
	writeFile(t, filepath.Join(commonRoot, "go.mod"), "module example.com/common\n\ngo 1.25.0\n")
	writeFile(t, filepath.Join(commonRoot, "pkg", "pkg.go"), "package pkg\n\nfunc Value() string { return \"main-version\" }\n")

	cfg := Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps,
		Dir: workspaceDir,
		Env: map[string]string{
			"GOOS":       runtime.GOOS,
			"GOARCH":     runtime.GOARCH,
			"GOROOT":     runtime.GOROOT(),
			"GOMODCACHE": modCacheDir,
		},
	}

	rsp, err := NewLoader(cfg).Load(context.Background(), []string{"."})
	if err != nil {
		t.Fatalf("Load failed: %s", err)
	}

	byID := make(map[string]*packages.Package, len(rsp.Packages))
	for _, pkg := range rsp.Packages {
		byID[pkg.ID] = pkg
	}

	commonPkg := byID["example.com/common/pkg"]
	if commonPkg == nil {
		t.Fatalf("missing common package resolved through main module requirements")
	}
	if len(commonPkg.GoFiles) == 0 {
		t.Fatalf("common package should have GoFiles")
	}
	if !strings.Contains(commonPkg.GoFiles[0], "@v1.2.0") {
		t.Fatalf("expected common package to use main module version v1.2.0, got file %q", commonPkg.GoFiles[0])
	}
	if len(commonPkg.Errors) > 0 {
		t.Fatalf("unexpected common package errors: %+v", commonPkg.Errors)
	}
}

func TestResolveImportCanonicalizesGorootVendorID(t *testing.T) {
	vendorDir := filepath.Join(runtime.GOROOT(), "src", "vendor", "golang.org", "x", "net", "http", "httpguts")
	if !dirExists(vendorDir) {
		t.Skip("GOROOT vendor package not available in current toolchain")
	}

	rt, err := newLoaderRuntime(Config{
		Dir: t.TempDir(),
		Env: map[string]string{
			"GOOS":   runtime.GOOS,
			"GOARCH": runtime.GOARCH,
			"GOROOT": runtime.GOROOT(),
		},
	})
	if err != nil {
		t.Fatalf("newLoaderRuntime failed: %s", err)
	}

	resolvedDir, pkgID, _, err := rt.resolveImport(
		"golang.org/x/net/http/httpguts",
		filepath.Join(runtime.GOROOT(), "src", "net", "http"),
	)
	if err != nil {
		t.Fatalf("resolveImport failed: %s", err)
	}
	if pkgID != "vendor/golang.org/x/net/http/httpguts" {
		t.Fatalf("unexpected vendored package ID: got %q", pkgID)
	}
	if resolvedDir != vendorDir {
		t.Fatalf("unexpected vendored package dir: got %q want %q", resolvedDir, vendorDir)
	}
}

func TestResolveImportCanonicalizesNestedGorootVendorID(t *testing.T) {
	srcDir := filepath.Join(runtime.GOROOT(), "src", "vendor", "golang.org", "x", "net", "idna")
	vendorDir := filepath.Join(runtime.GOROOT(), "src", "vendor", "golang.org", "x", "text", "unicode", "bidi")
	if !dirExists(srcDir) || !dirExists(vendorDir) {
		t.Skip("GOROOT nested vendor package not available in current toolchain")
	}

	rt, err := newLoaderRuntime(Config{
		Dir: t.TempDir(),
		Env: map[string]string{
			"GOOS":   runtime.GOOS,
			"GOARCH": runtime.GOARCH,
			"GOROOT": runtime.GOROOT(),
		},
	})
	if err != nil {
		t.Fatalf("newLoaderRuntime failed: %s", err)
	}

	resolvedDir, pkgID, _, err := rt.resolveImport(
		"golang.org/x/text/unicode/bidi",
		srcDir,
	)
	if err != nil {
		t.Fatalf("resolveImport failed: %s", err)
	}
	if pkgID != "vendor/golang.org/x/text/unicode/bidi" {
		t.Fatalf("unexpected nested vendored package ID: got %q", pkgID)
	}
	if resolvedDir != vendorDir {
		t.Fatalf("unexpected nested vendored package dir: got %q want %q", resolvedDir, vendorDir)
	}
}

func TestLoadByImportPrefersResolvedVendorIDOverCachedImportPath(t *testing.T) {
	srcDir := filepath.Join(runtime.GOROOT(), "src", "vendor", "golang.org", "x", "net", "idna")
	if !dirExists(srcDir) {
		t.Skip("GOROOT nested vendor package not available in current toolchain")
	}

	rt, err := newLoaderRuntime(Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps,
		Dir: t.TempDir(),
		Env: map[string]string{
			"GOOS":   runtime.GOOS,
			"GOARCH": runtime.GOARCH,
			"GOROOT": runtime.GOROOT(),
		},
	})
	if err != nil {
		t.Fatalf("newLoaderRuntime failed: %s", err)
	}

	st := newLoaderState(rt)
	st.packages["golang.org/x/text/unicode/bidi"] = &packages.Package{ID: "golang.org/x/text/unicode/bidi"}

	pkg, err := st.loadByImport(context.Background(), "golang.org/x/text/unicode/bidi", srcDir)
	if err != nil {
		t.Fatalf("loadByImport failed: %s", err)
	}
	if pkg.ID != "vendor/golang.org/x/text/unicode/bidi" {
		t.Fatalf("expected vendor package ID, got %q", pkg.ID)
	}
}

func TestNewLoaderRuntimeRespectsCGOEnabledEnv(t *testing.T) {
	rtEnabled, err := newLoaderRuntime(Config{
		Dir: t.TempDir(),
		Env: map[string]string{
			"GOOS":        runtime.GOOS,
			"GOARCH":      runtime.GOARCH,
			"GOROOT":      runtime.GOROOT(),
			"CGO_ENABLED": "1",
		},
	})
	if err != nil {
		t.Fatalf("newLoaderRuntime with CGO_ENABLED=1 failed: %s", err)
	}
	if !rtEnabled.buildCtx.CgoEnabled {
		t.Fatalf("expected CgoEnabled=true when CGO_ENABLED=1")
	}

	rtDisabled, err := newLoaderRuntime(Config{
		Dir: t.TempDir(),
		Env: map[string]string{
			"GOOS":        runtime.GOOS,
			"GOARCH":      runtime.GOARCH,
			"GOROOT":      runtime.GOROOT(),
			"CGO_ENABLED": "0",
		},
	})
	if err != nil {
		t.Fatalf("newLoaderRuntime with CGO_ENABLED=0 failed: %s", err)
	}
	if rtDisabled.buildCtx.CgoEnabled {
		t.Fatalf("expected CgoEnabled=false when CGO_ENABLED=0")
	}
}

func TestLoaderMapsCgoImportToRuntimeCgo(t *testing.T) {
	runtimeCgoDir := filepath.Join(runtime.GOROOT(), "src", "runtime", "cgo")
	if !dirExists(runtimeCgoDir) {
		t.Skip("runtime/cgo package not available in current toolchain")
	}

	tmpDir := t.TempDir()
	writeFile(t, filepath.Join(tmpDir, "go.mod"), "module example.com/cgomod\n\ngo 1.25.0\n")
	writeFile(t, filepath.Join(tmpDir, "cgofile.go"), "package cgomod\n\nimport \"C\"\n\nfunc Value() int { return 1 }\n")

	cfg := Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps,
		Dir: tmpDir,
		Env: map[string]string{
			"GOOS":        runtime.GOOS,
			"GOARCH":      runtime.GOARCH,
			"GOROOT":      runtime.GOROOT(),
			"CGO_ENABLED": "1",
		},
	}

	rsp, err := NewLoader(cfg).Load(context.Background(), []string{"."})
	if err != nil {
		t.Fatalf("Load failed: %s", err)
	}

	byID := make(map[string]*packages.Package, len(rsp.Packages))
	for _, pkg := range rsp.Packages {
		byID[pkg.ID] = pkg
	}

	rootPkg := byID["example.com/cgomod"]
	if rootPkg == nil {
		t.Fatalf("missing root package")
	}
	if !hasKey(rootPkg.Imports, "runtime/cgo") {
		t.Fatalf("expected runtime/cgo import mapping, got imports: %+v", rootPkg.Imports)
	}
	if !hasKey(byID, "runtime/cgo") {
		t.Fatalf("runtime/cgo package should be loaded")
	}
}

func TestLoaderTestsModeAddsTestVariantRoots(t *testing.T) {
	tmpDir := t.TempDir()
	writeFile(t, filepath.Join(tmpDir, "go.mod"), "module example.com/app\n\ngo 1.25.0\n")
	writeFile(t, filepath.Join(tmpDir, "app.go"), "package app\n\nfunc Value() string { return \"ok\" }\n")
	writeFile(t, filepath.Join(tmpDir, "app_test.go"), "package app\n\nimport \"testing\"\n\nfunc TestValue(t *testing.T) {\n\tif Value() == \"\" {\n\t\tt.Fatal(\"empty\")\n\t}\n}\n")

	cfg := Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedForTest,
		Dir:   tmpDir,
		Tests: true,
		Env: map[string]string{
			"GOOS":   runtime.GOOS,
			"GOARCH": runtime.GOARCH,
			"GOROOT": runtime.GOROOT(),
		},
	}

	rsp, err := NewLoader(cfg).Load(context.Background(), []string{"."})
	if err != nil {
		t.Fatalf("Load failed: %s", err)
	}

	rootSet := make(map[string]struct{}, len(rsp.Roots))
	byID := make(map[string]*packages.Package, len(rsp.Packages))
	for _, rootID := range rsp.Roots {
		rootSet[rootID] = struct{}{}
	}
	for _, pkg := range rsp.Packages {
		byID[pkg.ID] = pkg
	}

	variantID := "example.com/app [example.com/app.test]"
	testMainID := "example.com/app.test"

	if !hasKey(rootSet, variantID) {
		t.Fatalf("missing test variant root: %q", variantID)
	}
	if !hasKey(rootSet, testMainID) {
		t.Fatalf("missing test main root: %q", testMainID)
	}

	variantPkg := byID[variantID]
	if variantPkg == nil {
		t.Fatalf("missing test variant package")
	}
	if variantPkg.PkgPath != "example.com/app" {
		t.Fatalf("unexpected test variant PkgPath: got %q", variantPkg.PkgPath)
	}
	if len(variantPkg.GoFiles) != 2 {
		t.Fatalf("unexpected test variant GoFiles count: got %d want 2", len(variantPkg.GoFiles))
	}

	testMainPkg := byID[testMainID]
	if testMainPkg == nil {
		t.Fatalf("missing test main package")
	}
	if !hasKey(testMainPkg.Imports, "example.com/app") {
		t.Fatalf("test main should import root package path")
	}
	if testMainPkg.Imports["example.com/app"].ID != variantID {
		t.Fatalf("test main root import should point to test variant: got %q", testMainPkg.Imports["example.com/app"].ID)
	}
	if !hasKey(testMainPkg.Imports, "testing/internal/testdeps") {
		t.Fatalf("test main missing testing/internal/testdeps import")
	}
	if len(testMainPkg.Errors) > 0 {
		t.Fatalf("unexpected test main errors: %+v", testMainPkg.Errors)
	}
}

func writeFile(t *testing.T, filePath, content string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil {
		t.Fatalf("mkdir failed for %q: %s", filePath, err)
	}
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		t.Fatalf("write failed for %q: %s", filePath, err)
	}
}

func hasKey[V any](m map[string]V, key string) bool {
	_, ok := m[key]
	return ok
}

func moduleCachePath(t *testing.T, modulePath, version string) string {
	t.Helper()

	escapedPath, err := modmodule.EscapePath(modulePath)
	if err != nil {
		t.Fatalf("EscapePath failed: %s", err)
	}

	escapedVersion, err := modmodule.EscapeVersion(version)
	if err != nil {
		t.Fatalf("EscapeVersion failed: %s", err)
	}

	return escapedPath + "@" + escapedVersion
}
