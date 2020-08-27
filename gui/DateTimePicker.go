/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"time"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Native date and time picker control.
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
	parent Window, ctrlId, x, y int32, width, height uint32,
	exStyles co.WS_EX, styles co.WS, dtpStyles co.DTS) *DateTimePicker {

	x, y, width, height = _Util.MultiplyDpi(x, y, width, height)

	me._ControlNativeBase.create(exStyles, "SysDateTimePick32", "",
		styles|co.WS(dtpStyles), x, y, width, height, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with DTS_LONGDATEFORMAT.
//
// Position and size will be adjusted to the current system DPI.
func (me *DateTimePicker) CreateLongDate(
	parent Window, ctrlId, x, y int32, width uint32) *DateTimePicker {

	return me.Create(parent, ctrlId, x, y, width, 23, co.WS_EX_NONE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.DTS_LONGDATEFORMAT)
}

// Calls CreateWindowEx() with DTS_SHORTDATEFORMAT.
//
// Position and size will be adjusted to the current system DPI.
func (me *DateTimePicker) CreateShortDate(
	parent Window, ctrlId, x, y int32, width uint32) *DateTimePicker {

	return me.Create(parent, ctrlId, x, y, width, 23, co.WS_EX_NONE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.DTS_SHORTDATEFORMAT)
}

// Calls CreateWindowEx() with DTS_TIMEFORMAT.
//
// Position and size will be adjusted to the current system DPI.
func (me *DateTimePicker) CreateTime(
	parent Window, ctrlId, x, y int32, width uint32) *DateTimePicker {

	return me.Create(parent, ctrlId, x, y, width, 23, co.WS_EX_NONE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.DTS_TIMEFORMAT)
}

// Sets the format string with DTM_SETFORMAT. An empty string will reset it.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/date-and-time-picker-controls#format-strings
func (me *DateTimePicker) SetFormat(format string) *DateTimePicker {
	if len(format) > 0 {
		formatBuf := win.StrToSlice(format)

		if me.sendDtmMessage(co.DTM_SETFORMAT,
			0, win.LPARAM(uintptr(unsafe.Pointer(&formatBuf[0])))) == 0 {

			panic("DTM_SETFORMAT failed.")
		}
	} else {
		if me.sendDtmMessage(co.DTM_SETFORMAT, 0, 0) == 0 {
			panic("DTM_SETFORMAT failed.")
		}
	}

	return me
}

// Sets a new time with DTM_SETSYSTEMTIME.
func (me *DateTimePicker) SetTime(newTime time.Time) *DateTimePicker {
	st := win.SYSTEMTIME{}
	_Util.TimeToSystemtime(newTime, &st)
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

	return _Util.SystemtimeToTime(&st)
}

// Syntactic sugar.
func (me *DateTimePicker) sendDtmMessage(msg co.DTM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.Hwnd().SendMessage(co.WM(msg), wParam, lParam)
}
