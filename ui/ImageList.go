/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/co"
	"windigo/win"
)

// Native image list resource.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/image-lists
type ImageList struct {
	himl       win.HIMAGELIST
	resolution uint
}

// Automatically destroyed if attached to a ListView, unless created with
// LVS_SHAREIMAGELISTS. Other than that, must always be manually destroyed.
func (me *ImageList) Create(resolution, iconCount uint) *ImageList {
	if me.himl != 0 {
		panic("ImageList already created.")
	}
	me.himl = win.ImageListCreate(uint32(resolution), uint32(resolution),
		co.ILC_COLOR32, uint32(iconCount), 1)
	me.resolution = resolution
	return me
}

func (me *ImageList) AddIcon(icon win.HICON) *ImageList {
	me.himl.AddIcon(icon)
	return me
}

// Loads an icon resource and adds it to the image list.
func (me *ImageList) AddResourceIcon(resourceId int) *ImageList {
	return me.AddIcon(
		win.GetModuleHandle("").LoadIcon(co.IDI(resourceId)),
	)
}

// Loads an icon from Windows Shell.
//
// Extension must be in the format "*.mp3".
func (me *ImageList) AddShellIcon(fileExtension string) *ImageList {
	szFlag := co.SHGFI_NONE
	if me.resolution == 16 {
		szFlag = co.SHGFI_SMALLICON
	} else if me.resolution == 32 {
		szFlag = co.SHGFI_LARGEICON
	} else {
		panic("AddShellIcon implemented only for 16 and 32 icon resolutions.")
	}

	shfi := win.SHGetFileInfo(fileExtension, co.FILE_ATTRIBUTE_NORMAL,
		co.SHGFI_USEFILEATTRIBUTES|co.SHGFI_ICON|szFlag)
	return me.AddIcon(shfi.HIcon)
}

func (me *ImageList) Destroy() {
	me.himl.Destroy()
}

func (me *ImageList) Himagelist() win.HIMAGELIST {
	return me.himl
}
