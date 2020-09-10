/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"windigo/win/proc"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hicon
type HICON HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyicon
func (hIcon HICON) DestroyIcon() {
	if hIcon != 0 {
		syscall.Syscall(proc.DestroyIcon.Addr(), 1,
			uintptr(hIcon), 0, 0)
	}
}
