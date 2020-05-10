package proc

import (
	"syscall"
)

var (
	dllShCore = syscall.NewLazyDLL("shcore.dll")

	GetDpiForMonitor = dllShCore.NewProc("GetDpiForMonitor")
)
