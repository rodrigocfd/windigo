//go:build windows

package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win/co"
)

// [DrawIcon] function.
//
// [DrawIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawicon
func (hdc HDC) DrawIcon(x, y int, hIcon HICON) error {
	ret, _, err := syscall.SyscallN(_DrawIcon.Addr(),
		uintptr(hdc), uintptr(x), uintptr(y), uintptr(hIcon))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _DrawIcon = dll.User32.NewProc("DrawIcon")

// [DrawIconEx] function.
//
// [DrawIconEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawiconex
func (hdc HDC) DrawIconEx(
	pos POINT,
	hIcon HICON,
	size SIZE,
	frameIndex uint,
	hbrFlickerFree HBRUSH,
	diFlags co.DI,
) error {
	ret, _, err := syscall.SyscallN(_DrawIconEx.Addr(),
		uintptr(hdc), uintptr(pos.X), uintptr(pos.Y), uintptr(hIcon),
		uintptr(size.Cx), uintptr(size.Cy), uintptr(frameIndex),
		uintptr(hbrFlickerFree), uintptr(diFlags))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _DrawIconEx = dll.User32.NewProc("DrawIconEx")

// [EnumDisplayMonitors] function.
//
// [EnumDisplayMonitors]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func (hdc HDC) EnumDisplayMonitors(rcClip *RECT) ([]EnumDisplayMonitorsInfo, error) {
	pPack := &_EnumDisplayMonitorsPack{
		arr: make([]EnumDisplayMonitorsInfo, 0),
	}
	ret, _, _ := syscall.SyscallN(_EnumDisplayMonitors.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(rcClip)),
		enumDisplayMonitorsCallback(), uintptr(unsafe.Pointer(pPack)))
	runtime.KeepAlive(pPack)
	if ret == 0 {
		return nil, co.ERROR_INVALID_PARAMETER
	}
	return pPack.arr, nil
}

type (
	_EnumDisplayMonitorsPack struct{ arr []EnumDisplayMonitorsInfo }

	// Returned by [HDC.EnumDisplayMonitors].
	EnumDisplayMonitorsInfo struct {
		HMon   HMONITOR
		HdcMon HDC
		Rc     RECT
	}
)

var (
	_EnumDisplayMonitors         = dll.User32.NewProc("EnumDisplayMonitors")
	_enumDisplayMonitorsCallback uintptr
)

func enumDisplayMonitorsCallback() uintptr {
	if _enumDisplayMonitorsCallback == 0 {
		_enumDisplayMonitorsCallback = syscall.NewCallback(
			func(hMon HMONITOR, hdcMon HDC, rcMon *RECT, lParam LPARAM) uintptr {
				pPack := (*_EnumDisplayMonitorsPack)(unsafe.Pointer(lParam))
				pPack.arr = append(pPack.arr, EnumDisplayMonitorsInfo{hMon, hdcMon, *rcMon})
				return 1
			},
		)
	}
	return _enumDisplayMonitorsCallback
}

// [FrameRect] function.
//
// [FrameRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-framerect
func (hdc HDC) FrameRect(rc *RECT, hBrush HBRUSH) error {
	ret, _, err := syscall.SyscallN(_FrameRect.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(rc)), uintptr(hBrush))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _FrameRect = dll.User32.NewProc("FrameRect")

// [InvertRect] function.
//
// [InvertRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invertrect
func (hdc HDC) InvertRect(rc *RECT) error {
	ret, _, err := syscall.SyscallN(_InvertRect.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(rc)))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _InvertRect = dll.User32.NewProc("InvertRect")

// [PaintDesktop] function.
//
// [PaintDesktop]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-paintdesktop
func (hdc HDC) PaintDesktop() error {
	ret, _, err := syscall.SyscallN(_PaintDesktop.Addr(),
		uintptr(hdc))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _PaintDesktop = dll.User32.NewProc("PaintDesktop")

// [WindowFromDC] function.
//
// [WindowFromDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-windowfromdc
func (hdc HDC) WindowFromDC() HWND {
	ret, _, _ := syscall.SyscallN(_WindowFromDC.Addr(),
		uintptr(hdc))
	return HWND(ret)
}

var _WindowFromDC = dll.User32.NewProc("WindowFromDC")
