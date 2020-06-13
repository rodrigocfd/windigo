/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"strings"
	"wingows/co"
	"wingows/win"
)

// Any child control with HWND and ID.
type Control interface {
	Window
	CtrlId() co.ID
}

// Any window with a HWND handle.
type Window interface {
	Hwnd() win.HWND
}

//------------------------------------------------------------------------------

// Calculates the bound rectangle to fit the text with current system font.
func calcIdealSize(hReferenceDc win.HWND, text string,
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
