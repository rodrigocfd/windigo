package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragacceptfiles
func (hWnd HWND) DragAcceptFiles(accept bool) {
	syscall.Syscall(proc.DragAcceptFiles.Addr(), 2,
		uintptr(hWnd), util.BoolToUintptr(accept), 0)
}
