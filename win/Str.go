/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

type _StrT struct{}

// UTF-16 string conversion functions.
var Str _StrT

// Converts a pointer to a null-terminated UTF-16 into string.
//
// Copied from syscall_windows.go, utf16PtrToString() private function.
func (_StrT) FromUint16Ptr(p *uint16) string {
	if p == nil {
		return ""
	}

	// Find null terminator.
	pEnd := unsafe.Pointer(p)
	sLen := 0
	for *(*uint16)(pEnd) != 0 {
		pEnd = unsafe.Pointer(uintptr(pEnd) + unsafe.Sizeof(*p)) // pointer++
		sLen++
	}

	// Turn *uint16 into []uint16.
	// https://stackoverflow.com/a/43592538
	// https://golang.org/pkg/internal/unsafeheader/#Slice
	var sliceMem = struct { // slice memory layout
		addr unsafe.Pointer
		len  int
		cap  int
	}{unsafe.Pointer(p), sLen, sLen}

	// Decode []uint16 into string.
	return string(utf16.Decode(*(*[]uint16)(unsafe.Pointer(&sliceMem))))
}

// Converts a null-terminated UTF-16 slice into string.
//
// Simple wrapper to syscall.UTF16ToString().
func (_StrT) FromUint16Slice(s []uint16) string {
	return syscall.UTF16ToString(s)
}

// Converts string to *uint16.
//
// Ideal to pass strings to syscalls. We won't return an uintptr right away
// because it has no pointer semantics, it's just a number, so pointed memory
// can be garbage-collected.
//
// https://stackoverflow.com/a/51188315
//
// Wrapper to syscall.UTF16PtrFromString(). Panics on error.
func (_StrT) ToUint16Ptr(s string) *uint16 {
	pstr, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		panic(fmt.Sprintf("StrToPtr failed \"%s\": %s",
			s, err))
	}
	return pstr
}

// Converts string to null-terminated []uint16.
//
// Wrapper to syscall.UTF16FromString(). Panics on error.
func (_StrT) ToUint16Slice(s string) []uint16 {
	sli, err := syscall.UTF16FromString(s)
	if err != nil {
		panic(fmt.Sprintf("StrToSlice failed \"%s\": %s",
			s, err))
	}
	return sli
}

// Converts string to *uint16, or nil if string is empty.
//
// Wrapper to syscall.UTF16PtrFromString(). Panics on error.
func (_StrT) ToUint16PtrBlankIsNil(s string) *uint16 {
	if s != "" {
		return Str.ToUint16Ptr(s)
	}
	return nil
}
