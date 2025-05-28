//go:build windows

package ole

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// A [COM] object, derived from [IUnknown].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type ComObj interface {
	ComResource

	// Returns the unique [COM] [interface ID].
	//
	// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
	// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
	IID() co.IID

	// Returns the [COM] virtual table pointer.
	//
	// This is a low-level method, used internally by the library. Incorrect usage
	// may lead to segmentation faults.
	//
	// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
	Ppvt() **IUnknownVt
}

// [IUnknown] [COM] interface, base to all COM interfaces.
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type IUnknown struct {
	ppvt **IUnknownVt
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

// [AddRef] method.
//
// The returned object must have the same type of the caller.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	var folder *shell.IShellItem
//	shell.SHCreateItemFromParsingName(rel, "C:\\Temp", &folder)
//
//	var folderCopy *shell.IShellItem
//	folder.QueryInterface(rel, &folderCopy)
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IUnknown) AddRef(releaser *Releaser, ppOut interface{}) {
	pOut := utl.ComValidateAndRetrievePointedToObj(ppOut).(ComObj)
	releaser.ReleaseNow(pOut)

	syscall.SyscallN((*me.Ppvt()).AddRef,
		uintptr(unsafe.Pointer(me.Ppvt())))

	pOut = utl.ComCreateObj(ppOut, unsafe.Pointer(me.ppvt)).(ComObj)
	releaser.Add(pOut)
}

// [Release] method. Implements [ComResource].
//
// You usually don't need to call this method directly, since every function
// which returns a [COM] object will require a [Releaser] to manage the object's
// lifetime.
//
// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (me *IUnknown) Release() {
	if me.ppvt != nil {
		syscall.SyscallN((*me.ppvt).Release,
			uintptr(unsafe.Pointer(me.ppvt)))
		me.ppvt = nil
	}
}

// [QueryInterface] method.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	var item *shell.IShellItem
//	shell.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
//	var item2 *shell.IShellItem2
//	item.QueryInterface(rel, &item2)
//
// [QueryInterface]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (me *IUnknown) QueryInterface(releaser *Releaser, ppOut interface{}) error {
	pOut := utl.ComValidateAndRetrievePointedToObj(ppOut).(ComObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **IUnknownVt
	guidIid := win.GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN((*me.Ppvt()).QueryInterface,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pOut = utl.ComCreateObj(ppOut, unsafe.Pointer(ppvtQueried)).(ComObj)
		releaser.Add(pOut)
		return nil
	} else {
		return hr
	}
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
