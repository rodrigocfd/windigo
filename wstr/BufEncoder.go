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
// This buffer is used internally to speed up syscalls with string arguments.
//
// Example:
//
//	var buf1, buf2 wstr.BufEncoder
//	p1 := buf1.AllowEmpty("abc")
//	p2 := buf2.EmptyIsNil("def")
type BufEncoder struct {
	stack [BUF_MAX]uint16
}

// Encodes a Go string into a null-terminated UTF-16, returning a pointer to it.
// If the string is empty, the buffer will contain just nulls.
//
// If the number of UTF-16 words fit the internal stack buffer, the returned
// pointer will point to the internal stack buffer. Otherwise, the returned
// pointer will point to a new heap-allocated slice.
func (me *BufEncoder) AllowEmpty(s string) unsafe.Pointer {
	slice := me.Slice(s)
	return unsafe.Pointer(unsafe.SliceData(slice))
}

// Encodes a Go string into a null-terminated UTF-16, returning a pointer to it.
// If the string is empty, returns nil.
//
// If the number of UTF-16 words fit the internal stack buffer, the returned
// pointer will point to the internal stack buffer. Otherwise, the returned
// pointer will point to a new heap-allocated slice.
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
