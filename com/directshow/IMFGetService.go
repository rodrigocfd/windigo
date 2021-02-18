/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package directshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

type (
	// IMFGetService > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
	IMFGetService struct{ win.IUnknown }

	IMFGetServiceVtbl struct {
		win.IUnknownVtbl
		GetService uintptr
	}
)

// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nf-mfidl-imfgetservice-getservice
func (me *IMFGetService) GetService(guidService, riid *win.GUID) *win.IUnknown {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall6(
		(*IMFGetServiceVtbl)(unsafe.Pointer(*me.Ppv)).GetService, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(guidService)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppQueried)), 0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFGetService.GetService"))
	}
	return &win.IUnknown{Ppv: ppQueried}
}

// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/api/evr/nn-evr-imfvideodisplaycontrol
func (me *IMFGetService) GetIMFVideoDisplayControl() *IMFVideoDisplayControl {
	iUnk := me.GetService(
		win.NewGuid(0x1092a86c, 0xab1a, 0x459a, 0xa336, 0x831fbc4d11ff), // MR_VIDEO_RENDER_SERVICE
		win.NewGuid(0xa490b1e4, 0xab84, 0x4d31, 0xa1b2, 0x181e03b1077a), // IID_IMFVideoDisplayControl
	)
	return &IMFVideoDisplayControl{
		IUnknown: *iUnk,
	}
}
