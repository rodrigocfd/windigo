//go:build windows

package win

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// A [COM] object whose lifetime can be managed by an [OleReleaser], automating
// the cleanup.
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type OleResource interface {
	release()
}

// A [COM] object, derived from [IUnknown].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type OleObj interface {
	OleResource

	// Returns the unique [COM] [interface ID].
	//
	// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
	// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
	IID() co.IID

	// Returns the [COM] virtual table pointer.
	//
	// This is a low-level method, used internally by the library. Incorrect usage
	// may lead to segmentation faults.
	//
	// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
	Ppvt() **_IUnknownVt
}

// Returns the virtual table pointer, performing a nil check.
func com_ppvtOrNil(obj OleObj) unsafe.Pointer {
	if !utl.IsNil(obj) {
		return unsafe.Pointer(obj.Ppvt())
	}
	return nil
}

// Converts a GUID string into the GUID struct buffer, returning its pointer.
//
// If strGuid is [co.GUID_NULL], returns nil.
func com_guidStructPtrOrNil[T ~string](strGuid T, guidStructBuf *GUID) unsafe.Pointer {
	if co.GUID(strGuid) == co.GUID_NULL {
		return nil
	}
	*guidStructBuf = GuidFrom(strGuid)
	return unsafe.Pointer(guidStructBuf)
}

// Validates the pointer to pointer to COM object. Panics if fails.
//
// If the object is fine, calls [OleReleaser.ReleaseNow].
//
// Returns the [OleObj.IID] of the underlying pointed-to object.
func com_validateAndRelease(ppOut interface{}, releaser *OleReleaser) co.IID {
	ppTy := reflect.TypeOf(ppOut) // **IUnknown
	if ppTy.Kind() != reflect.Ptr {
		panic("You must a pass a pointer to a pointer COM object [**Ty failed].")
	}

	pTy := ppTy.Elem() // *IUnknown
	if pTy.Kind() != reflect.Ptr {
		panic("You must a pass a pointer to a pointer COM object [*Ty failed].")
	}

	ty := pTy.Elem() // IUnknown
	if ty.Kind() != reflect.Struct {
		panic("You must a pass a pointer to a pointer COM object [Ty failed].")
	}

	pTarget := reflect.ValueOf(ppOut).Elem() // *IUnknown
	if !pTarget.CanSet() {
		panic("You must a pass a pointer to a pointer COM object [target CanSet() failed].")
	}
	var emptyVal reflect.Value
	if pTarget.MethodByName("IID") == emptyVal {
		panic("You must a pass a pointer to a pointer COM object [target IID() failed].")
	}

	pObj := pTarget.Interface().(OleObj) // *IUnknown
	releaser.ReleaseNow(pObj)            // safety, because pOut will receive the new COM object
	return pObj.IID()
}

// Constructs a new COM object within the ppOut, writing ppvtQueried in its
// first field.
//
// Always returns a nil error.
func com_buildObj(ppOut interface{}, ppvtQueried **_IUnknownVt, releaser *OleReleaser) error {
	pTarget := reflect.ValueOf(ppOut).Elem()  // *IUnknown
	ty := reflect.TypeOf(ppOut).Elem().Elem() // IUnknown
	pTarget.Set(reflect.New(ty))              // instantiate new object on the heap and assign its pointer

	addrField0 := pTarget.Elem().Field(0).UnsafeAddr()
	*(*uintptr)(unsafe.Pointer(addrField0)) = uintptr(unsafe.Pointer(ppvtQueried)) // assign ppvt field

	pObj := pTarget.Interface().(OleObj) // *IUnknown
	releaser.Add(pObj)
	return nil
}

// Validates the HRESULT, and constructs a new COM object within ppOut.
//
// Returns HRESULT.
func com_buildObj_retHres(ret uintptr, ppOut interface{}, ppvtQueried **_IUnknownVt, releaser *OleReleaser) error {
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return com_buildObj(ppOut, ppvtQueried, releaser)
	} else {
		return hr
	}
}

// Validates the HRESULT, and constructs a new COM object within ppOut.
//
// Returns object and HRESULT.
func com_buildObj_retObjHres[T OleObj](ret uintptr, ppvtQueried **_IUnknownVt, releaser *OleReleaser) (T, error) {
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		var pObj T
		com_buildObj(&pObj, ppvtQueried, releaser)
		return pObj, nil
	} else {
		var dummy T // will be a nil pointer
		return dummy, hr
	}
}

// Calls the COM method, without parameters, returns error.
func com_callErr(me OleObj, pMethod uintptr) error {
	ret, _, _ := syscall.SyscallN(
		pMethod,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.HresultToError(ret)
}

// Calls the COM method which returns a single COM object.
func com_callObj[T OleObj](me OleObj, releaser *OleReleaser, pMethod uintptr) (T, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		pMethod,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[T](ret, ppvtQueried, releaser)
}

// Calls the COM method, without parameters, returning BSTR and error.
func com_callBstrGet(me OleObj, pMethod uintptr) (string, error) {
	var name BSTR
	defer name.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		pMethod,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&name)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return name.String(), nil
	} else {
		return "", hr
	}
}

// Calls COM method, receives a BSTR and returns error.
func com_callBstrSet(me OleObj, s string, pMethod uintptr) error {
	bstrS, err := SysAllocString(s)
	if err != nil {
		return err
	}
	defer bstrS.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		pMethod,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(bstrS))
	return utl.HresultToError(ret)
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
	objs []OleResource
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
	return new(OleReleaser)
}

// Adds a new [COM] resource to have its lifetime managed by the [OleReleaser].
func (me *OleReleaser) Add(objs ...OleResource) {
	me.objs = append(me.objs, objs...)
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
	for i := len(me.objs) - 1; i >= 0; i-- {
		me.objs[i].release()
	}
	me.objs = nil
}

// Releases the specific [COM] resources, if present, immediately.
//
// These objects will be removed from the internal list, thus not being released
// when [OleReleaser.Release] is further called.
//
// Panics if no object is passed.
func (me *OleReleaser) ReleaseNow(objs ...OleResource) {
	if len(objs) == 0 {
		panic("No objects passed to ReleaseNow.")
	}

NextHisObj:
	for _, hisObj := range objs {
		if utl.IsNil(hisObj) {
			continue // skip nil objects
		}
		hisObj.release() // release no matter what

		for ourIdx, ourObj := range me.objs {
			if ourObj == hisObj { // we found the passed object in our array
				copy(me.objs[ourIdx:len(me.objs)-1], me.objs[ourIdx+1:len(me.objs)]) // move subsequent elements into the gap
				me.objs[len(me.objs)-1] = nil
				me.objs = me.objs[:len(me.objs)-1] // shrink our slice over the same memory
				continue NextHisObj
			}
		}
	}
}
