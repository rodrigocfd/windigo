/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package api

import (
	"syscall"
	"wingows/api/proc"
)

type HPEN HANDLE

func (hPen HPEN) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hPen), 0, 0)
	return ret != 0
}
