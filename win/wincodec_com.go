//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [IWICBitmap] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IWICBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmap
type IWICBitmap struct{ IWICBitmapSource }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmap) IID() co.IID {
	return co.IID_IWICBitmap
}

// [Lock] method.
//
// [Lock]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmap-lock
func (me *IWICBitmap) Lock(releaser *OleReleaser, lock *WICRect, flags co.WICBMP_LOCK) (*IWICBitmapLock, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapVt)(unsafe.Pointer(*me.Ppvt())).Lock,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(lock)),
		uintptr(flags),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapLock{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [SetPalette] method.
//
// [SetPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmap-setpalette
func (me *IWICBitmap) SetPalette(palette *IWICPalette) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapVt)(unsafe.Pointer(*me.Ppvt())).SetPalette,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(palette.Ppvt())))
	return utl.HresultToError(ret)
}

// [SetResolution] method.
//
// [SetResolution]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmap-setresolution
func (me *IWICBitmap) SetResolution(dpiX, dpiY float64) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapVt)(unsafe.Pointer(*me.Ppvt())).SetResolution,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(dpiX),
		uintptr(dpiY))
	return utl.HresultToError(ret)
}

type _IWICBitmapVt struct {
	_IUnknownVt
	Lock          uintptr
	SetPalette    uintptr
	SetResolution uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICBitmapCodecInfo] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IWICBitmapCodecInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapcodecinfo
type IWICBitmapCodecInfo struct{ IWICComponentInfo }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapCodecInfo) IID() co.IID {
	return co.IID_IWICBitmapCodecInfo
}

// [DoesSupportAnimation] method.
//
// [DoesSupportAnimation]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-doessupportanimation
func (me *IWICBitmapCodecInfo) DoesSupportAnimation() (bool, error) {
	var supportAnim BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapCodecInfoVt)(unsafe.Pointer(*me.Ppvt())).DoesSupportAnimation,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&supportAnim)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return supportAnim.Get(), nil
	} else {
		return false, hr
	}
}

// [DoesSupportChromakey] method.
//
// [DoesSupportChromakey]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-doessupportchromakey
func (me *IWICBitmapCodecInfo) DoesSupportChromakey() (bool, error) {
	var supportChr BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapCodecInfoVt)(unsafe.Pointer(*me.Ppvt())).DoesSupportChromakey,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&supportChr)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return supportChr.Get(), nil
	} else {
		return false, hr
	}
}

// [DoesSupportLossless] method.
//
// [DoesSupportLossless]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-doessupportlossless
func (me *IWICBitmapCodecInfo) DoesSupportLossless() (bool, error) {
	var supportLossless BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapCodecInfoVt)(unsafe.Pointer(*me.Ppvt())).DoesSupportLossless,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&supportLossless)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return supportLossless.Get(), nil
	} else {
		return false, hr
	}
}

// [DoesSupportMultiframe] method.
//
// [DoesSupportMultiframe]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-doessupportmultiframe
func (me *IWICBitmapCodecInfo) DoesSupportMultiframe() (bool, error) {
	var supportMulti BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapCodecInfoVt)(unsafe.Pointer(*me.Ppvt())).DoesSupportMultiframe,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&supportMulti)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return supportMulti.Get(), nil
	} else {
		return false, hr
	}
}

// [GetContainerFormat] method.
//
// [GetContainerFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-getcontainerformat
func (me *IWICBitmapCodecInfo) GetContainerFormat() (co.WIC_CONTAINER, error) {
	var guid GUID
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapCodecInfoVt)(unsafe.Pointer(*me.Ppvt())).GetContainerFormat,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.WIC_CONTAINER(guid.String()), nil
	} else {
		return "", hr
	}
}

// [GetPixelFormats] method.
//
// [GetPixelFormats]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-getpixelformats
func (me *IWICBitmapCodecInfo) GetPixelFormats() ([]co.WIC_PIXELFORMAT, error) {
	var numFormats uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapCodecInfoVt)(unsafe.Pointer(*me.Ppvt())).GetPixelFormats,
		uintptr(unsafe.Pointer(me.Ppvt())),
		0, 0,
		uintptr(unsafe.Pointer(&numFormats)))

	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}

	formatGuids := make([]GUID, numFormats)
	ret, _, _ = syscall.SyscallN(
		(*_IWICBitmapCodecInfoVt)(unsafe.Pointer(*me.Ppvt())).GetPixelFormats,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(numFormats),
		uintptr(unsafe.Pointer(unsafe.SliceData(formatGuids))),
		uintptr(unsafe.Pointer(&numFormats)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		formats := make([]co.WIC_PIXELFORMAT, 0, numFormats)
		for _, guid := range formatGuids {
			formats = append(formats, co.WIC_PIXELFORMAT(guid.String()))
		}
		return formats, nil
	} else {
		return nil, hr
	}
}

