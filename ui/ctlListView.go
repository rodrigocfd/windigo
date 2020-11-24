/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Native ListView control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/list-view-controls-overview
type ListView struct {
	*_NativeControlBase
	events      *_EventsListView
	columns     *_ListViewColumnCollection
	items       *_ListViewItemCollection
	contextMenu *Menu // we don't own this menu; if set, will be shown with right-click
}

// Constructor. Optionally receives a control ID.
func NewListView(parent Parent, ctrlId ...int) *ListView {
	base := _NewNativeControlBase(parent, ctrlId...)
	me := &ListView{
		_NativeControlBase: base,
		events:             _NewEventsListView(base),
	}
	me.columns = _NewListViewColumnCollection(me)
	me.items = _NewListViewItemCollection(me)
	me.installSubclass()
	return me
}

// Calls CreateWindowEx(). With this method, you must also specify WS and WS_EX
// window styles.
//
// For safety, LVS_SHAREIMAGELISTS will be added automatically.
//
// Position and size will be adjusted to the current system DPI.
func (me *ListView) CreateWs(
	pos Pos, size Size,
	lvStyles co.LVS, lvExStyles co.LVS_EX,
	styles co.WS, exStyles co.WS_EX) *ListView {

	_global.MultiplyDpi(&pos, &size)
	me._NativeControlBase.create("SysListView32", "", pos, size,
		co.WS(lvStyles|co.LVS_SHAREIMAGELISTS)|styles, exStyles)

	if lvExStyles != co.LVS_EX_NONE {
		me.SetExtendedStyle(true, lvExStyles)
	}
	return me
}

// Calls CreateWindowEx() with WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE,
// and WS_EX_CLIENTEDGE.
//
// For safety, LVS_SHAREIMAGELISTS will be added automatically.
//
// A typical report ListView has LVS_REPORT | LVS_NOSORTHEADER | LVS_SHOWSELALWAYS.
//
// Position and size will be adjusted to the current system DPI.
func (me *ListView) Create(
	pos Pos, size Size, lvStyles co.LVS, lvExStyles co.LVS_EX) *ListView {

	return me.CreateWs(pos, size,
		lvStyles|co.LVS_SHAREIMAGELISTS, lvExStyles,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_CLIENTEDGE)
}

// Exposes all ListView notifications.
func (me *ListView) On() *_EventsListView {
	if me.hwnd != 0 {
		panic("Cannot add notifications after the ListView was created.")
	}
	return me.events
}

// Access to the columns.
func (me *ListView) Columns() *_ListViewColumnCollection {
	return me.columns
}

// Retrieves extended styles.
func (me *ListView) ExtendedStyle() co.LVS_EX {
	return co.LVS_EX(
		me.Hwnd().SendMessage(co.WM(co.LVM_GETEXTENDEDLISTVIEWSTYLE), 0, 0),
	)
}

// Access to the items.
func (me *ListView) Items() *_ListViewItemCollection {
	return me.items
}

// Sends LVM_ISGROUPVIEWENABLED.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-isgroupviewenabled
func (me *ListView) IsGroupViewEnabled() bool {
	return me.Hwnd().SendMessage(co.WM(co.LVM_ISGROUPVIEWENABLED), 0, 0) >= 0
}

// Scrolls the contents with LVM_SCROLL.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-scroll
func (me *ListView) Scroll(pxHorz, pxVert int) *ListView {
	ret := me.Hwnd().SendMessage(co.WM(co.LVM_SCROLL),
		win.WPARAM(pxHorz), win.LPARAM(pxVert))
	if ret == 0 {
		panic("LVM_SCROLL failed.")
	}
	return me
}

// Defines a menu to be shown as the context menu for the list view.
//
// The menu is shared, the ListView doesn't own it.
func (me *ListView) SetContextMenu(popupMenu *Menu) *ListView {
	me.contextMenu = popupMenu
	return me
}

// Sets or unsets extended styles.
func (me *ListView) SetExtendedStyle(isSet bool, exStyle co.LVS_EX) *ListView {
	mask := exStyle
	if !isSet {
		mask = 0
	}
	me.Hwnd().SendMessage(co.WM(co.LVM_SETEXTENDEDLISTVIEWSTYLE),
		win.WPARAM(mask), win.LPARAM(exStyle))
	return me
}

