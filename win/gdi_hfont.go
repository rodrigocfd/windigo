//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

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
	height int,
	width int,
	escapement int,
	orientation int,
	weight int,
	italic bool,
	underline bool,
	strikeOut bool,
	charSet co.CHARSET,
	outPrecision co.OUT_PRECIS,
	clipPrecision co.CLIP_PRECIS,
	quality co.QUALITY,
	pitchAndFamily co.FF,
	faceName string,
) (HFONT, error) {
	faceName16 := wstr.NewBufWith[wstr.Stack20](faceName, wstr.EMPTY_IS_NIL)
	ret, _, _ := syscall.SyscallN(_CreateFontW.Addr(),
		uintptr(height), uintptr(width), uintptr(escapement),
		uintptr(orientation), uintptr(height),
		wutil.BoolToUintptr(italic), wutil.BoolToUintptr(underline),
		wutil.BoolToUintptr(strikeOut),
		uintptr(charSet), uintptr(outPrecision), uintptr(clipPrecision),
		uintptr(quality), uintptr(pitchAndFamily),
		uintptr(faceName16.UnsafePtr()))
	if ret == 0 {
		return HFONT(0), co.ERROR_INVALID_PARAMETER
	}
	return HFONT(ret), nil
}

var _CreateFontW = dll.Gdi32.NewProc("CreateFontW")

// [CreateFontIndirect] function.
//
// ⚠️ You must defer [HFONT.DeleteObject].
//
// [CreateFontIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createfontindirectw
func CreateFontIndirect(lf *LOGFONT) (HFONT, error) {
	ret, _, _ := syscall.SyscallN(_CreateFontIndirectW.Addr(),
		uintptr(unsafe.Pointer(lf)))
	if ret == 0 {
		return HFONT(0), co.ERROR_INVALID_PARAMETER
	}
	return HFONT(ret), nil
}

var _CreateFontIndirectW = dll.Gdi32.NewProc("CreateFontIndirectW")

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
