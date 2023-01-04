//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a GDI object.
//
// This type is used as the base type for the specialized GDI objects, being
// rarely used as itself.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hgdiobj
type HGDIOBJ HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hGdiObj HGDIOBJ) DeleteObject() {
	ret, _, err := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hGdiObj), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
