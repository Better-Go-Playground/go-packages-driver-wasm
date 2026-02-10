// Generated from ./drivertypes.go

/**
 * Package research contains Go types to read "drvtrace.jsonl" trace file
 */

/**
 * A LoadMode controls the amount of detail to return when loading.
 * The bits below can be combined to specify which fields should be
 * filled in the result packages.
 *
 * The zero value is a special case, equivalent to combining
 * the NeedName, NeedFiles, and NeedCompiledGoFiles bits.
 *
 * ID and Errors (if present) will always be filled.
 * [Load] may return more information than requested.
 *
 * The Mode flag is a union of several bits named NeedName,
 * NeedFiles, and so on, each of which determines whether
 * a given field of Package (Name, Files, etc) should be
 * populated.
 */
export type LoadMode = number;

/**
 * DriverRequest defines the schema of a request for package metadata
 * from an external driver program. The JSON-encoded DriverRequest
 * message is provided to the driver program's standard input. The
 * query patterns are provided as command-line arguments.
 *
 * See the package documentation for an overview.
 */
export interface DriverRequest {
	mode: LoadMode;
	/**
	 * Env specifies the environment the underlying build system should be run in.
	 */
	env: string[];
	/**
	 * BuildFlags are flags that should be passed to the underlying build system.
	 */
	build_flags: string[];
	/**
	 * Tests specifies whether the patterns should also return test packages.
	 */
	tests: boolean;
	/**
	 * Overlay maps file paths (relative to the driver's working directory)
	 * to the contents of overlay files (see Config.Overlay).
	 */
	overlay: Record<string, string>;
}

/**
 * DriverResponse defines the schema of a response from an external
 * driver program, providing the results of a query for package
 * metadata. The driver program must write a JSON-encoded
 * DriverResponse message to its standard output.
 *
 * See the package documentation for an overview.
 */
export interface DriverResponse {
	/**
	 * NotHandled is returned if the request can't be handled by the current
	 * driver. If an external driver returns a response with NotHandled, the
	 * rest of the DriverResponse is ignored, and go/packages will fallback
	 * to the next driver. If go/packages is extended in the future to support
	 * lists of multiple drivers, go/packages will fall back to the next driver.
	 */
	NotHandled: boolean;
	/**
	 * Compiler and Arch are the arguments pass of types.SizesFor
	 * to get a types.Sizes to use when type checking.
	 */
	Compiler: string;
	Arch: string;
	/**
	 * Roots is the set of package IDs that make up the root packages.
	 * We have to encode this separately because when we encode a single package
	 * we cannot know if it is one of the roots as that requires knowledge of the
	 * graph it is part of.
	 */
	Roots?: string[];
	/**
	 * Packages is the full set of packages in the graph.
	 * The packages are not connected into a graph.
	 * The Imports if populated will be stubs that only have their ID set.
	 * Imports will be connected and then type and syntax information added in a
	 * later pass (see refine).
	 */
	Packages: Package[];
	/**
	 * GoVersion is the minor version number used by the driver
	 * (e.g. the go command on the PATH) when selecting .go files.
	 * Zero means unknown.
	 */
	GoVersion: number;
}

/**
 * A Package describes a loaded Go package.
 *
 * It also defines part of the JSON schema of [DriverResponse].
 * See the package documentation for an overview.
 */
