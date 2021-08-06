package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IFileSinkFilterVtbl struct {
	win.IUnknownVtbl
	SetFileName uintptr
	GetCurFile  uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesinkfilter
type IFileSinkFilter struct {
	win.IUnknown // Base IUnknown.
}

// Returns false if no file is opened.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesinkfilter-getcurfile
func (me *IFileSinkFilter) GetCurFile(pmt *AM_MEDIA_TYPE) (string, bool) {
	var pv *uint16
	ret, _, _ := syscall.Syscall(
		(*_IFileSinkFilterVtbl)(unsafe.Pointer(*me.Ppv)).GetCurFile, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&pv)), uintptr(unsafe.Pointer(pmt)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		name := win.Str.FromUint16Ptr(pv)
		win.CoTaskMemFree(unsafe.Pointer(pv))
		return name, true
	} else if hr == errco.E_FAIL {
		return "", false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesinkfilter-setfilename
func (me *IFileSinkFilter) SetFileName(pszFileName string, pmt *AM_MEDIA_TYPE) {
	ret, _, _ := syscall.Syscall(
		(*_IFileSinkFilterVtbl)(unsafe.Pointer(*me.Ppv)).SetFileName, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(pszFileName))),
		uintptr(unsafe.Pointer(pmt)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
