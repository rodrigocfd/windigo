//go:build windows

package ole

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// [IUnknown] COM interface, base to all COM interfaces.
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknown struct {
	ppvt **vt.IUnknown
}

// Returns the [IUnknown] virtual table.
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
func (me *IUnknown) Ppvt() **vt.IUnknown {
	return me.ppvt
}

// Calls [Release], then sets a new [IUnknown] virtual table.
//
// If you pass nil, you effectively release the object; the owning ole.Releaser
// will simply do nothing.
//
// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
func (me *IUnknown) Set(ppvt **vt.IUnknown) {
	vt.Release(me.ppvt)
	me.ppvt = ppvt
}

// Returns the unique COM interface identifier.
func (*IUnknown) IID() co.IID {
	return co.IID_IUnknown
}

// [QueryInterface] method. Not implemented as a method of [IUnknown] because
// Go doesn't support generic methods.
//
// [QueryInterface]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
func QueryInterface[T any, P ComCtor[T]](
	iUnknown ComPtr,
	releaser *Releaser,
) (*T, error) {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	var ppvtQueried **vt.IUnknown
	riidGuid := win.GuidFrom(pObj.IID())

	ret, _, _ := syscall.SyscallN((*iUnknown.Ppvt()).QueryInterface,
		uintptr(unsafe.Pointer(iUnknown.Ppvt())),
		uintptr(unsafe.Pointer(&riidGuid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj.Set(ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}
