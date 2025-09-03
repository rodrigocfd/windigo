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
		dll.Load(dll.USER32, &_DrawIcon, "DrawIcon"),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(hIcon))
	return utl.ZeroAsGetLastError(ret, err)
}

var _DrawIcon *syscall.Proc

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
		dll.Load(dll.USER32, &_DrawIconEx, "DrawIconEx"),
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

var _DrawIconEx *syscall.Proc

// [EnumDisplayMonitors] function.
//
// [EnumDisplayMonitors]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func (hdc HDC) EnumDisplayMonitors(rcClip *RECT) ([]EnumDisplayMonitorsInfo, error) {
	pPack := &_EnumDisplayMonitorsPack{
		arr: make([]EnumDisplayMonitorsInfo, 0),
	}
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_EnumDisplayMonitors, "EnumDisplayMonitors"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(rcClip)),
		enumDisplayMonitorsCallback(),
		uintptr(unsafe.Pointer(pPack)))
	runtime.KeepAlive(pPack)
	if ret == 0 {
		return nil, co.ERROR_INVALID_PARAMETER
	}
	return pPack.arr, nil
}

var _EnumDisplayMonitors *syscall.Proc

type (
	_EnumDisplayMonitorsPack struct{ arr []EnumDisplayMonitorsInfo }

	// Returned by [HDC.EnumDisplayMonitors].
	EnumDisplayMonitorsInfo struct {
		HMon   HMONITOR
		HdcMon HDC
		Rc     RECT
	}
)

var _enumDisplayMonitorsCallback uintptr

func enumDisplayMonitorsCallback() uintptr {
	if _enumDisplayMonitorsCallback != 0 {
		return _enumDisplayMonitorsCallback
	}

	_enumDisplayMonitorsCallback = syscall.NewCallback(
		func(hMon HMONITOR, hdcMon HDC, rcMon *RECT, lParam LPARAM) uintptr {
			pPack := (*_EnumDisplayMonitorsPack)(unsafe.Pointer(lParam))
			pPack.arr = append(pPack.arr, EnumDisplayMonitorsInfo{hMon, hdcMon, *rcMon})
			return 1
		},
	)
	return _enumDisplayMonitorsCallback
}

// [FrameRect] function.
//
// [FrameRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-framerect
func (hdc HDC) FrameRect(rc *RECT, hBrush HBRUSH) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_FrameRect, "FrameRect"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(rc)),
		uintptr(hBrush))
	return utl.ZeroAsGetLastError(ret, err)
}

var _FrameRect *syscall.Proc

// [InvertRect] function.
//
// [InvertRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invertrect
func (hdc HDC) InvertRect(rc *RECT) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_InvertRect, "InvertRect"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(rc)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _InvertRect *syscall.Proc

// [PaintDesktop] function.
//
// [PaintDesktop]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-paintdesktop
func (hdc HDC) PaintDesktop() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_PaintDesktop, "PaintDesktop"),
		uintptr(hdc))
	return utl.ZeroAsGetLastError(ret, err)
}

var _PaintDesktop *syscall.Proc

// [WindowFromDC] function.
//
// [WindowFromDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-windowfromdc
func (hdc HDC) WindowFromDC() HWND {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_WindowFromDC, "WindowFromDC"),
		uintptr(hdc))
	return HWND(ret)
}

var _WindowFromDC *syscall.Proc
