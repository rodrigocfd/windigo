//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [ShellAbout] function.
//
// [ShellAbout]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellaboutw
func (hWnd HWND) ShellAbout(app, otherStuff string, hIcon HICON) error {
	app16 := wstr.NewBufWith[wstr.Stack20](app, wstr.ALLOW_EMPTY)
	other16 := wstr.NewBufWith[wstr.Stack20](otherStuff, wstr.EMPTY_IS_NIL)

	ret, _, _ := syscall.SyscallN(_ShellAboutW.Addr(),
		uintptr(hWnd), uintptr(app16.UnsafePtr()), uintptr(other16.UnsafePtr()),
		uintptr(hIcon))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ShellAboutW = dll.Shell32.NewProc("ShellAboutW")
