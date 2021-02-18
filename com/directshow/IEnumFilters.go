/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package directshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

type (
	// IEnumFilters > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienumfilters
	IEnumFilters struct{ win.IUnknown }

	IEnumFiltersVtbl struct {
		win.IUnknownVtbl
		Next  uintptr
		Skip  uintptr
		Reset uintptr
		Clone uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-reset
func (me *IEnumFilters) Reset() {
	ret, _, _ := syscall.Syscall(
		(*IEnumFiltersVtbl)(unsafe.Pointer(*me.Ppv)).Reset, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IEnumFilters.Reset"))
	}
}
