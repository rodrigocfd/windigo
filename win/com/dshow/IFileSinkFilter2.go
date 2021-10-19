package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IFileSinkFilter2Vtbl struct {
	_IFileSinkFilterVtbl
	SetMode uintptr
	GetMode uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesinkfilter2
type IFileSinkFilter2 struct{ IFileSinkFilter }

// Constructs a COM object from a pointer to its COM virtual table.
//
// ⚠️ You must defer IFileSinkFilter2.Release().
func NewIFileSinkFilter2(ptr win.IUnknownPtr) IFileSinkFilter2 {
	return IFileSinkFilter2{
		IFileSinkFilter: NewIFileSinkFilter(ptr),
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesinkfilter2-getmode
func (me *IFileSinkFilter2) GetMode() dshowco.AM_FILE {
	var pdwFlags dshowco.AM_FILE
	ret, _, _ := syscall.Syscall(
		(*_IFileSinkFilter2Vtbl)(unsafe.Pointer(*me.Ptr())).GetMode, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pdwFlags)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return pdwFlags
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesinkfilter2-setmode
func (me *IFileSinkFilter2) SetMode(flags dshowco.AM_FILE) {
	ret, _, _ := syscall.Syscall(
		(*_IFileSinkFilter2Vtbl)(unsafe.Pointer(*me.Ptr())).SetMode, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(flags), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
