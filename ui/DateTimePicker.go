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

// Native date and time picker (calendar) control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/date-and-time-picker-controls
type DateTimePicker struct {
	_ControlNativeBase
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *DateTimePicker) Create(
	parent Window, ctrlId int, pos Pos, size Size,
	dtpStyles co.DTS, styles co.WS, exStyles co.WS_EX) *DateTimePicker {

	_Ui.MultiplyDpi(&pos, &size)
	me._ControlNativeBase.create(exStyles, "SysDateTimePick32", "",
		styles|co.WS(dtpStyles), pos, size, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with height 23, and DTS_LONGDATEFORMAT.
//
// Position and size will be adjusted to the current system DPI.
func (me *DateTimePicker) CreateLongDate(
	parent Window, ctrlId int, pos Pos, width uint) *DateTimePicker {

	return me.Create(parent, ctrlId, pos, Size{width, 23},
		co.DTS_LONGDATEFORMAT,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_NONE)
}

// Calls CreateWindowEx() with height 23, and DTS_SHORTDATEFORMAT.
//
// Position and size will be adjusted to the current system DPI.
func (me *DateTimePicker) CreateShortDate(
	parent Window, ctrlId int, pos Pos, width uint) *DateTimePicker {

	return me.Create(parent, ctrlId, pos, Size{width, 23},
		co.DTS_SHORTDATEFORMAT,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_NONE)
}

// Calls CreateWindowEx() with height 23, and DTS_TIMEFORMAT.
//
// Position and size will be adjusted to the current system DPI.
func (me *DateTimePicker) CreateTime(
	parent Window, ctrlId int, pos Pos, width uint) *DateTimePicker {

	return me.Create(parent, ctrlId, pos, Size{width, 23},
		co.DTS_TIMEFORMAT,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_NONE)
}

// Sets the format string with DTM_SETFORMAT. An empty string will reset it.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/date-and-time-picker-controls#format-strings
func (me *DateTimePicker) SetFormat(format string) *DateTimePicker {
	var pFormat unsafe.Pointer

	if len(format) > 0 {
		formatBuf := win.Str.ToUint16Slice(format)
		pFormat = unsafe.Pointer(&formatBuf[0])
	} else {
		pFormat = nil
	}

	if me.sendDtmMessage(co.DTM_SETFORMAT, 0, win.LPARAM(pFormat)) == 0 {
		panic("DTM_SETFORMAT failed.")
	}

	return me
}

// Sets a new time with DTM_SETSYSTEMTIME.
func (me *DateTimePicker) SetTime(newTime time.Time) *DateTimePicker {
	st := win.SYSTEMTIME{}
	_Ui.TimeToSystemtime(newTime, &st)
	me.sendDtmMessage(co.DTM_SETSYSTEMTIME,
		win.WPARAM(co.GDT_VALID), win.LPARAM(unsafe.Pointer(&st)))
	return me
}

// Retrieves the time.
func (me *DateTimePicker) Time() time.Time {
	st := win.SYSTEMTIME{}

	if co.GDT(me.sendDtmMessage(co.DTM_GETSYSTEMTIME,
		0, win.LPARAM(unsafe.Pointer(&st)))) != co.GDT_VALID {

		panic("DTM_GETSYSTEMTIME failed.")
	}

	return _Ui.SystemtimeToTime(&st)
}

// Syntactic sugar.
func (me *DateTimePicker) sendDtmMessage(
	msg co.DTM, wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.Hwnd().SendMessage(co.WM(msg), wParam, lParam)
}
