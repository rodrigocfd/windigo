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
	// IPersist > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
	IPersist struct{ win.IUnknown }

	IPersistVtbl struct {
		win.IUnknownVtbl
		GetClassID uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ipersist-getclassid
func (me *IPersist) GetClassID() *win.GUID {
	clsid := win.GUID{}
	ret, _, _ := syscall.Syscall(
		(*IPersistVtbl)(unsafe.Pointer(*me.Ppv)).GetClassID, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&clsid)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IPersist.GetClassID"))
	}
	return &clsid
}
