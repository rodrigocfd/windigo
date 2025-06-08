//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to a [display monitor].
//
// [display monitor]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hmonitor
type HMONITOR HANDLE

// [MonitorFromPoint] function.
//
// [MonitorFromPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfrompoint
func MonitorFromPoint(pt POINT, flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_MonitorFromPoint),
		uintptr(utl.Make64(uint32(pt.X), uint32(pt.Y))),
		uintptr(flags))
	return HMONITOR(ret)
}

// [MonitorFromRect] function.
//
// [MonitorFromRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfromrect
func MonitorFromRect(rc *RECT, flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_MonitorFromRect),
		uintptr(unsafe.Pointer(rc)),
		uintptr(flags))
	return HMONITOR(ret)
}
