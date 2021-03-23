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
	AnyControl

	isProgressBar() // disambiguate

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

// Creates a new ProgressBar specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewProgressBarOpts(parent AnyParent, opts ProgressBarOpts) ProgressBar {
	opts.fillBlankValuesWithDefault()

	me := _ProgressBar{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.isMarquee = false

	parent.internalOn().addMsgZero(_ParentCreateWm(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.Position, &opts.Size)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"msctls_progress32", "", opts.Styles|co.WS(opts.ProgressBarStyles),
			opts.Position, opts.Size, win.HMENU(opts.CtrlId))
	})

	return &me
}

// Creates a new ProgressBar from a dialog resource.
func NewProgressBarDlg(parent AnyParent, ctrlId int) ProgressBar {
	me := _ProgressBar{}
	me._NativeControlBase.new(parent, ctrlId)
	me.isMarquee = false

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return &me
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

// Options for NewProgressBarOpts().
type ProgressBarOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// Control size in pixels.
	// Defaults to 120x23. Will be adjusted to the current system DPI.
	Size win.SIZE
	// ProgressBar control styles, passed to CreateWindowEx().
	// Defaults to PBS_SMOOTH.
	ProgressBarStyles co.PBS
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_VISIBLE.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStyles co.WS_EX
}

func (opts *ProgressBarOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.Size.Cx == 0 {
		opts.Size.Cx = 120
	}
	if opts.Size.Cy == 0 {
		opts.Size.Cy = 23
	}

	if opts.ProgressBarStyles == 0 {
		opts.ProgressBarStyles = co.PBS_SMOOTH
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_NONE
	}
}
