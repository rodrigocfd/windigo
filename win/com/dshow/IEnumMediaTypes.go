package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
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

// ⚠️ You must defer IEnumMediaTypes.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-clone
func (me *IEnumMediaTypes) Clone() IEnumMediaTypes {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IEnumMediaTypesVtbl)(unsafe.Pointer(*me.Ppv)).Clone, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IEnumMediaTypes{
			win.IUnknown{Ppv: ppQueried},
		}
	} else {
		panic(hr)
	}
}

// Calls Skip() until the end of the enum to retrieve the actual number of media
// types, then calls Reset().
func (me *IEnumMediaTypes) Count() int {
	count := int(0)
	for {
		gotOne := me.Skip(1)
		if gotOne {
			count++
		} else {
			me.Reset()
			return count
		}
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-next
func (me *IEnumMediaTypes) Next(mt *AM_MEDIA_TYPE) bool {
	ret, _, _ := syscall.Syscall6(
		(*_IEnumMediaTypesVtbl)(unsafe.Pointer(*me.Ppv)).Next, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		1, uintptr(unsafe.Pointer(&mt)), 0, 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-reset
func (me *IEnumMediaTypes) Reset() {
	syscall.Syscall(
		(*_IEnumMediaTypesVtbl)(unsafe.Pointer(*me.Ppv)).Reset, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-skip
func (me *IEnumMediaTypes) Skip(numMediaTypes int) bool {
	ret, _, _ := syscall.Syscall(
		(*_IEnumMediaTypesVtbl)(unsafe.Pointer(*me.Ppv)).Skip, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(uint32(numMediaTypes)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
