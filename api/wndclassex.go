/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package api

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/api/proc"
	"wingows/co"
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
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}

// Generates a string from all fields, excluding CbSize and LpszClassName, that
// uniquely identifies a WNDCLASSEX object.
func (wcx *WNDCLASSEX) Hash() string {
	return fmt.Sprintf("%x.%x.%x.%x.%x.%x.%x.%x.%x.%x",
		wcx.Style, wcx.LpfnWndProc, wcx.CbClsExtra, wcx.CbWndExtra,
		wcx.HInstance, wcx.HIcon, wcx.HCursor, wcx.HbrBackground,
		wcx.LpszMenuName, wcx.HIconSm)
}

func (wcx *WNDCLASSEX) RegisterClassEx() (ATOM, syscall.Errno) {
	wcx.CbSize = uint32(unsafe.Sizeof(*wcx)) // safety

	ret, _, lerr := syscall.Syscall(proc.RegisterClassEx.Addr(), 1,
		uintptr(unsafe.Pointer(wcx)), 0, 0)
	return ATOM(ret), lerr
}
