//go:build windows

package win

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
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
// ⚠️ You must defer [HIMAGELIST.Destroy].
//
// # Example
//
//	hImg, _ := win.ImageListCreate(16, 16, co.ILC_COLOR32, 1, 1)
//	defer hImg.Destroy()
//
// [ImageList_Create]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_create
func ImageListCreate(cx, cy uint, flags co.ILC, szInitial, szGrow uint) (HIMAGELIST, error) {
	ret, _, _ := syscall.SyscallN(_ImageList_Create.Addr(),
		uintptr(int32(cx)),
		uintptr(int32(cy)),
		uintptr(flags),
		uintptr(int32(szInitial)),
		uintptr(int32(szGrow)))
	if ret == 0 {
		return HIMAGELIST(0), co.ERROR_INVALID_PARAMETER
	}
	return HIMAGELIST(ret), nil
}

var _ImageList_Create = dll.Comctl32.NewProc("ImageList_Create")

// [ImageList_Add] function.
//
// [ImageList_Add]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_add
func (hImg HIMAGELIST) Add(hbmp, hbmpMask HBITMAP) error {
	ret, _, _ := syscall.SyscallN(_ImageList_Add.Addr(),
		uintptr(hImg),
		uintptr(hbmp),
		uintptr(hbmpMask))
	return utl.Minus1AsSysInvalidParm(ret)
}

var _ImageList_Add = dll.Comctl32.NewProc("ImageList_Add")

// [ImageList_AddIcon] function.
//
// If icon was loaded from resource with [HINSTANCE.LoadIcon], it doesn't need
// to be destroyed, because all icon resources are automatically freed.
//
// [ImageList_AddIcon]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_addicon
func (hImg HIMAGELIST) AddIcon(hIcons ...HICON) error {
	for _, hIco := range hIcons {
		if err := hImg.ReplaceIcon(-1, hIco); err != nil {
			return err
		}
	}
	return nil
}

// Calls [HINSTANCE.LoadIcon] and [HIMAGELIST.AddIcon] to load an icon from the
// resource.
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

// Calls [SHGetFileInfo] and [HIMAGELIST.AddIcon] to load icons from the shell,
// used by Windows Explorer to represent the given file extensions, like "mp3".
//
// # Example
//
//	hImg, _ := win.ImageListCreate(16, 16, co.ILC_COLOR32, 1, 1)
//	defer hImg.Destroy()
//
//	hImg.AddIconFromShell("mp3", "wav")
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

// [ImageList_AddMasked] function.
//
// [ImageList_AddMasked]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_add
func (hImg HIMAGELIST) AddMasked(hbmp HBITMAP, mask COLORREF) error {
	ret, _, _ := syscall.SyscallN(_ImageList_AddMasked.Addr(),
		uintptr(hImg),
		uintptr(hbmp),
		uintptr(mask))
	return utl.Minus1AsSysInvalidParm(ret)
}

var _ImageList_AddMasked = dll.Comctl32.NewProc("ImageList_AddMasked")

