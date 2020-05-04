package parm

import (
	"unsafe"
	"winffi/api"
)

type LvnDeleteAllItems WmNotify

func (p *LvnDeleteAllItems) NmListView() *api.NMLISTVIEW {
	return (*api.NMLISTVIEW)(unsafe.Pointer(p.WParam))
}

type LvnItemChanged WmNotify

func (p *LvnItemChanged) NmListView() *api.NMLISTVIEW {
	return (*api.NMLISTVIEW)(unsafe.Pointer(p.WParam))
}
