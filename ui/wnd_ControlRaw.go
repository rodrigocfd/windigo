//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _ControlRaw struct {
	_BaseRaw
	ctrlId uint16
}

// Constructor.
func newControlRaw(parent Parent, opts *VarOptsControl) *_ControlRaw {
	setUniqueCtrlId(&opts.ctrlId)
	me := &_ControlRaw{
		_BaseRaw: newBaseRaw(),
		ctrlId:   opts.ctrlId,
	}

	parent.base().beforeUserEvents.Wm(parent.base().wndTy.initMsg(), func(_ Wm) uintptr {
		hInst, _ := parent.Hwnd().HInstance()
		atom := me.registerClass(hInst, opts.className, opts.classStyle,
			0, opts.classBrush, opts.classCursor)
		me.createWindow(opts.exStyle, atom, "", opts.style,
			opts.position, opts.size, parent.Hwnd(), win.HMENU(opts.ctrlId), hInst)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		return 0 // ignored
	})

	me.defaultMessageHandlers()
	return me
}

func (me *_ControlRaw) defaultMessageHandlers() {
	me.userEvents.WmNcPaint(func(p WmNcPaint) {
		paintThemedBorders(me.hWnd, p)
	})
}

// Options for [NewControl]; returned by [OptsControl].
type VarOptsControl struct {
	className   string
	classStyle  co.CS
	classCursor win.HCURSOR
	classBrush  win.HBRUSH

	ctrlId   uint16
	layout   LAY
	position win.POINT
	size     win.SIZE
	style    co.WS
	exStyle  co.WS_EX
}

// Options for [NewControl].
func OptsControl() *VarOptsControl {
	hCursor, _ := win.HINSTANCE(0).LoadCursor(win.CursorResIdc(co.IDC_ARROW))
	return &VarOptsControl{
		classStyle:  co.CS_DBLCLKS,
		classCursor: hCursor,
		classBrush:  win.HBRUSH(co.COLOR_WINDOW + 1),
		size:        win.SIZE{Cx: int32(DpiX(300)), Cy: int32(DpiY(200))},
		style:       co.WS_CHILD | co.WS_TABSTOP | co.WS_GROUP | co.WS_VISIBLE | co.WS_CLIPCHILDREN | co.WS_CLIPSIBLINGS,
		exStyle:     co.WS_EX_LEFT | co.WS_EX_CLIENTEDGE,
	}
}

// Class name registered with [RegisterClassEx].
//
// Defaults to a computed hash.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsControl) ClassName(s string) *VarOptsControl { o.className = s; return o }

// Window class style, passed to [RegisterClassEx].
//
// Defaults to co.CS_DBLCLKS.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsControl) ClassStyle(s co.CS) *VarOptsControl { o.classStyle = s; return o }

// Window cursor, passed to [RegisterClassEx].
//
// Defaults to stock co.IDC_ARROW.
//
// # Example
//
//	hCursor, _ := win.HINSTANCE(0).
//		LoadCursor(win.CursorResIdc(co.IDC_ARROW))
//
//	ui.OptsControl().
//		ClassCursor(hCursor)
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsControl) ClassCursor(h win.HCURSOR) *VarOptsControl { o.classCursor = h; return o }

// Window background brush, passed to [RegisterClassEx].
//
// Defaults to co.COLOR_WINDOW color.
//
// # Example
//
//	ui.OptsControl().
//		ClassBrush(win.HBRUSH(co.COLOR_WINDOW + 1))
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func (o *VarOptsControl) ClassBrush(h win.HBRUSH) *VarOptsControl { o.classBrush = h; return o }

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsControl) CtrlId(id uint16) *VarOptsControl { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsControl) Layout(l LAY) *VarOptsControl { o.layout = l; return o }

// Position coordinates within parent window client area, passed to
// [CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsControl) Position(x, y int) *VarOptsControl {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [CreateWindowEx].
//
// Defaults to ui.Dpi(300, 200).
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsControl) Size(cx int, cy int) *VarOptsControl {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Window style, passed to [CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_TABSTOP | co.WS_GROUP | co.WS_VISIBLE | co.WS_CLIPCHILDREN | co.WS_CLIPSIBLINGS.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsControl) Style(s co.WS) *VarOptsControl { o.style = s; return o }

// Extended window style, passed to [CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT | co.WS_EX_CLIENTEDGE.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsControl) ExStyle(s co.WS_EX) *VarOptsControl { o.exStyle = s; return o }
