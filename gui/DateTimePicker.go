/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
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
	return me
}
