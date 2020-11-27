/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"windigo/co"
	proc "windigo/win/internal"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdwp
type HDWP HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-begindeferwindowpos
func BeginDeferWindowPos(numWindows int32) HDWP {
	ret, _, lerr := syscall.Syscall(proc.BeginDeferWindowPos.Addr(), 1,
		uintptr(numWindows), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "BeginDeferWindowPos"))
	}
	return HDWP(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deferwindowpos
func (hDwp HDWP) DeferWindowPos(
	hWnd HWND, hWndInsertAfter HWND, x, y, cx, cy int32, uFlags co.SWP) HDWP {

	ret, _, lerr := syscall.Syscall9(proc.DeferWindowPos.Addr(), 8,
		uintptr(hDwp), uintptr(hWnd), uintptr(hWndInsertAfter),
		uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags),
		0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "DeferWindowPos"))
	}
	return HDWP(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddeferwindowpos
func (hDwp HDWP) EndDeferWindowPos() {
	ret, _, lerr := syscall.Syscall(proc.EndDeferWindowPos.Addr(), 1,
		uintptr(hDwp), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "EndDeferWindowPos"))
	}
}
