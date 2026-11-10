//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Dialog-based control.
type _DlgControl struct {
	dlgBase _DlgBase
	ctrlId  uint16
}

// Constructor.
func newControlDlg(parent Parent, opts *VarOptsControlDlg) *_DlgControl {
	setUniqueCtrlId(&opts.ctrlId)
	me := &_DlgControl{
		dlgBase: newBaseDlg(opts.dlgId),
		ctrlId:  opts.ctrlId,
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		hInst, _ := parent.Hwnd().HInstance()
		me.dlgBase.createDialogParam(hInst, parent.Hwnd())
		me.dlgBase.base.hWnd.SetWindowLongPtr(co.GWLP_ID, uintptr(opts.ctrlId)) // give the control its ID
		if opts.ownerTab == nil {
			me.dlgBase.base.hWnd.SetWindowPos(win.HWND(0), opts.position, win.SIZE{},
				co.SWP_NOZORDER|co.SWP_NOSIZE)
		}
		parent.base().layout.Add(parent, me.dlgBase.base.hWnd, opts.layout)
	})

	me.defaultMessageHandlers(opts.ownerTab)
	return me
}

func (me *_DlgControl) defaultMessageHandlers(ownerTab *Tab) {
	me.dlgBase.base.defaultMessageHandlers()

	if ownerTab != nil { // we are a Tab container
		me.dlgBase.base.beforeUserEvents.wmCreateOrInitdialog(func() {
			ownerTab.resizeChildContainer(me.dlgBase.base.hWnd) // must run right before our childrens'
		})
	}

	me.dlgBase.base.userEvents.WmNcPaint(func(p WmNcPaint) {
		paintThemedBorders(me.dlgBase.base.hWnd, p)
	})
}

// Options for [NewControlDlg]; returned by [OptsControlDlg].
type VarOptsControlDlg struct {
	dlgId    uint16
	ctrlId   uint16
	layout   LAY
	position win.POINT

	ownerTab *Tab // if the Control is a Tab container
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
	o.position = win.POINT{X: int32(x), Y: int32(y)}
	return o
}
