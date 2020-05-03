package api

import (
	"syscall"
	"unsafe"
	p "winffi/procs"
)

type HINSTANCE HANDLE

func (hinst HINSTANCE) GetClassInfo(className string,
	wcx *WNDCLASSEX) syscall.Errno {

	_, _, errno := syscall.Syscall(p.GetClassInfo.Addr(), 3,
		uintptr(hinst), toUtf16ToUintptr(className), uintptr(unsafe.Pointer(wcx)))
	return errno
}

func GetModuleHandle(moduleName string) HINSTANCE {
	ret, _, _ := syscall.Syscall(p.GetModuleHandle.Addr(), 1,
		toUtf16BlankIsNilToUintptr(moduleName), 0, 0)
	return HINSTANCE(ret)
}
