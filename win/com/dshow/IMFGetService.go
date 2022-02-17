package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
type IMFGetService struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IMFGetService.Release().
//
// Example:
//
//  var vmr dshow.IBaseFilter // initialized somewhere
//
//  gs := dshow.NewIMFGetService(
//      vmr.QueryInterface(dshowco.IID_IMFGetService),
//  )
//  defer gs.Release()
func NewIMFGetService(base com.IUnknown) IMFGetService {
	return IMFGetService{IUnknown: base}
}

// ⚠️ The returned pointer must be used to construct a COM object; you must
// defer its Release().
//
// Example for IMFVideoDisplayControl:
//
//  var gs dshow.IMFGetService // initialized somewhere
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
	guidService, riid *win.GUID) com.IUnknown {

	var ppQueried com.IUnknown
	ret, _, _ := syscall.Syscall6(
		(*dshowvt.IMFGetService)(unsafe.Pointer(*me.Ptr())).GetService, 4,
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
