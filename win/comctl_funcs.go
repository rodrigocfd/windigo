//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [ImageList_DragMove] function.
//
// [ImageList_DragMove]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_dragmove
func ImageListDragMove(x, y int) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_ImageList_DragMove, "ImageList_DragMove"),
		uintptr(int32(x)),
		uintptr(int32(y)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_DragMove *syscall.Proc

// [ImageList_DragShowNolock] function.
//
// [ImageList_DragShowNolock]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_dragshownolock
func ImageListDragShowNolock(show bool) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_ImageList_DragShowNolock, "ImageList_DragShowNolock"),
		utl.BoolToUintptr(show))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_DragShowNolock *syscall.Proc

// [ImageList_DrawIndirect] function.
//
// [ImageList_DrawIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_drawindirect
func ImageListDrawIndirect(imldp *IMAGELISTDRAWPARAMS) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_ImageList_DrawIndirect, "ImageList_DrawIndirect"),
		uintptr(unsafe.Pointer(imldp)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_DrawIndirect *syscall.Proc

// [ImageList_EndDrag] function.
//
// [ImageList_EndDrag]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_enddrag
func ImageListEndDrag() {
	syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_ImageList_EndDrag, "ImageList_EndDrag"))
}

var _ImageList_EndDrag *syscall.Proc

// [InitCommonControls] function.
//
// [InitCommonControls]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrols
func InitCommonControls() {
	syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_InitCommonControls, "InitCommonControls"))
}

var _InitCommonControls *syscall.Proc

// [InitCommonControlsEx] function.
//
// [InitCommonControlsEx]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrolsex
func InitCommonControlsEx(icc co.ICC) error {
	var iccx _INITCOMMONCONTROLSEX
	iccx.SetDwSize()
	iccx.DwICC = icc

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_InitCommonControlsEx, "InitCommonControlsEx"),
		uintptr(unsafe.Pointer(&iccx)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _InitCommonControlsEx *syscall.Proc

// [InitMUILanguage] function.
//
// [InitMUILanguage]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initmuilanguage
func InitMUILanguage(lang LANGID) {
	syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_InitMUILanguage, "InitMUILanguage"),
		uintptr(lang))
}

var _InitMUILanguage *syscall.Proc

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

	totTextUtf16Words := tdiStrLenIfAny(taskConfig.WindowTitle) + // counts terminating nulls
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

	szTdc := uint(TASKDIALOGCONFIG_SZ) // sizes of all blocks in bytes
	szBtns := uint(len(taskConfig.Buttons)) * TASKDIALOG_BUTTON_SZ
	szRads := uint(len(taskConfig.RadioButtons)) * TASKDIALOG_BUTTON_SZ
	szTexts := totTextUtf16Words * 2

	totSize := szTdc + szBtns + szRads +
		3*4 + // button, radio and check values returned (int32)
		szTexts // all strings, null-terminated

	buf := NewVecSized(totSize, byte(0)) // alloc a single buffer to keep everything
	defer buf.Free()

	tdcBuf := buf.HotSlice()[:szTdc] // subslices over the single buffer
	btnsBuf := buf.HotSlice()[szTdc : szTdc+szBtns]
	radsBuf := buf.HotSlice()[szTdc+szBtns : szTdc+szBtns+szRads]
	bufRetBtn := buf.HotSlice()[szTdc+szBtns+szRads : szTdc+szBtns+szRads+4]
	bufRetRad := buf.HotSlice()[szTdc+szBtns+szRads+4 : szTdc+szBtns+szRads+8]
	bufRetChk := buf.HotSlice()[szTdc+szBtns+szRads+8 : szTdc+szBtns+szRads+12]
	bufPtrStrs := buf.HotSlice()[szTdc+szBtns+szRads+12:]

	bufPtrStrs16 := unsafe.Slice((*uint16)(unsafe.Pointer(&bufPtrStrs[0])), len(bufPtrStrs)/2)

	taskConfig.serialize(tdcBuf, btnsBuf, radsBuf, bufPtrStrs16)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.COMCTL32, &_TaskDialogIndirect, "TaskDialogIndirect"),
		uintptr(unsafe.Pointer(&tdcBuf[0])),
		uintptr(unsafe.Pointer(&bufRetBtn[0])),
		uintptr(unsafe.Pointer(&bufRetRad[0])),
		uintptr(unsafe.Pointer(&bufRetChk[0])))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return co.ID(0), hr
	}

	pBtnId := (*int32)(unsafe.Pointer(&bufRetBtn[0]))
	return co.ID(*pBtnId), nil
}

var _TaskDialogIndirect *syscall.Proc

func tdiStrLenIfAny(s string) uint {
	if s != "" {
		return wstr.CountUtf16Len(s) + 1 // count terminating null
	}
	return 0
}
