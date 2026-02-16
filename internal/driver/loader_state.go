package driver

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

// loader_state.go contains the package graph loading state machine.

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

	if after, ok := strings.CutPrefix(pattern, "file="); ok {
		fileName := normalizeAbsolutePath(st.rt.cfg.Dir, after)
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

// loadTestVariants creates tests=true companion roots for a loaded root package.
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

			if importPath == pkg.ID {
				continue
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
