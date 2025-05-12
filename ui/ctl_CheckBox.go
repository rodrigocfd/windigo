//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [check box] control.
//
// [check box]: https://learn.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#check-boxes
type CheckBox struct {
	_BaseCtrl
	events EventsButton
}

// Creates a new [CheckBox] with [win.CreateWindowEx].
//
// # Example
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	chk := ui.NewCheckBox(
//		wndOwner,
//		ui.OptsCheckBox().
//			Text("&Click me").
//			Position(ui.Dpi(128, 75)).
//			State(co.BST_CHECKED),
//	)
func NewCheckBox(parent Parent, opts *VarOptsCheckBox) *CheckBox {
	setUniqueCtrlId(&opts.ctrlId)
	me := &CheckBox{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsButton{opts.ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.WmCreate(func(_ WmCreate) int {
		if opts.size.Cx == 0 && opts.size.Cy == 0 {
			opts.size, _ = calcTextBoundBoxWithCheck(utl.RemoveAccelAmpersands(opts.text))
		}
		me.createWindow(opts.wndExStyle, "BUTTON", opts.text,
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, true)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		me.SetState(opts.state)
		return 0 // ignored
	})

	return me
}

// Instantiates a new [CheckBox] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
//
// # Example
//
//	const ID_CHK uint16 = 0x100
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	chk := ui.NewCheckBoxDlg(
//		wndOwner, ID_CHK, ui.LAY_NONE_NONE)
func NewCheckBoxDlg(parent Parent, ctrlId uint16, layout LAY) *CheckBox {
	me := &CheckBox{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsButton{ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.WmInitDialog(func(_ WmInitDialog) bool {
		me.assignDialog(parent)
		parent.base().layout.Add(parent, me.hWnd, layout)
		return true // ignored
	})

	return me
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *CheckBox) On() *EventsButton {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Sends a [BM_GETCHECK] message and returns true if current check state is
// co.BST_CHECKED.
//
// [BM_GETCHECK]: https://learn.microsoft.com/en-us/windows/win32/controls/bm-getcheck
func (me *CheckBox) IsChecked() bool {
	return me.State() == co.BST_CHECKED
}

// Sets the current check state by sending a [BM_SETCHECK] message.
//
// A true value will apply [co.BST_CHECKED], otherwise [co.BST_UNCHECKED].
//
// Returns the same object, so further operations can be chained.
//
// [BM_SETCHECK]: https://learn.microsoft.com/en-us/windows/win32/controls/bm-setcheck
func (me *CheckBox) SetCheck(checked bool) *CheckBox {
	bst := co.BST_UNCHECKED
	if checked {
		bst = co.BST_CHECKED
	}
	return me.SetState(bst)
}

// Sets the current check state by sending a [BM_SETCHECK] message, then sends
// a [BN_CLICKED] notification.
//
// A true value will apply [co.BST_CHECKED], otherwise [co.BST_UNCHECKED].
//
// Returns the same object, so further operations can be chained.
//
// [BM_SETCHECK]: https://learn.microsoft.com/en-us/windows/win32/controls/bm-setcheck
// [BN_CLICKED]: https://learn.microsoft.com/en-us/windows/win32/controls/bn-clicked
func (me *CheckBox) SetCheckAndTrigger(checked bool) *CheckBox {
	bst := co.BST_UNCHECKED
	if checked {
		bst = co.BST_CHECKED
	}
	return me.SetStateAndTrigger(bst)
}

// Sets the current check state by sending a [BM_SETCHECK] message.
//
// Returns the same object, so further operations can be chained.
//
// [BM_SETCHECK]: https://learn.microsoft.com/en-us/windows/win32/controls/bm-setcheck
func (me *CheckBox) SetState(state co.BST) *CheckBox {
	me.hWnd.SendMessage(co.BM_SETCHECK, win.WPARAM(state), 0)
	return me
}

// Sets the current check state by sending a [BM_SETCHECK] message, then sends
// a [BN_CLICKED] notification.
//
// Returns the same object, so further operations can be chained.
//
// [BM_SETCHECK]: https://learn.microsoft.com/en-us/windows/win32/controls/bm-setcheck
// [BN_CLICKED]: https://learn.microsoft.com/en-us/windows/win32/controls/bn-clicked
func (me *CheckBox) SetStateAndTrigger(state co.BST) *CheckBox {
	me.SetState(state)
	hParent, _ := me.hWnd.GetParent()
	hParent.SendMessage(co.WM_COMMAND,
		win.MAKEWPARAM(uint16(me.CtrlId()), uint16(co.BN_CLICKED)),
		win.LPARAM(me.Hwnd()))
	return me
}

// Sets the current text and resizes the control to exactly fit it.
//
// Returns the same object, so further operations can be chained.
func (me *CheckBox) SetTextAndResize(text string) *CheckBox {
	me.hWnd.SetWindowText(text)
	boundBox, _ := calcTextBoundBoxWithCheck(utl.RemoveAccelAmpersands(text))
	me.hWnd.SetWindowPos(win.HWND(0), 0, 0,
		uint(boundBox.Cx), uint(boundBox.Cy), co.SWP_NOZORDER|co.SWP_NOMOVE)
	return me
}

// Returns the current check state by sending a [BM_GETCHECK] message.
//
// [BM_GETCHECK]: https://learn.microsoft.com/en-us/windows/win32/controls/bm-getcheck
func (me *CheckBox) State() co.BST {
	state, _ := me.hWnd.SendMessage(co.BM_GETCHECK, 0, 0)
	return co.BST(state)
}

// Calls [win.HWND.GetWindowText].
func (me *CheckBox) Text() string {
	t, _ := me.hWnd.GetWindowText()
	return t
}

// Options for [NewCheckBox]; returned by [OptsCheckBox].
type VarOptsCheckBox struct {
	ctrlId     uint16
	layout     LAY
	text       string
	position   win.POINT
	size       win.SIZE
	ctrlStyle  co.BS
	wndStyle   co.WS
	wndExStyle co.WS_EX
	state      co.BST
}

// Options for NewCheckBox].
func OptsCheckBox() *VarOptsCheckBox {
	return &VarOptsCheckBox{
		ctrlStyle: co.BS_AUTOCHECKBOX,
		wndStyle:  co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsCheckBox) CtrlId(id uint16) *VarOptsCheckBox { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsCheckBox) Layout(l LAY) *VarOptsCheckBox { o.layout = l; return o }

// Text to be displayed, passed to [win.CreateWindowEx].
//
// Defaults to empty string.
func (o *VarOptsCheckBox) Text(t string) *VarOptsCheckBox { o.text = t; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsCheckBox) Position(x, y int) *VarOptsCheckBox {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [win.CreateWindowEx].
//
// Defaults to fit current text.
func (o *VarOptsCheckBox) Size(cx int, cy int) *VarOptsCheckBox {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Check box control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.BS_AUTOCHECKBOX.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/button-styles
func (o *VarOptsCheckBox) CtrlStyle(s co.BS) *VarOptsCheckBox { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP.
func (o *VarOptsCheckBox) WndStyle(s co.WS) *VarOptsCheckBox { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
func (o *VarOptsCheckBox) WndExStyle(s co.WS_EX) *VarOptsCheckBox { o.wndExStyle = s; return o }

// Sets the initial state of the check box.
//
// Defaults to co.BST_UNCHECKED.
func (o *VarOptsCheckBox) State(s co.BST) *VarOptsCheckBox { o.state = s; return o }
