//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Custom control.
//
// Implements:
//   - [Window]
//   - [ChildControl]
//   - [Parent]
type Control struct {
	raw    *_RawControl
	dlg    *_DlgControl
	parent Parent
}

// Creates a new custom control with [CreateWindowEx].
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func NewControl(parent Parent, opts *VarOptsControl) *Control {
	return &Control{
		raw:    newControlRaw(parent, opts),
		parent: parent,
	}
}

// Creates a new dialog-based custom control with [CreateDialogParam].
//
// [CreateDialogParam]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createdialogparamw
func NewControlDlg(parent Parent, opts *VarOptsControlDlg) *Control {
	return &Control{
		dlg:    newControlDlg(parent, opts),
		parent: parent,
	}
}

// Returns the parent container of this control.
func (me *Control) Parent() Parent {
	return me.parent
}

// Returns the underlying HWND handle of this window.
//
// Implements [Window].
//
// Note that this handle is initially zero, existing only after window creation.
func (me *Control) Hwnd() win.HWND {
	if me.raw != nil {
		return me.raw.hWnd
	} else {
		return me.dlg.hWnd
	}
}

// Returns the control ID, unique within the same Parent.
//
// Implements [ChildControl].
func (me *Control) CtrlId() uint16 {
	if me.raw != nil {
		return me.raw.ctrlId
	} else {
		return me.dlg.ctrlId
	}
}

// If parent is a dialog, sets the focus by sending [WM_NEXTDLGCTL]. This
// draws the borders correctly in some undefined controls, like buttons.
// Otherwise, calls [SetFocus].
//
// Implements [ChildControl].
//
// [WM_NEXTDLGCTL]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-nextdlgctl
// [SetFocus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setfocus
func (me *Control) Focus() {
	hParent, _ := me.Hwnd().GetAncestor(co.GA_PARENT)
	isDialog, _ := hParent.IsDialog()
	if isDialog {
		hParent.SendMessage(co.WM_NEXTDLGCTL, win.WPARAM(me.Hwnd()), 1)
	} else {
		me.Hwnd().SetFocus()
	}
}

// Exposes all the window notifications the can be handled.
//
// Implements [Parent].
//
// Panics if called after the window has been created.
func (me *Control) On() *EventsWindow {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the window has been created.")
	}

	if me.raw != nil {
		return &me.raw.userEvents
	} else {
		return &me.dlg.userEvents
	}
}

// This method is analog to [SendMessage] (synchronous), but intended to be
// called from another thread, so a callback function can, tunelled by
// [WNDPROC], run in the original thread of the window, thus allowing GUI
// updates. With this, the user doesn't have to deal with a custom WM_ message.
//
// Implements [Parent].
//
// Example:
//
//	var wnd *ui.WindowModal // initialized somewhere
//
//	wnd.On().WmCreate(func(p WmCreate) int {
//		go func() {
//			// process to be done in a parallel goroutine...
//
//			wnd.UiThread(func() {
//				// update the UI in the original UI thread...
//			})
//		}()
//		return 0
//	})
//
// [SendMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew
// [WNDPROC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-wndproc
func (me *Control) UiThread(fun func()) {
	if me.raw != nil {
		me.raw.uiThread(fun)
	} else {
		me.dlg.uiThread(fun)
	}
}

// Implements [Parent].
func (me *Control) base() *_BaseContainer {
	if me.raw != nil {
		return &me.raw._BaseContainer
	} else {
		return &me.dlg._BaseContainer
	}
}
