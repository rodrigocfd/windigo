//go:build windows

package win

import (
	"sync/atomic"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IFileOperationProgressSink] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// # Example
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var op *win.IFileOperation
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_FileOperation,
//		nil,
//		co.CLSCTX_ALL, &op,
//	)
//
//	sink := win.NewIFileOperationProgressSinkImpl(rel)
//	sink.PreCopyItem(
//		func(
//			flags co.TSF,
//			item, destFolder *win.IShellItem,
//			newName string,
//		) co.HRESULT {
//			println("Pre-copy", newName)
//			return co.HRESULT_S_OK
//		},
//	)
//	_, _ = op.Advise(sink)
//
// [IFileOperationProgressSink]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileoperationprogresssink
type IFileOperationProgressSink struct{ IUnknown }

// Returns the unique [COM] [interface ID].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileOperationProgressSink) IID() co.IID {
	return co.IID_IFileOperationProgressSink
}

type _IFileOperationProgressSinkImpl struct {
	vt               _IFileOperationProgressSinkVt
	counter          uint32
	startOperations  func() co.HRESULT
	finishOperations func(result co.HRESULT) co.HRESULT
	preRenameItem    func(flags co.TSF, item *IShellItem, newName string) co.HRESULT
	postRenameItem   func(flags co.TSF, item *IShellItem, newName string, hrRename co.HRESULT, newlyCreated *IShellItem) co.HRESULT
	preMoveItem      func(flags co.TSF, item, destFolder *IShellItem, newName string) co.HRESULT
	postMoveItem     func(flags co.TSF, item, destFolder *IShellItem, newName string, hrMove co.HRESULT, newlyCreated *IShellItem) co.HRESULT
	preCopyItem      func(flags co.TSF, item, destFolder *IShellItem, newName string) co.HRESULT
	postCopyItem     func(flags co.TSF, item, destFolder *IShellItem, newName string, hrMove co.HRESULT, newlyCreated *IShellItem) co.HRESULT
	preDeleteItem    func(flags co.TSF, item *IShellItem) co.HRESULT
	postDeleteItem   func(flags co.TSF, item *IShellItem, hrDelete co.HRESULT, newlyCreated *IShellItem) co.HRESULT
	preNewItem       func(flags co.TSF, destFolder *IShellItem, newName string) co.HRESULT
	postNewItem      func(flags co.TSF, destFolder *IShellItem, newName, templateName string, attr co.FILE_ATTRIBUTE, hrNew co.HRESULT, newItem *IShellItem) co.HRESULT
	updateProgress   func(workTotal, workSoFar uint) co.HRESULT
	resetTimer       func() co.HRESULT
	pauseTimer       func() co.HRESULT
	resumeTimer      func() co.HRESULT
}

// Implements [IFileOperationProgressSink].
//
// # Example
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var op *win.IFileOperation
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_FileOperation,
//		nil,
//		co.CLSCTX_ALL, &op,
//	)
//
//	sink := win.NewIFileOperationProgressSinkImpl(rel)
//	sink.PreCopyItem(
//		func(
//			flags co.TSF,
//			item, destFolder *win.IShellItem,
//			newName string,
//		) co.HRESULT {
//			println("Pre-copy", newName)
//			return co.HRESULT_S_OK
//		},
//	)
//	_, _ = op.Advise(sink)
//
// [IFileOperationProgressSink]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileoperationprogresssink
func NewIFileOperationProgressSinkImpl(releaser *OleReleaser) *IFileOperationProgressSink {
	_iFileOperationProgressSinkVtPtrs.init()
	pImpl := &_IFileOperationProgressSinkImpl{ // has Go function pointers, so cannot be allocated on the OS heap
		vt:      _iFileOperationProgressSinkVtPtrs, // simply copy the syscall callback pointers
		counter: 1,
	}
	utl.PtrCache.Add(unsafe.Pointer(pImpl)) // keep ptr
	ppImpl := &pImpl
	utl.PtrCache.Add(unsafe.Pointer(ppImpl)) // also keep ptr ptr

	ppFakeVtbl := (**_IUnknownVt)(unsafe.Pointer(ppImpl))
	pObj := &IFileOperationProgressSink{IUnknown{ppFakeVtbl}}
	releaser.Add(pObj)
	return pObj
}

