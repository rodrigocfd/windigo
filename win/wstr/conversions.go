//go:build windows

package wstr

import (
	"syscall"
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"
)

const (
	_MASK_X      = 0b00111111
	_MASK_3      = 0b00001111
	_SURR_1      = 0xd800
	_SURR_3      = 0xe000
	_SURR_SELF   = 0x10000
	_MAX_RUNE    = '\U0010FFFF'
	_REPLAC_CHAR = '\uFFFD'
)

// Converts multiple Go strings into multiple null-terminated Windows UTF-16
// strings, with an additional null terminator. Writes directly to dest,
// assuming it's properly allocated.
//
// Panics on error.
func GoArrToWinBuf(strs []string, dest []uint16) {
	idxNextCh := 0
	for _, str := range strs {
		GoToWinBuf(str, dest[idxNextCh:])
		idxNextCh += utf8.RuneCountInString(str) + 1
	}
	dest[idxNextCh] = 0x0000 // additional terminating null
}

// Converts multiple Go strings into multiple null-terminated Windows UTF-16
// strings, with an additional null terminator. Returns a heap-allocated
// *uint16.
//
// Panics on error.
func GoArrToWinPtr(strs ...string) *uint16 {
	buf := GoArrToWinSlice(strs...)
	return &buf[0]
}

// Converts multiple Go strings into multiple null-terminated Windows UTF-16
// strings, with an additional null terminator. Returns a heap-allocated
// []uint16.
//
// Panics on error.
func GoArrToWinSlice(strs ...string) []uint16 {
	numChars := 1 // count additional terminating null
	for _, str := range strs {
		numChars += utf8.RuneCountInString(str) + 1
	}
	buf := make([]uint16, numChars)
	GoArrToWinBuf(strs, buf)
	return buf
}

// Converts a Go string into a null-terminated Windows UTF-16 string. Writes
// directly to dest, assuming it's properly allocated.
//
// Panics on error.
func GoToWinBuf(str string, dest []uint16) {
	curPos := 0

	for i := 0; i < len(str); { // adapted from wtf8_windows.go
		if str[i] == 0 {
			panic("The Go string contains a NUL byte, therefore it is invalid.")
		}

		r, size := utf8.DecodeRuneInString(str[i:])
		if r == utf8.RuneError {
			if sc := str[i:]; len(sc) >= 3 && sc[0] == 0xed && 0xa0 <= sc[1] && sc[1] <= 0xbf && 0x80 <= sc[2] && sc[2] <= 0xbf {
				r = rune(sc[0]&_MASK_3)<<12 + rune(sc[1]&_MASK_X)<<6 + rune(sc[2]&_MASK_X)

				dest[curPos] = uint16(r)
				curPos++

				i += 3
				continue
			}
		}
		i += size

		// Adapted from utf16.AppendRune().
		switch {
		case 0 <= r && r < _SURR_1, _SURR_3 <= r && r < _SURR_SELF:
			dest[curPos] = uint16(r)
			curPos++
		case _SURR_SELF <= r && r <= _MAX_RUNE:
			r1, r2 := utf16.EncodeRune(r)
			dest[curPos] = uint16(r1)
			dest[curPos+1] = uint16(r2)
			curPos += 2
		default:
			dest[curPos] = _REPLAC_CHAR
			curPos++
		}
	}

	dest[curPos] = 0x0000 // terminating null
}

// Converts a Go string into a null-terminated Windows UTF-16 string. Returns a
// heap-allocated *uint16.
//
// Panics on error.
func GoToWinPtr(s string) *uint16 {
	buf := GoToWinSlice(s)
	return &buf[0]
}

// Converts a Go string into a null-terminated Windows UTF-16 string. Returns a
// heap-allocated []uint16.
//
// Panics on error.
func GoToWinSlice(str string) []uint16 {
	numChars := utf8.RuneCountInString(str) + 1
	buf := make([]uint16, numChars)
	GoToWinBuf(str, buf)
	return buf
}

// Converts a pointer to a multi null-terminated Windows UTF-16 string into a Go
// []string.
//
// Source string must have 2 terminating nulls.
func WinArrPtrToGo(p *uint16) []string {
	values := make([]string, 0)
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
			values = append(values, WinSliceToGo(slice))

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
func WinPtrToGo(p *uint16) string {
	// Copied from syscall_windows.go, utf16PtrToString() private function.
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
	return WinSliceToGo(slice)
}

// Converts a []uint16 with a Windows UTF-16 string, null-terminated or not,
// into a Go string.
//
// Wraps [syscall.UTF16ToString].
func WinSliceToGo(s []uint16) string {
	trimAt := len(s)
	for i, ch := range s {
		if ch == 0 { // we found a terminating null
			trimAt = i
			break
		}
	}
	return syscall.UTF16ToString(s[:trimAt])
}
