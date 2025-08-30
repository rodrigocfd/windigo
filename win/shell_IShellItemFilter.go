//go:build windows

package win

import (
	"sync/atomic"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [IShellItemFilter] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IShellItemFilter]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitemfilter
type IShellItemFilter struct{ IUnknown }

// Returns the unique [COM] [interface ID].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItemFilter) IID() co.IID {
	return co.IID_IShellItemFilter
}

type _IShellItemFilterImpl struct {
	vt                  _IShellItemFilterVt
	counter             uint32
	includeItem         func(item *IShellItem) co.HRESULT
	getEnumFlagsForItem func(item *IShellItem, flags *co.SHCONTF) co.HRESULT
}

// Implements [IShellItemFilter].
//
// [IShellItemFilter]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitemfilter
func NewIShellItemFilterImpl(releaser *OleReleaser) *IShellItemFilter {
	_iShellItemFilterVtPtrs.init()
	pImpl := &_IShellItemFilterImpl{ // has Go function pointers, so cannot be allocated on the OS heap
		vt:      _iShellItemFilterVtPtrs, // simply copy the syscall callback pointers
		counter: 1,
	}
	utl.PtrCache.Add(unsafe.Pointer(pImpl)) // keep ptr
	ppImpl := &pImpl
	utl.PtrCache.Add(unsafe.Pointer(ppImpl)) // also keep ptr ptr

	ppFakeVtbl := (**_IUnknownVt)(unsafe.Pointer(ppImpl))
	pObj := &IShellItemFilter{IUnknown{ppFakeVtbl}}
	releaser.Add(pObj)
	return pObj
}

// Defines [GetEnumFlagsForItem] method.
//
// [GetEnumFlagsForItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemfilter-getenumflagsforitem
func (me *IShellItemFilter) GetEnumFlagsForItem(fun func(item *IShellItem, flags *co.SHCONTF) co.HRESULT) {
	(*(**_IShellItemFilterImpl)(unsafe.Pointer(me.Ppvt()))).getEnumFlagsForItem = fun
}

// Defines [IncludeItem] method.
//
// [IncludeItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemfilter-includeitem
func (me *IShellItemFilter) IncludeItem(fun func(item *IShellItem) co.HRESULT) {
	(*(**_IShellItemFilterImpl)(unsafe.Pointer(me.Ppvt()))).includeItem = fun
}

type _IShellItemFilterVt struct {
	_IUnknownVt
	IncludeItem         uintptr
	GetEnumFlagsForItem uintptr
}

var _iShellItemFilterVtPtrs _IShellItemFilterVt // Global to keep the syscall callback pointers.

func (me *_IShellItemFilterVt) init() {
	if me.QueryInterface != 0 {
		return
	}

	*me = _IShellItemFilterVt{
		_IUnknownVt: _IUnknownVt{
			QueryInterface: syscall.NewCallback(
				func(_p uintptr, _riid uintptr, ppv ***_IUnknownVt) uintptr {
					*ppv = nil
					return uintptr(co.HRESULT_E_NOTIMPL)
				},
			),
			AddRef: syscall.NewCallback(
				func(p uintptr) uintptr {
					ppImpl := (**_IDropTargetImpl)(unsafe.Pointer(p))
					newCount := atomic.AddUint32(&(**ppImpl).counter, 1)
					return uintptr(newCount)
				},
			),
			Release: syscall.NewCallback(
				func(p uintptr) uintptr {
					ppImpl := (**_IDropTargetImpl)(unsafe.Pointer(p))
					newCount := atomic.AddUint32(&(*ppImpl).counter, ^uint32(0)) // decrement 1
					if newCount == 0 {
						utl.PtrCache.Delete(unsafe.Pointer(*ppImpl)) // now GC can collect them
						utl.PtrCache.Delete(unsafe.Pointer(ppImpl))
					}
					return uintptr(newCount)
				},
			),
		},
		IncludeItem: syscall.NewCallback(
			func(p uintptr, psi **_IUnknownVt) uintptr {
				ppImpl := (**_IShellItemFilterImpl)(unsafe.Pointer(p))
				if fun := (**ppImpl).includeItem; fun == nil { // user didn't define a callback
					return uintptr(co.HRESULT_S_OK)
				} else {
					return uintptr(fun(&IShellItem{IUnknown{psi}}))
				}
			},
		),
		GetEnumFlagsForItem: syscall.NewCallback(
			func(p uintptr, psi **_IUnknownVt, pgrfFlags *uint32) uintptr {
				ppImpl := (**_IShellItemFilterImpl)(unsafe.Pointer(p))
				if fun := (**ppImpl).getEnumFlagsForItem; fun == nil { // user didn't define a callback
					return uintptr(co.HRESULT_S_OK)
				} else {
					return uintptr(fun(
						&IShellItem{IUnknown{psi}},
						(*co.SHCONTF)(pgrfFlags),
					))
				}
			},
		),
	}
}
