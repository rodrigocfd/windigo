//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [list view] control.
//
// [list view]: https://learn.microsoft.com/en-us/windows/win32/controls/list-view-controls-overview
type ListView struct {
	_BaseCtrl
	events       EventsListView
	hContextMenu win.HMENU
	itemsData    map[int]interface{} // data associated with each item; replaces LPARAM approach
	header       *Header
	Cols         CollectionListViewCols  // Methods to interact with the columns collection.
	Items        CollectionListViewItems // Methods to interact with the items collection.
}

// Creates a new [ListView] with [win.CreateWindowEx].
func NewListView(parent Parent, opts *VarOptsListView) *ListView {
	setUniqueCtrlId(&opts.ctrlId)
	me := &ListView{
		_BaseCtrl:    newBaseCtrl(opts.ctrlId),
		events:       EventsListView{opts.ctrlId, &parent.base().userEvents},
		hContextMenu: opts.contextMenu,
		itemsData:    make(map[int]interface{}),
		header:       newHeaderFromListView(parent),
	}
	me.Cols.owner = me
	me.Items.owner = me

	parent.base().beforeUserEvents.WmCreate(func(_ WmCreate) int {
		me.createWindow(opts.wndExStyle, "SysListView32", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, false)
		if opts.ctrlExStyle != co.LVS_EX(0) {
			me.SetExtendedStyle(true, opts.ctrlExStyle)
		}
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		me.assignOrClearHeader()
		return 0 // ignored
	})

	me.defaultMessageHandlers(parent)
	return me
}

// Instantiates a new [ListView] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
func NewListViewDlg(parent Parent, ctrlId uint16, contextMenuId uint16, layout LAY) *ListView {
	hInst, _ := win.GetModuleHandle("")

	var hMenu win.HMENU
	if contextMenuId != 0 {
		var err error
		hMenu, err = hInst.LoadMenu(win.ResIdInt(contextMenuId))
		if err != nil {
			panic(err)
		}
	}

	me := &ListView{
		_BaseCtrl:    newBaseCtrl(ctrlId),
		events:       EventsListView{ctrlId, &parent.base().userEvents},
		hContextMenu: hMenu,
		itemsData:    make(map[int]interface{}),
		header:       newHeaderFromListView(parent),
	}
	me.Cols.owner = me
	me.Items.owner = me

	parent.base().beforeUserEvents.WmInitDialog(func(_ WmInitDialog) bool {
		me.assignDialog(parent)
		parent.base().layout.Add(parent, me.hWnd, layout)
		me.assignOrClearHeader()
		return true // ignored
	})

	me.defaultMessageHandlers(parent)
	return me
}

