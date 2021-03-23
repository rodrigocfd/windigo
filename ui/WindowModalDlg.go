package ui

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// Implements WindowModal interface.
type _WindowModalDlg struct {
	_WindowBaseDlg
}

// Creates a new WindowModal by loading a dialog resource.
func NewWindowModalDlg(dialogId int) WindowModal {
	me := _WindowModalDlg{}
	me._WindowBaseDlg.new(dialogId)

	me.defaultMessages()
	return &me
}

func (me *_WindowModalDlg) ShowModal(parent AnyParent) {
	me._WindowBaseDlg.dialogBox(parent.Hwnd(), parent.Hwnd().Hinstance())
}

func (me *_WindowModalDlg) isDialog() bool {
	return true
}

func (me *_WindowModalDlg) defaultMessages() {
	me.On().WmClose(func() {
		me.Hwnd().EndDialog(uintptr(co.ID_CANCEL))
	})
}
