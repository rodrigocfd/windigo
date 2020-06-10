/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"wingows/win/proc"
)

type HFONT HANDLE

func (hFont HFONT) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hFont), 0, 0)
	return ret != 0
}
