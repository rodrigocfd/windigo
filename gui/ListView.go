/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
	"unsafe"
	"wingows/co"
	"wingows/gui/wm"
	"wingows/win"
)

// Native list view control.
type ListView struct {
	controlNativeBase
	hMenuContext win.HMENU // if set, will be shown with right-click
}

// Adds a new column; returns the newly inserted column.
func (me *ListView) AddColumn(text string, width uint32) *ListViewColumn {
	textBuf := win.StrToSlice(text)
	lvc := win.LVCOLUMN{
		Mask:    co.LVCF_TEXT | co.LVCF_WIDTH,
		PszText: uintptr(unsafe.Pointer(&textBuf[0])),
		Cx:      int32(width),
	}
	newIdx := me.sendLvmMessage(co.LVM_INSERTCOLUMN, 0xFFFF,
		win.LPARAM(unsafe.Pointer(&lvc)))
	if int(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTCOLUMN failed \"%s\".", text))
	}
	return &ListViewColumn{
		owner: me,
		index: uint32(newIdx),
	}
}

// Adds a new item; returns the newly inserted item.
func (me *ListView) AddItem(text string) *ListViewItem {
	textBuf := win.StrToSlice(text)
	lvi := win.LVITEM{
		Mask:    co.LVIF_TEXT,
		PszText: uintptr(unsafe.Pointer(&textBuf[0])),
		IItem:   0x0FFFFFFF, // insert as the last one
	}
	newIdx := me.sendLvmMessage(co.LVM_INSERTITEM, 0,
		win.LPARAM(unsafe.Pointer(&lvi)))
	if int(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM failed \"%s\".", text))
	}
	return &ListViewItem{
		owner: me,
		index: uint32(newIdx),
	}
}

// Adds a new item; returns the newly inserted item.
// Before call this method, attach an image list and load its icons.
func (me *ListView) AddItemWithIcon(text string,
	iconIndex uint32) *ListViewItem {

	textBuf := win.StrToSlice(text)
	lvi := win.LVITEM{
		Mask:    co.LVIF_TEXT | co.LVIF_IMAGE,
		PszText: uintptr(unsafe.Pointer(&textBuf[0])),
		IImage:  int32(iconIndex),
		IItem:   0x0FFFFFFF, // insert as the last one
	}
	newIdx := me.sendLvmMessage(co.LVM_INSERTITEM, 0,
		win.LPARAM(unsafe.Pointer(&lvi)))
	if int(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM failed \"%s\".", text))
	}
	return &ListViewItem{
		owner: me,
		index: uint32(newIdx),
	}
}

// Adds many items at once.
func (me *ListView) AddItems(texts []string) *ListView {
	for i := range texts {
		me.AddItem(texts[i])
	}
	return me
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *ListView) Create(parent Window, x, y int32, width, height uint32,
	exStyles co.WS_EX, styles co.WS,
	lvExStyles co.LVS_EX, lvStyles co.LVS) *ListView {

	me.installSubclass()

	x, y, width, height = globalDpi.multiply(x, y, width, height)
	me.controlNativeBase.create(exStyles,
		"SysListView32", "", styles|co.WS(lvStyles),
		x, y, width, height, parent)

	if lvExStyles != co.LVS_EX(0) {
		me.SetExtendedStyle(lvExStyles, lvExStyles)
	}
	return me
}

// Calls CreateWindowEx(). List view control will have LVS_REPORT style.
// Position and size will be adjusted to the current system DPI.
func (me *ListView) CreateReport(parent Window, x, y int32,
	width, height uint32) *ListView {

	return me.Create(parent, x, y, width, height,
		co.WS_EX_CLIENTEDGE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.LVS_EX_FULLROWSELECT,
		co.LVS_REPORT|co.LVS_NOSORTHEADER|co.LVS_SHOWSELALWAYS)
}

// Calls CreateWindowEx(). List view control will have LVS_REPORT and
// LVS_SORTASCENDING styles. Position and size will be adjusted to the current system DPI.
func (me *ListView) CreateSortedReport(parent Window, x, y int32,
	width, height uint32) *ListView {

	return me.Create(parent, x, y, width, height,
		co.WS_EX_CLIENTEDGE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.LVS_EX_FULLROWSELECT,
		co.LVS_REPORT|co.LVS_NOSORTHEADER|co.LVS_SHOWSELALWAYS|co.LVS_SORTASCENDING)
}

