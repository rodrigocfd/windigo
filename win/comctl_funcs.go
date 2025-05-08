//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [InitCommonControls] function.
//
// [InitCommonControls]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrols
func InitCommonControls() {
	syscall.SyscallN(_InitCommonControls.Addr())
}

var _InitCommonControls = dll.Comctl32.NewProc("InitCommonControls")

// [InitCommonControlsEx] function.
//
// [InitCommonControlsEx]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrolsex
func InitCommonControlsEx(icc co.ICC) error {
	var iccx _INITCOMMONCONTROLSEX
	iccx.SetDwSize()
	iccx.DwICC = icc

	ret, _, _ := syscall.SyscallN(_InitCommonControlsEx.Addr(),
		uintptr(unsafe.Pointer(&iccx)))
	return util.ZeroAsSysInvalidParm(ret)
}

var _InitCommonControlsEx = dll.Comctl32.NewProc("InitCommonControlsEx")

// [InitMUILanguage] function.
//
// [InitMUILanguage]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initmuilanguage
func InitMUILanguage(lang LANGID) {
	syscall.SyscallN(_InitMUILanguage.Addr(),
		uintptr(lang))
}

var _InitMUILanguage = dll.Comctl32.NewProc("InitMUILanguage")

// [TaskDialogIndirect] function.
//
// # Example
//
//	var hWnd win.HWND // initialized somewhere
//
//	win.TaskDialogIndirect(win.TASKDIALOGCONFIG{
//		HwndParent:      hWnd,
//		WindowTitle:     "Title",
//		MainInstruction: "Caption",
//		Content:         "Body",
//		HMainIcon:       win.TdcIconTdi(co.TDICON_INFORMATION),
//		CommonButtons:   co.TDCBF_OK,
//		Flags: co.TDF_ALLOW_DIALOG_CANCELLATION |
//			co.TDF_POSITION_RELATIVE_TO_WINDOW,
//	})
//
// [TaskDialogIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialogindirect
func TaskDialogIndirect(taskConfig TASKDIALOGCONFIG) (co.ID, error) {
	strs16 := wstr.NewArray()

	tdcBuf := NewVecSized(160, byte(0)) // packed TASKDIALOGCONFIG is 160 bytes
	defer tdcBuf.Free()

	btnsBuf := NewVec[[12]byte]() // packed TASKDIALOG_BUTTON is 12 bytes
	defer btnsBuf.Free()

	taskConfig.serialize(&strs16, &tdcBuf, &btnsBuf)

	pPnButton := NewVecSized(1, int32(0)) // OS-allocated; value to be returned
	defer pPnButton.Free()

	ret, _, _ := syscall.SyscallN(_TaskDialogIndirect.Addr(),
		uintptr(tdcBuf.UnsafePtr()), uintptr(pPnButton.UnsafePtr()))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return co.ID(0), hr
	}

	return co.ID(*pPnButton.Get(0)), nil
}

var _TaskDialogIndirect = dll.Comctl32.NewProc("TaskDialogIndirect")
