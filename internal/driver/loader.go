// Package driver implements package driver business logic.
package driver

import (
	"context"
	"errors"
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/mod/modfile"
	modmodule "golang.org/x/mod/module"
	"golang.org/x/mod/semver"
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
	safeArgMax = 32767 - 16384
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

type loaderState struct {
	rt *loaderRuntime

	packages map[string]*packages.Package
	loading  map[string]struct{}
	roots    map[string]struct{}
}

func newLoaderState(rt *loaderRuntime) *loaderState {
	return &loaderState{
		rt:       rt,
		packages: make(map[string]*packages.Package),
		loading:  make(map[string]struct{}),
		roots:    make(map[string]struct{}),
	}
}

func (st *loaderState) loadChunk(ctx context.Context, patterns []string) error {
	for _, pattern := range patterns {
		if err := ctx.Err(); err != nil {
			return err
		}

		rootIDs := st.loadPattern(ctx, pattern)
		for _, id := range rootIDs {
			st.roots[id] = struct{}{}
			for _, testRootID := range st.loadTestVariants(ctx, id) {
				st.roots[testRootID] = struct{}{}
			}
		}
	}

	return nil
}

func (st *loaderState) loadPattern(ctx context.Context, pattern string) []string {
	if pattern == "builtin" {
		pkg, err := st.loadByImport(ctx, "builtin", st.rt.cfg.Dir)
		if err != nil {
			return []string{st.addErrorPackage("builtin", err).ID}
		}
		return []string{pkg.ID}
	}

	if strings.HasPrefix(pattern, "file=") {
		fileName := normalizeAbsolutePath(st.rt.cfg.Dir, strings.TrimPrefix(pattern, "file="))
		pkg, err := st.loadByDir(ctx, filepath.Dir(fileName))
		if err != nil {
			return []string{st.addErrorPackage(pattern, err).ID}
		}
		return []string{pkg.ID}
	}

	if isAbsoluteRecursivePattern(pattern) {
		baseDir := strings.TrimSuffix(pattern, "/...")
		rootIDs, err := st.loadRecursivePattern(ctx, baseDir)
		if err != nil {
			return []string{st.addErrorPackage(pattern, err).ID}
		}
		if len(rootIDs) == 0 {
			return []string{st.addErrorPackage(pattern, errors.New("no packages matched pattern")).ID}
		}
		return rootIDs
	}

	if strings.HasSuffix(pattern, "/...") {
		importPrefix := strings.TrimSuffix(pattern, "/...")
		rootIDs, err := st.loadRecursiveImportPattern(ctx, importPrefix)
		if err != nil {
			return []string{st.addErrorPackage(pattern, err).ID}
		}
		if len(rootIDs) == 0 {
			return []string{st.addErrorPackage(pattern, errors.New("no packages matched pattern")).ID}
		}
		return rootIDs
	}

	if filepath.IsAbs(pattern) {
		target := pattern
		if strings.HasSuffix(target, ".go") {
			target = filepath.Dir(target)
		}

		pkg, err := st.loadByDir(ctx, target)
		if err != nil {
			return []string{st.addErrorPackage(pattern, err).ID}
		}
		return []string{pkg.ID}
	}

	pkg, err := st.loadByImport(ctx, pattern, st.rt.cfg.Dir)
	if err != nil {
		return []string{st.addErrorPackage(pattern, err).ID}
	}
	return []string{pkg.ID}
}

func (st *loaderState) loadRecursivePattern(ctx context.Context, baseDir string) ([]string, error) {
	baseDir = filepath.Clean(baseDir)
	if !dirExists(baseDir) {
		return nil, fmt.Errorf("directory %q does not exist", baseDir)
	}

	rootSet := make(map[string]struct{})
	walkErr := filepath.WalkDir(baseDir, func(curr string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if err := ctx.Err(); err != nil {
			return err
		}

		if d.IsDir() {
			name := d.Name()
			if curr != baseDir && (name == "vendor" || name == "testdata" || strings.HasPrefix(name, ".")) {
				return filepath.SkipDir
			}
		}

		if !d.IsDir() {
			return nil
		}
		if !st.rt.hasGoCandidates(curr) {
			return nil
		}

		pkg, err := st.loadByDir(ctx, curr)
		if err != nil {
			return nil
		}

		rootSet[pkg.ID] = struct{}{}
		return nil
	})
	if walkErr != nil {
		return nil, walkErr
	}

	rootIDs := keysSorted(rootSet)
	return rootIDs, nil
}

func (st *loaderState) loadRecursiveImportPattern(ctx context.Context, importPrefix string) ([]string, error) {
	baseDir, err := st.rt.resolveImportPrefixToDir(importPrefix)
	if err != nil {
		return nil, err
	}

	rootIDs, err := st.loadRecursivePattern(ctx, baseDir)
	if err != nil {
		return nil, err
	}

	filtered := make([]string, 0, len(rootIDs))
	for _, id := range rootIDs {
		if id == importPrefix || strings.HasPrefix(id, importPrefix+"/") {
			filtered = append(filtered, id)
		}
	}

	return filtered, nil
}

func (st *loaderState) ensureBuiltin(ctx context.Context) error {
	_, err := st.loadByImport(ctx, "builtin", st.rt.cfg.Dir)
	if err != nil {
		st.addErrorPackage("builtin", err)
	}
	return ctx.Err()
}

type testVariantIDs struct {
	variantID string
	testMain  string
}

func newTestVariantIDs(rootID string) testVariantIDs {
	return testVariantIDs{
		variantID: rootID + " [" + rootID + ".test]",
		testMain:  rootID + ".test",
	}
}

func (st *loaderState) loadTestVariants(ctx context.Context, rootID string) []string {
	if !st.rt.cfg.Tests || st.rt.mode&packages.NeedForTest == 0 {
		return nil
	}

	if rootID == "builtin" || isTestVariantID(rootID) || strings.HasSuffix(rootID, ".test") {
		return nil
	}

	rootPkg := st.packages[rootID]
	if rootPkg == nil || len(rootPkg.Errors) > 0 {
		return nil
	}

	rootDir, err := st.packageDir(rootPkg)
	if err != nil || rootDir == "" {
		return nil
	}

	if !st.rt.hasTestCandidates(rootDir) {
		return nil
	}

	ids := newTestVariantIDs(rootID)
	variantPkg, err := st.loadResolvedDir(ctx, rootDir, ids.variantID, rootPkg.Module)
	if err != nil || variantPkg == nil || len(variantPkg.Errors) > 0 {
		return nil
	}

	st.addTestMainPackage(ctx, rootPkg, ids)

	return []string{ids.variantID, ids.testMain}
}

func (st *loaderState) packageDir(pkg *packages.Package) (string, error) {
	if pkg == nil {
		return "", errors.New("package is nil")
	}

	if pkg.Dir != "" {
		return pkg.Dir, nil
	}

	importPath := packagePathForID(pkg.ID)
	dir, _, _, err := st.rt.resolveImport(importPath, st.rt.cfg.Dir)
	if err != nil {
		return "", err
	}

	return dir, nil
}

func (st *loaderState) addTestMainPackage(ctx context.Context, rootPkg *packages.Package, ids testVariantIDs) {
	if _, ok := st.packages[ids.testMain]; ok {
		return
	}

	testMainPkg := &packages.Package{
		ID:      ids.testMain,
		PkgPath: ids.testMain,
	}
	if st.rt.wantsName() {
		testMainPkg.Name = "main"
	}

	testMainFile := st.rt.syntheticTestMainPath(ids.testMain)
	if st.rt.wantsFiles() {
		testMainPkg.GoFiles = []string{testMainFile}
	}
	if st.rt.wantsCompiledGoFiles() {
		testMainPkg.CompiledGoFiles = []string{testMainFile}
	}

	if st.rt.wantsImports() {
		testMainPkg.Imports = make(map[string]*packages.Package, 5)

		rootImportPath := rootPkg.PkgPath
		if rootImportPath == "" {
			rootImportPath = packagePathForID(rootPkg.ID)
		}
		testMainPkg.Imports[rootImportPath] = &packages.Package{ID: ids.variantID}

		for _, importPath := range []string{"os", "reflect", "testing", "testing/internal/testdeps"} {
			importID := importPath
			if st.rt.wantsDeps() {
				depPkg, depErr := st.loadByImport(ctx, importPath, rootPkg.Dir)
				if depErr != nil {
					depPkg = st.addErrorPackage(importPath, depErr)
				}
				importID = depPkg.ID
			}

			testMainPkg.Imports[importPath] = &packages.Package{ID: importID}
		}
	}

	st.packages[ids.testMain] = testMainPkg
}

func (st *loaderState) loadByImport(ctx context.Context, importPath, srcDir string) (*packages.Package, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	dir, pkgID, mod, err := st.rt.resolveImport(importPath, srcDir)
	if err != nil {
		return nil, err
	}

	return st.loadResolvedDir(ctx, dir, pkgID, mod)
}

func (st *loaderState) loadByDir(ctx context.Context, dir string) (*packages.Package, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	absDir := normalizeAbsolutePath(st.rt.cfg.Dir, dir)
	pkgID, mod, err := st.rt.idFromDir(absDir)
	if err != nil {
		return nil, err
	}

	return st.loadResolvedDir(ctx, absDir, pkgID, mod)
}

func (st *loaderState) loadResolvedDir(ctx context.Context, dir, pkgID string, mod *packages.Module) (*packages.Package, error) {
	if pkg, ok := st.packages[pkgID]; ok {
		return pkg, nil
	}

	if _, ok := st.loading[pkgID]; ok {
		return st.packages[pkgID], nil
	}

	pkgPath := packagePathForID(pkgID)
	pkg := &packages.Package{
		ID:      pkgID,
		PkgPath: pkgPath,
	}
	if st.rt.wantsName() {
		pkg.Name = path.Base(pkgPath)
	}

	st.packages[pkgID] = pkg
	st.loading[pkgID] = struct{}{}
	defer delete(st.loading, pkgID)

	meta, err := st.rt.readPackageMetadata(dir, isTestVariantID(pkgID))
	if err != nil {
		pkg.Errors = append(pkg.Errors, packages.Error{
			Msg:  err.Error(),
			Kind: packages.ListError,
		})
		return pkg, nil
	}

	if meta.name != "" {
		pkg.Name = meta.name
	}

	if st.rt.wantsFiles() {
		pkg.Dir = dir
		pkg.GoFiles = meta.goFiles
		pkg.IgnoredFiles = meta.ignoredFiles
	}

	if st.rt.wantsCompiledGoFiles() {
		pkg.CompiledGoFiles = meta.compiledGoFiles
	}

	if st.rt.wantsModule() && mod != nil {
		pkg.Module = mod
	}

	if st.rt.wantsImports() {
		pkg.Imports = make(map[string]*packages.Package, len(meta.imports))
		for _, rawImportPath := range meta.imports {
			importPath := rawImportPath
			if rawImportPath == "C" {
				if st.rt.buildCtx.CgoEnabled {
					importPath = "runtime/cgo"
				} else {
					continue
				}
			}

			importID := importPath
			if st.rt.wantsDeps() {
				depPkg, depErr := st.loadByImport(ctx, importPath, dir)
				if depErr != nil {
					depPkg = st.addErrorPackage(importPath, depErr)
				}
				importID = depPkg.ID
			}

			pkg.Imports[importPath] = &packages.Package{
				ID: importID,
			}
		}
	}

	return pkg, nil
}

func (st *loaderState) addErrorPackage(pkgID string, err error) *packages.Package {
	if pkg, ok := st.packages[pkgID]; ok {
		pkg.Errors = append(pkg.Errors, packages.Error{
			Msg:  err.Error(),
			Kind: packages.ListError,
		})
		return pkg
	}

	pkg := &packages.Package{
		ID:      pkgID,
		Name:    path.Base(pkgID),
		PkgPath: pkgID,
		Errors: []packages.Error{
			{
				Msg:  err.Error(),
				Kind: packages.ListError,
			},
		},
	}
	st.packages[pkgID] = pkg
	return pkg
}

func (st *loaderState) buildResponse() *packages.DriverResponse {
	allPackages := make([]*packages.Package, 0, len(st.packages))
	for _, pkg := range st.packages {
		allPackages = append(allPackages, pkg)
	}

	sort.Slice(allPackages, func(i, j int) bool {
		return allPackages[i].ID < allPackages[j].ID
	})

	roots := keysSorted(st.roots)

	compiler := st.rt.cfg.GoVersion.Compiler
	if compiler == "" {
		compiler = runtime.Compiler
	}

	arch := st.rt.cfg.GoVersion.Arch
	if arch == "" {
		arch = st.rt.buildCtx.GOARCH
	}

	return &packages.DriverResponse{
		NotHandled: false,
		Compiler:   compiler,
		Arch:       arch,
		Roots:      roots,
		Packages:   allPackages,
		GoVersion:  st.rt.cfg.GoVersion.GoMinorVersion,
	}
}

type packageMetadata struct {
	name string

	goFiles         []string
	compiledGoFiles []string
	ignoredFiles    []string
	imports         []string
	hasTestFiles    bool
}

func (rt *loaderRuntime) readPackageMetadata(dir string, includeTestFiles bool) (*packageMetadata, error) {
	dir = filepath.Clean(dir)
	fileNames := make(map[string]struct{})

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasSuffix(name, ".go") {
			fileNames[name] = struct{}{}
		}
	}

	for _, name := range rt.overlayDir[dir] {
		if strings.HasSuffix(name, ".go") {
			fileNames[name] = struct{}{}
		}
	}

	if len(fileNames) == 0 {
		return nil, fmt.Errorf("no Go files found in %q", dir)
	}

	sortedNames := make([]string, 0, len(fileNames))
	for name := range fileNames {
		sortedNames = append(sortedNames, name)
	}
	sort.Strings(sortedNames)

	meta := &packageMetadata{
		goFiles:         make([]string, 0, len(sortedNames)),
		compiledGoFiles: make([]string, 0, len(sortedNames)),
		ignoredFiles:    make([]string, 0, len(sortedNames)),
	}

	importSet := make(map[string]struct{})
	var selectedPkgName string
	for _, name := range sortedNames {
		absPath := filepath.Join(dir, name)

		includeFile, err := rt.shouldIncludeGoFile(dir, name, absPath)
		if err != nil {
			meta.ignoredFiles = append(meta.ignoredFiles, absPath)
			continue
		}
		if !includeFile {
			meta.ignoredFiles = append(meta.ignoredFiles, absPath)
			continue
		}
		isTestFile := strings.HasSuffix(name, "_test.go")
		if isTestFile && !includeTestFiles {
			meta.ignoredFiles = append(meta.ignoredFiles, absPath)
			continue
		}

		pkgName, fileImports, parseErr := rt.readFileImports(absPath)
		if parseErr != nil {
			meta.ignoredFiles = append(meta.ignoredFiles, absPath)
			continue
		}

		if selectedPkgName == "" {
			selectedPkgName = pkgName
		}
		if selectedPkgName != "" && pkgName != "" && pkgName != selectedPkgName {
			meta.ignoredFiles = append(meta.ignoredFiles, absPath)
			continue
		}

		meta.goFiles = append(meta.goFiles, absPath)
		meta.compiledGoFiles = append(meta.compiledGoFiles, absPath)
		if isTestFile {
			meta.hasTestFiles = true
		}
		for _, importPath := range fileImports {
			importSet[importPath] = struct{}{}
		}
	}

	if len(meta.goFiles) == 0 {
		return nil, fmt.Errorf("no buildable Go files found in %q", dir)
	}

	meta.name = selectedPkgName
	meta.imports = keysSorted(importSet)
	meta.ignoredFiles = uniqueStrings(meta.ignoredFiles)

	return meta, nil
}

