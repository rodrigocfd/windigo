//go:build windows

package win

import (
	"fmt"
	"reflect"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [menu].
//
// [menu]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hmenu
type HMENU HANDLE

// [CreateMenu] function.
//
// ⚠️ You must defer HMENU.DestroyMenu(), unless it's attached to a window.
//
// [CreateMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createmenu
func CreateMenu() HMENU {
	ret, _, err := syscall.SyscallN(proc.CreateMenu.Addr())
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HMENU(ret)
}

// [CreatePopupMenu] function.
//
// ⚠️ You must defer HMENU.DestroyMenu(), unless it's attached to a window.
//
// [CreatePopupMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createpopupmenu
func CreatePopupMenu() HMENU {
	ret, _, err := syscall.SyscallN(proc.CreatePopupMenu.Addr())
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HMENU(ret)
}

// [AppendMenu] function.
//
// This function is rather tricky. Prefer using HMENU.AddItem(),
// HMENU.AddSeparator() or HMENU.AddSubmenu().
//
// ⚠️ uIDNewItem must be uint16 or HMENU.
//
// ⚠️ lpNewItem must be HBITMAP, LPARAM or string.
//
// [AppendMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-appendmenuw
func (hMenu HMENU) AppendMenu(
	uFlags co.MF, uIDNewItem interface{}, lpNewItem interface{}) {

	var pId uintptr
	switch v := uIDNewItem.(type) {
	case uint16:
		pId = uintptr(v)
	case HMENU:
		pId = uintptr(v)
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(uIDNewItem)))
	}

	var pItem uintptr
	var pLpNewItem *uint16
	switch v := lpNewItem.(type) {
	case HBITMAP:
		pItem = uintptr(v)
	case LPARAM:
		pItem = uintptr(v)
	case string:
		pLpNewItem = Str.ToNativePtr(v) // keep the buffer
		pItem = uintptr(unsafe.Pointer(pLpNewItem))
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(lpNewItem)))
	}

	ret, _, err := syscall.SyscallN(proc.AppendMenu.Addr(),
		uintptr(hMenu), uintptr(uFlags), pId, pItem)
	runtime.KeepAlive(pLpNewItem)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [CheckMenuItem] function.
//
// # Example:
//
//	var hMenu win.HMENU // initialized somewhere
//
//	hMenu.CheckMenuItem(win.MenuItemPos(0), true)
//
// [CheckMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-checkmenuitem
func (hMenu HMENU) CheckMenuItem(item MenuItem, check bool) bool {
	idPos, mf := item.raw()
	flags := util.Iif(check, co.MF_CHECKED, co.MF_UNCHECKED).(co.MF) | mf

	ret, _, err := syscall.SyscallN(proc.CheckMenuItem.Addr(),
		uintptr(hMenu), idPos, uintptr(flags))
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return co.MF(ret) == co.MF_CHECKED
}

