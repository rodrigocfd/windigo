//go:build windows

package ui

import (
	"fmt"
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
	AnyFocusControl
	implListView() // prevent public implementation

	// Exposes all the ListView notifications the can be handled.
	//
	// Panics if called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-list-view-control-reference-notifications
	On() *_ListViewEvents

	ContextMenu() win.HMENU                                           // Returns the associated context menu, if any.
	Columns() *_ListViewColumns                                       // Column methods.
	EditControl() win.HWND                                            // Retrieves a handle to the edit control being used.
	ExtendedStyle() co.LVS_EX                                         // Retrieves the extended style flags.
	ImageList(which co.LVSIL) win.HIMAGELIST                          // Retrieves one of the current image lists. If the list has LVS_SHAREIMAGELISTS, it's shared, otherwise it will be automatically destroyed.
	Items() *_ListViewItems                                           // Item methods.
	Scroll(horz, vert int)                                            // Scrolls the list view horizontally and vertically, in pixels, from its current position.
	SetExtendedStyle(doSet bool, styles co.LVS_EX)                    // Sets or unsets extended style flags.
	SetImageList(which co.LVSIL, himgl win.HIMAGELIST) win.HIMAGELIST // Sets one of the current image lists. If the list has LVS_SHAREIMAGELISTS, it's shared, otherwise it will be automatically destroyed.
	SetRedraw(allowRedraw bool)                                       // Sends WM_SETREDRAW to enable or disable UI updates.
	SetView(view co.LV_VIEW)                                          // Sets current view.
	View() co.LV_VIEW                                                 // Retrieves current view.
}

//------------------------------------------------------------------------------

type _ListView struct {
	_NativeControlBase
	events       _ListViewEvents
	columns      _ListViewColumns
	items        _ListViewItems
	hContextMenu win.HMENU
}

// Creates a new ListView. Call ui.ListViewOpts() to define the options to be
// passed to the underlying CreateWindowEx().
//
// Example:
//
//		var owner ui.AnyParent // initialized somewhere
//
//		myList := ui.NewListView(
//			owner,
//			ui.ListViewOpts(
//				Position(win.POINT{X: 10, Y: 20}).
//				Size(win.SIZE{Cx: 120}).
//				CtrlExStyles(co.LVS_EX_FULLROWSELECT),
//			),
//		)
func NewListView(parent AnyParent, opts *_ListViewO) ListView {
	if opts == nil {
		opts = ListViewOpts()
	}
	opts.lateDefaults()

	me := &_ListView{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)
	me.columns.new(me)
	me.items.new(me)
	me.hContextMenu = opts.contextMenu

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_ConvertDtuOrMultiplyDpi(parent, &opts.position, &opts.size)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			win.ClassNameStr("SysListView32"), win.StrOptNone(),
			opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		parent.addResizingChild(me, opts.horz, opts.vert)

		if opts.ctrlExStyles != co.LVS_EX_NONE {
			me.SetExtendedStyle(true, opts.ctrlExStyles)
		}
	})

	me.handledEvents()
	return me
}

// Creates a new ListView from a dialog resource.
func NewListViewDlg(
	parent AnyParent, ctrlId int,
	horz HORZ, vert VERT, contextMenuId int) ListView {

	hContextMenu := win.HMENU(0)
	if contextMenuId != 0 {
		hResMenu, found := win.HINSTANCE(0).
			LoadMenu(win.ResIdInt(contextMenuId)).GetSubMenu(0) // usually this is how it's set in the resources
		if !found {
			panic("ListView context menu not found.")
		}
		hContextMenu = hResMenu // menu resources are automatically freed by the system
	}

	me := &_ListView{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)
	me.columns.new(me)
	me.items.new(me)
	me.hContextMenu = hContextMenu

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
		parent.addResizingChild(me, horz, vert)
	})

	me.handledEvents()
	return me
}

// Implements ListView.
func (*_ListView) implListView() {}

