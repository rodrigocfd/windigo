package dshow

import (
	"syscall"
	"unsafe"

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
type IFileSinkFilter2 struct {
	IFileSinkFilter // Base IFileSinkFilter > IUnknown.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesinkfilter2-getmode
func (me *IFileSinkFilter2) GetMode() dshowco.AM_FILE {
	pdwFlags := dshowco.AM_FILE(0)
	ret, _, _ := syscall.Syscall(
		(*_IFileSinkFilter2Vtbl)(unsafe.Pointer(*me.Ppv)).GetMode, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&pdwFlags)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return pdwFlags
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesinkfilter2-setmode
func (me *IFileSinkFilter2) SetMode(dwFlags dshowco.AM_FILE) {
	ret, _, _ := syscall.Syscall(
		(*_IFileSinkFilter2Vtbl)(unsafe.Pointer(*me.Ppv)).SetMode, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(dwFlags), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
