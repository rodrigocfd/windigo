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

// [IWICBitmapDecoderInfo] COM interface.
//
// Implements [OleResource].
//
// [IWICBitmapDecoderInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapdecoderinfo
type IWICBitmapDecoderInfo struct{ IWICBitmapCodecInfo }

type _IWICBitmapDecoderInfoVt struct {
	_IWICBitmapCodecInfoVt
	GetPatterns    uintptr
	MatchesPattern uintptr
	CreateInstance uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapDecoderInfo) IID() *co.IID {
	return &cowic.IID_IWICBitmapDecoderInfo
}

// [CreateInstance] method.
//
// [CreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoderinfo-createinstance
func (me *IWICBitmapDecoderInfo) CreateInstance(releaser *win.OleReleaser) (*IWICBitmapDecoder, error) {
	return utl.OleNewFromCallWithoutParms[*IWICBitmapDecoder](me, releaser,
		utl.Vt[_IWICBitmapDecoderInfoVt](me.Ppvt()).CreateInstance)
}

// [MatchesPattern] method.
//
// [MatchesPattern]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapdecoderinfo-matchespattern
func (me *IWICBitmapDecoderInfo) MatchesPattern(stream *win.IStream) (bool, error) {
	var matches win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapDecoderInfoVt](me.Ppvt()).MatchesPattern,
		me.Ppvt(),
		stream.Ppvt(),
		uintptr(unsafe.Pointer(&matches)))
	return utl.HresultToBoolError(int32(matches), ret)
}