// Implements AnyFocusControl.
func (me *_ListView) Focus() {
	me._NativeControlBase.focus()
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

func (me *_ListView) EditControl() win.HWND {
	return win.HWND(me.Hwnd().SendMessage(co.LVM_GETEDITCONTROL, 0, 0))
}

func (me *_ListView) ExtendedStyle() co.LVS_EX {
	return co.LVS_EX(
		me.Hwnd().SendMessage(co.LVM_GETEXTENDEDLISTVIEWSTYLE, 0, 0),
	)
}

func (me *_ListView) ImageList(which co.LVSIL) win.HIMAGELIST {
	return win.HIMAGELIST(
		me.Hwnd().SendMessage(co.LVM_GETIMAGELIST, win.WPARAM(which), 0),
	)
}

func (me *_ListView) Items() *_ListViewItems {
	return &me.items
}

func (me *_ListView) Scroll(horz, vert int) {
	if me.Hwnd().SendMessage(co.LVM_SCROLL, win.WPARAM(horz), win.LPARAM(vert)) == 0 {
		panic(fmt.Sprintf("ListView scrolling failed: %d, %d.", horz, vert))
	}
}

func (me *_ListView) SetExtendedStyle(doSet bool, styles co.LVS_EX) {
	affected := util.Iif(doSet, styles, 0).(co.LVS_EX)
	me.Hwnd().SendMessage(co.LVM_SETEXTENDEDLISTVIEWSTYLE,
		win.WPARAM(affected), win.LPARAM(styles))
}

func (me *_ListView) SetImageList(
	which co.LVSIL, himgl win.HIMAGELIST) win.HIMAGELIST {

	return win.HIMAGELIST(
		me.Hwnd().SendMessage(co.LVM_SETIMAGELIST,
			win.WPARAM(which), win.LPARAM(himgl)),
	)
}

func (me *_ListView) SetRedraw(allowRedraw bool) {
	me.Hwnd().SendMessage(co.WM_SETREDRAW,
		win.WPARAM(util.BoolToUintptr(allowRedraw)), 0)
}

func (me *_ListView) SetView(view co.LV_VIEW) {
	ret := me.Hwnd().SendMessage(co.LVM_SETVIEW, win.WPARAM(view), 0)
	if int(ret) == -1 {
		panic(fmt.Sprintf("LVM_SETVIEW failed for %d.", view))
	}
}

func (me *_ListView) View() co.LV_VIEW {
	return co.LV_VIEW(me.Hwnd().SendMessage(co.LVM_GETVIEW, 0, 0))
}

func (me *_ListView) handledEvents() {
	me.Parent().internalOn().addNfyZero(me.CtrlId(), co.LVN_KEYDOWN, func(p unsafe.Pointer) {
		nmk := (*win.NMLVKEYDOWN)(p)
		hasCtrl := (win.GetAsyncKeyState(co.VK_CONTROL) & 0x8000) != 0
		hasShift := (win.GetAsyncKeyState(co.VK_SHIFT) & 0x8000) != 0

		if hasCtrl && nmk.WVKey == 'A' { // Ctrl+A pressed?
			me.Items().SelectAll(true)
		} else if nmk.WVKey == co.VK_APPS { // context menu key
			me.showContextMenu(false, hasCtrl, hasShift)
		}
	})

	me.Parent().internalOn().addNfyZero(me.CtrlId(), co.NM_RCLICK, func(p unsafe.Pointer) {
		nmi := (*win.NMITEMACTIVATE)(p)
		hasCtrl := (nmi.UKeyFlags & co.LVKF_CONTROL) != 0
		hasShift := (nmi.UKeyFlags & co.LVKF_SHIFT) != 0

		me.showContextMenu(true, hasCtrl, hasShift)
	})
}

func (me *_ListView) showContextMenu(followCursor, hasCtrl, hasShift bool) {
	if me.hContextMenu == win.HMENU(0) { // no menu, nothing to do
		return
	}

	var menuPos win.POINT // menu anchor coords, relative to list view

	if followCursor { // usually when fired by a right-click
		menuPos = win.GetCursorPos()         // relative to screen
		me.Hwnd().ScreenToClientPt(&menuPos) // now relative to list view

		if clickedItem, hasClickedItem := me.Items().HitTest(menuPos); !hasClickedItem {
			me.Items().SelectAll(false)
		} else {
			if !hasCtrl && !hasShift {
				clickedItem.Focus()
			}
		}
		me.Hwnd().SetFocus() // because a right-click won't set the focus by itself

	} else { // usually fired with the context keyboard key
		if focusItem, hasFocused := me.Items().Focused(); hasFocused && focusItem.IsVisible() {
			rcItem := focusItem.Rect(co.LVIR_BOUNDS)
			menuPos.X = rcItem.Left + 16 // arbitrary
			menuPos.Y = rcItem.Top + (rcItem.Bottom-rcItem.Top)/2
		} else { // no item is focused and visible
			menuPos.X = 6 // arbitrary anchor coords
			menuPos.Y = 10
		}
	}

	me.hContextMenu.ShowAtPoint(menuPos, me.Hwnd().GetParent(), me.Hwnd())
}

//------------------------------------------------------------------------------

type _ListViewO struct {
	ctrlId int

	position     win.POINT
	size         win.SIZE
	horz         HORZ
	vert         VERT
	ctrlStyles   co.LVS
	ctrlExStyles co.LVS_EX
	wndStyles    co.WS
	wndExStyles  co.WS_EX

	contextMenu win.HMENU
}

// Control ID.
//
// Defaults to an auto-generated ID.
func (o *_ListViewO) CtrlId(i int) *_ListViewO { o.ctrlId = i; return o }

// Position within parent's client area.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 0x0.
func (o *_ListViewO) Position(p win.POINT) *_ListViewO { _OwPt(&o.position, p); return o }

// Control size.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 120x120.
func (o *_ListViewO) Size(s win.SIZE) *_ListViewO { _OwSz(&o.size, s); return o }

// Horizontal behavior when the parent is resized.
//
// Defaults to HORZ_NONE.
func (o *_ListViewO) Horz(s HORZ) *_ListViewO { o.horz = s; return o }

// Vertical behavior when the parent is resized.
//
// Defaults to VERT_NONE.
func (o *_ListViewO) Vert(s VERT) *_ListViewO { o.vert = s; return o }

// ListView control styles, passed to CreateWindowEx().
//
// Defaults to LVS_REPORT | LVS_NOSORTHEADER | LVS_SHOWSELALWAYS.
func (o *_ListViewO) CtrlStyles(s co.LVS) *_ListViewO { o.ctrlStyles = s; return o }

// ListView extended control styles, passed to CreateWindowEx().
//
// Defaults to LVS_EX_NONE.
func (o *_ListViewO) CtrlExStyles(s co.LVS_EX) *_ListViewO { o.ctrlExStyles = s; return o }

// Window styles, passed to CreateWindowEx().
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *_ListViewO) WndStyles(s co.WS) *_ListViewO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
//
// Defaults to WS_EX_CLIENTEDGE.
func (o *_ListViewO) WndExStyles(s co.WS_EX) *_ListViewO { o.wndExStyles = s; return o }

