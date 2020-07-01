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

// Native radio button control.
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type RadioButton struct {
	controlNativeBase
}

// Optional; returns a RadioButton with a specific control ID.
func MakeRadioButton(ctrlId co.ID) RadioButton {
	return RadioButton{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *RadioButton) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles co.WS_EX, styles co.WS, btnStyles co.BS) *RadioButton {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles, "BUTTON", text, // radio button is, in fact, a button
		styles|co.WS(btnStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Creates the first radio button of a group, with
// WS_GROUP and BS_AUTORADIOBUTTON styles. Position will be adjusted to the
// current system DPI. The size will be calculated to fit the text exactly.
func (me *RadioButton) CreateFirst(parent Window, x, y int32,
	text string) *RadioButton {

	return me.createBase(parent, x, y, text,
		co.WS_GROUP|co.WS_TABSTOP|co.WS(co.BS_AUTORADIOBUTTON))
}

// Calls CreateWindowEx(). Creates a subsequent radio button of a group, with
// BS_AUTORADIOBUTTON style. Position will be adjusted to the current system
// DPI. The size will be calculated to fit the text exactly.
func (me *RadioButton) CreateSubsequent(parent Window, x, y int32,
	text string) *RadioButton {

	return me.createBase(parent, x, y, text, co.WS(co.BS_AUTORADIOBUTTON))
}

func (me *RadioButton) IsChecked() bool {
	return co.BST(me.Hwnd().
		SendMessage(co.WM(co.BM_GETCHECK), 0, 0)) == co.BST_CHECKED
}

func (me *RadioButton) SetCheck() *RadioButton {
	me.Hwnd().
		SendMessage(co.WM(co.BM_SETCHECK), win.WPARAM(co.BST_CHECKED), 0)
	return me
}

// SetWindowText() doesn't resize the control to fit the text. This method
// resizes the control to fit the text exactly.
func (me *RadioButton) SetText(text string) *RadioButton {
	cx, cy := me.calcRadioButtonIdealSize(me.Hwnd().GetParent(), text)

	me.Hwnd().SetWindowPos(co.SWP_HWND(0), 0, 0, cx, cy,
		co.SWP_NOZORDER|co.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
	return me
}

// Returns the text without the accelerator ampersands.
// For example: "&He && she" is returned as "He & she".
// Use HWND().GetWindowText() to retrieve the full text, with ampersands.
func (me *RadioButton) Text() string {
	return removeAccelAmpersands(me.Hwnd().GetWindowText())
}

func (me *RadioButton) calcRadioButtonIdealSize(hReferenceDc win.HWND,
	text string) (uint32, uint32) {

	cx, cy := calcIdealSize(hReferenceDc, text, true)
	cx += uint32(win.GetSystemMetrics(co.SM_CXMENUCHECK)) +
		uint32(win.GetSystemMetrics(co.SM_CXEDGE)) // https://stackoverflow.com/a/1165052/6923555

	cyCheck := uint32(win.GetSystemMetrics(co.SM_CYMENUCHECK))
	if cyCheck > cy {
		cy = cyCheck // if the check is taller than the font, use its height
	}

	return cx, cy
}

func (me *RadioButton) createBase(parent Window, x, y int32,
	text string, otherStyles co.WS) *RadioButton {

	x, y, _, _ = multiplyByDpi(x, y, 0, 0)
	cx, cy := me.calcRadioButtonIdealSize(parent.Hwnd(), text)

	me.controlNativeBase.create(co.WS_EX(0), "BUTTON", text,
		co.WS_CHILD|co.WS_VISIBLE|otherStyles,
		x, y, cx, cy, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Helper function to retrieve the index of the checked radio button.
// Returns -1 if none is checked.
func GetCheckedRadio(radios []RadioButton) int32 {
	for i := range radios {
		if radios[i].IsChecked() {
			return int32(i)
		}
	}
	return -1 // no checked one
}

// Simple utility conversion; useful with Resizer.
func RadioButtonAsControl(radios []RadioButton) []Control {
	ctrls := make([]Control, 0, len(radios))
	for i := range radios {
		ctrls = append(ctrls, &radios[i])
	}
	return ctrls
}
