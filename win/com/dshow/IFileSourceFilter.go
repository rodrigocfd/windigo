//go:build windows

package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [IFileSourceFilter] COM interface.
//
// [IFileSourceFilter]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesourcefilter
type IFileSourceFilter interface {
	com.IUnknown

	// [GetCurFile] COM method.
	//
	// Returns false if no file is opened.
	//
	// [GetCurFile]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesourcefilter-getcurfile
	GetCurFile(mt *AM_MEDIA_TYPE) (string, bool)

	// [Load] COM method.
	//
	// [Load]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesourcefilter-load
	Load(fileName string, mt *AM_MEDIA_TYPE)
}

type _IFileSourceFilter struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IFileSourceFilter.Release().
func NewIFileSourceFilter(base com.IUnknown) IFileSourceFilter {
	return &_IFileSourceFilter{IUnknown: base}
}

func (me *_IFileSourceFilter) GetCurFile(mt *AM_MEDIA_TYPE) (string, bool) {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFileSourceFilter)(unsafe.Pointer(*me.Ptr())).GetCurFile,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pv)), uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		defer win.HTASKMEM(pv).CoTaskMemFree()
		name := win.Str.FromNativePtr((*uint16)(unsafe.Pointer(pv)))
		return name, true
	} else if hr == errco.E_FAIL {
		return "", false
	} else {
		panic(hr)
	}
}

func (me *_IFileSourceFilter) Load(fileName string, mt *AM_MEDIA_TYPE) {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFileSourceFilter)(unsafe.Pointer(*me.Ptr())).Load,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(fileName))),
		uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
