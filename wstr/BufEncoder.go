//go:build windows

package wstr

import (
	"unsafe"
)

// Size of the stack buffer, equal to [MAX_PATH].
//
// [MAX_PATH]: https://stackoverflow.com/a/1880453/6923555
const BUF_MAX = 260

// Encodes a Go string into a null-terminated UTF-16. If the string fits the
// stack buffer, no allocation is performed.
//
// This buffer is used internally to speed up syscalls with string arguments,
// and it's prone to buffer overruns if used incorrectly.
//
// This struct contains a buffer intended to be stack-allocated, so don't move
// it.
//
// Example:
//
//	var wFoo, wBar wstr.BufEncoder
//
//	_, _, _ = syscall.SyscallN(
//		pProc,
//		uintptr(wFoo.AllowEmpty("foo")),
//	)
//
//	_, _, _ = syscall.SyscallN(
//		pProc,
//		uintptr(wFoo.EmptyIsNil("bar")),
//	)
type BufEncoder struct {
	stack [BUF_MAX]uint16
}

// Encodes a Go string into a null-terminated UTF-16, returning a pointer to it.
// If the string is empty, the buffer will contain just nulls.
//
// If the number of UTF-16 words fits the internal stack buffer, the returned
// pointer will point to the internal stack buffer. Otherwise, the returned
// pointer will point to a new heap-allocated slice.
//
// The returned pointer can be passed to syscalls.
//
// Example:
//
//	var wFoo wstr.BufEncoder
//	_, _, _ = syscall.SyscallN(
//		pProc,
//		uintptr(wFoo.AllowEmpty("foo")),
//	)
func (me *BufEncoder) AllowEmpty(s string) unsafe.Pointer {
	slice := me.Slice(s)
	return unsafe.Pointer(&slice[0])
}

// Encodes a Go string into a null-terminated UTF-16, returning a pointer to it.
// If the string is empty, returns nil.
//
// If the number of UTF-16 words fit the internal stack buffer, the returned
// pointer will point to the internal stack buffer. Otherwise, the returned
// pointer will point to a new heap-allocated slice.
//
// The returned pointer can be passed to syscalls.
//
// Example:
//
//	var wFoo wstr.BufEncoder
//	_, _, _ = syscall.SyscallN(
//		pProc,
//		uintptr(wFoo.EmptyIsNil("foo")),
//	)
func (me *BufEncoder) EmptyIsNil(s string) unsafe.Pointer {
	if s == "" {
		return nil
	}
	return me.AllowEmpty(s)
}

// Encodes a Go string into a null-terminated UTF-16, returning a slice with it.
//
// If the number of UTF-16 words fit the internal stack buffer, the returned
// slice will point to the internal stack buffer. Otherwise, the returned slice
// will point to a new heap-allocated slice.
func (me *BufEncoder) Slice(s string) []uint16 {
	numWords := CountUtf16Len(s) + 1 // plus terminating null
	if numWords > BUF_MAX {
		heap := make([]uint16, numWords) // overflows our stack buffer, use the heap
		EncodeToBuf(heap, s)
		return heap
	} else {
		EncodeToBuf(me.stack[:], s)
		return me.stack[:numWords]
	}
}