func (rt *loaderRuntime) shouldIncludeGoFile(dir, name, absPath string) (bool, error) {
	_, hasOverlay := rt.overlay[absPath]
	if hasOverlay && !fileExists(absPath) {
		return true, nil
	}

	return rt.buildCtx.MatchFile(dir, name)
}

func (rt *loaderRuntime) readFileImports(absPath string) (string, []string, error) {
	var src any
	if overlayData, ok := rt.overlay[absPath]; ok {
		src = overlayData
	}

	fileNode, err := parser.ParseFile(token.NewFileSet(), absPath, src, parser.ImportsOnly)
	if err != nil {
		return "", nil, err
	}

	importSet := make(map[string]struct{}, len(fileNode.Imports))
	for _, importSpec := range fileNode.Imports {
		quotedPath := importSpec.Path.Value
		importPath, unquoteErr := strconv.Unquote(quotedPath)
		if unquoteErr != nil {
			continue
		}
		importSet[importPath] = struct{}{}
	}

	imports := keysSorted(importSet)

	pkgName := ""
	if fileNode.Name != nil {
		pkgName = fileNode.Name.Name
	}
	return pkgName, imports, nil
}

func (rt *loaderRuntime) hasGoCandidates(dir string) bool {
	if len(rt.overlayDir[filepath.Clean(dir)]) > 0 {
		return true
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".go") {
			return true
		}
	}

	return false
}

