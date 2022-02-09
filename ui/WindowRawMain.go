package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Implements WindowMain interface.
type _WindowRawMain struct {
	_WindowRaw
	opts            *_WindowMainO
	hChildPrevFocus win.HWND // when window is inactivated
}

// Creates a new WindowMain. Call WindowMainOpts() to define the options to be
// passed to the underlying CreateWindowEx().
//
// Example:
//
//  myWindow := ui.NewWindowMain(
//      ui.WindowMainOpts(
//          Title("Hello world").
//          ClientArea(win.SIZE{Cx: 500, Cy: 400}).
//          WndStyles(co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN |
//              co.WS_BORDER | co.WS_VISIBLE | co.WS_MINIMIZEBOX |
//              co.WS_MAXIMIZEBOX | co.WS_SIZEBOX),
//      ),
//  )
func NewWindowMain(opts *_WindowMainO) WindowMain {
	if opts == nil {
		opts = WindowMainOpts()
	}
	opts.lateDefaults()

	me := &_WindowRawMain{}
	me._WindowRaw.new()
	me.opts = opts
	me.hChildPrevFocus = win.HWND(0)

	me.defaultMessages()
	return me
}

// Implements WindowMain.
func (me *_WindowRawMain) RunAsMain() int {
	_FirstMainStuff()
	_CreateGlobalUiFont()
	defer _globalUiFont.DeleteObject()

	hInst := win.GetModuleHandle(nil)
	var wcx win.WNDCLASSEX
	me.opts.className = me._WindowRaw.generateWcx(&wcx, hInst,
		me.opts.className, me.opts.classStyles, me.opts.hCursor,
		me.opts.hBrushBkgnd, me.opts.iconId)
	atom := me._WindowRaw.registerClass(&wcx)

	pos, size := me._WindowRaw.calcWndCoords(&me.opts.clientArea,
		me.opts.mainMenu, me.opts.wndStyles, me.opts.wndExStyles)
	me._WindowRaw.createWindow(me.opts.wndExStyles, win.ClassNameAtom(atom),
		win.StrVal(me.opts.title), me.opts.wndStyles, pos, size, win.HWND(0),
		me.opts.mainMenu, hInst)

	me.Hwnd().ShowWindow(me.opts.cmdShow)
	me.Hwnd().UpdateWindow()

	hAccel := win.HACCEL(0)
	if me.opts.accelTable != nil {
		defer me.opts.accelTable.Destroy()
		hAccel = me.opts.accelTable.Haccel()
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

type _WindowMainO struct {
	className   string // defined in RunAsMain()
	classStyles co.CS
	hCursor     win.HCURSOR
	hBrushBkgnd win.HBRUSH
	iconId      int

	wndStyles   co.WS
	wndExStyles co.WS_EX
	title       string
	clientArea  win.SIZE
	mainMenu    win.HMENU
	accelTable  AcceleratorTable

	cmdShow co.SW
}

// Class name registered with RegisterClassEx().
// Defaults to a computed hash.
func (o *_WindowMainO) ClassName(n string) *_WindowMainO { o.className = n; return o }

// Window class styles, passed to RegisterClassEx().
// Defaults to CS_DBLCLKS.
func (o *_WindowMainO) ClassStyles(s co.CS) *_WindowMainO { o.classStyles = s; return o }

// Window cursor, passed to RegisterClassEx().
// Defaults to stock IDC_ARROW.
func (o *_WindowMainO) HCursor(h win.HCURSOR) *_WindowMainO { o.hCursor = h; return o }

// Window background brush, passed to RegisterClassEx().
// Defaults to COLOR_BTNFACE color.
func (o *_WindowMainO) HBrushBkgnd(h win.HBRUSH) *_WindowMainO { o.hBrushBkgnd = h; return o }

// ID of the icon resource associated with the window, passed to RegisterClassEx().
// Defaults to none.
func (o *_WindowMainO) IconId(i int) *_WindowMainO { o.iconId = i; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to WS_CAPTION | WS_SYSMENU | WS_CLIPCHILDREN | WS_BORDER | WS_VISIBLE | WS_MINIMIZEBOX.
func (o *_WindowMainO) WndStyles(s co.WS) *_WindowMainO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
// Defaults to WS_EX_NONE.
func (o *_WindowMainO) WndExStyles(s co.WS_EX) *_WindowMainO { o.wndExStyles = s; return o }

// The title of the window, passed to CreateWindowEx().
// Defaults to empty string.
func (o *_WindowMainO) Title(t string) *_WindowMainO { o.title = t; return o }

// Size of client area in pixels, passed to CreateWindowEx().
// Defaults to 500x400. Will be adjusted to the current system DPI.
func (o *_WindowMainO) ClientArea(c win.SIZE) *_WindowMainO { _OwSz(&o.clientArea, c); return o }

// Horizontal main window menu, passed to CreateWindowEx().
// Defaults to none.
func (o *_WindowMainO) MainMenu(m win.HMENU) *_WindowMainO { o.mainMenu = m; return o }

// Accelerator table. Will be automatically destroyed.
// Defaults to none.
func (o *_WindowMainO) AccelTable(a AcceleratorTable) *_WindowMainO { o.accelTable = a; return o }

// Initial window exhibition state, passed to ShowWindow().
// Defaults to SW_SHOW.
func (o *_WindowMainO) CmdShow(c co.SW) *_WindowMainO { o.cmdShow = c; return o }

func (o *_WindowMainO) lateDefaults() {
	if o.hCursor == 0 {
		o.hCursor = win.HINSTANCE(0).LoadCursor(win.CursorResIdc(co.IDC_ARROW))
	}
}

// Options for NewWindowMain().
func WindowMainOpts() *_WindowMainO {
	return &_WindowMainO{
		classStyles: co.CS_DBLCLKS,
		hBrushBkgnd: win.CreateSysColorBrush(co.COLOR_BTNFACE),
		wndStyles: co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN |
			co.WS_BORDER | co.WS_VISIBLE | co.WS_MINIMIZEBOX,
		clientArea: win.SIZE{Cx: 500, Cy: 400},
		cmdShow:    co.SW_SHOW,
	}
}
