/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type WNDCLASSEX struct {
	CbSize        uint32
	Style         co.CS
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  uintptr // LPCWSTR
	LpszClassName uintptr // LPCWSTR
	HIconSm       HICON
}

func (wcx *WNDCLASSEX) RegisterClassEx() (ATOM, co.ERROR) {
	wcx.CbSize = uint32(unsafe.Sizeof(*wcx)) // safety
	ret, _, lerr := syscall.Syscall(proc.RegisterClassEx.Addr(), 1,
		uintptr(unsafe.Pointer(wcx)), 0, 0)
	return ATOM(ret), co.ERROR(lerr)
}
