//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/win"
)

// Modal window.
//
// Implements:
//   - [Window]
//   - [Parent]
type Modal struct {
	raw *_RawModal
	dlg *_DlgModal
}

// Creates a new modal window with [CreateWindowEx].
//
// # Example
//
//	var wndParent ui.Parent // initialized somewhere
//
//	wndModal := ui.NewModal(
//		wndParent,
//		ui.OptsModal().
//			Title("Hello modal"),
//	)
//	wndModal.ShowModal()
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func NewModal(parent Parent, opts *VarOptsModal) *Modal {
	return &Modal{
		raw: newModalRaw(parent, opts),
		dlg: nil,
	}
}

// Creates a new dialog-based Modal with [DialogBoxParam].
//
// # Example
//
//	const ID_MODAL_DLG uint16 = 2000
//
//	var wndParent ui.Parent // initialized somewhere
//
//	wndModal := ui.NewModalDlg(wndParent, 2000)
//	wndModal.ShowModal()
//
// [DialogBoxParam]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dialogboxparamw
func NewModalDlg(parent Parent, dlgId uint16) *Modal {
	return &Modal{
		raw: nil,
		dlg: newModalDlg(parent, dlgId),
	}
}

// Physically creates the window, then runs the modal loop. This method will
// block until the window is closed.
func (me *Modal) ShowModal() {
	if me.raw != nil {
		me.raw.showModal()
	} else {
		me.dlg.showModal()
	}
}

// Returns the underlying HWND handle of this window.
//
// Implements [Window].
//
// Note that this handle is initially zero, existing only after window creation.
func (me *Modal) Hwnd() win.HWND {
	if me.raw != nil {
		return me.raw.hWnd
	} else {
		return me.dlg.hWnd
	}
}

// Exposes all the window notifications the can be handled.
//
// Implements [Parent].
//
// Panics if called after the window has been created.
func (me *Modal) On() *EventsWindow {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the window has been created.")
	}

	if me.raw != nil {
		return &me.raw.userEvents
	} else {
		return &me.dlg.userEvents
	}
}

// This method is analog to [SendMessage] (synchronous), but intended to be
// called from another thread, so a callback function can, tunelled by
// [WNDPROC], run in the original thread of the window, thus allowing GUI
// updates. With this, the user doesn't have to deal with a custom WM_ message.
//
// Implements [Parent].
//
// # Example
//
//	var wnd *ui.WindowModal // initialized somewhere
//
//	wnd.On().WmCreate(func(p WmCreate) int {
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
func (me *Modal) UiThread(fun func()) {
	if me.raw != nil {
		me.raw.uiThread(fun)
	} else {
		me.dlg.uiThread(fun)
	}
}

// Implements [Parent].
func (me *Modal) base() *_BaseContainer {
	if me.raw != nil {
		return &me.raw._BaseContainer
	} else {
		return &me.dlg._BaseContainer
	}
}