// [MatchesMimeType] method.
//
// [MatchesMimeType]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-matchesmimetype
func (me *IWICBitmapCodecInfo) MatchesMimeType(mimeType string) (bool, error) {
	var wText wstr.BufEncoder
	var matches BOOL

	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapCodecInfoVt)(unsafe.Pointer(*me.Ppvt())).MatchesMimeType,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wText.AllowEmpty(mimeType)),
		uintptr(unsafe.Pointer(&matches)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return matches.Get(), nil
	} else {
		return false, hr
	}
}

type _IWICBitmapCodecInfoVt struct {
	_IWICComponentInfoVt
	GetContainerFormat        uintptr
	GetPixelFormats           uintptr
	GetColorManagementVersion uintptr
	GetDeviceManufacturer     uintptr
	GetDeviceModels           uintptr
	GetMimeTypes              uintptr
	GetFileExtensions         uintptr
	DoesSupportAnimation      uintptr
	DoesSupportChromakey      uintptr
	DoesSupportLossless       uintptr
	DoesSupportMultiframe     uintptr
	MatchesMimeType           uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICBitmapDecoder] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var factory *win.IWICImagingFactory
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_WICImagingFactory,
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
//		co.WICDECMETADATACACHE_OnDemand,
//	)
//
// [IWICBitmapDecoder]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapdecoder
type IWICBitmapDecoder struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapDecoder) IID() co.IID {
	return co.IID_IWICBitmapDecoder
}

// [CopyPalette] method.
//
// [CopyPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-copypalette
func (me *IWICBitmapDecoder) CopyPalette(palette *IWICPalette) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapDecoderVt)(unsafe.Pointer(*me.Ppvt())).CopyPalette,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(palette.Ppvt())))
	return utl.HresultToError(ret)
}

// [GetContainerFormat] method.
//
// [GetContainerFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-getcontainerformat
func (me *IWICBitmapDecoder) GetContainerFormat() (co.WIC_CONTAINER, error) {
	var guid GUID
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapDecoderVt)(unsafe.Pointer(*me.Ppvt())).GetContainerFormat,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.WIC_CONTAINER(guid.String()), nil
	} else {
		return "", hr
	}
}

// [GetDecoderInfo] method.
//
// [GetDecoderInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-getdecoderinfo
func (me *IWICBitmapDecoder) GetDecoderInfo(releaser *OleReleaser) (*IWICBitmapDecoderInfo, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapDecoderVt)(unsafe.Pointer(*me.Ppvt())).GetDecoderInfo,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapDecoderInfo{IWICBitmapCodecInfo{IWICComponentInfo{IUnknown{ppvtQueried}}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [GetFrame] method.
//
// [GetFrame]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-getframe
func (me *IWICBitmapDecoder) GetFrame(releaser *OleReleaser, index int) (*IWICBitmapFrameDecode, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapDecoderVt)(unsafe.Pointer(*me.Ppvt())).GetFrame,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapFrameDecode{IWICBitmapSource{IUnknown{ppvtQueried}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [GetFrameCount] method.
//
// [GetFrameCount]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-getframecount
func (me *IWICBitmapDecoder) GetFrameCount() (int, error) {
	var count uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapDecoderVt)(unsafe.Pointer(*me.Ppvt())).CopyPalette,
		uintptr(unsafe.Pointer(me.Ppvt())),
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
func (me *IWICBitmapDecoder) GetPreview(releaser *OleReleaser) (*IWICBitmapSource, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapDecoderVt)(unsafe.Pointer(*me.Ppvt())).GetPreview,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapSource{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [GetThumbnail] method.
//
// [GetThumbnail]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-getthumbnail
func (me *IWICBitmapDecoder) GetThumbnail(releaser *OleReleaser) (*IWICBitmapSource, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapDecoderVt)(unsafe.Pointer(*me.Ppvt())).GetThumbnail,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapSource{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [Initialize] method.
//
// [Initialize]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-initialize
func (me *IWICBitmapDecoder) Initialize(stream *IStream, cacheOpts co.WICDEC_METADATACACHE) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapDecoderVt)(unsafe.Pointer(*me.Ppvt())).Initialize,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		uintptr(cacheOpts))
	return utl.HresultToError(ret)
}

// [QueryCapability] method.
//
// [QueryCapability]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoder-querycapability
func (me *IWICBitmapDecoder) QueryCapability(stream *IStream) (co.WICDEC_CAP, error) {
	var capability uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapDecoderVt)(unsafe.Pointer(*me.Ppvt())).QueryCapability,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		uintptr(unsafe.Pointer(&capability)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.WICDEC_CAP(capability), nil
	} else {
		return co.WICDEC_CAP(0), hr
	}
}

type _IWICBitmapDecoderVt struct {
	_IUnknownVt
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

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICBitmapDecoderInfo] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IWICBitmapDecoderInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapdecoderinfo
type IWICBitmapDecoderInfo struct{ IWICBitmapCodecInfo }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapDecoderInfo) IID() co.IID {
	return co.IID_IWICBitmapDecoderInfo
}

// [CreateInstance] method.
//
// [CreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoderinfo-createinstance
func (me *IWICBitmapDecoderInfo) CreateInstance(releaser *OleReleaser) (*IWICBitmapDecoder, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapDecoderInfoVt)(unsafe.Pointer(*me.Ppvt())).CreateInstance,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapDecoder{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [MatchesPattern] method.
//
// [MatchesPattern]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoderinfo-matchespattern
func (me *IWICBitmapDecoderInfo) MatchesPattern(stream *IStream) (bool, error) {
	var matches BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IWICFormatConverterVt)(unsafe.Pointer(*me.Ppvt())).CanConvert,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		uintptr(unsafe.Pointer(&matches)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return matches.Get(), nil
	} else {
		return false, hr
	}
}

type _IWICBitmapDecoderInfoVt struct {
	_IWICBitmapCodecInfoVt
	GetPatterns    uintptr
	MatchesPattern uintptr
	CreateInstance uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICBitmapEncoder] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var factory *win.IWICImagingFactory
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_WICImagingFactory,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&factory,
//	)
//
//	encoder, _ := factory.CreateEncoder(rel, co.WICCONTAINER_Bmp, "")
//
// [IWICBitmapEncoder]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapencoder
type IWICBitmapEncoder struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapEncoder) IID() co.IID {
	return co.IID_IWICBitmapEncoder
}

