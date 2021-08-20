package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native toolbar control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/toolbar-controls-overview
type Toolbar interface {
	AnyNativeControl

	// Exposes all the StatusBar notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-toolbar-control-reference-notifications
	On() *_ToolbarEvents

	AutoSize() // Sends a TB_AUTOSIZE message to resize the toolbar.

	ExtendedStyle() co.TBSTYLE_EX                                // Retrieves the extended style flags.
	Items() *_ToolbarItems                                       // Item methods.
	SetExtendedStyle(doSet bool, styles co.TBSTYLE_EX)           // Sets or unsets extended style flags.
	SetImageList(index int, himgl win.HIMAGELIST) win.HIMAGELIST // Sets the nth image list for the control.
}

//------------------------------------------------------------------------------

type _Toolbar struct {
	_NativeControlBase
	events _ToolbarEvents
	items  _ToolbarItems
}

// Creates a new Toolbar. Call ToolbarOpts() to define the options to be passed
// to the underlying CreateWindowEx().
func NewToolbar(parent AnyParent, opts *_ToolbarO) Toolbar {
	opts.lateDefaults()

	me := &_Toolbar{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)
	me.items.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		me._NativeControlBase.createWindow(opts.wndExStyles,
			"ToolbarWindow32", "", opts.wndStyles|co.WS(opts.ctrlStyles),
			win.POINT{}, win.SIZE{}, win.HMENU(opts.ctrlId))

		me.Hwnd().SendMessage(co.TB_BUTTONSTRUCTSIZE,
			win.WPARAM(unsafe.Sizeof(win.TBBUTTON{})), 0)
		me.Hwnd().SendMessage(co.CCM_SETVERSION, 5, 0)

		if opts.ctrlExStyles != co.TBSTYLE_EX_NONE {
			me.SetExtendedStyle(true, opts.ctrlExStyles)
		}
	})

	return me
}

func (me *_Toolbar) On() *_ToolbarEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the Toolbar is created.")
	}
	return &me.events
}

func (me *_Toolbar) Items() *_ToolbarItems {
	return &me.items
}

func (me *_Toolbar) AutoSize() {
	me.Hwnd().SendMessage(co.TB_AUTOSIZE, 0, 0)
}

func (me *_Toolbar) ExtendedStyle() co.TBSTYLE_EX {
	return co.TBSTYLE_EX(
		me.Hwnd().SendMessage(co.TB_GETEXTENDEDSTYLE, 0, 0),
	)
}

func (me *_Toolbar) SetExtendedStyle(doSet bool, styles co.TBSTYLE_EX) {
	curStyles := me.ExtendedStyle()
	newStyles := util.Iif(doSet, curStyles|styles, curStyles & ^styles).(co.TBSTYLE_EX)

	me.Hwnd().SendMessage(co.TB_SETEXTENDEDSTYLE, 0, win.LPARAM(newStyles))
}

func (me *_Toolbar) SetImageList(
	index int, himgl win.HIMAGELIST) win.HIMAGELIST {

	return win.HIMAGELIST(
		me.Hwnd().SendMessage(co.TB_SETIMAGELIST,
			win.WPARAM(index), win.LPARAM(himgl)),
	)
}

//------------------------------------------------------------------------------

type _ToolbarO struct {
	ctrlId int

	ctrlStyles   co.TBSTYLE
	ctrlExStyles co.TBSTYLE_EX
	wndStyles    co.WS
	wndExStyles  co.WS_EX
}

// Control ID.
// Defaults to an auto-generated ID.
func (o *_ToolbarO) CtrlId(i int) *_ToolbarO { o.ctrlId = i; return o }

// Toolbar control styles, passed to CreateWindowEx().
// Defaults to TBSTYLE_BUTTON | TBSTYLE_FLAT
func (o *_ToolbarO) CtrlStyles(s co.TBSTYLE) *_ToolbarO { o.ctrlStyles = s; return o }

// Toolbar extended control styles, passed to CreateWindowEx().
// Defaults to TBSTYLE_EX_NONE.
func (o *_ToolbarO) CtrlExStyles(s co.TBSTYLE_EX) *_ToolbarO { o.ctrlExStyles = s; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
func (o *_ToolbarO) WndStyles(s co.WS) *_ToolbarO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
// Defaults to WS_EX_NONE.
func (o *_ToolbarO) WndExStyles(s co.WS_EX) *_ToolbarO { o.wndExStyles = s; return o }

func (o *_ToolbarO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewToolbar().
func ToolbarOpts() *_ToolbarO {
	return &_ToolbarO{
		ctrlStyles: co.TBSTYLE_BUTTON | co.TBSTYLE_FLAT,
		wndStyles:  co.WS_CHILD | co.WS_VISIBLE,
	}
}

//------------------------------------------------------------------------------

// Toolbar control notifications.
type _ToolbarEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_ToolbarEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.Parent().On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-char-toolbar
func (me *_ToolbarEvents) NmChar(userFunc func(p *win.NMCHAR) bool) {
	me.events.addNfyRet(me.ctrlId, co.NM_CHAR, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMCHAR)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-toolbar
func (me *_ToolbarEvents) NmClick(userFunc func(p *win.NMMOUSE) bool) {
	me.events.addNfyRet(me.ctrlId, co.NM_CLICK, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-customdraw-toolbar
func (me *_ToolbarEvents) NmCustomDraw(userFunc func(p *win.NMTBCUSTOMDRAW) co.CDRF) {
	me.events.addNfyRet(me.ctrlId, co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMTBCUSTOMDRAW)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-dblclk-toolbar
func (me *_ToolbarEvents) NmDblClk(userFunc func(p *win.NMMOUSE) bool) {
	me.events.addNfyRet(me.ctrlId, co.NM_DBLCLK, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-keydown-toolbar
func (me *_ToolbarEvents) NmKeyDown(userFunc func(p *win.NMKEY) int) {
	me.events.addNfyRet(me.ctrlId, co.NM_KEYDOWN, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMKEY)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-ldown-toolbar
func (me *_ToolbarEvents) NmLDown(userFunc func(p *win.NMMOUSE) bool) {
	me.events.addNfyRet(me.ctrlId, co.NM_LDOWN, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-rclick-toolbar
func (me *_ToolbarEvents) NmRClick(userFunc func(p *win.NMMOUSE) bool) {
	me.events.addNfyRet(me.ctrlId, co.NM_RCLICK, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-toolbar
func (me *_ToolbarEvents) NmRDblClk(userFunc func(p *win.NMMOUSE) bool) {
	me.events.addNfyRet(me.ctrlId, co.NM_RDBLCLK, func(p unsafe.Pointer) uintptr {
		return util.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-list-view-
func (me *_ToolbarEvents) NmReleasedCapture(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) {
		userFunc()
	})
}
