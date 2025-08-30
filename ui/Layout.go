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
	_LAYH_REPOS  LAY = 0b0000_0001
	_LAYH_RESIZE LAY = 0b0000_0010
	_LAYV_REPOS  LAY = 0b0000_0100
	_LAYV_RESIZE LAY = 0b0000_1000

	// When parent is resized, nothing happens.
	LAY_NONE_NONE LAY = 0
	// When parent resizes:
	//	- horizontal: nothing happens;
	//	- vertical: control moves anchored at bottom.
	LAY_NONE_REPOS = _LAYV_REPOS
	// When parent resizes:
	//	- horizontal: nothing happens;
	//	- vertical: control is resized together.
	LAY_NONE_RESIZE = _LAYV_RESIZE
	// When parent resizes:
	//	- horizontal: control moves anchored at right;
	//	- vertical: nothing happens.
	LAY_REPOS_NONE = _LAYH_REPOS
	// When parent resizes:
	//	- horizontal: control moves anchored at right;
	//	- vertical: control moves anchored at bottom.
	LAY_REPOS_REPOS = _LAYH_REPOS | _LAYV_REPOS
	// When parent resizes:
	//	- horizontal: control moves anchored at right;
	//	- vertical: control is resized together.
	LAY_REPOS_RESIZE = _LAYH_REPOS | _LAYV_RESIZE
	// When parent resizes:
	//	- horizontal: control is resized together;
	//	- vertical: nothing happens.
	LAY_RESIZE_NONE = _LAYH_RESIZE
	// When parent resizes:
	//	- horizontal: control is resized together;
	//	- vertical: control moves anchored at bottom.
	LAY_RESIZE_REPOS = _LAYH_RESIZE | _LAYV_REPOS
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
	if layout == LAY_NONE_NONE {
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

	hdwp, _ := win.BeginDeferWindowPos(uint(len(me.ctrls)))
	defer hdwp.EndDeferWindowPos()

	for i := range me.ctrls {
		ctl := me.ctrls[i]

		uFlags := co.SWP_NOZORDER
		switch ctl.layout {
		case LAY_REPOS_REPOS: // repos both horz and vert
			uFlags |= co.SWP_NOSIZE
		case LAY_RESIZE_RESIZE: // resize both horz and vert
			uFlags |= co.SWP_NOMOVE
		}

		szParent := parm.ClientAreaSize()

		x := ctl.rcOrig.Left // keep original left pos
		if (ctl.layout & _LAYH_REPOS) != 0 {
			x = szParent.Cx - me.szOrig.Cx + ctl.rcOrig.Left
		}

		y := ctl.rcOrig.Top // keep original top pos
		if (ctl.layout & _LAYV_REPOS) != 0 {
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
