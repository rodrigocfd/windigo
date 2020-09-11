/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win/proc"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdc
type HDC HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatibledc
func (hdc HDC) CreateCompatibleDC() HDC {
	ret, _, _ := syscall.Syscall(proc.CreateCompatibleDC.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "CreateCompatibleDC").Error())
	}
	return HDC(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func (hdc HDC) EnumDisplayMonitors(
	lprcClip *RECT,
	lpfnEnum func(hMon HMONITOR, hdcMon HDC, rcMon uintptr, lp LPARAM) bool,
	dwData LPARAM) *WinError {

	ret, _, _ := syscall.Syscall6(proc.EnumDisplayMonitors.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(lprcClip)),
		syscall.NewCallback(
			func(hMon HMONITOR, hdcMon HDC, rcMon uintptr, lp LPARAM) int32 {
				return _Win.BoolToInt32(lpfnEnum(hMon, hdcMon, rcMon, lp))
			}),
		0, 0, 0)
	if ret == 0 {
		return NewWinError(co.ERROR_E_UNEXPECTED, "EnumDisplayMonitors")
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deletedc
func (hdc HDC) DeleteDC() {
	if hdc != 0 {
		syscall.Syscall(proc.DeleteDC.Addr(), 1,
			uintptr(hdc), 0, 0)
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdevicecaps
func (hdc HDC) GetDeviceCaps(index co.GDC) int32 {
	ret, _, _ := syscall.Syscall(proc.GetDeviceCaps.Addr(), 2,
		uintptr(hdc), uintptr(index), 0)
	return int32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextextentpoint32w
func (hdc HDC) GetTextExtentPoint32(lpString string) *SIZE {
	sz := &SIZE{}
	ret, _, _ := syscall.Syscall6(proc.GetTextExtentPoint32.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(StrToPtr(lpString))),
		uintptr(len(lpString)), uintptr(unsafe.Pointer(sz)), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "GetTextExtentPoint32").Error())
	}
	return sz
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextfacew
func (hdc HDC) GetTextFace() string {
	buf := [32]uint16{} // LF_FACESIZE
	ret, _, _ := syscall.Syscall(proc.GetTextFace.Addr(), 3,
		uintptr(hdc), uintptr(len(buf)), uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "GetTextFace").Error())
	}
	return syscall.UTF16ToString(buf[:])
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lineto
func (hdc HDC) LineTo(x, y int32) {
	ret, _, _ := syscall.Syscall(proc.LineTo.Addr(), 3,
		uintptr(hdc), uintptr(x), uintptr(y))
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "LineTo").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polydraw
func (hdc HDC) PolyDraw(apt []POINT, aj []co.PT) {
	if len(apt) != len(aj) {
		panic(fmt.Sprintf("PolyDraw different slice sizes: %d, %d.",
			len(apt), len(aj)))
	}
	ret, _, _ := syscall.Syscall6(proc.PolyDraw.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])),
		uintptr(unsafe.Pointer(&aj[0])), uintptr(len(apt)),
		0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED,
			fmt.Sprintf("PolyDraw, %d points", len(apt))).Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polygon
func (hdc HDC) Polygon(apt []POINT) {
	ret, _, _ := syscall.Syscall(proc.Polygon.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])), uintptr(len(apt)))
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED,
			fmt.Sprintf("Polygon, %d points", len(apt))).Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polyline
func (hdc HDC) Polyline(apt []POINT) {
	ret, _, _ := syscall.Syscall(proc.Polyline.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])), uintptr(len(apt)))
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED,
			fmt.Sprintf("Polyline, %d points", len(apt))).Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polylineto
func (hdc HDC) PolylineTo(apt []POINT) {
	ret, _, _ := syscall.Syscall(proc.PolylineTo.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])), uintptr(len(apt)))
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED,
			fmt.Sprintf("PolylineTo failed for %d points.", len(apt))).Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-restoredc
func (hdc HDC) RestoreDC(nSavedDC int32) {
	syscall.Syscall(proc.RestoreDC.Addr(), 2,
		uintptr(hdc), uintptr(nSavedDC), 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-savedc
func (hdc HDC) SaveDC() int32 {
	ret, _, _ := syscall.Syscall(proc.SaveDC.Addr(), 1,
		uintptr(hdc), 0, 0)
	return int32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBitmap(b HBITMAP) HBITMAP {
	ret, _, _ := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(b), 0)
	if ret == _HGDI_ERROR {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "SelectObject, HBITMAP").Error())
	}
	return HBITMAP(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBrush(b HBRUSH) HBRUSH {
	ret, _, _ := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(b), 0)
	if ret == _HGDI_ERROR {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "SelectObject, HBRUSH").Error())
	}
	return HBRUSH(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectFont(f HFONT) HFONT {
	ret, _, _ := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(f), 0)
	if ret == _HGDI_ERROR {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "SelectObject, HFONT").Error())
	}
	return HFONT(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectPen(p HPEN) HPEN {
	ret, _, _ := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(p), 0)
	if ret == _HGDI_ERROR {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "SelectObject, HPEN").Error())
	}
	return HPEN(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectRgn(r HRGN) HRGN {
	ret, _, _ := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(r), 0)
	if ret == _HGDI_ERROR {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "SelectObject, HRGN").Error())
	}
	return HRGN(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkcolor
func (hdc HDC) SetBkColor(color COLORREF) COLORREF {
	ret, _, _ := syscall.Syscall(proc.SetBkColor.Addr(), 2,
		uintptr(hdc), uintptr(color), 0)
	if ret == _CLR_INVALID {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "SetBkColor").Error())
	}
	return COLORREF(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
func (hdc HDC) SetBkMode(mode co.BKMODE) co.BKMODE {
	ret, _, _ := syscall.Syscall(proc.SetBkMode.Addr(), 2,
		uintptr(hdc), uintptr(mode), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "SetBkMode").Error())
	}
	return co.BKMODE(ret)
}
