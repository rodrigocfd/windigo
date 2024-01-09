//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [tab] control.
//
// [tab]: https://learn.microsoft.com/en-us/windows/win32/controls/tab-controls
type Tab interface {
	AnyNativeControl
	AnyFocusControl
	implTab() // prevent public implementation

	// Exposes all the [tab notifications] that can be handled.
	//
	// Panics if called after the control was created.
	//
	// [tab notifications]: https://docs.microsoft.com/en-us/windows/win32/controls/bumper-tab-control-reference-notifications
	On() *_TabEvents

	Items() *_TabItems // Item methods.
	Child() AnyParent  // Returns the inner window
}

// ------------------------------------------------------------------------------

type _Tab struct {
	_NativeControlBase
	_TabWindow

	events _TabEvents
	items  _TabItems
}

type _TabWindow struct {
	_WindowBase
}

// Creates a new TreView. Call ui.TabOpts() to define the options to be
// passed to the underlying CreateWindowEx().
//
// # Example
//
//	var owner AnyParent // initialized somewhere
//
//	myTree := ui.NewTab(
//		owner,
//		ui.TabOpts().
//			Position(win.POINT{X: 10, Y: 240}).
//			Size(win.SIZE{Cx: 150, Cy: 100}),
//		),
//	)
func NewTab(parent AnyParent, opts *_TabO) Tab {
	if opts == nil {
		opts = TabOpts()
	}
	opts.lateDefaults()

	me := &_Tab{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me._TabWindow.new()
	me.events.new(&me._NativeControlBase)
	me.items.new(me)

	me._TabWindow.events = *me.events.events
	me._TabWindow.internalEvents = *parent.internalOn()

	parent.internalOn().addMsgNoRet(_CreateOrInitDialog(parent), func(_ wm.Any) {
		var icx win.INITCOMMONCONTROLSEX
		icx.SetDwSize()
		icx.DwICC = co.ICC_TAB_CLASSES
		win.InitCommonControlsEx(&icx)

		_ConvertDtuOrMultiplyDpi(parent, &opts.position, &opts.size)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			win.ClassNameStr("SysTabControl32"), win.StrOptNone(),
			opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		me._TabWindow.hWnd = me.Hwnd()

		parent.addResizingChild(me, opts.horz, opts.vert)

		if opts.ctrlExStyles != co.TCS_EX_NONE {
			me.Hwnd().SendMessage(co.TCM_SETEXTENDEDSTYLE,
				win.WPARAM(opts.ctrlExStyles),
				win.LPARAM(opts.ctrlExStyles))
		}
		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
	})

	return me
}

// Creates a new Tab from a dialog resource.
func NewTabDlg(
	parent AnyParent, ctrlId int,
	horz HORZ, vert VERT) Tab {

	me := &_Tab{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)
	me.items.new(me)

	parent.internalOn().addMsgNoRet(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
		parent.addResizingChild(me, horz, vert)
	})

	return me
}

// Implements Tab.
func (*_Tab) implTab() {}

func (me *_Tab) Hwnd() win.HWND {
	return me._NativeControlBase.Hwnd()
}

func (me *_Tab) Child() AnyParent {
	return &me._TabWindow
}

// Implements AnyFocusControl.
func (me *_Tab) Focus() {
	me._NativeControlBase.focus()
}

func (me *_Tab) On() *_TabEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the Tab is created.")
	}
	return &me.events
}

func (me *_TabWindow) isDialog() bool {
	return false
}

func (me *_Tab) Items() *_TabItems {
	return &me.items
}

// ------------------------------------------------------------------------------

type _TabO struct {
	ctrlId int

	position     win.POINT
	size         win.SIZE
	horz         HORZ
	vert         VERT
	ctrlStyles   co.TCS
	ctrlExStyles co.TCS_EX
	wndStyles    co.WS
	wndExStyles  co.WS_EX
}

// Control ID.
//
// Defaults to an auto-generated ID.
func (o *_TabO) CtrlId(i int) *_TabO { o.ctrlId = i; return o }

// Position within parent's client area.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 0x0.
func (o *_TabO) Position(p win.POINT) *_TabO { _OwPt(&o.position, p); return o }

// Control size.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 120x120.
func (o *_TabO) Size(s win.SIZE) *_TabO { _OwSz(&o.size, s); return o }

// Horizontal behavior when the parent is resized.
//
// Defaults to HORZ_NONE.
func (o *_TabO) Horz(s HORZ) *_TabO { o.horz = s; return o }

// Vertical behavior when the parent is resized.
//
// Defaults to VERT_NONE.
func (o *_TabO) Vert(s VERT) *_TabO { o.vert = s; return o }

// Tab control styles, passed to CreateWindowEx().
//
// Defaults to TCS_MULTILINE | TCS_TABS.
func (o *_TabO) CtrlStyles(s co.TCS) *_TabO { o.ctrlStyles = s; return o }

// Tab extended control styles, passed to CreateWindowEx().
//
// Defaults to TCS_EX_NONE.
func (o *_TabO) CtrlExStyles(s co.TCS_EX) *_TabO { o.ctrlExStyles = s; return o }

