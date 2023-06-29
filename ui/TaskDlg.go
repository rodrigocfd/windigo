//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _TaskDlgT struct{}

// Displays various modal prompts which the user must interact to.
//
// The methods are high-level wrappers to [TaskDialogIndirect], which is a
// modern replacement to the old [MessageBox] method.
//
// [TaskDialogIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialogindirect
// [MessageBox]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
var TaskDlg _TaskDlgT

// Displays an error modal prompt with an OK button.
func (_TaskDlgT) Error(
	parent AnyParent,
	title string,
	header win.StrOpt,
	body string) {

	TaskDlg.generate(parent, title, header, body,
		co.TDCBF_OK, co.TD_ICON_ERROR, nil)
}

// Displays an information modal prompt with an OK button.
func (_TaskDlgT) Info(
	parent AnyParent,
	title string,
	header win.StrOpt,
	body string) {

	TaskDlg.generate(parent, title, header, body,
		co.TDCBF_OK, co.TD_ICON_INFORMATION, nil)
}

// Displays a question modal prompt with OK and Cancel buttons.
//
// Returns true if the user clicked OK.
func (_TaskDlgT) OkCancel(
	parent AnyParent,
	title string,
	header win.StrOpt,
	body string) bool {

	return TaskDlg.generate(parent, title, header, body,
		co.TDCBF_OK|co.TDCBF_CANCEL, co.TD_ICON_WARNING, nil) == co.ID_OK
}

// Displays a question modal prompt with OK and Cancel buttons. Custom texts may
// be specified for the buttons.
//
// Returns true if the user clicked OK.
func (_TaskDlgT) OkCancelEx(
	parent AnyParent,
	title string,
	header win.StrOpt,
	body string,
	okText, cancelText win.StrOpt) bool {

	var btns co.TDCBF
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

	return TaskDlg.generate(parent, title, header, body,
		btns, co.TD_ICON_WARNING, customBtns) == co.ID_OK
}

// Displays a question modal prompt with Yes and No buttons.
//
// Returns true if the user clicked Yes.
func (_TaskDlgT) YesNo(
	parent AnyParent,
	title string,
	header win.StrOpt,
	body string) bool {

	return TaskDlg.generate(parent, title, header, body,
		co.TDCBF_YES|co.TDCBF_NO, co.TD_ICON_WARNING, nil) == co.ID_YES
}

// Displays a question modal prompt with Yes, No and Cancel buttons.
func (_TaskDlgT) YesNoCancel(
	parent AnyParent,
	title string,
	header win.StrOpt,
	body string) co.ID {

	return TaskDlg.generate(parent, title, header, body,
		co.TDCBF_YES|co.TDCBF_NO|co.TDCBF_CANCEL, co.TD_ICON_WARNING, nil)
}

func (_TaskDlgT) generate(
	parent AnyParent,
	title string,
	header win.StrOpt,
	body string,
	btns co.TDCBF,
	ico co.TD_ICON,
	customBtns []win.TASKDIALOG_BUTTON) co.ID {

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
		tdc.PszMainInstruction = header
	}

	return win.TaskDialogIndirect(&tdc)
}
