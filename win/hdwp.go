/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"wingows/co"
	"wingows/win/proc"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdwp
type HDWP HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-begindeferwindowpos
func BeginDeferWindowPos(numWindows uint32) HDWP {
	ret, _, lerr := syscall.Syscall(proc.BeginDeferWindowPos.Addr(), 1,
		uintptr(numWindows), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("BeginDeferWindowPos failed. %s",
			co.ERROR(lerr).Error()))
	}
	return HDWP(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deferwindowpos
func (hDwp HDWP) DeferWindowPos(hWnd HWND, hWndInsertAfter HWND, x, y int32,
	cx, cy uint32, uFlags co.SWP) HDWP {

	ret, _, lerr := syscall.Syscall9(proc.DeferWindowPos.Addr(), 8,
		uintptr(hDwp), uintptr(hWnd), uintptr(hWndInsertAfter),
		uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags),
		0)
	if ret == 0 {
		hDwp.endDeferWindowPosNoPanic() // free resource
		panic(fmt.Sprintf("DeferWindowPos failed. %s", co.ERROR(lerr).Error()))
	}
	return HDWP(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddeferwindowpos
func (hDwp HDWP) EndDeferWindowPos() {
	lerr := hDwp.endDeferWindowPosNoPanic()
	if lerr != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("EndDeferWindowPos failed. %s", co.ERROR(lerr).Error()))
	}
}

func (hDwp HDWP) endDeferWindowPosNoPanic() co.ERROR {
	if hDwp == 0 { // handle is null, do nothing
		return co.ERROR_SUCCESS
	}
	ret, _, lerr := syscall.Syscall(proc.EndDeferWindowPos.Addr(), 1,
		uintptr(hDwp), 0, 0)
	if ret == 0 { // an error occurred
		return co.ERROR(lerr)
	}
	return co.ERROR_SUCCESS
}
