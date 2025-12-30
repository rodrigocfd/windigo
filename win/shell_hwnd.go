//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [ShellAbout] function.
//
// [ShellAbout]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellaboutw
func (hWnd HWND) ShellAbout(app, otherStuff string, hIcon HICON) error {
	var wApp, wOtherStuff wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_shell_ShellAboutW, "ShellAboutW"),
		uintptr(hWnd),
		uintptr(wApp.AllowEmpty(app)),
		uintptr(wOtherStuff.EmptyIsNil(otherStuff)),
		uintptr(hIcon))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _shell_ShellAboutW *syscall.Proc
