//go:build windows

package windxgi

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [IDXGISurface] COM interface.
//
// Implements [OleResource].
//
// [IDXGISurface]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgisurface
type IDXGISurface struct{ IDXGIDeviceSubObject }

type _IDXGISurfaceVt struct {
	_IDXGIDeviceSubObjectVt
	GetDesc uintptr
	Map     uintptr
	Unmap   uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGISurface) IID() *co.IID {
	return &codxgi.IID_IDXGISurface
}

// [GetDesc] method.
//
// [GetDesc]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgisurface-getdesc
func (me *IDXGISurface) GetDesc() (DXGI_SURFACE_DESC, error) {
	return utl.OleCallReturnStruct[DXGI_SURFACE_DESC](me,
		utl.Vt[_IDXGISurfaceVt](me.Ppvt()).GetDesc)
}

// [Map] method.
//
// [Map]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgisurface-map
func (me *IDXGISurface) Map(flags codxgi.DXGI_MAP) (DXGI_MAPPED_RECT, error) {
	var lockedRect DXGI_MAPPED_RECT
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISurfaceVt](me.Ppvt()).Map,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&lockedRect)),
		uintptr(flags))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return lockedRect, nil
	} else {
		return DXGI_MAPPED_RECT{}, hr
	}
}

// [Unmap] method.
//
// [Unmap]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgisurface-unmap
func (me *IDXGISurface) Unmap() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IDXGISurfaceVt](me.Ppvt()).Unmap)
}
