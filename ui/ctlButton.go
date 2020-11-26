/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Native button control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#push-buttons
type Button struct {
	*_NativeControlBase
	events *_EventsButton
}

// Constructor. Optionally receives a control ID.
func NewButton(parent Parent, ctrlId ...int) *Button {
	base := _NewNativeControlBase(parent, ctrlId...)
	return &Button{
		_NativeControlBase: base,
		events:             _NewEventsButton(base),
	}
}

// Calls CreateWindowEx(). With this method, you must also specify WS and WS_EX
// window styles.
//
// Position and size will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *Button) CreateWs(
	text string, pos Pos, size Size,
	btnStyles co.BS, styles co.WS, exStyles co.WS_EX) *Button {

	_global.MultiplyDpi(&pos, &size)
	me._NativeControlBase.create("BUTTON", text, pos, size,
		co.WS(btnStyles)|styles, exStyles)
	_global.UiFont().SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
// Standard height is 23 pixels.
//
// A typical Button has BS_PUSHBUTTON, a default has BS_DEFPUSHBUTTON.
// For notifications beyond BN_CLICKED, use BS_NOTIFY.
//
// Position and width will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *Button) Create(
	text string, pos Pos, width int, btnStyles co.BS) *Button {

	return me.CreateWs(text, pos, Size{Cx: width, Cy: 23}, btnStyles,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_NONE)
}

func (me *Button) createAsDlgCtrl() { me._NativeControlBase.createAssignDlg() }

// Exposes all Button notifications.
//
// Cannot be called after the parent window was created.
func (me *Button) On() *_EventsButton {
	if me.hwnd != 0 {
		panic("Cannot add notifications after the Button was created.")
	}
	return me.events
}

//------------------------------------------------------------------------------

// Button control notifications.
type _EventsButton struct {
	ctrl *_NativeControlBase
}

// Constructor.
func _NewEventsButton(ctrl *_NativeControlBase) *_EventsButton {
	return &_EventsButton{
		ctrl: ctrl,
	}
}

// https://docs.microsoft.com/en-us/windows/win32/controls/bcn-dropdown
func (me *_EventsButton) BcnDropDown(userFunc func(p *win.NMBCDROPDOWN)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.BCN_DROPDOWN), func(p unsafe.Pointer) {
		userFunc((*win.NMBCDROPDOWN)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/bcn-hotitemchange
func (me *_EventsButton) BcnHotItemChange(userFunc func(p *win.NMBCHOTITEM)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.BCN_HOTITEMCHANGE), func(p unsafe.Pointer) {
		userFunc((*win.NMBCHOTITEM)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/bn-clicked
func (me *_EventsButton) BnClicked(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.BN_CLICKED), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/bn-dblclk
func (me *_EventsButton) BnDblClk(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.BN_DBLCLK), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/bn-killfocus
func (me *_EventsButton) BnKillFocus(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.BN_KILLFOCUS), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/bn-setfocus
func (me *_EventsButton) BnSetFocus(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.BN_SETFOCUS), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-customdraw-button
func (me *_EventsButton) NmCustomDraw(userFunc func(p *win.NMCUSTOMDRAW) co.CDRF) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMCUSTOMDRAW)(p)))
	})
}
