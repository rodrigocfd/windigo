//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// Dynamic, manually [heap-allocated] memory block array.
//
// Created with:
//   - [NewVec]
//   - [NewVecReserved]
//   - [NewVecSized]
//
// Do not store Go pointers in a Vec – this will make the GC believe they are no
// more in use, thus collecting them.
//
// [heap-allocated]: https://learn.microsoft.com/en-us/windows/win32/Memory/heap-functions
type Vec[T any] struct {
	data  []T  // Slice to the heap-allocated memory.
	inUse uint // Number of elements effectively being used.
}

// Creates a new, unallocated [Vec], which allows manual memory management from
// the heap.
//
// Do not store Go pointers in a Vec – this will make the GC believe they are no
// more in use, thus collecting them.
//
// ⚠️ You must defer [Vec.Free].
//
// Example:
//
//	pts := win.NewVec[win.POINT]()
//	defer pts.Free()
func NewVec[T any]() Vec[T] {
	return Vec[T]{
		data:  nil,
		inUse: 0,
	}
}

// Creates a new [Vec] with preallocated memory, but zero elements.
//
// Do not store Go pointers in a Vec – this will make the GC believe they are no
// more in use, thus collecting them.
//
// ⚠️ You must defer [Vec.Free].
//
// Example:
//
//	pts := win.NewVecReserved[win.POINT](30)
//	defer pts.Free()
func NewVecReserved[T any](numElems uint) Vec[T] {
	var me Vec[T]
	me.Reserve(numElems)
	return me
}

// Creates a new [Vec] with the given copies of the element.
//
// Do not store Go pointers in a Vec – this will make the GC believe they are no
// more in use, thus collecting them.
//
// ⚠️ You must defer [Vec.Free].
//
// Example:
//
//	pts := win.NewVecSized(30, win.POINT{})
//	defer pts.Free()
func NewVecSized[T any](numElems uint, elem T) Vec[T] {
	var me Vec[T]
	me.AppendN(numElems, elem)
	return me
}

// Appends new elements, increasing the buffer size if needed.
//
// Example:
//
//	bigNums := win.NewVec[uint64]()
//	defer bigNums.Free()
//
//	bigNums.Append(200)
//
//	others := []uint64{10, 20, 30}
//	bigNums.Append(others...)
func (me *Vec[T]) Append(elems ...T) {
	me.Reserve(me.inUse + uint(len(elems)))
	for _, elem := range elems {
		me.data[me.inUse] = elem
		me.inUse++
	}
}

// Appends N copies of the element, increasing the buffer size if needed.
func (me *Vec[T]) AppendN(numElems uint, elem T) {
	me.Reserve(me.inUse + numElems)
	for i := uint(0); i < numElems; i++ {
		me.data[me.inUse] = elem
		me.inUse++
	}
}

// Removes all elements, keeping the reserved size.
func (me *Vec[T]) Clear() {
	var dummy T
	for i := uint(0); i < me.inUse; i++ {
		me.data[i] = dummy
	}
	me.inUse = 0
}

// Releases the allocated heap memory, if any.
func (me *Vec[T]) Free() {
	if me.data != nil {
		hHeap, _ := GetProcessHeap()
		hHeap.HeapFree(co.HEAP_NS_NONE, unsafe.Pointer(&me.data[0]))
		me.data = nil
		me.inUse = 0
	}
}

// Returns a pointer the element at the given position.
//
// Does not perform bounds check.
//
// If the buffer is changed for whathever reason – like by adding an element or
// reserving more space –, this pointer will be no longer valid.
func (me *Vec[T]) Get(index uint) *T {
	return &me.data[index]
}

// Returns a slice over the current elements.
//
// If the data is changed for whathever reason – like by adding an element or
// reserving more space –, the slice will be no longer valid.
func (me *Vec[T]) HotSlice() []T {
	if me.inUse == 0 {
		return []T{}
	} else {
		return me.data[:me.inUse]
	}
}

// Returns true if there are no elements.
func (me *Vec[T]) IsEmpty() bool {
	return me.inUse == 0
}

// Returns the number of elements currently stored, not counting the reserved
// space.
func (me *Vec[T]) Len() uint {
	return me.inUse
}

// Allocates memory for the given number of elements, reserving the space,
// without adding elements.
//
// This method is intended for optimization purposes. If you want to create a
// buffer to receive data, use [Vec.Resize] instead.
//
// If amount is smaller than the current buffer size, does nothing; that is,
// this function only grows the buffer.
func (me *Vec[T]) Reserve(numElems uint) {
	if numElems > uint(len(me.data)) {
		newSizeBytes := numElems * me.szElem()
		hHeap, _ := GetProcessHeap()
		if me.data == nil {
			ptr, _ := hHeap.HeapAlloc(co.HEAP_ALLOC_ZERO_MEMORY, newSizeBytes)
			me.data = unsafe.Slice((*T)(ptr), numElems)
		} else {
			curPtr := unsafe.Pointer(&me.data[0])
			newPtr, _ := hHeap.HeapReAlloc(co.HEAP_REALLOC_ZERO_MEMORY, curPtr, newSizeBytes)
			me.data = unsafe.Slice((*T)(newPtr), numElems)
		}
	}
}

// Returns the actual number of allocated elements in the buffer.
func (me *Vec[T]) Reserved() uint {
	return uint(len(me.data))
}

// Resizes the internal buffer to the given number of elements. If increased,
// the given element is used to fill the new positions.
func (me *Vec[T]) Resize(numElements uint, elemToFill T) {
	if numElements > me.inUse { // enlarge
		me.AppendN(numElements-me.inUse, elemToFill)
	} else if me.inUse > numElements { // shrink
		var dummy T
		for i := numElements; i < me.inUse; i++ {
			me.data[i] = dummy // fill the unused memory
		}
		me.inUse = numElements
	}
}

// Returns a pointer to allocated memory block.
//
// If the buffer is changed for whathever reason – like by adding an element or
// reserving more space –, this pointer will be no longer valid.
func (me *Vec[T]) UnsafePtr() unsafe.Pointer {
	if me.IsEmpty() {
		return nil
	} else {
		return unsafe.Pointer(&me.data[0])
	}
}

// Size of a single element, in bytes.
func (me *Vec[T]) szElem() uint {
	var dummy T
	return uint(unsafe.Sizeof(dummy))
}
