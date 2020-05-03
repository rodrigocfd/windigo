package api

import (
	"syscall"
	"winffi/procs"
)

type HFONT HANDLE

func (hFont HFONT) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(procs.DeleteObject.Addr(), 1,
		uintptr(hFont), 0, 0)
	return ret != 0
}
