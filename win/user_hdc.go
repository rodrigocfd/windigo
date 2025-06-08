//go:build windows

package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// [DrawIcon] function.
//
// [DrawIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawicon
func (hdc HDC) DrawIcon(x, y int, hIcon HICON) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_DrawIcon),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(hIcon))
	return utl.ZeroAsGetLastError(ret, err)
}

// [DrawIconEx] function.
//
// [DrawIconEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawiconex
func (hdc HDC) DrawIconEx(
	pos POINT,
	hIcon HICON,
	size SIZE,
	frameIndexIfCursor uint,
	hbrFlickerFree HBRUSH,
	diFlags co.DI,
) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_DrawIconEx),
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

// [EnumDisplayMonitors] function.
//
// [EnumDisplayMonitors]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func (hdc HDC) EnumDisplayMonitors(rcClip *RECT) ([]EnumDisplayMonitorsInfo, error) {
	pPack := &_EnumDisplayMonitorsPack{
		arr: make([]EnumDisplayMonitorsInfo, 0),
	}
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_EnumDisplayMonitors),
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
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_FrameRect),
		uintptr(hdc),
		uintptr(unsafe.Pointer(rc)),
		uintptr(hBrush))
	return utl.ZeroAsGetLastError(ret, err)
}

// [InvertRect] function.
//
// [InvertRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invertrect
func (hdc HDC) InvertRect(rc *RECT) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_InvertRect),
		uintptr(hdc),
		uintptr(unsafe.Pointer(rc)))
	return utl.ZeroAsGetLastError(ret, err)
}

// [PaintDesktop] function.
//
// [PaintDesktop]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-paintdesktop
func (hdc HDC) PaintDesktop() error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_PaintDesktop),
		uintptr(hdc))
	return utl.ZeroAsGetLastError(ret, err)
}

// [WindowFromDC] function.
//
// [WindowFromDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-windowfromdc
func (hdc HDC) WindowFromDC() HWND {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_WindowFromDC),
		uintptr(hdc))
	return HWND(ret)
}
