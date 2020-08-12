/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// A single item of a menu.
type MenuItem struct {
	owner *Menu
	cmdId int32
}

// Returns the command ID of this menu item.
func (me *MenuItem) CmdId() int32 {
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
	textBuf := win.StrToSlice(text)
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
	return syscall.UTF16ToString(buf)
}
