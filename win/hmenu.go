/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type HMENU HANDLE

func (hMenu HMENU) AppendMenu(uFlags co.MF, uIDNewItem uintptr,
	lpNewItem uintptr) {

	ret, _, lerr := syscall.Syscall6(proc.AppendMenu.Addr(), 4,
		uintptr(hMenu), uintptr(uFlags), uIDNewItem, uintptr(lpNewItem),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("AppendMenu failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) CheckMenuItem(uIDCheckItem int32, uCheck co.MF) co.MF {
	ret, _, _ := syscall.Syscall(proc.CheckMenuItem.Addr(), 3,
		uintptr(hMenu), uintptr(uIDCheckItem), uintptr(uCheck))
	return co.MF(ret)
}

func (hMenu HMENU) CheckMenuRadioItem(first, last, check int32, flags co.MF) {
	ret, _, lerr := syscall.Syscall6(proc.CheckMenuRadioItem.Addr(), 5,
		uintptr(hMenu), uintptr(first), uintptr(last), uintptr(check),
		uintptr(flags), 0)
	if ret == 0 {
		panic(fmt.Sprintf("CheckMenuRadioItem failed: %d %s",
			lerr, lerr.Error()))
	}
}

// Creates a horizontal menu bar.
func CreateMenu() HMENU {
	ret, _, lerr := syscall.Syscall(proc.CreateMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("CreateMenu failed: %d %s",
			lerr, lerr.Error()))
	}
	return HMENU(ret)
}

// Creates a vertical popup menu.
func CreatePopupMenu() HMENU {
	ret, _, lerr := syscall.Syscall(proc.CreatePopupMenu.Addr(), 0, 0, 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("CreatePopupMenu failed: %d %s",
			lerr, lerr.Error()))
	}
	return HMENU(ret)
}

func (hMenu HMENU) DeleteMenu(uPosition int32, uFlags co.MF) {
	ret, _, lerr := syscall.Syscall(proc.DeleteMenu.Addr(), 3,
		uintptr(hMenu), uintptr(uPosition), uintptr(uFlags))
	if ret == 0 {
		panic(fmt.Sprintf("DeleteMenu failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) DestroyMenu() {
	ret, _, lerr := syscall.Syscall(proc.DestroyMenu.Addr(), 1,
		uintptr(hMenu), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("DestroyMenu failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) EnableMenuItem(uIDEnableItem int32, uEnable co.MF) co.MF {
	ret, _, _ := syscall.Syscall(proc.EnableMenuItem.Addr(), 3,
		uintptr(hMenu), uintptr(uIDEnableItem), uintptr(uEnable))
	return co.MF(ret)
}

func (hMenu HMENU) GetMenuInfo(mi *MENUINFO) {
	mi.CbSize = uint32(unsafe.Sizeof(*mi)) // safety

	ret, _, lerr := syscall.Syscall(proc.GetMenuInfo.Addr(), 2,
		uintptr(hMenu), uintptr(unsafe.Pointer(mi)), 0)
	if ret == 0 {
		panic(fmt.Sprintf("GetMenuInfo failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) GetMenuItemCount() uint32 {
	ret, _, lerr := syscall.Syscall(proc.GetMenuItemCount.Addr(), 1,
		uintptr(hMenu), 0, 0)
	if int32(ret) == -1 {
		panic(fmt.Sprintf("GetItemCount failed: %d %s",
			lerr, lerr.Error()))
	}
	return uint32(ret)
}

func (hMenu HMENU) GetMenuItemID(index uint32) int32 {
	ret, _, _ := syscall.Syscall(proc.GetMenuItemID.Addr(), 2,
		uintptr(hMenu), uintptr(index), 0)
	return int32(ret)
}

func (hMenu HMENU) GetMenuItemInfo(item int32, fByPosition bool,
	lpmii *MENUITEMINFO) {

	lpmii.CbSize = uint32(unsafe.Sizeof(*lpmii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.GetMenuItemInfo.Addr(), 4,
		uintptr(hMenu), uintptr(item), boolToUintptr(fByPosition),
		uintptr(unsafe.Pointer(lpmii)), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("GetMenuItemInfo failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) GetSubMenu(nPos uint32) HMENU {
	ret, _, _ := syscall.Syscall(proc.GetSubMenu.Addr(), 2,
		uintptr(hMenu), uintptr(nPos), 0)
	return HMENU(ret)
}

func (hMenu HMENU) InsertMenu(uPosition int32, uFlags co.MF,
	uIDNewItem uintptr, lpNewItem uintptr) {

	ret, _, lerr := syscall.Syscall6(proc.InsertMenu.Addr(), uintptr(5),
		uintptr(hMenu), uintptr(uPosition), uintptr(uFlags),
		uIDNewItem, lpNewItem, 0)
	if ret == 0 {
		panic(fmt.Sprintf("InsertMenu failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) InsertMenuItem(item int32, fByPosition bool,
	lpmi *MENUITEMINFO) {

	ret, _, lerr := syscall.Syscall6(proc.InsertMenuItem.Addr(), 4,
		uintptr(hMenu), uintptr(item), boolToUintptr(fByPosition),
		uintptr(unsafe.Pointer(lpmi)), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("InsertMenuItem failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) SetMenuDefaultItem(uItem int32, fByPos bool) {
	ret, _, lerr := syscall.Syscall(proc.SetMenuDefaultItem.Addr(), 3,
		uintptr(hMenu), uintptr(uItem), boolToUintptr(fByPos))
	if ret == 0 {
		panic(fmt.Sprintf("SetMenuDefaultItem failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) SetMenuItemInfo(item int32, fByPosition bool,
	lpmii *MENUITEMINFO) {

	lpmii.CbSize = uint32(unsafe.Sizeof(*lpmii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.SetMenuItemInfo.Addr(), 4,
		uintptr(hMenu), uintptr(item), boolToUintptr(fByPosition),
		uintptr(unsafe.Pointer(lpmii)), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("SetMenuItemInfo failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) TrackPopupMenu(uFlags co.TPM, x, y int32, hWnd HWND) int32 {
	ret, _, lerr := syscall.Syscall9(proc.TrackPopupMenu.Addr(), 7,
		uintptr(hMenu), uintptr(uFlags), uintptr(x), uintptr(y), 0, uintptr(hWnd),
		0, 0, 0)
	if (uFlags & co.TPM_RETURNCMD) != 0 {
		if ret == 0 && lerr != 0 {
			panic(fmt.Sprintf("TrackPopupMenu failed: %d %s",
				lerr, lerr.Error()))
		} else {
			return int32(ret)
		}
	} else {
		if ret == 0 {
			panic(fmt.Sprintf("TrackPopupMenu failed: %d %s",
				lerr, lerr.Error()))
		} else {
			return 0
		}
	}
}
