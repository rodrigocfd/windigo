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

	// Exposes all the Static notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-static-control-reference-notifications
	On() *_StaticEvents

	SetText(text string) // Sets the text and resizes the control to fit it exactly.
}

//------------------------------------------------------------------------------

type _Static struct {
	_NativeControlBase
	events _StaticEvents
}

// Creates a new Static specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewStaticRaw(parent AnyParent, opts StaticRawOpts) Static {
	opts.fillBlankValuesWithDefault()

	me := &_Static{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.Position, nil)
		boundBox := _CalcTextBoundBox(opts.Text, true)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"STATIC", opts.Text, opts.Styles|co.WS(opts.StaticStyles),
			opts.Position, boundBox, win.HMENU(opts.CtrlId))

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

func (me *_Static) On() *_StaticEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the Static is created.")
	}
	return &me.events
}

func (me *_Static) SetText(text string) {
	me.Hwnd().SetWindowText(text)
	boundBox := _CalcTextBoundBox(text, true)
	me.Hwnd().SetWindowPos(win.HWND(0), 0, 0,
		boundBox.Cx, boundBox.Cy, co.SWP_NOZORDER|co.SWP_NOMOVE)
}

//------------------------------------------------------------------------------

// Options for NewStaticRaw().
type StaticRawOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

	// Text to appear in the control, passed to CreateWindowEx().
	// Defaults to empty string.
	Text string
	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// Static control styles, passed to CreateWindowEx().
	// Defaults to SS_LEFT | SS_NOTIFY.
	StaticStyles co.SS
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_VISIBLE.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStyles co.WS_EX
}

func (opts *StaticRawOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.StaticStyles == 0 {
		opts.StaticStyles = co.SS_LEFT | co.SS_NOTIFY
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_NONE
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
	me.events = ctrl.parent.On()
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
