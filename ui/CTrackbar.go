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
	isTrackbar() // prevent public implementation

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
func NewTrackbar(parent AnyParent, opts *_TrackbarO) Trackbar {
	opts.lateDefaults()

	me := &_Trackbar{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.position, &opts.size)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			"msctls_trackbar32", "", opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		if opts.rangeMin != 0 {
			me.SetRangeMin(opts.rangeMin)
		}
		if opts.rangeMax != 0 {
			me.SetRangeMax(opts.rangeMax)
		}
		if opts.pageSize != 0 {
			me.SetPageSize(opts.pageSize)
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

func (me *_Trackbar) isTrackbar() {}

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
	return int(
		me.Hwnd().SendMessage(co.TBM_SETPAGESIZE, 1, win.LPARAM(pageSize)),
	)
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

type _TrackbarO struct {
	ctrlId int

	position    win.POINT
	size        win.SIZE
	ctrlStyles  co.TBS
	wndStyles   co.WS
	wndExStyles co.WS_EX

	pageSize int
	rangeMin int
	rangeMax int
}

// Control ID.
// Defaults to an auto-generated ID.
func (o *_TrackbarO) CtrlId(i int) *_TrackbarO { o.ctrlId = i; return o }

// Position within parent's client area in pixels.
// Defaults to 0x0. Will be adjusted to the current system DPI.
func (o *_TrackbarO) Position(p win.POINT) *_TrackbarO { _OwPt(&o.position, p); return o }

// Control size in pixels.
// Defaults to 175x28. Will be adjusted to the current system DPI.
func (o *_TrackbarO) Size(s win.SIZE) *_TrackbarO { _OwSz(&o.size, s); return o }

// Trackbar control styles, passed to CreateWindowEx().
// Defauls to TBS_HORZ | TBS_AUTOTICKS.
func (o *_TrackbarO) CtrlStyles(s co.TBS) *_TrackbarO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *_TrackbarO) WndStyles(s co.WS) *_TrackbarO { o.wndStyles = s; return o }

// Number of positions of page up/down.
// Defaults to RangeMax / 5.
func (o *_TrackbarO) PageSize(p int) *_TrackbarO { o.pageSize = p; return o }

// Minimum position value.
// Defaults to 0.
func (o *_TrackbarO) RangeMin(r int) *_TrackbarO { o.rangeMin = r; return o }

// Maximum position value.
// Defaults to 100.
func (o *_TrackbarO) RangeMax(r int) *_TrackbarO { o.rangeMax = r; return o }

func (o *_TrackbarO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewTrackbar().
func TrackbarOpts() *_TrackbarO {
	return &_TrackbarO{
		size:       win.SIZE{Cx: 175, Cy: 28},
		ctrlStyles: co.TBS_HORZ | co.TBS_AUTOTICKS,
		wndStyles:  co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
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
	me.events = ctrl.Parent().On()
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
