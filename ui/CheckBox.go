/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"wingows/api"
	c "wingows/consts"
)

// Native check box control.
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type CheckBox struct {
	controlNativeBase
}

// Optional; returns a CheckBox with a specific control ID.
func MakeCheckBox(ctrlId c.ID) CheckBox {
	return CheckBox{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *CheckBox) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles c.WS_EX, styles c.WS, btnStyles c.BS) *CheckBox {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles, "BUTTON", text, // check box is, in fact, a button
		styles|c.WS(btnStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Creates a check box with BS_AUTO3STATE style.
// Position will be adjusted to the current system DPI. The size will be
// calculated to fit the text exactly.
func (me *CheckBox) CreateThreeState(parent Window, x, y int32,
	text string) *CheckBox {

	return me.createBase(parent, x, y, text, c.BS_AUTO3STATE)
}

// Calls CreateWindowEx(). Creates a check box with BS_AUTOCHECKBOX style.
// Position will be adjusted to the current system DPI. The size will be
// calculated to fit the text exactly.
func (me *CheckBox) CreateTwoState(parent Window, x, y int32,
	text string) *CheckBox {

	return me.createBase(parent, x, y, text, c.BS_AUTOCHECKBOX)
}

func (me *CheckBox) IsChecked() bool {
	return me.State() == c.BST_CHECKED
}

func (me *CheckBox) SetCheck() *CheckBox {
	return me.SetState(c.BST_CHECKED)
}

// A BS_AUTOCHECKBOX can be only checked or unchecked, a BS_AUTO3STATE can also
// be indeterminate.
func (me *CheckBox) SetState(state c.BST) *CheckBox {
	me.Hwnd().SendMessage(c.WM(c.BM_SETCHECK), api.WPARAM(state), 0)
	return me
}

// SetWindowText() doesn't resize the control to fit the text. This method
// resizes the control to fit the text exactly.
func (me *CheckBox) SetText(text string) *CheckBox {
	cx, cy := me.calcCheckBoxIdealSize(me.Hwnd().GetParent(), text)

	me.Hwnd().SetWindowPos(c.SWP_HWND(0), 0, 0, cx, cy,
		c.SWP_NOZORDER|c.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
	return me
}

// A BS_AUTOCHECKBOX can be only checked or unchecked, a BS_AUTO3STATE can also
// be indeterminate.
func (me *CheckBox) State() c.BST {
	return c.BST(me.Hwnd().SendMessage(c.WM(c.BM_GETCHECK), 0, 0))
}

// Returns the text without the accelerator ampersands.
// For example: "&He && she" is returned as "He & she".
// Use HWND().GetWindowText() to retrieve the full text, with ampersands.
func (me *CheckBox) Text() string {
	return removeAccelAmpersands(me.Hwnd().GetWindowText())
}

func (me *CheckBox) calcCheckBoxIdealSize(hReferenceDc api.HWND,
	text string) (uint32, uint32) {

	cx, cy := calcIdealSize(hReferenceDc, text, true)
	cx += uint32(api.GetSystemMetrics(c.SM_CXMENUCHECK)) +
		uint32(api.GetSystemMetrics(c.SM_CXEDGE)) // https://stackoverflow.com/a/1165052/6923555

	cyCheck := uint32(api.GetSystemMetrics(c.SM_CYMENUCHECK))
	if cyCheck > cy {
		cy = cyCheck // if the check is taller than the font, use its height
	}

	return cx, cy
}

func (me *CheckBox) createBase(parent Window, x, y int32,
	text string, chbxStyles c.BS) *CheckBox {

	x, y, _, _ = multiplyByDpi(x, y, 0, 0)
	cx, cy := me.calcCheckBoxIdealSize(parent.Hwnd(), text)

	me.controlNativeBase.create(c.WS_EX(0), "BUTTON", text,
		c.WS_CHILD|c.WS_GROUP|c.WS_VISIBLE|c.WS(chbxStyles),
		x, y, cx, cy, parent)
	globalUiFont.SetOnControl(me)
	return me
}
