package ui

import (
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

func (lv *ListView) Create(parent Window, x, y int32, width, height uint32,
	exStyles c.WS_EX, styles c.WS,
	lvExStyles c.LVS_EX, lvStyles c.LVS) *ListView {

	if lv.hwnd != 0 {
		panic("Trying to create ListView twice.")
	}
	lv.hwnd = api.CreateWindowEx(exStyles|c.WS_EX(lvExStyles),
		"SysListView32", "", styles|c.WS(lvStyles),
		x, y, width, height, parent.Hwnd(), api.HMENU(lv.id),
		parent.Hwnd().GetInstance(), nil)
	return lv
}

func (lv *ListView) CreateReport(parent Window, x, y int32,
	width, height uint32) *ListView {

	return lv.Create(parent, x, y, width, height,
		c.WS_EX_CLIENTEDGE,
		c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.LVS_EX_FULLROWSELECT,
		c.LVS_REPORT|c.LVS_SHOWSELALWAYS)
}

func (lv *ListView) CtrlId() c.ID {
	return lv.id
}

func (lv *ListView) Enable(enabled bool) *ListView {
	lv.hwnd.EnableWindow(enabled)
	return lv
}

func (lv *ListView) Hwnd() api.HWND {
	return lv.hwnd
}

func (lv *ListView) IsEnabled() bool {
	return lv.hwnd.IsWindowEnabled()
}

func (lv *ListView) SetFocus() api.HWND {
	return lv.hwnd.SetFocus()
}
