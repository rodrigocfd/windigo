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

// ⚠️ You must defer HBITMAP.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbitmap
func CreateBitmap(width, height int32,
	numPlanes, bitCount uint32, bits *byte) HBITMAP {

	ret, _, err := syscall.Syscall6(proc.CreateBitmap.Addr(), 5,
		uintptr(width), uintptr(height), uintptr(numPlanes), uintptr(bitCount),
		uintptr(unsafe.Pointer(bits)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret)
}

// ⚠️ You must defer HBITMAP.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbitmapindirect
func CreateBitmapIndirect(bmp *BITMAP) HBITMAP {
	ret, _, err := syscall.Syscall(proc.CreateBitmapIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(bmp)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBmp HBITMAP) DeleteObject() {
	HGDIOBJ(hBmp).DeleteObject()
}

//------------------------------------------------------------------------------

// A handle to a brush.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbrush
type HBRUSH HGDIOBJ

// ⚠️ You must defer HBRUSH.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbrushindirect
func CreateBrushIndirect(lb *LOGBRUSH) HBRUSH {
	ret, _, err := syscall.Syscall(proc.CreateBrushIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lb)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// ⚠️ You must defer HBRUSH.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createhatchbrush
func CreateHatchBrush(hatch co.HS, color COLORREF) HBRUSH {
	ret, _, err := syscall.Syscall(proc.CreateHatchBrush.Addr(), 2,
		uintptr(hatch), uintptr(color), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// ⚠️ You must defer HBRUSH.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpatternbrush
func CreatePatternBrush(hBmp HBITMAP) HBRUSH {
	ret, _, err := syscall.Syscall(proc.CreatePatternBrush.Addr(), 1,
		uintptr(hBmp), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// ⚠️ You must defer HBRUSH.DeleteObject().
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
func GetSysColorBrush(index co.COLOR) HBRUSH {
	ret, _, err := syscall.Syscall(proc.GetSysColorBrush.Addr(), 1,
		uintptr(index), 0, 0)
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

// ⚠️ You must defer HFONT.DeleteObject().
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
func CreatePen(style co.PS, width int32, color COLORREF) HPEN {
	ret, _, err := syscall.Syscall(proc.CreatePen.Addr(), 3,
		uintptr(style), uintptr(width), uintptr(color))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HPEN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpenindirect
func CreatePenIndirect(lp *LOGPEN) HPEN {
	ret, _, err := syscall.Syscall(proc.CreatePenIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lp)), 0, 0)
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

// ⚠️ You must defer HRGN.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createrectrgnindirect
func CreateRectRgnIndirect(rc *RECT) HRGN {
	ret, _, err := syscall.Syscall(proc.CreateRectRgnIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(rc)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// ⚠️ You must defer HRGN.DeleteObject().
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
func (hRgn HRGN) CombineRgn(hrgnSrc1, hrgnSrc2 HRGN, mode co.RGN) co.REGION {
	ret, _, err := syscall.Syscall6(proc.CombineRgn.Addr(), 4,
		uintptr(hRgn), uintptr(hrgnSrc1), uintptr(hrgnSrc2), uintptr(mode), 0, 0)
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
