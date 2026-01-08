//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

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
