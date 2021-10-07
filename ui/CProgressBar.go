package ui

import (
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native progress bar control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/progress-bar-control
type ProgressBar interface {
	AnyNativeControl
	isProgressBar() // prevent public implementation

	Pos() int                  // Retrieves the current position.
	SetMarquee(isMarquee bool) // Sets indeterminate state, a graphic animation going back and forth.
	SetPos(pos int)            // Sets the current position.
	SetRange(min, max int)     // Sets the new range. Default is 0-100.
	SetState(state co.PBST)    // Sets the current state (green, yellow, red).
}

//------------------------------------------------------------------------------

type _ProgressBar struct {
	_NativeControlBase
	isMarquee bool
}

// Creates a new ProgressBar. Call ProgressBarOpts() to define the options to be
// passed to the underlying CreateWindowEx().
func NewProgressBar(parent AnyParent, opts *_ProgressBarO) ProgressBar {
	opts.lateDefaults()

	me := &_ProgressBar{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.isMarquee = false

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.position, &opts.size)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			"msctls_progress32", "", opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		parent.addResizerChild(me, opts.horz, opts.vert)
	})

	return me
}

// Creates a new ProgressBar from a dialog resource.
func NewProgressBarDlg(
	parent AnyParent, ctrlId int,
	horz HORZ, vert VERT) ProgressBar {

	me := &_ProgressBar{}
	me._NativeControlBase.new(parent, ctrlId)
	me.isMarquee = false

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
		parent.addResizerChild(me, horz, vert)
	})

	return me
}

func (me *_ProgressBar) isProgressBar() {}

func (me *_ProgressBar) Pos() int {
	return int(me.Hwnd().SendMessage(co.PBM_GETPOS, 0, 0))
}

func (me *_ProgressBar) SetMarquee(isMarquee bool) {
	if isMarquee {
		me.Hwnd().SetWindowLongPtr(co.GWLP_STYLE,
			me.Hwnd().GetWindowLongPtr(co.GWLP_STYLE)|uintptr(co.PBS_MARQUEE))
	}

	me.Hwnd().SendMessage(co.PBM_SETMARQUEE,
		win.WPARAM(util.BoolToUintptr(isMarquee)), 0)

	if !isMarquee {
		me.Hwnd().SetWindowLongPtr(co.GWLP_STYLE,
			me.Hwnd().GetWindowLongPtr(co.GWLP_STYLE)&^uintptr(co.PBS_MARQUEE))
	}

	me.isMarquee = isMarquee
}

func (me *_ProgressBar) SetPos(pos int) {
	if me.isMarquee {
		me.SetMarquee(false) // avoid crash
	}
	me.Hwnd().SendMessage(co.PBM_SETPOS, win.WPARAM(pos), 0)
}

func (me *_ProgressBar) SetRange(min, max int) {
	me.Hwnd().SendMessage(co.PBM_SETRANGE32,
		win.WPARAM(min), win.LPARAM(max))
}

func (me *_ProgressBar) SetState(state co.PBST) {
	me.Hwnd().SendMessage(co.PBM_SETSTATE, win.WPARAM(co.PBM_SETSTATE), 0)
}

//------------------------------------------------------------------------------

type _ProgressBarO struct {
	ctrlId int

	position    win.POINT
	size        win.SIZE
	horz        HORZ
	vert        VERT
	ctrlStyles  co.PBS
	wndStyles   co.WS
	wndExStyles co.WS_EX
}

// Control ID.
// Defaults to an auto-generated ID.
func (o *_ProgressBarO) CtrlId(i int) *_ProgressBarO { o.ctrlId = i; return o }

// Position within parent's client area in pixels.
// Defaults to 0x0. Will be adjusted to the current system DPI.
func (o *_ProgressBarO) Position(p win.POINT) *_ProgressBarO { _OwPt(&o.position, p); return o }

// Control size in pixels.
// Defaults to 140x26. Will be adjusted to the current system DPI.
func (o *_ProgressBarO) Size(s win.SIZE) *_ProgressBarO { _OwSz(&o.size, s); return o }

// Horizontal behavior when the parent is resized.
// Defaults to HORZ_NONE.
func (o *_ProgressBarO) Horz(s HORZ) *_ProgressBarO { o.horz = s; return o }

// Vertical behavior when the parent is resized.
// Defaults to VERT_NONE.
func (o *_ProgressBarO) Vert(s VERT) *_ProgressBarO { o.vert = s; return o }

// ProgressBar control styles, passed to CreateWindowEx().
// Defaults to PBS_SMOOTH.
func (o *_ProgressBarO) CtrlStyles(s co.PBS) *_ProgressBarO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
func (o *_ProgressBarO) WndStyles(s co.WS) *_ProgressBarO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
// Defaults to WS_EX_NONE.
func (o *_ProgressBarO) WndExStyles(s co.WS_EX) *_ProgressBarO { o.wndExStyles = s; return o }

func (o *_ProgressBarO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewProgressBar().
func ProgressBarOpts() *_ProgressBarO {
	return &_ProgressBarO{
		size:       win.SIZE{Cx: 140, Cy: 26},
		horz:       HORZ_NONE,
		vert:       VERT_NONE,
		ctrlStyles: co.PBS_SMOOTH,
		wndStyles:  co.WS_CHILD | co.WS_VISIBLE,
	}
}
