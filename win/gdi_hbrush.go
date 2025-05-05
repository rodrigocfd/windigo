//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to a [brush].
//
// [brush]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbrush
type HBRUSH HGDIOBJ

// [CreateBrushIndirect] function.
//
// ⚠️ You must defer HBRUSH.DeleteObject().
//
// [CreateBrushIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbrushindirect
func CreateBrushIndirect(lb *LOGBRUSH) (HBRUSH, error) {
	ret, _, _ := syscall.SyscallN(_CreateBrushIndirect.Addr(),
		uintptr(unsafe.Pointer(lb)))
	if ret == 0 {
		return HBRUSH(0), co.ERROR_INVALID_PARAMETER
	}
	return HBRUSH(ret), nil
}

var _CreateBrushIndirect = dll.Gdi32.NewProc("CreateBrushIndirect")

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBrush HBRUSH) DeleteObject() error {
	return HGDIOBJ(hBrush).DeleteObject()
}

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hBrush HBRUSH) GetObject() (LOGBRUSH, error) {
	var lb LOGBRUSH
	if err := HGDIOBJ(hBrush).GetObject(unsafe.Sizeof(lb), unsafe.Pointer(&lb)); err != nil {
		return LOGBRUSH{}, err
	} else {
		return lb, nil
	}
}
