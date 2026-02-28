//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// An icon to be loaded, either from resource or from a shell file extension.
type Ico struct {
	id  uint16 // Resource ID.
	ext string // File extension.
}

// Will load the icon with the given resource ID from the resource.
func IcoId(iconId uint16) Ico {
	return Ico{iconId, ""}
}

// Will load the icon of the given file extension, as displayed by the shell.
func IcoExt(fileExtension string) Ico {
	return Ico{0, fileExtension}
}

// Returns true if there is an icon ID or a shell extension.
func (me *Ico) isValid() bool {
	return me.id > 0 || len(me.ext) > 0
}

// If the icon is a resource ID, returns it and true.
func (me *Ico) Id() (uint16, bool) {
	return me.id, me.id > 0
}

// If the icon is a shell file extension, returns it and true.
func (me *Ico) Ext() (string, bool) {
	return me.ext, me.id <= 0
}

// Caches icons in a [win.HIMAGELIST], and returns them on-demand.
type _IconCacheImgList struct {
	hImgList win.HIMAGELIST
	entries  []Ico
}

// Constructor.
func newIconCacheImgList() _IconCacheImgList {
	return _IconCacheImgList{}
}

// Destroys the image list.
func (me *_IconCacheImgList) Release() {
	if me.hImgList != win.HIMAGELIST(0) {
		me.hImgList.Destroy()
		me.hImgList = win.HIMAGELIST(0)
	}
	me.entries = nil
}

// Returns the entry at the given zero-based index, if any.
func (me *_IconCacheImgList) EntryByIndex(index int) (Ico, bool) {
	if index >= 0 && index < len(me.entries) {
		return me.entries[index], true
	}
	return Ico{}, false
}

// Creates the image list, if not created yet. Loads the given icon, if not yet.
// Returns:
//   - the handle to the image list;
//   - true if the image list was created on this call;
//   - zero-based index of the icon within the image list.
func (me *_IconCacheImgList) IconIndex(
	resolution int,
	ico Ico,
) (hImgList win.HIMAGELIST, newImgList bool, idxIcon int) {
	for idx, entry := range me.entries {
		if entry == ico {
			return me.hImgList, false, idx // already cached
		}
	}

	justCreateImgList := false
	if me.hImgList == win.HIMAGELIST(0) { // image list not created yet?
		hImgList, err := win.ImageListCreate(resolution, resolution, co.ILC_COLOR32, 1, 4) // arbitrary grow size
		if err != nil {
			panic("ImageListCreate failed.")
		}
		me.hImgList = hImgList
		me.entries = make([]Ico, 0, 4) // arbitrary
		justCreateImgList = true
	}

	if ico.id > 0 {
		if err := me.hImgList.AddIconFromResource(ico.id); err != nil {
			panic("AddIconFromResource failed " + err.Error())
		}
	} else {
		if err := me.hImgList.AddIconFromShell(ico.ext); err != nil {
			panic("AddIconFromResource failed " + err.Error())
		}
	}

	me.entries = append(me.entries, ico)
	return me.hImgList, justCreateImgList, len(me.entries) - 1 // index of last icon
}

// Caches raw [win.HICON] handles.
type _IconCacheHicon struct {
	hIcons  []win.HICON
	entries []Ico
}

// Constructor.
func newIconCacheHicon() _IconCacheHicon {
	return _IconCacheHicon{}
}

// Destroys all icons.
func (me *_IconCacheHicon) Release() {
	for _, hIcon := range me.hIcons {
		hIcon.DestroyIcon()
	}
	me.hIcons = nil
	me.entries = nil
}

// Loads the icon, if not yet. Returns its handle.
func (me *_IconCacheHicon) Handle(resolution int, ico Ico) win.HICON {
	for idx, entry := range me.entries {
		if entry == ico {
			return me.hIcons[idx] // already cached
		}
	}

	var hIconNew win.HICON
	if ico.id > 0 {
		hInst, _ := win.GetModuleHandle("")
		hIcon, err := hInst.LoadIcon(win.IconResId(ico.id))
		if err != nil {
			panic("LoadIcon failed " + err.Error())
		}
		hIconNew = hIcon
	} else {
		shgfi := co.SHGFI_USEFILEATTRIBUTES | co.SHGFI_ICON
		if resolution == 16 {
			shgfi |= co.SHGFI_SMALLICON
		} else {
			shgfi |= co.SHGFI_LARGEICON
		}
		fi, err := win.SHGetFileInfo("*."+ico.ext, co.FILE_ATTRIBUTE_NORMAL, shgfi)
		if err != nil {
			panic("SHGetFileInfo failed: " + err.Error())
		}
		hIconNew = fi.HIcon
	}

	me.hIcons = append(me.hIcons, hIconNew)
	me.entries = append(me.entries, ico)
	return hIconNew
}
