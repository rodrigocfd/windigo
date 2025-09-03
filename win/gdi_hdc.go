//go:build windows

package win

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// Handle to a [device context] (DC).
//
// [device context]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdc
type HDC HANDLE

// [CreateDC] function.
//
// ⚠️ You must defer [HDC.DeleteDC].
//
// [CreateDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdcw
func CreateDC(driver, device string, dm *DEVMODE) (HDC, error) {
	var wDriver, wDevice wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateDCW, "CreateDCW"),
		uintptr(wDriver.EmptyIsNil(driver)),
		uintptr(wDevice.AllowEmpty(device)),
		0,
		uintptr(unsafe.Pointer(dm)))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _CreateDCW *syscall.Proc

// [CreateIC] function.
//
// ⚠️ You must defer [HDC.DeleteDC].
//
// [CreateIC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createicw
func CreateIC(driver, device string, dm *DEVMODE) (HDC, error) {
	var wDriver, wDevice wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateICW, "CreateICW"),
		uintptr(wDriver.AllowEmpty(driver)),
		uintptr(wDevice.AllowEmpty(device)),
		0,
		uintptr(unsafe.Pointer(dm)))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _CreateICW *syscall.Proc

// [AbortDoc] function.
//
// [AbortDoc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-abortdoc
func (hdc HDC) AbortDoc() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_AbortDoc, "AbortDoc"),
		uintptr(hdc))
	return utl.Minus1AsSysInvalidParm(ret)
}

var _AbortDoc *syscall.Proc

// [AbortPath] function.
//
// [AbortPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-abortpath
func (hdc HDC) AbortPath() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_AbortPath, "AbortPath"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _AbortPath *syscall.Proc

// [AlphaBlend] function.
//
// This method is called from the destination HDC.
//
// [AlphaBlend]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-alphablend
func (hdc HDC) AlphaBlend(
	originDest POINT,
	szDest SIZE,
	hdcSrc HDC,
	originSrc POINT,
	szSrc SIZE,
	ftn BLENDFUNCTION,
) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_AlphaBlend, "AlphaBlend"),
		uintptr(hdc),
		uintptr(originDest.X),
		uintptr(originDest.Y),
		uintptr(szDest.Cx),
		uintptr(szDest.Cy),
		uintptr(hdcSrc),
		uintptr(originSrc.X),
		uintptr(originSrc.Y),
		uintptr(szSrc.Cx),
		uintptr(szSrc.Cy),
		uintptr(
			utl.Make32(
				utl.Make16(ftn.BlendOp, ftn.BlendFlags),
				utl.Make16(ftn.SourceConstantAlpha, ftn.AlphaFormat),
			),
		))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _AlphaBlend *syscall.Proc

// [AngleArc] function.
//
// [AngleArc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-anglearc
func (hdc HDC) AngleArc(center POINT, r int, startAngle, sweepAngle float32) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_AngleArc, "AngleArc"),
		uintptr(hdc),
		uintptr(center.X),
		uintptr(center.Y),
		uintptr(uint32(r)),
		uintptr(startAngle),
		uintptr(sweepAngle))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _AngleArc *syscall.Proc

// [Arc] function.
//
// [Arc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arc
func (hdc HDC) Arc(bound RECT, radialStart, radialEnd POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_Arc, "Arc"),
		uintptr(hdc),
		uintptr(bound.Left),
		uintptr(bound.Top),
		uintptr(bound.Right),
		uintptr(bound.Bottom),
		uintptr(radialStart.X),
		uintptr(radialStart.Y),
		uintptr(radialEnd.X),
		uintptr(radialEnd.Y))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _Arc *syscall.Proc

// [ArcTo] function.
//
// [ArcTo]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arcto
func (hdc HDC) ArcTo(bound RECT, radialStart, radialEnd POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_ArcTo, "ArcTo"),
		uintptr(hdc),
		uintptr(bound.Left),
		uintptr(bound.Top),
		uintptr(bound.Right),
		uintptr(bound.Bottom),
		uintptr(radialStart.X),
		uintptr(radialStart.Y),
		uintptr(radialEnd.X),
		uintptr(radialEnd.Y))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ArcTo *syscall.Proc

// [BeginPath] function.
//
// ⚠️ You must defer [HDC.EndPath].
//
// [BeginPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-beginpath
func (hdc HDC) BeginPath() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_BeginPath, "BeginPath"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _BeginPath *syscall.Proc

// [BitBlt] function.
//
// This method is called from the destination HDC.
//
// [BitBlt]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-bitblt
func (hdc HDC) BitBlt(
	destTopLeft POINT,
	sz SIZE,
	hdcSrc HDC,
	srcTopLeft POINT,
	rop co.ROP,
) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.GDI32, &_BitBlt, "BitBlt"),
		uintptr(hdc),
		uintptr(destTopLeft.X),
		uintptr(destTopLeft.Y),
		uintptr(sz.Cx),
		uintptr(sz.Cy),
		uintptr(hdcSrc),
		uintptr(srcTopLeft.X),
		uintptr(srcTopLeft.Y),
		uintptr(rop))
	return utl.ZeroAsGetLastError(ret, err)
}

var _BitBlt *syscall.Proc

// [CancelDC] function.
//
// [CancelDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-canceldc
func (hdc HDC) CancelDC() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CancelDC, "CancelDC"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _CancelDC *syscall.Proc

// [ChoosePixelFormat] function.
//
// [ChoosePixelFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-choosepixelformat
func (hdc HDC) ChoosePixelFormat(pfd *PIXELFORMATDESCRIPTOR) (int, error) {
	pfd.SetNSize() // safety
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.GDI32, &_ChoosePixelFormat, "ChoosePixelFormat"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(pfd)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return int(int32(ret)), nil
}

var _ChoosePixelFormat *syscall.Proc

// [Chord] function.
//
// [Chord]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-chord
func (hdc HDC) Chord(bound RECT, radialStart, radialEnd POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_Chord, "Chord"),
		uintptr(hdc),
		uintptr(bound.Left),
		uintptr(bound.Top),
		uintptr(bound.Right),
		uintptr(bound.Bottom),
		uintptr(radialStart.X),
		uintptr(radialStart.Y),
		uintptr(radialEnd.X),
		uintptr(radialEnd.Y))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _Chord *syscall.Proc

