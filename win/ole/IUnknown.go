//go:build windows

package ole

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// [IUnknown] [COM] interface, base to all COM interfaces.
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type IUnknown struct {
	ppvt **IUnknownVt
}

// Calls [Release].
//
// You usually don't need to call this method directly, since every function
// which returns a [COM] object will require a [Releaser] to manage the object's
// lifetime.
//
// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func (me *IUnknown) Release() {
	me.Set(nil)
}

// Returns the unique [COM] [interface ID].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IUnknown) IID() co.IID {
	return co.IID_IUnknown
}

// Returns the [COM] virtual table pointer.
//
// This is a low-level method, used internally by the library. Incorrect usage
// may lead to segmentation faults.
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func (me *IUnknown) Ppvt() **IUnknownVt {
	return me.ppvt
}

// Calls [Release], then sets a new [COM] virtual table pointer.
//
// If you pass nil, you effectively release the object; the owning ole.Releaser
// will simply do nothing.
//
// This is a low-level method, used internally by the library. Incorrect usage
// may lead to segmentation faults.
//
// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func (me *IUnknown) Set(ppvt **IUnknownVt) {
	if me.ppvt != nil {
		syscall.SyscallN((*me.ppvt).Release,
			uintptr(unsafe.Pointer(me.ppvt)))
	}
	me.ppvt = ppvt
}

// [AddRef] method. Not implemented as a method of [IUnknown] because Go doesn't
// support generic methods.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func AddRef[T any, P ComCtor[T]](iUnknown P, releaser *Releaser) *T {
	syscall.SyscallN((*iUnknown.Ppvt()).AddRef,
		uintptr(unsafe.Pointer(iUnknown.Ppvt())))

	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	pObj.Set(iUnknown.Ppvt())
	releaser.Add(pObj)
	return pObj
}

// [QueryInterface] method. Not implemented as a method of [IUnknown] because
// Go doesn't support generic methods.
//
// [QueryInterface]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func QueryInterface[T any, P ComCtor[T]](
	iUnknown ComPtr,
	releaser *Releaser,
) (*T, error) {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	var ppvtQueried **IUnknownVt
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

// Syntactic sugar to create a new [COM] object from its virtual table.
//
// This is a low-level method, used internally by the library. Incorrect usage
// may lead to segmentation faults.
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func ComObj[T any, P interface {
	*T
	Set(**IUnknownVt)
}](ppvt **IUnknownVt) *T {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	pObj.Set(ppvt)
	return pObj
}

// [IUnknown] [COM] virtual table, base to all COM virtual tables.
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type IUnknownVt struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}
