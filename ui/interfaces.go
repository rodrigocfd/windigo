//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/win"
)

// Any window.
type Window interface {
	// Returns the underlying HWND handle of this window.
	//
	// Note that this handle is initially zero, existing only after window creation.
	Hwnd() win.HWND
}

// A child control window.
type ChildControl interface {
	Window

	// Returns the control ID, unique within the same Parent.
	CtrlId() uint16

	// If parent is a dialog, sets the focus by sending [WM_NEXTDLGCTL]. This
	// draws the borders correctly in some undefined controls, like buttons.
	//
	// Otherwise, calls [SetFocus].
	//
	// [WM_NEXTDLGCTL]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-nextdlgctl
	// [SetFocus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setfocus
	Focus()
}

// A parent window.
type Parent interface {
	Window

	// Exposes all the window notifications the can be handled.
	//
	// Panics if called after the window has been created.
	On() *EventsWindow

	// This method is analog to [SendMessage] (synchronous), but intended to be
	// called from another thread, so a callback function can, tunelled by
	// [WNDPROC], run in the original thread of the window, thus allowing GUI
	// updates. With this, the user doesn't have to deal with a custom WM_
	// message.
	//
	// Example:
	//
	//	var wnd ui.Parent // initialized somewhere
	//
	//	wnd.On().WmCreate(func(_ WmCreate) int {
	//		go func() {
	//			// process to be done in a parallel goroutine...
	//
	//			wnd.UiThread(func() {
	//				// update the UI in the original UI thread...
	//			})
	//		}()
	//		return 0
	//	})
	//
	// [SendMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew
	// [WNDPROC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-wndproc
	UiThread(fun func())

	base() *_BaseContainer
}
