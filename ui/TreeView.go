/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/co"
	"windigo/win"
)

// Native tree view control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/tree-view-controls
type TreeView struct {
	_ControlNativeBase
}

// Adds a new root item; returns the newly inserted item.
func (me *TreeView) AddRootItem(text string) *TreeViewItem {
	return me.Item(0).AddChild(text)
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *TreeView) Create(
	parent Window, ctrlId, x, y int, width, height uint,
	exStyles co.WS_EX, styles co.WS,
	tvExStyles co.TVS_EX, tvStyles co.TVS) *TreeView {

	x, y, width, height = _Util.MultiplyDpi(x, y, width, height)

	me._ControlNativeBase.create(exStyles,
		"SysTreeView32", "", styles|co.WS(tvStyles),
		x, y, width, height, parent, ctrlId)

	if tvExStyles != co.TVS_EX_NONE {
		me.SetExtendedStyle(tvExStyles, tvExStyles)
	}
	return me
}

// Calls CreateWindowEx() with TVS_HASLINES | TVS_LINESATROOT | TVS_SHOWSELALWAYS | TVS_HASBUTTONS.
//
// Position and size will be adjusted to the current system DPI.
func (me *TreeView) CreateSimple(
	parent Window, ctrlId, x, y int, width, height uint) *TreeView {

	return me.Create(parent, ctrlId, x, y, width, height,
		co.WS_EX_CLIENTEDGE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.TVS_EX_NONE,
		co.TVS_HASLINES|co.TVS_LINESATROOT|co.TVS_SHOWSELALWAYS|co.TVS_HASBUTTONS)
}

// Retrieves extended styles with TVM_GETEXTENDEDSTYLE.
func (me *TreeView) ExtendedStyle() co.TVS_EX {
	return co.TVS_EX(me.sendTvmMessage(co.TVM_GETEXTENDEDSTYLE, 0, 0))
}

// Sends TVM_GETNEXTITEM with TVGN_ROOT, returns nil if none.
func (me *TreeView) FirstRootItem() *TreeViewItem {
	return me.Item(0).nextItem(co.TVGN_ROOT)
}

// Returns nil if none.
func (me *TreeView) FirstVisibleItem() *TreeViewItem {
	return me.Item(0).nextItem(co.TVGN_FIRSTVISIBLE)
}

// Returns the item of the given HTREEITEM.
func (me *TreeView) Item(hTreeItem win.HTREEITEM) *TreeViewItem {
	return &TreeViewItem{
		owner:     me,
		hTreeItem: hTreeItem,
	}
}

// Returns the number of items with TVM_GETCOUNT.
func (me *TreeView) ItemCount() uint {
	ret := me.sendTvmMessage(co.TVM_GETCOUNT, 0, 0)
	if ret < 0 {
		panic("TVM_GETCOUNT failed.")
	}
	return uint(ret)
}

// Returns all root items.
func (me *TreeView) RootItems() []TreeViewItem {
	return me.Item(0).Children()
}

// Sends TVM_GETNEXTITEM with TVGN_CARET, returns nil if none.
func (me *TreeView) SelectedItem() *TreeViewItem {
	return me.Item(0).nextItem(co.TVGN_CARET)
}

// Sends TVM_SETEXTENDEDSTYLE.
func (me *TreeView) SetExtendedStyle(mask, exStyle co.TVS_EX) *TreeView {
	me.sendTvmMessage(co.TVM_SETEXTENDEDSTYLE,
		win.WPARAM((mask)), win.LPARAM(exStyle))
	return me
}

// Sends WM_SETREDRAW to enable or disable UI updates.
func (me *TreeView) SetRedraw(allowRedraw bool) *TreeView {
	me.hwnd.SendMessage(co.WM_SETREDRAW,
		win.WPARAM(_Util.BoolToUint32(allowRedraw)), 0)
	return me
}

// Syntactic sugar.
func (me *TreeView) sendTvmMessage(msg co.TVM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.Hwnd().SendMessage(co.WM(msg), wParam, lParam)
}
