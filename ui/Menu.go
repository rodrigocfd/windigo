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
//
// https://docs.microsoft.com/en-us/windows/win32/menurc/about-menus
type Menu struct {
	hMenu win.HMENU
}

// Appends a new item to the menu.
func (me *Menu) AppendItem(cmdId int, text string) *Menu {
	me.hMenu.AppendMenu(co.MF_STRING, uintptr(cmdId),
		unsafe.Pointer(win.StrToPtr(text)))
	return me
}

// Appends a new separator to the menu.
func (me *Menu) AppendSeparator() *Menu {
	me.hMenu.AppendMenu(co.MF_SEPARATOR, 0, nil)
	return me
}

// Appends a new submenu to the menu, and returns it.
func (me *Menu) AppendSubmenu(text string) *Menu {
	newMenu := &Menu{}
	newMenu.CreatePopup()
	me.hMenu.AppendMenu(co.MF_STRING|co.MF_POPUP, uintptr(newMenu.Hmenu()),
		unsafe.Pointer(win.StrToPtr(text)))
	return newMenu
}

// Calls CreateMenu(), used when creating the main menu of a window.
//
// If attached to a window, will be automatically destroyed by the system.
func (me *Menu) CreateMain() *Menu {
	if me.hMenu != 0 {
		panic("Menu already created, CreateMenu not called.")
	}
	me.hMenu = win.CreateMenu()
	return me
}

// Calls CreatePopupMenu().
//
// Must be manually destroyed, unless attached to an existing main menu which is
// attached to a window.
func (me *Menu) CreatePopup() *Menu {
	if me.hMenu != 0 {
		panic("Menu already created, CreatePopupMenu not called.")
	}
	me.hMenu = win.CreatePopupMenu()
	return me
}

// Calls DestroyMenu() to free the resources.
func (me *Menu) Destroy() {
	if me.hMenu != 0 {
		me.hMenu.DestroyMenu()
		me.hMenu = 0
	}
}

// Enables or disables many items at once, by command ID.
func (me *Menu) EnableItemsByCmdId(isEnabled bool, cmdIds []int) *Menu {
	for _, cmdId := range cmdIds {
		me.ItemByCmdId(cmdId).Enable(isEnabled)
	}
	return me
}

// Enables or disables many items at once, by zero-based position.
func (me *Menu) EnableItemsByPos(isEnabled bool, poss []uint) *Menu {
	for _, pos := range poss {
		me.ItemByPos(pos).Enable(isEnabled)
	}
	return me
}

// Returns the HMENU handle.
func (me *Menu) Hmenu() win.HMENU {
	return me.hMenu
}

// Returns the item with the given command ID.
//
// Does not validate if such item exists.
func (me *Menu) ItemByCmdId(cmdId int) *MenuItem {
	return &MenuItem{
		owner: me,
		cmdId: cmdId,
	}
}

// Returns the item at the given position.
//
// Does not perform bound checking.
func (me *Menu) ItemByPos(pos uint) *MenuItem {
	return me.ItemByCmdId(int(me.hMenu.GetMenuItemID(uint32(pos))))
}

// Replaces current HMENU with another one.
func (me *Menu) Set(hMenu win.HMENU) *Menu {
	me.hMenu = hMenu
	return me
}

// Shows the popup menu anchored at the given coordinates.
//
// If hCoordsRelativeTo is zero, coordinates must be relative to hParent.
//
// This function will block until the menu disappears.
func (me *Menu) ShowAtPoint(
	pos *win.POINT, hParent, hCoordsRelativeTo win.HWND) {

	if hCoordsRelativeTo == 0 {
		hCoordsRelativeTo = hParent
	}
	hCoordsRelativeTo.ClientToScreenPt(pos) // now relative to screen
	hParent.SetForegroundWindow()
	me.hMenu.TrackPopupMenu(co.TPM_LEFTBUTTON, pos.X, pos.Y, hParent)
	hParent.PostMessage(co.WM_NULL, 0, 0) // necessary according to TrackMenuPopup docs
}

// Returns the submenu at the given position.
//
// If pos is not a submenu, returns nil.
func (me *Menu) SubMenu(pos uint) *Menu {
	hSub := me.hMenu.GetSubMenu(uint32(pos))
	if hSub != 0 {
		return &Menu{hMenu: hSub}
	}
	return nil
}
