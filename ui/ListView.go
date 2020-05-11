package ui

import (
	"fmt"
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

// Native list view control.
type ListView struct {
	nativeControlBase
}

// Optional; returns a ListView with a specific control ID.
func MakeListView(ctrlId c.ID) ListView {
	return ListView{
		nativeControlBase: makeNativeControlBase(ctrlId),
	}
}

func (me *ListView) AddColumn(text string, width uint32) *ListViewColumn {
	lvc := api.LVCOLUMN{
		Mask:    c.LVCF_TEXT | c.LVCF_WIDTH,
		PszText: api.StrToUtf16Ptr(text),
		Cx:      int32(width),
	}
	newIdx := me.sendLvmMessage(c.LVM_INSERTCOLUMN, 0xFFFF,
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
		Mask:    c.LVIF_TEXT,
		PszText: api.StrToUtf16Ptr(text),
		IItem:   0x0FFFFFFF, // insert as the last one
	}
	newIdx := me.sendLvmMessage(c.LVM_INSERTITEM, 0,
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

func (me *ListView) Create(parent Window, x, y int32, width, height uint32,
	exStyles c.WS_EX, styles c.WS,
	lvExStyles c.LVS_EX, lvStyles c.LVS) *ListView {

	me.nativeControlBase.create(exStyles|c.WS_EX(lvExStyles),
		"SysListView32", "", styles|c.WS(lvStyles),
		x, y, width, height, parent)
	return me
}

func (me *ListView) CreateReport(parent Window, x, y int32,
	width, height uint32) *ListView {

	return me.Create(parent, x, y, width, height,
		c.WS_EX_CLIENTEDGE,
		c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.LVS_EX_FULLROWSELECT,
		c.LVS_REPORT|c.LVS_SHOWSELALWAYS)
}

func (me *ListView) Column(index uint32) *ListViewColumn {
	numCols := me.ColumnCount()
	if index >= numCols {
		panic("Trying to retrieve column with index out of bounds.")
	}
	return newListViewColumn(me, index)
}

func (me *ListView) ColumnCount() uint32 {
	hHeader := api.HWND(me.sendLvmMessage(c.LVM_GETHEADER, 0, 0))
	if hHeader == 0 {
		panic("LVM_GETHEADER failed.")
	}

	count := hHeader.SendMessage(c.WM(c.HDM_GETITEMCOUNT), 0, 0)
	if int32(count) == -1 {
		panic("HDM_GETITEMCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) DeleteAllItems() *ListView {
	ret := me.sendLvmMessage(c.LVM_DELETEALLITEMS, 0, 0)
	if ret == 0 {
		panic("LVM_DELETEALLITEMS failed.")
	}
	return me
}

func (me *ListView) Enable(enabled bool) *ListView {
	me.hwnd.EnableWindow(enabled)
	return me
}

func (me *ListView) IsEnabled() bool {
	return me.hwnd.IsWindowEnabled()
}

func (me *ListView) IsGroupViewEnabled() bool {
	return me.sendLvmMessage(c.LVM_ISGROUPVIEWENABLED, 0, 0) >= 0
}

func (me *ListView) Item(index uint32) *ListViewItem {
	numItems := me.ItemCount()
	if index >= numItems {
		panic("Trying to retrieve item with index out of bounds.")
	}
	return newListViewItem(me, index)
}

func (me *ListView) ItemCount() uint32 {
	count := me.sendLvmMessage(c.LVM_GETITEMCOUNT, 0, 0)
	if int32(count) == -1 {
		panic("LVM_GETITEMCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) SelectedItemCount() uint32 {
	count := me.sendLvmMessage(c.LVM_GETSELECTEDCOUNT, 0, 0)
	if int32(count) == -1 {
		panic("LVM_GETSELECTEDCOUNT failed.")
	}
	return uint32(count)
}

func (me *ListView) SetFocus() api.HWND {
	return me.hwnd.SetFocus()
}

func (me *ListView) SetRedraw(allowRedraw bool) *ListView {
	wp := 0
	if allowRedraw {
		wp = 1
	}
	me.hwnd.SendMessage(c.WM_SETREDRAW, api.WPARAM(wp), 0)
	return me
}

func (me *ListView) SetView(view c.LV_VIEW) *ListView {
	if int32(me.sendLvmMessage(c.LVM_SETVIEW, 0, 0)) == -1 {
		panic("LVM_SETVIEW failed.")
	}
	return me
}

func (me *ListView) View() c.LV_VIEW {
	return c.LV_VIEW(me.sendLvmMessage(c.LVM_GETVIEW, 0, 0))
}

func (me *ListView) sendLvmMessage(msg c.LVM,
	wParam api.WPARAM, lParam api.LPARAM) uintptr {

	return me.nativeControlBase.Hwnd().
		SendMessage(c.WM(msg), wParam, lParam) // simple wrapper
}
