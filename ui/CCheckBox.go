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
	AnyNativeControl
	isCheckBox() // prevent public implementation

	// Exposes all the Button notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
	On() *_ButtonEvents

	CheckState() co.BST                   // Retrieves the current check state.
	EmulateClick()                        // Emulates an user click.
	SetCheckState(state co.BST)           // Sets the current check state.
	SetCheckStateAndTrigger(state co.BST) // Sets the current check state and triggers the click event.
	SetText(text string)                  // Sets the text.
	SetTextAndResize(text string)         // Sets the text and resizes the control to fit it exactly.
	Text() string                         // Retrieves the text.
}

//------------------------------------------------------------------------------

type _CheckBox struct {
	_NativeControlBase
	events _ButtonEvents
}

// Creates a new CheckBox. Call CheckBoxOpts() to define the options to be
// passed to the underlying CreateWindowEx().
func NewCheckBox(parent AnyParent, opts *_CheckBoxO) CheckBox {
	opts.lateDefaults()

	me := &_CheckBox{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.position, &opts.size)
		if opts.size.Cx == 0 && opts.size.Cy == 0 {
			opts.size = _CalcTextBoundBoxWithCheck(opts.text, true)
		}

		me._NativeControlBase.createWindow(opts.wndExStyles,
			"BUTTON", opts.text, opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
		me.SetCheckState(opts.state)
	})

	return me
}

// Creates a new CheckBox from a dialog resource.
func NewCheckBoxDlg(parent AnyParent, ctrlId int) CheckBox {
	me := &_CheckBox{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return me
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

func (me *_CheckBox) SetCheckStateAndTrigger(state co.BST) {
	me.SetCheckState(state)
	me.Hwnd().GetParent().SendMessage(co.WM_COMMAND,
		win.MAKEWPARAM(uint16(me.CtrlId()), uint16(co.BN_CLICKED)),
		win.LPARAM(me.Hwnd()))
}

func (me *_CheckBox) SetText(text string) {
	me.Hwnd().SetWindowText(text)
}

func (me *_CheckBox) SetTextAndResize(text string) {
	me.SetText(text)
	boundBox := _CalcTextBoundBoxWithCheck(text, true)
	me.Hwnd().SetWindowPos(win.HWND(0), 0, 0,
		boundBox.Cx, boundBox.Cy, co.SWP_NOZORDER|co.SWP_NOMOVE)
}

func (me *_CheckBox) Text() string {
	return me.Hwnd().GetWindowText()
}

//------------------------------------------------------------------------------

type _CheckBoxO struct {
	ctrlId int

	text        string
	position    win.POINT
	size        win.SIZE
	ctrlStyles  co.BS
	wndStyles   co.WS
	wndExStyles co.WS_EX

	state co.BST
}

// Control ID.
// Defaults to an auto-generated ID.
func (o *_CheckBoxO) CtrlId(i int) *_CheckBoxO { o.ctrlId = i; return o }

// Text to appear in the control, passed to CreateWindowEx().
// Defaults to empty string.
func (o *_CheckBoxO) Text(t string) *_CheckBoxO { o.text = t; return o }

// Position within parent's client area in pixels.
// Defaults to 0x0. Will be adjusted to the current system DPI.
func (o *_CheckBoxO) Position(p win.POINT) *_CheckBoxO { _OwPt(&o.position, p); return o }

// Control size in pixels.
// Defaults to fit current text. Will be adjusted to the current system DPI.
func (o *_CheckBoxO) Size(s win.SIZE) *_CheckBoxO { _OwSz(&o.size, s); return o }

// CheckBox control styles, passed to CreateWindowEx().
// Defaults to BS_AUTOCHECKBOX.
func (o *_CheckBoxO) CtrlStyles(s co.BS) *_CheckBoxO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *_CheckBoxO) WndStyles(s co.WS) *_CheckBoxO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
// Defaults to WS_EX_NONE.
func (o *_CheckBoxO) WndExStyles(s co.WS_EX) *_CheckBoxO { o.wndExStyles = s; return o }

// CheckBox initial state.
// Defaults to BST_UNCHECKED.
func (o *_CheckBoxO) State(s co.BST) *_CheckBoxO { o.state = s; return o }

func (o *_CheckBoxO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewCheckBox().
func CheckBoxOpts() *_CheckBoxO {
	return &_CheckBoxO{
		ctrlStyles: co.BS_AUTOCHECKBOX,
		wndStyles:  co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_VISIBLE,
		state:      co.BST_UNCHECKED,
	}
}