// Retrieves the column at the given index.
func (me *ListView) Column(index uint32) *ListViewColumn {
	numCols := me.ColumnCount()
	if index >= numCols {
		panic("Trying to retrieve column with index out of bounds.")
	}
	return &ListViewColumn{
		owner: me,
		index: index,
	}
}

// Retrieves the number of columns with LVM_GETHEADER and HDM_GETITEMCOUNT.
func (me *ListView) ColumnCount() uint32 {
	hHeader := win.HWND(me.sendLvmMessage(co.LVM_GETHEADER, 0, 0))
	if hHeader == 0 {
		panic("LVM_GETHEADER failed.")
	}

	count := hHeader.SendMessage(co.WM(co.HDM_GETITEMCOUNT), 0, 0)
	if int(count) == -1 {
		panic("HDM_GETITEMCOUNT failed.")
	}
	return uint32(count)
}

// Deletes all items with LVM_DELETEALLITEMS.
func (me *ListView) DeleteAllItems() *ListView {
	ret := me.sendLvmMessage(co.LVM_DELETEALLITEMS, 0, 0)
	if ret == 0 {
		panic("LVM_DELETEALLITEMS failed.")
	}
	return me
}

// Retrieves extended styles with LVM_GETEXTENDEDLISTVIEWSTYLE.
func (me *ListView) ExtendedStyle() co.LVS_EX {
	return co.LVS_EX(me.sendLvmMessage(co.LVM_GETEXTENDEDLISTVIEWSTYLE, 0, 0))
}

// Searches for an item with the given exact text, case-insensitive.
func (me *ListView) FindItem(text string) *ListViewItem {
	buf := win.StrToSlice(text)
	lvfi := win.LVFINDINFO{
		Flags: co.LVFI_STRING,
		Psz:   uintptr(unsafe.Pointer(&buf[0])),
	}
	wp := -1
	idx := me.sendLvmMessage(co.LVM_FINDITEM,
		win.WPARAM(wp), win.LPARAM(unsafe.Pointer(&lvfi)))
	if int(idx) == -1 {
		return nil // not found
	}
	return &ListViewItem{
		owner: me,
		index: uint32(idx),
	}
}

// Sends LVM_ISGROUPVIEWENABLED.
func (me *ListView) GroupViewEnabled() bool {
	return me.sendLvmMessage(co.LVM_ISGROUPVIEWENABLED, 0, 0) >= 0
}

// Sends LVM_HITTEST to determine the item at specified position, if any.
// Pos coordinates must be relative to list view.
func (me *ListView) HitTest(pos *win.POINT) *win.LVHITTESTINFO {
	lvhti := &win.LVHITTESTINFO{
		Pt: *pos,
	}
	wp := -1 // Vista: retrieve iGroup and iSubItem
	me.sendLvmMessage(co.LVM_HITTEST,
		win.WPARAM(wp), win.LPARAM(unsafe.Pointer(lvhti)))
	return lvhti
}

// Retrieves the associated HIMAGELIST by sending LVM_GETIMAGELIST.
func (me *ListView) ImageList(typeImgList co.LVSIL) win.HIMAGELIST {
	return win.HIMAGELIST(
		me.sendLvmMessage(co.LVM_GETIMAGELIST, win.WPARAM(typeImgList), 0),
	)
}

// Retrieves the item at the given index.
func (me *ListView) Item(index uint32) *ListViewItem {
	numItems := me.ItemCount()
	if index >= numItems {
		panic("Trying to retrieve item with index out of bounds.")
	}
	return &ListViewItem{
		owner: me,
		index: index,
	}
}

// Retrieves the number of items with LVM_GETITEMCOUNT.
func (me *ListView) ItemCount() uint32 {
	count := me.sendLvmMessage(co.LVM_GETITEMCOUNT, 0, 0)
	if int(count) == -1 {
		panic("LVM_GETITEMCOUNT failed.")
	}
	return uint32(count)
}

