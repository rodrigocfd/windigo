//go:build windows

package wstr

import (
	"unicode/utf16"
	"unsafe"
)

// Converts multiple Go strings into multiple null-terminated UTF-16 strings,
// with a double null terminator. Writes to a previously allocated buffer. If
// the buffer isn't long enough, the output strings will be truncated.
//
// Returns the number of uint16 words written, including the double terminating
// null.
func EncodeArrToBuf(dest []uint16, strs ...string) int {
	if len(dest) == 0 {
		return 0
	}

	totWritten := 1 // count double terminating null
	for _, str := range strs {
		szDest := len(dest)
		if szDest == 1 { // we need a second terminating null
			break
		}

		numWritten := EncodeToBuf(dest[:szDest-1], str) // will write a terminating null
		dest = dest[numWritten:]
		totWritten += numWritten
	}
	dest[0] = 0x0000 // double terminating null
	return totWritten
}

// Converts multiple Go strings into multiple null-terminated UTF-16 strings,
// with a double null terminator. Returns a new heap-allocated *uint16.
func EncodeArrToPtr(strs ...string) *uint16 {
	buf := EncodeArrToSlice(strs...)
	return unsafe.SliceData(buf)
}

// Converts multiple Go strings into multiple null-terminated UTF-16 strings,
// with a double null terminator. Returns a new heap-allocated []uint16.
func EncodeArrToSlice(strs ...string) []uint16 {
	numWords := 1 // count double terminating null
	for _, s := range strs {
		numWords += CountUtf16Len(s) + 1 // count terminating null
	}

	buf := make([]uint16, numWords)
	EncodeArrToBuf(buf, strs...)
	return buf
}

// Converts a Go string into a null-terminated UTF-16 string, writing to a
// previously allocated buffer. If the buffer isn't long enough, the output
// string will be truncated.
//
// Returns the number of uint16 words written, including the terminating null.
//
// Adapted from [utf16.Encode]; performs no allocations.
//
// Example:
//
//	buf := make([]uint16, 10)
//	wstr.EncodeToBuf(buf, "abc")
func EncodeToBuf(dest []uint16, s string) int {
	const (
		_SURR_SELF   = 0x10000
		_REPLAC_CHAR = '\uFFFD'
	)

	szDest := len(dest)
	if szDest == 0 {
		return 0
	}

	idx := 0
EachRune:
	for _, ch := range s {
		if idx >= szDest-1 {
			break // truncate to prevent buffer overrun
		}

		switch utf16.RuneLen(ch) {
		case 1: // normal rune
			dest[idx] = uint16(ch)
			idx++
		case 2: // needs surrogate sequence
			if idx+1 >= szDest-1 {
				break EachRune // no room for surrogate pair
			} else {
				r1, r2 := utf16.EncodeRune(ch)
				dest[idx] = uint16(r1)
				dest[idx+1] = uint16(r2)
				idx += 2
			}
		default: // cannot be properly encoded
			dest[idx] = uint16(_REPLAC_CHAR)
			idx++
		}
	}
	dest[idx] = 0x0000 // terminating null
	return idx + 1
}

// Converts a Go string into a null-terminated UTF-16 string. Returns a new
// heap-allocated *uint16.
func EncodeToPtr(s string) *uint16 {
	buf := EncodeToSlice(s)
	return unsafe.SliceData(buf)
}

// Converts a Go string into a null-terminated UTF-16 string. Returns a new
// heap-allocated []uint16.
func EncodeToSlice(s string) []uint16 {
	buf := make([]uint16, CountUtf16Len(s)+1) // count terminating null
	EncodeToBuf(buf, s)
	return buf
}
