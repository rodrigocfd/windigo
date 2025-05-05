//go:build windows

package vt

import (
	"syscall"
	"unsafe"
)

// Syntactic sugar to create a new COM object and set its virtual table pointer.
func NewObj[T any, P interface {
	*T
	Set(**IUnknown)
}](ppvt **IUnknown) *T {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	pObj.Set(ppvt)
	return pObj
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func AddRef(ppvt **IUnknown) uint32 {
	refCount, _, _ := syscall.SyscallN((*ppvt).AddRef,
		uintptr(unsafe.Pointer(ppvt)))
	return uint32(refCount)
}

// [Release] method.
//
// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func Release(ppvt **IUnknown) uint32 {
	var refCount uintptr
	if ppvt != nil {
		refCount, _, _ = syscall.SyscallN((*ppvt).Release,
			uintptr(unsafe.Pointer(ppvt)))
	}
	return uint32(refCount)
}
