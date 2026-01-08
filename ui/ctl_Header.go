//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
)

// Native [header] control.
//
// [header]: https://learn.microsoft.com/en-us/windows/win32/controls/header-controls
type Header struct {
	_BaseCtrl
	events EventsHeader
	Items  CollectionHeaderItems // Methods to interact with the items collection.
}

// Creates a new [Header] with [win.CreateWindowEx].
//
// Example:
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	header := ui.NewHeader(
//		wndOwner,
//		ui.OptsHeader().
//			Position(ui.Dpi(260, 240)).
//			Size(ui.Dpi(200, 23)),
//	)
func NewHeader(parent Parent, opts *VarOptsHeader) *Header {
	setUniqueCtrlId(&opts.ctrlId)
	me := &Header{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsHeader{opts.ctrlId, &parent.base().userEvents},
	}
	me.Items.owner = me

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.createWindow(opts.wndExStyle, "SysHeader32", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, true)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
	})

	me.defaultMessageHandlers(parent)
	return me
}

// Instantiates a new [Header] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
//
// Example:
//
//	const ID_HEADER uint16 = 0x100
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	header := ui.NewHeaderDlg(
//		wndOwner, ID_HEADER, ui.LAY_HOLD_HOLD)
func NewHeaderDlg(parent Parent, ctrlId uint16, layout LAY) *Header {
	me := &Header{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsHeader{ctrlId, &parent.base().userEvents},
	}
	me.Items.owner = me

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.assignDialog(parent)
		parent.base().layout.Add(parent, me.hWnd, layout)
	})

	me.defaultMessageHandlers(parent)
	return me
}

// Instantiates a [Header] from a [ListView] control.
func newHeaderFromListView(parent Parent) *Header {
	ctrlId := nextCtrlId()
	me := &Header{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsHeader{ctrlId, &parent.base().userEvents},
	}
	me.Items.owner = me

	me.defaultMessageHandlers(parent)
	return me
}

// Assigns the HWND to the [Header] if it belongs to an existing [ListView].
func (me *Header) assignToListView(hHeader win.HWND) {
	me._BaseCtrl.hWnd = hHeader
	hHeader.SetWindowLongPtr(co.GWLP_ID, uintptr(me._BaseCtrl.ctrlId)) // give the header an ID, initially is zero
}

func (me *Header) defaultMessageHandlers(parent Parent) {
	parent.base().afterUserEvents.wm(co.WM_DESTROY, func(_ Wm) {
		kinds := []co.HDSIL{co.HDSIL_NORMAL, co.HDSIL_STATE}
		for _, kind := range kinds {
			h, _ := me.hWnd.SendMessage(co.HDM_GETIMAGELIST, win.WPARAM(kind), 0)
			if h != 0 {
				me.hWnd.SendMessage(co.HDM_SETIMAGELIST, win.WPARAM(kind), 0) // release image list
				win.HIMAGELIST(h).Destroy()
			}
		}
	})
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *Header) On() *EventsHeader {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Retrieves the given image list with [HDM_GETIMAGELIST]. The image lists are
// lazy-initialized: the first time you call this method for a given image list,
// it will be created and assigned with [HDM_SETIMAGELIST].
//
// The image lists will be automatically destroyed.
//
// [HDM_GETIMAGELIST]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-getimagelist
// [HDM_SETIMAGELIST]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-setimagelist
func (me *Header) ImageList(which co.HDSIL) win.HIMAGELIST {
	h, _ := me.hWnd.SendMessage(co.HDM_GETIMAGELIST, win.WPARAM(which), 0)
	hImg := win.HIMAGELIST(h)
	if hImg == win.HIMAGELIST(0) {
		hImg, _ = win.ImageListCreate(16, 16, co.ILC_COLOR32, 1, 1)
		me.hWnd.SendMessage(co.HDM_SETIMAGELIST, win.WPARAM(which), win.LPARAM(hImg))
	}
	return hImg
}

// Options for [NewHeader]; returned by [OptsHeader].
type VarOptsHeader struct {
	ctrlId     uint16
	layout     LAY
	position   win.POINT
	size       win.SIZE
	ctrlStyle  co.HDS
	wndStyle   co.WS
	wndExStyle co.WS_EX
}

// Options for [NewHeader].
func OptsHeader() *VarOptsHeader {
	return &VarOptsHeader{
		size:      win.SIZE{Cx: int32(DpiX(100)), Cy: int32(DpiY(23))},
		ctrlStyle: co.HDS_BUTTONS | co.HDS_HORZ,
		wndStyle:  co.WS_CHILD | co.WS_GROUP | co.WS_VISIBLE | co.WS_BORDER,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsHeader) CtrlId(id uint16) *VarOptsHeader { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_HOLD_HOLD.
func (o *VarOptsHeader) Layout(l LAY) *VarOptsHeader { o.layout = l; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsHeader) Position(x, y int) *VarOptsHeader {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.Dpi(100, 23).
func (o *VarOptsHeader) Size(cx, cy int) *VarOptsHeader {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Header control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.HDS_BUTTONS | co.HDS_HORZ.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/header-control-styles
func (o *VarOptsHeader) CtrlStyle(s co.HDS) *VarOptsHeader { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_VISIBLE | co.WS_BORDER.
func (o *VarOptsHeader) WndStyle(s co.WS) *VarOptsHeader { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
func (o *VarOptsHeader) WndExStyle(s co.WS_EX) *VarOptsHeader { o.wndExStyle = s; return o }

// Native [header] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [header]: https://learn.microsoft.com/en-us/windows/win32/controls/header-controls
type EventsHeader struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [HDN_BEGINDRAG] message handler.
//
// [HDN_BEGINDRAG]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-begindrag
func (me *EventsHeader) HdnBeginDrag(fun func(p *win.NMHEADER) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_BEGINDRAG, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMHEADER)(p)))
	})
}

// [HDN_BEGINFILTEREDIT] message handler.
//
// [HDN_BEGINFILTEREDIT]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-beginfilteredit
func (me *EventsHeader) HdnBeginFilterEdit(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_BEGINFILTEREDIT, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_BEGINTRACK] message handler.
//
// [HDN_BEGINTRACK]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-begintrack
func (me *EventsHeader) HdnBeginTrack(fun func(p *win.NMHEADER) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_BEGINTRACK, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMHEADER)(p)))
	})
}