// [CloseFigure] function.
//
// [CloseFigure]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-closefigure
func (hdc HDC) CloseFigure() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CloseFigure, "CloseFigure"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _CloseFigure *syscall.Proc

// [CreateCompatibleBitmap] function.
//
// ⚠️ You must defer [HBITMAP.DeleteObject].
//
// [CreateCompatibleBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatiblebitmap
func (hdc HDC) CreateCompatibleBitmap(cx, cy int) (HBITMAP, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateCompatibleBitmap, "CreateCompatibleBitmap"),
		uintptr(hdc),
		uintptr(int32(cx)),
		uintptr(int32(cy)))
	if ret == 0 {
		return HBITMAP(0), co.ERROR_INVALID_PARAMETER
	}
	return HBITMAP(ret), nil
}

var _CreateCompatibleBitmap *syscall.Proc

// [CreateCompatibleDC] function.
//
// ⚠️ You must defer [HDC.DeleteDC].
//
// [CreateCompatibleDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatibledc
func (hdc HDC) CreateCompatibleDC() (HDC, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateCompatibleDC, "CreateCompatibleDC"),
		uintptr(hdc))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _CreateCompatibleDC *syscall.Proc

// [CreateDIBitmap] function.
//
// [CreateDIBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdibitmap
func (hdc HDC) CreateDIBitmap(
	pBmih *BITMAPV5HEADER,
	initialBitmapData []byte,
	pBmi *BITMAPINFO,
	usage co.DIB_COLORS,
) error {
	pBmih.SetSize() // safety
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateDIBitmap, "CreateDIBitmap"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(pBmih)),
		utl.CBM_INIT,
		uintptr(unsafe.Pointer(&initialBitmapData[0])),
		uintptr(unsafe.Pointer(pBmi)),
		uintptr(usage))
	return utl.ZeroAsSysError(ret)
}

var _CreateDIBitmap *syscall.Proc

// [CreateDIBSection] function.
//
// ⚠️ You must defer [HBITMAP.DeleteObject].
//
// [CreateDIBSection]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdibsection
func (hdc HDC) CreateDIBSection(
	bmi *BITMAPINFO,
	usage co.DIB_COLORS,
	hSection HFILEMAP,
	offset int,
) (HBITMAP, *byte, error) {
	var ppvBits *byte
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateDIBSection, "CreateDIBSection"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(bmi)),
		uintptr(usage),
		uintptr(unsafe.Pointer(&ppvBits)),
		uintptr(hSection),
		uintptr(uint32(offset)))
	if ret == 0 {
		return HBITMAP(0), nil, co.ERROR(err)
	}
	return HBITMAP(ret), ppvBits, nil
}

var _CreateDIBSection *syscall.Proc

// [CreateHalftonePalette] function.
//
// ⚠️ You must defer [HPALETTE.DeleteObject].
//
// [CreateHalftonePalette]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createhalftonepalette
func (hdc HDC) CreateHalftonePalette() (HPALETTE, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateHalftonePalette, "CreateHalftonePalette"),
		uintptr(hdc))
	if ret == 0 {
		return HPALETTE(0), co.ERROR_INVALID_PARAMETER
	}
	return HPALETTE(ret), nil
}

var _CreateHalftonePalette *syscall.Proc

// [DeleteDC] function.
//
// [DeleteDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deletedc
func (hdc HDC) DeleteDC() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_DeleteDC, "DeleteDC"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _DeleteDC *syscall.Proc

// [DescribePixelFormat] function.
//
// [DescribePixelFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-describepixelformat
func (hdc HDC) DescribePixelFormat(index int) (PIXELFORMATDESCRIPTOR, error) {
	var pfd PIXELFORMATDESCRIPTOR
	pfd.SetNSize()

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.GDI32, &_DescribePixelFormat, "DeleteDC"),
		uintptr(hdc),
		uintptr(int32(index)),
		uintptr(uint32(unsafe.Sizeof(pfd))),
		uintptr(unsafe.Pointer(&pfd)))
	if ret == 0 {
		return PIXELFORMATDESCRIPTOR{}, co.ERROR(err)
	}
	return pfd, nil
}

var _DescribePixelFormat *syscall.Proc

// [Ellipse] function.
//
// [Ellipse]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ellipse
func (hdc HDC) Ellipse(bound RECT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_Ellipse, "Ellipse"),
		uintptr(hdc),
		uintptr(bound.Left),
		uintptr(bound.Top),
		uintptr(bound.Right),
		uintptr(bound.Bottom))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _Ellipse *syscall.Proc

// [EndDoc] function.
//
// [EndDoc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-enddoc
func (hdc HDC) EndDoc() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_EndDoc, "EndDoc"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _EndDoc *syscall.Proc

// [EndPage] function.
//
// [EndPage]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-endpage
func (hdc HDC) EndPage() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_EndPage, "EndPage"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _EndPage *syscall.Proc

// [EndPath] function.
//
// Paired with [HDC.BeginPath].
//
// [EndPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-endpath
func (hdc HDC) EndPath() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_EndPath, "EndPath"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _EndPath *syscall.Proc

// [ExcludeClipRect] function.
//
// [ExcludeClipRect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-excludecliprect
func (hdc HDC) ExcludeClipRect(rc RECT) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_ExcludeClipRect, "ExcludeClipRect"),
		uintptr(hdc),
		uintptr(rc.Left),
		uintptr(rc.Top),
		uintptr(rc.Right),
		uintptr(rc.Bottom))
	if ret == 0 {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _ExcludeClipRect *syscall.Proc

// [FillPath] function.
//
// [FillPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillpath
func (hdc HDC) FillPath() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_FillPath, "FillPath"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _FillPath *syscall.Proc

// [FillRect] function.
//
// [FillRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-fillrect
func (hdc HDC) FillRect(rc *RECT, hBrush HBRUSH) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_FillRect, "FillRect"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(rc)),
		uintptr(hBrush))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _FillRect *syscall.Proc