// Sets the image list associated with this list view.
// The image list is shared and must remain valid.
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-setimagelist
func (me *ListView) SetImageList(
	imgListType co.LVSIL, imgList *ImageList) *ListView {

	me.Hwnd().SendMessage(co.WM(co.LVM_SETIMAGELIST),
		win.WPARAM(imgListType), win.LPARAM(imgList.Himagelist()))
	return me
}

// Sends WM_SETREDRAW to enable or disable UI updates.
//
// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-setredraw
func (me *ListView) SetRedraw(allowRedraw bool) *ListView {
	me.Hwnd().SendMessage(co.WM_SETREDRAW,
		win.WPARAM(_global.BoolToUint32(allowRedraw)), 0)
	return me
}

// Sets the current view with LVM_SETVIEW.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-setview
func (me *ListView) SetView(view co.LV_VIEW) *ListView {
	if int(me.Hwnd().SendMessage(co.WM(co.LVM_SETVIEW), 0, 0)) == -1 {
		panic("LVM_SETVIEW failed.")
	}
	return me
}

// Returns the width of a string using current ListView current, with
// LVM_GETSTRINGWIDTH.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getstringwidth
func (me *ListView) StringWidth(text string) int {
	ret := int(
		me.Hwnd().SendMessage(co.WM(co.LVM_GETSTRINGWIDTH),
			0, win.LPARAM(unsafe.Pointer(win.Str.ToUint16Ptr(text)))),
	)
	if ret == 0 {
		panic("LVM_GETSTRINGWIDTH failed.")
	}
	return ret
}

// Retrieves current view with LVM_GETVIEW.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getview
func (me *ListView) View() co.LV_VIEW {
	return co.LV_VIEW(me.Hwnd().SendMessage(co.WM(co.LVM_GETVIEW), 0, 0))
}

// Adds default subclass message handlers.
func (me *ListView) installSubclass() {
	me.OnSubclass().WmRButtonDown(func(p WmMouse) {
		// WM_RBUTTONUP doesn't work, only NM_RCLICK on parent.
		// https://stackoverflow.com/a/30206896
		me.showContextMenu(true, p.HasCtrl(), p.HasShift())
	})

	me.OnSubclass().WmGetDlgCode(func(p WmGetDlgCode) co.DLGC {
		if !p.IsQuery() && p.VirtualKeyCode() == 'A' && p.HasCtrl() { // Ctrl+A to select all items
			me.Items().SelectAll(true)
			return co.DLGC_WANTCHARS

		} else if !p.IsQuery() && p.VirtualKeyCode() == co.VK_RETURN { // send Enter key to parent
			code := co.LVN_KEYDOWN
			nmlvk := win.NMLVKEYDOWN{
				Hdr: win.NMHDR{
					HWndFrom: me.Hwnd(),
					Code:     uint32(code),
					IdFrom:   uintptr(me.CtrlId()),
				},
				WVKey: co.VK_RETURN,
			}
			me.Hwnd().GetAncestor(co.GA_PARENT).
				SendMessage(co.WM_NOTIFY,
					win.WPARAM(me.Hwnd()), win.LPARAM(unsafe.Pointer(&nmlvk)))
			return co.DLGC_WANTALLKEYS

		} else if !p.IsQuery() && p.VirtualKeyCode() == co.VK_APPS { // context menu key
			me.showContextMenu(false, p.HasCtrl(), p.HasShift())
		}

		return co.DLGC(
			me.Hwnd().DefSubclassProc(co.WM_GETDLGCODE,
				p.Raw().WParam, p.Raw().LParam),
		)
	})
}

