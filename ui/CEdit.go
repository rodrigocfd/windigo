package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native edit control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/about-edit-controls
type Edit interface {
	AnyNativeControl
	isEdit() // prevent public implementation

	// Exposes all the Edit notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-edit-control-reference-notifications
	On() *_EditEvents

	LimitText(maxChars int)            // Limits the length of the text.
	ReplaceSelection(text string)      // Replaces the current text selection with the given text.
	SelectedRange() (int, int)         // Retrieves the index of first and last selected chars.
	SelectRange(idxFirst, idxLast int) // Sets the currently selected chars.
	SetText(text string)               // Sets the text.
	Text() string                      // Retrieves the text.
}

//------------------------------------------------------------------------------

type _Edit struct {
	_NativeControlBase
	events _EditEvents
}

// Creates a new Edit. Call EditOpts() to define the options to be passed to
// the underlying CreateWindowEx().
func NewEdit(parent AnyParent, opts *_EditO) Edit {
	opts.lateDefaults()

	me := &_Edit{}
	me._NativeControlBase.new(parent, opts.ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_ConvertDtuOrMultiplyDpi(parent, &opts.position, &opts.size)

		me._NativeControlBase.createWindow(opts.wndExStyles,
			win.ClassNameStr("EDIT"), win.StrVal(opts.text),
			opts.wndStyles|co.WS(opts.ctrlStyles),
			opts.position, opts.size, win.HMENU(opts.ctrlId))

		parent.addResizingChild(me, opts.horz, opts.vert)
		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
	})

	return me
}

// Creates a new Edit from a dialog resource.
func NewEditDlg(parent AnyParent, ctrlId int, horz HORZ, vert VERT) Edit {
	me := &_Edit{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
		parent.addResizingChild(me, horz, vert)
	})

	return me
}

func (me *_Edit) isEdit() {}

func (me *_Edit) On() *_EditEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the Edit is created.")
	}
	return &me.events
}

func (me *_Edit) LimitText(maxChars int) {
	me.Hwnd().SendMessage(co.EM_LIMITTEXT, win.WPARAM(maxChars), 0)
}

func (me *_Edit) ReplaceSelection(replacementText string) {
	me.Hwnd().SendMessage(co.EM_REPLACESEL,
		1, win.LPARAM(unsafe.Pointer(win.Str.ToNativePtr(replacementText))))
}

func (me *_Edit) SelectedRange() (idxFirst, idxLast int) {
	var idxFirstU, idxLastU uint32
	me.Hwnd().SendMessage(co.EM_GETSEL,
		win.WPARAM(unsafe.Pointer(&idxFirstU)),
		win.LPARAM(unsafe.Pointer(&idxLastU)))
	idxFirst, idxLast = int(idxFirstU), int(idxLastU)
	return
}

func (me *_Edit) SelectRange(idxFirst, idxLast int) {
	me.Hwnd().SendMessage(co.EM_SETSEL,
		win.WPARAM(idxFirst), win.LPARAM(idxLast))
}

func (me *_Edit) SetText(text string) {
	me.Hwnd().SetWindowText(text)
}

func (me *_Edit) Text() string {
	return me.Hwnd().GetWindowText()
}

//------------------------------------------------------------------------------

type _EditO struct {
	ctrlId int

	text        string
	position    win.POINT
	size        win.SIZE
	horz        HORZ
	vert        VERT
	ctrlStyles  co.ES
	wndStyles   co.WS
	wndExStyles co.WS_EX
}

// Control ID.
//
// Defaults to an auto-generated ID.
func (o *_EditO) CtrlId(i int) *_EditO { o.ctrlId = i; return o }

// Text to appear in the control, passed to CreateWindowEx().
//
// Defaults to empty string.
func (o *_EditO) Text(t string) *_EditO { o.text = t; return o }

// Position within parent's client area.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 0x0.
func (o *_EditO) Position(p win.POINT) *_EditO { _OwPt(&o.position, p); return o }

// Control size in pixels.
//
// If parent is a dialog box, coordinates are in Dialog Template Units;
// otherwise, they are in pixels and they will be adjusted to the current system
// DPI.
//
// Defaults to 100x23.
func (o *_EditO) Size(s win.SIZE) *_EditO { _OwSz(&o.size, s); return o }

// Horizontal behavior when the parent is resized.
//
// Defaults to HORZ_NONE.
func (o *_EditO) Horz(s HORZ) *_EditO { o.horz = s; return o }

// Vertical behavior when the parent is resized.
//
// Defaults to VERT_NONE.
func (o *_EditO) Vert(s VERT) *_EditO { o.vert = s; return o }

// Edit control styles, passed to CreateWindowEx().
//
// Defaults to ES_AUTOHSCROLL | ES_NOHIDESEL.
func (o *_EditO) CtrlStyles(s co.ES) *_EditO { o.ctrlStyles = s; return o }

// Window styles, passed to CreateWindowEx().
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *_EditO) WndStyles(s co.WS) *_EditO { o.wndStyles = s; return o }

// Extended window styles, passed to CreateWindowEx().
//
// Defaults to WS_EX_CLIENTEDGE.
func (o *_EditO) WndExStyles(s co.WS_EX) *_EditO { o.wndExStyles = s; return o }

func (o *_EditO) lateDefaults() {
	if o.ctrlId == 0 {
		o.ctrlId = _NextCtrlId()
	}
}

// Options for NewEdit().
func EditOpts() *_EditO {
	return &_EditO{
		size:        win.SIZE{Cx: 100, Cy: 23},
		horz:        HORZ_NONE,
		vert:        VERT_NONE,
		ctrlStyles:  co.ES_AUTOHSCROLL | co.ES_NOHIDESEL,
		wndStyles:   co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
		wndExStyles: co.WS_EX_CLIENTEDGE,
	}
}

//------------------------------------------------------------------------------

// Edit control notifications.
type _EditEvents struct {
	ctrlId int
	events *_EventsWmNfy
}

func (me *_EditEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.Parent().On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-align-ltr-ec
func (me *_EditEvents) EnAlignLtrEc(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_ALIGN_LTR_EC, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-align-rtl-ec
func (me *_EditEvents) EnAlignRtlEc(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_ALIGN_RTL_EC, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-change
func (me *_EditEvents) EnChange(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_CHANGE, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-errspace
func (me *_EditEvents) EnErrSpace(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_ERRSPACE, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-hscroll
func (me *_EditEvents) EnHScroll(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_HSCROLL, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-killfocus
func (me *_EditEvents) EnKillFocus(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_KILLFOCUS, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-maxtext
func (me *_EditEvents) EnMaxText(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_MAXTEXT, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-setfocus
func (me *_EditEvents) EnSetFocus(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_SETFOCUS, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-update
func (me *_EditEvents) EnUpdate(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_UPDATE, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-vscroll
func (me *_EditEvents) EnVScroll(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_VSCROLL, func(_ wm.Command) {
		userFunc()
	})
}