// [FillRgn] function.
//
// [FillRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillrgn
func (hdc HDC) FillRgn(hRgn HRGN, hBrush HBRUSH) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_FillRgn, "FillRgn"),
		uintptr(hdc),
		uintptr(hRgn),
		uintptr(hBrush))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _FillRgn *syscall.Proc

// [FlattenPath] function.
//
// [FlattenPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-flattenpath
func (hdc HDC) FlattenPath() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_FlattenPath, "FlattenPath"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _FlattenPath *syscall.Proc

// [FrameRgn] function.
//
// [FrameRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-framergn
func (hdc HDC) FrameRgn(hRgn HRGN, hBrush HBRUSH, width, height int) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_FrameRgn, "FrameRgn"),
		uintptr(hdc),
		uintptr(hRgn),
		uintptr(hBrush),
		uintptr(int32(width)),
		uintptr(int32(height)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _FrameRgn *syscall.Proc

// [GetBkColor] function.
//
// [GetBkColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getbkcolor
func (hdc HDC) GetBkColor() (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetBkColor, "GetBkColor"),
		uintptr(hdc))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

var _GetBkColor *syscall.Proc

// [GetBkMode] function.
//
// [GetBkMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getbkmode
func (hdc HDC) GetBkMode() (co.BKMODE, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetBkMode, "GetBkMode"),
		uintptr(hdc))
	if ret == 0 {
		return co.BKMODE(0), co.ERROR_INVALID_PARAMETER
	}
	return co.BKMODE(ret), nil
}

var _GetBkMode *syscall.Proc

// [GetCurrentPositionEx] function.
//
// [GetCurrentPositionEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getcurrentpositionex
func (hdc HDC) GetCurrentPositionEx() (POINT, error) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetCurrentPositionEx, "GetCurrentPositionEx"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR_INVALID_PARAMETER
	}
	return pt, nil
}

var _GetCurrentPositionEx *syscall.Proc

// [GetDCBrushColor] function.
//
// [GetDCBrushColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdcbrushcolor
func (hdc HDC) GetDCBrushColor() (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetDCBrushColor, "GetDCBrushColor"),
		uintptr(hdc))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

var _GetDCBrushColor *syscall.Proc

// [GetDCOrgEx] function.
//
// [GetDCOrgEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdcorgex
func (hdc HDC) GetDCOrgEx() (POINT, error) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetDCOrgEx, "GetDCOrgEx"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR_INVALID_PARAMETER
	}
	return pt, nil
}

var _GetDCOrgEx *syscall.Proc

// [GetDCPenColor] function.
//
// [GetDCPenColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdcpencolor
func (hdc HDC) GetDCPenColor() (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetDCPenColor, "GetDCPenColor"),
		uintptr(hdc))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

var _GetDCPenColor *syscall.Proc

// [GetDeviceCaps] function.
//
// [GetDeviceCaps]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdevicecaps
func (hdc HDC) GetDeviceCaps(index co.GDC) int32 {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetDeviceCaps, "GetDeviceCaps"),
		uintptr(hdc),
		uintptr(index))
	return int32(ret)
}

var _GetDeviceCaps *syscall.Proc

// [GetDIBColorTable] function.
//
// [GetDIBColorTable]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdibcolortable
func (hdc HDC) GetDIBColorTable(iStart int, buf []RGBQUAD) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetDIBColorTable, "GetDIBColorTable"),
		uintptr(hdc),
		uintptr(uint32(iStart)),
		uintptr(uint32(len(buf))),
		uintptr(unsafe.Pointer(&buf[0])))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _GetDIBColorTable *syscall.Proc

// [GetDIBits] function.
//
// Note that this method fails if bitmapDataBuffer is an ordinary Go slice; it
// must be allocated directly from the OS heap, for example with [GlobalAlloc].
//
// Example:
//
// Taking a screenshot and saving into a BMP file:
//
//	cxScreen := win.GetSystemMetrics(co.SM_CXSCREEN)
//	cyScreen := win.GetSystemMetrics(co.SM_CYSCREEN)
//
//	hdcScreen, _ := win.HWND(0).GetDC()
//	defer win.HWND(0).ReleaseDC(hdcScreen)
//
//	hBmp, _ := hdcScreen.CreateCompatibleBitmap(int(cxScreen), int(cyScreen))
//	defer hBmp.DeleteObject()
//
//	hdcMem, _ := hdcScreen.CreateCompatibleDC()
//	defer hdcMem.DeleteDC()
//
//	hBmpOld, _ := hdcMem.SelectObjectBmp(hBmp)
//	defer hdcMem.SelectObjectBmp(hBmpOld)
//
//	_ = hdcMem.BitBlt(
//		win.POINT{X: 0, Y: 0},
//		win.SIZE{Cx: cxScreen, Cy: cyScreen},
//		hdcScreen,
//		win.POINT{X: 0, Y: 0},
//		co.ROP_SRCCOPY,
//	)
//
//	bi := win.BITMAPINFO{
//		BmiHeader: win.BITMAPINFOHEADER{
//			Width:       cxScreen,
//			Height:      cyScreen,
//			Planes:      1,
//			BitCount:    32,
//			Compression: co.BI_RGB,
//		},
//	}
//	bi.BmiHeader.SetSize()
//
//	bmpObj, _ := hBmp.GetObject()
//	bmpSize := bmpObj.CalcBitmapSize(bi.BmiHeader.BitCount)
//
//	rawMem, _ := win.GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, bmpSize)
//	defer rawMem.GlobalFree()
//
//	bmpSlice, _ := rawMem.GlobalLockSlice()
//	defer rawMem.GlobalUnlock()
//
//	_, _ = hdcScreen.GetDIBits(
//		hBmp,
//		0,
//		int(cyScreen),
//		bmpSlice,
//		&bi,
//		co.DIB_COLORS_RGB,
//	)
//
//	var bfh win.BITMAPFILEHEADER
//	bfh.SetBfType()
//	bfh.SetBfOffBits(uint32(unsafe.Sizeof(bfh) + unsafe.Sizeof(bi.BmiHeader)))
//	bfh.SetBfSize(bfh.BfOffBits() + uint32(bmpSize))
//
//	fout, _ := win.FileOpen(
//		"C:\\Temp\\screenshot.bmp",
//		co.FOPEN_RW_OPEN_OR_CREATE,
//	)
//	defer fout.Close()
//
//	_, _ = fout.Write(bfh.Serialize())
//	_, _ = fout.Write(bi.BmiHeader.Serialize())
//	_, _ = fout.Write(bmpSlice)
//
// [GetDIBits]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdibits
func (hdc HDC) GetDIBits(
	hbm HBITMAP,
	firstScanLine, numScanLines int,
	bitmapDataBuffer []byte,
	bmi *BITMAPINFO,
	usage co.DIB_COLORS,
) (int, error) {
	var dataBufPtr *byte
	if bitmapDataBuffer != nil {
		dataBufPtr = &bitmapDataBuffer[0]
	}

	bmi.BmiHeader.SetSize() // safety

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetDIBits, "GetDIBits"),
		uintptr(hdc),
		uintptr(hbm),
		uintptr(uint32(firstScanLine)),
		uintptr(uint32(numScanLines)),
		uintptr(unsafe.Pointer(dataBufPtr)),
		uintptr(unsafe.Pointer(bmi)),
		uintptr(usage))

	if ret == 0 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return int(int32(ret)), nil
}

