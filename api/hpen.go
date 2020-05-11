package api

import (
	"syscall"
	"wingows/api/proc"
)

type HPEN HANDLE

func (hPen HPEN) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hPen), 0, 0)
	return ret != 0
}
