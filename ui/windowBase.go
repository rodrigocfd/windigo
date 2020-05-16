/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

// Base to all window types.
type windowBase struct {
	hwnd   api.HWND
	wndMsg windowMsg
}

const wM_UI_THREAD = c.WM_APP + 0x3FFF // used in UI thread handling
type threadPack struct{ userFunc func() }

// Returns the underlying HWND handle of this window.
func (me *windowBase) Hwnd() api.HWND {
	return me.hwnd
}

// Exposes all the window messages the can be handled.
func (me *windowBase) OnMsg() *windowMsg {
	return &me.wndMsg
}

// Runs a closure synchronously in the window original UI thread.
// When in a goroutine, you *MUST* use this method to update the UI.
func (me *windowBase) RunUiThread(userFunc func()) {
	// This method is analog to SendMessage (synchronous), but intended to be
	// called from another thread, so a callback function can, tunelled by
	// wndproc, run in the original thread of the window, thus allowing GUI
	// updates. This avoids the user to deal with a custom WM_ message.
	pack := &threadPack{userFunc: userFunc}
	me.hwnd.SendMessage(wM_UI_THREAD, 0xC0DEF00D, api.LPARAM(unsafe.Pointer(pack)))
}

func (me *windowBase) createWindow(uiName string, exStyle c.WS_EX,
	className, title string, style c.WS, x, y int32, width, height uint32,
	parent Window, menu api.HMENU, hInst api.HINSTANCE) {

	if me.hwnd != 0 {
		panic(fmt.Sprintf("Trying to create %s \"%s\" twice.",
			uiName, title))
	}

	hwndParent := api.HWND(0) // if no parent, pass zero to CreateWindowEx
	if parent != nil {
		hwndParent = parent.Hwnd()
	}

	me.wndMsg.addMsg(wM_UI_THREAD, func(p wmBase) uintptr { // handle our custom thread UI message
		if p.WParam == 0xC0DEF00D {
			pack := (*threadPack)(unsafe.Pointer(p.LParam))
			pack.userFunc()
		}
		return 0
	})

	// The hwnd member is saved in WM_NCCREATE processing in wndProc.
	api.CreateWindowEx(exStyle, className, title, style, x, y, width, height,
		hwndParent, menu, hInst, unsafe.Pointer(me)) // pass pointer to our object
}

func (me *windowBase) registerClass(wcx *api.WNDCLASSEX) api.ATOM {
	wcx.LpfnWndProc = syscall.NewCallback(wndProc)
	atom, lerr := wcx.RegisterClassEx()
	if lerr != 0 {
		// https://devblogs.microsoft.com/oldnewthing/20150429-00/?p=44984
		// https://devblogs.microsoft.com/oldnewthing/20041011-00/?p=37603
		if c.ERROR(lerr) == c.ERROR_CLASS_ALREADY_EXISTS {
			atom = wcx.HInstance.GetClassInfoEx(wcx.LpszClassName, wcx)
		} else {
			panic(fmt.Sprintf("RegisterClassEx failed with atom %d: %d %s\n",
				atom, lerr, lerr.Error()))
		}
	}
	return atom
}

func wndProc(hwnd api.HWND, msg c.WM, wParam api.WPARAM, lParam api.LPARAM) uintptr {
	if msg == c.WM_NCCREATE {
		cs := (*api.CREATESTRUCT)(unsafe.Pointer(lParam))
		base := (*windowBase)(unsafe.Pointer(cs.LpCreateParams))
		hwnd.SetWindowLongPtr(c.GWLP_USERDATA, uintptr(unsafe.Pointer(base)))
		base.hwnd = hwnd // assign actual HWND
	}

	// Retrieve passed pointer.
	pMe := (*windowBase)(unsafe.Pointer(hwnd.GetWindowLongPtr(c.GWLP_USERDATA)))

	// Save *windowBase from being collected by GC.
	hwnd.SetWindowLongPtr(c.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))

	// If no pointer stored, then no processing is done.
	// Prevents processing before WM_NCCREATE and after WM_NCDESTROY.
	if pMe == nil {
		return hwnd.DefWindowProc(msg, wParam, lParam)
	}

	// Try to process the message with an user handler.
	userResult, wasProcessed := pMe.wndMsg.processMessage(msg,
		wmBase{WParam: wParam, LParam: lParam})

	// No further messages processed after this one.
	if msg == c.WM_NCDESTROY {
		pMe.hwnd.SetWindowLongPtr(c.GWLP_USERDATA, 0) // clear passed pointer
		pMe.hwnd = api.HWND(0)
	}

	if wasProcessed {
		return userResult
	}
	return hwnd.DefWindowProc(msg, wParam, lParam) // message was not processed
}
