//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [toolbar] control.
//
// [toolbar]: https://learn.microsoft.com/en-us/windows/win32/controls/toolbar-controls-overview
type Toolbar struct {
	_BaseCtrl
	events  EventsToolbar
	Buttons CollectionToolbarButtons // Methods to interact with the buttons collection.
}

// Creates a new Trackbar with [CreateWindowEx].
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func NewToolbar(parent Parent, opts *VarOptsToolbar) *Toolbar {
	setUniqueCtrlId(&opts.ctrlId)
	me := &Toolbar{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsToolbar{opts.ctrlId, &parent.base().userEvents},
	}
	me.Buttons.owner = me

	parent.base().beforeUserEvents.Wm(parent.base().wndTy.initMsg(), func(_ Wm) uintptr {
		me.createWindow(opts.wndExStyle, "ToolbarWindow32", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), win.POINT{}, win.SIZE{}, parent, false)
		me.hWnd.SendMessage(co.TB_BUTTONSTRUCTSIZE,
			win.WPARAM(unsafe.Sizeof(win.TBBUTTON{})), 0) // necessary before TB_ADDBUTTONS
		me.hWnd.SendMessage(co.CCM_SETVERSION, 5, 0)
		if opts.ctrlExStyle != co.TBSTYLE_EX_NONE {
			me.SetExtendedStyle(true, opts.ctrlExStyle)
		}
		return 0 // ignored
	})

	me.defaultMessageHandlers(parent)
	return me
}

