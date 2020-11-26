/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/co"
	"windigo/win"
)

// Custom control window.
type WindowControl struct {
	*_WindowBase
	opts *_OptsWindowControl
}

// Constructor. Initializes the window with the given options.
func NewWindowControl(opts *_OptsWindowControl) *WindowControl {
	me := WindowControl{
		_WindowBase: _NewWindowBase(),
		opts:        opts,
	}

	me.defaultMessageHandling()
	return &me
}

// Creates the child window control, returning immediately.
func (me *WindowControl) Create(parent Parent, pos Pos, size Size) {
	hInst := parent.Hwnd().GetInstance()
	wcx, className := _global.GenerateWndclassex(hInst, me.opts.ClassName,
		me.opts.ClassStyles, me.opts.HCursor, me.opts.HBrushBackground,
		co.COLOR_WINDOW, 0)
	me.opts.ClassName = className // if not specified, is auto-generated
	me._WindowBase.registerClass(wcx)

	_global.MultiplyDpi(&pos, &size)
	me._WindowBase.createWindow("WindowControl", me.opts.ExStyles,
		me.opts.ClassName, "", me.opts.Styles, pos, size, parent,
		win.HMENU(me.opts.CtrlId), hInst)
}

// Returns the control ID.
func (me *WindowControl) CtrlId() int {
	return me.opts.CtrlId
}

// Adds the messages which have a default processing.
func (me *WindowControl) defaultMessageHandling() {
	me.On().WmNcPaint(func(p WmNcPaint) {
		me.Hwnd().DefWindowProc(co.WM_NCPAINT, p.Raw().WParam, p.Raw().LParam) // make system draw the scrollbar for us

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

		if hTheme, hasMatch := me.Hwnd().OpenThemeData("LISTVIEW"); hasMatch { // borrow style from listview
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

//------------------------------------------------------------------------------

type _OptsWindowControl struct {
	// Class name registered with RegisterClassEx().
	// Defaults to a computed hash.
	ClassName string
	// Window class styles, passed to RegisterClassEx().
	// Defaults to CS_DBLCLKS.
	ClassStyles co.CS
	// Window cursor, passed to RegisterClassEx().
	// Defaults to stock IDC_ARROW.
	HCursor win.HCURSOR
	// Window background brush, passed to RegisterClassEx().
	// Defaults to COLOR_WINDOW color.
	HBrushBackground win.HBRUSH

	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_VISIBLE | WS_CLIPCHILDREN | WS_CLIPSIBLINGS.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_CLIENTEDGE, giving the control a border.
	ExStyles co.WS_EX

	// Specific control ID. If defined, must be unique.
	// Defaults to an auto-generated number.
	CtrlId int
}

// Constructor. Returns an option set for NewWindowControl() with default values.
func DefOptsWindowControl() *_OptsWindowControl {
	return &_OptsWindowControl{
		ClassStyles:      co.CS_DBLCLKS,
		HCursor:          win.HINSTANCE(0).LoadCursor(co.IDC_ARROW),
		HBrushBackground: win.CreateSysColorBrush(co.COLOR_WINDOW),
		Styles:           co.WS_CHILD | co.WS_VISIBLE | co.WS_CLIPCHILDREN | co.WS_CLIPSIBLINGS,
		ExStyles:         co.WS_EX_CLIENTEDGE, // has a border by default
	}
}
