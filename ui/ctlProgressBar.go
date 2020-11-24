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
	*_NativeControlBase
	isMarquee bool
}

// Constructor. Optionally receives a control ID.
func NewProgressBar(parent Parent, ctrlId ...int) *ProgressBar {
	return &ProgressBar{
		_NativeControlBase: _NewNativeControlBase(parent, ctrlId...),
	}
}

// Calls CreateWindowEx(). With this method, you must also specify WS and WS_EX
// window styles.
//
// Position and size will be adjusted to the current system DPI.
func (me *ProgressBar) CreateWs(
	pos Pos, size Size,
	pbStyles co.PBS, styles co.WS, exStyles co.WS_EX) *ProgressBar {

	_global.MultiplyDpi(&pos, &size)
	me._NativeControlBase.create("msctls_progress32", "", pos, size,
		co.WS(pbStyles)|styles, exStyles)
	_global.UiFont().SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with WS_CHILD | WS_VISIBLE.
// Standard height is 23 pixels.
//
// A typical ProgressBar has PBS_SMOOTH.
//
// Position and width will be adjusted to the current system DPI.
func (me *ProgressBar) Create(
	pos Pos, width int, pbStyles co.PBS) *ProgressBar {

	return me.CreateWs(pos, Size{Cx: width, Cy: 23}, pbStyles,
		co.WS_CHILD|co.WS_VISIBLE,
		co.WS_EX_NONE)
}

// Retrieves the current position.
func (me *ProgressBar) Pos() int {
	return int(me.Hwnd().SendMessage(co.WM(co.PBM_GETPOS), 0, 0))
}

// Sets indeterminate state, a graphic animation going back and forth.
func (me *ProgressBar) SetMarquee(isMarquee bool) *ProgressBar {
	if isMarquee {
		me.Hwnd().SetStyle(
			me.Hwnd().GetStyle() | co.WS(co.PBS_MARQUEE))
	}

	me.Hwnd().SendMessage(co.WM(co.PBM_SETMARQUEE),
		win.WPARAM(_global.BoolToUint32(isMarquee)), 0)

	if !isMarquee {
		me.Hwnd().SetStyle(
			me.Hwnd().GetStyle() &^ co.WS(co.PBS_MARQUEE))
	}

	me.isMarquee = isMarquee
	return me
}

// Sets the current position.
func (me *ProgressBar) SetPos(pos int) *ProgressBar {
	if me.isMarquee {
		me.SetMarquee(false) // avoid crash
	}
	me.Hwnd().SendMessage(co.WM(co.PBM_SETPOS), win.WPARAM(pos), 0)
	return me
}

// Sets the new range. Default is 0-100.
func (me *ProgressBar) SetRange(min, max int) *ProgressBar {
	me.Hwnd().SendMessage(co.WM(co.PBM_SETRANGE32),
		win.WPARAM(min), win.LPARAM(max))
	return me
}

// Sets the current state (green, yellow, red) with PBM_SETSTATE.
func (me *ProgressBar) SetState(state co.PBST) *ProgressBar {
	me.Hwnd().SendMessage(co.WM(co.PBM_SETSTATE), win.WPARAM(co.PBM_SETSTATE), 0)
	return me
}