func (me *Toolbar) defaultMessageHandlers(parent Parent) {
	parent.base().afterUserEvents.WmDestroy(func() {
		h, _ := me.hWnd.SendMessage(co.TB_GETIMAGELIST, 0, 0)
		if h != 0 {
			me.hWnd.SendMessage(co.TB_SETIMAGELIST, 0, 0)
			win.HIMAGELIST(h).Destroy()
		}
	})
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *Toolbar) On() *EventsToolbar {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Retrieves the extended style with [TB_GETEXTENDEDSTYLE].
//
// [TB_GETEXTENDEDSTYLE]: https://learn.microsoft.com/en-us/windows/win32/controls/tb-getextendedstyle
func (me *Toolbar) ExtendedStyle() co.TBSTYLE_EX {
	ret, _ := me.hWnd.SendMessage(co.TB_GETEXTENDEDSTYLE, 0, 0)
	return co.TBSTYLE_EX(ret)
}

// Retrieves the image list with [TB_GETIMAGELIST]. The image list is
// lazy-initialized: the first time you call this method, it will be created and
// assigned with [TB_SETIMAGELIST].
//
// The icon size is used to create the image list on the first call. Subsequent
// calls will ignore cx and cy parameters.
//
// The image list will be automatically destroyed.
//
// [TB_GETIMAGELIST]: https://learn.microsoft.com/en-us/windows/win32/controls/tb-getimagelist
// [TB_SETIMAGELIST]: https://learn.microsoft.com/en-us/windows/win32/controls/tb-setimagelist
func (me *Toolbar) ImageList(cx, cy int) win.HIMAGELIST {
	h, _ := me.hWnd.SendMessage(co.TB_GETIMAGELIST, 0, 0)
	hImg := win.HIMAGELIST(h)
	if hImg == win.HIMAGELIST(0) {
		hImg, _ = win.ImageListCreate(uint(cx), uint(cy), co.ILC_COLOR32, 1, 1)
		me.hWnd.SendMessage(co.TB_SETIMAGELIST, 0, win.LPARAM(hImg))
	}
	return hImg
}

// Adds or removes extended styles with [TB_SETEXTENDEDSTYLE].
//
// Returns the same object, so further operations can be chained.
//
// [TVM_SETEXTENDEDSTYLE]: https://learn.microsoft.com/en-us/windows/win32/controls/tb-setextendedstyle
func (me *Toolbar) SetExtendedStyle(doSet bool, style co.TBSTYLE_EX) *Toolbar {
	newStyle := me.ExtendedStyle()
	if doSet {
		newStyle |= style
	} else {
		newStyle &= ^style
	}
	me.Hwnd().SendMessage(co.TB_SETEXTENDEDSTYLE, 0, win.LPARAM(newStyle))
	return me
}

// Options for ui.NewToolbar(); returned by ui.OptsToolbar().
type VarOptsToolbar struct {
	ctrlId      uint16
	ctrlStyle   co.TBSTYLE
	ctrlExStyle co.TBSTYLE_EX
	wndStyle    co.WS
	wndExStyle  co.WS_EX
}

// Options for ui.NewToolbar().
func OptsToolbar() *VarOptsToolbar {
	return &VarOptsToolbar{
		ctrlStyle: co.TBSTYLE_BUTTON | co.TBSTYLE_FLAT | co.TBSTYLE_LIST,
		wndStyle:  co.WS_CHILD | co.WS_VISIBLE,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsToolbar) CtrlId(id uint16) *VarOptsToolbar { o.ctrlId = id; return o }

// Toolbar control [style], passed to [CreateWindowEx].
//
// Defaults to co.TBSTYLE_BUTTON | co.TBSTYLE_FLAT | co.TBSTYLE_LIST.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/toolbar-control-and-button-styles
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsToolbar) CtrlStyle(s co.TBSTYLE) *VarOptsToolbar { o.ctrlStyle = s; return o }

// Toolbar control [extended style].
//
// Defaults to co.TBSTYLE_EX_NONE.
//
// [extended style]: https://learn.microsoft.com/en-us/windows/win32/controls/toolbar-extended-styles
func (o *VarOptsToolbar) CtrlExStyle(s co.TBSTYLE_EX) *VarOptsToolbar { o.ctrlExStyle = s; return o }

// Window style, passed to [CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsToolbar) WndStyle(s co.WS) *VarOptsToolbar { o.wndStyle = s; return o }

// Window extended style, passed to [CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsToolbar) WndExStyle(s co.WS_EX) *VarOptsToolbar { o.wndExStyle = s; return o }

// Native [toolbar] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [toolbar]: https://learn.microsoft.com/en-us/windows/win32/controls/toolbar-controls-overview
type EventsToolbar struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [TBN_BEGINADJUST] message handler.
//
// [TBN_BEGINADJUST]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-beginadjust
func (me *EventsToolbar) TbnBeginAdjust(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_BEGINADJUST, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [TBN_BEGINDRAG] message handler.
//
// [TBN_BEGINDRAG]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-begindrag
func (me *EventsToolbar) TbnBeginDrag(fun func(p *win.NMTOOLBAR)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_BEGINDRAG, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTOOLBAR)(p))
		return me.parentEvents.defProcVal
	})
}

// [TBN_CUSTHELP] message handler.
//
// [TBN_CUSTHELP]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-custhelp
func (me *EventsToolbar) TbnCustHelp(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_CUSTHELP, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [TBN_DELETINGBUTTON] message handler.
//
// [TBN_DELETINGBUTTON]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-deletingbutton
func (me *EventsToolbar) TbnDeletingButton(fun func(p *win.NMTOOLBAR)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_DELETINGBUTTON, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTOOLBAR)(p))
		return me.parentEvents.defProcVal
	})
}

// [TBN_DRAGOUT] message handler.
//
// [TBN_DRAGOUT]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-dragout
func (me *EventsToolbar) TbnDragOut(fun func(p *win.NMTOOLBAR)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_DRAGOUT, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTOOLBAR)(p))
		return me.parentEvents.defProcVal
	})
}

// [TBN_DRAGOVER] message handler.
//
// [TBN_DRAGOVER]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-dragover
func (me *EventsToolbar) TbnDragOver(fun func(p *win.NMTBHOTITEM) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_DRAGOVER, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTBHOTITEM)(p)))
	})
}

