//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
type LAY uint8

const (
	_LAYH_MOVE   LAY = 0b0000_0001
	_LAYH_RESIZE LAY = 0b0000_0010
	_LAYV_MOVE   LAY = 0b0000_0100
	_LAYV_RESIZE LAY = 0b0000_1000

	// When parent is resized, nothing happens.
	LAY_HOLD_HOLD LAY = 0
	// When parent resizes:
	//	- horizontal: nothing happens;
	//	- vertical: control moves anchored at bottom.
	LAY_HOLD_MOVE = _LAYV_MOVE
	// When parent resizes:
	//	- horizontal: nothing happens;
	//	- vertical: control is resized together.
	LAY_HOLD_RESIZE = _LAYV_RESIZE
	// When parent resizes:
	//	- horizontal: control moves anchored at right;
	//	- vertical: nothing happens.
	LAY_MOVE_HOLD = _LAYH_MOVE
	// When parent resizes:
	//	- horizontal: control moves anchored at right;
	//	- vertical: control moves anchored at bottom.
	LAY_MOVE_MOVE = _LAYH_MOVE | _LAYV_MOVE
	// When parent resizes:
	//	- horizontal: control moves anchored at right;
	//	- vertical: control is resized together.
	LAY_MOVE_RESIZE = _LAYH_MOVE | _LAYV_RESIZE
	// When parent resizes:
	//	- horizontal: control is resized together;
	//	- vertical: nothing happens.
	LAY_RESIZE_HOLD = _LAYH_RESIZE
	// When parent resizes:
	//	- horizontal: control is resized together;
	//	- vertical: control moves anchored at bottom.
	LAY_RESIZE_MOVE = _LAYH_RESIZE | _LAYV_MOVE
	// When parent resizes:
	//	- horizontal: control is resized together;
	//	- vertical: control is resized together.
	LAY_RESIZE_RESIZE = _LAYH_RESIZE | _LAYV_RESIZE
)

// When parent window is resized, resizes all children at once.
type _Layout struct {
	ctrls  []_LayoutCtrl
	szOrig win.SIZE // Original size of parent's client area.
}

type _LayoutCtrl struct {
	hCtrl  win.HWND
	rcOrig win.RECT
	layout LAY
}

// Constructor.
func newLayout() _Layout {
	return _Layout{
		ctrls: make([]_LayoutCtrl, 0, 8), // arbitrary
	}
}

// Adds a new control to be resized.
//
// Must be called after both the parent and the children were created, because
// both HWNDs are used.
func (me *_Layout) Add(parent Parent, hCtrl win.HWND, layout LAY) {
	if layout == LAY_HOLD_HOLD {
		return // nothing to do, don't even bother adding the control
	}

	if len(me.ctrls) == 0 { // first control being added?
		rcParent, _ := parent.Hwnd().GetClientRect()
		me.szOrig = win.SIZE{Cx: rcParent.Right, Cy: rcParent.Bottom} // save parent client area
	}

	rcOrig, _ := hCtrl.GetWindowRect()      // relative to screen
	parent.Hwnd().ScreenToClientRc(&rcOrig) // now relative to parent

	me.ctrls = append(me.ctrls, _LayoutCtrl{hCtrl, rcOrig, layout})
}

// Rearrange all children. To be called during WM_SIZE processing.
func (me *_Layout) Rearrange(parm WmSize) {
	if len(me.ctrls) == 0 || parm.Request() == co.SIZE_REQ_MINIMIZED {
		return // no need to resize if window is minimized
	}

	hdwp, _ := win.BeginDeferWindowPos(len(me.ctrls))
	defer hdwp.EndDeferWindowPos()

	for i := range me.ctrls {
		ctl := me.ctrls[i]

		uFlags := co.SWP_NOZORDER
		switch ctl.layout {
		case LAY_MOVE_MOVE: // repos both horz and vert
			uFlags |= co.SWP_NOSIZE
		case LAY_RESIZE_RESIZE: // resize both horz and vert
			uFlags |= co.SWP_NOMOVE
		}

		szParent := parm.ClientAreaSize()

		x := ctl.rcOrig.Left // keep original left pos
		if (ctl.layout & _LAYH_MOVE) != 0 {
			x = szParent.Cx - me.szOrig.Cx + ctl.rcOrig.Left
		}

		y := ctl.rcOrig.Top // keep original top pos
		if (ctl.layout & _LAYV_MOVE) != 0 {
			y = szParent.Cy - me.szOrig.Cy + ctl.rcOrig.Top
		}

		cx := ctl.rcOrig.Right - ctl.rcOrig.Left // keep original width
		if (ctl.layout & _LAYH_RESIZE) != 0 {
			cx = szParent.Cx - me.szOrig.Cx + ctl.rcOrig.Right - ctl.rcOrig.Left
		}

		cy := ctl.rcOrig.Bottom - ctl.rcOrig.Top // keep original height
		if (ctl.layout & _LAYV_RESIZE) != 0 {
			cy = szParent.Cy - me.szOrig.Cy + ctl.rcOrig.Bottom - ctl.rcOrig.Top
		}

		hdwp.DeferWindowPos(ctl.hCtrl, win.HWND(0), int(x), int(y), int(cx), int(cy), uFlags)
	}
}