// Window styles, passed to CreateWindowEx().
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *_TabO) WndStyles(s co.WS) *_TabO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
//
// Defaults to WS_EX_CLIENTEDGE.
func (o *_TabO) WndExStyles(s co.WS_EX) *_TabO { o.wndExStyles = s; return o }

func (o *_TabO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewTab().
func TabOpts() *_TabO {
	return &_TabO{
		size:        win.SIZE{Cx: 120, Cy: 120},
		horz:        HORZ_NONE,
		vert:        VERT_NONE,
		ctrlStyles:  co.TCS_MULTILINE | co.TCS_TABS,
		wndStyles:   co.WS_CHILD | co.WS_CLIPSIBLINGS | co.WS_VISIBLE,
		wndExStyles: co.WS_EX_NONE,
	}
}

// ------------------------------------------------------------------------------

// Tab control notifications.
type _TabEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_TabEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.Parent().On()
}

// [TVN_ASYNCDRAW] message handler.
//
// [TVN_ASYNCDRAW]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-asyncdraw
// func (me *_TabEvents) TvnAsyncDraw(userFunc func(p *win.NMTVASYNCDRAW)) {
// 	me.events.addNfyZero(me.ctrlId, co.TVN_ASYNCDRAW, func(p unsafe.Pointer) {
// 		userFunc((*win.NMTVASYNCDRAW)(p))
// 	})
// }

// [TCN_FOCUSCHANGE] message handler.
//
// [TCN_FOCUSCHANGE]: https://docs.microsoft.com/en-us/windows/win32/controls/tcn-focuschange
func (me *_TabEvents) TcnFocusChange(userFunc func(p *win.NMHDR)) {
	me.events.addNfyZero(me.ctrlId, co.TCN_FOCUSCHANGE, func(p unsafe.Pointer) {
		userFunc((*win.NMHDR)(p))
	})
}

// [TCN_GETOBJECT] message handler.
//
// [TCN_GETOBJECT]: https://docs.microsoft.com/en-us/windows/win32/controls/tcn-getobject
func (me *_TabEvents) TcnGetObject(userFunc func(p *win.NMOBJECTNOTIFY)) {
	me.events.addNfyRet(me.ctrlId, co.TCN_GETOBJECT, func(p unsafe.Pointer) uintptr {
		userFunc((*win.NMOBJECTNOTIFY)(p))
		return 0
	})
}

// [TCN_KEYDOWN] message handler.
//
// [TCN_KEYDOWN]: https://docs.microsoft.com/en-us/windows/win32/controls/tcn-keydown
func (me *_TabEvents) TcnKeyDown(userFunc func(p *win.NMTCKEYDOWN)) {
	me.events.addNfyZero(me.ctrlId, co.TCN_KEYDOWN, func(p unsafe.Pointer) {
		userFunc((*win.NMTCKEYDOWN)(p))
	})
}

// [TCN_SELCHANGE] message handler.
//
// [TCN_SELCHANGE]: https://docs.microsoft.com/en-us/windows/win32/controls/tcn-selchange
func (me *_TabEvents) TcnSelChange(userFunc func(p *win.NMHDR)) {
	me.events.addNfyZero(me.ctrlId, co.TCN_SELCHANGE, func(p unsafe.Pointer) {
		userFunc((*win.NMHDR)(p))
	})
}

// [TCN_SELCHANGING] message handler.
//
// [TCN_SELCHANGING]: https://docs.microsoft.com/en-us/windows/win32/controls/tcn-selchanging
func (me *_TabEvents) TcnSelChanging(userFunc func(p *win.NMHDR) bool) {
	me.events.addNfyRet(me.ctrlId, co.TCN_SELCHANGING, func(p unsafe.Pointer) uintptr {
		if userFunc((*win.NMHDR)(p)) {
			return 1
		}
		return 0
	})
}

// [NM_CLICK] message handler.
//
// [NM_CLICK]: https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-tab
func (me *_TabEvents) NmClick(userFunc func() int) {
	me.events.addNfyRet(me.ctrlId, co.NM_CLICK, func(_ unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// [NM_DBLCLK] message handler.
//
// [NM_DBLCLK]: https://docs.microsoft.com/en-us/windows/win32/controls/nm-dblclk-tab
func (me *_TabEvents) NmDblClk(userFunc func() int) {
	me.events.addNfyRet(me.ctrlId, co.NM_DBLCLK, func(_ unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// [NM_RCLICK] message handler.
//
// [NM_RCLICK]: https://docs.microsoft.com/en-us/windows/win32/controls/nm-rclick-tab
func (me *_TabEvents) NmRClick(userFunc func() int) {
	me.events.addNfyRet(me.ctrlId, co.NM_RCLICK, func(_ unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// [NM_RDBLCLK] message handler.
//
// [NM_RDBLCLK]: https://docs.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-tab
func (me *_TabEvents) NmRDoubleClick(userFunc func() int) {
	me.events.addNfyRet(me.ctrlId, co.NM_RDBLCLK, func(_ unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// [NM_RELEASEDCAPTURE] message handler.
//
// [NM_RELEASEDCAPTURE]: https://docs.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-tab
func (me *_TabEvents) NmReleasedCapture(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) {
		userFunc()
	})
}
