package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Implements WindowMain interface.
type _WindowRawMain struct {
	_WindowRaw
	opts            WindowMainRawOpts
	hChildPrevFocus win.HWND // when window is inactivated
}

// Creates a new WindowMain specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewWindowMainRaw(opts WindowMainRawOpts) WindowMain {
	opts.fillBlankValuesWithDefault()

	me := _WindowRawMain{}
	me._WindowRaw.new()
	me.opts = opts
	me.hChildPrevFocus = win.HWND(0)

	me.defaultMessages()
	return &me
}

// Implements WindowMain.
func (me *_WindowRawMain) RunAsMain() int {
	if win.IsWindowsVistaOrGreater() {
		win.SetProcessDPIAware()
	}
	win.InitCommonControls()
	_CreateGlobalUiFont()
	defer _globalUiFont.DeleteObject()

	hInst := win.GetModuleHandle("")
	wcx := win.WNDCLASSEX{}
	me.opts.ClassName = me._WindowRaw.generateWcx(&wcx, hInst,
		me.opts.ClassName, me.opts.ClassStyles, me.opts.HCursor,
		me.opts.HBrushBackground, me.opts.IconId)
	me._WindowRaw.registerClass(&wcx)

	pos, size := me._WindowRaw.calcWndCoords(&me.opts.ClientAreaSize,
		me.opts.MainMenu, me.opts.Styles, me.opts.ExStyles)
	me._WindowRaw.createWindow(me.opts.ExStyles, me.opts.ClassName,
		me.opts.Title, me.opts.Styles, pos, size, win.HWND(0),
		me.opts.MainMenu, hInst)

	me.Hwnd().ShowWindow(me.opts.CmdShow)
	me.Hwnd().UpdateWindow()

	hAccel := win.HACCEL(0)
	if me.opts.AccelTable != nil {
		defer me.opts.AccelTable.Destroy()
		hAccel = me.opts.AccelTable.Haccel()
	}

	return _RunMainLoop(me.Hwnd(), hAccel)
}

// Implements AnyParent.
func (me *_WindowRawMain) isDialog() bool {
	return false
}

func (me *_WindowRawMain) defaultMessages() {
	me.On().WmNcDestroy(func() {
		win.PostQuitMessage(0)
	})

	me.On().WmSetFocus(func(_ wm.SetFocus) {
		if me.Hwnd() == win.GetFocus() {
			// If window receives focus, delegate to first child.
			if hFirstChild := me.Hwnd().GetNextDlgTabItem(win.HWND(0), false); hFirstChild != 0 {
				hFirstChild.SetFocus()
			}
		}
	})

	me.On().WmActivate(func(p wm.Activate) {
		// https://devblogs.microsoft.com/oldnewthing/20140521-00/?p=943
		if !p.IsMinimized() {
			if p.Event() == co.WA_INACTIVE {
				if hCurFocus := win.GetFocus(); hCurFocus != 0 && me.Hwnd().IsChild(hCurFocus) {
					me.hChildPrevFocus = hCurFocus // save previously focused control
				}
			} else if me.hChildPrevFocus != 0 {
				me.hChildPrevFocus.SetFocus() // put focus back
			}
		}
	})
}

//------------------------------------------------------------------------------

// Options for NewWindowMainRaw().
type WindowMainRawOpts struct {
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
	// Defaults to WS_CAPTION | WS_SYSMENU | WS_CLIPCHILDREN | WS_BORDER | WS_VISIBLE | WS_MINIMIZEBOX.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStyles co.WS_EX
	// The title of the window, passed to CreateWindowEx().
	// Defaults to empty string.
	Title string
	// Size of client area in pixels, passed to CreateWindowEx().
	// Defaults to 500x400. Will be adjusted to the current system DPI.
	ClientAreaSize win.SIZE
	// Horizontal main window menu, passed to CreateWindowEx().
	// Defaults to none.
	MainMenu win.HMENU
	// Accelerator table. Will be automatically destroyed.
	// Defaults to none.
	AccelTable AcceleratorTable

	// Initial window exhibition state, passed to ShowWindow().
	// Defaults to SW_SHOW.
	CmdShow co.SW
}

func (opts *WindowMainRawOpts) fillBlankValuesWithDefault() {
	if opts.ClassStyles == 0 {
		opts.ClassStyles = co.CS_DBLCLKS
	}
	if opts.HCursor == 0 {
		opts.HCursor = win.HINSTANCE(0).LoadCursor(co.IDC_ARROW)
	}
	if opts.HBrushBackground == 0 {
		opts.HBrushBackground = win.CreateSysColorBrush(co.COLOR_BTNFACE)
	}

	if opts.Styles == 0 {
		opts.Styles = co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN |
			co.WS_BORDER | co.WS_VISIBLE | co.WS_MINIMIZEBOX
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_NONE
	}

	if opts.ClientAreaSize.Cx == 0 {
		opts.ClientAreaSize.Cx = 500
	}
	if opts.ClientAreaSize.Cy == 0 {
		opts.ClientAreaSize.Cy = 400
	}

	if opts.CmdShow == 0 { // note that SW_HIDE (zero) is not supported
		opts.CmdShow = co.SW_SHOW
	}
}
