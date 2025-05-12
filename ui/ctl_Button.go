//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [button] control.
//
// [button]: https://learn.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#push-buttons
type Button struct {
	_BaseCtrl
	events EventsButton
}

// Creates a new [Button] with [win.CreateWindowEx].
//
// # Example
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	btn := ui.NewButton(
//		wndOwner,
//		ui.OptsButton().
//			Text("Click").
//			Position(ui.Dpi(20, 10)),
//	)
func NewButton(parent Parent, opts *VarOptsButton) *Button {
	setUniqueCtrlId(&opts.ctrlId)
	me := &Button{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsButton{opts.ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.WmCreate(func(_ WmCreate) int {
		me.createWindow(opts.wndExStyle, "BUTTON", opts.text,
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, true)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		return 0 // ignored
	})

	return me
}

// Instantiates a new [Button] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
//
// # Example
//
//	const ID_BTN uint16 = 0x100
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	btn := ui.NewButtonDlg(
//		wndOwner, ID_BTN, ui.LAY_NONE_NONE)
func NewButtonDlg(parent Parent, ctrlId uint16, layout LAY) *Button {
	me := &Button{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsButton{ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.WmInitDialog(func(_ WmInitDialog) bool {
		me.assignDialog(parent)
		parent.base().layout.Add(parent, me.hWnd, layout)
		return true // ignored
	})

	return me
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *Button) On() *EventsButton {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Calls [win.HWND.SetWindowText].
//
// Returns the same object, so further operations can be chained.
func (me *Button) SetText(text string) *Button {
	me.hWnd.SetWindowText(text)
	return me
}

// Calls [win.HWND.GetWindowText].
func (me *Button) Text() string {
	t, _ := me.hWnd.GetWindowText()
	return t
}

// Fires the click event by sending a [BM_CLICK] message.
//
// Returns the same object, so further operations can be chained.
//
// [BM_CLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/bm-click
func (me *Button) TriggerClick() *Button {
	me.hWnd.SendMessage(co.BM_CLICK, 0, 0)
	return me
}

// Options for [NewButton]; returned by [OptsButton].
type VarOptsButton struct {
	ctrlId     uint16
	layout     LAY
	text       string
	position   win.POINT
	size       win.SIZE
	ctrlStyle  co.BS
	wndStyle   co.WS
	wndExStyle co.WS_EX
}

// Options for [NewButton].
func OptsButton() *VarOptsButton {
	return &VarOptsButton{
		size:      win.SIZE{Cx: int32(DpiX(88)), Cy: int32(DpiY(26))},
		ctrlStyle: co.BS_PUSHBUTTON,
		wndStyle:  co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsButton) CtrlId(id uint16) *VarOptsButton { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsButton) Layout(l LAY) *VarOptsButton { o.layout = l; return o }

// Text to be displayed, passed to [win.CreateWindowEx].
//
// Defaults to empty string.
func (o *VarOptsButton) Text(t string) *VarOptsButton { o.text = t; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsButton) Position(x, y int) *VarOptsButton {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control height in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.DpiX(88).
func (o *VarOptsButton) Width(w int) *VarOptsButton { o.size.Cx = int32(w); return o }

// Control width in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.DpiY(26).
func (o *VarOptsButton) Height(h int) *VarOptsButton { o.size.Cy = int32(h); return o }

// Button control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.BS_PUSHBUTTON.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/button-styles
func (o *VarOptsButton) CtrlStyle(s co.BS) *VarOptsButton { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP.
func (o *VarOptsButton) WndStyle(s co.WS) *VarOptsButton { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
func (o *VarOptsButton) WndExStyle(s co.WS_EX) *VarOptsButton { o.wndExStyle = s; return o }

// Native [button] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [button]: https://learn.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#push-buttons
type EventsButton struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [BCN_DROPDOWN] message handler.
//
// [BCN_DROPDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/bcn-dropdown
func (me *EventsButton) BcnDropDown(fun func(p *win.NMBCDROPDOWN)) {
	me.parentEvents.WmNotify(me.ctrlId, co.BCN_DROPDOWN, func(p unsafe.Pointer) uintptr {
		fun((*win.NMBCDROPDOWN)(p))
		return me.parentEvents.defProcVal
	})
}

// [BCN_HOTITEMCHANGE] message handler.
//
// [BCN_HOTITEMCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/bcn-hotitemchange
func (me *EventsButton) BcnHotItemChange(fun func(p *win.NMBCHOTITEM)) {
	me.parentEvents.WmNotify(me.ctrlId, co.BCN_HOTITEMCHANGE, func(p unsafe.Pointer) uintptr {
		fun((*win.NMBCHOTITEM)(p))
		return me.parentEvents.defProcVal
	})
}

// [BN_CLICKED] message handler.
//
// [BN_CLICKED]: https://learn.microsoft.com/en-us/windows/win32/controls/bn-clicked
func (me *EventsButton) BnClicked(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.BN_CLICKED, fun)
}

// [BN_DBLCLK] message handler.
//
// [BN_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/bn-dblclk
func (me *EventsButton) BnDblClk(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.BN_DBLCLK, fun)
}

// [BN_KILLFOCUS] message handler.
//
// [BN_KILLFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/bn-killfocus
func (me *EventsButton) BnKillFocus(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.BN_KILLFOCUS, fun)
}

// [BN_SETFOCUS] message handler.
//
// [BN_SETFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/bn-setfocus
func (me *EventsButton) BnSetFocus(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.BN_SETFOCUS, func() {
		fun()
	})
}

// [NM_CUSTOMDRAW] message handler.
//
// [NM_CUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-customdraw-button
func (me *EventsButton) NmCustomDraw(fun func(p *win.NMCUSTOMDRAW) co.CDRF) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMCUSTOMDRAW)(p)))
	})
}