var _GetDIBits *syscall.Proc

// [GetPixel] function.
//
// [GetPixel]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getpixel
func (hdc HDC) GetPixel() (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetPixel, "GetPixel"),
		uintptr(hdc))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

var _GetPixel *syscall.Proc

// [GetPixelFormat] function.
//
// [GetPixelFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getpixelformat
func (hdc HDC) GetPixelFormat() (int, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetPixelFormat, "GetPixelFormat"),
		uintptr(hdc))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return int(int32(ret)), nil
}

var _GetPixelFormat *syscall.Proc

// [GetPolyFillMode] function.
//
// [GetPolyFillMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getpolyfillmode
func (hdc HDC) GetPolyFillMode() (co.POLYF, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetPolyFillMode, "GetPolyFillMode"),
		uintptr(hdc))
	if ret == 0 {
		return co.POLYF(0), co.ERROR_INVALID_PARAMETER
	}
	return co.POLYF(ret), nil
}

var _GetPolyFillMode *syscall.Proc

// [GetTextColor] function.
//
// [GetTextColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextcolor
func (hdc HDC) GetTextColor() (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetTextColor, "GetTextColor"),
		uintptr(hdc))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

var _GetTextColor *syscall.Proc

// [GetTextExtentPoint32] function.
//
// [GetTextExtentPoint32]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextextentpoint32w
func (hdc HDC) GetTextExtentPoint32(text string) (SIZE, error) {
	var wText wstr.BufEncoder
	var sz SIZE

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetTextExtentPoint32W, "GetTextExtentPoint32W"),
		uintptr(hdc),
		uintptr(wText.AllowEmpty(text)),
		uintptr(int32(wstr.CountUtf16Len(text))),
		uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		return SIZE{}, co.ERROR_INVALID_PARAMETER
	}
	return sz, nil
}

var _GetTextExtentPoint32W *syscall.Proc

// [GetTextFace] function.
//
// [GetTextFace]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextfacew
func (hdc HDC) GetTextFace() (string, error) {
	var buf [utl.LF_FACESIZE]uint16
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetTextFaceW, "GetTextFaceW"),
		uintptr(hdc),
		uintptr(int32(len(buf))),
		uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 {
		return "", co.ERROR_INVALID_PARAMETER
	}
	return wstr.DecodeSlice(buf[:]), nil
}

var _GetTextFaceW *syscall.Proc

// [GetTextMetrics] function.
//
// [GetTextMetrics]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextmetricsw
func (hdc HDC) GetTextMetrics() (TEXTMETRIC, error) {
	var tm TEXTMETRIC
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetTextMetricsW, "GetTextMetricsW"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&tm)))
	if ret == 0 {
		return TEXTMETRIC{}, co.ERROR_INVALID_PARAMETER
	}
	return tm, nil
}

var _GetTextMetricsW *syscall.Proc

// [GetViewportExtEx] function.
//
// [GetViewportExtEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getviewportextex
func (hdc HDC) GetViewportExtEx() (SIZE, error) {
	var sz SIZE
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetViewportExtEx, "GetViewportExtEx"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		return SIZE{}, co.ERROR_INVALID_PARAMETER
	}
	return sz, nil
}

var _GetViewportExtEx *syscall.Proc

// [GetViewportOrgEx] function.
//
// [GetViewportOrgEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getviewportorgex
func (hdc HDC) GetViewportOrgEx() (POINT, error) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetViewportOrgEx, "GetViewportOrgEx"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR_INVALID_PARAMETER
	}
	return pt, nil
}

var _GetViewportOrgEx *syscall.Proc

// [GetWindowExtEx] function.
//
// [GetWindowExtEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getwindowextex
func (hdc HDC) GetWindowExtEx() (SIZE, error) {
	var sz SIZE
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetWindowExtEx, "GetWindowExtEx"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		return SIZE{}, co.ERROR_INVALID_PARAMETER
	}
	return sz, nil
}

var _GetWindowExtEx *syscall.Proc

// [GetWindowOrgEx] function.
//
// [GetWindowOrgEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getwindoworgex
func (hdc HDC) GetWindowOrgEx() (POINT, error) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GetWindowOrgEx, "GetWindowOrgEx"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR_INVALID_PARAMETER
	}
	return pt, nil
}

var _GetWindowOrgEx *syscall.Proc

