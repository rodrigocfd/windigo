//go:build windows

package windxgi

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [IDXGIResource1] COM interface.
//
// [IDXGIResource1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nn-dxgi1_2-idxgiresource1
type IDXGIResource1 struct{ IDXGIResource }

type _IDXGIResource1Vt struct {
	_IDXGIResourceVt
	CreateSubresourceSurface uintptr
	CreateSharedHandle       uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIResource1) IID() *co.IID {
	return &codxgi.IID_IDXGIResource1
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIResource1) AddRef(releaser *win.OleReleaser) *IDXGIResource1 {
	return utl.OleNewFromAddRef[*IDXGIResource1](me, releaser)
}

// [CreateSharedHandle] method.
//
// [CreateSharedHandle]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiresource1-createsharedhandle
func (me *IDXGIResource1) CreateSharedHandle(
	pSecurityAttributes *win.SECURITY_ATTRIBUTES,
	access codxgi.DXGI_SHARED,
	name string,
) (win.HANDLE, error) {
	var wName wstr.BufEncoder
	var handle win.HANDLE

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIResource1Vt](me.Ppvt()).CreateSharedHandle,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pSecurityAttributes)),
		uintptr(access),
		uintptr(wName.EmptyIsNil(name)),
		uintptr(unsafe.Pointer(&handle)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return win.HANDLE(0), hr
	}
	return handle, nil
}

// // [CreateSubresourceSurface] method.
// //
// // [CreateSubresourceSurface]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiresource1-createsubresourcesurface
// func (me *IDXGIResource1) CreateSubresourceSurface(
// 	releaser *win.OleReleaser,
// 	index int,
// ) (*IDXGISurface2, error) {
// 	var ppvtQueried uintptr
// 	ret, _, _ := syscall.SyscallN(
// 		utl.Vt[_IDXGIResource1Vt](me.Ppvt()).CreateSubresourceSurface,
// 		me.Ppvt(),
// 		uintptr(uint32(index)),
// 		uintptr(unsafe.Pointer(&ppvtQueried)))
// 	return utl.OleNewIfOk[*IDXGISurface2](ret, ppvtQueried, releaser)
// }
