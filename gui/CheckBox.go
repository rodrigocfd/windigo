/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
	"wingows/win"
)

// Native check box control.
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type CheckBox struct {
	controlNativeBase
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *CheckBox) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles co.WS_EX, styles co.WS, btnStyles co.BS) *CheckBox {

	x, y, width, height = globalDpi.multiply(x, y, width, height)

	me.controlNativeBase.create(exStyles, "BUTTON", text, // check box is, in fact, a button
		styles|co.WS(btnStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Creates a check box with BS_AUTO3STATE style.
// Position will be adjusted to the current system DPI. The size will be
// calculated to fit the text exactly.
func (me *CheckBox) CreateThreeState(parent Window, x, y int32,
	text string) *CheckBox {

	return me.createAutoSize(parent, x, y, text, co.BS_AUTO3STATE)
}

// Calls CreateWindowEx(). Creates a check box with BS_AUTOCHECKBOX style.
// Position will be adjusted to the current system DPI. The size will be
// calculated to fit the text exactly.
func (me *CheckBox) CreateTwoState(parent Window, x, y int32,
	text string) *CheckBox {

	return me.createAutoSize(parent, x, y, text, co.BS_AUTOCHECKBOX)
}

func (me *CheckBox) IsChecked() bool {
	return me.State() == co.BST_CHECKED
}

func (me *CheckBox) SetCheck() *CheckBox {
	return me.SetState(co.BST_CHECKED)
}

// A BS_AUTOCHECKBOX can be only checked or unchecked, a BS_AUTO3STATE can also
// be indeterminate.
func (me *CheckBox) SetState(state co.BST) *CheckBox {
	me.Hwnd().SendMessage(co.WM(co.BM_SETCHECK), win.WPARAM(state), 0)
	return me
}

// SetWindowText() doesn't resize the control to fit the text. This method
// resizes the control to fit the text exactly.
func (me *CheckBox) SetText(text string) *CheckBox {
	cx, cy := me.calcIdealSize(me.Hwnd().GetParent(), text)
	me.Hwnd().SetWindowPos(co.SWP_HWND(0), 0, 0, cx, cy,
		co.SWP_NOZORDER|co.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
	return me
}

// A BS_AUTOCHECKBOX can be only checked or unchecked, a BS_AUTO3STATE can also
// be indeterminate.
func (me *CheckBox) State() co.BST {
	return co.BST(me.Hwnd().SendMessage(co.WM(co.BM_GETCHECK), 0, 0))
}

// Returns the text without the accelerator ampersands.
// For example: "&He && she" is returned as "He & she".
// Use HWND().GetWindowText() to retrieve the full text, with ampersands.
func (me *CheckBox) Text() string {
	return removeAccelAmpersands(me.Hwnd().GetWindowText())
}

func (me *CheckBox) calcIdealSize(hReferenceDc win.HWND,
	text string) (uint32, uint32) {

	cx, cy := calcTextBoundBox(hReferenceDc, text, true)
	cx += uint32(win.GetSystemMetrics(co.SM_CXMENUCHECK)) +
		uint32(win.GetSystemMetrics(co.SM_CXEDGE)) // https://stackoverflow.com/a/1165052/6923555

	cyCheck := uint32(win.GetSystemMetrics(co.SM_CYMENUCHECK))
	if cyCheck > cy {
		cy = cyCheck // if the check is taller than the font, use its height
	}

	return cx, cy
}

func (me *CheckBox) createAutoSize(parent Window, x, y int32,
	text string, chbxStyles co.BS) *CheckBox {

	x, y, _, _ = globalDpi.multiply(x, y, 0, 0)
	cx, cy := me.calcIdealSize(parent.Hwnd(), text)

	me.controlNativeBase.create(co.WS_EX(0), "BUTTON", text,
		co.WS_CHILD|co.WS_TABSTOP|co.WS_GROUP|co.WS_VISIBLE|co.WS(chbxStyles),
		x, y, cx, cy, parent)
	globalUiFont.SetOnControl(me)
	return me
}
