//go:build windows

package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesinkfilter2
type IFileSinkFilter2 interface {
	IFileSinkFilter

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesinkfilter2-getmode
	GetMode() dshowco.AM_FILE

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesinkfilter2-setmode
	SetMode(flags dshowco.AM_FILE)
}

type _IFileSinkFilter2 struct{ IFileSinkFilter }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IFileSinkFilter2.Release().
func NewIFileSinkFilter2(base com.IUnknown) IFileSinkFilter2 {
	return &_IFileSinkFilter2{IFileSinkFilter: NewIFileSinkFilter(base)}
}

func (me *_IFileSinkFilter2) GetMode() dshowco.AM_FILE {
	var pdwFlags dshowco.AM_FILE
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFileSinkFilter2)(unsafe.Pointer(*me.Ptr())).GetMode,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pdwFlags)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return pdwFlags
	} else {
		panic(hr)
	}
}

func (me *_IFileSinkFilter2) SetMode(flags dshowco.AM_FILE) {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFileSinkFilter2)(unsafe.Pointer(*me.Ptr())).SetMode,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(flags))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
