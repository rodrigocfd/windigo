//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
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
		dll.Load(dll.GDI32, &_gdi_GetStockObject, "GetStockObject"),
		uintptr(ty))
	if ret == 0 {
		return HGDIOBJ(0), co.ERROR_INVALID_PARAMETER
	}
	return HGDIOBJ(ret), nil
}

var _gdi_GetStockObject *syscall.Proc

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hGdiObj HGDIOBJ) DeleteObject() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_DeleteObject, "DeleteObject"),
		uintptr(hGdiObj))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _gdi_DeleteObject *syscall.Proc

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hGdiObj HGDIOBJ) GetObject(szBuf uintptr, buf unsafe.Pointer) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_GetObjectW, "GetObjectW"),
		uintptr(hGdiObj),
		uintptr(int32(szBuf)),
		uintptr(buf))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _gdi_GetObjectW *syscall.Proc

// [SelectObject] function.
//
// [SelectObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObject(hGdiObj HGDIOBJ) (HGDIOBJ, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_SelectObject, "SelectObject"),
		uintptr(hdc),
		uintptr(hGdiObj))
	if ret == 0 {
		return HGDIOBJ(0), co.ERROR_INVALID_PARAMETER
	}
	return HGDIOBJ(ret), nil
}

var _gdi_SelectObject *syscall.Proc

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// Handle to a [bitmap].
//
// [bitmap]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbitmap
type HBITMAP HGDIOBJ

// [CreateBitmap] function.
//
// Panics if numPlanes or bitCount is negative.
//
// ⚠️ You must defer [HBITMAP.DeleteObject].
//
// [CreateBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbitmap
func CreateBitmap(width, height, numPlanes, bitCount int, bits []byte) (HBITMAP, error) {
	utl.PanicNeg(numPlanes, bitCount)
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_CreateBitmap, "CreateBitmap"),
		uintptr(int32(width)),
		uintptr(int32(height)),
		uintptr(uint32(numPlanes)),
		uintptr(uint32(bitCount)),
		uintptr(unsafe.Pointer(unsafe.SliceData(bits))))
	if ret == 0 {
		return HBITMAP(0), co.ERROR_INVALID_PARAMETER
	}
	return HBITMAP(ret), nil
}

var _gdi_CreateBitmap *syscall.Proc

// [CreateBitmapIndirect] function.
//
// ⚠️ You must defer [HBITMAP.DeleteObject].
//
// [CreateBitmapIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbitmapindirect
func CreateBitmapIndirect(bm *BITMAP) (HBITMAP, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_CreateBitmapIndirect, "CreateBitmapIndirect"),
		uintptr(unsafe.Pointer(bm)))
	if ret == 0 {
		return HBITMAP(0), co.ERROR_INVALID_PARAMETER
	}
	return HBITMAP(ret), nil
}

var _gdi_CreateBitmapIndirect *syscall.Proc

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBmp HBITMAP) DeleteObject() error {
	return HGDIOBJ(hBmp).DeleteObject()
}

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hBmp HBITMAP) GetObject() (BITMAP, error) {
	var bmp BITMAP
	if err := HGDIOBJ(hBmp).GetObject(unsafe.Sizeof(bmp), unsafe.Pointer(&bmp)); err != nil {
		return BITMAP{}, err
	} else {
		return bmp, nil
	}
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// Handle to a [brush].
//
// [brush]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbrush
type HBRUSH HGDIOBJ

// [CreateBrushIndirect] function.
//
// ⚠️ You must defer [HBRUSH.DeleteObject].
//
// [CreateBrushIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbrushindirect
func CreateBrushIndirect(lb *LOGBRUSH) (HBRUSH, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_CreateBrushIndirect, "CreateBrushIndirect"),
		uintptr(unsafe.Pointer(lb)))
	if ret == 0 {
		return HBRUSH(0), co.ERROR_INVALID_PARAMETER
	}
	return HBRUSH(ret), nil
}

var _gdi_CreateBrushIndirect *syscall.Proc

// [CreatePatternBrush] function.
//
// ⚠️ You must defer [HBRUSH.DeleteObject].
//
// [CreatePatternBrush]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpatternbrush
func CreatePatternBrush(hbm HBITMAP) (HBRUSH, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_CreatePatternBrush, "CreatePatternBrush"),
		uintptr(hbm))
	if ret == 0 {
		return HBRUSH(0), co.ERROR_INVALID_PARAMETER
	}
	return HBRUSH(ret), nil
}

