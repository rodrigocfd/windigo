package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native list view control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/list-view-controls-overview
type ListView interface {
	AnyNativeControl

	// Exposes all the ListView notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-list-view-control-reference-notifications
	On() *_ListViewEvents

	ContextMenu() win.HMENU     // Returns the associated context menu, if any.
	Columns() *_ListViewColumns // Column methods.
	Items() *_ListViewItems     // Item methods.
	SetRedraw(allowRedraw bool) // Sends WM_SETREDRAW to enable or disable UI updates.
}

//------------------------------------------------------------------------------

type _ListView struct {
	_NativeControlBase
	events       _ListViewEvents
	columns      _ListViewColumns
	items        _ListViewItems
	hContextMenu win.HMENU
}

// Creates a new ListView specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewListViewRaw(parent AnyParent, opts ListViewRawOpts) ListView {
	opts.fillBlankValuesWithDefault()

	me := _ListView{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.events.new(&me._NativeControlBase)
	me.columns.new(&me._NativeControlBase)
	me.items.new(&me._NativeControlBase)
	me.hContextMenu = opts.ContextMenu

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.Position, &opts.Size)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"SysListView32", "", opts.Styles|co.WS(opts.ListViewStyles),
			opts.Position, opts.Size, win.HMENU(opts.CtrlId))

		if opts.ListViewExStyles != co.LVS_EX_NONE {
			me.Hwnd().SendMessage(co.LVM_SETEXTENDEDLISTVIEWSTYLE,
				win.WPARAM(opts.ListViewExStyles),
				win.LPARAM(opts.ListViewExStyles))
		}
	})

	me.handledEvents()
	return &me
}

// Creates a new ListView from a dialog resource.
func NewListViewDlg(parent AnyParent, ctrlId, contextMenuId int) ListView {
	hContextMenu := win.HMENU(0)
	if contextMenuId != 0 {
		hContextMenu = win.HINSTANCE(0).LoadMenu(int32(contextMenuId))
	}

	me := _ListView{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)
	me.columns.new(&me._NativeControlBase)
	me.items.new(&me._NativeControlBase)
	me.hContextMenu = hContextMenu

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	me.handledEvents()
	return &me
}

func (me *_ListView) On() *_ListViewEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the ListView is created.")
	}
	return &me.events
}

func (me *_ListView) ContextMenu() win.HMENU {
	return me.hContextMenu
}

func (me *_ListView) Columns() *_ListViewColumns {
	return &me.columns
}

func (me *_ListView) Items() *_ListViewItems {
	return &me.items
}

func (me *_ListView) SetRedraw(allowRedraw bool) {
	me.Hwnd().SendMessage(co.WM_SETREDRAW,
		win.WPARAM(util.BoolToUintptr(allowRedraw)), 0)
}

func (me *_ListView) handledEvents() {
	me.Parent().internalOn().addNfyZero(me.CtrlId(), co.LVN_KEYDOWN, func(p unsafe.Pointer) {
		nmk := (*win.NMLVKEYDOWN)(p)
		hasCtrl := (win.GetAsyncKeyState(co.VK_CONTROL) & 0x8000) != 0
		hasShift := (win.GetAsyncKeyState(co.VK_SHIFT) & 0x8000) != 0

		if hasCtrl && nmk.WVKey == 'A' { // Ctrl+A pressed?
			me.Items().SetSelectedAll(true)
		} else if nmk.WVKey == co.VK_APPS { // context meny key
			me.showContextMenu(false, hasCtrl, hasShift)
		}
	})

	me.Parent().internalOn().addNfyZero(me.CtrlId(), co.NM_RCLICK, func(p unsafe.Pointer) {
		nmi := (*win.NMITEMACTIVATE)(p)
		hasCtrl := (nmi.UKeyFlags & co.LVKF_CONTROL) != 0
		hasShift := (nmi.UKeyFlags & co.LVKF_SHIFT) != 0

		me.showContextMenu(true, hasCtrl, hasShift)
	})

	me.Parent().internalOn().addMsgZero(co.WM_NCDESTROY, func(_ wm.Any) {
		if me.hContextMenu != 0 {
			me.hContextMenu.DestroyMenu()
		}
	})
}

