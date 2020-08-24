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

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hcursor
type HCURSOR HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycursor
func (hCursor HCURSOR) DestroyCursor() {
	ret, _, lerr := syscall.Syscall(proc.DestroyCursor.Addr(), 1,
		uintptr(hCursor), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("DestroyCursor failed. %s", co.ERROR(lerr).Error()))
	}
}
