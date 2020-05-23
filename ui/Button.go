/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	c "wingows/consts"
)

// Native button control.
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type Button struct {
	controlNativeBase
}

// Optional; returns a Button with a specific control ID.
func MakeButton(ctrlId c.ID) Button {
	return Button{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *Button) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles c.WS_EX, styles c.WS, btnStyles c.BS) *Button {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles, "BUTTON", text,
		styles|c.WS(btnStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Position and width will be adjusted to the current
// system DPI. Height will be standard.
func (me *Button) CreateSimple(parent Window, x, y int32,
	width uint32, text string) *Button {

	return me.Create(parent, x, y, width, 23, text,
		c.WS_EX(0), c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.BS(0))
}

// Calls CreateWindowEx(). Creates a button with BS_DEFPUSHBUTTON style.
// Position and width will be adjusted to the current system DPI. Height will be
// standard.
func (me *Button) CreateSimpleDef(parent Window, x, y int32,
	width uint32, text string) *Button {

	return me.Create(parent, x, y, width, 23, text,
		c.WS_EX(0), c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.BS_DEFPUSHBUTTON)
}

// Returns the text without the accelerator ampersands.
// For example: "&He && she" is returned as "He & she".
// Use HWND().GetWindowText() to retrieve the full text, with ampersands.
func (me *Button) Text() string {
	return removeAccelAmpersands(me.Hwnd().GetWindowText())
}
