//go:build windows

package ui

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _BaseCtrl struct {
	ctrlId uint16
	hWnd   win.HWND

	subclassEvents EventsWindow
	subclassProc   uintptr // Prevents InvalidInitCycle.
}

// Constructor.
func newBaseCtrl(ctrlId uint16) _BaseCtrl {
	return _BaseCtrl{
		ctrlId: ctrlId,
		hWnd:   win.HWND(0),
	}
}

// Returns the underlying HWND handle of this window.
//
// Note that this handle is initially zero, existing only after window creation.
func (me *_BaseCtrl) Hwnd() win.HWND {
	return me.hWnd
}

// Returns the control ID, unique within the same Parent.
func (me *_BaseCtrl) CtrlId() uint16 {
	return me.ctrlId
}

// If parent is a dialog, sets the focus by sending [WM_NEXTDLGCTL]. This draws
// the borders correctly in some undefined controls, like buttons.
//
// Otherwise, calls [SetFocus].
//
// [WM_NEXTDLGCTL]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-nextdlgctl
// [SetFocus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setfocus
func (me *_BaseCtrl) Focus() {
	hParent, _ := me.hWnd.GetAncestor(co.GA_PARENT)
	isDialog, _ := hParent.IsDialog()
	if isDialog {
		hParent.SendMessage(co.WM_NEXTDLGCTL, win.WPARAM(me.hWnd), 1)
	} else {
		me.hWnd.SetFocus()
	}
}

// Exposes all the subclass events the can be handled.
//
// Panics if called after the control was created.
//
// Note that subclassing is a potentially slow technique, prefer using ordinary
// events.
func (me *_BaseCtrl) OnSubclass() *EventsWindow {
	if me.hWnd != 0 {
		panic("Cannot subclass a control after it is created.")
	}
	return &me.subclassEvents
}

func (me *_BaseCtrl) createWindow(
	exStyle co.WS_EX,
	className string,
	title string,
	style co.WS,
	pos win.POINT,
	size win.SIZE,
	parent Parent,
	setGlobalUiFont bool,
) {
	if me.hWnd != 0 {
		panic("Cannot create control twice.")
	}

	hInst, _ := parent.Hwnd().HInstance()
	me.hWnd, _ = win.CreateWindowEx(exStyle, win.ClassNameStr(className),
		title, style, int(pos.X), int(pos.Y), uint(size.Cx), uint(size.Cy),
		parent.Hwnd(), win.HMENU(me.ctrlId), hInst, win.LPARAM(0))
	if setGlobalUiFont {
		me.hWnd.SendMessage(co.WM_SETFONT, win.WPARAM(globalUiFont), win.LPARAM(1))
	}
	me.installSubclass()
}

func (me *_BaseCtrl) assignDialog(parent Parent) {
	if parent.base().wndTy == _WNDTY_RAW {
		panic("Parent window is not a dialog, cannot instantiate control.")
	}
	if me.hWnd != 0 {
		panic("Cannot create control twice.")
	}
	if parent.Hwnd() == 0 {
		panic("Cannot create control before parent window creation.")
	}

	me.hWnd, _ = parent.Hwnd().GetDlgItem(me.ctrlId)
	me.installSubclass()
}

func (me *_BaseCtrl) panicIfAddingEventAfterCreated() {
	if me.hWnd != 0 {
		panic("Cannot add event handling after the control has been created.")
	}
}

func (me *_BaseCtrl) installSubclass() {
	if me.subclassEvents.hasMessage() {
		subclassProcCallback()
		_subclassId++
		me.subclassProc = subclassProcCallback()
		err := me.hWnd.SetWindowSubclass(me.subclassProc, _subclassId, unsafe.Pointer(me)) // pass pointer to object itself
		if err != nil {
			panic(err)
		}
	}
}

var (
	_subclassId           uint32 = 0
	_subclassProcCallback uintptr
)

func subclassProcCallback() uintptr {
	if _subclassProcCallback == 0 {
		_subclassProcCallback = syscall.NewCallback(
			func(
				hWnd win.HWND,
				uMsg co.WM,
				wParam win.WPARAM,
				lParam win.LPARAM,
				uIdSubclass, dwRefData uintptr,
			) uintptr {
				pMe := (*_BaseCtrl)(unsafe.Pointer(dwRefData)) // retrieve passed pointer

				userRet, hasUserRet := uintptr(0), false

				if pMe != nil {
					msg := Wm{uMsg, wParam, lParam}
					userRet, hasUserRet = pMe.subclassEvents.processLastMessage(msg)
				}

				if uMsg == co.WM_NCDESTROY { // always check
					hWnd.RemoveWindowSubclass(pMe.subclassProc, uint32(uIdSubclass)) // https://devblogs.microsoft.com/oldnewthing/20031111-00/?p=41883
					if pMe != nil {
						pMe.subclassEvents.clear()
					}
				}

				if hasUserRet {
					return userRet
				} else {
					return hWnd.DefSubclassProc(uMsg, wParam, lParam)
				}
			},
		)
	}
	return _subclassProcCallback
}
