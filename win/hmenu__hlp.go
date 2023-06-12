//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// This helper method appends a new item to the menu, by calling
// HMENU.AppendMenu(). Returns the same menu, so you can chain multiple calls.
func (hMenu HMENU) AddItem(cmdId int, text string) HMENU {
	hMenu.AppendMenu(co.MF_STRING, uint16(cmdId), text)
	return hMenu
}

// This helper method appends a new separator to the menu, by calling
// HMENU.AppendMenu(). Returns the same menu, so you can chain multiple calls.
func (hMenu HMENU) AddSeparator() HMENU {
	hMenu.AppendMenu(co.MF_SEPARATOR, HMENU(0), LPARAM(0))
	return hMenu
}

// This helper method appends a new submenu to the menu, by calling
// HMENU.AppendMenu().
func (hMenu HMENU) AddSubmenu(text string, hSubMenu HMENU) {
	hMenu.AppendMenu(co.MF_STRING|co.MF_POPUP, hSubMenu, text)
}

// This helper method shows the popup menu anchored at the given coordinates,
// using HMENU.TrackPopupMenu().
//
// If hCoordsRelativeTo is zero, coordinates must be relative to hParent.
//
// This function will block until the menu disappears.
func (hMenu HMENU) ShowAtPoint(pos POINT, hParent, hCoordsRelativeTo HWND) {
	if hCoordsRelativeTo == 0 {
		hCoordsRelativeTo = hParent
	}

	hCoordsRelativeTo.ClientToScreenPt(&pos) // now relative to screen
	hParent.SetForegroundWindow()
	hMenu.TrackPopupMenu(co.TPM_LEFTBUTTON, pos.X, pos.Y, hParent)
	hParent.PostMessage(co.WM_NULL, 0, 0) // necessary according to TrackMenuPopup docs
}
