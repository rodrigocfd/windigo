//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a [display monitor].
//
// [display monitor]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hmonitor
type HMONITOR HANDLE

// [MonitorFromPoint] function.
//
// [MonitorFromPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfrompoint
func MonitorFromPoint(pt POINT, flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_MonitorFromPoint, "MonitorFromPoint"),
		uintptr(utl.Make64(uint32(pt.X), uint32(pt.Y))),
		uintptr(flags))
	return HMONITOR(ret)
}

var _MonitorFromPoint *syscall.Proc

// [MonitorFromRect] function.
//
// [MonitorFromRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfromrect
func MonitorFromRect(rc *RECT, flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_MonitorFromRect, "MonitorFromRect"),
		uintptr(unsafe.Pointer(rc)),
		uintptr(flags))
	return HMONITOR(ret)
}

var _MonitorFromRect *syscall.Proc

// [GetMonitorInfo] function.
//
// [GetMonitorInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmonitorinfow
func (hMon HMONITOR) GetMonitorInfo() (MONITORINFOEX, error) {
	var mix MONITORINFOEX
	mix.SetCbSize()

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetMonitorInfoW, "GetMonitorInfoW"),
		uintptr(hMon),
		uintptr(unsafe.Pointer(&mix)))
	if ret == 0 {
		return MONITORINFOEX{}, co.ERROR_INVALID_PARAMETER
	}
	return mix, nil
}

var _GetMonitorInfoW *syscall.Proc
