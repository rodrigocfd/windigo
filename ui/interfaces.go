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
	if globalDpiRatioX == -1 {
		dc := api.HWND(0).GetDC()
		globalDpiRatioX = float32(dc.GetDeviceCaps(c.GDC_LOGPIXELSX)) / 96
		globalDpiRatioY = float32(dc.GetDeviceCaps(c.GDC_LOGPIXELSY)) / 96
	}
	return int32(float32(x) * globalDpiRatioX),
		int32(float32(y) * globalDpiRatioY),
		uint32(float32(cx) * globalDpiRatioX),
		uint32(float32(cy) * globalDpiRatioY)
}

//------------------------------------------------------------------------------

// Enables or disables many controls at once.
func EnableControls(enabled bool, ctrls []Control) {
	for _, ctrl := range ctrls {
		ctrl.Hwnd().EnableWindow(enabled)
	}
}
