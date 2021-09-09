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

// A handle to a device context (DC).
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdc
type HDC HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-abortpath
func (hdc HDC) AbortPath() {
	ret, _, err := syscall.Syscall(proc.AbortPath.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-anglearc
func (hdc HDC) AngleArc(center POINT, r uint32, startAngle, sweepAngle float32) {
	ret, _, err := syscall.Syscall6(proc.AngleArc.Addr(), 6,
		uintptr(hdc), uintptr(center.X), uintptr(center.Y), uintptr(r),
		uintptr(startAngle), uintptr(sweepAngle))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arc
func (hdc HDC) Arc(bound RECT, radialStart, radialEnd POINT) {
	ret, _, err := syscall.Syscall9(proc.Arc.Addr(), 9,
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(radialStart.X), uintptr(radialStart.Y),
		uintptr(radialEnd.X), uintptr(radialEnd.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arcto
func (hdc HDC) ArcTo(bound RECT, radialStart, radialEnd POINT) {
	ret, _, err := syscall.Syscall9(proc.ArcTo.Addr(), 9,
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(radialStart.X), uintptr(radialStart.Y),
		uintptr(radialEnd.X), uintptr(radialEnd.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HDC.EndPath().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-beginpath
func (hdc HDC) BeginPath() {
	ret, _, err := syscall.Syscall(proc.BeginPath.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// This method is called from the destination HDC.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-bitblt
func (hdc HDC) BitBlt(
	destTopLeft POINT, sz SIZE, hdcSrc HDC, srcTopLeft POINT, rop co.ROP) {

	ret, _, err := syscall.Syscall9(proc.BitBlt.Addr(), 9,
		uintptr(hdc), uintptr(destTopLeft.X), uintptr(destTopLeft.Y),
		uintptr(sz.Cx), uintptr(sz.Cy),
		uintptr(hdcSrc), uintptr(srcTopLeft.X), uintptr(srcTopLeft.Y),
		uintptr(rop))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-canceldc
func (hdc HDC) CancelDC() {
	ret, _, err := syscall.Syscall(proc.CancelDC.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-chord
func (hdc HDC) Chord(bound RECT, radialStart, radialEnd POINT) {
	ret, _, err := syscall.Syscall9(proc.Chord.Addr(), 9,
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(radialStart.X), uintptr(radialStart.Y),
		uintptr(radialEnd.X), uintptr(radialEnd.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-closefigure
func (hdc HDC) CloseFigure() {
	ret, _, err := syscall.Syscall(proc.CloseFigure.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HBITMAP.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatiblebitmap
func (hdc HDC) CreateCompatibleBitmap(cx, cy int32) HBITMAP {
	ret, _, err := syscall.Syscall(proc.CreateCompatibleBitmap.Addr(), 3,
		uintptr(hdc), uintptr(cx), uintptr(cy))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret)
}

// ⚠️ You must defer HDC.DeleteDC().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatibledc
func (hdc HDC) CreateCompatibleDC() HDC {
	ret, _, err := syscall.Syscall(proc.CreateCompatibleDC.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HDC(ret)
}

// ⚠️ You must defer HBITMAP.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdibsection
func (hdc HDC) CreateDIBSection(
	bmi *BITMAPINFO, usage co.DIB,
	hSection HFILEMAP, offset uint32) (HBITMAP, *byte) {

	var ppvBits *byte
	ret, _, err := syscall.Syscall6(proc.CreateDIBSection.Addr(), 6,
		uintptr(hdc), uintptr(unsafe.Pointer(bmi)), uintptr(usage),
		uintptr(unsafe.Pointer(&ppvBits)), uintptr(hSection), uintptr(offset))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret), ppvBits
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deletedc
func (hdc HDC) DeleteDC() {
	ret, _, err := syscall.Syscall(proc.DeleteDC.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawicon
func (hdc HDC) DrawIcon(x, y int32, hIcon HICON) {
	ret, _, err := syscall.Syscall6(proc.DrawIcon.Addr(), 4,
		uintptr(hdc), uintptr(x), uintptr(y), uintptr(hIcon), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ellipse
func (hdc HDC) Ellipse(bound RECT) {
	ret, _, err := syscall.Syscall6(proc.Ellipse.Addr(), 5,
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-endpath
func (hdc HDC) EndPath() {
	ret, _, err := syscall.Syscall(proc.EndPath.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func (hdc HDC) EnumDisplayMonitors(
	rcClip *RECT,
	enumFunc func(hMon HMONITOR, hdcMon HDC, rcMon *RECT) bool) {

	ret, _, err := syscall.Syscall6(proc.EnumDisplayMonitors.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(rcClip)),
		syscall.NewCallback(
			func(hMon HMONITOR, hdcMon HDC, rcMon *RECT, _ LPARAM) uintptr {
				return util.BoolToUintptr(enumFunc(hMon, hdcMon, rcMon))
			}),
		0, 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillpath
func (hdc HDC) FillPath() {
	ret, _, err := syscall.Syscall(proc.FillPath.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-fillrect
func (hdc HDC) FillRect(rc *RECT, hBrush HBRUSH) {
	ret, _, err := syscall.Syscall(proc.FillRect.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(rc)), uintptr(hBrush))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillrgn
func (hdc HDC) FillRgn(hRgn HRGN, hBrush HBRUSH) {
	ret, _, err := syscall.Syscall(proc.FillRgn.Addr(), 3,
		uintptr(hdc), uintptr(hRgn), uintptr(hBrush))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-flattenpath
func (hdc HDC) FlattenPath() {
	ret, _, err := syscall.Syscall(proc.FlattenPath.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-framerect
func (hdc HDC) FrameRect(rc *RECT, hBrush HBRUSH) {
	ret, _, err := syscall.Syscall(proc.FrameRect.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(rc)), uintptr(hBrush))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-framergn
func (hdc HDC) FrameRgn(hRgn HRGN, hBrush HBRUSH, w, h int32) {
	ret, _, err := syscall.Syscall6(proc.FrameRgn.Addr(), 5,
		uintptr(hdc), uintptr(hRgn), uintptr(hBrush), uintptr(w), uintptr(h),
		0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getcurrentpositionex
func (hdc HDC) GetCurrentPositionEx() POINT {
	pt := POINT{}
	ret, _, err := syscall.Syscall(proc.GetCurrentPositionEx.Addr(), 2,
		uintptr(hdc), uintptr(unsafe.Pointer(&pt)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdcbrushcolor
func (hdc HDC) GetDCBrushColor() COLORREF {
	ret, _, err := syscall.Syscall(proc.GetDCBrushColor.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == _CLR_INVALID {
		panic(errco.ERROR(err))
	}
	return COLORREF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdcpencolor
func (hdc HDC) GetDCPenColor() COLORREF {
	ret, _, err := syscall.Syscall(proc.GetDCPenColor.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == _CLR_INVALID {
		panic(errco.ERROR(err))
	}
	return COLORREF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdevicecaps
func (hdc HDC) GetDeviceCaps(index co.GDC) int32 {
	ret, _, _ := syscall.Syscall(proc.GetDeviceCaps.Addr(), 2,
		uintptr(hdc), uintptr(index), 0)
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getpolyfillmode
func (hdc HDC) GetPolyFillMode() co.POLYF {
	ret, _, err := syscall.Syscall(proc.GetPolyFillMode.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.POLYF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextextentpoint32w
func (hdc HDC) GetTextExtentPoint32(text string) SIZE {
	sz := SIZE{}
	lpString16 := Str.ToUint16Slice(text)
	ret, _, err := syscall.Syscall6(proc.GetTextExtentPoint32.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(&lpString16[0])),
		uintptr(len(lpString16)-1), uintptr(unsafe.Pointer(&sz)), 0, 0)
	runtime.KeepAlive(lpString16)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return sz
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextfacew
func (hdc HDC) GetTextFace() string {
	buf := [_LF_FACESIZE]uint16{}
	ret, _, err := syscall.Syscall(proc.GetTextFace.Addr(), 3,
		uintptr(hdc), uintptr(len(buf)), uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromUint16Slice(buf[:])
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextmetricsw
func (hdc HDC) GetTextMetrics(tm *TEXTMETRIC) {
	ret, _, err := syscall.Syscall(proc.GetTextMetrics.Addr(), 2,
		uintptr(hdc), uintptr(unsafe.Pointer(tm)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invertrect
func (hdc HDC) InvertRect(rc *RECT) {
	ret, _, err := syscall.Syscall(proc.InvertRect.Addr(), 2,
		uintptr(hdc), uintptr(unsafe.Pointer(rc)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-invertrgn
func (hdc HDC) InvertRgn(hRgn HRGN) {
	ret, _, err := syscall.Syscall(proc.InvertRgn.Addr(), 2,
		uintptr(hdc), uintptr(hRgn), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lineto
func (hdc HDC) LineTo(x, y int32) {
	ret, _, err := syscall.Syscall(proc.LineTo.Addr(), 3,
		uintptr(hdc), uintptr(x), uintptr(y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lptodp
func (hdc HDC) LPtoDP(pts []POINT) {
	ret, _, err := syscall.Syscall(proc.LPtoDP.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-movetoex
func (hdc HDC) MoveToEx(x, y int32, pt *POINT) {
	ret, _, err := syscall.Syscall6(proc.MoveToEx.Addr(), 4,
		uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(pt)),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-paintrgn
func (hdc HDC) PaintRgn(hRgn HRGN) {
	ret, _, err := syscall.Syscall(proc.PaintRgn.Addr(), 2,
		uintptr(hdc), uintptr(hRgn), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HRGN.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-pathtoregion
func (hdc HDC) PathToRegion() HRGN {
	ret, _, err := syscall.Syscall(proc.PathToRegion.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-pie
func (hdc HDC) Pie(bound RECT, endPointRadial1, endPointRadial2 POINT) {
	ret, _, err := syscall.Syscall9(proc.Pie.Addr(), 9,
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(endPointRadial1.X), uintptr(endPointRadial1.Y),
		uintptr(endPointRadial2.X), uintptr(endPointRadial2.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polydraw
func (hdc HDC) PolyDraw(pts []POINT, usage []co.PT) {
	if len(pts) != len(usage) {
		panic(fmt.Sprintf("PolyDraw different slice sizes: %d, %d.",
			len(pts), len(usage)))
	}
	ret, _, err := syscall.Syscall6(proc.PolyDraw.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])),
		uintptr(unsafe.Pointer(&usage[0])), uintptr(len(pts)),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polygon
func (hdc HDC) Polygon(pts []POINT) {
	ret, _, err := syscall.Syscall(proc.Polygon.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polyline
func (hdc HDC) Polyline(pts []POINT) {
	ret, _, err := syscall.Syscall(proc.Polyline.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polylineto
func (hdc HDC) PolylineTo(pts []POINT) {
	ret, _, err := syscall.Syscall(proc.PolylineTo.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polypolygon
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

	ret, _, err := syscall.Syscall6(proc.PolyPolygon.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(&flat[0])),
		uintptr(unsafe.Pointer(&blockCount[0])), uintptr(len(blockCount)),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polypolyline
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

	ret, _, err := syscall.Syscall6(proc.PolyPolyline.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(&flat[0])),
		uintptr(unsafe.Pointer(&blockCount[0])), uintptr(len(blockCount)),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ptvisible
func (hdc HDC) PtVisible(x, y int32) bool {
	ret, _, err := syscall.Syscall(proc.PtVisible.Addr(), 3,
		uintptr(hdc), uintptr(x), uintptr(y))
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rectangle
func (hdc HDC) Rectangle(bound RECT) {
	ret, _, err := syscall.Syscall6(proc.Rectangle.Addr(), 5,
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-restoredc
func (hdc HDC) RestoreDC(savedDC int32) {
	ret, _, err := syscall.Syscall(proc.RestoreDC.Addr(), 2,
		uintptr(hdc), uintptr(savedDC), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-roundrect
func (hdc HDC) RoundRect(bound RECT, sz SIZE) {
	ret, _, err := syscall.Syscall9(proc.RoundRect.Addr(), 7,
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(sz.Cx), uintptr(sz.Cy), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-savedc
func (hdc HDC) SaveDC() int32 {
	ret, _, err := syscall.Syscall(proc.SaveDC.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBitmap(hBmp HBITMAP) HBITMAP {
	ret, _, err := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(hBmp), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBrush(hBrush HBRUSH) HBRUSH {
	ret, _, err := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(hBrush), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectFont(hFont HFONT) HFONT {
	ret, _, err := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(hFont), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HFONT(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectPen(hPen HPEN) HPEN {
	ret, _, err := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(hPen), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HPEN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectRgn(hRgn HRGN) co.REGION {
	ret, _, err := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(hRgn), 0)
	if ret == _HGDI_ERROR {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkcolor
func (hdc HDC) SetBkColor(color COLORREF) COLORREF {
	ret, _, err := syscall.Syscall(proc.SetBkColor.Addr(), 2,
		uintptr(hdc), uintptr(color), 0)
	if ret == _CLR_INVALID {
		panic(errco.ERROR(err))
	}
	return COLORREF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
func (hdc HDC) SetBkMode(mode co.BKMODE) co.BKMODE {
	ret, _, err := syscall.Syscall(proc.SetBkMode.Addr(), 2,
		uintptr(hdc), uintptr(mode), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.BKMODE(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpolyfillmode
func (hdc HDC) SetPolyFillMode(mode co.POLYF) co.POLYF {
	ret, _, err := syscall.Syscall(proc.SetPolyFillMode.Addr(), 2,
		uintptr(hdc), uintptr(mode), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.POLYF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-settextalign
func (hdc HDC) SetTextAlign(align co.TA) {
	ret, _, err := syscall.Syscall(proc.SetTextAlign.Addr(), 2,
		uintptr(hdc), uintptr(align), 0)
	if ret == _GDI_ERR {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-strokeandfillpath
func (hdc HDC) StrokeAndFillPath() {
	ret, _, err := syscall.Syscall(proc.StrokeAndFillPath.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-strokepath
func (hdc HDC) StrokePath() {
	ret, _, err := syscall.Syscall(proc.StrokePath.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-textoutw
func (hdc HDC) TextOut(x, y int32, text string) {
	lpString16 := Str.ToUint16Slice(text)
	ret, _, err := syscall.Syscall6(proc.TextOut.Addr(), 5,
		uintptr(hdc), uintptr(x), uintptr(y),
		uintptr(unsafe.Pointer(&lpString16[0])),
		uintptr(len(lpString16)-1), 0)
	runtime.KeepAlive(lpString16)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// This method is called from the destination HDC.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-transparentblt
func (hdc HDC) TransparentBlt(
	destTopLeft POINT, destSz SIZE,
	hdcSrc HDC, srcTopLeft POINT, srcSz SIZE,
	colorTransparent COLORREF) {

	ret, _, err := syscall.Syscall12(proc.TransparentBlt.Addr(), 11,
		uintptr(hdc),
		uintptr(destTopLeft.X), uintptr(destTopLeft.Y),
		uintptr(destSz.Cx), uintptr(destSz.Cy),
		uintptr(hdcSrc),
		uintptr(srcTopLeft.X), uintptr(srcTopLeft.Y),
		uintptr(srcSz.Cx), uintptr(srcSz.Cy),
		uintptr(colorTransparent), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-widenpath
func (hdc HDC) WidenPath() {
	ret, _, err := syscall.Syscall(proc.WidenPath.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
