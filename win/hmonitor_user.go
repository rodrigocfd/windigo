//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Handle to a display monitor.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hmonitor
type HMONITOR HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfrompoint
func MonitorFromPoint(pt POINT, flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.SyscallN(proc.MonitorFromPoint.Addr(),
		uintptr(pt.X), uintptr(pt.Y), uintptr(flags))
	return HMONITOR(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfromrect
func MonitorFromRect(rc RECT, flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.SyscallN(proc.MonitorFromRect.Addr(),
		uintptr(unsafe.Pointer(&rc)), uintptr(flags))
	return HMONITOR(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmonitorinfow
func (hMon HMONITOR) GetMonitorInfo(mi *MONITORINFOEX) error {
	ret, _, err := syscall.SyscallN(proc.GetMonitorInfo.Addr(),
		uintptr(hMon), uintptr(unsafe.Pointer(mi)))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}
