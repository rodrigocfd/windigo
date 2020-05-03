package api

import (
	"syscall"
	"winffi/procs"
)

type HBRUSH HANDLE

func (hBrush HBRUSH) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(procs.DeleteObject.Addr(), 1,
		uintptr(hBrush), 0, 0)
	return ret != 0
}
