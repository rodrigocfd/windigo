/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"strings"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Manages an HMENU resource.
type Menu struct {
	text     string // the root will have an empty string
	hMenu    win.HMENU
	children []interface{} // Menu, menuItem, menuSeparator
}

// Inserts a new item with the given text.
// The command ID will be automatically generated.
func (me *Menu) AppendItem(text string) *Menu {
	newItem := &menuItem{text: text}
	me.hMenu.AppendMenu(co.MF_STRING, uintptr(newItem.Id()),
		uintptr(unsafe.Pointer(win.StrToPtr(text))))
	me.children = append(me.children, newItem)
	return me
}

// Inserts a new item with the given text, and a specific command ID.
func (me *Menu) AppendItemWithId(text string, cmdId int32) *Menu {
	newItem := &menuItem{
		text:  text,
		cmdId: controlId{id: cmdId},
	}
	me.hMenu.AppendMenu(co.MF_STRING, uintptr(cmdId),
		uintptr(unsafe.Pointer(win.StrToPtr(text))))
	me.children = append(me.children, newItem)
	return me
}

// Inserts a separator.
func (me *Menu) AppendSeparator() *Menu {
	me.hMenu.AppendMenu(co.MF_SEPARATOR, 0, 0)
	me.children = append(me.children, &menuSeparator{})
	return me
}

// Inserts a new submenu with the given text.
// Returns newly inserted submenu.
func (me *Menu) AppendSubmenu(text string) *Menu {
	newSubmeu := &Menu{
		text:  text,
		hMenu: win.CreatePopupMenu(),
	}
	me.hMenu.AppendMenu(co.MF_STRING|co.MF_POPUP, uintptr(newSubmeu.hMenu),
		uintptr(unsafe.Pointer(win.StrToPtr(text))))
	me.children = append(me.children, newSubmeu)
	return newSubmeu
}

// Creates an horizontal main window menu.
// If not attached to a window, must be destroyed.
func (me *Menu) CreateMain() *Menu {
	if me.hMenu != 0 {
		panic("CreateMain failed: can't create menu twice.")
	}
	me.hMenu = win.CreateMenu()
	return me
}

// Creates a vertical floating popup menu.
// If not attached to a window, must be destroyed.
func (me *Menu) CreatePopup() *Menu {
	if me.hMenu != 0 {
		panic("CreatePopup failed: can't create menu twice.")
	}
	me.hMenu = win.CreatePopupMenu()
	return me
}

func (me *Menu) Destroy() {
	if me.hMenu != 0 {
		me.hMenu.DestroyMenu()
	}
}

func (me *Menu) Hmenu() win.HMENU {
	return me.hMenu
}

// Returns the command ID of the given item hierarchy, pipe-separated.
// Example: IdOf("&File|&Open").
func (me *Menu) IdOf(path string) int32 {
	entries := strings.Split(path, "|")
	lastSubmenu := me
	for i := range entries {
		if i == len(entries)-1 {
			return lastSubmenu.Item(entries[i]).Id()
		}
		lastSubmenu = lastSubmenu.Submenu(entries[i])
	}
	return -1 // not found
}

// Returns the first item with the given string, or nil.
// Remember the accelerator ampersands in the string, if any.
func (me *Menu) Item(text string) *menuItem {
	for _, child := range me.children {
		if item, ok := child.(*menuItem); ok && item.text == text {
			return item
		}
	}
	return nil
}

// Returns the first submenu with the given string, or nil.
// Remember the accelerator ampersands in the string, if any.
func (me *Menu) Submenu(text string) *Menu {
	for _, child := range me.children {
		if submenu, ok := child.(*Menu); ok && submenu.text == text {
			return submenu
		}
	}
	return nil
}

//------------------------------------------------------------------------------

type menuItem struct {
	cmdId controlId
	text  string
}

func (me *menuItem) Id() int32 {
	return me.cmdId.Id()
}

//------------------------------------------------------------------------------

type menuSeparator struct {
}
