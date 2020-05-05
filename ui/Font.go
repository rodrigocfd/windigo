package ui

import (
	"unsafe"
	"winffi/api"
	c "winffi/consts"
)

var globalUiFont = NewFont() // managed in WindowMain's createWindow() and runMainLoop()

// Manages a font resource.
type Font struct {
	hFont api.HFONT
}

func NewFont() *Font {
	return &Font{
		hFont: api.HFONT(0),
	}
}

func (f *Font) Destroy() {
	if f.hFont != 0 {
		f.hFont.DeleteObject()
		f.hFont = api.HFONT(0)
	}
}

func (f *Font) Hfont() api.HFONT {
	return f.hFont
}

func (f *Font) CreateFontLogFont(lf *api.LOGFONT) *Font {
	f.Destroy()
	f.hFont = lf.CreateFontIndirect()
	return f
}

func (f *Font) CreateUi() *Font {
	ncm := api.NONCLIENTMETRICS{}
	ncm.Size = uint32(unsafe.Sizeof(ncm))

	if !api.IsWindowsVistaOrGreater() {
		ncm.Size -= uint32(unsafe.Sizeof(ncm.BorderWidth))
	}

	api.SystemParametersInfo(c.SPI_GETNONCLIENTMETRICS,
		ncm.Size, unsafe.Pointer(&ncm), 0)
	f.CreateFontLogFont(&ncm.MenuFont)
	return f
}

func (f *Font) SetOnControl(ctrl Window) *Font {
	ctrl.Hwnd().SendMessage(c.WM_SETFONT, api.WPARAM(f.hFont), 1)
	return f
}
