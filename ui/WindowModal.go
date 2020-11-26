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

// Modal popup window.
type WindowModal struct {
	*_WindowBase
	opts            *WindowModalOpts
	prevFocusParent win.HWND // child control last focused on parent
}

// Constructor. Initializes the window with the given options.
func NewWindowModal(opts *WindowModalOpts) *WindowModal {
	me := WindowModal{
		_WindowBase: _NewWindowBase(),
		opts:        opts,
	}

	me.opts.setDefaultValues()
	me.defaultMessageHandling()
	return &me
}

// Creates the modal window and disables the parent.
// Will block until the window is closed.
func (me *WindowModal) Show(parent Parent) {
	hInst := parent.Hwnd().GetInstance()
	wcx, className := _global.GenerateWndclassex(hInst, me.opts.ClassName,
		me.opts.ClassStyles, me.opts.HCursor, me.opts.HBrushBackground,
		co.COLOR_BTNFACE, 0)
	me.opts.ClassName = className // if not specified, is auto-generated
	me._WindowBase.registerClass(wcx)

	me.prevFocusParent = win.GetFocus() // currently focused control
	parent.Hwnd().EnableWindow(false)   // https://devblogs.microsoft.com/oldnewthing/20040227-00/?p=40463

	pos, size := me.calcCoords(parent)
	me._WindowBase.createWindow("WindowModal", me.opts.ExStylesOverride,
		me.opts.ClassName, me.opts.Title, me.opts.StylesOverride,
		pos, size, parent, win.HMENU(0), hInst)

	me.runModalLoop()
}

// Adds the messages which have a default processing.
func (me *WindowModal) defaultMessageHandling() {
	me.On().WmSetFocus(func(hwndLosingFocus win.HWND) {
		if me.Hwnd() == win.GetFocus() {
			// If window receive focus, delegate to first child.
			// This also happens right after the modal is created.
			me.Hwnd().
				GetNextDlgTabItem(win.HWND(0), false).
				SetFocus()
		}
	})

	me.On().WmClose(func() {
		me.Hwnd().GetWindow(co.GW_OWNER).EnableWindow(true) // re-enable parent
		me.Hwnd().DestroyWindow()                           // then destroy modal
		me.prevFocusParent.SetFocus()                       // could be on WM_DESTROY too
	})
}

// Calculates size and position of the window to be created, based on the options.
func (me *WindowModal) calcCoords(parent Parent) (Pos, Size) {
	_global.MultiplyDpi(nil, &me.opts.ClientAreaSize) // size adjusted to DPI

	rc := win.RECT{ // left and top are zero
		Right:  int32(me.opts.ClientAreaSize.Cx),
		Bottom: int32(me.opts.ClientAreaSize.Cy),
	}
	win.AdjustWindowRectEx(&rc, me.opts.StylesOverride, false, me.opts.ExStylesOverride)
	me.opts.ClientAreaSize = Size{
		Cx: int(rc.Right - rc.Left),
		Cy: int(rc.Bottom - rc.Top),
	}

	rcParent := parent.Hwnd().GetWindowRect() // relative to screen
	return Pos{
			X: int(
				rcParent.Left + (rcParent.Right-rcParent.Left)/2 - // center on parent
					int32(me.opts.ClientAreaSize.Cx)/2,
			),
			Y: int(
				rcParent.Top + (rcParent.Bottom-rcParent.Top)/2 -
					int32(me.opts.ClientAreaSize.Cy)/2,
			),
		},
		me.opts.ClientAreaSize
}

// Runs the modal loop.
// Will block until the loop ends.
func (me *WindowModal) runModalLoop() {
	msg := win.MSG{}
	for {
		if win.GetMessage(&msg, win.HWND(0), 0, 0) == 0 {
			// WM_QUIT was sent, exit modal loop now and signal parent.
			// If it returned -1, it will simply panic.
			// https://devblogs.microsoft.com/oldnewthing/20050222-00/?p=36393
			win.PostQuitMessage(int32(msg.WParam))
			break
		}

		// If a child window, will retrieve its top-level parent.
		// If a top-level, use itself.
		if msg.HWnd.GetAncestor(co.GA_ROOT).IsDialogMessage(&msg) {
			// Processed all keyboard actions for child controls.
			if me.hwnd == win.HWND(0) {
				break // our modal was destroyed, terminate loop
			} else {
				continue
			}
		}

		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)

		if me.Hwnd() == win.HWND(0) {
			break // our modal was destroyed, terminate loop
		}
	}
}

//------------------------------------------------------------------------------

// Options for NewWindowModal().
type WindowModalOpts struct {
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
	// Defaults to COLOR_BTNFACE color.
	HBrushBackground win.HBRUSH

	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CAPTION | WS_SYSMENU | WS_CLIPCHILDREN | WS_BORDER | WS_VISIBLE.
	StylesOverride co.WS
	// Window styles, passed to CreateWindowEx().
	// These styles will be added to StylesOverride.
	// If you want to keep the default styles while adding others, use this field.
	StylesAdd co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_DLGMODALFRAME.
	ExStylesOverride co.WS_EX
	// Extended window styles, passed to CreateWindowEx().
	// These styles will be added to ExStylesOverride.
	// If you want to keep the default extended styles while adding others, use this field.
	ExStylesAdd co.WS_EX
	// The Title of the window, passed to CreateWindowEx().
	// Defaults to empty string.
	Title string
	// Size of client area, passed to CreateWindowEx().
	// Defaults to 400x300 pixels. Will be adjusted to the current system DPI.
	ClientAreaSize Size
}

func (o *WindowModalOpts) setDefaultValues() {
	if o.ClassStyles == 0 {
		o.ClassStyles = co.CS_DBLCLKS
	}
	if o.HCursor == 0 {
		o.HCursor = win.HINSTANCE(0).LoadCursor(co.IDC_ARROW)
	}
	if o.HBrushBackground == 0 {
		o.HBrushBackground = win.CreateSysColorBrush(co.COLOR_BTNFACE)
	}

	if o.StylesOverride == 0 {
		o.StylesOverride = co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN | co.WS_BORDER | co.WS_VISIBLE
	}
	o.StylesOverride |= o.StylesAdd

	if o.ExStylesOverride == 0 {
		o.ExStylesOverride = co.WS_EX_DLGMODALFRAME
	}
	o.ExStylesOverride |= o.ExStylesAdd

	if o.ClientAreaSize.Cx == 0 {
		o.ClientAreaSize.Cx = 400
	}
	if o.ClientAreaSize.Cy == 0 {
		o.ClientAreaSize.Cy = 300
	}
}
