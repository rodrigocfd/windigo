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
	name16 := wstr.NewBufWith[wstr.Stack20](name, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_AddFontResourceExW),
		uintptr(name16.UnsafePtr()),
		uintptr(fl),
		0)
	if ret == 0 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return uint(ret), nil
}

// [GdiFlush] function.
//
// [GdiFlush]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gdiflush
func GdiFlush() bool {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GdiFlush))
	return ret == 0
}