// [CheckMenuRadioItem] function.
//
// Panics if the three item identifiers don't have the same variant type.
//
// # Example:
//
//	var hMenu win.HMENU // initialized somewhere
//
//	p.Hmenu().CheckMenuRadioItem(
//		win.MenuItemPos(0), win.MenuItemPos(4), win.MenuItemPos(1))
//
// [CheckMenuRadioItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-checkmenuradioitem
func (hMenu HMENU) CheckMenuRadioItem(
	firstItem, lastItem, checkedItem MenuItem) {

	idPosFirst, mfFirst := firstItem.raw()
	idPosLast, mfLast := lastItem.raw()
	idPosChecked, mfChecked := checkedItem.raw()

	if mfFirst != mfLast {
		panic("firstItem and lastItem have different variant types.")
	} else if mfFirst != mfChecked {
		panic("firstItem and checkedItem have different variant types.")
	}

	ret, _, err := syscall.SyscallN(proc.CheckMenuRadioItem.Addr(),
		uintptr(hMenu), idPosFirst, idPosLast, idPosChecked,
		uintptr(mfFirst))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [DeleteMenu] function.
//
// # Example:
//
//	var hMenu win.HMENU // initialized somewhere
//
//	hMenu.DeleteMenu(win.MenuItemPos(3))
//
// [DeleteMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deletemenu
func (hMenu HMENU) DeleteMenu(item MenuItem) {
	idPos, mf := item.raw()
	ret, _, err := syscall.SyscallN(proc.DeleteMenu.Addr(),
		uintptr(hMenu), idPos, uintptr(mf))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [DestroyMenu] function.
//
// [DestroyMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroymenu
func (hMenu HMENU) DestroyMenu() error {
	ret, _, err := syscall.SyscallN(proc.DestroyMenu.Addr(),
		uintptr(hMenu))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [EnableMenuItem] function.
//
// # Example:
//
//	var hMenu win.HMENU // initialized somewhere
//
//	hMenu.EnableMenuItem(win.MenuItemPos(0), false)
//
// [EnableMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enablemenuitem
func (hMenu HMENU) EnableMenuItem(item MenuItem, enable bool) bool {
	idPos, mf := item.raw()
	flags := util.Iif(enable, co.MF_ENABLED, co.MF_DISABLED).(co.MF) | mf

	ret, _, err := syscall.SyscallN(proc.EnableMenuItem.Addr(),
		uintptr(hMenu), idPos, uintptr(flags))
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return co.MF(ret) == co.MF_CHECKED
}

// [GetMenuDefaultItem] function.
//
// [GetMenuDefaultItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenudefaultitem
func (hMenu HMENU) GetMenuDefaultItem(gmdiFlags co.GMDI) (pos MenuItem) {
	ret, _, err := syscall.SyscallN(proc.GetMenuDefaultItem.Addr(),
		uintptr(hMenu), 1, uintptr(gmdiFlags))
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return MenuItemPos(int(ret))
}

// [GetMenuItemCount] function.
//
// [GetMenuItemCount]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuitemcount
func (hMenu HMENU) GetMenuItemCount() uint32 {
	ret, _, err := syscall.SyscallN(proc.GetMenuItemCount.Addr(),
		uintptr(hMenu))
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// [GetMenuItemID] function.
//
// [GetMenuItemID]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuitemid
func (hMenu HMENU) GetMenuItemID(pos uint32) int32 {
	ret, _, _ := syscall.SyscallN(proc.GetMenuItemID.Addr(),
		uintptr(hMenu), uintptr(pos))
	return int32(ret)
}

// [GetMenuItemInfo] function.
//
// [GetMenuItemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuiteminfow
func (hMenu HMENU) GetMenuItemInfo(item MenuItem, mii *MENUITEMINFO) {
	idPos, mf := item.raw()
	ret, _, err := syscall.SyscallN(proc.GetMenuItemInfo.Addr(),
		uintptr(hMenu), idPos, util.BoolToUintptr(mf == co.MF_BYPOSITION),
		uintptr(unsafe.Pointer(mii)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [GetSubMenu] function.
//
// [GetSubMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsubmenu
func (hMenu HMENU) GetSubMenu(pos uint32) (HMENU, bool) {
	ret, _, _ := syscall.SyscallN(proc.GetSubMenu.Addr(),
		uintptr(hMenu), uintptr(pos))
	hSub := HMENU(ret)
	return hSub, hSub != 0
}

// [InsertMenuItem] function.
//
// [InsertMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insertmenuitemw
func (hMenu HMENU) InsertMenuItem(itemBefore MenuItem, mii *MENUITEMINFO) {
	idPos, mf := itemBefore.raw()
	ret, _, err := syscall.SyscallN(proc.InsertMenuItem.Addr(),
		uintptr(hMenu), idPos, util.BoolToUintptr(mf == co.MF_BYPOSITION),
		uintptr(unsafe.Pointer(mii)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [RemoveMenu] function.
//
// [RemoveMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-removemenu
func (hMenu HMENU) RemoveMenu(item MenuItem) {
	idPos, mf := item.raw()
	ret, _, err := syscall.SyscallN(proc.RemoveMenu.Addr(),
		uintptr(hMenu), idPos, uintptr(mf))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SetMenuDefaultItem] function.
//
// [SetMenuDefaultItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenudefaultitem
func (hMenu HMENU) SetMenuDefaultItem(item MenuItem) {
	idPos, mf := item.raw()
	ret, _, err := syscall.SyscallN(proc.SetMenuDefaultItem.Addr(),
		uintptr(hMenu), idPos, util.BoolToUintptr(mf == co.MF_BYPOSITION))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SetMenuInfo] function.
//
// [SetMenuInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuinfo
func (hMenu HMENU) SetMenuInfo(info *MENUINFO) {
	ret, _, err := syscall.SyscallN(proc.SetMenuInfo.Addr(),
		uintptr(hMenu), uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SetMenuItemBitmaps] function.
//
// [SetMenuItemBitmaps]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuitembitmaps
func (hMenu HMENU) SetMenuItemBitmaps(
	item MenuItem, hBmpUnchecked, hBmpChecked HBITMAP) {

	idPos, mf := item.raw()
	ret, _, err := syscall.SyscallN(proc.SetMenuItemBitmaps.Addr(),
		uintptr(hMenu), idPos, uintptr(mf),
		uintptr(hBmpUnchecked), uintptr(hBmpChecked))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SetMenuItemInfo] function.
//
// [SetMenuItemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuiteminfow
func (hMenu HMENU) SetMenuItemInfo(item MenuItem, info *MENUITEMINFO) {
	info.SetCbSize() // safety
	idPos, mf := item.raw()

	ret, _, err := syscall.SyscallN(proc.SetMenuItemInfo.Addr(),
		uintptr(hMenu), idPos, util.BoolToUintptr(mf == co.MF_BYPOSITION),
		uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [TrackPopupMenu] function.
//
// This function will block until the menu disappears.
// If TPM_RETURNCMD is passed, returns the selected command ID.
//
// [TrackPopupMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-trackpopupmenu
func (hMenu HMENU) TrackPopupMenu(flags co.TPM, x, y int32, hWnd HWND) int {
	ret, _, err := syscall.SyscallN(proc.TrackPopupMenu.Addr(),
		uintptr(hMenu), uintptr(flags), uintptr(x), uintptr(y),
		0, uintptr(hWnd), 0)

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
