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

// [IWICBitmapLock] COM interface.
//
// [IWICBitmapLock]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwicbitmaplock
type IWICBitmapLock struct{ win.IUnknown }

type _IWICBitmapLockVt struct {
	utl.IUnknownVt
	GetSize        uintptr
	GetStride      uintptr
	GetDataPointer uintptr
	GetPixelFormat uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICBitmapLock) IID() *co.IID {
	return &cowic.IID_IWICBitmapLock
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IWICBitmapLock) AddRef(releaser *win.OleReleaser) *IWICBitmapLock {
	return utl.OleNewFromAddRef[*IWICBitmapLock](me, releaser)
}

// [GetDataPointer] method.
//
// [GetDataPointer]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmaplock-getdatapointer
func (me *IWICBitmapLock) GetDataPointer() ([]byte, error) {
	var szBuf uint32
	var pData *byte

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapLockVt](me.Ppvt()).GetDataPointer,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&szBuf)),
		uintptr(unsafe.Pointer(&pData)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}
	return unsafe.Slice(pData, szBuf), nil
}

// [GetSize] method.
//
// [GetSize]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmaplock-getsize
func (me *IWICBitmapLock) GetSize() (win.SIZE, error) {
	var cx, cy uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapLockVt](me.Ppvt()).GetSize,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&cx)),
		uintptr(unsafe.Pointer(&cy)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return win.SIZE{}, hr
	}
	return win.SIZE{Cx: int32(cx), Cy: int32(cy)}, nil
}

// [GetStride] method.
//
// [GetStride]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicbitmaplock-getstride
func (me *IWICBitmapLock) GetStride() (int, error) {
	var stride uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICBitmapLockVt](me.Ppvt()).GetStride,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&stride)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(stride), nil
}
