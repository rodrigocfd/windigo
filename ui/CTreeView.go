package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native tree view control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tree-view-controls
type TreeView interface {
	AnyNativeControl
	AnyFocusControl
	implTreeView() // prevent public implementation

	// Exposes all the TreeView notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-tree-view-control-reference-notifications
	On() *_TreeViewEvents

	Items() *_TreeViewItems // Item methods.
}

//------------------------------------------------------------------------------

type _TreeView struct {
	_NativeControlBase
	events _TreeViewEvents
	items  _TreeViewItems
}

// Creates a new TreeView specifying all options, which will be passed to the
// underlying CreateWindowEx().
//
// Example:
//
//  var owner AnyParent // initialized somewhere
//
//  myTree := ui.NewTreeView(
//      owner,
//      ui.TreeViewOpts(
//          Position(win.POINT{X: 10, Y: 240}).
//          Size(win.SIZE{Cx: 150, Cy: 100}),
//      ),
//  )
func NewTreeView(parent AnyParent, opts *_TreeViewO) TreeView {
	if opts == nil {
		opts = TreeViewOpts()
	}
	opts.lateDefaults()

	me := &_TreeView{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)
	me.items.new(me)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_ConvertDtuOrMultiplyDpi(parent, &opts.position, &opts.size)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			win.ClassNameStr("SysTreeView32"), win.StrOptNone(),
			opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		parent.addResizingChild(me, opts.horz, opts.vert)

		if opts.ctrlExStyles != co.TVS_EX_NONE {
			me.Hwnd().SendMessage(co.TVM_SETEXTENDEDSTYLE,
				win.WPARAM(opts.ctrlExStyles),
				win.LPARAM(opts.ctrlExStyles))
		}
	})

	return me
}

// Creates a new TreeView from a dialog resource.
func NewTreeViewDlg(
	parent AnyParent, ctrlId int,
	horz HORZ, vert VERT) TreeView {

	me := &_TreeView{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)
	me.items.new(me)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
		parent.addResizingChild(me, horz, vert)
	})

	return me
}

// Implements TreeView.
func (*_TreeView) implTreeView() {}

// Implements AnyFocusControl.
func (me *_TreeView) Focus() {
	me._NativeControlBase.focus()
}

func (me *_TreeView) On() *_TreeViewEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the TreeView is created.")
	}
	return &me.events
}

func (me *_TreeView) Items() *_TreeViewItems {
	return &me.items
}

//------------------------------------------------------------------------------

type _TreeViewO struct {
	ctrlId int

	position     win.POINT
	size         win.SIZE
	horz         HORZ
	vert         VERT
	ctrlStyles   co.TVS
	ctrlExStyles co.TVS_EX
	wndStyles    co.WS
	wndExStyles  co.WS_EX
}

// Control ID.
//
// Defaults to an auto-generated ID.
func (o *_TreeViewO) CtrlId(i int) *_TreeViewO { o.ctrlId = i; return o }

// Position within parent's client area.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 0x0.
func (o *_TreeViewO) Position(p win.POINT) *_TreeViewO { _OwPt(&o.position, p); return o }

// Control size.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 120x120.
func (o *_TreeViewO) Size(s win.SIZE) *_TreeViewO { _OwSz(&o.size, s); return o }

// Horizontal behavior when the parent is resized.
//
// Defaults to HORZ_NONE.
func (o *_TreeViewO) Horz(s HORZ) *_TreeViewO { o.horz = s; return o }

// Vertical behavior when the parent is resized.
//
// Defaults to VERT_NONE.
func (o *_TreeViewO) Vert(s VERT) *_TreeViewO { o.vert = s; return o }

// TreeView control styles, passed to CreateWindowEx().
//
// Defaults to TVS_HASLINES | TVS_LINESATROOT | TVS_SHOWSELALWAYS | TVS_HASBUTTONS.
func (o *_TreeViewO) CtrlStyles(s co.TVS) *_TreeViewO { o.ctrlStyles = s; return o }

