package ui

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// Implements WindowModal interface.
type _WindowDlgModal struct {
	_WindowDlgBase
}

// Creates a new WindowModal by loading a dialog resource.
func NewWindowModalDlg(dialogId int) WindowModal {
	me := _WindowDlgModal{}
	me._WindowDlgBase.new(dialogId)

	me.defaultMessages()
	return &me
}

func (me *_WindowDlgModal) ShowModal(parent AnyParent) {
	me._WindowDlgBase.dialogBox(parent.Hwnd(), parent.Hwnd().Hinstance())
}

func (me *_WindowDlgModal) isDialog() bool {
	return true
}

func (me *_WindowDlgModal) defaultMessages() {
	me.On().WmClose(func() {
		me.Hwnd().EndDialog(uintptr(co.ID_CANCEL))
	})
}
