package driver

import (
	"go/token"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Config is package laoder configuration.
//
// Partially inspired by [golang.org/x/tools/packages.Config] struct.
type Config struct {
	// GoVersion is Go runtime version to be reported in response.
	GoVersion GoVersion

	// Mode controls the level of information returned for each package.
	Mode packages.LoadMode

	// Dir is the directory in which to run the build system's query tool
	// that provides information about the packages.
	// If Dir is empty, the tool is run in the current directory.
	Dir string

	// Env is the environment to use when invoking the build system's query tool.
	Env map[string]string

	// BuildFlags is a list of command-line flags to be passed through to
	// the build system's query tool.
	BuildFlags []string

	// Fset provides source position information for syntax trees and types.
	// If Fset is nil, Load will use a new fileset, but preserve Fset's value.
	Fset *token.FileSet

	// If Tests is set, the loader includes not just the packages
	// matching a particular pattern but also any related test packages,
	// including test-only variants of the package and the test executable.
	//
	// For example, when using the go command, loading "fmt" with Tests=true
	// returns four packages, with IDs "fmt" (the standard package),
	// "fmt [fmt.test]" (the package as compiled for the test),
	// "fmt_test" (the test functions from source files in package fmt_test),
	// and "fmt.test" (the test binary).
	//
	// In build systems with explicit names for tests,
	// setting Tests may have no effect.
	Tests bool

	// Overlay is a mapping from absolute file paths to file contents.
	//
	// For each map entry, [Load] uses the alternative file
	// contents provided by the overlay mapping instead of reading
	// from the file system. This mechanism can be used to enable
	// editor-integrated tools to correctly analyze the contents
	// of modified but unsaved buffers, for example.
	//
	// The overlay mapping is passed to the build system's driver
	// (see "The driver protocol") so that it too can report
	// consistent package metadata about unsaved files. However,
	// drivers may vary in their level of support for overlays.
	Overlay map[string][]byte
}

func (cfg Config) WithDefaults() Config {
	if cfg.Dir == "" {
		wd, err := os.Getwd()
		if err != nil {
			wd = "."
		}

		cfg.Dir = wd
	}

	if cfg.Fset == nil {
		cfg.Fset = token.NewFileSet()
	}

	return cfg
}

// ConfigFromDriverRequest creates a config from Go driver request.
func ConfigFromDriverRequest(ver GoVersion, workDir string, req packages.DriverRequest) Config {
	return Config{
		Mode:       req.Mode,
		Dir:        workDir,
		Env:        MapEnv(req.Env),
		BuildFlags: req.BuildFlags,
		Fset:       token.NewFileSet(),
		Tests:      req.Tests,
		Overlay:    req.Overlay,
	}
}

// MapEnv builds a map from environment variables list.
func MapEnv(env []string) map[string]string {
	out := make(map[string]string, len(env))
	for _, str := range env {
		parts := strings.SplitN(str, "=", 2)
		if len(parts) != 2 || parts[1] == "" {
			continue
		}

		out[parts[0]] = parts[1]
	}

	return out
}
