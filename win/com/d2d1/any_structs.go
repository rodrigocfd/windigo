//go:build windows

package d2d1

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1co"
)

// [COLOR_F] struct;
//
// [COLOR_F]: https://learn.microsoft.com/en-us/windows/win32/Direct2D/d2d1-color-f
type COLOR_F struct {
	R float32
	G float32
	B float32
	A float32
}

// [FACTORY_OPTIONS] struct.
//
// [FACTORY_OPTIONS]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_factory_options
type FACTORY_OPTIONS struct {
	DebugLevel d2d1co.DEBUG_LEVEL
}

// [HWND_RENDER_TARGET_PROPERTIES] struct.
//
// [HWND_RENDER_TARGET_PROPERTIES]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_hwnd_render_target_properties
type HWND_RENDER_TARGET_PROPERTIES struct {
	Hwnd           win.HWND
	PixelSize      SIZE_U
	PresentOptions d2d1co.PRESENT_OPTIONS
}

// [PIXEL_FORMAT] struct.
//
// [PIXEL_FORMAT]: https://learn.microsoft.com/en-us/windows/win32/api/dcommon/ns-dcommon-d2d1_pixel_format
type PIXEL_FORMAT struct {
	Format    d2d1co.DXGI_FORMAT
	AlphaMode d2d1co.ALPHA_MODE
}

// [RENDER_TARGET_PROPERTIES] struct.
//
// [RENDER_TARGET_PROPERTIES]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_render_target_properties
type RENDER_TARGET_PROPERTIES struct {
	Type        d2d1co.RENDER_TARGET_TYPE
	PixelFormat PIXEL_FORMAT
	DpiX        float32
	DpiY        float32
	Usage       d2d1co.RENDER_TARGET_USAGE
	MinLevel    d2d1co.FEATURE_LEVEL
}

// [SIZE_F] struct.
//
// [SIZE_F]: https://learn.microsoft.com/en-us/windows/win32/direct2d/d2d1-size-f
type SIZE_F struct {
	Width, Height float32
}

// [SIZE_U] struct.
//
// [SIZE_U]: https://learn.microsoft.com/en-us/windows/win32/direct2d/d2d1-size-u
type SIZE_U struct {
	Width, Height uint32
}
