package proc

import (
	"syscall"
)

var (
	dllKernel32 = syscall.NewLazyDLL("kernel32.dll")

	MulDiv              = dllKernel32.NewProc("MulDiv")
	GetModuleHandle     = dllKernel32.NewProc("GetModuleHandleW")
	VerifyVersionInfo   = dllKernel32.NewProc("VerifyVersionInfoW")
	VerSetConditionMask = dllKernel32.NewProc("VerSetConditionMask")
)
