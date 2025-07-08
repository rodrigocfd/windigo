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
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pApp := wbuf.PtrAllowEmpty(app)
	pOtherStuff := wbuf.PtrEmptyIsNil(otherStuff)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_ShellAboutW, "ShellAboutW"),
		uintptr(hWnd),
		uintptr(pApp),
		uintptr(pOtherStuff),
		uintptr(hIcon))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ShellAboutW *syscall.Proc
