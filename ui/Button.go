/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/co"
)

// Native button control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#push-buttons
type Button struct {
	_ControlNativeBase
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *Button) Create(
	parent Window, ctrlId, x, y int, width, height uint,
	text string, exStyles co.WS_EX, styles co.WS, btnStyles co.BS) *Button {

	x, y, width, height = _Ui.MultiplyDpi(x, y, width, height)

	me._ControlNativeBase.create(exStyles, "BUTTON", text,
		styles|co.WS(btnStyles), x, y, width, height, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with height 23, and BS_PUSHBUTTON.
//
// Position and size will be adjusted to the current system DPI.
func (me *Button) CreateSimple(
	parent Window, ctrlId, x, y int, width uint, text string) *Button {

	return me.Create(parent, ctrlId, x, y, width, 23, text,
		co.WS_EX_NONE, co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.BS_PUSHBUTTON)
}

// Calls CreateWindowEx() with height 23, and BS_DEFPUSHBUTTON.
//
// Position and size will be adjusted to the current system DPI.
func (me *Button) CreateSimpleDef(
	parent Window, ctrlId, x, y int, width uint, text string) *Button {

	return me.Create(parent, ctrlId, x, y, width, 23, text,
		co.WS_EX_NONE, co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.BS_DEFPUSHBUTTON)
}

// Sets the text in the button control.
//
// Syntactic sugar to Hwnd().SetWindowText().
func (me *Button) SetText(text string) *Button {
	me.Hwnd().SetWindowText(text)
	return me
}

// Returns the text without the accelerator ampersands, for example:
// "&He && she" is returned as "He & she".
//
// Use Hwnd().GetWindowText() to retrieve the raw text, with unparsed
// accelerator ampersands.
func (me *Button) Text() string {
	return _Ui.RemoveAccelAmpersands(me.Hwnd().GetWindowText())
}
