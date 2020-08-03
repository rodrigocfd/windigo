/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
	"strings"
	"wingows/win"
)

// "&He && she" becomes "He & she".
func removeAccelAmpersands(text string) string {
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

// Generates a string from all fields, excluding CbSize and LpszClassName, that
// uniquely identifies a WNDCLASSEX object.
func wndclassex_hash(wcx *win.WNDCLASSEX) string {
	return fmt.Sprintf("%x.%x.%x.%x.%x.%x.%x.%x.%x.%x",
		wcx.Style, wcx.LpfnWndProc, wcx.CbClsExtra, wcx.CbWndExtra,
		wcx.HInstance, wcx.HIcon, wcx.HCursor, wcx.HbrBackground,
		wcx.LpszMenuName, wcx.HIconSm)
}
