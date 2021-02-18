/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Manages a font resource.
type Font struct {
	hFont win.HFONT
}

// Simplified options to create a Font object.
type FontOpts struct {
	Name      string
	Size      int
	Bold      bool
	Italic    bool
	StrikeOut bool
	Underline bool
}

// Constructor.
//
// You must defer Destroy().
func (me *Font) Create(setup *FontOpts) *Font {
	me.Destroy()
	lf := win.LOGFONT{}
	lf.LfHeight = -(int32(setup.Size) + 3)

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

	return me.CreateLogFont(&lf)
}

// Constructor.
// Creates a font based on a LOGFONT struct.
//
// You must defer Destroy().
func (me *Font) CreateLogFont(lf *win.LOGFONT) *Font {
	me.Destroy()
	me.hFont = win.CreateFontIndirect(lf)
	return me
}

// Constructor.
// Creates a font identical to the current system font, usually Tahoma or Segoe
// UI. Because we call SetProcessDPIAware(), higher DPI resolutions will be
// reflected in the font size.
//
// You must defer Destroy().
func (me *Font) CreateUi() *Font {
	ncm := win.NONCLIENTMETRICS{}
	ncm.CbSize = uint32(unsafe.Sizeof(ncm))

	if !win.IsWindowsVistaOrGreater() {
		ncm.CbSize -= uint32(unsafe.Sizeof(ncm.IBorderWidth))
	}

	win.SystemParametersInfo(co.SPI_GETNONCLIENTMETRICS,
		ncm.CbSize, unsafe.Pointer(&ncm), 0)
	me.CreateLogFont(&ncm.LfMenuFont)
	return me
}

// Releases the font resource.
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

// Sends a WM_SETFONT message to the child control to apply the font.
func (me *Font) SetOnControl(ctrl Control) *Font {
	ctrl.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(me.hFont), 1)
	return me
}
