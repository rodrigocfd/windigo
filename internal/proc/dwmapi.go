package proc

import (
	"syscall"
)

var (
	dwmapi = syscall.NewLazyDLL("dwmapi.dll")

	DwmExtendFrameIntoClientArea  = dwmapi.NewProc("DwmExtendFrameIntoClientArea")
	DwmGetColorizationColor       = dwmapi.NewProc("DwmGetColorizationColor")
	DwmIsCompositionEnabled       = dwmapi.NewProc("DwmIsCompositionEnabled")
	DwmSetIconicLivePreviewBitmap = dwmapi.NewProc("DwmSetIconicLivePreviewBitmap")
	DwmSetIconicThumbnail         = dwmapi.NewProc("DwmSetIconicThumbnail")
)
