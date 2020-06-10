/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type HINSTANCE HANDLE

func (hinst HINSTANCE) GetClassInfoEx(className *uint16,
	destBuf *WNDCLASSEX) ATOM {

	ret, _, lerr := syscall.Syscall(proc.GetClassInfoEx.Addr(), 3,
		uintptr(hinst),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(destBuf)))
	if ret == 0 {
		panic(fmt.Sprintf("GetClassInfoEx failed: %d %s\n",
			lerr, lerr.Error()))
	}
	return ATOM(ret)
}

func GetModuleHandle(moduleName string) HINSTANCE {
	ret, _, _ := syscall.Syscall(proc.GetModuleHandle.Addr(), 1,
		uintptr(unsafe.Pointer(StrToUtf16PtrBlankIsNil(moduleName))),
		0, 0)
	return HINSTANCE(ret)
}

func (hinst HINSTANCE) LoadCursor(lpCursorName co.IDC) HCURSOR {
	ret, _, lerr := syscall.Syscall(proc.LoadCursor.Addr(), 2,
		uintptr(hinst), uintptr(lpCursorName), 0)
	if ret == 0 {
		panic(fmt.Sprintf("LoadCursor failed: %d %s\n",
			lerr, lerr.Error()))
	}
	return HCURSOR(ret)
}

func (hinst HINSTANCE) LoadIcon(lpIconName co.IDI) HICON {
	ret, _, lerr := syscall.Syscall(proc.LoadIcon.Addr(), 2,
		uintptr(hinst), uintptr(lpIconName), 0)
	if ret == 0 {
		panic(fmt.Sprintf("LoadIcon failed: %d %s\n",
			lerr, lerr.Error()))
	}
	return HICON(ret)
}
