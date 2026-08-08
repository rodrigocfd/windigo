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

// [IDXGIObject] COM interface.
//
// Implements [OleResource].
//
// [IDXGIObject]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgiobject
type IDXGIObject struct{ win.IUnknown }

type _IDXGIObjectVt struct {
	utl.IUnknownVt
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIObject) IID() *co.IID {
	return &codxgi.IID_IDXGIObject
}

// [GetParent] method.
//
// [GetParent]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiobject-getparent
func (me *IDXGIObject) GetParent(releaser *win.OleReleaser, ppOut interface{}) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIObjectVt](me.Ppvt()).GetParent,
		me.Ppvt(),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

// [GetPrivateData] method.
//
// [GetPrivateData]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiobject-getprivatedata
func (me *IDXGIObject) GetPrivateData(pName *co.GUID, szData int, pData unsafe.Pointer) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIObjectVt](me.Ppvt()).GetPrivateData,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pName)),
		uintptr(uint32(szData)),
		uintptr(pData))
	return utl.HresultToError(ret)
}

// [GetPrivateData] method, specialized to return an [IUnknown]-derived object.
//
// [GetPrivateData]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiobject-getprivatedata
func (me *IDXGIObject) GetPrivateDataInterface(
	releaser *win.OleReleaser,
	pName *co.GUID,
	ppOut interface{},
) error {
	utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	if err := me.GetPrivateData(pName, int(unsafe.Sizeof(uintptr(0))), unsafe.Pointer(&ppvtQueried)); err != nil {
		return err
	}
	utl.OleInject(ppOut, ppvtQueried, releaser)
	return nil
}

// [SetPrivateData] method.
//
// [SetPrivateData]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiobject-setprivatedata
func (me *IDXGIObject) SetPrivateData(pName *co.GUID, szData int, pData unsafe.Pointer) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIObjectVt](me.Ppvt()).SetPrivateData,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pName)),
		uintptr(uint32(szData)),
		uintptr(pData))
	return utl.HresultToError(ret)
}

// [SetPrivateDataInterface] method.
//
// [SetPrivateDataInterface]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiobject-setprivatedatainterface
func (me *IDXGIObject) SetPrivateDataInterface(pName *co.GUID, obj *win.IUnknown) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIObjectVt](me.Ppvt()).SetPrivateDataInterface,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(obj.Ppvt())))
	return utl.HresultToError(ret)
}