func (rt *loaderRuntime) hasTestCandidates(dir string) bool {
	dir = filepath.Clean(dir)

	for _, name := range rt.overlayDir[dir] {
		if strings.HasSuffix(name, "_test.go") {
			absPath := filepath.Join(dir, name)
			includeFile, err := rt.shouldIncludeGoFile(dir, name, absPath)
			if err == nil && includeFile {
				return true
			}
		}
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !strings.HasSuffix(name, "_test.go") {
			continue
		}

		absPath := filepath.Join(dir, name)
		includeFile, err := rt.shouldIncludeGoFile(dir, name, absPath)
		if err == nil && includeFile {
			return true
		}
	}

	return false
}

func (rt *loaderRuntime) syntheticTestMainPath(testMainID string) string {
	cacheRoot := filepath.Clean(rt.cfg.Env["GOCACHE"])
	if cacheRoot == "" || cacheRoot == "." {
		cacheRoot = filepath.Clean(os.TempDir())
	}

	safeName := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		" ", "_",
		"[", "",
		"]", "",
		":", "_",
	).Replace(testMainID)

	return filepath.Join(cacheRoot, "gopackagesdriver", safeName+"-testmain.go")
}

func (rt *loaderRuntime) resolveImportPrefixToDir(importPrefix string) (string, error) {
	if importPrefix == "" {
		return "", errors.New("empty import prefix")
	}

	if rt.mainModule != nil && (importPrefix == rt.mainModule.Path || strings.HasPrefix(importPrefix, rt.mainModule.Path+"/")) {
		rel := strings.TrimPrefix(importPrefix, rt.mainModule.Path)
		rel = strings.TrimPrefix(rel, "/")
		return filepath.Join(rt.mainModule.Dir, filepath.FromSlash(rel)), nil
	}

	for _, gp := range rt.gopath {
		if gp == "" {
			continue
		}
		dir := filepath.Join(gp, "src", filepath.FromSlash(importPrefix))
		if dirExists(dir) {
			return dir, nil
		}
	}

	gorootDir := filepath.Join(rt.goroot, "src", filepath.FromSlash(importPrefix))
	if dirExists(gorootDir) {
		return gorootDir, nil
	}

	return "", fmt.Errorf("cannot resolve import prefix %q", importPrefix)
}

