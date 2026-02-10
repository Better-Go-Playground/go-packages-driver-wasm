// Package research contains Go types to read "drvtrace.jsonl" trace file
package research

import "time"

// A LoadMode controls the amount of detail to return when loading.
// The bits below can be combined to specify which fields should be
// filled in the result packages.
//
// The zero value is a special case, equivalent to combining
// the NeedName, NeedFiles, and NeedCompiledGoFiles bits.
//
// ID and Errors (if present) will always be filled.
// [Load] may return more information than requested.
//
// The Mode flag is a union of several bits named NeedName,
// NeedFiles, and so on, each of which determines whether
// a given field of Package (Name, Files, etc) should be
// populated.
//
// For convenience, we provide named constants for the most
// common combinations of Need flags:
//
//	[LoadFiles]     lists of files in each package
//	[LoadImports]   ... plus imports
//	[LoadTypes]     ... plus type information
//	[LoadSyntax]    ... plus type-annotated syntax
//	[LoadAllSyntax] ... for all dependencies
type LoadMode int

const (
	// NeedName adds Name and PkgPath.
	NeedName LoadMode = 1 << iota

	// NeedFiles adds Dir, GoFiles, OtherFiles, and IgnoredFiles
	NeedFiles

	// NeedCompiledGoFiles adds CompiledGoFiles.
	NeedCompiledGoFiles

	// NeedImports adds Imports. If NeedDeps is not set, the Imports field will contain
	// "placeholder" Packages with only the ID set.
	NeedImports

	// NeedDeps adds the fields requested by the LoadMode in the packages in Imports.
	NeedDeps

	// NeedExportFile adds ExportFile.
	NeedExportFile

	// NeedTypes adds Types, Fset, and IllTyped.
	NeedTypes

	// NeedSyntax adds Syntax and Fset.
	NeedSyntax

	// NeedTypesInfo adds TypesInfo and Fset.
	NeedTypesInfo

	// NeedTypesSizes adds TypesSizes.
	NeedTypesSizes

	// needInternalDepsErrors adds the internal deps errors field for use by gopls.
	needInternalDepsErrors

	// NeedForTest adds ForTest.
	//
	// Tests must also be set on the context for this field to be populated.
	NeedForTest

	// typecheckCgo enables full support for type checking cgo. Requires Go 1.15+.
	// Modifies CompiledGoFiles and Types, and has no effect on its own.
	typecheckCgo

	// NeedModule adds Module.
	NeedModule

	// NeedEmbedFiles adds EmbedFiles.
	NeedEmbedFiles

	// NeedEmbedPatterns adds EmbedPatterns.
	NeedEmbedPatterns

	// NeedTarget adds Target.
	NeedTarget

	// Be sure to update loadmode_string.go when adding new items!
)

const (
	// LoadFiles loads the name and file names for the initial packages.
	LoadFiles = NeedName | NeedFiles | NeedCompiledGoFiles

	// LoadImports loads the name, file names, and import mapping for the initial packages.
	LoadImports = LoadFiles | NeedImports

	// LoadTypes loads exported type information for the initial packages.
	LoadTypes = LoadImports | NeedTypes | NeedTypesSizes

	// LoadSyntax loads typed syntax for the initial packages.
	LoadSyntax = LoadTypes | NeedSyntax | NeedTypesInfo

	// LoadAllSyntax loads typed syntax for the initial packages and all dependencies.
	LoadAllSyntax = LoadSyntax | NeedDeps

	// Deprecated: NeedExportsFile is a historical misspelling of NeedExportFile.
	//
	//go:fix inline
	NeedExportsFile = NeedExportFile
)

// DriverRequest defines the schema of a request for package metadata
// from an external driver program. The JSON-encoded DriverRequest
// message is provided to the driver program's standard input. The
// query patterns are provided as command-line arguments.
//
// See the package documentation for an overview.
type DriverRequest struct {
	Mode LoadMode `json:"mode"`

	// Env specifies the environment the underlying build system should be run in.
	Env []string `json:"env"`

	// BuildFlags are flags that should be passed to the underlying build system.
	BuildFlags []string `json:"build_flags"`

	// Tests specifies whether the patterns should also return test packages.
	Tests bool `json:"tests"`

	// Overlay maps file paths (relative to the driver's working directory)
	// to the contents of overlay files (see Config.Overlay).
	Overlay map[string][]byte `json:"overlay"`
}