// Shows the popup menu anchored at cursor pos.
//
// This function will block until the menu disappears.
func (me *ListView) showContextMenu(followCursor, hasCtrl, hasShift bool) {
	if me.contextMenu.Hmenu() == 0 {
		return
	}

	var menuPos *win.POINT // menu anchor coords, relative to list view

	if followCursor { // usually when fired by a right-click
		menuPos = win.GetCursorPos()          // relative to screen
		me.Hwnd().ScreenToClientPt(menuPos)   // now relative to list view
		lvhti := me.Items().HitTest(*menuPos) // to find item below cursor, if any

		if lvhti.IItem != -1 { // an item was right-clicked
			if !hasCtrl && !hasShift {
				clickedItem := me.Items().Get(int(lvhti.IItem))
				if !clickedItem.IsSelected() {
					me.Items().SelectAll(false)
					clickedItem.Select(true)
				}
				clickedItem.Focus()
			}
		} else if !hasCtrl && !hasShift { // no item was right-clicked
			me.Items().SelectAll(false)
		}
		me.Hwnd().SetFocus() // because a right-click won't set the focus by itself

	} else { // usually fired with the context keyboard key
		focusedItem := me.Items().Focused()
		if focusedItem != nil && focusedItem.IsVisible() { // there is a focused item, and it's visible
			rcItem := focusedItem.Rect(co.LVIR_BOUNDS)
			menuPos.X = rcItem.Left + 16 // arbitrary
			menuPos.Y = rcItem.Top + (rcItem.Bottom-rcItem.Top)/2
		} else { // no item is focused and visible
			menuPos.X = 6 // arbitrary
			menuPos.Y = 10
		}
	}

	me.contextMenu.ShowAtPoint(*menuPos, me.Hwnd().GetParent(), me.Hwnd())
}

//------------------------------------------------------------------------------

type _ListViewColumnCollection struct {
	ctrl *ListView
}

// Constructor.
func _NewListViewColumnCollection(ctrl *ListView) *_ListViewColumnCollection {
	return &_ListViewColumnCollection{
		ctrl: ctrl,
	}
}

// Appends a new column, returning it.
//
// Width will be adjusted to the current system DPI.
func (me *_ListViewColumnCollection) Add(
	text string, width int) *ListViewColumn {

	colWidth := Size{Cx: width, Cy: 0}
	_global.MultiplyDpi(nil, &colWidth)

	textBuf := win.Str.ToUint16Slice(text)

	lvc := win.LVCOLUMN{
		Mask:    co.LVCF_TEXT | co.LVCF_WIDTH,
		PszText: &textBuf[0],
		Cx:      int32(colWidth.Cx),
	}
	newIdx := int(
		me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_INSERTCOLUMN),
			0xffff, win.LPARAM(unsafe.Pointer(&lvc))),
	)
	if newIdx == -1 {
		panic(fmt.Sprintf("LVM_INSERTCOLUMN failed \"%s\".", text))
	}
	return me.Get(newIdx)
}

// Appends many columns at once.
//
// Width will be adjusted to the current system DPI.
func (me *_ListViewColumnCollection) AddMany(
	texts []string, widths []int) *ListView {

	if len(texts) != len(widths) {
		panic("Columns().Add() texts/widths mismatch.")
	}

	for i := range texts {
		me.Add(texts[i], widths[i])
	}
	return me.ctrl
}

// Retrieves the number of columns.
func (me *_ListViewColumnCollection) Count() int {
	hHeader := win.HWND(me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETHEADER), 0, 0))
	if hHeader == 0 {
		panic("LVM_GETHEADER failed.")
	}

	count := int(
		hHeader.SendMessage(co.WM(co.HDM_GETITEMCOUNT), 0, 0),
	)
	if count == -1 {
		panic("HDM_GETITEMCOUNT failed.")
	}
	return count
}

// Returns the column at the given index.
//
// Does not perform bound checking.
func (me *_ListViewColumnCollection) Get(index int) *ListViewColumn {
	return _NewListViewColumn(me.ctrl, index)
}

//------------------------------------------------------------------------------

// A single column of a ListView control.
type ListViewColumn struct {
	ctrl  *ListView
	index int
}

// Constructor.
func _NewListViewColumn(ctrl *ListView, index int) *ListViewColumn {
	return &ListViewColumn{
		ctrl:  ctrl,
		index: index,
	}
}

// Returns the index of this column.
func (me *ListViewColumn) Index() int {
	return me.index
}

// Returns the texts of the selected items, under this column.
func (me *ListViewColumn) SelectedItemsTexts() []string {
	selItems := me.ctrl.Items().Selected() // retrieve all selected items
	texts := make([]string, 0, len(selItems))
	for i := range selItems {
		texts = append(texts, selItems[i].SubItemText(me.index))
	}
	return texts
}

