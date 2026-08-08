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

// [IWICBitmap] COM interface.
//
// Implements [OleResource].
//
// [IWICBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmap
type IWICBitmap struct{ IWICBitmapSource }

type _IWICBitmapVt struct {
	utl.IUnknownVt
	Lock          uintptr
	SetPalette    uintptr
	SetResolution uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmap) IID() *co.IID {
	return &cowic.IID_IWICBitmap
}

// [Lock] method.
//
// [Lock]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmap-lock
func (me *IWICBitmap) Lock(
	releaser *win.OleReleaser,
	lock *WICRect,
	flags cowic.WICBMP_LOCK,
) (*IWICBitmapLock, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapVt](me.Ppvt()).Lock,
		me.Ppvt(),
		uintptr(unsafe.Pointer(lock)),
		uintptr(flags),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmapLock](ret, ppvtQueried, releaser)
}

// [SetPalette] method.
//
// [SetPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmap-setpalette
func (me *IWICBitmap) SetPalette(palette *IWICPalette) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapVt](me.Ppvt()).SetPalette,
		me.Ppvt(),
		palette.Ppvt())
	return utl.HresultToError(ret)
}

// [SetResolution] method.
//
// [SetResolution]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmap-setresolution
func (me *IWICBitmap) SetResolution(dpiX, dpiY float64) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapVt](me.Ppvt()).SetResolution,
		me.Ppvt(),
		uintptr(dpiX),
		uintptr(dpiY))
	return utl.HresultToError(ret)
}
