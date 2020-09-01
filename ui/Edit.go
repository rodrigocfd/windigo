/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Native edit control (textbox).
//
// https://docs.microsoft.com/en-us/windows/win32/controls/about-edit-controls
type Edit struct {
	_ControlNativeBase
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *Edit) Create(
	parent Window, ctrlId, x, y int32, width, height uint32, initialText string,
	exStyles co.WS_EX, styles co.WS, editStyles co.ES) *Edit {

	x, y, width, height = _Util.MultiplyDpi(x, y, width, height)

	me._ControlNativeBase.create(exStyles, "EDIT", initialText,
		styles|co.WS(editStyles), x, y, width, height, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with ES_MULTILINE | ES_WANTRETURN.
//
// Position and size will be adjusted to the current system DPI.
func (me *Edit) CreateMultiLine(
	parent Window, ctrlId, x, y int32, width, height uint32,
	initialText string) *Edit {

	return me.Create(parent, ctrlId, x, y, width, height, initialText,
		co.WS_EX_CLIENTEDGE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.ES_MULTILINE|co.ES_WANTRETURN)
}

// Calls CreateWindowEx() with ES_AUTOHSCROLL | ES_PASSWORD.
//
// Position and size will be adjusted to the current system DPI.
func (me *Edit) CreatePassword(
	parent Window, ctrlId, x, y int32, width uint32, initialText string) *Edit {

	return me.Create(parent, ctrlId, x, y, width, 21, initialText,
		co.WS_EX_CLIENTEDGE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.ES_AUTOHSCROLL|co.ES_PASSWORD)
}

// Calls CreateWindowEx() with ES_AUTOHSCROLL.
//
// Position and size will be adjusted to the current system DPI.
func (me *Edit) CreateSimple(
	parent Window, ctrlId, x, y int32, width uint32, initialText string) *Edit {

	return me.Create(parent, ctrlId, x, y, width, 21, initialText,
		co.WS_EX_CLIENTEDGE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.ES_AUTOHSCROLL)
}

// Replaces the currently selected text in the edit control.
func (me *Edit) ReplaceSelection(newText string) *Edit {
	me.sendEmMessage(co.EM_REPLACESEL, 1,
		win.LPARAM(unsafe.Pointer(win.StrToPtr(newText))))
	return me
}

// Selects all the text in the edit control.
//
// Only has effect if edit control is focused.
func (me *Edit) SelectAll() *Edit {
	return me.SelectRange(0, -1)
}

// Retrieves the selected range of text in the edit control.
func (me *Edit) SelectedRange() (int32, int32) {
	start, firstAfter := int32(0), int32(0)
	me.sendEmMessage(co.EM_GETSEL, win.WPARAM(unsafe.Pointer(&start)),
		win.LPARAM(unsafe.Pointer(&firstAfter)))
	return start, firstAfter - start
}

// Selects a range of text in the edit control.
//
// Only has effect if edit control is focused.
func (me *Edit) SelectRange(start, length int32) *Edit {
	me.sendEmMessage(co.EM_SETSEL, win.WPARAM(start),
		win.LPARAM(start+length))
	return me
}

// Syntactic sugar.
func (me *Edit) sendEmMessage(msg co.EM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.Hwnd().SendMessage(co.WM(msg), wParam, lParam)
}
