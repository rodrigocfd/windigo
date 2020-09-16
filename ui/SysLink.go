/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Native SysLink control, which renders simple HTML.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/syslink-control-entry
type SysLink struct {
	_ControlNativeBase
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *SysLink) Create(
	parent Window, ctrlId, x, y int, width, height uint,
	text string, exStyles co.WS_EX, styles co.WS, slStyles co.LWS) *SysLink {

	x, y, width, height = _Ui.MultiplyDpi(x, y, width, height)

	me._ControlNativeBase.create(exStyles, "SysLink", text,
		styles|co.WS(slStyles), x, y, width, height, parent, ctrlId)
	return me
}

// Calls CreateWindowEx() with LWS_TRANSPARENT.
//
// Position will be adjusted to the current system DPI. The size will be
// calculated to fit the text exactly.
func (me *SysLink) CreateLText(
	parent Window, ctrlId, x, y int, text string) *SysLink {

	x, y, _, _ = _Ui.MultiplyDpi(x, y, 0, 0)

	me._ControlNativeBase.create(co.WS_EX_NONE, "SysLink", text,
		co.WS_CHILD|co.WS_VISIBLE|co.WS_TABSTOP|co.WS(co.LWS_TRANSPARENT),
		x, y, 0, 0, parent, ctrlId) // note zero width & height

	_globalUiFont.SetOnControl(me)

	sz := win.SIZE{}
	me.sendLmMessage(co.LM_GETIDEALSIZE, 0, win.LPARAM(unsafe.Pointer(&sz)))
	me.Hwnd().SetWindowPos(co.SWP_HWND_NONE, 0, 0, sz.Cx, sz.Cy, // calc width & height
		co.SWP_NOZORDER|co.SWP_NOMOVE)

	return me
}

// Sets the text, and resizes the control to fit it exactly.
//
// To set the text without resizing the control, use Hwnd().SetWindowText().
func (me *SysLink) SetText(text string) *SysLink {
	cx, cy := calcTextBoundBox(me.Hwnd().GetParent(), text, false)

	me.Hwnd().SetWindowPos(co.SWP_HWND_NONE, 0, 0, int32(cx), int32(cy),
		co.SWP_NOZORDER|co.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
	return me
}

// Returns the text in the SysLink control.
//
// Syntactic sugar to Hwnd().GetWindowText().
func (me *SysLink) Text() string {
	return me.Hwnd().GetWindowText()
}

// Syntactic sugar.
func (me *SysLink) sendLmMessage(msg co.LM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.Hwnd().SendMessage(co.WM(msg), wParam, lParam)
}
