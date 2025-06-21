//go:build windows

package win

import (
	"sync/atomic"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// [IFileDialogEvents] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IFileDialogEvents]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifiledialogevents
type IFileDialogEvents struct{ IUnknown }

// Returns the unique [COM] [interface ID].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileDialogEvents) IID() co.IID {
	return co.IID_IFileDialogEvents
}

type _IFileDialogEventsImpl struct {
	vt                _IFileDialogEventsVt
	counter           uint32
	onFileOk          func() co.HRESULT
	onFolderChanging  func(folder *IShellItem) co.HRESULT
	onFolderChange    func() co.HRESULT
	onSelectionChange func() co.HRESULT
	onShareViolation  func(item *IShellItem, pResponse *co.FDESVR) co.HRESULT
	onTypeChange      func() co.HRESULT
	onOverwrite       func(item *IShellItem, pResponse *co.FDEOR) co.HRESULT
}

// Implements [IFileDialogEvents].
//
// [IFileDialogEvents]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifiledialogevents
func NewIFileDialogEventsImpl(releaser *OleReleaser) *IFileDialogEvents {
	_iFileDialogEventsVtPtrs.init()
	pImpl := &_IFileDialogEventsImpl{ // has Go function pointers, so cannot be allocated on the OS heap
		vt:      _iFileDialogEventsVtPtrs, // simply copy the syscall callback pointers
		counter: 1,
	}
	utl.PtrCache.Add(unsafe.Pointer(pImpl)) // keep ptr
	ppImpl := &pImpl
	utl.PtrCache.Add(unsafe.Pointer(ppImpl)) // also keep ptr ptr

	ppFakeVtbl := (**_IUnknownVt)(unsafe.Pointer(ppImpl))
	pObj := &IFileDialogEvents{IUnknown{ppFakeVtbl}}
	releaser.Add(pObj)
	return pObj
}

// Defines [OnFileOk] method.
//
// [OnFileOk]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialogevents-onfileok
func (me *IFileDialogEvents) OnFileOk(fun func() co.HRESULT) {
	(*(**_IFileDialogEventsImpl)(unsafe.Pointer(me.Ppvt()))).onFileOk = fun
}

// Defines [OnFolderChanging] method.
//
// [OnFolderChanging]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialogevents-onfolderchanging
func (me *IFileDialogEvents) OnFolderChanging(fun func(item *IShellItem) co.HRESULT) {
	(*(**_IFileDialogEventsImpl)(unsafe.Pointer(me.Ppvt()))).onFolderChanging = fun
}

// Defines [OnFolderChange] method.
//
// [OnFolderChange]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialogevents-onfolderchange
func (me *IFileDialogEvents) OnFolderChange(fun func() co.HRESULT) {
	(*(**_IFileDialogEventsImpl)(unsafe.Pointer(me.Ppvt()))).onFolderChange = fun
}

// Defines [OnSelectionChange] method.
//
// [OnSelectionChange]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialogevents-onselectionchange
func (me *IFileDialogEvents) OnSelectionChange(fun func() co.HRESULT) {
	(*(**_IFileDialogEventsImpl)(unsafe.Pointer(me.Ppvt()))).onSelectionChange = fun
}

// Defines [OnShareViolation] method.
//
// [OnShareViolation]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialogevents-onshareviolation
func (me *IFileDialogEvents) OnShareViolation(fun func(item *IShellItem, pResponse *co.FDESVR) co.HRESULT) {
	(*(**_IFileDialogEventsImpl)(unsafe.Pointer(me.Ppvt()))).onShareViolation = fun
}

// Defines [OnTypeChange] method.
//
// [OnTypeChange]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialogevents-ontypechange
func (me *IFileDialogEvents) OnTypeChange(fun func() co.HRESULT) {
	(*(**_IFileDialogEventsImpl)(unsafe.Pointer(me.Ppvt()))).onTypeChange = fun
}

// Defines [OnOverwrite] method.
//
// [OnOverwrite]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialogevents-onoverwrite
func (me *IFileDialogEvents) OnOverwrite(fun func(item *IShellItem, pResponse *co.FDEOR) co.HRESULT) {
	(*(**_IFileDialogEventsImpl)(unsafe.Pointer(me.Ppvt()))).onOverwrite = fun
}

type _IFileDialogEventsVt struct {
	_IUnknownVt
	OnFileOk          uintptr
	OnFolderChanging  uintptr
	OnFolderChange    uintptr
	OnSelectionChange uintptr
	OnShareViolation  uintptr
	OnTypeChange      uintptr
	OnOverwrite       uintptr
}

var _iFileDialogEventsVtPtrs _IFileDialogEventsVt // Global to keep the syscall callback pointers.

func (me *_IFileDialogEventsVt) init() {
	if me.QueryInterface == 0 { // initialize only once
		*me = _IFileDialogEventsVt{
			_IUnknownVt: _IUnknownVt{
				QueryInterface: syscall.NewCallback(
					func(_p uintptr, _riid uintptr, ppv ***_IUnknownVt) uintptr {
						*ppv = nil
						return uintptr(co.HRESULT_E_NOTIMPL)
					},
				),
				AddRef: syscall.NewCallback(
					func(p uintptr) uintptr {
						ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
						newCount := atomic.AddUint32(&(**ppImpl).counter, 1)
						return uintptr(newCount)
					},
				),
				Release: syscall.NewCallback(
					func(p uintptr) uintptr {
						ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
						newCount := atomic.AddUint32(&(*ppImpl).counter, ^uint32(0)) // decrement 1
						if newCount == 0 {
							utl.PtrCache.Delete(unsafe.Pointer(*ppImpl)) // now GC can collect them
							utl.PtrCache.Delete(unsafe.Pointer(ppImpl))
						}
						return uintptr(newCount)
					},
				),
			},
			OnFileOk: syscall.NewCallback(
				func(p uintptr, _ **_IUnknownVt) uintptr {
					ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
					if fun := (*ppImpl).onFileOk; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun())
					}
				},
			),
			OnFolderChanging: syscall.NewCallback(
				func(p uintptr, _ **_IUnknownVt, vtSi **_IUnknownVt) uintptr {
					ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
					if fun := (*ppImpl).onFolderChanging; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(&IShellItem{IUnknown{vtSi}}))
					}
				},
			),
			OnFolderChange: syscall.NewCallback(
				func(p uintptr, _ **_IUnknownVt) uintptr {
					ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
					if fun := (*ppImpl).onFolderChange; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun())
					}
				},
			),
			OnSelectionChange: syscall.NewCallback(
				func(p uintptr, _ **_IUnknownVt) uintptr {
					ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
					if fun := (*ppImpl).onSelectionChange; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun())
					}
				},
			),
			OnShareViolation: syscall.NewCallback(
				func(p uintptr, _ **_IUnknownVt, vtSi **_IUnknownVt, pResponse *uint32) uintptr {
					ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
					if fun := (*ppImpl).onShareViolation; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							&IShellItem{IUnknown{vtSi}},
							(*co.FDESVR)(pResponse),
						))
					}
				},
			),
			OnTypeChange: syscall.NewCallback(
				func(p uintptr, _ **_IUnknownVt) uintptr {
					ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
					if fun := (*ppImpl).onTypeChange; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun())
					}
				},
			),
			OnOverwrite: syscall.NewCallback(
				func(p uintptr, _ **_IUnknownVt, vtSi **_IUnknownVt, pResponse *uint32) uintptr {
					ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
					if fun := (*ppImpl).onOverwrite; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							&IShellItem{IUnknown{vtSi}},
							(*co.FDEOR)(pResponse),
						))
					}
				},
			),
		}
	}
}
