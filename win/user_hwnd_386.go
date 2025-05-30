//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
)

// [GetClassLong] function.
//
// [GetClassLong]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclasslongw
func (hWnd HWND) GetClassLongPtr(index co.GCL) (uintptr, error) {
	ret, _, err := syscall.SyscallN(_GetClassLongW.Addr(),
		uintptr(hWnd),
		uintptr(index))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return ret, nil
}

var _GetClassLongW = dll.User32.NewProc("GetClassLongW")

// [GetWindowLong] function.
//
// [GetWindowLong]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongw
func (hWnd HWND) GetWindowLongPtr(index co.GWLP) (uintptr, error) {
	ret, _, err := syscall.SyscallN(_GetWindowLongW.Addr(),
		uintptr(hWnd),
		uintptr(index))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return ret, nil
}

var _GetWindowLongW = dll.User32.NewProc("GetWindowLongW")

// [SetWindowLong] function.
//
// [SetWindowLong]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowlongw
func (hWnd HWND) SetWindowLongPtr(index co.GWLP, newLong uintptr) (uintptr, error) {
	ret, _, err := syscall.SyscallN(_SetWindowLongW.Addr(),
		uintptr(hWnd),
		uintptr(index),
		uintptr(int32(newLong)))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return ret, nil
}

var _SetWindowLongW = dll.User32.NewProc("SetWindowLongW")
