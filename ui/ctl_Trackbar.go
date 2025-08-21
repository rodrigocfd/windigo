//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [trackbar] control.
//
// [trackbar]: https://learn.microsoft.com/en-us/windows/win32/controls/trackbar-controls
type Trackbar struct {
	_BaseCtrl
	events EventsTrackbar
}

// Creates a new [Trackbar] with [win.CreateWindowEx].
//
// Example:
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	trackbar := ui.NewTrackbar(
//		wndOwner,
//		ui.OptsTrackbar().
//			Position(ui.Dpi(100, 100)).
//			Range(0, 10),
//	)
//
//	trackbar.On().WmHScroll(func(_ WmScroll) {
//		println("Pos", trackbar.Pos())
//	})
func NewTrackbar(parent Parent, opts *VarOptsTrackbar) *Trackbar {
	setUniqueCtrlId(&opts.ctrlId)
	me := &Trackbar{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsTrackbar{opts.ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.createWindow(opts.wndExStyle, "msctls_trackbar32", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, false)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		if opts.rangeMin != 0 || opts.rangeMax != 100 {
			me.SetRange(opts.rangeMin, opts.rangeMax)
		}
		if opts.pageSize != 0 {
			me.SetPageSize(opts.pageSize)
		} else {
			me.SetPageSize(opts.rangeMax / 5)
		}
		if opts.value != 0 {
			me.SetPos(opts.value)
		}
	})

	return me
}

// Instantiates a new [Trackbar] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
func NewTrackbarDlg(parent Parent, ctrlId uint16, layout LAY) *Trackbar {
	me := &Trackbar{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsTrackbar{ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.assignDialog(parent)
		parent.base().layout.Add(parent, me.hWnd, layout)
	})

	return me
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *Trackbar) On() *EventsTrackbar {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Retrieves the page with with [TBM_GETPAGESIZE].
//
// [TBM_GETPAGESIZE]: https://learn.microsoft.com/en-us/windows/win32/controls/tbm-getpagesize
func (me *Trackbar) PageSize() int {
	s, _ := me.Hwnd().SendMessage(co.TBM_GETPAGESIZE, 0, 0)
	return int(s)
}

// Retrieves the current position with [TBM_GETPOS].
//
// [TBM_GETPOS]: https://learn.microsoft.com/en-us/windows/win32/controls/tbm-getpos
func (me *Trackbar) Pos() int {
	p, _ := me.Hwnd().SendMessage(co.TBM_GETPOS, 0, 0)
	return int(p)
}

// Retrieves the minimum and maximum position values with [TBM_GETRANGEMIN],
// [TBM_GETRANGEMAX].
//
// [TBM_GETRANGEMIN]: https://learn.microsoft.com/en-us/windows/win32/controls/tbm-getrangemin
// [TBM_GETRANGEMAX]: https://learn.microsoft.com/en-us/windows/win32/controls/tbm-getrangemax
func (me *Trackbar) Range() (int, int) {
	min, _ := me.Hwnd().SendMessage(co.TBM_GETRANGEMIN, 0, 0)
	max, _ := me.Hwnd().SendMessage(co.TBM_GETRANGEMAX, 0, 0)
	return int(min), int(max)
}

// Sets the page size with [TBM_SETPAGESIZE].
//
// [TBM_SETPAGESIZE]: https://learn.microsoft.com/en-us/windows/win32/controls/tbm-setpagesize
func (me *Trackbar) SetPageSize(pageSize int) *Trackbar {
	me.Hwnd().SendMessage(co.TBM_SETPAGESIZE, 1, win.LPARAM(int32(pageSize)))
	return me
}

// Sets the current position witn [TBM_SETPOS].
//
// [TBM_SETPOS]: https://learn.microsoft.com/en-us/windows/win32/controls/tbm-setpos
func (me *Trackbar) SetPos(pos int) *Trackbar {
	me.Hwnd().SendMessage(co.TBM_SETPOS, 1, win.LPARAM(int32(pos)))
	return me
}

// Sets the minimum and maximum position values with [TBM_SETRANGEMIN] and
// [TBM_SETRANGEMAX].
//
// [TBM_SETRANGEMIN]: https://learn.microsoft.com/en-us/windows/win32/controls/tbm-setrangemin
// [TBM_SETRANGEMAX]: https://learn.microsoft.com/en-us/windows/win32/controls/tbm-setrangemax
func (me *Trackbar) SetRange(min, max int) *Trackbar {
	me.Hwnd().SendMessage(co.TBM_SETRANGEMIN, 1, win.LPARAM(int32(min)))
	me.Hwnd().SendMessage(co.TBM_SETRANGEMAX, 1, win.LPARAM(int32(max)))
	return me
}

// Options for [NewTrackbar]; returned by [OptsTrackbar].
type VarOptsTrackbar struct {
	ctrlId     uint16
	layout     LAY
	position   win.POINT
	size       win.SIZE
	ctrlStyle  co.TBS
	wndStyle   co.WS
	wndExStyle co.WS_EX
	pageSize   int
	rangeMin   int
	rangeMax   int
	value      int
}

// Options for [NewTrackbar].
func OptsTrackbar() *VarOptsTrackbar {
	return &VarOptsTrackbar{
		size:      win.SIZE{Cx: int32(DpiX(175)), Cy: int32(DpiY(28))},
		ctrlStyle: co.TBS_AUTOTICKS | co.TBS_HORZ,
		wndStyle:  co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP,
		rangeMax:  100,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsTrackbar) CtrlId(id uint16) *VarOptsTrackbar { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsTrackbar) Layout(l LAY) *VarOptsTrackbar { o.layout = l; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsTrackbar) Position(x, y int) *VarOptsTrackbar {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.Dpi(175, 28).
func (o *VarOptsTrackbar) Size(cx int, cy int) *VarOptsTrackbar {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Trackbar control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.TBS_AUTOTICKS | co.TBS_HORZ.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/trackbar-control-styles
func (o *VarOptsTrackbar) CtrlStyle(s co.TBS) *VarOptsTrackbar { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP.
func (o *VarOptsTrackbar) WndStyle(s co.WS) *VarOptsTrackbar { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
func (o *VarOptsTrackbar) WndExStyle(s co.WS_EX) *VarOptsTrackbar { o.wndExStyle = s; return o }

// Number of positions of page up/down.
//
// Defaults to rangeMax / 5.
func (o *VarOptsTrackbar) PageSize(p int) *VarOptsTrackbar { o.pageSize = p; return o }

// Minimum and maximum position value.
//
// Defaults to 0, 100.
func (o *VarOptsTrackbar) Range(min, max int) *VarOptsTrackbar {
	o.rangeMin = min
	o.rangeMax = max
	return o
}

// Current trackbar position.
//
// Defaults to 0.
func (o *VarOptsTrackbar) Value(v int) *VarOptsTrackbar { o.value = v; return o }

// Native [trackbar] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [trackbar]: https://learn.microsoft.com/en-us/windows/win32/controls/trackbar-controls
type EventsTrackbar struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [TRBN_THUMBPOSCHANGING] message handler.
//
// [TRBN_THUMBPOSCHANGING]: https://learn.microsoft.com/en-us/windows/win32/controls/trbn-thumbposchanging
func (me *EventsTrackbar) ThumbPosChanging(fun func(p *win.NMTRBTHUMBPOSCHANGING) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TRBN_THUMBPOSCHANGING, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMTRBTHUMBPOSCHANGING)(p)))
	})
}

// [WM_HSCROLL] message handler.
//
// [WM_HSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-hscroll
func (me *EventsTrackbar) WmHScroll(fun func(p WmScroll)) {
	me.parentEvents.Wm(co.WM_HSCROLL, func(p Wm) uintptr {
		fun(WmScroll{p})
		return me.parentEvents.defProcVal
	})
}

// [WM_VSCROLL] message handler.
//
// [WM_VSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-vscroll
func (me *EventsTrackbar) WmVScroll(fun func(p WmScroll)) {
	me.parentEvents.Wm(co.WM_VSCROLL, func(p Wm) uintptr {
		fun(WmScroll{p})
		return me.parentEvents.defProcVal
	})
}

// [NM_CUSTOMDRAW] message handler.
//
// [NM_CUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-customdraw-trackbar
func (me *EventsTrackbar) NmCustomDraw(fun func(p *win.NMCUSTOMDRAW) co.CDRF) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMCUSTOMDRAW)(p)))
	})
}

// [NM_RELEASEDCAPTURE] message handler.
//
// [NM_RELEASEDCAPTURE]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-trackbar-
func (me *EventsTrackbar) NmReleasedCapture(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}
