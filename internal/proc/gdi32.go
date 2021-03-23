package proc

import (
	"syscall"
)

var (
	gdi32 = syscall.NewLazyDLL("gdi32.dll")

	CreateCompatibleDC   = gdi32.NewProc("CreateCompatibleDC")
	CreateFontIndirect   = gdi32.NewProc("CreateFontIndirectW")
	DeleteDC             = gdi32.NewProc("DeleteDC")
	DeleteObject         = gdi32.NewProc("DeleteObject")
	FillRgn              = gdi32.NewProc("FillRgn")
	FrameRgn             = gdi32.NewProc("FrameRgn")
	GetDeviceCaps        = gdi32.NewProc("GetDeviceCaps")
	GetPolyFillMode      = gdi32.NewProc("GetPolyFillMode")
	GetTextExtentPoint32 = gdi32.NewProc("GetTextExtentPoint32W")
	GetTextFace          = gdi32.NewProc("GetTextFaceW")
	GetTextMetrics       = gdi32.NewProc("GetTextMetricsW")
	InvertRgn            = gdi32.NewProc("InvertRgn")
	LineTo               = gdi32.NewProc("LineTo")
	PaintRgn             = gdi32.NewProc("PaintRgn")
	PolyDraw             = gdi32.NewProc("PolyDraw")
	Polygon              = gdi32.NewProc("Polygon")
	Polyline             = gdi32.NewProc("Polyline")
	PolylineTo           = gdi32.NewProc("PolylineTo")
	RestoreDC            = gdi32.NewProc("RestoreDC")
	SaveDC               = gdi32.NewProc("SaveDC")
	SelectObject         = gdi32.NewProc("SelectObject")
	SetBkColor           = gdi32.NewProc("SetBkColor")
	SetBkMode            = gdi32.NewProc("SetBkMode")
	SetPolyFillMode      = gdi32.NewProc("SetPolyFillMode")
	SetTextAlign         = gdi32.NewProc("SetTextAlign")
	TextOut              = gdi32.NewProc("TextOutW")
)
