/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Base to all window types: WindowControl, WindowMain and WindowModal.
type _WindowBase struct {
	hwnd  win.HWND
	depot _DepotWmCmdNfy
}

const _WM_UI_THREAD = co.WM_APP + 0x3FFF   // used in UI thread handling
type _ThreadPack struct{ UserFunc func() } // transports user closure

// Returns the underlying HWND handle of this window.
func (me *_WindowBase) Hwnd() win.HWND {
	return me.hwnd
}

// Exposes all the window messages the can be handled.
func (me *_WindowBase) OnMsg() *_DepotWmCmdNfy {
	if me.Hwnd() != 0 {
		panic("Cannot add message after the window was created.")
	}
	return &me.depot
}

// Runs a closure synchronously in the window original UI thread.
//
// When in a goroutine, you *MUST* use this method to update the UI.
func (me *_WindowBase) RunUiThread(userFunc func()) {
	// This method is analog to SendMessage (synchronous), but intended to be
	// called from another thread, so a callback function can, tunelled by
	// wndproc, run in the original thread of the window, thus allowing GUI
	// updates. This avoids the user to deal with a custom WM_ message.
	pack := &_ThreadPack{UserFunc: userFunc}
	me.hwnd.SendMessage(_WM_UI_THREAD, 0xc0def00d,
		win.LPARAM(unsafe.Pointer(pack)))
}

func (me *_WindowBase) registerClass(wcx *win.WNDCLASSEX) win.ATOM {
	wcx.LpfnWndProc = syscall.NewCallback(wndProc)
	atom, err := win.RegisterClassEx(wcx)

	if err != nil && err.Code() == co.ERROR_CLASS_ALREADY_EXISTS {
		// https://devblogs.microsoft.com/oldnewthing/20150429-00/?p=44984
		// https://devblogs.microsoft.com/oldnewthing/20041011-00/?p=37603
		atom, err = wcx.HInstance.GetClassInfoEx(
			(*uint16)(unsafe.Pointer(wcx.LpszClassName)), wcx)

		if err != nil {
			panic(fmt.Sprintf("GetClassInfoEx failed. %s", err.Error()))
		}

	} else if err != nil && err.Code() != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegisterClassEx failed. %s", err.Error()))
	}

	return atom
}

func (me *_WindowBase) createWindow(uiName string, exStyle co.WS_EX,
	className, title string, style co.WS, x, y int, width, height uint,
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
	win.CreateWindowEx(exStyle, className, title, style,
		int32(x), int32(y), int32(width), int32(height),
		hwndParent, menu, hInst, unsafe.Pointer(me)) // pass pointer to our object
}

func (me *_WindowBase) defaultMessageHandling() {
	me.depot.Wm(_WM_UI_THREAD, func(p Wm) uintptr { // handle our custom thread UI message
		if p.WParam == 0xc0def00d {
			pack := (*_ThreadPack)(unsafe.Pointer(p.LParam))
			pack.UserFunc()
		}
		return 0
	})
}

func wndProc(hwnd win.HWND, msg co.WM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	// https://devblogs.microsoft.com/oldnewthing/20050422-08/?p=35813
	if msg == co.WM_NCCREATE {
		cs := (*win.CREATESTRUCT)(unsafe.Pointer(lParam))
		base := (*_WindowBase)(unsafe.Pointer(cs.LpCreateParams))
		hwnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(base)))
		base.hwnd = hwnd // assign actual HWND
	}

	// Retrieve passed pointer.
	pMe := (*_WindowBase)(unsafe.Pointer(
		hwnd.GetWindowLongPtr(co.GWLP_USERDATA)))

	// If the retrieved *_WindowBase stays here, the GC will collect it.
	// Sending it away will prevent the GC collection.
	// https://stackoverflow.com/a/51188315
	hwnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))

	// If no pointer stored, then no processing is done.
	// Prevents processing before WM_NCCREATE and after WM_NCDESTROY.
	if pMe == nil {
		return hwnd.DefWindowProc(msg, wParam, lParam)
	}

	// Try to process the message with an user handler.
	userRet, wasHandled := pMe.depot.processMessage(msg,
		Wm{WParam: wParam, LParam: lParam})

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