func (me *_ListView) showContextMenu(followCursor, hasCtrl, hasShift bool) {
	if me.hContextMenu == 0 { // no menu, nothing to do
		return
	}

	var menuPos win.POINT // menu anchor coords, relative to list view

	if followCursor { // usually when fired by a right-click
		menuPos = win.GetCursorPos()         // relative to screen
		me.Hwnd().ScreenToClientPt(&menuPos) // now relative to list view
		lvhti := me.Items().HitTest(menuPos) // to find item below cursor, if any

		if lvhti.IItem != -1 { // an item was right-clicked
			if !hasCtrl && !hasShift {
				clickedIdx := int(lvhti.IItem)
				if !me.Items().IsSelected(clickedIdx) {
					me.Items().SetSelectedAll(false)
					me.Items().SetSelected(true, clickedIdx)
				}
				me.Items().SetFocused(clickedIdx)
			}
		} else if !hasCtrl && !hasShift { // no item was right-clicked
			me.Items().SetSelectedAll(false)
		}
		me.Hwnd().SetFocus() // because a right-click won't set the focus by itself

	} else { // usually fired with the context keyboard key
		if focusedIdx, hasFocused := me.Items().Focused(); hasFocused && me.Items().IsVisible(focusedIdx) {
			rcItem := me.Items().Rect(focusedIdx, co.LVIR_BOUNDS)
			menuPos.X = rcItem.Left + 16 // arbitrary
			menuPos.Y = rcItem.Top + (rcItem.Bottom-rcItem.Top)/2
		} else { // no item is focused and visible
			menuPos.X = 6 // arbitrary
			menuPos.Y = 10
		}
	}

	me.hContextMenu.ShowAtPoint(menuPos, me.Hwnd().GetParent(), me.Hwnd())
}

//------------------------------------------------------------------------------

// Options for NewListViewRaw().
type ListViewRawOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// Control size in pixels.
	// Defaults to 120x120. Will be adjusted to the current system DPI.
	Size win.SIZE
	// ListView control styles, passed to CreateWindowEx().
	// Defaults to LVS_REPORT | LVS_NOSORTHEADER | LVS_SHOWSELALWAYS | LVS_SHAREIMAGELISTS.
	ListViewStyles co.LVS
	// ListView extended control styles, passed to CreateWindowEx().
	// Defaults to LVS_EX_NONE.
	ListViewExStyles co.LVS_EX
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_CLIENTEDGE.
	ExStyles co.WS_EX

	// Context menu for the list view. Will be automatically destroyed.
	// Defaults to none.
	ContextMenu win.HMENU
}

func (opts *ListViewRawOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.Size.Cx == 0 {
		opts.Size.Cx = 120
	}
	if opts.Size.Cy == 0 {
		opts.Size.Cy = 120
	}

	if opts.ListViewStyles == 0 {
		opts.ListViewStyles = co.LVS_REPORT | co.LVS_NOSORTHEADER |
			co.LVS_SHOWSELALWAYS | co.LVS_SHAREIMAGELISTS
	}
	if opts.ListViewExStyles == 0 {
		opts.ListViewExStyles = co.LVS_EX_NONE
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_CLIENTEDGE
	}
}

//------------------------------------------------------------------------------

// ListView control notifications.
type _ListViewEvents struct {
	ctrlId int
	events *_EventsNfy
}

