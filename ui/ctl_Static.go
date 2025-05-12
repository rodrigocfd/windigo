//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [static] (label) control.
//
// [static]: https://learn.microsoft.com/en-us/windows/win32/controls/about-static-controls
type Static struct {
	_BaseCtrl
	events EventsStatic
}

// Creates a new [Static] with [win.CreateWindowEx].
func NewStatic(parent Parent, opts *VarOptsStatic) *Static {
	setUniqueCtrlId(&opts.ctrlId)
	me := &Static{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsStatic{opts.ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.WmCreate(func(_ WmCreate) int {
		if opts.size.Cx == 0 && opts.size.Cy == 0 {
			opts.size, _ = calcTextBoundBox(utl.RemoveAccelAmpersands(opts.text))
		}
		me.createWindow(opts.wndExStyle, "STATIC", opts.text,
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, true)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		return 0 // ignored
	})

	return me
}

// Instantiates a new [Static] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
func NewStaticDlg(parent Parent, ctrlId uint16, layout LAY) *Static {
	me := &Static{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsStatic{ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.WmInitDialog(func(_ WmInitDialog) bool {
		me.assignDialog(parent)
		parent.base().layout.Add(parent, me.hWnd, layout)
		return true // ignored
	})

	return me
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *Static) On() *EventsStatic {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Calls [win.HWND.SetWindowText] and resizes the control to exactly fit it.
func (me *Static) SetTextAndResize(text string) *Static {
	me.hWnd.SetWindowText(text)
	boundBox, _ := calcTextBoundBox(utl.RemoveAccelAmpersands(text))
	me.hWnd.SetWindowPos(win.HWND(0), 0, 0,
		uint(boundBox.Cx), uint(boundBox.Cy), co.SWP_NOZORDER|co.SWP_NOMOVE)
	return me
}

// Calls [win.HWND.GetWindowText].
func (me *Static) Text() string {
	t, _ := me.hWnd.GetWindowText()
	return t
}

// Options for [NewStatic]; returned by [OptsStatic].
type VarOptsStatic struct {
	ctrlId     uint16
	layout     LAY
	text       string
	position   win.POINT
	size       win.SIZE
	ctrlStyle  co.SS
	wndStyle   co.WS
	wndExStyle co.WS_EX
}

// Options for [NewStatic].
func OptsStatic() *VarOptsStatic {
	return &VarOptsStatic{
		ctrlStyle: co.SS_LEFT | co.SS_NOTIFY,
		wndStyle:  co.WS_CHILD | co.WS_VISIBLE,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsStatic) CtrlId(id uint16) *VarOptsStatic { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsStatic) Layout(l LAY) *VarOptsStatic { o.layout = l; return o }

// Text to be displayed, passed to [win.CreateWindowEx].
//
// Defaults to empty string.
func (o *VarOptsStatic) Text(t string) *VarOptsStatic { o.text = t; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsStatic) Position(x, y int) *VarOptsStatic {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [win.CreateWindowEx].
//
// Defaults to fit current text.
func (o *VarOptsStatic) Size(cx int, cy int) *VarOptsStatic {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Static control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.SS_LEFT | co.SS_NOTIFY.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/static-control-styles
func (o *VarOptsStatic) CtrlStyle(s co.SS) *VarOptsStatic { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE.
func (o *VarOptsStatic) WndStyle(s co.WS) *VarOptsStatic { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
func (o *VarOptsStatic) WndExStyle(s co.WS_EX) *VarOptsStatic { o.wndExStyle = s; return o }

// Native [static] (label) control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [static]: https://learn.microsoft.com/en-us/windows/win32/controls/about-static-controls
type EventsStatic struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [STN_CLICKED] message handler.
//
// [STN_CLICKED]: https://learn.microsoft.com/en-us/windows/win32/controls/stn-clicked
func (me *EventsStatic) StnClicked(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.STN_CLICKED, fun)
}

// [STN_DBLCLK] message handler.
//
// [STN_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/stn-dblclk
func (me *EventsStatic) StnDblClk(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.STN_DBLCLK, fun)
}

// [STN_DISABLE] message handler.
//
// [STN_DISABLE]: https://learn.microsoft.com/en-us/windows/win32/controls/stn-disable
func (me *EventsStatic) StnDisable(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.STN_DISABLE, fun)
}

// [STN_ENABLE] message handler.
//
// [STN_ENABLE]: https://learn.microsoft.com/en-us/windows/win32/controls/stn-enable
func (me *EventsStatic) StnEnable(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.STN_ENABLE, fun)
}
