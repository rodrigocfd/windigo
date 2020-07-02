/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"syscall"
	"unsafe"
)

var Guid_IUnknown = makeGuid(0x00000000, 0x0000, 0x0000, 0xC000000000000046)

type IUnknown struct {
	lpVtbl *iUnknownVtbl
}

type iUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

func (v *IUnknown) Release() uint32 {
	ret, _, _ := syscall.Syscall(v.lpVtbl.Release, 1,
		uintptr(unsafe.Pointer(v)), 0, 0)
	return uint32(ret)
}
