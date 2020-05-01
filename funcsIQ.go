package winffi

import (
	"syscall"
	"winffi/consts"
	"winffi/procs"
)

func InitCommonControls() {
	syscall.Syscall(procs.InitCommonControls.Addr(), 0,
		0, 0, 0)
}

func (hwnd HWND) MessageBox(message, caption string, flags consts.MB) int32 {
	ret, _, _ := syscall.Syscall6(procs.MessageBox.Addr(), 4,
		uintptr(0), toUtf16(message), toUtf16(caption), uintptr(flags),
		0, 0)
	return int32(ret)
}
