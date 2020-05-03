package winffi

import (
	"syscall"
	"unsafe"
	"winffi/consts"
	"winffi/procs"
)

func VerifyVersionInfo(versionInfo *OSVERSIONINFOEX,
	typeMask consts.VER, conditionMask uint64) (bool, syscall.Errno) {

	ret, _, errno := syscall.Syscall(procs.VerifyVersionInfo.Addr(), 3,
		uintptr(unsafe.Pointer(versionInfo)),
		uintptr(typeMask), uintptr(conditionMask))
	return ret != 0, errno
}

func VerSetConditionMask(conditionMask uint64, typeMask consts.VER,
	condition consts.VERCOND) uint64 {

	ret, _, _ := syscall.Syscall(procs.VerSetConditionMask.Addr(), 3,
		uintptr(conditionMask), uintptr(typeMask), uintptr(condition))
	return uint64(ret)
}
