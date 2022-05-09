//go:build windows

package proc

import (
	"syscall"
)

var (
	version = syscall.NewLazyDLL("version.dll")

	GetFileVersionInfo     = version.NewProc("GetFileVersionInfoW")
	GetFileVersionInfoSize = version.NewProc("GetFileVersionInfoSizeW")
	VerQueryValue          = version.NewProc("VerQueryValueW")
)
