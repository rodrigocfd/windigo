package winffi

import (
	"syscall"
	"winffi/proc"
)

func InitCommonControls() {
	syscall.Syscall(proc.InitCommonControls.Addr(), 0,
		0, 0, 0)
}
