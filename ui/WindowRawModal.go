package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Implements WindowModal interface.
type _WindowRawModal struct {
	_WindowRaw
	opts             WindowModalOpts
	hPrevFocusParent win.HWND // child control last focused on parent
}

// Creates a new WindowModal specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewWindowModal(opts WindowModalOpts) WindowModal {
	opts.fillBlankValuesWithDefault()

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
	wcx := win.WNDCLASSEX{}
	me.opts.ClassName = me._WindowRaw.generateWcx(&wcx, hInst,
		me.opts.ClassName, me.opts.ClassStyles, me.opts.HCursor,
		me.opts.HBrushBackground, 0)
	me._WindowRaw.registerClass(&wcx)

	me.hPrevFocusParent = win.GetFocus() // currently focused control in parent
	parent.Hwnd().EnableWindow(false)    // https://devblogs.microsoft.com/oldnewthing/20040227-00/?p=40463

	pos, size := me._WindowRaw.calcWndCoords(&me.opts.ClientAreaSize,
		win.HMENU(0), me.opts.Styles, me.opts.ExStyles)
	me._WindowRaw.createWindow(me.opts.ExStyles, me.opts.ClassName,
		me.opts.Title, me.opts.Styles, pos, size, parent.Hwnd(),
		win.HMENU(0), hInst)

	_RunModalLoop(me.Hwnd())
}

// Implements AnyParent.
func (me *_WindowRawModal) isDialog() bool {
	return false
}

func (me *_WindowRawModal) defaultMessages() {
	me.On().WmSetFocus(func(_ wm.SetFocus) {
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
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_DLGMODALFRAME.
	ExStyles co.WS_EX
	// The title of the window, passed to CreateWindowEx().
	// Defaults to empty string.
	Title string
	// Size of client area in pixels, passed to CreateWindowEx().
	// Defaults to 400x300. Will be adjusted to the current system DPI.
	ClientAreaSize win.SIZE
}

func (opts *WindowModalOpts) fillBlankValuesWithDefault() {
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
			co.WS_BORDER | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_DLGMODALFRAME
	}

	if opts.ClientAreaSize.Cx == 0 {
		opts.ClientAreaSize.Cx = 400
	}
	if opts.ClientAreaSize.Cy == 0 {
		opts.ClientAreaSize.Cy = 300
	}
}
