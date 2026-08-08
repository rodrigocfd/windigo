//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/internal/utl"
)

// // Validates the HRESULT, and constructs a new COM object within ppOut.
// //
// // Returns object and HRESULT.
// func com_buildObj_retObjHres[T _IIUnknown](
// 	ret uintptr,
// 	ppvtQueried **_IUnknownVt,
// 	releaser *OleReleaser,
// ) (T, error) {
// 	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
// 		var pObj T
// 		com_buildObj(&pObj, ppvtQueried, releaser)
// 		return pObj, nil
// 	} else {
// 		var dummy T // will be a nil pointer
// 		return dummy, hr
// 	}
// }

// // Calls the pointed COM method without parameters, returns struct or newtype,
// // and HRESULT.
// func com_callRetStruct[T any](me _IIUnknown, pMethod uintptr) (T, error) {
// 	var obj T
// 	ret, _, _ := syscall.SyscallN(
// 		pMethod,
// 		me.Ppvt(),
// 		uintptr(unsafe.Pointer(&obj)))

// 	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
// 		return obj, nil
// 	} else {
// 		var dummy T
// 		return dummy, hr
// 	}
// }

// // Calls the pointed COM method without parameters, returns a single COM object
// // and HRESULT.
// func com_callRetCom[T _IIUnknown](me _IIUnknown, releaser *OleReleaser, pMethod uintptr) (T, error) {
// 	var ppvtQueried **_IUnknownVt
// 	ret, _, _ := syscall.SyscallN(
// 		pMethod,
// 		me.Ppvt(),
// 		uintptr(unsafe.Pointer(&ppvtQueried)))
// 	return com_buildObj_retObjHres[T](ret, ppvtQueried, releaser)
// }

// // A [COM] object whose lifetime can be managed by an [OleReleaser], automating
// // the cleanup.
// //
// // [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// type OleResource interface {
// 	utl.CanRelease
// }

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