var _gdi_CreatePatternBrush *syscall.Proc

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBrush HBRUSH) DeleteObject() error {
	return HGDIOBJ(hBrush).DeleteObject()
}

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hBrush HBRUSH) GetObject() (LOGBRUSH, error) {
	var lb LOGBRUSH
	if err := HGDIOBJ(hBrush).GetObject(unsafe.Sizeof(lb), unsafe.Pointer(&lb)); err != nil {
		return LOGBRUSH{}, err
	} else {
		return lb, nil
	}
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// Handle to a [font].
//
// [font]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hfont
type HFONT HGDIOBJ

// [CreateFont] function.
//
// ⚠️ You must defer [HFONT.DeleteObject].
//
// [CreateFont]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createfontw
func CreateFont(
	height, width int,
	escapement int,
	orientation int,
	weight int,
	italic, underline, strikeOut bool,
	charSet co.CHARSET,
	outPrecision co.OUT_PRECIS,
	clipPrecision co.CLIP_PRECIS,
	quality co.QUALITY,
	pitch co.PITCH,
	family co.FF,
	faceName string,
) (HFONT, error) {
	var wFaceName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_CreateFontW, "CreateFontW"),
		uintptr(int32(height)),
		uintptr(int32(width)),
		uintptr(int32(escapement)),
		uintptr(int32(orientation)),
		uintptr(int32(height)),
		utl.BoolToUintptr(italic),
		utl.BoolToUintptr(underline),
		utl.BoolToUintptr(strikeOut),
		uintptr(charSet),
		uintptr(outPrecision),
		uintptr(clipPrecision),
		uintptr(quality),
		uintptr(uint8(pitch)|uint8(family)),
		uintptr(wFaceName.EmptyIsNil(faceName)))
	if ret == 0 {
		return HFONT(0), co.ERROR_INVALID_PARAMETER
	}
	return HFONT(ret), nil
}

var _gdi_CreateFontW *syscall.Proc

// [CreateFontIndirect] function.
//
// ⚠️ You must defer [HFONT.DeleteObject].
//
// [CreateFontIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createfontindirectw
func CreateFontIndirect(lf *LOGFONT) (HFONT, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_CreateFontIndirectW, "CreateFontIndirectW"),
		uintptr(unsafe.Pointer(lf)))
	if ret == 0 {
		return HFONT(0), co.ERROR_INVALID_PARAMETER
	}
	return HFONT(ret), nil
}

var _gdi_CreateFontIndirectW *syscall.Proc

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hFont HFONT) DeleteObject() error {
	return HGDIOBJ(hFont).DeleteObject()
}

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hFont HFONT) GetObject() (LOGFONT, error) {
	var lf LOGFONT
	if err := HGDIOBJ(hFont).GetObject(unsafe.Sizeof(lf), unsafe.Pointer(&lf)); err != nil {
		return LOGFONT{}, err
	} else {
		return lf, nil
	}
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// Handle to a [pen].
//
// [pen]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hpen
type HPEN HGDIOBJ

// [CreatePen] function.
//
// ⚠️ You must defer [HPEN.DeleteObject].
//
// [CreatePen]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpen
func CreatePen(style co.PS, width int, color COLORREF) (HPEN, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_CreatePen, "CreatePen"),
		uintptr(style),
		uintptr(int32(width)),
		uintptr(color))
	if ret == 0 {
		return HPEN(0), co.ERROR_INVALID_PARAMETER
	}
	return HPEN(ret), nil
}

var _gdi_CreatePen *syscall.Proc

// [CreatePenIndirect] function.
//
// ⚠️ You must defer [HPEN.DeleteObject].
//
// [CreatePenIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpenindirect
func CreatePenIndirect(lp *LOGPEN) (HPEN, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_CreatePenIndirect, "CreatePenIndirect"),
		uintptr(unsafe.Pointer(lp)))
	if ret == 0 {
		return HPEN(0), co.ERROR_INVALID_PARAMETER
	}
	return HPEN(ret), nil
}

var _gdi_CreatePenIndirect *syscall.Proc

