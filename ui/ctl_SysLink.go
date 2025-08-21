//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [syslink] control.
//
// [syslink]: https://learn.microsoft.com/en-us/windows/win32/controls/syslink-control-entry
type SysLink struct {
	_BaseCtrl
	events EventsSysLink
}

// Creates a new [SysLink] with [win.CreateWindowEx].
func NewSysLink(parent Parent, opts *VarOptsSysLink) *SysLink {
	setUniqueCtrlId(&opts.ctrlId)
	me := &SysLink{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsSysLink{opts.ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		if opts.size.Cx == 0 && opts.size.Cy == 0 {
			opts.size, _ = calcTextBoundBox(utl.RemoveAccelAmpersands(utl.RemoveHtmlAnchor(opts.text)))
		}
		me.createWindow(opts.wndExStyle, "SysLink", opts.text,
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, true)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
	})

	return me
}

// Instantiates a new [SysLink] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
func NewSysLinkDlg(parent Parent, ctrlId uint16, layout LAY) *SysLink {
	me := &SysLink{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsSysLink{ctrlId, &parent.base().userEvents},
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
func (me *SysLink) On() *EventsSysLink {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Calls [win.HWND.SetWindowText] and resizes the control to exactly fit it.
//
// Example:
//
//	var link ui.SysLink // initialized somewhere
//
//	link.SetTextAndResize(
//		"Link <a href=\"https://google.com\">here</a>")
func (me *SysLink) SetTextAndResize(text string) *SysLink {
	me.hWnd.SetWindowText(text)
	boundBox, _ := calcTextBoundBox(utl.RemoveAccelAmpersands(utl.RemoveHtmlAnchor(text)))
	me.hWnd.SetWindowPos(win.HWND(0), 0, 0,
		uint(boundBox.Cx), uint(boundBox.Cy), co.SWP_NOZORDER|co.SWP_NOMOVE)
	return me
}

// Calls [win.HWND.GetWindowText].
func (me *SysLink) Text() string {
	t, _ := me.hWnd.GetWindowText()
	return t
}

// Options for [NewSysLink]; returned by [OptsSysLink].
type VarOptsSysLink struct {
	ctrlId     uint16
	layout     LAY
	text       string
	position   win.POINT
	size       win.SIZE
	ctrlStyle  co.LWS
	wndStyle   co.WS
	wndExStyle co.WS_EX
}

// Options for [NewSysLink]
func OptsSysLink() *VarOptsSysLink {
	return &VarOptsSysLink{
		ctrlStyle: co.LWS_TRANSPARENT,
		wndStyle:  co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsSysLink) CtrlId(id uint16) *VarOptsSysLink { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsSysLink) Layout(l LAY) *VarOptsSysLink { o.layout = l; return o }

// Text to be displayed, passed to [win.CreateWindowEx]. URLs are embedded using
// HTML anchor syntax.
//
// Defaults to empty string.
//
// Example:
//
//	ui.OptsSysLink().
//		Text("Link <a href=\"https://google.com\">here</a>")
func (o *VarOptsSysLink) Text(t string) *VarOptsSysLink { o.text = t; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsSysLink) Position(x, y int) *VarOptsSysLink {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [win.CreateWindowEx].
//
// Defaults to fit current text.
func (o *VarOptsSysLink) Size(cx int, cy int) *VarOptsSysLink {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// SysLink control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.LWS_TRANSPARENT.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/syslink-control-styles
func (o *VarOptsSysLink) CtrlStyle(s co.LWS) *VarOptsSysLink { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP.
func (o *VarOptsSysLink) WndStyle(s co.WS) *VarOptsSysLink { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
func (o *VarOptsSysLink) WndExStyle(s co.WS_EX) *VarOptsSysLink { o.wndExStyle = s; return o }

// Native [syslink] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [syslink]: https://learn.microsoft.com/en-us/windows/win32/controls/syslink-control-entry
type EventsSysLink struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [NM_CLICK] message handler.
//
// [NM_CLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-click-syslink
func (me *EventsSysLink) NmClick(fun func(p *win.NMLINK)) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CLICK, func(p unsafe.Pointer) uintptr {
		fun((*win.NMLINK)(p))
		return me.parentEvents.defProcVal
	})
}
