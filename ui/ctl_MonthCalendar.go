//go:build windows

package ui

import (
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [month calendar] control.
//
// [month calendar]: https://learn.microsoft.com/en-us/windows/win32/controls/month-calendar-controls
type MonthCalendar struct {
	_BaseCtrl
	events EventsMonthCalendar
}

// Creates a new MonthCalendar with [CreateWindowEx].
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func NewMonthCalendar(parent Parent, opts *VarOptsMonthCalendar) *MonthCalendar {
	setUniqueCtrlId(&opts.ctrlId)
	me := &MonthCalendar{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsMonthCalendar{opts.ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.WmCreate(func(_ WmCreate) int {
		me.createWindow(opts.wndExStyle, "SysMonthCal32", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, win.SIZE{}, parent, false)

		var rcBound win.RECT
		me.hWnd.SendMessage(co.MCM_GETMINREQRECT, // request the ideal size
			0, win.LPARAM(unsafe.Pointer(&rcBound)))
		me.hWnd.SetWindowPos(win.HWND(0), 0, 0, uint(rcBound.Right), uint(rcBound.Bottom),
			co.SWP_NOZORDER|co.SWP_NOMOVE)

		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		if !opts.value.IsZero() {
			me.SetDate(opts.value)
		}
		return 0 // ignored
	})

	return me
}

// Instantiates a new MonthCalendar to be loaded from a dialog resource with
// [GetDlgItem].
//
// [GetDlgItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgitem
func NewMonthCalendarDlg(parent Parent, ctrlId uint16, layout LAY) *MonthCalendar {
	me := &MonthCalendar{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsMonthCalendar{ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.WmInitDialog(func(_ WmInitDialog) bool {
		me.assignDialog(parent)
		parent.base().layout.Add(parent, me.hWnd, layout)
		return true // ignored
	})

	return me
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *MonthCalendar) On() *EventsMonthCalendar {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Returns the current date with [MCM_GETCURSEL].
//
// [MCM_GETCURSEL]: https://learn.microsoft.com/en-us/windows/win32/controls/mcm-getcursel
func (me *MonthCalendar) Date() time.Time {
	var st win.SYSTEMTIME
	me.hWnd.SendMessage(co.MCM_GETCURSEL, 0, win.LPARAM(unsafe.Pointer(&st)))
	return st.ToTime()
}

// Sets the current date with [MCM_SETCURSEL].
//
// Returns the same object, so further operations can be chained.
//
// [MCM_SETCURSEL]: https://learn.microsoft.com/en-us/windows/win32/controls/mcm-setcursel
func (me *MonthCalendar) SetDate(date time.Time) *MonthCalendar {
	var st win.SYSTEMTIME
	st.SetTime(date)
	me.hWnd.SendMessage(co.MCM_SETCURSEL, 0, win.LPARAM(unsafe.Pointer(&st)))
	return me
}

// Options for ui.NewMonthCalendar(); returned by ui.OptsMonthCalendar().
type VarOptsMonthCalendar struct {
	ctrlId     uint16
	layout     LAY
	position   win.POINT
	ctrlStyle  co.MCS
	wndStyle   co.WS
	wndExStyle co.WS_EX
	value      time.Time
}

// Options for ui.NewMonthCalendar().
func OptsMonthCalendar() *VarOptsMonthCalendar {
	return &VarOptsMonthCalendar{
		wndStyle: co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsMonthCalendar) CtrlId(id uint16) *VarOptsMonthCalendar { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsMonthCalendar) Layout(l LAY) *VarOptsMonthCalendar { o.layout = l; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsMonthCalendar) Position(x, y int) *VarOptsMonthCalendar {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Month calendar control [style], passed to [CreateWindowEx].
//
// Defaults to co.MCS_NONE.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/month-calendar-control-styles
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsMonthCalendar) CtrlStyle(s co.MCS) *VarOptsMonthCalendar { o.ctrlStyle = s; return o }

// Window style, passed to [CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsMonthCalendar) WndStyle(s co.WS) *VarOptsMonthCalendar { o.wndStyle = s; return o }

// Window extended style, passed to [CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsMonthCalendar) WndExStyle(s co.WS_EX) *VarOptsMonthCalendar {
	o.wndExStyle = s
	return o
}

// Initial value.
//
// Defaults to time.Now().
func (o *VarOptsMonthCalendar) Value(t time.Time) *VarOptsMonthCalendar { o.value = t; return o }

// Native [month calendar] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [month calendar]: https://learn.microsoft.com/en-us/windows/win32/controls/month-calendar-controls
type EventsMonthCalendar struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [MCN_GETDAYSTATE] message handler.
//
// [MCN_GETDAYSTATE]: https://learn.microsoft.com/en-us/windows/win32/controls/mcn-getdaystate
func (me *EventsMonthCalendar) McnGetDayState(fun func(p *win.NMDAYSTATE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.MCN_GETDAYSTATE, func(p unsafe.Pointer) uintptr {
		fun((*win.NMDAYSTATE)(p))
		return me.parentEvents.defProcVal
	})
}

// [MCN_SELCHANGE] message handler.
//
// [MCN_SELCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/mcn-selchange
func (me *EventsMonthCalendar) McnSelChange(fun func(p *win.NMSELCHANGE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.MCN_SELCHANGE, func(p unsafe.Pointer) uintptr {
		fun((*win.NMSELCHANGE)(p))
		return me.parentEvents.defProcVal
	})
}

// [MCN_SELECT] message handler.
//
// [MCN_SELECT]: https://learn.microsoft.com/en-us/windows/win32/controls/mcn-select
func (me *EventsMonthCalendar) McnSelect(fun func(p *win.NMSELCHANGE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.MCN_SELECT, func(p unsafe.Pointer) uintptr {
		fun((*win.NMSELCHANGE)(p))
		return me.parentEvents.defProcVal
	})
}

// [MCN_VIEWCHANGE] message handler.
//
// [MCN_VIEWCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/mcn-viewchange
func (me *EventsMonthCalendar) McnViewChange(fun func(p *win.NMVIEWCHANGE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.MCN_VIEWCHANGE, func(p unsafe.Pointer) uintptr {
		fun((*win.NMVIEWCHANGE)(p))
		return me.parentEvents.defProcVal
	})
}

// [NM_RELEASEDCAPTURE] message handler.
//
// [NM_RELEASEDCAPTURE]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-monthcal-
func (me *EventsMonthCalendar) NmReleasedCapture(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}
