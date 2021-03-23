package ui

import (
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native date and time picker control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/date-and-time-picker-controls
type DateTimePicker interface {
	AnyControl

	// Exposes all the DateTimePicker notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-date-and-time-picker-control-reference-notifications
	On() *_DateTimePickerEvents

	SetTime(newTime time.Time) // Sets the current time.
	Time() time.Time           // Retrieves current time.
}

//------------------------------------------------------------------------------

type _DateTimePicker struct {
	_NativeControlBase
	events _DateTimePickerEvents
}

// Creates a new DateTimePicker specifying all options, which will be passed to
// the underlying CreateWindowEx().
func NewDateTimePickerOpts(parent AnyParent, opts DateTimePickerOpts) DateTimePicker {
	opts.fillBlankValuesWithDefault()

	me := _DateTimePicker{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_ParentCreateWm(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.Position, &opts.Size)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"SysDateTimePick32", "", opts.Styles|co.WS(opts.DateTimePickerStyles),
			opts.Position, opts.Size, win.HMENU(opts.CtrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
	})

	return &me
}

// Creates a new DateTimePicker from a dialog resource.
func NewDateTimePickerDlg(parent AnyParent, ctrlId int) DateTimePicker {
	me := _DateTimePicker{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return &me
}

func (me *_DateTimePicker) On() *_DateTimePickerEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the DateTimePicker is created.")
	}
	return &me.events
}

// Sets the current time.
func (me *_DateTimePicker) SetTime(newTime time.Time) {
	st := win.SYSTEMTIME{}
	win.Time.ToSystemtime(newTime, &st)
	me.Hwnd().SendMessage(co.DTM_SETSYSTEMTIME,
		win.WPARAM(co.GDT_VALID), win.LPARAM(unsafe.Pointer(&st)))
}

func (me *_DateTimePicker) Time() time.Time {
	st := win.SYSTEMTIME{}
	ret := co.GDT(me.Hwnd().SendMessage(co.DTM_GETSYSTEMTIME,
		0, win.LPARAM(unsafe.Pointer(&st))))

	if ret != co.GDT_VALID {
		panic("DTM_GETSYSTEMTIME failed.")
	}
	return win.Time.FromSystemtime(&st)
}

//------------------------------------------------------------------------------

// Options for NewDateTimePickerOpts().
type DateTimePickerOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// Control size in pixels.
	// Defaults to 230x21. Will be adjusted to the current system DPI.
	Size win.SIZE
	// DateTimePicker control styles, passed to CreateWindowEx().
	// Defaults to DTS_LONGDATEFORMAT.
	DateTimePickerStyles co.DTS
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_CLIENTEDGE.
	ExStyles co.WS_EX
}

func (opts *DateTimePickerOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.Size.Cx == 0 {
		opts.Size.Cx = 230
	}
	if opts.Size.Cy == 0 {
		opts.Size.Cy = 21
	}

	if opts.DateTimePickerStyles == 0 {
		opts.DateTimePickerStyles = co.DTS_LONGDATEFORMAT
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_CLIENTEDGE
	}
}

//------------------------------------------------------------------------------

// DateTimePicker control notifications.
type _DateTimePickerEvents struct {
	ctrlId int
	events *_EventsNfy
}

func (me *_DateTimePickerEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.parent.On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/dtn-closeup
func (me *_DateTimePickerEvents) DtnCloseUp(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.DTN_CLOSEUP, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/dtn-datetimechange
func (me *_DateTimePickerEvents) DtnDateTimeChange(userFunc func(p *win.NMDATETIMECHANGE)) {
	me.events.addNfyZero(me.ctrlId, co.DTN_DATETIMECHANGE, func(p unsafe.Pointer) {
		userFunc((*win.NMDATETIMECHANGE)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/dtn-dropdown
func (me *_DateTimePickerEvents) DtnDropDown(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.DTN_DROPDOWN, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/dtn-format
func (me *_DateTimePickerEvents) DtnFormat(userFunc func(p *win.NMDATETIMEFORMAT)) {
	me.events.addNfyZero(me.ctrlId, co.DTN_FORMAT, func(p unsafe.Pointer) {
		userFunc((*win.NMDATETIMEFORMAT)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/dtn-formatquery
func (me *_DateTimePickerEvents) DtnFormatQuery(userFunc func(p *win.NMDATETIMEFORMATQUERY)) {
	me.events.addNfyZero(me.ctrlId, co.DTN_FORMATQUERY, func(p unsafe.Pointer) {
		userFunc((*win.NMDATETIMEFORMATQUERY)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/dtn-userstring
func (me *_DateTimePickerEvents) DtnUserString(userFunc func(p *win.NMDATETIMESTRING)) {
	me.events.addNfyZero(me.ctrlId, co.DTN_USERSTRING, func(p unsafe.Pointer) {
		userFunc((*win.NMDATETIMESTRING)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/dtn-wmkeydown
func (me *_DateTimePickerEvents) DtnWmKeyDown(userFunc func(p *win.NMDATETIMEWMKEYDOWN)) {
	me.events.addNfyZero(me.ctrlId, co.DTN_WMKEYDOWN, func(p unsafe.Pointer) {
		userFunc((*win.NMDATETIMEWMKEYDOWN)(p))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-killfocus-date-time
func (me *_DateTimePickerEvents) NmKillFocus(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_KILLFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/nm-setfocus-date-time-
func (me *_DateTimePickerEvents) NmSetFocus(userFunc func()) {
	me.events.addNfyZero(me.ctrlId, co.NM_SETFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}
