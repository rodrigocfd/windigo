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

// [IWICPalette] COM interface.
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var factory *winwic.IWICImagingFactory
//	_ = win.CoCreateInstance(
//		rel,
//		&cowic.CLSID_WICImagingFactory,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&factory,
//	)
//
//	palette, _ := factory.CreatePalette(rel)
//
// [IWICPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicpalette
type IWICPalette struct{ win.IUnknown }

type _IWICPaletteVt struct {
	utl.IUnknownVt
	InitializePredefined  uintptr
	InitializeCustom      uintptr
	InitializeFromBitmap  uintptr
	InitializeFromPalette uintptr
	GetType               uintptr
	GetColorCount         uintptr
	GetColors             uintptr
	IsBlackWhite          uintptr
	IsGrayscale           uintptr
	HasAlpha              uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICPalette) IID() *co.IID {
	return &cowic.IID_IWICPalette
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IWICPalette) AddRef(releaser *win.OleReleaser) *IWICPalette {
	return utl.OleNewFromAddRef[*IWICPalette](me, releaser)
}

// [GetColorCount] method.
//
// [GetColorCount]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicpalette-getcolorcount
func (me *IWICPalette) GetColorCount() (int, error) {
	var count uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICPaletteVt](me.Ppvt()).GetColorCount,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&count)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(count), nil
	} else {
		return 0, hr
	}
}

// [HasAlpha] method.
//
// [HasAlpha]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicpalette-hasalpha
func (me *IWICPalette) HasAlpha() (bool, error) {
	var hasAlpha win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICPaletteVt](me.Ppvt()).HasAlpha,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&hasAlpha)))
	return utl.HresultToBoolError(int32(hasAlpha), ret)
}

// [InitializeFromBitmap] method.
//
// [InitializeFromBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicpalette-initializefrombitmap
func (me *IWICPalette) InitializeFromBitmap(
	surface *IWICBitmapSource,
	numColors int,
	addTransparentColor bool,
) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICPaletteVt](me.Ppvt()).InitializeFromBitmap,
		me.Ppvt(),
		surface.Ppvt(),
		uintptr(uint32(numColors)),
		utl.BoolToUintptr(addTransparentColor))
	return utl.HresultToError(ret)
}

// [InitializeFromPalette] method.
//
// [InitializeFromPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicpalette-initializefrompalette
func (me *IWICPalette) InitializeFromPalette(palette *IWICPalette) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICPaletteVt](me.Ppvt()).InitializeFromPalette,
		me.Ppvt(),
		palette.Ppvt())
	return utl.HresultToError(ret)
}

// [IsBlackWhite] method.
//
// [IsBlackWhite]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicpalette-isblackwhite
func (me *IWICPalette) IsBlackWhite() (bool, error) {
	var isBW win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICPaletteVt](me.Ppvt()).IsBlackWhite,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&isBW)))
	return utl.HresultToBoolError(int32(isBW), ret)
}

// [IsGrayscale] method.
//
// [IsGrayscale]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicpalette-isgrayscale
func (me *IWICPalette) IsGrayscale() (bool, error) {
	var isGrayscale win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICPaletteVt](me.Ppvt()).IsGrayscale,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&isGrayscale)))
	return utl.HresultToBoolError(int32(isGrayscale), ret)
}
