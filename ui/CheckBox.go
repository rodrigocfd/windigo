/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/co"
	"windigo/win"
)

// Native check box control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#check-boxes
type CheckBox struct {
	_ControlNativeBase
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *CheckBox) Create(
	parent Window, ctrlId, x, y int, width, height uint,
	text string, exStyles co.WS_EX, styles co.WS, btnStyles co.BS) *CheckBox {

	x, y, width, height = _Util.MultiplyDpi(x, y, width, height)

	me._ControlNativeBase.create(exStyles, "BUTTON", text, // check box is, in fact, a button
		styles|co.WS(btnStyles), x, y, width, height, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with BS_AUTO3STATE.
//
// Position will be adjusted to the current system DPI. The size will be
// calculated to fit the text exactly.
func (me *CheckBox) Create3State(
	parent Window, ctrlId, x, y int, text string) *CheckBox {

	return me.createAutoSize(parent, ctrlId, x, y, text, co.BS_AUTO3STATE)
}

// Calls CreateWindowEx() with BS_AUTOCHECKBOX.
//
// Position will be adjusted to the current system DPI. The size will be
// calculated to fit the text exactly.
func (me *CheckBox) Create2State(
	parent Window, ctrlId, x, y int, text string) *CheckBox {

	return me.createAutoSize(parent, ctrlId, x, y, text, co.BS_AUTOCHECKBOX)
}

// Tells if the current state is BST_CHECKED.
func (me *CheckBox) IsChecked() bool {
	return me.State() == co.BST_CHECKED
}

// Sets the current state to BST_CHECKED or BST_UNCHECKED.
func (me *CheckBox) SetCheck(isChecked bool) *CheckBox {
	state := co.BST_UNCHECKED
	if isChecked {
		state = co.BST_CHECKED
	}
	return me.SetState(state)
}

// A BS_AUTOCHECKBOX can be only checked or unchecked, but a BS_AUTO3STATE can
// also be indeterminate.
func (me *CheckBox) SetState(state co.BST) *CheckBox {
	me.Hwnd().SendMessage(co.WM(co.BM_SETCHECK), win.WPARAM(state), 0)
	return me
}

// SetWindowText() doesn't resize the control to fit the text. This method
// resizes the control to fit the text exactly.
func (me *CheckBox) SetText(text string) *CheckBox {
	cx, cy := me.calcIdealSize(me.Hwnd().GetParent(), text)
	me.Hwnd().SetWindowPos(co.SWP_HWND_NONE, 0, 0, uint32(cx), uint32(cy),
		co.SWP_NOZORDER|co.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
	return me
}

// A BS_AUTOCHECKBOX can be only checked or unchecked, a BS_AUTO3STATE can also
// be indeterminate.
func (me *CheckBox) State() co.BST {
	return co.BST(me.Hwnd().SendMessage(co.WM(co.BM_GETCHECK), 0, 0))
}

// Returns the text without the accelerator ampersands, for example:
// "&He && she" is returned as "He & she".
//
// Uses Hwnd().GetWindowText() to retrieve the raw text, with accelerator
// ampersands.
func (me *CheckBox) Text() string {
	return _Util.RemoveAccelAmpersands(me.Hwnd().GetWindowText())
}

func (me *CheckBox) calcIdealSize(hReferenceDc win.HWND,
	text string) (uint, uint) {

	cx, cy := calcTextBoundBox(hReferenceDc, text, true)
	cx += uint(win.GetSystemMetrics(co.SM_CXMENUCHECK)) +
		uint(win.GetSystemMetrics(co.SM_CXEDGE)) // https://stackoverflow.com/a/1165052/6923555

	cyCheck := uint(win.GetSystemMetrics(co.SM_CYMENUCHECK))
	if cyCheck > cy {
		cy = cyCheck // if the check is taller than the font, use its height
	}

	return cx, cy
}

func (me *CheckBox) createAutoSize(
	parent Window, ctrlId, x, y int, text string, chbxStyles co.BS) *CheckBox {

	x, y, _, _ = _Util.MultiplyDpi(x, y, 0, 0)
	cx, cy := me.calcIdealSize(parent.Hwnd(), text)

	me._ControlNativeBase.create(co.WS_EX_NONE, "BUTTON", text,
		co.WS_CHILD|co.WS_TABSTOP|co.WS_GROUP|co.WS_VISIBLE|co.WS(chbxStyles),
		x, y, cx, cy, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}