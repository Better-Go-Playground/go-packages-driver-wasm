package server

import (
	"golang.org/x/tools/go/packages"
)

type DriverRequestEnvelope struct {
	WorkDir       string                 `json:"workDir"`
	Patterns      []string               `json:"patterns"`
	DriverRequest packages.DriverRequest `json:"driverRequest"`
}
