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
	ret, _, _ := syscall.Syscall(proc.MonitorFromPoint.Addr(), 3,
		uintptr(pt.X), uintptr(pt.Y), uintptr(flags))
	return HMONITOR(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfromrect
func MonitorFromRect(rc RECT, flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.Syscall(proc.MonitorFromRect.Addr(), 2,
		uintptr(unsafe.Pointer(&rc)), uintptr(flags), 0)
	return HMONITOR(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmonitorinfow
func (hMon HMONITOR) GetMonitorInfo(mi *MONITORINFO) error {
	ret, _, err := syscall.Syscall(proc.GetMonitorInfo.Addr(), 2,
		uintptr(hMon), uintptr(unsafe.Pointer(mi)), 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}
