package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _PromptT struct{}

// Displays various modal prompts which the user must interact to.
//
// The methods are high-level wrappers to win.TaskDialogIndirect() function,
// which is a modern replacement to the old HWND.MessageBox() method.
var Prompt _PromptT

// Displays an error modal prompt with an OK button.
func (_PromptT) Error(
	parent AnyParent,
	title string, header win.StrOpt, body string) {

	Prompt.generate(parent, title, header, body,
		co.TDCBF_OK, co.TD_ICON_ERROR,
		win.StrOptNone(), win.StrOptNone())
}

// Displays an information modal prompt with an OK button.
func (_PromptT) Info(
	parent AnyParent,
	title string, header win.StrOpt, body string) {

	Prompt.generate(parent, title, header, body,
		co.TDCBF_OK, co.TD_ICON_INFORMATION,
		win.StrOptNone(), win.StrOptNone())
}

// Displays a question modal prompt with OK and Cancel buttons.
//
// Returns true if the user clicked OK.
func (_PromptT) OkCancel(
	parent AnyParent,
	title string, header win.StrOpt, body string) bool {

	return Prompt.generate(parent, title, header, body,
		co.TDCBF_OK|co.TDCBF_CANCEL, co.TD_ICON_WARNING,
		win.StrOptNone(), win.StrOptNone()) == co.ID_OK
}

// Displays a question modal prompt with OK and Cancel buttons. Custom texts may
// be specified for the buttons.
//
// Returns true if the user clicked OK.
func (_PromptT) OkCancelEx(
	parent AnyParent, title string, header win.StrOpt, body string,
	okText, cancelText win.StrOpt) bool {

	return Prompt.generate(parent, title, header, body,
		co.TDCBF(0), co.TD_ICON_WARNING,
		okText, cancelText) == co.ID_OK
}

// Displays a question modal prompt with Yes and No buttons.
//
// Returns true if the user clicked Yes.
func (_PromptT) YesNo(
	parent AnyParent,
	title string, header win.StrOpt, body string) bool {

	return Prompt.generate(parent, title, header, body,
		co.TDCBF_YES|co.TDCBF_NO, co.TD_ICON_WARNING,
		win.StrOptNone(), win.StrOptNone()) == co.ID_YES
}

// Displays a question modal prompt with Yes, No and Cancel buttons.
func (_PromptT) YesNoCancel(
	parent AnyParent,
	title string, header win.StrOpt, body string) co.ID {

	return Prompt.generate(parent, title, header, body,
		co.TDCBF_YES|co.TDCBF_NO|co.TDCBF_CANCEL, co.TD_ICON_WARNING,
		win.StrOptNone(), win.StrOptNone())
}

func (_PromptT) generate(
	parent AnyParent,
	title string, header win.StrOpt, body string,
	btns co.TDCBF, ico co.TD_ICON,
	okText, cancelText win.StrOpt) co.ID {

	var customBtns []win.TASKDIALOG_BUTTON
	if text, has := okText.Str(); has {
		customBtns = append(customBtns, win.TASKDIALOG_BUTTON{
			PszButtonText: text,
			NButtonID:     int32(co.ID_OK),
		})
	} else {
		btns |= co.TDCBF_OK
	}
	if text, has := cancelText.Str(); has {
		customBtns = append(customBtns, win.TASKDIALOG_BUTTON{
			PszButtonText: text,
			NButtonID:     int32(co.ID_CANCEL),
		})
	} else {
		btns |= co.TDCBF_CANCEL
	}

	tdc := win.TASKDIALOGCONFIG{
		DwFlags:         co.TDF_ALLOW_DIALOG_CANCELLATION | co.TDF_POSITION_RELATIVE_TO_WINDOW,
		DwCommonButtons: btns,
		PszWindowTitle:  title,
		HMainIcon:       win.TdcIconTdi(ico),
		PszContent:      body,
		PButtons:        customBtns,
	}
	if parent != nil {
		tdc.HwndParent = parent.Hwnd()
	}
	if header, ok := header.Str(); ok {
		tdc.PszMainInstruction = string(header)
	}

	return win.TaskDialogIndirect(&tdc)
}
