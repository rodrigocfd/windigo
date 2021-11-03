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

	_globalWindowDlgPtrs[me] = struct{}{} // store pointer in the set

	// The hwnd member is saved in WM_INITDIALOG processing in dlgProc.
	hInst.CreateDialogParam(win.ResIdInt(me.dialogId), hParent,
		_globalDlgProc, win.LPARAM(unsafe.Pointer(me))) // pass pointer to object itself
}

// Calls DialogBoxParam().
func (me *_WindowDlg) dialogBox(hParent win.HWND, hInst win.HINSTANCE) {
	if me.Hwnd() != 0 {
		panic(fmt.Sprintf("Dialog already created: %d.", me.dialogId))
	}

	_globalWindowDlgPtrs[me] = struct{}{} // store pointer in the set

	// The hwnd member is saved in WM_INITDIALOG processing in dlgProc.
	hInst.DialogBoxParam(win.ResIdInt(me.dialogId), hParent,
		_globalDlgProc, win.LPARAM(unsafe.Pointer(me))) // pass pointer to object itself
}

var (
	// A set keeping all *_WindowDlg that were retrieved in _DlgProc.
	_globalWindowDlgPtrs = make(map[*_WindowDlg]struct{}, 10)

	// Default dialog procedure.
	_globalDlgProc uintptr = syscall.NewCallback(_DlgProc)
)

func _DlgProc(
	hDlg win.HWND, uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM) uintptr {

	var pMe *_WindowDlg

	// https://devblogs.microsoft.com/oldnewthing/20050422-08/?p=35813
	if uMsg == co.WM_INITDIALOG {
		pMe = (*_WindowDlg)(unsafe.Pointer(lParam))
		pMe._WindowBase.hWnd = hDlg // assign actual HWND
		hDlg.SetWindowLongPtr(co.GWLP_DWLP_USER, uintptr(unsafe.Pointer(pMe)))
	} else {
		pMe = (*_WindowDlg)(unsafe.Pointer(hDlg.GetWindowLongPtr(co.GWLP_DWLP_USER)))
	}

	// If object pointer is not stored, then no processing is done.
	// Prevents processing before WM_NCCREATE and after WM_NCDESTROY.
	if _, isStored := _globalWindowDlgPtrs[pMe]; isStored {
		// Process all internal events.
		pMe.internalEvents.processMessages(uMsg, wParam, lParam)

		// Child controls are created in internalEvents closures, so we put the
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
			delete(_globalWindowDlgPtrs, pMe) // remove from set
			hDlg.SetWindowLongPtr(co.GWLP_DWLP_USER, 0)
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
