package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
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

// Example:
//
//  var hwnd win.HWND // initialized somewhere
//
//  isEnabled := hwnd.DwmGetWindowAttribute(
//    co.DWMWA_GET_NCRENDERING_ENABLED).(bool)
//
//  rect := hwnd.DwmGetWindowAttribute(
//    co.DWMWA_GET_CAPTION_BUTTON_BOUNDS).(win.RECT)
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmgetwindowattribute
func (hWnd HWND) DwmGetWindowAttribute(attr co.DWMWA_GET) interface{} {
	var boolBuf int32
	var rectBuf RECT
	var cloakBuf co.DWM_CLOAKED
	var ptrRaw unsafe.Pointer
	var cbSize uintptr

	switch attr {
	case co.DWMWA_GET_NCRENDERING_ENABLED:
		ptrRaw = unsafe.Pointer(&boolBuf)
		cbSize = unsafe.Sizeof(boolBuf)
	case co.DWMWA_GET_CAPTION_BUTTON_BOUNDS:
		ptrRaw = unsafe.Pointer(&rectBuf)
		cbSize = unsafe.Sizeof(rectBuf)
	case co.DWMWA_GET_EXTENDED_FRAME_BOUNDS:
		ptrRaw = unsafe.Pointer(&rectBuf)
		cbSize = unsafe.Sizeof(rectBuf)
	case co.DWMWA_GET_CLOAKED:
		ptrRaw = unsafe.Pointer(&cloakBuf)
		cbSize = unsafe.Sizeof(cloakBuf)
	}
	defer runtime.KeepAlive(ptrRaw)

	ret, _, _ := syscall.Syscall6(proc.DwmGetWindowAttribute.Addr(), 4,
		uintptr(hWnd), uintptr(attr), uintptr(ptrRaw), cbSize,
		0, 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}

	switch attr {
	case co.DWMWA_GET_NCRENDERING_ENABLED:
		return boolBuf != 0
	case co.DWMWA_GET_CAPTION_BUTTON_BOUNDS:
		return rectBuf
	case co.DWMWA_GET_EXTENDED_FRAME_BOUNDS:
		return rectBuf
	case co.DWMWA_GET_CLOAKED:
		return cloakBuf
	default:
		panic("Invalid co.DWMWA_GET value.")
	}
}

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

// Example:
//
//  var hwnd win.HWND // initialized somewhere
//
//  hwnd.DwmSetWindowAttribute(
//    co.DWMWA_SET_NCRENDERING_POLICY, co.DWMNCRP_DISABLED)
//
//  hwnd.DwmSetWindowAttribute(
//    co.DWMWA_SET_TRANSITIONS_FORCEDISABLED, true)
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmsetwindowattribute
func (hWnd HWND) DwmSetWindowAttribute(attr co.DWMWA_SET, val interface{}) {
	var ptrRaw unsafe.Pointer
	var cbSize uintptr

	switch attr {
	case co.DWMWA_SET_NCRENDERING_POLICY:
		if val, ok := val.(co.DWMNCRP); !ok {
			panic("DWMWA_SET_NCRENDERING_POLICY must have a co.DWMNCRP.")
		} else {
			ptrRaw = unsafe.Pointer(&val)
			cbSize = unsafe.Sizeof(co.DWMNCRP_DISABLED)
		}
	case co.DWMWA_SET_TRANSITIONS_FORCEDISABLED:
		if val, ok := val.(bool); !ok {
			panic("DWMWA_SET_TRANSITIONS_FORCEDISABLED must have a bool.")
		} else {
			boolBuf := int32(util.BoolToUintptr(val))
			ptrRaw = unsafe.Pointer(&boolBuf)
			cbSize = unsafe.Sizeof(boolBuf)
		}
	case co.DWMWA_SET_ALLOW_NCPAINT:
		if val, ok := val.(bool); !ok {
			panic("DWMWA_SET_ALLOW_NCPAINT must have a bool.")
		} else {
			boolBuf := int32(util.BoolToUintptr(val))
			ptrRaw = unsafe.Pointer(&boolBuf)
			cbSize = unsafe.Sizeof(boolBuf)
		}
	case co.DWMWA_SET_FORCE_ICONIC_REPRESENTATION:
		if val, ok := val.(bool); !ok {
			panic("DWMWA_SET_FORCE_ICONIC_REPRESENTATION must have a bool.")
		} else {
			boolBuf := int32(util.BoolToUintptr(val))
			ptrRaw = unsafe.Pointer(&boolBuf)
			cbSize = unsafe.Sizeof(boolBuf)
		}
	case co.DWMWA_SET_FLIP3D_POLICY:
		if val, ok := val.(co.DWMFLIP3D); !ok {
			panic("DWMWA_SET_FLIP3D_POLICY must have a co.DWMFLIP3D.")
		} else {
			ptrRaw = unsafe.Pointer(&val)
			cbSize = unsafe.Sizeof(co.DWMFLIP3D_DEFAULT)
		}
	case co.DWMWA_SET_HAS_ICONIC_BITMAP:
		if val, ok := val.(bool); !ok {
			panic("DWMWA_SET_HAS_ICONIC_BITMAP must have a bool.")
		} else {
			boolBuf := int32(util.BoolToUintptr(val))
			ptrRaw = unsafe.Pointer(&boolBuf)
			cbSize = unsafe.Sizeof(boolBuf)
		}
	case co.DWMWA_SET_DISALLOW_PEEK:
		if val, ok := val.(bool); !ok {
			panic("DWMWA_SET_DISALLOW_PEEK must have a bool.")
		} else {
			boolBuf := int32(util.BoolToUintptr(val))
			ptrRaw = unsafe.Pointer(&boolBuf)
			cbSize = unsafe.Sizeof(boolBuf)
		}
	case co.DWMWA_SET_EXCLUDED_FROM_PEEK:
		if val, ok := val.(bool); !ok {
			panic("DWMWA_SET_EXCLUDED_FROM_PEEK must have a bool.")
		} else {
			boolBuf := int32(util.BoolToUintptr(val))
			ptrRaw = unsafe.Pointer(&boolBuf)
			cbSize = unsafe.Sizeof(boolBuf)
		}
	case co.DWMWA_SET_CLOAK:
		ptrRaw = nil
		cbSize = 0
	case co.DWMWA_SET_FREEZE_REPRESENTATION:
		ptrRaw = nil
		cbSize = 0
	}

	ret, _, _ := syscall.Syscall6(proc.DwmSetWindowAttribute.Addr(), 4,
		uintptr(hWnd), uintptr(attr), uintptr(ptrRaw), cbSize,
		0, 0)
	runtime.KeepAlive(ptrRaw)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
