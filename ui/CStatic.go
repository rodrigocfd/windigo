//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native static control.
//
// 📑 https://learn.microsoft.com/en-us/windows/win32/controls/about-static-controls
type Static interface {
	AnyNativeControl
	AnyTextControl
	implStatic() // prevent public implementation

	// Exposes all the [Static notifications] the can be handled.
	//
	// Panics if called after the control was created.
	//
	// [Static notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-static-control-reference-notifications
	On() *_StaticEvents

	SetTextAndResize(text string) // Sets the text and resizes the control to fit it exactly.
}

//------------------------------------------------------------------------------

type _Static struct {
	_NativeControlBase
	events _StaticEvents
}

// Creates a new Static. Call ui.StaticOpts() to define the options to be passed
// to the underlying CreateWindowEx().
//
// Example:
//
//	var owner ui.AnyParent // initialized somewhere
//
//	myLabel := ui.NewStatic(
//		owner,
//		ui.StaticOpts().
//			Text("Some label").
//			Position(win.POINT{X: 20, Y: 10}),
//		),
//	)
func NewStatic(parent AnyParent, opts *_StaticO) Static {
	if opts == nil {
		opts = StaticOpts()
	}
	opts.lateDefaults()

	me := &_Static{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_ConvertDtuOrMultiplyDpi(parent, &opts.position, &opts.size)
		if opts.size.Cx == 0 && opts.size.Cy == 0 {
			opts.size = _CalcTextBoundBox(opts.text, true)
		}

		me._NativeControlBase.createWindow(opts.wndExStyles,
			win.ClassNameStr("STATIC"), win.StrOptSome(opts.text),
			opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		parent.addResizingChild(me, opts.horz, opts.vert)
		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
	})

	return me
}

// Creates a new Static from a dialog resource.
func NewStaticDlg(parent AnyParent, ctrlId int, horz HORZ, vert VERT) Static {
	me := &_Static{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
		parent.addResizingChild(me, horz, vert)
	})

	return me
}

// Implements Static.
func (*_Static) implStatic() {}

// Implements AnyTextControl.
func (me *_Static) SetText(text string) {
	me.Hwnd().SetWindowText(text)
}

// Implements AnyTextControl.
func (me *_Static) Text() string {
	return me.Hwnd().GetWindowText()
}

func (me *_Static) On() *_StaticEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the Static is created.")
	}
	return &me.events
}

func (me *_Static) SetTextAndResize(text string) {
	me.SetText(text)
	boundBox := _CalcTextBoundBox(text, true)
	me.Hwnd().SetWindowPos(win.HWND(0), 0, 0,
		boundBox.Cx, boundBox.Cy, co.SWP_NOZORDER|co.SWP_NOMOVE)
}

//------------------------------------------------------------------------------

type _StaticO struct {
	ctrlId int

	text        string
	position    win.POINT
	size        win.SIZE
	horz        HORZ
	vert        VERT
	ctrlStyles  co.SS
	wndStyles   co.WS
	wndExStyles co.WS_EX
}

// Control ID.
//
// Defaults to an auto-generated ID.
func (o *_StaticO) CtrlId(i int) *_StaticO { o.ctrlId = i; return o }

// Text to appear in the control, passed to CreateWindowEx().
//
// Defaults to empty string.
func (o *_StaticO) Text(t string) *_StaticO { o.text = t; return o }

// Position within parent's client area.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 0x0.
func (o *_StaticO) Position(p win.POINT) *_StaticO { _OwPt(&o.position, p); return o }

// Control size.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to fit current text.
func (o *_StaticO) Size(s win.SIZE) *_StaticO { _OwSz(&o.size, s); return o }

// Horizontal behavior when the parent is resized.
//
// Defaults to HORZ_NONE.
func (o *_StaticO) Horz(s HORZ) *_StaticO { o.horz = s; return o }

// Vertical behavior when the parent is resized.
//
// Defaults to VERT_NONE.
func (o *_StaticO) Vert(s VERT) *_StaticO { o.vert = s; return o }

// Static control styles, passed to CreateWindowEx().
//
// Defaults to SS_LEFT | SS_NOTIFY.
func (o *_StaticO) CtrlStyles(s co.SS) *_StaticO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
func (o *_StaticO) WndStyles(s co.WS) *_StaticO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
//
// Defaults to WS_EX_NONE.
func (o *_StaticO) WndExStyles(s co.WS_EX) *_StaticO { o.wndExStyles = s; return o }

func (o *_StaticO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewStatic().
func StaticOpts() *_StaticO {
	return &_StaticO{
		horz:       HORZ_NONE,
		vert:       VERT_NONE,
		ctrlStyles: co.SS_LEFT | co.SS_NOTIFY,
		wndStyles:  co.WS_CHILD | co.WS_VISIBLE,
	}
}

//------------------------------------------------------------------------------

// Static control notifications.
type _StaticEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_StaticEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.Parent().On()
}

// [STN_CLICKED] message handler.
//
// [STN_CLICKED]: https://learn.microsoft.com/en-us/windows/win32/controls/stn-clicked
func (me *_StaticEvents) StnClicked(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.STN_CLICKED, func(_ wm.Command) {
		userFunc()
	})
}

// [STN_DBLCLK] message handler.
//
// [STN_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/stn-dblclk
func (me *_StaticEvents) StnDblClk(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.STN_DBLCLK, func(_ wm.Command) {
		userFunc()
	})
}

// [STN_DISABLE] message handler.
//
// [STN_DISABLE]: https://learn.microsoft.com/en-us/windows/win32/controls/stn-disable
func (me *_StaticEvents) StnDisable(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.STN_DISABLE, func(_ wm.Command) {
		userFunc()
	})
}

// [STN_ENABLE] message handler.
//
// [STN_ENABLE]: https://learn.microsoft.com/en-us/windows/win32/controls/stn-enable
func (me *_StaticEvents) StnEnable(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.STN_ENABLE, func(_ wm.Command) {
		userFunc()
	})
}
