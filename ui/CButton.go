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
	isButton() // prevent public implementation

	// Exposes all the Button notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
	On() *_ButtonEvents

	EmulateClick() // Emulates an user click.
}

//------------------------------------------------------------------------------

type _Button struct {
	_NativeControlBase
	events _ButtonEvents
}

// Creates a new Button. Call ButtonOpts() to define the options to be passed to
// the underlying CreateWindowEx().
func NewButton(parent AnyParent, opts *_ButtonO) Button {
	opts.lateDefaults()

	me := &_Button{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_ConvertDtuOrMultiplyDpi(parent, &opts.position, &opts.size)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			win.ClassNameStr("BUTTON"), win.StrVal(opts.text),
			opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		parent.addResizingChild(me, opts.horz, opts.vert)
		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
	})

	return me
}

// Creates a new Button from a dialog resource.
func NewButtonDlg(parent AnyParent, ctrlId int, horz HORZ, vert VERT) Button {
	me := &_Button{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
		parent.addResizingChild(me, horz, vert)
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

type _ButtonO struct {
	ctrlId int

	text        string
	position    win.POINT
	size        win.SIZE
	horz        HORZ
	vert        VERT
	ctrlStyles  co.BS
	wndStyles   co.WS
	wndExStyles co.WS_EX
}

// Control ID.
//
// Defaults to an auto-generated ID.
func (o *_ButtonO) CtrlId(i int) *_ButtonO { o.ctrlId = i; return o }

// Text to appear in the control, passed to CreateWindowEx().
//
// Defaults to empty string.
func (o *_ButtonO) Text(t string) *_ButtonO { o.text = t; return o }

// Position within parent's client area.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 0x0.
func (o *_ButtonO) Position(p win.POINT) *_ButtonO { _OwPt(&o.position, p); return o }

// Control size.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 88x26.
func (o *_ButtonO) Size(s win.SIZE) *_ButtonO { _OwSz(&o.size, s); return o }

// Horizontal behavior when the parent is resized.
//
// Defaults to HORZ_NONE.
func (o *_ButtonO) Horz(s HORZ) *_ButtonO { o.horz = s; return o }

// Vertical behavior when the parent is resized.
//
// Defaults to VERT_NONE.
func (o *_ButtonO) Vert(s VERT) *_ButtonO { o.vert = s; return o }

// Button control styles, passed to CreateWindowEx().
//
// Defaults to BS_PUSHBUTTON.
func (o *_ButtonO) CtrlStyles(s co.BS) *_ButtonO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *_ButtonO) WndStyles(s co.WS) *_ButtonO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
//
// Defaults to WS_EX_NONE.
func (o *_ButtonO) WndExStyles(s co.WS_EX) *_ButtonO { o.wndExStyles = s; return o }

func (o *_ButtonO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewButton().
func ButtonOpts() *_ButtonO {
	return &_ButtonO{
		size:       win.SIZE{Cx: 88, Cy: 26},
		horz:       HORZ_NONE,
		vert:       VERT_NONE,
		ctrlStyles: co.BS_PUSHBUTTON,
		wndStyles:  co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
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
	me.events = ctrl.Parent().On()
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
