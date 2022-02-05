package proc

import (
	"syscall"
)

var (
	dwmapi = syscall.NewLazyDLL("dwmapi.dll")

	DwmExtendFrameIntoClientArea  = dwmapi.NewProc("DwmExtendFrameIntoClientArea")
	DwmGetColorizationColor       = dwmapi.NewProc("DwmGetColorizationColor")
	DwmGetWindowAttribute         = dwmapi.NewProc("DwmGetWindowAttribute")
	DwmIsCompositionEnabled       = dwmapi.NewProc("DwmIsCompositionEnabled")
	DwmSetIconicLivePreviewBitmap = dwmapi.NewProc("DwmSetIconicLivePreviewBitmap")
	DwmSetIconicThumbnail         = dwmapi.NewProc("DwmSetIconicThumbnail")
	DwmSetWindowAttribute         = dwmapi.NewProc("DwmSetWindowAttribute")
)