// Returns the texts of all items, under this column.
func (me *ListViewColumn) ItemTexts() []string {
	numItems := me.ctrl.Items().Count()
	texts := make([]string, 0, numItems)
	for i := 0; i < numItems; i++ {
		texts = append(texts, me.ctrl.Items().Get(i).SubItemText(me.index))
	}
	return texts
}

// Sets the text of this column.
func (me *ListViewColumn) SetText(text string) *ListViewColumn {
	textBuf := win.Str.ToUint16Slice(text)
	lvc := win.LVCOLUMN{
		ISubItem: int32(me.index),
		Mask:     co.LVCF_TEXT,
		PszText:  &textBuf[0],
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_SETCOLUMN),
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMN failed to set text \"%s\".", text))
	}
	return me
}

// Sets the width of the column, adjusted to the current system DPI.
func (me *ListViewColumn) SetWidth(width int) *ListViewColumn {
	colWidth := Size{Cx: width, Cy: 0}
	_global.MultiplyDpi(nil, &colWidth)

	me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_SETCOLUMNWIDTH),
		win.WPARAM(me.index), win.LPARAM(colWidth.Cx))
	return me
}

// Resizes the column to fill the remaining space.
func (me *ListViewColumn) SetWidthToFill() *ListViewColumn {
	numCols := me.ctrl.Columns().Count()
	cxUsed := 0

	for i := 0; i < numCols; i++ {
		if i != me.index {
			cxUsed += me.ctrl.Columns().Get(i).Width() // retrieve cx of each column, but us
		}
	}

	rc := me.ctrl.Hwnd().GetClientRect() // list view client area
	me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_SETCOLUMNWIDTH),
		win.WPARAM(me.index), win.LPARAM(int(rc.Right)-cxUsed)) // fill available space
	return me
}

// Retrieves the text of this column.
func (me *ListViewColumn) Text() string {
	buf := [128]uint16{} // arbitrary
	lvc := win.LVCOLUMN{
		ISubItem:   int32(me.index),
		Mask:       co.LVCF_TEXT,
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETCOLUMN),
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic("LVM_GETCOLUMN failed to get text.")
	}
	return syscall.UTF16ToString(buf[:])
}

// Retrieves the width of the column.
func (me *ListViewColumn) Width() int {
	cx := int(
		me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETCOLUMNWIDTH),
			win.WPARAM(me.index), 0),
	)
	if cx == 0 {
		panic("LVM_GETCOLUMNWIDTH failed.")
	}
	return cx
}

//------------------------------------------------------------------------------

type _ListViewItemCollection struct {
	ctrl *ListView
}

// Constructor.
func _NewListViewItemCollection(ctrl *ListView) *_ListViewItemCollection {
	return &_ListViewItemCollection{
		ctrl: ctrl,
	}
}

// Adds a new item, returning it.
func (me *_ListViewItemCollection) Add(
	text string, subItemTexts ...string) *ListViewItem {

	return me.AddWithIcon(-1, text, subItemTexts...)
}

// Adds a new item, returning it. Receives the zero-based index of the icon from
// the associated ImageList, which must have been set with SetImageList().
func (me *_ListViewItemCollection) AddWithIcon(
	iconIndex int, text string, subItemTexts ...string) *ListViewItem {

	textBuf := win.Str.ToUint16Slice(text)
	lvi := win.LVITEM{
		Mask:    co.LVIF_TEXT | co.LVIF_IMAGE,
		PszText: &textBuf[0],
		IImage:  int32(iconIndex),
		IItem:   0x0fff_ffff, // insert as the last one
	}
	newIdx := int(
		me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_INSERTITEM), 0,
			win.LPARAM(unsafe.Pointer(&lvi))),
	)
	if newIdx == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM failed \"%s\".", text))
	}

	newItem := me.Get(int(newIdx))
	for i, subItemText := range subItemTexts { // for the sub items, if any
		newItem.SetSubItemText(i+1, subItemText)
	}
	return newItem
}

// Retrieves all items at once.
func (me *_ListViewItemCollection) All() []*ListViewItem {
	items := make([]*ListViewItem, 0, me.Count())
	idx := -1
	for {
		idx = int(
			me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETNEXTITEM),
				win.WPARAM(idx), win.LPARAM(co.LVNI_ALL)),
		)
		if idx == -1 {
			break
		}
		items = append(items, me.Get(idx))
	}
	return items
}

