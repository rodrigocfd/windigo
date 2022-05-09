//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Horizontal action to be performed when the parent window is resized.
type HORZ uint8

const (
	HORZ_NONE   HORZ = iota // When parent is resized, nothing happens to the control.
	HORZ_REPOS              // When parent is resized, control moves anchored at right.
	HORZ_RESIZE             // When parent is resized, control is resized together.
)

// Vertical action to be performed when the parent window is resized.
type VERT uint8

const (
	VERT_NONE   VERT = iota // When parent is resized, nothing happens to the control.
	VERT_REPOS              // When parent is resized, control moves anchored at bottom.
	VERT_RESIZE             // When parent is resized, control is resized together.
)

//------------------------------------------------------------------------------

type _ResizerChildrenCtrl struct {
	ctrl   AnyControl
	rcOrig win.RECT
	horz   HORZ
	vert   VERT
}

type _ResizerChildren struct {
	ctrls  []_ResizerChildrenCtrl
	szOrig win.SIZE // original size of parent's client area
}

func (me *_ResizerChildren) new() {
	me.ctrls = make([]_ResizerChildrenCtrl, 0, 16) // arbitrary
	me.szOrig = win.SIZE{Cx: 0, Cy: 0}
}

func (me *_ResizerChildren) add(
	hParent win.HWND, ctrl AnyControl, horz HORZ, vert VERT) {

	// Must be called after both the parent and the children were created,
	// because both HWNDs are used.

	if len(me.ctrls) == 0 { // first control being added?
		rcParent := hParent.GetClientRect()
		me.szOrig = win.SIZE{Cx: rcParent.Right, Cy: rcParent.Bottom} // save parent client area
	}

	rcOrig := ctrl.Hwnd().GetWindowRect() // relative to screen
	hParent.ScreenToClientRc(&rcOrig)     // now relative to parent

	me.ctrls = append(me.ctrls, _ResizerChildrenCtrl{
		ctrl:   ctrl,
		rcOrig: rcOrig,
		horz:   horz,
		vert:   vert,
	})
}

func (me *_ResizerChildren) resizeChildren(parm wm.Size) {
	if len(me.ctrls) == 0 || parm.Request() == co.SIZE_REQ_MINIMIZED {
		return // no need to resize if window is minimized
	}

	hdwp := win.BeginDeferWindowPos(int32(len(me.ctrls)))
	defer hdwp.EndDeferWindowPos()

	for i := range me.ctrls {
		ctl := me.ctrls[i]

		uFlags := co.SWP_NOZORDER
		if ctl.horz == HORZ_REPOS && ctl.vert == VERT_REPOS { // repos both horz and vert
			uFlags |= co.SWP_NOSIZE
		} else if ctl.horz == HORZ_RESIZE && ctl.vert == VERT_RESIZE { // resize both horz and vert
			uFlags |= co.SWP_NOMOVE
		}

		szParent := parm.ClientAreaSize()

		x := ctl.rcOrig.Left // keep original left pos
		if ctl.horz == HORZ_REPOS {
			x = szParent.Cx - me.szOrig.Cx + ctl.rcOrig.Left
		}

		y := ctl.rcOrig.Top // keep original top pos
		if ctl.vert == VERT_REPOS {
			y = szParent.Cy - me.szOrig.Cy + ctl.rcOrig.Top
		}

		cx := ctl.rcOrig.Right - ctl.rcOrig.Left // keep original width
		if ctl.horz == HORZ_RESIZE {
			cx = szParent.Cx - me.szOrig.Cx + ctl.rcOrig.Right - ctl.rcOrig.Left
		}

		cy := ctl.rcOrig.Bottom - ctl.rcOrig.Top // keep original height
		if ctl.vert == VERT_RESIZE {
			cy = szParent.Cy - me.szOrig.Cy + ctl.rcOrig.Bottom - ctl.rcOrig.Top
		}

		hdwp.DeferWindowPos(ctl.ctrl.Hwnd(), win.HWND(0), x, y, cx, cy, uFlags)
	}
}
