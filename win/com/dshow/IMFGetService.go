package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/err"
)

type _IMFGetServiceVtbl struct {
	win.IUnknownVtbl
	GetService uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
type IMFGetService struct {
	win.IUnknown // Base IUnknown.
}

// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nf-mfidl-imfgetservice-getservice
func (me *IMFGetService) GetService(
	guidService, riid *win.GUID) (win.IUnknown, error) {

	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall6(
		(*_IMFGetServiceVtbl)(unsafe.Pointer(*me.Ppv)).GetService, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(guidService)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppQueried)), 0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		return win.IUnknown{}, lerr
	}
	return win.IUnknown{Ppv: ppQueried}, nil
}

// Calls IMFGetService.GetService() to return IMFVideoDisplayControl.
//
// ⚠️ You must defer Release().
func (me *IMFGetService) GetServiceIMFVideoDisplayControl() (IMFVideoDisplayControl, error) {
	iUnk, lerr := me.GetService(win.NewGuidFromClsid(CLSID.MR_VideoRenderService),
		win.NewGuidFromIid(IID.IMFVideoDisplayControl))
	if lerr != nil {
		return IMFVideoDisplayControl{}, lerr
	}
	return IMFVideoDisplayControl{IUnknown: iUnk}, nil
}
