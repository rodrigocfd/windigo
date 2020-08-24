/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type HMONITOR HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfrompoint
func MonitorFromPoint(pt POINT, dwFlags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.Syscall(proc.MonitorFromPoint.Addr(), 3,
		uintptr(pt.X), uintptr(pt.Y), uintptr(dwFlags))
	return HMONITOR(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/shellscalingapi/nf-shellscalingapi-getdpiformonitor
//
// Available in Windows 8.1.
func (hMon HMONITOR) GetDpiForMonitor(dpiType co.MDT) (uint32, uint32) {
	dpiX, dpiY := uint32(0), uint32(0)
	ret, _, _ := syscall.Syscall6(proc.GetDpiForMonitor.Addr(), 4,
		uintptr(hMon), uintptr(dpiType),
		uintptr(unsafe.Pointer(&dpiX)), uintptr(unsafe.Pointer(&dpiY)),
		0, 0)
	if ret != 0 {
		panic("GetDpiForMonitor failed.")
	}
	return dpiX, dpiY
}
