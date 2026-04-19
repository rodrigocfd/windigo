//go:build windows

package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [DrawIcon] function.
//
// [DrawIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawicon
func (hdc HDC) DrawIcon(x, y int, hIcon HICON) error {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_DrawIcon, "DrawIcon"),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(hIcon))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_DrawIcon *syscall.Proc

// [DrawIconEx] function.
//
// [DrawIconEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawiconex
func (hdc HDC) DrawIconEx(
	pos POINT,
	hIcon HICON,
	size SIZE,
	frameIndexIfCursor int,
	hbrFlickerFree HBRUSH,
	diFlags co.DI,
) error {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_DrawIconEx, "DrawIconEx"),
		uintptr(hdc),
		uintptr(pos.X),
		uintptr(pos.Y),
		uintptr(hIcon),
		uintptr(size.Cx),
		uintptr(size.Cy),
		uintptr(uint32(frameIndexIfCursor)),
		uintptr(hbrFlickerFree),
		uintptr(diFlags))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_DrawIconEx *syscall.Proc

// [DrawText] function.
//
// [DrawText]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawtextw
func (hdc HDC) DrawText(text string, pRc *RECT, format co.DT) (int, error) {
	var wText wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.User.Load(&_user_DrawTextW, "DrawTextW"),
		uintptr(hdc),
		uintptr(wText.AllowEmpty(text)),
		uintptr(int32(len(text))),
		uintptr(unsafe.Pointer(pRc)),
		uintptr(format))
	if ret == 0 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return int(ret), nil
}

var _user_DrawTextW *syscall.Proc

// [DrawTextEx] function.
//
// [DrawTextEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawtextexw
func (hdc HDC) DrawTextEx(text string, pRc *RECT, format co.DT, pDtp *DRAWTEXTPARAMS) (int, error) {
	var wText wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.User.Load(&_user_DrawTextExW, "DrawTextExW"),
		uintptr(hdc),
		uintptr(wText.AllowEmpty(text)),
		uintptr(int32(len(text))),
		uintptr(unsafe.Pointer(pRc)),
		uintptr(format),
		uintptr(unsafe.Pointer(pDtp)))
	if ret == 0 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return int(ret), nil
}

var _user_DrawTextExW *syscall.Proc

// [EnumDisplayMonitors] function.
//
// [EnumDisplayMonitors]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func (hdc HDC) EnumDisplayMonitors(pRcClip *RECT) ([]EnumDisplayMonitorsInfo, error) {
	pPack := &callbackPack_EnumDisplayMonitors{
		arr: make([]EnumDisplayMonitorsInfo, 0),
	}
	ret, _, _ := syscall.SyscallN(
		dll.User.Load(&_user_EnumDisplayMonitors, "EnumDisplayMonitors"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(pRcClip)),
		callbackGet_EnumDisplayMonitors(),
		uintptr(unsafe.Pointer(pPack)))
	runtime.KeepAlive(pPack)
	if ret == 0 {
		return nil, co.ERROR_INVALID_PARAMETER
	}
	return pPack.arr, nil
}

var _user_EnumDisplayMonitors *syscall.Proc

// Returned by [HDC.EnumDisplayMonitors].
type EnumDisplayMonitorsInfo struct {
	HMon   HMONITOR
	HdcMon HDC
	Rc     RECT
}

type callbackPack_EnumDisplayMonitors struct{ arr []EnumDisplayMonitorsInfo }

var callback_EnumDisplayMonitors uintptr

func callbackGet_EnumDisplayMonitors() uintptr {
	if callback_EnumDisplayMonitors == 0 {
		callback_EnumDisplayMonitors = syscall.NewCallback(
			func(hMon HMONITOR, hdcMon HDC, rcMon *RECT, lParam LPARAM) uintptr {
				pPack := (*callbackPack_EnumDisplayMonitors)(unsafe.Pointer(lParam))
				pPack.arr = append(pPack.arr, EnumDisplayMonitorsInfo{hMon, hdcMon, *rcMon})
				return 1
			},
		)
	}
	return callback_EnumDisplayMonitors
}

// [FillRect] function.
//
// [FillRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-fillrect
func (hdc HDC) FillRect(pRc *RECT, hBrush HBRUSH) error {
	ret, _, _ := syscall.SyscallN(
		dll.User.Load(&_user_FillRect, "FillRect"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(pRc)),
		uintptr(hBrush))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _user_FillRect *syscall.Proc

// [FrameRect] function.
//
// [FrameRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-framerect
func (hdc HDC) FrameRect(pRc *RECT, hBrush HBRUSH) error {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_FrameRect, "FrameRect"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(pRc)),
		uintptr(hBrush))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_FrameRect *syscall.Proc

// [InvertRect] function.
//
// [InvertRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invertrect
func (hdc HDC) InvertRect(pRc *RECT) error {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_InvertRect, "InvertRect"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(pRc)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_InvertRect *syscall.Proc

// [PaintDesktop] function.
//
// [PaintDesktop]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-paintdesktop
func (hdc HDC) PaintDesktop() error {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_PaintDesktop, "PaintDesktop"),
		uintptr(hdc))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_PaintDesktop *syscall.Proc

// [WindowFromDC] function.
//
// [WindowFromDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-windowfromdc
func (hdc HDC) WindowFromDC() HWND {
	ret, _, _ := syscall.SyscallN(
		dll.User.Load(&_user_WindowFromDC, "WindowFromDC"),
		uintptr(hdc))
	return HWND(ret)
}

var _user_WindowFromDC *syscall.Proc
