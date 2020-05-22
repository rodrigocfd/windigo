/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"strings"
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

// Multiplies position and size by current DPI factor.
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

// Calculates the bound rectangle to fit the text with current system font.
func calcIdealSize(hReferenceDc api.HWND, text string,
	considerAccelerators bool) (uint32, uint32) {

	parentDc := hReferenceDc.GetDC()
	cloneDc := parentDc.CreateCompatibleDC()
	prevFont := cloneDc.SelectObjectFont(globalUiFont.Hfont()) // system font; already adjusted to current DPI

	if considerAccelerators {
		text = removeAcceleratorAmpersands(text)
	}

	bounds := cloneDc.GetTextExtentPoint32(text)
	cloneDc.SelectObjectFont(prevFont)
	cloneDc.DeleteDC()
	hReferenceDc.ReleaseDC(parentDc)

	return uint32(bounds.Cx), uint32(bounds.Cy)
}

// "&He && she" becomes "He & she".
func removeAcceleratorAmpersands(text string) string {
	buf := strings.Builder{}
	for i := 0; i < len(text)-1; i++ {
		if text[i] == '&' && text[i+1] != '&' {
			continue
		}
		buf.WriteByte(text[i])
	}
	if text[len(text)-1] != '&' {
		buf.WriteByte(text[len(text)-1])
	}
	return buf.String()
}

//------------------------------------------------------------------------------

// Enables or disables many controls at once.
func EnableControls(enabled bool, ctrls []Control) {
	for _, ctrl := range ctrls {
		ctrl.Hwnd().EnableWindow(enabled)
	}
}
