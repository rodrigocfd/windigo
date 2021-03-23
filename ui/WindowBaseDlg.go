package ui

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Base to all dialog windows.
type _WindowBaseDlg struct {
	_WindowBase
	dialogId int
}

func (me *_WindowBaseDlg) new(dialogId int) {
	me._WindowBase.new()
	me.dialogId = dialogId
}

// Calls CreateDialogParam().
func (me *_WindowBaseDlg) createDialog(hParent win.HWND, hInst win.HINSTANCE) {
	if me.Hwnd() != 0 {
		panic(fmt.Sprintf("Dialog already created: %d.", me.dialogId))
	}

	hInst.CreateDialogParam(int32(me.dialogId), hParent,
		syscall.NewCallback(_DlgProc), win.LPARAM(unsafe.Pointer(me))) // pass pointer to object itself
}

// Calls DialogBoxParam().
func (me *_WindowBaseDlg) dialogBox(hParent win.HWND, hInst win.HINSTANCE) {
	if me.Hwnd() != 0 {
		panic(fmt.Sprintf("Dialog already created: %d.", me.dialogId))
	}

	hInst.DialogBoxParam(int32(me.dialogId), hParent,
		syscall.NewCallback(_DlgProc), win.LPARAM(unsafe.Pointer(me))) // pass pointer to object itself
}

// Default dialog procedure.
func _DlgProc(
	hDlg win.HWND, uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM) uintptr {

	// https://devblogs.microsoft.com/oldnewthing/20050422-08/?p=35813
	if uMsg == co.WM_INITDIALOG {
		pMe := (*_WindowBaseDlg)(unsafe.Pointer(lParam))
		hDlg.SetWindowLongPtr(co.GWLP_DWLP_USER, uintptr(unsafe.Pointer(pMe)))
		pMe._WindowBase.hWnd = hDlg // assign actual HWND
	}

	// Retrieve passed pointer.
	pMe := (*_WindowBaseDlg)(unsafe.Pointer(hDlg.GetWindowLongPtr(co.GWLP_DWLP_USER)))

	// If the retrieved *_WindowBaseDlg stays here, the GC will collect it.
	// Sending it away will prevent the GC collection.
	// https://stackoverflow.com/a/51188315
	hDlg.SetWindowLongPtr(co.GWLP_DWLP_USER, uintptr(unsafe.Pointer(pMe)))

	// If no pointer stored, then no processing is done.
	// Prevents processing before WM_NCCREATE and after WM_NCDESTROY.
	if pMe != nil {
		// Process all internal events.
		pMe.internalEvents.processMessages(uMsg, wParam, lParam)

		// Child controls are created in internalEvents closures, so we set the
		// system font only after running them.
		if uMsg == co.WM_INITDIALOG {
			hDlg.SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 0)
			hDlg.EnumChildWindows(func(hChild win.HWND, _ win.LPARAM) bool {
				hChild.SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 0)
				return true
			}, 0)
		}

		// Try to process the message with an user handler.
		retVal, meaningfulRet, wasHandled :=
			pMe._WindowBase.events.processMessage(uMsg, wParam, lParam)

		// No further messages processed after this one.
		if uMsg == co.WM_NCDESTROY {
			pMe._WindowBase.hWnd.SetWindowLongPtr(co.GWLP_DWLP_USER, 0) // clear passed pointer
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
