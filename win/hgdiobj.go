package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a GDI object.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hgdiobj
type HGDIOBJ HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hGdiObj HGDIOBJ) DeleteObject() {
	ret, _, err := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hGdiObj), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

//------------------------------------------------------------------------------

// A handle to a bitmap.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbitmap
type HBITMAP HGDIOBJ

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBmp HBITMAP) DeleteObject() {
	HGDIOBJ(hBmp).DeleteObject()
}

//------------------------------------------------------------------------------

// A handle to a brush.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbrush
type HBRUSH HGDIOBJ

// ⚠️ You must defer DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createsolidbrush
func CreateSolidBrush(color COLORREF) HBRUSH {
	ret, _, err := syscall.Syscall(proc.CreateSolidBrush.Addr(), 1,
		uintptr(color), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// Not an actual Win32 function, just a tricky conversion to create a brush from
// a system color, particularly used when registering a window class.
func CreateSysColorBrush(sysColor co.COLOR) HBRUSH {
	return HBRUSH(sysColor + 1)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolorbrush
func GetSysColorBrush(nIndex co.COLOR) HBRUSH {
	ret, _, err := syscall.Syscall(proc.GetSysColorBrush.Addr(), 1,
		uintptr(nIndex), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBrush HBRUSH) DeleteObject() {
	HGDIOBJ(hBrush).DeleteObject()
}

//------------------------------------------------------------------------------

// A handle to a font.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hfont
type HFONT HGDIOBJ

// ⚠️ You must defer DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createfontindirectw
func CreateFontIndirect(lf *LOGFONT) HFONT {
	ret, _, err := syscall.Syscall(proc.CreateFontIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lf)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HFONT(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hFont HFONT) DeleteObject() {
	HGDIOBJ(hFont).DeleteObject()
}

//------------------------------------------------------------------------------

// A handle to a pen.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hpen
type HPEN HGDIOBJ

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpen
func CreatePen(iStyle co.PS, cWidth int32, color COLORREF) HPEN {
	ret, _, err := syscall.Syscall(proc.CreatePen.Addr(), 3,
		uintptr(iStyle), uintptr(cWidth), uintptr(color))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HPEN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpenindirect
func CreatePenIndirect(plpen *LOGPEN) HPEN {
	ret, _, err := syscall.Syscall(proc.CreatePenIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(plpen)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HPEN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hPen HPEN) DeleteObject() {
	HGDIOBJ(hPen).DeleteObject()
}

//------------------------------------------------------------------------------

// A handle to a region.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hrgn
type HRGN HGDIOBJ

// ⚠️ You must defer DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createrectrgnindirect
func CreateRectRgnIndirect(lprect *RECT) HRGN {
	ret, _, err := syscall.Syscall(proc.CreateRectRgnIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lprect)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// ⚠️ You must defer DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createroundrectrgn
func CreateRoundRectRgn(x1, y1, x2, y2, w, h int32) HRGN {
	ret, _, err := syscall.Syscall6(proc.CreateRoundRectRgn.Addr(), 6,
		uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2),
		uintptr(w), uintptr(h))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hRgn HRGN) DeleteObject() {
	HGDIOBJ(hRgn).DeleteObject()
}

// Combines the two regions and stores the result in current region.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-combinergn
func (hRgn HRGN) CombineRgn(hrgnSrc1, hrgnSrc2 HRGN, iMode co.RGN) co.REGION {
	ret, _, err := syscall.Syscall6(proc.CombineRgn.Addr(), 4,
		uintptr(hRgn), uintptr(hrgnSrc1), uintptr(hrgnSrc2), uintptr(iMode), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-offsetrgn
func (hRgn HRGN) OffsetRgn(x, y int32) co.REGION {
	ret, _, err := syscall.Syscall(proc.OffsetRgn.Addr(), 3,
		uintptr(hRgn), uintptr(x), uintptr(y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}
