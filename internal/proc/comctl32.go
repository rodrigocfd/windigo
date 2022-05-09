//go:build windows

package proc

import (
	"syscall"
)

var (
	comctl32 = syscall.NewLazyDLL("comctl32.dll")

	DefSubclassProc         = comctl32.NewProc("DefSubclassProc")
	ImageList_Create        = comctl32.NewProc("ImageList_Create")
	ImageList_Destroy       = comctl32.NewProc("ImageList_Destroy")
	ImageList_GetIconSize   = comctl32.NewProc("ImageList_GetIconSize")
	ImageList_GetImageCount = comctl32.NewProc("ImageList_GetImageCount")
	ImageList_ReplaceIcon   = comctl32.NewProc("ImageList_ReplaceIcon")
	InitCommonControls      = comctl32.NewProc("InitCommonControls")
	InitCommonControlsEx    = comctl32.NewProc("InitCommonControlsEx")
	RemoveWindowSubclass    = comctl32.NewProc("RemoveWindowSubclass")
	SetWindowSubclass       = comctl32.NewProc("SetWindowSubclass")
	TaskDialog              = comctl32.NewProc("TaskDialog")
	TaskDialogIndirect      = comctl32.NewProc("TaskDialogIndirect")
)
