//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to a [pen].
//
// [pen]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hpen
type HPEN HGDIOBJ

// [CreatePen] function.
//
// ⚠️ You must defer HPEN.DeleteObject().
//
// [CreatePen]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpen
func CreatePen(style co.PS, width uint, color COLORREF) (HPEN, error) {
	ret, _, _ := syscall.SyscallN(_CreatePen.Addr(),
		uintptr(style), uintptr(width), uintptr(color))
	if ret == 0 {
		return HPEN(0), co.ERROR_INVALID_PARAMETER
	}
	return HPEN(ret), nil
}

var _CreatePen = dll.Gdi32.NewProc("CreatePen")

// [CreatePenIndirect] function.
//
// ⚠️ You must defer HPEN.DeleteObject().
//
// [CreatePenIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpenindirect
func CreatePenIndirect(lp *LOGPEN) (HPEN, error) {
	ret, _, _ := syscall.SyscallN(_CreatePenIndirect.Addr(),
		uintptr(unsafe.Pointer(lp)))
	if ret == 0 {
		return HPEN(0), co.ERROR_INVALID_PARAMETER
	}
	return HPEN(ret), nil
}

var _CreatePenIndirect = dll.Gdi32.NewProc("CreatePenIndirect")

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hPen HPEN) DeleteObject() error {
	return HGDIOBJ(hPen).DeleteObject()
}

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hPen HPEN) GetObject() (LOGPEN, error) {
	var lp LOGPEN
	if err := HGDIOBJ(hPen).GetObject(unsafe.Sizeof(lp), unsafe.Pointer(&lp)); err != nil {
		return LOGPEN{}, err
	} else {
		return lp, nil
	}
}
