//go:build windows

package wstr

import (
	"fmt"
	"strings"
	"unicode/utf16"
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

// Counts the number of runes in the string, which can be less than the number
// of bytes returned by len(s).
//
// Adapted from [utf8.RuneCountInString].
func CountRunes(s string) uint {
	n := uint(0)
	for range s {
		n++
	}
	return n
}

// Counts the number of uint16 words in the string, when encoded to UTF-16. This
// can be greater than the number of runes, because surrogate pairs are also
// counted.
//
// This function doesn't count a terminating null, an empty string will return
// zero.
func CountUtf16Len(s string) uint {
	numWords := 0
	for _, ch := range s {
		switch utf16.RuneLen(ch) {
		case 2:
			numWords += 2 // surrogate sequence
		default:
			// If 1, is an ordinary char.
			// If -1, it's a rune which cannot be encoded to UTF-16; in such
			// case, encoding will simply place a '\uFFFD' word.
			numWords++
		}
	}
	return uint(numWords)
}

// Formats a number of bytes into KB, MB, GB, TB or PB.
func FmtBytes(numBytes uint) string {
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

// Converts the number to a string with thousand separators.
func FmtThousands(n uint) string {
	if n == 0 {
		return "0"
	}

	final := ""
	for {
		thou := n % 1000
		final = fmt.Sprintf("%03d,%s", thou, final)
		n = (n - thou) / 1000
		if n == 0 {
			return strings.TrimLeft(final[:len(final)-1], "0")
		}
	}
}

// Parses a string into an uint number.
//
// Panics if an invalid character is found.
func ParseUint(strNumber string) uint {
	nChars := uint(len(strNumber))
	var out uint

	for idx, ch := range strNumber {
		if ch < '0' || ch > '9' {
			panic(fmt.Sprintf("ParseUint: invalid character found - '%c'", ch))
		}

		out += (uint(ch) - uint('0')) * powUint(10, nChars-uint(idx)-1)
	}
	return out
}

func powUint(a, b uint) uint {
	out := uint(1)
	for i := uint(0); i < b; i++ {
		out *= a
	}
	return out
}

// Returns a new string with all diacritics removed.
func RemoveDiacritics(s string) string {
	diacs := []rune("ÁáÀàÃãÂâÄäÉéÈèÊêËëÍíÌìÎîÏïÓóÒòÕõÔôÖöÚúÙùÛûÜüÇçÅåÐðÑñØøÝýÿ")
	repls := []rune("AaAaAaAaAaEeEeEeEeIiIiIiIiOoOoOoOoOoUuUuUuUuCcAaDdNnOoYyy")

	var buf strings.Builder
	buf.Grow(int(CountRunes(s)))

EachChar:
	for _, ch := range s {
		for i, diac := range diacs {
			if ch == diac {
				buf.WriteRune(rune(repls[i]))
				continue EachChar
			}
		}
		buf.WriteRune(ch)
	}
	return buf.String()
}

// Splits the string into lines, considering LF or CR+LF.
func SplitLines(s string) []string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lineLen := len(line)
		if lineLen > 0 && line[lineLen-1] == '\r' {
			lines[i] = line[:lineLen-1] // strip away the trailing '\r'
		}
	}
	return lines
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
