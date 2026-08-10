//go:build windows

package winwic

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cowic"
)

// [IWICBitmapEncoder] COM interface.
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
//	encoder, _ := factory.CreateEncoder(
//		rel,
//		&cowic.WIC_CONTAINER_Bmp,
//		nil,
//	)
//
// [IWICBitmapEncoder]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapencoder
type IWICBitmapEncoder struct{ win.IUnknown }

type _IWICBitmapEncoderVt struct {
	utl.IUnknownVt
	Initialize             uintptr
	GetContainerFormat     uintptr
	GetEncoderInfo         uintptr
	SetColorContexts       uintptr
	SetPalette             uintptr
	SetThumbnail           uintptr
	SetPreview             uintptr
	CreateNewFrame         uintptr
	Commit                 uintptr
	GetMetadataQueryWriter uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapEncoder) IID() *co.IID {
	return &cowic.IID_IWICBitmapEncoder
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IWICBitmapEncoder) AddRef(releaser *win.OleReleaser) *IWICBitmapEncoder {
	return utl.OleNewFromAddRef[*IWICBitmapEncoder](me, releaser)
}

// [Commit] method.
//
// [Commit]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-commit
func (me *IWICBitmapEncoder) Commit() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IWICBitmapEncoderVt](me.Ppvt()).Commit)
}

// [GetContainerFormat] method.
//
// [GetContainerFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-getcontainerformat
func (me *IWICBitmapEncoder) GetContainerFormat() (cowic.WIC_CONTAINER, error) {
	return utl.OleCallReturnStruct[cowic.WIC_CONTAINER](me,
		utl.Vt[_IWICBitmapEncoderVt](me.Ppvt()).GetContainerFormat)
}

// [GetEncoderInfo] method.
//
// [GetEncoderInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-getencoderinfo
func (me *IWICBitmapEncoder) GetEncoderInfo(releaser *win.OleReleaser) (*IWICBitmapEncoderInfo, error) {
	return utl.OleNewFromCallWithoutParms[*IWICBitmapEncoderInfo](me, releaser,
		utl.Vt[_IWICBitmapEncoderVt](me.Ppvt()).GetEncoderInfo)
}

// [Initialize] method.
//
// [Initialize]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-initialize
func (me *IWICBitmapEncoder) Initialize(stream *win.IStream, cacheOpt cowic.WICENC_CACHE) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapEncoderVt](me.Ppvt()).Initialize,
		me.Ppvt(),
		stream.Ppvt(),
		uintptr(cacheOpt))
	return utl.HresultToError(ret)
}

// [SetPalette] method.
//
// [SetPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-setpalette
func (me *IWICBitmapEncoder) SetPalette(palette *IWICPalette) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapEncoderVt](me.Ppvt()).SetPalette,
		me.Ppvt(),
		palette.Ppvt())
	return utl.HresultToError(ret)
}

// [SetPreview] method.
//
// [SetPreview]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-setpreview
func (me *IWICBitmapEncoder) SetPreview(preview *IWICBitmapSource) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapEncoderVt](me.Ppvt()).SetPreview,
		me.Ppvt(),
		preview.Ppvt())
	return utl.HresultToError(ret)
}

// [SetThumbnail] method.
//
// [SetThumbnail]:
func (me *IWICBitmapEncoder) SetThumbnail(thumbnail *IWICBitmapSource) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapEncoderVt](me.Ppvt()).SetThumbnail,
		me.Ppvt(),
		thumbnail.Ppvt())
	return utl.HresultToError(ret)
}
