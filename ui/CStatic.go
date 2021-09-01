package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native static control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/about-static-controls
type Static interface {
	AnyNativeControl
	isStatic() // prevent public implementation

	// Exposes all the Static notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-static-control-reference-notifications
	On() *_StaticEvents

	SetText(text string)          // Sets the text.
	SetTextAndResize(text string) // Sets the text and resizes the control to fit it exactly.
	Text() string                 // Retrieves the text.
}

//------------------------------------------------------------------------------

type _Static struct {
	_NativeControlBase
	events _StaticEvents
}

// Creates a new Static. Call StaticOpts() to define the options to be passed to
// the underlying CreateWindowEx().
func NewStatic(parent AnyParent, opts *_StaticO) Static {
	opts.lateDefaults()

	me := &_Static{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.position, nil)
		boundBox := _CalcTextBoundBox(opts.text, true)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			"STATIC", opts.text, opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, boundBox, win.HMENU(opts.ctrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
	})

	return me
}

// Creates a new Static from a dialog resource.
func NewStaticDlg(parent AnyParent, ctrlId int) Static {
	me := &_Static{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return me
}

func (me *_Static) isStatic() {}

func (me *_Static) On() *_StaticEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the Static is created.")
	}
	return &me.events
}

func (me *_Static) SetText(text string) {
	me.Hwnd().SetWindowText(text)
}

func (me *_Static) SetTextAndResize(text string) {
	me.Hwnd().SetWindowText(text)
	boundBox := _CalcTextBoundBox(text, true)
	me.Hwnd().SetWindowPos(win.HWND(0), 0, 0,
		boundBox.Cx, boundBox.Cy, co.SWP_NOZORDER|co.SWP_NOMOVE)
}

func (me *_Static) Text() string {
	return me.Hwnd().GetWindowText()
}

//------------------------------------------------------------------------------

type _StaticO struct {
	ctrlId int

	text        string
	position    win.POINT
	ctrlStyles  co.SS
	wndStyles   co.WS
	wndExStyles co.WS_EX
}

// Control ID.
// Defaults to an auto-generated ID.
func (o *_StaticO) CtrlId(i int) *_StaticO { o.ctrlId = i; return o }

// Text to appear in the control, passed to CreateWindowEx().
// Defaults to empty string.
func (o *_StaticO) Text(t string) *_StaticO { o.text = t; return o }

// Position within parent's client area in pixels.
// Defaults to 0x0. Will be adjusted to the current system DPI.
func (o *_StaticO) Position(p win.POINT) *_StaticO { _OwPt(&o.position, p); return o }

// Static control styles, passed to CreateWindowEx().
// Defaults to SS_LEFT | SS_NOTIFY.
func (o *_StaticO) CtrlStyles(s co.SS) *_StaticO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
func (o *_StaticO) WndStyles(s co.WS) *_StaticO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/stn-clicked
func (me *_StaticEvents) StnClicked(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.STN_CLICKED, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/stn-dblclk
func (me *_StaticEvents) StnDblClk(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.STN_DBLCLK, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/stn-disable
func (me *_StaticEvents) StnDisable(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.STN_DISABLE, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/stn-enable
func (me *_StaticEvents) StnEnable(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.STN_ENABLE, func(_ wm.Command) {
		userFunc()
	})
}
