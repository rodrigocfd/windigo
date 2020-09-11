/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/co"
)

// Native static control (label).
//
// https://docs.microsoft.com/en-us/windows/win32/controls/about-static-controls
type Static struct {
	_ControlNativeBase
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *Static) Create(
	parent Window, ctrlId, x, y int, width, height uint,
	text string, exStyles co.WS_EX, styles co.WS, staStyles co.SS) *Static {

	x, y, width, height = _Ui.MultiplyDpi(x, y, width, height)

	me._ControlNativeBase.create(exStyles, "STATIC", text,
		styles|co.WS(staStyles), x, y, width, height, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with SS_LEFT.
//
// Position will be adjusted to the current system DPI. The size will be
// calculated to fit the text exactly.
func (me *Static) CreateLText(
	parent Window, ctrlId, x, y int, text string) *Static {

	x, y, _, _ = _Ui.MultiplyDpi(x, y, 0, 0)
	cx, cy := calcTextBoundBox(parent.Hwnd(), text, true)

	me._ControlNativeBase.create(co.WS_EX_NONE, "STATIC", text,
		co.WS_CHILD|co.WS_VISIBLE|co.WS(co.SS_LEFT),
		x, y, cx, cy, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Sets the text, and resizes the control to fit it exactly.
//
// To set the text without resizing the control, use Hwnd().SetWindowText().
func (me *Static) SetText(text string) {
	hasAccel := (co.SS(me.Hwnd().GetStyle()) & co.SS_NOPREFIX) == 0
	cx, cy := calcTextBoundBox(me.Hwnd().GetParent(), text, hasAccel)

	me.Hwnd().SetWindowPos(co.SWP_HWND_NONE, 0, 0, int32(cx), int32(cy),
		co.SWP_NOZORDER|co.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
}

// Returns the text without the accelerator ampersands, for example:
// "&He && she" is returned as "He & she".
//
// Use Hwnd().GetWindowText() to retrieve the raw text, with unparsed
// accelerator ampersands.
func (me *Static) Text() string {
	return _Ui.RemoveAccelAmpersands(me.Hwnd().GetWindowText())
}
