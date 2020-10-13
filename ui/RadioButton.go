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

// Native radio button control.
//
// Prefer using RadioGroup instead of manually managing each radio button.
type RadioButton struct {
	_ControlNativeBase
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *RadioButton) Create(
	parent Window, ctrlId int, pos Pos, size Size,
	text string, exStyles co.WS_EX, styles co.WS, btnStyles co.BS) *RadioButton {

	_Ui.MultiplyDpi(&pos, &size)
	me._ControlNativeBase.create(exStyles, "BUTTON", text, // radio button is, in fact, a button
		styles|co.WS(btnStyles), pos, size, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with BS_AUTORADIOBUTTON | WS_GROUP.
//
// Call this method to create the first radio button of a group.
//
// Position will be adjusted to the current system DPI. The size will be
// calculated to fit the text exactly.
func (me *RadioButton) CreateFirst(
	parent Window, ctrlId int, pos Pos, text string) *RadioButton {

	return me.createAutoSize(parent, ctrlId, pos, text,
		co.WS_GROUP|co.WS_TABSTOP|co.WS(co.BS_AUTORADIOBUTTON))
}

// Calls CreateWindowEx() with BS_AUTORADIOBUTTON.
//
// Position will be adjusted to the current system DPI. The size will be
// calculated to fit the text exactly.
func (me *RadioButton) CreateSubsequent(
	parent Window, ctrlId int, pos Pos, text string) *RadioButton {

	return me.createAutoSize(parent, ctrlId, pos, text,
		co.WS(co.BS_AUTORADIOBUTTON))
}

// Tells if the current state is BST_CHECKED.
func (me *RadioButton) IsChecked() bool {
	return co.BST(me.Hwnd().
		SendMessage(co.WM(co.BM_GETCHECK), 0, 0)) == co.BST_CHECKED
}

// Sets the state to BST_CHECKED.
//
// The currently checked radio button won't be cleared. Prefer using
// RadioGroup.SetCheck().
func (me *RadioButton) SetCheck() *RadioButton {
	me.Hwnd().
		SendMessage(co.WM(co.BM_SETCHECK), win.WPARAM(co.BST_CHECKED), 0)
	return me
}

// Sets the state to BST_CHECKED and emulates the user click.
func (me *RadioButton) SetCheckAndTrigger() *RadioButton {
	me.SetCheck()
	me.Hwnd().SendMessage(co.WM_COMMAND,
		win.MakeWParam(uint16(me.Hwnd().GetDlgCtrlID()), 0),
		win.LPARAM(me.Hwnd()))
	return me
}

// Sets the text, and resizes the control to fit it exactly.
//
// To set the text without resizing the control, use Hwnd().SetWindowText().
func (me *RadioButton) SetText(text string) *RadioButton {
	size := me.calcIdealSize(me.Hwnd().GetParent(), text)
	me.Hwnd().SetWindowPos(co.SWP_HWND_NONE,
		0, 0, int32(size.Cx), int32(size.Cy),
		co.SWP_NOZORDER|co.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
	return me
}

// Returns the text without the accelerator ampersands, for example:
// "&He && she" is returned as "He & she".
//
// Use Hwnd().GetWindowText() to retrieve the raw text, with unparsed
// accelerator ampersands.
func (me *RadioButton) Text() string {
	return _Ui.RemoveAccelAmpersands(me.Hwnd().GetWindowText())
}

func (me *RadioButton) calcIdealSize(
	hReferenceDc win.HWND, text string) Size {

	boundBox := _Ui.CalcTextBoundBox(hReferenceDc, text, true)
	boundBox.Cx += uint(win.GetSystemMetrics(co.SM_CXMENUCHECK)) +
		uint(win.GetSystemMetrics(co.SM_CXEDGE)) // https://stackoverflow.com/a/1165052/6923555

	cyCheck := uint(win.GetSystemMetrics(co.SM_CYMENUCHECK))
	if cyCheck > boundBox.Cy {
		boundBox.Cy = cyCheck // if the check is taller than the font, use its height
	}

	return boundBox
}

func (me *RadioButton) createAutoSize(
	parent Window, ctrlId int, pos Pos,
	text string, otherStyles co.WS) *RadioButton {

	_Ui.MultiplyDpi(&pos, nil)
	size := me.calcIdealSize(parent.Hwnd(), text)

	me._ControlNativeBase.create(co.WS_EX_NONE, "BUTTON", text,
		co.WS_CHILD|co.WS_VISIBLE|otherStyles,
		pos, size, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}
