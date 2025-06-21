//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to a deferred window position [structure].
//
// [structure]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdwp
type HDWP HANDLE

// [BeginDeferWindowPos] function.
//
// ⚠️ You must defer [HDWP.EndDeferWindowPos].
//
// [BeginDeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-begindeferwindowpos
func BeginDeferWindowPos(numWindows uint) (HDWP, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_BeginDeferWindowPos, "BeginDeferWindowPos"),
		uintptr(numWindows))
	if ret == 0 {
		return HDWP(0), co.ERROR(err)
	}
	return HDWP(ret), nil
}

var _BeginDeferWindowPos *syscall.Proc

// [DeferWindowPos] function.
//
// [DeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deferwindowpos
func (hDwp HDWP) DeferWindowPos(
	hWnd, hwndInsertAfter HWND,
	x, y, cx, cy int,
	uFlags co.SWP,
) (HDWP, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_DeferWindowPos, "DeferWindowPos"),
		uintptr(hDwp),
		uintptr(hWnd),
		uintptr(hwndInsertAfter),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(int32(cx)),
		uintptr(int32(cy)),
		uintptr(uFlags))
	if ret == 0 {
		return HDWP(0), co.ERROR(err)
	}
	return HDWP(ret), nil
}

var _DeferWindowPos *syscall.Proc

// [EndDeferWindowPos] function.
//
// Paired with [BeginDeferWindowPos].
//
// [EndDeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddeferwindowpos
// [BeginDeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-begindeferwindowpos
func (hDwp HDWP) EndDeferWindowPos() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_EndDeferWindowPos, "EndDeferWindowPos"),
		uintptr(hDwp))
	return utl.ZeroAsGetLastError(ret, err)
}

var _EndDeferWindowPos *syscall.Proc
