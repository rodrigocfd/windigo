//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/co"
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
	On() *WindowEvents

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

// Returns the value adjusted according to the current horizontal system DPI.
func DpiX(x int) int {
	initalGuiSetup()
	return x * dpiX / 96
}

// Returns the value adjusted according to the current vertical system DPI.
func DpiY(y int) int {
	initalGuiSetup()
	return y * dpiY / 96
}

// Returns the value adjusted according to the current system DPI.
func Dpi(x, y int) (int, int) {
	return DpiX(x), DpiY(y)
}

// Syntactic sugar to [win.TaskDialogIndirect] to display a message box
// indicating an error.
//
// Panics on error.
//
// Example:
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	ui.MsgError(
//		wndOwner,
//		"Title",
//		"Big caption above text",
//		"Here goes the text",
//	)
func MsgError(wnd Parent, title, caption, body string) {
	msgBuild(wnd, title, caption, body, co.TDICON_ERROR, "", false)
}

// Syntactic sugar to [win.TaskDialogIndirect] to display a message box
// indicating a warning.
//
// Panics on error.
//
// Example:
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	ui.MsgWarn(
//		wndOwner,
//		"Title",
//		"Big caption above text",
//		"Here goes the text",
//	)
func MsgWarn(wnd Parent, title, caption, body string) {
	msgBuild(wnd, title, caption, body, co.TDICON_WARNING, "", false)
}

// Syntactic sugar to [win.TaskDialogIndirect] to display a message box
// indicating a successful operation.
//
// Panics on error.
//
// Example:
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	ui.MsgOk(
//		wndOwner,
//		"Title",
//		"Big caption above text",
//		"Here goes the text",
//	)
func MsgOk(wnd Parent, title, caption, body string) {
	msgBuild(wnd, title, caption, body, co.TDICON_INFORMATION, "", false)
}

// Syntactic sugar to [win.TaskDialogIndirect] to display a message box prompting
// the user to choose "Ok" or "Cancel". The "Ok" text can be customized.
//
// Returns true if the user clicks "Ok".
//
// Panics on error.
//
// Example:
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	clickedOk := ui.MsgOkCancel(
//		wndOwner,
//		"Title",
//		"Big caption above text",
//		"Here goes the text",
//		"&Confirm",
//	)
//	if clickedOk {
//		// ...
//	}
func MsgOkCancel(wnd Parent, title, caption, body, okText string) bool {
	return msgBuild(wnd, title, caption, body, co.TDICON_WARNING, okText, true) == co.ID_OK
}

func msgBuild(
	wnd Parent,
	title, caption, body string,
	icon co.TDICON,
	okText string,
	hasCancel bool,
) co.ID {
	var hParent win.HWND
	if wnd != nil {
		hParent = wnd.Hwnd()
	}

	var commonButtons co.TDCBF
	var buttons []win.TASKDIALOG_BUTTON
	if hasCancel {
		finalOkText := okText
		if finalOkText == "" {
			finalOkText = "&OK"
		}
		buttons = []win.TASKDIALOG_BUTTON{
			{Id: co.ID_OK, Text: finalOkText},
			{Id: co.ID_CANCEL, Text: "&Cancel"},
		}
	} else {
		commonButtons = co.TDCBF_OK
	}

	ret, err := win.TaskDialogIndirect(
		&win.TASKDIALOGCONFIG{
			HwndParent:      hParent,
			WindowTitle:     title,
			MainInstruction: caption,
			Content:         body,
			HMainIcon:       win.TdcIconTdi(icon),
			CommonButtons:   commonButtons,
			Flags:           co.TDF_ALLOW_DIALOG_CANCELLATION | co.TDF_POSITION_RELATIVE_TO_WINDOW,
			Buttons:         buttons,
		},
	)

	if err != nil {
		panic(err)
	}
	return ret
}
