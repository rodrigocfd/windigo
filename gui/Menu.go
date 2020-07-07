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

// Helps managing a menu resource.
type Menu struct {
	submenus map[string]win.HMENU   // flat list with each submenu in this menu hierarchy
	items    map[string]menuItemMap // flat list with each menu item, separators are not kept
}

type menuItemMap struct { // we don't simply use menuItem to avoid a *Menu circular reference
	hMenu win.HMENU
	cmdId int32
}

// Creates the root submenu as the horizontal main menu of a window.
// If not attached to any window, it must be destroyed.
func (me *Menu) CreateMain() *menuStrip {
	return me.createRaw(false)
}

// Creates the root submenu as a floating popup menu.
// If not attached to any window, it must be destroyed.
func (me *Menu) CreatePopup() *menuStrip {
	return me.createRaw(true)
}

func (me *Menu) createRaw(isPopup bool) *menuStrip {
	if me.submenus != nil {
		if isPopup {
			panic("CreatePopup failed, menu already created.")
		} else {
			panic("CreateMain failed, menu already created.")
		}
	}
	me.submenus = make(map[string]win.HMENU) // init both maps
	me.items = make(map[string]menuItemMap)

	var hMenuRoot win.HMENU
	if isPopup {
		hMenuRoot = win.CreatePopupMenu()
	} else {
		hMenuRoot = win.CreateMenu()
	}

	rootStrip := &menuStrip{
		janitor: me,
		hMenu:   hMenuRoot,
	}
	me.submenus[""] = rootStrip.hMenu // root is empty textId
	return rootStrip
}

// Deletes the menu item and updates the internal cache.
func (me *Menu) DeleteItem(textId string) *Menu {
	item := me.items[textId]
	item.hMenu.DeleteMenu(item.cmdId, co.MF_BYCOMMAND)
	delete(me.items, textId)
	return me
}

// Calls DestroyMenu() on the root.
// Necessary if the menu wasn't attached to any window.
func (me *Menu) Destroy() {
	me.submenus[""].DestroyMenu()
}

// Calls EnableMenuItem() for many items at once.
func (me *Menu) EnableMany(doEnable bool, textIds []string) *Menu {
	flags := co.MF_BYCOMMAND
	if doEnable {
		flags |= co.MF_ENABLED
	} else {
		flags |= co.MF_GRAYED
	}
	for _, textId := range textIds {
		item := me.items[textId]
		item.hMenu.EnableMenuItem(item.cmdId, flags)
	}
	return me
}

// Returns the handle to the root submenu.
func (me *Menu) Hmenu() win.HMENU {
	return me.submenus[""]
}

// Returns the item for the given textId.
func (me *Menu) Item(textId string) *menuItem {
	storedItem := me.items[textId]
	return &menuItem{
		janitor: me,
		hMenu:   storedItem.hMenu,
		cmdId:   storedItem.cmdId,
	}
}

// Returns the submenu for the given textId.
func (me *Menu) Submenu(textId string) *menuStrip {
	return &menuStrip{
		janitor: me,
		hMenu:   me.submenus[textId],
	}
}

//------------------------------------------------------------------------------

type menuStrip struct {
	janitor *Menu
	hMenu   win.HMENU
}

// Appends a new item to the submenu, with an auto-generated command ID.
func (me *menuStrip) AddItem(textId, text string) *menuStrip {
	newCmdId := controlId{}
	return me.AddItemWithCmdId(textId, text, newCmdId.Id())
}

// Appends a new item to the submenu, with an specific command ID.
func (me *menuStrip) AddItemWithCmdId(textId, text string,
	cmdId int32) *menuStrip {

	newItemMap := menuItemMap{
		hMenu: me.hMenu,
		cmdId: cmdId,
	}
	me.janitor.items[textId] = newItemMap
	me.hMenu.AppendMenu(co.MF_STRING, uintptr(newItemMap.cmdId),
		uintptr(unsafe.Pointer(win.StrToPtr(text))))
	return me
}

// Adds a separator to this submenu. There are no high-level methods to retrieve
// separators, because they don't have an associated command ID.
func (me *menuStrip) AddSeparator() *menuStrip {
	me.hMenu.AppendMenu(co.MF_SEPARATOR, 0, 0) // separators are not kept in children map
	return me
}

// Appends a new submenu onto this submenu. Returns newly created submenu.
func (me *menuStrip) AddSubmenu(textId, text string) *menuStrip {
	newStrip := &menuStrip{
		janitor: me.janitor,
		hMenu:   win.CreatePopupMenu(),
	}
	me.janitor.submenus[textId] = newStrip.hMenu
	me.hMenu.AppendMenu(co.MF_STRING|co.MF_POPUP, uintptr(newStrip.hMenu),
		uintptr(unsafe.Pointer(win.StrToPtr(text))))
	return newStrip
}

// Returns the HMENU handle for this submenu.
func (me *menuStrip) Hmenu() win.HMENU {
	return me.hMenu
}

// Shows the popup menu anchored at the given coordinates.
// If hCoordsRelativeTo is zero, coordinates must be relative to hParent.
// This function will block until the menu disappears.
func (me *menuStrip) ShowAtPoint(pos win.POINT,
	hParent, hCoordsRelativeTo win.HWND) {

	if hCoordsRelativeTo == 0 {
		hCoordsRelativeTo = hParent
	}
	hCoordsRelativeTo.ClientToScreenPt(&pos) // now relative to screen
	hParent.SetForegroundWindow()
	me.hMenu.TrackPopupMenu(co.TPM_LEFTBUTTON, pos.X, pos.Y, hParent)
	hParent.PostMessage(co.WM_NULL, 0, 0) // necessary according to TrackMenuPopup docs
}

//------------------------------------------------------------------------------

type menuItem struct {
	janitor *Menu
	hMenu   win.HMENU
	cmdId   int32
}

// Returns the command ID for this item.
func (me *menuItem) CmdId() int32 {
	return me.cmdId
}

// Calls EnableMenuItem().
func (me *menuItem) Enable(doEnable bool) *menuItem {
	flags := co.MF_BYCOMMAND
	if doEnable {
		flags |= co.MF_ENABLED
	} else {
		flags |= co.MF_GRAYED
	}
	me.hMenu.EnableMenuItem(me.cmdId, flags)
	return me
}

// Returns the HMENU to which this item belongs.
func (me *menuItem) Hmenu() win.HMENU {
	return me.hMenu
}

// Sets the item text.
func (me *menuItem) SetText(text string) *menuItem {
	mii := win.MENUITEMINFO{
		FMask:      co.MIIM_STRING,
		DwTypeData: uintptr(unsafe.Pointer(win.StrToPtr(text))),
	}
	me.hMenu.SetMenuItemInfo(me.cmdId, false, &mii)
	return me
}

// Retrieves the item text.
func (me *menuItem) Text() string {
	mii := win.MENUITEMINFO{
		FMask: co.MIIM_STRING,
	}
	me.hMenu.GetMenuItemInfo(me.cmdId, false, &mii) // retrieve length
	mii.Cch++
	buf := make([]uint16, mii.Cch)
	mii.DwTypeData = uintptr(unsafe.Pointer(&buf[0])) // retrieve text
	me.hMenu.GetMenuItemInfo(me.cmdId, false, &mii)
	return syscall.UTF16ToString(buf)
}
