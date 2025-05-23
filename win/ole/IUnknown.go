//go:build windows

package ole

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

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
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IUnknown) AddRef(releaser *Releaser, ppOut interface{}) {

	// TODO: validate ppOut as receiver type

	utl.ComValidateOutPtr(ppOut)

	syscall.SyscallN((*me.Ppvt()).AddRef,
		uintptr(unsafe.Pointer(me.Ppvt())))

	utl.ComCreateObj(ppOut, unsafe.Pointer(me.ppvt))
	releaser.Add(reflect.ValueOf(ppOut).Elem().Interface().(ComResource))
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

// [IUnknown] [COM] virtual table, base to all COM virtual tables.
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type IUnknownVt struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}
