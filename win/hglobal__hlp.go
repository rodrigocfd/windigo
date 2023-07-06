//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// This helper method calls GlobalAlloc to alloc a null-terminated *uint16.
//
// With co.GMEM_FIXED, the handle itself is the pointer to the memory block, and
// it can optionally be passed to unsafe.Slice() to create a slice over the
// memory block.
//
// With co.GMEM_MOVEABLE, you must call HGLOBAL.GlobalLock() to retrieve the
// pointer.
//
// ⚠️ You must defer HGLOBAL.GlobalFree().
//
// # Example
//
//	hMem := win.GlobalAllocStr(co.GMEM_FIXED, "my text")
//	defer hMem.GlobalFree()
//
//	charSlice := hMem.GlobalLock(hMem.GlobalSize())
//	defer hMem.GlobalUnlock()
func GlobalAllocStr(uFlags co.GMEM, s string) HGLOBAL {
	sliceStr16 := Str.ToNativeSlice(s) // null-terminated
	sliceStr8 := unsafe.Slice((*byte)(unsafe.Pointer(&sliceStr16[0])), len(sliceStr16)*2)

	hMem := GlobalAlloc(uFlags, len(sliceStr8))
	if (uFlags & co.GMEM_MOVEABLE) != 0 {
		dest := hMem.GlobalLock(len(sliceStr8))
		copy(dest, sliceStr8)
		hMem.GlobalUnlock()
	} else {
		dest := unsafe.Slice((*byte)(unsafe.Pointer(hMem)), len(sliceStr8))
		copy(dest, sliceStr8)
	}
	return hMem
}
