package ui

import (
	"gowinui/api"
	c "gowinui/consts"
	"unsafe"
)

// Modal popup window.
type WindowModal struct {
	windowBase
	prevFocus api.HWND
	Setup     windowModalSetup // Parameters that will be used to create the window.
}

func NewWindowModal() *WindowModal {
	return &WindowModal{
		windowBase: makeWindowBase(),
		prevFocus:  api.HWND(0),
		Setup:      makeWindowModalSetup(),
	}
}

// Creates the modal window and disables the parent. This function will return
// only after the modal is closed.
func (me *WindowModal) Show(parent Window) {
	hInst := parent.Hwnd().GetInstance()
	me.windowBase.registerClass(me.Setup.genWndClassEx(hInst))

	me.prevFocus = api.GetFocus()     // currently focused control
	parent.Hwnd().EnableWindow(false) // https://devblogs.microsoft.com/oldnewthing/20040227-00/?p=40463

	me.windowBase.On.WmClose(func() { // default WM_CLOSE handling
		me.windowBase.Hwnd().GetWindow(c.GW_OWNER).EnableWindow(true) // re-enable parent
		me.windowBase.Hwnd().DestroyWindow()                          // then destroy modal
		me.prevFocus.SetFocus()
	})

	cxScreen := api.GetSystemMetrics(c.SM_CXSCREEN)
	cyScreen := api.GetSystemMetrics(c.SM_CYSCREEN)

	api.CreateWindowEx(me.Setup.ExStyle, // hwnd member is saved in WM_NCCREATE processing
		me.Setup.ClassName, me.Setup.Title, me.Setup.Style,
		cxScreen/2-int32(me.Setup.Width)/2, // center window on screen
		cyScreen/2-int32(me.Setup.Height)/2,
		me.Setup.Width, me.Setup.Height,
		parent.Hwnd(), api.HMENU(0), hInst,
		unsafe.Pointer(&me.windowBase)) // pass pointer to windowBase object
}

func (me *WindowModal) SetTitle(title string) *WindowModal {
	me.windowBase.Hwnd().SetWindowText(title)
	return me
}

func (me *WindowModal) Title() string {
	return me.windowBase.Hwnd().GetWindowText()
}
