//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a [menu].
//
// [menu]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hmenu
type HMENU HANDLE

// [CreateMenu] function.
//
// ⚠️ You must defer [HMENU.DestroyMenu], unless it's attached to a window.
//
// [CreateMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createmenu
func CreateMenu() (HMENU, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_CreateMenu, "CreateMenu"))
	if ret == 0 {
		return HMENU(0), co.ERROR(err)
	}
	return HMENU(ret), nil
}

var _CreateMenu *syscall.Proc

// [CreatePopupMenu] function.
//
// ⚠️ You must defer [HMENU.DestroyMenu], unless it's attached to a window.
//
// [CreatePopupMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createpopupmenu
func CreatePopupMenu() (HMENU, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_CreatePopupMenu, "CreatePopupMenu"))
	if ret == 0 {
		return HMENU(0), co.ERROR(err)
	}
	return HMENU(ret), nil
}

var _CreatePopupMenu *syscall.Proc

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

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_CheckMenuItem, "CheckMenuItem"),
		uintptr(hMenu),
		uintptr(item),
		uintptr(flagPosCmd))
	return utl.Minus1AsSysInvalidParm(ret)
}

var _CheckMenuItem *syscall.Proc

// [DeleteMenu] function for multiple items, using the item command ID.
//
// [DeleteMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deletemenu
func (hMenu HMENU) DeleteMenuByCmd(cmdIds ...uint16) error {
	for _, cmdId := range cmdIds {
		ret, _, err := syscall.SyscallN(
			dll.Load(dll.USER32, &_DeleteMenu, "DeleteMenu"),
			uintptr(hMenu),
			uintptr(uint32(cmdId)),
			uintptr(co.MF_BYCOMMAND))
		if ret == 0 {
			return co.ERROR(err)
		}
	}
	return nil
}

var _DeleteMenu *syscall.Proc

// [DeleteMenu] function for multiple items, using the zero-based item position.
//
// [DeleteMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deletemenu
func (hMenu HMENU) DeleteMenuByPos(indexes ...uint) error {
	for _, index := range indexes {
		ret, _, err := syscall.SyscallN(
			dll.Load(dll.USER32, &_DeleteMenu, "DeleteMenu"),
			uintptr(hMenu),
			uintptr(uint32(index)),
			uintptr(co.MF_BYPOSITION))
		if ret == 0 {
			return co.ERROR(err)
		}
	}
	return nil
}

// [DestroyMenu] function.
//
// [DestroyMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroymenu
func (hMenu HMENU) DestroyMenu() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_DestroyMenu, "DestroyMenu"),
		uintptr(hMenu))
	return utl.ZeroAsGetLastError(ret, err)
}

var _DestroyMenu *syscall.Proc

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

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_EnableMenuItem, "EnableMenuItem"),
		uintptr(hMenu),
		uintptr(uint32(item)),
		uintptr(flagPosCmd))
	return utl.Minus1AsSysInvalidParm(ret)
}

var _EnableMenuItem *syscall.Proc

// [GetMenuDefaultItem] function.
//
// Returns the zero-based index of the item.
//
// [GetMenuDefaultItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenudefaultitem
func (hMenu HMENU) GetMenuDefaultItem(gmdiFlags co.GMDI) (uint, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetMenuDefaultItem, "GetMenuDefaultItem"),
		uintptr(hMenu),
		1,
		uintptr(gmdiFlags))
	if wErr := co.ERROR(err); int32(ret) == -1 || wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint(ret), nil
}

var _GetMenuDefaultItem *syscall.Proc

// [GetMenuItemID] function.
//
// Given the zero-based index, returns its command ID.
//
// [GetMenuItemID]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuitemid
func (hMenu HMENU) GetMenuItemID(index uint) (uint16, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetMenuItemID, "GetMenuItemID"),
		uintptr(hMenu),
		uintptr(int32(index)))
	if int32(ret) == -1 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return uint16(ret), nil
}

var _GetMenuItemID *syscall.Proc

// [GetMenuItemCount] function.
//
// [GetMenuItemCount]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuitemcount
func (hMenu HMENU) GetMenuItemCount() (uint, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetMenuItemCount, "GetMenuItemCount"),
		uintptr(hMenu))
	if int32(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return uint(ret), nil
}

