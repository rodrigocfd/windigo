//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [status bar] control.
//
// [status bar]: https://learn.microsoft.com/en-us/windows/win32/controls/status-bars
type StatusBar struct {
	_BaseCtrl
	events EventsStatusBar
	Parts  CollectionStatusBarParts // Methods to interact with the parts collection.
}

// Creates a new [StatusBar] with [win.CreateWindowEx].
func NewStatusBar(parent Parent) *StatusBar {
	ctrlId := nextCtrlId() // always give it an auto ID
	me := &StatusBar{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsStatusBar{ctrlId, &parent.base().userEvents},
	}
	me.Parts.owner = me

	parent.base().beforeUserEvents.Wm(parent.base().wndTy.initMsg(), func(_ Wm) uintptr {
		sbStyle := co.WS_CHILD | co.WS_VISIBLE | co.WS(co.SBARS_TOOLTIPS)
		parentStyle, _ := parent.Hwnd().Style()
		isParentResizable := (parentStyle&co.WS_MAXIMIZEBOX) != 0 ||
			(parentStyle&co.WS_SIZEBOX) != 0
		if isParentResizable {
			sbStyle |= co.WS(co.SBARS_SIZEGRIP)
		}

		me.createWindow(co.WS_EX_NONE, "msctls_statusbar32", "",
			sbStyle, win.POINT{}, win.SIZE{}, parent, false)
		return 0 // ignored
	})

	parent.base().beforeUserEvents.WmSize(func(p WmSize) {
		me.Parts.resizeToFitParent(p)
	})

	return me
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *StatusBar) On() *EventsStatusBar {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Native [status bar] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [status bar]: https://learn.microsoft.com/en-us/windows/win32/controls/status-bars
type EventsStatusBar struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [NM_CLICK] message handler.
//
// [NM_CLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-click-status-bar
func (me *EventsStatusBar) NmClick(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CLICK, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_DBLCLK] message handler.
//
// [NM_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-dblclk-status-bar
func (me *EventsStatusBar) NmDblClk(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_DBLCLK, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_RCLICK] message handler.
//
// [NM_RCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rclick-status-bar
func (me *EventsStatusBar) NmRClick(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RCLICK, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_RDBLCLK] message handler.
//
// [NM_RDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-status-bar
func (me *EventsStatusBar) NmRDblClk(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RDBLCLK, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [SBN_SIMPLEMODECHANGE] message handler.
//
// [SBN_SIMPLEMODECHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/sbn-simplemodechange
func (me *EventsStatusBar) SbnSimpleModeChange(fun func(p *win.NMMOUSE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.SBN_SIMPLEMODECHANGE, func(p unsafe.Pointer) uintptr {
		fun((*win.NMMOUSE)(p))
		return me.parentEvents.defProcVal
	})
}
