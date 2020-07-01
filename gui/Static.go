/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
)

// Native static control (label).
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type Static struct {
	controlNativeBase
}

// Optional; returns a Static control with a specific control ID.
func MakeStatic(ctrlId co.ID) Static {
	return Static{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *Static) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles co.WS_EX, styles co.WS, staStyles co.SS) *Static {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles, "STATIC", text,
		styles|co.WS(staStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Position will be adjusted to the current system DPI.
// The size will be calculated to fit the text exactly.
func (me *Static) CreateLText(parent Window, x, y int32, text string) *Static {
	x, y, _, _ = multiplyByDpi(x, y, 0, 0)
	cx, cy := calcIdealSize(parent.Hwnd(), text, true)

	me.controlNativeBase.create(co.WS_EX(0), "STATIC", text,
		co.WS_CHILD|co.WS_VISIBLE|co.WS(co.SS_LEFT),
		x, y, cx, cy, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Sets the text and resizes the control to fit the text exactly.
func (me *Static) SetText(text string) {
	hasAccel := (co.SS(me.Hwnd().GetStyle()) & co.SS_NOPREFIX) == 0
	cx, cy := calcIdealSize(me.Hwnd().GetParent(), text, hasAccel)

	me.Hwnd().SetWindowPos(co.SWP_HWND(0), 0, 0, cx, cy,
		co.SWP_NOZORDER|co.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
}

// Returns the text without the accelerator ampersands.
// For example: "&He && she" is returned as "He & she".
// Use HWND().GetWindowText() to retrieve the full text, with ampersands.
func (me *Static) Text() string {
	return removeAccelAmpersands(me.Hwnd().GetWindowText())
}
