package api

import (
	"syscall"
	"unsafe"
	c "winffi/consts"
	p "winffi/procs"
)

func InitCommonControls() {
	syscall.Syscall(p.InitCommonControls.Addr(), 0,
		0, 0, 0)
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
			hiByte(uint16(c.WIN32_WINNT_WINTHRESHOLD))),
		uint32(
			loByte(uint16(c.WIN32_WINNT_WINTHRESHOLD))),
		0)
}

func IsWindows7OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(
			hiByte(uint16(c.WIN32_WINNT_WIN7))),
		uint32(
			loByte(uint16(c.WIN32_WINNT_WIN7))),
		0)
}

func IsWindowsVistaOrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(
			hiByte(uint16(c.WIN32_WINNT_VISTA))),
		uint32(
			loByte(uint16(c.WIN32_WINNT_VISTA))),
		0)
}