// [TBN_DROPDOWN] message handler.
//
// [TBN_DROPDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-dropdown
func (me *EventsToolbar) TbnDropDown(fun func(p *win.NMTOOLBAR) co.TBDDRET) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_DROPDOWN, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMTOOLBAR)(p)))
	})
}

// [TBN_DUPACCELERATOR] message handler.
//
// [TBN_DUPACCELERATOR]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-dupaccelerator
func (me *EventsToolbar) TbnDupAccelerator(fun func(p *win.NMTBDUPACCELERATOR) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_DUPACCELERATOR, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTBDUPACCELERATOR)(p)))
	})
}

// [TBN_ENDADJUST] message handler.
//
// [TBN_ENDADJUST]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-endadjust
func (me *EventsToolbar) TbnEndAdjust(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_ENDADJUST, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [TBN_ENDDRAG] message handler.
//
// [TBN_ENDDRAG]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-enddrag
func (me *EventsToolbar) TbnEndDrag(fun func(p *win.NMTOOLBAR)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_ENDDRAG, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTOOLBAR)(p))
		return me.parentEvents.defProcVal
	})
}

// [TBN_GETBUTTONINFO] message handler.
//
// [TBN_GETBUTTONINFO]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-getbuttoninfo
func (me *EventsToolbar) TbnGetButtonInfo(fun func(p *win.NMTOOLBAR) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_GETBUTTONINFO, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTOOLBAR)(p)))
	})
}

// [TBN_GETDISPINFO] message handler.
//
// [TBN_GETDISPINFO]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-getdispinfo
func (me *EventsToolbar) TbnGetDispInfo(fun func(p *win.NMTBDISPINFO)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_GETDISPINFO, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTBDISPINFO)(p))
		return me.parentEvents.defProcVal
	})
}

// [TBN_GETINFOTIP] message handler.
//
// [TBN_GETINFOTIP]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-getinfotip
func (me *EventsToolbar) TbnGetInfoTip(fun func(p *win.NMTBGETINFOTIP)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_GETINFOTIP, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTBGETINFOTIP)(p))
		return me.parentEvents.defProcVal
	})
}

// [TBN_GETOBJECT] message handler.
//
// [TBN_GETOBJECT]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-getobject
func (me *EventsToolbar) TbnGetObject(fun func(p *win.NMOBJECTNOTIFY)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_GETOBJECT, func(p unsafe.Pointer) uintptr {
		fun((*win.NMOBJECTNOTIFY)(p))
		return me.parentEvents.defProcVal
	})
}

// [TBN_HOTITEMCHANGE] message handler.
//
// [TBN_HOTITEMCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-hotitemchange
func (me *EventsToolbar) TbnHotItemChange(fun func(*win.NMTBHOTITEM) int) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_HOTITEMCHANGE, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMTBHOTITEM)(p)))
	})
}

// [TBN_INITCUSTOMIZE] message handler.
//
// [TBN_INITCUSTOMIZE]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-initcustomize
func (me *EventsToolbar) TbnInitCustomize(fun func() co.TBNRF) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_INITCUSTOMIZE, func(p unsafe.Pointer) uintptr {
		return uintptr(fun())
	})
}

// [TBN_MAPACCELERATOR] message handler.
//
// [TBN_MAPACCELERATOR]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-mapaccelerator
func (me *EventsToolbar) TbnMapAccelerator(fun func(p *win.NMCHAR) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_MAPACCELERATOR, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMCHAR)(p)))
	})
}

// [TBN_QUERYDELETE] message handler.
//
// [TBN_QUERYDELETE]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-querydelete
func (me *EventsToolbar) TbnQueryDelete(fun func(p *win.NMTOOLBAR) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_QUERYDELETE, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTOOLBAR)(p)))
	})
}

// [TBN_QUERYINSERT] message handler.
//
// [TBN_QUERYINSERT]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-queryinsert
func (me *EventsToolbar) TbnQueryInsert(fun func(p *win.NMTOOLBAR) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_QUERYINSERT, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTOOLBAR)(p)))
	})
}