func (rt *loaderRuntime) resolveImport(importPath, srcDir string) (string, string, *packages.Module, error) {
	if importPath == "" {
		return "", "", nil, errors.New("empty import path")
	}

	if importPath == "builtin" {
		dir := filepath.Join(rt.goroot, "src", "builtin")
		return dir, "builtin", nil, nil
	}

	if rt.mainModule != nil && (importPath == rt.mainModule.Path || strings.HasPrefix(importPath, rt.mainModule.Path+"/")) {
		rel := strings.TrimPrefix(importPath, rt.mainModule.Path)
		rel = strings.TrimPrefix(rel, "/")
		dir := filepath.Join(rt.mainModule.Dir, filepath.FromSlash(rel))
		if dirExists(dir) {
			return dir, importPath, rt.mainModule.toPackageModule(), nil
		}
	}

	if srcDir != "" {
		mod, err := rt.moduleForDir(srcDir)
		if err == nil && mod != nil && (importPath == mod.Path || strings.HasPrefix(importPath, mod.Path+"/")) {
			rel := strings.TrimPrefix(importPath, mod.Path)
			rel = strings.TrimPrefix(rel, "/")
			dir := filepath.Join(mod.Dir, filepath.FromSlash(rel))
			if dirExists(dir) {
				return dir, importPath, mod.toPackageModule(), nil
			}
		}
	}

	vendored, err := rt.resolveVendoredImport(importPath, srcDir)
	if err != nil {
		return "", "", nil, err
	}
	if vendored != nil {
		return vendored.dir, vendored.pkgID, vendored.mod, nil
	}

	for _, gp := range rt.gopath {
		if gp == "" {
			continue
		}
		dir := filepath.Join(gp, "src", filepath.FromSlash(importPath))
		if dirExists(dir) {
			return dir, importPath, nil, nil
		}
	}

	gorootDir := filepath.Join(rt.goroot, "src", filepath.FromSlash(importPath))
	if dirExists(gorootDir) {
		return gorootDir, importPath, nil, nil
	}

	moduleCacheHit, err := rt.resolveModuleCacheImport(importPath, srcDir)
	if err != nil {
		return "", "", nil, err
	}
	if moduleCacheHit != nil {
		return moduleCacheHit.dir, moduleCacheHit.pkgID, moduleCacheHit.mod, nil
	}

	return "", "", nil, fmt.Errorf("cannot resolve import path %q", importPath)
}