// DriverResponse defines the schema of a response from an external
// driver program, providing the results of a query for package
// metadata. The driver program must write a JSON-encoded
// DriverResponse message to its standard output.
//
// See the package documentation for an overview.
type DriverResponse struct {
	// NotHandled is returned if the request can't be handled by the current
	// driver. If an external driver returns a response with NotHandled, the
	// rest of the DriverResponse is ignored, and go/packages will fallback
	// to the next driver. If go/packages is extended in the future to support
	// lists of multiple drivers, go/packages will fall back to the next driver.
	NotHandled bool

	// Compiler and Arch are the arguments pass of types.SizesFor
	// to get a types.Sizes to use when type checking.
	Compiler string
	Arch     string

	// Roots is the set of package IDs that make up the root packages.
	// We have to encode this separately because when we encode a single package
	// we cannot know if it is one of the roots as that requires knowledge of the
	// graph it is part of.
	Roots []string `json:",omitempty"`

	// Packages is the full set of packages in the graph.
	// The packages are not connected into a graph.
	// The Imports if populated will be stubs that only have their ID set.
	// Imports will be connected and then type and syntax information added in a
	// later pass (see refine).
	Packages []*Package

	// GoVersion is the minor version number used by the driver
	// (e.g. the go command on the PATH) when selecting .go files.
	// Zero means unknown.
	GoVersion int
}

// A Package describes a loaded Go package.
//
// It also defines part of the JSON schema of [DriverResponse].
// See the package documentation for an overview.
type Package struct {
	// ID is a unique identifier for a package,
	// in a syntax provided by the underlying build system.
	//
	// Because the syntax varies based on the build system,
	// clients should treat IDs as opaque and not attempt to
	// interpret them.
	ID string

	// Name is the package name as it appears in the package source code.
	Name string

	// PkgPath is the package path as used by the go/types package.
	PkgPath string

	// Dir is the directory associated with the package, if it exists.
	//
	// For packages listed by the go command, this is the directory containing
	// the package files.
	Dir string

	// Errors contains any errors encountered querying the metadata
	// of the package, or while parsing or type-checking its files.
	Errors []Error

	// TypeErrors contains the subset of errors produced during type checking.
	TypeErrors []TypeError

	// GoFiles lists the absolute file paths of the package's Go source files.
	// It may include files that should not be compiled, for example because
	// they contain non-matching build tags, are documentary pseudo-files such as
	// unsafe/unsafe.go or builtin/builtin.go, or are subject to cgo preprocessing.
	GoFiles []string

	// CompiledGoFiles lists the absolute file paths of the package's source
	// files that are suitable for type checking.
	// This may differ from GoFiles if files are processed before compilation.
	CompiledGoFiles []string

	// OtherFiles lists the absolute file paths of the package's non-Go source files,
	// including assembly, C, C++, Fortran, Objective-C, SWIG, and so on.
	OtherFiles []string

	// EmbedFiles lists the absolute file paths of the package's files
	// embedded with go:embed.
	EmbedFiles []string

	// EmbedPatterns lists the absolute file patterns of the package's
	// files embedded with go:embed.
	EmbedPatterns []string

	// IgnoredFiles lists source files that are not part of the package
	// using the current build configuration but that might be part of
	// the package using other build configurations.
	IgnoredFiles []string

	// ExportFile is the absolute path to a file containing type
	// information for the package as provided by the build system.
	ExportFile string

	// Target is the absolute install path of the .a file, for libraries,
	// and of the executable file, for binaries.
	Target string

	// Imports maps import paths appearing in the package's Go source files
	// to corresponding loaded Packages.
	Imports map[string]*Package

	// Module is the module information for the package if it exists.
	//
	// Note: it may be missing for std and cmd; see Go issue #65816.
	Module *Module
}

// An Error describes a problem with a package's metadata, syntax, or types.
type Error struct {
	Pos  string // "file:line:col" or "file:line" or "" or "-"
	Msg  string
	Kind ErrorKind
}

// ErrorKind describes the source of the error, allowing the user to
// differentiate between errors generated by the driver, the parser, or the
// type-checker.
type ErrorKind int

const (
	UnknownError ErrorKind = iota
	ListError
	ParseError
	TypeError
)

// TypeError describes a type-checking error; it implements the error interface.
// A "soft" error is an error that still permits a valid interpretation of a
// package (such as "unused variable"); "hard" errors may lead to unpredictable
// behavior if ignored.
type TypeError struct {
	Pos  int    // error position
	Msg  string // error message
	Soft bool   // if set, error is "soft"
}

// Module provides module information for a package.
//
// It also defines part of the JSON schema of [DriverResponse].
// See the package documentation for an overview.
type Module struct {
	Path      string       // module path
	Version   string       // module version
	Replace   *Module      // replaced by this module
	Time      *time.Time   // time version was created
	Main      bool         // is this the main module?
	Indirect  bool         // is this module only an indirect dependency of main module?
	Dir       string       // directory holding files for this module, if any
	GoMod     string       // path to go.mod file used when loading this module, if any
	GoVersion string       // go version used in module
	Error     *ModuleError // error loading module
}

// ModuleError holds errors loading a module.
type ModuleError struct {
	Err string // the error itself
}