// [GradientFill] function.
//
// You must specify one mesh, either meshTriangle or meshRect. If you specify
// both or none, the function panics.
//
// [GradientFill]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gradientfill
func (hdc HDC) GradientFill(
	vertex []TRIVERTEX,
	meshTriangle []GRADIENT_TRIANGLE,
	meshRect []GRADIENT_RECT,
	mode co.GRADIENT_FILL,
) error {
	if (meshTriangle == nil && meshRect == nil) ||
		(meshTriangle != nil && meshRect != nil) {
		panic("You must specify one: meshTriangle or meshRect.")
	}

	var pMesh unsafe.Pointer
	var nMesh int
	if meshTriangle != nil {
		pMesh = unsafe.Pointer(&meshTriangle[0])
		nMesh = len(meshTriangle)
	} else {
		pMesh = unsafe.Pointer(&meshRect[0])
		nMesh = len(meshRect)
	}

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_GradientFill, "GradientFill"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&vertex[0])),
		uintptr(uint32(len(vertex))),
		uintptr(pMesh),
		uintptr(uint32(nMesh)),
		uintptr(mode))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _GradientFill *syscall.Proc

// [AtlHiMetricToPixel] function. Converts HIMETRIC units to pixels.
//
// [AtlHiMetricToPixel]: https://learn.microsoft.com/en-us/cpp/atl/reference/pixel-himetric-conversion-global-functions?view=msvc-170#atlhimetrictopixel
func (hdc HDC) HiMetricToPixel(himetricX, himetricY int) (pixelX, pixelY int) {
	// http://www.verycomputer.com/5_5f2f75dc2d090ee8_1.htm
	// https://forums.codeguru.com/showthread.php?109554-Unresizable-activeX-control
	pixelX = int(
		(int64(himetricX) * int64(hdc.GetDeviceCaps(co.GDC_LOGPIXELSX))) /
			int64(utl.HIMETRIC_PER_INCH),
	)
	pixelY = int(
		(int64(himetricY) * int64(hdc.GetDeviceCaps(co.GDC_LOGPIXELSY))) /
			int64(utl.HIMETRIC_PER_INCH),
	)
	return
}

// [IntersectClipRect] function.
//
// [IntersectClipRect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-intersectcliprect
func (hdc HDC) IntersectClipRect(coords RECT) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_IntersectClipRect, "IntersectClipRect"),
		uintptr(hdc),
		uintptr(coords.Left),
		uintptr(coords.Top),
		uintptr(coords.Right),
		uintptr((coords.Bottom)))
	if ret == 0 {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _IntersectClipRect *syscall.Proc

// [InvertRgn] function.
//
// [InvertRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-invertrgn
func (hdc HDC) InvertRgn(hRgn HRGN) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_InvertRgn, "InvertRgn"),
		uintptr(hdc),
		uintptr(hRgn))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _InvertRgn *syscall.Proc

// [LineTo] function.
//
// [LineTo]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lineto
func (hdc HDC) LineTo(x, y int) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_LineTo, "LineTo"),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _LineTo *syscall.Proc

// [LPtoDP] function.
//
// [LPtoDP]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lptodp
func (hdc HDC) LPtoDP(pts []POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_LPtoDP, "LPtoDP"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(int32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _LPtoDP *syscall.Proc

// [MaskBlt] function.
//
// [MaskBlt]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-maskblt
func (hdc HDC) MaskBlt(
	destTopLeft POINT,
	sz SIZE,
	hdcSrc HDC,
	srcTopLeft POINT,
	hbmMask HBITMAP,
	maskOffset POINT,
	rop co.ROP,
) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_MaskBlt, "MaskBlt"),
		uintptr(hdc),
		uintptr(destTopLeft.X),
		uintptr(destTopLeft.Y),
		uintptr(sz.Cx),
		uintptr(sz.Cy),
		uintptr(hdcSrc),
		uintptr(srcTopLeft.X),
		uintptr(srcTopLeft.Y),
		uintptr(hbmMask),
		uintptr(maskOffset.X),
		uintptr(maskOffset.Y),
		uintptr(rop))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _MaskBlt *syscall.Proc

// [MoveToEx] function.
//
// [MoveToEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-movetoex
func (hdc HDC) MoveToEx(x, y int) (POINT, error) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_MoveToEx, "MoveToEx"),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR_INVALID_PARAMETER
	}
	return pt, nil
}

var _MoveToEx *syscall.Proc

// [PaintRgn] function.
//
// [PaintRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-paintrgn
func (hdc HDC) PaintRgn(hRgn HRGN) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_PaintRgn, "PaintRgn"),
		uintptr(hdc),
		uintptr(hRgn))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _PaintRgn *syscall.Proc

// [PatBlt] function.
//
// [PatBlt]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-patblt
func (hdc HDC) PatBlt(topLeft POINT, sz SIZE, rop co.ROP) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_PatBlt, "PatBlt"),
		uintptr(hdc),
		uintptr(topLeft.X),
		uintptr(topLeft.Y),
		uintptr(sz.Cx),
		uintptr(sz.Cy),
		uintptr(rop))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _PatBlt *syscall.Proc

// [PathToRegion] function.
//
// ⚠️ You must defer [HRGN.DeleteObject] on the returned HRGN.
//
// [PathToRegion]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-pathtoregion
func (hdc HDC) PathToRegion() (HRGN, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_PathToRegion, "PathToRegion"),
		uintptr(hdc))
	if ret == 0 {
		return HRGN(0), co.ERROR_INVALID_PARAMETER
	}
	return HRGN(ret), nil
}

var _PathToRegion *syscall.Proc

// [Pie] function.
//
// [Pie]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-pie
func (hdc HDC) Pie(bound RECT, endPointRadial1, endPointRadial2 POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_Pie, "Pie"),
		uintptr(hdc),
		uintptr(bound.Left),
		uintptr(bound.Top),
		uintptr(bound.Right),
		uintptr(bound.Bottom),
		uintptr(endPointRadial1.X),
		uintptr(endPointRadial1.Y),
		uintptr(endPointRadial2.X),
		uintptr(endPointRadial2.Y))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _Pie *syscall.Proc

