//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _MainRaw struct {
	_BaseRaw
	opts            *VarOptsMain
	hChildPrevFocus win.HWND
}

// Constructor.
func newMainRaw(opts *VarOptsMain) *_MainRaw {
	me := &_MainRaw{
		_BaseRaw:        newBaseRaw(),
		opts:            opts,
		hChildPrevFocus: win.HWND(0),
	}
	me.defaultMessageHandlers()
	return me
}

func (me *_MainRaw) runAsMain(hInst win.HINSTANCE) int {
	atom := me.registerClass(hInst, me.opts.className, me.opts.classStyle,
		me.opts.classIconId, me.opts.classBrush, me.opts.classCursor)

	szScreen := win.SIZE{
		Cx: win.GetSystemMetrics(co.SM_CXSCREEN),
		Cy: win.GetSystemMetrics(co.SM_CYSCREEN),
	}

	ptWnd := win.POINT{
		X: szScreen.Cx/2 - me.opts.size.Cx/2, // center on screen
		Y: szScreen.Cy/2 - me.opts.size.Cy/2,
	}

	rcWnd := win.RECT{ // client area, will be adjusted to size with title bar and borders
		Left:   ptWnd.X,
		Top:    ptWnd.Y,
		Right:  ptWnd.X + me.opts.size.Cx,
		Bottom: ptWnd.Y + me.opts.size.Cy,
	}
	win.AdjustWindowRectEx(&rcWnd, me.opts.style, me.opts.menu != 0, me.opts.exStyle)

	me.createWindow(me.opts.exStyle, atom, me.opts.title, me.opts.style,
		win.POINT{X: rcWnd.Left, Y: rcWnd.Top},
		win.SIZE{Cx: rcWnd.Right - rcWnd.Left, Cy: rcWnd.Bottom - rcWnd.Top},
		win.HWND(0), me.opts.menu, hInst)

	me.hWnd.ShowWindow(me.opts.cmdShow)
	me.hWnd.UpdateWindow()

	accelTable := me.opts.accelTable
	processDlgMsgs := me.opts.processDlgMsgs
	me.opts = nil
	return me.runMainLoop(accelTable, processDlgMsgs)
}

func (me *_MainRaw) defaultMessageHandlers() {
	me._BaseRaw._BaseContainer.defaultMessageHandlers()

	me.beforeUserEvents.WmActivate(func(p WmActivate) {
		if !p.IsMinimized() { // https://devblogs.microsoft.com/oldnewthing/20140521-00/?p=943
			if p.Event() == co.WA_INACTIVE {
				if hCurFocus := win.GetFocus(); hCurFocus != 0 && me.hWnd.IsChild(hCurFocus) {
					me.hChildPrevFocus = hCurFocus // save previously focused control
				}
			} else if me.hChildPrevFocus != 0 {
				me.hChildPrevFocus.SetFocus() // put focus back
			}
		}
	})

	me.beforeUserEvents.WmSetFocus(func(p WmSetFocus) {
		me.delegateFocusToFirstChild()
	})

	me.userEvents.WmNcDestroy(func() {
		win.PostQuitMessage(0)
	})
}

// Options for ui.NewMain(); returned by ui.OptsMain().
type VarOptsMain struct {
	className   string
	classStyle  co.CS
	classIconId uint16
	classCursor win.HCURSOR
	classBrush  win.HBRUSH

	title      string
	size       win.SIZE
	style      co.WS
	exStyle    co.WS_EX
	menu       win.HMENU
	accelTable win.HACCEL

	cmdShow        co.SW
	processDlgMsgs bool
}

