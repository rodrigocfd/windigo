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

// Optional; returns a ListView with a specific control ID.
func MakeListView(ctrlId int32) ListView {
	return ListView{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

func (me *ListView) AddColumn(text string, width uint32) *ListViewColumn {
	lvc := win.LVCOLUMN{
		Mask:    co.LVCF_TEXT | co.LVCF_WIDTH,
		PszText: win.StrToUtf16Ptr(text),
		Cx:      int32(width),
	}
	newIdx := me.sendLvmMessage(co.LVM_INSERTCOLUMN, 0xFFFF,
		win.LPARAM(unsafe.Pointer(&lvc)))
	if int32(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTCOLUMN failed \"%s\".", text))
	}
	return newListViewColumn(me, uint32(newIdx)) // return newly inserted column
}

func (me *ListView) AddColumns(texts []string, widths []uint32) *ListView {
	if len(texts) != len(widths) {
		panic("Mismatch lenghts of texts and widths.")
	}
	for i := range texts {
		me.AddColumn(texts[i], widths[i])
	}
	return me
}

func (me *ListView) AddItem(text string) *ListViewItem {
	lvi := win.LVITEM{
		Mask:    co.LVIF_TEXT,
		PszText: win.StrToUtf16Ptr(text),
		IItem:   0x0FFFFFFF, // insert as the last one
	}
	newIdx := me.sendLvmMessage(co.LVM_INSERTITEM, 0,
		win.LPARAM(unsafe.Pointer(&lvi)))
	if int32(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM failed \"%s\".", text))
	}
	return newListViewItem(me, uint32(newIdx)) // return newly inserted item
}

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

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles|co.WS_EX(lvExStyles),
		"SysListView32", "", styles|co.WS(lvStyles),
		x, y, width, height, parent)
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
	return newListViewColumn(me, index)
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

func (me *ListView) IsGroupViewEnabled() bool {
	return me.sendLvmMessage(co.LVM_ISGROUPVIEWENABLED, 0, 0) >= 0
}

func (me *ListView) Item(index uint32) *ListViewItem {
	numItems := me.ItemCount()
	if index >= numItems {
		panic("Trying to retrieve item with index out of bounds.")
	}
	return newListViewItem(me, index)
}

func (me *ListView) ItemCount() uint32 {
	count := me.sendLvmMessage(co.LVM_GETITEMCOUNT, 0, 0)
	if int32(count) == -1 {
		panic("LVM_GETITEMCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) SelectedItemCount() uint32 {
	count := me.sendLvmMessage(co.LVM_GETSELECTEDCOUNT, 0, 0)
	if int32(count) == -1 {
		panic("LVM_GETSELECTEDCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) SetRedraw(allowRedraw bool) *ListView {
	wp := 0
	if allowRedraw {
		wp = 1
	}
	me.hwnd.SendMessage(co.WM_SETREDRAW, win.WPARAM(wp), 0)
	return me
}

func (me *ListView) SetView(view co.LV_VIEW) *ListView {
	if int32(me.sendLvmMessage(co.LVM_SETVIEW, 0, 0)) == -1 {
		panic("LVM_SETVIEW failed.")
	}
	return me
}

func (me *ListView) View() co.LV_VIEW {
	return co.LV_VIEW(me.sendLvmMessage(co.LVM_GETVIEW, 0, 0))
}

func (me *ListView) sendLvmMessage(msg co.LVM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.controlNativeBase.Hwnd().
		SendMessage(co.WM(msg), wParam, lParam) // simple wrapper
}
