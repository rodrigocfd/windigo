package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/err"
)

type _IEnumMediaTypesVtbl struct {
	win.IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienummediatypes
type IEnumMediaTypes struct {
	win.IUnknown // Base IUnknown.
}

// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-clone
func (me *IEnumMediaTypes) Clone() (IEnumMediaTypes, error) {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IEnumMediaTypesVtbl)(unsafe.Pointer(*me.Ppv)).Clone, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		return IEnumMediaTypes{}, lerr
	}
	return IEnumMediaTypes{
		win.IUnknown{Ppv: ppQueried},
	}, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-reset
func (me *IEnumMediaTypes) Reset() {
	syscall.Syscall(
		(*_IEnumMediaTypesVtbl)(unsafe.Pointer(*me.Ppv)).Reset, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)
}
