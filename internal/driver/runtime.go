package driver

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
)

var goMinorVerRegex = regexp.MustCompile(`(?m)^go1\.(\d+)`)

type GoVersion struct {
	GoMinorVersion int
	Compiler       string
	Arch           string
}

// GoVersionFromRuntime builds [GoVersion] from [runtime] package information.
func GoVersionFromRuntime() (ri GoVersion, err error) {
	rv := runtime.Version()
	matches := goMinorVerRegex.FindStringSubmatch(rv)
	if matches == nil {
		return ri, fmt.Errorf("failed to parse Go version: %q", rv)
	}

	minorVer, err := strconv.Atoi(matches[1])
	if err != nil {
		return ri, fmt.Errorf("go minor version is not a number: %q (in %q)", matches[1], rv)
	}

	return GoVersion{
		GoMinorVersion: minorVer,
		Compiler:       runtime.Compiler,
		Arch:           runtime.GOARCH,
	}, nil
}
