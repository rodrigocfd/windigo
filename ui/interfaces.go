/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
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

var osDpiX, osDpiY float32 = -1, -1

func multiplyByDpi(x, y int32, cx, cy uint32) (int32, int32, uint32, uint32) {
	if osDpiX == -1 {
		osDpiX = float32(api.HWND(0).GetDC().GetDeviceCaps(c.GDC_LOGPIXELSX))
		osDpiY = float32(api.HWND(0).GetDC().GetDeviceCaps(c.GDC_LOGPIXELSY))
	}
	xFac, yFac := osDpiX/96, osDpiY/96
	return int32(float32(x) * xFac), int32(float32(y) * yFac),
		uint32(float32(cx) * xFac), uint32(float32(cy) * yFac)
}

//------------------------------------------------------------------------------

// Enables or disables many controls at once.
func EnableControls(enabled bool, ctrls []Control) {
	for _, ctrl := range ctrls {
		ctrl.Hwnd().EnableWindow(enabled)
	}
}