// Options for ui.NewMain().
func OptsMain() *VarOptsMain {
	hCursor, _ := win.HINSTANCE(0).LoadCursor(win.CursorResIdc(co.IDC_ARROW))
	return &VarOptsMain{
		classStyle:     co.CS_DBLCLKS,
		classCursor:    hCursor,
		classBrush:     win.HBRUSH(co.COLOR_BTNFACE + 1),
		style:          co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN | co.WS_BORDER | co.WS_VISIBLE | co.WS_MINIMIZEBOX,
		size:           win.SIZE{Cx: int32(DpiX(500)), Cy: int32(DpiY(300))},
		cmdShow:        co.SW_SHOW,
		processDlgMsgs: true,
	}
}

// Class name registered with [RegisterClassEx].
//
// Defaults to a computed hash.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsMain) ClassName(s string) *VarOptsMain { o.className = s; return o }

// Window class style, passed to [RegisterClassEx].
//
// Defaults to co.CS_DBLCLKS.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsMain) ClassStyle(s co.CS) *VarOptsMain { o.classStyle = s; return o }

// Icon associated to the window, passed to [RegisterClassEx]. This icon is
// loaded from the resources with [LoadIcon], using the given resource ID.
//
// Defaults to none.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
// [LoadIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
func (o *VarOptsMain) ClassIconId(i uint16) *VarOptsMain { o.classIconId = i; return o }

// Window cursor, passed to [RegisterClassEx].
//
// Defaults to stock co.IDC_ARROW.
//
// # Example
//
//	hCursor, _ := win.HINSTANCE(0).
//		LoadCursor(win.CursorResIdc(co.IDC_ARROW))
//
//	ui.OptsMain().
//		ClassCursor(hCursor)
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsMain) ClassCursor(h win.HCURSOR) *VarOptsMain { o.classCursor = h; return o }

// Window background brush, passed to [RegisterClassEx].
//
// Defaults to co.COLOR_BTNFACE color.
//
// # Example
//
//	ui.OptsControl().
//		ClassBrush(win.HBRUSH(co.COLOR_BTNFACE + 1))
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsMain) ClassBrush(h win.HBRUSH) *VarOptsMain { o.classBrush = h; return o }

// Title of the window, passed to [CreateWindowEx].
//
// Defaults to empty string.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsMain) Title(t string) *VarOptsMain { o.title = t; return o }

// Size of client area in pixels, passed to [CreateWindowEx].
//
// Defaults to ui.Dpi(500, 400).
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsMain) Size(cx int, cy int) *VarOptsMain {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Window style, passed to [CreateWindowEx].
//
// Defaults to co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN | co.WS_BORDER | co.WS_VISIBLE | co.WS_MINIMIZEBOX.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsMain) Style(s co.WS) *VarOptsMain { o.style = s; return o }

// Extended window style, passed to [CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsMain) ExStyle(s co.WS_EX) *VarOptsMain { o.exStyle = s; return o }

// Main window menu, passed to [CreateWindowEx].
//
// Defaults to none.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsMain) Menu(m win.HMENU) *VarOptsMain { o.menu = m; return o }

// Main accelerator table to the window, passed to [CreateWindowEx].
// Defaults to none.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsMain) AccelTable(a win.HACCEL) *VarOptsMain { o.accelTable = a; return o }

// Initial window exhibition state, passed to [ShowWindow].
//
// Defaults to co.SW_SHOW.
//
// [ShowWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
func (o *VarOptsMain) CmdShow(c co.SW) *VarOptsMain { o.cmdShow = c; return o }

// In most applications, the window loop calls [IsDialogMessage] so child
// control messages will properly work. However, this has the side effect of
// inhibiting [WM_CHAR] messages from being sent to the window procedure. So,
// applications which do not have child controls and deal directly with
// character processing – like text editors – will never be able to receive
// WM_CHAR.
//
// This flag, when true, will enable the normal IsDialogMessage call in the
// window loop. When false, the call will be suppressed.
//
// Defaults to true.
//
// [IsDialogMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdialogmessagew
// [WM_CHAR]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-char
func (o *VarOptsMain) ProcessDlgMsgs(p bool) *VarOptsMain { o.processDlgMsgs = p; return o }
