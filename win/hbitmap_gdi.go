//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [bitmap].
//
// [bitmap]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbitmap
type HBITMAP HGDIOBJ

// [CreateBitmap] function.
//
// ⚠️ You must defer HBITMAP.DeleteObject().
//
// [CreateBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbitmap
func CreateBitmap(width, height int32,
	numPlanes, bitCount uint32, bits *byte) HBITMAP {

	ret, _, err := syscall.SyscallN(proc.CreateBitmap.Addr(),
		uintptr(width), uintptr(height), uintptr(numPlanes), uintptr(bitCount),
		uintptr(unsafe.Pointer(bits)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret)
}

// [CreateBitmapIndirect] function.
//
// ⚠️ You must defer HBITMAP.DeleteObject().
//
// [CreateBitmapIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbitmapindirect
func CreateBitmapIndirect(bmp *BITMAP) HBITMAP {
	ret, _, err := syscall.SyscallN(proc.CreateBitmapIndirect.Addr(),
		uintptr(unsafe.Pointer(bmp)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret)
}

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBmp HBITMAP) DeleteObject() error {
	return HGDIOBJ(hBmp).DeleteObject()
}

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hBmp HBITMAP) GetObject(bmp *BITMAP) {
	ret, _, err := syscall.SyscallN(proc.GetObject.Addr(),
		uintptr(hBmp), unsafe.Sizeof(*bmp), uintptr(unsafe.Pointer(bmp)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
