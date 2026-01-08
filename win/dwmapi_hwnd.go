//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [DwmExtendFrameIntoClientArea] function.
//
// [DwmExtendFrameIntoClientArea]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea
func (hWnd HWND) DwmExtendFrameIntoClientArea(marginsInset *MARGINS) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_dwmapi_DwmExtendFrameIntoClientArea, "DwmExtendFrameIntoClientArea"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(marginsInset)))
	return utl.HresultToError(ret)
}

var _dwmapi_DwmExtendFrameIntoClientArea *syscall.Proc

// [DwmGetWindowAttribute] function.
//
// Example:
//
//	var hWnd win.HWND // initialized somewhere
//
//	attr, _ := hWnd.DwmGetWindowAttribute(co.DWMWA_EXTENDED_FRAME_BOUNDS)
//	rc, _ := attr.ExtendedFrameBounds()
//
// [DwmGetWindowAttribute]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmgetwindowattribute
func (hWnd HWND) DwmGetWindowAttribute(attr co.DWMWA) (DwmAttr, error) {
	var dwBuf uint32
	var rcBuf RECT
	var ptrBuf unsafe.Pointer
	var szBuf uintptr

	switch attr {
	case co.DWMWA_CAPTION_BUTTON_BOUNDS,
		co.DWMWA_EXTENDED_FRAME_BOUNDS:
		ptrBuf, szBuf = unsafe.Pointer(&rcBuf), unsafe.Sizeof(rcBuf)
	default:
		ptrBuf, szBuf = unsafe.Pointer(&dwBuf), unsafe.Sizeof(dwBuf)
	}

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_dwmapi_DwmGetWindowAttribute, "DwmGetWindowAttribute"),
		uintptr(hWnd),
		uintptr(attr),
		uintptr(ptrBuf),
		szBuf)
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return DwmAttr{}, hr
	}
	return dwmAttrFromRaw(attr, dwBuf, rcBuf), nil
}

var _dwmapi_DwmGetWindowAttribute *syscall.Proc

// [DwmInvalidateIconicBitmaps] function.
//
// [DwmInvalidateIconicBitmaps]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwminvalidateiconicbitmaps
func (hWnd HWND) DwmInvalidateIconicBitmaps() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_dwmapi_DwmInvalidateIconicBitmaps, "DwmInvalidateIconicBitmaps"),
		uintptr(hWnd))
	return utl.HresultToError(ret)
}

var _dwmapi_DwmInvalidateIconicBitmaps *syscall.Proc

// [DwmModifyPreviousDxFrameDuration] function.
//
// Panics if numRefreshes is negative.
//
// [DwmModifyPreviousDxFrameDuration]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmmodifypreviousdxframeduration
func (hWnd HWND) DwmModifyPreviousDxFrameDuration(numRefreshes int, relative bool) error {
	utl.PanicNeg(numRefreshes)
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_dwmapi_DwmModifyPreviousDxFrameDuration, "DwmModifyPreviousDxFrameDuration"),
		uintptr(hWnd),
		uintptr(int32(numRefreshes)),
		utl.BoolToUintptr(relative))
	return utl.HresultToError(ret)
}

var _dwmapi_DwmModifyPreviousDxFrameDuration *syscall.Proc

// [DwmSetIconicLivePreviewBitmap] function.
//
// [DwmSetIconicLivePreviewBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmseticoniclivepreviewbitmap
func (hWnd HWND) DwmSetIconicLivePreviewBitmap(hBmp HBITMAP, ptClient POINT, sitFlags co.DWM_SIT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_dwmapi_DwmSetIconicLivePreviewBitmap, "DwmSetIconicLivePreviewBitmap"),
		uintptr(hWnd),
		uintptr(hBmp),
		uintptr(unsafe.Pointer(&ptClient)),
		uintptr(sitFlags))
	return utl.HresultToError(ret)
}

var _dwmapi_DwmSetIconicLivePreviewBitmap *syscall.Proc

// [DwmSetIconicThumbnail] function.
//
// [DwmSetIconicThumbnail]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmseticonicthumbnail
func (hWnd HWND) DwmSetIconicThumbnail(hBmp HBITMAP, sitFlags co.DWM_SIT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_dwmapi_DwmSetIconicThumbnail, "DwmSetIconicThumbnail"),
		uintptr(hWnd),
		uintptr(hBmp),
		uintptr(sitFlags))
	return utl.HresultToError(ret)
}

var _dwmapi_DwmSetIconicThumbnail *syscall.Proc

// [DwmSetWindowAttribute] function.
//
// Example:
//
//	var hWnd win.HWND // initialized somewhere
//
//	attr := win.DwmAttrExtendedFrameBounds(
//		win.RECT{Left: 10, Top: 10, Right: 10, Bottom: 10},
//	)
//	_ = hWnd.DwmSetWindowAttribute(attr)
//
// [DwmSetWindowAttribute]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmsetwindowattribute
func (hWnd HWND) DwmSetWindowAttribute(attr DwmAttr) error {
	var ptrBuf unsafe.Pointer
	var szBuf uintptr

	switch attr.tag {
	case co.DWMWA_CAPTION_BUTTON_BOUNDS,
		co.DWMWA_EXTENDED_FRAME_BOUNDS:
		ptrBuf, szBuf = unsafe.Pointer(&attr.rc), unsafe.Sizeof(attr.rc)
	default:
		ptrBuf, szBuf = unsafe.Pointer(&attr.dw), unsafe.Sizeof(attr.dw)
	}

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_dwmapi_DwmSetWindowAttribute, "DwmSetWindowAttribute"),
		uintptr(hWnd),
		uintptr(attr.tag),
		uintptr(ptrBuf),
		szBuf)
	return utl.HresultToError(ret)
}

var _dwmapi_DwmSetWindowAttribute *syscall.Proc
