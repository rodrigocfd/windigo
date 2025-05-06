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
	ret, _, _ := syscall.SyscallN(_AddFontResourceExW.Addr(),
		uintptr(name16.UnsafePtr()), uintptr(fl), 0)
	if ret == 0 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return uint(ret), nil
}

var _AddFontResourceExW = dll.Gdi32.NewProc("AddFontResourceExW")

// [GdiFlush] function.
//
// [GdiFlush]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gdiflush
func GdiFlush() bool {
	ret, _, _ := syscall.SyscallN(_GdiFlush.Addr())
	return ret == 0
}

var _GdiFlush = dll.Gdi32.NewProc("GdiFlush")
