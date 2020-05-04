package api

import (
	"syscall"
	"winffi/api/proc"
	c "winffi/consts"
)

type HBRUSH HANDLE

func NewBrushFromSysColor(sysColor c.COLOR) HBRUSH {
	return HBRUSH(sysColor + 1)
}

func (hBrush HBRUSH) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hBrush), 0, 0)
	return ret != 0
}
