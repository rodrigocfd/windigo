//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [progress bar] control.
//
// [progress bar]: https://learn.microsoft.com/en-us/windows/win32/controls/progress-bar-control
type ProgressBar struct {
	_BaseCtrl
	isMarquee bool
}

// Creates a new [ProgressBar] with [win.CreateWindowEx].
func NewProgressBar(parent Parent, opts *VarOptsProgressBar) *ProgressBar {
	setUniqueCtrlId(&opts.ctrlId)
	me := &ProgressBar{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		isMarquee: false,
	}

	parent.base().beforeUserEvents.WmCreate(func(_ WmCreate) int {
		me.createWindow(opts.wndExStyle, "msctls_progress32", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, false)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		if opts.rangeMin != 0 || opts.rangeMax != 100 {
			me.SetRange(opts.rangeMin, opts.rangeMax)
		}
		if opts.value != 0 {
			me.SetPos(opts.value)
		}
		if opts.state != 0 {
			me.SetState(opts.state)
		}
		return 0 // ignored
	})

	return me
}

// Instantiates a new [ProgressBar] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
func NewProgressBarDlg(parent Parent, ctrlId uint16, layout LAY) *ProgressBar {
	me := &ProgressBar{
		_BaseCtrl: newBaseCtrl(ctrlId),
	}

	parent.base().beforeUserEvents.WmInitDialog(func(_ WmInitDialog) bool {
		me.assignDialog(parent)
		parent.base().layout.Add(parent, me.hWnd, layout)
		return true // ignored
	})

	return me
}

// Retrieves the current position with [PBM_GETPOS].
//
// [PBM_GETPOS]: https://learn.microsoft.com/en-us/windows/win32/controls/pbm-getpos
func (me *ProgressBar) Pos() int {
	pos, _ := me.hWnd.SendMessage(co.PBM_GETPOS, 0, 0)
	return int(pos)
}

// Retrieves the range with [PBM_GETRANGE].
//
// [PBM_GETRANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/pbm-getrange
func (me *ProgressBar) Range() (int, int) {
	var r win.PBRANGE
	me.hWnd.SendMessage(co.PBM_GETRANGE, 0, win.LPARAM(unsafe.Pointer(&r)))
	return int(r.ILow), int(r.IHigh)
}

// Sets indeterminate state, a graphic animation going back and forth, by
// toggling the [PBS_MARQUEE] style.
//
// Returns the same object, so further operations can be chained.
//
// [PBS_MARQUEE]: https://learn.microsoft.com/en-us/windows/win32/controls/progress-bar-control-styles
func (me *ProgressBar) SetMarquee(isMarquee bool) *ProgressBar {
	curStyle, _ := me.hWnd.Style()

	if isMarquee {
		me.hWnd.SetWindowLongPtr(co.GWLP_STYLE,
			uintptr(curStyle)|uintptr(co.PBS_MARQUEE))
	}

	me.hWnd.SendMessage(co.PBM_SETMARQUEE,
		win.WPARAM(wutil.BoolToUintptr(isMarquee)), 0)

	if !isMarquee {
		me.hWnd.SetWindowLongPtr(co.GWLP_STYLE,
			uintptr(curStyle)&^uintptr(co.PBS_MARQUEE))
	}

	me.isMarquee = isMarquee
	return me
}

// Sets the current position with [PBM_SETPOS].
//
// Returns the same object, so further operations can be chained.
//
// [PBM_SETPOS]: https://learn.microsoft.com/en-us/windows/win32/controls/pbm-setpos
func (me *ProgressBar) SetPos(pos int) *ProgressBar {
	if me.isMarquee {
		me.SetMarquee(false) // avoid crash
	}
	me.hWnd.SendMessage(co.PBM_SETPOS, win.WPARAM(int32(pos)), 0)
	return me
}

// Sets the range with [PBM_SETRANGE32].
//
// Returns the same object, so further operations can be chained.
//
// [PBM_SETRANGE32]: https://learn.microsoft.com/en-us/windows/win32/controls/pbm-setrange32
func (me *ProgressBar) SetRange(min, max int) *ProgressBar {
	me.hWnd.SendMessage(co.PBM_SETRANGE32,
		win.WPARAM(int32(min)), win.LPARAM(int32(max)))
	return me
}

// Sets the state with [PBM_SETSTATE].
//
// Returns the same object, so further operations can be chained.
//
// [PBM_SETSTATE]: https://learn.microsoft.com/en-us/windows/win32/controls/pbm-setstate
func (me *ProgressBar) SetState(state co.PBST) *ProgressBar {
	me.hWnd.SendMessage(co.PBM_SETSTATE, win.WPARAM(state), 0)
	return me
}

// Options for [NewProgressBar]; returned by [OptsProgressBar].
type VarOptsProgressBar struct {
	ctrlId     uint16
	layout     LAY
	position   win.POINT
	size       win.SIZE
	ctrlStyle  co.PBS
	wndStyle   co.WS
	wndExStyle co.WS_EX
	rangeMin   int
	rangeMax   int
	value      int
	state      co.PBST
}

// Options for [NewProgressBar].
func OptsProgressBar() *VarOptsProgressBar {
	return &VarOptsProgressBar{
		size:      win.SIZE{Cx: int32(DpiX(140)), Cy: int32(DpiY(26))},
		ctrlStyle: co.PBS_SMOOTH,
		wndStyle:  co.WS_CHILD | co.WS_VISIBLE,
		rangeMax:  100,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsProgressBar) CtrlId(id uint16) *VarOptsProgressBar { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsProgressBar) Layout(l LAY) *VarOptsProgressBar { o.layout = l; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsProgressBar) Position(x, y int) *VarOptsProgressBar {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.Dpi(140, 26).
func (o *VarOptsProgressBar) Size(cx int, cy int) *VarOptsProgressBar {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Progress bar control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.PBS_SMOOTH.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/progress-bar-control-styles
func (o *VarOptsProgressBar) CtrlStyle(s co.PBS) *VarOptsProgressBar { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
func (o *VarOptsProgressBar) WndStyle(s co.WS) *VarOptsProgressBar { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
func (o *VarOptsProgressBar) WndExStyle(s co.WS_EX) *VarOptsProgressBar { o.wndExStyle = s; return o }

// Minimum and maximum range.
//
// Defaults to 0, 100.
func (o *VarOptsProgressBar) Range(min, max int) *VarOptsProgressBar {
	o.rangeMin = min
	o.rangeMax = max
	return o
}

// Current progress position.
//
// Defaults to 0.
func (o *VarOptsProgressBar) Value(v int) *VarOptsProgressBar { o.value = v; return o }

// State (normal, error or paused).
//
// Defaults to normal.
func (o *VarOptsProgressBar) State(s co.PBST) *VarOptsProgressBar { o.state = s; return o }
