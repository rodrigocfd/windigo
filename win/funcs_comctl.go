//go:build windows

package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrols
func InitCommonControls() {
	syscall.SyscallN(proc.InitCommonControls.Addr())
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrolsex
func InitCommonControlsEx(icce *INITCOMMONCONTROLSEX) bool {
	ret, _, _ := syscall.SyscallN(proc.InitCommonControlsEx.Addr(),
		uintptr(unsafe.Pointer(icce)))
	return ret != 0
}

// Prefer using ui.TaskDlg wrappers, which deals with the most commons cases of
// this function in a safer, easier way.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialogindirect
func TaskDialogIndirect(taskConfig *TASKDIALOGCONFIG) co.ID {
	serialized, ptrs := taskConfig.serializePacked()

	memPnButton := GlobalAlloc(co.GMEM_FIXED, 4) // sizeof(int)
	defer memPnButton.GlobalFree()

	ret, _, _ := syscall.SyscallN(proc.TaskDialogIndirect.Addr(),
		uintptr(unsafe.Pointer(&serialized[0])),
		uintptr(unsafe.Pointer(memPnButton)))

	if wErr := errco.ERROR(ret); wErr != errco.S_OK {
		panic(wErr)
	}

	runtime.KeepAlive(serialized)
	runtime.KeepAlive(ptrs)

	return co.ID(*(*int32)(unsafe.Pointer(memPnButton)))
}
