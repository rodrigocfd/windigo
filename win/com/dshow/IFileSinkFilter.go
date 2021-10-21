package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesinkfilter
type IFileSinkFilter struct{ win.IUnknown }

// Constructs a COM object from a pointer to its COM virtual table.
//
// ⚠️ You must defer IFileSinkFilter.Release().
func NewIFileSinkFilter(ptr win.IUnknownPtr) IFileSinkFilter {
	return IFileSinkFilter{
		IUnknown: win.NewIUnknown(ptr),
	}
}

// Returns false if no file is opened.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesinkfilter-getcurfile
func (me *IFileSinkFilter) GetCurFile(mt *AM_MEDIA_TYPE) (string, bool) {
	var pv uintptr
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IFileSinkFilter)(unsafe.Pointer(*me.Ptr())).GetCurFile, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pv)), uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		name := win.Str.FromNativePtr((*uint16)(unsafe.Pointer(pv)))
		win.CoTaskMemFree(pv)
		return name, true
	} else if hr == errco.E_FAIL {
		return "", false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesinkfilter-setfilename
func (me *IFileSinkFilter) SetFileName(fileName string, mt *AM_MEDIA_TYPE) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IFileSinkFilter)(unsafe.Pointer(*me.Ptr())).SetFileName, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(fileName))),
		uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
