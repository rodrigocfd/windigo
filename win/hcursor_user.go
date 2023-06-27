//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [cursor].
//
// [cursor]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hcursor
type HCURSOR HANDLE

// [CopyCursor] function.
//
// ⚠️ You must defer HCURSOR.DestroyCursor().
//
// [CopyCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copycursor
func (hCursor HCURSOR) CopyCursor() HCURSOR {
	return (HCURSOR)(((HICON)(hCursor)).CopyIcon())
}

// [DestroyCursor] function.
//
// [DestroyCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycursor
func (hCursor HCURSOR) DestroyCursor() error {
	ret, _, err := syscall.SyscallN(proc.DestroyCursor.Addr(),
		uintptr(hCursor))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [SetSystemCursor] function.
//
// [SetSystemCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setsystemcursor
func (hCursor HCURSOR) SetSystemCursor(id co.OCR) {
	ret, _, err := syscall.SyscallN(proc.SetSystemCursor.Addr(),
		uintptr(hCursor), uintptr(id))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
