//go:build windows

package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nn-evr-imfvideodisplaycontrol
type IMFVideoDisplayControl interface {
	com.IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getaspectratiomode
	GetAspectRatioMode() dshowco.MFVideoARMode

	// Returns the minimum and maximum ideal sizes.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getidealvideosize
	GetIdealVideoSize() (min, max win.SIZE)

	// Returns video rectangle and aspect ratio.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getnativevideosize
	GetNativeVideoSize() (size, aspectRatio win.SIZE)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getvideoposition
	GetVideoPosition() (source MFVideoNormalizedRect, dest win.RECT)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getvideowindow
	GetVideoWindow() win.HWND

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-repaintvideo
	RepaintVideo()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setaspectratiomode
	SetAspectRatioMode(mode dshowco.MFVideoARMode) error

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setvideoposition
	SetVideoPosition(nrcSource *MFVideoNormalizedRect, rcDest *win.RECT)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setvideowindow
	SetVideoWindow(hwndVideo win.HWND) error
}

type _IMFVideoDisplayControl struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IMFVideoDisplayControl.Release().
//
// Example:
//
//		var gs dshow.IMFGetService // initialized somewhere
//
//		vdc := dshow.NewIMFVideoDisplayControl(
//			gs.GetService(
//				win.NewGuidFromClsid(dshowco.CLSID_MR_VideoRenderService),
//				win.NewGuidFromIid(dshowco.IID_IMFVideoDisplayControl),
//			),
//		)
//		defer vdc.Release()
func NewIMFVideoDisplayControl(base com.IUnknown) IMFVideoDisplayControl {
	return &_IMFVideoDisplayControl{IUnknown: base}
}

func (me *_IMFVideoDisplayControl) GetAspectRatioMode() dshowco.MFVideoARMode {
	var aspectRatioMode dshowco.MFVideoARMode
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMFVideoDisplayControl)(unsafe.Pointer(*me.Ptr())).GetAspectRatioMode, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&aspectRatioMode)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return aspectRatioMode
	} else {
		panic(hr)
	}
}

func (me *_IMFVideoDisplayControl) GetIdealVideoSize() (min, max win.SIZE) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMFVideoDisplayControl)(unsafe.Pointer(*me.Ptr())).GetIdealVideoSize, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&min)),
		uintptr(unsafe.Pointer(&max)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return
	} else {
		panic(hr)
	}
}

func (me *_IMFVideoDisplayControl) GetNativeVideoSize() (size, aspectRatio win.SIZE) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMFVideoDisplayControl)(unsafe.Pointer(*me.Ptr())).GetNativeVideoSize, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&size)),
		uintptr(unsafe.Pointer(&aspectRatio)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return
	} else {
		panic(hr)
	}
}

func (me *_IMFVideoDisplayControl) GetVideoPosition() (source MFVideoNormalizedRect, dest win.RECT) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMFVideoDisplayControl)(unsafe.Pointer(*me.Ptr())).GetVideoPosition, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&source)),
		uintptr(unsafe.Pointer(&dest)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return
	} else {
		panic(hr)
	}
}

func (me *_IMFVideoDisplayControl) GetVideoWindow() win.HWND {
	var hwndVideo win.HWND
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMFVideoDisplayControl)(unsafe.Pointer(*me.Ptr())).GetVideoWindow, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&hwndVideo)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return hwndVideo
	} else {
		panic(hr)
	}
}

func (me *_IMFVideoDisplayControl) RepaintVideo() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMFVideoDisplayControl)(unsafe.Pointer(*me.Ptr())).RepaintVideo, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IMFVideoDisplayControl) SetAspectRatioMode(
	mode dshowco.MFVideoARMode) error {

	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMFVideoDisplayControl)(unsafe.Pointer(*me.Ptr())).SetAspectRatioMode, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(mode), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

func (me *_IMFVideoDisplayControl) SetVideoPosition(
	nrcSource *MFVideoNormalizedRect, rcDest *win.RECT) {

	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMFVideoDisplayControl)(unsafe.Pointer(*me.Ptr())).SetVideoPosition, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(nrcSource)),
		uintptr(unsafe.Pointer(rcDest)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IMFVideoDisplayControl) SetVideoWindow(hwndVideo win.HWND) error {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMFVideoDisplayControl)(unsafe.Pointer(*me.Ptr())).SetVideoWindow, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndVideo), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}
