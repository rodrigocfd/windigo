/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"wingows/co"
	"wingows/win/proc"
)

type HDWP HANDLE

func BeginDeferWindowPos(numWindows uint32) HDWP {
	ret, _, lerr := syscall.Syscall(proc.BeginDeferWindowPos.Addr(), 1,
		uintptr(numWindows), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("BeginDeferWindowPos failed."))
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
		hDwp.endDeferWindowPosNoPanic() // free resource
		panic(co.ERROR(lerr).Format("DeferWindowPos failed."))
	}
	return HDWP(ret)
}

func (hDwp HDWP) EndDeferWindowPos() {
	lerr := hDwp.endDeferWindowPosNoPanic()
	if lerr != co.ERROR_SUCCESS {
		panic(lerr.Format("EndDeferWindowPos failed."))
	}
}

func (hDwp HDWP) endDeferWindowPosNoPanic() co.ERROR {
	return freeNoPanic(HANDLE(hDwp), proc.EndDeferWindowPos)
}
