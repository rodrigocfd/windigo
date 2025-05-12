//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/wutil"
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
	ret, _, err := syscall.SyscallN(_BeginDeferWindowPos.Addr(),
		uintptr(numWindows))
	if ret == 0 {
		return HDWP(0), co.ERROR(err)
	}
	return HDWP(ret), nil
}

var _BeginDeferWindowPos = dll.User32.NewProc("BeginDeferWindowPos")

// [DeferWindowPos] function.
//
// [DeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deferwindowpos
func (hDwp HDWP) DeferWindowPos(
	hWnd, hwndInsertAfter HWND,
	x, y, cx, cy int,
	uFlags co.SWP,
) (HDWP, error) {
	ret, _, err := syscall.SyscallN(_DeferWindowPos.Addr(),
		uintptr(hDwp), uintptr(hWnd), uintptr(hwndInsertAfter),
		uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	if ret == 0 {
		return HDWP(0), co.ERROR(err)
	}
	return HDWP(ret), nil
}

var _DeferWindowPos = dll.User32.NewProc("DeferWindowPos")

// [EndDeferWindowPos] function.
//
// Paired with [BeginDeferWindowPos].
//
// [EndDeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddeferwindowpos
// [BeginDeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-begindeferwindowpos
func (hDwp HDWP) EndDeferWindowPos() error {
	ret, _, err := syscall.SyscallN(_EndDeferWindowPos.Addr(),
		uintptr(hDwp))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _EndDeferWindowPos = dll.User32.NewProc("EndDeferWindowPos")
