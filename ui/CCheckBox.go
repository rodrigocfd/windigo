package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native check box control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#check-boxes
type CheckBox interface {
	AnyControl

	// Exposes all the Button notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
	On() *_ButtonEvents

	isCheckBox() // disambiguate

	CheckState() co.BST         // Retrieves the current check state.
	EmulateClick()              // Emulates an user click.
	SetCheckState(state co.BST) // Sets the current check state.
	SetText(text string)        // Sets the text and resizes the control to fit it exactly.
}

//------------------------------------------------------------------------------

type _CheckBox struct {
	_NativeControlBase
	events _ButtonEvents
}

// Creates a new CheckBox specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewCheckBoxOpts(parent AnyParent, opts CheckBoxOpts) CheckBox {
	opts.fillBlankValuesWithDefault()

	me := _CheckBox{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_ParentCreateWm(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.Position, nil)
		boundBox := _CalcTextBoundBoxWithCheck(opts.Text, true)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"BUTTON", opts.Text, opts.Styles|co.WS(opts.ButtonStyles),
			opts.Position, boundBox, win.HMENU(opts.CtrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
		me.SetCheckState(opts.State)
	})

	return &me
}

// Creates a new CheckBox from a dialog resource.
func NewCheckBoxDlg(parent AnyParent, ctrlId int) CheckBox {
	me := _CheckBox{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return &me
}

func (me *_CheckBox) isCheckBox() {}

func (me *_CheckBox) On() *_ButtonEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the CheckBox is created.")
	}
	return &me.events
}

func (me *_CheckBox) CheckState() co.BST {
	return co.BST(me.Hwnd().SendMessage(co.BM_GETCHECK, 0, 0))
}

func (me *_CheckBox) EmulateClick() {
	me.Hwnd().SendMessage(co.BM_CLICK, 0, 0)
}

func (me *_CheckBox) SetCheckState(state co.BST) {
	me.Hwnd().SendMessage(co.BM_SETCHECK, win.WPARAM(state), 0)
}

func (me *_CheckBox) SetText(text string) {
	me.Hwnd().SetWindowText(text)
	boundBox := _CalcTextBoundBoxWithCheck(text, true)
	me.Hwnd().SetWindowPos(win.HWND(0), 0, 0,
		boundBox.Cx, boundBox.Cy, co.SWP_NOZORDER|co.SWP_NOMOVE)
}

//------------------------------------------------------------------------------

// Options for NewCheckBoxOpts().
type CheckBoxOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

	// Text to appear in the control, passed to CreateWindowEx().
	// Defaults to empty string.
	Text string
	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// Button control styles, passed to CreateWindowEx().
	// Defaults to BS_AUTOCHECKBOX.
	ButtonStyles co.BS
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStyles co.WS_EX

	// CheckBox initial state.
	// Defaults to BST_UNCHECKED.
	State co.BST
}

func (opts *CheckBoxOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.ButtonStyles == 0 {
		opts.ButtonStyles = co.BS_AUTOCHECKBOX
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_NONE
	}

	if opts.State == 0 {
		opts.State = co.BST_UNCHECKED
	}
}