func (me *ListView) defaultMessageHandlers(parent Parent) {
	me.subclassEvents.WmGetDlgCode(func(p WmGetDlgCode) co.DLGC {
		if !p.IsQuery() && p.VirtualKeyCode() == co.VK_RETURN { // Enter key
			iCode := int32(co.LVN_KEYDOWN)
			nmlvkd := win.NMLVKEYDOWN{
				Hdr: win.NMHDR{
					HWndFrom: me.hWnd,
					IdFrom:   uintptr(me.ctrlId),
					Code:     uint32(iCode),
				},
				WVKey: co.VK_RETURN,
			}
			hParent, _ := me.hWnd.GetAncestor(co.GA_PARENT)
			hParent.SendMessage(co.WM_NOTIFY, win.WPARAM(me.ctrlId), // send Enter key to parent
				win.LPARAM(unsafe.Pointer(&nmlvkd)))
		}
		dlgcSystem := me.hWnd.DefSubclassProc(co.WM_GETDLGCODE, p.Raw.WParam, p.Raw.LParam)
		return co.DLGC(dlgcSystem)
	})

	parent.base().beforeUserEvents.WmNotify(me.ctrlId, co.LVN_KEYDOWN, func(p unsafe.Pointer) uintptr {
		nmk := (*win.NMLVKEYDOWN)(p)
		hasCtrl := (win.GetAsyncKeyState(co.VK_CONTROL) & 0x8000) != 0
		hasShift := (win.GetAsyncKeyState(co.VK_SHIFT) & 0x8000) != 0

		if hasCtrl && nmk.WVKey == 'A' { // Ctrl+A pressed?
			me.Items.SelectAll(true)
		} else if nmk.WVKey == co.VK_APPS { // context menu key
			me.showContextMenu(false, hasCtrl, hasShift)
		}
		return 0 // ignored
	})

	parent.base().beforeUserEvents.WmNotify(me.ctrlId, co.NM_RCLICK, func(p unsafe.Pointer) uintptr {
		nmi := (*win.NMITEMACTIVATE)(p)
		hasCtrl := (nmi.UKeyFlags & co.LVKF_CONTROL) != 0
		hasShift := (nmi.UKeyFlags & co.LVKF_SHIFT) != 0

		me.showContextMenu(true, hasCtrl, hasShift)
		return 0 // ignored
	})

	parent.base().afterUserEvents.WmNotify(me.ctrlId, co.LVN_DELETEITEM, func(p unsafe.Pointer) uintptr {
		nmlv := (*win.NMLISTVIEW)(p)
		item := me.Items.Get(int(nmlv.IItem))
		delete(me.itemsData, item.Uid())
		return 0 // ignored
	})

	parent.base().afterUserEvents.WmDestroy(func() {
		if me.hContextMenu != 0 {
			me.hContextMenu.DestroyMenu()
		}
	})
}

func (me *ListView) showContextMenu(followCursor, hasCtrl, hasShift bool) {
	if me.hContextMenu == 0 { // no menu, nothing to do
		return
	}

	var menuPos win.POINT // menu anchor coords, relative to list view

	if followCursor { // usually when fired by a right-click
		menuPos, _ = win.GetCursorPos()    // relative to screen
		me.hWnd.ScreenToClientPt(&menuPos) // now relative to list view

		if clickedItem, hasClickedItem := me.Items.HitTest(menuPos); !hasClickedItem {
			me.Items.SelectAll(false)
		} else {
			if !hasCtrl && !hasShift {
				clickedItem.Focus()
			}
		}
		me.Focus() // because a right-click won't set the focus by itself

	} else { // usually fired with the context keyboard key
		if focusItem, hasFocused := me.Items.Focused(); hasFocused && focusItem.IsVisible() {
			rcItem := focusItem.ItemRect(co.LVIR_BOUNDS)
			menuPos.X = rcItem.Left + 16 // arbitrary
			menuPos.Y = rcItem.Top + (rcItem.Bottom-rcItem.Top)/2
		} else { // no item is focused and visible
			menuPos.X = 6 // arbitrary anchor coords
			menuPos.Y = 10
		}
	}

	hParent, _ := me.hWnd.GetAncestor(co.GA_PARENT)
	hSubMenu0, _ := me.hContextMenu.GetSubMenu(0)
	hSubMenu0.ShowAtPoint(menuPos, hParent, me.hWnd)
}

