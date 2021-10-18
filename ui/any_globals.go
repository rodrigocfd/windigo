package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Returns WM_INITDIALOG if parent is a dialog window, otherwise WM_CREATE.
func _CreateOrInitDialog(parent AnyParent) co.WM {
	if parent.isDialog() {
		return co.WM_INITDIALOG
	}
	return co.WM_CREATE
}

//------------------------------------------------------------------------------

var _globalUiFont win.HFONT // Global font, usually Segoe UI.

func _CreateGlobalUiFont() {
	ncm := win.NONCLIENTMETRICS{}
	ncm.SetCbSize()

	win.SystemParametersInfo(co.SPI_GETNONCLIENTMETRICS,
		ncm.CbSize(), unsafe.Pointer(&ncm), 0)
	_globalUiFont = win.CreateFontIndirect(&ncm.LfMenuFont)
}

//------------------------------------------------------------------------------

var _globalCtrlId int = 20_000 // in-between Visual Studio Resource Editor values

func _NextCtrlId() int {
	_globalCtrlId++
	return _globalCtrlId
}

//------------------------------------------------------------------------------

var _globalDpi win.POINT // Global system DPI.

// Multiplies position and size by current DPI factor.
func _MultiplyDpi(pos *win.POINT, size *win.SIZE) {
	if _globalDpi.X == 0 { // not initialized yet?
		dc := win.HWND(0).GetDC()
		_globalDpi.X = dc.GetDeviceCaps(co.GDC_LOGPIXELSX) // cache
		_globalDpi.Y = dc.GetDeviceCaps(co.GDC_LOGPIXELSY)
		win.HWND(0).ReleaseDC(dc)
	}

	if pos != nil {
		pos.X = int32((int64(pos.X) * int64(_globalDpi.X)) / int64(96)) // MulDiv
		pos.Y = int32((int64(pos.Y) * int64(_globalDpi.Y)) / int64(96))
	}
	if size != nil {
		size.Cx = int32((int64(size.Cx) * int64(_globalDpi.X)) / int64(96))
		size.Cy = int32((int64(size.Cy) * int64(_globalDpi.Y)) / int64(96))
	}
}

// If parent is a dialog, converts Dialog Template Units to pixels; otherwise
// multiplies by current DPI factor.
func _ConvertDtuOrMultiplyDpi(parent AnyParent, pos *win.POINT, size *win.SIZE) {
	if parent.isDialog() {
		rc := win.RECT{}
		if pos != nil {
			rc.Left, rc.Top = pos.X, pos.Y
		}
		if size != nil {
			rc.Right, rc.Bottom = size.Cx, size.Cy
		}

		parent.Hwnd().MapDialogRect(&rc)
		if pos != nil {
			pos.X, pos.Y = rc.Left, rc.Top
		}
		if size != nil {
			size.Cx, size.Cy = rc.Right, rc.Bottom
		}
	} else {
		_MultiplyDpi(pos, size)
	}
}

//------------------------------------------------------------------------------

// Overwrites X and Y if they are different from zero.
func _OwPt(point *win.POINT, newVal win.POINT) {
	if newVal.X != 0 {
		point.X = newVal.X
	}
	if newVal.Y != 0 {
		point.Y = newVal.Y
	}
}

// Overwrites Cx and Cy if they are different from zero.
func _OwSz(size *win.SIZE, newVal win.SIZE) {
	if newVal.Cx != 0 {
		size.Cx = newVal.Cx
	}
	if newVal.Cy != 0 {
		size.Cy = newVal.Cy
	}
}

//------------------------------------------------------------------------------

// Calculates the bound rectangle to fit the text with current system font.
func _CalcTextBoundBox(text string, considerAccelerators bool) win.SIZE {
	isTextEmpty := false
	if len(text) == 0 {
		isTextEmpty = true
		text = "Pj" // just a placeholder to get the text height
	}

	if considerAccelerators {
		text = util.RemoveAccelAmpersands(text)
	}

	hwndDesktop := win.GetDesktopWindow()
	hdcDesktop := hwndDesktop.GetDC()
	defer hwndDesktop.ReleaseDC(hdcDesktop)

	hdcCloned := hdcDesktop.CreateCompatibleDC()
	defer hdcCloned.DeleteDC()

	prevFont := hdcCloned.SelectObjectFont(_globalUiFont) // system font; already adjusted to current DPI
	defer hdcCloned.SelectObjectFont(prevFont)

	bounds := hdcCloned.GetTextExtentPoint32(text)

	if isTextEmpty {
		bounds.Cx = 0 // if no text was given, return just the height
	}
	return bounds
}

