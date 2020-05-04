package api

import (
	"syscall"
	"unsafe"
	c "winffi/consts"
	p "winffi/procs"
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

func IsWindowsVersionOrGreater(majorVersion, minorVersion uint32,
	servicePackMajor uint16) bool {

	ovi := OSVERSIONINFOEX{
		MajorVersion:     majorVersion,
		MinorVersion:     minorVersion,
		ServicePackMajor: servicePackMajor,
	}
	ovi.OsVersionInfoSize = uint32(unsafe.Sizeof(ovi))

	conditionMask := VerSetConditionMask(
		VerSetConditionMask(
			VerSetConditionMask(0, c.VER_MAJORVERSION,
				c.VER_GREATER_EQUAL),
			c.VER_MINORVERSION, c.VER_GREATER_EQUAL),
		c.VER_SERVICEPACKMAJOR, c.VER_GREATER_EQUAL)

	ret, _ := ovi.VerifyVersionInfo(
		c.VER_MAJORVERSION|c.VER_MINORVERSION|
			c.VER_SERVICEPACKMAJOR,
		conditionMask)
	return ret
}

func IsWindows10OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(
			HiByte(uint16(c.WIN32_WINNT_WINTHRESHOLD))),
		uint32(
			LoByte(uint16(c.WIN32_WINNT_WINTHRESHOLD))),
		0)
}

func IsWindows7OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(
			HiByte(uint16(c.WIN32_WINNT_WIN7))),
		uint32(
			LoByte(uint16(c.WIN32_WINNT_WIN7))),
		0)
}

func IsWindowsVistaOrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(
			HiByte(uint16(c.WIN32_WINNT_VISTA))),
		uint32(
			LoByte(uint16(c.WIN32_WINNT_VISTA))),
		0)
}

func (ovi *OSVERSIONINFOEX) VerifyVersionInfo(typeMask c.VER,
	conditionMask uint64) (bool, syscall.Errno) {

	ret, _, errno := syscall.Syscall(p.VerifyVersionInfo.Addr(), 3,
		uintptr(unsafe.Pointer(ovi)),
		uintptr(typeMask), uintptr(conditionMask))
	return ret != 0, errno
}

func VerSetConditionMask(conditionMask uint64, typeMask c.VER,
	condition c.VERCOND) uint64 {

	ret, _, _ := syscall.Syscall(p.VerSetConditionMask.Addr(), 3,
		uintptr(conditionMask), uintptr(typeMask), uintptr(condition))
	return uint64(ret)
}
