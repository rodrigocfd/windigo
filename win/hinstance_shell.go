//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-duplicateicon
func (hInst HINSTANCE) DuplicateIcon(hIcon HICON) HICON {
	ret, _, err := syscall.Syscall(proc.DuplicateIcon.Addr(), 2,
		uintptr(hInst), uintptr(hIcon), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HICON(ret)
}
