package ui

import (
	"fmt"
	"unsafe"
	"winffi/api"
	c "winffi/consts"
)

type ListView struct {
	hwnd api.HWND
	id   c.ID
}

func MakeListView() ListView {
	return MakeListViewWithId(NextAutoCtrlId())
}

func MakeListViewWithId(ctrlId c.ID) ListView {
	return ListView{
		hwnd: api.HWND(0),
		id:   ctrlId,
	}
}

func (me *ListView) AddColumn(text string, width uint32) ListViewColumn {
	lvc := api.LVCOLUMN{
		Mask:    c.LVCF_TEXT | c.LVCF_WIDTH,
		PszText: api.StrToUtf16Ptr(text),
		Cx:      int32(width),
	}
	newIdx := me.hwnd.SendMessage(c.WM(c.LVM_INSERTCOLUMN), 0xFFFF,
		api.LPARAM(unsafe.Pointer(&lvc)))
	if int32(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTCOLUMN failed \"%s\".", text))
	}
	return MakeListViewColumn(me, uint32(newIdx))
}

func (me *ListView) Create(parent Window, x, y int32, width, height uint32,
	exStyles c.WS_EX, styles c.WS,
	lvExStyles c.LVS_EX, lvStyles c.LVS) *ListView {

	if me.hwnd != 0 {
		panic("Trying to create ListView twice.")
	}
	me.hwnd = api.CreateWindowEx(exStyles|c.WS_EX(lvExStyles),
		"SysListView32", "", styles|c.WS(lvStyles),
		x, y, width, height, parent.Hwnd(), api.HMENU(me.id),
		parent.Hwnd().GetInstance(), nil)
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

func (me *ListView) CtrlId() c.ID {
	return me.id
}

func (me *ListView) Enable(enabled bool) *ListView {
	me.hwnd.EnableWindow(enabled)
	return me
}

func (me *ListView) Hwnd() api.HWND {
	return me.hwnd
}

func (me *ListView) IsEnabled() bool {
	return me.hwnd.IsWindowEnabled()
}

func (me *ListView) SetFocus() api.HWND {
	return me.hwnd.SetFocus()
}
