/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
	"wingows/win"
)

//------------------------------------------------------------------------------

var globalDpi = win.POINT{X: 0, Y: 0} // set once by multiplyByDpi

// Multiplies position and size by current DPI factor.
func multiplyByDpi(x, y int32, cx, cy uint32) (int32, int32, uint32, uint32) {
	if globalDpi.X == 0 {
		dc := win.HWND(0).GetDC()
		globalDpi.X = dc.GetDeviceCaps(co.GDC_LOGPIXELSX)
		globalDpi.Y = dc.GetDeviceCaps(co.GDC_LOGPIXELSY)
	}
	return win.MulDiv(x, globalDpi.X, 96),
		win.MulDiv(y, globalDpi.Y, 96),
		uint32(win.MulDiv(int32(cx), globalDpi.X, 96)),
		uint32(win.MulDiv(int32(cy), globalDpi.Y, 96))
}
