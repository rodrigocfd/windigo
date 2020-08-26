/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Native status bar control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/status-bars
type StatusBar struct {
	_ControlNativeBase
	parts       []statusBarPart
	firstAdjust bool
}

type statusBarPart struct { // describes each added part
	sizePixels   uint32
	resizeWeight uint32
}

// Adds a part which has a fixed width.
func (me *StatusBar) AddFixedPart(sizePixels uint32) *StatusBar {
	me.parts = append(me.parts, statusBarPart{
		sizePixels: sizePixels,
	})
	return me
}

// Adds a part which resizes according to the parent window.
// How resizeWeight works:
// Suppose you have 3 parts, respectively with weights of 1, 1 and 2.
// If available client area is 400px, respective part widths will be 100, 100
// and 200px.
func (me *StatusBar) AddResizablePart(resizeWeight uint32) *StatusBar {
	me.parts = append(me.parts, statusBarPart{
		resizeWeight: resizeWeight,
	})
	return me
}

// Call during WM_SIZE processing.
func (me *StatusBar) Adjust(p WmSize) {
	if p.Request() == co.SIZE_MINIMIZED {
		return // no need to adjust when minimized
	}
	me.firstAdjust = true

	cxParent := uint32(p.ClientAreaSize().Cx) // available width
	me.Hwnd().SendMessage(co.WM_SIZE, 0, 0)   // tell statusbar to fit parent

	// Find the space to be divided among variable-width parts, and total weight
	// of variable-width parts.
	totalWeight := uint32(0)
	cxVariable := cxParent

	for _, part := range me.parts {
		if part.resizeWeight == 0 { // fixed width
			cxVariable -= part.sizePixels
		} else { // variable size
			totalWeight += part.resizeWeight
		}
	}

	// Fill right edges array with the right edge of each part.
	rightEdges := make([]uint32, len(me.parts))
	cxTotal := cxParent

	for i := len(me.parts) - 1; i >= 0; i-- {
		rightEdges[i] = cxTotal

		if me.parts[i].resizeWeight == 0 { // fixed width
			cxTotal -= me.parts[i].sizePixels
		} else { // variable size
			cxTotal -= (cxVariable / totalWeight) * me.parts[i].resizeWeight
		}
	}

	me.sendSbMessage(co.SB_SETPARTS,
		win.WPARAM(len(me.parts)), win.LPARAM(unsafe.Pointer(&rightEdges[0])))
}

// Calls CreateWindowEx().
//
// Control will be docked at bottom of parent window.
func (me *StatusBar) Create(parent Window, ctrlId int32) *StatusBar {
	style := co.WS_CHILD | co.WS_VISIBLE

	parentStyle := parent.Hwnd().GetStyle()
	if (parentStyle&co.WS_MAXIMIZEBOX) != 0 ||
		(parentStyle&co.WS_SIZEBOX) != 0 {
		// Parent window can change its size.
		style |= co.WS(co.SBARS_SIZEGRIP)
	}

	me._ControlNativeBase.create(co.WS_EX(0), "msctls_statusbar32", "", style,
		0, 0, 0, 0, parent, ctrlId)
	return me
}

// Retrieves the HICON of the part.
//
// The status bar won't destroy the icon after use.
func (me *StatusBar) Icon(part uint32) win.HICON {
	return win.HICON(
		me.sendSbMessage(co.SB_GETICON, win.WPARAM(part), 0),
	)
}

// The status bar won't destroy the icon after use.
func (me *StatusBar) SetIcon(part uint32, hIcon win.HICON) *StatusBar {
	me.sendSbMessage(co.SB_SETICON,
		win.WPARAM(part), win.LPARAM(hIcon))
	return me
}

// Sets the text of the part.
func (me *StatusBar) SetText(part uint32, text string) *StatusBar {
	if !me.firstAdjust { // text is painted only after first adjust
		me.Adjust(WmSize{ // manually construct param
			_Wm{
				WParam: win.WPARAM(co.SIZE_RESTORED),
				LParam: win.LPARAM(me.Hwnd().GetParent().GetClientRect().Right),
			},
		})
	}
	me.sendSbMessage(co.SB_SETTEXT,
		win.WPARAM(part), win.LPARAM(unsafe.Pointer(win.StrToPtr(text))))
	return me
}

// Retrieves the text of the part.
func (me *StatusBar) Text(part uint32) string {
	len := uint16(me.sendSbMessage(co.SB_GETTEXTLENGTH, win.WPARAM(part), 0))
	if len == 0 {
		return ""
	}

	buf := make([]uint16, len+1)
	me.sendSbMessage(co.SB_GETTEXT,
		win.WPARAM(part), win.LPARAM(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf)
}

// Syntactic sugar.
func (me *StatusBar) sendSbMessage(msg co.SB,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.Hwnd().SendMessage(co.WM(msg), wParam, lParam)
}
