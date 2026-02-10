package driver

import (
	"context"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/packages"
)

// Config is package load configuration.
//
// Inspired by [golang.org/x/tools/packages.Config] struct.
type Config struct {
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

	// ParseFile is called to read and parse each file
	// when preparing a package's type-checked syntax tree.
	// It must be safe to call ParseFile simultaneously from multiple goroutines.
	// If ParseFile is nil, the loader will uses parser.ParseFile.
	//
	// ParseFile should parse the source from src and use filename only for
	// recording position information.
	//
	// An application may supply a custom implementation of ParseFile
	// to change the effective file contents or the behavior of the parser,
	// or to modify the syntax tree. For example, selectively eliminating
	// unwanted function bodies can significantly accelerate type checking.
	ParseFile func(fset *token.FileSet, filename string, src []byte) (*ast.File, error)

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

func Load(cfg *Config, patterns []string) (*packages.DriverResponse, error) {
	return nil, nil
}
