//go:build windows

package ui

import (
	"strings"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Global UI font.
var globalUiFont win.HFONT

func createGlobalUiFont() error {
	if globalUiFont == 0 {
		var err error
		var ncm win.NONCLIENTMETRICS
		ncm.SetCbSize()

		if err = win.SystemParametersInfo(
			co.SPI_GETNONCLIENTMETRICS,
			uint32(unsafe.Sizeof(ncm)),
			unsafe.Pointer(&ncm),
			co.SPIF(0),
		); err != nil {
			return err
		}

		globalUiFont, err = win.CreateFontIndirect(&ncm.LfMenuFont)
		if err != nil {
			return err
		}
	}
	return nil
}

var globalNextCtrlId uint16 = 0xdfff // https://stackoverflow.com/a/18192766/6923555

// Returns an unique child control ID.
func nextCtrlId() uint16 {
	nextId := globalNextCtrlId
	globalNextCtrlId-- // go down
	return nextId
}

// If ctrlId is zero, assigns a new unique control ID.
func setUniqueCtrlId(pCtrlId *uint16) {
	if *pCtrlId == 0 {
		*pCtrlId = nextCtrlId()
	}
}

// Calculates the bound rectangle to fit the text with current UI font.
func calcTextBoundBox(text string) (win.SIZE, error) {
	isTextEmpty := false
	if len(strings.TrimSpace(text)) == 0 {
		isTextEmpty = true
		text = "Pj" // just a placeholder to get the text height
	}

	hwndDesktop := win.GetDesktopWindow()
	hdcDesktop, err := hwndDesktop.GetDC()
	if err != nil {
		return win.SIZE{}, err
	}
	defer hwndDesktop.ReleaseDC(hdcDesktop)

	hdcCloned, err := hdcDesktop.CreateCompatibleDC()
	if err != nil {
		return win.SIZE{}, err
	}
	defer hdcCloned.DeleteDC()

	prevFont, err := hdcCloned.SelectObjectFont(globalUiFont)
	if err != nil {
		return win.SIZE{}, err
	}
	defer hdcCloned.SelectObjectFont(prevFont)

	bounds, err := hdcCloned.GetTextExtentPoint32(text)
	if err != nil {
		return win.SIZE{}, err
	}

	if isTextEmpty {
		bounds.Cx = 0 // if no text was given, return just the height
	}
	return bounds, nil
}

// Calculates the bound rectangle to fit the text with current UI font,
// including the check box for a checkbox/radio.
func calcTextBoundBoxWithCheck(text string) (win.SIZE, error) {
	boundBox, err := calcTextBoundBox(text)
	if err != nil {
		return win.SIZE{}, err
	}

	boundBox.Cx += win.GetSystemMetrics(co.SM_CXMENUCHECK) + // https://stackoverflow.com/a/1165052/6923555
		win.GetSystemMetrics(co.SM_CXEDGE)

	cyCheck := win.GetSystemMetrics(co.SM_CYMENUCHECK)
	if cyCheck > boundBox.Cy {
		boundBox.Cy = cyCheck // if the check is taller than the font, use its height
	}
	return boundBox, nil
}

// Paints the border of a child control according to the system theme.
func paintThemedBorders(hWnd win.HWND, p WmNcPaint) {
	hWnd.DefWindowProc(co.WM_NCPAINT, p.Raw.WParam, p.Raw.LParam) // make system draw the scrollbar for us

	exStyle, _ := hWnd.ExStyle()
	if (exStyle&co.WS_EX_CLIENTEDGE) == 0 || // has no border
		!win.IsThemeActive() ||
		!win.IsAppThemed() {
		// No themed borders to be painted.
		return
	}

	rc, _ := hWnd.GetWindowRect() // window outmost coordinates, including margins
	hWnd.ScreenToClientRc(&rc)
	win.OffsetRect(&rc, 2, 2) // because it comes up anchored at -2,-2

	hdc, _ := hWnd.GetWindowDC()
	defer hWnd.ReleaseDC(hdc)

	// The HRGN which comes in WM_NCPAINT seems to be invalid, so we carve our own.
	rcHole := rc
	win.InflateRect(&rcHole, -2, -2)
	hRgnHole, _ := win.CreateRectRgnIndirect(rcHole)
	defer hRgnHole.DeleteObject()

	hRgnClip, _ := win.CreateRectRgnIndirect(rc)
	defer hRgnClip.DeleteObject()
	hRgnClip.CombineRgn(hRgnClip, hRgnHole, co.RGN_DIFF)
	hdc.SelectClipRgn(hRgnClip)

	if hTheme, err := hWnd.OpenThemeData("EDIT"); err == nil {
		defer hTheme.CloseThemeData()
		hTheme.DrawThemeBackground(hdc, co.VS_EDIT_EDITTEXT_NORMAL, &rc, nil)
	}
}
