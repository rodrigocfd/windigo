package api

import (
	"syscall"
	p "winffi/procs"
)

type HPEN HANDLE

func (hPen HPEN) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(p.DeleteObject.Addr(), 1,
		uintptr(hPen), 0, 0)
	return ret != 0
}
