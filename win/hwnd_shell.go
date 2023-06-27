//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
)

// [DragAcceptFiles] function.
//
// [DragAcceptFiles]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragacceptfiles
func (hWnd HWND) DragAcceptFiles(accept bool) {
	syscall.SyscallN(proc.DragAcceptFiles.Addr(),
		uintptr(hWnd), util.BoolToUintptr(accept))
}