func (me *ListView) assignOrClearHeader() {
	hHeader, err := me.hWnd.SendMessage(co.LVM_GETHEADER, 0, 0)
	if hHeader != 0 && err == nil { // the list has a header
		me.header.assignToListView(win.HWND(hHeader))
	} else {
		me.header = nil // no header, free the object
	}
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *ListView) On() *EventsListView {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Returns the context menu associated to this control, if any.
func (me *ListView) ContextMenu() win.HMENU {
	return me.hContextMenu
}

// Returns the embedded [Header] of the list view, or nil if none.
func (me *ListView) Header() *Header {
	return me.header
}

// Retrieves the given image list with [LVM_GETIMAGELIST]. The image lists are
// lazy-initialized: the first time you call this method for a given image list,
// it will be created and assigned with [LVM_SETIMAGELIST].
//
// Since [LVS_SHAREIMAGELISTS] style is not allowed, image lists will be
// automatically destroyed by the OS.
//
// [LVM_GETIMAGELIST]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getimagelist
// [LVM_SETIMAGELIST]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-setimagelist
// [LVS_SHAREIMAGELISTS]: https://learn.microsoft.com/en-us/windows/win32/controls/list-view-window-styles
func (me *ListView) ImageList(which co.LVSIL) win.HIMAGELIST {
	h, _ := me.hWnd.SendMessage(co.LVM_GETIMAGELIST, win.WPARAM(which), 0)
	hImg := win.HIMAGELIST(h)
	if hImg == win.HIMAGELIST(0) {
		cx, cy := 16, 16
		if which == co.LVSIL_NORMAL {
			cx, cy = 32, 32
		}
		hImg, _ = win.ImageListCreate(uint(cx), uint(cy), co.ILC_COLOR32, 1, 1)
		me.hWnd.SendMessage(co.LVM_SETIMAGELIST, win.WPARAM(which), win.LPARAM(hImg))
	}
	return hImg
}

// Scrolls the control with [LVM_SCROLL].
//
// Returns the same object, so further operations can be chained.
//
// Panics on error.
//
// [LVM_SCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-scroll
func (me *ListView) Scroll(horz, vert int) *ListView {
	ret, err := me.hWnd.SendMessage(co.LVM_SCROLL, win.WPARAM(horz), win.LPARAM(vert))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("ListView scrolling failed: %d, %d.", horz, vert))
	}
	return me
}

// Adds or removes extended styles with [LVM_SETEXTENDEDLISTVIEWSTYLE].
//
// Returns the same object, so further operations can be chained.
//
// [LVM_SETEXTENDEDLISTVIEWSTYLE]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-setextendedlistviewstyle
func (me *ListView) SetExtendedStyle(doSet bool, style co.LVS_EX) *ListView {
	affected := co.LVS_EX(0)
	if doSet {
		affected = style
	}
	me.hWnd.SendMessage(co.LVM_SETEXTENDEDLISTVIEWSTYLE,
		win.WPARAM(affected), win.LPARAM(style))
	return me
}

// Enables or disables redrawing with [WM_SETREDRAW].
//
// Use this method to disable redrawing while you're updating multiple items at
// once.
//
// Returns the same object, so further operations can be chained.
//
// [WM_SETREDRAW]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-setredraw
func (me *ListView) SetRedraw(allowRedraw bool) *ListView {
	me.hWnd.SendMessage(co.WM_SETREDRAW,
		win.WPARAM(utl.BoolToUintptr(allowRedraw)), 0)
	return me
}

// Sets the current view with [LVM_SETVIEW].
//
// Returns the same object, so further operations can be chained.
//
// Panics on error.
//
// [LVM_SETVIEW]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-setview
func (me *ListView) SetView(view co.LV_VIEW) *ListView {
	ret, err := me.hWnd.SendMessage(co.LVM_SETVIEW, win.WPARAM(view), 0)
	if err != nil || int32(ret) == -1 {
		panic(fmt.Sprintf("LVM_SETVIEW failed for %d.", view))
	}
	return me
}

// Retrieves the current view with [LVM_GETVIEW].
//
// [LVM_GETVIEW]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getview
func (me *ListView) View() co.LV_VIEW {
	viewRet, _ := me.hWnd.SendMessage(co.LVM_GETVIEW, 0, 0)
	return co.LV_VIEW(viewRet)
}

// Options for [NewListView]; returned by [OptsListView].
type VarOptsListView struct {
	ctrlId      uint16
	layout      LAY
	position    win.POINT
	size        win.SIZE
	ctrlStyle   co.LVS
	ctrlExStyle co.LVS_EX
	wndStyle    co.WS
	wndExStyle  co.WS_EX
	contextMenu win.HMENU
}

