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

// [IWICBitmapDecoder] COM interface.
//
// Implements [OleResource].
//
// Example:
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
//	decoder, _ := factory.CreateDecoderFromFilename(
//		rel,
//		"C:\\Temp\\foo.png",
//		"",
//		co.GENERIC_READ,
//		cowic.WICDEC_METADATACACHE_OnDemand,
//	)
//
// [IWICBitmapDecoder]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapdecoder
type IWICBitmapDecoder struct{ win.IUnknown }

type _IWICBitmapDecoderVt struct {
	utl.IUnknownVt
	QueryCapability        uintptr
	Initialize             uintptr
	GetContainerFormat     uintptr
	GetDecoderInfo         uintptr
	CopyPalette            uintptr
	GetMetadataQueryReader uintptr
	GetPreview             uintptr
	GetColorContexts       uintptr
	GetThumbnail           uintptr
	GetFrameCount          uintptr
	GetFrame               uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapDecoder) IID() *co.IID {
	return &cowic.IID_IWICBitmapDecoder
}

// [CopyPalette] method.
//
// [CopyPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-copypalette
func (me *IWICBitmapDecoder) CopyPalette(palette *IWICPalette) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapDecoderVt](me.Ppvt()).CopyPalette,
		me.Ppvt(),
		palette.Ppvt())
	return utl.HresultToError(ret)
}

// [GetContainerFormat] method.
//
// [GetContainerFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-getcontainerformat
func (me *IWICBitmapDecoder) GetContainerFormat() (cowic.WIC_CONTAINER, error) {
	return utl.OleCallReturnStruct[cowic.WIC_CONTAINER](me,
		utl.Vt[_IWICBitmapDecoderVt](me.Ppvt()).GetContainerFormat)
}

// [GetDecoderInfo] method.
//
// [GetDecoderInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-getdecoderinfo
func (me *IWICBitmapDecoder) GetDecoderInfo(releaser *win.OleReleaser) (*IWICBitmapDecoderInfo, error) {
	return utl.OleNewFromCallWithoutParms[*IWICBitmapDecoderInfo](me, releaser,
		utl.Vt[_IWICBitmapDecoderVt](me.Ppvt()).GetDecoderInfo)
}

// [GetFrame] method.
//
// [GetFrame]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-getframe
func (me *IWICBitmapDecoder) GetFrame(releaser *win.OleReleaser, index int) (*IWICBitmapFrameDecode, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapDecoderVt](me.Ppvt()).GetFrame,
		me.Ppvt(),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmapFrameDecode](ret, ppvtQueried, releaser)
}

// [GetFrameCount] method.
//
// [GetFrameCount]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-getframecount
func (me *IWICBitmapDecoder) GetFrameCount() (int, error) {
	var count uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapDecoderVt](me.Ppvt()).CopyPalette,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&count)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(count), nil
	} else {
		return 0, hr
	}
}

// [GetPreview] method.
//
// [GetPreview]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-getpreview
func (me *IWICBitmapDecoder) GetPreview(releaser *win.OleReleaser) (*IWICBitmapSource, error) {
	return utl.OleNewFromCallWithoutParms[*IWICBitmapSource](me, releaser,
		utl.Vt[_IWICBitmapDecoderVt](me.Ppvt()).GetPreview)
}

// [GetThumbnail] method.
//
// [GetThumbnail]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-getthumbnail
func (me *IWICBitmapDecoder) GetThumbnail(releaser *win.OleReleaser) (*IWICBitmapSource, error) {
	return utl.OleNewFromCallWithoutParms[*IWICBitmapSource](me, releaser,
		utl.Vt[_IWICBitmapDecoderVt](me.Ppvt()).GetThumbnail)
}

// [Initialize] method.
//
// [Initialize]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-initialize
func (me *IWICBitmapDecoder) Initialize(stream *win.IStream, cacheOpts cowic.WICDEC_METADATACACHE) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapDecoderVt](me.Ppvt()).Initialize,
		me.Ppvt(),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		uintptr(cacheOpts))
	return utl.HresultToError(ret)
}

// [QueryCapability] method.
//
// [QueryCapability]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-querycapability
func (me *IWICBitmapDecoder) QueryCapability(stream *win.IStream) (cowic.WICDEC_CAP, error) {
	var capability uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapDecoderVt](me.Ppvt()).QueryCapability,
		me.Ppvt(),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		uintptr(unsafe.Pointer(&capability)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return cowic.WICDEC_CAP(capability), nil
	} else {
		return cowic.WICDEC_CAP(0), hr
	}
}
