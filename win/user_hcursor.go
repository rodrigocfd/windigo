//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a [cursor].
//
// [cursor]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hcursor
type HCURSOR HANDLE

// [CreateIconFromResourceEx] function for cursor.
//
// This function creates [HCURSOR] only. The [HICON] variation is
// [CreateIconFromResourceEx].
//
// ⚠️ You must defer [HCURSOR.DestroyCursor].
//
// [CreateIconFromResourceEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconfromresourceex
func CreateCursorFromResourceEx(
	resBits []byte,
	fmtVersion uint32,
	cxDesired, cyDesired int,
	flags co.LR,
) (HCURSOR, error) {
	hIcon, err := CreateIconFromResourceEx(resBits, fmtVersion, cxDesired, cyDesired, flags)
	return HCURSOR(hIcon), err
}

// [GetCursor] function.
//
// [GetCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursor
func GetCursor() HCURSOR {
	ret, _, _ := syscall.SyscallN(
		dll.User.Load(&_user_GetCursor, "GetCursor"))
	return HCURSOR(ret)
}

var _user_GetCursor *syscall.Proc

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
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_DestroyCursor, "DestroyCursor"),
		uintptr(hCursor))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_DestroyCursor *syscall.Proc

// [SetCursor] function.
//
// [SetCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setcursor
func (hCursor HCURSOR) SetCursor() (HCURSOR, error) {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_SetCursor, "SetCursor"),
		uintptr(hCursor))
	if wErr := co.ERROR(err); wErr != co.ERROR_SUCCESS {
		return HCURSOR(0), err
	} else {
		return HCURSOR(ret), nil
	}
}

var _user_SetCursor *syscall.Proc
