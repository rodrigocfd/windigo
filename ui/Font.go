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

func (me *Font) Destroy() {
	if me.hFont != 0 {
		me.hFont.DeleteObject()
		me.hFont = api.HFONT(0)
	}
}

func (me *Font) Hfont() api.HFONT {
	return me.hFont
}

func (me *Font) CreateFontLogFont(lf *api.LOGFONT) *Font {
	me.Destroy()
	me.hFont = lf.CreateFontIndirect()
	return me
}

func (me *Font) CreateUi() *Font {
	ncm := api.NONCLIENTMETRICS{}
	ncm.CbSize = uint32(unsafe.Sizeof(ncm))

	if !api.IsWindowsVistaOrGreater() {
		ncm.CbSize -= uint32(unsafe.Sizeof(ncm.IBorderWidth))
	}

	api.SystemParametersInfo(c.SPI_GETNONCLIENTMETRICS,
		ncm.CbSize, unsafe.Pointer(&ncm), 0)
	me.CreateFontLogFont(&ncm.LfMenuFont)
	return me
}

func (me *Font) SetOnControl(ctrl Window) *Font {
	ctrl.Hwnd().SendMessage(c.WM_SETFONT, api.WPARAM(me.hFont), 1)
	return me
}
