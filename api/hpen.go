package api

import (
	"syscall"
	"winffi/procs"
)

type HPEN HANDLE

func (hPen HPEN) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(procs.DeleteObject.Addr(), 1,
		uintptr(hPen), 0, 0)
	return ret != 0
}
