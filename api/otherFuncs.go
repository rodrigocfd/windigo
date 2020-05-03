package api

import (
	"syscall"
	"unsafe"
	"winffi/consts"
	"winffi/procs"
)

func InitCommonControls() {
	syscall.Syscall(procs.InitCommonControls.Addr(), 0,
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
			VerSetConditionMask(0, consts.VER_MAJORVERSION,
				consts.VER_GREATER_EQUAL),
			consts.VER_MINORVERSION, consts.VER_GREATER_EQUAL),
		consts.VER_SERVICEPACKMAJOR, consts.VER_GREATER_EQUAL)

	ret, _ := ovi.VerifyVersionInfo(
		consts.VER_MAJORVERSION|consts.VER_MINORVERSION|
			consts.VER_SERVICEPACKMAJOR,
		conditionMask)
	return ret
}

func IsWindows10OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(
			hiByte(uint16(consts.WIN32_WINNT_WINTHRESHOLD))),
		uint32(
			loByte(uint16(consts.WIN32_WINNT_WINTHRESHOLD))),
		0)
}

func IsWindows7OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(
			hiByte(uint16(consts.WIN32_WINNT_WIN7))),
		uint32(
			loByte(uint16(consts.WIN32_WINNT_WIN7))),
		0)
}

func IsWindowsVistaOrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(
			hiByte(uint16(consts.WIN32_WINNT_VISTA))),
		uint32(
			loByte(uint16(consts.WIN32_WINNT_VISTA))),
		0)
}
