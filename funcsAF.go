package winffi

import (
	"syscall"
	"unsafe"
	"winffi/procs"
)

func (lf *LOGFONT) CreateFontIndirect() HFONT {
	ret, _, _ := syscall.Syscall(procs.CreateFontIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lf)), 0, 0)
	return HFONT(ret)
}
