//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
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
	height, width int,
	escapement int,
	orientation int,
	weight int,
	italic, underline, strikeOut bool,
	charSet co.CHARSET,
	outPrecision co.OUT_PRECIS,
	clipPrecision co.CLIP_PRECIS,
	quality co.QUALITY,
	pitch co.PITCH,
	family co.FF,
	faceName string,
) (HFONT, error) {
	var wFaceName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateFontW, "CreateFontW"),
		uintptr(int32(height)),
		uintptr(int32(width)),
		uintptr(int32(escapement)),
		uintptr(int32(orientation)),
		uintptr(int32(height)),
		utl.BoolToUintptr(italic),
		utl.BoolToUintptr(underline),
		utl.BoolToUintptr(strikeOut),
		uintptr(charSet),
		uintptr(outPrecision),
		uintptr(clipPrecision),
		uintptr(quality),
		uintptr(uint8(pitch)|uint8(family)),
		uintptr(wFaceName.EmptyIsNil(faceName)))
	if ret == 0 {
		return HFONT(0), co.ERROR_INVALID_PARAMETER
	}
	return HFONT(ret), nil
}

var _CreateFontW *syscall.Proc

// [CreateFontIndirect] function.
//
// ⚠️ You must defer [HFONT.DeleteObject].
//
// [CreateFontIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createfontindirectw
func CreateFontIndirect(lf *LOGFONT) (HFONT, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateFontIndirectW, "CreateFontIndirectW"),
		uintptr(unsafe.Pointer(lf)))
	if ret == 0 {
		return HFONT(0), co.ERROR_INVALID_PARAMETER
	}
	return HFONT(ret), nil
}

var _CreateFontIndirectW *syscall.Proc

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
