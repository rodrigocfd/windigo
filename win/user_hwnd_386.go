//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
)

// [GetClassLong] function.
//
// [GetClassLong]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclasslongw
func (hWnd HWND) GetClassLongPtr(index co.GCL) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_GetClassLongW, "GetClassLongW"),
		uintptr(hWnd),
		uintptr(index))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return ret, nil
}

var _user_GetClassLongW *syscall.Proc

// [GetWindowLong] function.
//
// [GetWindowLong]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongw
func (hWnd HWND) GetWindowLongPtr(index co.GWLP) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_GetWindowLongW, "GetWindowLongW"),
		uintptr(hWnd),
		uintptr(index))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return ret, nil
}

var _user_GetWindowLongW *syscall.Proc

// [SetWindowLong] function.
//
// [SetWindowLong]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowlongw
func (hWnd HWND) SetWindowLongPtr(index co.GWLP, newLong uintptr) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_SetWindowLongW, "SetWindowLongW"),
		uintptr(hWnd),
		uintptr(index),
		uintptr(int32(newLong)))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return ret, nil
}

var _user_SetWindowLongW *syscall.Proc