// Calculates the bound rectangle to fit the text with current system font,
// including the check box for a checkbox/radio.
func _CalcTextBoundBoxWithCheck(
	text string, considerAccelerators bool) win.SIZE {

	boundBox := _CalcTextBoundBox(text, considerAccelerators)
	boundBox.Cx += win.GetSystemMetrics(co.SM_CXMENUCHECK) + // https://stackoverflow.com/a/1165052/6923555
		win.GetSystemMetrics(co.SM_CXEDGE)

	cyCheck := win.GetSystemMetrics(co.SM_CYMENUCHECK)
	if cyCheck > boundBox.Cy {
		boundBox.Cy = cyCheck // if the check is taller than the font, use its height
	}
	return boundBox
}

//------------------------------------------------------------------------------

// Runs the main window loop synchronously.
func _RunMainLoop(hWnd win.HWND, hAccel win.HACCEL) int {
	msg := win.MSG{}
	for {
		if res, err := win.GetMessage(&msg, win.HWND(0), 0, 0); err != nil {
			panic(err)
		} else if res == 0 {
			// WM_QUIT was sent, gracefully terminate the program.
			// If it returned -1, it will simply panic.
			// WParam has the program exit code.
			// https://docs.microsoft.com/en-us/windows/win32/winmsg/using-messages-and-message-queues
			return int(msg.WParam)
		}

		// Check if modeless...

		// If a child window, will retrieve its top-level parent.
		// If a top-level, use itself.
		hTopLevel := msg.HWnd.GetAncestor(co.GA_ROOT)

		// If we have an accelerator table, try to translate the message.
		if hAccel != 0 && hTopLevel.TranslateAccelerator(hAccel, &msg) == nil {
			continue // message translated, no further processing is done
		}

		if hTopLevel.IsDialogMessage(&msg) {
			// Processed all keyboard actions for child controls.
			continue
		}

		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}
}

// Runs the modal window loop synchronously.
func _RunModalLoop(hWnd win.HWND) {
	msg := win.MSG{}
	for {
		if res, err := win.GetMessage(&msg, win.HWND(0), 0, 0); err != nil {
			panic(err)
		} else if res == 0 {
			// WM_QUIT was sent, exit modal loop now and signal parent.
			// If it returned -1, it will simply panic.
			// https://devblogs.microsoft.com/oldnewthing/20050222-00/?p=36393
			win.PostQuitMessage(int32(msg.WParam))
			break
		}

		// If a child window, will retrieve its top-level parent.
		// If a top-level, use itself.
		if msg.HWnd.GetAncestor(co.GA_ROOT).IsDialogMessage(&msg) {
			// Processed all keyboard actions for child controls.
			if hWnd == 0 {
				break // our modal was destroyed, terminate loop
			} else {
				continue
			}
		}

		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)

		if hWnd == 0 {
			break // our modal was destroyed, terminate loop
		}
	}
}

//------------------------------------------------------------------------------

// Paints the themed border for child controls.
func _PaintThemedBorders(hWnd win.HWND, p wm.NcPaint) {
	hWnd.DefWindowProc(co.WM_NCPAINT, p.Msg.WParam, p.Msg.LParam) // make system draw the scrollbar for us

	exStyle := co.WS_EX(hWnd.GetWindowLongPtr(co.GWLP_EXSTYLE))

	if (exStyle&co.WS_EX_CLIENTEDGE) == 0 || // has no border
		!win.IsThemeActive() ||
		!win.IsAppThemed() {
		// No themed borders to be painted.
		return
	}

	rc := hWnd.GetWindowRect() // window outmost coordinates, including margins
	hWnd.ScreenToClientRc(&rc)
	rc.Left += 2 // manual fix, because it comes up anchored at -2,-2
	rc.Top += 2
	rc.Right += 2
	rc.Bottom += 2

	hdc := hWnd.GetWindowDC()
	defer hWnd.ReleaseDC(hdc)

	if hTheme, err := hWnd.OpenThemeData("LISTVIEW"); err == nil { // borrow style from listview
		defer hTheme.CloseThemeData()

		// Clipping region; will draw only within this rectangle.
		// Draw only the borders to avoid flickering.
		rc2 := win.RECT{Left: rc.Left, Top: rc.Top, Right: rc.Left + 2, Bottom: rc.Bottom}
		hTheme.DrawThemeBackground(hdc, co.VS_LISTVIEW_LISTGROUP, &rc, &rc2) // draw themed left border

		rc2 = win.RECT{Left: rc.Left, Top: rc.Top, Right: rc.Right, Bottom: rc.Top + 2}
		hTheme.DrawThemeBackground(hdc, co.VS_LISTVIEW_LISTGROUP, &rc, &rc2) // draw themed top border

		rc2 = win.RECT{Left: rc.Right - 2, Top: rc.Top, Right: rc.Right, Bottom: rc.Bottom}
		hTheme.DrawThemeBackground(hdc, co.VS_LISTVIEW_LISTGROUP, &rc, &rc2) // draw themed right border

		rc2 = win.RECT{Left: rc.Left, Top: rc.Bottom - 2, Right: rc.Right, Bottom: rc.Bottom}
		hTheme.DrawThemeBackground(hdc, co.VS_LISTVIEW_LISTGROUP, &rc, &rc2) // draw themed bottom border
	}
}
