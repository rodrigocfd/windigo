//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
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
	ret, _, _ := syscall.SyscallN(_CreateRectRgnIndirect.Addr(),
		uintptr(unsafe.Pointer(&bounds)))
	if ret == 0 {
		return HRGN(0), co.ERROR_INVALID_PARAMETER
	}
	return HRGN(ret), nil
}

var _CreateRectRgnIndirect = dll.Gdi32.NewProc("CreateRectRgnIndirect")

// [CombineRgn] function.
//
// [CombineRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-combinergn
func (hRgn HRGN) CombineRgn(src1, src2 HRGN, mode co.RGN) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(_CombineRgn.Addr(),
		uintptr(hRgn),
		uintptr(src1),
		uintptr(src2),
		uintptr(mode))
	if ret == 0 {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _CombineRgn = dll.Gdi32.NewProc("CombineRgn")

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
	ret, _, _ := syscall.SyscallN(_EqualRgn.Addr(),
		uintptr(hRgn),
		uintptr(other))
	return ret != 0
}

var _EqualRgn = dll.Gdi32.NewProc("EqualRgn")

// [GetRgnBox] function.
//
// [GetRgnBox]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getrgnbox
func (hRgn HRGN) GetRgnBox() (RECT, co.REGION, error) {
	var rc RECT
	ret, _, _ := syscall.SyscallN(_GetRgnBox.Addr(),
		uintptr(hRgn),
		uintptr(unsafe.Pointer(&rc)))
	if ret == utl.REGION_ERROR {
		return RECT{}, co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return rc, co.REGION(ret), nil
}

var _GetRgnBox = dll.Gdi32.NewProc("GetRgnBox")

// [OffsetClipRgn] function.
//
// [OffsetClipRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-offsetcliprgn
func (hRgn HRGN) OffsetClipRgn(x, y int32) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(_OffsetClipRgn.Addr(),
		uintptr(hRgn),
		uintptr(int32(x)),
		uintptr(int32(y)))
	if ret == 0 {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _OffsetClipRgn = dll.Gdi32.NewProc("OffsetClipRgn")

// [OffsetRgn] function.
//
// [OffsetRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-offsetrgn
func (hRgn HRGN) OffsetRgn(x, y int32) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(_OffsetRgn.Addr(),
		uintptr(hRgn),
		uintptr(int32(x)),
		uintptr(int32(y)))
	if ret == 0 {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _OffsetRgn = dll.Gdi32.NewProc("OffsetRgn")

// [PtInRegion] function.
//
// [PtInRegion]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ptinregion
func (hRgn HRGN) PtInRegion(x, y int32) bool {
	ret, _, _ := syscall.SyscallN(_PtInRegion.Addr(),
		uintptr(hRgn),
		uintptr(int32(x)),
		uintptr(int32(y)))
	return ret != 0
}

var _PtInRegion = dll.Gdi32.NewProc("PtInRegion")

// [RectInRegion] function.
//
// [RectInRegion]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rectinregion
func (hRgn HRGN) RectInRegion(rc RECT) bool {
	ret, _, _ := syscall.SyscallN(_RectInRegion.Addr(),
		uintptr(hRgn),
		uintptr(unsafe.Pointer(&rc)))
	return ret != 0
}

var _RectInRegion = dll.Gdi32.NewProc("RectInRegion")
