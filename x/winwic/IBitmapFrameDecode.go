//go:build windows

package winwic

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cowic"
)

// [IWICBitmapFrameDecode] COM interface.
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
//	frame, _ := decoder.GetFrame(rel, 0)
//
// [IWICBitmapFrameDecode]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapframedecode
type IWICBitmapFrameDecode struct{ IWICBitmapSource }

type _IWICBitmapFrameDecodeVt struct {
	_IWICBitmapSourceVt
	GetMetadataQueryReader uintptr
	GetColorContexts       uintptr
	GetThumbnail           uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapFrameDecode) IID() *co.IID {
	return &cowic.IID_IWICBitmapFrameDecode
}

// [GetThumbnail] method.
//
// [GetThumbnail]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapframedecode-getthumbnail
func (me *IWICBitmapFrameDecode) GetThumbnail(releaser *win.OleReleaser) (*IWICBitmapSource, error) {
	return utl.OleNewFromCallWithoutParms[*IWICBitmapSource](me, releaser,
		utl.Vt[_IWICBitmapFrameDecodeVt](me.Ppvt()).GetThumbnail)
}
