//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
)

// [GetClassLongPtr] function.
//
// [GetClassLongPtr]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclasslongptrw
func (hWnd HWND) GetClassLongPtr(index co.GCL) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_GetClassLongPtrW, "GetClassLongPtrW"),
		uintptr(hWnd),
		uintptr(index))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return ret, nil
}

var _user_GetClassLongPtrW *syscall.Proc

// [GetWindowLongPtr] function.
//
// [GetWindowLongPtr]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
func (hWnd HWND) GetWindowLongPtr(index co.GWLP) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_GetWindowLongPtrW, "GetWindowLongPtrW"),
		uintptr(hWnd),
		uintptr(index))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return ret, nil
}

var _user_GetWindowLongPtrW *syscall.Proc

// [SetWindowLongPtr] function.
//
// [SetWindowLongPtr]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowlongptrw
func (hWnd HWND) SetWindowLongPtr(index co.GWLP, newLong uintptr) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_SetWindowLongPtrW, "SetWindowLongPtrW"),
		uintptr(hWnd),
		uintptr(index),
		newLong)
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return ret, nil
}

var _user_SetWindowLongPtrW *syscall.Proc
