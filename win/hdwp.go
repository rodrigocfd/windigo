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

type HDWP HANDLE

func BeginDeferWindowPos(numWindows uint32) HDWP {
	ret, _, lerr := syscall.Syscall(proc.BeginDeferWindowPos.Addr(), 1,
		uintptr(numWindows), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("BeginDeferWindowPos failed: %d %s",
			lerr, lerr.Error()))
	}
	return HDWP(ret)
}

func (hDwp HDWP) DeferWindowPos(hWnd HWND, hWndInsertAfter HWND, x, y int32,
	cx, cy uint32, uFlags co.SWP) HDWP {

	ret, _, lerr := syscall.Syscall9(proc.DeferWindowPos.Addr(), 8,
		uintptr(hDwp), uintptr(hWnd), uintptr(hWndInsertAfter),
		uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags),
		0)
	if ret == 0 {
		hDwp.endDeferWindowPosNoPanic() // cleanup
		panic(fmt.Sprintf("DeferWindowPos failed: %d %s",
			lerr, lerr.Error()))
	}
	return HDWP(ret)
}

func (hDwp HDWP) EndDeferWindowPos() {
	ret, lerr := hDwp.endDeferWindowPosNoPanic()
	if ret == 0 {
		panic(fmt.Sprintf("EndDeferWindowPos failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hDwp HDWP) endDeferWindowPosNoPanic() (uintptr, syscall.Errno) {
	if hDwp == 0 {
		return 1, syscall.Errno(co.ERROR_SUCCESS)
	}
	ret, _, lerr := syscall.Syscall(proc.EndDeferWindowPos.Addr(), 1,
		uintptr(hDwp), 0, 0)
	return ret, lerr
}
