//go:build windows

package winwic

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/cowic"
)

// [IWICImagingFactory] COM interface.
//
// Implements [OleResource].
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
// [IWICImagingFactory]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicimagingfactory
type IWICImagingFactory struct{ win.IUnknown }

type _IWICImagingFactoryVt struct {
	utl.IUnknownVt
	CreateDecoderFromFilename                uintptr
	CreateDecoderFromStream                  uintptr
	CreateDecoderFromFileHandle              uintptr
	CreateComponentInfo                      uintptr
	CreateDecoder                            uintptr
	CreateEncoder                            uintptr
	CreatePalette                            uintptr
	CreateFormatConverter                    uintptr
	CreateBitmapScaler                       uintptr
	CreateBitmapClipper                      uintptr
	CreateBitmapFlipRotator                  uintptr
	CreateStream                             uintptr
	CreateColorContext                       uintptr
	CreateColorTransformer                   uintptr
	CreateBitmap                             uintptr
	CreateBitmapFromSource                   uintptr
	CreateBitmapFromSourceRect               uintptr
	CreateBitmapFromMemory                   uintptr
	CreateBitmapFromHBITMAP                  uintptr
	CreateBitmapFromHICON                    uintptr
	CreateComponentEnumerator                uintptr
	CreateFastMetadataEncoderFromDecoder     uintptr
	CreateFastMetadataEncoderFromFrameDecode uintptr
	CreateQueryWriter                        uintptr
	CreateQueryWriterFromReader              uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICImagingFactory) IID() *co.IID {
	return &cowic.IID_IWICImagingFactory
}

// [CreateBitmap] method.
//
// [CreateBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmap
func (me *IWICImagingFactory) CreateBitmap(
	releaser *win.OleReleaser,
	sz win.SIZE,
	pPixelFormat *cowic.WIC_PIXELFORMAT,
	option cowic.WICBMP_CACHE,
) (*IWICBitmap, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateBitmap,
		me.Ppvt(),
		uintptr(uint32(sz.Cx)),
		uintptr(uint32(sz.Cy)),
		uintptr(unsafe.Pointer(pPixelFormat)),
		uintptr(option),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmap](ret, ppvtQueried, releaser)
}

// [CreateBitmapFromHBITMAP] method.
//
// [CreateBitmapFromHBITMAP]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmapfromhbitmap
func (me *IWICImagingFactory) CreateBitmapFromHBITMAP(
	releaser *win.OleReleaser,
	hBmp win.HBITMAP,
	hPal win.HPALETTE,
	options cowic.WICBMP_ALPHACH,
) (*IWICBitmap, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateBitmapFromHBITMAP,
		me.Ppvt(),
		uintptr(hBmp),
		uintptr(hPal),
		uintptr(options),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmap](ret, ppvtQueried, releaser)
}

// [CreateBitmapFromHICON] method.
//
// [CreateBitmapFromHICON]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmapfromhicon
func (me *IWICImagingFactory) CreateBitmapFromHICON(
	releaser *win.OleReleaser,
	hIcon win.HICON,
) (*IWICBitmap, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateBitmapFromHICON,
		me.Ppvt(),
		uintptr(hIcon),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmap](ret, ppvtQueried, releaser)
}

// [CreateBitmapFromMemory] method.
//
// [CreateBitmapFromMemory]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmapfrommemory
func (me *IWICImagingFactory) CreateBitmapFromMemory(
	releaser *win.OleReleaser,
	newBmpSz win.SIZE,
	pPixelFormat *cowic.WIC_PIXELFORMAT,
	stride int,
	srcBuf []byte,
) (*IWICBitmap, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateBitmapFromMemory,
		me.Ppvt(),
		uintptr(uint32(newBmpSz.Cx)),
		uintptr(uint32(newBmpSz.Cy)),
		uintptr(unsafe.Pointer(pPixelFormat)),
		uintptr(uint32(stride)),
		uintptr(uint32(len(srcBuf))),
		uintptr(unsafe.Pointer(&srcBuf[0])),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmap](ret, ppvtQueried, releaser)
}

// [CreateBitmapFromSource] method.
//
// [CreateBitmapFromSource]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmapfromsource
func (me *IWICImagingFactory) CreateBitmapFromSource(
	releaser *win.OleReleaser,
	source *IWICBitmapSource,
	option cowic.WICBMP_CACHE,
) (*IWICBitmap, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateBitmapFromSource,
		me.Ppvt(),
		uintptr(unsafe.Pointer(source.Ppvt())),
		uintptr(option),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmap](ret, ppvtQueried, releaser)
}

// [CreateBitmapFromSourceRect] method.
//
// [CreateBitmapFromSourceRect]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmapfromsourcerect
func (me *IWICImagingFactory) CreateBitmapFromSourceRect(
	releaser *win.OleReleaser,
	source *IWICBitmapSource,
	pos win.POINT,
	size win.SIZE,
) (*IWICBitmap, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateBitmapFromSourceRect,
		me.Ppvt(),
		uintptr(unsafe.Pointer(source.Ppvt())),
		uintptr(uint32(pos.X)),
		uintptr(uint32(pos.Y)),
		uintptr(uint32(size.Cx)),
		uintptr(uint32(size.Cy)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmap](ret, ppvtQueried, releaser)
}