func (rt *loaderRuntime) resolveVendoredImport(importPath, srcDir string) (*resolvedImport, error) {
	vendorRoot, err := rt.vendorRootForSourceDir(srcDir)
	if err != nil || vendorRoot == "" {
		return nil, err
	}

	vendoredDir := filepath.Join(vendorRoot, "vendor", filepath.FromSlash(importPath))
	if !dirExists(vendoredDir) {
		return nil, nil
	}

	pkgID, mod, err := rt.idFromDir(vendoredDir)
	if err != nil {
		return nil, err
	}

	return &resolvedImport{
		dir:   vendoredDir,
		pkgID: pkgID,
		mod:   mod,
	}, nil
}

func (rt *loaderRuntime) vendorRootForSourceDir(srcDir string) (string, error) {
	if srcDir == "" {
		return "", nil
	}

	cleanDir := filepath.Clean(srcDir)
	gorootSrc := filepath.Join(rt.goroot, "src")
	if _, ok := relativeToBase(gorootSrc, cleanDir); ok {
		for curr := cleanDir; ; curr = filepath.Dir(curr) {
			if _, ok := relativeToBase(gorootSrc, curr); !ok {
				break
			}
			if dirExists(filepath.Join(curr, "vendor")) {
				return curr, nil
			}
			if curr == gorootSrc {
				break
			}
		}
	}

	return "", nil
}

func (rt *loaderRuntime) resolveModuleCacheImport(importPath, srcDir string) (*resolvedImport, error) {
	if len(rt.goModCache) == 0 {
		return nil, nil
	}

	requirementMaps, err := rt.moduleRequirementSources(srcDir)
	if err != nil {
		return nil, err
	}
	if len(requirementMaps) == 0 {
		return nil, nil
	}

	for _, modulePath := range modulePathCandidates(importPath) {
		version := resolveRequiredVersion(modulePath, requirementMaps)
		if version == "" {
			continue
		}

		depModule, depErr := rt.moduleFromVersionRef(modulePath, version)
		if depErr != nil || depModule == nil {
			continue
		}

		depDir := depModule.Dir
		rel := strings.TrimPrefix(importPath, modulePath)
		rel = strings.TrimPrefix(rel, "/")
		if rel != "" {
			depDir = filepath.Join(depDir, filepath.FromSlash(rel))
		}

		if !dirExists(depDir) {
			continue
		}

		return &resolvedImport{
			dir:   depDir,
			pkgID: importPath,
			mod:   depModule.toPackageModule(),
		}, nil
	}

	return nil, nil
}

