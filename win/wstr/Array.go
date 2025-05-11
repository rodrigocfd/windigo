//go:build windows

package wstr

import (
	"unicode/utf8"
	"unsafe"
)

// Converts multiple Go strings into null-terminated Windows UTF-16 strings,
// stored in a single buffer, with a double terminating null.
//
// Uses short string optimization, allocating either on the stack or on the GC
// heap.
//
// Don't move this object, otherwise you'll invalidate the stack pointer.
//
// Used internally by the library to make syscalls.
//
// [HeapAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapalloc
type Array struct {
	buf   Buf[Stack20]
	count uint // Number of stored strings.
}

// Converts multiple Go strings into null-terminated Windows UTF-16 strings,
// stored in a single buffer, with a double terminating null.
//
// Uses short string optimization, allocating either on the stack or on the GC
// heap.
//
// Don't move this object, otherwise you'll invalidate the stack pointer.
//
// Used internally by the library to make syscalls.
//
// # Example
//
//	buf := wstr.NewArray("abc", "def")
func NewArray(strs ...string) Array {
	me := Array{
		buf:   NewBuf[Stack20](),
		count: 0,
	}
	me.Append(strs...)
	return me
}

// Appends the given strings to the buffer.
//
// Since the pointers are invalidated when the data is changed, first insert all
// the strings, then grab their pointers.
func (me *Array) Append(strs ...string) {
	if len(strs) == 0 {
		return
	}

	if me.count > 1 {
		me.buf.Resize(me.buf.Len() - 1) // if we have more than 1 string, remove double terminating null
	}
	curSz := me.buf.Len()

	neededSz := uint(0)
	for _, str := range strs {
		neededSz += uint(utf8.RuneCountInString(str)) + 1 // room for terminating null
	}
	if me.count+uint(len(strs)) > 1 {
		neededSz += 1 // double terminating null if multiple strings
	}
	me.buf.Resize(me.buf.Len() + neededSz)

	dest := me.buf.HotSlice()[curSz:] // slice to receive the strings
	for _, str := range strs {
		StrToWstrBuf(str, dest) // empty strings are also added
		strLen := utf8.RuneCountInString(str)
		dest = dest[strLen+1:] // advance the slice to receive the next string
	}

	me.count += uint(len(strs))
	if me.count > 1 {
		dest[0] = 0x0000 // double terminating null if we have multiple strings
	}
}

// Returns the number of strings actually stored.
func (me *Array) Count() uint {
	return me.count
}

// Returns the pointer to the given stored string, or nil if out of bounds.
//
// If the data is changed for whathever reason – like by adding an element –,
// the pointer will be no longer valid.
func (me *Array) PtrOf(index uint) *uint16 {
	if me.count == 0 {
		return nil
	} else if index == 0 { // first string is simply the beginning of the block
		return me.buf.At(0)
	} else { // skip each terminating null until we reach the desired index
		iNextStr := uint(1)
		for i := uint(0); i < me.buf.Len(); i++ {
			if *me.buf.At(i) == 0x0000 { // found a terminating null
				if iNextStr == index { // this is the index we're after
					return me.buf.At(i + 1) // 1st char past null
				} else {
					iNextStr++
				}
			}
		}
		return nil // index out of bounds
	}
}

// Returns the pointer to the given stored string, or nil if out of bounds.
//
// If the data is changed for whathever reason – like by adding an element –,
// the pointer will be no longer valid.
func (me *Array) UnsafePtrOf(index uint) unsafe.Pointer {
	return unsafe.Pointer(me.PtrOf(index))
}
