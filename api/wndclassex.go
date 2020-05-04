package api

import (
	"syscall"
	"unsafe"
	"winffi/api/proc"
	c "winffi/consts"
)

type WNDCLASSEX struct {
	Size          uint32
	Style         c.CS
	WndProc       uintptr
	ClsExtra      int32
	WndExtra      int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}

func (wcx *WNDCLASSEX) RegisterClassEx() (ATOM, syscall.Errno) {
	ret, _, errno := syscall.Syscall(proc.RegisterClassEx.Addr(), 1,
		uintptr(unsafe.Pointer(wcx)), 0, 0)
	return ATOM(ret), errno
}
