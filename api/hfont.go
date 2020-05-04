package api

import (
	"syscall"
	"winffi/api/proc"
)

type HFONT HANDLE

func (hFont HFONT) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hFont), 0, 0)
	return ret != 0
}
