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

// Creates a new SysLink. Call SysLinkOpts() to define the options to be passed
// to the underlying CreateWindowEx().
func NewSysLink(parent AnyParent, opts *_SysLinkO) SysLink {
	opts.lateDefaults()

	me := &_SysLink{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.position, nil)
		boundBox := _CalcTextBoundBox(opts.text, true)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			"SysLink", opts.text, opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, boundBox, win.HMENU(opts.ctrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
	})

	return me
}

// Creates a new SysLink from a dialog resource.
func NewSysLinkDlg(parent AnyParent, ctrlId int) SysLink {
	me := &_SysLink{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return me
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

type _SysLinkO struct {
	ctrlId int

	text        string
	position    win.POINT
	ctrlStyles  co.LWS
	wndStyles   co.WS
	wndExStyles co.WS_EX
}

// Control ID.
// Defaults to an auto-generated ID.
func (o *_SysLinkO) CtrlId(i int) *_SysLinkO { o.ctrlId = i; return o }

// Text to appear in the control, passed to CreateWindowEx().
// Defaults to empty string.
func (o *_SysLinkO) Text(t string) *_SysLinkO { o.text = t; return o }

// Position within parent's client area in pixels.
// Defaults to 0x0. Will be adjusted to the current system DPI.
func (o *_SysLinkO) Position(p win.POINT) *_SysLinkO { _OwPt(&o.position, p); return o }

// SysLink control styles, passed to CreateWindowEx().
// Defaults to LWS_TRANSPARENT.
func (o *_SysLinkO) CtrlStyles(s co.LWS) *_SysLinkO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
func (o *_SysLinkO) WndStyles(s co.WS) *_SysLinkO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
// Defaults to WS_EX_NONE.
func (o *_SysLinkO) WndExStyles(s co.WS_EX) *_SysLinkO { o.wndExStyles = s; return o }

func (o *_SysLinkO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewSysLink().
func SysLinkOpts() *_SysLinkO {
	return &_SysLinkO{
		ctrlStyles: co.LWS_TRANSPARENT,
		wndStyles:  co.WS_CHILD | co.WS_VISIBLE,
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
