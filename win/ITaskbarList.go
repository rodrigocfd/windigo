/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

type iTaskbarList struct {
	lpVtbl *iTaskbarListVtbl
}

type iTaskbarListVtbl struct {
	iUnknownVtbl
	HrInit       uintptr
	AddTab       uintptr
	DeleteTab    uintptr
	ActivateTab  uintptr
	SetActiveAlt uintptr
}

//------------------------------------------------------------------------------

type iTaskbarList2 struct {
	lpVtbl *iTaskbarList2Vtbl
}

type iTaskbarList2Vtbl struct {
	iTaskbarListVtbl
	MarkFullscreenWindow uintptr
}

//------------------------------------------------------------------------------

type ITaskbarList3 struct {
	lpVtbl *iTaskbarList3Vtbl
}

type iTaskbarList3Vtbl struct {
	iTaskbarList2Vtbl
	SetProgressValue      uintptr
	SetProgressState      uintptr
	RegisterTab           uintptr
	UnregisterTab         uintptr
	SetTabOrder           uintptr
	SetTabActive          uintptr
	ThumbBarAddButtons    uintptr
	ThumbBarUpdateButtons uintptr
	ThumbBarSetImageList  uintptr
	SetOverlayIcon        uintptr
	SetThumbnailTooltip   uintptr
	SetThumbnailClip      uintptr
}
