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

// Validates the pointer to pointer to COM object. Panics if fails.
//
// Returns the underlying pointed-to object.
func OleValidateObj(ppOut interface{}) interface{} {
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

	return pTarget.Interface()
}

// Constructs a new COM object, assigns it to the pointer, and sets its
// **IUnknownVt.
//
// Returns the underlying pointed-to object.
func OleCreateObj(ppOut interface{}, ppIUnknownVt unsafe.Pointer) interface{} {
	pTarget := reflect.ValueOf(ppOut).Elem()  // *IUnknown
	ty := reflect.TypeOf(ppOut).Elem().Elem() // IUnknown
	pTarget.Set(reflect.New(ty))              // instantiate new object on the heap and assign its pointer

	addrField0 := pTarget.Elem().Field(0).UnsafeAddr()
	*(*uintptr)(unsafe.Pointer(addrField0)) = uintptr(ppIUnknownVt) // assign ppvt field

	return pTarget.Interface()
}
