//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a deferred window position structure.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdwp
type HDWP HANDLE

// ⚠️ You must defer HDWP.EndDeferWindowPos().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-begindeferwindowpos
func BeginDeferWindowPos(numWindows int32) HDWP {
	ret, _, err := syscall.SyscallN(proc.BeginDeferWindowPos.Addr(),
		uintptr(numWindows))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HDWP(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deferwindowpos
func (hDwp HDWP) DeferWindowPos(
	hWnd, hwndInsertAfter HWND, x, y, cx, cy int32, uFlags co.SWP) HDWP {

	ret, _, err := syscall.SyscallN(proc.DeferWindowPos.Addr(),
		uintptr(hDwp), uintptr(hWnd), uintptr(hwndInsertAfter),
		uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HDWP(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddeferwindowpos
func (hDwp HDWP) EndDeferWindowPos() error {
	ret, _, err := syscall.SyscallN(proc.EndDeferWindowPos.Addr(),
		uintptr(hDwp))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}
