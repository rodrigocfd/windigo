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

// Native SysLink control, which renders simple anchor markup.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/syslink-control-entry
type SysLink struct {
	*_NativeControlBase
	events *_EventsSysLink
}

// Constructor. Optionally receives a control ID.
func NewSysLink(parent Parent, ctrlId ...int) *SysLink {
	base := _NewNativeControlBase(parent, ctrlId...)
	return &SysLink{
		_NativeControlBase: base,
		events:             _NewEventsSysLink(base),
	}
}

// Calls CreateWindowEx(). With this method, you must also specify WS and WS_EX
// window styles.
//
// Position and size will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *SysLink) CreateWs(
	text string, pos Pos, size Size,
	slStyles co.LWS, styles co.WS, exStyles co.WS_EX) *SysLink {

	_global.MultiplyDpi(&pos, &size)
	me._NativeControlBase.create("SysLink", text, pos, size,
		co.WS(slStyles)|styles, exStyles)
	_global.UiFont().SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
// Size will be calculated to fit the text exactly.
//
// A typical SysLink has LWS_TRANSPARENT.
//
// Position will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *SysLink) Create(text string, pos Pos, slStyles co.LWS) *SysLink {
	me.CreateWs(text, pos, Size{}, slStyles, // zero width & height
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_NONE)

	sz := win.SIZE{}
	me.Hwnd().SendMessage(co.WM(co.LM_GETIDEALSIZE),
		0, win.LPARAM(unsafe.Pointer(&sz)))
	me.Hwnd().SetWindowPos(co.SWP_HWND_NONE, 0, 0, sz.Cx, sz.Cy,
		co.SWP_NOZORDER|co.SWP_NOMOVE)

	return me
}

func (me *SysLink) createAsDlgCtrl() { me._NativeControlBase.createAssignDlg() }

// Exposes all SysLink notifications.
//
// Cannot be called after the parent window was created.
func (me *SysLink) On() *_EventsSysLink {
	if me.hwnd != 0 {
		panic("Cannot add notifications after the SysLink was created.")
	}
	return me.events
}

// Sets the text, and resizes the control to fit it exactly.
//
// To set the text without resizing the control, use Hwnd().SetWindowText().
func (me *SysLink) SetText(text string) *SysLink {
	size := _global.CalcTextBoundBox(text, false)
	me.Hwnd().SetWindowPos(co.SWP_HWND_NONE,
		0, 0, int32(size.Cx), int32(size.Cy),
		co.SWP_NOZORDER|co.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
	return me
}

//------------------------------------------------------------------------------

// SysLink control notifications.
type _EventsSysLink struct {
	ctrl *_NativeControlBase
}

// Constructor.
func _NewEventsSysLink(ctrl *_NativeControlBase) *_EventsSysLink {
	return &_EventsSysLink{
		ctrl: ctrl,
	}
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-syslink
func (me *_EventsSysLink) NmClick(userFunc func(p *win.NMLINK)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_CLICK, func(p unsafe.Pointer) {
		userFunc((*win.NMLINK)(p))

	})
}
