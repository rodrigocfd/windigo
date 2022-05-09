//go:build windows

package proc

import (
	"syscall"
)

var (
	psapi = syscall.NewLazyDLL("psapi.dll")

	EnumProcesses      = psapi.NewProc("EnumProcesses")
	EnumProcessModules = psapi.NewProc("EnumProcessModules")
	GetModuleBaseName  = psapi.NewProc("GetModuleBaseNameW")
)
