package parm

import (
	"gowinui/api"
	"unsafe"
)

type LvnDeleteAllItems WmNotify

func (p LvnDeleteAllItems) NmListView() *api.NMLISTVIEW {
	return (*api.NMLISTVIEW)(unsafe.Pointer(p.LParam))
}

type LvnDeleteItem struct{ LvnDeleteAllItems } // inherit
type LvnInsertItem struct{ LvnDeleteAllItems }
type LvnItemChanged struct{ LvnDeleteAllItems }

type LvnItemActivate WmNotify

func (p LvnItemActivate) NmItemActivate() *api.NMITEMACTIVATE {
	return (*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam))
}

type LvnKeyDown WmNotify

func (p LvnKeyDown) NmLvKeyDown() *api.NMLVKEYDOWN {
	return (*api.NMLVKEYDOWN)(unsafe.Pointer(p.LParam))
}
