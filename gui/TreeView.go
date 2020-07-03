/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
	"wingows/win"
)

// Native tree view control.
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type TreeView struct {
	controlNativeBase
}

// Optional; returns a TreeView with a specific control ID.
func MakeTreeView(ctrlId int32) TreeView {
	return TreeView{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

// Adds a new root item; returns the newly inserted item.
func (me *TreeView) AddRoot(text string) *TreeViewItem {
	return me.Item(0).AddChild(text)
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *TreeView) Create(parent Window, x, y int32, width, height uint32,
	exStyles co.WS_EX, styles co.WS,
	tvExStyles co.TVS_EX, tvStyles co.TVS) *TreeView {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles,
		"SysTreeView32", "", styles|co.WS(tvStyles),
		x, y, width, height, parent)

	if tvExStyles != co.TVS_EX(0) {
		me.SetExtendedStyle(tvExStyles)
	}
	return me
}

// Calls CreateWindowEx(). Tree view control will have lines and buttons.
// Position and size will be adjusted to the current system DPI.
func (me *TreeView) CreateSimple(parent Window, x, y int32,
	width, height uint32) *TreeView {

	return me.Create(parent, x, y, width, height,
		co.WS_EX_CLIENTEDGE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.TVS_EX(0),
		co.TVS_HASLINES|co.TVS_LINESATROOT|co.TVS_SHOWSELALWAYS|co.TVS_HASBUTTONS)
}

func (me *TreeView) ExtendedStyle() co.TVS_EX {
	return co.TVS_EX(me.sendTvmMessage(co.TVM_GETEXTENDEDSTYLE, 0, 0))
}

// Returns nil if none.
func (me *TreeView) FirstRoot() *TreeViewItem {
	return me.Item(0).NextItem(co.TVGN_ROOT)
}

// Returns nil if none.
func (me *TreeView) FirstVisible() *TreeViewItem {
	return me.Item(0).NextItem(co.TVGN_FIRSTVISIBLE)
}

func (me *TreeView) Item(hTreeItem win.HTREEITEM) *TreeViewItem {
	return &TreeViewItem{
		owner:     me,
		hTreeItem: hTreeItem,
	}
}

func (me *TreeView) ItemCount() uint32 {
	ret := me.sendTvmMessage(co.TVM_GETCOUNT, 0, 0)
	if ret < 0 {
		panic("TVM_GETCOUNT failed.")
	}
	return uint32(ret)
}

// Returns nil if none.
func (me *TreeView) NextItem(flags co.TVGN) *TreeViewItem {
	return me.Item(0).NextItem(flags)
}

func (me *TreeView) Roots() []TreeViewItem {
	return me.Item(0).Children()
}

// Returns nil if none.
func (me *TreeView) SelectedItem() *TreeViewItem {
	return me.Item(0).NextItem(co.TVGN_CARET)
}

func (me *TreeView) SetExtendedStyle(exStyle co.TVS_EX) *TreeView {
	me.sendTvmMessage(co.TVM_SETEXTENDEDSTYLE, 0, win.LPARAM(exStyle))
	return me
}

func (me *TreeView) SetRedraw(allowRedraw bool) *TreeView {
	wp := 0
	if allowRedraw {
		wp = 1
	}
	me.hwnd.SendMessage(co.WM_SETREDRAW, win.WPARAM(wp), 0)
	return me
}

// Simple wrapper.
func (me *TreeView) sendTvmMessage(msg co.TVM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.controlNativeBase.Hwnd().
		SendMessage(co.WM(msg), wParam, lParam)
}
