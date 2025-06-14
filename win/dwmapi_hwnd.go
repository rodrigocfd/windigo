//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// [DwmExtendFrameIntoClientArea] function.
//
// [DwmExtendFrameIntoClientArea]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea
func (hWnd HWND) DwmExtendFrameIntoClientArea(marginsInset *MARGINS) error {
	ret, _, _ := syscall.SyscallN(dll.Dwmapi(dll.PROC_DwmExtendFrameIntoClientArea),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(marginsInset)))
	return utl.ErrorAsHResult(ret)
}

// [DwmGetWindowAttribute] function.
//
// # Example
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

	ret, _, _ := syscall.SyscallN(dll.Dwmapi(dll.PROC_DwmGetWindowAttribute),
		uintptr(hWnd),
		uintptr(attr),
		uintptr(ptrBuf),
		szBuf)
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return DwmAttr{}, hr
	}
	return dwmAttrFromRaw(attr, dwBuf, rcBuf), nil
}

// [DwmInvalidateIconicBitmaps] function.
//
// [DwmInvalidateIconicBitmaps]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwminvalidateiconicbitmaps
func (hWnd HWND) DwmInvalidateIconicBitmaps() error {
	ret, _, _ := syscall.SyscallN(dll.Dwmapi(dll.PROC_DwmInvalidateIconicBitmaps),
		uintptr(hWnd))
	return utl.ErrorAsHResult(ret)
}

// [DwmSetIconicLivePreviewBitmap] function.
//
// [DwmSetIconicLivePreviewBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmseticoniclivepreviewbitmap
func (hWnd HWND) DwmSetIconicLivePreviewBitmap(hBmp HBITMAP, ptClient POINT, sitFlags co.DWM_SIT) error {
	ret, _, _ := syscall.SyscallN(dll.Dwmapi(dll.PROC_DwmSetIconicLivePreviewBitmap),
		uintptr(hWnd),
		uintptr(hBmp),
		uintptr(unsafe.Pointer(&ptClient)),
		uintptr(sitFlags))
	return utl.ErrorAsHResult(ret)
}

// [DwmSetIconicThumbnail] function.
//
// [DwmSetIconicThumbnail]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmseticonicthumbnail
func (hWnd HWND) DwmSetIconicThumbnail(hBmp HBITMAP, sitFlags co.DWM_SIT) error {
	ret, _, _ := syscall.SyscallN(dll.Dwmapi(dll.PROC_DwmSetIconicThumbnail),
		uintptr(hWnd),
		uintptr(hBmp),
		uintptr(sitFlags))
	return utl.ErrorAsHResult(ret)
}

// [DwmSetWindowAttribute] function.
//
// # Example
//
//	var hWnd win.HWND // initialized somewhere
//
//	attr := win.DwmAttrExtendedFrameBounds(RECT{10, 10, 10, 10})
//	hWnd.DwmSetWindowAttribute(attr)
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

	ret, _, _ := syscall.SyscallN(dll.Dwmapi(dll.PROC_DwmSetWindowAttribute),
		uintptr(hWnd),
		uintptr(attr.tag),
		uintptr(ptrBuf),
		szBuf)
	return utl.ErrorAsHResult(ret)
}
