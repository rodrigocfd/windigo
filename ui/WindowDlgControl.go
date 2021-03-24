package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Implements WindowControl interface.
type _WindowDlgControl struct {
	_WindowDlgBase
	parent AnyParent
	ctrlId int
}

// Creates a new WindowControl by loading a dialog resource, with an auto-generated control ID.
//
// Position will be adjusted to the current system DPI.
func NewWindowControlDlg(
	parent AnyParent, dialogId int, position win.POINT) WindowControl {

	return _NewWindowControlDlg(parent, dialogId, position, 0)
}

// Creates a new WindowControl by loading a dialog resource, specifying a control ID.
//
// Position will be adjusted to the current system DPI.
func NewWindowControlDlgWithId(
	parent AnyParent, dialogId int, position win.POINT, ctrlId int) WindowControl {

	return _NewWindowControlDlg(parent, dialogId, position, ctrlId)
}

func _NewWindowControlDlg(
	parent AnyParent, dialogId int, position win.POINT, ctrlId int) WindowControl {

	me := _WindowDlgControl{}
	me._WindowDlgBase.new(dialogId)
	me.parent = parent
	me.ctrlId = 0

	if ctrlId == 0 {
		me.ctrlId = _NextCtrlId()
	}

	parent.internalOn().addMsgZero(_ParentCreateWm(parent), func(_ wm.Any) {
		_MultiplyDpi(&position, nil)

		me._WindowDlgBase.createDialog(parent.Hwnd(), parent.Hwnd().Hinstance())
		me.Hwnd().SetWindowPos(win.HWND(0), position.X, position.Y, 0, 0,
			co.SWP_NOZORDER|co.SWP_NOSIZE)
	})

	me.defaultMessages()
	return &me
}

func (me *_WindowDlgControl) CtrlId() int {
	return me.ctrlId
}

func (me *_WindowDlgControl) Parent() AnyParent {
	return me.parent
}

func (me *_WindowDlgControl) isDialog() bool {
	return true
}

func (me *_WindowDlgControl) defaultMessages() {
	me.On().WmNcPaint(func(p wm.NcPaint) {
		_PaintThemedBorders(me.Hwnd(), p)
	})
}
