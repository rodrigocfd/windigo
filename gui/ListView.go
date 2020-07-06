/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Native list view control.
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type ListView struct {
	controlNativeBase
}

// Adds a new column; returns the newly inserted column.
func (me *ListView) AddColumn(text string, width uint32) *ListViewColumn {
	lvc := win.LVCOLUMN{
		Mask:    co.LVCF_TEXT | co.LVCF_WIDTH,
		PszText: win.StrToPtr(text),
		Cx:      int32(width),
	}
	newIdx := me.sendLvmMessage(co.LVM_INSERTCOLUMN, 0xFFFF,
		win.LPARAM(unsafe.Pointer(&lvc)))
	if int32(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTCOLUMN failed \"%s\".", text))
	}
	return &ListViewColumn{
		owner: me,
		index: uint32(newIdx),
	}
}

// Adds many columns at once.
func (me *ListView) AddColumns(texts []string, widths []uint32) *ListView {
	if len(texts) != len(widths) {
		panic("Mismatch lenghts of texts and widths.")
	}
	for i := range texts {
		me.AddColumn(texts[i], widths[i])
	}
	return me
}

// Adds a new item; returns the newly inserted item.
func (me *ListView) AddItem(text string) *ListViewItem {
	lvi := win.LVITEM{
		Mask:    co.LVIF_TEXT,
		PszText: win.StrToPtr(text),
		IItem:   0x0FFFFFFF, // insert as the last one
	}
	newIdx := me.sendLvmMessage(co.LVM_INSERTITEM, 0,
		win.LPARAM(unsafe.Pointer(&lvi)))
	if int32(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM failed \"%s\".", text))
	}
	return &ListViewItem{
		owner: me,
		index: uint32(newIdx),
	}
}

// Adds many items at once.
func (me *ListView) AddItems(texts []string) *ListView {
	for i := range texts {
		me.AddItem(texts[i])
	}
	return me
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *ListView) Create(parent Window, x, y int32, width, height uint32,
	exStyles co.WS_EX, styles co.WS,
	lvExStyles co.LVS_EX, lvStyles co.LVS) *ListView {

	x, y, width, height = globalDpi.multiply(x, y, width, height)

	me.controlNativeBase.create(exStyles,
		"SysListView32", "", styles|co.WS(lvStyles),
		x, y, width, height, parent)

	if lvExStyles != co.LVS_EX(0) {
		me.SetExtendedStyle(lvExStyles)
	}
	return me
}

// Calls CreateWindowEx(). List view control will have LVS_REPORT style.
// Position and size will be adjusted to the current system DPI.
func (me *ListView) CreateReport(parent Window, x, y int32,
	width, height uint32) *ListView {

	return me.Create(parent, x, y, width, height,
		co.WS_EX_CLIENTEDGE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.LVS_EX_FULLROWSELECT,
		co.LVS_REPORT|co.LVS_SHOWSELALWAYS)
}

func (me *ListView) Column(index uint32) *ListViewColumn {
	numCols := me.ColumnCount()
	if index >= numCols {
		panic("Trying to retrieve column with index out of bounds.")
	}
	return &ListViewColumn{
		owner: me,
		index: index,
	}
}

func (me *ListView) ColumnCount() uint32 {
	hHeader := win.HWND(me.sendLvmMessage(co.LVM_GETHEADER, 0, 0))
	if hHeader == 0 {
		panic("LVM_GETHEADER failed.")
	}

	count := hHeader.SendMessage(co.WM(co.HDM_GETITEMCOUNT), 0, 0)
	if int32(count) == -1 {
		panic("HDM_GETITEMCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) DeleteAllItems() *ListView {
	ret := me.sendLvmMessage(co.LVM_DELETEALLITEMS, 0, 0)
	if ret == 0 {
		panic("LVM_DELETEALLITEMS failed.")
	}
	return me
}

// Sends LVM_ISGROUPVIEWENABLED.
func (me *ListView) GroupViewEnabled() bool {
	return me.sendLvmMessage(co.LVM_ISGROUPVIEWENABLED, 0, 0) >= 0
}

// Sends LVM_HITTEST to determine the item at specified position, if any.
func (me *ListView) HitTest(pos win.POINT) *win.LVHITTESTINFO {
	lvhti := &win.LVHITTESTINFO{
		Pt: pos,
	}
	wp := int32(-1) // Vista: retrieve iGroup and iSubItem
	me.sendLvmMessage(co.LVM_HITTEST,
		win.WPARAM(wp), win.LPARAM(unsafe.Pointer(&lvhti)))
	return lvhti
}

func (me *ListView) Item(index uint32) *ListViewItem {
	numItems := me.ItemCount()
	if index >= numItems {
		panic("Trying to retrieve item with index out of bounds.")
	}
	return &ListViewItem{
		owner: me,
		index: index,
	}
}

func (me *ListView) ItemCount() uint32 {
	count := me.sendLvmMessage(co.LVM_GETITEMCOUNT, 0, 0)
	if int32(count) == -1 {
		panic("LVM_GETITEMCOUNT failed.")
	}
	return uint32(count)
}

// Returns nil if none.
func (me *ListView) NextItem(relationship co.LVNI) *ListViewItem {
	idx := int32(-1)
	allItems := ListViewItem{
		owner: me,
		index: uint32(idx),
	}
	return allItems.NextItem(relationship)
}

func (me *ListView) SelectedItemCount() uint32 {
	count := me.sendLvmMessage(co.LVM_GETSELECTEDCOUNT, 0, 0)
	if int32(count) == -1 {
		panic("LVM_GETSELECTEDCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) SetExtendedStyle(exStyle co.LVS_EX) *ListView {
	me.sendLvmMessage(co.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, win.LPARAM(exStyle))
	return me
}

// Returns image list previously associated.
func (me *ListView) SetImageList(typeImgList co.LVSIL,
	hImgList win.HIMAGELIST) win.HIMAGELIST {

	return win.HIMAGELIST(
		me.sendLvmMessage(co.LVM_SETIMAGELIST,
			win.WPARAM(typeImgList), win.LPARAM(hImgList)),
	)
}

func (me *ListView) SetRedraw(allowRedraw bool) *ListView {
	wp := 0
	if allowRedraw {
		wp = 1
	}
	me.hwnd.SendMessage(co.WM_SETREDRAW, win.WPARAM(wp), 0)
	return me
}

func (me *ListView) SetStateAllItems(
	state co.LVIS, stateMask co.LVIS) *ListView {

	idx := int32(-1)
	allItems := ListViewItem{
		owner: me,
		index: uint32(idx),
	}
	allItems.SetState(state, stateMask)
	return me
}

func (me *ListView) SetView(view co.LV_VIEW) *ListView {
	if int32(me.sendLvmMessage(co.LVM_SETVIEW, 0, 0)) == -1 {
		panic("LVM_SETVIEW failed.")
	}
	return me
}

// Returns the width of a string using list view current font.
func (me *ListView) StringWidth(text string) uint32 {
	ret := me.sendLvmMessage(co.LVM_GETSTRINGWIDTH,
		0, win.LPARAM(unsafe.Pointer(win.StrToPtr(text))))
	if ret == 0 {
		panic("LVM_GETSTRINGWIDTH failed.")
	}
	return uint32(ret)
}

func (me *ListView) View() co.LV_VIEW {
	return co.LV_VIEW(me.sendLvmMessage(co.LVM_GETVIEW, 0, 0))
}

// Simple wrapper.
func (me *ListView) sendLvmMessage(msg co.LVM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.controlNativeBase.Hwnd().
		SendMessage(co.WM(msg), wParam, lParam)
}