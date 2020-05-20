/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"wingows/api"
	c "wingows/consts"
)

// Any child control with HWND and ID.
type Control interface {
	Window
	CtrlId() c.ID
}

// Any window with a HWND handle.
type Window interface {
	Hwnd() api.HWND
}

//------------------------------------------------------------------------------

func multiplyByDpi(x, y int32, cx, cy uint32) (int32, int32, uint32, uint32) {
	if globalDpi.X == 0 {
		dc := api.HWND(0).GetDC()
		globalDpi.X = dc.GetDeviceCaps(c.GDC_LOGPIXELSX)
		globalDpi.Y = dc.GetDeviceCaps(c.GDC_LOGPIXELSY)
	}
	return api.MulDiv(x, globalDpi.X, 96),
		api.MulDiv(y, globalDpi.Y, 96),
		uint32(api.MulDiv(int32(cx), globalDpi.X, 96)),
		uint32(api.MulDiv(int32(cy), globalDpi.Y, 96))
}

//------------------------------------------------------------------------------

// Enables or disables many controls at once.
func EnableControls(enabled bool, ctrls []Control) {
	for _, ctrl := range ctrls {
		ctrl.Hwnd().EnableWindow(enabled)
	}
}
