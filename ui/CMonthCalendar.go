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

// Creates a new MonthCalendar specifying all options, which will be passed to
// the underlying CreateWindowEx().
func NewMonthCalendarRaw(parent AnyParent, opts MonthCalendarRawOpts) MonthCalendar {
	opts.fillBlankValuesWithDefault()

	me := _MonthCalendar{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.Position, nil)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"SysMonthCal32", "", opts.Styles|co.WS(opts.MonthCalendarStyles),
			opts.Position, win.SIZE{}, win.HMENU(opts.CtrlId))

		rcBound := win.RECT{}
		me.Hwnd().SendMessage(co.MCM_GETMINREQRECT,
			0, win.LPARAM(unsafe.Pointer(&rcBound)))
		me.Hwnd().SetWindowPos(win.HWND(0), 0, 0, rcBound.Right, rcBound.Bottom,
			co.SWP_NOZORDER|co.SWP_NOMOVE)
	})

	return &me
}

// Creates a new MonthCalendar from a dialog resource.
func NewMonthCalendarDlg(parent AnyParent, ctrlId int) MonthCalendar {
	me := _MonthCalendar{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return &me
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

// Options for NewMonthCalendarRaw().
type MonthCalendarRawOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// MonthCalendar control styles, passed to CreateWindowEx().
	// Defaults to MCS_NONE.
	MonthCalendarStyles co.MCS
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStyles co.WS_EX
}

func (opts *MonthCalendarRawOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.MonthCalendarStyles == 0 {
		opts.MonthCalendarStyles = co.MCS_NONE
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_NONE
	}
}

//------------------------------------------------------------------------------

// MonthCalendar control notifications.
type _MonthCalendarEvents struct {
	ctrlId int
	events *_EventsNfy
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