func (me *_ListViewEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.parent.On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-begindrag
func (me *_ListViewEvents) LvnBeginDrag(userFunc func(p *win.NMLISTVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_BEGINDRAG, func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-beginlabeledit
func (me *_ListViewEvents) LvnBeginLabelEdit(userFunc func(p *win.NMLVDISPINFO) bool) {
	me.events.addNfyRet(me.ctrlId, co.LVN_BEGINLABELEDIT, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMLVDISPINFO)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-beginrdrag
func (me *_ListViewEvents) LvnBeginRDrag(userFunc func(p *win.NMLISTVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_BEGINRDRAG, func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-beginscroll
func (me *_ListViewEvents) LvnBeginScroll(userFunc func(p *win.NMLVSCROLL)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_BEGINSCROLL, func(p unsafe.Pointer) {
		userFunc((*win.NMLVSCROLL)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-columnclick
func (me *_ListViewEvents) LvnColumnClick(userFunc func(p *win.NMLISTVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_COLUMNCLICK, func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-columndropdown
func (me *_ListViewEvents) LvnColumnDropDown(userFunc func(p *win.NMLISTVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_COLUMNDROPDOWN, func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-columnoverflowclick
func (me *_ListViewEvents) LvnColumnOverflowClick(userFunc func(p *win.NMLISTVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_COLUMNOVERFLOWCLICK, func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-deleteallitems
func (me *_ListViewEvents) LvnDeleteAllItems(userFunc func(p *win.NMLISTVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_DELETEALLITEMS, func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-deleteitem
func (me *_ListViewEvents) LvnDeleteItem(userFunc func(p *win.NMLISTVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_DELETEITEM, func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-endlabeledit
func (me *_ListViewEvents) LvnEndLabelEdit(userFunc func(p *win.NMLVDISPINFO) bool) {
	me.events.addNfyRet(me.ctrlId, co.LVN_ENDLABELEDIT, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMLVDISPINFO)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-endscroll
func (me *_ListViewEvents) LvnEndScroll(userFunc func(p *win.NMLVSCROLL)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_ENDSCROLL, func(p unsafe.Pointer) {
		userFunc((*win.NMLVSCROLL)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-getdispinfo
func (me *_ListViewEvents) LvnGetDispInfo(userFunc func(p *win.NMLVDISPINFO)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_GETDISPINFO, func(p unsafe.Pointer) {
		userFunc((*win.NMLVDISPINFO)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-getemptymarkup
func (me *_ListViewEvents) LvnGetEmptyMarkup(userFunc func(p *win.NMLVEMPTYMARKUP) bool) {
	me.events.addNfyRet(me.ctrlId, co.LVN_GETEMPTYMARKUP, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMLVEMPTYMARKUP)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-getinfotip
func (me *_ListViewEvents) LvnGetInfoTip(userFunc func(p *win.NMLVGETINFOTIP)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_GETINFOTIP, func(p unsafe.Pointer) {
		userFunc((*win.NMLVGETINFOTIP)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-hottrack
func (me *_ListViewEvents) LvnHotTrack(userFunc func(p *win.NMLISTVIEW) int) {
	me.events.addNfyRet(me.ctrlId, co.LVN_HOTTRACK, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMLISTVIEW)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-incrementalsearch
func (me *_ListViewEvents) LvnIncrementalSearch(userFunc func(p *win.NMLVFINDITEM) int) {
	me.events.addNfyRet(me.ctrlId, co.LVN_INCREMENTALSEARCH, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMLVFINDITEM)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-insertitem
func (me *_ListViewEvents) LvnInsertItem(userFunc func(p *win.NMLISTVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_INSERTITEM, func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-itemactivate
func (me *_ListViewEvents) LvnItemActivate(userFunc func(p *win.NMITEMACTIVATE)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_ITEMACTIVATE, func(p unsafe.Pointer) {
		userFunc((*win.NMITEMACTIVATE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-itemchanged
func (me *_ListViewEvents) LvnItemChanged(userFunc func(p *win.NMLISTVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_ITEMCHANGED, func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-itemchanging
func (me *_ListViewEvents) LvnItemChanging(userFunc func(p *win.NMLISTVIEW) bool) {
	me.events.addNfyRet(me.ctrlId, co.LVN_ITEMCHANGING, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMLISTVIEW)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-keydown
func (me *_ListViewEvents) LvnKeyDown(userFunc func(p *win.NMLVKEYDOWN)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_KEYDOWN, func(p unsafe.Pointer) {
		userFunc((*win.NMLVKEYDOWN)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-linkclick
func (me *_ListViewEvents) LvnLinkClick(userFunc func(p *win.NMLVLINK)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_LINKCLICK, func(p unsafe.Pointer) {
		userFunc((*win.NMLVLINK)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-marqueebegin
func (me *_ListViewEvents) LvnMarqueeBegin(userFunc func() uint) {
	me.events.addNfyRet(me.ctrlId, co.LVN_MARQUEEBEGIN, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-odcachehint
func (me *_ListViewEvents) LvnODCacheHint(userFunc func(p *win.NMLVCACHEHINT)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_ODCACHEHINT, func(p unsafe.Pointer) {
		userFunc((*win.NMLVCACHEHINT)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-odfinditem
func (me *_ListViewEvents) LvnODFindItem(userFunc func(p *win.NMLVFINDITEM) int) {
	me.events.addNfyRet(me.ctrlId, co.LVN_ODFINDITEM, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMLVFINDITEM)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-odstatechanged
func (me *_ListViewEvents) LvnODStateChanged(userFunc func(p *win.NMLVODSTATECHANGE)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_ODSTATECHANGED, func(p unsafe.Pointer) {
		userFunc((*win.NMLVODSTATECHANGE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvn-setdispinfo
func (me *_ListViewEvents) LvnSetDispInfo(userFunc func(p *win.NMLVDISPINFO)) {
	me.events.addNfyZero(me.ctrlId, co.LVN_SETDISPINFO, func(p unsafe.Pointer) {
		userFunc((*win.NMLVDISPINFO)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-list-view
func (me *_ListViewEvents) NmClick(userFunc func(p *win.NMITEMACTIVATE)) {
	me.events.addNfyZero(me.ctrlId, co.NM_CLICK, func(p unsafe.Pointer) {
		userFunc((*win.NMITEMACTIVATE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-customdraw-list-view
func (me *_ListViewEvents) NmCustomDraw(userFunc func(p *win.NMLVCUSTOMDRAW) co.CDRF) {
	me.events.addNfyRet(me.ctrlId, co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMLVCUSTOMDRAW)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-dblclk-list-view
func (me *_ListViewEvents) NmDblClk(userFunc func(p *win.NMITEMACTIVATE)) {
	me.events.addNfyZero(me.ctrlId, co.NM_DBLCLK, func(p unsafe.Pointer) {
		userFunc((*win.NMITEMACTIVATE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-hover-list-view
func (me *_ListViewEvents) NmHover(userFunc func() uint) {
	me.events.addNfyRet(me.ctrlId, co.NM_HOVER, func(_ unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-killfocus-list-view
func (me *_ListViewEvents) NmKillFocus(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_KILLFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-rclick-list-view
func (me *_ListViewEvents) NmRClick(userFunc func(p *win.NMITEMACTIVATE)) {
	me.events.addNfyZero(me.ctrlId, co.NM_RCLICK, func(p unsafe.Pointer) {
		userFunc((*win.NMITEMACTIVATE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-list-view
func (me *_ListViewEvents) LvnRDblClk(userFunc func(p *win.NMITEMACTIVATE)) {
	me.events.addNfyZero(me.ctrlId, co.NM_RDBLCLK, func(p unsafe.Pointer) {
		userFunc((*win.NMITEMACTIVATE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-list-view-
func (me *_ListViewEvents) LvnReleasedCapture(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-return-list-view-
func (me *_ListViewEvents) LvnReturn(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_RETURN, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-setfocus-list-view-
func (me *_ListViewEvents) LvnSetFocus(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_SETFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}
