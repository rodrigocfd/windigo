//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
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
		dll.User.Load(&_user_MonitorFromPoint, "MonitorFromPoint"),
		pt.serializeUint64(),
		uintptr(flags))
	return HMONITOR(ret)
}

var _user_MonitorFromPoint *syscall.Proc

// [MonitorFromRect] function.
//
// [MonitorFromRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfromrect
func MonitorFromRect(pRc *RECT, flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.SyscallN(
		dll.User.Load(&_user_MonitorFromRect, "MonitorFromRect"),
		uintptr(unsafe.Pointer(pRc)),
		uintptr(flags))
	return HMONITOR(ret)
}

var _user_MonitorFromRect *syscall.Proc

// [GetMonitorInfo] function.
//
// [GetMonitorInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmonitorinfow
func (hMon HMONITOR) GetMonitorInfo() (MONITORINFOEX, error) {
	var mix MONITORINFOEX
	mix.SetCbSize()

	ret, _, _ := syscall.SyscallN(
		dll.User.Load(&_user_GetMonitorInfoW, "GetMonitorInfoW"),
		uintptr(hMon),
		uintptr(unsafe.Pointer(&mix)))
	if ret == 0 {
		return MONITORINFOEX{}, co.ERROR_INVALID_PARAMETER
	}
	return mix, nil
}

var _user_GetMonitorInfoW *syscall.Proc