func (rt *loaderRuntime) moduleRequirementSources(srcDir string) ([]map[string]string, error) {
	sources := make([]map[string]string, 0, 2)

	if rt.mainModule != nil {
		requires, err := rt.moduleRequirements(rt.mainModule)
		if err != nil {
			return nil, err
		}
		if len(requires) > 0 {
			sources = append(sources, requires)
		}
	}

	currentModule, err := rt.moduleForSourceDir(srcDir)
	if err != nil {
		return nil, err
	}
	if currentModule == nil {
		return sources, nil
	}

	if rt.mainModule != nil && currentModule.GoModPath == rt.mainModule.GoModPath {
		return sources, nil
	}

	requires, err := rt.moduleRequirements(currentModule)
	if err != nil {
		return nil, err
	}
	if len(requires) > 0 {
		sources = append(sources, requires)
	}

	return sources, nil
}

func resolveRequiredVersion(modulePath string, requirementMaps []map[string]string) string {
	for _, requires := range requirementMaps {
		version, ok := requires[modulePath]
		if ok && version != "" {
			return version
		}
	}

	return ""
}

func (rt *loaderRuntime) moduleForSourceDir(srcDir string) (*moduleInfo, error) {
	if srcDir != "" {
		mod, err := rt.moduleForDir(srcDir)
		if err != nil {
			return nil, err
		}
		if mod != nil {
			return mod, nil
		}
	}

	return rt.mainModule, nil
}

func (rt *loaderRuntime) moduleRequirements(mod *moduleInfo) (map[string]string, error) {
	if mod == nil || mod.GoModPath == "" {
		return nil, nil
	}

	if cached, ok := rt.moduleReqs[mod.GoModPath]; ok {
		return cached, nil
	}

	summary, err := readGoModSummary(mod.GoModPath)
	if err != nil {
		return nil, err
	}

	rt.moduleReqs[mod.GoModPath] = summary.requires
	return summary.requires, nil
}

func (rt *loaderRuntime) moduleFromVersionRef(modulePath, version string) (*moduleInfo, error) {
	if modulePath == "" || version == "" {
		return nil, nil
	}

	cacheKey := modulePath + "@" + version
	if cached, ok := rt.moduleByRef[cacheKey]; ok {
		return cached, nil
	}

	escapedPath, err := modmodule.EscapePath(modulePath)
	if err != nil {
		return nil, err
	}

	escapedVersion, err := modmodule.EscapeVersion(version)
	if err != nil {
		return nil, err
	}

	modDir := rt.findModuleDir(escapedPath, escapedVersion)
	resolvedVersion := version
	if modDir == "" {
		var bestVersion string
		modDir, bestVersion, err = rt.findBestModuleDir(escapedPath)
		if err != nil {
			return nil, err
		}
		if modDir == "" {
			return nil, nil
		}
		resolvedVersion = bestVersion
	}

	goModPath := filepath.Join(modDir, "go.mod")
	modGoVersion := ""
	parsedModulePath := modulePath
	if fileExists(goModPath) {
		summary, parseErr := readGoModSummary(goModPath)
		if parseErr == nil {
			if summary.modulePath != "" {
				parsedModulePath = summary.modulePath
			}
			modGoVersion = summary.goVersion
			rt.moduleReqs[goModPath] = summary.requires
		}
	}

	mi := &moduleInfo{
		Path:      parsedModulePath,
		Dir:       modDir,
		GoModPath: goModPath,
		GoVersion: modGoVersion,
		Main:      false,
	}

	rt.moduleByRef[cacheKey] = mi
	if resolvedVersion != "" {
		resolvedCacheKey := modulePath + "@" + resolvedVersion
		rt.moduleByRef[resolvedCacheKey] = mi
	}

	return mi, nil
}

func (rt *loaderRuntime) findModuleDir(escapedPath, escapedVersion string) string {
	cacheDirName := escapedPath + "@" + escapedVersion
	for _, root := range rt.goModCache {
		modDir := filepath.Join(root, cacheDirName)
		if dirExists(modDir) {
			return modDir
		}
	}

	return ""
}

func (rt *loaderRuntime) findBestModuleDir(escapedPath string) (string, string, error) {
	bestDir := ""
	bestVersion := ""

	for _, root := range rt.goModCache {
		matches, err := filepath.Glob(filepath.Join(root, escapedPath+"@*"))
		if err != nil {
			return "", "", err
		}

		for _, match := range matches {
			if !dirExists(match) {
				continue
			}

			base := filepath.Base(match)
			at := strings.LastIndex(base, "@")
			if at <= 0 {
				continue
			}

			escapedVersion := base[at+1:]
			unescapedVersion, err := modmodule.UnescapeVersion(escapedVersion)
			if err != nil {
				continue
			}

			if bestDir == "" || compareModuleVersions(unescapedVersion, bestVersion) > 0 {
				bestDir = match
				bestVersion = unescapedVersion
			}
		}
	}

	return bestDir, bestVersion, nil
}

