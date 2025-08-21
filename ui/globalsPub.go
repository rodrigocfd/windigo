//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Global system DPI factor.
var dpiX, dpiY int = 0, 0

func cacheSystemDpi() error {
	if dpiX == 0 || dpiY == 0 { // not cached yet?
		hdcScreen, err := win.HWND(0).GetDC()
		if err != nil {
			return err
		}
		defer win.HWND(0).ReleaseDC(hdcScreen)

		dpiX = int(hdcScreen.GetDeviceCaps(co.GDC_LOGPIXELSX))
		dpiY = int(hdcScreen.GetDeviceCaps(co.GDC_LOGPIXELSY))
	}
	return nil
}

// Returns the value adjusted according to the current horizontal system DPI.
func DpiX(x int) int {
	cacheSystemDpi()
	return x * dpiX / 96
}

// Returns the value adjusted according to the current vertical system DPI.
func DpiY(y int) int {
	cacheSystemDpi()
	return y * dpiY / 96
}

// Returns the value adjusted according to the current system DPI.
func Dpi(x, y int) (int, int) {
	return DpiX(x), DpiY(y)
}

// Syntactic sugar to [TaskDialogIndirect] to display a message box indicating
// an error.
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
//
// [TaskDialogIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialogindirect
func MsgError(wnd Parent, title, caption, body string) {
	msgBuild(wnd, title, caption, body, co.TDICON_ERROR, "", false)
}

// Syntactic sugar to [TaskDialogIndirect] to display a message box indicating
// a warning.
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
//
// [TaskDialogIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialogindirect
func MsgWarn(wnd Parent, title, caption, body string) {
	msgBuild(wnd, title, caption, body, co.TDICON_WARNING, "", false)
}

// Syntactic sugar to [TaskDialogIndirect] to display a message box indicating
// a successful operation.
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
//
// [TaskDialogIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialogindirect
func MsgOk(wnd Parent, title, caption, body string) {
	msgBuild(wnd, title, caption, body, co.TDICON_INFORMATION, "", false)
}

// Syntactic sugar to [TaskDialogIndirect] to display a message box prompting
// the user to choose "Ok" or "Cancel". The "Ok" text can be customized.
//
// Returns co.ID_OK or co.ID_CANCEL.
//
// Panics on error.
//
// Example:
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	ret := ui.MsgOkCancel(
//		wndOwner,
//		"Title",
//		"Big caption above text",
//		"Here goes the text",
//		"&Confirm",
//	)
//	if ret == co.ID_OK {
//		// ...
//	}
//
// [TaskDialogIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialogindirect
func MsgOkCancel(wnd Parent, title, caption, body, okText string) co.ID {
	return msgBuild(wnd, title, caption, body, co.TDICON_INFORMATION, okText, true)
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

	ret, err := win.TaskDialogIndirect(win.TASKDIALOGCONFIG{
		HwndParent:      hParent,
		WindowTitle:     title,
		MainInstruction: caption,
		Content:         body,
		HMainIcon:       win.TdcIconTdi(icon),
		CommonButtons:   commonButtons,
		Flags:           co.TDF_ALLOW_DIALOG_CANCELLATION | co.TDF_POSITION_RELATIVE_TO_WINDOW,
		Buttons:         buttons,
	})

	if err != nil {
		panic(err)
	}
	return ret
}
