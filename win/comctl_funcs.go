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
	syscall.Syscall(proc.InitCommonControls.Addr(), 0, 0, 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrolsex
func InitCommonControlsEx(icce *INITCOMMONCONTROLSEX) bool {
	ret, _, _ := syscall.Syscall(proc.InitCommonControlsEx.Addr(), 1,
		uintptr(unsafe.Pointer(icce)), 0, 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialogindirect
func TaskDialogIndirect(taskConfig *TASKDIALOGCONFIG) co.ID {
	serialized, ptrs := taskConfig.serializePacked()

	memPnButton := GlobalAlloc(co.GMEM_FIXED, 4) // sizeof(int)
	defer memPnButton.GlobalFree()

	ret, _, _ := syscall.Syscall6(proc.TaskDialogIndirect.Addr(), 4,
		uintptr(unsafe.Pointer(&serialized[0])),
		uintptr(unsafe.Pointer(memPnButton)),
		0, 0, 0, 0)

	if wErr := errco.ERROR(ret); wErr != errco.S_OK {
		panic(wErr)
	}

	runtime.KeepAlive(serialized)
	runtime.KeepAlive(ptrs)

	return co.ID(*(*int32)(unsafe.Pointer(memPnButton)))
}
