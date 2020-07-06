/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	dllComCtl32 = syscall.NewLazyDLL("comctl32.dll")

	DefSubclassProc         = dllComCtl32.NewProc("DefSubclassProc")
	ImageList_Create        = dllComCtl32.NewProc("ImageList_Create")
	ImageList_Duplicate     = dllComCtl32.NewProc("ImageList_Duplicate")
	ImageList_GetIcon       = dllComCtl32.NewProc("ImageList_GetIcon")
	ImageList_GetIconSize   = dllComCtl32.NewProc("ImageList_GetIconSize")
	ImageList_GetImageCount = dllComCtl32.NewProc("ImageList_GetImageCount")
	ImageList_GetImageInfo  = dllComCtl32.NewProc("ImageList_GetImageInfo ")
	ImageList_Destroy       = dllComCtl32.NewProc("ImageList_Destroy")
	ImageList_ReplaceIcon   = dllComCtl32.NewProc("ImageList_ReplaceIcon")
	InitCommonControls      = dllComCtl32.NewProc("InitCommonControls")
	RemoveWindowSubclass    = dllComCtl32.NewProc("RemoveWindowSubclass")
	SetWindowSubclass       = dllComCtl32.NewProc("SetWindowSubclass")
)
