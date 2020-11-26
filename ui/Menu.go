/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Native menu resource.
// https://docs.microsoft.com/en-us/windows/win32/menurc/about-menus
type Menu struct {
	hMenu win.HMENU
}

// Calls CreatePopupMenu().
//
// Must be manually destroyed, unless attached to an existing menu which is
// attached to an existing window.
//
// You must defer Destroy().
func NewMenu() *Menu {
	return &Menu{
		hMenu: win.CreatePopupMenu(),
	}
}

// Calls DestroyMenu() to free the resources.
func (me *Menu) Destroy() {
	if me.hMenu != 0 {
		me.hMenu.DestroyMenu()
		me.hMenu = 0
	}
}

// Appends a new item to the menu.
func (me *Menu) AppendItem(cmdId int, text string) *Menu {
	me.hMenu.AppendMenu(co.MF_STRING, uintptr(cmdId),
		unsafe.Pointer(win.Str.ToUint16Ptr(text)))
	return me
}

// Appends a new separator to the menu.
func (me *Menu) AppendSeparator() *Menu {
	me.hMenu.AppendMenu(co.MF_SEPARATOR, 0, nil)
	return me
}

// Appends a new submenu to the menu, returning the newly appended submenu.
func (me *Menu) AppendSubmenu(text string) *Menu {
	newMenu := NewMenu()
	me.hMenu.AppendMenu(co.MF_STRING|co.MF_POPUP, uintptr(newMenu.Hmenu()),
		unsafe.Pointer(win.Str.ToUint16Ptr(text)))
	return newMenu
}

// Enables or disables many items at once, by command ID.
func (me *Menu) EnableItemsByCmdId(isEnabled bool, cmdIds ...int) *Menu {
	for _, cmdId := range cmdIds {
		me.ItemByCmdId(cmdId).Enable(isEnabled)
	}
	return me
}

// Enables or disables many items at once, by zero-based position.
func (me *Menu) EnableItemsByPos(isEnabled bool, indexes ...int) *Menu {
	for _, idx := range indexes {
		me.ItemByPos(idx).Enable(isEnabled)
	}
	return me
}

// Returns the underlying HMENU handle.
func (me *Menu) Hmenu() win.HMENU {
	return me.hMenu
}

// Returns the item with the given command ID.
//
// Does not validate if such item exists.
func (me *Menu) ItemByCmdId(cmdId int) *MenuItem {
	return _NewMenuItem(me, cmdId)
}

// Returns the item at the given position.
//
// Does not perform bound checking.
func (me *Menu) ItemByPos(pos int) *MenuItem {
	cmdId := me.hMenu.GetMenuItemID(uint32(pos))
	return me.ItemByCmdId(int(cmdId))
}

// Retrieves the number if items.
func (me *Menu) ItemCount() int {
	return int(me.hMenu.GetMenuItemCount())
}

// Shows the popup menu anchored at the given coordinates.
//
// If hCoordsRelativeTo is zero, coordinates must be relative to hParent.
//
// This function will block until the menu disappears.
func (me *Menu) ShowAtPoint(
	pos win.POINT, hParent, hCoordsRelativeTo win.HWND) {

	if hCoordsRelativeTo == 0 {
		hCoordsRelativeTo = hParent
	}
	hCoordsRelativeTo.ClientToScreenPt(&pos) // now relative to screen
	hParent.SetForegroundWindow()
	me.hMenu.TrackPopupMenu(co.TPM_LEFTBUTTON, pos.X, pos.Y, hParent)
	hParent.PostMessage(co.WM_NULL, 0, 0) // necessary according to TrackMenuPopup docs
}

// Returns the submenu at the given position.
//
// If pos is not a submenu, returns nil.
func (me *Menu) SubMenu(pos int) *Menu {
	hSub := me.hMenu.GetSubMenu(uint32(pos))
	if hSub != 0 {
		return &Menu{hMenu: hSub}
	}
	return nil
}

//------------------------------------------------------------------------------

// A single item of a menu.
type MenuItem struct {
	owner *Menu
	cmdId int
}

// Constructor.
func _NewMenuItem(owner *Menu, cmdId int) *MenuItem {
	return &MenuItem{
		owner: owner,
		cmdId: cmdId,
	}
}

// Returns the command ID of this menu item.
func (me *MenuItem) CmdId() int {
	return me.cmdId
}

// Calls DeleteMenu() on this item.
func (me *MenuItem) Delete() {
	me.owner.Hmenu().DeleteMenu(uintptr(me.cmdId), co.MF_BYCOMMAND)
}

// Calls EnableMenuItem().
func (me *MenuItem) Enable(isEnabled bool) *MenuItem {
	flags := co.MF_BYCOMMAND
	if isEnabled {
		flags |= co.MF_ENABLED
	} else {
		flags |= co.MF_GRAYED
	}
	me.owner.Hmenu().EnableMenuItem(uintptr(me.cmdId), flags)
	return me
}

// Returns the menu to which this item belongs.
func (me *MenuItem) Owner() *Menu {
	return me.owner
}

// Sets the text to this menu item.
func (me *MenuItem) SetText(text string) *MenuItem {
	textBuf := win.Str.ToUint16Slice(text)
	mii := win.MENUITEMINFO{
		FMask:      co.MIIM_STRING,
		DwTypeData: uintptr(unsafe.Pointer(&textBuf[0])),
	}
	me.owner.Hmenu().SetMenuItemInfo(uintptr(me.cmdId), false, &mii)
	return me
}

// Retrieves the text of this menu item.
func (me *MenuItem) Text() string {
	mii := win.MENUITEMINFO{
		FMask: co.MIIM_STRING,
	}
	me.owner.Hmenu().GetMenuItemInfo(uintptr(me.cmdId), false, &mii) // retrieve length
	mii.Cch++
	buf := make([]uint16, mii.Cch)
	mii.DwTypeData = uintptr(unsafe.Pointer(&buf[0])) // retrieve text
	me.owner.Hmenu().GetMenuItemInfo(uintptr(me.cmdId), false, &mii)
	return win.Str.FromUint16Slice(buf)
}
