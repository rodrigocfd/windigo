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
	AnyEnabledControl
	AnyFocusControl
	AnyTextControl
	isCheckBox() // prevent public implementation

	// Exposes all the Button notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
	On() *_ButtonEvents

	CheckState() co.BST                   // Retrieves the current check state.
	EmulateClick()                        // Emulates an user click.
	IsChecked() bool                      // Tells whether the check box state is checked.
	SetCheckState(state co.BST)           // Sets the current check state.
	SetCheckStateAndTrigger(state co.BST) // Sets the current check state and triggers the click event.
	SetTextAndResize(text string)         // Sets the text and resizes the control to fit it exactly.
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
		_ConvertDtuOrMultiplyDpi(parent, &opts.position, &opts.size)
		if opts.size.Cx == 0 && opts.size.Cy == 0 {
			opts.size = _CalcTextBoundBoxWithCheck(opts.text, true)
		}

		me._NativeControlBase.createWindow(opts.wndExStyles,
			win.ClassNameStr("BUTTON"), win.StrVal(opts.text),
			opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		parent.addResizingChild(me, opts.horz, opts.vert)
		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
		me.SetCheckState(opts.state)
	})

	return me
}

// Creates a new CheckBox from a dialog resource.
func NewCheckBoxDlg(
	parent AnyParent, ctrlId int,
	horz HORZ, vert VERT) CheckBox {

	me := &_CheckBox{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
		parent.addResizingChild(me, horz, vert)
	})

	return me
}

// Implements CheckBox.
func (me *_CheckBox) isCheckBox() {}

// Implements AnyEnabledControl.
func (me *_CheckBox) Enable(enable bool) {
	me.Hwnd().EnableWindow(enable)
}

// Implements AnyFocusControl.
func (me *_CheckBox) Focus() {
	me._NativeControlBase.focus()
}

// Implements AnyTextControl.
func (me *_CheckBox) SetText(text string) {
	me.Hwnd().SetWindowText(text)
}

// Implements AnyTextControl.
func (me *_CheckBox) Text() string {
	return me.Hwnd().GetWindowText()
}

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

func (me *_CheckBox) IsChecked() bool {
	return me.CheckState() == co.BST_CHECKED
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

func (me *_CheckBox) SetTextAndResize(text string) {
	me.SetText(text)
	boundBox := _CalcTextBoundBoxWithCheck(text, true)
	me.Hwnd().SetWindowPos(win.HWND(0), 0, 0,
		boundBox.Cx, boundBox.Cy, co.SWP_NOZORDER|co.SWP_NOMOVE)
}

//------------------------------------------------------------------------------

type _CheckBoxO struct {
	ctrlId int

	text        string
	position    win.POINT
	size        win.SIZE
	horz        HORZ
	vert        VERT
	ctrlStyles  co.BS
	wndStyles   co.WS
	wndExStyles co.WS_EX

	state co.BST
}

// Control ID.
//
// Defaults to an auto-generated ID.
func (o *_CheckBoxO) CtrlId(i int) *_CheckBoxO { o.ctrlId = i; return o }

// Text to appear in the control, passed to CreateWindowEx().
//
// Defaults to empty string.
func (o *_CheckBoxO) Text(t string) *_CheckBoxO { o.text = t; return o }

// Position within parent's client area.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 0x0.
func (o *_CheckBoxO) Position(p win.POINT) *_CheckBoxO { _OwPt(&o.position, p); return o }

// Control size.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to fit current text.
func (o *_CheckBoxO) Size(s win.SIZE) *_CheckBoxO { _OwSz(&o.size, s); return o }

// Horizontal behavior when the parent is resized.
//
// Defaults to HORZ_NONE.
func (o *_CheckBoxO) Horz(s HORZ) *_CheckBoxO { o.horz = s; return o }

// Vertical behavior when the parent is resized.
//
// Defaults to VERT_NONE.
func (o *_CheckBoxO) Vert(s VERT) *_CheckBoxO { o.vert = s; return o }

// CheckBox control styles, passed to CreateWindowEx().
//
// Defaults to BS_AUTOCHECKBOX.
func (o *_CheckBoxO) CtrlStyles(s co.BS) *_CheckBoxO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *_CheckBoxO) WndStyles(s co.WS) *_CheckBoxO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
//
// Defaults to WS_EX_NONE.
func (o *_CheckBoxO) WndExStyles(s co.WS_EX) *_CheckBoxO { o.wndExStyles = s; return o }

// CheckBox initial state.
//
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
		horz:       HORZ_NONE,
		vert:       VERT_NONE,
		ctrlStyles: co.BS_AUTOCHECKBOX,
		wndStyles:  co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_VISIBLE,
		state:      co.BST_UNCHECKED,
	}
}
