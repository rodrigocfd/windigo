//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [hook].
//
// [hook]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hhook
type HHOOK HANDLE

// [SetWindowsHookEx] function.
//
// Note that the callback is recreated each function call, and the number of
// system callbacks is limited somewhere by the Go runtime.
//
// SetWindowsHookEx() doesn't have a context argument, so everything inside of
// it depends on global objects.
//
// ⚠️ You must defer HHOOK.UnhookWindowsHookEx().
//
// [SetWindowsHookEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowshookexw
func SetWindowsHookEx(idHook co.WH,
	callback func(code int32, wp WPARAM, lp LPARAM) uintptr,
	hMod HINSTANCE,
	threadId uint32) (HHOOK, error) {

	ret, _, err := syscall.SyscallN(proc.SetWindowsHookEx.Addr(),
		uintptr(idHook), syscall.NewCallback(callback),
		uintptr(hMod), uintptr(threadId))
	if ret == 0 {
		return HHOOK(0), errco.ERROR(err)
	}
	return HHOOK(ret), nil
}

// [CallNextHookEx] function.
//
// [CallNextHookEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-callnexthookex
func (hHook HHOOK) CallNextHookEx(nCode int32, wp WPARAM, lp LPARAM) uintptr {
	ret, _, _ := syscall.SyscallN(proc.CallNextHookEx.Addr(),
		uintptr(hHook), uintptr(nCode), uintptr(wp), uintptr(lp))
	return uintptr(ret)
}

// [UnhookWindowsHookEx] function.
//
// [UnhookWindowsHookEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unhookwindowshookex
func (hHook HHOOK) UnhookWindowsHookEx() error {
	ret, _, err := syscall.SyscallN(proc.UnhookWindowsHookEx.Addr(),
		uintptr(hHook))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}
