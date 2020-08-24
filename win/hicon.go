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

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hicon
type HICON HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyicon
func (hIcon HICON) DestroyIcon() {
	ret, _, lerr := syscall.Syscall(proc.DestroyIcon.Addr(), 1,
		uintptr(hIcon), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("DestroyIcon failed. %s", co.ERROR(lerr).Error()))
	}
}
