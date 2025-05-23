//go:build windows

package ole

import (
	"sync/atomic"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
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

// Returns the unique [COM] [interface ID].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDropTarget) IID() co.IID {
	return co.IID_IDropTarget
}

type _IDropTargetImpl struct {
	vt        _IDropTargetVt
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
	_iDropTargetVtPtrs.init()
	pImpl := &_IDropTargetImpl{ // has Go function pointers, so cannot be allocated on the OS heap
		vt:      _iDropTargetVtPtrs, // simply copy the syscall callback pointers
		counter: 1,
	}
	utl.PtrCache.Add(unsafe.Pointer(pImpl)) // keep ptr
	ppImpl := &pImpl
	utl.PtrCache.Add(unsafe.Pointer(ppImpl)) // also keep ptr ptr

	ppFakeVtbl := (**IUnknownVt)(unsafe.Pointer(ppImpl))
	var pObj *IDropTarget
	utl.ComCreateObj(&pObj, unsafe.Pointer(ppFakeVtbl))
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

type _IDropTargetVt struct {
	IUnknownVt
	DragEnter uintptr
	DragOver  uintptr
	DragLeave uintptr
	Drop      uintptr
}

var _iDropTargetVtPtrs _IDropTargetVt // Global to keep the syscall callback pointers.

func (me *_IDropTargetVt) init() {
	if me.QueryInterface == 0 { // initialize only once
		*me = _IDropTargetVt{
			IUnknownVt: IUnknownVt{
				QueryInterface: syscall.NewCallback(
					func(_p uintptr, _riid uintptr, ppv ***IUnknownVt) uintptr {
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
			DragEnter: syscall.NewCallback(
				func(p uintptr, vtDataObj **IUnknownVt, grfKeyState uint32, pt win.POINT, pdwEffect *uint32) uintptr {
					ppImpl := (**_IDropTargetImpl)(unsafe.Pointer(p))
					var pDataObj *IDataObject
					utl.ComCreateObj(&pDataObj, unsafe.Pointer(vtDataObj))
					if fun := (*ppImpl).dragEnter; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						effect := co.DROPEFFECT(*pdwEffect)
						ret := fun(pDataObj, co.MK(grfKeyState), pt, &effect)
						*pdwEffect = uint32(effect)
						return uintptr(ret)
					}
				},
			),
			DragOver: syscall.NewCallback(
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
			),
			DragLeave: syscall.NewCallback(
				func(p uintptr) uintptr {
					ppImpl := (**_IDropTargetImpl)(unsafe.Pointer(p))
					if fun := (*ppImpl).dragLeave; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun())
					}
				},
			),
			Drop: syscall.NewCallback(
				func(p uintptr, vtDataObj **IUnknownVt, grfKeyState uint32, pt win.POINT, pdwEffect *uint32) uintptr {
					ppImpl := (**_IDropTargetImpl)(unsafe.Pointer(p))
					var pDataObj *IDataObject
					utl.ComCreateObj(&pDataObj, unsafe.Pointer(vtDataObj))
					if fun := (*ppImpl).drop; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						effect := co.DROPEFFECT(*pdwEffect)
						ret := fun(pDataObj, co.MK(grfKeyState), pt, &effect)
						*pdwEffect = uint32(effect)
						return uintptr(ret)
					}
				},
			),
		}
	}
}