// Defines [StartOperations] method.
//
// [StartOperations]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-startoperations
func (me *IFileOperationProgressSink) StartOperations(fun func() co.HRESULT) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).startOperations = fun
}

// Defines [FinishOperations] method.
//
// [FinishOperations]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-finishoperations
func (me *IFileOperationProgressSink) FinishOperations(fun func(result co.HRESULT) co.HRESULT) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).finishOperations = fun
}

// Defines [PreRenameItem] method.
//
// [PreRenameItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-prerenameitem
func (me *IFileOperationProgressSink) PreRenameItem(
	fun func(flags co.TSF, item *IShellItem, newName string) co.HRESULT,
) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).preRenameItem = fun
}

// Defines [PostRenameItem] method.
//
// [PostRenameItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-postrenameitem
func (me *IFileOperationProgressSink) PostRenameItem(
	fun func(flags co.TSF, item *IShellItem, newName string, hrRename co.HRESULT, newlyCreated *IShellItem) co.HRESULT,
) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).postRenameItem = fun
}

// Defines [PreMoveItem] method.
//
// [PreMoveItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-premoveitem
func (me *IFileOperationProgressSink) PreMoveItem(
	fun func(flags co.TSF, item, destFolder *IShellItem, newName string) co.HRESULT,
) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).preMoveItem = fun
}

// Defines [PostMoveItem] method.
//
// [PostMoveItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-postmoveitem
func (me *IFileOperationProgressSink) PostMoveItem(
	fun func(flags co.TSF, item, destFolder *IShellItem, newName string, hrMove co.HRESULT, newlyCreated *IShellItem) co.HRESULT,
) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).postMoveItem = fun
}

// Defines [PreCopyItem] method.
//
// [PreCopyItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-precopyitem
func (me *IFileOperationProgressSink) PreCopyItem(
	fun func(flags co.TSF, item, destFolder *IShellItem, newName string) co.HRESULT,
) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).preCopyItem = fun
}

// Defines [PostCopyItem] method.
//
// [PostCopyItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-postcopyitem
func (me *IFileOperationProgressSink) PostCopyItem(
	fun func(flags co.TSF, item, destFolder *IShellItem, newName string, hrMove co.HRESULT, newlyCreated *IShellItem) co.HRESULT,
) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).postCopyItem = fun
}

// Defines [PreDeleteItem] method.
//
// [PreDeleteItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-predeleteitem
func (me *IFileOperationProgressSink) PreDeleteItem(
	fun func(flags co.TSF, item *IShellItem) co.HRESULT,
) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).preDeleteItem = fun
}

// Defines [PostDeleteItem] method.
//
// [PostDeleteItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-postdeleteitem
func (me *IFileOperationProgressSink) PostDeleteItem(
	fun func(flags co.TSF, item *IShellItem, hrDelete co.HRESULT, newlyCreated *IShellItem) co.HRESULT,
) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).postDeleteItem = fun
}

// Defines [PreNewItem] method.
//
// [PreNewItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-prenewitem
func (me *IFileOperationProgressSink) PreNewItem(
	fun func(flags co.TSF, destFolder *IShellItem, newName string) co.HRESULT,
) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).preNewItem = fun
}

// Defines [PostNewItem] method.
//
// [PostNewItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-postnewitem
func (me *IFileOperationProgressSink) PostNewItem(
	fun func(flags co.TSF, destFolder *IShellItem, newName, templateName string, attr co.FILE_ATTRIBUTE, hrNew co.HRESULT, newItem *IShellItem) co.HRESULT,
) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).postNewItem = fun
}

// Defines [UpdateProgress] method.
//
// [UpdateProgress]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-updateprogress
func (me *IFileOperationProgressSink) UpdateProgress(fun func(workTotal, workSoFar uint) co.HRESULT) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).updateProgress = fun
}

// Defines [ResetTimer] method.
//
// [ResetTimer]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-resettimer
func (me *IFileOperationProgressSink) ResetTimer(fun func() co.HRESULT) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).resetTimer = fun
}

// Defines [PauseTimer] method.
//
// [PauseTimer]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-pausetimer
func (me *IFileOperationProgressSink) PauseTimer(fun func() co.HRESULT) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).pauseTimer = fun
}