// Retrieves the number of items.
func (me *_ListViewItemCollection) Count() int {
	count := int(
		me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETITEMCOUNT), 0, 0),
	)
	if count == -1 {
		panic("LVM_GETITEMCOUNT failed.")
	}
	return count
}

// Deletes all items at once.
func (me *_ListViewItemCollection) DeleteAll() *ListView {
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_DELETEALLITEMS), 0, 0)
	if ret == 0 {
		panic("LVM_DELETEALLITEMS failed.")
	}
	return me.ctrl
}

// Deletes all selected items at once.
func (me *_ListViewItemCollection) DeleteSelected() *ListView {
	selItems := me.Selected()
	for i := len(selItems) - 1; i >= 0; i-- { // from last to first because indexes are sequential
		selItems[i].Delete()
	}
	return me.ctrl
}

// Searches for an item with the given exact text, case-insensitive.
//
// Returns nil if not found.
func (me *_ListViewItemCollection) Find(text string) *ListViewItem {
	buf := win.Str.ToUint16Slice(text)
	lvfi := win.LVFINDINFO{
		Flags: co.LVFI_STRING,
		Psz:   &buf[0],
	}
	wp := -1
	idx := int(
		me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_FINDITEM),
			win.WPARAM(wp), win.LPARAM(unsafe.Pointer(&lvfi))),
	)
	if idx == -1 {
		return nil // not found
	}
	return me.Get(idx)
}

// Retrieves the currently focused item, or nil if none.
func (me *_ListViewItemCollection) Focused() *ListViewItem {
	idx := -1
	idx = int(
		me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETNEXTITEM),
			win.WPARAM(idx), win.LPARAM(co.LVNI_FOCUSED)),
	)
	if idx == -1 {
		return nil
	}
	return me.Get(idx)
}

// Returns the item at the given index.
//
// Does not perform bound checking.
func (me *_ListViewItemCollection) Get(index int) *ListViewItem {
	return _NewListViewItem(me.ctrl, index)
}

// Sends LVM_HITTEST to determine the item at specified position, if any. Pos
// coordinates must be relative to list view.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-hittest
func (me *_ListViewItemCollection) HitTest(pos win.POINT) *win.LVHITTESTINFO {
	lvhti := win.LVHITTESTINFO{
		Pt: pos,
	}
	wp := -1 // Vista: retrieve iGroup and iSubItem
	me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_HITTEST),
		win.WPARAM(wp), win.LPARAM(unsafe.Pointer(&lvhti)))
	return &lvhti
}

// Selects or deselects all items at once.
func (me *_ListViewItemCollection) SelectAll(isSelected bool) *ListView {
	state := co.LVIS_NONE
	if isSelected {
		state = co.LVIS_SELECTED
	}

	lvi := win.LVITEM{
		State:     state,
		StateMask: co.LVIS_SELECTED,
	}
	idx := -1
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_SETITEMSTATE),
		win.WPARAM(idx), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
	return me.ctrl
}

// Retrieves the currently selected items, sorted by index.
func (me *_ListViewItemCollection) Selected() []*ListViewItem {
	items := make([]*ListViewItem, 0, me.SelectedCount())
	idx := -1
	for {
		idx = int(
			me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETNEXTITEM),
				win.WPARAM(idx), win.LPARAM(co.LVNI_SELECTED)),
		)
		if idx == -1 {
			break
		}
		items = append(items, me.Get(idx))
	}
	return items
}

// Retrieves the number of selected items.
func (me *_ListViewItemCollection) SelectedCount() int {
	count := int(
		me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETSELECTEDCOUNT), 0, 0),
	)
	if count == -1 {
		panic("LVM_GETSELECTEDCOUNT failed.")
	}
	return count
}

// Retrieves the topmost visible item, or nil of none.
func (me *_ListViewItemCollection) TopmostVisible() *ListViewItem {
	idx := int(
		me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETTOPINDEX), 0, 0),
	)
	if idx == -1 {
		return nil
	}
	return me.Get(idx)
}

//------------------------------------------------------------------------------

// A single item of a ListView control.
type ListViewItem struct {
	ctrl  *ListView
	index int
}

// Constructor.
func _NewListViewItem(ctrl *ListView, index int) *ListViewItem {
	return &ListViewItem{
		ctrl:  ctrl,
		index: index,
	}
}

