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

	CombineRgn                = dllGdi32.NewProc("CombineRgn")
	CreateCompatibleDC        = dllGdi32.NewProc("CreateCompatibleDC")
	CreateEllipticRgn         = dllGdi32.NewProc("CreateEllipticRgn")
	CreateEllipticRgnIndirect = dllGdi32.NewProc("CreateEllipticRgnIndirect")
	CreateFontIndirect        = dllGdi32.NewProc("CreateFontIndirectW")
	CreatePatternBrush        = dllGdi32.NewProc("CreatePatternBrush")
	CreatePolygonRgn          = dllGdi32.NewProc("CreatePolygonRgn")
	CreatePolyPolygonRgn      = dllGdi32.NewProc("CreatePolyPolygonRgn")
	CreateRectRgn             = dllGdi32.NewProc("CreateRectRgn")
	CreateRectRgnIndirect     = dllGdi32.NewProc("CreateRectRgnIndirect")
	CreateRoundRectRgn        = dllGdi32.NewProc("CreateRoundRectRgn")
	DeleteDC                  = dllGdi32.NewProc("DeleteDC")
	DeleteObject              = dllGdi32.NewProc("DeleteObject")
	EqualRgn                  = dllGdi32.NewProc("EqualRgn")
	FillRgn                   = dllGdi32.NewProc("FillRgn")
	FrameRgn                  = dllGdi32.NewProc("FrameRgn")
	GetDeviceCaps             = dllGdi32.NewProc("GetDeviceCaps")
	GetPolyFillMode           = dllGdi32.NewProc("GetPolyFillMode")
	GetTextExtentPoint32      = dllGdi32.NewProc("GetTextExtentPoint32W")
	GetTextFace               = dllGdi32.NewProc("GetTextFaceW")
	InvertRgn                 = dllGdi32.NewProc("InvertRgn")
	LineTo                    = dllGdi32.NewProc("LineTo")
	OffsetRgn                 = dllGdi32.NewProc("OffsetRgn")
	PaintRgn                  = dllGdi32.NewProc("PaintRgn")
	PolyDraw                  = dllGdi32.NewProc("PolyDraw")
	Polygon                   = dllGdi32.NewProc("Polygon")
	Polyline                  = dllGdi32.NewProc("Polyline")
	PolylineTo                = dllGdi32.NewProc("PolylineTo")
	PtInRegion                = dllGdi32.NewProc("PtInRegion")
	RectInRegion              = dllGdi32.NewProc("RectInRegion")
	RestoreDC                 = dllGdi32.NewProc("RestoreDC")
	SaveDC                    = dllGdi32.NewProc("SaveDC")
	SelectObject              = dllGdi32.NewProc("SelectObject")
	SetBkColor                = dllGdi32.NewProc("SetBkColor")
	SetBkMode                 = dllGdi32.NewProc("SetBkMode")
	SetPolyFillMode           = dllGdi32.NewProc("SetPolyFillMode")
	SetRectRgn                = dllGdi32.NewProc("SetRectRgn")
)
