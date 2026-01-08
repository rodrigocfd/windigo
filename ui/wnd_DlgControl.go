//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Dialog-based control.
type _DlgControl struct {
	_DlgBase
	ctrlId uint16
}

// Constructor.
func newControlDlg(parent Parent, opts *VarOptsControlDlg) *_DlgControl {
	setUniqueCtrlId(&opts.ctrlId)
	me := &_DlgControl{
		_DlgBase: newBaseDlg(opts.dlgId),
		ctrlId:   opts.ctrlId,
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		hInst, _ := parent.Hwnd().HInstance()
		me.createDialogParam(hInst, parent.Hwnd())
		me.hWnd.SetWindowLongPtr(co.GWLP_ID, uintptr(opts.ctrlId)) // give the control its ID
		me.hWnd.SetWindowPos(win.HWND(0),
			int(opts.position.X), int(opts.position.Y), 0, 0,
			co.SWP_NOZORDER|co.SWP_NOSIZE)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
	})

	me.defaultMessageHandlers()
	return me
}

func (me *_DlgControl) defaultMessageHandlers() {
	me.userEvents.WmNcPaint(func(p WmNcPaint) {
		paintThemedBorders(me.hWnd, p)
	})
}

// Options for [NewControlDlg]; returned by [OptsControlDlg].
type VarOptsControlDlg struct {
	dlgId    uint16
	ctrlId   uint16
	layout   LAY
	position win.POINT
}

// Options for [NewControlDlg].
func OptsControlDlg() *VarOptsControlDlg {
	return &VarOptsControlDlg{}
}

// Dialog resource ID passed to [CreateDialogParam].
//
// Panics if not informed.
//
// [CreateDialogParam]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createdialogparamw
func (o *VarOptsControlDlg) DlgId(id uint16) *VarOptsControlDlg { o.dlgId = id; return o }

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsControlDlg) CtrlId(id uint16) *VarOptsControlDlg { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_HOLD_HOLD.
func (o *VarOptsControlDlg) Layout(l LAY) *VarOptsControlDlg { o.layout = l; return o }

// Position coordinates within parent window client area.
//
// Defaults to ui.Dpi(0, 0).
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsControlDlg) Position(x, y int) *VarOptsControlDlg {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}
