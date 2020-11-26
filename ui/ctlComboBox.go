/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Native combo box control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/about-combo-boxes
type ComboBox struct {
	*_NativeControlBase
	events *_EventsComboBox
}

// Constructor. Optionally receives a control ID.
func NewComboBox(parent Parent, ctrlId ...int) *ComboBox {
	base := _NewNativeControlBase(parent, ctrlId...)
	return &ComboBox{
		_NativeControlBase: base,
		events:             _NewEventsComboBox(base),
	}
}

// Calls CreateWindowEx(). With this method, you must also specify WS and WS_EX
// window styles.
//
// Position and width will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *ComboBox) CreateWs(
	pos Pos, width int,
	cbStyles co.CBS, styles co.WS, exStyles co.WS_EX) *ComboBox {

	size := Size{width, 0}
	_global.MultiplyDpi(&pos, &size)
	me._NativeControlBase.create("COMBOBOX", "", pos, size,
		co.WS(cbStyles)|styles, exStyles)
	_global.UiFont().SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
//
// A typical editable/sorted ComboBox has CBS_DROPDOWN | CBS_SORT,
// a non-editable has CBS_DROPDOWNLIST.
//
// Position and width will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *ComboBox) Create(pos Pos, width int, cbStyles co.CBS) *ComboBox {
	return me.CreateWs(pos, width, cbStyles,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_NONE)
}

func (me *ComboBox) createAsDlgCtrl() { me._NativeControlBase.createAssignDlg() }

// Exposes all Button notifications.
//
// Cannot be called after the parent window was created.
func (me *ComboBox) On() *_EventsComboBox {
	if me.hwnd != 0 {
		panic("Cannot add notifications after the ComboBox was created.")
	}
	return me.events
}

// Adds items to the list.
func (me *ComboBox) Add(texts ...string) *ComboBox {
	for _, text := range texts {
		me.Hwnd().SendMessage(co.WM(co.CB_ADDSTRING),
			0, win.LPARAM(unsafe.Pointer(win.Str.ToUint16Ptr(text))))
	}
	return me
}

// Returns the number of items.
func (me *ComboBox) Count() int {
	return int(me.Hwnd().SendMessage(co.WM(co.CB_GETCOUNT), 0, 0))
}

// Limits the length of the text the user may type with CB_LIMITTEXT. Pass zero
// to remove the limitation.
//
// Works only if the ComboBox is editable.
func (me *ComboBox) LimitEditText(numChars int) *ComboBox {
	me.Hwnd().SendMessage(co.WM(co.CB_LIMITTEXT), win.WPARAM(numChars), 0)
	return me
}

// Returns the index of the selected item, or -1 if none.
func (me *ComboBox) SelectedIndex() int {
	return int(me.Hwnd().SendMessage(co.WM(co.CB_GETCURSEL), 0, 0))
}

// Sets the index of the selected item, or -1 to clear.
func (me *ComboBox) SelectIndex(index int) *ComboBox {
	me.Hwnd().SendMessage(co.WM(co.CB_SETCURSEL), win.WPARAM(index), 0)
	return me
}

// Returns the selected text, if any.
func (me *ComboBox) SelectedText() (string, bool) {
	idx := me.SelectedIndex()
	if idx < 0 {
		return "", false
	}
	return me.Text(idx)
}

// Returns the string at the given index, if any.
//
// In an editable ComboBox, the text typed by the user can be retrieved with
// Hwnd().GetWindowText().
func (me *ComboBox) Text(index int) (string, bool) {
	len := int(me.Hwnd().SendMessage(co.WM(co.CB_GETLBTEXTLEN),
		win.WPARAM(index), 0))

	if len == -1 {
		return "", false
	} else if len == 0 {
		return "", true
	}

	buf := make([]uint16, len+1)
	me.Hwnd().SendMessage(co.WM(co.CB_GETLBTEXT),
		win.WPARAM(index), win.LPARAM(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf), true
}

//------------------------------------------------------------------------------

// ComboBox control notifications.
type _EventsComboBox struct {
	ctrl *_NativeControlBase
}

// Constructor.
func _NewEventsComboBox(ctrl *_NativeControlBase) *_EventsComboBox {
	return &_EventsComboBox{
		ctrl: ctrl,
	}
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cbn-closeup
func (me *_EventsComboBox) CbnCloseUp(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.CBN_CLOSEUP), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cbn-dblclk
func (me *_EventsComboBox) CbnDblClk(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.CBN_DBLCLK), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cbn-dropdown
func (me *_EventsComboBox) CbnDropDown(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.CBN_DROPDOWN), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cbn-editchange
func (me *_EventsComboBox) CbnEditChange(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.CBN_EDITCHANGE), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cbn-editupdate
func (me *_EventsComboBox) CbnEditUpdate(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.CBN_EDITUPDATE), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cbn-errspace
func (me *_EventsComboBox) CbnErrSpace(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.CBN_ERRSPACE), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cbn-killfocus
func (me *_EventsComboBox) CbnKillFocus(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.CBN_KILLFOCUS), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cbn-selchange
func (me *_EventsComboBox) CbnSelChange(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.CBN_SELCHANGE), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cbn-selendcancel
func (me *_EventsComboBox) CbnSelEndCancel(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.CBN_SELENDCANCEL), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cbn-selendok
func (me *_EventsComboBox) CbnSelEndOk(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.CBN_SELENDOK), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cbn-setfocus
func (me *_EventsComboBox) CbnSetFocus(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.CBN_SETFOCUS), func(_ WmCommand) {
		userFunc()
	})
}
