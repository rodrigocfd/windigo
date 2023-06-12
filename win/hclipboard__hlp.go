//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// This helper method writes a bitmap to the clipboard with
// HCLIPBOARD.SetClipboardData().
//
// ⚠️ hBmp will be owned by the clipboard, do not call HBITMAP.DeleteObject()
// anymore.
func (hClip HCLIPBOARD) WriteBitmap(hBmp HBITMAP) {
	hClip.SetClipboardData(co.CF_BITMAP, HGLOBAL(hBmp))
}

// This helper method writes a string to the clipboard with
// HCLIPBOARD.SetClipboardData().
func (hClip HCLIPBOARD) WriteString(text string) {
	hMem := GlobalAllocStr(co.GMEM_MOVEABLE, text)
	hClip.SetClipboardData(co.CF_UNICODETEXT, hMem) // pass pointer ownership
}
