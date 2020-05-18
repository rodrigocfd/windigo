/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package api

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/api/proc"
	c "wingows/consts"
)

type HMENU HANDLE

func (hmenu HMENU) AppendMenu(uFlags c.MF, uIDNewItem uintptr,
	lpNewItem uintptr) {

	ret, _, lerr := syscall.Syscall6(proc.AppendMenu.Addr(), 4,
		uintptr(hmenu), uintptr(uFlags), uIDNewItem, uintptr(lpNewItem),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("AppendMenu failed: %d %s\n",
			lerr, lerr.Error()))
	}
}

func CreateMenu() HMENU {
	ret, _, lerr := syscall.Syscall(proc.CreateMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("CreateMenu failed: %d %s\n",
			lerr, lerr.Error()))
	}
	return HMENU(ret)
}

func (hmenu HMENU) DeleteMenuById(id c.ID) {
	ret, _, lerr := syscall.Syscall(proc.DeleteMenu.Addr(), 3,
		uintptr(hmenu), uintptr(id), uintptr(c.MF_BYCOMMAND))
	if ret == 0 {
		panic(fmt.Sprintf("DeleteMeny by ID failed: %d %s\n",
			lerr, lerr.Error()))
	}
}

func (hmenu HMENU) DeleteMenuByPos(index uint32) {
	ret, _, lerr := syscall.Syscall(proc.DeleteMenu.Addr(), 3,
		uintptr(hmenu), uintptr(index), uintptr(c.MF_BYPOSITION))
	if ret == 0 {
		panic(fmt.Sprintf("DeleteMeny by pos failed: %d %s\n",
			lerr, lerr.Error()))
	}
}

func (hmenu HMENU) DestroyMenu() {
	ret, _, lerr := syscall.Syscall(proc.DestroyMenu.Addr(), 1,
		uintptr(hmenu), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("DestroyMenu failed: %d %s\n",
			lerr, lerr.Error()))
	}
}

func (hmenu HMENU) EnableMenuItem(uIDEnableItem uint32, uEnable c.MF) {
	syscall.Syscall(proc.EnableMenuItem.Addr(), 3,
		uintptr(hmenu), uintptr(uIDEnableItem), uintptr(uEnable))
}

func (hmenu HMENU) GetMenuItemCount() uint32 {
	ret, _, lerr := syscall.Syscall(proc.GetMenuItemCount.Addr(), 1,
		uintptr(hmenu), 0, 0)
	if int32(ret) == -1 {
		panic(fmt.Sprintf("GetItemCount failed: %d %s\n",
			lerr, lerr.Error()))
	}
	return uint32(ret)
}

func (hmenu HMENU) GetMenuInfo(mi *MENUINFO) {
	mi.CbSize = uint32(unsafe.Sizeof(*mi)) // safety

	ret, _, lerr := syscall.Syscall(proc.GetMenuInfo.Addr(), 2,
		uintptr(hmenu), uintptr(unsafe.Pointer(mi)), 0)
	if ret == 0 {
		panic(fmt.Sprintf("GetMenuInfo failed: %d %s\n",
			lerr, lerr.Error()))
	}
}

func (hmenu HMENU) GetMenuItemID(index uint32) c.ID {
	ret, _, _ := syscall.Syscall(proc.GetMenuItemID.Addr(), 2,
		uintptr(hmenu), uintptr(index), 0)
	return c.ID(ret)
}

func (hmenu HMENU) GetMenuItemInfoById(id c.ID, mii *MENUITEMINFO) {
	mii.CbSize = uint32(unsafe.Sizeof(*mii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.GetMenuItemInfo.Addr(), 4,
		uintptr(hmenu), uintptr(id), 0, uintptr(unsafe.Pointer(mii)),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("GetMenuItemInfo by ID failed: %d %s\n",
			lerr, lerr.Error()))
	}
}

func (hmenu HMENU) GetMenuItemInfoByPos(index uint32, mii *MENUITEMINFO) {
	mii.CbSize = uint32(unsafe.Sizeof(*mii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.GetMenuItemInfo.Addr(), 4,
		uintptr(hmenu), uintptr(index), 1, uintptr(unsafe.Pointer(mii)),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("GetMenuItemInfo by pos failed: %d %s\n",
			lerr, lerr.Error()))
	}
}

func (hmenu HMENU) GetSubMenu(nPos uint32) HMENU {
	ret, _, _ := syscall.Syscall(proc.GetSubMenu.Addr(), 2,
		uintptr(hmenu), uintptr(nPos), 0)
	return HMENU(ret)
}

func (hmenu HMENU) SetMenuItemInfoById(id c.ID, mii *MENUITEMINFO) {
	mii.CbSize = uint32(unsafe.Sizeof(*mii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.SetMenuItemInfo.Addr(), 4,
		uintptr(hmenu), uintptr(id), 0, uintptr(unsafe.Pointer(mii)),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("SetMenuItemInfo by ID failed: %d %s\n",
			lerr, lerr.Error()))
	}
}

func (hmenu HMENU) SetMenuItemInfoByPos(index uint32, mii *MENUITEMINFO) {
	mii.CbSize = uint32(unsafe.Sizeof(*mii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.SetMenuItemInfo.Addr(), 4,
		uintptr(hmenu), uintptr(index), 1, uintptr(unsafe.Pointer(mii)),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("SetMenuItemInfo by pos failed: %d %s\n",
			lerr, lerr.Error()))
	}
}
