//go:build windows

package d2d1vt

import (
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
)

// [ID2D1Factory] virtual table.
//
// [ID2D1Factory]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1factory
type ID2D1Factory struct {
	comvt.IUnknown
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
}

// [ID2D1HwndRenderTarget] virtual table.
//
// [ID2D1HwndRenderTarget]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1hwndrendertarget
type ID2D1HwndRenderTarget struct {
	ID2D1RenderTarget
	CheckWindowState uintptr
	Resize           uintptr
	GetHwnd          uintptr
}

// [ID2D1RenderTarget] virtual table.
//
// [ID2D1RenderTarget]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1rendertarget
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
}

// [ID2D1Resource] virtual table.
//
// [ID2D1Resource]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1resource
type ID2D1Resource struct {
	comvt.IUnknown
	GetFactory uintptr
}
