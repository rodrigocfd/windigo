/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

// Native edit control (textbox).
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type Edit struct {
	controlNativeBase
}

// Optional; returns a, Edit with a specific control ID.
func MakeEdit(ctrlId c.ID) Edit {
	return Edit{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *Edit) Create(parent Window, x, y int32, width, height uint32,
	initialText string, exStyles c.WS_EX, styles c.WS, editStyles c.ES) *Edit {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles, "EDIT", initialText,
		styles|c.WS(editStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Edit control will have ES_MULTILINE and ES_WANTRETURN
// styles. Position and size will be adjusted to the current system DPI.
func (me *Edit) CreateMultiLine(parent Window, x, y int32,
	width, height uint32, initialText string) *Edit {

	return me.Create(parent, x, y, width, height, initialText,
		c.WS_EX_CLIENTEDGE,
		c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.ES_MULTILINE|c.ES_WANTRETURN)
}

// Calls CreateWindowEx(). Edit control will have ES_PASSWORD style. Position
// and width will be adjusted to the current system DPI. Height will be
// standard.
func (me *Edit) CreatePassword(parent Window, x, y int32, width uint32,
	initialText string) *Edit {

	return me.Create(parent, x, y, width, 21, initialText,
		c.WS_EX_CLIENTEDGE,
		c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.ES_AUTOHSCROLL|c.ES_PASSWORD)
}

// Calls CreateWindowEx(). Position and width will be adjusted to the current
// system DPI. Height will be standard.
func (me *Edit) CreateSimple(parent Window, x, y int32, width uint32,
	initialText string) *Edit {

	return me.Create(parent, x, y, width, 21, initialText,
		c.WS_EX_CLIENTEDGE,
		c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.ES_AUTOHSCROLL)
}

// Replaces the currently selected text in the edit control.
func (me *Edit) ReplaceSelection(newText string) *Edit {
	me.sendEmMessage(c.EM_REPLACESEL, 1,
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
	me.sendEmMessage(c.EM_GETSEL, api.WPARAM(unsafe.Pointer(&start)),
		api.LPARAM(unsafe.Pointer(&firstAfter)))
	return start, firstAfter - start
}

// Selects a range of text in the edit control.
// Only has effect if edit control is focused.
func (me *Edit) SelectRange(start, length int32) *Edit {
	me.sendEmMessage(c.EM_SETSEL, api.WPARAM(start),
		api.LPARAM(start+length))
	return me
}

func (me *Edit) sendEmMessage(msg c.EM,
	wParam api.WPARAM, lParam api.LPARAM) uintptr {

	return me.controlNativeBase.Hwnd().
		SendMessage(c.WM(msg), wParam, lParam) // simple wrapper
}
