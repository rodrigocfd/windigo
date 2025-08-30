//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
)

// Handle to a [bitmap].
//
// [bitmap]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbitmap
type HBITMAP HGDIOBJ

// [CreateBitmap] function.
//
// ⚠️ You must defer [HBITMAP.DeleteObject].
//
// [CreateBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbitmap
func CreateBitmap(width, height int, numPlanes, bitCount uint, bits []byte) (HBITMAP, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateBitmap, "CreateBitmap"),
		uintptr(int32(width)),
		uintptr(int32(height)),
		uintptr(uint32(numPlanes)),
		uintptr(uint32(bitCount)),
		uintptr(unsafe.Pointer(&bits[0])))
	if ret == 0 {
		return HBITMAP(0), co.ERROR_INVALID_PARAMETER
	}
	return HBITMAP(ret), nil
}

var _CreateBitmap *syscall.Proc

// [CreateBitmapIndirect] function.
//
// ⚠️ You must defer [HBITMAP.DeleteObject].
//
// [CreateBitmapIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbitmapindirect
func CreateBitmapIndirect(bm *BITMAP) (HBITMAP, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_CreateBitmapIndirect, "CreateBitmapIndirect"),
		uintptr(unsafe.Pointer(bm)))
	if ret == 0 {
		return HBITMAP(0), co.ERROR_INVALID_PARAMETER
	}
	return HBITMAP(ret), nil
}

var _CreateBitmapIndirect *syscall.Proc

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBmp HBITMAP) DeleteObject() error {
	return HGDIOBJ(hBmp).DeleteObject()
}

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hBmp HBITMAP) GetObject() (BITMAP, error) {
	var bmp BITMAP
	if err := HGDIOBJ(hBmp).GetObject(unsafe.Sizeof(bmp), unsafe.Pointer(&bmp)); err != nil {
		return BITMAP{}, err
	} else {
		return bmp, nil
	}
}
