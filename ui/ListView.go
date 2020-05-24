/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"unsafe"
	"wingows/api"
)

// Native list view control.
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type ListView struct {
	controlNativeBase
}

// Optional; returns a ListView with a specific control ID.
func MakeListView(ctrlId api.ID) ListView {
	return ListView{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

func (me *ListView) AddColumn(text string, width uint32) *ListViewColumn {
	lvc := api.LVCOLUMN{
		Mask:    api.LVCF_TEXT | api.LVCF_WIDTH,
		PszText: api.StrToUtf16Ptr(text),
		Cx:      int32(width),
	}
	newIdx := me.sendLvmMessage(api.LVM_INSERTCOLUMN, 0xFFFF,
		api.LPARAM(unsafe.Pointer(&lvc)))
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
	lvi := api.LVITEM{
		Mask:    api.LVIF_TEXT,
		PszText: api.StrToUtf16Ptr(text),
		IItem:   0x0FFFFFFF, // insert as the last one
	}
	newIdx := me.sendLvmMessage(api.LVM_INSERTITEM, 0,
		api.LPARAM(unsafe.Pointer(&lvi)))
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
	exStyles api.WS_EX, styles api.WS,
	lvExStyles api.LVS_EX, lvStyles api.LVS) *ListView {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles|api.WS_EX(lvExStyles),
		"SysListView32", "", styles|api.WS(lvStyles),
		x, y, width, height, parent)
	return me
}

// Calls CreateWindowEx(). List view control will have LVS_REPORT style.
// Position and size will be adjusted to the current system DPI.
func (me *ListView) CreateReport(parent Window, x, y int32,
	width, height uint32) *ListView {

	return me.Create(parent, x, y, width, height,
		api.WS_EX_CLIENTEDGE,
		api.WS_CHILD|api.WS_GROUP|api.WS_TABSTOP|api.WS_VISIBLE,
		api.LVS_EX_FULLROWSELECT,
		api.LVS_REPORT|api.LVS_SHOWSELALWAYS)
}

func (me *ListView) Column(index uint32) *ListViewColumn {
	numCols := me.ColumnCount()
	if index >= numCols {
		panic("Trying to retrieve column with index out of bounds.")
	}
	return newListViewColumn(me, index)
}

func (me *ListView) ColumnCount() uint32 {
	hHeader := api.HWND(me.sendLvmMessage(api.LVM_GETHEADER, 0, 0))
	if hHeader == 0 {
		panic("LVM_GETHEADER failed.")
	}

	count := hHeader.SendMessage(api.WM(api.HDM_GETITEMCOUNT), 0, 0)
	if int32(count) == -1 {
		panic("HDM_GETITEMCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) DeleteAllItems() *ListView {
	ret := me.sendLvmMessage(api.LVM_DELETEALLITEMS, 0, 0)
	if ret == 0 {
		panic("LVM_DELETEALLITEMS failed.")
	}
	return me
}

func (me *ListView) IsGroupViewEnabled() bool {
	return me.sendLvmMessage(api.LVM_ISGROUPVIEWENABLED, 0, 0) >= 0
}

func (me *ListView) Item(index uint32) *ListViewItem {
	numItems := me.ItemCount()
	if index >= numItems {
		panic("Trying to retrieve item with index out of bounds.")
	}
	return newListViewItem(me, index)
}

func (me *ListView) ItemCount() uint32 {
	count := me.sendLvmMessage(api.LVM_GETITEMCOUNT, 0, 0)
	if int32(count) == -1 {
		panic("LVM_GETITEMCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) SelectedItemCount() uint32 {
	count := me.sendLvmMessage(api.LVM_GETSELECTEDCOUNT, 0, 0)
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
	me.hwnd.SendMessage(api.WM_SETREDRAW, api.WPARAM(wp), 0)
	return me
}

func (me *ListView) SetView(view api.LV_VIEW) *ListView {
	if int32(me.sendLvmMessage(api.LVM_SETVIEW, 0, 0)) == -1 {
		panic("LVM_SETVIEW failed.")
	}
	return me
}

func (me *ListView) View() api.LV_VIEW {
	return api.LV_VIEW(me.sendLvmMessage(api.LVM_GETVIEW, 0, 0))
}

func (me *ListView) sendLvmMessage(msg api.LVM,
	wParam api.WPARAM, lParam api.LPARAM) uintptr {

	return me.controlNativeBase.Hwnd().
		SendMessage(api.WM(msg), wParam, lParam) // simple wrapper
}
