//go:build windows

package proc

import (
	"syscall"
)

var (
	dwmapi = syscall.NewLazyDLL("dwmapi.dll")

	DwmEnableMMCSS                = dwmapi.NewProc("DwmEnableMMCSS")
	DwmExtendFrameIntoClientArea  = dwmapi.NewProc("DwmExtendFrameIntoClientArea")
	DwmFlush                      = dwmapi.NewProc("DwmFlush")
	DwmGetColorizationColor       = dwmapi.NewProc("DwmGetColorizationColor")
	DwmGetWindowAttribute         = dwmapi.NewProc("DwmGetWindowAttribute")
	DwmInvalidateIconicBitmaps    = dwmapi.NewProc("DwmInvalidateIconicBitmaps")
	DwmIsCompositionEnabled       = dwmapi.NewProc("DwmIsCompositionEnabled")
	DwmSetIconicLivePreviewBitmap = dwmapi.NewProc("DwmSetIconicLivePreviewBitmap")
	DwmSetIconicThumbnail         = dwmapi.NewProc("DwmSetIconicThumbnail")
	DwmSetWindowAttribute         = dwmapi.NewProc("DwmSetWindowAttribute")
)
