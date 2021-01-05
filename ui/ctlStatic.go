/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/co"
)

// Native Static control (label).
//
// https://docs.microsoft.com/en-us/windows/win32/controls/about-static-controls
type Static struct {
	*_NativeControlBase
	events *_EventsStatic
}

// Constructor. Optionally receives a control ID.
func NewStatic(parent Parent, ctrlId ...int) *Static {
	base := _NewNativeControlBase(parent, ctrlId...)
	return &Static{
		_NativeControlBase: base,
		events:             _NewEventsStatic(base),
	}
}

// Calls CreateWindowEx(). With this method, you must also specify WS and WS_EX
// window styles.
//
// Position and size will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *Static) CreateWs(
	text string, pos Pos, size Size,
	sStyles co.SS, styles co.WS, exStyles co.WS_EX) *Static {

	_global.MultiplyDpi(&pos, &size)
	return me.createNoDpi(text, pos, size, sStyles, styles, exStyles)
}

// Calls CreateWindowEx() with WS_CHILD | WS_VISIBLE.
// Size will be calculated to fit the text exactly.
//
// A typical Static has SS_LEFT | SS_NOTIFY.
//
// Position will be adjusted to the current system DPI.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *Static) Create(text string, pos Pos, sStyles co.SS) *Static {
	_global.MultiplyDpi(&pos, nil)
	size := _global.CalcTextBoundBox(text, true)
	return me.createNoDpi(text, pos, size, sStyles,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_NONE)
}

func (me *Static) createAsDlgCtrl() { me._NativeControlBase.createAssignDlg() }

// Exposes all Static notifications.
//
// Cannot be called after the parent window was created.
func (me *Static) On() *_EventsStatic {
	if me.hwnd != 0 {
		panic("Cannot add notifications after the Static was created.")
	}
	return me.events
}

// Sets the text, and resizes the control to fit it exactly.
//
// To set the text without resizing the control, use Hwnd().SetWindowText().
func (me *Static) SetText(text string) *Static {
	hasAccel := (co.SS(me.Hwnd().GetStyle()) & co.SS_NOPREFIX) == 0
	size := _global.CalcTextBoundBox(text, hasAccel)

	me.Hwnd().SetWindowPos(co.SWP_HWND_NONE,
		0, 0, int32(size.Cx), int32(size.Cy),
		co.SWP_NOZORDER|co.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
	return me
}

// Returns the text without the accelerator ampersands, for example:
// "&He && she" is returned as "He & she".
//
// Use Hwnd().GetWindowText() to retrieve the raw text, with unparsed
// accelerator ampersands.
func (me *Static) Text() string {
	return _global.RemoveAccelAmpersands(me.Hwnd().GetWindowText())
}

func (me *Static) createNoDpi(
	text string, pos Pos, size Size,
	sStyles co.SS, styles co.WS, exStyles co.WS_EX) *Static {

	me._NativeControlBase.create("STATIC", text, pos, size,
		co.WS(sStyles)|styles, exStyles)
	_global.UiFont().SetOnControl(me)
	return me
}

//------------------------------------------------------------------------------

// Static control notifications.
type _EventsStatic struct {
	ctrl *_NativeControlBase
}

// Constructor.
func _NewEventsStatic(ctrl *_NativeControlBase) *_EventsStatic {
	return &_EventsStatic{
		ctrl: ctrl,
	}
}

// https://docs.microsoft.com/en-us/windows/win32/controls/stn-clicked
func (me *_EventsStatic) StnClicked(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.STN_CLICKED), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/stn-dblclk
func (me *_EventsStatic) StnDblClk(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.STN_DBLCLK), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/stn-disable
func (me *_EventsStatic) StnDisable(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.STN_DISABLE), func(_ WmCommand) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/stn-enable
func (me *_EventsStatic) StnEnable(userFunc func()) {
	me.ctrl.parent.On().WmCommand(me.ctrl.CtrlId(), int(co.STN_ENABLE), func(_ WmCommand) {
		userFunc()
	})
}
