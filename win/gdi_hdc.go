//go:build windows

package win

import (
	"fmt"
	"syscall"
	"unicode/utf8"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
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
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pDriver := wbuf.PtrEmptyIsNil(driver)
	pDevice := wbuf.PtrAllowEmpty(device)

	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_CreateDCW),
		uintptr(pDriver),
		uintptr(pDevice),
		0,
		uintptr(unsafe.Pointer(dm)))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

// [CreateIC] function.
//
// ⚠️ You must defer [HDC.DeleteDC].
//
// [CreateIC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createicw
func CreateIC(driver, device string, dm *DEVMODE) (HDC, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pDriver := wbuf.PtrAllowEmpty(driver)
	pDevice := wbuf.PtrAllowEmpty(device)

	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_CreateICW),
		uintptr(pDriver),
		uintptr(pDevice),
		0,
		uintptr(unsafe.Pointer(dm)))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

// [AbortDoc] function.
//
// [AbortDoc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-abortdoc
func (hdc HDC) AbortDoc() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_AbortDoc),
		uintptr(hdc))
	return utl.Minus1AsSysInvalidParm(ret)
}

// [AbortPath] function.
//
// [AbortPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-abortpath
func (hdc HDC) AbortPath() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_AbortPath),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

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
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_AlphaBlend),
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

// [AngleArc] function.
//
// [AngleArc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-anglearc
func (hdc HDC) AngleArc(center POINT, r uint, startAngle, sweepAngle float32) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_AngleArc),
		uintptr(hdc),
		uintptr(center.X),
		uintptr(center.Y),
		uintptr(uint32(r)),
		uintptr(startAngle),
		uintptr(sweepAngle))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [Arc] function.
//
// [Arc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arc
func (hdc HDC) Arc(bound RECT, radialStart, radialEnd POINT) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_Arc),
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

// [ArcTo] function.
//
// [ArcTo]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arcto
func (hdc HDC) ArcTo(bound RECT, radialStart, radialEnd POINT) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_ArcTo),
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

// [BeginPath] function.
//
// ⚠️ You must defer [HDC.EndPath].
//
// [BeginPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-beginpath
func (hdc HDC) BeginPath() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_BeginPath),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

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
	ret, _, err := syscall.SyscallN(dll.Gdi(dll.PROC_BitBlt),
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

// [CancelDC] function.
//
// [CancelDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-canceldc
func (hdc HDC) CancelDC() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_CancelDC),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [Chord] function.
//
// [Chord]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-chord
func (hdc HDC) Chord(bound RECT, radialStart, radialEnd POINT) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_Chord),
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

// [CloseFigure] function.
//
// [CloseFigure]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-closefigure
func (hdc HDC) CloseFigure() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_CloseFigure),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [ChoosePixelFormat] function.
//
// [ChoosePixelFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-choosepixelformat
func (hdc HDC) ChoosePixelFormat(pfd *PIXELFORMATDESCRIPTOR) (int, error) {
	ret, _, err := syscall.SyscallN(dll.Gdi(dll.PROC_ChoosePixelFormat),
		uintptr(hdc),
		uintptr(unsafe.Pointer(pfd)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return int(ret), nil
}

// [CreateCompatibleBitmap] function.
//
// ⚠️ You must defer [HBITMAP.DeleteObject].
//
// [CreateCompatibleBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatiblebitmap
func (hdc HDC) CreateCompatibleBitmap(cx, cy uint) (HBITMAP, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_CreateCompatibleBitmap),
		uintptr(hdc),
		uintptr(int32(cx)),
		uintptr(int32(cy)))
	if ret == 0 {
		return HBITMAP(0), co.ERROR_INVALID_PARAMETER
	}
	return HBITMAP(ret), nil
}

// [CreateCompatibleDC] function.
//
// ⚠️ You must defer [HDC.DeleteDC].
//
// [CreateCompatibleDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatibledc
func (hdc HDC) CreateCompatibleDC() (HDC, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_CreateCompatibleDC),
		uintptr(hdc))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

