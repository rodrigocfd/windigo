/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"windigo/co"
	proc "windigo/win/internal"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#haccel
type HACCEL HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyacceleratortablew
func (hAccel HACCEL) CopyAcceleratorTable() []ACCEL {
	szRet, _, _ := syscall.Syscall(proc.CopyAcceleratorTable.Addr(), 3,
		uintptr(hAccel), 0, 0)
	if szRet == 0 {
		return []ACCEL{}
	}
	accelList := make([]ACCEL, uint32(szRet))
	syscall.Syscall(proc.CopyAcceleratorTable.Addr(), 3,
		uintptr(hAccel), uintptr(unsafe.Pointer(&accelList[0])), szRet)
	return accelList
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createacceleratortablew
func CreateAcceleratorTable(accelList []ACCEL) HACCEL {
	ret, _, lerr := syscall.Syscall(proc.CreateAcceleratorTable.Addr(), 2,
		uintptr(unsafe.Pointer(&accelList[0])), uintptr(len(accelList)),
		0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "CreateAcceleratorTable"))
	}
	return HACCEL(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyacceleratortable
func (hAccel HACCEL) DestroyAcceleratorTable() bool {
	ret, _, _ := syscall.Syscall(proc.DestroyAcceleratorTable.Addr(), 1,
		uintptr(hAccel), 0, 0)
	return ret != 0
}
