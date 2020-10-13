/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/co"
	"windigo/win"
)

// Native progress bar control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/progress-bar-control
type ProgressBar struct {
	_ControlNativeBase
	isMarquee bool
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *ProgressBar) Create(
	parent Window, ctrlId int, pos Pos, size Size,
	exStyles co.WS_EX, styles co.WS, pbStyles co.PBS) *ProgressBar {

	_Ui.MultiplyDpi(&pos, &size)
	me._ControlNativeBase.create(exStyles, "msctls_progress32", "",
		styles|co.WS(pbStyles), pos, size, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with PBS_SMOOTH.
//
// Position and size will be adjusted to the current system DPI.
func (me *ProgressBar) CreateHorizontal(
	parent Window, ctrlId int, pos Pos, size Size) *ProgressBar {

	return me.Create(parent, ctrlId, pos, size,
		co.WS_EX_NONE,
		co.WS_CHILD|co.WS_VISIBLE,
		co.PBS_SMOOTH)
}

// Calls CreateWindowEx() with PBS_SMOOTH, PBS_VERTICAL.
//
// Position and size will be adjusted to the current system DPI.
func (me *ProgressBar) CreateVertical(
	parent Window, ctrlId int, pos Pos, size Size) *ProgressBar {

	return me.Create(parent, ctrlId, pos, size,
		co.WS_EX_NONE,
		co.WS_CHILD|co.WS_VISIBLE,
		co.PBS_SMOOTH|co.PBS_VERTICAL)
}

// Retrieves the current position with PBM_GETPOS.
func (me *ProgressBar) Pos() uint {
	return uint(me.sendPbmMessage(co.PBM_GETPOS, 0, 0))
}

// Sets indeterminate state, a graphic animation going back and forth.
func (me *ProgressBar) SetMarquee(isMarquee bool) *ProgressBar {
	if isMarquee {
		me.Hwnd().SetStyle(
			me.Hwnd().GetStyle() | co.WS(co.PBS_MARQUEE))
	}

	me.sendPbmMessage(co.PBM_SETMARQUEE,
		win.WPARAM(_Ui.BoolToUint32(isMarquee)), 0)

	if !isMarquee {
		me.Hwnd().SetStyle(
			me.Hwnd().GetStyle() &^ co.WS(co.PBS_MARQUEE))
	}

	me.isMarquee = isMarquee
	return me
}

// Sets the new position with PBM_SETPOS.
func (me *ProgressBar) SetPos(pos uint) *ProgressBar {
	if me.isMarquee {
		me.SetMarquee(false) // avoid crash
	}
	me.sendPbmMessage(co.PBM_SETPOS, win.WPARAM(pos), 0)
	return me
}

// Sets the new range with PBM_SETRANGE32. Default is 0-100.
func (me *ProgressBar) SetRange(min, max uint) *ProgressBar {
	me.sendPbmMessage(co.PBM_SETRANGE32, win.WPARAM(min), win.LPARAM(max))
	return me
}

// Sets the current state (green, yellow, red) with PBM_SETSTATE.
func (me *ProgressBar) SetState(state co.PBST) *ProgressBar {
	me.sendPbmMessage(co.PBM_SETSTATE, win.WPARAM(co.PBM_SETSTATE), 0)
	return me
}

// Syntactic sugar.
func (me *ProgressBar) sendPbmMessage(msg co.PBM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.Hwnd().SendMessage(co.WM(msg), wParam, lParam)
}