// [CreateDIBSection] function.
//
// ⚠️ You must defer [HBITMAP.DeleteObject].
//
// [CreateDIBSection]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdibsection
func (hdc HDC) CreateDIBSection(
	bmi *BITMAPINFO,
	usage co.DIB,
	hSection HFILEMAP,
	offset uint,
) (HBITMAP, *byte, error) {
	var ppvBits *byte
	ret, _, err := syscall.SyscallN(dll.Gdi(dll.PROC_CreateDIBSection),
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

// [CreateHalftonePalette] function.
//
// ⚠️ You must defer [HPALETTE.DeleteObject].
//
// [CreateHalftonePalette]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createhalftonepalette
func (hdc HDC) CreateHalftonePalette() (HPALETTE, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_CreateHalftonePalette),
		uintptr(hdc))
	if ret == 0 {
		return HPALETTE(0), co.ERROR_INVALID_PARAMETER
	}
	return HPALETTE(ret), nil
}

// [DeleteDC] function.
//
// [DeleteDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deletedc
func (hdc HDC) DeleteDC() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_DeleteDC),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [Ellipse] function.
//
// [Ellipse]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ellipse
func (hdc HDC) Ellipse(bound RECT) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_Ellipse),
		uintptr(hdc),
		uintptr(bound.Left),
		uintptr(bound.Top),
		uintptr(bound.Right),
		uintptr(bound.Bottom))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [EndDoc] function.
//
// [EndDoc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-enddoc
func (hdc HDC) EndDoc() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_EndDoc),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [EndPage] function.
//
// [EndPage]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-endpage
func (hdc HDC) EndPage() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_EndPage),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [EndPath] function.
//
// Paired with [HDC.BeginPath].
//
// [EndPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-endpath
func (hdc HDC) EndPath() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_EndPath),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [ExcludeClipRect] function.
//
// [ExcludeClipRect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-excludecliprect
func (hdc HDC) ExcludeClipRect(rc RECT) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_ExcludeClipRect),
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

// [FillPath] function.
//
// [FillPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillpath
func (hdc HDC) FillPath() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_FillPath),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [FillRect] function.
//
// [FillRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-fillrect
func (hdc HDC) FillRect(rc *RECT, hBrush HBRUSH) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_FillRect),
		uintptr(hdc),
		uintptr(unsafe.Pointer(rc)),
		uintptr(hBrush))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [FillRgn] function.
//
// [FillRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillrgn
func (hdc HDC) FillRgn(hRgn HRGN, hBrush HBRUSH) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_FillRgn),
		uintptr(hdc),
		uintptr(hRgn),
		uintptr(hBrush))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [FlattenPath] function.
//
// [FlattenPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-flattenpath
func (hdc HDC) FlattenPath() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_FlattenPath),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [FrameRgn] function.
//
// [FrameRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-framergn
func (hdc HDC) FrameRgn(hRgn HRGN, hBrush HBRUSH, width, height uint) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_FrameRgn),
		uintptr(hdc),
		uintptr(hRgn),
		uintptr(hBrush),
		uintptr(int32(width)),
		uintptr(int32(height)))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [GetBkColor] function.
//
// [GetBkColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getbkcolor
func (hdc HDC) GetBkColor() (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetBkColor),
		uintptr(hdc))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

// [GetBkMode] function.
//
// [GetBkMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getbkmode
func (hdc HDC) GetBkMode() (co.BKMODE, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetBkMode),
		uintptr(hdc))
	if ret == 0 {
		return co.BKMODE(0), co.ERROR_INVALID_PARAMETER
	}
	return co.BKMODE(ret), nil
}

// [GetCurrentPositionEx] function.
//
// [GetCurrentPositionEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getcurrentpositionex
func (hdc HDC) GetCurrentPositionEx() (POINT, error) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetCurrentPositionEx),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR_INVALID_PARAMETER
	}
	return pt, nil
}

// [GetDCBrushColor] function.
//
// [GetDCBrushColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdcbrushcolor
func (hdc HDC) GetDCBrushColor() (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetDCBrushColor),
		uintptr(hdc))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

// [GetDCPenColor] function.
//
// [GetDCPenColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdcpencolor
func (hdc HDC) GetDCPenColor() (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetDCPenColor),
		uintptr(hdc))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

