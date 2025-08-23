//go:build windows

package wstr

import (
	"unicode/utf16"
	"unsafe"
)

// Converts a pointer to a multi null-terminated Windows UTF-16 string into a Go
// []string.
//
// Source string must have 2 terminating nulls.
func DecodeArrPtr(p *uint16) []string {
	values := make([]string, 0, 1)
	if p == nil {
		return values
	}

	pRun := unsafe.Pointer(p)
	sLen := 0
	for {
		if *(*uint16)(pRun) == 0 { // terminating null found
			if sLen == 0 {
				break // two terminating nulls
			}

			slice := unsafe.Slice(p, sLen) // create slice without terminating null
			values = append(values, string(utf16.Decode(slice)))

			pRun = unsafe.Add(pRun, unsafe.Sizeof(*p)) // pRun++
			p = (*uint16)(pRun)
			sLen = 0

		} else {
			pRun = unsafe.Add(pRun, unsafe.Sizeof(*p)) // pRun++
			sLen++
		}
	}

	return values
}

// Converts a pointer to a null-terminated Windows UTF-16 string into a Go
// string.
func DecodePtr(p *uint16) string {
	// Adapted from syscall_windows.go, utf16PtrToString() private function.
	if p == nil {
		return ""
	}

	// Find null terminator.
	pRun := unsafe.Pointer(p)
	sLen := 0
	for *(*uint16)(pRun) != 0 {
		pRun = unsafe.Add(pRun, unsafe.Sizeof(*p)) // pRun++
		sLen++
	}

	slice := unsafe.Slice(p, sLen) // create slice without terminating null
	return string(utf16.Decode(slice))
}

// Converts a []uint16 with a Windows UTF-16 string, null-terminated or not,
// into a Go string.
//
// Wraps [utf16.Decode].
func DecodeSlice(str []uint16) string {
	for idx, ch := range str {
		if ch == 0x0000 {
			return string(utf16.Decode(str[:idx])) // stop before first terminating null
		}
	}
	return string(utf16.Decode(str))
}

// Converts multiple Go strings into multiple null-terminated Windows UTF-16
// strings, with an additional null terminator. Writes directly to dest; if it's
// not long enough the string will be truncated.
func EncodeArrToBuf(strs []string, dest []uint16) {
	idxNextCh := 0
	for _, str := range strs {
		EncodeToBuf(str, dest[idxNextCh:])
		idxNextCh += len([]rune(str)) + 1
	}
	dest[idxNextCh] = 0x0000 // additional terminating null
}

// Converts multiple Go strings into multiple null-terminated Windows UTF-16
// strings, with an additional null terminator. Returns a heap-allocated
// *uint16.
func EncodeArrToPtr(strs ...string) *uint16 {
	buf := EncodeArrToSlice(strs...)
	return &buf[0]
}

// Converts multiple Go strings into multiple null-terminated Windows UTF-16
// strings, with an additional null terminator. Returns a heap-allocated
// []uint16.
func EncodeArrToSlice(strs ...string) []uint16 {
	numChars := 1 // count additional terminating null
	for _, str := range strs {
		numChars += len([]rune(str)) + 1
	}
	buf := make([]uint16, numChars)
	EncodeArrToBuf(strs, buf)
	return buf
}

// Converts a Go string into a null-terminated Windows UTF-16 string. Writes
// directly to dest; if it's not long enough the string will be truncated.
func EncodeToBuf(str string, dest []uint16) {
	rawEncodeToBuf([]rune(str), dest)
}

// Converts a Go string into a null-terminated Windows UTF-16 string. Returns a
// heap-allocated *uint16.
func EncodeToPtr(str string) *uint16 {
	buf := EncodeToSlice(str)
	return &buf[0]
}

// Converts a Go string into a null-terminated Windows UTF-16 string. Returns a
// heap-allocated []uint16.
func EncodeToSlice(str string) []uint16 {
	runes := []rune(str)
	buf := make([]uint16, len(runes)+1) // room for terminating null
	rawEncodeToBuf(runes, buf)
	return buf
}

// Adapted from [utf16.Encode].
func rawEncodeToBuf(s []rune, dest []uint16) {
	const (
		_SURR_SELF   = 0x10000
		_REPLAC_CHAR = '\uFFFD'
	)

	n := len(s)
	for _, v := range s {
		if v >= _SURR_SELF {
			n++ // count each extra surrogate char we'll need
		}
	}

	szDest := len(dest)
	n = 0
	for _, ch := range s {
		if n >= szDest-1 {
			break // truncate to prevent buffer overrun
		}

		switch utf16.RuneLen(ch) {
		case 1: // normal rune
			dest[n] = uint16(ch)
			n++
		case 2: // needs surrogate sequence
			r1, r2 := utf16.EncodeRune(ch)
			dest[n] = uint16(r1)
			dest[n+1] = uint16(r2)
			n += 2
		default:
			dest[n] = uint16(_REPLAC_CHAR)
			n++
		}
	}
	dest[n] = 0x0000 // terminating null
}
