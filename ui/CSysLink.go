package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native SysLink control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/syslink-control-entry
type SysLink interface {
	AnyNativeControl

	// Exposes all the SysLink notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-syslink-control-reference-notifications
	On() *_SysLinkEvents

	SetText(text string) // Sets the text and resizes the control to fit it exactly.
}

//------------------------------------------------------------------------------

type _SysLink struct {
	_NativeControlBase
	events _SysLinkEvents
}

// Creates a new SysLink specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewSysLinkRaw(parent AnyParent, opts SysLinkRawOpts) SysLink {
	opts.fillBlankValuesWithDefault()

	me := _SysLink{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.Position, nil)
		boundBox := _CalcTextBoundBox(opts.Text, true)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"SysLink", opts.Text, opts.Styles|co.WS(opts.SysLinkStyles),
			opts.Position, boundBox, win.HMENU(opts.CtrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
	})

	return &me
}

// Creates a new SysLink from a dialog resource.
func NewSysLinkDlg(parent AnyParent, ctrlId int) SysLink {
	me := _SysLink{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return &me
}

func (me *_SysLink) On() *_SysLinkEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the SysLink is created.")
	}
	return &me.events
}

func (me *_SysLink) SetText(text string) {
	me.Hwnd().SetWindowText(text)
	boundBox := _CalcTextBoundBox(text, true)
	me.Hwnd().SetWindowPos(win.HWND(0), 0, 0,
		boundBox.Cx, boundBox.Cy, co.SWP_NOZORDER|co.SWP_NOMOVE)
}

//------------------------------------------------------------------------------

// Options for NewSysLinkRaw().
type SysLinkRawOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

	// Text to appear in the control, passed to CreateWindowEx().
	// Defaults to empty string.
	Text string
	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// SysLink control styles, passed to CreateWindowEx().
	// Defaults to LWS_TRANSPARENT.
	SysLinkStyles co.LWS
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_VISIBLE.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStyles co.WS_EX
}

func (opts *SysLinkRawOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.SysLinkStyles == 0 {
		opts.SysLinkStyles = co.LWS_TRANSPARENT
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_NONE
	}
}

//------------------------------------------------------------------------------

// SysLink control notifications.
type _SysLinkEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_SysLinkEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.parent.On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-syslink
func (me *_SysLinkEvents) NmClick(userFunc func(p *win.NMLINK)) {
	me.events.addNfyZero(me.ctrlId, co.NM_CLICK, func(p unsafe.Pointer) {
		userFunc((*win.NMLINK)(p))
	})
}
