package winffi

import (
	"syscall"
	"winffi/proc"
)

func (hWnd HWND) MessageBox(message, caption string, flags MB) int32 {
	ret, _, _ := syscall.Syscall6(proc.MessageBox.Addr(), 4,
		uintptr(hWnd), toUtf16(message), toUtf16(caption), uintptr(flags),
		0, 0)
	return int32(ret)
}
