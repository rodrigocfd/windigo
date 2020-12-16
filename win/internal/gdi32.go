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
	gdi32Dll = syscall.NewLazyDLL("gdi32.dll")

	CombineRgn                = gdi32Dll.NewProc("CombineRgn")
	CreateCompatibleDC        = gdi32Dll.NewProc("CreateCompatibleDC")
	CreateEllipticRgn         = gdi32Dll.NewProc("CreateEllipticRgn")
	CreateEllipticRgnIndirect = gdi32Dll.NewProc("CreateEllipticRgnIndirect")
	CreateFontIndirect        = gdi32Dll.NewProc("CreateFontIndirectW")
	CreatePatternBrush        = gdi32Dll.NewProc("CreatePatternBrush")
	CreatePolygonRgn          = gdi32Dll.NewProc("CreatePolygonRgn")
	CreatePolyPolygonRgn      = gdi32Dll.NewProc("CreatePolyPolygonRgn")
	CreateRectRgn             = gdi32Dll.NewProc("CreateRectRgn")
	CreateRectRgnIndirect     = gdi32Dll.NewProc("CreateRectRgnIndirect")
	CreateRoundRectRgn        = gdi32Dll.NewProc("CreateRoundRectRgn")
	DeleteDC                  = gdi32Dll.NewProc("DeleteDC")
	DeleteObject              = gdi32Dll.NewProc("DeleteObject")
	EqualRgn                  = gdi32Dll.NewProc("EqualRgn")
	FillRgn                   = gdi32Dll.NewProc("FillRgn")
	FrameRgn                  = gdi32Dll.NewProc("FrameRgn")
	GetDeviceCaps             = gdi32Dll.NewProc("GetDeviceCaps")
	GetPolyFillMode           = gdi32Dll.NewProc("GetPolyFillMode")
	GetTextExtentPoint32      = gdi32Dll.NewProc("GetTextExtentPoint32W")
	GetTextFace               = gdi32Dll.NewProc("GetTextFaceW")
	InvertRgn                 = gdi32Dll.NewProc("InvertRgn")
	LineTo                    = gdi32Dll.NewProc("LineTo")
	OffsetRgn                 = gdi32Dll.NewProc("OffsetRgn")
	PaintRgn                  = gdi32Dll.NewProc("PaintRgn")
	PolyDraw                  = gdi32Dll.NewProc("PolyDraw")
	Polygon                   = gdi32Dll.NewProc("Polygon")
	Polyline                  = gdi32Dll.NewProc("Polyline")
	PolylineTo                = gdi32Dll.NewProc("PolylineTo")
	PtInRegion                = gdi32Dll.NewProc("PtInRegion")
	RectInRegion              = gdi32Dll.NewProc("RectInRegion")
	RestoreDC                 = gdi32Dll.NewProc("RestoreDC")
	SaveDC                    = gdi32Dll.NewProc("SaveDC")
	SelectObject              = gdi32Dll.NewProc("SelectObject")
	SetBkColor                = gdi32Dll.NewProc("SetBkColor")
	SetBkMode                 = gdi32Dll.NewProc("SetBkMode")
	SetPolyFillMode           = gdi32Dll.NewProc("SetPolyFillMode")
	SetRectRgn                = gdi32Dll.NewProc("SetRectRgn")
)
