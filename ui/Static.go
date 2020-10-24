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
	parent Window, ctrlId int, pos Pos, size Size, text string,
	staStyles co.SS, styles co.WS, exStyles co.WS_EX) *Static {

	_Ui.MultiplyDpi(&pos, &size)
	me._ControlNativeBase.create(exStyles, "STATIC", text,
		styles|co.WS(staStyles), pos, size, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with SS_LEFT.
//
// Position will be adjusted to the current system DPI. The size will be
// calculated to fit the text exactly.
func (me *Static) CreateLText(
	parent Window, ctrlId int, pos Pos, text string) *Static {

	_Ui.MultiplyDpi(&pos, nil)
	size := _Ui.CalcTextBoundBox(parent.Hwnd(), text, true)

	me._ControlNativeBase.create(co.WS_EX_NONE, "STATIC", text,
		co.WS_CHILD|co.WS_VISIBLE|co.WS(co.SS_LEFT),
		pos, size, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Sets the text, and resizes the control to fit it exactly.
//
// To set the text without resizing the control, use Hwnd().SetWindowText().
func (me *Static) SetText(text string) *Static {
	hasAccel := (co.SS(me.Hwnd().GetStyle()) & co.SS_NOPREFIX) == 0
	size := _Ui.CalcTextBoundBox(me.Hwnd().GetParent(), text, hasAccel)

	me.Hwnd().SetWindowPos(co.SWP_HWND_NONE,
		0, 0, int32(size.Cx), int32(size.Cy),
		co.SWP_NOZORDER|co.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
	return me
}

// Returns the text without the accelerator ampersands, for example:
// "&He && she" is returned as "He & she".
//
// Use Hwnd().GetWindowText() to retrieve the raw text, with unparsed
// accelerator ampersands.
func (me *Static) Text() string {
	return _Ui.RemoveAccelAmpersands(me.Hwnd().GetWindowText())
}
