//go:build windows

package win

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [device context] (DC).
//
// [device context]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdc
type HDC HANDLE

// [AbortDoc] function.
//
// [AbortDoc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-abortdoc
func (hdc HDC) AbortDoc() {
	ret, _, _ := syscall.SyscallN(proc.AbortDoc.Addr(),
		uintptr(hdc))
	if int32(ret) == -1 {
		panic(errco.INVALID_PRINTER_COMMAND)
	}
}

// [AbortPath] function.
//
// [AbortPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-abortpath
func (hdc HDC) AbortPath() {
	ret, _, err := syscall.SyscallN(proc.AbortPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
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
	ftn BLENDFUNCTION) {

	ret, _, err := syscall.SyscallN(proc.AlphaBlend.Addr(),
		uintptr(hdc), uintptr(originDest.X), uintptr(originDest.Y),
		uintptr(szDest.Cx), uintptr(szDest.Cy),
		uintptr(hdcSrc), uintptr(originSrc.X), uintptr(originSrc.Y),
		uintptr(szSrc.Cx), uintptr(szSrc.Cy),
		uintptr(
			util.Make32(
				util.Make16(ftn.BlendOp, ftn.BlendFlags),
				util.Make16(ftn.SourceConstantAlpha, ftn.AlphaFormat),
			),
		))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [AngleArc] function.
//
// [AngleArc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-anglearc
func (hdc HDC) AngleArc(center POINT, r uint32, startAngle, sweepAngle float32) {
	ret, _, err := syscall.SyscallN(proc.AngleArc.Addr(),
		uintptr(hdc), uintptr(center.X), uintptr(center.Y), uintptr(r),
		uintptr(startAngle), uintptr(sweepAngle))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [Arc] function.
//
// [Arc]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arc
func (hdc HDC) Arc(bound RECT, radialStart, radialEnd POINT) {
	ret, _, err := syscall.SyscallN(proc.Arc.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(radialStart.X), uintptr(radialStart.Y),
		uintptr(radialEnd.X), uintptr(radialEnd.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [ArcTo] function.
//
// [ArcTo]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arcto
func (hdc HDC) ArcTo(bound RECT, radialStart, radialEnd POINT) {
	ret, _, err := syscall.SyscallN(proc.ArcTo.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(radialStart.X), uintptr(radialStart.Y),
		uintptr(radialEnd.X), uintptr(radialEnd.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [BeginPath] function.
//
// ⚠️ You must defer HDC.EndPath().
//
// [BeginPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-beginpath
func (hdc HDC) BeginPath() {
	ret, _, err := syscall.SyscallN(proc.BeginPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
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
	rop co.ROP) {

	ret, _, err := syscall.SyscallN(proc.BitBlt.Addr(),
		uintptr(hdc), uintptr(destTopLeft.X), uintptr(destTopLeft.Y),
		uintptr(sz.Cx), uintptr(sz.Cy),
		uintptr(hdcSrc), uintptr(srcTopLeft.X), uintptr(srcTopLeft.Y),
		uintptr(rop))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [CancelDC] function.
//
// [CancelDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-canceldc
func (hdc HDC) CancelDC() {
	ret, _, err := syscall.SyscallN(proc.CancelDC.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [Chord] function.
//
// [Chord]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-chord
func (hdc HDC) Chord(bound RECT, radialStart, radialEnd POINT) {
	ret, _, err := syscall.SyscallN(proc.Chord.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(radialStart.X), uintptr(radialStart.Y),
		uintptr(radialEnd.X), uintptr(radialEnd.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [CloseFigure] function.
//
// [CloseFigure]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-closefigure
func (hdc HDC) CloseFigure() {
	ret, _, err := syscall.SyscallN(proc.CloseFigure.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [CreateCompatibleBitmap] function.
//
// ⚠️ You must defer HBITMAP.DeleteObject().
//
// [CreateCompatibleBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatiblebitmap
func (hdc HDC) CreateCompatibleBitmap(cx, cy int32) HBITMAP {
	ret, _, err := syscall.SyscallN(proc.CreateCompatibleBitmap.Addr(),
		uintptr(hdc), uintptr(cx), uintptr(cy))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret)
}

// [CreateCompatibleDC] function.
//
// ⚠️ You must defer HDC.DeleteDC().
//
// [CreateCompatibleDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatibledc
func (hdc HDC) CreateCompatibleDC() HDC {
	ret, _, err := syscall.SyscallN(proc.CreateCompatibleDC.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HDC(ret)
}

// [CreateDIBSection] function.
//
// ⚠️ You must defer HBITMAP.DeleteObject().
//
// [CreateDIBSection]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdibsection
func (hdc HDC) CreateDIBSection(
	bmi *BITMAPINFO,
	usage co.DIB,
	hSection HFILEMAP,
	offset uint32) (HBITMAP, *byte) {

	var ppvBits *byte
	ret, _, err := syscall.SyscallN(proc.CreateDIBSection.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(bmi)), uintptr(usage),
		uintptr(unsafe.Pointer(&ppvBits)), uintptr(hSection), uintptr(offset))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret), ppvBits
}

// [DeleteDC] function.
//
// [DeleteDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deletedc
func (hdc HDC) DeleteDC() {
	ret, _, err := syscall.SyscallN(proc.DeleteDC.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [Ellipse] function.
//
// [Ellipse]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ellipse
func (hdc HDC) Ellipse(bound RECT) {
	ret, _, err := syscall.SyscallN(proc.Ellipse.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [EndPath] function.
//
// [EndPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-endpath
func (hdc HDC) EndPath() {
	ret, _, err := syscall.SyscallN(proc.EndPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [FillPath] function.
//
// [FillPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillpath
func (hdc HDC) FillPath() {
	ret, _, err := syscall.SyscallN(proc.FillPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [FillRect] function.
//
// [FillRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-fillrect
func (hdc HDC) FillRect(rc *RECT, hBrush HBRUSH) {
	ret, _, err := syscall.SyscallN(proc.FillRect.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(rc)), uintptr(hBrush))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [FillRgn] function.
//
// [FillRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillrgn
func (hdc HDC) FillRgn(hRgn HRGN, hBrush HBRUSH) {
	ret, _, err := syscall.SyscallN(proc.FillRgn.Addr(),
		uintptr(hdc), uintptr(hRgn), uintptr(hBrush))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [FlattenPath] function.
//
// [FlattenPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-flattenpath
func (hdc HDC) FlattenPath() {
	ret, _, err := syscall.SyscallN(proc.FlattenPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [FrameRgn] function.
//
// [FrameRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-framergn
func (hdc HDC) FrameRgn(hRgn HRGN, hBrush HBRUSH, w, h int32) {
	ret, _, err := syscall.SyscallN(proc.FrameRgn.Addr(),
		uintptr(hdc), uintptr(hRgn), uintptr(hBrush), uintptr(w), uintptr(h))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [GetCurrentPositionEx] function.
//
// [GetCurrentPositionEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getcurrentpositionex
func (hdc HDC) GetCurrentPositionEx() POINT {
	var pt POINT
	ret, _, err := syscall.SyscallN(proc.GetCurrentPositionEx.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// [GetDCBrushColor] function.
//
// [GetDCBrushColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdcbrushcolor
func (hdc HDC) GetDCBrushColor() COLORREF {
	ret, _, err := syscall.SyscallN(proc.GetDCBrushColor.Addr(),
		uintptr(hdc))
	if ret == _CLR_INVALID {
		panic(errco.ERROR(err))
	}
	return COLORREF(ret)
}

// [GetDCPenColor] function.
//
// [GetDCPenColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdcpencolor
func (hdc HDC) GetDCPenColor() COLORREF {
	ret, _, err := syscall.SyscallN(proc.GetDCPenColor.Addr(),
		uintptr(hdc))
	if ret == _CLR_INVALID {
		panic(errco.ERROR(err))
	}
	return COLORREF(ret)
}

// [GetDeviceCaps] function.
//
// [GetDeviceCaps]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdevicecaps
func (hdc HDC) GetDeviceCaps(index co.GDC) int32 {
	ret, _, _ := syscall.SyscallN(proc.GetDeviceCaps.Addr(),
		uintptr(hdc), uintptr(index))
	return int32(ret)
}

// [GetDIBits] function.
//
// Note that this method fails if bitmapDataBuffer is an ordinary Go slice; it
// must be allocated directly from the OS heap.
//
// # Example
//
// Taking a screenshot and saving into a BMP file:
//
//	cxScreen := win.GetSystemMetrics(co.SM_CXSCREEN)
//	cyScreen := win.GetSystemMetrics(co.SM_CYSCREEN)
//
//	hdcScreen := win.HWND(0).GetDC()
//	defer win.HWND(0).ReleaseDC(hdcScreen)
//
//	hBmp := hdcScreen.CreateCompatibleBitmap(cxScreen, cyScreen)
//	defer hBmp.DeleteObject()
//
//	hdcMem := hdcScreen.CreateCompatibleDC()
//	defer hdcMem.DeleteDC()
//
//	hBmpOld := hdcMem.SelectObjectBitmap(hBmp)
//	defer hdcMem.SelectObjectBitmap(hBmpOld)
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
//	var bmpObj win.BITMAP
//	hBmp.GetObject(&bmpObj)
//	bmpSize := bmpObj.CalcBitmapSize(bi.BmiHeader.BiBitCount)
//
//	rawMem := win.GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, bmpSize)
//	defer rawMem.GlobalFree()
//
//	bmpSlice := rawMem.GlobalLock(bmpSize)
//	defer rawMem.GlobalUnlock()
//
//	hdcScreen.GetDIBits(hBmp, 0, int(cyScreen), bmpSlice, &bi, co.DIB_RGB_COLORS)
//
//	var bfh win.BITMAPFILEHEADER
//	bfh.SetBfType()
//	bfh.SetBfOffBits(uint32(unsafe.Sizeof(bfh) + unsafe.Sizeof(bi.BmiHeader)))
//	bfh.SetBfSize(bfh.BfOffBits() + uint32(bmpSize))
//
//	fo, _ := win.FileOpen("C:\\Temp\\foo.bmp", co.FILE_OPEN_RW_OPEN_OR_CREATE)
//	defer fo.Close()
//
//	fo.Write(bfh.Serialize())
//	fo.Write(bi.BmiHeader.Serialize())
//	fo.Write(bmpSlice)
//
// [GetDIBits]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdibits
func (hdc HDC) GetDIBits(
	hbm HBITMAP,
	firstScanLine, numScanLines int,
	bitmapDataBuffer []byte,
	bmi *BITMAPINFO,
	usage co.DIB) int {

	var dataBufPtr *byte
	if bitmapDataBuffer != nil {
		dataBufPtr = &bitmapDataBuffer[0]
	}

	bmi.BmiHeader.SetBiSize() // safety

	ret, _, err := syscall.SyscallN(proc.GetDIBits.Addr(),
		uintptr(hdc), uintptr(hbm), uintptr(firstScanLine), uintptr(numScanLines),
		uintptr(unsafe.Pointer(dataBufPtr)), uintptr(unsafe.Pointer(bmi)),
		uintptr(usage))

	if wErr := errco.ERROR(ret); wErr == errco.INVALID_PARAMETER {
		panic(wErr)
	} else if ret == 0 {
		panic(errco.ERROR(err))
	}

	return int(ret)
}

// [GetPolyFillMode] function.
//
// [GetPolyFillMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getpolyfillmode
func (hdc HDC) GetPolyFillMode() co.POLYF {
	ret, _, err := syscall.SyscallN(proc.GetPolyFillMode.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.POLYF(ret)
}

// [GetTextExtentPoint32] function.
//
// [GetTextExtentPoint32]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextextentpoint32w
func (hdc HDC) GetTextExtentPoint32(text string) SIZE {
	var sz SIZE
	lpString16 := Str.ToNativeSlice(text)
	ret, _, err := syscall.SyscallN(proc.GetTextExtentPoint32.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&lpString16[0])),
		uintptr(len(lpString16)-1), uintptr(unsafe.Pointer(&sz)))
	runtime.KeepAlive(lpString16)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return sz
}

// [GetTextFace] function.
//
// [GetTextFace]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextfacew
func (hdc HDC) GetTextFace() string {
	var buf [_LF_FACESIZE]uint16
	ret, _, err := syscall.SyscallN(proc.GetTextFace.Addr(),
		uintptr(hdc), uintptr(len(buf)), uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromNativeSlice(buf[:])
}

// [GetTextMetrics] function.
//
// [GetTextMetrics]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextmetricsw
func (hdc HDC) GetTextMetrics(tm *TEXTMETRIC) {
	ret, _, err := syscall.SyscallN(proc.GetTextMetrics.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(tm)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [GetViewportExtEx] function.
//
// [GetViewportExtEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getviewportextex
func (hdc HDC) GetViewportExtEx() SIZE {
	var sz SIZE
	ret, _, err := syscall.SyscallN(proc.GetViewportExtEx.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return sz
}

// [GetViewportOrgEx] function.
//
// [GetViewportOrgEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getviewportorgex
func (hdc HDC) GetViewportOrgEx() POINT {
	var pt POINT
	ret, _, err := syscall.SyscallN(proc.GetViewportOrgEx.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// [GetWindowExtEx] function.
//
// [GetWindowExtEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getwindowextex
func (hdc HDC) GetWindowExtEx() SIZE {
	var sz SIZE
	ret, _, err := syscall.SyscallN(proc.GetWindowExtEx.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return sz
}

// [GetWindowOrgEx] function.
//
// [GetWindowOrgEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getwindoworgex
func (hdc HDC) GetWindowOrgEx() POINT {
	var pt POINT
	ret, _, err := syscall.SyscallN(proc.GetWindowOrgEx.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// [IntersectClipRect] function.
//
// [IntersectClipRect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-intersectcliprect
func (hdc HDC) IntersectClipRect(coords RECT) co.REGION {
	ret, _, err := syscall.SyscallN(proc.IntersectClipRect.Addr(),
		uintptr(hdc), uintptr(coords.Left), uintptr(coords.Top),
		uintptr(coords.Right), uintptr((coords.Bottom)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}

// [InvertRgn] function.
//
// [InvertRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-invertrgn
func (hdc HDC) InvertRgn(hRgn HRGN) {
	ret, _, err := syscall.SyscallN(proc.InvertRgn.Addr(),
		uintptr(hdc), uintptr(hRgn))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [LineTo] function.
//
// [LineTo]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lineto
func (hdc HDC) LineTo(x, y int32) {
	ret, _, err := syscall.SyscallN(proc.LineTo.Addr(),
		uintptr(hdc), uintptr(x), uintptr(y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [LPtoDP] function.
//
// [LPtoDP]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lptodp
func (hdc HDC) LPtoDP(pts []POINT) {
	ret, _, err := syscall.SyscallN(proc.LPtoDP.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
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
	rop co.ROP) {

	ret, _, err := syscall.SyscallN(proc.MaskBlt.Addr(),
		uintptr(hdc), uintptr(destTopLeft.X), uintptr(destTopLeft.Y),
		uintptr(sz.Cx), uintptr(sz.Cy),
		uintptr(hdcSrc), uintptr(srcTopLeft.X), uintptr(srcTopLeft.Y),
		uintptr(hbmMask), uintptr(maskOffset.X), uintptr(maskOffset.Y),
		uintptr(rop))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [MoveToEx] function.
//
// [MoveToEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-movetoex
func (hdc HDC) MoveToEx(x, y int32, pt *POINT) {
	ret, _, err := syscall.SyscallN(proc.MoveToEx.Addr(),
		uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [PaintRgn] function.
//
// [PaintRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-paintrgn
func (hdc HDC) PaintRgn(hRgn HRGN) {
	ret, _, err := syscall.SyscallN(proc.PaintRgn.Addr(),
		uintptr(hdc), uintptr(hRgn))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [PathToRegion] function.
//
// ⚠️ You must defer HRGN.DeleteObject().
//
// [PathToRegion]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-pathtoregion
func (hdc HDC) PathToRegion() HRGN {
	ret, _, err := syscall.SyscallN(proc.PathToRegion.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// [Pie] function.
//
// [Pie]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-pie
func (hdc HDC) Pie(bound RECT, endPointRadial1, endPointRadial2 POINT) {
	ret, _, err := syscall.SyscallN(proc.Pie.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(endPointRadial1.X), uintptr(endPointRadial1.Y),
		uintptr(endPointRadial2.X), uintptr(endPointRadial2.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [PolyDraw] function.
//
// [PolyDraw]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polydraw
func (hdc HDC) PolyDraw(pts []POINT, usage []co.PT) {
	if len(pts) != len(usage) {
		panic(fmt.Sprintf("PolyDraw different slice sizes: %d, %d.",
			len(pts), len(usage)))
	}
	ret, _, err := syscall.SyscallN(proc.PolyDraw.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])),
		uintptr(unsafe.Pointer(&usage[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [Polygon] function.
//
// [Polygon]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polygon
func (hdc HDC) Polygon(pts []POINT) {
	ret, _, err := syscall.SyscallN(proc.Polygon.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [Polyline] function.
//
// [Polyline]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polyline
func (hdc HDC) Polyline(pts []POINT) {
	ret, _, err := syscall.SyscallN(proc.Polyline.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [PolylineTo] function.
//
// [PolylineTo]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polylineto
func (hdc HDC) PolylineTo(pts []POINT) {
	ret, _, err := syscall.SyscallN(proc.PolylineTo.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [PolyPolygon] function.
//
// [PolyPolygon]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polypolygon
func (hdc HDC) PolyPolygon(pts [][]POINT) {
	totalPoints := 0
	for _, block := range pts {
		totalPoints += len(block)
	}

	flat := make([]POINT, 0, totalPoints)    // flat slice of all points
	blockCount := make([]int32, 0, len(pts)) // lengths of each block of points
	for _, block := range pts {
		flat = append(flat, block...)
		blockCount = append(blockCount, int32(len(block)))
	}

	ret, _, err := syscall.SyscallN(proc.PolyPolygon.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&flat[0])),
		uintptr(unsafe.Pointer(&blockCount[0])), uintptr(len(blockCount)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [PolyPolyline] function.
//
// [PolyPolyline]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polypolyline
func (hdc HDC) PolyPolyline(pts [][]POINT) {
	totalPoints := 0
	for _, block := range pts {
		totalPoints += len(block)
	}

	flat := make([]POINT, 0, totalPoints)     // flat slice of all points
	blockCount := make([]uint32, 0, len(pts)) // lengths of each block of points
	for _, block := range pts {
		flat = append(flat, block...)
		blockCount = append(blockCount, uint32(len(block)))
	}

	ret, _, err := syscall.SyscallN(proc.PolyPolyline.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&flat[0])),
		uintptr(unsafe.Pointer(&blockCount[0])), uintptr(len(blockCount)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [PtVisible] function.
//
// [PtVisible]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ptvisible
func (hdc HDC) PtVisible(x, y int32) bool {
	ret, _, err := syscall.SyscallN(proc.PtVisible.Addr(),
		uintptr(hdc), uintptr(x), uintptr(y))
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return ret != 0
}

// [Rectangle] function.
//
// [Rectangle]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rectangle
func (hdc HDC) Rectangle(bound RECT) {
	ret, _, err := syscall.SyscallN(proc.Rectangle.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [RestoreDC] function.
//
// Used together with HDC.SaveDC().
//
// [RestoreDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-restoredc
func (hdc HDC) RestoreDC(savedDC int32) {
	ret, _, err := syscall.SyscallN(proc.RestoreDC.Addr(),
		uintptr(hdc), uintptr(savedDC))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [RoundRect] function.
//
// [RoundRect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-roundrect
func (hdc HDC) RoundRect(bound RECT, sz SIZE) {
	ret, _, err := syscall.SyscallN(proc.RoundRect.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(sz.Cx), uintptr(sz.Cy))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SaveDC] function.
//
// Used together with HDC.RestoreDC().
//
// [SaveDC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-savedc
func (hdc HDC) SaveDC() int32 {
	ret, _, err := syscall.SyscallN(proc.SaveDC.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return int32(ret)
}

// [SelectClipPath] function.
//
// [SelectClipPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectclippath
func (hdc HDC) SelectClipPath(mode co.RGN) {
	ret, _, err := syscall.SyscallN(proc.SelectClipPath.Addr(),
		uintptr(hdc), uintptr(mode))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SelectClipRgn] function.
//
// [SelectClipRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectcliprgn
func (hdc HDC) SelectClipRgn(hRgn HRGN) co.REGION {
	ret, _, err := syscall.SyscallN(proc.SelectClipRgn.Addr(),
		uintptr(hdc), uintptr(hRgn))
	if ret == _REGION_ERROR {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}

// [SelectObjectBitmap] function.
//
// [SelectObjectBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBitmap(hBmp HBITMAP) HBITMAP {
	ret, _, err := syscall.SyscallN(proc.SelectObject.Addr(),
		uintptr(hdc), uintptr(hBmp))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret)
}

// [SelectObjectBrush] function.
//
// [SelectObjectBrush]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBrush(hBrush HBRUSH) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.SelectObject.Addr(),
		uintptr(hdc), uintptr(hBrush))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// [SelectObjectFont] function.
//
// [SelectObjectFont]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectFont(hFont HFONT) HFONT {
	ret, _, err := syscall.SyscallN(proc.SelectObject.Addr(),
		uintptr(hdc), uintptr(hFont))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HFONT(ret)
}

// [SelectObjectPen] function.
//
// [SelectObjectPen]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectPen(hPen HPEN) HPEN {
	ret, _, err := syscall.SyscallN(proc.SelectObject.Addr(),
		uintptr(hdc), uintptr(hPen))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HPEN(ret)
}

// [SelectObjectRgn] function.
//
// [SelectObjectRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectRgn(hRgn HRGN) co.REGION {
	ret, _, err := syscall.SyscallN(proc.SelectObject.Addr(),
		uintptr(hdc), uintptr(hRgn))
	if ret == _HGDI_ERROR || ret == _REGION_ERROR {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}

// [SetArcDirection] function.
//
// [SetArcDirection]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setarcdirection
func (hdc HDC) SetArcDirection(direction co.AD) co.AD {
	ret, _, err := syscall.SyscallN(proc.SetArcDirection.Addr(),
		uintptr(hdc), uintptr(direction))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.AD(ret)
}

// [SetBkColor] function.
//
// [SetBkColor]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkcolor
func (hdc HDC) SetBkColor(color COLORREF) COLORREF {
	ret, _, err := syscall.SyscallN(proc.SetBkColor.Addr(),
		uintptr(hdc), uintptr(color))
	if ret == _CLR_INVALID {
		panic(errco.ERROR(err))
	}
	return COLORREF(ret)
}

// [SetBkMode] function.
//
// [SetBkMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
func (hdc HDC) SetBkMode(mode co.BKMODE) co.BKMODE {
	ret, _, err := syscall.SyscallN(proc.SetBkMode.Addr(),
		uintptr(hdc), uintptr(mode))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.BKMODE(ret)
}

// [SetPolyFillMode] function.
//
// [SetPolyFillMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpolyfillmode
func (hdc HDC) SetPolyFillMode(mode co.POLYF) co.POLYF {
	ret, _, err := syscall.SyscallN(proc.SetPolyFillMode.Addr(),
		uintptr(hdc), uintptr(mode))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.POLYF(ret)
}

// [SetStretchBltMode] function.
//
// [SetStretchBltMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setstretchbltmode
func (hdc HDC) SetStretchBltMode(mode co.STRETCH) co.STRETCH {
	ret, _, err := syscall.SyscallN(proc.SetStretchBltMode.Addr(),
		uintptr(hdc), uintptr(mode))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.STRETCH(ret)
}

// [SetTextAlign] function.
//
// [SetTextAlign]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-settextalign
func (hdc HDC) SetTextAlign(align co.TA) {
	ret, _, err := syscall.SyscallN(proc.SetTextAlign.Addr(),
		uintptr(hdc), uintptr(align))
	if ret == _GDI_ERR {
		panic(errco.ERROR(err))
	}
}

// [StretchBlt] function.
//
// This method is called from the destination HDC.
//
// [StretchBlt]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-stretchblt
func (hdc HDC) StretchBlt(
	destTopLeft POINT, destSz SIZE,
	hdcSrc HDC, srcTopLeft POINT, srcSz SIZE, rop co.ROP) {

	ret, _, err := syscall.SyscallN(proc.StretchBlt.Addr(),
		uintptr(hdc), uintptr(destTopLeft.X), uintptr(destTopLeft.Y),
		uintptr(destSz.Cx), uintptr(destSz.Cy),
		uintptr(hdcSrc), uintptr(srcTopLeft.X), uintptr(srcTopLeft.Y),
		uintptr(srcSz.Cx), uintptr(srcSz.Cy), uintptr(rop))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [StrokeAndFillPath] function.
//
// [StrokeAndFillPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-strokeandfillpath
func (hdc HDC) StrokeAndFillPath() {
	ret, _, err := syscall.SyscallN(proc.StrokeAndFillPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [StrokePath] function.
//
// [StrokePath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-strokepath
func (hdc HDC) StrokePath() {
	ret, _, err := syscall.SyscallN(proc.StrokePath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [TextOut] function.
//
// [TextOut]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-textoutw
func (hdc HDC) TextOut(x, y int32, text string) {
	lpString16 := Str.ToNativeSlice(text)
	ret, _, err := syscall.SyscallN(proc.TextOut.Addr(),
		uintptr(hdc), uintptr(x), uintptr(y),
		uintptr(unsafe.Pointer(&lpString16[0])),
		uintptr(len(lpString16)-1))
	runtime.KeepAlive(lpString16)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
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
	colorTransparent COLORREF) {

	ret, _, err := syscall.SyscallN(proc.TransparentBlt.Addr(),
		uintptr(hdc),
		uintptr(destTopLeft.X), uintptr(destTopLeft.Y),
		uintptr(destSz.Cx), uintptr(destSz.Cy),
		uintptr(hdcSrc),
		uintptr(srcTopLeft.X), uintptr(srcTopLeft.Y),
		uintptr(srcSz.Cx), uintptr(srcSz.Cy),
		uintptr(colorTransparent))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [WidenPath] function.
//
// [WidenPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-widenpath
func (hdc HDC) WidenPath() {
	ret, _, err := syscall.SyscallN(proc.WidenPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
