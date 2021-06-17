package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// A single item of a TreeView.
type TreeViewItem interface {
	AddChild(text string) TreeViewItem
	Children() []TreeViewItem
	Delete()
	EnsureVisible()
	Expand(doExpand bool)
	Htreeitem() win.HTREEITEM
	IsExpanded() bool
	IsRoot() bool
	LParam() win.LPARAM
	NextSibling() (TreeViewItem, bool)
	Parent() (TreeViewItem, bool)
	PrevSibling() (TreeViewItem, bool)
	SetLParam(lp win.LPARAM)
	SetText(text string)
	Text() string
}

//------------------------------------------------------------------------------

type _TreeViewItem struct {
	pHwnd *win.HWND
	hItem win.HTREEITEM
}

func (me *_TreeViewItem) new(pHwnd *win.HWND, hItem win.HTREEITEM) {
	me.pHwnd = pHwnd
	me.hItem = hItem
}

func (me *_TreeViewItem) AddChild(text string) TreeViewItem {
	textBuf := win.Str.ToUint16Slice(text)
	tvi := win.TVINSERTSTRUCT{
		HParent:      me.hItem,
		HInsertAfter: win.HTREEITEM(co.HTREEITEM_LAST),
		Itemex: win.TVITEMEX{
			Mask:    co.TVIF_TEXT,
			PszText: &textBuf[0],
		},
	}

	hNewItem := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_INSERTITEM,
			0, win.LPARAM(unsafe.Pointer(&tvi))),
	)
	if hNewItem == 0 {
		panic(fmt.Sprintf("TVM_INSERTITEM failed \"%s\".", text))
	}

	return &_TreeViewItem{
		pHwnd: me.pHwnd,
		hItem: win.HTREEITEM(hNewItem),
	}
}

func (me *_TreeViewItem) Children() []TreeViewItem {
	hChildren := make([]TreeViewItem, 0)
	hItem := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_CHILD), win.LPARAM(me.hItem)), // retrieve first child
	)
	hasSibling := hItem != 0 // has 1st child?

	for hasSibling {
		hChildren = append(hChildren, &_TreeViewItem{
			pHwnd: me.pHwnd,
			hItem: hItem,
		})

		hItem = win.HTREEITEM(
			me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
				win.WPARAM(co.TVGN_NEXT), win.LPARAM(me.hItem)),
		)
		hasSibling = hItem != 0
	}

	return hChildren
}

func (me *_TreeViewItem) Delete() {
	if me.pHwnd.SendMessage(co.TVM_DELETEITEM, 0, win.LPARAM(me.hItem)) == 0 {
		panic("TVM_DELETEITEM failed.")
	}
}

func (me *_TreeViewItem) EnsureVisible() {
	me.pHwnd.SendMessage(co.TVM_ENSUREVISIBLE, 0, win.LPARAM(me.hItem))
}

func (me *_TreeViewItem) Expand(doExpand bool) {
	flag := co.TVE_EXPAND
	if !doExpand {
		flag = co.TVE_COLLAPSE
	}
	me.pHwnd.SendMessage(co.TVM_EXPAND, win.WPARAM(flag), win.LPARAM(me.hItem))
}

func (me *_TreeViewItem) Htreeitem() win.HTREEITEM {
	return me.hItem
}

func (me *_TreeViewItem) IsExpanded() bool {
	return (co.TVIS(
		me.pHwnd.SendMessage(co.TVM_GETITEMSTATE,
			win.WPARAM(me.hItem), win.LPARAM(co.TVIS_EXPANDED)),
	) & co.TVIS_EXPANDED) != 0
}

func (me *_TreeViewItem) IsRoot() bool {
	_, hasParent := me.Parent()
	return !hasParent
}

func (me *_TreeViewItem) LParam() win.LPARAM {
	tvi := win.TVITEMEX{
		HItem: me.hItem,
		Mask:  co.TVIF_PARAM,
	}

	ret := me.pHwnd.SendMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return tvi.LParam
}

func (me *_TreeViewItem) NextSibling() (TreeViewItem, bool) {
	hSibling := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_NEXT), win.LPARAM(me.hItem)),
	)

	if hSibling == 0 {
		return nil, false
	}
	return &_TreeViewItem{
		pHwnd: me.pHwnd,
		hItem: hSibling,
	}, true
}

func (me *_TreeViewItem) Parent() (TreeViewItem, bool) {
	hParent := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_PARENT), win.LPARAM(me.hItem)),
	)

	if hParent == 0 {
		return nil, false
	}
	return &_TreeViewItem{
		pHwnd: me.pHwnd,
		hItem: hParent,
	}, true
}

func (me *_TreeViewItem) PrevSibling() (TreeViewItem, bool) {
	hSibling := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_PREVIOUS), win.LPARAM(me.hItem)),
	)

	if hSibling == 0 {
		return nil, false
	}
	return &_TreeViewItem{
		pHwnd: me.pHwnd,
		hItem: hSibling,
	}, true
}

func (me *_TreeViewItem) SetLParam(lp win.LPARAM) {
	tvi := win.TVITEMEX{
		HItem:  me.hItem,
		Mask:   co.TVIF_PARAM,
		LParam: lp,
	}

	ret := me.pHwnd.SendMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_SETITEM failed.")
	}
}

func (me *_TreeViewItem) SetText(text string) {
	textBuf := win.Str.ToUint16Slice(text)
	tvi := win.TVITEMEX{
		HItem:   me.hItem,
		Mask:    co.TVIF_TEXT,
		PszText: &textBuf[0],
	}

	ret := me.pHwnd.SendMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_SETITEM failed \"%s\".", text))
	}
}

func (me *_TreeViewItem) Text() string {
	buf := [256]uint16{} // arbitrary
	tvi := win.TVITEMEX{
		HItem:      me.hItem,
		Mask:       co.TVIF_TEXT,
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}

	ret := me.pHwnd.SendMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return win.Str.FromUint16Slice(buf[:])
}
