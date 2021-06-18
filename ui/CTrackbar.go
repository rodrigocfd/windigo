package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native trackbar control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/trackbar-controls
type Trackbar interface {
	AnyNativeControl

	// Exposes all the Trackbar notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-trackbar-control-reference-notifications
	On() *_TrackbarEvents

	PageSize() int                // Retrieves the number of positions of page up/down.
	Pos() int                     // Retrieves the current position.
	RangeMax() int                // Retrieves the maximum position.
	RangeMin() int                // Retrieves the mininum position.
	SetPageSize(pageSize int) int // Sets the number of positions of page up/down.
	SetPos(pos int)               // Sets the current position.
	SetRangeMax(max int)          // Sets the maximum position.
	SetRangeMin(min int)          // Sets the minimum position.
}

//------------------------------------------------------------------------------

type _Trackbar struct {
	_NativeControlBase
	events _TrackbarEvents
}

// Creates a new Trackbar specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewTrackbar(parent AnyParent, opts TrackbarOpts) Trackbar {
	opts.fillBlankValuesWithDefault()

	me := &_Trackbar{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.Position, &opts.Size)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"msctls_trackbar32", "", opts.Styles|co.WS(opts.TrackbarStyles),
			opts.Position, opts.Size, win.HMENU(opts.CtrlId))

		if opts.RangeMin != 0 {
			me.SetRangeMin(opts.RangeMin)
		}
		if opts.RangeMax != 0 {
			me.SetRangeMax(opts.RangeMax)
		}
		if opts.PageSize != 0 {
			me.SetPageSize(opts.PageSize)
		}
	})

	return me
}

// Creates a new Trackbar from a dialog resource.
func NewTrackbarDlg(parent AnyParent, ctrlId int) Trackbar {
	me := &_Trackbar{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return me
}

func (me *_Trackbar) On() *_TrackbarEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the Trackbar is created.")
	}
	return &me.events
}

func (me *_Trackbar) PageSize() int {
	return int(me.Hwnd().SendMessage(co.TBM_GETPAGESIZE, 0, 0))
}

func (me *_Trackbar) Pos() int {
	return int(me.Hwnd().SendMessage(co.TBM_GETPOS, 0, 0))
}

func (me *_Trackbar) RangeMax() int {
	return int(me.Hwnd().SendMessage(co.TBM_GETRANGEMAX, 0, 0))
}

func (me *_Trackbar) RangeMin() int {
	return int(me.Hwnd().SendMessage(co.TBM_GETRANGEMIN, 0, 0))
}

func (me *_Trackbar) SetPageSize(pageSize int) int {
	return int(me.Hwnd().SendMessage(co.TBM_SETPAGESIZE, 1, win.LPARAM(pageSize)))
}

func (me *_Trackbar) SetPos(pos int) {
	me.Hwnd().SendMessage(co.TBM_SETPOS, 1, win.LPARAM(pos))
}

func (me *_Trackbar) SetRangeMax(max int) {
	me.Hwnd().SendMessage(co.TBM_SETRANGEMAX, 1, win.LPARAM(max))
}

func (me *_Trackbar) SetRangeMin(min int) {
	me.Hwnd().SendMessage(co.TBM_SETRANGEMIN, 1, win.LPARAM(min))
}

//------------------------------------------------------------------------------

// Options for NewTrackbar().
type TrackbarOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// Control size in pixels.
	// Defaults to 120x23. Will be adjusted to the current system DPI.
	Size win.SIZE
	// Trackbar control styles, passed to CreateWindowEx().
	// Defauls to TBS_HORZ | TBS_AUTOTICKS.
	TrackbarStyles co.TBS
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStyles co.WS_EX

	// Number of positions of page up/down.
	// Defaults to RangeMax / 5.
	PageSize int
	// Minimum position value.
	// Defaults to 0.
	RangeMin int
	// Maximum position value.
	// Defaults to 100.
	RangeMax int
}

func (opts *TrackbarOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.Size.Cx == 0 {
		opts.Size.Cx = 120
	}
	if opts.Size.Cy == 0 {
		opts.Size.Cy = 23
	}

	if opts.TrackbarStyles == 0 {
		opts.TrackbarStyles = co.TBS_HORZ | co.TBS_AUTOTICKS
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_NONE
	}
}

//------------------------------------------------------------------------------

// Trackbar control notifications.
type _TrackbarEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_TrackbarEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.parent.On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/trbn-thumbposchanging
func (me *_TrackbarEvents) ThumbPosChanging(userFunc func(p *win.NMTRBTHUMBPOSCHANGING)) {
	me.events.addNfyZero(me.ctrlId, co.TRBN_THUMBPOSCHANGING, func(p unsafe.Pointer) {
		userFunc((*win.NMTRBTHUMBPOSCHANGING)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-trackbar-
func (me *_TrackbarEvents) NmReleasedCapture(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) {
		userFunc()
	})
}
