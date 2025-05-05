//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to a [GDI object].
//
// This type is used as the base type for the specialized GDI objects, being
// rarely used as itself.
//
// [GDI object]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hgdiobj
type HGDIOBJ HANDLE

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hGdiObj HGDIOBJ) DeleteObject() error {
	ret, _, _ := syscall.SyscallN(_DeleteObject.Addr(),
		uintptr(hGdiObj))
	return util.ZeroAsSysInvalidParm(ret)
}

var _DeleteObject = dll.Gdi32.NewProc("DeleteObject")

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hGdiObj HGDIOBJ) GetObject(szBuf uintptr, buf unsafe.Pointer) error {
	ret, _, _ := syscall.SyscallN(_GetObjectW.Addr(),
		uintptr(hGdiObj), szBuf, uintptr(buf))
	return util.ZeroAsSysInvalidParm(ret)
}

var _GetObjectW = dll.Gdi32.NewProc("GetObjectW")

// [SelectObject] function.
//
// [SelectObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObject(hGdiObj HGDIOBJ) (HGDIOBJ, error) {
	ret, _, _ := syscall.SyscallN(_SelectObject.Addr(),
		uintptr(hdc), uintptr(hGdiObj))
	if ret == 0 {
		return HGDIOBJ(0), co.ERROR_INVALID_PARAMETER
	}
	return HGDIOBJ(ret), nil
}

var _SelectObject = dll.Gdi32.NewProc("SelectObject")
