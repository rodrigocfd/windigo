//go:build windows

package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [DrawIcon] function.
//
// [DrawIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawicon
func (hdc HDC) DrawIcon(x, y int, hIcon HICON) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_DrawIcon, "DrawIcon"),
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
		dll.Load(dll.USER32, &_user_DrawIconEx, "DrawIconEx"),
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

// [EnumDisplayMonitors] function.
//
// [EnumDisplayMonitors]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func (hdc HDC) EnumDisplayMonitors(rcClip *RECT) ([]EnumDisplayMonitorsInfo, error) {
	pPack := &callbackPack_EnumDisplayMonitors{
		arr: make([]EnumDisplayMonitorsInfo, 0),
	}
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_EnumDisplayMonitors, "EnumDisplayMonitors"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(rcClip)),
		callbackGet_EnumDisplayMonitors(),
		uintptr(unsafe.Pointer(pPack)))
	runtime.KeepAlive(pPack)
	if ret == 0 {
		return nil, co.ERROR_INVALID_PARAMETER
	}
	return pPack.arr, nil
}

var _user_EnumDisplayMonitors *syscall.Proc

type (
	// Returned by [HDC.EnumDisplayMonitors].
	EnumDisplayMonitorsInfo struct {
		HMon   HMONITOR
		HdcMon HDC
		Rc     RECT
	}

	callbackPack_EnumDisplayMonitors struct{ arr []EnumDisplayMonitorsInfo }
)

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

// [FrameRect] function.
//
// [FrameRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-framerect
func (hdc HDC) FrameRect(rc *RECT, hBrush HBRUSH) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_FrameRect, "FrameRect"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(rc)),
		uintptr(hBrush))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_FrameRect *syscall.Proc

// [InvertRect] function.
//
// [InvertRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invertrect
func (hdc HDC) InvertRect(rc *RECT) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_InvertRect, "InvertRect"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(rc)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_InvertRect *syscall.Proc

// [PaintDesktop] function.
//
// [PaintDesktop]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-paintdesktop
func (hdc HDC) PaintDesktop() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_PaintDesktop, "PaintDesktop"),
		uintptr(hdc))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_PaintDesktop *syscall.Proc

// [WindowFromDC] function.
//
// [WindowFromDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-windowfromdc
func (hdc HDC) WindowFromDC() HWND {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_WindowFromDC, "WindowFromDC"),
		uintptr(hdc))
	return HWND(ret)
}

var _user_WindowFromDC *syscall.Proc