func compareModuleVersions(a, b string) int {
	aNorm := normalizeModuleVersion(a)
	bNorm := normalizeModuleVersion(b)

	if aNorm != "" && bNorm != "" {
		return semver.Compare(aNorm, bNorm)
	}

	if aNorm != "" {
		return 1
	}
	if bNorm != "" {
		return -1
	}

	return strings.Compare(a, b)
}

func normalizeModuleVersion(version string) string {
	if version == "" {
		return ""
	}

	if semver.IsValid(version) {
		return version
	}

	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}

	if semver.IsValid(version) {
		return version
	}

	return ""
}

func modulePathCandidates(importPath string) []string {
	if importPath == "" {
		return nil
	}

	parts := strings.Split(importPath, "/")
	out := make([]string, 0, len(parts))
	for idx := len(parts); idx > 0; idx-- {
		out = append(out, strings.Join(parts[:idx], "/"))
	}

	return out
}

func readGoModSummary(goModPath string) (*goModSummary, error) {
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return nil, err
	}

	moduleFile, err := modfile.ParseLax(goModPath, data, nil)
	if err != nil {
		return nil, err
	}

	if moduleFile.Module == nil || moduleFile.Module.Mod.Path == "" {
		return nil, fmt.Errorf("go.mod in %q does not declare module path", filepath.Dir(goModPath))
	}

	moduleGoVersion := ""
	if moduleFile.Go != nil {
		moduleGoVersion = moduleFile.Go.Version
	}

	requires := make(map[string]string, len(moduleFile.Require))
	for _, req := range moduleFile.Require {
		if req == nil || req.Mod.Path == "" || req.Mod.Version == "" {
			continue
		}
		requires[req.Mod.Path] = req.Mod.Version
	}

	return &goModSummary{
		modulePath: moduleFile.Module.Mod.Path,
		goVersion:  moduleGoVersion,
		requires:   requires,
	}, nil
}

func (rt *loaderRuntime) idFromDir(dir string) (string, *packages.Module, error) {
	dir = filepath.Clean(dir)
	stdRoot := filepath.Join(rt.goroot, "src")
	if rel, ok := relativeToBase(stdRoot, dir); ok {
		if rel == "." || rel == "" {
			return "", nil, fmt.Errorf("directory %q is not a package directory", dir)
		}
		return filepath.ToSlash(rel), nil, nil
	}

	mod, err := rt.moduleForDir(dir)
	if err != nil {
		return "", nil, err
	}
	if mod != nil {
		rel, ok := relativeToBase(mod.Dir, dir)
		if !ok {
			return "", nil, fmt.Errorf("directory %q is outside module %q", dir, mod.Dir)
		}

		pkgPath := mod.Path
		if rel != "." && rel != "" {
			pkgPath = path.Join(mod.Path, filepath.ToSlash(rel))
		}

		return pkgPath, mod.toPackageModule(), nil
	}

	for _, gp := range rt.gopath {
		if gp == "" {
			continue
		}

		srcRoot := filepath.Join(gp, "src")
		rel, ok := relativeToBase(srcRoot, dir)
		if ok {
			return filepath.ToSlash(rel), nil, nil
		}
	}

	return filepath.ToSlash(dir), nil, nil
}

func (rt *loaderRuntime) moduleForDir(dir string) (*moduleInfo, error) {
	dir = filepath.Clean(dir)
	if cached, ok := rt.moduleCache[dir]; ok {
		return cached, nil
	}

	visited := make([]string, 0, 8)
	curr := dir
	for {
		if cached, ok := rt.moduleCache[curr]; ok {
			for _, v := range visited {
				rt.moduleCache[v] = cached
			}
			return cached, nil
		}

		visited = append(visited, curr)
		goModPath := filepath.Join(curr, "go.mod")
		if fileExists(goModPath) {
			summary, err := readGoModSummary(goModPath)
			if err != nil {
				return nil, err
			}

			mi := &moduleInfo{
				Path:      summary.modulePath,
				Dir:       curr,
				GoModPath: goModPath,
				GoVersion: summary.goVersion,
				Main:      false,
			}
			rt.moduleReqs[goModPath] = summary.requires
			for _, v := range visited {
				rt.moduleCache[v] = mi
			}
			return mi, nil
		}

		parent := filepath.Dir(curr)
		if parent == curr {
			break
		}
		curr = parent
	}

	for _, v := range visited {
		rt.moduleCache[v] = nil
	}
	return nil, nil
}