// Deletes this item.
func (me *ListViewItem) Delete() {
	if me.index >= me.ctrl.Items().Count() { // index out of bounds: ignore
		return
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_DELETEITEM),
		win.WPARAM(me.index), 0)
	if ret == 0 {
		panic(fmt.Sprintf("LVM_DELETEITEM failed, index %d.", me.index))
	}
}

// Scrolls the list view so this item becomes visible.
func (me *ListViewItem) EnsureVisible() *ListViewItem {
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_ENSUREVISIBLE),
		win.WPARAM(me.index), win.LPARAM(1)) // always entirely visible
	if ret == 0 {
		panic("LVM_ENSUREVISIBLE failed.")
	}
	return me
}

// Sets the item as the currently focused one.
func (me *ListViewItem) Focus() *ListViewItem {
	lvi := win.LVITEM{
		State:     co.LVIS_FOCUSED,
		StateMask: co.LVIS_FOCUSED,
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_SETITEMSTATE),
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
	return me
}

// Retrieves the image index, or -1 if no image.
func (me *ListViewItem) IconIndex() int {
	lvi := win.LVITEM{
		IItem: int32(me.index),
		Mask:  co.LVIF_IMAGE,
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETITEM),
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_GETITEM failed.")
	}
	return int(lvi.IImage)
}

// Returns the index of this item.
func (me *ListViewItem) Index() int {
	return me.index
}

// Tells if the item is the currently focused one.
func (me *ListViewItem) IsFocused() bool {
	return co.LVIS(
		me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETITEMSTATE),
			win.WPARAM(me.index), win.LPARAM(co.LVIS_FOCUSED)),
	) == co.LVIS_FOCUSED
}

// Tells if the item is currently selected.
func (me *ListViewItem) IsSelected() bool {
	return co.LVIS(
		me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETITEMSTATE),
			win.WPARAM(me.index), win.LPARAM(co.LVIS_SELECTED)),
	) == co.LVIS_SELECTED
}

// Checks if this item is visible.
func (me *ListViewItem) IsVisible() bool {
	return me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_ISITEMVISIBLE),
		win.WPARAM(me.index), 0) != 0
}

// Retrieves the LPARAM associated to this item.
func (me *ListViewItem) LParam() win.LPARAM {
	lvi := win.LVITEM{
		IItem: int32(me.index),
		Mask:  co.LVIF_PARAM,
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETITEM),
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_GETITEM failed.")
	}
	return lvi.LParam
}

// Retrieves bound coordinates of the item with LVM_GETITEMRECT. Coordinates are
// relative to list view.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getitemrect
func (me *ListViewItem) Rect(portion co.LVIR) *win.RECT {
	rcItem := &win.RECT{
		Left: int32(portion),
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETITEMRECT),
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(rcItem)))
	if ret == 0 {
		panic("LVM_GETITEMRECT failed.")
	}
	return rcItem
}

// Selects or unselects the item.
func (me *ListViewItem) Select(isSelected bool) *ListViewItem {
	state := co.LVIS_NONE
	if isSelected {
		state = co.LVIS_SELECTED
	}

	lvi := win.LVITEM{
		State:     state,
		StateMask: co.LVIS_SELECTED,
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_SETITEMSTATE),
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
	return me
}

// Sets the image index, or -1 for no image.
func (me *ListViewItem) SetIconIndex(index int) *ListViewItem {
	lvi := win.LVITEM{
		IItem:  int32(me.index),
		Mask:   co.LVIF_IMAGE,
		IImage: int32(index),
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_SETITEM),
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_GETITEM failed.")
	}
	return me
}

// Sets the LPARAM associated to this item.
func (me *ListViewItem) SetLParam(lParam win.LPARAM) *ListViewItem {
	lvi := win.LVITEM{
		IItem:  int32(me.index),
		Mask:   co.LVIF_PARAM,
		LParam: lParam,
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_SETITEM),
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_GETITEM failed.")
	}
	return me
}

// Sets the text under the given column.
func (me *ListViewItem) SetSubItemText(
	columnIndex int, text string) *ListViewItem {

	textBuf := win.Str.ToUint16Slice(text)
	lvi := win.LVITEM{
		ISubItem: int32(columnIndex),
		PszText:  &textBuf[0],
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_SETITEMTEXT),
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMTEXT failed \"%s\".", text))
	}
	return me
}