// Options for [NewListView].
func OptsListView() *VarOptsListView {
	return &VarOptsListView{
		size:       win.SIZE{Cx: int32(DpiX(120)), Cy: int32(DpiY(120))},
		ctrlStyle:  co.LVS_REPORT | co.LVS_NOSORTHEADER | co.LVS_SHOWSELALWAYS,
		wndStyle:   co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
		wndExStyle: co.WS_EX_LEFT | co.WS_EX_CLIENTEDGE,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsListView) CtrlId(id uint16) *VarOptsListView { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsListView) Layout(l LAY) *VarOptsListView { o.layout = l; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsListView) Position(x, y int) *VarOptsListView {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.Dpi(120, 120).
func (o *VarOptsListView) Size(cx int, cy int) *VarOptsListView {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// List view control [style], passed to [win.CreateWindowEx].
//
// Since the image lists are managed by the control, co.LVS_SHAREIMAGELISTS
// won't be allowed.
//
// Defaults to co.LVS_REPORT | co.LVS_NOSORTHEADER | co.LVS_SHOWSELALWAYS.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/list-view-window-styles
func (o *VarOptsListView) CtrlStyle(s co.LVS) *VarOptsListView {
	o.ctrlStyle = s &^ co.LVS_SHAREIMAGELISTS
	return o
}

// List view control [extended style].
//
// Defaults to co.LVS_EX_NONE.
//
// [extended style]: https://learn.microsoft.com/en-us/windows/win32/controls/extended-list-view-styles
func (o *VarOptsListView) CtrlExStyle(s co.LVS_EX) *VarOptsListView { o.ctrlExStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *VarOptsListView) WndStyle(s co.WS) *VarOptsListView { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT | co.WS_EX_CLIENTEDGE.
func (o *VarOptsListView) WndExStyle(s co.WS_EX) *VarOptsListView { o.wndExStyle = s; return o }

// Context menu popup.
//
// This menu is owned by the list view. It will be automatically destroyed.
//
// Defaults to none.
//
// # Example
//
//	hInst, _ := win.GetModuleHandle("")
//	hMenu, _ := hInst.LoadMenu(win.ResIdInt(0x101))
//
//	ui.ListViewOpts().
//		ContextMenu(hMenu)
func (o *VarOptsListView) ContextMenu(h win.HMENU) *VarOptsListView { o.contextMenu = h; return o }

// Native [list view] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [list view]: https://learn.microsoft.com/en-us/windows/win32/controls/list-view-controls-overview
type EventsListView struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [LVN_BEGINDRAG] message handler.
//
// [LVN_BEGINDRAG]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-begindrag
func (me *EventsListView) LvnBeginDrag(fun func(p *win.NMLISTVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_BEGINDRAG, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLISTVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_BEGINLABELEDIT] message handler.
//
// [LVN_BEGINLABELEDIT]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-beginlabeledit
func (me *EventsListView) LvnBeginLabelEdit(fun func(p *win.NMLVDISPINFO) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_BEGINLABELEDIT, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMLVDISPINFO)(p)))
	})
}

// [LVN_BEGINRDRAG] message handler.
//
// [LVN_BEGINRDRAG]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-beginrdrag
func (me *EventsListView) LvnBeginRDrag(fun func(p *win.NMLISTVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_BEGINRDRAG, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLISTVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_BEGINSCROLL] message handler.
//
// [LVN_BEGINSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-beginscroll
func (me *EventsListView) LvnBeginScroll(fun func(p *win.NMLVSCROLL)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_BEGINSCROLL, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLVSCROLL)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_COLUMNCLICK] message handler.
//
// [LVN_COLUMNCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-columnclick
func (me *EventsListView) LvnColumnClick(fun func(p *win.NMLISTVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_COLUMNCLICK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLISTVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_COLUMNDROPDOWN] message handler.
//
// [LVN_COLUMNDROPDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-columndropdown
func (me *EventsListView) LvnColumnDropDown(fun func(p *win.NMLISTVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_COLUMNDROPDOWN, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLISTVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_COLUMNOVERFLOWCLICK] message handler.
//
// [LVN_COLUMNOVERFLOWCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-columnoverflowclick
func (me *EventsListView) LvnColumnOverflowClick(fun func(p *win.NMLISTVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_COLUMNOVERFLOWCLICK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLISTVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_DELETEALLITEMS] message handler.
//
// [LVN_DELETEALLITEMS]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-deleteallitems
func (me *EventsListView) LvnDeleteAllItems(fun func(p *win.NMLISTVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_DELETEALLITEMS, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLISTVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_DELETEITEM] message handler.
//
// [LVN_DELETEITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-deleteitem
func (me *EventsListView) LvnDeleteItem(fun func(p *win.NMLISTVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_DELETEITEM, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLISTVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_ENDLABELEDIT] message handler.
//
// [LVN_ENDLABELEDIT]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-endlabeledit
func (me *EventsListView) LvnEndLabelEdit(fun func(p *win.NMLVDISPINFO) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_ENDLABELEDIT, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMLVDISPINFO)(p)))
	})
}

