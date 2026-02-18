//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
)

// Native [tab] control.
//
// Before creating the tab, you must create an [ui.Control] to each tab content.
// They behave like ordinary child control windows, being shown/hidden by the
// tab control.
//
// [tab]: https://learn.microsoft.com/en-us/windows/win32/controls/tab-controls
type Tab struct {
	_BaseCtrl
	events   TabEvents
	children []*Control
	Items    TabItemCollection // Methods to interact with the items collection.
}

// Creates a new [Tab] with [win.CreateWindowEx].
//
// Panics if no titles are defined for the tab items.
//
// Example:
//
//	runtime.LockOSThread()
//
//	wnd := ui.NewMain(
//		ui.OptsMain().
//			Title("Hello world"),
//	)
//	tab := ui.NewTab(
//		wnd,
//		ui.OptsTab().
//			Title("First", "Second").
//			Position(ui.Dpi(10, 10)).
//			Size(ui.Dpi(200, 200)),
//	)
//	btn := ui.NewButton(
//		tab.Items.Get(0).Child(), // this button belongs to 1st tab
//		ui.OptsButton().
//			Text("Click").
//			Position(ui.Dpi(10, 10)),
//	)
//	wnd.RunAsMain()
func NewTab(parent Parent, opts *VarOptsTab) *Tab {
	if len(opts.titles) == 0 {
		panic("Cannot create a Tab control without tab items.")
	}

	setUniqueCtrlId(&opts.ctrlId)
	me := &Tab{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		children:  make([]*Control, 0, len(opts.titles)),
		events:    TabEvents{opts.ctrlId, &parent.base().userEvents},
	}
	me.Items.owner = me
	for i := 0; i < len(opts.titles); i++ {
		me.children = append(me.children,
			NewControl(parent, OptsControl().ExStyle(co.WS_EX_LEFT))) // create the Control containers
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.createWindow(opts.wndExStyle, "SysTabControl32", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, true)
		if opts.ctrlExStyle != co.TCS_EX(0) {
			me.SetExtendedStyle(true, opts.ctrlExStyle)
		}
		for _, title := range opts.titles {
			me.Items.add(title) // add each tab item
		}
		me.Items.Get(opts.selected).Select()
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
	})

	me.defaultMessageHandlers(parent)
	return me
}

// Instantiates a new [Tab] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
//
// Panics if no titles are defined for the tab items.
//
// Example:
//
//	const (
//		ID_MAIN_DLG uint16 = 1000
//		ID_TAB_FOO  uint16 = 1001
//		ID_BTN_FOO  uint16 = 1002
//	)
//
//	runtime.LockOSThread()
//
//	wnd := ui.NewMainDlg(
//		ui.OptsMainDlg().
//			DlgId(ID_MAIN_DLG),
//	)
//	tab := ui.NewTabDlg(wnd, ID_TAB_FOO, ui.LAY_HOLD_HOLD, "First", "Second")
//	btn := ui.NewButtonDlg(tab.Items.Get(0).Child(), ID_BTN_FOO, ui.LAY_HOLD_HOLD)
//	wnd.RunAsMain()
func NewTabDlg(parent Parent, ctrlId uint16, layout LAY, titles ...string) *Tab {
	if len(titles) == 0 {
		panic("Cannot create a Tab control without tab items.")
	}

	me := &Tab{
		_BaseCtrl: newBaseCtrl(ctrlId),
		children:  make([]*Control, 0, len(titles)),
		events:    TabEvents{ctrlId, &parent.base().userEvents},
	}
	me.Items.owner = me
	for i := 0; i < len(titles); i++ {
		me.children = append(me.children,
			NewControl(parent, OptsControl().ExStyle(co.WS_EX_LEFT))) // create the Control containers
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.assignDialog(parent)
		for _, title := range titles {
			me.Items.add(title) // add each tab item
		}
		me.Items.Get(0).displayContent() // 1st tab selected by default
		parent.base().layout.Add(parent, me.hWnd, layout)
	})

	me.defaultMessageHandlers(parent)
	return me
}

