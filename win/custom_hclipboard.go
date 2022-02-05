package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// ⚠️ hBmp will be owned by the clipboard, do not call HBITMAP.DeleteObject() anymore.
//
// Writes a bitmap to the clipboard with HCLIPBOARD.SetClipboardData().
func (hClip HCLIPBOARD) WriteBitmap(hBmp HBITMAP) {
	hClip.SetClipboardData(co.CF_BITMAP, HGLOBAL(hBmp))
}

// Writes a string to the clipboard with HCLIPBOARD.SetClipboardData().
func (hClip HCLIPBOARD) WriteString(text string) {
	text16 := Str.ToNativeSlice(text)
	text8 := unsafe.Slice((*byte)(unsafe.Pointer(&text16[0])), len(text16)*2) // direct pointer conversion

	hGlob := GlobalAlloc(co.GMEM_MOVEABLE, uint64(len(text8)))
	pMem := hGlob.GlobalLock()
	pMemSlice := unsafe.Slice((*byte)(pMem), len(text8))
	copy(pMemSlice, text8)
	hGlob.GlobalUnlock()

	hClip.SetClipboardData(co.CF_UNICODETEXT, hGlob)
}