// [LVN_ENDSCROLL] message handler.
//
// [LVN_ENDSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-endscroll
func (me *EventsListView) LvnEndScroll(fun func(p *win.NMLVSCROLL)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_ENDSCROLL, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLVSCROLL)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_GETDISPINFO] message handler.
//
// [LVN_GETDISPINFO]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-getdispinfo
func (me *EventsListView) LvnGetDispInfo(fun func(p *win.NMLVDISPINFO)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_GETDISPINFO, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLVDISPINFO)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_GETEMPTYMARKUP] message handler.
//
// [LVN_GETEMPTYMARKUP]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-getemptymarkup
func (me *EventsListView) LvnGetEmptyMarkup(fun func(p *win.NMLVEMPTYMARKUP) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_GETEMPTYMARKUP, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMLVEMPTYMARKUP)(p)))
	})
}

// [LVN_GETINFOTIP] message handler.
//
// [LVN_GETINFOTIP]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-getinfotip
func (me *EventsListView) LvnGetInfoTip(fun func(p *win.NMLVGETINFOTIP)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_GETINFOTIP, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLVGETINFOTIP)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_HOTTRACK] message handler.
//
// [LVN_HOTTRACK]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-hottrack
func (me *EventsListView) LvnHotTrack(fun func(p *win.NMLISTVIEW) int) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_HOTTRACK, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMLISTVIEW)(p)))
	})
}

// [LVN_INCREMENTALSEARCH] message handler.
//
// [LVN_INCREMENTALSEARCH]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-incrementalsearch
func (me *EventsListView) LvnIncrementalSearch(fun func(p *win.NMLVFINDITEM) int) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_INCREMENTALSEARCH, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMLVFINDITEM)(p)))
	})
}

// [LVN_INSERTITEM] message handler.
//
// [LVN_INSERTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-insertitem
func (me *EventsListView) LvnInsertItem(fun func(p *win.NMLISTVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_INSERTITEM, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLISTVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_ITEMACTIVATE] message handler.
//
// [LVN_ITEMACTIVATE]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-itemactivate
func (me *EventsListView) LvnItemActivate(fun func(p *win.NMITEMACTIVATE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_ITEMACTIVATE, func(p unsafe.Pointer) uintptr {
		fun((*win.NMITEMACTIVATE)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_ITEMCHANGED] message handler.
//
// [LVN_ITEMCHANGED]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-itemchanged
func (me *EventsListView) LvnItemChanged(fun func(p *win.NMLISTVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_ITEMCHANGED, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLISTVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_ITEMCHANGING] message handler.
//
// [LVN_ITEMCHANGING]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-itemchanging
func (me *EventsListView) LvnItemChanging(fun func(p *win.NMLISTVIEW) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_ITEMCHANGING, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMLISTVIEW)(p)))
	})
}

// [LVN_KEYDOWN] message handler.
//
// [LVN_KEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-keydown
func (me *EventsListView) LvnKeyDown(fun func(p *win.NMLVKEYDOWN)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_KEYDOWN, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLVKEYDOWN)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_LINKCLICK] message handler.
//
// [LVN_LINKCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-linkclick
func (me *EventsListView) LvnLinkClick(fun func(p *win.NMLVLINK)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_LINKCLICK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLVLINK)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_MARQUEEBEGIN] message handler.
//
// [LVN_MARQUEEBEGIN]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-marqueebegin
func (me *EventsListView) LvnMarqueeBegin(fun func() uint) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_MARQUEEBEGIN, func(p unsafe.Pointer) uintptr {
		return uintptr(fun())
	})
}

