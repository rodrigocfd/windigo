//go:build windows

package windxgi

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [IDXGISurface2] COM interface.
//
// [IDXGISurface2]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nn-dxgi1_2-idxgisurface2
type IDXGISurface2 struct{ IDXGISurface1 }

type _IDXGISurface2Vt struct {
	_IDXGISurface1Vt
	GetResource uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGISurface2) IID() *co.IID {
	return &codxgi.IID_IDXGISurface2
}

// [GetResource] method.
//
// [GetResource]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgisurface2-getresource
func (me *IDXGISurface2) GetResource(
	releaser *win.OleReleaser,
	pGuid *co.GUID,
	ppOut interface{},
) (int, error) {
	_ = utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr
	var subresourceIndex uint32

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISurface2Vt](me.Ppvt()).GetResource,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pGuid)),
		uintptr(unsafe.Pointer(&ppvtQueried)),
		uintptr(unsafe.Pointer(&subresourceIndex)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	utl.OleInject(ppOut, ppvtQueried, releaser)
	return int(subresourceIndex), nil
}