// [Commit] method.
//
// [Commit]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-commit
func (me *IWICBitmapEncoder) Commit() error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapEncoderVt)(unsafe.Pointer(*me.Ppvt())).Commit,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.HresultToError(ret)
}

// [GetContainerFormat] method.
//
// [GetContainerFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-getcontainerformat
func (me *IWICBitmapEncoder) GetContainerFormat() (co.WIC_CONTAINER, error) {
	var guid GUID
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapEncoderVt)(unsafe.Pointer(*me.Ppvt())).GetContainerFormat,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.WIC_CONTAINER(guid.String()), nil
	} else {
		return "", hr
	}
}

// [GetEncoderInfo] method.
//
// [GetEncoderInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-getencoderinfo
func (me *IWICBitmapEncoder) GetEncoderInfo(releaser *OleReleaser) (*IWICBitmapEncoderInfo, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapEncoderVt)(unsafe.Pointer(*me.Ppvt())).GetEncoderInfo,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapEncoderInfo{IWICBitmapCodecInfo{IWICComponentInfo{IUnknown{ppvtQueried}}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [Initialize] method.
//
// [Initialize]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-initialize
func (me *IWICBitmapEncoder) Initialize(stream *IStream, cacheOpt co.WICENC_CACHE) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapEncoderVt)(unsafe.Pointer(*me.Ppvt())).Initialize,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		uintptr(cacheOpt))
	return utl.HresultToError(ret)
}

// [SetPalette] method.
//
// [SetPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-setpalette
func (me *IWICBitmapEncoder) SetPalette(palette *IWICPalette) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapEncoderVt)(unsafe.Pointer(*me.Ppvt())).SetPalette,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(palette.Ppvt())))
	return utl.HresultToError(ret)
}

// [SetPreview] method.
//
// [SetPreview]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoder-setpreview
func (me *IWICBitmapEncoder) SetPreview(preview *IWICBitmapSource) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapEncoderVt)(unsafe.Pointer(*me.Ppvt())).SetPreview,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(preview.Ppvt())))
	return utl.HresultToError(ret)
}

// [SetThumbnail] method.
//
// [SetThumbnail]:
func (me *IWICBitmapEncoder) SetThumbnail(thumbnail *IWICBitmapSource) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapEncoderVt)(unsafe.Pointer(*me.Ppvt())).SetThumbnail,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(thumbnail.Ppvt())))
	return utl.HresultToError(ret)
}

type _IWICBitmapEncoderVt struct {
	_IUnknownVt
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

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICBitmapEncoderInfo] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IWICBitmapEncoderInfo]:
type IWICBitmapEncoderInfo struct{ IWICBitmapCodecInfo }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapEncoderInfo) IID() co.IID {
	return co.IID_IWICBitmapEncoderInfo
}

// [CreateInstance] method.
//
// [CreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoderinfo-createinstance
func (me *IWICBitmapEncoderInfo) CreateInstance(releaser *OleReleaser) (*IWICBitmapEncoder, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapEncoderInfoVt)(unsafe.Pointer(*me.Ppvt())).CreateInstance,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapEncoder{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

type _IWICBitmapEncoderInfoVt struct {
	_IWICBitmapCodecInfoVt
	CreateInstance uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICBitmapFrameDecode] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var factory *win.IWICImagingFactory
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_WICImagingFactory,
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
//		co.WICDECMETADATACACHE_OnDemand,
//	)
//
//	frame, _ := decoder.GetFrame(rel, 0)
//
// [IWICBitmapFrameDecode]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapframedecode
type IWICBitmapFrameDecode struct{ IWICBitmapSource }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapFrameDecode) IID() co.IID {
	return co.IID_IWICBitmapFrameDecode
}

