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
	// IMFVideoDisplayControl > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/evr/nn-evr-imfvideodisplaycontrol
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

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getnativevideosize
func (me *IMFVideoDisplayControl) GetNativeVideoSize() (*win.SIZE, *win.SIZE) {
	nativeSize, aspectRatio := win.SIZE{}, win.SIZE{}
	ret, _, _ := syscall.Syscall(
		(*IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).GetNativeVideoSize, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&nativeSize)),
		uintptr(unsafe.Pointer(&aspectRatio)))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFVideoDisplayControl.GetNativeVideoSize"))
	}
	return &nativeSize, &aspectRatio
}

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getidealvideosize
func (me *IMFVideoDisplayControl) GetIdealVideoSize() (*win.SIZE, *win.SIZE) {
	min, max := win.SIZE{}, win.SIZE{}
	ret, _, _ := syscall.Syscall(
		(*IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).GetIdealVideoSize, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&min)),
		uintptr(unsafe.Pointer(&max)))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFVideoDisplayControl.GetIdealVideoSize"))
	}
	return &min, &max
}

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setvideoposition
func (me *IMFVideoDisplayControl) SetVideoPosition(
	pnrcSource *MFVideoNormalizedRect, prcDest *win.RECT) {

	ret, _, _ := syscall.Syscall(
		(*IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).SetVideoPosition, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pnrcSource)),
		uintptr(unsafe.Pointer(prcDest)))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFVideoDisplayControl.SetVideoPosition"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getvideoposition
func (me *IMFVideoDisplayControl) GetVideoPosition() (*MFVideoNormalizedRect, *win.RECT) {
	pnrcSource, prcDest := MFVideoNormalizedRect{}, win.RECT{}
	ret, _, _ := syscall.Syscall(
		(*IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).GetVideoPosition, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&pnrcSource)),
		uintptr(unsafe.Pointer(&prcDest)))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFVideoDisplayControl.GetVideoPosition"))
	}
	return &pnrcSource, &prcDest
}

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setaspectratiomode
func (me *IMFVideoDisplayControl) SetAspectRatioMode(mode MFVideoARMode) {
	ret, _, _ := syscall.Syscall(
		(*IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).SetAspectRatioMode, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(mode), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFVideoDisplayControl.SetAspectRatioMode"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getaspectratiomode
func (me *IMFVideoDisplayControl) GetAspectRatioMode() MFVideoARMode {
	aspectRatioMode := MFVideoARMode(0)
	ret, _, _ := syscall.Syscall(
		(*IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).GetAspectRatioMode, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&aspectRatioMode)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFVideoDisplayControl.GetAspectRatioMode"))
	}
	return aspectRatioMode
}

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setvideowindow
func (me *IMFVideoDisplayControl) SetVideoWindow(hwndVideo win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).SetVideoWindow, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndVideo), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFVideoDisplayControl.SetVideoWindow"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getvideowindow
func (me *IMFVideoDisplayControl) GetVideoWindow() win.HWND {
	hwndVideo := win.HWND(0)
	ret, _, _ := syscall.Syscall(
		(*IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).GetVideoWindow, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&hwndVideo)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFVideoDisplayControl.GetVideoWindow"))
	}
	return hwndVideo
}

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-repaintvideo
func (me *IMFVideoDisplayControl) RepaintVideo() {
	ret, _, _ := syscall.Syscall(
		(*IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).RepaintVideo, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMFVideoDisplayControl.RepaintVideo"))
	}
}