// Sends LVM_GETNEXTITEM with -1 as item index.
// Returns nil if none.
func (me *ListView) NextItem(relationship co.LVNI) *ListViewItem {
	idx := -1
	allItems := ListViewItem{
		owner: me,
		index: uint32(idx),
	}
	return allItems.NextItem(relationship)
}

// Calls LVM_GETNEXTITEM sequentially to retrieve all matched items.
func (me *ListView) NextItemAll(relationship co.LVNI) []ListViewItem {
	idx := -1
	items := make([]ListViewItem, 0)
	for {
		idx = int(me.sendLvmMessage(co.LVM_GETNEXTITEM,
			win.WPARAM(idx), win.LPARAM(relationship)))
		if idx == -1 {
			break
		}
		items = append(items, ListViewItem{
			owner: me,
			index: uint32(idx),
		})
	}
	return items
}

// Sends LVM_SCROLL.
func (me *ListView) Scroll(pxHorz, pxVert int32) *ListView {
	ret := me.sendLvmMessage(co.LVM_SCROLL,
		win.WPARAM(pxHorz), win.LPARAM(pxVert))
	if ret == 0 {
		panic("LVM_SCROLL failed.")
	}
	return me
}

// Retrieves the number of selected items with LVM_GETSELECTEDCOUNT.
func (me *ListView) SelectedItemCount() uint32 {
	count := me.sendLvmMessage(co.LVM_GETSELECTEDCOUNT, 0, 0)
	if int(count) == -1 {
		panic("LVM_GETSELECTEDCOUNT failed.")
	}
	return uint32(count)
}

// Defines a menu to be shown as the context menu for the list view.
func (me *ListView) SetContextMenu(hMenu win.HMENU) *ListView {
	me.hMenuContext = hMenu
	return me
}

// Sends LVM_SETEXTENDEDLISTVIEWSTYLE.
func (me *ListView) SetExtendedStyle(mask, exStyle co.LVS_EX) *ListView {
	me.sendLvmMessage(co.LVM_SETEXTENDEDLISTVIEWSTYLE,
		win.WPARAM(mask), win.LPARAM(exStyle))
	return me
}

// Sends LVM_SETIMAGELIST.
// Returns image list previously associated.
func (me *ListView) SetImageList(typeImgList co.LVSIL,
	hImgList win.HIMAGELIST) win.HIMAGELIST {

	return win.HIMAGELIST(
		me.sendLvmMessage(co.LVM_SETIMAGELIST,
			win.WPARAM(typeImgList), win.LPARAM(hImgList)),
	)
}

// Sends WM_SETREDRAW to enable or disable UI updates.
func (me *ListView) SetRedraw(allowRedraw bool) *ListView {
	wp := 0
	if allowRedraw {
		wp = 1
	}
	me.hwnd.SendMessage(co.WM_SETREDRAW, win.WPARAM(wp), 0)
	return me
}

// Sends LVM_SETITEMSTATE with index -1, which affects all items.
func (me *ListView) SetStateAllItems(
	state co.LVIS, stateMask co.LVIS) *ListView {

	idx := -1
	allItems := ListViewItem{
		owner: me,
		index: uint32(idx),
	}
	allItems.SetState(state, stateMask)
	return me
}

// Sends LVM_SETVIEW.
func (me *ListView) SetView(view co.LV_VIEW) *ListView {
	if int(me.sendLvmMessage(co.LVM_SETVIEW, 0, 0)) == -1 {
		panic("LVM_SETVIEW failed.")
	}
	return me
}

// Returns the width of a string using list view current font, with
// LVM_GETSTRINGWIDTH.
func (me *ListView) StringWidth(text string) uint32 {
	ret := me.sendLvmMessage(co.LVM_GETSTRINGWIDTH,
		0, win.LPARAM(unsafe.Pointer(win.StrToPtr(text))))
	if ret == 0 {
		panic("LVM_GETSTRINGWIDTH failed.")
	}
	return uint32(ret)
}

// Retrieves the index of the topmost visible item with LVM_GETTOPINDEX.
func (me *ListView) TopIndex() uint32 {
	return uint32(me.sendLvmMessage(co.LVM_GETTOPINDEX, 0, 0))
}

// Retrieves current view with LVM_GETVIEW.
func (me *ListView) View() co.LV_VIEW {
	return co.LV_VIEW(me.sendLvmMessage(co.LVM_GETVIEW, 0, 0))
}