// [AtlPixelToHiMetric] function. Converts pixels to HIMETRIC units.
//
// [AtlPixelToHiMetric]: https://learn.microsoft.com/en-us/cpp/atl/reference/pixel-himetric-conversion-global-functions?view=msvc-170#atlpixeltohimetri
func (hdc HDC) PixelToHiMetric(pixelX, pixelY int) (himetricX, himetricY int) {
	himetricX = int(
		(int64(pixelX) * int64(utl.HIMETRIC_PER_INCH)) /
			int64(hdc.GetDeviceCaps(co.GDC_LOGPIXELSX)),
	)
	himetricY = int(
		(int64(pixelY) * int64(utl.HIMETRIC_PER_INCH)) /
			int64(hdc.GetDeviceCaps(co.GDC_LOGPIXELSY)),
	)
	return
}

// [PolyBezier] function.
//
// [PolyBezier]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polybezier
func (hdc HDC) PolyBezier(pts []POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_PolyBezier, "PolyBezier"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(uint32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _PolyBezier *syscall.Proc

// [PolyBezierTo] function.
//
// [PolyBezierTo]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polybezierto
func (hdc HDC) PolyBezierTo(pts []POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_PolyBezierTo, "PolyBezierTo"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(uint32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _PolyBezierTo *syscall.Proc

// [PolyDraw] function.
//
// Panics if pts and usage don't have the same length.
//
// [PolyDraw]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polydraw
func (hdc HDC) PolyDraw(pts []POINT, usage []co.PT) error {
	if len(pts) != len(usage) {
		panic(fmt.Sprintf("PolyDraw different slice sizes: %d, %d.",
			len(pts), len(usage)))
	}

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_PolyDraw, "PolyDraw"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(unsafe.Pointer(&usage[0])),
		uintptr(int32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _PolyDraw *syscall.Proc

// [Polygon] function.
//
// [Polygon]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polygon
func (hdc HDC) Polygon(pts []POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_Polygon, "Polygon"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(uint32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _Polygon *syscall.Proc

// [Polyline] function.
//
// [Polyline]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polyline
func (hdc HDC) Polyline(pts []POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_Polyline, "Polyline"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(int32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _Polyline *syscall.Proc

// [PolylineTo] function.
//
// [PolylineTo]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polylineto
func (hdc HDC) PolylineTo(pts []POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_PolylineTo, "PolylineTo"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(uint32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _PolylineTo *syscall.Proc

// [PolyPolygon] function.
//
// [PolyPolygon]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polypolygon
func (hdc HDC) PolyPolygon(polygons [][]POINT) error {
	numPolygons := len(polygons)
	totalPts := 0
	for _, polygon := range polygons {
		totalPts += len(polygon)
	}

	allPtsFlat := make([]POINT, 0, totalPts)
	ptsPerPolygon := make([]int32, 0, numPolygons)

	for _, polygon := range polygons {
		allPtsFlat = append(allPtsFlat, polygon...)
		ptsPerPolygon = append(ptsPerPolygon, int32(len(polygon)))
	}

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_PolyPolygon, "PolyPolygon"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&allPtsFlat[0])),
		uintptr(unsafe.Pointer(&ptsPerPolygon[0])),
		uintptr(int32(len(polygons))))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _PolyPolygon *syscall.Proc

// [PolyPolyline] function.
//
// [PolyPolyline]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polypolyline
func (hdc HDC) PolyPolyline(polyLines [][]POINT) error {
	numPolyLines := len(polyLines)
	totalPts := 0
	for _, polygon := range polyLines {
		totalPts += len(polygon)
	}

	allPtsFlat := make([]POINT, 0, totalPts)
	ptsPerPolyLine := make([]int32, 0, numPolyLines)

	for _, polyLine := range polyLines {
		allPtsFlat = append(allPtsFlat, polyLine...)
		ptsPerPolyLine = append(ptsPerPolyLine, int32(len(polyLine)))
	}

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_PolyPolyline, "PolyPolyline"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&allPtsFlat[0])),
		uintptr(unsafe.Pointer(&ptsPerPolyLine[0])),
		uintptr(uint32(len(polyLines))))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _PolyPolyline *syscall.Proc

// [PtVisible] function.
//
// [PtVisible]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ptvisible
func (hdc HDC) PtVisible(x, y int) (bool, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_PtVisible, "PtVisible"),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)))
	if int32(ret) == -1 {
		return false, co.ERROR_INVALID_PARAMETER
	}
	return ret != 0, nil
}

var _PtVisible *syscall.Proc

// [RealizePalette] function.
//
// [RealizePalette]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-realizepalette
func (hdc HDC) RealizePalette() (int, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_RealizePalette, "RealizePalette"),
		uintptr(hdc))
	if ret == utl.GDI_ERR {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return int(uint32(ret)), nil
}

var _RealizePalette *syscall.Proc

// [Rectangle] function.
//
// [Rectangle]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rectangle
func (hdc HDC) Rectangle(bound RECT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_Rectangle, "Rectangle"),
		uintptr(hdc),
		uintptr(bound.Left),
		uintptr(bound.Top),
		uintptr(bound.Right),
		uintptr(bound.Bottom))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _Rectangle *syscall.Proc

// [ResetDC] function.
//
// [ResetDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-resetdcw
func (hdc HDC) ResetDC(dm *DEVMODE) (HDC, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_ResetDCW, "ResetDCW"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(dm)))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _ResetDCW *syscall.Proc

// [RestoreDC] function.
//
// Paired with [HDC.SaveDC].
//
// [RestoreDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-restoredc
func (hdc HDC) RestoreDC(savedDC int32) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_RestoreDC, "RestoreDC"),
		uintptr(hdc),
		uintptr(savedDC))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _RestoreDC *syscall.Proc

// [RoundRect] function.
//
// [RoundRect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-roundrect
func (hdc HDC) RoundRect(bound RECT, sz SIZE) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_RoundRect, "RoundRect"),
		uintptr(hdc),
		uintptr(bound.Left),
		uintptr(bound.Top),
		uintptr(bound.Right),
		uintptr(bound.Bottom),
		uintptr(sz.Cx),
		uintptr(sz.Cy))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _RoundRect *syscall.Proc

