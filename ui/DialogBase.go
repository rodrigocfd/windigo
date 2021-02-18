/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Modal popup dialog.
type _DialogBase struct {
	hwnd           win.HWND
	eventsWmCmdNfy *_EventsWmCmdNfy
	dlgId          int
}

// Constructor.
func _NewDialogBase(dlgId int) *_DialogBase {
	return &_DialogBase{
		eventsWmCmdNfy: _NewEventsWmCmdNfy(),
		dlgId:          dlgId,
	}
}

// Creates all child controls declared in the dialog resource.
//
// Should be called at On().WmInitDialog().
func (me *_DialogBase) CreateChildren(children ...ControlResource) {
	for _, child := range children {
		child.createAsDlgCtrl()
	}
}

// Returns the underlying HWND handle of this window.
func (me *_DialogBase) Hwnd() win.HWND {
	return me.hwnd
}

// Exposes all the window messages the can be handled.
//
// Cannot be called after the window was created.
func (me *_DialogBase) On() *_EventsWmCmdNfy {
	if me.hwnd != 0 {
		panic("Cannot add message after the window was created.")
	}
	return me.eventsWmCmdNfy
}

// Calls CreateDialogParam().
func (me *_DialogBase) createDialogParam(
	hInst win.HINSTANCE, parent Parent) win.HWND {

	hParent := win.HWND(0)
	if parent != nil {
		hParent = parent.Hwnd()
	}

	return hInst.CreateDialogParam(int32(me.dlgId), hParent,
		syscall.NewCallback(_globalDlgProc),
		win.LPARAM(unsafe.Pointer(me))) // pass pointer to our object
}

// Calls DialogBoxParam().
func (me *_DialogBase) dialogBoxParam(hInst win.HINSTANCE, parent Parent) int {
	hParent := win.HWND(0)
	if parent != nil {
		hParent = parent.Hwnd()
	}

	return int(
		hInst.DialogBoxParam(int32(me.dlgId), hParent,
			syscall.NewCallback(_globalDlgProc),
			win.LPARAM(unsafe.Pointer(me))), // pass pointer to our object
	)
}

// DLGPROC for all dialog windows.
func _globalDlgProc(
	hwnd win.HWND, msg co.WM, wParam win.WPARAM, lParam win.LPARAM) uintptr {

	if msg == co.WM_INITDIALOG {
		base := (*_DialogBase)(unsafe.Pointer(lParam))
		hwnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(base)))

		hwnd.EnumChildWindows(func(hChild win.HWND, _ win.LPARAM) bool { // set font on children
			hChild.SendMessage(co.WM_SETFONT,
				win.WPARAM(_global.UiFont().Hfont()), win.MakeLParam(0, 0)) // FALSE, 0
			return true
		}, 0)

		base.hwnd = hwnd // assign actual HWND
	}

	// Retrieve passed pointer.
	pMe := (*_WindowBase)(unsafe.Pointer(
		hwnd.GetWindowLongPtr(co.GWLP_USERDATA)))

	// If the retrieved *_DialogBase stays here, the GC will collect it.
	// Sending it away will prevent the GC collection.
	// https://stackoverflow.com/a/51188315
	hwnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))

	// If no pointer stored, then no processing is done.
	// Prevents processing before WM_INITDIALOG and after WM_NCDESTROY.
	if pMe == nil {
		return 0 // FALSE
	}

	// Try to process the message with an user handler.
	retVal, useRetVal, wasHandled := pMe.eventsWmCmdNfy.processMessage(
		msg, Wm{WParam: wParam, LParam: lParam})

	// No further messages processed after this one.
	if msg == co.WM_NCDESTROY {
		pMe.hwnd.SetWindowLongPtr(co.GWLP_DWLP_USER, 0) // clear passed pointer
		pMe.hwnd = win.HWND(0)
	}

	if wasHandled {
		if useRetVal {
			return retVal
		}
		return 1 // TRUE
	}
	return 0 // FALSE; message was not processed
}
