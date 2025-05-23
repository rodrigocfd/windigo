//go:build windows

package utl

import (
	"reflect"
	"unsafe"
)

// Checks whether the interface carries an actual [nil value].
//
// [nil value]: https://blog.devtrovert.com/p/go-secret-interface-nil-is-not-nil
func IsNil(x interface{}) bool {
	if x == nil {
		return true
	}

	value := reflect.ValueOf(x)
	kind := value.Kind()

	switch kind {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Ptr,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}

// Validates the receiving pointer to pointer to COM object; panics if fails.
func ComValidateOutPtr(ppOut interface{}) {
	ppTy := reflect.TypeOf(ppOut) // **IUnknown
	if ppTy.Kind() != reflect.Ptr {
		panic("You must a pass a pointer to a pointer COM object [**Ty failed].")
	}

	pTy := ppTy.Elem() // *IUnknown
	if pTy.Kind() != reflect.Ptr {
		panic("You must a pass a pointer to a pointer COM object [*Ty failed].")
	}

	// ty := pTy.Elem() // IUnknown

	pTarget := reflect.ValueOf(ppOut).Elem() // *IUnknown
	if !pTarget.CanSet() {
		panic("You must a pass a pointer to a pointer COM object [target CanSet() failed].")
	}
	var emptyVal reflect.Value
	if pTarget.MethodByName("IID") == emptyVal {
		panic("You must a pass a pointer to a pointer COM object [target IID() failed].")
	}
}

// Calls the IID() method to retrieve the co.IID constant from a COM object.
func ComRetrieveIid(ppOut interface{}) string {
	pTarget := reflect.ValueOf(ppOut).Elem() // *IUnknown
	rets := pTarget.MethodByName("IID").Call([]reflect.Value{})
	return rets[0].String()
}

// Creates a COM object, assign it to the pointer, and sets is **IUnknownVt.
func ComCreateObj(ppOut interface{}, ppIUnknownVt unsafe.Pointer) {
	pTarget := reflect.ValueOf(ppOut).Elem()  // *IUnknown
	ty := reflect.TypeOf(ppOut).Elem().Elem() // IUnknown
	pTarget.Set(reflect.New(ty))              // instantiate new object on the heap and assign its pointer

	addrField0 := pTarget.Elem().Field(0).UnsafeAddr()
	*(*uintptr)(unsafe.Pointer(addrField0)) = uintptr(ppIUnknownVt)
}
