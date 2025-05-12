//go:build windows

package utl

import (
	"strings"
	"unicode/utf8"
)

// "&He && she" becomes "He & she".
func RemoveAccelAmpersands(text string) string {
	runes := []rune(text)
	var buf strings.Builder
	buf.Grow(len(runes)) // prealloc for performance

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

// Removes anchor markdowns from SysLink texts. Will fail if '<' char is present
// without delimiting a tag.
func RemoveHtmlAnchor(text string) string {
	var buf strings.Builder
	buf.Grow(utf8.RuneCountInString(text))

	withinAnchor := false
	for _, ch := range text {
		if withinAnchor {
			if ch == '>' {
				withinAnchor = false
				continue
			}
		}
		if ch == '<' {
			withinAnchor = true
			continue
		}
		if !withinAnchor {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}
