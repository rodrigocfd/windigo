//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// Native [edit] (text box) control.
//
// [edit]: https://learn.microsoft.com/en-us/windows/win32/controls/about-edit-controls
type Edit struct {
	_BaseCtrl
	events EventsEdit
}

// Creates a new [Edit] with [win.CreateWindowEx].
//
// Example:
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	myEdit := ui.NewEdit(
//		wndOwner,
//		ui.OptsEdit().
//			Position(ui.Dpi(20, 10)),
//	)
func NewEdit(parent Parent, opts *VarOptsEdit) *Edit {
	setUniqueCtrlId(&opts.ctrlId)
	me := &Edit{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsEdit{opts.ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		me.createWindow(opts.wndExStyle, "EDIT", opts.text,
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, true)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
	})

	return me
}

// Instantiates a new [Edit] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
//
// Example:
//
//	const ID_TXT uint16 = 0x100
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	txt := ui.NewEditDlg(
//		wndOwner, ID_TXT, ui.LAY_NONE_NONE)
func NewEditDlg(parent Parent, ctrlId uint16, layout LAY) *Edit {
	me := &Edit{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsEdit{ctrlId, &parent.base().userEvents},
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
func (me *Edit) On() *EventsEdit {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Calls [EM_HIDEBALLOONTIP].
//
// Returns the same object, so further operations can be chained.
//
// [EM_HIDEBALLOONTIP]: https://learn.microsoft.com/en-us/windows/win32/controls/em-hideballoontip
func (me *Edit) HideBalloonTip() *Edit {
	me.hWnd.SendMessage(co.EM_HIDEBALLOONTIP, 0, 0)
	return me
}

// Calls [EM_SETLIMITTEXT].
//
// Returns the same object, so further operations can be chained.
//
// [EM_SETLIMITTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/em-setlimittext
func (me *Edit) LimitText(maxChars uint) *Edit {
	me.hWnd.SendMessage(co.EM_SETLIMITTEXT, win.WPARAM(maxChars), 0)
	return me
}

// Calls [EM_SETSEL].
//
// If the start is 0 and the end is -1, all the text in the edit control is
// selected. If the start is -1, any current selection is deselected.
//
// Returns the same object, so further operations can be chained.
//
// [EM_SETSEL]: https://learn.microsoft.com/en-us/windows/win32/controls/em-setsel
func (me *Edit) SetSelection(startPos, endPos int) *Edit {
	me.hWnd.SendMessage(co.EM_SETSEL, win.WPARAM(startPos), win.LPARAM(endPos))
	return me
}

// Calls [win.HWND.SetWindowText].
//
// Returns the same object, so further operations can be chained.
func (me *Edit) SetText(text string) *Edit {
	me.hWnd.SetWindowText(text)
	return me
}

// Calls [EM_SHOWBALLOONTIP].
//
// Returns the same object, so further operations can be chained.
//
// [EM_SHOWBALLOONTIP]: https://learn.microsoft.com/en-us/windows/win32/controls/em-showballoontip
func (me *Edit) ShowBalloonTip(title, text string, icon co.TTI) *Edit {
	var wTitle, wText wstr.BufEncoder
	ebt := win.EDITBALLOONTIP{
		PszTitle: (*uint16)(wTitle.AllowEmpty(title)),
		PszText:  (*uint16)(wText.AllowEmpty(text)),
		TtiIcon:  icon,
	}
	ebt.SetCbStruct()
	me.hWnd.SendMessage(co.EM_SHOWBALLOONTIP, 0, win.LPARAM(unsafe.Pointer(&ebt)))
	return me
}

// Calls [win.HWND.GetWindowText].
func (me *Edit) Text() string {
	t, _ := me.hWnd.GetWindowText()
	return t
}

// Options for [NewEdit]; returned by [OptsEdit].
type VarOptsEdit struct {
	ctrlId     uint16
	layout     LAY
	text       string
	position   win.POINT
	size       win.SIZE
	ctrlStyle  co.ES
	wndStyle   co.WS
	wndExStyle co.WS_EX
}

// Options for [NewEdit].
func OptsEdit() *VarOptsEdit {
	return &VarOptsEdit{
		size:       win.SIZE{Cx: int32(DpiX(100)), Cy: int32(DpiY(23))},
		ctrlStyle:  co.ES_AUTOHSCROLL | co.ES_NOHIDESEL,
		wndStyle:   co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP,
		wndExStyle: co.WS_EX_LEFT | co.WS_EX_CLIENTEDGE,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsEdit) CtrlId(id uint16) *VarOptsEdit { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsEdit) Layout(l LAY) *VarOptsEdit { o.layout = l; return o }

// Text to be displayed, passed to [win.CreateWindowEx].
//
// Defaults to empty string.
func (o *VarOptsEdit) Text(t string) *VarOptsEdit { o.text = t; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsEdit) Position(x, y int) *VarOptsEdit {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control width in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.DpiX(100).
func (o *VarOptsEdit) Width(w int) *VarOptsEdit { o.size.Cx = int32(w); return o }

// Control height in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.DpiY(23).
func (o *VarOptsEdit) Height(h int) *VarOptsEdit { o.size.Cy = int32(h); return o }

// Edit control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.ES_AUTOHSCROLL | co.ES_NOHIDESEL.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/edit-control-styles
func (o *VarOptsEdit) CtrlStyle(s co.ES) *VarOptsEdit { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP.
func (o *VarOptsEdit) WndStyle(s co.WS) *VarOptsEdit { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT | co.WS_EX_CLIENTEDGE.
func (o *VarOptsEdit) WndExStyle(s co.WS_EX) *VarOptsEdit { o.wndExStyle = s; return o }

// Native [edit] (text box) control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [edit]: https://learn.microsoft.com/en-us/windows/win32/controls/about-edit-controls
type EventsEdit struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [EN_ALIGN_LTR_EC] message handler.
//
// [EN_ALIGN_LTR_EC]: https://learn.microsoft.com/en-us/windows/win32/controls/en-align-ltr-ec
func (me *EventsEdit) EnAlignLtrEc(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.EN_ALIGN_LTR_EC, fun)
}

// [EN_ALIGN_RTL_EC] message handler.
//
// [EN_ALIGN_RTL_EC]: https://learn.microsoft.com/en-us/windows/win32/controls/en-align-rtl-ec
func (me *EventsEdit) EnAlignRtlEc(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.EN_ALIGN_RTL_EC, fun)
}

// [EN_CHANGE] message handler.
//
// [EN_CHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/en-change
func (me *EventsEdit) EnChange(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.EN_CHANGE, fun)
}

// [EN_ERRSPACE] message handler.
//
// [EN_ERRSPACE]: https://learn.microsoft.com/en-us/windows/win32/controls/en-errspace
func (me *EventsEdit) EnErrSpace(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.EN_ERRSPACE, fun)
}

// [EN_HSCROLL] message handler.
//
// [EN_HSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/en-hscroll
func (me *EventsEdit) EnHScroll(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.EN_HSCROLL, fun)
}

// [EN_KILLFOCUS] message handler.
//
// [EN_KILLFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/en-killfocus
func (me *EventsEdit) EnKillFocus(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.EN_KILLFOCUS, fun)
}

// [EN_MAXTEXT] message handler.
//
// [EN_MAXTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/en-maxtext
func (me *EventsEdit) EnMaxText(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.EN_MAXTEXT, fun)
}

// [EN_SETFOCUS] message handler.
//
// [EN_SETFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/en-setfocus
func (me *EventsEdit) EnSetFocus(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.EN_SETFOCUS, fun)
}

// [EN_UPDATE] message handler.
//
// [EN_UPDATE]: https://learn.microsoft.com/en-us/windows/win32/controls/en-update
func (me *EventsEdit) EnUpdate(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.EN_UPDATE, fun)
}

// [EN_VSCROLL] message handler.
//
// [EN_VSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/en-vscroll
func (me *EventsEdit) EnVScroll(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.EN_VSCROLL, fun)
}
