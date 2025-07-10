//go:build windows

package ui

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Base to all dialog-based windows created with CreateDialogParam and
// DialogBoxParam.
type _BaseDlg struct {
	_BaseContainer
	dlgId uint16
}

// Constructor.
func newBaseDlg(dlgId uint16) _BaseDlg {
	if dlgId == 0 {
		panic("Dialog ID must be specified.")
	}

	return _BaseDlg{
		_BaseContainer: newBaseContainer(_WNDTY_DLG),
		dlgId:          dlgId,
	}
}

func (me *_BaseDlg) createDialogParam(hInst win.HINSTANCE, hParent win.HWND) {
	if me.hWnd != 0 {
		panic("Cannot create dialog twice.")
	}
	dlgProcCallback()

	// The hWnd member is saved in WM_INITDIALOG processing in dlgProc.
	_, err := hInst.CreateDialogParam(win.ResIdInt(me.dlgId), hParent, dlgProcCallback(),
		win.LPARAM(unsafe.Pointer(me))) // pass pointer to object itself
	if err != nil {
		panic(err)
	}
}

func (me *_BaseDlg) dialogBoxParam(hInst win.HINSTANCE, hParent win.HWND) {
	if me.hWnd != 0 {
		panic("Cannot create dialog twice.")
	}

	// The hWnd member is saved in WM_INITDIALOG processing in dlgProc.
	_, err := hInst.DialogBoxParam(win.ResIdInt(me.dlgId), hParent, dlgProcCallback(),
		win.LPARAM(unsafe.Pointer(me))) // pass pointer to object itself
	if err != nil {
		panic(err)
	}
}

func (me *_BaseDlg) setIcon(hInst win.HINSTANCE, iconId uint16) error {
	hGdiobjIcon, err := hInst.LoadImage(win.ResIdInt(iconId),
		co.IMAGE_ICON, 16, 16, co.LR_DEFAULTCOLOR|co.LR_SHARED)
	if err != nil {
		return err
	}
	hGdi32, err := hInst.LoadImage(win.ResIdInt(iconId),
		co.IMAGE_ICON, 32, 32, co.LR_DEFAULTCOLOR|co.LR_SHARED)
	if err != nil {
		return err
	}
	hIcon16, hIcon32 := win.HICON(hGdiobjIcon), win.HICON(hGdi32)

	me.hWnd.SendMessage(co.WM_SETICON, win.WPARAM(co.ICON_SZ_SMALL), win.LPARAM(hIcon16))
	me.hWnd.SendMessage(co.WM_SETICON, win.WPARAM(co.ICON_SZ_BIG), win.LPARAM(hIcon32))
	return nil
}

var _dlgProcCallback uintptr

func dlgProcCallback() uintptr {
	if _dlgProcCallback != 0 {
		return _dlgProcCallback
	}

	_dlgProcCallback = syscall.NewCallback(
		func(hDlg win.HWND, uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM) uintptr {
			var pMe *_BaseDlg

			if uMsg == co.WM_INITDIALOG {
				pMe = (*_BaseDlg)(unsafe.Pointer(lParam))
				pMe.hWnd = hDlg
				hDlg.SetWindowLongPtr(co.GWLP_DWLP_USER, uintptr(unsafe.Pointer(pMe)))
			} else {
				ptr, _ := hDlg.GetWindowLongPtr(co.GWLP_DWLP_USER) // retrieve
				pMe = (*_BaseDlg)(unsafe.Pointer(ptr))
			}

			// If no pointer stored, then no processing is done.
			// Prevents processing before WM_INITDIALOG and after WM_NCDESTROY.
			if pMe == nil {
				return 0 // FALSE
			}

			// Execute before-user closures, keep track if at least one was executed.
			msg := Wm{uMsg, wParam, lParam}
			atLeastOneBeforeUser := pMe.beforeUserEvents.processAllMessages(msg)

			// Execute user closure, if any.
			userRet, hasUserRet := pMe.userEvents.processLastMessage(msg)

			// Execute post-user closures, keep track if at least one was executed.
			atLeastOneAfterUser := pMe.afterUserEvents.processAllMessages(msg)

			switch uMsg {
			case co.WM_INITDIALOG:
				pMe.removeWmCreateInitdialog() // will release all memory in these closures
			case co.WM_NCDESTROY: // always check
				hDlg.SetWindowLongPtr(co.GWLP_DWLP_USER, 0)
				pMe.hWnd = win.HWND(0)
				pMe.clearMessages()
			}

			if hasUserRet {
				return userRet
			} else if atLeastOneBeforeUser || atLeastOneAfterUser {
				return 1 // TRUE
			} else {
				return 0 // FALSE
			}
		},
	)
	return _dlgProcCallback
}
