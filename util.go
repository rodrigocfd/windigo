package winffi

import (
	"syscall"
	"unsafe"
)

func toUtf16(s string) uintptr {
	pstr, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(pstr))
}
