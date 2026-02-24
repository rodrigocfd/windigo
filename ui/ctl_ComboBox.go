//go:build windows

package ui

import (
	"fmt"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
)

// Native [combo box] control.
//
// [combo box]: https://learn.microsoft.com/en-us/windows/win32/controls/about-combo-boxes
type ComboBox struct {
	_BaseCtrl
	events ComboBoxEvents
}

// Creates a new [ComboBox] with [win.CreateWindowEx].
//
// Example:
//
//	runtime.LockOSThread()
//
//	wnd := ui.NewMain(
//		ui.OptsMain().
//			Title("Hello world"),
//	)
//	cmb := ui.NewComboBox(
//		wnd,
//		ui.OptsComboBox().
//			Position(ui.Dpi(10, 10)).
//			Texts("Avocado", "Banana", "Pineapple").
//			Select(2),
//	)
//	wnd.RunAsMain()
func NewComboBox(parent Parent, opts *VarOptsComboBox) *ComboBox {
	setUniqueCtrlId(&opts.ctrlId)
	me := &ComboBox{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    ComboBoxEvents{opts.ctrlId, &parent.base().userEvents},
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		sz := win.SIZE{Cx: int32(opts.width)}
		me.createWindow(opts.wndExStyle, "COMBOBOX", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, sz, parent, true)
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		me.AddItem(opts.texts...)
		me.SelectIndex(opts.selected)
	})

	return me
}

