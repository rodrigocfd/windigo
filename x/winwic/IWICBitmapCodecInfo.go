//go:build windows

package winwic

import (
	"strings"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/cowic"
)

// [IWICBitmapCodecInfo] COM interface.
//
// [IWICBitmapCodecInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapcodecinfo
type IWICBitmapCodecInfo struct{ IWICComponentInfo }

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

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapCodecInfo) IID() *co.IID {
	return &cowic.IID_IWICBitmapCodecInfo
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IWICBitmapCodecInfo) AddRef(releaser *win.OleReleaser) *IWICBitmapCodecInfo {
	return utl.OleNewFromAddRef[*IWICBitmapCodecInfo](me, releaser)
}

// [DoesSupportAnimation] method.
//
// [DoesSupportAnimation]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-doessupportanimation
func (me *IWICBitmapCodecInfo) DoesSupportAnimation() (bool, error) {
	var supportAnim win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).DoesSupportAnimation,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&supportAnim)))
	return utl.HresultToBoolError(int32(supportAnim), ret)
}

// [DoesSupportChromakey] method.
//
// [DoesSupportChromakey]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-doessupportchromakey
func (me *IWICBitmapCodecInfo) DoesSupportChromakey() (bool, error) {
	var supportChr win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).DoesSupportChromakey,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&supportChr)))
	return utl.HresultToBoolError(int32(supportChr), ret)
}

// [DoesSupportLossless] method.
//
// [DoesSupportLossless]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-doessupportlossless
func (me *IWICBitmapCodecInfo) DoesSupportLossless() (bool, error) {
	var supportLossless win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).DoesSupportLossless,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&supportLossless)))
	return utl.HresultToBoolError(int32(supportLossless), ret)
}

// [DoesSupportMultiframe] method.
//
// [DoesSupportMultiframe]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-doessupportmultiframe
func (me *IWICBitmapCodecInfo) DoesSupportMultiframe() (bool, error) {
	var supportMulti win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).DoesSupportMultiframe,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&supportMulti)))
	return utl.HresultToBoolError(int32(supportMulti), ret)
}

// [GetColorManagementVersion] method.
//
// [GetColorManagementVersion]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-getcolormanagementversion
func (me *IWICBitmapCodecInfo) GetColorManagementVersion() (string, error) {
	return oleCallAllocBufRetStr(me, utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).GetColorManagementVersion)
}

// [GetContainerFormat] method.
//
// [GetContainerFormat]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-getcontainerformat
func (me *IWICBitmapCodecInfo) GetContainerFormat() (cowic.WIC_CONTAINER, error) {
	return utl.OleCallReturnStruct[cowic.WIC_CONTAINER](me,
		utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).GetContainerFormat)
}

// [GetDeviceManufacturer] method.
//
// [GetDeviceManufacturer]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-getdevicemanufacturer
func (me *IWICBitmapCodecInfo) GetDeviceManufacturer() (string, error) {
	return oleCallAllocBufRetStr(me, utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).GetDeviceManufacturer)
}

// [GetDeviceModels] method.
//
// [GetDeviceModels]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-getdevicemodels
func (me *IWICBitmapCodecInfo) GetDeviceModels() ([]string, error) {
	str, err := oleCallAllocBufRetStr(me, utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).GetDeviceModels)
	if err != nil {
		return nil, err
	}
	return strings.Split(str, ","), nil
}

// [GetFileExtensions] method.
//
// [GetFileExtensions]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-getfileextensions
func (me *IWICBitmapCodecInfo) GetFileExtensions() ([]string, error) {
	str, err := oleCallAllocBufRetStr(me, utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).GetFileExtensions)
	if err != nil {
		return nil, err
	}
	return strings.Split(str, ","), nil
}

// [GetMimeTypes] method.
//
// [GetMimeTypes]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-getmimetypes
func (me *IWICBitmapCodecInfo) GetMimeTypes() ([]string, error) {
	str, err := oleCallAllocBufRetStr(me, utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).GetMimeTypes)
	if err != nil {
		return nil, err
	}
	return strings.Split(str, ","), nil
}

// [GetPixelFormats] method.
//
// [GetPixelFormats]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-getpixelformats
func (me *IWICBitmapCodecInfo) GetPixelFormats() ([]cowic.WIC_PIXELFORMAT, error) {
	var numFormats uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).GetPixelFormats,
		me.Ppvt(),
		0, 0,
		uintptr(unsafe.Pointer(&numFormats)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}

	formats := make([]cowic.WIC_PIXELFORMAT, numFormats)
	ret, _, _ = syscall.SyscallN(
		utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).GetPixelFormats,
		me.Ppvt(),
		uintptr(numFormats),
		uintptr(unsafe.Pointer(&formats[0])),
		uintptr(unsafe.Pointer(&numFormats)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}
	return formats, nil
}

// [MatchesMimeType] method.
//
// [MatchesMimeType]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapcodecinfo-matchesmimetype
func (me *IWICBitmapCodecInfo) MatchesMimeType(mimeType string) (bool, error) {
	var wText wstr.BufEncoder
	var matches win.BOOL

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapCodecInfoVt](me.Ppvt()).MatchesMimeType,
		me.Ppvt(),
		uintptr(wText.AllowEmpty(mimeType)),
		uintptr(unsafe.Pointer(&matches)))
	return utl.HresultToBoolError(int32(matches), ret)
}
