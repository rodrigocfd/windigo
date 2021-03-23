package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

// A handle to a menu.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hmenu
type HMENU HANDLE

// Appends a new item to the menu.
func (hMenu HMENU) AddItem(cmdId int, text string) {
	hMenu.AppendMenu(co.MF_STRING, uintptr(cmdId),
		unsafe.Pointer(Str.ToUint16Ptr(text)))
}

// Appends a new separator to the menu.
func (hMenu HMENU) AddSeparator() {
	hMenu.AppendMenu(co.MF_SEPARATOR, 0, nil)
}

// Appends a new submenu to the menu.
func (hMenu HMENU) AddSubmenu(text string, hSubMenu HMENU) {
	hMenu.AppendMenu(co.MF_STRING|co.MF_POPUP, uintptr(hSubMenu),
		unsafe.Pointer(Str.ToUint16Ptr(text)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-appendmenuw
func (hMenu HMENU) AppendMenu(
	uFlags co.MF, uIDNewItem uintptr, lpNewItem unsafe.Pointer) {

	ret, _, lerr := syscall.Syscall6(proc.AppendMenu.Addr(), 4,
		uintptr(hMenu), uintptr(uFlags), uIDNewItem, uintptr(lpNewItem),
		0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-checkmenuitem
func (hMenu HMENU) CheckMenuItem(idOrPos uint32, uCheck co.MF) co.MF {
	ret, _, lerr := syscall.Syscall(proc.CheckMenuItem.Addr(), 3,
		uintptr(hMenu), uintptr(idOrPos), uintptr(uCheck))
	if int(ret) == -1 {
		panic(err.ERROR(lerr))
	}
	return co.MF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-checkmenuradioitem
func (hMenu HMENU) CheckMenuRadioItem(
	firstIdOrPos, lastIdOrPos, checkedIdOrPos uint32, flags co.MF) {

	ret, _, lerr := syscall.Syscall6(proc.CheckMenuRadioItem.Addr(), 5,
		uintptr(hMenu), uintptr(firstIdOrPos), uintptr(lastIdOrPos),
		uintptr(checkedIdOrPos), uintptr(flags), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deletemenu
func (hMenu HMENU) DeleteMenu(idOrPos uint32, uFlags co.MF) {
	ret, _, lerr := syscall.Syscall(proc.DeleteMenu.Addr(), 3,
		uintptr(hMenu), uintptr(idOrPos), uintptr(uFlags))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroymenu
func (hMenu HMENU) DestroyMenu() {
	ret, _, lerr := syscall.Syscall(proc.DestroyMenu.Addr(), 1,
		uintptr(hMenu), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// Enables or disables many items at once, by command ID.
func (hMenu HMENU) EnableByCmdId(isEnabled bool, cmdIds ...int) {
	flags := co.MF_BYCOMMAND
	if isEnabled {
		flags |= co.MF_ENABLED
	} else {
		flags |= co.MF_GRAYED
	}

	for _, cmdId := range cmdIds {
		hMenu.EnableMenuItem(uint32(cmdId), flags)
	}
}

// Enables or disables many items at once, by zero-based position.
func (hMenu HMENU) EnableByPos(isEnabled bool, indexes ...int) {
	flags := co.MF_BYCOMMAND
	if isEnabled {
		flags |= co.MF_ENABLED
	} else {
		flags |= co.MF_GRAYED
	}

	for _, index := range indexes {
		hMenu.EnableMenuItem(uint32(index), flags)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enablemenuitem
func (hMenu HMENU) EnableMenuItem(uIDEnableItem uint32, uEnable co.MF) co.MF {
	ret, _, lerr := syscall.Syscall(proc.EnableMenuItem.Addr(), 3,
		uintptr(hMenu), uintptr(uIDEnableItem), uintptr(uEnable))
	if int(ret) == -1 {
		panic(err.ERROR(lerr))
	}
	return co.MF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuitemcount
func (hMenu HMENU) GetMenuItemCount() uint32 {
	ret, _, lerr := syscall.Syscall(proc.GetMenuItemCount.Addr(), 1,
		uintptr(hMenu), 0, 0)
	if int(ret) == -1 {
		panic(err.ERROR(lerr))
	}
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuitemid
func (hMenu HMENU) GetMenuItemID(nPos uint32) int32 {
	ret, _, _ := syscall.Syscall(proc.GetMenuItemID.Addr(), 2,
		uintptr(hMenu), uintptr(nPos), 0)
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsubmenu
func (hMenu HMENU) GetSubMenu(nPos uint32) (HMENU, bool) {
	ret, _, _ := syscall.Syscall(proc.GetSubMenu.Addr(), 2,
		uintptr(hMenu), uintptr(nPos), 0)
	if ret == 0 {
		return HMENU(0), false
	}
	return HMENU(ret), true
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insertmenuitemw
func (hMenu HMENU) InsertMenuItem(
	item uint32, fByPosition bool, lpmi *MENUITEMINFO) {

	ret, _, lerr := syscall.Syscall6(proc.InsertMenuItem.Addr(), 4,
		uintptr(hMenu), uintptr(item), util.BoolToUintptr(fByPosition),
		uintptr(unsafe.Pointer(lpmi)), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenudefaultitem
func (hMenu HMENU) SetMenuDefaultItem(uItem uint32, fByPos bool) {
	ret, _, lerr := syscall.Syscall(proc.SetMenuDefaultItem.Addr(), 3,
		uintptr(hMenu), uintptr(uItem), util.BoolToUintptr(fByPos))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// Shows the popup menu anchored at the given coordinates, using TrackPopupMenu().
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

// This function will block until the menu disappears.
// If TPM_RETURNCMD is passed, returns the selected command ID.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-trackpopupmenu
func (hMenu HMENU) TrackPopupMenu(uFlags co.TPM, x, y int32, hWnd HWND) int {
	ret, _, lerr := syscall.Syscall9(proc.TrackPopupMenu.Addr(), 7,
		uintptr(hMenu), uintptr(uFlags), uintptr(x), uintptr(y), 0, uintptr(hWnd),
		0, 0, 0)

	if (uFlags & co.TPM_RETURNCMD) != 0 {
		if ret == 0 && lerr != 0 {
			panic(err.ERROR(lerr))
		} else {
			return int(ret)
		}
	} else {
		if ret == 0 {
			panic(err.ERROR(lerr))
		} else {
			return 0
		}
	}
}
