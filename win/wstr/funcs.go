//go:build windows

package wstr

import (
	"fmt"
	"strings"
	"syscall"
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"
)

// Compares two strings [lexographically].
//
// [lexographically]: https://stackoverflow.com/a/52831144/6923555
func Cmp(a, b string) int {
	switch {
	case a == b:
		return 0
	case a < b:
		return -1
	default:
		return 1
	}
}

// Compares two strings [lexographically], case insensitive.
//
// [lexographically]: https://stackoverflow.com/a/52831144/6923555
func CmpI(a, b string) int {
	return Cmp(strings.ToUpper(a), strings.ToUpper(b))
}

// Formats a number of bytes into KB, MB, GB, TB or PB.
func FmtBytes(numBytes uint64) string {
	switch {
	case numBytes < 1024:
		return fmt.Sprintf("%d bytes", numBytes)
	case numBytes < 1024*1024:
		return fmt.Sprintf("%.2f KB", float64(numBytes)/1024)
	case numBytes < 1024*1024*1024:
		return fmt.Sprintf("%.2f MB", float64(numBytes)/1024/1024)
	case numBytes < 1024*1024*1024*1024:
		return fmt.Sprintf("%.2f GB", float64(numBytes)/1024/1024/1024)
	case numBytes < 1024*1024*1024*1024*1024:
		return fmt.Sprintf("%.2f TB", float64(numBytes)/1024/1024/1024/1024)
	default:
		return fmt.Sprintf("%.2f PB", float64(numBytes)/1024/1024/1024/1024/1024)
	}
}

// Splits the string into lines, considering LF or CR+LF.
func SplitLines(s string) []string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lineLen := utf8.RuneCountInString(line)
		if lineLen > 0 && line[lineLen-1] == '\r' {
			lines[i] = line[:lineLen-1]
		}
	}
	return lines
}

const (
	_MASK_X      = 0b00111111
	_MASK_3      = 0b00001111
	_SURR_1      = 0xd800
	_SURR_3      = 0xe000
	_SURR_SELF   = 0x10000
	_MAX_RUNE    = '\U0010FFFF'
	_REPLAC_CHAR = '\uFFFD'
)

// Converts a Go string into a null-terminated Windows UTF-16 string, writing
// directly to dest, assuming dest is properly allocated.
//
// Panics on error.
func StrToUtf16(str string, dest []uint16) {
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

// Converts string to null-terminated *uint16. Wrapper to
// syscall.UTF16PtrFromString().
//
// Panics on error.
func StrToUtf16Ptr(s string) *uint16 {
	pstr, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		panic(fmt.Sprintf("wstr.StrToUtf16Ptr() failed \"%s\": %s", s, err))
	}
	return pstr
}

// Returns a slice over the string, starting at the given index, and with the
// given length. Counts [runes], not bytes.
//
// This function is useful if your string contains multi-byte UTF-8 chars.
//
// [runes]: https://stackoverflow.com/a/38537764/6923555
func SubstrRunes(s string, start, length uint) string {
	startStrIdx := 0
	i := uint(0)
	for j := range s {
		if i == start {
			startStrIdx = j
		}
		if i == start+length {
			return s[startStrIdx:j]
		}
		i++
	}
	return s[startStrIdx:]
}

// Returns a new string with all diacritics removed.
func RemoveDiacritics(s string) string {
	diacs := []rune("ÁáÀàÃãÂâÄäÉéÈèÊêËëÍíÌìÎîÏïÓóÒòÕõÔôÖöÚúÙùÛûÜüÇçÅåÐðÑñØøÝý")
	repls := []rune("AaAaAaAaAaEeEeEeEeIiIiIiIiOoOoOoOoOoUuUuUuUuCcAaDdNnOoYy")

	var strBuf strings.Builder
	strBuf.Grow(utf8.RuneCountInString(s))

	for _, ch := range s {
		replaced := false
		for i, diac := range diacs {
			if ch == diac {
				strBuf.WriteRune(repls[i])
				replaced = true
				break
			}
		}
		if !replaced {
			strBuf.WriteRune(ch)
		}
	}
	return strBuf.String()
}

// Converts a multi null-terminated *uint16 to []string.
//
// Source must have 2 terminating nulls.
func Utf16PtrMultiToStr(p *uint16) []string {
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
			values = append(values, Utf16SliceToStr(slice))

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

// Converts a null-terminated *uint16 to string.
//
// Copied from syscall_windows.go, utf16PtrToString() private function.
func Utf16PtrToStr(p *uint16) string {
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
	return Utf16SliceToStr(slice)
}

// Converts an []uint16 to string, which may be null-terminated or not.
//
// Wraps syscall.UTF16ToString().
func Utf16SliceToStr(s []uint16) string {
	trimAt := len(s)
	for i, ch := range s {
		if ch == 0 { // we found a terminating null
			trimAt = i
			break
		}
	}
	return syscall.UTF16ToString(s[:trimAt])
}
