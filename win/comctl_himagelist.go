//go:build windows

package win

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to an [image list].
//
// # Example
//
//	hImg, _ := win.ImageListCreate(16, 16, co.ILC_COLOR32, 1, 1)
//	defer hImg.Destroy()
//
// [image list]: https://learn.microsoft.com/en-us/windows/win32/controls/image-lists
type HIMAGELIST HANDLE

// [ImageList_Create] function.
//
// ⚠️ You must defer HIMAGELIST.Destroy().
//
// # Example
//
//	hImg, _ := win.ImageListCreate(16, 16, co.ILC_COLOR32, 1, 1)
//	defer hImg.Destroy()
//
// [ImageList_Create]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_create
func ImageListCreate(cx, cy uint, flags co.ILC, szInitial, szGrow uint) (HIMAGELIST, error) {
	ret, _, _ := syscall.SyscallN(_ImageList_Create.Addr(),
		uintptr(cx), uintptr(cy), uintptr(flags),
		uintptr(szInitial), uintptr(szGrow))
	if ret == 0 {
		return HIMAGELIST(0), co.ERROR_INVALID_PARAMETER
	}
	return HIMAGELIST(ret), nil
}

var _ImageList_Create = dll.Comctl32.NewProc("ImageList_Create")

// [ImageList_AddIcon] function.
//
// If icon was loaded from resource with [LoadIcon], it doesn't need to be
// destroyed, because all icon resources are automatically freed.
// Otherwise, if loaded with CreateIcon(), it must be destroyed.
//
// [ImageList_AddIcon]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_addicon
// [LoadIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
func (hImg HIMAGELIST) AddIcon(hIcons ...HICON) error {
	for _, hIco := range hIcons {
		if _, err := hImg.ReplaceIcon(-1, hIco); err != nil {
			return err
		}
	}
	return nil
}

// Calls [LoadIcon] and [ImageList_AddIcon] to load an icon from the resource.
//
// [LoadIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
// [ImageList_AddIcon]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_addicon
func (hImg HIMAGELIST) AddIconFromResource(iconIds ...uint16) error {
	hInst, err := GetModuleHandle("")
	if err != nil {
		return err
	}

	for _, iconId := range iconIds {
		hIco, err := hInst.LoadIcon(IconResId(iconId))
		if err != nil {
			return err
		}
		if err := hImg.AddIcon(hIco); err != nil {
			return err
		}
	}
	return nil
}

// Calls [SHGetFileInfo] to load icons from the shell, used by Windows Explorer
// to represent the given file extensions, like "mp3".
//
// # Example
//
//	hImg, _ := win.ImageListCreate(16, 16, co.ILC_COLOR32, 1, 1)
//	defer hImg.Destroy()
//
//	hImg.AddIconFromShell("mp3", "wav")
//
// [SHGetFileInfo]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
func (hImg HIMAGELIST) AddIconFromShell(fileExtensions ...string) error {
	sz, err := hImg.GetIconSize()
	if err != nil {
		return err
	}

	isIco16 := sz.Cx == 16 && sz.Cy == 16
	isIco32 := sz.Cx == 32 && sz.Cy == 32
	if !isIco16 && !isIco32 {
		return errors.New("AddIconFromShell can load only 16x16 or 32x32 icons")
	}

	shgfi := co.SHGFI_USEFILEATTRIBUTES | co.SHGFI_ICON
	if isIco16 {
		shgfi |= co.SHGFI_SMALLICON
	} else if isIco32 {
		shgfi |= co.SHGFI_LARGEICON
	}

	for _, fileExt := range fileExtensions {
		fi, err := SHGetFileInfo("*."+fileExt, co.FILE_ATTRIBUTE_NORMAL, shgfi)
		if err != nil {
			return err
		}
		defer fi.HIcon.DestroyIcon()

		if err := hImg.AddIcon(fi.HIcon); err != nil {
			return err
		}
	}

	return nil
}

// [ImageList_Destroy] function.
//
// [ImageList_Destroy]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_destroy
func (hImg HIMAGELIST) Destroy() error {
	// http://www.catch22.net/tuts/win32/system-image-list
	// https://www.autohotkey.com/docs/commands/ListView.htm
	ret, _, _ := syscall.SyscallN(_ImageList_Destroy.Addr(),
		uintptr(hImg))
	return util.ZeroAsSysInvalidParm(ret)
}

var _ImageList_Destroy = dll.Comctl32.NewProc("ImageList_Destroy")

// [ImageList_GetIconSize] function.
//
// [ImageList_GetIconSize]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_geticonsize
func (hImg HIMAGELIST) GetIconSize() (SIZE, error) {
	var sz SIZE
	ret, _, _ := syscall.SyscallN(_ImageList_GetIconSize.Addr(),
		uintptr(hImg),
		uintptr(unsafe.Pointer(&sz.Cx)), uintptr(unsafe.Pointer(&sz.Cy)))
	if ret == 0 {
		return SIZE{}, co.ERROR_INVALID_PARAMETER
	}
	return sz, nil
}

var _ImageList_GetIconSize = dll.Comctl32.NewProc("ImageList_GetIconSize")

// [ImageList_GetImageCount] function.
//
// [ImageList_GetImageCount]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_getimagecount
func (hImg HIMAGELIST) GetImageCount() uint {
	ret, _, _ := syscall.SyscallN(_ImageList_GetImageCount.Addr(),
		uintptr(hImg))
	return uint(ret)
}

var _ImageList_GetImageCount = dll.Comctl32.NewProc("ImageList_GetImageCount")

// [ImageList_ReplaceIcon] function.
//
// If icon was loaded from resource with [LoadIcon], it doesn't need to be
// destroyed, because all icon resources are automatically freed. Otherwise, if
// loaded with CreateIcon(), it must be destroyed.
//
// [ImageList_ReplaceIcon]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_replaceicon
// [LoadIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
func (hImg HIMAGELIST) ReplaceIcon(i int, hIcon HICON) (int, error) {
	ret, _, _ := syscall.SyscallN(_ImageList_ReplaceIcon.Addr(),
		uintptr(hImg), uintptr(i), uintptr(hIcon))
	if int32(ret) == -1 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return int(ret), nil
}

var _ImageList_ReplaceIcon = dll.Comctl32.NewProc("ImageList_ReplaceIcon")

// [ImageList_SetImageCount] function.
//
// [ImageList_SetImageCount]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_setimagecount
func (hImg HIMAGELIST) SetImageCount(count uint) error {
	ret, _, _ := syscall.SyscallN(_ImageList_SetImageCount.Addr(),
		uintptr(hImg), uintptr(count))
	return util.ZeroAsSysInvalidParm(ret)
}

var _ImageList_SetImageCount = dll.Comctl32.NewProc("ImageList_SetImageCount")
