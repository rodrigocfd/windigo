/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"syscall"
	"time"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Global private variables and methods.
type _GlobalT struct {
	uiFont     Font
	dpi        win.POINT
	autoCtrlId int
}

var _global _GlobalT

// Syntactic sugar; converts bool to 0 or 1.
func (*_GlobalT) BoolToUint32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func (*_GlobalT) BoolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// Calculates the bound rectangle to fit the text with current system font.
func (*_GlobalT) CalcTextBoundBox(
	hReferenceDc win.HWND, text string, considerAccelerators bool) Size {

	isTextEmpty := false
	if len(text) == 0 {
		isTextEmpty = true
		text = "Pj" // just a placeholder to get the text height
	}

	if considerAccelerators {
		text = _global.RemoveAccelAmpersands(text)
	}

	parentDc := hReferenceDc.GetDC()
	defer hReferenceDc.ReleaseDC(parentDc)

	cloneDc := parentDc.CreateCompatibleDC()
	defer cloneDc.DeleteDC()

	prevFont := cloneDc.SelectObjectFont(_global.UiFont().Hfont()) // system font; already adjusted to current DPI
	defer cloneDc.SelectObjectFont(prevFont)

	bounds := cloneDc.GetTextExtentPoint32(text)

	if isTextEmpty {
		bounds.Cx = 0 // if no text was given, return just the height
	}
	return Size{Cx: int(bounds.Cx), Cy: int(bounds.Cy)}
}

// Returns the global UI font, creates if not yet.
func (me *_GlobalT) UiFont() *Font {
	if me.uiFont.Hfont() == win.HFONT(0) { // not initialized yet?
		me.uiFont.CreateUi()
	}
	return &me.uiFont
}

// Returns a WNDCLASSEX structure filled with the given parameters, and the
// class name, which is auto-generated if not specified.
func (*_GlobalT) GenerateWndclassex(
	hInst win.HINSTANCE, className string, classStyles co.CS,
	hCursor win.HCURSOR, hBrushBg win.HBRUSH,
	defBrushBgColor co.COLOR, iconId int) (*win.WNDCLASSEX, string) {

	wcx := win.WNDCLASSEX{}
	wcx.CbSize = uint32(unsafe.Sizeof(wcx))
	wcx.LpfnWndProc = syscall.NewCallback(_globalWndProc)
	wcx.HInstance = hInst
	wcx.Style = classStyles

	if hCursor != 0 {
		wcx.HCursor = hCursor
	} else {
		wcx.HCursor = win.HINSTANCE(0).LoadCursor(co.IDC_ARROW) // default cursor
	}

	if hBrushBg != 0 {
		wcx.HbrBackground = hBrushBg
	} else {
		wcx.HbrBackground = win.CreateSysColorBrush(defBrushBgColor)
	}

	if iconId != 0 {
		// If an icon ID was specified, load it from the resource.
		// Resource icons are automatically released by the system.
		wcx.HIcon = win.HICON(
			hInst.LoadImage(int32(iconId), co.IMAGE_ICON,
				32, 32, co.LR_DEFAULTCOLOR))
		wcx.HIconSm = win.HICON(
			hInst.LoadImage(int32(iconId), co.IMAGE_ICON,
				16, 16, co.LR_DEFAULTCOLOR))
	}

	// After all the fields are set, if no class name, we generate one by hashing
	// all WNDCLASSEX fields. That's why it must be the last thing to be done.
	classNameStr := ""
	if className == "" {
		classNameStr = fmt.Sprintf("%x.%x.%x.%x.%x.%x.%x.%x.%x.%x",
			wcx.Style, wcx.LpfnWndProc, wcx.CbClsExtra, wcx.CbWndExtra,
			wcx.HInstance, wcx.HIcon, wcx.HCursor, wcx.HbrBackground,
			wcx.LpszMenuName, wcx.HIconSm)

		classNameSlice := win.Str.ToUint16Slice(classNameStr)
		wcx.LpszClassName = &classNameSlice[0]
	}

	return &wcx, classNameStr
}

// Multiplies position and size by current DPI factor.
func (me *_GlobalT) MultiplyDpi(pos *Pos, size *Size) {
	if me.dpi.X == 0 { // not initialized yet?
		dc := win.HWND(0).GetDC()
		me.dpi.X = dc.GetDeviceCaps(co.GDC_LOGPIXELSX) // cache
		me.dpi.Y = dc.GetDeviceCaps(co.GDC_LOGPIXELSY)
		win.HWND(0).ReleaseDC(dc)
	}

	if pos != nil {
		pos.X = int(win.MulDiv(int32(pos.X), me.dpi.X, 96))
		pos.Y = int(win.MulDiv(int32(pos.Y), me.dpi.Y, 96))
	}
	if size != nil {
		size.Cx = int(win.MulDiv(int32(size.Cx), me.dpi.X, 96))
		size.Cy = int(win.MulDiv(int32(size.Cy), me.dpi.Y, 96))
	}
}

// Returns a new unique auto-generated control ID.
func (me *_GlobalT) NewAutoCtrlId() int {
	if me.autoCtrlId == 0 { // not initialized yet?
		me.autoCtrlId = 20_000 // in-between Visual Studio Resource Editor values
	}
	me.autoCtrlId++
	return me.autoCtrlId
}

// "&He && she" becomes "He & she".
func (*_GlobalT) RemoveAccelAmpersands(text string) string {
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
func (*_GlobalT) SystemtimeToTime(stLocalTime *win.SYSTEMTIME) time.Time {
	return time.Date(int(stLocalTime.WYear),
		time.Month(stLocalTime.WMonth), int(stLocalTime.WDay),
		int(stLocalTime.WHour), int(stLocalTime.WMinute), int(stLocalTime.WSecond),
		int(stLocalTime.WMilliseconds)*1_000_000,
		time.Local)
}

// Displays the panic message, if any.
func (*_GlobalT) TreatPanic(r interface{}) {
	var msg string

	switch v := r.(type) {
	case *win.WinError:
		msg = fmt.Sprintf("A panic has occurred, Win32 error:\n\n%s", v.Error())
	case error:
		msg = fmt.Sprintf("A panic has occurred, error:\n\n%s", v.Error())
	case string:
		msg = fmt.Sprintf("A panic occurred:\n\n%s", v)
	default:
		msg = "An unspecified panic has occurred."
	}
	msg += "\n\n" + string(debug.Stack())

	fmt.Fprintln(os.Stderr, msg)
	win.HWND(0).MessageBox(msg, "Panic", co.MB_ICONERROR)
}

// Converts time.Time into current timezone SYSTEMTIME, millisecond precision.
func (*_GlobalT) TimeToSystemtime(t time.Time, stLocalTime *win.SYSTEMTIME) {
	// https://support.microsoft.com/en-ca/help/167296/how-to-convert-a-unix-time-t-to-a-win32-filetime-or-systemtime
	epoch := t.UnixNano()/100 + 116_444_736_000_000_000

	ft := win.FILETIME{}
	ft.DwLowDateTime = uint32(epoch & 0xffff_ffff)
	ft.DwHighDateTime = uint32(epoch >> 32)

	stUtc := win.SYSTEMTIME{}
	win.FileTimeToSystemTime(&ft, &stUtc)
	win.SystemTimeToTzSpecificLocalTime(nil, &stUtc, stLocalTime)
}
