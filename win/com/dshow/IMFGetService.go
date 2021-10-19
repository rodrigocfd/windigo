package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IMFGetServiceVtbl struct {
	win.IUnknownVtbl
	GetService uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
type IMFGetService struct{ win.IUnknown }

// Constructs a COM object from a pointer to its COM virtual table.
//
// ⚠️ You must defer IMFGetService.Release().
func NewIMFGetService(ptr win.IUnknownPtr) IMFGetService {
	return IMFGetService{
		IUnknown: win.NewIUnknown(ptr),
	}
}

// ⚠️ The returned pointer must be used to construct a COM object; you must
// defer its Release().
//
// Example for IMFVideoDisplayControl:
//
//  var gs dshow.IMFGetService
//
//  vdc := dshow.NewIMFVideoDisplayControl(
//      gs.GetService(
//          win.NewGuidFromClsid(dshowco.CLSID_MR_VideoRenderService),
//          win.NewGuidFromIid(dshowco.IID_IMFVideoDisplayControl),
//      ),
//  )
//  defer vdc.Release()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nf-mfidl-imfgetservice-getservice
func (me *IMFGetService) GetService(
	guidService, riid *win.GUID) win.IUnknownPtr {

	var ppQueried win.IUnknownPtr
	ret, _, _ := syscall.Syscall6(
		(*_IMFGetServiceVtbl)(unsafe.Pointer(*me.Ptr())).GetService, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(guidService)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppQueried)), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return ppQueried
	} else {
		panic(hr)
	}
}
