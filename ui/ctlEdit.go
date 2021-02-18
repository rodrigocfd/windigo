/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Native edit control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/about-edit-controls
type Edit struct {
	*_NativeControlBase
	events *_EventsEdit
}

// Constructor. Optionally receives a control ID.
func NewEdit(parent Parent, ctrlId ...int) *Edit {
	base := _NewNativeControlBase(parent, ctrlId...)
	return &Edit{
		_NativeControlBase: base,
		events:             _NewEventsEdit(base),
	}
}

// Calls CreateWindowEx(). With this method, you must also specify WS and WS_EX
// window styles.
//
// Position and size will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *Edit) CreateWs(
	initialText string, pos Pos, size Size,
	editStyles co.ES, styles co.WS, exStyles co.WS_EX) *Edit {

	_global.MultiplyDpi(&pos, &size)
	me._NativeControlBase.create("EDIT", initialText, pos, size,
		co.WS(editStyles)|styles, exStyles)
	_global.UiFont().SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE,
// and WS_EX_CLIENTEDGE. Standard height is 23 pixels.
//
// A typical Edit has ES_AUTOHSCROLL, a password adds ES_PASSWORD.
//
// Position and width will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *Edit) Create(
	initialText string, pos Pos, width int, editStyles co.ES) *Edit {

	return me.CreateWs(initialText, pos, Size{Cx: width, Cy: 21}, editStyles,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_CLIENTEDGE)
}

// Calls CreateWindowEx() with WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE,
// and WS_EX_CLIENTEDGE. Both width and height must be specified.
//
// A typical multi-line Edit has ES_MULTILINE | ES_WANTRETURN.
//
// Position and size will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *Edit) CreateSz(
	initialText string, pos Pos, size Size, editStyles co.ES) *Edit {

	return me.CreateWs(initialText, pos, size, editStyles,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_CLIENTEDGE)
}

func (me *Edit) createAsDlgCtrl() { me._NativeControlBase.createAssignDlg() }

// Exposes all Edit notifications.
//
// Cannot be called after the parent window was created.
func (me *Edit) On() *_EventsEdit {
	if me.hwnd != 0 {
		panic("Cannot add notifications after the Edit was created.")
	}
	return me.events
}

// Replaces the currently selected text in the edit control.
func (me *Edit) ReplaceSelection(newText string) *Edit {
	me.Hwnd().SendMessage(co.WM(co.EM_REPLACESEL),
		1, win.LPARAM(unsafe.Pointer(win.Str.ToUint16Ptr(newText))))
	return me
}

// Selects all the text in the edit control.
//
// Only has effect if edit control is focused.
func (me *Edit) SelectAll() *Edit {
	return me.SelectRange(0, -1)
}

// Retrieves the selected range of text in the edit control.
func (me *Edit) SelectedRange() (int, int) {
	start, firstAfter := int(0), int(0)
	me.Hwnd().SendMessage(co.WM(co.EM_GETSEL),
		win.WPARAM(unsafe.Pointer(&start)),
		win.LPARAM(unsafe.Pointer(&firstAfter)))
	return start, firstAfter - start
}

// Selects a range of text in the edit control.
//
// Only has effect if edit control is focused.
func (me *Edit) SelectRange(start, length int) *Edit {
	me.Hwnd().SendMessage(co.WM(co.EM_SETSEL),
		win.WPARAM(start), win.LPARAM(start+length))
	return me
}

//------------------------------------------------------------------------------

// Edit control notifications.
type _EventsEdit struct {
	ctrl *_NativeControlBase
}

// Constructor.
func _NewEventsEdit(ctrl *_NativeControlBase) *_EventsEdit {
	return &_EventsEdit{
		ctrl: ctrl,
	}
}

// https://docs.microsoft.com/en-us/windows/win32/controls/en-align-ltr-ec
func (me *_EventsEdit) EnAlignLtrEc(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.EN_ALIGN_LTR_EC), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/en-align-rtl-ec
func (me *_EventsEdit) EnAlignRtlEc(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.EN_ALIGN_RTL_EC), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/en-change
func (me *_EventsEdit) EnChange(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.EN_CHANGE), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/en-errspace
func (me *_EventsEdit) EnErrSpace(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.EN_ERRSPACE), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/en-hscroll
func (me *_EventsEdit) EnHScroll(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.EN_HSCROLL), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/en-killfocus
func (me *_EventsEdit) EnKillFocus(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.EN_KILLFOCUS), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/en-maxtext
func (me *_EventsEdit) EnMaxText(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.EN_MAXTEXT), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/en-setfocus
func (me *_EventsEdit) EnSetFocus(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.EN_SETFOCUS), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/en-update
func (me *_EventsEdit) EnUpdate(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.EN_UPDATE), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/en-vscroll
func (me *_EventsEdit) EnVScroll(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.EN_VSCROLL), func(_ WmCommand) {
		userFunc()
	})
}
