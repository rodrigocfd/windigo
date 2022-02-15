package d2d1vt

import (
	"github.com/rodrigocfd/windigo/win"
)

// ID2D1Factory virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1factory
type ID2D1Factory struct {
	win.IUnknownVtbl
	ReloadSystemMetrics            uintptr
	GetDesktopDpi                  uintptr
	CreateRectangleGeometry        uintptr
	CreateRoundedRectangleGeometry uintptr
	CreateEllipseGeometry          uintptr
	CreateGeometryGroup            uintptr
	CreateTransformedGeometry      uintptr
	CreatePathGeometry             uintptr
	CreateStrokeStyle              uintptr
	CreateDrawingStateBlock        uintptr
	CreateWicBitmapRenderTarget    uintptr
	CreateHwndRenderTarget         uintptr
	CreateDxgiSurfaceRenderTarget  uintptr
	CreateDCRenderTarget           uintptr

	CreateRectangleGeometry2        uintptr
	CreateRoundedRectangleGeometry2 uintptr
	CreateEllipseGeometry2          uintptr
	CreateTransformedGeometry2      uintptr
	CreateStrokeStyle2              uintptr
	CreateDrawingStateBlock2        uintptr
	CreateDrawingStateBlock3        uintptr
	CreateWicBitmapRenderTarget2    uintptr
	CreateHwndRenderTarget2         uintptr
	CreateDxgiSurfaceRenderTarget2  uintptr
}

// ID2D1HwndRenderTarget virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1hwndrendertarget
type ID2D1HwndRenderTarget struct {
	ID2D1RenderTarget
	CheckWindowState uintptr
	Resize           uintptr
	GetHwnd          uintptr

	Resize2 uintptr
}

// ID2D1RenderTarget virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1rendertarget
type ID2D1RenderTarget struct {
	ID2D1Resource
	CreateBitmap                 uintptr
	CreateBitmapFromWicBitmap    uintptr
	CreateSharedBitmap           uintptr
	CreateBitmapBrush            uintptr
	CreateSolidColorBrush        uintptr
	CreateGradientStopCollection uintptr
	CreateLinearGradientBrush    uintptr
	CreateRadialGradientBrush    uintptr
	CreateCompatibleRenderTarget uintptr
	CreateLayer                  uintptr
	CreateMesh                   uintptr
	DrawLine                     uintptr
	DrawRectangle                uintptr
	FillRectangle                uintptr
	DrawRoundedRectangle         uintptr
	FillRoundedRectangle         uintptr
	DrawEllipse                  uintptr
	FillEllipse                  uintptr
	DrawGeometry                 uintptr
	FillGeometry                 uintptr
	FillMesh                     uintptr
	FillOpacityMask              uintptr
	DrawBitmap                   uintptr
	DrawText                     uintptr
	DrawTextLayout               uintptr
	DrawGlyphRun                 uintptr
	SetTransform                 uintptr
	GetTransform                 uintptr
	SetAntialiasMode             uintptr
	GetAntialiasMode             uintptr
	SetTextAntialiasMode         uintptr
	GetTextAntialiasMode         uintptr
	SetTextRenderingParams       uintptr
	GetTextRenderingParams       uintptr
	SetTags                      uintptr
	GetTags                      uintptr
	PushLayer                    uintptr
	PopLayer                     uintptr
	Flush                        uintptr
	SaveDrawingState             uintptr
	RestoreDrawingState          uintptr
	PushAxisAlignedClip          uintptr
	PopAxisAlignedClip           uintptr
	Clear                        uintptr
	BeginDraw                    uintptr
	EndDraw                      uintptr
	GetPixelFormat               uintptr
	SetDpi                       uintptr
	GetDpi                       uintptr
	GetSize                      uintptr
	GetPixelSize                 uintptr
	GetMaximumBitmapSize         uintptr
	IsSupported                  uintptr

	CreateBitmap2                 uintptr
	CreateBitmap3                 uintptr
	CreateBitmapFromWicBitmap2    uintptr
	CreateBitmapFromWicBitmap3    uintptr
	CreateBitmapBrush2            uintptr
	CreateBitmapBrush3            uintptr
	CreateBitmapBrush4            uintptr
	CreateSolidColorBrush2        uintptr
	CreateSolidColorBrush3        uintptr
	CreateGradientStopCollection2 uintptr
	CreateLinearGradientBrush2    uintptr
	CreateLinearGradientBrush3    uintptr
	CreateRadialGradientBrush2    uintptr
	CreateRadialGradientBrush3    uintptr
	CreateCompatibleRenderTarget2 uintptr
	CreateCompatibleRenderTarget3 uintptr
	CreateCompatibleRenderTarget4 uintptr
	CreateCompatibleRenderTarget5 uintptr
	CreateCompatibleRenderTarget6 uintptr
	CreateLayer2                  uintptr
	CreateLayer3                  uintptr
	DrawRectangle2                uintptr
	FillRectangle2                uintptr
	DrawRoundedRectangle2         uintptr
	FillRoundedRectangle2         uintptr
	DrawEllipse2                  uintptr
	FillEllipse2                  uintptr
	FillOpacityMask2              uintptr
	DrawBitmap2                   uintptr
	DrawBitmap3                   uintptr
	SetTransform2                 uintptr
	PushLayer2                    uintptr
	PushAxisAlignedClip2          uintptr
	Clear2                        uintptr
	DrawText2                     uintptr
	IsSupported2                  uintptr
}

// ID2D1Resource virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1resource
type ID2D1Resource struct {
	win.IUnknownVtbl
	GetFactory uintptr
}
