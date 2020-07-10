/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/gui/wm"
	"wingows/win"
)

type windowBaseDepot struct { // aglutinate both msg and nfy into one façade
	windowDepotMsg
	windowDepotNfy
}

// Base to all window types: WindowControl, WindowMain and WindowModal.
type windowBase struct {
	hwnd  win.HWND
	depot windowBaseDepot
}

const wM_UI_THREAD = co.WM_APP + 0x3FFF // used in UI thread handling
type threadPack struct{ userFunc func() }

// Returns the underlying HWND handle of this window.
func (me *windowBase) Hwnd() win.HWND {
	return me.hwnd
}

// Exposes all the window messages the can be handled.
func (me *windowBase) OnMsg() *windowBaseDepot {
	if me.Hwnd() != 0 {
		panic("Cannot add message after the window was created.")
	}
	return &me.depot
}

// Runs a closure synchronously in the window original UI thread.
// When in a goroutine, you *MUST* use this method to update the UI.
func (me *windowBase) RunUiThread(userFunc func()) {
	// This method is analog to SendMessage (synchronous), but intended to be
	// called from another thread, so a callback function can, tunelled by
	// wndproc, run in the original thread of the window, thus allowing GUI
	// updates. This avoids the user to deal with a custom WM_ message.
	pack := &threadPack{userFunc: userFunc}
	me.hwnd.SendMessage(wM_UI_THREAD, 0xC0DEF00D,
		win.LPARAM(unsafe.Pointer(pack)))
}

func (me *windowBase) registerClass(wcx *win.WNDCLASSEX) win.ATOM {
	wcx.LpfnWndProc = syscall.NewCallback(wndProc)
	atom, lerr := wcx.RegisterClassEx()

	if lerr == co.ERROR_CLASS_ALREADY_EXISTS {
		// https://devblogs.microsoft.com/oldnewthing/20150429-00/?p=44984
		// https://devblogs.microsoft.com/oldnewthing/20041011-00/?p=37603
		atom = wcx.HInstance.GetClassInfoEx(
			(*uint16)(unsafe.Pointer(wcx.LpszClassName)), wcx)

	} else if lerr != co.ERROR_SUCCESS {
		panic(lerr.Format("RegisterClassEx failed."))
	}

	return atom
}

func (me *windowBase) createWindow(uiName string, exStyle co.WS_EX,
	className, title string, style co.WS, x, y int32, width, height uint32,
	parent Window, menu win.HMENU, hInst win.HINSTANCE) {

	if me.hwnd != 0 {
		panic(fmt.Sprintf("Trying to create %s \"%s\" twice.",
			uiName, title))
	}

	hwndParent := win.HWND(0) // if no parent, pass zero to CreateWindowEx
	if parent != nil {
		hwndParent = parent.Hwnd()
	}

	me.defaultMessageHandling()

	// The hwnd member is saved in WM_NCCREATE processing in wndProc.
	win.CreateWindowEx(exStyle, className, title, style, x, y, width, height,
		hwndParent, menu, hInst, unsafe.Pointer(me)) // pass pointer to our object
}

func (me *windowBase) defaultMessageHandling() {
	me.depot.addMsg(wM_UI_THREAD, func(p wm.Base) uintptr { // handle our custom thread UI message
		if p.WParam == 0xC0DEF00D {
			pack := (*threadPack)(unsafe.Pointer(p.LParam))
			pack.userFunc()
		}
		return 0
	})
}

func wndProc(hwnd win.HWND, msg co.WM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	// https://devblogs.microsoft.com/oldnewthing/20050422-08/?p=35813
	if msg == co.WM_NCCREATE {
		cs := (*win.CREATESTRUCT)(unsafe.Pointer(lParam))
		base := (*windowBase)(unsafe.Pointer(cs.LpCreateParams))
		hwnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(base)))
		base.hwnd = hwnd // assign actual HWND
	}

	// Retrieve passed pointer.
	pMe := (*windowBase)(unsafe.Pointer(
		hwnd.GetWindowLongPtr(co.GWLP_USERDATA)))

	// If the retrieved *windowBase stays here, the GC will collect it.
	// Sending it away will prevent the GC collection.
	// https://stackoverflow.com/a/51188315
	hwnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))

	// If no pointer stored, then no processing is done.
	// Prevents processing before WM_NCCREATE and after WM_NCDESTROY.
	if pMe == nil {
		return hwnd.DefWindowProc(msg, wParam, lParam)
	}

	// Try to process the message with an user handler.
	parm := wm.Base{WParam: wParam, LParam: lParam}
	userRet, wasHandled := pMe.depot.windowDepotMsg.processMessage(msg, parm)
	if !wasHandled {
		userRet, wasHandled = pMe.depot.windowDepotNfy.processMessage(msg, parm)
	}

	// No further messages processed after this one.
	if msg == co.WM_NCDESTROY {
		pMe.hwnd.SetWindowLongPtr(co.GWLP_USERDATA, 0) // clear passed pointer
		pMe.hwnd = win.HWND(0)
	}

	if wasHandled {
		return userRet
	}
	return hwnd.DefWindowProc(msg, wParam, lParam) // message was not processed
}