// [GetThumbnail] method.
//
// [GetThumbnail]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapframedecode-getthumbnail
func (me *IWICBitmapFrameDecode) GetThumbnail(releaser *OleReleaser) (*IWICBitmapSource, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapFrameDecodeVt)(unsafe.Pointer(*me.Ppvt())).GetThumbnail,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapSource{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

type _IWICBitmapFrameDecodeVt struct {
	_IWICBitmapSourceVt
	GetMetadataQueryReader uintptr
	GetColorContexts       uintptr
	GetThumbnail           uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICBitmapLock] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IWICBitmapLock]:
type IWICBitmapLock struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapLock) IID() co.IID {
	return co.IID_IWICBitmapLock
}

// [GetDataPointer] method.
//
// [GetDataPointer]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmaplock-getdatapointer
func (me *IWICBitmapLock) GetDataPointer() ([]byte, error) {
	var szBuf uint32
	var pData *byte

	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapLockVt)(unsafe.Pointer(*me.Ppvt())).GetDataPointer,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&szBuf)),
		uintptr(unsafe.Pointer(&pData)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return unsafe.Slice(pData, szBuf), nil
	} else {
		return nil, hr
	}
}

// [GetSize] method.
//
// [GetSize]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmaplock-getsize
func (me *IWICBitmapLock) GetSize() (SIZE, error) {
	var cx, cy uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapLockVt)(unsafe.Pointer(*me.Ppvt())).GetSize,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&cx)),
		uintptr(unsafe.Pointer(&cy)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return SIZE{int32(cx), int32(cy)}, nil
	} else {
		return SIZE{}, hr
	}
}

// [GetStride] method.
//
// [GetStride]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmaplock-getstride
func (me *IWICBitmapLock) GetStride() (int, error) {
	var stride uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapLockVt)(unsafe.Pointer(*me.Ppvt())).GetStride,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&stride)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(stride), nil
	} else {
		return 0, hr
	}
}

type _IWICBitmapLockVt struct {
	_IUnknownVt
	GetSize        uintptr
	GetStride      uintptr
	GetDataPointer uintptr
	GetPixelFormat uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICBitmapSource] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IWICBitmapSource]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapsource
type IWICBitmapSource struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapSource) IID() co.IID {
	return co.IID_IWICBitmapSource
}

// [CopyPalette] method.
//
// [CopyPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapsource-copypalette
func (me *IWICBitmapSource) CopyPalette(palette *IWICPalette) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapSourceVt)(unsafe.Pointer(*me.Ppvt())).CopyPalette,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(palette.Ppvt())))
	return utl.HresultToError(ret)
}

// [CopyPixels] method.
//
// [CopyPixels]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapsource-copypixels
func (me *IWICBitmapSource) CopyPixels(rc *WICRect, stride, szBuffer int, pBuffer *byte) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapSourceVt)(unsafe.Pointer(*me.Ppvt())).CopyPixels,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(rc)),
		uintptr(uint32(stride)),
		uintptr(uint32(szBuffer)),
		uintptr(unsafe.Pointer(pBuffer)))
	return utl.HresultToError(ret)
}

// [GetPixelFormat] method.
//
// [GetPixelFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapsource-getpixelformat
func (me *IWICBitmapSource) GetPixelFormat() (co.WIC_PIXELFORMAT, error) {
	var guid GUID
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapSourceVt)(unsafe.Pointer(*me.Ppvt())).GetPixelFormat,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.WIC_PIXELFORMAT(guid.String()), nil
	} else {
		return "", hr
	}
}

// [GetResolution] method.
//
// [GetResolution]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapsource-getresolution
func (me *IWICBitmapSource) GetResolution() (float64, float64, error) {
	var cx, cy float64
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapSourceVt)(unsafe.Pointer(*me.Ppvt())).GetResolution,
		uintptr(unsafe.Pointer(me.Ppvt())),
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
func (me *IWICBitmapSource) GetSize() (SIZE, error) {
	var cx, cy uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICBitmapSourceVt)(unsafe.Pointer(*me.Ppvt())).GetSize,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&cx)),
		uintptr(unsafe.Pointer(&cy)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return SIZE{int32(cx), int32(cy)}, nil
	} else {
		return SIZE{}, hr
	}
}

type _IWICBitmapSourceVt struct {
	_IUnknownVt
	GetSize        uintptr
	GetPixelFormat uintptr
	GetResolution  uintptr
	CopyPalette    uintptr
	CopyPixels     uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICComponentInfo] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IWICComponentInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwiccomponentinfo
