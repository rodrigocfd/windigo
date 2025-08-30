//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a [GDI object].
//
// This type is used as the base type for the specialized GDI objects, being
// rarely used as itself.
//
// [GDI object]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hgdiobj
type HGDIOBJ HANDLE

// [GetStockObject] function.
//
// ⚠️ The returned HGDIOBJ must be cast into the proper GDI object.
//
// [GetStockObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getstockobject
func GetStockObject(ty co.STOCK) (HGDIOBJ, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetStockObject, "GetStockObject"),
		uintptr(ty))
	if ret == 0 {
		return HGDIOBJ(0), co.ERROR_INVALID_PARAMETER
	}
	return HGDIOBJ(ret), nil
}

var _GetStockObject *syscall.Proc

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hGdiObj HGDIOBJ) DeleteObject() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_DeleteObject, "DeleteObject"),
		uintptr(hGdiObj))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _DeleteObject *syscall.Proc

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hGdiObj HGDIOBJ) GetObject(szBuf uintptr, buf unsafe.Pointer) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetObjectW, "GetObjectW"),
		uintptr(hGdiObj),
		uintptr(int32(szBuf)),
		uintptr(buf))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _GetObjectW *syscall.Proc

// [SelectObject] function.
//
// [SelectObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObject(hGdiObj HGDIOBJ) (HGDIOBJ, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SelectObject, "SelectObject"),
		uintptr(hdc),
		uintptr(hGdiObj))
	if ret == 0 {
		return HGDIOBJ(0), co.ERROR_INVALID_PARAMETER
	}
	return HGDIOBJ(ret), nil
}

var _SelectObject *syscall.Proc
