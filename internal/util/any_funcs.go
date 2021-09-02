package util

import (
	"fmt"
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

// Syntactic sugar; converts bool to 0 or 1.
func BoolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// Returns first value if condition is true, otherwise returns second.
//
// Return type must be cast accordingly.
func Iif(cond bool, ifTrue, ifFalse interface{}) interface{} {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

// Converts val to *uint16 or string; any other type will panic.
//
// Use runtime.KeepAlive() to make sure an eventual string will stay reachable.
func PullUint16String(val interface{}) uintptr {
	switch v := val.(type) {
	case uint16:
		return uintptr(v)

	case string:
		pStr, err := syscall.UTF16PtrFromString(v)
		if err != nil {
			panic(fmt.Sprintf("PullUint16String() failed \"%s\": %s", v, err))
		}
		return uintptr(unsafe.Pointer(pStr)) // runtime.KeepAlive()

	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(val)))
	}
}

// "&He && she" becomes "He & she".
func RemoveAccelAmpersands(text string) string {
	runes := []rune(text)
	buf := strings.Builder{}
	buf.Grow(len(runes)) // prealloc for performance

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == '&' && runes[i+1] != '&' {
			continue
		}
		buf.WriteRune(runes[i])
	}
	if runes[len(runes)-1] != '&' {
		buf.WriteRune(runes[len(runes)-1])
	}
	return buf.String()
}
