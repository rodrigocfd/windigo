package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native button control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#push-buttons
type Button interface {
	AnyNativeControl

	// Exposes all the Button notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
	On() *_ButtonEvents

	isButton() // disambiguate

	EmulateClick() // Emulates an user click.
}

//------------------------------------------------------------------------------

type _Button struct {
	_NativeControlBase
	events _ButtonEvents
}

// Creates a new Button specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewButton(parent AnyParent, opts ButtonOpts) Button {
	opts.fillBlankValuesWithDefault()

	me := &_Button{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.Position, &opts.Size)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"BUTTON", opts.Text, opts.Styles|co.WS(opts.ButtonStyles),
			opts.Position, opts.Size, win.HMENU(opts.CtrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
	})

	return me
}

// Creates a new Button from a dialog resource.
func NewButtonDlg(parent AnyParent, ctrlId int) Button {
	me := &_Button{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return me
}

func (me *_Button) isButton() {}

func (me *_Button) On() *_ButtonEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the Button is created.")
	}
	return &me.events
}

func (me *_Button) EmulateClick() {
	me.Hwnd().SendMessage(co.BM_CLICK, 0, 0)
}

//------------------------------------------------------------------------------

// Options for NewButton().
type ButtonOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

	// Text to appear in the control, passed to CreateWindowEx().
	// Defaults to empty string.
	Text string
	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// Control size in pixels.
	// Defaults to 80x23. Will be adjusted to the current system DPI.
	Size win.SIZE
	// Button control styles, passed to CreateWindowEx().
	// Defaults to BS_PUSHBUTTON.
	ButtonStyles co.BS
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_NONE.
	ExStyles co.WS_EX
}

func (opts *ButtonOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.Size.Cx == 0 {
		opts.Size.Cx = 80
	}
	if opts.Size.Cy == 0 {
		opts.Size.Cy = 23
	}

	if opts.ButtonStyles == 0 {
		opts.ButtonStyles = co.BS_PUSHBUTTON
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_NONE
	}
}

//------------------------------------------------------------------------------

// Button control notifications.
type _ButtonEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_ButtonEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.parent.On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bcn-dropdown
func (me *_ButtonEvents) BcnDropDown(userFunc func(p *win.NMBCDROPDOWN)) {
	me.events.addNfyZero(me.ctrlId, co.BCN_DROPDOWN, func(p unsafe.Pointer) {
		userFunc((*win.NMBCDROPDOWN)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bcn-hotitemchange
func (me *_ButtonEvents) BcnHotItemChange(userFunc func(p *win.NMBCHOTITEM)) {
	me.events.addNfyZero(me.ctrlId, co.BCN_HOTITEMCHANGE, func(p unsafe.Pointer) {
		userFunc((*win.NMBCHOTITEM)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bn-clicked
func (me *_ButtonEvents) BnClicked(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.BN_CLICKED, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bn-dblclk
func (me *_ButtonEvents) BnDblClk(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.BN_DBLCLK, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bn-killfocus
func (me *_ButtonEvents) BnKillFocus(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.BN_KILLFOCUS, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bn-setfocus
func (me *_ButtonEvents) BnSetFocus(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.BN_SETFOCUS, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-customdraw-button
func (me *_ButtonEvents) NmCustomDraw(userFunc func(p *win.NMCUSTOMDRAW) co.CDRF) {
	me.events.addNfyRet(me.ctrlId, co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMCUSTOMDRAW)(p)))
	})
}
