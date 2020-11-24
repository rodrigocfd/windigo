/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"time"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Native date and time picker control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/date-and-time-picker-controls
type DateTimePicker struct {
	*_NativeControlBase
	events *_EventsDateTimePicker
}

// Constructor. Optionally receives a control ID.
func NewDateTimePicker(parent Parent, ctrlId ...int) *DateTimePicker {
	base := _NewNativeControlBase(parent, ctrlId...)
	return &DateTimePicker{
		_NativeControlBase: base,
		events:             _NewEventsDateTimePicker(base),
	}
}

// Calls CreateWindowEx(). With this method, you must also specify WS and WS_EX
// window styles.
//
// Position and size will be adjusted to the current system DPI.
func (me *DateTimePicker) CreateWs(
	pos Pos, size Size,
	dtpStyles co.DTS, styles co.WS, exStyles co.WS_EX) *DateTimePicker {

	_global.MultiplyDpi(&pos, &size)
	me._NativeControlBase.create("SysDateTimePick32", "", pos, size,
		co.WS(dtpStyles)|styles, exStyles)
	_global.UiFont().SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
// Standard height is 21 pixels.
//
// A typical long date DateTimePicker has DTS_LONGDATEFORMAT, a short date has
// DTS_SHORTDATEFORMAT, and a time has DTS_TIMEFORMAT.
//
// Position and width will be adjusted to the current system DPI.
func (me *DateTimePicker) Create(
	pos Pos, width int, dtpStyles co.DTS) *DateTimePicker {

	return me.CreateWs(pos, Size{Cx: width, Cy: 21}, dtpStyles,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_NONE)
}

// Exposes all DateTimePicker notifications.
func (me *DateTimePicker) On() *_EventsDateTimePicker {
	if me.hwnd != 0 {
		panic("Cannot add notifications after the DateTimePicker was created.")
	}
	return me.events
}

// Sets the format string with DTM_SETFORMAT. An empty string will reset it.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/date-and-time-picker-controls#format-strings
func (me *DateTimePicker) SetFormat(format string) *DateTimePicker {
	pFmt := unsafe.Pointer(nil)
	if len(format) > 0 {
		fmtBuf := win.Str.ToUint16Slice(format)
		pFmt = unsafe.Pointer(&fmtBuf[0])
	}

	if me.Hwnd().SendMessage(co.WM(co.DTM_SETFORMAT), 0, win.LPARAM(pFmt)) == 0 {
		panic("DTM_SETFORMAT failed.")
	}
	return me
}

// Sets a new time with DTM_SETSYSTEMTIME.
func (me *DateTimePicker) SetTime(newTime time.Time) *DateTimePicker {
	st := win.SYSTEMTIME{}
	_global.TimeToSystemtime(newTime, &st)
	me.Hwnd().SendMessage(co.WM(co.DTM_SETSYSTEMTIME),
		win.WPARAM(co.GDT_VALID), win.LPARAM(unsafe.Pointer(&st)))
	return me
}

// Retrieves the time.
func (me *DateTimePicker) Time() time.Time {
	st := win.SYSTEMTIME{}
	ret := co.GDT(me.Hwnd().SendMessage(co.WM(co.DTM_GETSYSTEMTIME),
		0, win.LPARAM(unsafe.Pointer(&st))))

	if ret != co.GDT_VALID {
		panic("DTM_GETSYSTEMTIME failed.")
	}
	return _global.SystemtimeToTime(&st)
}

//------------------------------------------------------------------------------

// DateTimePicker control notifications.
type _EventsDateTimePicker struct {
	ctrl *_NativeControlBase
}

// Constructor.
func _NewEventsDateTimePicker(ctrl *_NativeControlBase) *_EventsDateTimePicker {
	return &_EventsDateTimePicker{
		ctrl: ctrl,
	}
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-closeup
func (me *_EventsDateTimePicker) DtnCloseUp(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.DTN_CLOSEUP), func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-datetimechange
func (me *_EventsDateTimePicker) DtnDateTimeChange(userFunc func(p *win.NMDATETIMECHANGE)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.DTN_DATETIMECHANGE), func(p unsafe.Pointer) {
		userFunc((*win.NMDATETIMECHANGE)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-dropdown
func (me *_EventsDateTimePicker) DtnDropDown(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.DTN_DROPDOWN), func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-format
func (me *_EventsDateTimePicker) DtnFormat(userFunc func(p *win.NMDATETIMEFORMAT)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.DTN_FORMAT), func(p unsafe.Pointer) {
		userFunc((*win.NMDATETIMEFORMAT)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-formatquery
func (me *_EventsDateTimePicker) DtnFormatQuery(userFunc func(p *win.NMDATETIMEFORMATQUERY)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.DTN_FORMATQUERY), func(p unsafe.Pointer) {
		userFunc((*win.NMDATETIMEFORMATQUERY)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-userstring
func (me *_EventsDateTimePicker) DtnUserString(userFunc func(p *win.NMDATETIMESTRING)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.DTN_USERSTRING), func(p unsafe.Pointer) {
		userFunc((*win.NMDATETIMESTRING)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-wmkeydown
func (me *_EventsDateTimePicker) DtnWmKeyDown(userFunc func(p *win.NMDATETIMEWMKEYDOWN)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.DTN_WMKEYDOWN), func(p unsafe.Pointer) {
		userFunc((*win.NMDATETIMEWMKEYDOWN)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-killfocus-date-time
func (me *_EventsDateTimePicker) NmKillFocus(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_KILLFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-setfocus-date-time-
func (me *_EventsDateTimePicker) NmSetFocus(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_SETFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}
