package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Implements WindowControl interface.
type _WindowRawControl struct {
	_WindowRaw
	opts   WindowControlRawOpts
	parent AnyParent
}

// Creates a new WindowControl specifying all options, which will be passed to
// the underlying CreateWindowEx().
func NewWindowControlRaw(parent AnyParent, opts WindowControlRawOpts) WindowControl {
	opts.fillBlankValuesWithDefault()

	me := &_WindowRawControl{}
	me._WindowRaw.new()
	me.opts = opts
	me.parent = parent

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		hInst := parent.Hwnd().Hinstance()
		wcx := win.WNDCLASSEX{}
		me.opts.ClassName = me._WindowRaw.generateWcx(&wcx, hInst,
			me.opts.ClassName, me.opts.ClassStyles, me.opts.HCursor,
			me.opts.HBrushBackground, 0)
		me._WindowRaw.registerClass(&wcx)

		_MultiplyDpi(&me.opts.Position, &me.opts.Size)
		me._WindowRaw.createWindow(me.opts.ExStyles, me.opts.ClassName,
			"", me.opts.Styles, me.opts.Position, me.opts.Size, parent.Hwnd(),
			win.HMENU(me.opts.CtrlId), hInst)
	})

	me.defaultMessages()
	return me
}

// Implements AnyControl.
func (me *_WindowRawControl) CtrlId() int {
	return me.opts.CtrlId
}

// Implements AnyControl.
func (me *_WindowRawControl) Parent() AnyParent {
	return me.parent
}

// Implements AnyParent.
func (me *_WindowRawControl) isDialog() bool {
	return false
}

func (me *_WindowRawControl) defaultMessages() {
	me.On().WmNcPaint(func(p wm.NcPaint) {
		_PaintThemedBorders(me.Hwnd(), p)
	})
}

//------------------------------------------------------------------------------

// Options for NewWindowControlRaw().
type WindowControlRawOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

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
	// Defaults to WS_CHILD | WS_TABSTOP | WS_GROUP | WS_VISIBLE | WS_CLIPCHILDREN | WS_CLIPSIBLINGS.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStyles co.WS_EX
	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// Control size in pixels.
	// Defaults to 300x200. Will be adjusted to the current system DPI.
	Size win.SIZE
}

func (opts *WindowControlRawOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.ClassStyles == 0 {
		opts.ClassStyles = co.CS_DBLCLKS
	}
	if opts.HCursor == 0 {
		opts.HCursor = win.HINSTANCE(0).LoadCursor(co.IDC_ARROW)
	}
	if opts.HBrushBackground == 0 {
		opts.HBrushBackground = win.CreateSysColorBrush(co.COLOR_WINDOW)
	}

	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_TABSTOP | co.WS_GROUP | co.WS_VISIBLE |
			co.WS_CLIPCHILDREN | co.WS_CLIPSIBLINGS
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_NONE
	}

	if opts.Size.Cx == 0 {
		opts.Size.Cx = 300
	}
	if opts.Size.Cy == 0 {
		opts.Size.Cy = 200
	}
}