// [GetDeviceCaps] function.
//
// [GetDeviceCaps]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdevicecaps
func (hdc HDC) GetDeviceCaps(index co.GDC) int32 {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetDeviceCaps),
		uintptr(hdc),
		uintptr(index))
	return int32(ret)
}

// [GetDIBits] function.
//
// Note that this method fails if bitmapDataBuffer is an ordinary Go slice; it
// must be allocated directly from the OS heap, for example with [GlobalAlloc].
//
// # Example
//
// Taking a screenshot and saving into a BMP file:
//
//	cxScreen := win.GetSystemMetrics(co.SM_CXSCREEN)
//	cyScreen := win.GetSystemMetrics(co.SM_CYSCREEN)
//
//	hdcScreen, _ := win.HWND(0).GetDC()
//	defer win.HWND(0).ReleaseDC(hdcScreen)
//
//	hBmp, _ := hdcScreen.CreateCompatibleBitmap(uint(cxScreen), uint(cyScreen))
//	defer hBmp.DeleteObject()
//
//	hdcMem, _ := hdcScreen.CreateCompatibleDC()
//	defer hdcMem.DeleteDC()
//
//	hBmpOld, _ := hdcMem.SelectObjectBmp(hBmp)
//	defer hdcMem.SelectObjectBmp(hBmpOld)
//
//	hdcMem.BitBlt(
//		win.POINT{X: 0, Y: 0},
//		win.SIZE{Cx: cxScreen, Cy: cyScreen},
//		hdcScreen,
//		win.POINT{X: 0, Y: 0},
//		co.ROP_SRCCOPY,
//	)
//
//	bi := win.BITMAPINFO{
//		BmiHeader: win.BITMAPINFOHEADER{
//			BiWidth:       cxScreen,
//			BiHeight:      cyScreen,
//			BiPlanes:      1,
//			BiBitCount:    32,
//			BiCompression: co.BI_RGB,
//		},
//	}
//	bi.BmiHeader.SetBiSize()
//
//	bmpObj, _ := hBmp.GetObject()
//	bmpSize := bmpObj.CalcBitmapSize(bi.BmiHeader.BiBitCount)
//
//	rawMem, _ := win.GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, bmpSize)
//	defer rawMem.GlobalFree()
//
//	bmpSlice, _ := rawMem.GlobalLockSlice()
//	defer rawMem.GlobalUnlock()
//
//	hdcScreen.GetDIBits(hBmp, 0, uint(cyScreen), bmpSlice, &bi, co.DIB_RGB_COLORS)
//
//	var bfh win.BITMAPFILEHEADER
//	bfh.SetBfType()
//	bfh.SetBfOffBits(uint32(unsafe.Sizeof(bfh) + unsafe.Sizeof(bi.BmiHeader)))
//	bfh.SetBfSize(bfh.BfOffBits() + uint32(bmpSize))
//
//	fo, _ := win.FileOpen("C:\\Temp\\screenshot.bmp", co.FOPEN_RW_OPEN_OR_CREATE)
//	defer fo.Close()
//
//	fo.Write(bfh.Serialize())
//	fo.Write(bi.BmiHeader.Serialize())
//	fo.Write(bmpSlice)
//
// [GetDIBits]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdibits
func (hdc HDC) GetDIBits(
	hbm HBITMAP,
	firstScanLine, numScanLines uint,
	bitmapDataBuffer []byte,
	bmi *BITMAPINFO,
	usage co.DIB,
) (int, error) {
	var dataBufPtr *byte
	if bitmapDataBuffer != nil {
		dataBufPtr = &bitmapDataBuffer[0]
	}

	bmi.BmiHeader.SetBiSize() // safety

	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetDIBits),
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
	return int(ret), nil
}

// [GetPixel] function.
//
// [GetPixel]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getpixel
func (hdc HDC) GetPixel() (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetPixel),
		uintptr(hdc))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

