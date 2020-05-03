package api

import (
	"syscall"
	"unsafe"
	p "winffi/procs"
)

type WNDCLASSEX struct {
	Size          uint32
	Style         uint32
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
	ret, _, errno := syscall.Syscall(p.RegisterClassEx.Addr(), 1,
		uintptr(unsafe.Pointer(wcx)), 0, 0)
	return ATOM(ret), errno
}
