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

// Native radio button control.
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type RadioButton struct {
	controlNativeBase
}

// Helper function to retrieve the index of the selected radio button.
func GetCheckedRadio(radioGroup []RadioButton) int {
	for i := range radioGroup {
		if radioGroup[i].IsChecked() {
			return i
		}
	}
	return -1 // no checked one
}

// Optional; returns a RadioButton with a specific control ID.
func MakeRadioButton(ctrlId c.ID) RadioButton {
	return RadioButton{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *RadioButton) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles c.WS_EX, styles c.WS, btnStyles c.BS) *RadioButton {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles, "BUTTON", text, // radio button is, in fact, a button
		styles|c.WS(btnStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Creates the first radio button of a group, with
// WS_GROUP and BS_AUTORADIOBUTTON styles. Position will be adjusted to the
// current system DPI. The size will be calculated to fit the text exactly.
func (me *RadioButton) CreateFirst(parent Window, x, y int32,
	text string) *RadioButton {

	return me.createBase(parent, x, y, text,
		c.WS_GROUP|c.WS(c.BS_AUTORADIOBUTTON))
}

// Calls CreateWindowEx(). Creates a subsequent radio button of a group, with
// BS_AUTORADIOBUTTON style. Position will be adjusted to the current system
// DPI. The size will be calculated to fit the text exactly.
func (me *RadioButton) CreateSubsequent(parent Window, x, y int32,
	text string) *RadioButton {

	return me.createBase(parent, x, y, text, c.WS(c.BS_AUTORADIOBUTTON))
}

func (me *RadioButton) IsChecked() bool {
	return c.BST(me.Hwnd().
		SendMessage(c.WM(c.BM_GETCHECK), 0, 0)) == c.BST_CHECKED
}

func (me *RadioButton) SetCheck() *RadioButton {
	me.Hwnd().SendMessage(c.WM(c.BM_SETCHECK), api.WPARAM(c.BST_CHECKED), 0)
	return me
}

// SetWindowText() doesn't resize the control to fit the text. This method
// resizes the control to fit the text exactly.
func (me *RadioButton) SetText(text string) *RadioButton {
	cx, cy := me.calcRadioButtonIdealSize(me.Hwnd().GetParent(), text)

	me.Hwnd().SetWindowPos(c.SWP_HWND(0), 0, 0, cx, cy,
		c.SWP_NOZORDER|c.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
	return me
}

// Returns the text without the accelerator ampersands.
// For example: "&He && she" is returned as "He & she".
// Use HWND().GetWindowText() to retrieve the full text, with ampersands.
func (me *RadioButton) Text() string {
	return removeAccelAmpersands(me.Hwnd().GetWindowText())
}

func (me *RadioButton) calcRadioButtonIdealSize(hReferenceDc api.HWND,
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

func (me *RadioButton) createBase(parent Window, x, y int32,
	text string, otherStyles c.WS) *RadioButton {

	x, y, _, _ = multiplyByDpi(x, y, 0, 0)
	cx, cy := me.calcRadioButtonIdealSize(parent.Hwnd(), text)

	me.controlNativeBase.create(c.WS_EX(0), "BUTTON", text,
		c.WS_CHILD|c.WS_VISIBLE|otherStyles,
		x, y, cx, cy, parent)
	globalUiFont.SetOnControl(me)
	return me
}
