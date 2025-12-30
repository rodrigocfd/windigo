//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a deferred window position [structure].
//
// [structure]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdwp
type HDWP HANDLE

// [BeginDeferWindowPos] function.
//
// Panics if numWindows is negative.
//
// ⚠️ You must defer [HDWP.EndDeferWindowPos].
//
// [BeginDeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-begindeferwindowpos
func BeginDeferWindowPos(numWindows int) (HDWP, error) {
	utl.PanicNeg(numWindows)
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_BeginDeferWindowPos, "BeginDeferWindowPos"),
		uintptr(int32(numWindows)))
	if ret == 0 {
		return HDWP(0), co.ERROR(err)
	}
	return HDWP(ret), nil
}

var _user_BeginDeferWindowPos *syscall.Proc

// [DeferWindowPos] function.
//
// [DeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deferwindowpos
func (hDwp HDWP) DeferWindowPos(
	hWnd, hwndInsertAfter HWND,
	x, y, cx, cy int,
	uFlags co.SWP,
) (HDWP, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_DeferWindowPos, "DeferWindowPos"),
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

var _user_DeferWindowPos *syscall.Proc

// [EndDeferWindowPos] function.
//
// Paired with [BeginDeferWindowPos].
//
// [EndDeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddeferwindowpos
// [BeginDeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-begindeferwindowpos
func (hDwp HDWP) EndDeferWindowPos() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_EndDeferWindowPos, "EndDeferWindowPos"),
		uintptr(hDwp))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_EndDeferWindowPos *syscall.Proc
