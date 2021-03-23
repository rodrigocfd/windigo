package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/err"
)

// A handle to an accelerator table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#haccel
type HACCEL HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyacceleratortablew
func (hAccel HACCEL) CopyAcceleratorTable() []ACCEL {
	szRet, _, _ := syscall.Syscall(proc.CopyAcceleratorTable.Addr(), 3,
		uintptr(hAccel), 0, 0)
	if szRet == 0 {
		return []ACCEL{}
	}
	accelList := make([]ACCEL, szRet)
	syscall.Syscall(proc.CopyAcceleratorTable.Addr(), 3,
		uintptr(hAccel), uintptr(unsafe.Pointer(&accelList[0])), szRet)
	return accelList
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyacceleratortable
func (hAccel HACCEL) DestroyAcceleratorTable() {
	ret, _, lerr := syscall.Syscall(proc.DestroyAcceleratorTable.Addr(), 1,
		uintptr(hAccel), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}