// [TBN_RESET] message handler.
//
// [TBN_RESET]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-reset
func (me *EventsToolbar) TbnReset(fun func() co.TBNRF) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_RESET, func(p unsafe.Pointer) uintptr {
		return uintptr(fun())
	})
}

// [TBN_RESTORE] message handler.
//
// [TBN_RESTORE]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-restore
func (me *EventsToolbar) TbnRestore(fun func(p *win.NMTBRESTORE) int) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_RESTORE, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMTBRESTORE)(p)))
	})
}

// [TBN_SAVE] message handler.
//
// [TBN_SAVE]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-save
func (me *EventsToolbar) TbnSave(fun func(p *win.NMTBSAVE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_SAVE, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTBSAVE)(p))
		return me.parentEvents.defProcVal
	})
}

// [TBN_TOOLBARCHANGE] message handler.
//
// [TBN_TOOLBARCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-toolbarchange
func (me *EventsToolbar) TbnToolbarChange(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_TOOLBARCHANGE, func(p unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [TBN_WRAPACCELERATOR] message handler.
//
// [TBN_WRAPACCELERATOR]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-wrapaccelerator
func (me *EventsToolbar) TbnWrapAccelerator(fun func(p *win.NMTBWRAPACCELERATOR) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_WRAPACCELERATOR, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTBWRAPACCELERATOR)(p)))
	})
}

// [TBN_WRAPHOTITEM] message handler.
//
// [TBN_WRAPHOTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-wraphotitem
func (me *EventsToolbar) TbnWrapHotItem(fun func(p *win.NMTBWRAPHOTITEM) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TBN_WRAPHOTITEM, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTBWRAPHOTITEM)(p)))
	})
}

// [NM_CHAR] message handler.
//
// [NM_CHAR]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-char-toolbar
func (me *EventsToolbar) NmChar(fun func(p *win.NMCHAR) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CHAR, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMCHAR)(p)))
	})
}

// [NM_CLICK] message handler.
//
// [NM_CLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-click-toolbar
func (me *EventsToolbar) NmClick(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CLICK, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_CUSTOMDRAW] message handler.
//
// [NM_CUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-customdraw-toolbar
func (me *EventsToolbar) NmCustomDraw(fun func(p *win.NMTBCUSTOMDRAW) co.CDRF) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMTBCUSTOMDRAW)(p)))
	})
}

// [NM_DBLCLK] message handler.
//
// [NM_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-dblclk-toolbar
func (me *EventsToolbar) NmDblClk(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_DBLCLK, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_KEYDOWN] message handler.
//
// [NM_KEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-keydown-toolbar
func (me *EventsToolbar) NmKeyDown(fun func(p *win.NMKEY) int) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_KEYDOWN, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMKEY)(p)))
	})
}

// [NM_LDOWN] message handler.
//
// [NM_LDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-ldown-toolbar
func (me *EventsToolbar) NmLDown(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_LDOWN, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_RCLICK] message handler.
//
// [NM_RCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rclick-toolbar
func (me *EventsToolbar) NmRClick(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RCLICK, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_RDBLCLK] message handler.
//
// [NM_RDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-toolbar
func (me *EventsToolbar) NmRDblClk(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RDBLCLK, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_RELEASEDCAPTURE] message handler.
//
// [NM_RELEASEDCAPTURE]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-list-view-
func (me *EventsToolbar) NmReleasedCapture(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [NM_TOOLTIPSCREATED] message handler.
//
// [NM_TOOLTIPSCREATED]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-tooltipscreated-toolbar-
func (me *EventsToolbar) NmTooltipsCreated(fun func(p *win.NMTOOLTIPSCREATED)) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_TOOLTIPSCREATED, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTOOLTIPSCREATED)(p))
		return me.parentEvents.defProcVal
	})
}
