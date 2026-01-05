//go:build windows

package wstr

import (
	"unsafe"
)

// A buffer to receive UTF-16 strings and convert them to Go strings.
//
// This buffer is used internally to speed up syscalls that return strings, and
// it's prone to buffer overruns. Be sure to allocate the needed space.
//
// This struct contains a buffer intended to be stack-allocated, so don't move
// it.
//
// Example:
//
//	var buf wstr.BufDecoder
//	buf.Alloc(20)
//	ptr := buf.Ptr()
type BufDecoder struct {
	stack [BUF_MAX]uint16
	heap  []uint16
	sz    int
}

// Makes sure there is enough room for the given number of chars. If an
// allocation is necessary, any previous content will be lost.
//
// Panics if numChars is negative.
func (me *BufDecoder) Alloc(numChars int) {
	if numChars > BUF_MAX {
		me.heap = make([]uint16, numChars)
	} else {
		me.heap = nil
	}
	me.sz = numChars
}

// Makes sure there is enough room for the given number of chars. If an
// allocation is necessary, any previous content will be lost.
//
// In addition, zeroes the whole buffer.
//
// Panics if numChars is negative.
func (me *BufDecoder) AllocAndZero(numChars int) {
	me.Alloc(numChars)
	me.Zero()
}

// Returns a slice over the internal memory block, up to latest
// [BufDecoder.Alloc] call.
func (me *BufDecoder) HotSlice() []uint16 {
	if me.heap != nil {
		return me.heap
	} else {
		return me.stack[:me.sz]
	}
}

// Returns the size of the last call to [BufDecoder.Alloc].
func (me *BufDecoder) Len() int {
	return me.sz
}

// Returns a pointer to the internal memory block, either stack or heap.
func (me *BufDecoder) Ptr() unsafe.Pointer {
	if me.heap != nil {
		return unsafe.Pointer(unsafe.SliceData(me.heap))
	} else {
		return unsafe.Pointer(unsafe.SliceData(me.stack[:]))
	}
}

// Converts the contents to a Go string.
func (me *BufDecoder) String() string {
	return DecodeSlice(me.HotSlice())
}

// Zeroes the whole buffer.
func (me *BufDecoder) Zero() {
	if me.heap != nil {
		for i := 0; i < len(me.heap); i++ {
			me.heap[i] = 0x0000
		}
	} else {
		for i := 0; i < len(me.stack); i++ {
			me.stack[i] = 0x0000
		}
	}
}