// [GetPolyFillMode] function.
//
// [GetPolyFillMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getpolyfillmode
func (hdc HDC) GetPolyFillMode() (co.POLYF, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetPolyFillMode),
		uintptr(hdc))
	if ret == 0 {
		return co.POLYF(0), co.ERROR_INVALID_PARAMETER
	}
	return co.POLYF(ret), nil
}

// [GetTextColor] function.
//
// [GetTextColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextcolor
func (hdc HDC) GetTextColor() (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetTextColor),
		uintptr(hdc))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

// [GetTextExtentPoint32] function.
//
// [GetTextExtentPoint32]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextextentpoint32w
func (hdc HDC) GetTextExtentPoint32(text string) (SIZE, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pText := wbuf.PtrAllowEmpty(text)

	textLen := utf8.RuneCountInString(text)
	var sz SIZE

	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetTextExtentPoint32W),
		uintptr(hdc),
		uintptr(pText),
		uintptr(textLen),
		uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		return SIZE{}, co.ERROR_INVALID_PARAMETER
	}
	return sz, nil
}

// [GetTextFace] function.
//
// [GetTextFace]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextfacew
func (hdc HDC) GetTextFace() (string, error) {
	var buf [utl.LF_FACESIZE]uint16
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetTextFaceW),
		uintptr(hdc),
		uintptr(len(buf)),
		uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 {
		return "", co.ERROR_INVALID_PARAMETER
	}
	return wstr.WinSliceToGo(buf[:]), nil
}

// [GetTextMetrics] function.
//
// [GetTextMetrics]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextmetricsw
func (hdc HDC) GetTextMetrics() (TEXTMETRIC, error) {
	var tm TEXTMETRIC
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetTextMetricsW),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&tm)))
	if ret == 0 {
		return TEXTMETRIC{}, co.ERROR_INVALID_PARAMETER
	}
	return tm, nil
}

// [GetViewportExtEx] function.
//
// [GetViewportExtEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getviewportextex
func (hdc HDC) GetViewportExtEx() (SIZE, error) {
	var sz SIZE
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetViewportExtEx),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		return SIZE{}, co.ERROR_INVALID_PARAMETER
	}
	return sz, nil
}

// [GetViewportOrgEx] function.
//
// [GetViewportOrgEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getviewportorgex
func (hdc HDC) GetViewportOrgEx() (POINT, error) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetViewportOrgEx),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR_INVALID_PARAMETER
	}
	return pt, nil
}

// [GetWindowExtEx] function.
//
// [GetWindowExtEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getwindowextex
func (hdc HDC) GetWindowExtEx() (SIZE, error) {
	var sz SIZE
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetWindowExtEx),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		return SIZE{}, co.ERROR_INVALID_PARAMETER
	}
	return sz, nil
}

// [GetWindowOrgEx] function.
//
// [GetWindowOrgEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getwindoworgex
func (hdc HDC) GetWindowOrgEx() (POINT, error) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GetWindowOrgEx),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR_INVALID_PARAMETER
	}
	return pt, nil
}

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

	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_GradientFill),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&vertex[0])),
		uintptr(uint32(len(vertex))),
		uintptr(pMesh),
		uintptr(uint32(nMesh)),
		uintptr(mode))
	return utl.ZeroAsSysInvalidParm(ret)
}

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
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_IntersectClipRect),
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

// [InvertRgn] function.
//
// [InvertRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-invertrgn
func (hdc HDC) InvertRgn(hRgn HRGN) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_InvertRgn),
		uintptr(hdc),
		uintptr(hRgn))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [LineTo] function.
//
// [LineTo]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lineto
func (hdc HDC) LineTo(x, y int) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_LineTo),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [LPtoDP] function.
//
// [LPtoDP]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lptodp
func (hdc HDC) LPtoDP(pts []POINT) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_LPtoDP),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(int32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

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
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_MaskBlt),
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

// [MoveToEx] function.
//
// [MoveToEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-movetoex
func (hdc HDC) MoveToEx(x, y int) (POINT, error) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_MoveToEx),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR_INVALID_PARAMETER
	}
	return pt, nil
}

// [PaintRgn] function.
//
// [PaintRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-paintrgn
func (hdc HDC) PaintRgn(hRgn HRGN) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_PaintRgn),
		uintptr(hdc),
		uintptr(hRgn))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [PatBlt] function.
