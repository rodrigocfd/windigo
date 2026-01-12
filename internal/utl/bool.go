//go:build windows

package utl

import (
	"reflect"
)

// Syntactic sugar; converts bool to 0 or 1.
func BoolToInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func BoolToUint32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func BoolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

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
