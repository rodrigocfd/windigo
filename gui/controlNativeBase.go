/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/gui/wm"
	"wingows/win"
)

var (
	globalBaseSubclassId  = uint32(0)  // incremented at each subclass installed
	globalSubclassProcPtr = uintptr(0) // necessary for RemoveWindowSubclass
)

// Base to all native child control types, like Button and Edit.
// Allows control subclassing.
type controlNativeBase struct {
	controlIdGuard
	hwnd       win.HWND
	msgs       windowDepotMsg
	subclassId uint32
}

func makeNativeControlBase(ctrlId int32) controlNativeBase {
	return controlNativeBase{
		controlIdGuard: makeCtrlIdGuard(ctrlId),
	}
}

// Returns the underlying HWND handle of this native control.
func (me *controlNativeBase) Hwnd() win.HWND {
	return me.hwnd
}

// Exposes all the control subclass methods that can be handled.
// The subclass will be installed in create() if at least 1 message was added.
func (me *controlNativeBase) OnSubclassMsg() *windowDepotMsg {
	if me.hwnd != 0 {
		panic("Cannot add subclass message after the control was created.")
	}
	return &me.msgs
}

func (me *controlNativeBase) create(exStyle co.WS_EX, className, title string,
	style co.WS, x, y int32, width, height uint32, parent Window) {

	if me.hwnd != 0 {
		panic(fmt.Sprintf("Trying to create %s twice.", className))
	}

	me.hwnd = win.CreateWindowEx(exStyle, className, title, style,
		x, y, width, height, parent.Hwnd(), win.HMENU(me.controlIdGuard.Id()),
		parent.Hwnd().GetInstance(), nil)

	if len(me.msgs.mapMsgs) > 0 || // at last 1 subclass message was added?
		len(me.msgs.mapCmds) > 0 {

		if globalSubclassProcPtr == 0 {
			globalSubclassProcPtr = syscall.NewCallback(subclassProc)
		}
		globalBaseSubclassId++
		me.subclassId = globalBaseSubclassId

		// Subclass is installed after window creation, thus WM_CREATE can never
		// be handled for a subclassed control.
		me.hwnd.SetWindowSubclass(globalSubclassProcPtr,
			me.subclassId, unsafe.Pointer(me))
	}
}

func subclassProc(hwnd win.HWND, msg co.WM,
	wParam win.WPARAM, lParam win.LPARAM,
	uIdSubclass, dwRefData uintptr) uintptr {

	// Retrieve passed pointer.
	pMe := (*controlNativeBase)(unsafe.Pointer(dwRefData))

	// Save *nativeControlBase from being collected by GC; stored won't be used.
	hwnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))

	if pMe != nil && pMe.hwnd != 0 {
		userRet, wasHandled := pMe.msgs.processMessage(msg,
			wm.Base{WParam: wParam, LParam: lParam}) // try to process the message with an user handler

		if msg == co.WM_NCDESTROY { // even if the user handles WM_NCDESTROY, we must ensure cleanup
			pMe.hwnd.RemoveWindowSubclass(globalSubclassProcPtr, pMe.subclassId)
		}
		if wasHandled && msg != co.WM_LBUTTONUP {
			// For some reason, if we don't call DefSubclassProc with WM_LBUTTONUP,
			// all parent window messages are routed to this proc, and it becomes
			// unresponsive. So we return user result only if not WM_LBUTTONUP.
			return userRet
		}
	} else if msg == co.WM_NCDESTROY { // https://devblogs.microsoft.com/oldnewthing/20031111-00/?p=41883
		hwnd.RemoveWindowSubclass(globalSubclassProcPtr, pMe.subclassId)
	}

	return hwnd.DefSubclassProc(msg, wParam, lParam) // message was not processed
}

// Calculates the bound rectangle to fit the text with current system font.
func calcIdealSize(hReferenceDc win.HWND, text string,
	considerAccelerators bool) (uint32, uint32) {

	isTextEmpty := false
	if len(text) == 0 {
		isTextEmpty = true
		text = "Pj" // just a placeholder to get the text height
	}

	if considerAccelerators {
		text = removeAccelAmpersands(text)
	}

	parentDc := hReferenceDc.GetDC()
	cloneDc := parentDc.CreateCompatibleDC()
	prevFont := cloneDc.SelectObjectFont(globalUiFont.Hfont()) // system font; already adjusted to current DPI
	bounds := cloneDc.GetTextExtentPoint32(text)
	cloneDc.SelectObjectFont(prevFont)
	cloneDc.DeleteDC()
	hReferenceDc.ReleaseDC(parentDc)

	if isTextEmpty {
		bounds.Cx = 0 // if no text was given, return just the height
	}
	return uint32(bounds.Cx), uint32(bounds.Cy)
}

// "&He && she" becomes "He & she".
func removeAccelAmpersands(text string) string {
	buf := strings.Builder{}
	for i := 0; i < len(text)-1; i++ {
		if text[i] == '&' && text[i+1] != '&' {
			continue
		}
		buf.WriteByte(text[i])
	}
	if text[len(text)-1] != '&' {
		buf.WriteByte(text[len(text)-1])
	}
	return buf.String()
}