export interface Package {
	/**
	 * ID is a unique identifier for a package,
	 * in a syntax provided by the underlying build system.
	 *
	 * Because the syntax varies based on the build system,
	 * clients should treat IDs as opaque and not attempt to
	 * interpret them.
	 */
	ID: string;
	/**
	 * Name is the package name as it appears in the package source code.
	 */
	Name: string;
	/**
	 * PkgPath is the package path as used by the go/types package.
	 */
	PkgPath: string;
	/**
	 * Dir is the directory associated with the package, if it exists.
	 *
	 * For packages listed by the go command, this is the directory containing
	 * the package files.
	 */
	Dir: string;
	/**
	 * Errors contains any errors encountered querying the metadata
	 * of the package, or while parsing or type-checking its files.
	 */
	Errors: Error[];
	/**
	 * TypeErrors contains the subset of errors produced during type checking.
	 */
	TypeErrors: TypeError[];
	/**
	 * GoFiles lists the absolute file paths of the package's Go source files.
	 * It may include files that should not be compiled, for example because
	 * they contain non-matching build tags, are documentary pseudo-files such as
	 * unsafe/unsafe.go or builtin/builtin.go, or are subject to cgo preprocessing.
	 */
	GoFiles: string[];
	/**
	 * CompiledGoFiles lists the absolute file paths of the package's source
	 * files that are suitable for type checking.
	 * This may differ from GoFiles if files are processed before compilation.
	 */
	CompiledGoFiles: string[];
	/**
	 * OtherFiles lists the absolute file paths of the package's non-Go source files,
	 * including assembly, C, C++, Fortran, Objective-C, SWIG, and so on.
	 */
	OtherFiles: string[];
	/**
	 * EmbedFiles lists the absolute file paths of the package's files
	 * embedded with go:embed.
	 */
	EmbedFiles: string[];
	/**
	 * EmbedPatterns lists the absolute file patterns of the package's
	 * files embedded with go:embed.
	 */
	EmbedPatterns: string[];
	/**
	 * IgnoredFiles lists source files that are not part of the package
	 * using the current build configuration but that might be part of
	 * the package using other build configurations.
	 */
	IgnoredFiles: string[];
	/**
	 * ExportFile is the absolute path to a file containing type
	 * information for the package as provided by the build system.
	 */
	ExportFile: string;
	/**
	 * Target is the absolute install path of the .a file, for libraries,
	 * and of the executable file, for binaries.
	 */
	Target: string;
	/**
	 * Imports maps import paths appearing in the package's Go source files
	 * to corresponding loaded Packages.
	 */
	Imports: Record<string, Package>;
	/**
	 * Module is the module information for the package if it exists.
	 *
	 * Note: it may be missing for std and cmd; see Go issue #65816.
	 */
	Module: Module | null;
}

/**
 * An Error describes a problem with a package's metadata, syntax, or types.
 */
export interface Error {
	/** "file:line:col" or "file:line" or "" or "-" */
	Pos: string;
	Msg: string;
	Kind: ErrorKind;
}

/**
 * ErrorKind describes the source of the error, allowing the user to
 * differentiate between errors generated by the driver, the parser, or the
 * type-checker.
 */
export type ErrorKind = number;

/**
 * TypeError describes a type-checking error; it implements the error interface.
 * A "soft" error is an error that still permits a valid interpretation of a
 * package (such as "unused variable"); "hard" errors may lead to unpredictable
 * behavior if ignored.
 */
export interface TypeError {
	/** error position */
	Pos: number;
	/** error message */
	Msg: string;
	/** if set, error is "soft" */
	Soft: boolean;
}

/**
 * Module provides module information for a package.
 *
 * It also defines part of the JSON schema of [DriverResponse].
 * See the package documentation for an overview.
 */
export interface Module {
	/** module path */
	Path: string;
	/** module version */
	Version: string;
	/** replaced by this module */
	Replace: Module | null;
	/** time version was created */
	Time: string | null;
	/** is this the main module? */
	Main: boolean;
	/** is this module only an indirect dependency of main module? */
	Indirect: boolean;
	/** directory holding files for this module, if any */
	Dir: string;
	/** path to go.mod file used when loading this module, if any */
	GoMod: string;
	/** go version used in module */
	GoVersion: string;
	/** error loading module */
	Error: ModuleError | null;
}

/**
 * ModuleError holds errors loading a module.
 */
export interface ModuleError {
	/** the error itself */
	Err: string;
}
