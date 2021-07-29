package util

import (
	"strings"
	"unsafe"
)

// Syntactic sugar; converts bool to 0 or 1.
func BoolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// Converts *byte to []byte, with the given length.
func PtrToSliceByte(ptr *byte, sz int) []byte {
	// https://stackoverflow.com/a/43592538
	// https://golang.org/pkg/internal/unsafeheader/#Slice
	var sliceMem = struct { // slice memory layout
		addr unsafe.Pointer
		len  int
		cap  int
	}{unsafe.Pointer(ptr), int(sz), int(sz)}

	return *(*[]byte)(unsafe.Pointer(&sliceMem)) // convert to slice itself
}

// Converts *uint16 to []uint16, with the given length.
func PtrToSliceUint16(ptr *uint16, length int) []uint16 {
	// https://stackoverflow.com/a/43592538
	// https://golang.org/pkg/internal/unsafeheader/#Slice
	var sliceMem = struct { // slice memory layout
		addr unsafe.Pointer
		len  int
		cap  int
	}{unsafe.Pointer(ptr), length, length}

	return *(*[]uint16)(unsafe.Pointer(&sliceMem)) // convert to slice itself
}

// Converts **uint16 to []*uint16, with the given length.
func PtrToSliceUint16Ptr(ptr **uint16, sz int) []*uint16 {
	// https://stackoverflow.com/a/43592538
	// https://golang.org/pkg/internal/unsafeheader/#Slice
	var sliceMem = struct { // slice memory layout
		addr unsafe.Pointer
		len  int
		cap  int
	}{unsafe.Pointer(ptr), int(sz), int(sz)}

	return *(*[]*uint16)(unsafe.Pointer(&sliceMem)) // convert to slice itself
}

// "&He && she" becomes "He & she".
func RemoveAccelAmpersands(text string) string {
	runes := []rune(text)
	buf := strings.Builder{}
	buf.Grow(len(text)) // prealloc for performance

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
