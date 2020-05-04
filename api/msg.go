package api

import (
	"fmt"
	"syscall"
	"unsafe"
	"winffi/api/proc"
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
	ret, _, _ := syscall.Syscall(proc.DispatchMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret
}

func (msg *MSG) GetMessage(hWnd HWND, msgFilterMin, msgFilterMax uint32) int32 {
	ret, _, errno := syscall.Syscall6(proc.GetMessage.Addr(), 4,
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax),
		0, 0)
	if int32(ret) == -1 {
		panic(fmt.Sprintf("GetMessage failed: %d %s\n",
			errno, errno.Error()))
	}
	return int32(ret)
}

func (msg *MSG) TranslateMessage() bool {
	ret, _, _ := syscall.Syscall(proc.TranslateMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret != 0
}
