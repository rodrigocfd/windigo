//go:build windows

package wstr

import (
	"unicode/utf8"
	"unsafe"
)

// How to interpret the empty string value.
type EMPTY_STR uint8

const (
	// An empty string will yield a 1-char buffer with just the terminating
	// null.
	ALLOW_EMPTY EMPTY_STR = 0x1
	// An empty string will yield a nil pointer. No allocation will be made, and
	// the buffer will remain empty.
	EMPTY_IS_NIL EMPTY_STR = 0x2
)

type (
	Stack20  = [20]uint16  // Stack size of 20 chars.
	Stack64  = [64]uint16  // Stack size of 64 chars.
	Stack260 = [260]uint16 // Stack size of 260 chars (MAX_PATH).

	// Short string optimization size for wstr.Buf buffers.
	//
	// Used internally by the library to make syscalls.
	Sso interface {
		Stack20 | Stack64 | Stack260
	}
)

// A buffer with parametrized small string optimization to receive Windows
// UTF-16 strings, using the GC heap when needed. The buffer size always grows.
//
// Be careful when choosing the buffer size: if it won't fit the stack, it will
// escape to the GC heap, making all the effort useless. If you're unsure, just
// use 20.
//
// Don't move this object, otherwise you'll invalidate the stack pointer.
//
// Used internally by the library to make syscalls.
type Buf[B Sso] struct {
	stackBuf B
	heapBuf  []uint16
	inUse    uint
}

// Creates an empty wstr.Buf with a small buffer optimization of the given size.
//
// Be careful when choosing the buffer size: if it won't fit the stack, it will
// escape to the GC heap, making all the effort useless. If you're unsure, just
// use 20.
//
// Don't move this object, otherwise you'll invalidate the stack pointer.
//
// # Example
//
//	buf := wstr.NewBuf[wstr.Stack20]()
func NewBuf[B Sso]() Buf[B] {
	return Buf[B]{}
}

// Creates an empty wstr.Buf with a small buffer optimization of the given size,
// resizing it to the given number of UTF-16 chars.
//
// Be careful when choosing the buffer size: if it won't fit the stack, it will
// escape to the GC heap, making all the effort useless. If you're unsure, just
// use 20.
//
// Don't move this object, otherwise you'll invalidate the stack pointer.
//
// # Example
//
//	buf := wstr.NewBufSized[wstr.Stack20](20)
func NewBufSized[B Sso](numChars uint) Buf[B] {
	var me Buf[B]
	me.Resize(numChars)
	return me
}

// Creates an empty wstr.Buf with a small buffer optimization of the given size,
// converting the given string to an UTF-16 string, with a terminating null. The
// buffer is resized to accommodate it.
//
// Be careful when choosing the buffer size: if it won't fit the stack, it will
// escape to the GC heap, making all the effort useless. If you're unsure, just
// use 20.
//
// Don't move this object, otherwise you'll invalidate the stack pointer.
//
// # Example
//
//	buf := wstr.NewBufWith[wstr.Stack20]("abc", wstr.ALLOW_EMPTY)
func NewBufWith[B Sso](str string, emptyStr EMPTY_STR) Buf[B] {
	var me Buf[B]
	me.Set(str, emptyStr)
	return me
}

// Returns a pointer at the given char position, without checking bounds.
//
// If the buffer is changed for whathever reason, the pointer will be no longer
// valid.
func (me *Buf[B]) At(index uint) *uint16 {
	if me.IsStack() {
		return &me.stackBuf[index]
	} else {
		return &me.heapBuf[index]
	}
}

// Returns a slice over the current memory, either stack or heap. An empty
// buffer yields nil.
//
// If the buffer is changed for whathever reason, the slice will be no longer
// valid.
func (me *Buf[B]) HotSlice() []uint16 {
	if me.inUse == 0 {
		return nil
	}
	return unsafe.Slice(me.At(0), me.inUse)
}

// Returns true if the buffer is still using the stack allocation.
func (me *Buf[B]) IsStack() bool {
	return len(me.heapBuf) == 0 // once we launch the heap, we never go back to stack
}

// Returns the buffer length.
func (me *Buf[B]) Len() uint {
	return me.inUse
}

// Returns a pointer to the memory block, either stack or heap. An empty buffer
// yields nil.
//
// If the buffer is changed for whathever reason, the pointer will be no longer
// valid.
func (me *Buf[B]) Ptr() *uint16 {
	if me.inUse == 0 {
		return nil
	}
	return me.At(0)
}

// Resizes the buffer, growing or shrinking. Moves to the heap if the stack size
// isn't enough.
//
// Once we move to the heap, we never go back to the stack.
func (me *Buf[B]) Resize(numChars uint) {
	if me.IsStack() {
		if numChars > uint(len(me.stackBuf)) { // move from stack to heap
			me.heapBuf = make([]uint16, numChars)
			for i := uint(0); i < me.inUse; i++ { // copy from stack to heap
				me.heapBuf[i] = me.stackBuf[i]
				me.stackBuf[i] = 0x0000 // zero the stack
			}
		}
	} else if numChars > uint(len(me.heapBuf)) { // already in heap, requesting more room
		newHeap := make([]uint16, numChars)
		for i := uint(0); i < me.inUse; i++ { // copy to the newly allocated buffer
			newHeap[i] = me.heapBuf[i]
		}
		me.heapBuf = newHeap
	}

	if me.IsStack() { // if shrinking, fill the rest with zeros
		for i := numChars; i < me.inUse; i++ {
			me.stackBuf[i] = 0x0000
		}
	} else {
		for i := numChars; i < me.inUse; i++ {
			me.heapBuf[i] = 0x0000
		}
	}

	me.inUse = numChars
}

// Truncates the buffer and stores the string converted to UTF-16, with a
// terminating null.
func (me *Buf[B]) Set(str string, emptyStr EMPTY_STR) {
	me.Resize(0)
	if emptyStr == ALLOW_EMPTY || str != "" {
		strLen := utf8.RuneCountInString(str)
		me.Resize(uint(strLen) + 1) // room for terminating null
		StrToWstrBuf(str, me.HotSlice())
	}
}

// Returns a pointer to the memory block, either stack or heap. An empty buffer
// yields nil.
//
// If the buffer is changed for whathever reason, the pointer will be no longer
// valid.
func (me *Buf[B]) UnsafePtr() unsafe.Pointer {
	return unsafe.Pointer(me.Ptr())
}

// Sets all bytes to zero.
func (me *Buf[B]) ZeroBuffer() {
	if me.IsStack() {
		for i := uint(0); i < me.inUse; i++ {
			me.stackBuf[i] = 0x0000
		}
	} else {
		for i := uint(0); i < me.inUse; i++ {
			me.heapBuf[i] = 0x0000
		}
	}
}
