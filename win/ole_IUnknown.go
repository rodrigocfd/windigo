//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
)

// [IUnknown] [COM] interface, base to all COM interfaces.
//
// Implements [OleObj] and [OleResource].
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type IUnknown struct {
	ppvt **_IUnknownVt
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
func (me *IUnknown) Ppvt() **_IUnknownVt {
	return me.ppvt
}

// [AddRef] method.
//
// The returned object must have the same type of the caller.
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
//	var folder *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp", &folder)
//
//	var folderCopy *win.IShellItem
//	folder.AddRef(rel, &folderCopy)
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IUnknown) AddRef(releaser *OleReleaser, ppOut interface{}) {
	com_validateAndRelease(ppOut, releaser)
	_, _, _ = syscall.SyscallN(
		(*me.Ppvt()).AddRef,
		uintptr(unsafe.Pointer(me.Ppvt())))
	com_buildObj(ppOut, me.ppvt, releaser)
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
//	var item *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
//	var item2 *win.IShellItem2
//	_ = item.QueryInterface(rel, &item2)
//
// [QueryInterface]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (me *IUnknown) QueryInterface(releaser *OleReleaser, ppOut interface{}) error {
	iid := com_validateAndRelease(ppOut, releaser)
	var ppvtQueried **_IUnknownVt
	guidIid := GuidFrom(iid)

	ret, _, _ := syscall.SyscallN(
		(*me.Ppvt()).QueryInterface,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retHres(ret, ppOut, ppvtQueried, releaser)
}

// Implements [OleResource].
func (me *IUnknown) release() {
	if me.ppvt != nil {
		_, _, _ = syscall.SyscallN(
			(*me.ppvt).Release,
			uintptr(unsafe.Pointer(me.ppvt)))
		me.ppvt = nil
	}
}

// [IUnknown] [COM] virtual table, base to all COM virtual tables.
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type _IUnknownVt struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

// IUnknown.QueryInterface method for custom-implemented interfaces.
var _iunknownQueryInterfaceImpl uintptr

func com_iunknownQueryInterfaceImpl() uintptr {
	if _iunknownQueryInterfaceImpl == 0 {
		_iunknownQueryInterfaceImpl = syscall.NewCallback(
			func(_p uintptr, _riid uintptr, ppv ***_IUnknownVt) uintptr {
				*ppv = nil
				return uintptr(co.HRESULT_E_NOTIMPL)
			},
		)
	}
	return _iunknownQueryInterfaceImpl
}