// [CreateComponentEnumerator] method.
//
// [CreateComponentEnumerator]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createcomponentenumerator
func (me *IWICImagingFactory) CreateComponentEnumerator(
	releaser *win.OleReleaser,
	componentTypes cowic.WIC_COMPONENTTYPE,
	options cowic.WIC_COMPONENTENUM,
) (*win.IEnumUnknown, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateComponentEnumerator,
		me.Ppvt(),
		uintptr(componentTypes),
		uintptr(options),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*win.IEnumUnknown](ret, ppvtQueried, releaser)
}

// [CreateComponentInfo] method.
//
// [CreateComponentInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createcomponentinfo
func (me *IWICImagingFactory) CreateComponentInfo(
	releaser *win.OleReleaser,
	pClsidComponent *co.CLSID,
) (*IWICComponentInfo, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateComponentInfo,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pClsidComponent)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICComponentInfo](ret, ppvtQueried, releaser)
}

// [CreateDecoder] method.
//
// [CreateDecoder]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createdecoder
func (me *IWICImagingFactory) CreateDecoder(
	releaser *win.OleReleaser,
	pGuidContainerFormat *cowic.WIC_CONTAINER,
	pGuidVendor *cowic.WIC_VENDOR,
) (*IWICBitmapDecoder, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateDecoder,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pGuidContainerFormat)),
		uintptr(unsafe.Pointer(pGuidVendor)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmapDecoder](ret, ppvtQueried, releaser)
}

// [CreateDecoderFromFileHandle] method.
//
// [CreateDecoderFromFileHandle]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createdecoderfromfilehandle
func (me *IWICImagingFactory) CreateDecoderFromFileHandle(
	releaser *win.OleReleaser,
	hFile win.HFILE,
	pGuidVendor *co.GUID,
	metadataOpts cowic.WICDEC_METADATACACHE,
) (*IWICBitmapDecoder, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateDecoderFromFileHandle,
		me.Ppvt(),
		uintptr(hFile),
		uintptr(unsafe.Pointer(pGuidVendor)),
		uintptr(metadataOpts),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmapDecoder](ret, ppvtQueried, releaser)
}

// [CreateDecoderFromFilename] method.
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
//		nil,
//		co.GENERIC_READ,
//		cowic.WICDEC_METADATACACHE_OnDemand,
//	)
//
// [CreateDecoderFromFilename]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createdecoderfromfilename
func (me *IWICImagingFactory) CreateDecoderFromFilename(
	releaser *win.OleReleaser,
	filename string,
	pGuidVendor *co.GUID,
	desiredAccess co.GENERIC,
	metadataOpts cowic.WICDEC_METADATACACHE,
) (*IWICBitmapDecoder, error) {
	var ppvtQueried uintptr
	var wFilename wstr.BufEncoder

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateDecoderFromFilename,
		me.Ppvt(),
		uintptr(wFilename.AllowEmpty(filename)),
		uintptr(unsafe.Pointer(pGuidVendor)),
		uintptr(desiredAccess),
		uintptr(metadataOpts),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmapDecoder](ret, ppvtQueried, releaser)
}

// [CreateDecoderFromStream] method.
//
// [CreateDecoderFromStream]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createdecoderfromstream
func (me *IWICImagingFactory) CreateDecoderFromStream(
	releaser *win.OleReleaser,
	stream *win.IStream,
	pGuidVendor *co.GUID,
	metadataOpts cowic.WICDEC_METADATACACHE,
) (*IWICBitmapDecoder, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateDecoderFromStream,
		me.Ppvt(),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		uintptr(unsafe.Pointer(pGuidVendor)),
		uintptr(metadataOpts),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmapDecoder](ret, ppvtQueried, releaser)
}

// [CreateEncoder] method.
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
// [CreateEncoder]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createencoder
func (me *IWICImagingFactory) CreateEncoder(
	releaser *win.OleReleaser,
	pGuidContainer *cowic.WIC_CONTAINER,
	pGuidVendor *co.GUID,
) (*IWICBitmapEncoder, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateEncoder,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pGuidContainer)),
		uintptr(unsafe.Pointer(pGuidVendor)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IWICBitmapEncoder](ret, ppvtQueried, releaser)
}

// [CreateFormatConverter] method.
//
// [CreateFormatConverter]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createformatconverter
func (me *IWICImagingFactory) CreateFormatConverter(releaser *win.OleReleaser) (*IWICFormatConverter, error) {
	return utl.OleNewFromCallWithoutParms[*IWICFormatConverter](me, releaser,
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateFormatConverter)
}

// [CreatePalette] method.
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
// [CreatePalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createpalette
func (me *IWICImagingFactory) CreatePalette(releaser *win.OleReleaser) (*IWICPalette, error) {
	return utl.OleNewFromCallWithoutParms[*IWICPalette](me, releaser,
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreatePalette)
}

// [CreateStream] method.
//
// [CreateStream]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createstream
func (me *IWICImagingFactory) CreateStream(releaser *win.OleReleaser) (*IWICStream, error) {
	return utl.OleNewFromCallWithoutParms[*IWICStream](me, releaser,
		utl.Vt[_IWICImagingFactoryVt](me.Ppvt()).CreateStream)
}
