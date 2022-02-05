package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea
func (hWnd HWND) DwmExtendFrameIntoClientArea(marginsInset *MARGINS) {
	ret, _, _ := syscall.Syscall(proc.DwmExtendFrameIntoClientArea.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(marginsInset)), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmgetwindowattribute
// func (hWnd HWND) DwmGetWindowAttribute(attr co.DWMWA_GET) {
// 	ret, _, _ := syscall.Syscall6(proc.DwmGetWindowAttribute.Addr(), 4,
// 		uintptr(hWnd), uintptr(attr), 0, 0,
// 		0, 0)
// }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmseticoniclivepreviewbitmap
func (hWnd HWND) DwmSetIconicLivePreviewBitmap(
	hBmp HBITMAP, ptClient POINT, sitFlags co.DWM_SIT) {

	ret, _, _ := syscall.Syscall6(proc.DwmSetIconicLivePreviewBitmap.Addr(), 4,
		uintptr(hWnd), uintptr(hBmp), uintptr(unsafe.Pointer(&ptClient)),
		uintptr(sitFlags), 0, 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmseticonicthumbnail
func (hWnd HWND) DwmSetIconicThumbnail(hBmp HBITMAP, sitFlags co.DWM_SIT) {
	ret, _, _ := syscall.Syscall(proc.DwmSetIconicThumbnail.Addr(), 3,
		uintptr(hWnd), uintptr(hBmp), uintptr(sitFlags))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmsetwindowattribute
// func (hWnd HWND) DwmSetWindowAttribute(attribute co.DWMWA_SET) {

// }
