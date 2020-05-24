/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"wingows/api"
)

// Native edit control (textbox).
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type Edit struct {
	controlNativeBase
}

// Optional; returns a, Edit with a specific control ID.
func MakeEdit(ctrlId api.ID) Edit {
	return Edit{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *Edit) Create(parent Window, x, y int32, width, height uint32,
	initialText string, exStyles api.WS_EX, styles api.WS,
	editStyles api.ES) *Edit {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles, "EDIT", initialText,
		styles|api.WS(editStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Edit control will have ES_MULTILINE and ES_WANTRETURN
// styles. Position and size will be adjusted to the current system DPI.
func (me *Edit) CreateMultiLine(parent Window, x, y int32,
	width, height uint32, initialText string) *Edit {

	return me.Create(parent, x, y, width, height, initialText,
		api.WS_EX_CLIENTEDGE,
		api.WS_CHILD|api.WS_GROUP|api.WS_TABSTOP|api.WS_VISIBLE,
		api.ES_MULTILINE|api.ES_WANTRETURN)
}

// Calls CreateWindowEx(). Edit control will have ES_PASSWORD style. Position
// and width will be adjusted to the current system DPI. Height will be
// standard.
func (me *Edit) CreatePassword(parent Window, x, y int32, width uint32,
	initialText string) *Edit {

	return me.Create(parent, x, y, width, 21, initialText,
		api.WS_EX_CLIENTEDGE,
		api.WS_CHILD|api.WS_GROUP|api.WS_TABSTOP|api.WS_VISIBLE,
		api.ES_AUTOHSCROLL|api.ES_PASSWORD)
}

// Calls CreateWindowEx(). Position and width will be adjusted to the current
// system DPI. Height will be standard.
func (me *Edit) CreateSimple(parent Window, x, y int32, width uint32,
	initialText string) *Edit {

	return me.Create(parent, x, y, width, 21, initialText,
		api.WS_EX_CLIENTEDGE,
		api.WS_CHILD|api.WS_GROUP|api.WS_TABSTOP|api.WS_VISIBLE,
		api.ES_AUTOHSCROLL)
}

// Replaces the currently selected text in the edit control.
func (me *Edit) ReplaceSelection(newText string) *Edit {
	me.sendEmMessage(api.EM_REPLACESEL, 1,
		api.LPARAM(unsafe.Pointer(api.StrToUtf16Ptr(newText))))
	return me
}

// Selects all the text in the edit control.
// Only has effect if edit control is focused.
func (me *Edit) SelectAll() *Edit {
	return me.SelectRange(0, -1)
}

// Retrieves the selected range of text in the edit control.
func (me *Edit) SelectedRange() (int32, int32) {
	start, firstAfter := int32(0), int32(0)
	me.sendEmMessage(api.EM_GETSEL, api.WPARAM(unsafe.Pointer(&start)),
		api.LPARAM(unsafe.Pointer(&firstAfter)))
	return start, firstAfter - start
}

// Selects a range of text in the edit control.
// Only has effect if edit control is focused.
func (me *Edit) SelectRange(start, length int32) *Edit {
	me.sendEmMessage(api.EM_SETSEL, api.WPARAM(start),
		api.LPARAM(start+length))
	return me
}

func (me *Edit) sendEmMessage(msg api.EM,
	wParam api.WPARAM, lParam api.LPARAM) uintptr {

	return me.controlNativeBase.Hwnd().
		SendMessage(api.WM(msg), wParam, lParam) // simple wrapper
}