var _GetMenuItemCount *syscall.Proc

// [GetMenuItemInfo] function, using the item command ID.
//
// [GetMenuItemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuiteminfow
func (hMenu HMENU) GetMenuItemInfoByCmd(cmdId uint16, mii *MENUITEMINFO) error {
	mii.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetMenuItemInfoW, "GetMenuItemInfoW"),
		uintptr(hMenu),
		uintptr(uint32(cmdId)),
		uintptr(co.MF_BYCOMMAND),
		uintptr(unsafe.Pointer(mii)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _GetMenuItemInfoW *syscall.Proc

// [GetMenuItemInfo] function, using the zero-based item position.
//
// [GetMenuItemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenuiteminfow
func (hMenu HMENU) GetMenuItemInfoByPos(index uint, mii *MENUITEMINFO) error {
	mii.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetMenuItemInfoW, "GetMenuItemInfoW"),
		uintptr(hMenu),
		uintptr(uint32(index)),
		uintptr(co.MF_BYPOSITION),
		uintptr(unsafe.Pointer(mii)))
	return utl.ZeroAsGetLastError(ret, err)
}

// [GetSubMenu] function.
//
// [GetSubMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsubmenu
func (hMenu HMENU) GetSubMenu(index uint) (HMENU, bool) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetSubMenu, "GetSubMenu"),
		uintptr(hMenu),
		uintptr(int32(index)))
	hSub := HMENU(ret)
	return hSub, hSub != 0
}

var _GetSubMenu *syscall.Proc

// [InsertMenuItem] function, using the item command ID.
//
// [InsertMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insertmenuitemw
func (hMenu HMENU) InsertMenuItemByCmd(cmdId uint16, mii *MENUITEMINFO) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_InsertMenuItemW, "InsertMenuItemW"),
		uintptr(hMenu),
		uintptr(uint32(cmdId)),
		uintptr(co.MF_BYCOMMAND),
		uintptr(unsafe.Pointer(mii)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _InsertMenuItemW *syscall.Proc

// [InsertMenuItem] function, using the zero-based item position.
//
// [InsertMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insertmenuitemw
func (hMenu HMENU) InsertMenuItemByPos(index uint, mii *MENUITEMINFO) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_InsertMenuItemW, "InsertMenuItemW"),
		uintptr(hMenu),
		uintptr(uint32(index)),
		uintptr(co.MF_BYPOSITION),
		uintptr(unsafe.Pointer(mii)))
	return utl.ZeroAsGetLastError(ret, err)
}

// [RemoveMenu] function for multiple items, using the item command ID.
//
// [RemoveMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-removemenu
func (hMenu HMENU) RemoveMenuByCmd(cmdIds ...uint16) error {
	for _, cmdId := range cmdIds {
		ret, _, err := syscall.SyscallN(
			dll.Load(dll.USER32, &_RemoveMenu, "RemoveMenu"),
			uintptr(hMenu),
			uintptr(uint32(cmdId)),
			uintptr(co.MF_BYCOMMAND))
		if ret == 0 {
			return co.ERROR(err)
		}
	}
	return nil
}

var _RemoveMenu *syscall.Proc

// [RemoveMenu] function for multiple items, using the zero-based item position.
//
// [RemoveMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-removemenu
func (hMenu HMENU) RemoveMenuByPos(indexes ...uint) error {
	for _, index := range indexes {
		ret, _, err := syscall.SyscallN(
			dll.Load(dll.USER32, &_RemoveMenu, "RemoveMenu"),
			uintptr(hMenu),
			uintptr(uint32(index)),
			uintptr(co.MF_BYPOSITION))
		if ret == 0 {
			return co.ERROR(err)
		}
	}
	return nil
}

// [SetMenuDefaultItem] function, using the item command ID.
//
// [SetMenuDefaultItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenudefaultitem
func (hMenu HMENU) SetMenuDefaultItemByCmd(cmdId uint16) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetMenuDefaultItem, "SetMenuDefaultItem"),
		uintptr(hMenu),
		uintptr(uint32(cmdId)),
		uintptr(co.MF_BYCOMMAND))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetMenuDefaultItem *syscall.Proc

