package win

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// Appends a new item to the menu. Returns the same menu, so you can chain multiple calls.
//
// Wrapper to HMENU.AppendMenu().
func (hMenu HMENU) AddItem(cmdId int, text string) HMENU {
	hMenu.AppendMenu(co.MF_STRING, uint16(cmdId), text)
	return hMenu
}

// Appends a new separator to the menu. Returns the same menu, so you can chain multiple calls.
//
// Wrapper to HMENU.AppendMenu().
func (hMenu HMENU) AddSeparator() HMENU {
	hMenu.AppendMenu(co.MF_SEPARATOR, HMENU(0), LPARAM(0))
	return hMenu
}

// Appends a new submenu to the menu.
//
// Wrapper to HMENU.AppendMenu().
func (hMenu HMENU) AddSubmenu(text string, hSubMenu HMENU) {
	hMenu.AppendMenu(co.MF_STRING|co.MF_POPUP, hSubMenu, text)
}