//
// [PatBlt]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-patblt
func (hdc HDC) PatBlt(topLeft POINT, sz SIZE, rop co.ROP) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_PatBlt),
		uintptr(hdc),
		uintptr(topLeft.X),
		uintptr(topLeft.Y),
		uintptr(sz.Cx),
		uintptr(sz.Cy),
		uintptr(rop))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [PathToRegion] function.
//
// ⚠️ You must defer [HRGN.DeleteObject] on the returned HRGN.
//
// [PathToRegion]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-pathtoregion
func (hdc HDC) PathToRegion() (HRGN, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_PathToRegion),
		uintptr(hdc))
	if ret == 0 {
		return HRGN(0), co.ERROR_INVALID_PARAMETER
	}
	return HRGN(ret), nil
}

// [Pie] function.
//
// [Pie]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-pie
func (hdc HDC) Pie(bound RECT, endPointRadial1, endPointRadial2 POINT) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_Pie),
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
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_PolyBezier),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(uint32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [PolyBezierTo] function.
//
// [PolyBezierTo]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polybezierto
func (hdc HDC) PolyBezierTo(pts []POINT) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_PolyBezierTo),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(uint32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

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

	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_PolyDraw),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(unsafe.Pointer(&usage[0])),
		uintptr(int32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [Polygon] function.
//
// [Polygon]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polygon
func (hdc HDC) Polygon(pts []POINT) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_Polygon),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(uint32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [Polyline] function.
//
// [Polyline]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polyline
func (hdc HDC) Polyline(pts []POINT) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_Polyline),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(int32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [PolylineTo] function.
//
// [PolylineTo]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polylineto
func (hdc HDC) PolylineTo(pts []POINT) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_PolylineTo),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pts[0])),
		uintptr(uint32(len(pts))))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [PolyPolygon] function.
//
// [PolyPolygon]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polypolygon
func (hdc HDC) PolyPolygon(polygons [][]POINT) error {
	numPolygons := uint(len(polygons))
	totalPts := uint(0)
	for _, polygon := range polygons {
		totalPts += uint(len(polygon))
	}

	allPtsFlat := make([]POINT, 0, totalPts)
	ptsPerPolygon := make([]int32, 0, numPolygons)

	for _, polygon := range polygons {
		allPtsFlat = append(allPtsFlat, polygon...)
		ptsPerPolygon = append(ptsPerPolygon, int32(len(polygon)))
	}

	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_PolyPolygon),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&allPtsFlat[0])),
		uintptr(unsafe.Pointer(&ptsPerPolygon[0])),
		uintptr(int32(len(polygons))))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [PolyPolyline] function.
//
// [PolyPolyline]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polypolyline
func (hdc HDC) PolyPolyline(polyLines [][]POINT) error {
	numPolyLines := uint(len(polyLines))
	totalPts := uint(0)
	for _, polygon := range polyLines {
		totalPts += uint(len(polygon))
	}

	allPtsFlat := make([]POINT, 0, totalPts)
	ptsPerPolyLine := make([]int32, 0, numPolyLines)

	for _, polyLine := range polyLines {
		allPtsFlat = append(allPtsFlat, polyLine...)
		ptsPerPolyLine = append(ptsPerPolyLine, int32(len(polyLine)))
	}

	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_PolyPolyline),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&allPtsFlat[0])),
		uintptr(unsafe.Pointer(&ptsPerPolyLine[0])),
		uintptr(uint32(len(polyLines))))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [PtVisible] function.
//
// [PtVisible]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ptvisible
func (hdc HDC) PtVisible(x, y int) (bool, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_PtVisible),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)))
	if int32(ret) == -1 {
		return false, co.ERROR_INVALID_PARAMETER
	}
	return ret != 0, nil
}

// [RealizePalette] function.
//
// [RealizePalette]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-realizepalette
func (hdc HDC) RealizePalette() (uint, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_RealizePalette),
		uintptr(hdc))
	if ret == utl.GDI_ERR {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return uint(ret), nil
}

