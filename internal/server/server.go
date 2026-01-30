package server

import (
	"context"
	"fmt"
	"regexp"
	"runtime"
	"strconv"

	"golang.org/x/tools/go/packages"

	"github.com/Better-Go-Playground/go-packages-driver-wasm/internal/jsonrpc"
)

var goMinorVerRegex = regexp.MustCompile(`(?m)^go1\.(\d+)`)

type DriverRequestEnvelope struct {
	WorkDir       string                 `json:"workDir"`
	Patterns      []string               `json:"patterns"`
	DriverRequest packages.DriverRequest `json:"driverRequest"`
}

type Service struct {
	ri RuntimeInfo
}

func NewService(ri RuntimeInfo) Service {
	return Service{
		ri: ri,
	}
}

func (svc Service) handleDriverRequest(ctx context.Context, req *DriverRequestEnvelope) (*packages.DriverResponse, error) {
	// TODO: implement driver logic
	return &packages.DriverResponse{
		NotHandled: true,
		Compiler:   svc.ri.Compiler,
		Arch:       svc.ri.Arch,
		Roots:      []string{},
		Packages:   []*packages.Package{},
		GoVersion:  svc.ri.GoMinorVersion,
	}, nil
}

func (svc Service) Listener() *jsonrpc.Listener {
	handlers := map[string]jsonrpc.RequestHandler{
		"goPackageDriver/query": jsonrpc.NewHandler(svc.handleDriverRequest),
	}

	return jsonrpc.NewListener(handlers)
}

type RuntimeInfo struct {
	GoMinorVersion int
	Compiler       string
	Arch           string
}

// BuildRuntimeInfo builds [RuntimeInfo] from [runtime] package information.
func BuildRuntimeInfo() (ri RuntimeInfo, err error) {
	rv := runtime.Version()
	matches := goMinorVerRegex.FindStringSubmatch(rv)
	if matches == nil {
		return ri, fmt.Errorf("failed to parse Go version: %q", rv)
	}

	minorVer, err := strconv.Atoi(matches[1])
	if err != nil {
		return ri, fmt.Errorf("go minor version is not a number: %q (in %q)", matches[1], rv)
	}

	return RuntimeInfo{
		GoMinorVersion: minorVer,
		Compiler:       runtime.Compiler,
		Arch:           runtime.GOARCH,
	}, nil
}
