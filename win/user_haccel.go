//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to an [accelerator table].
//
// [accelerator table]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#haccel
type HACCEL HANDLE

// [CreateAcceleratorTable] function.
//
// ⚠️ You must defer [HACCEL.DestroyAcceleratorTable].
//
// [CreateAcceleratorTable]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createacceleratortablew
func CreateAcceleratorTable(accelList []ACCEL) (HACCEL, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_CreateAcceleratorTableW, "CreateAcceleratorTableW"),
		uintptr(unsafe.Pointer(&accelList[0])),
		uintptr(int32(len(accelList))))
	if ret == 0 {
		return HACCEL(0), co.ERROR(err)
	}
	return HACCEL(ret), nil
}

var _CreateAcceleratorTableW *syscall.Proc

// [CopyAcceleratorTable] function.
//
// [CopyAcceleratorTable]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyacceleratortablew
func (hAccel HACCEL) CopyAcceleratorTable() []ACCEL {
	szRet, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_CopyAcceleratorTableW, "CopyAcceleratorTableW"),
		uintptr(hAccel),
		0, 0)
	if szRet == 0 {
		return []ACCEL{}
	}
	accelList := make([]ACCEL, szRet)
	syscall.SyscallN(
		dll.Load(dll.USER32, &_CopyAcceleratorTableW, "CopyAcceleratorTableW"),
		uintptr(hAccel),
		uintptr(unsafe.Pointer(&accelList[0])),
		szRet)
	return accelList
}

var _CopyAcceleratorTableW *syscall.Proc

// [DestroyAcceleratorTable] function.
//
// Paired with [CreateAcceleratorTable].
//
// [DestroyAcceleratorTable]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyacceleratortable
func (hAccel HACCEL) DestroyAcceleratorTable() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_DestroyAcceleratorTable, "DestroyAcceleratorTable"),
		uintptr(hAccel))
	return utl.ZeroAsGetLastError(ret, err)
}

var _DestroyAcceleratorTable *syscall.Proc