// [LVN_ODCACHEHINT] message handler.
//
// [LVN_ODCACHEHINT]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-odcachehint
func (me *EventsListView) LvnODCacheHint(fun func(p *win.NMLVCACHEHINT)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_ODCACHEHINT, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLVCACHEHINT)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_ODFINDITEM] message handler.
//
// [LVN_ODFINDITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-odfinditem
func (me *EventsListView) LvnODFindItem(fun func(p *win.NMLVFINDITEM) int) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_ODFINDITEM, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMLVFINDITEM)(p)))
	})
}

// [LVN_ODSTATECHANGED] message handler.
//
// [LVN_ODSTATECHANGED]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-odstatechanged
func (me *EventsListView) LvnODStateChanged(fun func(p *win.NMLVODSTATECHANGE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_ODSTATECHANGED, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLVODSTATECHANGE)(p))
		return me.parentEvents.defProcVal
	})
}

// [LVN_SETDISPINFO] message handler.
//
// [LVN_SETDISPINFO]: https://learn.microsoft.com/en-us/windows/win32/controls/lvn-setdispinfo
func (me *EventsListView) LvnSetDispInfo(fun func(p *win.NMLVDISPINFO)) {
	me.parentEvents.WmNotify(me.ctrlId, co.LVN_SETDISPINFO, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLVDISPINFO)(p))
		return me.parentEvents.defProcVal
	})
}

// [NM_CLICK] message handler.
//
// [NM_CLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-click-list-view
func (me *EventsListView) NmClick(fun func(p *win.NMITEMACTIVATE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CLICK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMITEMACTIVATE)(p))
		return me.parentEvents.defProcVal
	})
}

// [NM_CUSTOMDRAW] message handler.
//
// [NM_CUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-customdraw-list-view
func (me *EventsListView) NmCustomDraw(fun func(p *win.NMLVCUSTOMDRAW) co.CDRF) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMLVCUSTOMDRAW)(p)))
	})
}

// [NM_DBLCLK] message handler.
//
// [NM_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-dblclk-list-view
func (me *EventsListView) NmDblClk(fun func(p *win.NMITEMACTIVATE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_DBLCLK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMITEMACTIVATE)(p))
		return me.parentEvents.defProcVal
	})
}

// [NM_HOVER] message handler.
//
// [NM_HOVER]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-hover-list-view
func (me *EventsListView) NmHover(fun func() uint) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_HOVER, func(_ unsafe.Pointer) uintptr {
		return uintptr(fun())
	})
}

// [NM_KILLFOCUS] message handler.
//
// [NM_KILLFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-killfocus-list-view
func (me *EventsListView) NmKillFocus(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_KILLFOCUS, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [NM_RCLICK] message handler.
//
// [NM_RCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rclick-list-view
func (me *EventsListView) NmRClick(fun func(p *win.NMITEMACTIVATE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RCLICK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMITEMACTIVATE)(p))
		return me.parentEvents.defProcVal
	})
}

// [NM_RDBLCLK] message handler.
//
// [NM_RDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-list-view
func (me *EventsListView) NmRDblClk(fun func(p *win.NMITEMACTIVATE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RDBLCLK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMITEMACTIVATE)(p))
		return me.parentEvents.defProcVal
	})
}

// [NM_RELEASEDCAPTURE] message handler.
//
// [NM_RELEASEDCAPTURE]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-list-view-
func (me *EventsListView) NmReleasedCapture(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [NM_RETURN] message handler.
//
// [NM_RETURN]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-return-list-view-
func (me *EventsListView) NmReturn(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RETURN, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [NM_SETFOCUS] message handler.
//
// [NM_SETFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-setfocus-list-view-
func (me *EventsListView) NmSetFocus(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_SETFOCUS, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}
