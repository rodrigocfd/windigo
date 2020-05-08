package parm

import (
	"gowinui/api"
	"unsafe"
)

type LvnDeleteAllItems WmNotify

func (p LvnDeleteAllItems) NmListView() *api.NMLISTVIEW {
	return (*api.NMLISTVIEW)(unsafe.Pointer(p.LParam))
}

type LvnDeleteItem WmNotify

func (p LvnDeleteItem) NmListView() *api.NMLISTVIEW {
	return (*api.NMLISTVIEW)(unsafe.Pointer(p.LParam))
}

type LvnInsertItem WmNotify

func (p LvnInsertItem) NmListView() *api.NMLISTVIEW {
	return (*api.NMLISTVIEW)(unsafe.Pointer(p.LParam))
}

type LvnItemChanged WmNotify

func (p LvnItemChanged) NmListView() *api.NMLISTVIEW {
	return (*api.NMLISTVIEW)(unsafe.Pointer(p.LParam))
}
