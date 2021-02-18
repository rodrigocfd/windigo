/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	proc "github.com/rodrigocfd/windigo/win/internal"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hicon
type HICON HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyicon
func (hIcon HICON) CopyIcon() HICON {
	ret, _, lerr := syscall.Syscall(proc.CopyIcon.Addr(), 1,
		uintptr(hIcon), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "CopyIcon"))
	}
	return HICON(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyicon
func (hIcon HICON) DestroyIcon() {
	if hIcon != 0 {
		syscall.Syscall(proc.DestroyIcon.Addr(), 1,
			uintptr(hIcon), 0, 0)
	}
}
