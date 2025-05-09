//go:build windows

package ole

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
)

// Stores multiple COM pointers and VARIANT values, allowing releasing the
// resources all at once:
//
//   - COM pointers will call [Release];
//   - VARIANT objects will call [VariantClear].
//
// Every function or method which creates a COM pointer or VARIANT will require
// a Releaser to manage the lifetime of that COM pointer or VARIANT.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
// [VariantClear]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-variantclear
type Releaser struct {
	ptrs []ComPtr
	vars []*VARIANT
}

// Stores multiple COM pointers and VARIANT values, allowing releasing the
// resources all at once:
//
//   - COM pointers will call [Release];
//   - VARIANT objects will call [VariantClear].
//
// Every function or method which creates a COM pointer or VARIANT will require
// a Releaser to manage the lifetime of that COM pointer or VARIANT.
//
// ⚠️ You must defer Releaser.Release().
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
// [VariantClear]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-variantclear
func NewReleaser() *Releaser {
	return new(Releaser)
}

// Adds new COM pointers to be released with [Release].
//
// Usually, this method is called automatically by functions which return COM
// objects.
//
// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (me *Releaser) Add(ptrs ...ComPtr) {
	me.ptrs = append(me.ptrs, ptrs...)
}

// Adds new VARIANT objects to be released with [VariantClear].
//
// Usually, this method is called automatically by functions which return
// VARIANT objects.
//
// [VariantClear]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-variantclear
func (me *Releaser) AddVar(va ...*VARIANT) {
	me.vars = append(me.vars, va...)
}

// Calls [Release] on each stored COM pointer, in the reverse order they were
// added. Then frees the internal slice.
//
// Calls [VariantClear] on each stored VARIANT object, in the reverse order they
// were added. Then frees the internal slice.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
// [VariantClear]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-variantclear
func (me *Releaser) Release() {
	for i := len(me.ptrs) - 1; i >= 0; i-- {
		me.ptrs[i].Set(nil)
	}
	me.ptrs = nil

	for i := len(me.vars) - 1; i >= 0; i-- {
		syscall.SyscallN(_VariantClear.Addr(),
			uintptr(unsafe.Pointer(me.vars[i]))) // ignore errors
	}
	me.vars = nil
}

var _VariantClear = dll.Oleaut32.NewProc("VariantClear")
