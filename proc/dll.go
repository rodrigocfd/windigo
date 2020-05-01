package proc

import "syscall"

var (
	dllComCtl32 = syscall.NewLazyDLL("comctl32.dll")
	dllUser32   = syscall.NewLazyDLL("user32.dll")
)
