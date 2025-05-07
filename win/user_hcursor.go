//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to a [cursor].
//
// [cursor]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hcursor
type HCURSOR HANDLE

// [CopyCursor] function.
//
// [CopyCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copycursor
func (hCursor HCURSOR) CopyCursor() (HCURSOR, error) {
	hIcon, err := HICON(hCursor).CopyIcon()
	return HCURSOR(hIcon), err
}

// [DestroyCursor] function.
//
// [DestroyCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycursor
func (hCursor HCURSOR) DestroyCursor() error {
	ret, _, err := syscall.SyscallN(_DestroyCursor.Addr(),
		uintptr(hCursor))
	return util.ZeroAsGetLastError(ret, err)
}

var _DestroyCursor = dll.User32.NewProc("DestroyCursor")

// [SetCursor] function.
//
// [SetCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setcursor
func (hCursor HCURSOR) SetCursor() (HCURSOR, error) {
	ret, _, err := syscall.SyscallN(_SetCursor.Addr(),
		uintptr(hCursor))
	if wErr := co.ERROR(err); wErr != co.ERROR_SUCCESS {
		return HCURSOR(0), err
	} else {
		return HCURSOR(ret), nil
	}
}

var _SetCursor = dll.User32.NewProc("SetCursor")
