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
	AnyNativeControl
	isRadioButton() // prevent public implementation

	// Exposes all the Button notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
	On() *_ButtonEvents

	EmulateClick()                // Emulates an user click.
	IsChecked() bool              // Retrieves the current check state.
	SetChecked()                  // Checks the radio button.
	SetTextAndResize(text string) // Sets the text and resizes the control to fit it exactly.
}

//------------------------------------------------------------------------------

type _RadioButton struct {
	_NativeControlBase
	events _ButtonEvents
}

// Creates a new RadioButton. Call RadioButtonOpts() to define the options to be
// passed to the underlying CreateWindowEx().
func NewRadioButton(parent AnyParent, opts *_RadioButtonO) RadioButton {
	opts.lateDefaults()

	me := &_RadioButton{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.position, &opts.size)
		if opts.size.Cx == 0 && opts.size.Cy == 0 {
			opts.size = _CalcTextBoundBox(opts.text, true)
		}

		me._NativeControlBase.createWindow(opts.wndExStyles,
			"BUTTON", opts.text, opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)

		if opts.checked {
			me.SetChecked()
		}
	})

	return me
}

// Creates a new RadioButton from a dialog resource.
func NewRadioButtonDlg(parent AnyParent, ctrlId int) RadioButton {
	me := &_RadioButton{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return me
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
	bst := co.BST(me.Hwnd().SendMessage(co.BM_GETSTATE, 0, 0))
	return (bst & co.BST_CHECKED) != 0
}

func (me *_RadioButton) SetChecked() {
	me.Hwnd().SendMessage(co.BM_SETCHECK, win.WPARAM(co.BST_CHECKED), 0)
}

func (me *_RadioButton) SetTextAndResize(text string) {
	me.Hwnd().SetWindowText(text)
	boundBox := _CalcTextBoundBoxWithCheck(text, true)
	me.Hwnd().SetWindowPos(win.HWND(0), 0, 0,
		boundBox.Cx, boundBox.Cy, co.SWP_NOZORDER|co.SWP_NOMOVE)
}

//------------------------------------------------------------------------------

type _RadioButtonO struct {
	ctrlId int

	text        string
	position    win.POINT
	size        win.SIZE
	ctrlStyles  co.BS
	wndStyles   co.WS
	wndExStyles co.WS_EX

	checked bool
}

// Control ID.
// Defaults to an auto-generated ID.
func (o *_RadioButtonO) CtrlId(i int) *_RadioButtonO { o.ctrlId = i; return o }

// Text to appear in the control, passed to CreateWindowEx().
// Defaults to empty string.
func (o *_RadioButtonO) Text(t string) *_RadioButtonO { o.text = t; return o }

// Position within parent's client area in pixels.
// Defaults to 0x0. Will be adjusted to the current system DPI.
func (o *_RadioButtonO) Position(p win.POINT) *_RadioButtonO { _OwPt(&o.position, p); return o }

// Control size in pixels.
// Defaults to fit current text. Will be adjusted to the current system DPI.
func (o *_RadioButtonO) Size(s win.SIZE) *_RadioButtonO { _OwSz(&o.size, s); return o }

// Button control styles, passed to CreateWindowEx().
// Defaults to BS_AUTORADIOBUTTON.
func (o *_RadioButtonO) CtrlStyles(s co.BS) *_RadioButtonO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
// Note that the first radio button of a group should also have WS_TABSTOP | WS_GROUP.
func (o *_RadioButtonO) WndStyles(s co.WS) *_RadioButtonO { o.wndStyles = s; return o }

// RadioButton initial checked state.
// Defaults to false.
func (o *_RadioButtonO) Checked(c bool) *_RadioButtonO { o.checked = c; return o }

func (o *_RadioButtonO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewRadioButton().
func RadioButtonOpts() *_RadioButtonO {
	return &_RadioButtonO{
		ctrlStyles: co.BS_AUTORADIOBUTTON,
		wndStyles:  co.WS_CHILD | co.WS_VISIBLE,
	}
}
