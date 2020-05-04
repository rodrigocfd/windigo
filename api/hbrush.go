package api

import (
	"syscall"
	"winffi/api/proc"
)

type HBRUSH HANDLE

func (hBrush HBRUSH) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hBrush), 0, 0)
	return ret != 0
}