// Adds all subclass message handlers; must be called before creation.
func (me *ListView) installSubclass() {
	me.OnSubclassMsg().WmRButtonDown(func(p wm.RButtonDown) {
		// WM_RBUTTONUP doesn't work, only NM_RCLICK on parent.
		// https://stackoverflow.com/a/30206896
		me.showContextMenu(true, p.HasCtrl(), p.HasShift())
	})

	me.OnSubclassMsg().WmGetDlgCode(func(p wm.GetDlgCode) co.DLGC {
		if !p.IsQuery() && p.VirtualKeyCode() == 'A' && p.HasCtrl() { // Ctrl+A to select all items
			me.SetStateAllItems(co.LVIS_SELECTED, co.LVIS_SELECTED)
			return co.DLGC_WANTCHARS

		} else if !p.IsQuery() && p.VirtualKeyCode() == co.VK_RETURN { // send Enter key to parent
			code := co.LVN_KEYDOWN
			nmlvk := win.NMLVKEYDOWN{
				Hdr: win.NMHDR{
					HWndFrom: me.Hwnd(),
					Code:     uint32(code),
					IdFrom:   uintptr(me.Id()),
				},
				WVKey: co.VK_RETURN,
			}
			me.Hwnd().GetAncestor(co.GA_PARENT).SendMessage(co.WM_NOTIFY,
				win.WPARAM(me.Hwnd()), win.LPARAM(unsafe.Pointer(&nmlvk)))
			return co.DLGC_WANTALLKEYS

		} else if !p.IsQuery() && p.VirtualKeyCode() == co.VK_APPS { // context menu key
			me.showContextMenu(false, p.HasCtrl(), p.HasShift())
		}

		return co.DLGC(
			me.Hwnd().DefSubclassProc(co.WM_GETDLGCODE, p.WParam, p.LParam),
		)
	})
}

// Shows the popup menu anchored at cursor pos.
// This function will block until the menu disappears.
func (me *ListView) showContextMenu(followCursor, hasCtrl, hasShift bool) {
	if me.hMenuContext == 0 {
		return
	}

	menuPos := &win.POINT{} // menu anchor coords, relative to list view

	if followCursor { // usually when fired by a right-click
		menuPos = win.GetCursorPos()        // relative to screen
		me.Hwnd().ScreenToClientPt(menuPos) // now relative to list view
		lvhti := me.HitTest(menuPos)        // to find item below cursor, if any

		if lvhti.IItem != -1 { // an item was right-clicked
			if !hasCtrl && !hasShift {
				clickedItem := me.Item(uint32(lvhti.IItem))
				if (clickedItem.State() & co.LVIS_SELECTED) == 0 { // not selected?
					me.SetStateAllItems(0, co.LVIS_SELECTED)                 // unselect all
					clickedItem.SetState(co.LVIS_SELECTED, co.LVIS_SELECTED) // select it
				}
				clickedItem.SetState(co.LVIS_FOCUSED, co.LVIS_FOCUSED) // and focus it
			}
		} else if !hasCtrl && !hasShift { // no item was right-clicked
			me.SetStateAllItems(0, co.LVIS_SELECTED) // unselect all
		}
		me.Hwnd().SetFocus() // because a right-click won't set the focus by itself

	} else { // usually fired with the context keyboard key
		focusedItem := me.NextItem(co.LVNI_FOCUSED)
		if focusedItem != nil && focusedItem.Visible() { // there is a focused item, and it's visible
			rcItem := focusedItem.Rect(co.LVIR_BOUNDS)
			menuPos.X = rcItem.Left + 16 // arbitrary
			menuPos.Y = rcItem.Top + (rcItem.Bottom-rcItem.Top)/2
		} else { // no item is focused and visible
			menuPos.X = 6 // arbitrary
			menuPos.Y = 10
		}
	}

	fakeMenuStrip := menuStrip{
		hMenu: me.hMenuContext, // note: no janitor member
	}
	fakeMenuStrip.ShowAtPoint(menuPos, me.Hwnd().GetParent(), me.Hwnd())
}

// Simple wrapper.
func (me *ListView) sendLvmMessage(msg co.LVM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.Hwnd().SendMessage(co.WM(msg), wParam, lParam)
}
