package api

import (
	"syscall"
	p "winffi/procs"
)

func InitCommonControls() {
	syscall.Syscall(p.InitCommonControls.Addr(), 0,
		0, 0, 0)
}

func PostQuitMessage(exitCode int32) {
	syscall.Syscall(p.PostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
}
