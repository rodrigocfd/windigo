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
	if globalDpi.X == 0 { // not initialized yet?
		dc := win.HWND(0).GetDC()
		globalDpi.X = dc.GetDeviceCaps(co.GDC_LOGPIXELSX)
		globalDpi.Y = dc.GetDeviceCaps(co.GDC_LOGPIXELSY)
	}
	if x != 0 || y != 0 {
		x = win.MulDiv(x, globalDpi.X, 96)
		y = win.MulDiv(y, globalDpi.Y, 96)
	}
	if cx != 0 || cy != 0 {
		cx = uint32(win.MulDiv(int32(cx), globalDpi.X, 96))
		cy = uint32(win.MulDiv(int32(cy), globalDpi.Y, 96))
	}
	return x, y, cx, cy
}
