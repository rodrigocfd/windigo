package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IFileSourceFilterVtbl struct {
	win.IUnknownVtbl
	Load       uintptr
	GetCurFile uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesourcefilter
type IFileSourceFilter struct {
	win.IUnknown // Base IUnknown.
}

// Returns false if no file is opened.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesourcefilter-getcurfile
func (me *IFileSourceFilter) GetCurFile(mt *AM_MEDIA_TYPE) (string, bool) {
	var pv *uint16
	ret, _, _ := syscall.Syscall(
		(*_IFileSourceFilterVtbl)(unsafe.Pointer(*me.Ppv)).GetCurFile, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&pv)), uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		name := win.Str.FromNativePtr(pv)
		win.CoTaskMemFree(unsafe.Pointer(pv))
		return name, true
	} else if hr == errco.E_FAIL {
		return "", false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesourcefilter-load
func (me *IFileSourceFilter) Load(fileName string, mt *AM_MEDIA_TYPE) {
	ret, _, _ := syscall.Syscall(
		(*_IFileSourceFilterVtbl)(unsafe.Pointer(*me.Ppv)).Load, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(fileName))),
		uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
