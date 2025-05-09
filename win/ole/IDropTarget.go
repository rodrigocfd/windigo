//go:build windows

package ole

import (
	"sync/atomic"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// [IDropTarget] COM interface.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	dropTarget := ole.NewIDropTargetImpl(rel)
//
// [IDropTarget]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nn-oleidl-idroptarget
type IDropTarget struct{ IUnknown }

type _IDropTargetImpl struct {
	vt        vt.IDropTarget
	counter   uint32
	dragEnter func(dataObj *IDataObject, keyState co.MK, pt win.POINT, effect *co.DROPEFFECT) co.HRESULT
	dragOver  func(keyState co.MK, pt win.POINT, effect *co.DROPEFFECT) co.HRESULT
	dragLeave func() co.HRESULT
	drop      func(dataObj *IDataObject, keyState co.MK, pt win.POINT, effect *co.DROPEFFECT) co.HRESULT
}

// Implements [IDropTarget].
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	dropTarget := ole.NewIDropTargetImpl(rel)
//
// [IDropTarget]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nn-oleidl-idroptarget
func NewIDropTargetImpl(releaser *Releaser) *IDropTarget {
	iDropTargetCallbacks()
	pImpl := &_IDropTargetImpl{ // has Go function pointers, so cannot be allocated on the OS heap
		vt: vt.IDropTarget{
			IUnknown: vt.IUnknown{
				QueryInterface: _iDropTargetQueryInterface,
				AddRef:         _iDropTargetAddRef,
				Release:        _iDropTargetRelease,
			},
			DragEnter: _iDropTargetDragEnter,
			DragOver:  _iDropTargetDragOver,
			DragLeave: _iDropTargetDragLeave,
			Drop:      _iDropTargetDrop,
		},
		counter: 1,
	}
	wutil.PtrCache.Add(unsafe.Pointer(pImpl)) // keep ptr
	ppImpl := &pImpl
	wutil.PtrCache.Add(unsafe.Pointer(ppImpl)) // also keep ptr ptr

	ppFakeVtbl := (**vt.IUnknown)(unsafe.Pointer(ppImpl))
	pObj := vt.NewObj[IDropTarget](ppFakeVtbl)
	releaser.Add(pObj)
	return pObj
}

// Defines [DragEnter] method.
//
// [DragEnter]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragenter
func (me *IDropTarget) DragEnter(
	fun func(dataObj *IDataObject, keyState co.MK, pt win.POINT, effect *co.DROPEFFECT) co.HRESULT,
) {
	(*(**_IDropTargetImpl)(unsafe.Pointer(me.Ppvt()))).dragEnter = fun
}

// Defines [DragOver] method.
//
// [DragOver]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragover
func (me *IDropTarget) DragOver(
	fun func(keyState co.MK, pt win.POINT, effect *co.DROPEFFECT) co.HRESULT,
) {
	(*(**_IDropTargetImpl)(unsafe.Pointer(me.Ppvt()))).dragOver = fun
}

// Defines [DragLeave] method.
//
// [DragLeave]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragleave
func (me *IDropTarget) DragLeave(fun func() co.HRESULT) {
	(*(**_IDropTargetImpl)(unsafe.Pointer(me.Ppvt()))).dragLeave = fun
}

// Defines [Drop] method.
//
// # Example
//
//	var dropTarget *ole.IDropTarget // initialized somewhere
//
//	dropTarget.Drop(
//		func(
//			dataObj *ole.IDataObject,
//			keyState co.MK,
//			pt win.POINT,
//			effect *co.DROPEFFECT,
//		) co.HRESULT {
//			fetc := ole.FORMATETC{
//				CfFormat: co.CF_HDROP,
//				Aspect:   co.DVASPECT_CONTENT,
//				Lindex:   -1,
//				Tymed:    co.TYMED_HGLOBAL,
//			}
//
//			stg, err := dataObj.GetData(&fetc)
//			if err != nil {
//				panic(err)
//			}
//			defer ole.ReleaseStgMedium(&stg)
//
//			if hGlobal, ok := stg.HGlobal(); ok {
//				hMem, _ := hGlobal.GlobalLock()
//				defer hGlobal.GlobalUnlock()
//
//				// DragFinish() crashes ReleaseStgMedium(), don't call
//				hDrop := win.HDROP(hMem)
//				for path, _ := range hDrop.Iter() {
//					println(path)
//				}
//			}
//			return co.HRESULT_S_OK
//		},
//	)
//
// [Drop]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-idroptarget-drop
func (me *IDropTarget) Drop(
	fun func(dataObj *IDataObject, keyState co.MK, pt win.POINT, effect *co.DROPEFFECT) co.HRESULT,
) {
	(*(**_IDropTargetImpl)(unsafe.Pointer(me.Ppvt()))).drop = fun
}

var (
	_iDropTargetQueryInterface uintptr
	_iDropTargetAddRef         uintptr
	_iDropTargetRelease        uintptr

	_iDropTargetDragEnter uintptr
	_iDropTargetDragOver  uintptr
	_iDropTargetDragLeave uintptr
	_iDropTargetDrop      uintptr
)

func iDropTargetCallbacks() {
	if _iDropTargetQueryInterface != 0 {
		return
	}

	_iDropTargetQueryInterface = syscall.NewCallback(
		func(_p uintptr, _riid uintptr, ppv ***vt.IUnknown) uintptr {
			*ppv = nil
			return uintptr(co.HRESULT_E_NOTIMPL)
		},
	)
	_iDropTargetAddRef = syscall.NewCallback(
		func(p uintptr) uintptr {
			ppImpl := (**_IDropTargetImpl)(unsafe.Pointer(p))
			newCount := atomic.AddUint32(&(**ppImpl).counter, 1)
			return uintptr(newCount)
		},
	)
	_iDropTargetRelease = syscall.NewCallback(
		func(p uintptr) uintptr {
			ppImpl := (**_IDropTargetImpl)(unsafe.Pointer(p))
			newCount := atomic.AddUint32(&(*ppImpl).counter, ^uint32(0)) // decrement 1
			if newCount == 0 {
				wutil.PtrCache.Delete(unsafe.Pointer(*ppImpl)) // now GC can collect them
				wutil.PtrCache.Delete(unsafe.Pointer(ppImpl))
			}
			return uintptr(newCount)
		},
	)

	_iDropTargetDragEnter = syscall.NewCallback(
		func(p uintptr, vtDataObj **vt.IUnknown, grfKeyState uint32, pt win.POINT, pdwEffect *uint32) uintptr {
			ppImpl := (**_IDropTargetImpl)(unsafe.Pointer(p))
			pDataObj := vt.NewObj[IDataObject](vtDataObj)
			if fun := (*ppImpl).dragEnter; fun == nil { // user didn't define a callback
				return uintptr(co.HRESULT_S_OK)
			} else {
				effect := co.DROPEFFECT(*pdwEffect)
				ret := fun(pDataObj, co.MK(grfKeyState), pt, &effect)
				*pdwEffect = uint32(effect)
				return uintptr(ret)
			}
		},
	)
	_iDropTargetDragOver = syscall.NewCallback(
		func(p uintptr, grfKeyState uint32, pt win.POINT, pdwEffect *uint32) uintptr {
			ppImpl := (**_IDropTargetImpl)(unsafe.Pointer(p))
			if fun := (*ppImpl).dragOver; fun == nil { // user didn't define a callback
				return uintptr(co.HRESULT_S_OK)
			} else {
				effect := co.DROPEFFECT(*pdwEffect)
				ret := fun(co.MK(grfKeyState), pt, &effect)
				*pdwEffect = uint32(effect)
				return uintptr(ret)
			}
		},
	)
	_iDropTargetDragLeave = syscall.NewCallback(
		func(p uintptr) uintptr {
			ppImpl := (**_IDropTargetImpl)(unsafe.Pointer(p))
			if fun := (*ppImpl).dragLeave; fun == nil { // user didn't define a callback
				return uintptr(co.HRESULT_S_OK)
			} else {
				return uintptr(fun())
			}
		},
	)
	_iDropTargetDrop = syscall.NewCallback(
		func(p uintptr, vtDataObj **vt.IUnknown, grfKeyState uint32, pt win.POINT, pdwEffect *uint32) uintptr {
			ppImpl := (**_IDropTargetImpl)(unsafe.Pointer(p))
			pDataObj := vt.NewObj[IDataObject](vtDataObj)
			if fun := (*ppImpl).drop; fun == nil { // user didn't define a callback
				return uintptr(co.HRESULT_S_OK)
			} else {
				effect := co.DROPEFFECT(*pdwEffect)
				ret := fun(pDataObj, co.MK(grfKeyState), pt, &effect)
				*pdwEffect = uint32(effect)
				return uintptr(ret)
			}
		},
	)
}
