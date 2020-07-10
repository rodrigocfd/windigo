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

type HICON HANDLE

func (hIcon HICON) DestroyIcon() {
	ret, _, lerr := syscall.Syscall(proc.DestroyIcon.Addr(), 1,
		uintptr(hIcon), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("DestroyIcon failed."))
	}
}
