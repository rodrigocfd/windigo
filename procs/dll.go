package procs

import "syscall"

var (
	dllComCtl32 = syscall.NewLazyDLL("comctl32.dll")
	dllGdi32    = syscall.NewLazyDLL("gdi32.dll")
	dllKernel32 = syscall.NewLazyDLL("kernel32.dll")
	dllUser32   = syscall.NewLazyDLL("user32.dll")
	dllUxTheme  = syscall.NewLazyDLL("uxtheme.dll")
)
