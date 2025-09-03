//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Native [up-down] control.
//
// [up-down]: https://learn.microsoft.com/en-us/windows/win32/controls/up-down-controls
type UpDown struct {
	_BaseCtrl
	events EventsUpDown
}

// Creates a new [UpDown] with [win.CreateWindowEx].
//
// If co.UDS_AUTOBUDDY control style is set, the UpDown will use the immediately
// previous control – usually an Edit – as its buddy, attaching itself to it.
//
// Example:
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	txt := ui.NewEdit( // will be taken as the buddy control
//		wndOwner,
//		ui.OptsEdit().
//			Position(ui.Dpi(10, 10)).
//			Width(ui.DpiY(50)),
//	)
//
//	upDn := ui.NewUpDown(
//		wndOwner,
//		ui.OptsUpDown().
//			Range(0, 10),
//	)
func NewUpDown(parent Parent, opts *VarOptsUpDown) *UpDown {
	setUniqueCtrlId(&opts.ctrlId)
	me := &UpDown{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsUpDown{opts.ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.createWindow(opts.wndExStyle, "msctls_updown32", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position,
			win.SIZE{Cy: opts.height}, parent, true)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		if opts.radix != 0 {
			me.SetRadix(opts.radix)
		}
		if opts.rangeMin != 0 || opts.rangeMax != 100 {
			me.SetRange(opts.rangeMin, opts.rangeMax)
		}
		if opts.value != 0 {
			me.SetValue(opts.value)
		}
	})

	return me
}

// Instantiates a new [UpDown] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
//
// If co.UDS_AUTOBUDDY control style is set, the UpDown will use the immediately
// previous control – usually an Edit – as its buddy, attaching itself to it.
func NewUpDownDlg(parent Parent, ctrlId uint16, layout LAY) *UpDown {
	me := &UpDown{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsUpDown{ctrlId, &parent.base().userEvents},
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
func (me *UpDown) On() *EventsUpDown {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Retrieves the radix with [UDM_GETBASE]. Can be 10 (decimal) or 16
// (hexadecimal).
//
// [UDM_GETBASE]: https://learn.microsoft.com/en-us/windows/win32/controls/udm-getbase
func (me *UpDown) Radix() int {
	ret, _ := me.hWnd.SendMessage(co.UDM_GETBASE, 0, 0)
	return int(ret)
}

// Retrieves the range with [UDM_GETRANGE32].
//
// [UDM_GETRANGE32]: https://learn.microsoft.com/en-us/windows/win32/controls/udm-getrange32
func (me *UpDown) Range() (int, int) {
	var min, max int32
	me.hWnd.SendMessage(co.UDM_GETRANGE32,
		win.WPARAM(unsafe.Pointer(&min)), win.LPARAM(unsafe.Pointer(&max)))
	return int(min), int(max)
}

// Sets the radix with [UDM_SETBASE]. Must be 10 (decimal) or 16 (hexadecimal).
//
// Returns the same object, so further operations can be chained.
//
// [UDM_SETBASE]: https://learn.microsoft.com/en-us/windows/win32/controls/udm-setbase
func (me *UpDown) SetRadix(radix int) *UpDown {
	me.hWnd.SendMessage(co.UDM_SETBASE, win.WPARAM(uint32(radix)), 0)
	return me
}

// Sets the range with [UDM_SETRANGE32].
//
// Returns the same object, so further operations can be chained.
//
// [UDM_SETRANGE32]: https://learn.microsoft.com/en-us/windows/win32/controls/udm-setrange32
func (me *UpDown) SetRange(min, max int) *UpDown {
	me.hWnd.SendMessage(co.UDM_SETRANGE32,
		win.WPARAM(int32(min)), win.LPARAM(int32(max)))
	return me
}

// Sets the value with [UDM_SETPOS32].
//
// Returns the same object, so further operations can be chained.
//
// [UDM_SETPOS32]: https://learn.microsoft.com/en-us/windows/win32/controls/udm-setpos32
func (me *UpDown) SetValue(val int) *UpDown {
	me.hWnd.SendMessage(co.UDM_SETPOS32, 0, win.LPARAM(int32(val)))
	return me
}

// Retrieves the value with [UDM_GETPOS32], which may be invalid.
//
// [UDM_GETPOS32]: https://learn.microsoft.com/en-us/windows/win32/controls/udm-getpos32
func (me *UpDown) Value() (int, bool) {
	var valid int32 // BOOL
	ret, err := me.hWnd.SendMessage(co.UDM_GETPOS32,
		0, win.LPARAM(unsafe.Pointer(&valid)))
	if valid == 0 || err != nil {
		return 0, false
	}
	return int(ret), true
}

// Options for [NewUpDown]; returned by [OptsUpDown].
type VarOptsUpDown struct {
	ctrlId     uint16
	layout     LAY
	position   win.POINT
	height     int32
	ctrlStyle  co.UDS
	wndStyle   co.WS
	wndExStyle co.WS_EX
	radix      int
	rangeMin   int
	rangeMax   int
	value      int
}

// Options for [NewUpDown].
func OptsUpDown() *VarOptsUpDown {
	return &VarOptsUpDown{
		height:    int32(DpiY(50)),
		ctrlStyle: co.UDS_AUTOBUDDY | co.UDS_SETBUDDYINT | co.UDS_ALIGNRIGHT | co.UDS_ARROWKEYS | co.UDS_HOTTRACK,
		wndStyle:  co.WS_CHILD | co.WS_VISIBLE,
		rangeMax:  100,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsUpDown) CtrlId(id uint16) *VarOptsUpDown { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsUpDown) Layout(l LAY) *VarOptsUpDown { o.layout = l; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsUpDown) Position(x, y int) *VarOptsUpDown {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control height in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.DpiY(50).
func (o *VarOptsUpDown) Height(h int) *VarOptsUpDown { o.height = int32(h); return o }

// Up-down control [style], passed to [win.CreateWindowEx].
//
// If co.UDS_AUTOBUDDY is set, the UpDown will use the immediately previous
// control – usually an Edit – as its buddy, attaching itself to it.
//
// Defaults to co.UDS_AUTOBUDDY | co.UDS_SETBUDDYINT | co.UDS_ALIGNRIGHT | co.UDS_ARROWKEYS | co.UDS_HOTTRACK.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/up-down-control-styles
func (o *VarOptsUpDown) CtrlStyle(s co.UDS) *VarOptsUpDown { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
func (o *VarOptsUpDown) WndStyle(s co.WS) *VarOptsUpDown { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
func (o *VarOptsUpDown) WndExStyle(s co.WS_EX) *VarOptsUpDown { o.wndExStyle = s; return o }

// Value radix. Must be 10 (decimal) or 16 (hexadecimal).
//
// Defaults to 10.
func (o *VarOptsUpDown) Radix(r int) *VarOptsUpDown { o.radix = r; return o }

// Minimum and maximum range.
//
// Defaults to 0, 100.
func (o *VarOptsUpDown) Range(min, max int) *VarOptsUpDown {
	o.rangeMin = min
	o.rangeMax = max
	return o
}

// Current value.
//
// Defaults to 0.
func (o *VarOptsUpDown) Value(v int) *VarOptsUpDown { o.value = v; return o }

// Native [up-down] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [up-down]: https://learn.microsoft.com/en-us/windows/win32/controls/up-down-controls
type EventsUpDown struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [NM_RELEASEDCAPTURE] message handler.
//
// [NM_RELEASEDCAPTURE]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-up-down-
func (me *EventsUpDown) NmReleasedCapture(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [UDN_DELTAPOS] message handler.
//
// [UDN_DELTAPOS]: https://learn.microsoft.com/en-us/windows/win32/controls/udn-deltapos
func (me *EventsUpDown) UdnDeltaPos(fun func(p *win.NMUPDOWN) int) {
	me.parentEvents.WmNotify(me.ctrlId, co.UDN_DELTAPOS, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMUPDOWN)(p)))
	})
}
