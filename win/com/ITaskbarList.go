/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

var Guid_ITaskbarList = makeGuid(0x56FDF344, 0xFD6D, 0x11d0, 0x958A006097C9A090)
var Guid_ITaskbarList3 = makeGuid(0xea1afb91, 0x9e28, 0x4b86, 0x90e99e9f8a5eefaf)

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

type iTaskbarList2 struct {
	lpVtbl *iTaskbarList2Vtbl
}

type iTaskbarList2Vtbl struct {
	iTaskbarListVtbl
	MarkFullscreenWindow uintptr
}

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
