//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Implements WindowModal interface.
type _WindowRawModal struct {
	_WindowRaw
	opts             *_WindowModalO
	hPrevFocusParent win.HWND // child control last focused on parent
}

// Creates a new WindowModal. Call WindowModalOpts() to define the options to be
// passed to the underlying CreateWindowEx().
//
// # Example
//
//	myModal := ui.NewWindowModal(
//		ui.WindowModalOpts().
//			Title("My modal window"),
//		),
//	)
func NewWindowModal(opts *_WindowModalO) WindowModal {
	if opts == nil {
		opts = WindowModalOpts()
	}
	opts.lateDefaults()

	me := &_WindowRawModal{}
	me._WindowRaw.new()
	me.opts = opts
	me.hPrevFocusParent = win.HWND(0)

	me.defaultMessages()
	return me
}

// Implements WindowModal.
func (me *_WindowRawModal) ShowModal(parent AnyParent) {
	hInst := parent.Hwnd().Hinstance()
	var wcx win.WNDCLASSEX
	me.opts.className = me._WindowRaw.generateWcx(&wcx, hInst,
		me.opts.className, me.opts.classStyles, me.opts.hCursor,
		me.opts.hBrushBkgnd, 0)
	atom := me._WindowRaw.registerClass(&wcx)

	me.hPrevFocusParent = win.GetFocus() // currently focused control in parent
	parent.Hwnd().EnableWindow(false)    // https://devblogs.microsoft.com/oldnewthing/20040227-00/?p=40463

	pos, size := me._WindowRaw.calcWndCoords(&me.opts.clientArea,
		win.HMENU(0), me.opts.wndStyles, me.opts.wndExStyles)
	me._WindowRaw.createWindow(me.opts.wndExStyles, win.ClassNameAtom(atom),
		win.StrOptSome(me.opts.title), me.opts.wndStyles, pos, size, parent.Hwnd(),
		win.HMENU(0), hInst)

	_RunModalLoop(me.Hwnd())
}

// Implements AnyParent.
func (me *_WindowRawModal) isDialog() bool {
	return false
}

func (me *_WindowRawModal) defaultMessages() {
	me.internalOn().addMsgNoRet(co.WM_SETFOCUS, func(_ wm.Any) {
		if me.Hwnd() == win.GetFocus() {
			// If window receive focus, delegate to first child.
			// This also happens right after the modal is created.
			if hFirstChild := me.Hwnd().GetNextDlgTabItem(win.HWND(0), false); hFirstChild != 0 {
				hFirstChild.SetFocus()
			}
		}
	})

	me.On().WmClose(func() {
		me.Hwnd().GetWindow(co.GW_OWNER).EnableWindow(true) // re-enable parent
		me.Hwnd().DestroyWindow()                           // then destroy modal

		if me.hPrevFocusParent != 0 {
			me.hPrevFocusParent.SetFocus() // could be on WM_DESTROY too
		}
	})
}

//------------------------------------------------------------------------------

type _WindowModalO struct {
	className   string // defined in Show()
	classStyles co.CS
	hCursor     win.HCURSOR
	hBrushBkgnd win.HBRUSH

	wndStyles   co.WS
	wndExStyles co.WS_EX
	title       string
	clientArea  win.SIZE
}

// Class name registered with RegisterClassEx().
// Defaults to a computed hash.
func (o *_WindowModalO) ClassName(n string) *_WindowModalO { o.className = n; return o }

// Window class styles, passed to RegisterClassEx().
// Defaults to CS_DBLCLKS.
func (o *_WindowModalO) ClassStyles(s co.CS) *_WindowModalO { o.classStyles = s; return o }

// Window cursor, passed to RegisterClassEx().
// Defaults to stock IDC_ARROW.
func (o *_WindowModalO) HCursor(h win.HCURSOR) *_WindowModalO { o.hCursor = h; return o }

// Window background brush, passed to RegisterClassEx().
// Defaults to COLOR_BTNFACE color.
func (o *_WindowModalO) HBrushBkgnd(h win.HBRUSH) *_WindowModalO { o.hBrushBkgnd = h; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to WS_CAPTION | WS_SYSMENU | WS_CLIPCHILDREN | WS_BORDER | WS_VISIBLE.
func (o *_WindowModalO) WndStyles(s co.WS) *_WindowModalO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
// Defaults to WS_EX_DLGMODALFRAME.
func (o *_WindowModalO) WndExStyles(s co.WS_EX) *_WindowModalO { o.wndExStyles = s; return o }

// The title of the window, passed to CreateWindowEx().
// Defaults to empty string.
func (o *_WindowModalO) Title(t string) *_WindowModalO { o.title = t; return o }

// Size of client area in pixels, passed to CreateWindowEx().
// Defaults to 400x300. Will be adjusted to the current system DPI.
func (o *_WindowModalO) ClientArea(c win.SIZE) *_WindowModalO { _OwSz(&o.clientArea, c); return o }

func (o *_WindowModalO) lateDefaults() {
	if o.hCursor == 0 {
		o.hCursor = win.HINSTANCE(0).LoadCursor(win.CursorResIdc(co.IDC_ARROW))
	}
}

// Options for NewWindowModal().
func WindowModalOpts() *_WindowModalO {
	return &_WindowModalO{
		classStyles: co.CS_DBLCLKS,
		hBrushBkgnd: win.CreateSysColorBrush(co.COLOR_BTNFACE),
		wndStyles: co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN |
			co.WS_BORDER | co.WS_VISIBLE,
		wndExStyles: co.WS_EX_DLGMODALFRAME,
		clientArea:  win.SIZE{Cx: 400, Cy: 300},
	}
}