// [Rectangle] function.
//
// [Rectangle]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rectangle
func (hdc HDC) Rectangle(bound RECT) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_Rectangle),
		uintptr(hdc),
		uintptr(bound.Left),
		uintptr(bound.Top),
		uintptr(bound.Right),
		uintptr(bound.Bottom))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [ResetDC] function.
//
// [ResetDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-resetdcw
func (hdc HDC) ResetDC(dm *DEVMODE) (HDC, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_ResetDCW),
		uintptr(hdc),
		uintptr(unsafe.Pointer(dm)))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

// [RestoreDC] function.
//
// Paired with [HDC.SaveDC].
//
// [RestoreDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-restoredc
func (hdc HDC) RestoreDC(savedDC int32) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_RestoreDC),
		uintptr(hdc),
		uintptr(savedDC))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [RoundRect] function.
//
// [RoundRect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-roundrect
func (hdc HDC) RoundRect(bound RECT, sz SIZE) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_RoundRect),
		uintptr(hdc),
		uintptr(bound.Left),
		uintptr(bound.Top),
		uintptr(bound.Right),
		uintptr(bound.Bottom),
		uintptr(sz.Cx),
		uintptr(sz.Cy))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [SaveDC] function.
//
// Paired with [HDC.RestoreDC].
//
// [SaveDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-savedc
func (hdc HDC) SaveDC() (int32, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SaveDC),
		uintptr(hdc))
	if ret == 0 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return int32(ret), nil
}

// [SelectClipPath] function.
//
// [SelectClipPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectclippath
func (hdc HDC) SelectClipPath(mode co.RGN) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SelectClipPath),
		uintptr(hdc),
		uintptr(mode))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [SelectClipRgn] function.
//
// [SelectClipRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectcliprgn
func (hdc HDC) SelectClipRgn(hRgn HRGN) (co.REGION, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SelectClipRgn),
		uintptr(hdc),
		uintptr(hRgn))
	if ret == utl.REGION_ERROR {
		return co.REGION(0), co.ERROR_INVALID_PARAMETER
	}
	return co.REGION(ret), nil
}

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
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SelectObject),
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
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SelectPalette),
		uintptr(hdc),
		uintptr(hPal),
		utl.BoolToUintptr(forceBkgd))
	if ret == 0 {
		return HPALETTE(0), co.ERROR_INVALID_PARAMETER
	}
	return HPALETTE(ret), nil
}

// [SetArcDirection] function.
//
// [SetArcDirection]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setarcdirection
func (hdc HDC) SetArcDirection(direction co.AD) (co.AD, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SetArcDirection),
		uintptr(hdc),
		uintptr(direction))
	if ret == 0 {
		return co.AD(0), co.ERROR_INVALID_PARAMETER
	}
	return co.AD(ret), nil
}

// [SetBkColor] function.
//
// [SetBkColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkcolor
func (hdc HDC) SetBkColor(color COLORREF) (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SetBkColor),
		uintptr(hdc),
		uintptr(color))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

// [SetBkMode] function.
//
// [SetBkMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
func (hdc HDC) SetBkMode(mode co.BKMODE) (co.BKMODE, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SetBkMode),
		uintptr(hdc),
		uintptr(mode))
	if ret == 0 {
		return co.BKMODE(0), co.ERROR_INVALID_PARAMETER
	}
	return co.BKMODE(ret), nil
}

// [SetBrushOrgEx] function.
//
// [SetBrushOrgEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbrushorgex
func (hdc HDC) SetBrushOrgEx(newOrigin POINT) (POINT, error) {
	var oldOrigin POINT
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SetBrushOrgEx),
		uintptr(hdc),
		uintptr(newOrigin.X),
		uintptr(newOrigin.Y),
		uintptr(unsafe.Pointer(&oldOrigin)))
	if ret == 0 {
		return POINT{}, co.ERROR_INVALID_PARAMETER
	}
	return oldOrigin, nil
}

