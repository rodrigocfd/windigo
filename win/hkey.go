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

type HKEY HANDLE

func (hkey HKEY) RegCloseKey() {
	ret, _, _ := syscall.Syscall(proc.RegCloseKey.Addr(), 1,
		uintptr(hkey), 0, 0)
	lerr := syscall.Errno(ret)
	if co.ERROR(lerr) != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegCloseKey failed: %d %s\n",
			ret, lerr.Error()))
	}
}
