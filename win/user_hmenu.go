//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to a [menu].
//
// [menu]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hmenu
type HMENU HANDLE

// [CreateMenu] function.
//
// ⚠️ You must defer HMENU.DestroyMenu(), unless it's attached to a window.
//
// [CreateMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createmenu
func CreateMenu() (HMENU, error) {
	ret, _, err := syscall.SyscallN(_CreateMenu.Addr())
	if ret == 0 {
		return HMENU(0), co.ERROR(err)
	}
	return HMENU(ret), nil
}

var _CreateMenu = dll.User32.NewProc("CreateMenu")

// [CreatePopupMenu] function.
//
// ⚠️ You must defer HMENU.DestroyMenu(), unless it's attached to a window.
//
// [CreatePopupMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createpopupmenu
func CreatePopupMenu() (HMENU, error) {
	ret, _, err := syscall.SyscallN(_CreatePopupMenu.Addr())
	if ret == 0 {
		return HMENU(0), co.ERROR(err)
	}
	return HMENU(ret), nil
}

var _CreatePopupMenu = dll.User32.NewProc("CreatePopupMenu")

// [CheckMenuItem] function for multiple items, using the item command ID.
//
// [CheckMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-checkmenuitem
func (hMenu HMENU) CheckMenuItemByCmd(check bool, cmdIds ...uint16) error {
	for _, cmdId := range cmdIds {
		if err := hMenu.checkMenuItem(check, co.MF_BYCOMMAND, uint(cmdId)); err != nil {
			return err
		}
	}
	return nil
}

// [CheckMenuItem] function for multiple items, using the zero-based item
// position.
//
// [CheckMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-checkmenuitem
func (hMenu HMENU) CheckMenuItemByPos(check bool, indexes ...uint) error {
	for _, index := range indexes {
		if err := hMenu.checkMenuItem(check, co.MF_BYPOSITION, index); err != nil {
			return err
		}
	}
	return nil
}

func (hMenu HMENU) checkMenuItem(check bool, flagPosCmd co.MF, item uint) error {
	if check {
		flagPosCmd |= co.MF_CHECKED
	} else {
		flagPosCmd |= co.MF_UNCHECKED
	}

	ret, _, _ := syscall.SyscallN(_CheckMenuItem.Addr(),
		uintptr(hMenu), uintptr(item), uintptr(flagPosCmd))
	if int(ret) == -1 {
		return co.ERROR_INVALID_PARAMETER
	}
	return nil
}

var _CheckMenuItem = dll.User32.NewProc("CheckMenuItem")

// [DeleteMenu] function for multiple items, using the item command ID.
//
// [DeleteMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deletemenu
func (hMenu HMENU) DeleteMenuByCmd(cmdIds ...uint16) error {
	for _, cmdId := range cmdIds {
		ret, _, err := syscall.SyscallN(_DeleteMenu.Addr(),
			uintptr(hMenu), uintptr(cmdId), uintptr(co.MF_BYCOMMAND))
		if ret == 0 {
			return co.ERROR(err)
		}
	}
	return nil
}

// [DeleteMenu] function for multiple items, using the zero-based item position.
//
// [DeleteMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deletemenu
func (hMenu HMENU) DeleteMenuByPos(indexes ...uint) error {
	for _, index := range indexes {
		ret, _, err := syscall.SyscallN(_DeleteMenu.Addr(),
			uintptr(hMenu), uintptr(index), uintptr(co.MF_BYPOSITION))
		if ret == 0 {
			return co.ERROR(err)
		}
	}
	return nil
}

var _DeleteMenu = dll.User32.NewProc("DeleteMenu")

// [DestroyMenu] function.
//
// [DestroyMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroymenu
func (hMenu HMENU) DestroyMenu() error {
	ret, _, err := syscall.SyscallN(_DestroyMenu.Addr(),
		uintptr(hMenu))
	return util.ZeroToGetLastError(ret, err)
}

var _DestroyMenu = dll.User32.NewProc("DestroyMenu")

