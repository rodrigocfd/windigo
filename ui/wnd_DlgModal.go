//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Dialog-based modal window.
type _DlgModal struct {
	dlgBase _DlgBase
	parent  Parent
}

func newModalDlg(parent Parent, dlgId uint16) *_DlgModal {
	me := &_DlgModal{
		dlgBase: newBaseDlg(dlgId),
		parent:  parent,
	}
	me.defaultMessageHandlers()
	return me
}

func (me *_DlgModal) showModal() {
	hInst, _ := me.parent.Hwnd().HInstance()
	me.dlgBase.dialogBoxParam(hInst, me.parent.Hwnd())
}

func (me *_DlgModal) defaultMessageHandlers() {
	me.dlgBase.base.defaultMessageHandlers()

	me.dlgBase.base.beforeUserEvents.wmCreateOrInitdialog(func() {
		rcModal, _ := me.dlgBase.base.hWnd.GetWindowRect()
		rcParent, _ := me.parent.Hwnd().GetWindowRect()

		me.dlgBase.base.hWnd.SetWindowPos(win.HWND(0),
			win.POINT{
				X: rcParent.Left + ((rcParent.Right - rcParent.Left) / 2) - (rcModal.Right-rcModal.Left)/2,
				Y: rcParent.Top + ((rcParent.Bottom - rcParent.Top) / 2) - (rcModal.Bottom-rcModal.Top)/2,
			},
			win.SIZE{}, co.SWP_NOSIZE|co.SWP_NOZORDER)
	})

	me.dlgBase.base.userEvents.WmClose(func() {
		me.dlgBase.base.hWnd.EndDialog(0)
	})
}