type IWICComponentInfo struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICComponentInfo) IID() co.IID {
	return co.IID_IWICComponentInfo
}

// [GetAuthor] method.
//
// [GetAuthor]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getauthor
func (me *IWICComponentInfo) GetAuthor() (string, error) {
	var szRequired uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICComponentInfoVt)(unsafe.Pointer(*me.Ppvt())).GetAuthor,
		uintptr(unsafe.Pointer(me.Ppvt())),
		0, 0,
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	buf := make([]uint16, szRequired)
	ret, _, _ = syscall.SyscallN(
		(*_IWICComponentInfoVt)(unsafe.Pointer(*me.Ppvt())).GetAuthor,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(szRequired),
		uintptr(unsafe.Pointer(unsafe.SliceData(buf))),
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.DecodeSlice(buf), nil
	} else {
		return "", hr
	}
}

// [GetCLSID] method.
//
// [GetCLSID]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getclsid
func (me *IWICComponentInfo) GetCLSID() (co.CLSID, error) {
	var guid GUID
	ret, _, _ := syscall.SyscallN(
		(*_IWICComponentInfoVt)(unsafe.Pointer(*me.Ppvt())).GetCLSID,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.CLSID(guid.String()), nil
	} else {
		return "", hr
	}
}

// [GetFriendlyName] method.
//
// [GetFriendlyName]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getfriendlyname
func (me *IWICComponentInfo) GetFriendlyName() (string, error) {
	var szRequired uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICComponentInfoVt)(unsafe.Pointer(*me.Ppvt())).GetFriendlyName,
		uintptr(unsafe.Pointer(me.Ppvt())),
		0, 0,
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	buf := make([]uint16, szRequired)
	ret, _, _ = syscall.SyscallN(
		(*_IWICComponentInfoVt)(unsafe.Pointer(*me.Ppvt())).GetFriendlyName,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(szRequired),
		uintptr(unsafe.Pointer(unsafe.SliceData(buf))),
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.DecodeSlice(buf), nil
	} else {
		return "", hr
	}
}

// [GetSigningStatus] method.
//
// [GetSigningStatus]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getsigningstatus
func (me *IWICComponentInfo) GetSigningStatus() (co.WIC_COMPONENTSIGN, error) {
	var status uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICComponentInfoVt)(unsafe.Pointer(*me.Ppvt())).GetSigningStatus,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&status)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.WIC_COMPONENTSIGN(status), nil
	} else {
		return co.WIC_COMPONENTSIGN(0), hr
	}
}

// [GetSpecVersion] method.
//
// [GetSpecVersion]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getspecversion
func (me *IWICComponentInfo) GetSpecVersion() (string, error) {
	var szRequired uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICComponentInfoVt)(unsafe.Pointer(*me.Ppvt())).GetSpecVersion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		0, 0,
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	buf := make([]uint16, szRequired)
	ret, _, _ = syscall.SyscallN(
		(*_IWICComponentInfoVt)(unsafe.Pointer(*me.Ppvt())).GetSpecVersion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(szRequired),
		uintptr(unsafe.Pointer(unsafe.SliceData(buf))),
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.DecodeSlice(buf), nil
	} else {
		return "", hr
	}
}

// [GetVendorGUID] method.
//
// [GetVendorGUID]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getvendorguid
func (me *IWICComponentInfo) GetVendorGUID() (co.GUID, error) {
	var guid GUID
	ret, _, _ := syscall.SyscallN(
		(*_IWICComponentInfoVt)(unsafe.Pointer(*me.Ppvt())).GetVendorGUID,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.GUID(guid.String()), nil
	} else {
		return "", hr
	}
}

// [GetVersion] method.
//
// [GetVersion]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getversion
func (me *IWICComponentInfo) GetVersion() (string, error) {
	var szRequired uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICComponentInfoVt)(unsafe.Pointer(*me.Ppvt())).GetVersion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		0, 0,
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	buf := make([]uint16, szRequired)
	ret, _, _ = syscall.SyscallN(
		(*_IWICComponentInfoVt)(unsafe.Pointer(*me.Ppvt())).GetVersion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(szRequired),
		uintptr(unsafe.Pointer(unsafe.SliceData(buf))),
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.DecodeSlice(buf), nil
	} else {
		return "", hr
	}
}

type _IWICComponentInfoVt struct {
	_IUnknownVt
	GetComponentType uintptr
	GetCLSID         uintptr
	GetSigningStatus uintptr
	GetAuthor        uintptr
	GetVendorGUID    uintptr
	GetVersion       uintptr
	GetSpecVersion   uintptr
	GetFriendlyName  uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICFormatConverter] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IWICFormatConverter]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicformatconverter
type IWICFormatConverter struct{ IWICBitmapSource }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICFormatConverter) IID() co.IID {
	return co.IID_IWICFormatConverter
}

