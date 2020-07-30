/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Manages a menu resource.
type Menu struct {
	hMenu win.HMENU
}

func (me *Menu) AppendItem(cmdId int32, text string) *Menu {
	me.hMenu.AppendMenu(co.MF_STRING, uintptr(cmdId),
		unsafe.Pointer(win.StrToPtr(text)))
	return me
}

func (me *Menu) AppendSeparator() *Menu {
	me.hMenu.AppendMenu(co.MF_SEPARATOR, 0, nil)
	return me
}

// Returns the newly appended menu.
func (me *Menu) AppendSubmenu(text string) *Menu {
	newMenu := &Menu{}
	newMenu.CreatePopup()
	me.hMenu.AppendMenu(co.MF_STRING|co.MF_POPUP, uintptr(newMenu.Hmenu()),
		unsafe.Pointer(win.StrToPtr(text)))
	return newMenu
}

func (me *Menu) CreateMain() *Menu {
	if me.hMenu != 0 {
		panic("Menu already created, CreateMenu not called.")
	}
	me.hMenu = win.CreateMenu()
	return me
}

func (me *Menu) CreatePopup() *Menu {
	if me.hMenu != 0 {
		panic("Menu already created, CreatePopupMenu not called.")
	}
	me.hMenu = win.CreatePopupMenu()
	return me
}

// Only necessary if the menu is not attached to any window.
func (me *Menu) Destroy() {
	if me.hMenu != 0 {
		me.hMenu.DestroyMenu()
		me.hMenu = 0
	}
}

func (me *Menu) EnableManyByCmdId(isEnabled bool, cmdIds []int32) *Menu {
	for _, cmdId := range cmdIds {
		me.ItemByCmdId(cmdId).Enable(isEnabled)
	}
	return me
}

func (me *Menu) EnableManyByPos(isEnabled bool, poss []uint32) *Menu {
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
// Does not validate if such item exists.
func (me *Menu) ItemByCmdId(cmdId int32) *MenuItem {
	return &MenuItem{
		owner: me,
		cmdId: cmdId,
	}
}

// Returns the item at the given position.
// Does not perform bounds checking.
func (me *Menu) ItemByPos(pos uint32) *MenuItem {
	return me.ItemByCmdId(me.hMenu.GetMenuItemID(pos))
}

// Replaces current HMENU with another one.
func (me *Menu) Set(hMenu win.HMENU) *Menu {
	me.hMenu = hMenu
	return me
}

// Shows the popup menu anchored at the given coordinates.
// If hCoordsRelativeTo is zero, coordinates must be relative to hParent.
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
// If pos is not a submenu, returns nil.
func (me *Menu) SubMenu(pos uint32) *Menu {
	hSub := me.hMenu.GetSubMenu(pos)
	if hSub != 0 {
		return &Menu{hMenu: hSub}
	}
	return nil
}
