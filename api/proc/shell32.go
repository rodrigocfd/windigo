package proc

import (
	"syscall"
)

var (
	dllShell32 = syscall.NewLazyDLL("shell32.dll")

	DragFinish     = dllShell32.NewProc("DragFinish")
	DragQueryFile  = dllShell32.NewProc("DragQueryFileW")
	DragQueryPoint = dllShell32.NewProc("DragQueryPoint")
)
