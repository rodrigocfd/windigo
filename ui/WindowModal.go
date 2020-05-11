package ui

import (
	"wingows/api"
	c "wingows/consts"
)

// Modal popup window.
type WindowModal struct {
	windowBase
	prevFocus api.HWND
	Setup     windowModalSetup // Parameters that will be used to create the window.
}

func NewWindowModal() *WindowModal {
	me := WindowModal{
		windowBase: makeWindowBase(),
		prevFocus:  api.HWND(0),
		Setup:      makeWindowModalSetup(),
	}

	me.windowBase.On.WmClose(func() { // default WM_CLOSE handling
		me.windowBase.Hwnd().GetWindow(c.GW_OWNER).EnableWindow(true) // re-enable parent
		me.windowBase.Hwnd().DestroyWindow()                          // then destroy modal
		me.prevFocus.SetFocus()
	})

	return &me
}

// Creates the modal window and disables the parent. This function will return
// only after the modal is closed.
func (me *WindowModal) Show(parent Window) {
	hInst := parent.Hwnd().GetInstance()
	me.windowBase.registerClass(me.Setup.genWndClassEx(hInst))

	me.prevFocus = api.GetFocus()     // currently focused control
	parent.Hwnd().EnableWindow(false) // https://devblogs.microsoft.com/oldnewthing/20040227-00/?p=40463

	cxScreen := api.GetSystemMetrics(c.SM_CXSCREEN) // retrieve screen size
	cyScreen := api.GetSystemMetrics(c.SM_CYSCREEN)

	me.windowBase.createWindow(me.Setup.ExStyle, me.Setup.ClassName,
		me.Setup.Title, me.Setup.Style,
		cxScreen/2-int32(me.Setup.Width)/2, // center window on screen
		cyScreen/2-int32(me.Setup.Height)/2,
		me.Setup.Width, me.Setup.Height, parent, api.HMENU(0), hInst)
}

func (me *WindowModal) SetTitle(title string) *WindowModal {
	me.windowBase.Hwnd().SetWindowText(title)
	return me
}

func (me *WindowModal) Title() string {
	return me.windowBase.Hwnd().GetWindowText()
}
