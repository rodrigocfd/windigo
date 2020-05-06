package api

import (
	"fmt"
	"gowinui/api/proc"
	c "gowinui/consts"
	"syscall"
	"unsafe"
)

type HINSTANCE HANDLE

func (hinst HINSTANCE) GetClassInfo(className *uint16,
	destBuf *WNDCLASSEX) syscall.Errno {

	_, _, errno := syscall.Syscall(proc.GetClassInfo.Addr(), 3,
		uintptr(hinst),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(destBuf)))
	return errno
}

func GetModuleHandle(moduleName string) HINSTANCE {
	ret, _, _ := syscall.Syscall(proc.GetModuleHandle.Addr(), 1,
		uintptr(unsafe.Pointer(StrToUtf16PtrBlankIsNil(moduleName))),
		0, 0)
	return HINSTANCE(ret)
}

func (hinst HINSTANCE) LoadCursor(lpCursorName c.IDC) HCURSOR {
	ret, _, errno := syscall.Syscall(proc.LoadCursor.Addr(), 2,
		uintptr(hinst), uintptr(lpCursorName), 0)
	if ret == 0 {
		panic(fmt.Sprintf("LoadCursor failed: %d %s\n",
			errno, errno.Error()))
	}
	return HCURSOR(ret)
}
