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

// [IWICStream] COM interface.
//
// [IWICStream]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicstream
type IWICStream struct{ win.IStream }

type _IWICStreamVt struct {
	utl.IStreamVt
	InitializeFromIStream       uintptr
	InitializeFromFilename      uintptr
	InitializeFromMemory        uintptr
	InitializeFromIStreamRegion uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICStream) IID() *co.IID {
	return &cowic.IID_IWICStream
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IWICStream) AddRef(releaser *win.OleReleaser) *IWICStream {
	return utl.OleNewFromAddRef[*IWICStream](me, releaser)
}

// [InitializeFromFilename] method.
//
// [InitializeFromFilename]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicstream-initializefromfilename
func (me *IWICStream) InitializeFromFilename(fileName string, desiredAccess co.GENERIC) error {
	var wFileName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICStreamVt](me.Ppvt()).InitializeFromFilename,
		me.Ppvt(),
		uintptr(wFileName.AllowEmpty(fileName)),
		uintptr(desiredAccess))
	return utl.HresultToError(ret)
}

// [InitializeFromIStream] method.
//
// [InitializeFromIStream]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicstream-initializefromistream
func (me *IWICStream) InitializeFromIStream(stream *win.IStream) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICStreamVt](me.Ppvt()).InitializeFromIStream,
		me.Ppvt(),
		stream.Ppvt())
	return utl.HresultToError(ret)
}

// [InitializeFromIStreamRegion] method.
//
// [InitializeFromIStreamRegion]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicstream-initializefromistreamregion
func (me *IWICStream) InitializeFromIStreamRegion(stream *win.IStream, offset, maxSize int) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICStreamVt](me.Ppvt()).InitializeFromIStreamRegion,
		me.Ppvt(),
		stream.Ppvt(),
		uintptr(uint64(offset)),
		uintptr(uint64(maxSize)))
	return utl.HresultToError(ret)
}

// [InitializeFromMemory] method.
//
// [InitializeFromMemory]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicstream-initializefrommemory
func (me *IWICStream) InitializeFromMemory(buf []byte) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICStreamVt](me.Ppvt()).InitializeFromMemory,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(uint32(len(buf))))
	return utl.HresultToError(ret)
}