// [SaveDC] function.
//
// Paired with [HDC.RestoreDC].
//
// [SaveDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-savedc
func (hdc HDC) SaveDC() (int32, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SaveDC, "SaveDC"),
		uintptr(hdc))
	if ret == 0 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return int32(ret), nil
}

var _SaveDC *syscall.Proc

// [SelectClipPath] function.
//
// [SelectClipPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectclippath
func (hdc HDC) SelectClipPath(mode co.RGN) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SelectClipPath, "SelectClipPath"),
		uintptr(hdc),
		uintptr(mode))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _SelectClipPath *syscall.Proc

// [SelectClipRgn] function.
//
// [SelectClipRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectcliprgn
func (hdc HDC) SelectClipRgn(hRgn HRGN) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SelectClipRgn, "SelectClipRgn"),
		uintptr(hdc),
		uintptr(hRgn))
	if ret == utl.REGION_ERROR {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

var _SelectClipRgn *syscall.Proc

// [SelectObject] function for [HBITMAP].
//
// [SelectObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBmp(hBmp HBITMAP) (HBITMAP, error) {
	hGdiObj, err := hdc.SelectObject(HGDIOBJ(hBmp))
	if err != nil {
		return HBITMAP(0), err
	}
	return HBITMAP(hGdiObj), nil
}

// [SelectObject] function for [HBRUSH].
//
// [SelectObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBrush(hBrush HBRUSH) (HBRUSH, error) {
	hGdiObj, err := hdc.SelectObject(HGDIOBJ(hBrush))
	if err != nil {
		return HBRUSH(0), err
	}
	return HBRUSH(hGdiObj), nil
}

// [SelectObject] function for [HFONT].
//
// [SelectObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectFont(hFont HFONT) (HFONT, error) {
	hGdiObj, err := hdc.SelectObject(HGDIOBJ(hFont))
	if err != nil {
		return HFONT(0), err
	}
	return HFONT(hGdiObj), nil
}

// [SelectObject] function for [HPEN].
//
// [HPEN]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hpen
func (hdc HDC) SelectObjectPen(hPen HPEN) (HPEN, error) {
	hGdiObj, err := hdc.SelectObject(HGDIOBJ(hPen))
	if err != nil {
		return HPEN(0), err
	}
	return HPEN(hGdiObj), nil
}

// [SelectObject] function for [HRGN].
//
// [SelectObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectRgn(hRgn HRGN) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SelectObject, "SelectObject"),
		uintptr(hdc),
		uintptr(hRgn))
	if ret == utl.HGDI_ERROR || ret == utl.REGION_ERROR {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

// [SelectPalette] function.
//
// [SelectPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectpalette
func (hdc HDC) SelectPalette(hPal HPALETTE, forceBkgd bool) (HPALETTE, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SelectPalette, "SelectPalette"),
		uintptr(hdc),
		uintptr(hPal),
		utl.BoolToUintptr(forceBkgd))
	if ret == 0 {
		return HPALETTE(0), co.ERROR_INVALID_PARAMETER
	}
	return HPALETTE(ret), nil
}

var _SelectPalette *syscall.Proc

// [SetArcDirection] function.
//
// [SetArcDirection]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setarcdirection
func (hdc HDC) SetArcDirection(direction co.AD) (co.AD, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetArcDirection, "SetArcDirection"),
		uintptr(hdc),
		uintptr(direction))
	if ret == 0 {
		return co.AD(0), co.ERROR_INVALID_PARAMETER
	}
	return co.AD(ret), nil
}

var _SetArcDirection *syscall.Proc

// [SetBkColor] function.
//
// [SetBkColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkcolor
func (hdc HDC) SetBkColor(color COLORREF) (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetBkColor, "SetBkColor"),
		uintptr(hdc),
		uintptr(color))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

var _SetBkColor *syscall.Proc

// [SetBkMode] function.
//
// [SetBkMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
func (hdc HDC) SetBkMode(mode co.BKMODE) (co.BKMODE, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetBkMode, "SetBkMode"),
		uintptr(hdc),
		uintptr(mode))
	if ret == 0 {
		return co.BKMODE(0), co.ERROR_INVALID_PARAMETER
	}
	return co.BKMODE(ret), nil
}

var _SetBkMode *syscall.Proc

// [SetBrushOrgEx] function.
//
// [SetBrushOrgEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbrushorgex
func (hdc HDC) SetBrushOrgEx(newOrigin POINT) (POINT, error) {
	var oldOrigin POINT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetBrushOrgEx, "SetBrushOrgEx"),
		uintptr(hdc),
		uintptr(newOrigin.X),
		uintptr(newOrigin.Y),
		uintptr(unsafe.Pointer(&oldOrigin)))
	if ret == 0 {
		return POINT{}, co.ERROR_INVALID_PARAMETER
	}
	return oldOrigin, nil
}

var _SetBrushOrgEx *syscall.Proc

// [SetDIBitsToDevice] function.
//
// [SetDIBitsToDevice]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setdibitstodevice
func (hdc HDC) SetDIBitsToDevice(
	upperLeftDest POINT,
	imgSz SIZE,
	imgLowerLeft POINT,
	startScanLine int,
	numDibScanLines int,
	pColorData unsafe.Pointer,
	pBmi *BITMAPINFO,
	usage co.DIB_COLORS,
) (int, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetDIBitsToDevice, "SetDIBitsToDevice"),
		uintptr(hdc),
		uintptr(upperLeftDest.X),
		uintptr(upperLeftDest.Y),
		uintptr(uint16(imgSz.Cx)),
		uintptr(uint16(imgSz.Cy)),
		uintptr(imgLowerLeft.X),
		uintptr(imgLowerLeft.Y),
		uintptr(uint32(startScanLine)),
		uintptr(uint32(numDibScanLines)),
		uintptr(pColorData),
		uintptr(unsafe.Pointer(pBmi)),
		uintptr(usage))

	if ret == 0 || ret == utl.GDI_ERR {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return int(int32(ret)), nil
}

var _SetDIBitsToDevice *syscall.Proc

