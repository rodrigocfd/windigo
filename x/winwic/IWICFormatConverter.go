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

// [IWICFormatConverter] COM interface.
//
// [IWICFormatConverter]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicformatconverter
type IWICFormatConverter struct{ IWICBitmapSource }

type _IWICFormatConverterVt struct {
	_IWICBitmapSourceVt
	Initialize uintptr
	CanConvert uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICFormatConverter) IID() *co.IID {
	return &cowic.IID_IWICFormatConverter
}

// [CanConvert] method.
//
// [CanConvert]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicformatconverter-canconvert
func (me *IWICFormatConverter) CanConvert(pSrcPixelFormat, pDestPixelFormat *cowic.WIC_PIXELFORMAT) (bool, error) {
	var canConvert win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICFormatConverterVt](me.Ppvt()).CanConvert,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pSrcPixelFormat)),
		uintptr(unsafe.Pointer(pDestPixelFormat)),
		uintptr(unsafe.Pointer(&canConvert)))
	return utl.HresultToBoolError(int32(canConvert), ret)
}

// [Initialize] method.
//
// [Initialize]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicformatconverter-initialize
func (me *IWICFormatConverter) Initialize(
	source *IWICBitmapSource,
	pDestFormat *cowic.WIC_PIXELFORMAT,
	dither cowic.WICBMP_DITHER,
	palette *IWICPalette,
	alphaThresholdPercent float64,
	paletteTranslate cowic.WICBMP_PAL,
) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICFormatConverterVt](me.Ppvt()).Initialize,
		me.Ppvt(),
		uintptr(unsafe.Pointer(source.Ppvt())),
		uintptr(unsafe.Pointer(pDestFormat)),
		uintptr(unsafe.Pointer(&dither)),
		utl.OlePpvtOrNil(palette),
		uintptr(alphaThresholdPercent),
		uintptr(paletteTranslate))
	return utl.HresultToError(ret)
}
