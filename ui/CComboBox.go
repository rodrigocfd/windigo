package ui

import (
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native combo box control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/about-combo-boxes
type ComboBox interface {
	AnyNativeControl
	isComboBox() // prevent public implementation

	// Exposes all the ComboBox notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-combobox-control-reference-notifications
	On() *_ComboBoxEvents

	Items() *_ComboBoxItems // Item methods.
}

//------------------------------------------------------------------------------

type _ComboBox struct {
	_NativeControlBase
	events _ComboBoxEvents
	items  _ComboBoxItems
}

// Creates a new ComboBox. Call ComboBox() to define the options to be passed to
// the underlying CreateWindowEx().
func NewComboBox(parent AnyParent, opts *_ComboBoxO) ComboBox {
	opts.lateDefaults()

	me := &_ComboBox{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)
	me.items.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		size := win.SIZE{Cx: int32(opts.width), Cy: 0}
		_MultiplyDpi(&opts.position, &size)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			"COMBOBOX", "", opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, size, win.HMENU(opts.ctrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)

		if opts.texts != nil {
			me.Items().Add(opts.texts...)
		}
	})

	return me
}

// Creates a new ComboBox from a dialog resource.
func NewComboBoxDlg(parent AnyParent, ctrlId int) ComboBox {
	me := &_ComboBox{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)
	me.items.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return me
}

func (me *_ComboBox) isComboBox() {}

func (me *_ComboBox) On() *_ComboBoxEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the ComboBox is created.")
	}
	return &me.events
}

func (me *_ComboBox) Items() *_ComboBoxItems {
	return &me.items
}

//------------------------------------------------------------------------------

type _ComboBoxO struct {
	ctrlId int

	position    win.POINT
	width       int
	ctrlStyles  co.CBS
	wndStyles   co.WS
	wndExStyles co.WS_EX

	texts []string
}

// Control ID.
// Defaults to an auto-generated ID.
func (o *_ComboBoxO) CtrlId(i int) *_ComboBoxO { o.ctrlId = i; return o }

// Position within parent's client area in pixels.
// Defaults to 0x0. Will be adjusted to the current system DPI.
func (o *_ComboBoxO) Position(p win.POINT) *_ComboBoxO { _OwPt(&o.position, p); return o }

// Control width in pixels.
// Defaults to 100. Will be adjusted to the current system DPI.
func (o *_ComboBoxO) Width(w int) *_ComboBoxO { o.width = w; return o }

// ComboBox control styles, passed to CreateWindowEx().
// Defaults to CBS_DROPDOWNLIST.
func (o *_ComboBoxO) CtrlStyles(s co.CBS) *_ComboBoxO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *_ComboBoxO) WndStyles(s co.WS) *_ComboBoxO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
// Defaults to WS_EX_NONE.
func (o *_ComboBoxO) WndExStyles(s co.WS_EX) *_ComboBoxO { o.wndExStyles = s; return o }

// Texts to be added to the ComboBox.
// Defaults to none.
func (o *_ComboBoxO) Texts(t ...string) *_ComboBoxO { o.texts = t; return o }

func (o *_ComboBoxO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewComboBox().
func ComboBoxOpts() *_ComboBoxO {
	return &_ComboBoxO{
		width:      100,
		ctrlStyles: co.CBS_DROPDOWNLIST,
		wndStyles:  co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
	}
}

//------------------------------------------------------------------------------

// ComboBox control notifications.
type _ComboBoxEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_ComboBoxEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.Parent().On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/cbn-closeup
func (me *_ComboBoxEvents) CbnCloseUp(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.CBN_CLOSEUP, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/cbn-dblclk
func (me *_ComboBoxEvents) CbnDblClk(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.CBN_DBLCLK, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/cbn-dropdown
func (me *_ComboBoxEvents) CbnDropDown(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.CBN_DROPDOWN, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/cbn-editchange
func (me *_ComboBoxEvents) CbnEditChange(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.CBN_EDITCHANGE, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/cbn-editupdate
func (me *_ComboBoxEvents) CbnEditUpdate(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.CBN_EDITUPDATE, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/cbn-errspace
func (me *_ComboBoxEvents) CbnErrSpace(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.CBN_ERRSPACE, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/cbn-killfocus
func (me *_ComboBoxEvents) CbnKillFocus(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.CBN_KILLFOCUS, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/cbn-selchange
func (me *_ComboBoxEvents) CbnSelChange(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.CBN_SELCHANGE, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/cbn-selendcancel
func (me *_ComboBoxEvents) CbnSelEndCancel(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.CBN_SELENDCANCEL, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/cbn-selendok
func (me *_ComboBoxEvents) CbnSelEndOk(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.CBN_SELENDOK, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/cbn-setfocus
func (me *_ComboBoxEvents) CbnSetFocus(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.CBN_SETFOCUS, func(_ wm.Command) {
		userFunc()
	})
}
