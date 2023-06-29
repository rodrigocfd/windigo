//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Implements WindowControl interface.
type _WindowDlgControl struct {
	_WindowDlg
	parent AnyParent
	ctrlId int
}

// Creates a new WindowControl by loading a dialog resource, with an auto-generated control ID.
//
// If parent is a dialog box, position coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
func NewWindowControlDlg(
	parent AnyParent, dialogId int,
	position win.POINT, horz HORZ, vert VERT) WindowControl {

	return _NewWindowControlDlg(
		parent, dialogId, position, 0, horz, vert)
}

// Creates a new WindowControl by loading a dialog resource, specifying a control ID.
//
// If parent is a dialog box, position coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
func NewWindowControlDlgWithId(
	parent AnyParent, dialogId int,
	position win.POINT, ctrlId int, horz HORZ, vert VERT) WindowControl {

	return _NewWindowControlDlg(
		parent, dialogId, position, ctrlId, horz, vert)
}

func _NewWindowControlDlg(
	parent AnyParent, dialogId int,
	position win.POINT, ctrlId int, horz HORZ, vert VERT) WindowControl {

	me := &_WindowDlgControl{}
	me._WindowDlg.new(dialogId)
	me.parent = parent
	me.ctrlId = 0

	if ctrlId == 0 {
		me.ctrlId = _NextCtrlId()
	}

	parent.internalOn().addMsgNoRet(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_ConvertDtuOrMultiplyDpi(parent, &position, nil)

		me._WindowDlg.createDialog(parent.Hwnd(), parent.Hwnd().Hinstance())
		me.Hwnd().SetWindowPos(win.HWND(0), position.X, position.Y, 0, 0,
			co.SWP_NOZORDER|co.SWP_NOSIZE)
		parent.addResizingChild(me, horz, vert)
	})

	me.defaultMessages()
	return me
}

// Implements AnyControl.
func (me *_WindowDlgControl) CtrlId() int {
	return me.ctrlId
}

// Implements AnyControl.
func (me *_WindowDlgControl) Parent() AnyParent {
	return me.parent
}

// Implements AnyParent.
func (me *_WindowDlgControl) isDialog() bool {
	return true
}

func (me *_WindowDlgControl) defaultMessages() {
	me.On().WmNcPaint(func(p wm.NcPaint) {
		_PaintThemedBorders(me.Hwnd(), p)
	})
}
