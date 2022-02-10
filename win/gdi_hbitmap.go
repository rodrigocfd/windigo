package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hBmp HBITMAP) GetObject(bmp *BITMAP) {
	ret, _, err := syscall.Syscall(proc.GetObject.Addr(), 3,
		uintptr(hBmp), unsafe.Sizeof(*bmp), uintptr(unsafe.Pointer(bmp)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
