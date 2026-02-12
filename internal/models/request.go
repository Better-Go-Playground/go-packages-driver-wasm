package models

import "golang.org/x/tools/go/packages"

// DriverServerRequest represents a request from gopls send to a driver which runs in a server mode.
type DriverServerRequest struct {
	WorkDir       string                 `json:"workDir"`
	Patterns      []string               `json:"patterns"`
	DriverRequest packages.DriverRequest `json:"driverRequest"`
}
