/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package ui

import (
	"wingows/api"
	c "wingows/consts"
)

// Custom user control.
type WindowControl struct {
	windowBase
	ctrlIdGuard
	setup windowControlSetup // Parameters that will be used to create the window.
}

// Optional; returns a WindowControl with a specific control ID.
func MakeWindowControl(ctrlId c.ID) WindowControl {
	return WindowControl{
		ctrlIdGuard: makeCtrlIdGuard(ctrlId),
	}
}

// Exposes parameters that will be used to create the child window control.
func (me *WindowControl) Setup() *windowControlSetup {
	me.setup.initOnce() // guard
	return &me.setup
}

// Creates the child window control.
func (me *WindowControl) Create(parent Window, x, y int32, width, height uint32) {
	me.setup.initOnce() // guard
	hInst := parent.Hwnd().GetInstance()
	me.windowBase.registerClass(me.setup.genWndClassEx(hInst))

	me.windowBase.OnMsg().WmNcPaint(func(p WmNcPaint) { // default WM_NCPAINT handling
		me.paintThemedBorders(p.base)
	})

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.windowBase.createWindow("WindowControl", me.setup.ExStyle,
		me.setup.ClassName, "", me.setup.Style, x, y, width, height, parent,
		api.HMENU(me.ctrlIdGuard.CtrlId()), hInst)
}

func (me *WindowControl) paintThemedBorders(p wmBase) {
	me.Hwnd().DefWindowProc(c.WM_NCPAINT, p.WParam, p.LParam) // make system draw the scrollbar for us

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
