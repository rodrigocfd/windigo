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

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbrush
type HBRUSH HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpatternbrush
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

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBrush HBRUSH) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hBrush), 0, 0)
	return ret != 0
}
