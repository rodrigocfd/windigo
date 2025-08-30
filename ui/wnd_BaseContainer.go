//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
)

const _WM_UI_THREAD = 0xbfff // Internal message to run closures from another thread (last WM_APP value).

// Base to raw and dialog container windows.
type _BaseContainer struct {
	hWnd   win.HWND
	layout _Layout

	beforeUserEvents _EventsWindowLib
	userEvents       EventsWindow
	afterUserEvents  _EventsWindowLib
}

// Constructor.
func newBaseContainer(wndTy _WNDTY) _BaseContainer {
	return _BaseContainer{
		hWnd:   win.HWND(0),
		layout: newLayout(),

		beforeUserEvents: newEventsWindowLib(),
		userEvents:       newEventsWindow(wndTy),
		afterUserEvents:  newEventsWindowLib(),
	}
}

func (me *_BaseContainer) removeWmCreateInitdialog() {
	me.beforeUserEvents.removeWmCreateInitdialog()
	me.userEvents.removeWmCreateInitdialog()
	me.afterUserEvents.removeWmCreateInitdialog()
}
func (me *_BaseContainer) clearMessages() {
	me.beforeUserEvents.clear()
	me.userEvents.clear()
	me.afterUserEvents.clear()
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
	me.beforeUserEvents.wm(_WM_UI_THREAD, func(p Wm) {
		if p.WParam == win.WPARAM(_WM_UI_THREAD) { // additional safety check
			pPack := (*_ThreadPack)(unsafe.Pointer(p.LParam))
			utl.PtrCache.Delete(unsafe.Pointer(pPack)) // now GC will be able to collect it
			pPack.fun()
		}
	})

	me.beforeUserEvents.wm(co.WM_SIZE, func(p Wm) {
		me.layout.Rearrange(WmSize{p})
	})
}

// Runs the message loop for the main window, blocking until it's closed.
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

// Runs the message loop for a modal window, blocking until it's closed.
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