// Sets the text under the first column.
func (me *ListViewItem) SetText(text string) *ListViewItem {
	return me.SetSubItemText(0, text)
}

// Retrieves the text under the given column.
func (me *ListViewItem) SubItemText(columnIndex int) string {
	buf := [256]uint16{} // arbitrary
	lvi := win.LVITEM{
		ISubItem:   int32(columnIndex),
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_GETITEMTEXT),
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret < 0 {
		panic("LVM_GETITEMTEXT failed.")
	}
	return syscall.UTF16ToString(buf[:])
}

// Retrieves the text under the first column.
func (me *ListViewItem) Text() string {
	return me.SubItemText(0)
}

// Updates the item with LVM_UPDATE.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-update
func (me *ListViewItem) Update() *ListViewItem {
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.LVM_UPDATE),
		win.WPARAM(me.index), 0)
	if ret == 0 {
		panic("LVM_UPDATE failed.")
	}
	return me
}

//------------------------------------------------------------------------------

// ListView control notifications.
type _EventsListView struct {
	ctrl *_NativeControlBase
}

// Constructor.
func _NewEventsListView(ctrl *_NativeControlBase) *_EventsListView {
	return &_EventsListView{
		ctrl: ctrl,
	}
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-begindrag
func (me *_EventsListView) LvnBeginDrag(userFunc func(p *win.NMLISTVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_BEGINDRAG), func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-beginlabeledit
func (me *_EventsListView) LvnBeginLabelEdit(userFunc func(p *win.NMLVDISPINFO) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.LVN_BEGINLABELEDIT), func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMLVDISPINFO)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-beginrdrag
func (me *_EventsListView) LvnBeginRDrag(userFunc func(p *win.NMLISTVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_BEGINRDRAG), func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-beginscroll
func (me *_EventsListView) LvnBeginScroll(userFunc func(p *win.NMLVSCROLL)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_BEGINSCROLL), func(p unsafe.Pointer) {
		userFunc((*win.NMLVSCROLL)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-columnclick
func (me *_EventsListView) LvnColumnClick(userFunc func(p *win.NMLISTVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_COLUMNCLICK), func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-columndropdown
func (me *_EventsListView) LvnColumnDropDown(userFunc func(p *win.NMLISTVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_COLUMNDROPDOWN), func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-columnoverflowclick
func (me *_EventsListView) LvnColumnOverflowClick(userFunc func(p *win.NMLISTVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_COLUMNOVERFLOWCLICK), func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-deleteallitems
func (me *_EventsListView) LvnDeleteAllItems(userFunc func(p *win.NMLISTVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_DELETEALLITEMS), func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-deleteitem
func (me *_EventsListView) LvnDeleteItem(userFunc func(p *win.NMLISTVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_DELETEITEM), func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-endlabeledit
func (me *_EventsListView) LvnEndLabelEdit(userFunc func(p *win.NMLVDISPINFO) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.LVN_ENDLABELEDIT), func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMLVDISPINFO)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-endscroll
func (me *_EventsListView) LvnEndScroll(userFunc func(p *win.NMLVSCROLL)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_ENDSCROLL), func(p unsafe.Pointer) {
		userFunc((*win.NMLVSCROLL)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-getdispinfo
func (me *_EventsListView) LvnGetDispInfo(userFunc func(p *win.NMLVDISPINFO)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_GETDISPINFO), func(p unsafe.Pointer) {
		userFunc((*win.NMLVDISPINFO)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-getemptymarkup
func (me *_EventsListView) LvnGetEmptyMarkup(userFunc func(p *win.NMLVEMPTYMARKUP) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.LVN_GETEMPTYMARKUP), func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMLVEMPTYMARKUP)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-getinfotip
func (me *_EventsListView) LvnGetInfoTip(userFunc func(p *win.NMLVGETINFOTIP)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_GETINFOTIP), func(p unsafe.Pointer) {
		userFunc((*win.NMLVGETINFOTIP)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-hottrack
func (me *_EventsListView) LvnHotTrack(userFunc func(p *win.NMLISTVIEW) int) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.LVN_HOTTRACK), func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMLISTVIEW)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-incrementalsearch
func (me *_EventsListView) LvnIncrementalSearch(userFunc func(p *win.NMLVFINDITEM) int) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.LVN_INCREMENTALSEARCH), func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMLVFINDITEM)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-insertitem
func (me *_EventsListView) LvnInsertItem(userFunc func(p *win.NMLISTVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_INSERTITEM), func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-itemactivate
func (me *_EventsListView) LvnItemActivate(userFunc func(p *win.NMITEMACTIVATE)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_ITEMACTIVATE), func(p unsafe.Pointer) {
		userFunc((*win.NMITEMACTIVATE)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-itemchanged
func (me *_EventsListView) LvnItemChanged(userFunc func(p *win.NMLISTVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_ITEMCHANGED), func(p unsafe.Pointer) {
		userFunc((*win.NMLISTVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-itemchanging
func (me *_EventsListView) LvnItemChanging(userFunc func(p *win.NMLISTVIEW) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.LVN_ITEMCHANGING), func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMLISTVIEW)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-keydown
func (me *_EventsListView) LvnKeyDown(userFunc func(p *win.NMLVKEYDOWN)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_KEYDOWN), func(p unsafe.Pointer) {
		userFunc((*win.NMLVKEYDOWN)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-linkclick
func (me *_EventsListView) LvnLinkClick(userFunc func(p *win.NMLVLINK)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_LINKCLICK), func(p unsafe.Pointer) {
		userFunc((*win.NMLVLINK)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-marqueebegin
func (me *_EventsListView) LvnMarqueeBegin(userFunc func() uint) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.LVN_MARQUEEBEGIN), func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-odcachehint
func (me *_EventsListView) LvnODCacheHint(userFunc func(p *win.NMLVCACHEHINT)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_ODCACHEHINT), func(p unsafe.Pointer) {
		userFunc((*win.NMLVCACHEHINT)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-odfinditem
func (me *_EventsListView) LvnODFindItem(userFunc func(p *win.NMLVFINDITEM) int) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.LVN_ODFINDITEM), func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMLVFINDITEM)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-odstatechanged
func (me *_EventsListView) LvnODStateChanged(userFunc func(p *win.NMLVODSTATECHANGE)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_ODSTATECHANGED), func(p unsafe.Pointer) {
		userFunc((*win.NMLVODSTATECHANGE)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-setdispinfo
func (me *_EventsListView) LvnSetDispInfo(userFunc func(p *win.NMLVDISPINFO)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.LVN_SETDISPINFO), func(p unsafe.Pointer) {
		userFunc((*win.NMLVDISPINFO)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-list-view
func (me *_EventsListView) NmClick(userFunc func(p *win.NMITEMACTIVATE)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_CLICK, func(p unsafe.Pointer) {
		userFunc((*win.NMITEMACTIVATE)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-customdraw-list-view
func (me *_EventsListView) NmCustomDraw(userFunc func(p *win.NMLVCUSTOMDRAW) co.CDRF) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMLVCUSTOMDRAW)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-dblclk-list-view
func (me *_EventsListView) NmDblClk(userFunc func(p *win.NMITEMACTIVATE)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_DBLCLK, func(p unsafe.Pointer) {
		userFunc((*win.NMITEMACTIVATE)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-hover-list-view
func (me *_EventsListView) NmHover(userFunc func() uint) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM_HOVER, func(_ unsafe.Pointer) uintptr {
		return uintptr(userFunc())
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-killfocus-list-view
func (me *_EventsListView) NmKillFocus(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_KILLFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rclick-list-view
func (me *_EventsListView) NmRClick(userFunc func(p *win.NMITEMACTIVATE)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_RCLICK, func(p unsafe.Pointer) {
		userFunc((*win.NMITEMACTIVATE)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-list-view
func (me *_EventsListView) LvnRDblClk(userFunc func(p *win.NMITEMACTIVATE)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_RDBLCLK, func(p unsafe.Pointer) {
		userFunc((*win.NMITEMACTIVATE)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-list-view-
func (me *_EventsListView) LvnReleasedCapture(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_RELEASEDCAPTURE, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-return-list-view-
func (me *_EventsListView) LvnReturn(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_RETURN, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-setfocus-list-view-
func (me *_EventsListView) LvnSetFocus(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_SETFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}
