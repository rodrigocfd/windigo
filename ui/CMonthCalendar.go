package ui

import (
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native month calendar control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/month-calendar-controls
type MonthCalendar interface {
	AnyNativeControl

	// Exposes all the MonthCalendar notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-month-calendar-control-reference-notifications
	On() *_MonthCalendarEvents

	Selected() time.Time        // Retrieves the selected date.
	SetSelected(date time.Time) // Sets the selected date.
}

//------------------------------------------------------------------------------

type _MonthCalendar struct {
	_NativeControlBase
	events _MonthCalendarEvents
}

// Creates a new MonthCalendar. Call MonthCalendarOpts() to define the options
// to be passed to the underlying CreateWindowEx().
func NewMonthCalendar(parent AnyParent, opts *_MonthCalendarO) MonthCalendar {
	opts.lateDefaults()

	me := &_MonthCalendar{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.position, nil)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			"SysMonthCal32", "", opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, win.SIZE{}, win.HMENU(opts.ctrlId))

		rcBound := win.RECT{}
		me.Hwnd().SendMessage(co.MCM_GETMINREQRECT,
			0, win.LPARAM(unsafe.Pointer(&rcBound)))
		me.Hwnd().SetWindowPos(win.HWND(0), 0, 0, rcBound.Right, rcBound.Bottom,
			co.SWP_NOZORDER|co.SWP_NOMOVE)
	})

	return me
}

// Creates a new MonthCalendar from a dialog resource.
func NewMonthCalendarDlg(parent AnyParent, ctrlId int) MonthCalendar {
	me := &_MonthCalendar{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return me
}

func (me *_MonthCalendar) On() *_MonthCalendarEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the MonthCalendar is created.")
	}
	return &me.events
}

func (me *_MonthCalendar) Selected() time.Time {
	st := win.SYSTEMTIME{}
	me.Hwnd().SendMessage(co.MCM_GETCURSEL, 0, win.LPARAM(unsafe.Pointer(&st)))
	return st.ToTime()
}

func (me *_MonthCalendar) SetSelected(date time.Time) {
	st := win.SYSTEMTIME{}
	st.FromTime(date)
	me.Hwnd().SendMessage(co.MCM_SETCURSEL, 0, win.LPARAM(unsafe.Pointer(&st)))
}

//------------------------------------------------------------------------------

type _MonthCalendarO struct {
	ctrlId int

	position    win.POINT
	ctrlStyles  co.MCS
	wndStyles   co.WS
	wndExStyles co.WS_EX
}

// Control ID.
// Defaults to an auto-generated ID.
func (o *_MonthCalendarO) CtrlId(i int) *_MonthCalendarO { o.ctrlId = i; return o }

// Position within parent's client area in pixels.
// Defaults to 0x0. Will be adjusted to the current system DPI.
func (o *_MonthCalendarO) Position(p win.POINT) *_MonthCalendarO { _OwPt(&o.position, p); return o }

// MonthCalendar control styles, passed to CreateWindowEx().
// Defaults to MCS_NONE.
func (o *_MonthCalendarO) CtrlStyles(s co.MCS) *_MonthCalendarO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *_MonthCalendarO) WndStyles(s co.WS) *_MonthCalendarO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
// Defaults to WS_EX_NONE.
func (o *_MonthCalendarO) WndExStyles(s co.WS_EX) *_MonthCalendarO { o.wndExStyles = s; return o }

func (o *_MonthCalendarO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewMonthCalendar().
func MonthCalendarOpts() *_MonthCalendarO {
	return &_MonthCalendarO{
		ctrlStyles: co.MCS_NONE,
		wndStyles:  co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
	}
}

//------------------------------------------------------------------------------

// MonthCalendar control notifications.
type _MonthCalendarEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_MonthCalendarEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.parent.On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/mcn-getdaystate
func (me *_MonthCalendarEvents) McnGetDayState(userFunc func(p *win.NMDAYSTATE)) {
	me.events.addNfyZero(me.ctrlId, co.MCN_GETDAYSTATE, func(p unsafe.Pointer) {
		userFunc((*win.NMDAYSTATE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/mcn-selchange
func (me *_MonthCalendarEvents) McnSelChange(userFunc func(p *win.NMSELCHANGE)) {
	me.events.addNfyZero(me.ctrlId, co.MCN_SELCHANGE, func(p unsafe.Pointer) {
		userFunc((*win.NMSELCHANGE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/mcn-select
func (me *_MonthCalendarEvents) McnSelect(userFunc func(p *win.NMSELCHANGE)) {
	me.events.addNfyZero(me.ctrlId, co.MCN_SELECT, func(p unsafe.Pointer) {
		userFunc((*win.NMSELCHANGE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/mcn-viewchange
func (me *_MonthCalendarEvents) McnViewChange(userFunc func(p *win.NMVIEWCHANGE)) {
	me.events.addNfyZero(me.ctrlId, co.MCN_VIEWCHANGE, func(p unsafe.Pointer) {
		userFunc((*win.NMVIEWCHANGE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-monthcal-
func (me *_MonthCalendarEvents) NmReleasedCapture(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) {
		userFunc()
	})
}