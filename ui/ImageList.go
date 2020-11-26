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

// Manages an image list resource.
type ImageList struct {
	himl win.HIMAGELIST
}

// Constructor.
//
// You must defer Destroy().
func NewImageList(cx, cy int) *ImageList {
	return &ImageList{
		himl: win.ImageListCreate(uint32(cx), uint32(cy), co.ILC_COLOR32, 0, 1),
	}
}

// Releases the image list resource.
func (me *ImageList) Destroy() {
	if me.himl != 0 {
		me.himl.Destroy()
	}
}

// Creates a copy of the icon, and adds this copy to the image list.
//
// If icon was loaded with LoadIcon(), it doesn't need to be destroyed, because
// all icon resources are automatically freed.
// Otherwise, if loaded with CreateIcon(), it must be destroyed.
func (me *ImageList) AddIcon(hIcons ...win.HICON) *ImageList {
	me.himl.AddIcon(hIcons...)
	return me
}

// Loads into the image list an icon from the shell, used by Windows Explorer to
// represent the given file extensions, like "mp3".
// The icon can have 16x16 or 32x32 pixels only.
func (me *ImageList) AddShellIcon(fileExtensions ...string) *ImageList {
	sz := me.ImageSize()
	if !sz.equals(Size{Cx: 16, Cy: 16}) && !sz.equals(Size{Cx: 32, Cy: 32}) {
		panic("ImageList can load only 16x16 or 32x32 shell icons.")
	}

	shgfi := co.SHGFI_SMALLICON
	if sz.equals(Size{Cx: 32, Cy: 32}) {
		shgfi = co.SHGFI_LARGEICON
	}

	for _, fileExt := range fileExtensions {
		shfi := win.SHGetFileInfo("*."+fileExt, co.FILE_ATTRIBUTE_NORMAL,
			co.SHGFI_USEFILEATTRIBUTES|co.SHGFI_ICON|shgfi)
		me.himl.AddIcon(shfi.HIcon)
		shfi.HIcon.DestroyIcon()
	}
	return me
}

// Retrieves the number of stored images.
func (me *ImageList) Count() int {
	return int(me.himl.GetImageCount())
}

// Returns the underlying HIMAGELIST handle.
func (me *ImageList) Himagelist() win.HIMAGELIST {
	return me.himl
}

// Retrieves the image size being stored in this image list.
func (me *ImageList) ImageSize() Size {
	sz := me.himl.GetIconSize()
	return Size{Cx: int(sz.Cx), Cy: int(sz.Cy)}
}
