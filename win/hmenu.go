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

func CreateMenu() HMENU {
	ret, _, lerr := syscall.Syscall(proc.CreateMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("CreateMenu failed: %d %s",
			lerr, lerr.Error()))
	}
	return HMENU(ret)
}

func (hMenu HMENU) DeleteMenu(uPosition uintptr, uFlags co.MF) {
	ret, _, lerr := syscall.Syscall(proc.DeleteMenu.Addr(), 3,
		uintptr(hMenu), uPosition, uintptr(uFlags))
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

func (hMenu HMENU) EnableMenuItem(uIDEnableItem uintptr, uEnable co.MF) {
	syscall.Syscall(proc.EnableMenuItem.Addr(), 3,
		uintptr(hMenu), uIDEnableItem, uintptr(uEnable))
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

func (hMenu HMENU) GetMenuInfo(mi *MENUINFO) {
	mi.CbSize = uint32(unsafe.Sizeof(*mi)) // safety

	ret, _, lerr := syscall.Syscall(proc.GetMenuInfo.Addr(), 2,
		uintptr(hMenu), uintptr(unsafe.Pointer(mi)), 0)
	if ret == 0 {
		panic(fmt.Sprintf("GetMenuInfo failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) GetMenuItemID(index uint32) co.ID {
	ret, _, _ := syscall.Syscall(proc.GetMenuItemID.Addr(), 2,
		uintptr(hMenu), uintptr(index), 0)
	return co.ID(ret)
}

func (hMenu HMENU) GetMenuItemInfo(item uintptr, fByPosition bool,
	lpmii *MENUITEMINFO) {

	lpmii.CbSize = uint32(unsafe.Sizeof(*lpmii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.GetMenuItemInfo.Addr(), 4,
		uintptr(hMenu), item, boolToUintptr(fByPosition),
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

func (hMenu HMENU) SetMenuItemInfo(item uintptr, fByPosition bool,
	lpmii *MENUITEMINFO) {

	lpmii.CbSize = uint32(unsafe.Sizeof(*lpmii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.SetMenuItemInfo.Addr(), 4,
		uintptr(hMenu), item, boolToUintptr(fByPosition),
		uintptr(unsafe.Pointer(lpmii)), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("SetMenuItemInfo failed: %d %s",
			lerr, lerr.Error()))
	}
}
