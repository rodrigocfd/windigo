package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a cursor.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hcursor
type HCURSOR HANDLE

// ⚠️ You must defer HCURSOR.DestroyCursor().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copycursor
func (hCursor HCURSOR) CopyCursor() HCURSOR {
	return (HCURSOR)(((HICON)(hCursor)).CopyIcon())
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycursor
func (hCursor HCURSOR) DestroyCursor() {
	ret, _, err := syscall.Syscall(proc.DestroyCursor.Addr(), 1,
		uintptr(hCursor), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setsystemcursor
func (hCursor HCURSOR) SetSystemCursor(id co.OCR) {
	ret, _, err := syscall.Syscall(proc.SetSystemCursor.Addr(), 2,
		uintptr(hCursor), uintptr(id), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
