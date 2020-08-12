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

type HBRUSH HANDLE

func CreatePatternBrush(hbm HBITMAP) HBRUSH {
	ret, _, _ := syscall.Syscall(proc.CreatePatternBrush.Addr(), 1,
		uintptr(hbm), 0, 0)
	if ret == 0 {
		panic("CreatePatternBrush failed.")
	}
	return HBRUSH(ret)
}

// This is not an actual Win32 function, it's just a tricky conversion to create
// a brush from a system color.
func CreateSysColorBrush(sysColor co.COLOR) HBRUSH {
	return HBRUSH(sysColor + 1)
}

func (hBrush HBRUSH) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hBrush), 0, 0)
	return ret != 0
}