// [EnableMenuItem] function for multiple items, using the item command ID
//
// Panics if cmdIds is empty.
//
// [EnableMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enablemenuitem
func (hMenu HMENU) EnableMenuItemByCmd(enable bool, cmdIds ...uint16) error {
	if len(cmdIds) == 0 {
		panic("No cmdId for EnableMenuItemByCmd.")
	}
	for _, cmdId := range cmdIds {
		if err := hMenu.enableMenuItem(enable, co.MF_BYCOMMAND, uint(cmdId)); err != nil {
			return err
		}
	}
	return nil
}

// [EnableMenuItem] function for multiple items, using the zero-based item
// position.
//
// Panics if indexes is empty.
//
// [EnableMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enablemenuitem
func (hMenu HMENU) EnableMenuItemByPos(enable bool, indexes ...uint) error {
	if len(indexes) == 0 {
		panic("No index for EnableMenuItemByPos.")
	}
	for _, index := range indexes {
		if err := hMenu.enableMenuItem(enable, co.MF_BYPOSITION, index); err != nil {
			return err
		}
	}
	return nil
}

func (hMenu HMENU) enableMenuItem(enable bool, flagPosCmd co.MF, item uint) error {
	if enable {
		flagPosCmd |= co.MF_ENABLED
	} else {
		flagPosCmd |= co.MF_DISABLED
	}

	ret, _, _ := syscall.SyscallN(_EnableMenuItem.Addr(),
		uintptr(hMenu), uintptr(item), uintptr(flagPosCmd))
	if int(ret) == -1 {
		return co.ERROR_INVALID_PARAMETER
	}
	return nil
}

var _EnableMenuItem = dll.User32.NewProc("EnableMenuItem")

// [GetMenuDefaultItem] function.
//
// Returns the zero-based index of the item.
//
// [GetMenuDefaultItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenudefaultitem
func (hMenu HMENU) GetMenuDefaultItem(gmdiFlags co.GMDI) (uint, error) {
	ret, _, err := syscall.SyscallN(_GetMenuDefaultItem.Addr(),
		uintptr(hMenu), 1, uintptr(gmdiFlags))
	if wErr := co.ERROR(err); int(ret) == -1 || wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint(ret), nil
}

var _GetMenuDefaultItem = dll.User32.NewProc("GetMenuDefaultItem")

// [GetMenuItemID] function.
//
// Given the zero-based index, returns its command ID.
//
// [GetMenuItemID]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuitemid
func (hMenu HMENU) GetMenuItemID(index uint) (uint16, error) {
	ret, _, _ := syscall.SyscallN(_GetMenuItemID.Addr(),
		uintptr(hMenu), uintptr(index))
	if int(ret) == -1 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return uint16(ret), nil
}

var _GetMenuItemID = dll.User32.NewProc("GetMenuItemID")

