//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [AddFontResourceEx] function.
//
// [AddFontResourceEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-addfontresourceexw
func AddFontResourceEx(name string, fl co.FR) (uint, error) {
	var wName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_AddFontResourceExW, "AddFontResourceExW"),
		uintptr(wName.AllowEmpty(name)),
		uintptr(fl),
		0)
	if ret == 0 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return uint(ret), nil
}

var _AddFontResourceExW *syscall.Proc

// [GdiFlush] function.
//
// [GdiFlush]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gdiflush
func GdiFlush() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GdiFlush, "GdiFlush"))
	return ret == 0
}

var _GdiFlush *syscall.Proc