// Defines [ResumeTimer] method.
//
// [ResumeTimer]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperationprogresssink-resumetimer
func (me *IFileOperationProgressSink) ResumeTimer(fun func() co.HRESULT) {
	(*(**_IFileOperationProgressSinkImpl)(unsafe.Pointer(me.Ppvt()))).resumeTimer = fun
}

type _IFileOperationProgressSinkVt struct {
	_IUnknownVt
	StartOperations  uintptr
	FinishOperations uintptr
	PreRenameItem    uintptr
	PostRenameItem   uintptr
	PreMoveItem      uintptr
	PostMoveItem     uintptr
	PreCopyItem      uintptr
	PostCopyItem     uintptr
	PreDeleteItem    uintptr
	PostDeleteItem   uintptr
	PreNewItem       uintptr
	PostNewItem      uintptr
	UpdateProgress   uintptr
	ResetTimer       uintptr
	PauseTimer       uintptr
	ResumeTimer      uintptr
}

var _iFileOperationProgressSinkVtPtrs _IFileOperationProgressSinkVt // Global to keep the syscall callback pointers.

func (me *_IFileOperationProgressSinkVt) init() {
	if me.QueryInterface == 0 { // initialize only once
		*me = _IFileOperationProgressSinkVt{
			_IUnknownVt: _IUnknownVt{
				QueryInterface: syscall.NewCallback(
					func(_p uintptr, _riid uintptr, ppv ***_IUnknownVt) uintptr {
						*ppv = nil
						return uintptr(co.HRESULT_E_NOTIMPL)
					},
				),
				AddRef: syscall.NewCallback(
					func(p uintptr) uintptr {
						ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
						newCount := atomic.AddUint32(&(**ppImpl).counter, 1)
						return uintptr(newCount)
					},
				),
				Release: syscall.NewCallback(
					func(p uintptr) uintptr {
						ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
						newCount := atomic.AddUint32(&(*ppImpl).counter, ^uint32(0)) // decrement 1
						if newCount == 0 {
							utl.PtrCache.Delete(unsafe.Pointer(*ppImpl)) // now GC can collect them
							utl.PtrCache.Delete(unsafe.Pointer(ppImpl))
						}
						return uintptr(newCount)
					},
				),
			},
			StartOperations: syscall.NewCallback(
				func(p uintptr) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).startOperations; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun())
					}
				},
			),
			FinishOperations: syscall.NewCallback(
				func(p uintptr, hrResult co.HRESULT) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).finishOperations; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(hrResult))
					}
				},
			),
			PreRenameItem: syscall.NewCallback(
				func(p uintptr, dwFlags uint32, psiItem **_IUnknownVt, pszNewName *uint16) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).preRenameItem; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							co.TSF(dwFlags),
							&IShellItem{IUnknown{psiItem}},
							wstr.DecodePtr(pszNewName),
						))
					}
				},
			),
			PostRenameItem: syscall.NewCallback(
				func(p uintptr,
					dwFlags uint32,
					psiItem **_IUnknownVt,
					pszNewName *uint16,
					hrRename co.HRESULT,
					psiNewlyCreated **_IUnknownVt,
				) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).postRenameItem; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							co.TSF(dwFlags),
							&IShellItem{IUnknown{psiItem}},
							wstr.DecodePtr(pszNewName),
							hrRename,
							&IShellItem{IUnknown{psiNewlyCreated}},
						))
					}
				},
			),
			PreMoveItem: syscall.NewCallback(
				func(p uintptr,
					dwFlags uint32,
					psiItem, psiDestinationFolder **_IUnknownVt,
					pszNewName *uint16,
				) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).preMoveItem; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							co.TSF(dwFlags),
							&IShellItem{IUnknown{psiItem}},
							&IShellItem{IUnknown{psiDestinationFolder}},
							wstr.DecodePtr(pszNewName),
						))
					}
				},
			),
			PostMoveItem: syscall.NewCallback(
				func(p uintptr,
					dwFlags uint32,
					psiItem, psiDestinationFolder **_IUnknownVt,
					pszNewName *uint16,
					hrMove co.HRESULT,
					psiNewlyCreated **_IUnknownVt,
				) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).postMoveItem; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							co.TSF(dwFlags),
							&IShellItem{IUnknown{psiItem}},
							&IShellItem{IUnknown{psiDestinationFolder}},
							wstr.DecodePtr(pszNewName),
							hrMove,
							&IShellItem{IUnknown{psiNewlyCreated}},
						))
					}
				},
			),
			PreCopyItem: syscall.NewCallback(
				func(p uintptr,
					dwFlags uint32,
					psiItem, psiDestinationFolder **_IUnknownVt,
					pszNewName *uint16,
				) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).preCopyItem; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							co.TSF(dwFlags),
							&IShellItem{IUnknown{psiItem}},
							&IShellItem{IUnknown{psiDestinationFolder}},
							wstr.DecodePtr(pszNewName),
						))
					}
				},
			),
			PostCopyItem: syscall.NewCallback(
				func(p uintptr,
					dwFlags uint32,
					psiItem, psiDestinationFolder **_IUnknownVt,
					pszNewName *uint16,
					hrMove co.HRESULT,
					psiNewlyCreated **_IUnknownVt,
				) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).postCopyItem; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							co.TSF(dwFlags),
							&IShellItem{IUnknown{psiItem}},
							&IShellItem{IUnknown{psiDestinationFolder}},
							wstr.DecodePtr(pszNewName),
							hrMove,
							&IShellItem{IUnknown{psiNewlyCreated}},
						))
					}
				},
			),
			PreDeleteItem: syscall.NewCallback(
				func(p uintptr, dwFlags uint32, psiItem **_IUnknownVt) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).preDeleteItem; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							co.TSF(dwFlags),
							&IShellItem{IUnknown{psiItem}},
						))
					}
				},
			),
			PostDeleteItem: syscall.NewCallback(
				func(p uintptr,
					dwFlags uint32,
					psiItem **_IUnknownVt,
					hrMove co.HRESULT,
					psiNewlyCreated **_IUnknownVt,
				) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).postDeleteItem; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							co.TSF(dwFlags),
							&IShellItem{IUnknown{psiItem}},
							hrMove,
							&IShellItem{IUnknown{psiNewlyCreated}},
						))
					}
				},
			),
			PreNewItem: syscall.NewCallback(
				func(p uintptr,
					dwFlags uint32,
					psiDestinationFolder **_IUnknownVt,
					pszNewName *uint16,
				) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).preNewItem; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							co.TSF(dwFlags),
							&IShellItem{IUnknown{psiDestinationFolder}},
							wstr.DecodePtr(pszNewName),
						))
					}
				},
			),
			PostNewItem: syscall.NewCallback(
				func(p uintptr,
					dwFlags uint32,
					psiDestinationFolder **_IUnknownVt,
					pszNewName, pszTemplateName *uint16,
					fileAttributes co.FILE_ATTRIBUTE,
					hrNew co.HRESULT,
					psiNewItem **_IUnknownVt,
				) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).postNewItem; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(
							co.TSF(dwFlags),
							&IShellItem{IUnknown{psiDestinationFolder}},
							wstr.DecodePtr(pszNewName),
							wstr.DecodePtr(pszTemplateName),
							fileAttributes,
							hrNew,
							&IShellItem{IUnknown{psiNewItem}},
						))
					}
				},
			),
			UpdateProgress: syscall.NewCallback(
				func(p uintptr, iWorkTotal, iWorkSoFar uint32) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).updateProgress; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun(uint(iWorkTotal), uint(iWorkSoFar)))
					}
				},
			),
			ResetTimer: syscall.NewCallback(
				func(p uintptr, iWorkTotal, iWorkSoFar uint32) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).resetTimer; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun())
					}
				},
			),
			PauseTimer: syscall.NewCallback(
				func(p uintptr, iWorkTotal, iWorkSoFar uint32) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).pauseTimer; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun())
					}
				},
			),
			ResumeTimer: syscall.NewCallback(
				func(p uintptr, iWorkTotal, iWorkSoFar uint32) uintptr {
					ppImpl := (**_IFileOperationProgressSinkImpl)(unsafe.Pointer(p))
					if fun := (**ppImpl).resumeTimer; fun == nil { // user didn't define a callback
						return uintptr(co.HRESULT_S_OK)
					} else {
						return uintptr(fun())
					}
				},
			),
		}
	}
}
