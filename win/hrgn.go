/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"windigo/co"
	proc "windigo/win/internal"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hrgn
type HRGN HGDIOBJ

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createellipticrgn
func CreateEllipticRgn(x1, y1, x2, y2 int32) HRGN {
	ret, _, _ := syscall.Syscall6(proc.CreateEllipticRgn.Addr(), 4,
		uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2),
		0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "CreateEllipticRgn"))
	}
	return HRGN(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createellipticrgnindirect
func CreateEllipticRgnIndirect(lprect *RECT) HRGN {
	ret, _, _ := syscall.Syscall(proc.CreateEllipticRgnIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lprect)), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "CreateEllipticRgnIndirect"))
	}
	return HRGN(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpolygonrgn
func CreatePolygonRgn(pptl []POINT, iMode co.POLYF) HRGN {
	ret, _, _ := syscall.Syscall(proc.CreatePolygonRgn.Addr(), 3,
		uintptr(unsafe.Pointer(&pptl[0])), uintptr(len(pptl)), uintptr(iMode))
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "CreatePolygonRgn"))
	}
	return HRGN(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpolypolygonrgn
func CreatePolyPolygonRgn(pptl [][]POINT, iMode co.POLYF) HRGN {
	counts := make([]int32, 0, len(pptl)) // store length of each array
	for _, ptArr := range pptl {
		counts = append(counts, int32(len(ptArr)))
	}
	ret, _, _ := syscall.Syscall6(proc.CreatePolyPolygonRgn.Addr(), 4,
		uintptr(unsafe.Pointer(&pptl[0])),
		uintptr(unsafe.Pointer(&counts[0])),
		uintptr(len(counts)),
		uintptr(iMode), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "CreatePolyPolygonRgn"))
	}
	return HRGN(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createrectrgn
func CreateRectRgn(x1, y1, x2, y2 int32) HRGN {
	ret, _, _ := syscall.Syscall6(proc.CreateRectRgn.Addr(), 4,
		uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2),
		0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "CreateRectRgn"))
	}
	return HRGN(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createrectrgnindirect
func CreateRectRgnIndirect(lprect *RECT) HRGN {
	ret, _, _ := syscall.Syscall(proc.CreateRectRgnIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lprect)), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "CreateRectRgnIndirect"))
	}
	return HRGN(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createroundrectrgn
func CreateRoundRectRgn(x1, y1, x2, y2, w, h int32) HRGN {
	ret, _, _ := syscall.Syscall6(proc.CreateRoundRectRgn.Addr(), 6,
		uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2),
		uintptr(w), uintptr(h))
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "CreateRoundRectRgn"))
	}
	return HRGN(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-combinergn
func (hRgn HRGN) CombineRgn(hrgnSrc1, hrgnSrc2 HRGN, iMode co.RGN) {
	ret, _, _ := syscall.Syscall6(proc.CombineRgn.Addr(), 4,
		uintptr(hRgn), uintptr(hrgnSrc1), uintptr(hrgnSrc2), uintptr(iMode),
		0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "CombineRgn"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-equalrgn
func (hRgn HRGN) EqualRgn(hrgn2 HRGN) bool {
	ret, _, _ := syscall.Syscall(proc.EqualRgn.Addr(), 2,
		uintptr(hRgn), uintptr(hrgn2), 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-offsetrgn
func (hRgn HRGN) OffsetRgn(x, y int32) {
	ret, _, _ := syscall.Syscall(proc.OffsetRgn.Addr(), 3,
		uintptr(hRgn), uintptr(x), uintptr(y))
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "OffsetRgn"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ptinregion
func (hRgn HRGN) PtInRegion(x, y int32) bool {
	ret, _, _ := syscall.Syscall(proc.PtInRegion.Addr(), 3,
		uintptr(hRgn), uintptr(x), uintptr(y))
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rectinregion
func (hRgn HRGN) RectInRegion(lprect *RECT) bool {
	ret, _, _ := syscall.Syscall(proc.RectInRegion.Addr(), 2,
		uintptr(hRgn), uintptr(unsafe.Pointer(lprect)), 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setrectrgn
func (hRgn HRGN) SetRectRgn(left, top, right, bottom int32) {
	ret, _, _ := syscall.Syscall6(proc.SetRectRgn.Addr(), 5,
		uintptr(hRgn),
		uintptr(left), uintptr(top), uintptr(right), uintptr(bottom),
		0)
	if ret == 0 {
		panic(NewWinError(co.ERROR_E_UNEXPECTED, "SetRectRgn"))
	}
}
