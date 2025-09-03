//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a [region].
//
// [region]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hrgn
type HRGN HGDIOBJ

// [CreateRectRgnIndirect] function.
//
// ⚠️ You must defer [HRGN.DeleteObject].
//
// [CreateRectRgnIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createrectrgnindirect
func CreateRectRgnIndirect(bounds RECT) (HRGN, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateRectRgnIndirect, "CreateRectRgnIndirect"),
		uintptr(unsafe.Pointer(&bounds)))
	if ret == 0 {
		return HRGN(0), co.ERROR_INVALID_PARAMETER
	}
	return HRGN(ret), nil
}

var _CreateRectRgnIndirect *syscall.Proc

// [CombineRgn] function.
//
// [CombineRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-combinergn
func (hRgn HRGN) CombineRgn(src1, src2 HRGN, mode co.RGN) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CombineRgn, "CombineRgn"),
		uintptr(hRgn),
		uintptr(src1),
		uintptr(src2),
		uintptr(mode))
	if ret == 0 {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _CombineRgn *syscall.Proc

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hRgn HRGN) DeleteObject() error {
	return HGDIOBJ(hRgn).DeleteObject()
}

// [EqualRgn] function.
//
// [EqualRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-equalrgn
func (hRgn HRGN) EqualRgn(other HRGN) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_EqualRgn, "EqualRgn"),
		uintptr(hRgn),
		uintptr(other))
	return ret != 0
}

var _EqualRgn *syscall.Proc

// [GetRgnBox] function.
//
// [GetRgnBox]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getrgnbox
func (hRgn HRGN) GetRgnBox() (RECT, co.REGION, error) {
	var rc RECT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetRgnBox, "GetRgnBox"),
		uintptr(hRgn),
		uintptr(unsafe.Pointer(&rc)))
	if ret == utl.REGION_ERROR {
		return RECT{}, co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return rc, co.REGION(ret), nil
}

var _GetRgnBox *syscall.Proc

// [OffsetClipRgn] function.
//
// [OffsetClipRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-offsetcliprgn
func (hRgn HRGN) OffsetClipRgn(x, y int) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_OffsetClipRgn, "OffsetClipRgn"),
		uintptr(hRgn),
		uintptr(int32(x)),
		uintptr(int32(y)))
	if ret == 0 {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _OffsetClipRgn *syscall.Proc

// [OffsetRgn] function.
//
// [OffsetRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-offsetrgn
func (hRgn HRGN) OffsetRgn(x, y int) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_OffsetRgn, "OffsetRgn"),
		uintptr(hRgn),
		uintptr(int32(x)),
		uintptr(int32(y)))
	if ret == 0 {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _OffsetRgn *syscall.Proc

// [PtInRegion] function.
//
// [PtInRegion]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ptinregion
func (hRgn HRGN) PtInRegion(x, y int) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_PtInRegion, "PtInRegion"),
		uintptr(hRgn),
		uintptr(int32(x)),
		uintptr(int32(y)))
	return ret != 0
}

var _PtInRegion *syscall.Proc

// [RectInRegion] function.
//
// [RectInRegion]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rectinregion
func (hRgn HRGN) RectInRegion(rc RECT) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_RectInRegion, "RectInRegion"),
		uintptr(hRgn),
		uintptr(unsafe.Pointer(&rc)))
	return ret != 0
}

var _RectInRegion *syscall.Proc
