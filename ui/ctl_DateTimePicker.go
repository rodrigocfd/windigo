//go:build windows

package ui

import (
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [date and time picker] control.
//
// [date and time picker]: https://learn.microsoft.com/en-us/windows/win32/controls/date-and-time-picker-controls
type DateTimePicker struct {
	_BaseCtrl
	events EventsDateTimePicker
}

// Creates a new [DateTimePicker] with [win.CreateWindowEx].
//
// # Example
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	dtp := ui.NewDateTimePicker(
//		wndOwner,
//		ui.OptsDateTimePicker().
//			Position(ui.Dpi(210, 20)).
//			Value(time.Date(1981, 4, 26, 5, 0, 0, 0, time.Local)),
//	)
func NewDateTimePicker(parent Parent, opts *VarOptsDateTimePicker) *DateTimePicker {
	setUniqueCtrlId(&opts.ctrlId)
	me := &DateTimePicker{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsDateTimePicker{opts.ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.createWindow(opts.wndExStyle, "SysDateTimePick32", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, true)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		if !opts.value.IsZero() {
			me.SetTime(opts.value)
		}
	})

	return me
}

// Instantiates a new [DateTimePicker] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
//
// # Example
//
//	const ID_DTP uint16 = 0x100
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	dtp := ui.NewDateTimePickerDlg(
//		wndOwner, ID_DTP, ui.LAY_NONE_NONE)
func NewDateTimePickerDlg(parent Parent, ctrlId uint16, layout LAY) *DateTimePicker {
	me := &DateTimePicker{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsDateTimePicker{ctrlId, &parent.base().userEvents},
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
func (me *DateTimePicker) On() *EventsDateTimePicker {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Calls [DTM_SETSYSTEMTIME].
//
// Returns the same object, so further operations can be chained.
//
// [DTM_SETSYSTEMTIME]: https://learn.microsoft.com/en-us/windows/win32/controls/dtm-setsystemtime
func (me *DateTimePicker) SetTime(newTime time.Time) *DateTimePicker {
	var st win.SYSTEMTIME
	st.SetTime(newTime)
	me.hWnd.SendMessage(co.DTM_SETSYSTEMTIME,
		win.WPARAM(co.GDT_VALID), win.LPARAM(unsafe.Pointer(&st)))
	return me
}

// Calls [DTM_GETSYSTEMTIME].
//
// Panics on error.
//
// [DTM_GETSYSTEMTIME]: https://learn.microsoft.com/en-us/windows/win32/controls/dtm-getsystemtime
func (me *DateTimePicker) Time() time.Time {
	var st win.SYSTEMTIME
	ret, _ := me.hWnd.SendMessage(co.DTM_GETSYSTEMTIME,
		0, win.LPARAM(unsafe.Pointer(&st)))
	if co.GDT(ret) != co.GDT_VALID {
		panic("DTM_GETSYSTEMTIME failed.")
	}
	return st.ToTime()
}

// Options for [NewDateTimePicker]; returned by [OptsDateTimePicker].
type VarOptsDateTimePicker struct {
	ctrlId     uint16
	layout     LAY
	position   win.POINT
	size       win.SIZE
	ctrlStyle  co.DTS
	wndStyle   co.WS
	wndExStyle co.WS_EX
	value      time.Time
}

// Options for [NewDateTimePicker].
func OptsDateTimePicker() *VarOptsDateTimePicker {
	return &VarOptsDateTimePicker{
		size:       win.SIZE{Cx: int32(DpiX(230)), Cy: int32(DpiY(23))},
		ctrlStyle:  co.DTS_LONGDATEFORMAT,
		wndStyle:   co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
		wndExStyle: co.WS_EX_LEFT | co.WS_EX_CLIENTEDGE,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsDateTimePicker) CtrlId(id uint16) *VarOptsDateTimePicker { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsDateTimePicker) Layout(l LAY) *VarOptsDateTimePicker { o.layout = l; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsDateTimePicker) Position(x, y int) *VarOptsDateTimePicker {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.Dpi(230, 23).
func (o *VarOptsDateTimePicker) Size(cx int, cy int) *VarOptsDateTimePicker {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Date and time picker control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.DTS_LONGDATEFORMAT.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/date-and-time-picker-control-styles
func (o *VarOptsDateTimePicker) CtrlStyle(s co.DTS) *VarOptsDateTimePicker { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *VarOptsDateTimePicker) WndStyle(s co.WS) *VarOptsDateTimePicker { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT | co.WS_EX_CLIENTEDGE.
func (o *VarOptsDateTimePicker) WndExStyle(s co.WS_EX) *VarOptsDateTimePicker {
	o.wndExStyle = s
	return o
}

// Initial value.
//
// Defaults to [time.Now].
func (o *VarOptsDateTimePicker) Value(t time.Time) *VarOptsDateTimePicker { o.value = t; return o }

// Native [date and time picker] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [date and time picker]: https://learn.microsoft.com/en-us/windows/win32/controls/date-and-time-picker-controls
type EventsDateTimePicker struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [DTN_CLOSEUP] message handler.
//
// [DTN_CLOSEUP]: https://learn.microsoft.com/en-us/windows/win32/controls/dtn-closeup
func (me *EventsDateTimePicker) DtnCloseUp(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.DTN_CLOSEUP, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [DTN_DATETIMECHANGE] message handler.
//
// [DTN_DATETIMECHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/dtn-datetimechange
func (me *EventsDateTimePicker) DtnDateTimeChange(fun func(p *win.NMDATETIMECHANGE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.DTN_DATETIMECHANGE, func(p unsafe.Pointer) uintptr {
		fun((*win.NMDATETIMECHANGE)(p))
		return me.parentEvents.defProcVal
	})
}

// [DTN_DROPDOWN] message handler.
//
// [DTN_DROPDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/dtn-dropdown
func (me *EventsDateTimePicker) DtnDropDown(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.DTN_DROPDOWN, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [DTN_FORMAT] message handler.
//
// [DTN_FORMAT]: https://learn.microsoft.com/en-us/windows/win32/controls/dtn-format
func (me *EventsDateTimePicker) DtnFormat(fun func(p *win.NMDATETIMEFORMAT)) {
	me.parentEvents.WmNotify(me.ctrlId, co.DTN_FORMAT, func(p unsafe.Pointer) uintptr {
		fun((*win.NMDATETIMEFORMAT)(p))
		return me.parentEvents.defProcVal
	})
}

// [DTN_FORMATQUERY] message handler.
//
// [DTN_FORMATQUERY]: https://learn.microsoft.com/en-us/windows/win32/controls/dtn-formatquery
func (me *EventsDateTimePicker) DtnFormatQuery(fun func(p *win.NMDATETIMEFORMATQUERY)) {
	me.parentEvents.WmNotify(me.ctrlId, co.DTN_FORMATQUERY, func(p unsafe.Pointer) uintptr {
		fun((*win.NMDATETIMEFORMATQUERY)(p))
		return me.parentEvents.defProcVal
	})
}

// [DTN_USERSTRING] message handler.
//
// [DTN_USERSTRING]: https://learn.microsoft.com/en-us/windows/win32/controls/dtn-userstring
func (me *EventsDateTimePicker) DtnUserString(fun func(p *win.NMDATETIMESTRING)) {
	me.parentEvents.WmNotify(me.ctrlId, co.DTN_USERSTRING, func(p unsafe.Pointer) uintptr {
		fun((*win.NMDATETIMESTRING)(p))
		return me.parentEvents.defProcVal
	})
}

// [DTN_WMKEYDOWN] message handler.
//
// [DTN_WMKEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/dtn-wmkeydown
func (me *EventsDateTimePicker) DtnWmKeyDown(fun func(p *win.NMDATETIMEWMKEYDOWN)) {
	me.parentEvents.WmNotify(me.ctrlId, co.DTN_WMKEYDOWN, func(p unsafe.Pointer) uintptr {
		fun((*win.NMDATETIMEWMKEYDOWN)(p))
		return me.parentEvents.defProcVal
	})
}

// [NM_KILLFOCUS] message handler.
//
// [NM_KILLFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-killfocus-date-time
func (me *EventsDateTimePicker) NmKillFocus(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_KILLFOCUS, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [NM_SETFOCUS] message handler.
//
// [NM_SETFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-setfocus-date-time-
func (me *EventsDateTimePicker) NmSetFocus(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_SETFOCUS, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}