// [SetMenuDefaultItem] function, using the zero-based item position.
//
// [SetMenuDefaultItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenudefaultitem
func (hMenu HMENU) SetMenuDefaultItemByPos(index uint) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetMenuDefaultItem, "SetMenuDefaultItem"),
		uintptr(hMenu),
		uintptr(uint32(index)),
		uintptr(co.MF_BYPOSITION))
	return utl.ZeroAsGetLastError(ret, err)
}

// [SetMenuInfo] function.
//
// [SetMenuInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuinfo
func (hMenu HMENU) SetMenuInfo(info *MENUINFO) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetMenuInfo, "SetMenuInfo"),
		uintptr(hMenu),
		uintptr(unsafe.Pointer(info)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetMenuInfo *syscall.Proc

// [SetMenuItemBitmaps] function, using the item command ID.
//
// [SetMenuItemBitmaps]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuitembitmaps
func (hMenu HMENU) SetMenuItemBitmapsByCmd(cmdId uint16, hBmpUnchecked, hBmpChecked HBITMAP) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetMenuItemBitmaps, "SetMenuItemBitmaps"),
		uintptr(hMenu),
		uintptr(uint32(cmdId)),
		uintptr(co.MF_BYCOMMAND),
		uintptr(hBmpUnchecked),
		uintptr(hBmpChecked))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetMenuItemBitmaps *syscall.Proc

// [SetMenuItemBitmaps] function, using the zero-based item position.
//
// [SetMenuItemBitmaps]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuitembitmaps
func (hMenu HMENU) SetMenuItemBitmapsByPos(index uint, hBmpUnchecked, hBmpChecked HBITMAP) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetMenuItemBitmaps, "SetMenuItemBitmaps"),
		uintptr(hMenu),
		uintptr(uint32(index)),
		uintptr(co.MF_BYPOSITION),
		uintptr(hBmpUnchecked),
		uintptr(hBmpChecked))
	return utl.ZeroAsGetLastError(ret, err)
}

// [SetMenuItemInfo] function, using the item command ID.
//
// [SetMenuItemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuiteminfow
func (hMenu HMENU) SetMenuItemInfoByCmd(cmdId uint16, info *MENUITEMINFO) error {
	info.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetMenuItemInfo, "SetMenuItemInfo"),
		uintptr(hMenu),
		uintptr(uint32(cmdId)),
		uintptr(co.MF_BYCOMMAND),
		uintptr(unsafe.Pointer(info)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetMenuItemInfo *syscall.Proc

// [SetMenuItemInfo] function, using the zero-based item position.
//
// [SetMenuItemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuiteminfow
func (hMenu HMENU) SetMenuItemInfoByPos(index uint, info *MENUITEMINFO) error {
	info.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetMenuItemInfo, "SetMenuItemInfo"),
		uintptr(hMenu),
		uintptr(uint32(index)),
		uintptr(co.MF_BYPOSITION),
		uintptr(unsafe.Pointer(info)))
	return utl.ZeroAsGetLastError(ret, err)
}

// Shows the popup menu anchored at the given coordinates using
// [HMENU.TrackPopupMenu].
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
	hMenu.TrackPopupMenu(co.TPM_LEFTBUTTON, int(pos.X), int(pos.Y), hParent)
	hParent.PostMessage(co.WM_NULL, 0, 0) // necessary according to TrackMenuPopup docs
}

// [TrackPopupMenu] function.
//
// This function will block until the menu disappears.
//
// If [co.TPM_RETURNCMD] is passed, returns the selected command ID.
//
// [TrackPopupMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-trackpopupmenu
func (hMenu HMENU) TrackPopupMenu(flags co.TPM, x, y int, hWnd HWND) (int, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_TrackPopupMenu, "TrackPopupMenu"),
		uintptr(hMenu),
		uintptr(flags),
		uintptr(int32(x)),
		uintptr(int32(y)),
		0,
		uintptr(hWnd),
		0)

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

var _TrackPopupMenu *syscall.Proc