// [CanConvert] method.
//
// [CanConvert]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicformatconverter-canconvert
func (me *IWICFormatConverter) CanConvert(srcPixelFormat, destPixelFormat co.WIC_PIXELFORMAT) (bool, error) {
	guidSrc := GuidFrom(srcPixelFormat)
	guidDest := GuidFrom(destPixelFormat)
	var canConvert BOOL

	ret, _, _ := syscall.SyscallN(
		(*_IWICFormatConverterVt)(unsafe.Pointer(*me.Ppvt())).CanConvert,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidSrc)),
		uintptr(unsafe.Pointer(&guidDest)),
		uintptr(unsafe.Pointer(&canConvert)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return canConvert.Get(), nil
	} else {
		return false, hr
	}
}

// [Initialize] method.
//
// [Initialize]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicformatconverter-initialize
func (me *IWICFormatConverter) Initialize(
	source *IWICBitmapSource,
	destFormat co.WIC_PIXELFORMAT,
	dither co.WICBMP_DITHER,
	palette *IWICPalette,
	alphaThresholdPercent float64,
	paletteTranslate co.WICBMP_PAL,
) error {
	guidDestFormat := GuidFrom(destFormat)
	ret, _, _ := syscall.SyscallN(
		(*_IWICFormatConverterVt)(unsafe.Pointer(*me.Ppvt())).Initialize,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(source.Ppvt())),
		uintptr(unsafe.Pointer(&guidDestFormat)),
		uintptr(unsafe.Pointer(&dither)),
		uintptr(unsafe.Pointer(palette.Ppvt())),
		uintptr(alphaThresholdPercent),
		uintptr(paletteTranslate))
	return utl.HresultToError(ret)
}

type _IWICFormatConverterVt struct {
	_IWICBitmapSourceVt
	Initialize uintptr
	CanConvert uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICImagingFactory] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	var factory *win.IWICImagingFactory
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_WICImagingFactory,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&factory,
//	)
//
// [IWICImagingFactory]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicimagingfactory
type IWICImagingFactory struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICImagingFactory) IID() co.IID {
	return co.IID_IWICImagingFactory
}

