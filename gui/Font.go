/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Manages a font resource. Must be default-initialized, then call one of the
// creation methods on the object.
type Font struct {
	hFont win.HFONT
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

// Calls DeleteObject and sets the HFONT to zero.
func (me *Font) Destroy() {
	if me.hFont != 0 {
		me.hFont.DeleteObject()
		me.hFont = win.HFONT(0)
	}
}

// Returns the HFONT handle.
func (me *Font) Hfont() win.HFONT {
	return me.hFont
}

// Creates a font based on a FontSetup struct.
func (me *Font) Create(setup *FontSetup) *Font {
	me.Destroy()
	lf := win.LOGFONT{}
	lf.LfHeight = -(setup.Size + 3)

	if len(setup.Name) > len(lf.LfFaceName)-1 {
		panic(fmt.Sprintf("Font name can't be longer than %d chars.",
			len(lf.LfFaceName)-1))
	}
	copy(lf.LfFaceName[:], syscall.StringToUTF16(setup.Name))

	if setup.Bold {
		lf.LfWeight = co.FW_BOLD
	} else {
		lf.LfWeight = co.FW_DONTCARE
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

// Creates a font based on a LOGFONT struct.
func (me *Font) CreateFromLogFont(lf *win.LOGFONT) *Font {
	me.Destroy()
	me.hFont = lf.CreateFontIndirect()
	return me
}

// Creates a font identical to the current system font, usually Tahoma or Segoe
// UI. Because we call SetProcessDPIAware(), higher DPI resolutions will be
// reflected in the font size.
func (me *Font) CreateUi() *Font {
	ncm := win.NONCLIENTMETRICS{}
	ncm.CbSize = uint32(unsafe.Sizeof(ncm))

	if !win.IsWindowsVistaOrGreater() {
		ncm.CbSize -= uint32(unsafe.Sizeof(ncm.IBorderWidth))
	}

	win.SystemParametersInfo(co.SPI_GETNONCLIENTMETRICS,
		ncm.CbSize, unsafe.Pointer(&ncm), 0)
	me.CreateFromLogFont(&ncm.LfMenuFont)
	return me
}

// Sends a WM_SETFONT message to the child control to apply the font.
func (me *Font) SetOnControl(ctrl Window) *Font {
	ctrl.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(me.hFont), 1)
	return me
}
