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

// [InitCommonControls] function.
//
// [InitCommonControls]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrols
func InitCommonControls() {
	syscall.SyscallN(proc.InitCommonControls.Addr())
}

// [InitCommonControlsEx] function.
//
// [InitCommonControlsEx]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrolsex
func InitCommonControlsEx(icce *INITCOMMONCONTROLSEX) bool {
	ret, _, _ := syscall.SyscallN(proc.InitCommonControlsEx.Addr(),
		uintptr(unsafe.Pointer(icce)))
	return ret != 0
}

// [TaskDialogIndirect] function.
//
// Prefer using ui.TaskDlg wrappers, which deals with the most commons cases of
// this function in a safer, easier way.
//
// [TaskDialogIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialogindirect
func TaskDialogIndirect(taskConfig *TASKDIALOGCONFIG) co.ID {
	serialized, ptrs := taskConfig.serializePacked()

	hHeap := GetProcessHeap()
	memPnButton, _ := hHeap.HeapAlloc(co.HEAP_ALLOC_ZERO_MEMORY, uint(unsafe.Sizeof(int32(0))))
	defer hHeap.HeapFree(0, memPnButton)

	ret, _, _ := syscall.SyscallN(proc.TaskDialogIndirect.Addr(),
		uintptr(unsafe.Pointer(&serialized[0])),
		uintptr(unsafe.Pointer(&memPnButton[0])))

	if wErr := errco.ERROR(ret); wErr != errco.S_OK {
		panic(wErr)
	}

	runtime.KeepAlive(serialized)
	runtime.KeepAlive(ptrs)

	return co.ID(*(*int32)(unsafe.Pointer(&memPnButton[0])))
}
