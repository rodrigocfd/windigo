/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	proc "windigo/win/internal"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hcursor
type HCURSOR HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycursor
func (hCursor HCURSOR) DestroyCursor() {
	if hCursor != 0 {
		syscall.Syscall(proc.DestroyCursor.Addr(), 1,
			uintptr(hCursor), 0, 0)
	}
}
