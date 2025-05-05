//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
)

// [GdiFlush] function.
//
// [GdiFlush]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gdiflush
func GdiFlush() bool {
	ret, _, _ := syscall.SyscallN(_GdiFlush.Addr())
	return ret == 0
}

var _GdiFlush = dll.Gdi32.NewProc("GdiFlush")
