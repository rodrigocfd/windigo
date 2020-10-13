/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"strings"
	"time"
	"windigo/co"
	"windigo/win"
)

type _UiT struct {
	globalDpi win.POINT
}

// Internal ui package functions.
var _Ui _UiT

// Syntactic sugar; converts bool to 0 or 1.
func (_UiT) BoolToUint32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func (_UiT) BoolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// Multiplies position and size by current DPI factor.
func (me *_UiT) MultiplyDpi(pos *Pos, size *Size) {
	if me.globalDpi.X == 0 { // not initialized yet?
		dc := win.HWND(0).GetDC()
		me.globalDpi.X = dc.GetDeviceCaps(co.GDC_LOGPIXELSX) // cache
		me.globalDpi.Y = dc.GetDeviceCaps(co.GDC_LOGPIXELSY)
		win.HWND(0).ReleaseDC(dc)
	}

	if pos != nil {
		pos.X = int(win.MulDiv(int32(pos.X), me.globalDpi.X, 96))
		pos.Y = int(win.MulDiv(int32(pos.Y), me.globalDpi.Y, 96))
	}
	if size != nil {
		size.Cx = uint(win.MulDiv(int32(size.Cx), me.globalDpi.X, 96))
		size.Cy = uint(win.MulDiv(int32(size.Cy), me.globalDpi.Y, 96))
	}
}

// Calculates the bound rectangle to fit the text with current system font.
func (_UiT) CalcTextBoundBox(
	hReferenceDc win.HWND, text string, considerAccelerators bool) Size {

	isTextEmpty := false
	if len(text) == 0 {
		isTextEmpty = true
		text = "Pj" // just a placeholder to get the text height
	}

	if considerAccelerators {
		text = _Ui.RemoveAccelAmpersands(text)
	}

	parentDc := hReferenceDc.GetDC()
	defer hReferenceDc.ReleaseDC(parentDc)

	cloneDc := parentDc.CreateCompatibleDC()
	defer cloneDc.DeleteDC()

	prevFont := cloneDc.SelectObjectFont(_globalUiFont.Hfont()) // system font; already adjusted to current DPI
	defer cloneDc.SelectObjectFont(prevFont)

	bounds := cloneDc.GetTextExtentPoint32(text)

	if isTextEmpty {
		bounds.Cx = 0 // if no text was given, return just the height
	}
	return Size{Cx: uint(bounds.Cx), Cy: uint(bounds.Cy)}
}

// "&He && she" becomes "He & she".
func (_UiT) RemoveAccelAmpersands(text string) string {
	runes := []rune(text)
	buf := strings.Builder{}
	buf.Grow(len(text)) // prealloc for performance

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == '&' && runes[i+1] != '&' {
			continue
		}
		buf.WriteRune(runes[i])
	}
	if runes[len(runes)-1] != '&' {
		buf.WriteRune(runes[len(runes)-1])
	}
	return buf.String()
}

// Converts current timezone SYSTEMTIME into time.Time, millisecond precision.
func (_UiT) SystemtimeToTime(stLocalTime *win.SYSTEMTIME) time.Time {
	return time.Date(int(stLocalTime.WYear),
		time.Month(stLocalTime.WMonth), int(stLocalTime.WDay),
		int(stLocalTime.WHour), int(stLocalTime.WMinute), int(stLocalTime.WSecond),
		int(stLocalTime.WMilliseconds)*1_000_000,
		time.Local)
}

// Converts time.Time into current timezone SYSTEMTIME, millisecond precision.
func (_UiT) TimeToSystemtime(t time.Time, stLocalTime *win.SYSTEMTIME) {
	// https://support.microsoft.com/en-ca/help/167296/how-to-convert-a-unix-time-t-to-a-win32-filetime-or-systemtime
	epoch := t.UnixNano()/100 + 116_444_736_000_000_000

	ft := win.FILETIME{}
	ft.DwLowDateTime = uint32(epoch & 0xffff_ffff)
	ft.DwHighDateTime = uint32(epoch >> 32)

	stUtc := win.SYSTEMTIME{}
	win.FileTimeToSystemTime(&ft, &stUtc)
	win.SystemTimeToTzSpecificLocalTime(nil, &stUtc, stLocalTime)
}
