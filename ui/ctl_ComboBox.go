//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [combo box] control.
//
// [combo box]: https://learn.microsoft.com/en-us/windows/win32/controls/about-combo-boxes
type ComboBox struct {
	_BaseCtrl
	events EventsComboBox
	Items  CollectionComboBoxItems // Methods to interact with the items collection.
}

// Creates a new ComboBox with [CreateWindowEx].
//
// # Example
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	cmb := ui.NewComboBox(
//		wndOwner,
//		ui.OptsComboBox().
//			Position(ui.Dpi(20, 92)).
//			Texts("Avocado", "Banana", "Pineapple").
//			Selected(2),
//	)
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func NewComboBox(parent Parent, opts *VarOptsComboBox) *ComboBox {
	setUniqueCtrlId(&opts.ctrlId)
	me := &ComboBox{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsComboBox{opts.ctrlId, &parent.base().userEvents},
	}
	me.Items.owner = me

	parent.base().beforeUserEvents.WmCreate(func(_ WmCreate) int {
		sz := win.SIZE{Cx: int32(opts.width)}
		me.createWindow(opts.wndExStyle, "COMBOBOX", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, sz, parent, true)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		me.Items.Add(opts.texts...)
		me.Items.Select(opts.selected)
		return 0 // ignored
	})

	return me
}

// Instantiates a new ComboBox to be loaded from a dialog resource with
// [GetDlgItem].
//
// # Example
//
//	const ID_CMB uint16 = 0x100
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	cmb := ui.NewComboBoxDlg(
//		wndOwner, ID_CMB, ui.LAY_NONE_NONE)
//
// [GetDlgItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgitem
func NewComboBoxDlg(parent Parent, ctrlId uint16, layout LAY) *ComboBox {
	me := &ComboBox{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsComboBox{ctrlId, &parent.base().userEvents},
	}
	me.Items.owner = me

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
func (me *ComboBox) On() *EventsComboBox {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Returns the text currently on display.
func (me *ComboBox) Text() string {
	txt, _ := me.hWnd.GetWindowText()
	return txt
}

// Options for ui.NewComboBox(); returned by ui.OptsComboBox().
type VarOptsComboBox struct {
	ctrlId     uint16
	layout     LAY
	position   win.POINT
	width      int
	ctrlStyle  co.CBS
	wndStyle   co.WS
	wndExStyle co.WS_EX

	texts    []string
	selected int
}

// Options for ui.NewComboBox().
func OptsComboBox() *VarOptsComboBox {
	return &VarOptsComboBox{
		width:     DpiX(100),
		ctrlStyle: co.CBS_DROPDOWNLIST,
		wndStyle:  co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP,
		selected:  -1,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsComboBox) CtrlId(id uint16) *VarOptsComboBox { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsComboBox) Layout(l LAY) *VarOptsComboBox { o.layout = l; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsComboBox) Position(x, y int) *VarOptsComboBox {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control width in pixels, passed to [CreateWindowEx].
//
// Defaults to ui.Dpi(100).
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsComboBox) Width(w int) *VarOptsComboBox { o.width = w; return o }

// Combo box control [style], passed to [CreateWindowEx].
//
// Defaults to co.CBS_DROPDOWNLIST.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/combo-box-styles
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsComboBox) CtrlStyle(s co.CBS) *VarOptsComboBox { o.ctrlStyle = s; return o }

// Window style, passed to [CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsComboBox) WndStyle(s co.WS) *VarOptsComboBox { o.wndStyle = s; return o }

// Window extended style, passed to [CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (o *VarOptsComboBox) WndExStyle(s co.WS_EX) *VarOptsComboBox { o.wndExStyle = s; return o }

// Texts to be added to the ComboBox.
//
// Defaults to none.
func (o *VarOptsComboBox) Texts(t ...string) *VarOptsComboBox { o.texts = t; return o }

// Zero-based index of the item initially selected.
//
// Defaults to -1 (none).
func (o *VarOptsComboBox) Selected(i int) *VarOptsComboBox { o.selected = i; return o }

// Native [combo box] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [combo box]: https://learn.microsoft.com/en-us/windows/win32/controls/about-combo-boxes
type EventsComboBox struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [CBN_CLOSEUP] message handler.
//
// [CBN_CLOSEUP]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-closeup
func (me *EventsComboBox) CbnCloseUp(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_CLOSEUP, fun)
}

// [CBN_DBLCLK] message handler.
//
// [CBN_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-dblclk
func (me *EventsComboBox) CbnDblClk(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_DBLCLK, fun)
}

// [CBN_DROPDOWN] message handler.
//
// [CBN_DROPDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-dropdown
func (me *EventsComboBox) CbnDropDown(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_DROPDOWN, fun)
}

// [CBN_EDITCHANGE] message handler.
//
// [CBN_EDITCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-editchange
func (me *EventsComboBox) CbnEditChange(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_EDITCHANGE, fun)
}

// [CBN_EDITUPDATE] message handler.
//
// [CBN_EDITUPDATE]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-editupdate
func (me *EventsComboBox) CbnEditUpdate(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_EDITUPDATE, fun)
}

// [CBN_ERRSPACE] message handler.
//
// [CBN_ERRSPACE]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-errspace
func (me *EventsComboBox) CbnErrSpace(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_ERRSPACE, fun)
}

// [CBN_KILLFOCUS] message handler.
//
// [CBN_KILLFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-killfocus
func (me *EventsComboBox) CbnKillFocus(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_KILLFOCUS, fun)
}

// [CBN_SELCHANGE] message handler.
//
// [CBN_SELCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-selchange
func (me *EventsComboBox) CbnSelChange(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_SELCHANGE, fun)
}

// [CBN_SELENDCANCEL] message handler.
//
// [CBN_SELENDCANCEL]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-selendcancel
func (me *EventsComboBox) CbnSelEndCancel(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_SELENDCANCEL, fun)
}

// [CBN_SELENDOK] message handler.
//
// [CBN_SELENDOK]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-selendok
func (me *EventsComboBox) CbnSelEndOk(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_SELENDOK, fun)
}

// [CBN_SETFOCUS] message handler.
//
// [CBN_SETFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-setfocus
func (me *EventsComboBox) CbnSetFocus(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_SETFOCUS, fun)
}
