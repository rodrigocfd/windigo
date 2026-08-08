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

// [IDXGIDeviceSubObject] COM interface.
//
// Implements [OleResource].
//
// [IDXGIDeviceSubObject]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgidevicesubobject
type IDXGIDeviceSubObject struct{ IDXGIObject }

type _IDXGIDeviceSubObjectVt struct {
	_IDXGIObjectVt
	GetDevice uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIDeviceSubObject) IID() *co.IID {
	return &codxgi.IID_IDXGIDeviceSubObject
}

// [GetDevice] method.
//
// [GetDevice]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgidevicesubobject-getdevice
func (me *IDXGIDeviceSubObject) GetDevice(releaser *win.OleReleaser, ppOut interface{}) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIDeviceSubObjectVt](me.Ppvt()).GetDevice,
		me.Ppvt(),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}
