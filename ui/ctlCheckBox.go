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
	*_NativeControlBase
	events *_EventsButton // the CheckBox is just a Button type
}

// Constructor. Optionally receives a control ID.
func NewCheckBox(parent Parent, ctrlId ...int) *CheckBox {
	base := _NewNativeControlBase(parent, ctrlId...)
	return &CheckBox{
		_NativeControlBase: base,
		events:             _NewEventsButton(base),
	}
}

// Calls CreateWindowEx(). With this method, you must also specify WS and WS_EX
// window styles.
//
// Position and size will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *CheckBox) CreateWs(
	text string, pos Pos, size Size,
	btnStyles co.BS, styles co.WS, exStyles co.WS_EX) *CheckBox {

	_global.MultiplyDpi(&pos, &size)
	return me.createNoDpi(text, pos, size, btnStyles, styles, exStyles)
}

// Calls CreateWindowEx() with WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
// Size will be calculated to fit the text exactly.
//
// A typical CheckBox has BS_AUTOCHECKBOX, a 3-state has BS_AUTO3STATE.
//
// Position will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *CheckBox) Create(text string, pos Pos, btnStyles co.BS) *CheckBox {
	_global.MultiplyDpi(&pos, nil)
	size := me.calcIdealSize(text)
	return me.createNoDpi(text, pos, size, btnStyles,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_NONE)
}

func (me *CheckBox) createAsDlgCtrl() { me._NativeControlBase.createAssignDlg() }

// Exposes all CheckBox notifications.
//
// Cannot be called after the parent window was created.
func (me *CheckBox) On() *_EventsButton {
	if me.hwnd != 0 {
		panic("Cannot add notifications after the CheckBox was created.")
	}
	return me.events
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

// Sets the text, and resizes the control to fit it exactly.
//
// To set the text without resizing the control, use Hwnd().SetWindowText().
func (me *CheckBox) SetText(text string) *CheckBox {
	size := me.calcIdealSize(text)
	me.Hwnd().SetWindowPos(co.SWP_HWND_NONE,
		0, 0, int32(size.Cx), int32(size.Cy),
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
// Use Hwnd().GetWindowText() to retrieve the raw text, with unparsed
// accelerator ampersands.
func (me *CheckBox) Text() string {
	return _global.RemoveAccelAmpersands(me.Hwnd().GetWindowText())
}

// Calculates the exact size to fit the text.
func (me *CheckBox) calcIdealSize(text string) Size {
	boundBox := _global.CalcTextBoundBox(text, true)
	boundBox.Cx += int(
		win.GetSystemMetrics(co.SM_CXMENUCHECK) + // https://stackoverflow.com/a/1165052/6923555
			win.GetSystemMetrics(co.SM_CXEDGE),
	)

	cyCheck := int(win.GetSystemMetrics(co.SM_CYMENUCHECK))
	if cyCheck > boundBox.Cy {
		boundBox.Cy = cyCheck // if the check is taller than the font, use its height
	}
	return boundBox
}

// Creates the control without adjusting DPI.
func (me *CheckBox) createNoDpi(
	text string, pos Pos, size Size,
	btnStyles co.BS, styles co.WS, exStyles co.WS_EX) *CheckBox {

	me._NativeControlBase.create("BUTTON", text, pos, size, // the CheckBox is just a Button type
		co.WS(btnStyles)|styles, exStyles)
	_global.UiFont().SetOnControl(me)
	return me
}
