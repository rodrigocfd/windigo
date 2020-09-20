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
	// https://docs.microsoft.com/en-us/windows/win32/api/evr/nn-evr-imfvideodisplaycontrol
	//
	// IMFVideoDisplayControl > IUnknown.
	IMFVideoDisplayControl struct{ win.IUnknown }

	IMFVideoDisplayControlVtbl struct {
		win.IUnknownVtbl
		GetNativeVideoSize uintptr
		GetIdealVideoSize  uintptr
		SetVideoPosition   uintptr
		GetVideoPosition   uintptr
		SetAspectRatioMode uintptr
		GetAspectRatioMode uintptr
		SetVideoWindow     uintptr
		GetVideoWindow     uintptr
		RepaintVideo       uintptr
		GetCurrentImage    uintptr
		SetBorderColor     uintptr
		GetBorderColor     uintptr
		SetRenderingPrefs  uintptr
		GetRenderingPrefs  uintptr
		SetFullscreen      uintptr
		GetFullscreen      uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setaspectratiomode
func (me *IMFVideoDisplayControl) SetAspectRatioMode(
	mode MFVideoARMode) *IMFVideoDisplayControl {

	ret, _, _ := syscall.Syscall(
		(*IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).SetAspectRatioMode, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(mode), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFVideoDisplayControl.SetAspectRatioMode").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setvideowindow
func (me *IMFVideoDisplayControl) SetVideoWindow(
	hwndVideo win.HWND) *IMFVideoDisplayControl {

	ret, _, _ := syscall.Syscall(
		(*IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).SetVideoWindow, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndVideo), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFVideoDisplayControl.SetVideoWindow").Error())
	}
	return me
}
