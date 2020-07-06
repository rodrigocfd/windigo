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

type systemDpi struct {
	x, y int32
}

var globalDpi = systemDpi{x: 0, y: 0}

func (me *systemDpi) init() {
	if me.x == 0 { // not initialized yet?
		dc := win.HWND(0).GetDC()
		me.x = dc.GetDeviceCaps(co.GDC_LOGPIXELSX)
		me.y = dc.GetDeviceCaps(co.GDC_LOGPIXELSY)
	}
}

// Multiplies position and size by current DPI factor.
func (me *systemDpi) multiply(x, y int32,
	cx, cy uint32) (int32, int32, uint32, uint32) {

	me.init()
	if x != 0 || y != 0 {
		x = win.MulDiv(x, me.x, 96)
		y = win.MulDiv(y, me.y, 96)
	}
	if cx != 0 || cy != 0 {
		cx = uint32(win.MulDiv(int32(cx), me.x, 96))
		cy = uint32(win.MulDiv(int32(cy), me.y, 96))
	}
	return x, y, cx, cy
}