// Context menu for the list view. This handle is shared, the control won't destroy it.
// Defaults to none.
func (o *_ListViewO) ContextMenu(m win.HMENU) *_ListViewO { o.contextMenu = m; return o }

func (o *_ListViewO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewListView().
func ListViewOpts() *_ListViewO {
	return &_ListViewO{
		size:        win.SIZE{Cx: 120, Cy: 120},
		horz:        HORZ_NONE,
		vert:        VERT_NONE,
		ctrlStyles:  co.LVS_REPORT | co.LVS_NOSORTHEADER | co.LVS_SHOWSELALWAYS,
		wndStyles:   co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
		wndExStyles: co.WS_EX_CLIENTEDGE,
	}
}

//------------------------------------------------------------------------------

// ListView control notifications.
type _ListViewEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_ListViewEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.Parent().On()
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
func (me *_ListViewEvents) NmRDblClk(userFunc func(p *win.NMITEMACTIVATE)) {
	me.events.addNfyZero(me.ctrlId, co.NM_RDBLCLK, func(p unsafe.Pointer) {
		userFunc((*win.NMITEMACTIVATE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-list-view-
func (me *_ListViewEvents) NmReleasedCapture(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-return-list-view-
func (me *_ListViewEvents) NmReturn(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_RETURN, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-setfocus-list-view-
func (me *_ListViewEvents) NmSetFocus(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_SETFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}
