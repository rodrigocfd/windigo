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
	dllGdi32 = syscall.NewLazyDLL("gdi32.dll")

	CreateFontIndirect   = dllGdi32.NewProc("CreateFontIndirectW")
	CreatePatternBrush   = dllGdi32.NewProc("CreatePatternBrush")
	DeleteObject         = dllGdi32.NewProc("DeleteObject")
	GetDeviceCaps        = dllGdi32.NewProc("GetDeviceCaps")
	GetTextExtentPoint32 = dllGdi32.NewProc("GetTextExtentPoint32W")
	GetTextFace          = dllGdi32.NewProc("GetTextFaceW")
	LineTo               = dllGdi32.NewProc("LineTo")
	SelectObject         = dllGdi32.NewProc("SelectObject")
	SetBkColor           = dllGdi32.NewProc("SetBkColor")
	SetBkMode            = dllGdi32.NewProc("SetBkMode")
)
