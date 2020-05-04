package api

import (
	"syscall"
	"unsafe"
	p "winffi/procs"
)

type HINSTANCE HANDLE

func (hinst HINSTANCE) GetClassInfo(className *uint16,
	destBuf *WNDCLASSEX) syscall.Errno {

	_, _, errno := syscall.Syscall(p.GetClassInfo.Addr(), 3,
		uintptr(hinst),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(destBuf)))
	return errno
}

func GetModuleHandle(moduleName string) HINSTANCE {
	ret, _, _ := syscall.Syscall(p.GetModuleHandle.Addr(), 1,
		uintptr(unsafe.Pointer(toUtf16PtrBlankIsNil(moduleName))),
		0, 0)
	return HINSTANCE(ret)
}