// [HDN_DIVIDERDBLCLICK] message handler.
//
// [HDN_DIVIDERDBLCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-dividerdblclick
func (me *EventsHeader) HdnDividerDblClick(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_DIVIDERDBLCLICK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_DROPDOWN] message handler.
//
// [HDN_DROPDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-dropdown
func (me *EventsHeader) HdnDropDown(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_DROPDOWN, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_ENDDRAG] message handler.
//
// [HDN_ENDDRAG]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-enddrag
func (me *EventsHeader) HdnEndDrag(fun func(p *win.NMHEADER) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_ENDDRAG, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMHEADER)(p)))
	})
}

// [HDN_ENDFILTEREDIT] message handler.
//
// [HDN_ENDFILTEREDIT]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-endfilteredit
func (me *EventsHeader) HdnEndFilterEdit(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_ENDFILTEREDIT, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_ENDTRACK] message handler.
//
// [HDN_ENDTRACK]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-endtrack
func (me *EventsHeader) HdnEndTrack(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_ENDTRACK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_FILTERBTNCLICK] message handler.
//
// [HDN_FILTERBTNCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-filterbtnclick
func (me *EventsHeader) HdnFilterBtnClick(fun func(p *win.NMHDFILTERBTNCLICK) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_FILTERBTNCLICK, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMHDFILTERBTNCLICK)(p)))
	})
}

// [HDN_FILTERCHANGE] message handler.
//
// [HDN_FILTERCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-filterchange
func (me *EventsHeader) HdnFilterChange(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_FILTERCHANGE, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_GETDISPINFO] message handler.
//
// [HDN_GETDISPINFO]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-getdispinfo
func (me *EventsHeader) HdnGetDispInfo(fun func(p *win.NMHDDISPINFO) uintptr) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_GETDISPINFO, func(p unsafe.Pointer) uintptr {
		return fun((*win.NMHDDISPINFO)(p))
	})
}

// [HDN_ITEMCHANGED] message handler.
//
// [HDN_ITEMCHANGED]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-itemchanged
func (me *EventsHeader) HdnItemChanged(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_ITEMCHANGED, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_ITEMCHANGING] message handler.
//
// [HDN_ITEMCHANGING]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-itemchanging
func (me *EventsHeader) HdnItemChanging(fun func(p *win.NMHEADER) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_ITEMCHANGING, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMHEADER)(p)))
	})
}

// [HDN_ITEMCLICK] message handler.
//
// [HDN_ITEMCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-itemclick
func (me *EventsHeader) HdnItemClick(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_ITEMCLICK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_ITEMDBLCLICK] message handler.
//
// [HDN_ITEMDBLCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-itemdblclick
func (me *EventsHeader) HdnItemDblClick(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_ITEMDBLCLICK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_ITEMKEYDOWN] message handler.
//
// [HDN_ITEMKEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-itemkeydown
func (me *EventsHeader) HdnItemKeyDown(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_ITEMKEYDOWN, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_ITEMSTATEICONCLICK] message handler.
//
// [HDN_ITEMSTATEICONCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-itemstateiconclick
func (me *EventsHeader) HdnItemStateIconClick(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_ITEMSTATEICONCLICK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_OVERFLOWCLICK] message handler.
//
// [HDN_OVERFLOWCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-overflowclick
func (me *EventsHeader) HdnOverflowClick(fun func(p *win.NMHEADER)) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_OVERFLOWCLICK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMHEADER)(p))
		return me.parentEvents.defProcVal
	})
}

// [HDN_TRACK] message handler.
//
// [HDN_TRACK]: https://learn.microsoft.com/en-us/windows/win32/controls/hdn-track
func (me *EventsHeader) HdnTrack(fun func(p *win.NMHEADER) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.HDN_TRACK, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMHEADER)(p)))
	})
}

// [NM_CUSTOMDRAW] message handler.
//
// [NM_CUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-customdraw-header
func (me *EventsHeader) NmCustomDraw(fun func(p *win.NMCUSTOMDRAW) co.CDRF) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMCUSTOMDRAW)(p)))
	})
}

// [NM_RCLICK] message handler.
//
// [NM_RCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rclick-header
func (me *EventsHeader) NmRClick(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RCLICK, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [NM_RELEASEDCAPTURE] message handler.
//
// [NM_RELEASEDCAPTURE]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-header-
func (me *EventsHeader) NmReleasedCapture(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}
