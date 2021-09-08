package util

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"
)

func _ToUint16Ptr(s string) unsafe.Pointer {
	pStr, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		panic(fmt.Sprintf("Variant conversion from string failed \"%s\": %s",
			s, err))
	}
	return unsafe.Pointer(pStr)
}

// Converts from nil or string; any other type will panic.
func VariantNilString(val interface{}) unsafe.Pointer {
	if val != nil {
		switch v := val.(type) {
		case string:
			return _ToUint16Ptr(v)
		default:
			panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(val)))
		}
	}
	return nil
}

// Converts from uint16 or string; any other type will panic.
func VariantUint16String(val interface{}) unsafe.Pointer {
	switch v := val.(type) {
	case uint16:
		return unsafe.Pointer(uintptr(v))
	case string:
		return _ToUint16Ptr(v)
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(val)))
	}
}
