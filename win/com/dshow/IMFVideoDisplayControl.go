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
type IMFVideoDisplayControl struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IMFVideoDisplayControl.Release().
//
// Example:
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
func NewIMFVideoDisplayControl(base com.IUnknown) IMFVideoDisplayControl {
	return IMFVideoDisplayControl{IUnknown: base}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getaspectratiomode
func (me *IMFVideoDisplayControl) GetAspectRatioMode() dshowco.MFVideoARMode {
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

// Returns the minimum and maximum ideal sizes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getidealvideosize
func (me *IMFVideoDisplayControl) GetIdealVideoSize() (min, max win.SIZE) {
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

// Returns video rectangle and aspect ratio.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getnativevideosize
func (me *IMFVideoDisplayControl) GetNativeVideoSize() (size, aspectRatio win.SIZE) {
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getvideoposition
func (me *IMFVideoDisplayControl) GetVideoPosition() (source MFVideoNormalizedRect, dest win.RECT) {
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-getvideowindow
func (me *IMFVideoDisplayControl) GetVideoWindow() win.HWND {
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-repaintvideo
func (me *IMFVideoDisplayControl) RepaintVideo() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMFVideoDisplayControl)(unsafe.Pointer(*me.Ptr())).RepaintVideo, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setaspectratiomode
func (me *IMFVideoDisplayControl) SetAspectRatioMode(
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setvideoposition
func (me *IMFVideoDisplayControl) SetVideoPosition(
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setvideowindow
func (me *IMFVideoDisplayControl) SetVideoWindow(hwndVideo win.HWND) error {
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
