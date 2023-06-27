//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [DuplicateIcon] function.
//
// [DuplicateIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-duplicateicon
func (hInst HINSTANCE) DuplicateIcon(hIcon HICON) HICON {
	ret, _, err := syscall.SyscallN(proc.DuplicateIcon.Addr(),
		uintptr(hInst), uintptr(hIcon))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HICON(ret)
}
