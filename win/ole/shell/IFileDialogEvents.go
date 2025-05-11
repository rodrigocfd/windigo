//go:build windows

package shell

import (
	"sync/atomic"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
)

// [IFileDialogEvents] COM interface.
//
// [IFileDialogEvents]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifiledialogevents
type IFileDialogEvents struct{ ole.IUnknown }

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
func NewIFileDialogEventsImpl(releaser *ole.Releaser) *IFileDialogEvents {
	iFileDialogEventsCallbacks()
	pImpl := &_IFileDialogEventsImpl{ // has Go function pointers, so cannot be allocated on the OS heap
		vt: _IFileDialogEventsVt{
			IUnknownVt: ole.IUnknownVt{
				QueryInterface: _iFileDialogEventsQueryInterface,
				AddRef:         _iFileDialogEventsAddRef,
				Release:        _iFileDialogEventsRelease,
			},
			OnFileOk:          _iFileDialogEventsOnFileOk,
			OnFolderChanging:  _iFileDialogEventsOnFolderChanging,
			OnFolderChange:    _iFileDialogEventsOnFolderChange,
			OnSelectionChange: _iFileDialogEventsOnSelectionChange,
			OnShareViolation:  _iFileDialogEventsOnShareViolation,
			OnTypeChange:      _iFileDialogEventsOnTypeChange,
			OnOverwrite:       _iFileDialogEventsOnOverwrite,
		},
		counter: 1,
	}
	wutil.PtrCache.Add(unsafe.Pointer(pImpl)) // keep ptr
	ppImpl := &pImpl
	wutil.PtrCache.Add(unsafe.Pointer(ppImpl)) // also keep ptr ptr

	ppFakeVtbl := (**ole.IUnknownVt)(unsafe.Pointer(ppImpl))
	pObj := ole.ComObj[IFileDialogEvents](ppFakeVtbl)
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

var (
	_iFileDialogEventsQueryInterface uintptr
	_iFileDialogEventsAddRef         uintptr
	_iFileDialogEventsRelease        uintptr

	_iFileDialogEventsOnFileOk          uintptr
	_iFileDialogEventsOnFolderChanging  uintptr
	_iFileDialogEventsOnFolderChange    uintptr
	_iFileDialogEventsOnSelectionChange uintptr
	_iFileDialogEventsOnShareViolation  uintptr
	_iFileDialogEventsOnTypeChange      uintptr
	_iFileDialogEventsOnOverwrite       uintptr
)

func iFileDialogEventsCallbacks() {
	if _iFileDialogEventsQueryInterface != 0 {
		return
	}

	_iFileDialogEventsQueryInterface = syscall.NewCallback(
		func(_p uintptr, _riid uintptr, ppv ***ole.IUnknownVt) uintptr {
			*ppv = nil
			return uintptr(co.HRESULT_E_NOTIMPL)
		},
	)
	_iFileDialogEventsAddRef = syscall.NewCallback(
		func(p uintptr) uintptr {
			ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
			newCount := atomic.AddUint32(&(**ppImpl).counter, 1)
			return uintptr(newCount)
		},
	)
	_iFileDialogEventsRelease = syscall.NewCallback(
		func(p uintptr) uintptr {
			ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
			newCount := atomic.AddUint32(&(*ppImpl).counter, ^uint32(0)) // decrement 1
			if newCount == 0 {
				wutil.PtrCache.Delete(unsafe.Pointer(*ppImpl)) // now GC can collect them
				wutil.PtrCache.Delete(unsafe.Pointer(ppImpl))
			}
			return uintptr(newCount)
		},
	)

	_iFileDialogEventsOnFileOk = syscall.NewCallback(
		func(p uintptr, _ **ole.IUnknownVt) uintptr {
			ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
			if fun := (*ppImpl).onFileOk; fun == nil { // user didn't define a callback
				return uintptr(co.HRESULT_S_OK)
			} else {
				return uintptr(fun())
			}
		},
	)
	_iFileDialogEventsOnFolderChanging = syscall.NewCallback(
		func(p uintptr, _ **ole.IUnknownVt, vtSi **ole.IUnknownVt) uintptr {
			ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
			if fun := (*ppImpl).onFolderChanging; fun == nil { // user didn't define a callback
				return uintptr(co.HRESULT_S_OK)
			} else {
				pItem := ole.ComObj[IShellItem](vtSi)
				return uintptr(fun(pItem))
			}
		},
	)
	_iFileDialogEventsOnFolderChange = syscall.NewCallback(
		func(p uintptr, _ **ole.IUnknownVt) uintptr {
			ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
			if fun := (*ppImpl).onFolderChange; fun == nil { // user didn't define a callback
				return uintptr(co.HRESULT_S_OK)
			} else {
				return uintptr(fun())
			}
		},
	)
	_iFileDialogEventsOnSelectionChange = syscall.NewCallback(
		func(p uintptr, _ **ole.IUnknownVt) uintptr {
			ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
			if fun := (*ppImpl).onSelectionChange; fun == nil { // user didn't define a callback
				return uintptr(co.HRESULT_S_OK)
			} else {
				return uintptr(fun())
			}
		},
	)
	_iFileDialogEventsOnShareViolation = syscall.NewCallback(
		func(p uintptr, _ **ole.IUnknownVt, vtSi **ole.IUnknownVt, pResponse *uint32) uintptr {
			ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
			if fun := (*ppImpl).onShareViolation; fun == nil { // user didn't define a callback
				return uintptr(co.HRESULT_S_OK)
			} else {
				pItem := ole.ComObj[IShellItem](vtSi)
				response := co.FDESVR(*pResponse)
				ret := fun(pItem, &response)
				*pResponse = uint32(response)
				return uintptr(ret)
			}
		},
	)
	_iFileDialogEventsOnTypeChange = syscall.NewCallback(
		func(p uintptr, _ **ole.IUnknownVt) uintptr {
			ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
			if fun := (*ppImpl).onTypeChange; fun == nil { // user didn't define a callback
				return uintptr(co.HRESULT_S_OK)
			} else {
				return uintptr(fun())
			}
		},
	)
	_iFileDialogEventsOnOverwrite = syscall.NewCallback(
		func(p uintptr, _ **ole.IUnknownVt, vtSi **ole.IUnknownVt, pResponse *uint32) uintptr {
			ppImpl := (**_IFileDialogEventsImpl)(unsafe.Pointer(p))
			if fun := (*ppImpl).onOverwrite; fun == nil { // user didn't define a callback
				return uintptr(co.HRESULT_S_OK)
			} else {
				pItem := ole.ComObj[IShellItem](vtSi)
				response := co.FDEOR(*pResponse)
				ret := fun(pItem, &response)
				*pResponse = uint32(response)
				return uintptr(ret)
			}
		},
	)
}

type _IFileDialogEventsVt struct {
	ole.IUnknownVt
	OnFileOk          uintptr
	OnFolderChanging  uintptr
	OnFolderChange    uintptr
	OnSelectionChange uintptr
	OnShareViolation  uintptr
	OnTypeChange      uintptr
	OnOverwrite       uintptr
}
