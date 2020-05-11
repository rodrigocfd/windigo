package ui

import (
	"gowinui/api"
	c "gowinui/consts"
	"unsafe"
)

// Custom user control.
type WindowControl struct {
	windowBase
	ctrlId c.ID
	Setup  windowControlSetup // Parameters that will be used to create the window.
}

func NewWindowControl() *WindowControl {
	return NewWindowControlWithId(nextAutoCtrlId())
}

func NewWindowControlWithId(ctrlId c.ID) *WindowControl {
	return &WindowControl{
		windowBase: makeWindowBase(),
		ctrlId:     ctrlId,
		Setup:      makeWindowControlSetup(),
	}
}

// Returns the control ID of this child window control.
func (me *WindowControl) CtrlId() c.ID {
	return me.ctrlId
}

// Creates the child control window.
func (me *WindowControl) Create(parent Window,
	x, y int32, width, height uint32) {

	hInst := parent.Hwnd().GetInstance()
	me.windowBase.registerClass(me.Setup.genWndClassEx(hInst))

	me.windowBase.On.WmNcPaint(func(p *WmNcPaint) { // default WM_NCPAINT handling
		me.paintThemedBorders(p)
	})

	api.CreateWindowEx(me.Setup.ExStyle, // hwnd member is saved in WM_NCCREATE processing
		me.Setup.ClassName, "", me.Setup.Style, x, y, width, height,
		parent.Hwnd(), api.HMENU(me.ctrlId), hInst,
		unsafe.Pointer(&me.windowBase)) // pass pointer to windowBase object
}

func (me *WindowControl) paintThemedBorders(p *WmNcPaint) {
	me.Hwnd().DefWindowProc(c.WM_NCPAINT, api.WPARAM(p.Hrgn), 0) // make system draw the scrollbar for us

	if (me.Hwnd().GetExStyle()&c.WS_EX_CLIENTEDGE) == 0 ||
		!api.IsThemeActive() ||
		!api.IsAppThemed() {

		return
	}

	rc := me.Hwnd().GetWindowRect() // window outmost coordinates, including margins
	me.Hwnd().ScreenToClientRc(rc)
	rc.Left += 2 // manual fix, because it comes up anchored at -2,-2
	rc.Top += 2
	rc.Right += 2
	rc.Bottom += 2

	hdc := me.Hwnd().GetWindowDC()
	defer me.Hwnd().ReleaseDC(hdc)

	hTheme := me.Hwnd().OpenThemeData("LISTVIEW") // borrow style from listview
	if hTheme != 0 {
		// Clipping region; will draw only within this rectangle.
		// Draw only the borders to avoid flickering.
		rc2 := api.RECT{Left: rc.Left, Top: rc.Top, Right: rc.Left + 2, Bottom: rc.Bottom}
		hTheme.DrawThemeBackground(hdc, c.LVP_LISTGROUP, 0, rc, &rc2) // draw themed left border

		rc2 = api.RECT{Left: rc.Left, Top: rc.Top, Right: rc.Right, Bottom: rc.Top + 2}
		hTheme.DrawThemeBackground(hdc, c.LVP_LISTGROUP, 0, rc, &rc2) // draw themed top border

		rc2 = api.RECT{Left: rc.Right - 2, Top: rc.Top, Right: rc.Right, Bottom: rc.Bottom}
		hTheme.DrawThemeBackground(hdc, c.LVP_LISTGROUP, 0, rc, &rc2) // draw themed right border

		rc2 = api.RECT{Left: rc.Left, Top: rc.Bottom - 2, Right: rc.Right, Bottom: rc.Bottom}
		hTheme.DrawThemeBackground(hdc, c.LVP_LISTGROUP, 0, rc, &rc2) // draw themed bottom border

		hTheme.CloseThemeData()
	}
}
