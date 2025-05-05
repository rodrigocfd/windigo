//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to an [accelerator table].
//
// [accelerator table]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#haccel
type HACCEL HANDLE

// [CreateAcceleratorTable] function.
//
// ⚠️ You must defer HACCEL.DestroyAcceleratorTable().
//
// [CreateAcceleratorTable]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createacceleratortablew
func CreateAcceleratorTable(accelList []ACCEL) (HACCEL, error) {
	ret, _, err := syscall.SyscallN(_CreateAcceleratorTableW.Addr(),
		uintptr(unsafe.Pointer(&accelList[0])), uintptr(len(accelList)))
	if ret == 0 {
		return HACCEL(0), co.ERROR(err)
	}
	return HACCEL(ret), nil
}

var _CreateAcceleratorTableW = dll.User32.NewProc("CreateAcceleratorTableW")

// [CopyAcceleratorTable] function.
//
// [CopyAcceleratorTable]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyacceleratortablew
func (hAccel HACCEL) CopyAcceleratorTable() []ACCEL {
	szRet, _, _ := syscall.SyscallN(_CopyAcceleratorTableW.Addr(),
		uintptr(hAccel), 0, 0)
	if szRet == 0 {
		return []ACCEL{}
	}
	accelList := make([]ACCEL, szRet)
	syscall.SyscallN(_CopyAcceleratorTableW.Addr(),
		uintptr(hAccel), uintptr(unsafe.Pointer(&accelList[0])), szRet)
	return accelList
}

var _CopyAcceleratorTableW = dll.User32.NewProc("CopyAcceleratorTableW")

// [DestroyAcceleratorTable] function.
//
// Paired with [CreateAcceleratorTable].
//
// [DestroyAcceleratorTable]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyacceleratortable
// [CreateAcceleratorTable]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createacceleratortablew
func (hAccel HACCEL) DestroyAcceleratorTable() error {
	ret, _, err := syscall.SyscallN(_DestroyAcceleratorTable.Addr(),
		uintptr(hAccel))
	return util.ZeroToGetLastError(ret, err)
}

var _DestroyAcceleratorTable = dll.User32.NewProc("DestroyAcceleratorTable")
