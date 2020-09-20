/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package directshow

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
	//
	// IMFGetService > IUnknown.
	IMFGetService struct{ win.IUnknown }

	IMFGetServiceVtbl struct {
		win.IUnknownVtbl
		GetService uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nf-mfidl-imfgetservice-getservice
func (me *IMFGetService) GetService(
	guidService, riid *win.GUID) **win.IUnknownVtbl {

	var ppvObject **win.IUnknownVtbl = nil
	ret, _, _ := syscall.Syscall6(
		(*IMFGetServiceVtbl)(unsafe.Pointer(*me.Ppv)).GetService, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(guidService)), uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppvObject)), 0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFGetService.GetService").Error())
	}
	return ppvObject
}

// https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nf-mfidl-imfgetservice-getservice
//
// https://docs.microsoft.com/en-us/windows/win32/api/evr/nn-evr-imfvideodisplaycontrol
func (me *IMFGetService) GetIMFVideoDisplayControl() IMFVideoDisplayControl {
	return IMFVideoDisplayControl{
		win.IUnknown{
			Ppv: me.GetService(
				win.NewGuid(0x1092a86c, 0xab1a, 0x459a, 0xa336_831fbc4d11ff),  // MR_VIDEO_RENDER_SERVICE
				win.NewGuid(0xa490b1e4, 0xab84, 0x4d31, 0xa1b2_181e03b1077a)), // IID_IMFVideoDisplayControl
		},
	}
}
