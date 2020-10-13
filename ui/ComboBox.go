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
	_ControlNativeBase
}

// Adds a single item to the combo box.
func (me *ComboBox) Add(text string) *ComboBox {
	me.sendCbMessage(co.CB_ADDSTRING,
		0, win.LPARAM(unsafe.Pointer(win.Str.ToUint16Ptr(text))))
	return me
}

// Adds many items to the combo box.
func (me *ComboBox) AddMany(texts []string) *ComboBox {
	for _, text := range texts {
		me.Add(text)
	}
	return me
}

// Returns the number of items in the list box of a combo box.
func (me *ComboBox) Count() uint {
	return uint(me.sendCbMessage(co.CB_GETCOUNT, 0, 0))
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *ComboBox) Create(
	parent Window, ctrlId int, pos Pos, width uint,
	exStyles co.WS_EX, styles co.WS, cbStyles co.CBS) *ComboBox {

	size := Size{width, 0}
	_Ui.MultiplyDpi(&pos, &size)

	me._ControlNativeBase.create(exStyles, "COMBOBOX", "",
		styles|co.WS(cbStyles), pos, size, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with CBS_DROPDOWN.
//
// Position and size will be adjusted to the current system DPI.
func (me *ComboBox) CreateEditable(
	parent Window, ctrlId int, pos Pos, width uint) *ComboBox {

	return me.Create(parent, ctrlId, pos, width, co.WS_EX_NONE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.CBS_DROPDOWN)
}

// Calls CreateWindowEx() with CBS_DROPDOWN | CBS_SORT.
//
// Position and size will be adjusted to the current system DPI.
func (me *ComboBox) CreateEditableSorted(
	parent Window, ctrlId int, pos Pos, width uint) *ComboBox {

	return me.Create(parent, ctrlId, pos, width, co.WS_EX_NONE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.CBS_DROPDOWN|co.CBS_SORT)
}

// Calls CreateWindowEx() with CBS_DROPDOWNLIST.
//
// Position and size will be adjusted to the current system DPI.
func (me *ComboBox) CreateFixed(
	parent Window, ctrlId int, pos Pos, width uint) *ComboBox {

	return me.Create(parent, ctrlId, pos, width, co.WS_EX_NONE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.CBS_DROPDOWNLIST)
}

// Calls CreateWindowEx() with CBS_DROPDOWNLIST | CBS_SORT.
//
// Position and size will be adjusted to the current system DPI.
func (me *ComboBox) CreateFixedSorted(
	parent Window, ctrlId int, pos Pos, width uint) *ComboBox {

	return me.Create(parent, ctrlId, pos, width, co.WS_EX_NONE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.CBS_DROPDOWNLIST|co.CBS_SORT)
}

// Limits the length of the text the user may type with CB_LIMITTEXT. Pass zero
// to remove the limitation.
//
// Works only if the combo box is editable.
func (me *ComboBox) LimitEditText(numChars uint) *ComboBox {
	me.sendCbMessage(co.CB_LIMITTEXT, win.WPARAM(numChars), 0)
	return me
}

// Returns the index of the selected item, or -1 if none.
func (me *ComboBox) SelectedIndex() int {
	return int(me.sendCbMessage(co.CB_GETCURSEL, 0, 0))
}

// Sets the index of the selected item, or -1 to clear.
func (me *ComboBox) SelectIndex(index int) *ComboBox {
	me.sendCbMessage(co.CB_SETCURSEL, win.WPARAM(index), 0)
	return me
}

// Returns the selected text, if any.
func (me *ComboBox) SelectedText() (string, bool) {
	idx := me.SelectedIndex()
	if idx < 0 {
		return "", false
	}
	return me.Text(uint(idx))
}

// Returns the string at the given index, if any.
//
// In an editable combo box, the text typed by the user can be retrieved with
// Hwnd().GetWindowText().
func (me *ComboBox) Text(index uint) (string, bool) {
	len := int(me.sendCbMessage(co.CB_GETLBTEXTLEN, win.WPARAM(index), 0))
	if len == -1 {
		return "", false
	} else if len == 0 {
		return "", true
	}

	buf := make([]uint16, len+1)
	me.sendCbMessage(co.CB_GETLBTEXT,
		win.WPARAM(index), win.LPARAM(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf), true
}

// Syntactic sugar.
func (me *ComboBox) sendCbMessage(
	msg co.CB, wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.Hwnd().SendMessage(co.WM(msg), wParam, lParam)
}
