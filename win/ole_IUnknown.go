//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [IUnknown] [COM] interface, base to all COM interfaces.
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type IUnknown struct {
	ppvt uintptr
}

// Returns the unique [COM] [interface ID].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IUnknown) IID() *co.IID {
	return &co.IID_IUnknown
}

// Returns the [COM] virtual table pointer.
//
// This is a low-level method, used internally by the library. Incorrect usage
// may lead to segmentation faults.
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func (me *IUnknown) Ppvt() uintptr {
	return me.ppvt
}

// [AddRef] method.
//
// Essentially clones the COM object, so its lifetime can be managed by
// different [OleReleaser].
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var folder *winsh.IShellItem
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp", &folder)
//
//	rel2 := win.NewOleReleaser()
//	defer rel2.Release()
//
//	folderClone := folder.AddRef(rel2)
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IUnknown) AddRef(releaser *OleReleaser) *IUnknown {
	return utl.OleNewFromAddRef[*IUnknown](me, releaser)
}

// [QueryInterface] method.
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var item *winsh.IShellItem
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
//	var item2 *winsh.IShellItem2
//	_ = item.QueryInterface(rel, &item2)
//
// [QueryInterface]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (me *IUnknown) QueryInterface(releaser *OleReleaser, ppOut interface{}) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[utl.IUnknownVt](me.ppvt).QueryInterface,
		me.ppvt,
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

// [Release] method.
//
// You usually don't need to call this method, it's called automatically by
// [OleReleaser].
//
// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (me *IUnknown) Release() {
	if me.ppvt != 0 {
		_, _, _ = syscall.SyscallN(
			utl.Vt[utl.IUnknownVt](me.ppvt).Release,
			me.ppvt)
		me.ppvt = 0
	}
}

// Stores multiple [COM] resources, releasing all them at once.
//
// Every function which returns a COM resource will require an [OleReleaser]
// to manage the object's lifetime.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type OleReleaser struct {
	r utl.OleBatchReleaser
}

// Constructs a new [OleReleaser] to store multiple [COM] resources, releasing
// them all at once.
//
// Every function which returns a COM resource will require an [OleReleaser] to
// manage the object's lifetime.
//
// ⚠️ You must defer [OleReleaser.Release].
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func NewOleReleaser() *OleReleaser {
	return &OleReleaser{
		r: utl.NewOleBatchReleaser(),
	}
}

// Adds a new [COM] resource to have its lifetime managed by the [OleReleaser].
func (me *OleReleaser) Add(obj interface{ Release() }) {
	me.r.Add(obj)
}

// Releases all added [COM] resource, in the reverse order they were added.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func (me *OleReleaser) Release() {
	me.r.Release()
}
