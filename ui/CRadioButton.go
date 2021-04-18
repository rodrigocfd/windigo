package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native radio button control.
//
// Prefer using a RadioGroup instead.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#radio-buttons
type RadioButton interface {
	AnyControl

	// Exposes all the Button notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
	On() *_ButtonEvents

	isRadioButton() // disambiguate

	EmulateClick()       // Emulates an user click.
	IsChecked() bool     // Retrieves the current check state.
	SetChecked()         // Checks the radio button.
	SetText(text string) // Sets the text and resizes the control to fit it exactly.
}

//------------------------------------------------------------------------------

type _RadioButton struct {
	_NativeControlBase
	events _ButtonEvents
}

// Creates a new RadioButton specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewRadioButtonRaw(parent AnyParent, opts RadioButtonRawOpts) RadioButton {
	opts.fillBlankValuesWithDefault()

	me := _RadioButton{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.Position, nil)
		boundBox := _CalcTextBoundBoxWithCheck(opts.Text, true)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"BUTTON", opts.Text, opts.Styles|co.WS(opts.ButtonStyles),
			opts.Position, boundBox, win.HMENU(opts.CtrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)

		if opts.Checked {
			me.SetChecked()
		}
	})

	return &me
}

// Creates a new RadioButton from a dialog resource.
func NewRadioButtonDlg(parent AnyParent, ctrlId int) RadioButton {
	me := _RadioButton{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return &me
}

func (me *_RadioButton) isRadioButton() {}

func (me *_RadioButton) On() *_ButtonEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the RadioButton is created.")
	}
	return &me.events
}

func (me *_RadioButton) EmulateClick() {
	me.Hwnd().SendMessage(co.BM_CLICK, 0, 0)
}

func (me *_RadioButton) IsChecked() bool {
	return (co.BST(me.Hwnd().SendMessage(co.BM_GETSTATE, 0, 0)) &
		co.BST_CHECKED) != 0
}

func (me *_RadioButton) SetChecked() {
	me.Hwnd().SendMessage(co.BM_SETCHECK, win.WPARAM(co.BST_CHECKED), 0)
}

func (me *_RadioButton) SetText(text string) {
	me.Hwnd().SetWindowText(text)
	boundBox := _CalcTextBoundBoxWithCheck(text, true)
	me.Hwnd().SetWindowPos(win.HWND(0), 0, 0,
		boundBox.Cx, boundBox.Cy, co.SWP_NOZORDER|co.SWP_NOMOVE)
}

//------------------------------------------------------------------------------

// Options for NewRadioButtonRaw().
type RadioButtonRawOpts struct {
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
	// Defaults to BS_AUTORADIOBUTTON.
	ButtonStyles co.BS
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_VISIBLE.
	// Note that the first radio button of a group should have WS_TABSTOP | WS_GROUP too.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStyles co.WS_EX

	// RadioButton initial checked state.
	// Defaults to false.
	Checked bool
}

func (opts *RadioButtonRawOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.ButtonStyles == 0 {
		opts.ButtonStyles = co.BS_AUTORADIOBUTTON
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_NONE
	}
}
