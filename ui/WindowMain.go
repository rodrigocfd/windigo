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

// Main application window.
type WindowMain struct {
	*_WindowBase
	opts           *WindowMainOpts
	mainLoop       *_WindowMainLoop
	childPrevFocus win.HWND // when window is inactivated
	mainMenu       *Menu
	accelTable     AccelTable
}

// Constructor. Initializes the window with the given options.
func NewWindowMain(opts *WindowMainOpts) *WindowMain {
	me := WindowMain{
		_WindowBase: _NewWindowBase(),
		opts:        opts,
		mainLoop:    _NewWindowMainLoop(),
		mainMenu:    &Menu{hMenu: win.CreateMenu()}, // passed to CreateWindowEx()
	}

	me.opts.setDefaultValues()
	me.defaultMessageHandling()
	return &me
}

// Returns the horizontal main window menu.
// Automatically destroyed when the window is closed.
func (me *WindowMain) MainMenu() *Menu {
	return me.mainMenu
}

// Returns the accelerator table, which keeps the keyboard shortcut combinations.
// Automatically destroyed when the window is closed.
func (me *WindowMain) AccelTable() *AccelTable {
	return &me.accelTable
}

// Creates the main window and runs the main application loop.
// Will block until the window is closed.
func (me *WindowMain) RunAsMain() int {
	defer func() {
		// Recover from any panic within GUI thread.
		// Panics in other threads can't be recovered.
		if r := recover(); r != nil {
			_global.TreatPanic(r)
		}
		me.accelTable.Destroy() // free resources
		_global.uiFont.Destroy()
	}()

	if win.IsWindowsVistaOrGreater() {
		win.SetProcessDPIAware()
	}
	win.InitCommonControls()

	hInst := win.GetModuleHandle("")
	wcx, className := _global.GenerateWndclassex(hInst, me.opts.ClassName,
		me.opts.ClassStyles, me.opts.HCursor, me.opts.HBrushBackground,
		co.COLOR_BTNFACE, me.opts.IconId)
	me.opts.ClassName = className // if not specified, is auto-generated
	me._WindowBase.registerClass(wcx)

	pos, size := me.calcCoords()
	me._WindowBase.createWindow("WindowMain", me.opts.ExStylesOverride,
		me.opts.ClassName, me.opts.Title, me.opts.StylesOverride,
		pos, size, nil, me.mainMenu.Hmenu(), hInst)

	me.Hwnd().ShowWindow(me.opts.CmdShow)
	me.Hwnd().UpdateWindow()
	return me.mainLoop.RunLoop(me.Hwnd(), me.accelTable.Haccel())
}

// Adds the messages which have a default processing.
func (me *WindowMain) defaultMessageHandling() {
	me.On().WmNcDestroy(func() {
		win.PostQuitMessage(0)
	})

	me.On().WmSetFocus(func(hwndLosingFocus win.HWND) {
		if me.Hwnd() == win.GetFocus() {
			// If window receives focus, delegate to first child.
			me.Hwnd().
				GetNextDlgTabItem(win.HWND(0), false).
				SetFocus()
		}
	})

	me.On().WmActivate(func(p WmActivate) {
		// https://devblogs.microsoft.com/oldnewthing/20140521-00/?p=943
		if !p.IsMinimized() {
			if p.Event() == co.WA_INACTIVE {
				curFocus := win.GetFocus()
				if curFocus != 0 && me.Hwnd().IsChild(curFocus) {
					me.childPrevFocus = curFocus // save previously focused control
				}
			} else if me.childPrevFocus != 0 {
				me.childPrevFocus.SetFocus() // put focus back
			}
		}
	})
}

// Calculates size and position of the window to be created, based on the options.
func (me *WindowMain) calcCoords() (Pos, Size) {
	screenSize := Size{
		Cx: int(win.GetSystemMetrics(co.SM_CXSCREEN)),
		Cy: int(win.GetSystemMetrics(co.SM_CYSCREEN)),
	}

	_global.MultiplyDpi(nil, &me.opts.ClientAreaSize) // size adjusted to DPI

	pos := Pos{
		X: int(screenSize.Cx/2 - me.opts.ClientAreaSize.Cx/2), // center on screen
		Y: int(screenSize.Cy/2 - me.opts.ClientAreaSize.Cy/2),
	}

	rc := win.RECT{
		Left:   int32(pos.X),
		Top:    int32(pos.Y),
		Right:  int32(int(me.opts.ClientAreaSize.Cx) + pos.X),
		Bottom: int32(int(me.opts.ClientAreaSize.Cy) + pos.Y),
	}
	win.AdjustWindowRectEx(&rc, me.opts.StylesOverride,
		me.mainMenu.ItemCount() > 0, me.opts.ExStylesOverride)
	me.opts.ClientAreaSize = Size{
		Cx: int(rc.Right - rc.Left),
		Cy: int(rc.Bottom - rc.Top),
	}

	return Pos{int(rc.Left), int(rc.Top)},
		me.opts.ClientAreaSize
}

//------------------------------------------------------------------------------

// Options for NewWindowMain().
type WindowMainOpts struct {
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
	// ID of the icon resource associated with the window, passed to RegisterClassEx().
	// Defaults to none.
	IconId int

	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CAPTION | WS_SYSMENU | WS_CLIPCHILDREN | WS_BORDER | WS_VISIBLE.
	StylesOverride co.WS
	// Window styles, passed to CreateWindowEx().
	// These styles will be added to StylesOverride.
	// If you want to keep the default styles while adding others, use this field.
	StylesAdd co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStylesOverride co.WS_EX
	// Extended window styles, passed to CreateWindowEx().
	// These styles will be added to ExStylesOverride.
	// If you want to keep the default extended styles while adding others, use this field.
	ExStylesAdd co.WS_EX
	// The title of the window, passed to CreateWindowEx().
	// Defaults to empty string.
	Title string
	// Size of client area, passed to CreateWindowEx().
	// Defaults to 500x400 pixels. Will be adjusted to the current system DPI.
	ClientAreaSize Size

	// Initial window exhibition state, passed to ShowWindow().
	// Defaults to SW_SHOW.
	CmdShow co.SW
}

func (o *WindowMainOpts) setDefaultValues() {
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

	o.ExStylesOverride |= o.ExStylesAdd

	if o.ClientAreaSize.Cx == 0 {
		o.ClientAreaSize.Cx = 500
	}
	if o.ClientAreaSize.Cy == 0 {
		o.ClientAreaSize.Cy = 400
	}

	if o.CmdShow == 0 { // note that SW_HIDE (zero) is not supported
		o.CmdShow = co.SW_SHOW
	}
}
