package ui

import (
	"wingows/api"
	c "wingows/consts"
)

// Modal popup window.
type WindowModal struct {
	windowBase
	prevFocus api.HWND // child control last focused on parent
	setup     windowModalSetup
}

// Parameters that will be used to create the window.
func (me *WindowModal) Setup() *windowModalSetup {
	me.setup.initOnce() // guard
	return &me.setup
}

// Creates the modal window and disables the parent. This function will return
// only after the modal is closed.
func (me *WindowModal) Show(parent Window) {
	me.setup.initOnce() // guard
	hInst := parent.Hwnd().GetInstance()
	me.windowBase.registerClass(me.setup.genWndClassEx(hInst))

	me.windowBase.OnMsg().WmClose(func() { // default WM_CLOSE handling
		me.windowBase.Hwnd().GetWindow(c.GW_OWNER).EnableWindow(true) // re-enable parent
		me.windowBase.Hwnd().DestroyWindow()                          // then destroy modal
		me.prevFocus.SetFocus()
	})

	me.prevFocus = api.GetFocus()     // currently focused control
	parent.Hwnd().EnableWindow(false) // https://devblogs.microsoft.com/oldnewthing/20040227-00/?p=40463

	me.windowBase.createWindow("WindowModal", me.setup.ExStyle,
		me.setup.ClassName, me.setup.Title, me.setup.Style,
		0, 0, // initially anchored at zero
		me.setup.Width, me.setup.Height, parent, api.HMENU(0), hInst)

	rc := me.windowBase.Hwnd().GetWindowRect()
	rcParent := parent.Hwnd().GetWindowRect() // both rc relative to screen

	me.windowBase.Hwnd().SetWindowPos(c.SWP_HWND(0), // center modal over parent (warning: happens after WM_CREATE processing)
		rcParent.Left+(rcParent.Right-rcParent.Left)/2-(rc.Right-rc.Left)/2,
		rcParent.Top+(rcParent.Bottom-rcParent.Top)/2-(rc.Bottom-rc.Top)/2,
		0, 0, c.SWP_NOZORDER|c.SWP_NOSIZE)
}