// [ExtCreatePen] function.
//
// ⚠️ You must defer [HPEN.DeleteObject].
//
// [ExtCreatePen]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-extcreatepen
func ExtCreatePen(
	penType co.PS_TYPE,
	penStyle co.PS_STYLE,
	endCap co.PS_ENDCAP,
	width int,
	brush *LOGBRUSH,
	styleLengths []int,
) (HPEN, error) {
	var pLens unsafe.Pointer
	if styleLengths != nil {
		lens32 := make([]uint32, 0, len(styleLengths))
		for _, sl := range styleLengths {
			lens32 = append(lens32, uint32(sl))
		}
		pLens = unsafe.Pointer(unsafe.SliceData(lens32))
	}

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_ExtCreatePen, "ExtCreatePen"),
		uintptr(uint32(penType)|uint32(penStyle)|uint32(endCap)),
		uintptr(uint32(width)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(uint32(len(styleLengths))),
		uintptr(pLens))
	if ret == 0 {
		return HPEN(0), co.ERROR_INVALID_PARAMETER
	}
	return HPEN(ret), nil
}

var _gdi_ExtCreatePen *syscall.Proc

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hPen HPEN) DeleteObject() error {
	return HGDIOBJ(hPen).DeleteObject()
}

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hPen HPEN) GetObject() (LOGPEN, error) {
	var lp LOGPEN
	if err := HGDIOBJ(hPen).GetObject(unsafe.Sizeof(lp), unsafe.Pointer(&lp)); err != nil {
		return LOGPEN{}, err
	} else {
		return lp, nil
	}
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

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
		dll.Load(dll.GDI32, &_gdi_CreateRectRgnIndirect, "CreateRectRgnIndirect"),
		uintptr(unsafe.Pointer(&bounds)))
	if ret == 0 {
		return HRGN(0), co.ERROR_INVALID_PARAMETER
	}
	return HRGN(ret), nil
}

var _gdi_CreateRectRgnIndirect *syscall.Proc

// [CombineRgn] function.
//
// [CombineRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-combinergn
func (hRgn HRGN) CombineRgn(src1, src2 HRGN, mode co.RGN) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_CombineRgn, "CombineRgn"),
		uintptr(hRgn),
		uintptr(src1),
		uintptr(src2),
		uintptr(mode))
	if ret == 0 {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _gdi_CombineRgn *syscall.Proc

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
		dll.Load(dll.GDI32, &_gdi_EqualRgn, "EqualRgn"),
		uintptr(hRgn),
		uintptr(other))
	return ret != 0
}

var _gdi_EqualRgn *syscall.Proc

// [GetRgnBox] function.
//
// [GetRgnBox]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getrgnbox
func (hRgn HRGN) GetRgnBox() (RECT, co.REGION, error) {
	var rc RECT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_GetRgnBox, "GetRgnBox"),
		uintptr(hRgn),
		uintptr(unsafe.Pointer(&rc)))
	if ret == utl.REGION_ERROR {
		return RECT{}, co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return rc, co.REGION(ret), nil
}

var _gdi_GetRgnBox *syscall.Proc

// [OffsetClipRgn] function.
//
// [OffsetClipRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-offsetcliprgn
func (hRgn HRGN) OffsetClipRgn(x, y int) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_OffsetClipRgn, "OffsetClipRgn"),
		uintptr(hRgn),
		uintptr(int32(x)),
		uintptr(int32(y)))
	if ret == 0 {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _gdi_OffsetClipRgn *syscall.Proc

// [OffsetRgn] function.
//
// [OffsetRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-offsetrgn
func (hRgn HRGN) OffsetRgn(x, y int) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_OffsetRgn, "OffsetRgn"),
		uintptr(hRgn),
		uintptr(int32(x)),
		uintptr(int32(y)))
	if ret == 0 {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _gdi_OffsetRgn *syscall.Proc

// [PtInRegion] function.
//
// [PtInRegion]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ptinregion
func (hRgn HRGN) PtInRegion(x, y int) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_PtInRegion, "PtInRegion"),
		uintptr(hRgn),
		uintptr(int32(x)),
		uintptr(int32(y)))
	return ret != 0
}

var _gdi_PtInRegion *syscall.Proc

// [RectInRegion] function.
//
// [RectInRegion]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rectinregion
func (hRgn HRGN) RectInRegion(rc RECT) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_gdi_RectInRegion, "RectInRegion"),
		uintptr(hRgn),
		uintptr(unsafe.Pointer(&rc)))
	return ret != 0
}

var _gdi_RectInRegion *syscall.Proc
