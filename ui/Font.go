package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"winffi/api"
	c "winffi/consts"
)

var globalUiFont = NewFont() // managed in WindowMain's createWindow() and runMainLoop()

// Manages a font resource.
type Font struct {
	hFont api.HFONT
}

// Simplified options to create a Font object.
type FontSetup struct {
	Name      string
	Size      int32
	Bold      bool
	Italic    bool
	StrikeOut bool
	Underline bool
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

func (me *Font) Create(setup *FontSetup) *Font {
	me.Destroy()
	lf := api.LOGFONT{}
	lf.LfHeight = -(setup.Size + 3)

	if len(setup.Name) > len(lf.LfFaceName)-1 {
		panic(fmt.Sprintf("Font name can't be longer than %d chars.",
			len(lf.LfFaceName)-1))
	}
	copy(lf.LfFaceName[:], syscall.StringToUTF16(setup.Name))

	if setup.Bold {
		lf.LfWeight = c.FW_BOLD
	} else {
		lf.LfWeight = c.FW_DONTCARE
	}

	if setup.Italic {
		lf.LfItalic = 1
	}
	if setup.Underline {
		lf.LfUnderline = 1
	}
	if setup.StrikeOut {
		lf.LfStrikeOut = 1
	}

	return me.CreateFromLogFont(&lf)
}

func (me *Font) CreateFromLogFont(lf *api.LOGFONT) *Font {
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
	me.CreateFromLogFont(&ncm.LfMenuFont)
	return me
}

func (me *Font) SetOnControl(ctrl Window) *Font {
	ctrl.Hwnd().SendMessage(c.WM_SETFONT, api.WPARAM(me.hFont), 1)
	return me
}
