/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type _WndclassexCommon struct {
	wasInit bool

	classNameBuf     []uint16
	ClassName        string      // Class name registered with RegisterClassEx(). Defaults to a computed hash.
	ClassStyle       co.CS       // Window class style, passed to RegisterClassEx(). Defaults to CS_DBLCLKS.
	HCursor          win.HCURSOR // Window cursor, passed to RegisterClassEx(). Defaults to IDC_ARROW.
	HBrushBackground win.HBRUSH  // Window background brush, passed to RegisterClassEx().
}

func (me *_WndclassexCommon) genWndclassex(
	hInst win.HINSTANCE,
	defaultBackgroundColor co.COLOR,
	otherFieldInits func(wcx *win.WNDCLASSEX)) *win.WNDCLASSEX {

	wcx := win.WNDCLASSEX{}

	wcx.CbSize = uint32(unsafe.Sizeof(wcx))
	wcx.HInstance = hInst
	wcx.Style = me.ClassStyle

	if me.HCursor != 0 { // user specified a cursor
		wcx.HCursor = me.HCursor
	} else {
		wcx.HCursor = win.HINSTANCE(0).LoadCursor(co.IDC_ARROW)
	}

	if me.HBrushBackground != 0 { // user specified a background brush
		wcx.HbrBackground = me.HBrushBackground
	} else {
		wcx.HbrBackground = win.CreateSysColorBrush(defaultBackgroundColor)
	}

	otherFieldInits(&wcx)

	// After all the fields are set, if user didn't choose a class name, we
	// generate one by hashing all the WNDCLASSEX fields. That's why it must be
	// the last thing to be done.
	if me.ClassName == "" {
		me.ClassName = fmt.Sprintf("%x.%x.%x.%x.%x.%x.%x.%x.%x.%x",
			wcx.Style, wcx.LpfnWndProc, wcx.CbClsExtra, wcx.CbWndExtra,
			wcx.HInstance, wcx.HIcon, wcx.HCursor, wcx.HbrBackground,
			wcx.LpszMenuName, wcx.HIconSm)
	}

	me.classNameBuf = win.StrToSlice(me.ClassName) // keep the buffer, we'll use a pointer to it
	wcx.LpszClassName = &me.classNameBuf[0]

	return &wcx
}

//------------------------------------------------------------------------------

// Setup parameters for WindowMain.
type _WindowSetupMain struct {
	_WndclassexCommon

	HIcon      win.HICON // Icon associated with the window, passed to RegisterClassEx(). Defaults to none.
	HIconSmall win.HICON // Small icon associated with the window, passed to RegisterClassEx(). Defaults to none.

	Style    co.WS    // Window style, passed to CreateWindowEx(). Defaults to WS_CAPTION | WS_SYSMENU | WS_CLIPCHILDREN | WS_BORDER.
	ExStyle  co.WS_EX // Window extended style, passed to CreateWindowEx(). Defaults to WS_EX_NONE.
	Title    string   // The title of the window, passed to CreateWindowEx(). Defaults to empty string.
	Width    uint     // Initial width of the window, passed to CreateWindowEx(). Defaults to 500px.
	Height   uint     // Initial height of the window, passed to CreateWindowEx(). Defaults to 400px.
	MainMenu Menu     // Main window menu, passed to CreateWindowEx(). Automatically destroyed.

	AcceleratorTable AccelTable // Accelerator table with keyboard shortcuts. Automatically destroyed.
	CmdShow          co.SW      // Passed to ShowWindow(). Defaults to SW_SHOW.
}

func (me *_WindowSetupMain) initOnce() {
	if !me.wasInit { // so it can be called multiple times
		me.wasInit = true

		me.ClassStyle = co.CS_DBLCLKS

		me.Width = 500 // arbitrary dimensions
		me.Height = 400
		me.Style = co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN | co.WS_BORDER
		me.ExStyle = co.WS_EX_NONE

		me.CmdShow = co.SW_SHOW
	}
}

func (me *_WindowSetupMain) genWndclassex(hInst win.HINSTANCE) *win.WNDCLASSEX {
	return me._WndclassexCommon.genWndclassex(hInst, co.COLOR_BTNFACE,
		func(wcx *win.WNDCLASSEX) {
			wcx.HIcon = me.HIcon
			wcx.HIconSm = me.HIconSmall
		})
}

//------------------------------------------------------------------------------

// Setup parameters for WindowModal.
type _WindowSetupModal struct {
	_WndclassexCommon

	Style   co.WS    // Window style, passed to CreateWindowEx(). Defaults to WS_CAPTION | WS_SYSMENU | WS_CLIPCHILDREN | WS_BORDER | WS_VISIBLE.
	ExStyle co.WS_EX // Window extended style, passed to CreateWindowEx(). Defaults to WS_EX_NONE.
	Title   string   // The title of the window, passed to CreateWindowEx(). Defaults to empty string.
	Width   uint     // Initial width of the window, passed to CreateWindowEx(). Defaults to 400px.
	Height  uint     // Initial height of the window, passed to CreateWindowEx(). Defaults to 300px.
}

func (me *_WindowSetupModal) initOnce() {
	if !me.wasInit { // so it can be called multiple times
		me.wasInit = true

		me.ClassStyle = co.CS_DBLCLKS

		me.Width = 400 // arbitrary dimensions
		me.Height = 300
		me.Style = co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN | co.WS_BORDER | co.WS_VISIBLE
		me.ExStyle = co.WS_EX_NONE
	}
}

func (me *_WindowSetupModal) genWndclassex(hInst win.HINSTANCE) *win.WNDCLASSEX {
	return me._WndclassexCommon.genWndclassex(
		hInst, co.COLOR_BTNFACE, func(wcx *win.WNDCLASSEX) {})
}

//------------------------------------------------------------------------------

// Setup parameters for WindowControl.
type _WindowSetupControl struct {
	_WndclassexCommon

	Style   co.WS    // Window style, passed to CreateWindowEx(). Defaults to WS_CHILD | WS_VISIBLE | WS_CLIPCHILDREN | WS_CLIPSIBLINGS.
	ExStyle co.WS_EX // Window extended style, passed to CreateWindowEx(). Defaults to WS_EX_NONE, for a border use WS_EX_CLIENTEDGE.
}

func (me *_WindowSetupControl) initOnce() {
	if !me.wasInit { // so it can be called multiple times
		me.wasInit = true

		me.ClassStyle = co.CS_DBLCLKS

		me.Style = co.WS_CHILD | co.WS_VISIBLE | co.WS_CLIPCHILDREN | co.WS_CLIPSIBLINGS
		me.ExStyle = co.WS_EX_NONE
	}
}

func (me *_WindowSetupControl) genWndclassex(hInst win.HINSTANCE) *win.WNDCLASSEX {
	return me._WndclassexCommon.genWndclassex(
		hInst, co.COLOR_WINDOW, func(wcx *win.WNDCLASSEX) {})
}
