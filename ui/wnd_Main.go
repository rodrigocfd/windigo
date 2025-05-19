//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Main application window.
//
// Implements:
//   - [Window]
//   - [Parent]
type Main struct {
	raw *_MainRaw
	dlg *_MainDlg
}

// Creates a new main window with [CreateWindowEx].
//
// # Example
//
//	runtime.LockOSThread()
//
//	wnd := ui.NewMain(
//		ui.OptsMain().
//			Title("Hello world").
//			Size(ui.Dpi(500, 400)).
//			Style(co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN |
//				co.WS_BORDER | co.WS_VISIBLE | co.WS_MINIMIZEBOX |
//				co.WS_MAXIMIZEBOX | co.WS_SIZEBOX),
//	)
//	wnd.RunAsMain()
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func NewMain(opts *VarOptsMain) *Main {
	return &Main{
		raw: newMainRaw(opts),
		dlg: nil,
	}
}

// Creates a new dialog-based Main with [CreateDialogParam].
//
// # Example
//
//	const (
//		ID_MAIN_DLG    uint16 = 1000
//		ID_MAIN_ICON   uint16 = 101
//		ID_MAIN_ACCTBL uint16 = 102
//	)
//
//	runtime.LockOSThread()
//
//	wnd := ui.NewMainDlg(
//		ui.OptsMainDlg().
//			DlgId(ID_MAIN_DLG).
//			IconId(ID_MAIN_ICON).
//			AccelTableId(ID_MAIN_ACCTBL),
//	)
//	wnd.RunAsMain()
//
// [CreateDialogParam]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createdialogparamw
func NewMainDlg(opts *VarOptsMainDlg) *Main {
	return &Main{
		raw: nil,
		dlg: newMainDlg(opts),
	}
}

// Physically creates the window, then runs the main application loop. This
// method will block until the window is closed.
//
// Panics on error.
func (me *Main) RunAsMain() int {
	if win.IsWindowsVistaOrGreater() {
		if err := win.SetProcessDPIAware(); err != nil {
			panic(err)
		}
	}

	win.InitCommonControls()

	if win.IsWindows8OrGreater() {
		bVal := int32(0) // BOOL=FALSE; SetTimer() safety
		err := win.GetCurrentProcess().SetUserObjectInformation(
			co.UOI_TIMERPROC_EXCEPTION_SUPPRESSION, unsafe.Pointer(&bVal), unsafe.Sizeof(bVal))
		if err != nil {
			if wErr, _ := err.(co.ERROR); wErr != co.ERROR_INVALID_PARAMETER {
				// Wine doesn't support SetUserObjectInformation() and returns
				// INVALID_PARAMETER, so we let it pass. Otherwise, we crash.
				// https://bugs.winehq.org/show_bug.cgi?id=54951
				panic(wErr)
			}
		}
	}

	createGlobalUiFont() // will be applied to native controls
	defer globalUiFont.DeleteObject()

	hInst, _ := win.GetModuleHandle("")
	if me.raw != nil {
		return me.raw.runAsMain(hInst)
	} else {
		return me.dlg.runAsMain(hInst)
	}
}

// Returns the underlying HWND handle of this window.
//
// Implements [Window].
//
// Note that this handle is initially zero, existing only after window creation.
func (me *Main) Hwnd() win.HWND {
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
func (me *Main) On() *EventsWindow {
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
//	var wnd *ui.WindowMain // initialized somewhere
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
func (me *Main) UiThread(fun func()) {
	if me.raw != nil {
		me.raw.uiThread(fun)
	} else {
		me.dlg.uiThread(fun)
	}
}

// Implements [Parent].
func (me *Main) base() *_BaseContainer {
	if me.raw != nil {
		return &me.raw._BaseContainer
	} else {
		return &me.dlg._BaseContainer
	}
}