// TreeView extended control styles, passed to CreateWindowEx().
//
// Defaults to LVS_EX_NONE.
func (o *_TreeViewO) CtrlExStyles(s co.TVS_EX) *_TreeViewO { o.ctrlExStyles = s; return o }

// Window styles, passed to CreateWindowEx().
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *_TreeViewO) WndStyles(s co.WS) *_TreeViewO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
//
// Defaults to WS_EX_CLIENTEDGE.
func (o *_TreeViewO) WndExStyles(s co.WS_EX) *_TreeViewO { o.wndExStyles = s; return o }

func (o *_TreeViewO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewTreeView().
func TreeViewOpts() *_TreeViewO {
	return &_TreeViewO{
		size: win.SIZE{Cx: 120, Cy: 120},
		horz: HORZ_NONE,
		vert: VERT_NONE,
		ctrlStyles: co.TVS_HASLINES | co.TVS_LINESATROOT |
			co.TVS_SHOWSELALWAYS | co.TVS_HASBUTTONS,
		wndStyles:   co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
		wndExStyles: co.WS_EX_CLIENTEDGE,
	}
}

//------------------------------------------------------------------------------

// TreeView control notifications.
type _TreeViewEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_TreeViewEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.Parent().On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-asyncdraw
func (me *_TreeViewEvents) TvnAsyncDraw(userFunc func(p *win.NMTVASYNCDRAW)) {
	me.events.addNfyZero(me.ctrlId, co.TVN_ASYNCDRAW, func(p unsafe.Pointer) {
		userFunc((*win.NMTVASYNCDRAW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-begindrag
func (me *_TreeViewEvents) TvnBeginDrag(userFunc func(p *win.NMTREEVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.TVN_BEGINDRAG, func(p unsafe.Pointer) {
		userFunc((*win.NMTREEVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-beginlabeledit
func (me *_TreeViewEvents) TvnBeginLabelEdit(userFunc func(p *win.NMTVDISPINFO) bool) {
	me.events.addNfyRet(me.ctrlId, co.TVN_BEGINLABELEDIT, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMTVDISPINFO)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-beginrdrag
func (me *_TreeViewEvents) TvnBeginRDrag(userFunc func(p *win.NMTREEVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.TVN_BEGINRDRAG, func(p unsafe.Pointer) {
		userFunc((*win.NMTREEVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-deleteitem
func (me *_TreeViewEvents) TvnDeleteItem(userFunc func(p *win.NMTREEVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.TVN_DELETEITEM, func(p unsafe.Pointer) {
		userFunc((*win.NMTREEVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-endlabeledit
func (me *_TreeViewEvents) TvnEndLabelEdit(userFunc func(p *win.NMTVDISPINFO) bool) {
	me.events.addNfyRet(me.ctrlId, co.TVN_ENDLABELEDIT, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMTVDISPINFO)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-getdispinfo
func (me *_TreeViewEvents) TvnGetDispInfo(userFunc func(p *win.NMTVDISPINFO)) {
	me.events.addNfyZero(me.ctrlId, co.TVN_GETDISPINFO, func(p unsafe.Pointer) {
		userFunc((*win.NMTVDISPINFO)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-getinfotip
func (me *_TreeViewEvents) TvnGetInfoTip(userFunc func(p *win.NMTVGETINFOTIP)) {
	me.events.addNfyZero(me.ctrlId, co.TVN_GETINFOTIP, func(p unsafe.Pointer) {
		userFunc((*win.NMTVGETINFOTIP)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemchanged
func (me *_TreeViewEvents) TvnItemChanged(userFunc func(p *win.NMTVITEMCHANGE)) {
	me.events.addNfyZero(me.ctrlId, co.TVN_ITEMCHANGED, func(p unsafe.Pointer) {
		userFunc((*win.NMTVITEMCHANGE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemchanging
func (me *_TreeViewEvents) TvnItemChanging(userFunc func(p *win.NMTVITEMCHANGE) bool) {
	me.events.addNfyRet(me.ctrlId, co.TVN_ITEMCHANGING, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMTVITEMCHANGE)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemexpanded
func (me *_TreeViewEvents) TvnItemExpanded(userFunc func(p *win.NMTREEVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.TVN_ITEMEXPANDED, func(p unsafe.Pointer) {
		userFunc((*win.NMTREEVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemexpanding
func (me *_TreeViewEvents) TvnItemExpanding(userFunc func(p *win.NMTREEVIEW) bool) {
	me.events.addNfyRet(me.ctrlId, co.TVN_ITEMEXPANDING, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMTREEVIEW)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-keydown
func (me *_TreeViewEvents) TvnKeyDown(userFunc func(p *win.NMTVKEYDOWN) int) {
	me.events.addNfyRet(me.ctrlId, co.TVN_KEYDOWN, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMTVKEYDOWN)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-selchanged
func (me *_TreeViewEvents) TvnSelChanged(userFunc func(p *win.NMTREEVIEW)) {
	me.events.addNfyZero(me.ctrlId, co.TVN_SELCHANGED, func(p unsafe.Pointer) {
		userFunc((*win.NMTREEVIEW)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-selchanging
func (me *_TreeViewEvents) TvnSelChanging(userFunc func(p *win.NMTREEVIEW) bool) {
	me.events.addNfyRet(me.ctrlId, co.TVN_SELCHANGING, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMTREEVIEW)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-setdispinfo
func (me *_TreeViewEvents) TvnSetDispInfo(userFunc func(p *win.NMTVDISPINFO)) {
	me.events.addNfyZero(me.ctrlId, co.TVN_SETDISPINFO, func(p unsafe.Pointer) {
		userFunc((*win.NMTVDISPINFO)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-singleexpand
func (me *_TreeViewEvents) TvnSingleExpand(userFunc func(p *win.NMTREEVIEW) co.TVNRET) {
	me.events.addNfyRet(me.ctrlId, co.TVN_SINGLEEXPAND, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMTREEVIEW)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-tree-view
func (me *_TreeViewEvents) NmClick(userFunc func() int) {
	me.events.addNfyRet(me.ctrlId, co.NM_CLICK, func(_ unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-customdraw-tree-view
func (me *_TreeViewEvents) NmCustomDraw(userFunc func(p *win.NMTVCUSTOMDRAW) co.CDRF) {
	me.events.addNfyRet(me.ctrlId, co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMTVCUSTOMDRAW)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-dblclk-tree-view
func (me *_TreeViewEvents) NmDblClk(userFunc func() int) {
	me.events.addNfyRet(me.ctrlId, co.NM_DBLCLK, func(_ unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-killfocus-tree-view
func (me *_TreeViewEvents) NmKillFocus(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_KILLFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-rclick-tree-view
func (me *_TreeViewEvents) NmRClick(userFunc func() int) {
	me.events.addNfyRet(me.ctrlId, co.NM_RCLICK, func(_ unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-tree-view
func (me *_TreeViewEvents) NmRDblClk(userFunc func() int) {
	me.events.addNfyRet(me.ctrlId, co.NM_RDBLCLK, func(_ unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-return-tree-view-
func (me *_TreeViewEvents) NmReturn(userFunc func() int) {
	me.events.addNfyRet(me.ctrlId, co.NM_RETURN, func(_ unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-setcursor-tree-view-
func (me *_TreeViewEvents) NmSetCursor(userFunc func(p *win.NMMOUSE) int) {
	me.events.addNfyRet(me.ctrlId, co.NM_SETCURSOR, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-setfocus-tree-view-
func (me *_TreeViewEvents) NmSetFocus(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_SETFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}
