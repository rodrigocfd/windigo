//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/wstr"
)

// [AddFontResourceEx] function.
//
// [AddFontResourceEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-addfontresourceexw
func AddFontResourceEx(name string, fl co.FR) (int, error) {
	var wName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_AddFontResourceExW, "AddFontResourceExW"),
		uintptr(wName.AllowEmpty(name)),
		uintptr(fl),
		0)
	if ret == 0 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return int(int32(ret)), nil
}

var _gdi_AddFontResourceExW *syscall.Proc

// [GdiFlush] function.
//
// [GdiFlush]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gdiflush
func GdiFlush() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_GdiFlush, "GdiFlush"))
	return ret == 0
}

var _gdi_GdiFlush *syscall.Proc
