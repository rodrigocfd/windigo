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
	AnyNativeControl
	isDateTimePicker() // prevent public implementation

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

// Creates a new DateTimePicker. Call DateTimePickerOpts() to define the options
// to be passed to the underlying CreateWindowEx().
func NewDateTimePicker(parent AnyParent, opts *_DateTimePickerO) DateTimePicker {
	opts.lateDefaults()

	me := &_DateTimePicker{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_ConvertDtuOrMultiplyDpi(parent, &opts.position, &opts.size)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			win.ClassNameStr("SysDateTimePick32"), nil,
			opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		parent.addResizingChild(me, opts.horz, opts.vert)
		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
	})

	return me
}

// Creates a new DateTimePicker from a dialog resource.
func NewDateTimePickerDlg(
	parent AnyParent, ctrlId int,
	horz HORZ, vert VERT) DateTimePicker {

	me := &_DateTimePicker{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
		parent.addResizingChild(me, horz, vert)
	})

	return me
}

func (me *_DateTimePicker) isDateTimePicker() {}

func (me *_DateTimePicker) On() *_DateTimePickerEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the DateTimePicker is created.")
	}
	return &me.events
}

func (me *_DateTimePicker) SetTime(newTime time.Time) {
	st := win.SYSTEMTIME{}
	st.FromTime(newTime)
	me.Hwnd().SendMessage(co.DTM_SETSYSTEMTIME,
		win.WPARAM(co.GDT_VALID), win.LPARAM(unsafe.Pointer(&st)))
}

func (me *_DateTimePicker) Time() time.Time {
	st := win.SYSTEMTIME{}
	ret := co.GDT(
		me.Hwnd().SendMessage(co.DTM_GETSYSTEMTIME,
			0, win.LPARAM(unsafe.Pointer(&st))),
	)

	if ret != co.GDT_VALID {
		panic("DTM_GETSYSTEMTIME failed.")
	}
	return st.ToTime()
}

//------------------------------------------------------------------------------

type _DateTimePickerO struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	ctrlId int

	position    win.POINT
	size        win.SIZE
	horz        HORZ
	vert        VERT
	ctrlStyles  co.DTS
	wndStyles   co.WS
	wndExStyles co.WS_EX
}

// Control ID.
//
// Defaults to an auto-generated ID.
func (o *_DateTimePickerO) CtrlId(i int) *_DateTimePickerO { o.ctrlId = i; return o }

// Position within parent's client area.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 0x0.
func (o *_DateTimePickerO) Position(p win.POINT) *_DateTimePickerO { _OwPt(&o.position, p); return o }

// Control size.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 230x23.
func (o *_DateTimePickerO) Size(s win.SIZE) *_DateTimePickerO { _OwSz(&o.size, s); return o }

// Horizontal behavior when the parent is resized.
//
// Defaults to HORZ_NONE.
func (o *_DateTimePickerO) Horz(s HORZ) *_DateTimePickerO { o.horz = s; return o }

// Vertical behavior when the parent is resized.
//
// Defaults to VERT_NONE.
func (o *_DateTimePickerO) Vert(s VERT) *_DateTimePickerO { o.vert = s; return o }

// DateTimePicker control styles, passed to CreateWindowEx().
//
// Defaults to DTS_LONGDATEFORMAT.
func (o *_DateTimePickerO) CtrlStyles(s co.DTS) *_DateTimePickerO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *_DateTimePickerO) WndStyles(s co.WS) *_DateTimePickerO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
//
// Defaults to WS_EX_CLIENTEDGE.
func (o *_DateTimePickerO) WndExStyles(s co.WS_EX) *_DateTimePickerO { o.wndExStyles = s; return o }

func (o *_DateTimePickerO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewDateTimePicker().
func DateTimePickerOpts() *_DateTimePickerO {
	return &_DateTimePickerO{
		size:        win.SIZE{Cx: 230, Cy: 23},
		horz:        HORZ_NONE,
		vert:        VERT_NONE,
		ctrlStyles:  co.DTS_LONGDATEFORMAT,
		wndStyles:   co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
		wndExStyles: co.WS_EX_CLIENTEDGE,
	}
}

//------------------------------------------------------------------------------

// DateTimePicker control notifications.
type _DateTimePickerEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_DateTimePickerEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.Parent().On()
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
