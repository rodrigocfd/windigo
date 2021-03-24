package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

const _WM_UI_THREAD = co.WM_APP + 0x3fff // sent by RunUiThread()

// Base to _WindowOptsBase and _WindowDlgBase.
type _WindowBase struct {
	hWnd           win.HWND
	events         _EventsNfy
	internalEvents _EventsInternal
}

func (me *_WindowBase) new() {
	me.hWnd = win.HWND(0)
	me.events.new()
	me.internalEvents.new()

	me.defaultMessages()
}

func (me *_WindowBase) Hwnd() win.HWND {
	return me.hWnd
}

func (me *_WindowBase) On() *_EventsNfy {
	if me.hWnd != 0 {
		panic("Cannot add event handling after the window is created.")
	}
	return &me.events
}

func (me *_WindowBase) internalOn() *_EventsInternal {
	return &me.internalEvents
}

func (me *_WindowBase) RunUiThread(userFunc func()) {
	// This method is analog to SendMessage (synchronous), but intended to be
	// called from another thread, so a callback function can, tunelled by
	// wndproc, run in the original thread of the window, thus allowing GUI
	// updates. This avoids the user to deal with a custom WM_ message.
	me.hWnd.SendMessage(_WM_UI_THREAD, win.WPARAM(_WM_UI_THREAD),
		win.LPARAM(unsafe.Pointer(&userFunc)))
}

func (me *_WindowBase) defaultMessages() {
	me.events.Wm(_WM_UI_THREAD, func(p wm.Any) uintptr { // handle our custom thread UI message
		if p.WParam == win.WPARAM(_WM_UI_THREAD) {
			pUserFunc := (*func())(unsafe.Pointer(p.LParam))
			(*pUserFunc)()
		}
		return 0
	})
}

func (me *_WindowBase) loadIcons(
	hInst win.HINSTANCE, iconId int) (hIcon16, hIcon32 win.HICON) {

	// Resource icons are automatically released by the system.
	hIcon16 = win.HICON(
		hInst.LoadImage(int32(iconId),
			co.IMAGE_ICON, 16, 16, co.LR_DEFAULTCOLOR))
	hIcon32 = win.HICON(
		hInst.LoadImage(int32(iconId),
			co.IMAGE_ICON, 32, 32, co.LR_DEFAULTCOLOR))
	return
}
