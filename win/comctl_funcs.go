//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [ImageList_DragMove] function.
//
// [ImageList_DragMove]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_dragmove
func ImageListDragMove(x, y int) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_comctl_ImageList_DragMove, "ImageList_DragMove"),
		uintptr(int32(x)),
		uintptr(int32(y)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _comctl_ImageList_DragMove *syscall.Proc

// [ImageList_DragShowNolock] function.
//
// [ImageList_DragShowNolock]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_dragshownolock
func ImageListDragShowNolock(show bool) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_comctl_ImageList_DragShowNolock, "ImageList_DragShowNolock"),
		utl.BoolToUintptr(show))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _comctl_ImageList_DragShowNolock *syscall.Proc

// [ImageList_DrawIndirect] function.
//
// [ImageList_DrawIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_drawindirect
func ImageListDrawIndirect(imldp *IMAGELISTDRAWPARAMS) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_comctl_ImageList_DrawIndirect, "ImageList_DrawIndirect"),
		uintptr(unsafe.Pointer(imldp)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _comctl_ImageList_DrawIndirect *syscall.Proc

// [ImageList_EndDrag] function.
//
// [ImageList_EndDrag]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_enddrag
func ImageListEndDrag() {
	syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_comctl_ImageList_EndDrag, "ImageList_EndDrag"))
}

var _comctl_ImageList_EndDrag *syscall.Proc

// [InitCommonControls] function.
//
// [InitCommonControls]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrols
func InitCommonControls() {
	syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_comctl_InitCommonControls, "InitCommonControls"))
}

var _comctl_InitCommonControls *syscall.Proc

// [InitCommonControlsEx] function.
//
// [InitCommonControlsEx]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrolsex
func InitCommonControlsEx(icc co.ICC) error {
	var iccx _INITCOMMONCONTROLSEX
	iccx.SetDwSize()
	iccx.DwICC = icc

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_comctl_InitCommonControlsEx, "InitCommonControlsEx"),
		uintptr(unsafe.Pointer(&iccx)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _comctl_InitCommonControlsEx *syscall.Proc

// [InitMUILanguage] function.
//
// [InitMUILanguage]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initmuilanguage
func InitMUILanguage(lang LANGID) {
	syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_comctl_InitMUILanguage, "InitMUILanguage"),
		uintptr(lang))
}

var _comctl_InitMUILanguage *syscall.Proc

// [TaskDialogIndirect] function.
//
// Example:
//
//	var hWnd win.HWND // initialized somewhere
//
//	_, _ = win.TaskDialogIndirect(win.TASKDIALOGCONFIG{
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
	const TASKDIALOGCONFIG_SZ = 160
	const TASKDIALOG_BUTTON_SZ = 12

	// Sizes of all texts written by the user, counting terminating nulls.
	totTextUtf16Words := tdiStrLenIfAny(taskConfig.WindowTitle) +
		tdiStrLenIfAny(taskConfig.MainInstruction) +
		tdiStrLenIfAny(taskConfig.Content) +
		tdiStrLenIfAny(taskConfig.VerificationText) +
		tdiStrLenIfAny(taskConfig.ExpandedInformation) +
		tdiStrLenIfAny(taskConfig.ExpandedControlText) +
		tdiStrLenIfAny(taskConfig.CollapsedControlText) +
		tdiStrLenIfAny(taskConfig.Footer)
	for _, btn := range taskConfig.Buttons {
		totTextUtf16Words += tdiStrLenIfAny(btn.Text)
	}
	for _, btn := range taskConfig.RadioButtons {
		totTextUtf16Words += tdiStrLenIfAny(btn.Text)
	}

	// Sizes of each block, in bytes.
	szTdc := TASKDIALOGCONFIG_SZ
	szBtns := len(taskConfig.Buttons) * TASKDIALOG_BUTTON_SZ
	szRads := len(taskConfig.RadioButtons) * TASKDIALOG_BUTTON_SZ
	szTexts := totTextUtf16Words * 2

	totSize := szTdc + szBtns + szRads +
		3*4 + // button, radio and check values returned (int32)
		szTexts // all strings, null-terminated

	// Alloc a single buffer to keep everything.
	buf := NewVecSized(totSize, byte(0))
	defer buf.Free()

	// Subslices over the single buffer.
	tdcBuf := buf.HotSlice()[:szTdc]
	btnsBuf := buf.HotSlice()[szTdc : szTdc+szBtns]
	radsBuf := buf.HotSlice()[szTdc+szBtns : szTdc+szBtns+szRads]
	bufRetBtn := buf.HotSlice()[szTdc+szBtns+szRads : szTdc+szBtns+szRads+4]
	bufRetRad := buf.HotSlice()[szTdc+szBtns+szRads+4 : szTdc+szBtns+szRads+8]
	bufRetChk := buf.HotSlice()[szTdc+szBtns+szRads+8 : szTdc+szBtns+szRads+12]
	bufPtrStrs := buf.HotSlice()[szTdc+szBtns+szRads+12:]

	bufPtrStrs16 := unsafe.Slice( // wchar buffer for all strings
		(*uint16)(unsafe.Pointer(unsafe.SliceData(bufPtrStrs))), len(bufPtrStrs)/2)

	taskConfig.serialize(tdcBuf, btnsBuf, radsBuf, bufPtrStrs16)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_comctl_TaskDialogIndirect, "TaskDialogIndirect"),
		uintptr(unsafe.Pointer(unsafe.SliceData(tdcBuf))),
		uintptr(unsafe.Pointer(unsafe.SliceData(bufRetBtn))),
		uintptr(unsafe.Pointer(unsafe.SliceData(bufRetRad))),
		uintptr(unsafe.Pointer(unsafe.SliceData(bufRetChk))))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return co.ID(0), hr
	}

	pBtnId := (*int32)(unsafe.Pointer(unsafe.SliceData(bufRetBtn)))
	return co.ID(*pBtnId), nil
}

var _comctl_TaskDialogIndirect *syscall.Proc

func tdiStrLenIfAny(s string) int {
	if s != "" {
		return wstr.CountUtf16Len(s) + 1 // count terminating null
	}
	return 0
}