// [ImageList_BeginDrag] function.
//
// [ImageList_BeginDrag]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_begindrag
func (hImg HIMAGELIST) BeginDrag(index, dxHotspot, dyHotspot int) error {
	ret, _, _ := syscall.SyscallN(_ImageList_BeginDrag.Addr(),
		uintptr(hImg),
		uintptr(int32(index)),
		uintptr(int32(dxHotspot)),
		uintptr(int32(dyHotspot)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_BeginDrag = dll.Comctl32.NewProc("ImageList_BeginDrag")

// [ImageList_Destroy] function.
//
// [ImageList_Destroy]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_destroy
func (hImg HIMAGELIST) Destroy() error {
	// http://www.catch22.net/tuts/win32/system-image-list
	// https://www.autohotkey.com/docs/commands/ListView.htm
	ret, _, _ := syscall.SyscallN(_ImageList_Destroy.Addr(),
		uintptr(hImg))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_Destroy = dll.Comctl32.NewProc("ImageList_Destroy")

// [ImageList_DrawEx] function.
//
// [ImageList_DrawEx]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_drawex
func (hImg HIMAGELIST) DrawEx(
	index int,
	hdcDest HDC,
	coords POINT,
	sz SIZE,
	bk, fg COLORREF,
	style co.ILD,
) error {
	ret, _, _ := syscall.SyscallN(_ImageList_DrawEx.Addr(),
		uintptr(hImg),
		uintptr(int32(index)),
		uintptr(hdcDest),
		uintptr(coords.X),
		uintptr(coords.Y),
		uintptr(sz.Cx),
		uintptr(sz.Cy),
		uintptr(bk),
		uintptr(fg),
		uintptr(style))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_DrawEx = dll.Comctl32.NewProc("ImageList_DrawEx")

// [ImageList_Duplicate] function.
//
// ⚠️ You must defer [HIMAGELIST.Destroy].
//
// # Example
//
//	hImg, _ := win.ImageListCreate(16, 16, co.ILC_COLOR32, 1, 1)
//	defer hImg.Destroy()
//
//	hImgCopy, _ := hImg.Duplicate()
//	defer hImgCopy.Destroy()
//
// [ImageList_Duplicate]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_duplicate
func (hImg HIMAGELIST) Duplicate() (HIMAGELIST, error) {
	ret, _, _ := syscall.SyscallN(_ImageList_Duplicate.Addr(),
		uintptr(hImg))
	if ret == 0 {
		return HIMAGELIST(0), co.ERROR_INVALID_PARAMETER
	}
	return HIMAGELIST(ret), nil
}

var _ImageList_Duplicate = dll.Comctl32.NewProc("ImageList_Duplicate")

// [ImageList_GetBkColor] function.
//
// [ImageList_GetBkColor]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_getbkcolor
func (hImg HIMAGELIST) GetBkColor() COLORREF {
	ret, _, _ := syscall.SyscallN(_ImageList_GetBkColor.Addr(),
		uintptr(hImg))
	return COLORREF(ret)
}

var _ImageList_GetBkColor = dll.Comctl32.NewProc("ImageList_GetBkColor")

// [ImageList_GetIconSize] function.
//
// [ImageList_GetIconSize]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_geticonsize
func (hImg HIMAGELIST) GetIconSize() (SIZE, error) {
	var sz SIZE
	ret, _, _ := syscall.SyscallN(_ImageList_GetIconSize.Addr(),
		uintptr(hImg),
		uintptr(unsafe.Pointer(&sz.Cx)),
		uintptr(unsafe.Pointer(&sz.Cy)))
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

// [ImageList_GetImageInfo] function.
//
// [ImageList_GetImageInfo]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_getimageinfo
func (hImg HIMAGELIST) GetImageInfo(index int) (IMAGEINFO, error) {
	var nfo IMAGEINFO
	ret, _, _ := syscall.SyscallN(_ImageList_GetImageInfo.Addr(),
		uintptr(hImg),
		uintptr(int32(index)),
		uintptr(unsafe.Pointer(&nfo)))
	if ret == 0 {
		return IMAGEINFO{}, co.ERROR_INVALID_PARAMETER
	}
	return nfo, nil
}

var _ImageList_GetImageInfo = dll.Comctl32.NewProc("ImageList_GetImageInfo")

// [ImageList_Remove] function.
//
// [ImageList_Remove]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_remove
func (hImg HIMAGELIST) Remove(index int) error {
	ret, _, _ := syscall.SyscallN(_ImageList_Remove.Addr(),
		uintptr(hImg),
		uintptr(int32(index)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_Remove = dll.Comctl32.NewProc("ImageList_Remove")

// [ImageList_RemoveAll] macro.
//
// [ImageList_RemoveAll]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_removeall
func (hImg HIMAGELIST) RemoveAll() error {
	return hImg.Remove(-1)
}

// [ImageList_ReplaceIcon] function.
//
// If icon was loaded from resource with [HINSTANCE.LoadIcon], it doesn't need
// to be destroyed, because all icon resources are automatically freed.
//
// [ImageList_ReplaceIcon]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_replaceicon
func (hImg HIMAGELIST) ReplaceIcon(index int, hIcon HICON) error {
	ret, _, _ := syscall.SyscallN(_ImageList_ReplaceIcon.Addr(),
		uintptr(hImg),
		uintptr(int32(index)),
		uintptr(hIcon))
	return utl.Minus1AsSysInvalidParm(ret)
}

var _ImageList_ReplaceIcon = dll.Comctl32.NewProc("ImageList_ReplaceIcon")

// [ImageList_SetDragCursorImage] function.
//
// [ImageList_SetDragCursorImage]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_setdragcursorimage
func (hImg HIMAGELIST) SetDragCursorImage(index int, dxHotspot, dyHotspot int) error {
	ret, _, _ := syscall.SyscallN(_ImageList_SetDragCursorImage.Addr(),
		uintptr(hImg),
		uintptr(int32(index)),
		uintptr(int32(dxHotspot)),
		uintptr(int32(dyHotspot)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_SetDragCursorImage = dll.Comctl32.NewProc("ImageList_SetDragCursorImage")

// [ImageList_SetIconSize] function.
//
// [ImageList_SetIconSize]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_seticonsize
func (hImg HIMAGELIST) SetIconSize(cx, cy int) error {
	ret, _, _ := syscall.SyscallN(_ImageList_SetIconSize.Addr(),
		uintptr(hImg),
		uintptr(int32(cx)),
		uintptr(int32(cy)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_SetIconSize = dll.Comctl32.NewProc("ImageList_SetIconSize")

// [ImageList_SetImageCount] function.
//
// [ImageList_SetImageCount]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_setimagecount
func (hImg HIMAGELIST) SetImageCount(count uint) error {
	ret, _, _ := syscall.SyscallN(_ImageList_SetImageCount.Addr(),
		uintptr(hImg),
		uintptr(uint32(count)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_SetImageCount = dll.Comctl32.NewProc("ImageList_SetImageCount")

// [ImageList_SetOverlayImage] function.
//
// [ImageList_SetOverlayImage]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_setoverlayimage
func (hImg HIMAGELIST) SetOverlayImage(index, overlayIndex int) error {
	ret, _, _ := syscall.SyscallN(_ImageList_SetOverlayImage.Addr(),
		uintptr(hImg),
		uintptr(int32(index)),
		uintptr(int32(overlayIndex)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_SetOverlayImage = dll.Comctl32.NewProc("ImageList_SetOverlayImage")
