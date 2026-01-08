//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a [hook].
//
// [hook]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hhook
type HHOOK HANDLE

// [SetWindowsHookEx] function.
//
// Note that [syscall.NewCallback] is called each time you call this function,
// which means only a limited number of calls can be made. See its documentation
// for further details.
//
// ⚠️ You must defer [HHHOK.UnhookWindowsHookEx].
//
// [SetWindowsHookEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowshookexw
func SetWindowsHookEx(
	idHook co.WH,
	hMod HINSTANCE,
	threadId int32,
	hookProc func(code int32, wp WPARAM, lp LPARAM) uintptr,
) (HHOOK, error) {
	hookCallback := syscall.NewCallback(
		func(code int32, wParam WPARAM, lParam LPARAM) uintptr {
			return hookProc(code, wParam, lParam)
		},
	)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_SetWindowsHookExW, "SetWindowsHookExW"),
		uintptr(idHook),
		hookCallback,
		uintptr(hMod),
		uintptr(threadId))
	if ret == 0 {
		return HHOOK(0), co.ERROR(err)
	}
	return HHOOK(ret), nil
}

var _user_SetWindowsHookExW *syscall.Proc

// [UnhookWindowsHookEx] function.
//
// [UnhookWindowsHookEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unhookwindowshookex
func (me HHOOK) UnhookWindowsHookEx() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_UnhookWindowsHookEx, "UnhookWindowsHookEx"))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_UnhookWindowsHookEx *syscall.Proc
