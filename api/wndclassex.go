package api

import (
	"gowinui/api/proc"
	c "gowinui/consts"
	"syscall"
	"unsafe"
)

type WNDCLASSEX struct {
	CbSize        uint32
	Style         c.CS
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
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
