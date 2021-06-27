package proc

import (
	"syscall"
)

var (
	dwmapi = syscall.NewLazyDLL("dwmapi.dll")

	DwmIsCompositionEnabled = dwmapi.NewProc("DwmIsCompositionEnabled")
)
