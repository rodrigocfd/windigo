//go:build windows

package ui

import (
	"sync"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

const _WM_UI_THREAD = co.WM_APP + 0x3fff // Sent by RunUiThread().
var (
	_globalUiThreadCache = make(map[int]func(), 20) // User functions of RunUiThread().
	_globalUiThreadCount = 0
	_globalUiThreadMutex = sync.Mutex{}
)

// Base to _WindowRaw and _WindowDlg; the root of all parent windows.
type _WindowBase struct {
	hWnd            win.HWND
	internalEvents  _EventsInternal // Events added internally by the library.
	events          _EventsWmNfy    // Ordinary window events, added by user.
	resizerChildren _ResizerChildren
}

func (me *_WindowBase) new() {
	me.hWnd = win.HWND(0)
	me.internalEvents.new()
	me.events.new()
	me.resizerChildren.new()

	me.defaultMessages()
}

// Implements AnyWindow.
func (me *_WindowBase) Hwnd() win.HWND {
	return me.hWnd
}

// Implements AnyParent.
func (me *_WindowBase) On() *_EventsWmNfy {
	if me.hWnd != 0 {
		panic("Cannot add event handling after the window is created.")
	}
	return &me.events
}

// Implements AnyParent.
func (me *_WindowBase) internalOn() *_EventsInternal {
	return &me.internalEvents
}

// Implements AnyParent.
func (me *_WindowBase) addResizingChild(ctrl AnyControl, horz HORZ, vert VERT) {
	// Must be called after the control was created, because its HWND is used.
	me.resizerChildren.add(me.Hwnd(), ctrl, horz, vert)
}

// Implements AnyParent.
func (me *_WindowBase) RunUiThread(userFunc func()) {
	// This method is analog to SendMessage (synchronous), but intended to be
	// called from another thread, so a callback function can, tunelled by
	// wndproc, run in the original thread of the window, thus allowing GUI
	// updates. This avoids the user to deal with a custom WM_ message.

	_globalUiThreadMutex.Lock()
	_globalUiThreadCount++
	_globalUiThreadCache[_globalUiThreadCount] = userFunc // cache
	_globalUiThreadMutex.Unlock()

	// Bypass any modals and send straight to main window. This avoids any blind
	// spots of unhandled messages by a modal being created/destroyed.
	me.hWnd.GetAncestor(co.GA_ROOTOWNER).
		SendMessage(_WM_UI_THREAD,
			win.WPARAM(_WM_UI_THREAD), win.LPARAM(_globalUiThreadCount))
}

func (me *_WindowBase) defaultMessages() {
	me.internalOn().addMsgZero(_WM_UI_THREAD, func(p wm.Any) { // handle our custom thread UI message
		if p.WParam == win.WPARAM(_WM_UI_THREAD) { // additional safety check
			_globalUiThreadMutex.Lock()
			userFunc := _globalUiThreadCache[int(p.LParam)] // retrieve from cache
			delete(_globalUiThreadCache, int(p.LParam))     // clear from cache
			_globalUiThreadMutex.Unlock()

			userFunc()
		}
	})

	me.internalOn().addMsgZero(co.WM_SIZE, func(p wm.Any) {
		me.resizerChildren.resizeChildren(wm.Size{Msg: p})
	})
}

func (me *_WindowBase) loadIcons(
	hInst win.HINSTANCE, iconId int) (hIcon16, hIcon32 win.HICON) {

	// Resource icons are automatically released by the system.
	hIcon16 = win.HICON(
		hInst.LoadImage(win.ResIdInt(iconId), co.IMAGE_ICON, 16, 16, co.LR_DEFAULTCOLOR))
	hIcon32 = win.HICON(
		hInst.LoadImage(win.ResIdInt(iconId), co.IMAGE_ICON, 32, 32, co.LR_DEFAULTCOLOR))
	return
}
