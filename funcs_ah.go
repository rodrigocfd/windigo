package winffi

import (
	"syscall"
	"winffi/procs"
)

func InitCommonControls() {
	syscall.Syscall(procs.InitCommonControls.Addr(), 0,
		0, 0, 0)
}
