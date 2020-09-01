/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Raw window message parameters.
type Wm struct {
	_Wm
}

type _Wm struct {
	WParam win.WPARAM
	LParam win.LPARAM
}

type _WmButton struct{ _Wm }

func (p _WmButton) HasCtrl() bool      { return (co.MK(p.WParam) & co.MK_CONTROL) != 0 }
func (p _WmButton) HasLeftBtn() bool   { return (co.MK(p.WParam) & co.MK_LBUTTON) != 0 }
func (p _WmButton) HasMiddleBtn() bool { return (co.MK(p.WParam) & co.MK_MBUTTON) != 0 }
func (p _WmButton) HasRightBtn() bool  { return (co.MK(p.WParam) & co.MK_RBUTTON) != 0 }
func (p _WmButton) HasShift() bool     { return (co.MK(p.WParam) & co.MK_SHIFT) != 0 }
func (p _WmButton) HasXBtn1() bool     { return (co.MK(p.WParam) & co.MK_XBUTTON1) != 0 }
func (p _WmButton) HasXBtn2() bool     { return (co.MK(p.WParam) & co.MK_XBUTTON2) != 0 }
func (p _WmButton) Pos() win.POINT     { return p.LParam.MakePoint() }

type _WmChar struct{ _Wm }

func (p _WmChar) CharCode() uint16          { return uint16(p.WParam) }
func (p _WmChar) RepeatCount() uint16       { return p.LParam.LoWord() }
func (p _WmChar) ScanCode() uint8           { return p.LParam.LoByteHiWord() }
func (p _WmChar) IsExtendedKey() bool       { return (p.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p _WmChar) HasAltKey() bool           { return (p.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p _WmChar) IsKeyDownBeforeSend() bool { return (p.LParam.HiByteHiWord() & 0b0100_0000) != 0 }
func (p _WmChar) IsKeyBeingReleased() bool  { return (p.LParam.HiByteHiWord() & 0b1000_0000) != 0 }

type _WmCtlColor struct{ _Wm }

func (p _WmCtlColor) Hdc() win.HDC       { return win.HDC(p.WParam) }
func (p _WmCtlColor) HControl() win.HWND { return win.HWND(p.LParam) }

type _WmKey struct{ _Wm }

func (p _WmKey) VirtualKeyCode() co.VK     { return co.VK(p.WParam) }
func (p _WmKey) RepeatCount() uint16       { return p.LParam.LoWord() }
func (p _WmKey) ScanCode() uint8           { return p.LParam.LoByteHiWord() }
func (p _WmKey) IsExtendedKey() bool       { return (p.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p _WmKey) HasAltKey() bool           { return (p.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p _WmKey) IsKeyDownBeforeSend() bool { return (p.LParam.HiByteHiWord() & 0b0100_0000) != 0 }

type _WmScroll struct{ _Wm }

func (p _WmScroll) ScrollBoxPos() uint16    { return p.WParam.HiWord() }
func (p _WmScroll) Request() co.SBR         { return co.SBR(p.WParam.LoWord()) }
func (p _WmScroll) HwndScrollbar() win.HWND { return win.HWND(p.LParam) }

type WmNotify struct{ _Wm }

func (p WmNotify) NmHdr() *win.NMHDR { return (*win.NMHDR)(unsafe.Pointer(p.LParam)) }
