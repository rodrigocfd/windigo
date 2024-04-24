package main

import (
	"github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Child window which displays the playback progress.
type Tracker struct {
	wnd           ui.WindowControl
	onClickCB     func(pct float32)
	onSpaceCB     func()
	onLeftRightCB func(key co.VK)
	elapsed       float32
}

func NewTracker(
	parent ui.AnyParent, pos win.POINT, sz win.SIZE,
	horz ui.HORZ, vert ui.VERT) *Tracker {

	wnd := ui.NewWindowControl(
		parent,
		ui.WindowControlOpts().
			WndExStyles(co.WS_EX_NONE).
			Position(pos).
			Size(sz).
			Horz(horz).
			Vert(vert).
			HCursor(win.HINSTANCE(0).LoadCursor(win.CursorResIdc(co.IDC_HAND))),
	)

	me := &Tracker{
		wnd: wnd,
	}

	me.events()
	return me
}

func (me *Tracker) OnClick(fun func(pct float32)) {
	me.onClickCB = fun
}

func (me *Tracker) OnSpace(fun func()) {
	me.onSpaceCB = fun
}

func (me *Tracker) OnLeftRight(fun func(key co.VK)) {
	me.onLeftRightCB = fun
}

func (me *Tracker) SetElapsed(pct float32) {
	me.elapsed = pct
	me.wnd.Hwnd().InvalidateRect(nil, true)
}

func (me *Tracker) events() {
	me.wnd.On().WmPaint(func() {
		hwnd := me.wnd.Hwnd()
		hasFocus := win.GetFocus() == hwnd

		ps := win.PAINTSTRUCT{}
		hdc := hwnd.BeginPaint(&ps)
		defer hwnd.EndPaint(&ps)

		var fillColor win.COLORREF
		if hasFocus {
			fillColor = win.GetSysColor(co.COLOR_ACTIVECAPTION)
		} else {
			fillColor = win.GetSysColor(co.COLOR_ACTIVEBORDER)
		}

		myPen := win.CreatePen(co.PS_SOLID, 1, fillColor)
		defer myPen.DeleteObject()
		defPen := hdc.SelectObjectPen(myPen)
		defer hdc.SelectObjectPen(defPen)

		myBrush := win.CreateSolidBrush(fillColor)
		defer myBrush.DeleteObject()
		defBrush := hdc.SelectObjectBrush(myBrush)
		defer hdc.SelectObjectBrush(defBrush)

		rcClient := hwnd.GetClientRect()
		hdc.Rectangle(win.RECT{
			Left:   0,
			Top:    0,
			Right:  int32(float32(rcClient.Right) * me.elapsed),
			Bottom: rcClient.Bottom,
		})
	})

	me.wnd.On().WmLButtonDown(func(p wm.Mouse) {
		me.wnd.Hwnd().SetFocus()

		if me.onClickCB != nil {
			rcClient := me.wnd.Hwnd().GetClientRect()
			pct := float32(p.Pos().X) / float32(rcClient.Right)
			me.onClickCB(pct)
		}
	})

	me.wnd.On().WmKeyDown(func(p wm.Key) {
		if p.VirtualKeyCode() == co.VK_SPACE {
			if me.onSpaceCB != nil {
				me.onSpaceCB()
			}
		}
	})

	me.wnd.On().WmGetDlgCode(func(p wm.GetDlgCode) co.DLGC {
		if p.VirtualKeyCode() == co.VK_LEFT || p.VirtualKeyCode() == co.VK_RIGHT {
			if me.onLeftRightCB != nil {
				me.onLeftRightCB(p.VirtualKeyCode())
			}
			return co.DLGC_WANTARROWS
		}
		return co.DLGC_NONE
	})
}
