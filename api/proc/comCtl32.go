package proc

import (
	"syscall"
)

var (
	dllComCtl32 = syscall.NewLazyDLL("comctl32.dll")

	InitCommonControls = dllComCtl32.NewProc("InitCommonControls")
)
