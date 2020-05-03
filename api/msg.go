package api

import (
	"syscall"
	"unsafe"
	p "winffi/procs"
)

type MSG struct {
	HWnd   HWND
	Msg    uint32
	WParam WPARAM
	LParam LPARAM
	Time   uint32
	Pt     POINT
}

func (msg *MSG) DispatchMessage() uintptr {
	ret, _, _ := syscall.Syscall(p.DispatchMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret
}

func (msg *MSG) GetMessage(hWnd HWND,
	msgFilterMin, msgFilterMax uint32) (int32, syscall.Errno) {

	ret, _, errno := syscall.Syscall6(p.GetMessage.Addr(), 4,
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax),
		0, 0)
	return int32(ret), errno
}
