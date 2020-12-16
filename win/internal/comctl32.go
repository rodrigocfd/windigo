/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	comctl32Dll = syscall.NewLazyDLL("comctl32.dll")

	DefSubclassProc         = comctl32Dll.NewProc("DefSubclassProc")
	ImageList_Create        = comctl32Dll.NewProc("ImageList_Create")
	ImageList_Duplicate     = comctl32Dll.NewProc("ImageList_Duplicate")
	ImageList_GetIcon       = comctl32Dll.NewProc("ImageList_GetIcon")
	ImageList_GetIconSize   = comctl32Dll.NewProc("ImageList_GetIconSize")
	ImageList_GetImageCount = comctl32Dll.NewProc("ImageList_GetImageCount")
	ImageList_GetImageInfo  = comctl32Dll.NewProc("ImageList_GetImageInfo ")
	ImageList_Destroy       = comctl32Dll.NewProc("ImageList_Destroy")
	ImageList_ReplaceIcon   = comctl32Dll.NewProc("ImageList_ReplaceIcon")
	InitCommonControls      = comctl32Dll.NewProc("InitCommonControls")
	RemoveWindowSubclass    = comctl32Dll.NewProc("RemoveWindowSubclass")
	SetWindowSubclass       = comctl32Dll.NewProc("SetWindowSubclass")
)
