//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Raw modal window.
type _RawModal struct {
	_RawBase
	parent                Parent
	opts                  *VarOptsModal
	hChildPrevFocusParent win.HWND
}

func newModalRaw(parent Parent, opts *VarOptsModal) *_RawModal {
	me := &_RawModal{
		_RawBase:              newBaseRaw(),
		parent:                parent,
		opts:                  opts,
		hChildPrevFocusParent: win.HWND(0),
	}
	me.defaultMessageHandlers()
	return me
}

func (me *_RawModal) showModal() {
	hInst, _ := me.parent.Hwnd().HInstance()
	atom := me.registerClass(hInst, me.opts.className, me.opts.classStyle,
		me.opts.classIconId, me.opts.classBrush, me.opts.classCursor)

	me.hChildPrevFocusParent = win.GetFocus()
	me.parent.Hwnd().EnableWindow(false) // https://devblogs.microsoft.com/oldnewthing/20040227-00/?p=40463

	rcWnd := win.RECT{ // client area, will be adjusted to size with title bar and borders
		Left:   0,
		Top:    0,
		Right:  me.opts.size.Cx,
		Bottom: me.opts.size.Cy,
	}
	win.AdjustWindowRectEx(&rcWnd, me.opts.style, false, me.opts.exStyle)

	rcParent, _ := me.parent.Hwnd().GetWindowRect() // relative to screen
	ptWnd := win.POINT{
		X: rcParent.Left + (rcParent.Right-rcParent.Left)/2 - me.opts.size.Cx/2, // center on parent
		Y: rcParent.Top + (rcParent.Bottom-rcParent.Top)/2 - me.opts.size.Cy/2,
	}

	me.createWindow(me.opts.exStyle, atom, me.opts.title, me.opts.style, ptWnd,
		win.SIZE{Cx: rcWnd.Right - rcWnd.Left, Cy: rcWnd.Bottom - rcWnd.Top},
		me.parent.Hwnd(), win.HMENU(0), hInst)

	processDlgMsgs := me.opts.processDlgMsgs
	me.opts = nil
	me.runModalLoop(processDlgMsgs)
}

func (me *_RawModal) defaultMessageHandlers() {
	me._RawBase._BaseContainer.defaultMessageHandlers()

	me.beforeUserEvents.wm(co.WM_SETFOCUS, func(_ Wm) {
		me.delegateFocusToFirstChild()
	})

	me.userEvents.WmClose(func() {
		hParent, _ := me.hWnd.GetWindow(co.GW_OWNER)
		hParent.EnableWindow(true) // re-enable parent
		me.hWnd.DestroyWindow()    // then destroy modal
		if me.hChildPrevFocusParent != 0 {
			me.hChildPrevFocusParent.SetFocus() // this focus could be set on WM_DESTROY as well
		}
	})
}

// Options for [NewModal]; returned by [OptsModal].
type VarOptsModal struct {
	className   string
	classStyle  co.CS
	classIconId uint16
	classCursor win.HCURSOR
	classBrush  win.HBRUSH

	title   string
	size    win.SIZE
	style   co.WS
	exStyle co.WS_EX

	processDlgMsgs bool
}

// Options for [NewModal].
func OptsModal() *VarOptsModal {
	hCursor, _ := win.HINSTANCE(0).LoadCursor(win.CursorResIdc(co.IDC_ARROW))
	return &VarOptsModal{
		classStyle:     co.CS_DBLCLKS,
		classCursor:    hCursor,
		classBrush:     win.HBRUSH(co.COLOR_BTNFACE + 1),
		style:          co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN | co.WS_BORDER | co.WS_VISIBLE,
		exStyle:        co.WS_EX_LEFT | co.WS_EX_DLGMODALFRAME,
		size:           win.SIZE{Cx: int32(DpiX(400)), Cy: int32(DpiY(200))},
		processDlgMsgs: true,
	}
}

// Class name registered with [RegisterClassEx].
//
// Defaults to a computed hash.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsModal) ClassName(s string) *VarOptsModal { o.className = s; return o }

// Window class style, passed to [RegisterClassEx].
//
// Defaults to co.CS_DBLCLKS.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsModal) ClassStyle(s co.CS) *VarOptsModal { o.classStyle = s; return o }

// ID of the resource icon to be associated to the window. The icon will be
// automatically loaded from the resource with [LoadIcon], then passed to
// [RegisterClassEx].
//
// Defaults to none.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
// [LoadIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
func (o *VarOptsModal) ClassIconId(i uint16) *VarOptsModal { o.classIconId = i; return o }

// Window cursor, passed to [RegisterClassEx].
//
// Defaults to stock co.IDC_ARROW.
//
// Example:
//
//	hCursor, _ := win.HINSTANCE(0).
//		LoadCursor(win.CursorResIdc(co.IDC_ARROW))
//
//	ui.OptsModal().
//		ClassCursor(hCursor)
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsModal) ClassCursor(h win.HCURSOR) *VarOptsModal { o.classCursor = h; return o }

// Window background brush, passed to [RegisterClassEx].
//
// Defaults to co.COLOR_BTNFACE color.
//
// Example:
//
//	classBrush: win.HBRUSH(co.COLOR_BTNFACE + 1),
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsModal) ClassBrush(h win.HBRUSH) *VarOptsModal { o.classBrush = h; return o }

// Title of the window, passed to [CreateWindowEx].
//
// Defaults to empty string.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsModal) Title(t string) *VarOptsModal { o.title = t; return o }

// Size of client area in pixels, passed to [CreateWindowEx].
//
// Defaults to ui.Dpi(400, 200).
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsModal) Size(cx, cy int) *VarOptsModal {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Window style, passed to [CreateWindowEx].
//
// Defaults to co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN | co.WS_BORDER | co.WS_VISIBLE.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsModal) Style(s co.WS) *VarOptsModal { o.style = s; return o }

// Extended window style, passed to [CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT | co.WS_EX_DLGMODALFRAME.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsModal) ExStyle(s co.WS_EX) *VarOptsModal { o.exStyle = s; return o }

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
func (o *VarOptsModal) ProcessDlgMsgs(p bool) *VarOptsModal { o.processDlgMsgs = p; return o }