// [GetMenuItemCount] function.
//
// [GetMenuItemCount]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuitemcount
func (hMenu HMENU) GetMenuItemCount() (uint, error) {
	ret, _, err := syscall.SyscallN(_GetMenuItemCount.Addr(),
		uintptr(hMenu))
	if int(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return uint(ret), nil
}

var _GetMenuItemCount = dll.User32.NewProc("GetMenuItemCount")

// [GetMenuItemInfo] function, using the item command ID.
//
// [GetMenuItemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuiteminfow
func (hMenu HMENU) GetMenuItemInfoByCmd(cmdId uint16, mii *MENUITEMINFO) error {
	mii.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(_GetMenuItemInfoW.Addr(),
		uintptr(hMenu), uintptr(cmdId), uintptr(co.MF_BYCOMMAND),
		uintptr(unsafe.Pointer(mii)))
	return util.ZeroToGetLastError(ret, err)
}

// [GetMenuItemInfo] function, using the zero-based item position.
//
// [GetMenuItemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuiteminfow
func (hMenu HMENU) GetMenuItemInfoByPos(index uint, mii *MENUITEMINFO) error {
	mii.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(_GetMenuItemInfoW.Addr(),
		uintptr(hMenu), uintptr(index), uintptr(co.MF_BYPOSITION),
		uintptr(unsafe.Pointer(mii)))
	return util.ZeroToGetLastError(ret, err)
}

var _GetMenuItemInfoW = dll.User32.NewProc("GetMenuItemInfoW")

// [GetSubMenu] function.
//
// [GetSubMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsubmenu
func (hMenu HMENU) GetSubMenu(index uint) (HMENU, bool) {
	ret, _, _ := syscall.SyscallN(_GetSubMenu.Addr(),
		uintptr(hMenu), uintptr(index))
	hSub := HMENU(ret)
	return hSub, hSub != 0
}

var _GetSubMenu = dll.User32.NewProc("GetSubMenu")

// [InsertMenuItem] function, using the item command ID.
//
// [InsertMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insertmenuitemw
func (hMenu HMENU) InsertMenuItemByCmd(cmdId uint16, mii *MENUITEMINFO) error {
	ret, _, err := syscall.SyscallN(_InsertMenuItemW.Addr(),
		uintptr(hMenu), uintptr(cmdId), uintptr(co.MF_BYCOMMAND),
		uintptr(unsafe.Pointer(mii)))
	return util.ZeroToGetLastError(ret, err)
}

// [InsertMenuItem] function, using the zero-based item position.
//
// [InsertMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insertmenuitemw
func (hMenu HMENU) InsertMenuItemByPos(index uint, mii *MENUITEMINFO) error {
	ret, _, err := syscall.SyscallN(_InsertMenuItemW.Addr(),
		uintptr(hMenu), uintptr(index), uintptr(co.MF_BYPOSITION),
		uintptr(unsafe.Pointer(mii)))
	return util.ZeroToGetLastError(ret, err)
}

var _InsertMenuItemW = dll.User32.NewProc("InsertMenuItemW")

// [RemoveMenu] function for multiple items, using the item command ID.
//
// [RemoveMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-removemenu
func (hMenu HMENU) RemoveMenuByCmd(cmdIds ...uint16) error {
	for _, cmdId := range cmdIds {
		ret, _, err := syscall.SyscallN(_RemoveMenu.Addr(),
			uintptr(hMenu), uintptr(cmdId), uintptr(co.MF_BYCOMMAND))
		if ret == 0 {
			return co.ERROR(err)
		}
	}
	return nil
}

// [RemoveMenu] function for multiple items, using the zero-based item position.
//
// [RemoveMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-removemenu
func (hMenu HMENU) RemoveMenuByPos(indexes ...uint) error {
	for _, index := range indexes {
		ret, _, err := syscall.SyscallN(_RemoveMenu.Addr(),
			uintptr(hMenu), uintptr(index), uintptr(co.MF_BYPOSITION))
		if ret == 0 {
			return co.ERROR(err)
		}
	}
	return nil
}

var _RemoveMenu = dll.User32.NewProc("RemoveMenu")

// [SetMenuDefaultItem] function, using the item command ID.
//
// [SetMenuDefaultItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenudefaultitem
func (hMenu HMENU) SetMenuDefaultItemByCmd(cmdId uint16) error {
	ret, _, err := syscall.SyscallN(_SetMenuDefaultItem.Addr(),
		uintptr(hMenu), uintptr(cmdId), uintptr(co.MF_BYCOMMAND))
	return util.ZeroToGetLastError(ret, err)
}

// [SetMenuDefaultItem] function, using the zero-based item position.
//
// [SetMenuDefaultItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenudefaultitem
func (hMenu HMENU) SetMenuDefaultItemByPos(index uint) error {
	ret, _, err := syscall.SyscallN(_SetMenuDefaultItem.Addr(),
		uintptr(hMenu), uintptr(index), uintptr(co.MF_BYPOSITION))
	return util.ZeroToGetLastError(ret, err)
}

var _SetMenuDefaultItem = dll.User32.NewProc("SetMenuDefaultItem")

// [SetMenuInfo] function.
//
// [SetMenuInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuinfo
func (hMenu HMENU) SetMenuInfo(info *MENUINFO) error {
	ret, _, err := syscall.SyscallN(_SetMenuInfo.Addr(),
		uintptr(hMenu), uintptr(unsafe.Pointer(info)))
	return util.ZeroToGetLastError(ret, err)
}

var _SetMenuInfo = dll.User32.NewProc("SetMenuInfo")

// [SetMenuItemBitmaps] function, using the item command ID.
//
// [SetMenuItemBitmaps]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuitembitmaps
func (hMenu HMENU) SetMenuItemBitmapsByCmd(cmdId uint16, hBmpUnchecked, hBmpChecked HBITMAP) error {
	ret, _, err := syscall.SyscallN(_SetMenuItemBitmaps.Addr(),
		uintptr(hMenu), uintptr(cmdId), uintptr(co.MF_BYCOMMAND),
		uintptr(hBmpUnchecked), uintptr(hBmpChecked))
	return util.ZeroToGetLastError(ret, err)
}

// [SetMenuItemBitmaps] function, using the zero-based item position.
//
// [SetMenuItemBitmaps]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuitembitmaps
func (hMenu HMENU) SetMenuItemBitmapsByPos(index uint, hBmpUnchecked, hBmpChecked HBITMAP) error {
	ret, _, err := syscall.SyscallN(_SetMenuItemBitmaps.Addr(),
		uintptr(hMenu), uintptr(index), uintptr(co.MF_BYPOSITION),
		uintptr(hBmpUnchecked), uintptr(hBmpChecked))
	return util.ZeroToGetLastError(ret, err)
}

var _SetMenuItemBitmaps = dll.User32.NewProc("SetMenuItemBitmaps")

// [SetMenuItemInfo] function, using the item command ID.
//
// [SetMenuItemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuiteminfow
func (hMenu HMENU) SetMenuItemInfoByCmd(cmdId uint16, info *MENUITEMINFO) error {
	info.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(_SetMenuItemInfo.Addr(),
		uintptr(hMenu), uintptr(cmdId), uintptr(co.MF_BYCOMMAND),
		uintptr(unsafe.Pointer(info)))
	return util.ZeroToGetLastError(ret, err)
}

// [SetMenuItemInfo] function, using the zero-based item position.
//
// [SetMenuItemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuiteminfow
func (hMenu HMENU) SetMenuItemInfoByPos(index uint, info *MENUITEMINFO) error {
	info.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(_SetMenuItemInfo.Addr(),
		uintptr(hMenu), uintptr(index), uintptr(co.MF_BYPOSITION),
		uintptr(unsafe.Pointer(info)))
	return util.ZeroToGetLastError(ret, err)
}

var _SetMenuItemInfo = dll.User32.NewProc("SetMenuItemInfo")

// Shows the popup menu anchored at the given coordinates using
// [TrackPopupMenu].
//
// If hCoordsRelativeTo is zero, coordinates must be relative to hParent.
//
// This function will block until the menu disappears.
//
// [TrackPopupMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-trackpopupmenu
func (hMenu HMENU) ShowAtPoint(pos POINT, hParent, hCoordsRelativeTo HWND) {
	if hCoordsRelativeTo == 0 {
		hCoordsRelativeTo = hParent
	}

	hCoordsRelativeTo.ClientToScreenPt(&pos) // now relative to screen
	hParent.SetForegroundWindow()
	hMenu.TrackPopupMenu(co.TPM_LEFTBUTTON, int(pos.X), int(pos.Y), hParent)
	hParent.PostMessage(co.WM_NULL, 0, 0) // necessary according to TrackMenuPopup docs
}

// [TrackPopupMenu] function.
//
// This function will block until the menu disappears.
//
// If TPM_RETURNCMD is passed, returns the selected command ID.
//
// [TrackPopupMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-trackpopupmenu
func (hMenu HMENU) TrackPopupMenu(flags co.TPM, x, y int, hWnd HWND) (int, error) {
	ret, _, err := syscall.SyscallN(_TrackPopupMenu.Addr(),
		uintptr(hMenu), uintptr(flags), uintptr(x), uintptr(y),
		0, uintptr(hWnd), 0)

	if (flags & co.TPM_RETURNCMD) != 0 {
		if ret == 0 && err != 0 {
			return 0, co.ERROR(err)
		} else {
			return int(ret), nil
		}
	} else {
		if ret == 0 {
			return 0, co.ERROR(err)
		} else {
			return 0, nil
		}
	}
}

var _TrackPopupMenu = dll.User32.NewProc("TrackPopupMenu")
