/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	c "wingows/consts"
)

// Native static control (label).
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type Static struct {
	controlNativeBase
}

// Optional; returns a Button with a specific control ID.
func MakeStatic(ctrlId c.ID) Static {
	return Static{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *Static) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles c.WS_EX, styles c.WS, staStyles c.SS) *Static {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles, "STATIC", text,
		styles|c.WS(staStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Position will be adjusted to the current system DPI.
// The size will be calculated to fit the text exactly.
func (me *Static) CreateLText(parent Window, x, y int32, text string) *Static {
	x, y, _, _ = multiplyByDpi(x, y, 0, 0)
	cx, cy := calcIdealSize(parent.Hwnd(), text, true)

	me.controlNativeBase.create(c.WS_EX(0), "STATIC", text,
		c.WS_CHILD|c.WS_GROUP|c.WS_VISIBLE|c.WS(c.SS_LEFT), x, y, cx, cy, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Sets the text and resizes the control to fit the text exactly.
func (me *Static) SetText(text string) {
	hasAccel := (c.SS(me.Hwnd().GetStyle()) & c.SS_NOPREFIX) == 0
	cx, cy := calcIdealSize(me.Hwnd().GetParent(), text, hasAccel)

	me.Hwnd().SetWindowPos(c.SWP_HWND(0), 0, 0, cx, cy,
		c.SWP_NOZORDER|c.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
}

// Returns the text without the accelerator ampersands.
// For example: "&He && she" is returned as "He & she".
// Use HWND().GetWindowText() to retrieve the full text, with ampersands.
func (me *Static) Text() string {
	return removeAccelAmpersands(me.Hwnd().GetWindowText())
}
