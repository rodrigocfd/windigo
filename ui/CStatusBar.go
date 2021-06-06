package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native status bar control
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/status-bars
type StatusBar interface {
	AnyNativeControl

	// Exposes all the StatusBar notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-status-bars-reference-notifications
	On() *_StatusBarEvents

	// Parts methods.
	Parts() *_StatusBarParts
}

//------------------------------------------------------------------------------

type _StatusBar struct {
	_NativeControlBase
	events _StatusBarEvents
	parts  _StatusBarParts
}

// Creates a new StatusBar.
func NewStatusBar(parent AnyParent) StatusBar {
	me := _StatusBar{}
	me._NativeControlBase.new(parent, _NextCtrlId()) // always automatic
	me.events.new(&me._NativeControlBase)
	me.parts.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		sbStyles := co.WS_CHILD | co.WS_VISIBLE | co.WS(co.SBARS_TOOLTIPS)

		parentStyle := co.WS(parent.Hwnd().GetWindowLongPtr(co.GWLP_STYLE))
		isParentResizable := (parentStyle&co.WS_MAXIMIZEBOX) != 0 ||
			(parentStyle&co.WS_SIZEBOX) != 0

		if isParentResizable {
			sbStyles |= co.WS(co.SBARS_SIZEGRIP)
		}

		me._NativeControlBase.createWindow(co.WS_EX(0), "msctls_statusbar32", "",
			sbStyles, win.POINT{}, win.SIZE{}, win.HMENU(me.CtrlId()))
	})

	parent.internalOn().addMsgZero(co.WM_SIZE, func(p wm.Any) {
		me.parts.resizeToFitParent(wm.Size{Msg: p})
	})

	return &me
}

func (me *_StatusBar) On() *_StatusBarEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the StatusBar is created.")
	}
	return &me.events
}

func (me *_StatusBar) Parts() *_StatusBarParts {
	return &me.parts
}

//------------------------------------------------------------------------------

// StatusBar control notifications.
type _StatusBarEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_StatusBarEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.parent.On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-status-bar
func (me *_StatusBarEvents) NmClick(userFunc func(p *win.NMMOUSE) bool) {
	me.events.addNfyRet(me.ctrlId, co.NM_CLICK, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-dblclk-status-bar
func (me *_StatusBarEvents) NmDblClk(userFunc func(p *win.NMMOUSE) bool) {
	me.events.addNfyRet(me.ctrlId, co.NM_DBLCLK, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-rclick-status-bar
func (me *_StatusBarEvents) NmRClick(userFunc func(p *win.NMMOUSE) bool) {
	me.events.addNfyRet(me.ctrlId, co.NM_RCLICK, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-status-bar
func (me *_StatusBarEvents) NmRDblClk(userFunc func(p *win.NMMOUSE) bool) {
	me.events.addNfyRet(me.ctrlId, co.NM_RDBLCLK, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/sbn-simplemodechange
func (me *_StatusBarEvents) SbnSimpleModeChange(userFunc func(p *win.NMMOUSE)) {
	me.events.addNfyZero(me.ctrlId, co.SBN_SIMPLEMODECHANGE, func(p unsafe.Pointer) {
		userFunc((*win.NMMOUSE)(p))
	})
}
