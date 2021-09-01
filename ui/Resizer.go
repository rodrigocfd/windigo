package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Used in Resizer.Add().
//
// Action to be done on a child control when the parent is resized.
type RESZ uint8

const (
	RESZ_REPOS   RESZ = iota // When parent resizes, change control left/top position.
	RESZ_RESIZE              // When parent resizes, change control width/height.
	RESZ_NOTHING             // When parent resizes, do nothing.
)

//------------------------------------------------------------------------------

type Resizer interface {
	isResizer() // prevent public implementation

	// Adds child controls, and their behavior when the parent is resized.
	//
	// Should be called before the window is created.
	Add(horzBehavior, vertBehavior RESZ, ctrls ...AnyControl) Resizer
}

type _ResizerCtrl struct {
	hChild     AnyControl
	rcOrig     win.RECT
	horzAction RESZ
	vertAction RESZ
}

//------------------------------------------------------------------------------

type _Resizer struct {
	parent AnyParent
	ctrls  []_ResizerCtrl
	szOrig win.SIZE // Original client area of parent.
}

// Creates a new Resizer. This function must be called after all controls are
// created themselves.
func NewResizer(parent AnyParent) Resizer {
	me := &_Resizer{
		parent: parent,
	}

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		rcParent := me.parent.Hwnd().GetClientRect()
		me.szOrig = win.SIZE{Cx: rcParent.Right, Cy: rcParent.Bottom} // save parent client area

		for i := range me.ctrls {
			me.ctrls[i].rcOrig = me.ctrls[i].hChild.Hwnd().GetWindowRect() // relative to screen
			parent.Hwnd().ScreenToClientRc(&me.ctrls[i].rcOrig)            // now relative to parent
		}
	})

	parent.internalOn().addMsgZero(co.WM_SIZE, func(p wm.Any) {
		me.adjustToParent(wm.Size{Msg: p})
	})

	return me
}

func (me *_Resizer) isResizer() {}

func (me *_Resizer) Add(
	horzBehavior, vertBehavior RESZ, ctrls ...AnyControl) Resizer {

	if me.parent.Hwnd() != 0 {
		panic("Cannot add Resizer controls after the window is created.")
	}

	for _, ctrl := range ctrls {
		me.ctrls = append(me.ctrls, _ResizerCtrl{
			hChild:     ctrl,
			rcOrig:     win.RECT{}, // will be set during WM_CREATE or WM_INITDIALOG
			horzAction: horzBehavior,
			vertAction: vertBehavior,
		})
	}
	return me
}

func (me *_Resizer) adjustToParent(parm wm.Size) {
	if len(me.ctrls) == 0 || parm.Request() == co.SIZE_REQ_MINIMIZED {
		return // no need to resize if window is minimized
	}

	hdwp := win.BeginDeferWindowPos(int32(len(me.ctrls)))
	defer hdwp.EndDeferWindowPos()

	for i := range me.ctrls {
		ctl := me.ctrls[i]

		uFlags := co.SWP_NOZORDER
		if ctl.horzAction == RESZ_REPOS && ctl.vertAction == RESZ_REPOS { // repos both horz and vert
			uFlags |= co.SWP_NOSIZE
		} else if ctl.horzAction == RESZ_RESIZE && ctl.vertAction == RESZ_RESIZE { // resize both horz and vert
			uFlags |= co.SWP_NOMOVE
		}

		szParent := parm.ClientAreaSize()

		x := ctl.rcOrig.Left // keep original left pos
		if ctl.horzAction == RESZ_REPOS {
			x = szParent.Cx - me.szOrig.Cx + ctl.rcOrig.Left
		}

		y := ctl.rcOrig.Top // keep original top pos
		if ctl.vertAction == RESZ_REPOS {
			y = szParent.Cy - me.szOrig.Cy + ctl.rcOrig.Top
		}

		cx := ctl.rcOrig.Right - ctl.rcOrig.Left // keep original width
		if ctl.horzAction == RESZ_RESIZE {
			cx = szParent.Cx - me.szOrig.Cx + ctl.rcOrig.Right - ctl.rcOrig.Left
		}

		cy := ctl.rcOrig.Bottom - ctl.rcOrig.Top // keep original height
		if ctl.vertAction == RESZ_RESIZE {
			cy = szParent.Cy - me.szOrig.Cy + ctl.rcOrig.Bottom - ctl.rcOrig.Top
		}

		hdwp.DeferWindowPos(ctl.hChild.Hwnd(), win.HWND(0), x, y, cx, cy, uFlags)
	}
}
