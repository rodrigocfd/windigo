package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

type _IMFVideoDisplayControlVtbl struct {
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

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nn-evr-imfvideodisplaycontrol
type IMFVideoDisplayControl struct {
	win.IUnknown // Base IUnknown.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getaspectratiomode
func (me *IMFVideoDisplayControl) GetAspectRatioMode() co.MFVideoARMode {
	aspectRatioMode := co.MFVideoARMode_None
	ret, _, _ := syscall.Syscall(
		(*_IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).GetAspectRatioMode, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&aspectRatioMode)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return aspectRatioMode
}

// Returns the minimum and maximum ideal sizes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getidealvideosize
func (me *IMFVideoDisplayControl) GetIdealVideoSize() (min, max win.SIZE) {
	ret, _, _ := syscall.Syscall(
		(*_IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).GetIdealVideoSize, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&min)),
		uintptr(unsafe.Pointer(&max)))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return
}

// Returns video rectangle and aspect ratio.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getnativevideosize
func (me *IMFVideoDisplayControl) GetNativeVideoSize() (size, aspectRatio win.SIZE) {
	ret, _, _ := syscall.Syscall(
		(*_IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).GetNativeVideoSize, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&size)),
		uintptr(unsafe.Pointer(&aspectRatio)))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getvideoposition
func (me *IMFVideoDisplayControl) GetVideoPosition() (source MFVideoNormalizedRect, dest win.RECT) {
	ret, _, _ := syscall.Syscall(
		(*_IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).GetVideoPosition, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&source)),
		uintptr(unsafe.Pointer(&dest)))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getvideowindow
func (me *IMFVideoDisplayControl) GetVideoWindow() win.HWND {
	hwndVideo := win.HWND(0)
	ret, _, _ := syscall.Syscall(
		(*_IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).GetVideoWindow, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&hwndVideo)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return hwndVideo
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-repaintvideo
func (me *IMFVideoDisplayControl) RepaintVideo() {
	ret, _, _ := syscall.Syscall(
		(*_IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).RepaintVideo, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setaspectratiomode
func (me *IMFVideoDisplayControl) SetAspectRatioMode(mode co.MFVideoARMode) error {
	ret, _, _ := syscall.Syscall(
		(*_IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).SetAspectRatioMode, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(mode), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		return lerr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setvideoposition
func (me *IMFVideoDisplayControl) SetVideoPosition(
	pnrcSource *MFVideoNormalizedRect, prcDest *win.RECT) {

	ret, _, _ := syscall.Syscall(
		(*_IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).SetVideoPosition, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pnrcSource)),
		uintptr(unsafe.Pointer(prcDest)))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setvideowindow
func (me *IMFVideoDisplayControl) SetVideoWindow(hwndVideo win.HWND) error {
	ret, _, _ := syscall.Syscall(
		(*_IMFVideoDisplayControlVtbl)(unsafe.Pointer(*me.Ppv)).SetVideoWindow, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndVideo), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		return lerr
	}
	return nil
}
