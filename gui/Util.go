/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"strings"
	"wingows/co"
	"wingows/win"
)

type _UtilT struct {
	globalDpi win.POINT
}

// Internal package utilities.
var _Util _UtilT

// Multiplies position and size by current DPI factor.
func (me *_UtilT) MultiplyDpi(x, y int32,
	cx, cy uint32) (int32, int32, uint32, uint32) {

	if me.globalDpi.X == 0 { // not initialized yet?
		dc := win.HWND(0).GetDC()
		me.globalDpi.X = dc.GetDeviceCaps(co.GDC_LOGPIXELSX)
		me.globalDpi.Y = dc.GetDeviceCaps(co.GDC_LOGPIXELSY)
		win.HWND(0).ReleaseDC(dc)
	}

	if x != 0 || y != 0 {
		x = win.MulDiv(x, me.globalDpi.X, 96)
		y = win.MulDiv(y, me.globalDpi.Y, 96)
	}
	if cx != 0 || cy != 0 {
		cx = uint32(win.MulDiv(int32(cx), me.globalDpi.X, 96))
		cy = uint32(win.MulDiv(int32(cy), me.globalDpi.Y, 96))
	}
	return x, y, cx, cy
}

// "&He && she" becomes "He & she".
func (_UtilT) RemoveAccelAmpersands(text string) string {
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
