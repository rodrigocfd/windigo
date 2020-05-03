package api

import (
	"syscall"
	"unsafe"
	"winffi/consts"
	"winffi/procs"
)

type OSVERSIONINFOEX struct {
	OsVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformId        uint32
	CSDVersion        [128]uint16
	ServicePackMajor  uint16
	ServicePackMinor  uint16
	SuiteMask         uint16
	ProductType       uint8
	Reserve           uint8
}

func (ovi *OSVERSIONINFOEX) VerifyVersionInfo(typeMask consts.VER,
	conditionMask uint64) (bool, syscall.Errno) {

	ret, _, errno := syscall.Syscall(procs.VerifyVersionInfo.Addr(), 3,
		uintptr(unsafe.Pointer(ovi)),
		uintptr(typeMask), uintptr(conditionMask))
	return ret != 0, errno
}

func VerSetConditionMask(conditionMask uint64, typeMask consts.VER,
	condition consts.VERCOND) uint64 {

	ret, _, _ := syscall.Syscall(procs.VerSetConditionMask.Addr(), 3,
		uintptr(conditionMask), uintptr(typeMask), uintptr(condition))
	return uint64(ret)
}
