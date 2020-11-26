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

// Native RadioButton control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#radio-buttons
type RadioButton struct {
	*_NativeControlBase
	events *_EventsButton // the RadioButton is just a Button type
}

// Constructor. Optionally receives a control ID.
func NewRadioButton(parent Parent, ctrlId ...int) *RadioButton {
	base := _NewNativeControlBase(parent, ctrlId...)
	return &RadioButton{
		_NativeControlBase: _NewNativeControlBase(parent, ctrlId...),
		events:             _NewEventsButton(base),
	}
}

// Calls CreateWindowEx(). With this method, you must also specify WS and WS_EX
// window styles.
//
// Position and size will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *RadioButton) CreateWs(
	text string, pos Pos, size Size,
	btnStyles co.BS, styles co.WS, exStyles co.WS_EX) *RadioButton {

	_global.MultiplyDpi(&pos, &size)
	return me.createNoDpi(text, pos, size, btnStyles, styles, exStyles)
}

// Calls CreateWindowEx() with WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
// Size will be calculated to fit the text exactly.
//
// A typical RadioButton has BS_AUTORADIOBUTTON.
//
// Position will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *RadioButton) CreateFirst(
	text string, pos Pos, btnStyles co.BS) *RadioButton {

	_global.MultiplyDpi(&pos, nil)
	size := me.calcIdealSize(me.parent.Hwnd(), text)
	return me.createNoDpi(text, pos, size, btnStyles,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_NONE)
}

// Calls CreateWindowEx() with WS_CHILD | WS_VISIBLE.
// Size will be calculated to fit the text exactly.
//
// A typical RadioButton has BS_AUTORADIOBUTTON.
//
// Position will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *RadioButton) CreateSubsequent(
	text string, pos Pos, btnStyles co.BS) *RadioButton {

	_global.MultiplyDpi(&pos, nil)
	size := me.calcIdealSize(me.parent.Hwnd(), text)
	return me.createNoDpi(text, pos, size, btnStyles,
		co.WS_CHILD|co.WS_VISIBLE,
		co.WS_EX_NONE)
}

func (me *RadioButton) createAsDlgCtrl() { me._NativeControlBase.createAssignDlg() }

// Exposes all RadioButton notifications.
//
// Cannot be called after the parent window was created.
func (me *RadioButton) On() *_EventsButton {
	if me.hwnd != 0 {
		panic("Cannot add notifications after the RadioButton was created.")
	}
	return me.events
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
	return _global.RemoveAccelAmpersands(me.Hwnd().GetWindowText())
}

// Calculates the exact size to fit the text.
func (me *RadioButton) calcIdealSize(hReferenceDc win.HWND, text string) Size {
	boundBox := _global.CalcTextBoundBox(hReferenceDc, text, true)
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
func (me *RadioButton) createNoDpi(
	text string, pos Pos, size Size,
	btnStyles co.BS, styles co.WS, exStyles co.WS_EX) *RadioButton {

	me._NativeControlBase.create("BUTTON", text, pos, size, // the RadioButton is just a Button type
		co.WS(btnStyles)|styles, exStyles)
	_global.UiFont().SetOnControl(me)
	return me
}

//------------------------------------------------------------------------------

// Manages a group of RadioButton controls.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#radio-buttons
type RadioGroup struct {
	radioButtons []*RadioButton
}

// Constructor. Receives the number of RadioButton controls to be created.
// Calls the constructor of each RadioButton.
func NewRadioGroupCount(parent Parent, numRadios int) *RadioGroup {
	me := RadioGroup{
		radioButtons: make([]*RadioButton, numRadios),
	}
	for i := range me.radioButtons {
		me.radioButtons[i] = NewRadioButton(parent)
	}
	return &me
}

// Constructor. Receives the control ID of each RadioButton.
// Calls the constructor of each RadioButton.
func NewRadioGroupIds(parent Parent, ctrlIds ...int) *RadioGroup {
	me := RadioGroup{
		radioButtons: make([]*RadioButton, len(ctrlIds)),
	}
	for i := range me.radioButtons {
		me.radioButtons[i] = NewRadioButton(parent, ctrlIds[i])
	}
	return &me
}

// Returns a slice with Control interfaces.
func (me *RadioGroup) AsControl() []Control {
	ctrls := make([]Control, 0, len(me.radioButtons))
	for i := range me.radioButtons {
		ctrls = append(ctrls, me.radioButtons[i])
	}
	return ctrls
}

// Returns the checked RadioButton, if any.
func (me *RadioGroup) Checked() (*RadioButton, bool) {
	for i := range me.radioButtons {
		if me.radioButtons[i].IsChecked() {
			return me.radioButtons[i], true
		}
	}
	return nil, false
}

// Returns the RadioButton at the given index.
//
// Does not perform bound checking.
func (me *RadioGroup) Get(index int) *RadioButton {
	return me.radioButtons[index]
}
