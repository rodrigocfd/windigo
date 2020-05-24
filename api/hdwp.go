/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package api

import (
	"fmt"
	"syscall"
	"wingows/api/proc"
)

type HDWP HANDLE

func BeginDeferWindowPos(numWindows uint32) HDWP {
	ret, _, lerr := syscall.Syscall(proc.BeginDeferWindowPos.Addr(), 1,
		uintptr(numWindows), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("BeginDeferWindowPos failed: %d %s\n",
			lerr, lerr.Error()))
	}
	return HDWP(ret)
}

func (hdwp HDWP) DeferWindowPos(hWnd HWND, hWndInsertAfter HWND, x, y int32,
	cx, cy uint32, uFlags SWP) HDWP {

	ret, _, lerr := syscall.Syscall9(proc.DeferWindowPos.Addr(), 8,
		uintptr(hdwp), uintptr(hWnd), uintptr(hWndInsertAfter),
		uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags),
		0)
	if ret == 0 {
		panic(fmt.Sprintf("DeferWindowPos failed: %d %s\n",
			lerr, lerr.Error()))
	}
	return HDWP(ret)
}

func (hdwp HDWP) EndDeferWindowPos() {
	ret, _, lerr := syscall.Syscall(proc.EndDeferWindowPos.Addr(), 1,
		uintptr(hdwp), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("EndDeferWindowPos failed: %d %s\n",
			lerr, lerr.Error()))
	}
}