func (me *Tab) defaultMessageHandlers(parent Parent) {
	parent.base().beforeUserEvents.wmNotify(me.ctrlId, co.TCN_SELCHANGE, func(_ unsafe.Pointer) {
		if selTab, ok := me.Items.Selected(); ok {
			selTab.displayContent()
		}
	})
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *Tab) On() *TabEvents {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Adds or removes extended styles with [TCM_SETEXTENDEDSTYLE].
//
// Returns the same object, so further operations can be chained.
//
// [TCM_SETEXTENDEDSTYLE]: https://learn.microsoft.com/en-us/windows/win32/controls/tcm-setextendedstyle
func (me *Tab) SetExtendedStyle(doSet bool, style co.TCS_EX) *Tab {
	affected := co.TCS_EX(0)
	if doSet {
		affected = style
	}
	me.hWnd.SendMessage(co.TCM_SETEXTENDEDSTYLE,
		win.WPARAM(affected), win.LPARAM(style))
	return me
}

// Options for [NewTab]; returned by [OptsTab].
type VarOptsTab struct {
	ctrlId      uint16
	layout      LAY
	position    win.POINT
	size        win.SIZE
	ctrlStyle   co.TCS
	ctrlExStyle co.TCS_EX
	wndStyle    co.WS
	wndExStyle  co.WS_EX

	titles   []string
	selected int
}

// Options for [NewTab].
func OptsTab() *VarOptsTab {
	return &VarOptsTab{
		size:     win.SIZE{Cx: int32(DpiX(80)), Cy: int32(DpiY(50))},
		wndStyle: co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsTab) CtrlId(id uint16) *VarOptsTab { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_HOLD_HOLD.
func (o *VarOptsTab) Layout(l LAY) *VarOptsTab { o.layout = l; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsTab) Position(x, y int) *VarOptsTab {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.Dpi(80, 50).
func (o *VarOptsTab) Size(cx, cy int) *VarOptsTab {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Tab control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.TCS_NONE.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/tab-control-styles
func (o *VarOptsTab) CtrlStyle(s co.TCS) *VarOptsTab { o.ctrlStyle = s; return o }

// Tab control [extended style].
//
// Defaults to co.TCS_EX_NONE.
//
// [extended style]: https://learn.microsoft.com/en-us/windows/win32/controls/tab-control-extended-styles
func (o *VarOptsTab) CtrlExStyle(s co.TCS_EX) *VarOptsTab { o.ctrlExStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *VarOptsTab) WndStyle(s co.WS) *VarOptsTab { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
func (o *VarOptsTab) WndExStyle(s co.WS_EX) *VarOptsTab { o.wndExStyle = s; return o }

// Titles of the tab items to be created.
func (o *VarOptsTab) Titles(t ...string) *VarOptsTab { o.titles = t; return o }

// Zero-based index of the item initially selected.
//
// Defaults to 0 (first tab).
func (o *VarOptsTab) Selected(i int) *VarOptsTab { o.selected = i; return o }

// Native [tab] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [tab]: https://learn.microsoft.com/en-us/windows/win32/controls/tab-controls
type TabEvents struct {
	ctrlId       uint16
	parentEvents *WindowEvents
}

// [NM_CLICK] message handler.
//
// [NM_CLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-click-tab
func (me *TabEvents) NmClick(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CLICK, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [NM_DBLCLK] message handler.
//
// [NM_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-dblclk-tab
func (me *TabEvents) NmDblClk(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_DBLCLK, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [NM_RCLICK] message handler.
//
// [NM_RCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rclick-tab
func (me *TabEvents) NmRClick(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RCLICK, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [NM_RDBLCLK] message handler.
//
// [NM_RDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-tab
func (me *TabEvents) NmRDblClk(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RDBLCLK, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [NM_RELEASEDCAPTURE] message handler.
//
// [NM_RELEASEDCAPTURE]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-tab
func (me *TabEvents) NmReleasedCapture(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [TCN_FOCUSCHANGE] message handler.
//
// [TCN_FOCUSCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/tcn-focuschange
func (me *TabEvents) TcnFocusChange(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.TCN_FOCUSCHANGE, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [TCN_GETOBJECT] message handler.
//
// [TCN_GETOBJECT]: https://learn.microsoft.com/en-us/windows/win32/controls/tcn-getobject
func (me *TabEvents) TcnGetObject(fun func(p *win.NMOBJECTNOTIFY)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TCN_GETOBJECT, func(p unsafe.Pointer) uintptr {
		fun((*win.NMOBJECTNOTIFY)(p))
		return me.parentEvents.defProcVal
	})
}

// [TCN_KEYDOWN] message handler.
//
// [TCN_KEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/tcn-keydown
func (me *TabEvents) TcnKeyDown(fun func(p *win.NMTCKEYDOWN)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TCN_KEYDOWN, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTCKEYDOWN)(p))
		return me.parentEvents.defProcVal
	})
}

// [TCN_SELCHANGE] message handler.
//
// [TCN_SELCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/tcn-selchange
func (me *TabEvents) TcnSelChange(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.TCN_SELCHANGE, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [TCN_SELCHANGING] message handler.
//
// [TCN_SELCHANGING]: https://learn.microsoft.com/en-us/windows/win32/controls/tcn-selchanging
func (me *TabEvents) TcnSelChanging(fun func() bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TCN_SELCHANGING, func(_ unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun())
	})
}
