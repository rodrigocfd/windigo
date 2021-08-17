package ui

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Base to all dialog windows.
type _WindowDlg struct {
	_WindowBase
	dialogId int
}

func (me *_WindowDlg) new(dialogId int) {
	me._WindowBase.new()
	me.dialogId = dialogId
}

// Calls CreateDialogParam().
func (me *_WindowDlg) createDialog(hParent win.HWND, hInst win.HINSTANCE) {
	if me.Hwnd() != 0 {
		panic(fmt.Sprintf("Dialog already created: %d.", me.dialogId))
	}

	hInst.CreateDialogParam(uint16(me.dialogId), hParent,
		syscall.NewCallback(_DlgProc), win.LPARAM(unsafe.Pointer(me))) // pass pointer to object itself
}

// Calls DialogBoxParam().
func (me *_WindowDlg) dialogBox(hParent win.HWND, hInst win.HINSTANCE) {
	if me.Hwnd() != 0 {
		panic(fmt.Sprintf("Dialog already created: %d.", me.dialogId))
	}

	hInst.DialogBoxParam(int32(me.dialogId), hParent,
		syscall.NewCallback(_DlgProc), win.LPARAM(unsafe.Pointer(me))) // pass pointer to object itself
}

// Keeps all *_WindowDlg that were retrieved in _DlgProc.
var _globalWindowDlgPtrs = make(map[win.HWND]*_WindowDlg, 10)

// Default dialog procedure.
func _DlgProc(
	hDlg win.HWND, uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM) uintptr {

	// https://devblogs.microsoft.com/oldnewthing/20050422-08/?p=35813
	if uMsg == co.WM_INITDIALOG {
		pMe := (*_WindowDlg)(unsafe.Pointer(lParam))
		_globalWindowDlgPtrs[hDlg] = pMe
		pMe._WindowBase.hWnd = hDlg // assign actual HWND
	}

	// Retrieve passed pointer.
	// If no pointer stored, then no processing is done.
	// Prevents processing before WM_NCCREATE and after WM_NCDESTROY.
	if pMe, hasPtr := _globalWindowDlgPtrs[hDlg]; hasPtr {
		// Process all internal events.
		pMe.internalEvents.processMessages(uMsg, wParam, lParam)

		// Child controls are created in internalEvents closures, so we set the
		// system font only after running them.
		if uMsg == co.WM_INITDIALOG {
			hDlg.SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 0)
			hDlg.EnumChildWindows(func(hChild win.HWND) bool {
				hChild.SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 0)
				return true
			})
		}

		// Try to process the message with an user handler.
		retVal, meaningfulRet, wasHandled :=
			pMe._WindowBase.events.processMessage(uMsg, wParam, lParam)

		// No further messages processed after this one.
		if uMsg == co.WM_NCDESTROY {
			delete(_globalWindowDlgPtrs, hDlg) // clear our pointer
			pMe._WindowBase.hWnd = win.HWND(0)
		}

		if wasHandled {
			if meaningfulRet {
				return retVal
			}
			return 1 // message processed, default return value
		}
	}

	return 0 // message not processed
}
