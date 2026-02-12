package driver

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

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
