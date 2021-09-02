package win

import (
	"fmt"
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
func (hdc HDC) AngleArc(x, y int32, r uint32, startAngle, sweepAngle float32) {
	ret, _, err := syscall.Syscall6(proc.AngleArc.Addr(), 6,
		uintptr(hdc), uintptr(x), uintptr(y), uintptr(r),
		uintptr(startAngle), uintptr(sweepAngle))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arc
func (hdc HDC) Arc(left, top, right, bottom, xr1, yr1, xr2, yr2 int32) {
	ret, _, err := syscall.Syscall9(proc.Arc.Addr(), 9,
		uintptr(hdc), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom),
		uintptr(xr1), uintptr(yr1), uintptr(xr2), uintptr(yr2))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arcto
func (hdc HDC) ArcTo(left, top, right, bottom, xr1, yr1, xr2, yr2 int32) {
	ret, _, err := syscall.Syscall9(proc.ArcTo.Addr(), 9,
		uintptr(hdc), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom),
		uintptr(xr1), uintptr(yr1), uintptr(xr2), uintptr(yr2))
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
func (hdc HDC) BitBlt(x, y, cx, cy, hdcSrc HDC, x1, y1 int32, rop co.ROP) {
	ret, _, err := syscall.Syscall9(proc.BitBlt.Addr(), 9,
		uintptr(hdc), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy),
		uintptr(hdcSrc), uintptr(x1), uintptr(y1), uintptr(rop))
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
func (hdc HDC) Chord(x1, y1, x2, y2, x3, y3, x4, y4 int32) {
	ret, _, err := syscall.Syscall9(proc.Chord.Addr(), 9,
		uintptr(hdc),
		uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2),
		uintptr(x3), uintptr(y3), uintptr(x4), uintptr(y4))
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
func (hdc HDC) Ellipse(left, top, right, bottom int32) {
	ret, _, err := syscall.Syscall6(proc.Ellipse.Addr(), 5,
		uintptr(hdc), uintptr(left), uintptr(top),
		uintptr(right), uintptr(bottom), 0)
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
	lprcClip *RECT,
	lpfnEnum func(hMon HMONITOR, hdcMon HDC, rcMon *RECT, lp LPARAM) bool,
	dwData LPARAM) {

	ret, _, err := syscall.Syscall6(proc.EnumDisplayMonitors.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(lprcClip)),
		syscall.NewCallback(
			func(hMon HMONITOR, hdcMon HDC, rcMon *RECT, lp LPARAM) uintptr {
				return util.BoolToUintptr(lpfnEnum(hMon, hdcMon, rcMon, lp))
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
func (hdc HDC) FillRect(lprc *RECT, hbr HBRUSH) {
	ret, _, err := syscall.Syscall(proc.FillRect.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(lprc)), uintptr(hbr))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillrgn
func (hdc HDC) FillRgn(hrgn HRGN, hbr HBRUSH) {
	ret, _, err := syscall.Syscall(proc.FillRgn.Addr(), 3,
		uintptr(hdc), uintptr(hrgn), uintptr(hbr))
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
func (hdc HDC) FrameRect(lprc *RECT, hbr HBRUSH) {
	ret, _, err := syscall.Syscall(proc.FrameRect.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(lprc)), uintptr(hbr))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-framergn
func (hdc HDC) FrameRgn(hrgn HRGN, hbr HBRUSH, w, h int32) {
	ret, _, err := syscall.Syscall6(proc.FrameRgn.Addr(), 5,
		uintptr(hdc), uintptr(hrgn), uintptr(hbr), uintptr(w), uintptr(h),
		0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
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
func (hdc HDC) GetTextExtentPoint32(lpString string) SIZE {
	sz := SIZE{}
	lpString16 := Str.ToUint16Slice(lpString)
	ret, _, err := syscall.Syscall6(proc.GetTextExtentPoint32.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(&lpString16[0])),
		uintptr(len(lpString16)-1), uintptr(unsafe.Pointer(&sz)), 0, 0)
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
func (hdc HDC) GetTextMetrics(lptm *TEXTMETRIC) {
	ret, _, err := syscall.Syscall(proc.GetTextMetrics.Addr(), 2,
		uintptr(hdc), uintptr(unsafe.Pointer(lptm)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invertrect
func (hdc HDC) InvertRect(lprc RECT) {
	ret, _, err := syscall.Syscall(proc.InvertRect.Addr(), 2,
		uintptr(hdc), uintptr(unsafe.Pointer(&lprc)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-invertrgn
func (hdc HDC) InvertRgn(hrgn HRGN) {
	ret, _, err := syscall.Syscall(proc.InvertRgn.Addr(), 2,
		uintptr(hdc), uintptr(hrgn), 0)
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
func (hdc HDC) LPtoDP(lppt []POINT) {
	ret, _, err := syscall.Syscall(proc.LPtoDP.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&lppt[0])), uintptr(len(lppt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-movetoex
func (hdc HDC) MoveToEx(x, y int32, lppt *POINT) {
	ret, _, err := syscall.Syscall6(proc.MoveToEx.Addr(), 4,
		uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(lppt)),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-paintrgn
func (hdc HDC) PaintRgn(hrgn HRGN) {
	ret, _, err := syscall.Syscall(proc.PaintRgn.Addr(), 2,
		uintptr(hdc), uintptr(hrgn), 0)
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
func (hdc HDC) Pie(left, top, right, bottom, xr1, yr1, xr2, yr2 int32) {
	ret, _, err := syscall.Syscall9(proc.Pie.Addr(), 9,
		uintptr(hdc), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom),
		uintptr(xr1), uintptr(yr1), uintptr(xr2), uintptr(yr2))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polydraw
func (hdc HDC) PolyDraw(apt []POINT, aj []co.PT) {
	if len(apt) != len(aj) {
		panic(fmt.Sprintf("PolyDraw different slice sizes: %d, %d.",
			len(apt), len(aj)))
	}
	ret, _, err := syscall.Syscall6(proc.PolyDraw.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])),
		uintptr(unsafe.Pointer(&aj[0])), uintptr(len(apt)),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polygon
func (hdc HDC) Polygon(apt []POINT) {
	ret, _, err := syscall.Syscall(proc.Polygon.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])), uintptr(len(apt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polyline
func (hdc HDC) Polyline(apt []POINT) {
	ret, _, err := syscall.Syscall(proc.Polyline.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])), uintptr(len(apt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polylineto
func (hdc HDC) PolylineTo(apt []POINT) {
	ret, _, err := syscall.Syscall(proc.PolylineTo.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])), uintptr(len(apt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polypolygon
func (hdc HDC) PolyPolygon(apt [][]POINT) {
	totalPoints := 0
	for _, block := range apt {
		totalPoints += len(block)
	}

	flat := make([]POINT, 0, totalPoints)    // flat slice of all points
	blockCount := make([]int32, 0, len(apt)) // lengths of each block of points
	for _, block := range apt {
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
func (hdc HDC) PolyPolyline(apt [][]POINT) {
	totalPoints := 0
	for _, block := range apt {
		totalPoints += len(block)
	}

	flat := make([]POINT, 0, totalPoints)     // flat slice of all points
	blockCount := make([]uint32, 0, len(apt)) // lengths of each block of points
	for _, block := range apt {
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
func (hdc HDC) Rectangle(left, top, right, bottom int32) {
	ret, _, err := syscall.Syscall6(proc.Rectangle.Addr(), 5,
		uintptr(hdc), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom),
		0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-restoredc
func (hdc HDC) RestoreDC(nSavedDC int32) {
	ret, _, err := syscall.Syscall(proc.RestoreDC.Addr(), 2,
		uintptr(hdc), uintptr(nSavedDC), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-roundrect
func (hdc HDC) RoundRect(left, top, right, bottom, width, height int32) {
	ret, _, err := syscall.Syscall9(proc.RoundRect.Addr(), 7,
		uintptr(hdc), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom),
		uintptr(width), uintptr(height), 0, 0)
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
func (hdc HDC) SelectObjectBitmap(b HBITMAP) HBITMAP {
	ret, _, err := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(b), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBrush(b HBRUSH) HBRUSH {
	ret, _, err := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(b), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectFont(f HFONT) HFONT {
	ret, _, err := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(f), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HFONT(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectPen(p HPEN) HPEN {
	ret, _, err := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(p), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HPEN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectRgn(r HRGN) co.REGION {
	ret, _, err := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(r), 0)
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
func (hdc HDC) SetPolyFillMode(iMode co.POLYF) co.POLYF {
	ret, _, err := syscall.Syscall(proc.SetPolyFillMode.Addr(), 2,
		uintptr(hdc), uintptr(iMode), 0)
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
func (hdc HDC) TextOut(x, y int32, lpString string) {
	lpString16 := Str.ToUint16Slice(lpString)
	ret, _, err := syscall.Syscall6(proc.TextOut.Addr(), 5,
		uintptr(hdc), uintptr(x), uintptr(y),
		uintptr(unsafe.Pointer(&lpString16[0])),
		uintptr(len(lpString16)-1), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// This method is called from the destination HDC.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-transparentblt
func (hdc HDC) TransparentBlt(
	xoriginDest, yoriginDest, wDest, hDest,
	hdcSrc HDC, xoriginSrc, yoriginSrc, wSrc, hSrc int32,
	crTransparent COLORREF) {

	ret, _, err := syscall.Syscall12(proc.TransparentBlt.Addr(), 11,
		uintptr(hdc), uintptr(xoriginDest), uintptr(yoriginDest), uintptr(wDest), uintptr(hDest),
		uintptr(hdcSrc), uintptr(xoriginSrc), uintptr(yoriginSrc), uintptr(wSrc), uintptr(hSrc),
		uintptr(crTransparent), 0)
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
