/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
)

// Native button control.
type Button struct {
	controlNativeBase
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *Button) Create(
	parent Window, ctrlId, x, y int32, width, height uint32,
	text string, exStyles co.WS_EX, styles co.WS, btnStyles co.BS) *Button {

	x, y, width, height = globalDpi.multiply(x, y, width, height)

	me.controlNativeBase.create(exStyles, "BUTTON", text,
		styles|co.WS(btnStyles), x, y, width, height, parent, ctrlId)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Position and width will be adjusted to the current
// system DPI. Height will be standard.
func (me *Button) CreateSimple(
	parent Window, ctrlId, x, y int32, width uint32, text string) *Button {

	return me.Create(parent, ctrlId, x, y, width, 23, text,
		co.WS_EX(0), co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.BS(0))
}

// Calls CreateWindowEx(). Creates a button with BS_DEFPUSHBUTTON style.
// Position and width will be adjusted to the current system DPI. Height will be
// standard.
func (me *Button) CreateSimpleDef(
	parent Window, ctrlId, x, y int32, width uint32, text string) *Button {

	return me.Create(parent, ctrlId, x, y, width, 23, text,
		co.WS_EX(0), co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.BS_DEFPUSHBUTTON)
}

// Returns the text without the accelerator ampersands.
// For example: "&He && she" is returned as "He & she".
// Use HWND().GetWindowText() to retrieve the full text, with ampersands.
func (me *Button) Text() string {
	return removeAccelAmpersands(me.Hwnd().GetWindowText())
}
