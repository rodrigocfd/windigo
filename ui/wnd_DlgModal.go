//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _DlgModal struct {
	_DlgBase
	parent Parent
}

func newModalDlg(parent Parent, dlgId uint16) *_DlgModal {
	me := &_DlgModal{
		_DlgBase: newBaseDlg(dlgId),
		parent:   parent,
	}
	me.defaultMessageHandlers()
	return me
}

func (me *_DlgModal) showModal() {
	hInst, _ := me.parent.Hwnd().HInstance()
	me.dialogBoxParam(hInst, me.parent.Hwnd())
}

func (me *_DlgModal) defaultMessageHandlers() {
	me._DlgBase._BaseContainer.defaultMessageHandlers()

	me.beforeUserEvents.WmInitDialog(func(_ WmInitDialog) bool {
		rcModal, _ := me.hWnd.GetWindowRect()
		rcParent, _ := me.parent.Hwnd().GetWindowRect()

		x := rcParent.Left + ((rcParent.Right - rcParent.Left) / 2) - (rcModal.Right-rcModal.Left)/2
		y := rcParent.Top + ((rcParent.Bottom - rcParent.Top) / 2) - (rcModal.Bottom-rcModal.Top)/2

		me.hWnd.SetWindowPos(win.HWND(0), int(x), int(y), 0, 0, co.SWP_NOSIZE|co.SWP_NOZORDER)

		return true // ignored
	})

	me.userEvents.WmClose(func() {
		me.hWnd.EndDialog(0)
	})
}
