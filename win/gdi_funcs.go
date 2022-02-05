package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gdiflush
func GdiFlush() bool {
	ret, _, _ := syscall.Syscall(proc.GdiFlush.Addr(), 0,
		0, 0, 0)
	return ret == 0
}
