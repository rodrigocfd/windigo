/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"strings"
	"wingows/api"
	"wingows/co"
)

// Any child control with HWND and ID.
type Control interface {
	Window
	CtrlId() co.ID
}

// Any window with a HWND handle.
type Window interface {
	Hwnd() api.HWND
}

//------------------------------------------------------------------------------

var globalDpi = api.POINT{X: 0, Y: 0} // set once by multiplyByDpi

// Multiplies position and size by current DPI factor.
func multiplyByDpi(x, y int32, cx, cy uint32) (int32, int32, uint32, uint32) {
	if globalDpi.X == 0 {
		dc := api.HWND(0).GetDC()
		globalDpi.X = dc.GetDeviceCaps(co.GDC_LOGPIXELSX)
		globalDpi.Y = dc.GetDeviceCaps(co.GDC_LOGPIXELSY)
	}
	return api.MulDiv(x, globalDpi.X, 96),
		api.MulDiv(y, globalDpi.Y, 96),
		uint32(api.MulDiv(int32(cx), globalDpi.X, 96)),
		uint32(api.MulDiv(int32(cy), globalDpi.Y, 96))
}

// Calculates the bound rectangle to fit the text with current system font.
func calcIdealSize(hReferenceDc api.HWND, text string,
	considerAccelerators bool) (uint32, uint32) {

	isTextEmpty := false
	if len(text) == 0 {
		isTextEmpty = true
		text = "Pj" // just a placeholder to get the text height
	}

	if considerAccelerators {
		text = removeAccelAmpersands(text)
	}

	parentDc := hReferenceDc.GetDC()
	cloneDc := parentDc.CreateCompatibleDC()
	prevFont := cloneDc.SelectObjectFont(globalUiFont.Hfont()) // system font; already adjusted to current DPI
	bounds := cloneDc.GetTextExtentPoint32(text)
	cloneDc.SelectObjectFont(prevFont)
	cloneDc.DeleteDC()
	hReferenceDc.ReleaseDC(parentDc)

	if isTextEmpty {
		bounds.Cx = 0 // if no text was given, return just the height
	}
	return uint32(bounds.Cx), uint32(bounds.Cy)
}

// "&He && she" becomes "He & she".
func removeAccelAmpersands(text string) string {
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
