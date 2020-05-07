package api

import (
	"gowinui/api/proc"
	c "gowinui/consts"
	"syscall"
	"unsafe"
)

type OSVERSIONINFOEX struct {
	DwOsVersionInfoSize uint32
	DwMajorVersion      uint32
	DwMinorVersion      uint32
	DwBuildNumber       uint32
	DWPlatformId        uint32
	SzCSDVersion        [128]uint16
	WServicePackMajor   uint16
	WServicePackMinor   uint16
	WSuiteMask          uint16
	WProductType        uint8
	WReserve            uint8
}

func IsWindowsVersionOrGreater(majorVersion, minorVersion uint32,
	servicePackMajor uint16) bool {

	ovi := OSVERSIONINFOEX{
		DwMajorVersion:    majorVersion,
		DwMinorVersion:    minorVersion,
		WServicePackMajor: servicePackMajor,
	}
	ovi.DwOsVersionInfoSize = uint32(unsafe.Sizeof(ovi))

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

	ret, _, lerr := syscall.Syscall(proc.VerifyVersionInfo.Addr(), 3,
		uintptr(unsafe.Pointer(ovi)),
		uintptr(typeMask), uintptr(conditionMask))
	return ret != 0, lerr
}

func VerSetConditionMask(conditionMask uint64, typeMask c.VER,
	condition c.VERCOND) uint64 {

	ret, _, _ := syscall.Syscall(proc.VerSetConditionMask.Addr(), 3,
		uintptr(conditionMask), uintptr(typeMask), uintptr(condition))
	return uint64(ret)
}