// [CreateBitmap] method.
//
// [CreateBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmap
func (me *IWICImagingFactory) CreateBitmap(
	releaser *OleReleaser,
	sz SIZE,
	pixelFormat co.WIC_PIXELFORMAT,
	option co.WICBMP_CACHE,
) (*IWICBitmap, error) {
	var ppvtQueried **_IUnknownVt
	guidPixelFormat := GuidFrom(pixelFormat)

	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateBitmap,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(sz.Cx)),
		uintptr(uint32(sz.Cy)),
		uintptr(unsafe.Pointer(&guidPixelFormat)),
		uintptr(option),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmap{IWICBitmapSource{IUnknown{ppvtQueried}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [CreateBitmapFromHBITMAP] method.
//
// [CreateBitmapFromHBITMAP]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmapfromhbitmap
func (me *IWICImagingFactory) CreateBitmapFromHBITMAP(
	releaser *OleReleaser,
	hBmp HBITMAP,
	hPal HPALETTE,
	options co.WICBMP_ALPHACH,
) (*IWICBitmap, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateBitmapFromHBITMAP,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hBmp),
		uintptr(hPal),
		uintptr(options),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmap{IWICBitmapSource{IUnknown{ppvtQueried}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [CreateBitmapFromHICON] method.
//
// [CreateBitmapFromHICON]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmapfromhicon
func (me *IWICImagingFactory) CreateBitmapFromHICON(releaser *OleReleaser, hIcon HICON) (*IWICBitmap, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateBitmapFromHICON,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hIcon),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmap{IWICBitmapSource{IUnknown{ppvtQueried}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [CreateBitmapFromMemory] method.
//
// [CreateBitmapFromMemory]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmapfrommemory
func (me *IWICImagingFactory) CreateBitmapFromMemory(
	releaser *OleReleaser,
	newBmpSz SIZE,
	pixelFormat co.WIC_PIXELFORMAT,
	stride int,
	srcBuf []byte,
) (*IWICBitmap, error) {
	var ppvtQueried **_IUnknownVt
	guidPixelFormat := GuidFrom(pixelFormat)

	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateBitmapFromMemory,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(newBmpSz.Cx)),
		uintptr(uint32(newBmpSz.Cy)),
		uintptr(unsafe.Pointer(&guidPixelFormat)),
		uintptr(uint32(stride)),
		uintptr(uint32(len(srcBuf))),
		uintptr(unsafe.Pointer(unsafe.SliceData(srcBuf))),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmap{IWICBitmapSource{IUnknown{ppvtQueried}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [CreateBitmapFromSource] method.
//
// [CreateBitmapFromSource]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmapfromsource
func (me *IWICImagingFactory) CreateBitmapFromSource(
	releaser *OleReleaser,
	source *IWICBitmapSource,
	option co.WICBMP_CACHE,
) (*IWICBitmap, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateBitmapFromSource,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(source.Ppvt())),
		uintptr(option),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmap{IWICBitmapSource{IUnknown{ppvtQueried}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [CreateBitmapFromSourceRect] method.
//
// [CreateBitmapFromSourceRect]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createbitmapfromsourcerect
func (me *IWICImagingFactory) CreateBitmapFromSourceRect(
	releaser *OleReleaser,
	source *IWICBitmapSource,
	pos POINT,
	size SIZE,
) (*IWICBitmap, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateBitmapFromSourceRect,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(source.Ppvt())),
		uintptr(uint32(pos.X)),
		uintptr(uint32(pos.Y)),
		uintptr(uint32(size.Cx)),
		uintptr(uint32(size.Cy)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmap{IWICBitmapSource{IUnknown{ppvtQueried}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [CreateDecoderFromFileHandle] method.
//
// For a null guidVendor, pass an empty string.
//
// [CreateDecoderFromFileHandle]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createdecoderfromfilehandle
func (me *IWICImagingFactory) CreateDecoderFromFileHandle(
	releaser *OleReleaser,
	hFile HFILE,
	guidVendor co.GUID,
	metadataOpts co.WICDEC_METADATACACHE,
) (*IWICBitmapDecoder, error) {
	var ppvtQueried **_IUnknownVt
	var guidGuidVendor GUID

	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateDecoderFromFileHandle,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hFile),
		uintptr(guidStructPtrOrNil(guidVendor, &guidGuidVendor)),
		uintptr(metadataOpts),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapDecoder{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [CreateDecoderFromFilename] method.
//
// For a null guidVendor, pass an empty string.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var factory *win.IWICImagingFactory
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_WICImagingFactory,
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
//		co.WICDECMETADATACACHE_OnDemand,
//	)
//
// [CreateDecoderFromFilename]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createdecoderfromfilename
func (me *IWICImagingFactory) CreateDecoderFromFilename(
	releaser *OleReleaser,
	filename string,
	guidVendor co.GUID,
	desiredAccess co.GENERIC,
	metadataOpts co.WICDEC_METADATACACHE,
) (*IWICBitmapDecoder, error) {
	var ppvtQueried **_IUnknownVt
	var wFilename wstr.BufEncoder
	var guidGuidVendor GUID

	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateDecoderFromFilename,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wFilename.AllowEmpty(filename)),
		uintptr(guidStructPtrOrNil(guidVendor, &guidGuidVendor)),
		uintptr(desiredAccess),
		uintptr(metadataOpts),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapDecoder{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [CreateDecoderFromStream] method.
//
// For a null guidVendor, pass an empty string.
//
// [CreateDecoderFromStream]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createdecoderfromstream
func (me *IWICImagingFactory) CreateDecoderFromStream(
	releaser *OleReleaser,
	stream *IStream,
	guidVendor co.GUID,
	metadataOpts co.WICDEC_METADATACACHE,
) (*IWICBitmapDecoder, error) {
	var ppvtQueried **_IUnknownVt
	var guidGuidVendor GUID

	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateDecoderFromStream,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		uintptr(guidStructPtrOrNil(guidVendor, &guidGuidVendor)),
		uintptr(metadataOpts),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapDecoder{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [CreateEncoder] method.
//
// For a null guidVendor, pass an empty string.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var factory *win.IWICImagingFactory
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_WICImagingFactory,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&factory,
//	)
//
//	encoder, _ := factory.CreateEncoder(rel, co.WICCONTAINER_Bmp, "")
//
// [CreateEncoder]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createencoder
func (me *IWICImagingFactory) CreateEncoder(
	releaser *OleReleaser,
	guidContainer co.WIC_CONTAINER,
	guidVendor co.GUID,
) (*IWICBitmapEncoder, error) {
	var ppvtQueried **_IUnknownVt
	guidGuidContainer := GuidFrom(guidContainer)
	var guidGuidVendor GUID

	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateEncoder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidGuidContainer)),
		uintptr(guidStructPtrOrNil(guidVendor, &guidGuidVendor)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICBitmapEncoder{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [CreateFormatConverter] method.
//
// [CreateFormatConverter]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createformatconverter
func (me *IWICImagingFactory) CreateFormatConverter(releaser *OleReleaser) (*IWICFormatConverter, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateFormatConverter,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICFormatConverter{IWICBitmapSource{IUnknown{ppvtQueried}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
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
//	var factory *win.IWICImagingFactory
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_WICImagingFactory,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&factory,
//	)
//
//	palette, _ := factory.CreatePalette(rel)
//
// [CreatePalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createpalette
func (me *IWICImagingFactory) CreatePalette(releaser *OleReleaser) (*IWICPalette, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreatePalette,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICPalette{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [CreateStream] method.
//
// [CreateStream]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createstream
func (me *IWICImagingFactory) CreateStream(releaser *OleReleaser) (*IWICStream, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IWICImagingFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateStream,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IWICStream{IStream{ISequentialStream{IUnknown{ppvtQueried}}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

type _IWICImagingFactoryVt struct {
	_IUnknownVt
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

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICPalette] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	var factory *win.IWICImagingFactory
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_WICImagingFactory,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&factory,
//	)
//
//	palette, _ := factory.CreatePalette(rel)
//
// [IWICPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicpalette
type IWICPalette struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICPalette) IID() co.IID {
	return co.IID_IWICPalette
}

// [GetColorCount] method.
//
// [GetColorCount]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicpalette-getcolorcount
func (me *IWICPalette) GetColorCount() (int, error) {
	var count uint32
	ret, _, _ := syscall.SyscallN(
		(*_IWICPaletteVt)(unsafe.Pointer(*me.Ppvt())).GetColorCount,
		uintptr(unsafe.Pointer(me.Ppvt())),
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
	var hasAlpha BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IWICPaletteVt)(unsafe.Pointer(*me.Ppvt())).HasAlpha,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hasAlpha)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hasAlpha.Get(), nil
	} else {
		return false, hr
	}
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
		(*_IWICPaletteVt)(unsafe.Pointer(*me.Ppvt())).InitializeFromBitmap,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(surface.Ppvt())),
		uintptr(uint32(numColors)),
		utl.BoolToUintptr(addTransparentColor))
	return utl.HresultToError(ret)
}

// [InitializeFromPalette] method.
//
// [InitializeFromPalette]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicpalette-initializefrompalette
func (me *IWICPalette) InitializeFromPalette(palette *IWICPalette) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICPaletteVt)(unsafe.Pointer(*me.Ppvt())).InitializeFromPalette,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(palette.Ppvt())))
	return utl.HresultToError(ret)
}

// [IsBlackWhite] method.
//
// [IsBlackWhite]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicpalette-isblackwhite
func (me *IWICPalette) IsBlackWhite() (bool, error) {
	var isBW BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IWICPaletteVt)(unsafe.Pointer(*me.Ppvt())).IsBlackWhite,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&isBW)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return isBW.Get(), nil
	} else {
		return false, hr
	}
}

// [IsGrayscale] method.
//
// [IsGrayscale]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicpalette-isgrayscale
func (me *IWICPalette) IsGrayscale() (bool, error) {
	var isGrayscale BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IWICPaletteVt)(unsafe.Pointer(*me.Ppvt())).IsGrayscale,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&isGrayscale)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return isGrayscale.Get(), nil
	} else {
		return false, hr
	}
}

type _IWICPaletteVt struct {
	_IUnknownVt
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

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IWICStream] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IWICStream]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicstream
type IWICStream struct{ IStream }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICStream) IID() co.IID {
	return co.IID_IWICStream
}

// [InitializeFromFilename] method.
//
// [InitializeFromFilename]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicstream-initializefromfilename
func (me *IWICStream) InitializeFromFilename(fileName string, desiredAccess co.GENERIC) error {
	var wFileName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IWICStreamVt)(unsafe.Pointer(*me.Ppvt())).InitializeFromFilename,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wFileName.AllowEmpty(fileName)),
		uintptr(desiredAccess))
	return utl.HresultToError(ret)
}

// [InitializeFromIStream] method.
//
// [InitializeFromIStream]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicstream-initializefromistream
func (me *IWICStream) InitializeFromIStream(stream *IStream) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICStreamVt)(unsafe.Pointer(*me.Ppvt())).InitializeFromIStream,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(stream.Ppvt())))
	return utl.HresultToError(ret)
}

// [InitializeFromIStreamRegion] method.
//
// [InitializeFromIStreamRegion]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicstream-initializefromistreamregion
func (me *IWICStream) InitializeFromIStreamRegion(stream *IStream, offset, maxSize int) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICStreamVt)(unsafe.Pointer(*me.Ppvt())).InitializeFromIStreamRegion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		uintptr(uint64(offset)),
		uintptr(uint64(maxSize)))
	return utl.HresultToError(ret)
}

// [InitializeFromMemory] method.
//
// [InitializeFromMemory]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicstream-initializefrommemory
func (me *IWICStream) InitializeFromMemory(buf []byte) error {
	ret, _, _ := syscall.SyscallN(
		(*_IWICStreamVt)(unsafe.Pointer(*me.Ppvt())).InitializeFromMemory,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(unsafe.SliceData(buf))),
		uintptr(uint32(len(buf))))
	return utl.HresultToError(ret)
}

type _IWICStreamVt struct {
	_IStreamVt
	InitializeFromIStream       uintptr
	InitializeFromFilename      uintptr
	InitializeFromMemory        uintptr
	InitializeFromIStreamRegion uintptr
}