func normalizePatterns(workDir string, patterns []string) []string {
	out := make([]string, 0, len(patterns))
	for _, in := range patterns {
		p := strings.TrimSpace(in)
		if p == "" {
			continue
		}

		if p == "builtin" {
			out = append(out, p)
			continue
		}

		if strings.HasPrefix(p, "file=") {
			target := strings.TrimPrefix(p, "file=")
			out = append(out, "file="+normalizeAbsolutePath(workDir, target))
			continue
		}

		if p == "..." {
			out = append(out, filepath.ToSlash(filepath.Clean(workDir))+"/...")
			continue
		}

		if strings.HasSuffix(p, "/...") {
			base := strings.TrimSuffix(p, "/...")
			if base == "" || base == "." {
				base = workDir
			}

			if filepath.IsAbs(base) || strings.HasPrefix(base, ".") {
				out = append(out, filepath.ToSlash(normalizeAbsolutePath(workDir, base))+"/...")
				continue
			}
			out = append(out, path.Clean(p))
			continue
		}

		if p == "." || p == ".." || strings.HasPrefix(p, "./") || strings.HasPrefix(p, "../") {
			out = append(out, normalizeAbsolutePath(workDir, p))
			continue
		}

		if filepath.IsAbs(p) {
			out = append(out, filepath.Clean(p))
			continue
		}

		out = append(out, path.Clean(p))
	}
	return out
}

func chunkPatterns(patterns []string, maxChunkLen int) [][]string {
	if len(patterns) == 0 {
		return nil
	}

	if maxChunkLen <= 0 {
		maxChunkLen = safeArgMax
	}

	chunks := make([][]string, 0, 1)
	chunk := make([]string, 0, len(patterns))
	chunkLen := 0
	for _, p := range patterns {
		plen := len(p) + 1
		if len(chunk) > 0 && chunkLen+plen > maxChunkLen {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0, len(patterns))
			chunkLen = 0
		}
		chunk = append(chunk, p)
		chunkLen += plen
	}

	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}

	return chunks
}

func isAbsoluteRecursivePattern(pattern string) bool {
	return strings.HasSuffix(pattern, "/...") && filepath.IsAbs(strings.TrimSuffix(pattern, "/..."))
}

func isTestVariantID(pkgID string) bool {
	if !strings.Contains(pkgID, " [") || !strings.HasSuffix(pkgID, ".test]") {
		return false
	}

	openIdx := strings.Index(pkgID, " [")
	return openIdx > 0
}

func packagePathForID(pkgID string) string {
	if !isTestVariantID(pkgID) {
		return pkgID
	}

	openIdx := strings.Index(pkgID, " [")
	if openIdx <= 0 {
		return pkgID
	}

	return pkgID[:openIdx]
}

func normalizeAbsolutePath(baseDir, path string) string {
	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}
	return filepath.Clean(filepath.Join(baseDir, path))
}

func splitPathList(pathValue string) []string {
	if pathValue == "" {
		return nil
	}

	items := filepath.SplitList(pathValue)
	out := make([]string, 0, len(items))
	for _, item := range items {
		if item == "" {
			continue
		}
		out = append(out, filepath.Clean(item))
	}
	return out
}

func resolveGoModCacheDirs(env map[string]string, gopath []string) []string {
	caches := make([]string, 0, len(gopath)+2)

	if gomodcache := filepath.Clean(env["GOMODCACHE"]); gomodcache != "" && gomodcache != "." {
		caches = append(caches, gomodcache)
	}

	if gomodcache := filepath.Clean(os.Getenv("GOMODCACHE")); gomodcache != "" && gomodcache != "." {
		caches = append(caches, gomodcache)
	}

	for _, gp := range gopath {
		if gp == "" {
			continue
		}
		caches = append(caches, filepath.Join(gp, "pkg", "mod"))
	}

	if len(caches) == 0 {
		return nil
	}

	seen := make(map[string]struct{}, len(caches))
	unique := make([]string, 0, len(caches))
	for _, dir := range caches {
		if dir == "" {
			continue
		}
		if _, ok := seen[dir]; ok {
			continue
		}
		seen[dir] = struct{}{}
		unique = append(unique, dir)
	}

	return unique
}

func keysSorted[T any](m map[string]T) []string {
	out := make([]string, 0, len(m))
	for key := range m {
		out = append(out, key)
	}
	sort.Strings(out)
	return out
}

func uniqueStrings(in []string) []string {
	if len(in) == 0 {
		return nil
	}

	sort.Strings(in)

	out := make([]string, 0, len(in))
	var prev string
	for i, item := range in {
		if i == 0 || item != prev {
			out = append(out, item)
			prev = item
		}
	}
	return out
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func relativeToBase(baseDir, target string) (string, bool) {
	rel, err := filepath.Rel(baseDir, target)
	if err != nil {
		return "", false
	}
	if rel == "." {
		return ".", true
	}
	if strings.HasPrefix(rel, ".."+string(filepath.Separator)) || rel == ".." {
		return "", false
	}
	return rel, true
}
