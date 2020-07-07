/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
	"strings"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Manages an HMENU resource.
type Menu struct {
	hMenu    win.HMENU
	children map[string]interface{} // Menu or menuItem
}

// Inserts a new item with the given text.
// The command ID will be automatically generated.
func (me *Menu) AppendItem(text string) *Menu {
	newItem := &menuItem{}
	me.hMenu.AppendMenu(co.MF_STRING, uintptr(newItem.Id()),
		uintptr(unsafe.Pointer(win.StrToPtr(text))))
	me.children[removeAccelAmpersands(text)] = newItem // text kept without ampersands
	return me
}

// Inserts a new item with the given text, and a specific command ID.
func (me *Menu) AppendItemWithId(text string, cmdId int32) *Menu {
	newItem := &menuItem{
		cmdId: controlId{id: cmdId},
	}
	me.hMenu.AppendMenu(co.MF_STRING, uintptr(cmdId),
		uintptr(unsafe.Pointer(win.StrToPtr(text))))
	me.children[removeAccelAmpersands(text)] = newItem // text kept without ampersands
	return me
}

// Inserts a separator.
func (me *Menu) AppendSeparator() *Menu {
	me.hMenu.AppendMenu(co.MF_SEPARATOR, 0, 0) // separators are not kept in children map
	return me
}

// Inserts a new submenu with the given text.
// Returns newly inserted submenu.
func (me *Menu) AppendSubmenu(text string) *Menu {
	newSubmenu := &Menu{
		hMenu:    win.CreatePopupMenu(),
		children: make(map[string]interface{}),
	}
	me.hMenu.AppendMenu(co.MF_STRING|co.MF_POPUP, uintptr(newSubmenu.hMenu),
		uintptr(unsafe.Pointer(win.StrToPtr(text))))
	me.children[removeAccelAmpersands(text)] = newSubmenu // text kept without ampersands
	return newSubmenu
}

// Creates an horizontal main window menu.
// If not attached to a window, must be destroyed.
func (me *Menu) CreateMain() *Menu {
	if me.hMenu != 0 {
		panic("CreateMain failed: can't create menu twice.")
	}
	me.hMenu = win.CreateMenu()
	me.children = make(map[string]interface{})
	return me
}

// Creates a vertical floating popup menu.
// If not attached to a window, must be destroyed.
func (me *Menu) CreatePopup() *Menu {
	if me.hMenu != 0 {
		panic("CreatePopup failed: can't create menu twice.")
	}
	me.hMenu = win.CreatePopupMenu()
	me.children = make(map[string]interface{})
	return me
}

// Recursively destroys all submenus.
// Call only if the menu wasn't attached to a window.
func (me *Menu) Destroy() {
	if me.hMenu != 0 {
		me.hMenu.DestroyMenu()
	}
}

func (me *Menu) Hmenu() win.HMENU {
	return me.hMenu
}

// Returns the command ID of the given item hierarchy, pipe-separated.
// Don't pass accelerator ampersands in the string.
// Example: IdOf("File|Open").
func (me *Menu) IdOf(path string) int32 {
	pathEntries := strings.Split(path, "|")
	lastSubmenu := me
	for i := range pathEntries {
		if i == len(pathEntries)-1 {
			return lastSubmenu.Item(pathEntries[i]).Id()
		}
		lastSubmenu = lastSubmenu.Submenu(pathEntries[i])
	}
	return -1 // not found
}

// Returns the first item with the given string, or panics.
// Don't pass accelerator ampersands in the string.
func (me *Menu) Item(text string) *menuItem {
	if value, ok := me.children[text]; ok { // key exists?
		if item, ok := value.(*menuItem); ok { // key holds a menuItem?
			return item
		}
	}
	panic(fmt.Sprintf("Inexistent menu item text: %s", text))
}

// Returns the first submenu with the given string, or panics.
// Don't pass accelerator ampersands in the string.
func (me *Menu) Submenu(text string) *Menu {
	if value, ok := me.children[text]; ok { // key exists?
		if submenu, ok := value.(*Menu); ok { // key holds a Menu?
			return submenu
		}
	}
	panic(fmt.Sprintf("Inexistent submenu text: %s", text))
}

//------------------------------------------------------------------------------

type menuItem struct {
	cmdId controlId
}

func (me *menuItem) Id() int32 {
	return me.cmdId.Id()
}