// [SetPixel] function.
//
// [SetPixel]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpixel
func (hdc HDC) SetPixel(x, y int, color COLORREF) (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetPixel, "SetPixel"),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(color))
	if int32(ret) == -1 {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

var _SetPixel *syscall.Proc

// [SetPixelFormat] function.
//
// [SetPixelFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpixelformat
func (hdc HDC) SetPixelFormat(format int, pfd *PIXELFORMATDESCRIPTOR) error {
	pfd.SetNSize() // safety
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetPixelFormat, "SetPixelFormat"),
		uintptr(hdc),
		uintptr(int32(format)),
		uintptr(unsafe.Pointer(pfd)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetPixelFormat *syscall.Proc

// [SetPolyFillMode] function.
//
// [SetPolyFillMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpolyfillmode
func (hdc HDC) SetPolyFillMode(mode co.POLYF) (co.POLYF, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetPolyFillMode, "SetPolyFillMode"),
		uintptr(hdc),
		uintptr(mode))
	if ret == 0 {
		return co.POLYF(0), co.ERROR_INVALID_PARAMETER
	}
	return co.POLYF(ret), nil
}

var _SetPolyFillMode *syscall.Proc

// [SetStretchBltMode] function.
//
// [SetStretchBltMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setstretchbltmode
func (hdc HDC) SetStretchBltMode(mode co.STRETCH) (co.STRETCH, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetStretchBltMode, "SetStretchBltMode"),
		uintptr(hdc),
		uintptr(mode))
	if ret == 0 {
		return co.STRETCH(0), co.ERROR_INVALID_PARAMETER
	}
	return co.STRETCH(ret), nil
}

var _SetStretchBltMode *syscall.Proc

// [SetTextAlign] function.
//
// [SetTextAlign]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-settextalign
func (hdc HDC) SetTextAlign(align co.TA) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetTextAlign, "SetTextAlign"),
		uintptr(hdc),
		uintptr(align))
	if ret == utl.GDI_ERR {
		return co.ERROR_INVALID_PARAMETER
	}
	return nil
}

var _SetTextAlign *syscall.Proc

// [SetTextColor] function.
//
// [SetTextColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-settextcolor
func (hdc HDC) SetTextColor(color COLORREF) (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetTextColor, "SetTextColor"),
		uintptr(hdc),
		uintptr(color))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

var _SetTextColor *syscall.Proc

// [SetViewportExtEx] function.
//
// [SetViewportExtEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setviewportextex
func (hdc HDC) SetViewportExtEx(x, y int) (SIZE, error) {
	var sz SIZE
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SetViewportExtEx, "SetViewportExtEx"),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		return SIZE{}, co.ERROR_INVALID_PARAMETER
	}
	return sz, nil
}

var _SetViewportExtEx *syscall.Proc

// [StartDoc] function.
//
// [StartDoc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-startdocw
func (hdc HDC) StartDoc(di *DOCINFO) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_StartDocW, "StartDocW"),
		uintptr(hdc),
		uintptr(unsafe.Pointer(di)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _StartDocW *syscall.Proc

// [StartPage] function.
//
// [StartPage]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-startpage
func (hdc HDC) StartPage() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_StartPage, "StartPage"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _StartPage *syscall.Proc

// [StretchBlt] function.
//
// This method is called from the destination HDC.
//
// [StretchBlt]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-stretchblt
func (hdc HDC) StretchBlt(
	destTopLeft POINT,
	destSz SIZE,
	hdcSrc HDC,
	srcTopLeft POINT,
	srcSz SIZE,
	rop co.ROP,
) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_StretchBlt, "StretchBlt"),
		uintptr(hdc),
		uintptr(destTopLeft.X),
		uintptr(destTopLeft.Y),
		uintptr(destSz.Cx),
		uintptr(destSz.Cy),
		uintptr(hdcSrc),
		uintptr(srcTopLeft.X),
		uintptr(srcTopLeft.Y),
		uintptr(srcSz.Cx),
		uintptr(srcSz.Cy),
		uintptr(rop))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _StretchBlt *syscall.Proc

// [StrokeAndFillPath] function.
//
// [StrokeAndFillPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-strokeandfillpath
func (hdc HDC) StrokeAndFillPath() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_StrokeAndFillPath, "StrokeAndFillPath"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _StrokeAndFillPath *syscall.Proc

// [StrokePath] function.
//
// [StrokePath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-strokepath
func (hdc HDC) StrokePath() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_StrokePath, "StrokePath"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _StrokePath *syscall.Proc

// [SwapBuffers] function.
//
// [SwapBuffers]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-swapbuffers
func (hdc HDC) SwapBuffers() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.GDI32, &_SwapBuffers, "SwapBuffers"),
		uintptr(hdc))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SwapBuffers *syscall.Proc

// [TextOut] function.
//
// [TextOut]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-textoutw
func (hdc HDC) TextOut(x, y int, text string) error {
	var wText wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_TextOutW, "TextOutW"),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(wText.AllowEmpty(text)),
		uintptr(int32(wstr.CountUtf16Len(text))))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _TextOutW *syscall.Proc

// [TransparentBlt] function.
//
// This method is called from the destination HDC.
//
// [TransparentBlt]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-transparentblt
func (hdc HDC) TransparentBlt(
	destTopLeft POINT,
	destSz SIZE,
	hdcSrc HDC,
	srcTopLeft POINT,
	srcSz SIZE,
	colorTransparent COLORREF,
) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_TransparentBlt, "TransparentBlt"),
		uintptr(hdc),
		uintptr(destTopLeft.X),
		uintptr(destTopLeft.Y),
		uintptr(destSz.Cx),
		uintptr(destSz.Cy),
		uintptr(hdcSrc),
		uintptr(srcTopLeft.X),
		uintptr(srcTopLeft.Y),
		uintptr(srcSz.Cx),
		uintptr(srcSz.Cy),
		uintptr(colorTransparent))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _TransparentBlt *syscall.Proc

// [WidenPath] function.
//
// [WidenPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-widenpath
func (hdc HDC) WidenPath() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_WidenPath, "WidenPath"),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _WidenPath *syscall.Proc
