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
	dllGdi32 = syscall.NewLazyDLL("gdi32.dll")

	CreateCompatibleDC   = dllGdi32.NewProc("CreateCompatibleDC")
	CreateFontIndirect   = dllGdi32.NewProc("CreateFontIndirectW")
	CreatePatternBrush   = dllGdi32.NewProc("CreatePatternBrush")
	DeleteDC             = dllGdi32.NewProc("DeleteDC")
	DeleteObject         = dllGdi32.NewProc("DeleteObject")
	GetDeviceCaps        = dllGdi32.NewProc("GetDeviceCaps")
	GetTextExtentPoint32 = dllGdi32.NewProc("GetTextExtentPoint32W")
	GetTextFace          = dllGdi32.NewProc("GetTextFaceW")
	LineTo               = dllGdi32.NewProc("LineTo")
	PolyDraw             = dllGdi32.NewProc("PolyDraw")
	Polygon              = dllGdi32.NewProc("Polygon")
	Polyline             = dllGdi32.NewProc("Polyline")
	PolylineTo           = dllGdi32.NewProc("PolylineTo")
	RestoreDC            = dllGdi32.NewProc("RestoreDC")
	SaveDC               = dllGdi32.NewProc("SaveDC")
	SelectObject         = dllGdi32.NewProc("SelectObject")
	SetBkColor           = dllGdi32.NewProc("SetBkColor")
	SetBkMode            = dllGdi32.NewProc("SetBkMode")
)
