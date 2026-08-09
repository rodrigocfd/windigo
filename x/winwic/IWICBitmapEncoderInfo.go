//go:build windows

package winwic

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cowic"
)

// [IWICBitmapEncoderInfo] COM interface.
//
// [IWICBitmapEncoderInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmapencoderinfo
type IWICBitmapEncoderInfo struct{ IWICBitmapCodecInfo }

type _IWICBitmapEncoderInfoVt struct {
	_IWICBitmapCodecInfoVt
	CreateInstance uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapEncoderInfo) IID() *co.IID {
	return &cowic.IID_IWICBitmapEncoderInfo
}

// [CreateInstance] method.
//
// [CreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmapencoderinfo-createinstance
func (me *IWICBitmapEncoderInfo) CreateInstance(releaser *win.OleReleaser) (*IWICBitmapEncoder, error) {
	return utl.OleNewFromCallWithoutParms[*IWICBitmapEncoder](me, releaser,
		utl.Vt[_IWICBitmapEncoderInfoVt](me.Ppvt()).CreateInstance)
}
