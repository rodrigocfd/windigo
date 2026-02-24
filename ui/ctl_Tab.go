//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
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
		events:    TabEvents{opts.ctrlId, &parent.base().userEvents},
		children:  make([]*Control, 0, len(opts.titles)),
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.createWindow(opts.wndExStyle, "SysTabControl32", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, true)
		if opts.ctrlExStyle != co.TCS_EX(0) {
			me.SetExtendedStyle(true, opts.ctrlExStyle)
		}
		for _, title := range opts.titles {
			me.addTab(title) // add each tab item
		}
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
	})

	me.defaultMessageHandlers(parent, opts.titles, opts.selected, opts.layout != LAY_HOLD_HOLD)
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
		events:    TabEvents{ctrlId, &parent.base().userEvents},
		children:  make([]*Control, 0, len(titles)),
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.assignDialog(parent)
		for _, title := range titles {
			me.addTab(title) // add each tab item
		}
		parent.base().layout.Add(parent, me.hWnd, layout)
	})

	me.defaultMessageHandlers(parent, titles, 0, layout != LAY_HOLD_HOLD)
	return me
}

func (me *Tab) defaultMessageHandlers(parent Parent, titles []string, selTab int, resizable bool) {
	for i := 0; i < len(titles); i++ {
		me.children = append(me.children, // create the Control containers
			NewControl(parent, OptsControl().
				ExStyle(co.WS_EX_LEFT).
				tabOwner(me)))
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.Item(selTab).Select() // must run after all containers are created and resized
	})

	parent.base().beforeUserEvents.wmNotify(me.ctrlId, co.TCN_SELCHANGE, func(_ unsafe.Pointer) {
		if selTab, ok := me.ItemSelected(); ok {
			me.displayContent(selTab.Index())
		}
	})

	if resizable { // when the Tab itself is resized, we resize the container as well
		parent.base().beforeUserEvents.wm(co.WM_SIZE, func(_ Wm) {
			if selTab, ok := me.ItemSelected(); ok {
				me.resizeChildContainer(selTab.Child().Hwnd())
			}
		})
	}
}

func (me *Tab) addTab(title string) TabItem {
	tci := win.TCITEM{
		Mask: co.TCIF_TEXT,
	}

	var wBuf wstr.BufEncoder
	tci.SetPszText(wBuf.Slice(title))

	newIdxRet, err := me.Hwnd().SendMessage(co.TCM_INSERTITEM,
		0x0fff_ffff, win.LPARAM(unsafe.Pointer(&tci)))
	newIdx := int(newIdxRet)
	if err != nil || newIdx == -1 {
		panic(fmt.Sprintf("TCM_INSERTITEM \"%s\" failed.", title))
	}

	return me.Item(newIdx)
}

func (me *Tab) resizeChildContainer(hChild win.HWND) {
	rcTab, _ := me.Hwnd().GetWindowRect()
	hParent, _ := me.Hwnd().GetParent()
	hParent.ScreenToClientRc(&rcTab)
	me.Hwnd().SendMessage(co.TCM_ADJUSTRECT, 0, win.LPARAM(unsafe.Pointer(&rcTab))) // ideal child size
	hChild.SetWindowPos(win.HWND(0),
		win.POINT{X: rcTab.Left, Y: rcTab.Top},
		win.SIZE{Cx: rcTab.Right - rcTab.Left, Cy: rcTab.Bottom - rcTab.Top},
		co.SWP_NOZORDER|co.SWP_SHOWWINDOW)
}

func (me *Tab) displayContent(index int) {
	if len(me.children) == 0 {
		return
	}
	for idxChild, child := range me.children {
		if idxChild != index {
			child.Hwnd().ShowWindow(co.SW_HIDE) // hide all others
		}
	}
	me.resizeChildContainer(me.children[index].Hwnd())
	me.Hwnd().SetWindowPos(me.children[index].Hwnd(),
		win.POINT{}, win.SIZE{}, co.SWP_NOSIZE|co.SWP_NOMOVE) // container above the Tab
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *Tab) On() *TabEvents {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Returns the item at the given index.
//
// A negative index will give you an invalid column.
func (me *Tab) Item(index int) TabItem {
	return TabItem{me, int32(index)}
}

// Retrieves the number of items with [TCM_GETITEMCOUNT].
//
// Panics on error.
//
// [TCM_GETITEMCOUNT]: https://learn.microsoft.com/en-us/windows/win32/controls/tcm-getitemcount
func (me *Tab) ItemCount() int {
	countRet, err := me.Hwnd().SendMessage(co.TCM_GETITEMCOUNT, 0, 0)
	count := int(countRet)
	if err != nil || count == -1 {
		panic("TCM_GETITEMCOUNT failed.")
	}
	return count
}

// Retrieves the focused item with [TCM_GETCURFOCUS], if any
//
// [TCM_GETCURFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/tcm-getcurfocus
func (me *Tab) ItemFocused() (TabItem, bool) {
	idxRet, _ := me.Hwnd().SendMessage(co.TCM_GETCURFOCUS, 0, 0)
	idx := int(idxRet)
	if idx == -1 {
		return TabItem{}, false
	}
	return me.Item(idx), true
}

// Retrieves the selected item with [TCM_GETCURSEL], if any
//
// [TCM_GETCURSEL]: https://learn.microsoft.com/en-us/windows/win32/controls/tcm-getcursel
func (me *Tab) ItemSelected() (TabItem, bool) {
	idxRet, _ := me.Hwnd().SendMessage(co.TCM_GETCURSEL, 0, 0)
	idx := int(idxRet)
	if idx == -1 {
		return TabItem{}, false
	}
	return me.Item(idx), true
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
