package api

import (
	"syscall"
	"unsafe"
	"winffi/procs"
)

type HACCEL HANDLE

func (hAccel HACCEL) CopyAcceleratorTable() []ACCEL {
	szRet, _, _ := syscall.Syscall(procs.CopyAcceleratorTable.Addr(), 3,
		uintptr(hAccel), 0, 0)
	if szRet == 0 {
		return []ACCEL{}
	}

	accelList := make([]ACCEL, uint32(szRet))
	syscall.Syscall(procs.CopyAcceleratorTable.Addr(), 3,
		uintptr(hAccel), uintptr(unsafe.Pointer(&accelList[0])), szRet)
	return accelList
}

func CreateAcceleratorTable(accelList []ACCEL) (HACCEL, syscall.Errno) {
	ret, _, errno := syscall.Syscall(procs.CreateAcceleratorTable.Addr(), 2,
		uintptr(unsafe.Pointer(&accelList[0])), uintptr(len(accelList)),
		0)
	return HACCEL(ret), errno
}

func (hAccel HACCEL) DestroyAcceleratorTable() bool {
	ret, _, _ := syscall.Syscall(procs.DestroyAcceleratorTable.Addr(), 1,
		uintptr(hAccel), 0, 0)
	return ret != 0
}