// Instantiates a new [ComboBox] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
//
// Example:
//
//	const (
//		ID_MAIN_DLG uint16 = 1000
//		ID_CMB_FOO  uint16 = 1001
//	)
//
//	runtime.LockOSThread()
//
//	wnd := ui.NewTreeViewDlg(
//		ui.OptsMainDlg().
//			DlgId(ID_MAIN_DLG),
//	)
//	cmb := ui.NewComboBoxDlg(wnd, ID_CMB_FOO, ui.LAY_HOLD_HOLD)
//	wnd.RunAsMain()
func NewComboBoxDlg(parent Parent, ctrlId uint16, layout LAY) *ComboBox {
	me := &ComboBox{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    ComboBoxEvents{ctrlId, &parent.base().userEvents},
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
func (me *ComboBox) On() *ComboBoxEvents {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Adds one or more texts using [CB_ADDSTRING].
//
// Panics on error.
//
// [CB_ADDSTRING]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-addstring
func (me *ComboBox) AddItem(texts ...string) {
	var wText wstr.BufEncoder
	for _, text := range texts {
		ret, _ := me.Hwnd().SendMessage(co.CB_ADDSTRING,
			0, win.LPARAM(wText.AllowEmpty(text)))

		if int32(ret) == utl.CB_ERR || int32(ret) == utl.CB_ERRSPACE {
			panic("CB_ADDSTRING failed.")
		}
	}
}

// Retrieves all texts with [CB_GETLBTEXT].
//
// [CB_GETLBTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getlbtext
func (me *ComboBox) AllItems() []string {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	nItems := me.ItemCount()
	items := make([]string, 0, nItems)

	for i := 0; i < nItems; i++ {
		nChars, _ := me.Hwnd().SendMessage(co.CB_GETLBTEXTLEN, win.WPARAM(int32(i)), 0)
		wBuf.AllocAndZero(int(nChars) + 1)

		me.Hwnd().SendMessage(co.CB_GETLBTEXT,
			win.WPARAM(int32(i)), win.LPARAM(wBuf.Ptr()))

		items = append(items, wBuf.String())
	}

	return items
}

// Returns the text currently on display.
//
// If the ComboBox doesn't have the co.CBS_DROPDOWNLIST style, the user can type
// freely, so this text may not be on the list.
func (me *ComboBox) CurrentText() string {
	txt, _ := me.hWnd.GetWindowText()
	return txt
}

// Deletes all texts with [CB_RESETCONTENT].
//
// [CB_RESETCONTENT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-resetcontent
func (me *ComboBox) DeleteAllItems() {
	me.Hwnd().SendMessage(co.CB_RESETCONTENT, 0, 0)
}

// Returns the text at the given zero-based index with [CB_GETLBTEXT].
//
// Panics if the index is not valid.
//
// [CB_GETLBTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getlbtext
func (me *ComboBox) Item(index int) string {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	nChars, _ := me.Hwnd().SendMessage(co.CB_GETLBTEXTLEN, win.WPARAM(int32(index)), 0)
	if int32(nChars) == utl.CB_ERR {
		panic(fmt.Sprintf("Invalid ComboBox index: %d", index))
	}
	wBuf.Alloc(int(nChars) + 1)

	me.Hwnd().SendMessage(co.CB_GETLBTEXT,
		win.WPARAM(int32(index)), win.LPARAM(wBuf.Ptr()))
	return wBuf.String()
}

// Retrieves the number of items with [CB_GETCOUNT].
//
// [CB_GETCOUNT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getcount
func (me *ComboBox) ItemCount() int {
	n, _ := me.Hwnd().SendMessage(co.CB_GETCOUNT, 0, 0)
	return int(n)
}

// Returns the last text with [CB_GETLBTEXT].
//
// Panics if empty.
//
// [CB_GETLBTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getlbtext
func (me *ComboBox) LastItem() string {
	return me.Item(int(me.ItemCount()) - 1)
}

// Selects the text with the given zero-based index with [CB_SETCURSEL].
//
// If index is -1, selection is cleared.
//
// [CB_SETCURSEL]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-setcursel
func (me *ComboBox) SelectIndex(index int) {
	me.Hwnd().SendMessage(co.CB_SETCURSEL, win.WPARAM(int32(index)), 0)
}

// Retrieves the selected zero-based index with [CB_GETCURSEL].
//
// If no item is selected, returns -1.
//
// [CB_GETCURSEL]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getcursel
func (me *ComboBox) SelectedIndex() int {
	n, _ := me.Hwnd().SendMessage(co.CB_GETCURSEL, 0, 0)
	return int(n)
}

// Options for [NewComboBox]; returned by [OptsComboBox].
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

// Options for [NewComboBox].
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
// Defaults to ui.LAY_HOLD_HOLD.
func (o *VarOptsComboBox) Layout(l LAY) *VarOptsComboBox { o.layout = l; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsComboBox) Position(x, y int) *VarOptsComboBox {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control width in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.Dpi(100).
func (o *VarOptsComboBox) Width(w int) *VarOptsComboBox { o.width = w; return o }

// Combo box control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.CBS_DROPDOWNLIST.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/combo-box-styles
func (o *VarOptsComboBox) CtrlStyle(s co.CBS) *VarOptsComboBox { o.ctrlStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_VISIBLE | co.WS_TABSTOP | co.WS_GROUP.
func (o *VarOptsComboBox) WndStyle(s co.WS) *VarOptsComboBox { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT.
func (o *VarOptsComboBox) WndExStyle(s co.WS_EX) *VarOptsComboBox { o.wndExStyle = s; return o }

// Texts to be added to the ComboBox.
//
// Defaults to none.
func (o *VarOptsComboBox) Texts(t ...string) *VarOptsComboBox { o.texts = t; return o }

// Selects the item at the given zero-based index.
//
// Defaults to -1 (none).
func (o *VarOptsComboBox) Select(i int) *VarOptsComboBox { o.selected = i; return o }

// Native [combo box] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [combo box]: https://learn.microsoft.com/en-us/windows/win32/controls/about-combo-boxes
type ComboBoxEvents struct {
	ctrlId       uint16
	parentEvents *WindowEvents
}

// [CBN_CLOSEUP] message handler.
//
// [CBN_CLOSEUP]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-closeup
func (me *ComboBoxEvents) CbnCloseUp(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_CLOSEUP, fun)
}

// [CBN_DBLCLK] message handler.
//
// [CBN_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-dblclk
func (me *ComboBoxEvents) CbnDblClk(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_DBLCLK, fun)
}

// [CBN_DROPDOWN] message handler.
//
// [CBN_DROPDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-dropdown
func (me *ComboBoxEvents) CbnDropDown(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_DROPDOWN, fun)
}

// [CBN_EDITCHANGE] message handler.
//
// [CBN_EDITCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-editchange
func (me *ComboBoxEvents) CbnEditChange(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_EDITCHANGE, fun)
}

// [CBN_EDITUPDATE] message handler.
//
// [CBN_EDITUPDATE]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-editupdate
func (me *ComboBoxEvents) CbnEditUpdate(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_EDITUPDATE, fun)
}

// [CBN_ERRSPACE] message handler.
//
// [CBN_ERRSPACE]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-errspace
func (me *ComboBoxEvents) CbnErrSpace(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_ERRSPACE, fun)
}

// [CBN_KILLFOCUS] message handler.
//
// [CBN_KILLFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-killfocus
func (me *ComboBoxEvents) CbnKillFocus(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_KILLFOCUS, fun)
}

// [CBN_SELCHANGE] message handler.
//
// [CBN_SELCHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-selchange
func (me *ComboBoxEvents) CbnSelChange(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_SELCHANGE, fun)
}

// [CBN_SELENDCANCEL] message handler.
//
// [CBN_SELENDCANCEL]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-selendcancel
func (me *ComboBoxEvents) CbnSelEndCancel(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_SELENDCANCEL, fun)
}

// [CBN_SELENDOK] message handler.
//
// [CBN_SELENDOK]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-selendok
func (me *ComboBoxEvents) CbnSelEndOk(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_SELENDOK, fun)
}

// [CBN_SETFOCUS] message handler.
//
// [CBN_SETFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/cbn-setfocus
func (me *ComboBoxEvents) CbnSetFocus(fun func()) {
	me.parentEvents.WmCommand(me.ctrlId, co.CBN_SETFOCUS, fun)
}
