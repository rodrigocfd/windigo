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

type HIMAGELIST HANDLE

func (hImg HIMAGELIST) AddIcon(hIcon HICON) {
	hImg.ReplaceIcon(-1, hIcon)
}

func (hImg HIMAGELIST) Destroy() {
	// http://www.catch22.net/tuts/win32/system-image-list
	// https://www.autohotkey.com/docs/commands/ListView.htm
	syscall.Syscall(proc.ImageList_Destroy.Addr(), 1,
		uintptr(hImg), 0, 0)
}

func (hImg HIMAGELIST) Duplicate() HIMAGELIST {
	ret, _, _ := syscall.Syscall(proc.ImageList_Duplicate.Addr(), 1,
		uintptr(hImg), 0, 0)
	if ret == 0 {
		panic("ImageList_Duplicate failed.")
	}
	return HIMAGELIST(ret)
}

// Returned icon must be destroyed.
func (hImg HIMAGELIST) GetIcon(index uint32, flags co.ILD) HICON {
	ret, _, _ := syscall.Syscall(proc.ImageList_GetIcon.Addr(), 3,
		uintptr(hImg), uintptr(index), uintptr(flags))
	if ret == 0 {
		panic("ImageList_GetIcon failed.")
	}
	return HICON(ret)
}

func (hImg HIMAGELIST) GetIconSize() *SIZE {
	sz := &SIZE{}
	ret, _, _ := syscall.Syscall(proc.ImageList_GetIconSize.Addr(), 3,
		uintptr(hImg),
		uintptr(unsafe.Pointer(&sz.Cx)), uintptr(unsafe.Pointer(&sz.Cy)))
	if ret == 0 {
		panic("ImageList_GetIconSize failed.")
	}
	return sz
}

func (hImg HIMAGELIST) GetImageCount() uint32 {
	ret, _, _ := syscall.Syscall(proc.ImageList_GetImageCount.Addr(), 1,
		uintptr(hImg), 0, 0)
	return uint32(ret)
}

func (hImg HIMAGELIST) GetImageInfo(index uint32) *IMAGEINFO {
	ii := &IMAGEINFO{}
	ret, _, _ := syscall.Syscall(proc.ImageList_GetImageInfo.Addr(), 3,
		uintptr(hImg), uintptr(index), uintptr(unsafe.Pointer(ii)))
	if ret == 0 {
		panic("ImageList_GetImageInfo failed.")
	}
	return ii
}

// Automatically destroyed if attached to a ListView, unless created with
// LVS_SHAREIMAGELISTS. For TreeView, must be manually destroyed.
func ImageListCreate(cx, cy uint32, flags co.ILC,
	cInitial, cGrow uint32) HIMAGELIST {

	ret, _, _ := syscall.Syscall6(proc.ImageList_Create.Addr(), 5,
		uintptr(cx), uintptr(cy), uintptr(flags),
		uintptr(cInitial), uintptr(cGrow), 0)
	if ret == 0 {
		panic("ImageList_Create failed.")
	}
	return HIMAGELIST(ret)
}

// If icon was loaded with LoadIcon(), it doesn't need to be destroyed, because
// all icon resources are automatically freed.
// Otherwise, if CreateIcon(), it can be destroyed after the function returns.
func (hImg HIMAGELIST) ReplaceIcon(i int32, hIcon HICON) int32 {
	ret, _, _ := syscall.Syscall(proc.ImageList_ReplaceIcon.Addr(), 3,
		uintptr(hImg), uintptr(i), uintptr(hIcon))
	return int32(ret)
}
