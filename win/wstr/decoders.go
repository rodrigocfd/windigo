//go:build windows

package wstr

import (
	"unicode/utf16"
	"unsafe"
)

// Converts a pointer to a multi null-terminated UTF-16 string into a Go
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

// Converts a pointer to a null-terminated UTF-16 string into a Go string.
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

// Converts a []uint16 with an UTF-16 string, null-terminated or not, into a Go
// string.
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
