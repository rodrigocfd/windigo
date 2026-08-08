//go:build windows

package winwic

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cowic"
)

// [IWICBitmapSource] COM interface.
//
// Implements [OleResource].
//
// [IWICBitmapSource]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapsource
type IWICBitmapSource struct{ win.IUnknown }

type _IWICBitmapSourceVt struct {
	utl.IUnknownVt
	GetSize        uintptr
	GetPixelFormat uintptr
	GetResolution  uintptr
	CopyPalette    uintptr
	CopyPixels     uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapSource) IID() *co.IID {
	return &cowic.IID_IWICBitmapSource
}

// [CopyPalette] method.
//
// [CopyPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapsource-copypalette
func (me *IWICBitmapSource) CopyPalette(palette *IWICPalette) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapSourceVt](me.Ppvt()).CopyPalette,
		me.Ppvt(),
		palette.Ppvt())
	return utl.HresultToError(ret)
}

// [CopyPixels] method.
//
// [CopyPixels]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapsource-copypixels
func (me *IWICBitmapSource) CopyPixels(pRc *WICRect, stride, szBuffer int, pBuffer *byte) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapSourceVt](me.Ppvt()).CopyPixels,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pRc)),
		uintptr(uint32(stride)),
		uintptr(uint32(szBuffer)),
		uintptr(unsafe.Pointer(pBuffer)))
	return utl.HresultToError(ret)
}

// [GetPixelFormat] method.
//
// [GetPixelFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapsource-getpixelformat
func (me *IWICBitmapSource) GetPixelFormat() (cowic.WIC_PIXELFORMAT, error) {
	return utl.OleCallReturnStruct[cowic.WIC_PIXELFORMAT](me,
		utl.Vt[_IWICBitmapSourceVt](me.Ppvt()).GetPixelFormat)
}

// [GetResolution] method.
//
// [GetResolution]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapsource-getresolution
func (me *IWICBitmapSource) GetResolution() (dpiX float64, dpiY float64, hr error) {
	var cx, cy float64
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapSourceVt](me.Ppvt()).GetResolution,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&cx)),
		uintptr(unsafe.Pointer(&cy)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return cx, cy, nil
	} else {
		return 0, 0, hr
	}
}

// [GetSize] method.
//
// [GetSize]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapsource-getsize
func (me *IWICBitmapSource) GetSize() (win.SIZE, error) {
	var cx, cy uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapSourceVt](me.Ppvt()).GetSize,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&cx)),
		uintptr(unsafe.Pointer(&cy)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return win.SIZE{Cx: int32(cx), Cy: int32(cy)}, nil
	} else {
		return win.SIZE{}, hr
	}
}
