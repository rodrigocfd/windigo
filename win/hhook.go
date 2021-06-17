package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

// A handle to a hook.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hhook
type HHOOK HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowshookexw
func SetWindowsHookEx(idHook co.WH,
	lpfn func(code int32, wp WPARAM, lp LPARAM) uintptr,
	hmod HINSTANCE, dwThreadId uint32) HHOOK {

	ret, _, lerr := syscall.Syscall6(proc.SetWindowsHookEx.Addr(), 4,
		uintptr(idHook), syscall.NewCallback(lpfn),
		uintptr(hmod), uintptr(dwThreadId), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HHOOK(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-callnexthookex
func (hHook HHOOK) CallNextHookEx(nCode int32, wp WPARAM, lp LPARAM) uintptr {
	ret, _, _ := syscall.Syscall6(proc.CallNextHookEx.Addr(), 4,
		uintptr(hHook), uintptr(nCode), uintptr(wp), uintptr(lp), 0, 0)
	return uintptr(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unhookwindowshookex
func (hHook HHOOK) UnhookWindowsHookEx() {
	ret, _, lerr := syscall.Syscall(proc.UnhookWindowsHookEx.Addr(), 1,
		uintptr(hHook), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}
