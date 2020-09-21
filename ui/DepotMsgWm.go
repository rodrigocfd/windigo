/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/co"
	"windigo/win"
)

// Raw window message parameters, WPARAM and LPARAM.
type Wm struct {
	WParam win.WPARAM
	LParam win.LPARAM
}

type WmChar struct{ m Wm }

func (p WmChar) CharCode() rune            { return rune(p.m.WParam) }
func (p WmChar) RepeatCount() uint         { return uint(p.m.LParam.LoWord()) }
func (p WmChar) ScanCode() uint            { return uint(p.m.LParam.LoByteHiWord()) }
func (p WmChar) IsExtendedKey() bool       { return (p.m.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p WmChar) HasAltKey() bool           { return (p.m.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p WmChar) IsKeyDownBeforeSend() bool { return (p.m.LParam.HiByteHiWord() & 0b0100_0000) != 0 }
func (p WmChar) IsKeyBeingReleased() bool  { return (p.m.LParam.HiByteHiWord() & 0b1000_0000) != 0 }

type WmCtlColor struct{ m Wm }

func (p WmCtlColor) Hdc() win.HDC       { return win.HDC(p.m.WParam) }
func (p WmCtlColor) HControl() win.HWND { return win.HWND(p.m.LParam) }

type WmKey struct{ m Wm }

func (p WmKey) VirtualKeyCode() co.VK     { return co.VK(p.m.WParam) }
func (p WmKey) RepeatCount() uint         { return uint(p.m.LParam.LoWord()) }
func (p WmKey) ScanCode() uint            { return uint(p.m.LParam.LoByteHiWord()) }
func (p WmKey) IsExtendedKey() bool       { return (p.m.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p WmKey) HasAltKey() bool           { return (p.m.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p WmKey) IsKeyDownBeforeSend() bool { return (p.m.LParam.HiByteHiWord() & 0b0100_0000) != 0 }

type WmMenu struct{ m Wm }

func (p WmMenu) ItemIndex() uint  { return uint(p.m.WParam) }
func (p WmMenu) Hmenu() win.HMENU { return win.HMENU(p.m.LParam) }

type WmMouse struct{ m Wm }

func (p WmMouse) VirtualKeys() co.MK { return co.MK(p.m.WParam.LoWord()) }
func (p WmMouse) HasCtrl() bool      { return (p.VirtualKeys() & co.MK_CONTROL) != 0 }
func (p WmMouse) HasShift() bool     { return (p.VirtualKeys() & co.MK_SHIFT) != 0 }
func (p WmMouse) IsLeftBtn() bool    { return (p.VirtualKeys() & co.MK_LBUTTON) != 0 }
func (p WmMouse) IsMiddleBtn() bool  { return (p.VirtualKeys() & co.MK_MBUTTON) != 0 }
func (p WmMouse) IsRightBtn() bool   { return (p.VirtualKeys() & co.MK_RBUTTON) != 0 }
func (p WmMouse) IsXBtn1() bool      { return (p.VirtualKeys() & co.MK_XBUTTON1) != 0 }
func (p WmMouse) IsXBtn2() bool      { return (p.VirtualKeys() & co.MK_XBUTTON2) != 0 }
func (p WmMouse) Pos() win.POINT     { return p.m.LParam.MakePoint() }

type WmNcMouse struct{ m Wm }

func (p WmNcMouse) HitTest() co.HT { return co.HT(p.m.WParam) }
func (p WmNcMouse) Pos() win.POINT { return p.m.LParam.MakePoint() }

type WmScroll struct{ m Wm }

func (p WmScroll) ScrollBoxPos() uint      { return uint(p.m.WParam.HiWord()) }
func (p WmScroll) Request() co.SBR         { return co.SBR(p.m.WParam.LoWord()) }
func (p WmScroll) HwndScrollbar() win.HWND { return win.HWND(p.m.LParam) }