// [SetPixel] function.
//
// [SetPixel]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpixel
func (hdc HDC) SetPixel(x, y int, color COLORREF) (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SetPixel),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(color))
	if int32(ret) == -1 {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

// [SetPixelFormat] function.
//
// [SetPixelFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpixelformat
func (hdc HDC) SetPixelFormat(format int, pfd *PIXELFORMATDESCRIPTOR) error {
	ret, _, err := syscall.SyscallN(dll.Gdi(dll.PROC_SetPixelFormat),
		uintptr(hdc),
		uintptr(format),
		uintptr(unsafe.Pointer(pfd)))
	return utl.ZeroAsGetLastError(ret, err)
}

// [SetPolyFillMode] function.
//
// [SetPolyFillMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpolyfillmode
func (hdc HDC) SetPolyFillMode(mode co.POLYF) (co.POLYF, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SetPolyFillMode),
		uintptr(hdc),
		uintptr(mode))
	if ret == 0 {
		return co.POLYF(0), co.ERROR_INVALID_PARAMETER
	}
	return co.POLYF(ret), nil
}

// [SetStretchBltMode] function.
//
// [SetStretchBltMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setstretchbltmode
func (hdc HDC) SetStretchBltMode(mode co.STRETCH) (co.STRETCH, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SetStretchBltMode),
		uintptr(hdc),
		uintptr(mode))
	if ret == 0 {
		return co.STRETCH(0), co.ERROR_INVALID_PARAMETER
	}
	return co.STRETCH(ret), nil
}

// [SetTextAlign] function.
//
// [SetTextAlign]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-settextalign
func (hdc HDC) SetTextAlign(align co.TA) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SetTextAlign),
		uintptr(hdc),
		uintptr(align))
	if ret == utl.GDI_ERR {
		return co.ERROR_INVALID_PARAMETER
	}
	return nil
}

// [SetTextColor] function.
//
// [SetTextColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-settextcolor
func (hdc HDC) SetTextColor(color COLORREF) (COLORREF, error) {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SetTextColor),
		uintptr(hdc),
		uintptr(color))
	if ret == utl.CLR_INVALID {
		return COLORREF(0), co.ERROR_INVALID_PARAMETER
	}
	return COLORREF(ret), nil
}

// [SetViewportExtEx] function.
//
// [SetViewportExtEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setviewportextex
func (hdc HDC) SetViewportExtEx(x, y int) (SIZE, error) {
	var sz SIZE
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_SetViewportExtEx),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		return SIZE{}, co.ERROR_INVALID_PARAMETER
	}
	return sz, nil
}

// [StartDoc] function.
//
// [StartDoc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-startdocw
func (hdc HDC) StartDoc(di *DOCINFO) error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_StartDocW),
		uintptr(hdc),
		uintptr(unsafe.Pointer(di)))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [StartPage] function.
//
// [StartPage]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-startpage
func (hdc HDC) StartPage() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_StartPage),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

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
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_StretchBlt),
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

// [StrokeAndFillPath] function.
//
// [StrokeAndFillPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-strokeandfillpath
func (hdc HDC) StrokeAndFillPath() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_StrokeAndFillPath),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [StrokePath] function.
//
// [StrokePath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-strokepath
func (hdc HDC) StrokePath() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_StrokePath),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [SwapBuffers] function.
//
// [SwapBuffers]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-swapbuffers
func (hdc HDC) SwapBuffers() error {
	ret, _, err := syscall.SyscallN(dll.Gdi(dll.PROC_SwapBuffers),
		uintptr(hdc))
	return utl.ZeroAsGetLastError(ret, err)
}

// [TextOut] function.
//
// [TextOut]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-textoutw
func (hdc HDC) TextOut(x, y int, text string) error {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pText := wbuf.PtrAllowEmpty(text)

	textLen := utf8.RuneCountInString(text)

	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_TextOutW),
		uintptr(hdc),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(pText),
		uintptr(int32(textLen-1)))
	return utl.ZeroAsSysInvalidParm(ret)
}

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
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_TransparentBlt),
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

// [WidenPath] function.
//
// [WidenPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-widenpath
func (hdc HDC) WidenPath() error {
	ret, _, _ := syscall.SyscallN(dll.Gdi(dll.PROC_WidenPath),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}
