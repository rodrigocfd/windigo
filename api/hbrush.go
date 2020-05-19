/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package api

import (
	"syscall"
	"wingows/api/proc"
	c "wingows/consts"
)

type HBRUSH HANDLE

func NewBrushFromSysColor(sysColor c.COLOR) HBRUSH {
	return HBRUSH(sysColor + 1)
}

func CreatePatternBrush(hbm HBITMAP) HBRUSH {
	ret, _, _ := syscall.Syscall(proc.CreatePatternBrush.Addr(), 1,
		uintptr(hbm), 0, 0)
	if ret == 0 {
		panic("CreatePatternBrush failed.")
	}
	return HBRUSH(ret)
}

func (hBrush HBRUSH) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hBrush), 0, 0)
	return ret != 0
}
