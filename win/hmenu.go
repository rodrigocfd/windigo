package win

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a menu.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hmenu
type HMENU HANDLE

// ⚠️ You must defer HMENU.DestroyMenu(), unless it's attached to a window.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createmenu
func CreateMenu() HMENU {
	ret, _, err := syscall.Syscall(proc.CreateMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HMENU(ret)
}

// ⚠️ You must defer HMENU.DestroyMenu(), unless it's attached to a window.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createpopupmenu
func CreatePopupMenu() HMENU {
	ret, _, err := syscall.Syscall(proc.CreatePopupMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HMENU(ret)
}

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

// This function is rather tricky. Prefer using HMENU.AddItem(),
// HMENU.AddSeparator() or HMENU.AddSubmenu().
//
// ⚠️ uIDNewItem must be uint16 or HMENU.
//
// ⚠️ lpNewItem must be HBITMAP, LPARAM or string.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-appendmenuw
func (hMenu HMENU) AppendMenu(
	uFlags co.MF, uIDNewItem interface{}, lpNewItem interface{}) {

	pId := uintptr(0)
	switch v := uIDNewItem.(type) {
	case uint16:
		pId = uintptr(v)
	case HMENU:
		pId = uintptr(v)
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(uIDNewItem)))
	}

	var pItem unsafe.Pointer
	switch v := lpNewItem.(type) {
	case HBITMAP:
		pItem = unsafe.Pointer(uintptr(v))
	case LPARAM:
		pItem = unsafe.Pointer(uintptr(v))
	case string:
		pItem = unsafe.Pointer(Str.ToNativePtr(v))
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(lpNewItem)))
	}

	ret, _, err := syscall.Syscall6(proc.AppendMenu.Addr(), 4,
		uintptr(hMenu), uintptr(uFlags), pId, uintptr(pItem),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-checkmenuitem
func (hMenu HMENU) CheckMenuItem(idOrPos uint32, uCheck co.MF) co.MF {
	ret, _, err := syscall.Syscall(proc.CheckMenuItem.Addr(), 3,
		uintptr(hMenu), uintptr(idOrPos), uintptr(uCheck))
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return co.MF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-checkmenuradioitem
func (hMenu HMENU) CheckMenuRadioItem(
	firstIdOrPos, lastIdOrPos, checkedIdOrPos uint32, flags co.MF) {

	ret, _, err := syscall.Syscall6(proc.CheckMenuRadioItem.Addr(), 5,
		uintptr(hMenu), uintptr(firstIdOrPos), uintptr(lastIdOrPos),
		uintptr(checkedIdOrPos), uintptr(flags), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deletemenu
func (hMenu HMENU) DeleteMenu(idOrPos uint32, flags co.MF) {
	ret, _, err := syscall.Syscall(proc.DeleteMenu.Addr(), 3,
		uintptr(hMenu), uintptr(idOrPos), uintptr(flags))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroymenu
func (hMenu HMENU) DestroyMenu() {
	ret, _, err := syscall.Syscall(proc.DestroyMenu.Addr(), 1,
		uintptr(hMenu), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
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
func (hMenu HMENU) EnableMenuItem(idOrPos uint32, enable co.MF) co.MF {
	ret, _, err := syscall.Syscall(proc.EnableMenuItem.Addr(), 3,
		uintptr(hMenu), uintptr(idOrPos), uintptr(enable))
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return co.MF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuitemcount
func (hMenu HMENU) GetMenuItemCount() uint32 {
	ret, _, err := syscall.Syscall(proc.GetMenuItemCount.Addr(), 1,
		uintptr(hMenu), 0, 0)
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuitemid
func (hMenu HMENU) GetMenuItemID(pos uint32) int32 {
	ret, _, _ := syscall.Syscall(proc.GetMenuItemID.Addr(), 2,
		uintptr(hMenu), uintptr(pos), 0)
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuiteminfow
func (hMenu HMENU) GetMenuItemInfo(
	idOrPos uint32, byPosition bool, mii *MENUITEMINFO) {

	ret, _, err := syscall.Syscall6(proc.GetMenuItemInfo.Addr(), 4,
		uintptr(hMenu), uintptr(idOrPos), util.BoolToUintptr(byPosition),
		uintptr(unsafe.Pointer(mii)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsubmenu
func (hMenu HMENU) GetSubMenu(pos uint32) (HMENU, bool) {
	ret, _, _ := syscall.Syscall(proc.GetSubMenu.Addr(), 2,
		uintptr(hMenu), uintptr(pos), 0)
	hSub := HMENU(ret)
	return hSub, hSub != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insertmenuitemw
func (hMenu HMENU) InsertMenuItem(
	idOrPos uint32, byPosition bool, mii *MENUITEMINFO) {

	ret, _, err := syscall.Syscall6(proc.InsertMenuItem.Addr(), 4,
		uintptr(hMenu), uintptr(idOrPos), util.BoolToUintptr(byPosition),
		uintptr(unsafe.Pointer(mii)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenudefaultitem
func (hMenu HMENU) SetMenuDefaultItem(idOrPos uint32, byPosition bool) {
	ret, _, err := syscall.Syscall(proc.SetMenuDefaultItem.Addr(), 3,
		uintptr(hMenu), uintptr(idOrPos), util.BoolToUintptr(byPosition))
	if ret == 0 {
		panic(errco.ERROR(err))
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
func (hMenu HMENU) TrackPopupMenu(flags co.TPM, x, y int32, hWnd HWND) int {
	ret, _, err := syscall.Syscall9(proc.TrackPopupMenu.Addr(), 7,
		uintptr(hMenu), uintptr(flags), uintptr(x), uintptr(y), 0, uintptr(hWnd),
		0, 0, 0)

	if (flags & co.TPM_RETURNCMD) != 0 {
		if ret == 0 && err != 0 {
			panic(errco.ERROR(err))
		} else {
			return int(ret)
		}
	} else {
		if ret == 0 {
			panic(errco.ERROR(err))
		} else {
			return 0
		}
	}
}
