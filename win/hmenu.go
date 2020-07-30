/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type HMENU HANDLE

func (hMenu HMENU) AppendMenu(
	uFlags co.MF, idOrPos uintptr, bmpOrDataOrStr unsafe.Pointer) {

	ret, _, lerr := syscall.Syscall6(proc.AppendMenu.Addr(), 4,
		uintptr(hMenu), uintptr(uFlags), idOrPos, uintptr(bmpOrDataOrStr),
		0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("AppendMenu failed."))
	}
}

func (hMenu HMENU) CheckMenuItem(idOrPos uintptr, uCheck co.MF) co.MF {
	ret, _, _ := syscall.Syscall(proc.CheckMenuItem.Addr(), 3,
		uintptr(hMenu), idOrPos, uintptr(uCheck))
	return co.MF(ret)
}

// Radio group goes from first to last, check is the selected one.
func (hMenu HMENU) CheckMenuRadioItem(
	firstIdOrPos, lastIdOrPos, checkedIdOrPos uintptr, flags co.MF) {

	ret, _, lerr := syscall.Syscall6(proc.CheckMenuRadioItem.Addr(), 5,
		uintptr(hMenu), firstIdOrPos, lastIdOrPos, checkedIdOrPos,
		uintptr(flags), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("CheckMenuRadioItem failed."))
	}
}

// Creates a horizontal menu bar.
func CreateMenu() HMENU {
	ret, _, lerr := syscall.Syscall(proc.CreateMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("CreateMenu failed."))
	}
	return HMENU(ret)
}

// Creates a vertical popup menu.
func CreatePopupMenu() HMENU {
	ret, _, lerr := syscall.Syscall(proc.CreatePopupMenu.Addr(), 0, 0, 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("CreatePopupMenu failed."))
	}
	return HMENU(ret)
}

func (hMenu HMENU) DeleteMenu(idOrPos uintptr, uFlags co.MF) {
	ret, _, lerr := syscall.Syscall(proc.DeleteMenu.Addr(), 3,
		uintptr(hMenu), idOrPos, uintptr(uFlags))
	if ret == 0 {
		panic(co.ERROR(lerr).Format("DeleteMenu failed."))
	}
}

func (hMenu HMENU) DestroyMenu() {
	ret, _, lerr := syscall.Syscall(proc.DestroyMenu.Addr(), 1,
		uintptr(hMenu), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("DestroyMenu failed."))
	}
}

func (hMenu HMENU) EnableMenuItem(idOrPos uintptr, uEnable co.MF) co.MF {
	ret, _, _ := syscall.Syscall(proc.EnableMenuItem.Addr(), 3,
		uintptr(hMenu), idOrPos, uintptr(uEnable))
	return co.MF(ret)
}

func (hMenu HMENU) GetMenuInfo(mi *MENUINFO) {
	mi.CbSize = uint32(unsafe.Sizeof(*mi)) // safety

	ret, _, lerr := syscall.Syscall(proc.GetMenuInfo.Addr(), 2,
		uintptr(hMenu), uintptr(unsafe.Pointer(mi)), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("GetMenuInfo failed."))
	}
}

func (hMenu HMENU) GetMenuItemCount() uint32 {
	ret, _, lerr := syscall.Syscall(proc.GetMenuItemCount.Addr(), 1,
		uintptr(hMenu), 0, 0)
	if int(ret) == -1 {
		panic(co.ERROR(lerr).Format("GetItemCount failed."))
	}
	return uint32(ret)
}

func (hMenu HMENU) GetMenuItemID(nPos uint32) int32 {
	ret, _, _ := syscall.Syscall(proc.GetMenuItemID.Addr(), 2,
		uintptr(hMenu), uintptr(nPos), 0)
	return int32(ret)
}

func (hMenu HMENU) GetMenuItemInfo(
	idOrPos uintptr, fByPosition bool, lpmii *MENUITEMINFO) {

	lpmii.CbSize = uint32(unsafe.Sizeof(*lpmii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.GetMenuItemInfo.Addr(), 4,
		uintptr(hMenu), idOrPos, boolToUintptr(fByPosition),
		uintptr(unsafe.Pointer(lpmii)), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("GetMenuItemInfo failed."))
	}
}

// If nPos is not a submenu, returns zero.
func (hMenu HMENU) GetSubMenu(nPos uint32) HMENU {
	ret, _, _ := syscall.Syscall(proc.GetSubMenu.Addr(), 2,
		uintptr(hMenu), uintptr(nPos), 0)
	return HMENU(ret)
}

func (hMenu HMENU) InsertMenu(beforeIdOrPos uintptr, uFlags co.MF,
	idOrHmenu uintptr, bmpOrDataOrStr unsafe.Pointer) {

	ret, _, lerr := syscall.Syscall6(proc.InsertMenu.Addr(), uintptr(5),
		uintptr(hMenu), beforeIdOrPos, uintptr(uFlags),
		idOrHmenu, uintptr(bmpOrDataOrStr), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("InsertMenu failed."))
	}
}

func (hMenu HMENU) InsertMenuItem(
	beforeIdOrPos uintptr, fByPosition bool, lpmi *MENUITEMINFO) {

	ret, _, lerr := syscall.Syscall6(proc.InsertMenuItem.Addr(), 4,
		uintptr(hMenu), beforeIdOrPos, boolToUintptr(fByPosition),
		uintptr(unsafe.Pointer(lpmi)), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("InsertMenuItem failed."))
	}
}

func (hMenu HMENU) SetMenuDefaultItem(idOrPos uintptr, fByPos bool) {
	ret, _, lerr := syscall.Syscall(proc.SetMenuDefaultItem.Addr(), 3,
		uintptr(hMenu), idOrPos, boolToUintptr(fByPos))
	if ret == 0 {
		panic(co.ERROR(lerr).Format("SetMenuDefaultItem failed."))
	}
}

func (hMenu HMENU) SetMenuItemInfo(
	idOrPos uintptr, fByPosition bool, lpmii *MENUITEMINFO) {

	lpmii.CbSize = uint32(unsafe.Sizeof(*lpmii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.SetMenuItemInfo.Addr(), 4,
		uintptr(hMenu), idOrPos, boolToUintptr(fByPosition),
		uintptr(unsafe.Pointer(lpmii)), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("SetMenuItemInfo failed."))
	}
}

// This function will block until the menu disappears.
// If TPM_RETURNCMD is passed, returns the selected command ID.
func (hMenu HMENU) TrackPopupMenu(uFlags co.TPM, x, y int32, hWnd HWND) int {
	ret, _, lerr := syscall.Syscall9(proc.TrackPopupMenu.Addr(), 7,
		uintptr(hMenu), uintptr(uFlags), uintptr(x), uintptr(y), 0, uintptr(hWnd),
		0, 0, 0)

	if (uFlags & co.TPM_RETURNCMD) != 0 {
		if ret == 0 && lerr != 0 {
			panic(co.ERROR(lerr).Format("TrackPopupMenu failed."))
		} else {
			return int(ret)
		}
	} else {
		if ret == 0 {
			panic(co.ERROR(lerr).Format("TrackPopupMenu failed."))
		} else {
			return 0
		}
	}
}
