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
// ⚠️ You must defer [HPEN.DeleteObject].
//
// [CreatePen]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpen
func CreatePen(style co.PS, width uint, color COLORREF) (HPEN, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Gdi(&_CreatePen, "CreatePen"),
		uintptr(style),
		uintptr(int32(width)),
		uintptr(color))
	if ret == 0 {
		return HPEN(0), co.ERROR_INVALID_PARAMETER
	}
	return HPEN(ret), nil
}

var _CreatePen *syscall.Proc

// [CreatePenIndirect] function.
//
// ⚠️ You must defer [HPEN.DeleteObject].
//
// [CreatePenIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpenindirect
func CreatePenIndirect(lp *LOGPEN) (HPEN, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Gdi(&_CreatePenIndirect, "CreatePenIndirect"),
		uintptr(unsafe.Pointer(lp)))
	if ret == 0 {
		return HPEN(0), co.ERROR_INVALID_PARAMETER
	}
	return HPEN(ret), nil
}

var _CreatePenIndirect *syscall.Proc

// [ExtCreatePen] function.
//
// ⚠️ You must defer [HPEN.DeleteObject].
//
// [ExtCreatePen]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-extcreatepen
func ExtCreatePen(
	penType co.PS_TYPE,
	penStyle co.PS_STYLE,
	endCap co.PS_ENDCAP,
	width uint,
	brush *LOGBRUSH,
	styleLengths []uint,
) (HPEN, error) {
	var nLens int
	var pLens unsafe.Pointer
	if styleLengths != nil {
		nLens = len(styleLengths)
		pLens = unsafe.Pointer(&styleLengths[0])
	}

	ret, _, _ := syscall.SyscallN(
		dll.Gdi(&_ExtCreatePen, "ExtCreatePen"),
		uintptr(uint32(penType)|uint32(penStyle)|uint32(endCap)),
		uintptr(width),
		uintptr(unsafe.Pointer(brush)),
		uintptr(nLens),
		uintptr(pLens))
	if ret == 0 {
		return HPEN(0), co.ERROR_INVALID_PARAMETER
	}
	return HPEN(ret), nil
}

var _ExtCreatePen *syscall.Proc

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
