//go:build windows

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
// 📑 https://learn.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#radio-buttons
type RadioButton interface {
	AnyNativeControl
	AnyFocusControl
	AnyTextControl
	implRadioButton() // prevent public implementation

	// Exposes all the [Button notifications] the can be handled.
	//
	// Panics if called after the control was created.
	//
	// [Button notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
	On() *_ButtonEvents

	EmulateClick()                // Emulates an user click.
	IsSelected() bool             // Tells whether this radio button is the selected one in its group.
	Select()                      // Selects the radio button.
	SelectAndTrigger()            // Selects the radio button and triggers the click event.
	SetTextAndResize(text string) // Sets the text and resizes the control to fit it exactly.
}

//------------------------------------------------------------------------------

type _RadioButton struct {
	_NativeControlBase
	events _ButtonEvents
}

// Creates a new RadioButton. Call ui.RadioButtonOpts() to define the options to
// be passed to the underlying CreateWindowEx().
//
// Prefer using the RadioGroup, which manages multiple radio buttons at once.
//
// Example:
//
//	var owner ui.AnyParent // initialized somewhere
//
//	myRadio := ui.NewRadioButton(
//		owner,
//		ui.RadioButtonOpts().
//			Text("Some option").
//			Position(win.POINT{X: 10, Y: 40}).
//			WndStyles(co.WS_VISIBLE|co.WS_CHILD|co.WS_TABSTOP|co.WS_GROUP),
//		),
//	)
func NewRadioButton(parent AnyParent, opts *_RadioButtonO) RadioButton {
	if opts == nil {
		opts = RadioButtonOpts()
	}
	opts.lateDefaults()

	me := &_RadioButton{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_ConvertDtuOrMultiplyDpi(parent, &opts.position, &opts.size)
		if opts.size.Cx == 0 && opts.size.Cy == 0 {
			opts.size = _CalcTextBoundBoxWithCheck(opts.text, true)
		}

		me._NativeControlBase.createWindow(opts.wndExStyles,
			win.ClassNameStr("BUTTON"), win.StrOptSome(opts.text),
			opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		parent.addResizingChild(me, opts.horz, opts.vert)
		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)

		if opts.selected {
			me.Select()
		}
	})

	return me
}

// Creates a new RadioButton from a dialog resource.
func NewRadioButtonDlg(
	parent AnyParent, ctrlId int,
	horz HORZ, vert VERT) RadioButton {

	me := &_RadioButton{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
		parent.addResizingChild(me, horz, vert)
	})

	return me
}

// Implements RadioButton.
func (*_RadioButton) implRadioButton() {}

// Implements AnyFocusControl.
func (me *_RadioButton) Focus() {
	me._NativeControlBase.focus()
}

// Implements AnyTextControl.
func (me *_RadioButton) SetText(text string) {
	me.Hwnd().SetWindowText(text)
}

// Implements AnyTextControl.
func (me *_RadioButton) Text() string {
	return me.Hwnd().GetWindowText()
}

func (me *_RadioButton) On() *_ButtonEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the RadioButton is created.")
	}
	return &me.events
}

func (me *_RadioButton) EmulateClick() {
	me.Hwnd().SendMessage(co.BM_CLICK, 0, 0)
}

func (me *_RadioButton) IsSelected() bool {
	bst := co.BST(me.Hwnd().SendMessage(co.BM_GETSTATE, 0, 0))
	return (bst & co.BST_CHECKED) != 0
}

func (me *_RadioButton) Select() {
	me.Hwnd().SendMessage(co.BM_SETCHECK, win.WPARAM(co.BST_CHECKED), 0)
}

func (me *_RadioButton) SelectAndTrigger() {
	me.Select()
	me.Hwnd().GetParent().SendMessage(co.WM_COMMAND,
		win.MAKEWPARAM(uint16(me.CtrlId()), uint16(co.BN_CLICKED)),
		win.LPARAM(me.Hwnd()))
}

func (me *_RadioButton) SetTextAndResize(text string) {
	me.SetText(text)
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
	horz        HORZ
	vert        VERT
	ctrlStyles  co.BS
	wndStyles   co.WS
	wndExStyles co.WS_EX

	selected bool
}

// Control ID.
//
// Defaults to an auto-generated ID.
func (o *_RadioButtonO) CtrlId(i int) *_RadioButtonO { o.ctrlId = i; return o }

// Text to appear in the control, passed to CreateWindowEx().
//
// Defaults to empty string.
func (o *_RadioButtonO) Text(t string) *_RadioButtonO { o.text = t; return o }

// Position within parent's client area.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 0x0.
func (o *_RadioButtonO) Position(p win.POINT) *_RadioButtonO { _OwPt(&o.position, p); return o }

// Control size.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to fit current text.
func (o *_RadioButtonO) Size(s win.SIZE) *_RadioButtonO { _OwSz(&o.size, s); return o }

// Horizontal behavior when the parent is resized.
//
// Defaults to HORZ_NONE.
func (o *_RadioButtonO) Horz(s HORZ) *_RadioButtonO { o.horz = s; return o }

// Vertical behavior when the parent is resized.
//
// Defaults to VERT_NONE.
func (o *_RadioButtonO) Vert(s VERT) *_RadioButtonO { o.vert = s; return o }

// Button control styles, passed to CreateWindowEx().
//
// Defaults to BS_AUTORADIOBUTTON.
func (o *_RadioButtonO) CtrlStyles(s co.BS) *_RadioButtonO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
//
// Note that the first radio button of a group should also have WS_TABSTOP | WS_GROUP.
func (o *_RadioButtonO) WndStyles(s co.WS) *_RadioButtonO { o.wndStyles = s; return o }

// Defines this radio button as the one initially selected.
//
// Defaults to false.
func (o *_RadioButtonO) Select(c bool) *_RadioButtonO { o.selected = c; return o }

func (o *_RadioButtonO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewRadioButton().
func RadioButtonOpts() *_RadioButtonO {
	return &_RadioButtonO{
		horz:       HORZ_NONE,
		vert:       VERT_NONE,
		ctrlStyles: co.BS_AUTORADIOBUTTON,
		wndStyles:  co.WS_CHILD | co.WS_VISIBLE,
	}
}
