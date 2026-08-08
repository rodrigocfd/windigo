//go:build windows

package utl

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
)

// Converts a pointer to virtual table to the specific type.
func Vt[T interface{}](ppvt uintptr) *T {
	return *(**T)(unsafe.Pointer(ppvt))
}

// Calls Release() in multiple objects at once.
type OleBatchReleaser struct {
	objs []interface{ Release() }
}

func NewOleBatchReleaser() OleBatchReleaser {
	return OleBatchReleaser{
		objs: make([]interface{ Release() }, 0, 10), // arbitrary
	}
}
func (me *OleBatchReleaser) Add(obj interface{ Release() }) {
	me.objs = append(me.objs, obj)
}
func (me *OleBatchReleaser) Release() {
	for i := len(me.objs) - 1; i >= 0; i-- { // backwards
		me.objs[i].Release()
	}
	me.objs = nil
}

// Returns the virtual table pointer, performing a nil check.
func OlePpvtOrNil(obj interface{ Ppvt() uintptr }) uintptr {
	if !IsNil(obj) {
		return obj.Ppvt()
	}
	return 0
}

// Panics if ppOut is invalid, calls ppOut.Release(), returns IID.
func OleValidateRelease(ppOut interface{}) *co.IID {
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
	if pTarget.MethodByName("Release") == emptyVal {
		panic("You must a pass a pointer to a pointer COM object [target Release() failed].")
	} else if pTarget.MethodByName("IID") == emptyVal {
		panic("You must a pass a pointer to a pointer COM object [target IID() failed].")
	}

	pObj := pTarget.Interface().(interface { // *IUnknown
		Release()
		IID() *co.IID
	})
	if !pTarget.IsNil() { // object already constructed
		pObj.Release()
	}
	return pObj.IID()
}

func oleInjectWithOptionalReleaser(
	ppOut interface{}, // type must be **IUnknown or derived, will be modified through reflection
	ppvtQueried uintptr,
	releaserOptional interface {
		Add(obj interface{ Release() })
	},
) {
	pTarget := reflect.ValueOf(ppOut).Elem()  // *IUnknown
	ty := reflect.TypeOf(ppOut).Elem().Elem() // IUnknown
	pTarget.Set(reflect.New(ty))              // instantiate new object on the heap and assign its pointer

	addrField0 := pTarget.Elem().Field(0).UnsafeAddr()
	*(*uintptr)(unsafe.Pointer(addrField0)) = uintptr(unsafe.Pointer(ppvtQueried)) // assign ppvt field

	if releaserOptional != nil {
		pObj := pTarget.Interface().(interface{ Release() }) // *IUnknown
		releaserOptional.Add(pObj)
	}
}

// Actions:
//   - validates ppOut through reflection;
//   - injects ppvtQueried into ppOut;
//   - adds ppOut to releaser.
func OleInject(
	ppOut interface{}, // type must be **IUnknown or derived, will be modified through reflection
	ppvtQueried uintptr,
	releaser interface {
		Add(obj interface{ Release() })
	},
) {
	if releaser == nil {
		panic("Releaser cannot be nil.")
	}
	oleInjectWithOptionalReleaser(ppOut, ppvtQueried, releaser)
}

// Actions:
//   - continues if ret is S_OK, otherwise returns HRESULT;
//   - validates ppOut through reflection;
//   - injects ppvtQueried into ppOut;
//   - adds ppOut to releaser.
func OleInjectIfOk(
	ret uintptr,
	ppOut interface{},
	ppvtQueried uintptr,
	releaser interface {
		Add(obj interface{ Release() })
	},
) error {
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		OleInject(ppOut, ppvtQueried, releaser)
		return nil
	} else {
		return hr
	}
}

// Actions:
//   - creates a new object of type T, which must be *IUnknown or derived;
//   - injects ppvtQueried;
//   - returns it.
func OleNewWithoutReleaser[T interface{ IID() *co.IID }](ppvtQueried uintptr) T {
	var pObj T // nil pointer
	oleInjectWithOptionalReleaser(&pObj, ppvtQueried, nil)
	return pObj
}

// Actions:
//   - creates a new object of type T, which must be *IUnknown or derived;
//   - injects ppvtQueried;
//   - adds to the releaser;
//   - returns it.
func OleNew[T interface{ IID() *co.IID }](
	ppvtQueried uintptr,
	releaser interface {
		Add(obj interface{ Release() })
	},
) T {
	var pObj T // nil pointer
	OleInject(&pObj, ppvtQueried, releaser)
	return pObj
}

// Actions:
//   - continues if ret is S_OK, otherwise returns HRESULT;
//   - creates a new object of type T, which must be *IUnknown or derived;
//   - injects ppvtQueried;
//   - adds to the releaser;
//   - returns it.
func OleNewIfOk[T interface{ IID() *co.IID }](
	ret uintptr,
	ppvtQueried uintptr,
	releaser interface {
		Add(obj interface{ Release() })
	},
) (T, error) {
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return OleNew[T](ppvtQueried, releaser), nil
	} else {
		var dummy T // nil pointer
		return dummy, hr
	}
}

// Actions:
//   - calls pMethod without parameters;
//   - continues if S_OK, otherwise returns HRESULT;
//   - returns the object returned by pMethod.
func OleNewFromCallWithoutParms[T interface{ IID() *co.IID }](
	me interface{ Ppvt() uintptr },
	releaser interface {
		Add(obj interface{ Release() })
	},
	pMethod uintptr,
) (T, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		pMethod,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return OleNewIfOk[T](ret, ppvtQueried, releaser)
}

// Calls the given method without parameters.
func OleCallWithoutParms(me interface{ Ppvt() uintptr }, pMethod uintptr) error {
	ret, _, _ := syscall.SyscallN(
		pMethod,
		me.Ppvt())
	return HresultToError(ret)
}

// Calls the given method without paramters, and returns struct or newtype.
func OleCallReturnStruct[T interface{}](me interface{ Ppvt() uintptr }, pMethod uintptr) (T, error) {
	var obj T
	ret, _, _ := syscall.SyscallN(
		pMethod,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&obj)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return obj, nil
	} else {
		var dummy T
		return dummy, hr
	}
}

// So it can be used by HIMAGELIST and ui without causing a cyclic dependency.
var Shell_SHGetFileInfoW *syscall.Proc
