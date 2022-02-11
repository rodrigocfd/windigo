package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Implements WindowControl interface.
type _WindowRawControl struct {
	_WindowRaw
	opts   *_WindowControlO
	parent AnyParent
}

// Creates a new WindowControl. Call WindowControlOpts() to define the options
// to be passed to the underlying CreateWindowEx().
//
// Example:
//
//  var owner AnyParent // initialized somewhere
//
//  myControl := ui.NewWindowControl(
//      owner,
//      ui.WindowControlOpts(
//          Position(win.POINT{X: 100, Y: 100}).
//          Size(win.SIZE{Cx: 300, Cy: 200}).
//          RangeMax(4),
//      ),
//  )
func NewWindowControl(parent AnyParent, opts *_WindowControlO) WindowControl {
	if opts == nil {
		opts = WindowControlOpts()
	}
	opts.lateDefaults()

	me := &_WindowRawControl{}
	me._WindowRaw.new()
	me.opts = opts
	me.parent = parent

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		hInst := parent.Hwnd().Hinstance()
		var wcx win.WNDCLASSEX
		me.opts.className = me._WindowRaw.generateWcx(&wcx, hInst,
			me.opts.className, me.opts.classStyles, me.opts.hCursor,
			me.opts.hBrushBkgnd, 0)
		atom := me._WindowRaw.registerClass(&wcx)

		_ConvertDtuOrMultiplyDpi(parent, &me.opts.position, &me.opts.size)
		me._WindowRaw.createWindow(me.opts.wndExStyles, win.ClassNameAtom(atom),
			win.StrOptNone{}, me.opts.wndStyles, me.opts.position, me.opts.size,
			parent.Hwnd(), win.HMENU(me.opts.ctrlId), hInst)

		parent.addResizingChild(me, opts.horz, opts.vert)
	})

	me.defaultMessages()
	return me
}

// Implements AnyControl.
func (me *_WindowRawControl) CtrlId() int {
	return me.opts.ctrlId
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

type _WindowControlO struct {
	ctrlId int

	className   string // define in NewWindowControl()
	classStyles co.CS
	hCursor     win.HCURSOR
	hBrushBkgnd win.HBRUSH

	wndStyles   co.WS
	wndExStyles co.WS_EX
	position    win.POINT
	size        win.SIZE
	horz        HORZ
	vert        VERT
}

// Control ID.
// Defaults to an auto-generated ID.
func (o *_WindowControlO) CtrlId(i int) *_WindowControlO { o.ctrlId = i; return o }

// Class name registered with RegisterClassEx().
// Defaults to a computed hash.
func (o *_WindowControlO) ClassName(n string) *_WindowControlO { o.className = n; return o }

// Window class styles, passed to RegisterClassEx().
// Defaults to CS_DBLCLKS.
func (o *_WindowControlO) ClassStyles(s co.CS) *_WindowControlO { o.classStyles = s; return o }

// Window cursor, passed to RegisterClassEx().
// Defaults to stock IDC_ARROW.
func (o *_WindowControlO) HCursor(h win.HCURSOR) *_WindowControlO { o.hCursor = h; return o }

// Window background brush, passed to RegisterClassEx().
// Defaults to COLOR_BTNFACE color.
func (o *_WindowControlO) HBrushBkgnd(h win.HBRUSH) *_WindowControlO { o.hBrushBkgnd = h; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to WS_CHILD | WS_TABSTOP | WS_GROUP | WS_VISIBLE | WS_CLIPCHILDREN | WS_CLIPSIBLINGS.
func (o *_WindowControlO) WndStyles(s co.WS) *_WindowControlO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
// Defaults to WS_EX_CLIENTEDGE.
func (o *_WindowControlO) WndExStyles(s co.WS_EX) *_WindowControlO { o.wndExStyles = s; return o }

// Position within parent's client area in pixels.
// Defaults to 0x0. Will be adjusted to the current system DPI.
func (o *_WindowControlO) Position(p win.POINT) *_WindowControlO { _OwPt(&o.position, p); return o }

// Control size in pixels.
// Defaults to 300x200. Will be adjusted to the current system DPI.
func (o *_WindowControlO) Size(s win.SIZE) *_WindowControlO { _OwSz(&o.size, s); return o }

// Horizontal behavior when the parent is resized.
// Defaults to HORZ_NONE.
func (o *_WindowControlO) Horz(s HORZ) *_WindowControlO { o.horz = s; return o }

// Vertical behavior when the parent is resized.
// Defaults to VERT_NONE.
func (o *_WindowControlO) Vert(s VERT) *_WindowControlO { o.vert = s; return o }

func (o *_WindowControlO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
	if o.hCursor == 0 {
		o.hCursor = win.HINSTANCE(0).LoadCursor(win.CursorResIdc(co.IDC_ARROW))
	}
}

// Options for NewWindowControl().
func WindowControlOpts() *_WindowControlO {
	return &_WindowControlO{
		classStyles: co.CS_DBLCLKS,
		hBrushBkgnd: win.CreateSysColorBrush(co.COLOR_WINDOW),
		wndStyles: co.WS_CHILD | co.WS_TABSTOP | co.WS_GROUP | co.WS_VISIBLE |
			co.WS_CLIPCHILDREN | co.WS_CLIPSIBLINGS,
		wndExStyles: co.WS_EX_CLIENTEDGE,
		size:        win.SIZE{Cx: 300, Cy: 200},
		horz:        HORZ_NONE,
		vert:        VERT_NONE,
	}
}
