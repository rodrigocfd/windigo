/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"wingows/co"
	"wingows/win"
)

// Custom user control.
//
// Allows message and notification handling.
type WindowControl struct {
	_WindowBase
	setup _WindowSetupControl // Parameters that will be used to create the window.
}

// Retrieves the command ID for this control.
func (me *WindowControl) Id() int32 {
	return me.hwnd.GetDlgCtrlID()
}

// Exposes parameters that will be used to create the child window control.
func (me *WindowControl) Setup() *_WindowSetupControl {
	if me.Hwnd() != 0 {
		panic("Cannot change setup after the control was created.")
	}
	me.setup.initOnce() // guard
	return &me.setup
}

// Creates the child window control.
func (me *WindowControl) Create(
	parent Window, ctrlId, x, y int32, width, height uint32) {

	me.setup.initOnce() // guard
	hInst := parent.Hwnd().GetInstance()
	me._WindowBase.registerClass(me.setup.genWndclassex(hInst))

	me.defaultMessageHandling()

	x, y, width, height = _Util.MultiplyDpi(x, y, width, height)

	me._WindowBase.createWindow("WindowControl", me.setup.ExStyle,
		me.setup.ClassName, "", me.setup.Style, x, y, width, height, parent,
		win.HMENU(ctrlId), hInst)
}

func (me *WindowControl) defaultMessageHandling() {
	me.OnMsg().WmNcPaint(func(p WmNcPaint) {
		me.Hwnd().DefWindowProc(co.WM_NCPAINT, p.WParam, p.LParam) // make system draw the scrollbar for us

		if (me.Hwnd().GetExStyle()&co.WS_EX_CLIENTEDGE) == 0 || // has no border
			!win.IsThemeActive() ||
			!win.IsAppThemed() {
			// No themed borders to be painted.
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
			defer hTheme.CloseThemeData()

			// Clipping region; will draw only within this rectangle.
			// Draw only the borders to avoid flickering.
			rc2 := win.RECT{Left: rc.Left, Top: rc.Top, Right: rc.Left + 2, Bottom: rc.Bottom}
			hTheme.DrawThemeBackground(hdc, co.VS_PART_LVP_LISTGROUP, co.VS_STATE_NONE, rc, &rc2) // draw themed left border

			rc2 = win.RECT{Left: rc.Left, Top: rc.Top, Right: rc.Right, Bottom: rc.Top + 2}
			hTheme.DrawThemeBackground(hdc, co.VS_PART_LVP_LISTGROUP, co.VS_STATE_NONE, rc, &rc2) // draw themed top border

			rc2 = win.RECT{Left: rc.Right - 2, Top: rc.Top, Right: rc.Right, Bottom: rc.Bottom}
			hTheme.DrawThemeBackground(hdc, co.VS_PART_LVP_LISTGROUP, co.VS_STATE_NONE, rc, &rc2) // draw themed right border

			rc2 = win.RECT{Left: rc.Left, Top: rc.Bottom - 2, Right: rc.Right, Bottom: rc.Bottom}
			hTheme.DrawThemeBackground(hdc, co.VS_PART_LVP_LISTGROUP, co.VS_STATE_NONE, rc, &rc2) // draw themed bottom border
		}
	})
}
