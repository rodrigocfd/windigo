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

// Creates a new ComboBox specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewComboBox(parent AnyParent, opts ComboBoxOpts) ComboBox {
	opts.fillBlankValuesWithDefault()

	me := &_ComboBox{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.events.new(&me._NativeControlBase)
	me.items.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		size := win.SIZE{Cx: int32(opts.Width), Cy: 0}
		_MultiplyDpi(&opts.Position, &size)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"COMBOBOX", "", opts.Styles|co.WS(opts.ComboBoxStyles),
			opts.Position, size, win.HMENU(opts.CtrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)

		if opts.Texts != nil {
			me.Items().Add(opts.Texts...)
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

// Options for NewComboBox().
type ComboBoxOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// Control width in pixels.
	// Defaults to 100. Will be adjusted to the current system DPI.
	Width int
	// ComboBox control styles, passed to CreateWindowEx().
	// Defaults to CBS_DROPDOWNLIST.
	ComboBoxStyles co.CBS
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStyles co.WS_EX

	// Texts to be added to the ComboBox.
	// Defaults to none.
	Texts []string
}

func (opts *ComboBoxOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.Width == 0 {
		opts.Width = 100
	}

	if opts.ComboBoxStyles == 0 {
		opts.ComboBoxStyles = co.CBS_DROPDOWNLIST
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_NONE
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
	me.events = ctrl.parent.On()
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
