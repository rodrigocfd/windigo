package procs

import (
	"syscall"
)

var (
	dllKernel32 = syscall.NewLazyDLL("kernel32.dll")

	GetModuleHandle     = dllKernel32.NewProc("GetModuleHandleW")
	VerifyVersionInfo   = dllKernel32.NewProc("VerifyVersionInfoW")
	VerSetConditionMask = dllKernel32.NewProc("VerSetConditionMask")
)
