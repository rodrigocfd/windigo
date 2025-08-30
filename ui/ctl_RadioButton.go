//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
)

// Native [radio button] control.
//
// You cannot create this control directly, it must be created through a
// [RadioGroup].
//
// [radio button]: https://learn.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#radio-buttons
type RadioButton struct {
	_BaseCtrl
	events EventsButton
	index  uint
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
//
// Prefer using the [RadioGroup] notifications, which can handle all radio
// buttons in the group at once.
func (me *RadioButton) On() *EventsButton {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Zero-based index of this radio button within its [RadioGroup].
func (me *RadioButton) Index() uint {
	return me.index
}

// Returns true if the radio button is currently selected, with [BM_GETSTATE].
//
// [BM_GETSTATE]: https://learn.microsoft.com/en-us/windows/win32/controls/bm-getstate
func (me *RadioButton) IsSelected() bool {
	s, _ := me.hWnd.SendMessage(co.BM_GETSTATE, 0, 0)
	return (co.BST(s) & co.BST_CHECKED) != 0
}

// Selects this radio button with [BM_SETCHECK].
//
// Returns the same object, so further operations can be chained.
//
// [BM_SETCHECK]: https://learn.microsoft.com/en-us/windows/win32/controls/bm-setcheck
func (me *RadioButton) Select() *RadioButton {
	me.hWnd.SendMessage(co.BM_SETCHECK, win.WPARAM(co.BST_CHECKED), 0)
	return me
}

// Selects this radio button with [BM_SETCHECK], then sends
// a [BN_CLICKED] notification.
//
// Returns the same object, so further operations can be chained.
//
// [BM_SETCHECK]: https://learn.microsoft.com/en-us/windows/win32/controls/bm-setcheck
// [BN_CLICKED]: https://learn.microsoft.com/en-us/windows/win32/controls/bn-clicked
func (me *RadioButton) SelectAndTrigger() *RadioButton {
	me.Select()
	hParent, _ := me.hWnd.GetParent()
	hParent.SendMessage(co.WM_COMMAND,
		win.MAKEWPARAM(uint16(me.CtrlId()), uint16(co.BN_CLICKED)),
		win.LPARAM(me.Hwnd()))
	return me
}

// Sets the current text and resizes the control to exactly fit it.
//
// Returns the same object, so further operations can be chained.
func (me *RadioButton) SetTextAndResize(text string) *RadioButton {
	me.hWnd.SetWindowText(text)
	boundBox, _ := calcTextBoundBoxWithCheck(utl.RemoveAccelAmpersands(text))
	me.hWnd.SetWindowPos(win.HWND(0), 0, 0,
		uint(boundBox.Cx), uint(boundBox.Cy), co.SWP_NOZORDER|co.SWP_NOMOVE)
	return me
}

// Options for [NewRadioGroup]; returned by [OptsRadioButton].
type VarOptsRadioButton struct {
	ctrlId     uint16
	layout     LAY
	text       string
	position   win.POINT
	size       win.SIZE
	ctrlStyle  co.BS
	wndStyle   co.WS
	wndExStyle co.WS_EX
	selected   bool
}

// Options for [NewRadioGroup].
func OptsRadioButton() *VarOptsRadioButton {
	return &VarOptsRadioButton{
		ctrlStyle: co.BS_AUTORADIOBUTTON,
		wndStyle:  co.WS_CHILD | co.WS_VISIBLE,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsRadioButton) CtrlId(id uint16) *VarOptsRadioButton { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsRadioButton) Layout(l LAY) *VarOptsRadioButton { o.layout = l; return o }

// Text to be displayed, passed to [win.CreateWindowEx].
//
// Defaults to empty string.
func (o *VarOptsRadioButton) Text(t string) *VarOptsRadioButton { o.text = t; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsRadioButton) Position(x, y int) *VarOptsRadioButton {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [win.CreateWindowEx].
//
// Defaults to fit current text.
func (o *VarOptsRadioButton) Size(cx int, cy int) *VarOptsRadioButton {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Radio button control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.BS_AUTORADIOBUTTON.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/button-styles
func (o *VarOptsRadioButton) CtrlStyle(s co.BS) *VarOptsRadioButton { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
func (o *VarOptsRadioButton) WndStyle(s co.WS) *VarOptsRadioButton { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
func (o *VarOptsRadioButton) WndExStyle(s co.WS_EX) *VarOptsRadioButton { o.wndExStyle = s; return o }

// Defines this radio button as initially selected.
//
// Defaults to false.
func (o *VarOptsRadioButton) Selected(c bool) *VarOptsRadioButton { o.selected = c; return o }
