package proc

import (
	"syscall"
)

var (
	dllGdi32 = syscall.NewLazyDLL("gdi32.dll")

	CreateFontIndirect = dllGdi32.NewProc("CreateFontIndirectW")
	DeleteObject       = dllGdi32.NewProc("DeleteObject")
)
