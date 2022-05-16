//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ole2/nf-ole2-registerdragdrop
func RegisterDragDrop(hWnd win.HWND, dropTarget *shellvt.IDropTarget) {
	ret, _, _ := syscall.Syscall(proc.RegisterDragDrop.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(&dropTarget)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ole2/nf-ole2-revokedragdrop
func RevokeDragDrop(hWnd win.HWND) {
	ret, _, _ := syscall.Syscall(proc.RevokeDragDrop.Addr(), 1,
		uintptr(hWnd), 0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
