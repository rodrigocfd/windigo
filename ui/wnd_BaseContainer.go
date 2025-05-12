//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

const _WM_UI_THREAD = co.WM_APP + 0x3fff // Internal message to run closures from another thread.

// Base to raw and dialog container windows.
type _BaseContainer struct {
	hWnd   win.HWND
	wndTy  _WNDTY
	layout _Layout

	beforeUserEvents EventsWindow
	userEvents       EventsWindow
	afterUserEvents  EventsWindow
}

// Constructor.
func newBaseContainer(wndTy _WNDTY) _BaseContainer {
	return _BaseContainer{
		hWnd:   win.HWND(0),
		wndTy:  wndTy,
		layout: newLayout(),

		beforeUserEvents: newEventsWindow(_WNDTY_DLG),
		userEvents:       newEventsWindow(_WNDTY_DLG),
		afterUserEvents:  newEventsWindow(_WNDTY_DLG),
	}
}

func (me *_BaseContainer) clearMessages() {
	me.beforeUserEvents.clear()
	me.userEvents.clear()
	me.afterUserEvents.clear()
}
func (me *_BaseContainer) removeWmCreateInitdialog() {
	me.beforeUserEvents.removeWmCreateInitdialog()
	me.userEvents.removeWmCreateInitdialog()
	me.afterUserEvents.removeWmCreateInitdialog()
}

func (me *_BaseContainer) uiThread(fun func()) {
	pPack := &_ThreadPack{fun}
	utl.PtrCache.Add(unsafe.Pointer(pPack))

	hWndRoot, _ := me.hWnd.GetAncestor(co.GA_ROOTOWNER)
	hWndRoot.SendMessage(_WM_UI_THREAD,
		win.WPARAM(_WM_UI_THREAD), win.LPARAM(unsafe.Pointer(pPack)))
}

type _ThreadPack struct{ fun func() }

func (me *_BaseContainer) defaultMessageHandlers() {
	me.beforeUserEvents.Wm(_WM_UI_THREAD, func(p Wm) uintptr {
		if p.WParam == win.WPARAM(_WM_UI_THREAD) { // additional safety check
			pPack := (*_ThreadPack)(unsafe.Pointer(p.LParam))
			utl.PtrCache.Delete(unsafe.Pointer(pPack)) // now GC will be able to collect it
			pPack.fun()
		}
		return 0 // ignored
	})

	me.beforeUserEvents.WmSize(func(p WmSize) {
		me.layout.Rearrange(p)
	})
}

func (me *_BaseContainer) runMainLoop(hAccel win.HACCEL, processDlgMsgs bool) int {
	vecMsg := win.NewVecSized(1, win.MSG{})
	defer vecMsg.Free()
	pMsg := vecMsg.Get(0) // OS-allocated

	for {
		if res, err := win.GetMessage(pMsg, win.HWND(0), 0, 0); err != nil {
			panic(err)
		} else if res == 0 {
			// WM_QUIT was sent, gracefully terminate the program.
			// If it returned -1, it will simply panic.
			// WParam has the program exit code.
			// https://learn.microsoft.com/en-us/windows/win32/winmsg/using-messages-and-message-queues
			return int(pMsg.WParam)
		}

		// If a child window, will retrieve its top-level parent.
		// If a top-level, use itself.
		hTopLevel, _ := pMsg.HWnd.GetAncestor(co.GA_ROOT)
		if hTopLevel == 0 {
			hTopLevel = pMsg.HWnd
		}

		// If we have an accelerator table, try to translate the message.
		if hAccel != 0 && hTopLevel.TranslateAccelerator(hAccel, pMsg) == nil {
			continue // message translated
		}

		// Try to process keyboard actions for child controls.
		if processDlgMsgs && hTopLevel.IsDialogMessage(pMsg) {
			continue
		}

		win.TranslateMessage(pMsg)
		win.DispatchMessage(pMsg)
	}
}

func (me *_BaseContainer) runModalLoop(processDlgMsgs bool) {
	vecMsg := win.NewVecSized(1, win.MSG{})
	defer vecMsg.Free()
	pMsg := vecMsg.Get(0) // OS-allocated

	for {
		if res, err := win.GetMessage(pMsg, win.HWND(0), 0, 0); err != nil {
			panic(err)
		} else if res == 0 {
			return // our modal was destroyed, terminate loop
		}

		if me.hWnd == 0 || !me.hWnd.IsWindow() {
			return // our modal was destroyed, terminate loop
		}

		// If a child window, will retrieve its top-level parent.
		// If a top-level, use itself.
		hWndTopLevel, _ := pMsg.HWnd.GetAncestor(co.GA_ROOT)
		if hWndTopLevel != 0 {
			hWndTopLevel = me.hWnd
		}

		// Try to process keyboard actions for child controls.
		if processDlgMsgs && hWndTopLevel.IsDialogMessage(pMsg) {
			// Processed all keyboard actions for child controls.
			if me.hWnd == 0 {
				return // our modal was destroyed, terminate loop
			} else {
				continue
			}
		}

		win.TranslateMessage(pMsg)
		win.DispatchMessage(pMsg)

		if me.hWnd == 0 || !me.hWnd.IsWindow() {
			return // our modal was destroyed, terminate loop
		}
	}
}
