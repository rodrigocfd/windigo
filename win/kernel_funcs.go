//go:build windows

package win

import (
	"strings"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [CopyFile] function.
//
// [CopyFile]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-copyfilew
func CopyFile(existingFile, newFile string, failIfExists bool) error {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pExistingFile := wbuf.PtrEmptyIsNil(existingFile)
	pNewFile := wbuf.PtrEmptyIsNil(newFile)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_CopyFileW, "CopyFileW"),
		uintptr(pExistingFile),
		uintptr(pNewFile),
		utl.BoolToUintptr(failIfExists))
	return utl.ZeroAsGetLastError(ret, err)
}

var _CopyFileW *syscall.Proc

// [CreateDirectory] function.
//
// [CreateDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createdirectoryw
func CreateDirectory(pathName string, securityAttributes *SECURITY_ATTRIBUTES) error {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pPathName := wbuf.PtrEmptyIsNil(pathName)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_CreateDirectoryW, "CreateDirectoryW"),
		uintptr(pPathName),
		uintptr(unsafe.Pointer(securityAttributes)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _CreateDirectoryW *syscall.Proc

// [CreateProcess] function.
//
// ⚠️ You must defer [HPROCESS.CloseHandle] and [HTHREAD.CloseHandle] on
// HProcess and HThread members returned in [PROCESS_INFORMATION].
//
// # Example
//
//	var si win.STARTUPINFO
//	si.SetCb()
//
//	pi, _ := win.CreateProcess("C:\\Windows\\notepad.exe", "",
//		nil, nil, false, co.CREATE_NONE, []string{"FOO=bar", "BAR=44"},
//		"", &si)
//
//	defer pi.HProcess.CloseHandle()
//	defer pi.HThread.CloseHandle()
//
// [CreateProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-createprocessw
func CreateProcess(
	applicationName, commandLine string,
	processAttributes, threadAttributes *SECURITY_ATTRIBUTES,
	inheritHandles bool,
	creationFlags co.CREATE,
	environment []string,
	currentDirectory string,
	startupInfo *STARTUPINFO,
) (PROCESS_INFORMATION, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pApplicationName := wbuf.PtrEmptyIsNil(applicationName)
	pCommandLine := wbuf.PtrEmptyIsNil(commandLine)
	pCurrentDirectory := wbuf.PtrEmptyIsNil(currentDirectory)

	pEnvironment := wstr.GoArrToWinPtr(environment...)
	var pi PROCESS_INFORMATION

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_CreateProcessW, "CreateProcessW"),
		uintptr(pApplicationName),
		uintptr(pCommandLine),
		uintptr(unsafe.Pointer(processAttributes)),
		uintptr(unsafe.Pointer(threadAttributes)),
		utl.BoolToUintptr(inheritHandles),
		uintptr(creationFlags|co.CREATE_UNICODE_ENVIRONMENT), // env strings are always UTF-16
		uintptr(unsafe.Pointer(pEnvironment)),
		uintptr(pCurrentDirectory),
		uintptr(unsafe.Pointer(startupInfo)),
		uintptr(unsafe.Pointer(&pi)))
	if ret == 0 {
		return PROCESS_INFORMATION{}, co.ERROR(err)
	}
	return pi, nil
}

var _CreateProcessW *syscall.Proc

// [DeleteFile] function.
//
// [DeleteFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-deletefilew
func DeleteFile(fileName string) error {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pFileName := wbuf.PtrEmptyIsNil(fileName)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_DeleteFileW, "DeleteFileW"),
		uintptr(pFileName))
	return utl.ZeroAsGetLastError(ret, err)
}

var _DeleteFileW *syscall.Proc

// [ExitProcess] function.
//
// [ExitProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-exitprocess
func ExitProcess(exitCode uint32) {
	syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_ExitProcess, "ExitProcess"),
		uintptr(exitCode))
}

var _ExitProcess *syscall.Proc

// [ExpandEnvironmentStrings] function.
//
// [ExpandEnvironmentStrings]: https://learn.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-expandenvironmentstringsw
func ExpandEnvironmentStrings(s string) (string, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pS := wbuf.PtrAllowEmpty(s)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_ExpandEnvironmentStringsW, "ExpandEnvironmentStringsW"),
		uintptr(pS),
		0, 0) // 1st call to retrieve the required length
	if ret == 0 {
		return "", co.ERROR(err)
	}

	szBuf := uint(ret) // includes terminating null
	recvBuf := wstr.NewBufReceiver(szBuf)
	defer recvBuf.Free()

	for {
		recvBuf.Resize(szBuf)

		ret, _, err = syscall.SyscallN(
			dll.Load(dll.KERNEL32, &_ExpandEnvironmentStringsW, "ExpandEnvironmentStringsW"),
			uintptr(pS),
			uintptr(recvBuf.UnsafePtr()),
			uintptr(uint32(szBuf)))
		if ret == 0 {
			return "", co.ERROR(err)
		}
		required := uint(ret) // plus terminating null count

		if required <= szBuf {
			return recvBuf.String(), nil
		}

		szBuf = required // set new buffer size to try again
	}
}

var _ExpandEnvironmentStringsW *syscall.Proc

// [FileTimeToSystemTime] function.
//
// [FileTimeToSystemTime]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-filetimetosystemtime
func FileTimeToSystemTime(ft *FILETIME) (SYSTEMTIME, error) {
	var st SYSTEMTIME
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_FileTimeToSystemTime, "FileTimeToSystemTime"),
		uintptr(unsafe.Pointer(ft)),
		uintptr(unsafe.Pointer(&st)))
	if ret == 0 {
		return SYSTEMTIME{}, co.ERROR(err)
	}
	return st, nil
}

var _FileTimeToSystemTime *syscall.Proc

// [GetCommandLine] function.
//
// [GetCommandLine]: https://learn.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-getcommandlinew
func GetCommandLine() string {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetCommandLineW, "GetCommandLineW"))
	return wstr.WinPtrToGo((*uint16)(unsafe.Pointer(ret)))
}

var _GetCommandLineW *syscall.Proc

// [GetCurrentProcessId] function.
//
// [GetCurrentProcessId]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocessid
func GetCurrentProcessId() uint32 {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetCurrentProcessId, "GetCurrentProcessId"))
	return uint32(ret)
}

var _GetCurrentProcessId *syscall.Proc

// [GetCurrentThreadId] function.
//
// [GetCurrentThreadId]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthreadid
func GetCurrentThreadId() uint32 {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetCurrentThreadId, "GetCurrentThreadId"))
	return uint32(ret)
}

var _GetCurrentThreadId *syscall.Proc

// [GetEnvironmentStrings] function.
//
// [FreeEnvironmentStrings] is automatically called after the data retrieval.
//
// [GetEnvironmentStrings]: https://learn.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-getenvironmentstringsw
// [FreeEnvironmentStrings]: https://learn.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-freeenvironmentstringsw
func GetEnvironmentStrings() (map[string]string, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetEnvironmentStringsW, "GetEnvironmentStringsW"))
	if ret == 0 {
		return nil, co.ERROR_NOT_CAPABLE
	}
	rawEntries := wstr.WinArrPtrToGo((*uint16)(unsafe.Pointer(ret)))

	ret, _, _ = syscall.SyscallN( // free right away
		dll.Load(dll.KERNEL32, &_FreeEnvironmentStringsW, "FreeEnvironmentStringsW"),
		ret)
	if ret == 0 {
		return nil, co.ERROR_NOT_CAPABLE
	}

	mapEntries := make(map[string]string, len(rawEntries))
	for _, entry := range rawEntries {
		keyVal := strings.SplitN(entry, "=", 2)
		mapEntries[keyVal[0]] = keyVal[1]
	}
	return mapEntries, nil
}

var (
	_GetEnvironmentStringsW  *syscall.Proc
	_FreeEnvironmentStringsW *syscall.Proc
)

// [GetFileAttributes] function.
//
// [GetFileAttributes]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfileattributesw
func GetFileAttributes(fileName string) (co.FILE_ATTRIBUTE, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pFileName := wbuf.PtrEmptyIsNil(fileName)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetFileAttributesW, "GetFileAttributesW"),
		uintptr(pFileName))

	if retAttr := co.FILE_ATTRIBUTE(ret); retAttr == co.FILE_ATTRIBUTE_INVALID {
		return retAttr, co.ERROR(err) // err is extended error information
	} else {
		return retAttr, nil
	}
}

var _GetFileAttributesW *syscall.Proc

// [GetLocalTime] function.
//
// [GetLocalTime]: https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getlocaltime
func GetLocalTime() SYSTEMTIME {
	var st SYSTEMTIME
	syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetLocalTime, "GetLocalTime"),
		uintptr(unsafe.Pointer(&st)))
	return st
}

var _GetLocalTime *syscall.Proc

// [GetStartupInfo] function.
//
// [GetStartupInfo]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getstartupinfow
func GetStartupInfo() STARTUPINFO {
	var si STARTUPINFO
	si.SetCb()

	syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetStartupInfoW, "GetStartupInfoW"),
		uintptr(unsafe.Pointer(&si)))
	return si
}

var _GetStartupInfoW *syscall.Proc

// [GetTimeZoneInformation] function.
//
// [GetTimeZoneInformation]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-gettimezoneinformation
func GetTimeZoneInformation() (TIME_ZONE_INFORMATION, co.TIME_ZONE_ID, error) {
	var tzi TIME_ZONE_INFORMATION
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetTimeZoneInformation, "GetTimeZoneInformation"),
		uintptr(unsafe.Pointer(&tzi)))
	if ret == utl.TIME_ZONE_INVALID {
		return TIME_ZONE_INFORMATION{}, co.TIME_ZONE_ID(0), co.ERROR(err)
	}
	return tzi, co.TIME_ZONE_ID(ret), nil
}

var _GetTimeZoneInformation *syscall.Proc

// [GetSystemInfo] function.
//
// [GetSystemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getsysteminfo
func GetSystemInfo() SYSTEM_INFO {
	var si SYSTEM_INFO
	syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetSystemInfo, "GetSystemInfo"),
		uintptr(unsafe.Pointer(&si)))
	return si
}

var _GetSystemInfo *syscall.Proc

// [IsWindows10OrGreater] function.
//
// Panics on error.
//
// [IsWindows10OrGreater]: https://learn.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows10orgreater
func IsWindows10OrGreater() bool {
	ret, err := IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_WINTHRESHOLD))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_WINTHRESHOLD))),
		0)
	if err != nil {
		panic(err)
	}
	return ret
}

// [IsWindows7OrGreater] function.
//
// Panics on error.
//
// [IsWindows7OrGreater]: https://learn.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows7orgreater
func IsWindows7OrGreater() bool {
	ret, err := IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_WIN7))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_WIN7))),
		0)
	if err != nil {
		panic(err)
	}
	return ret
}

// [IsWindows8OrGreater] function.
//
// Panics on error.
//
// [IsWindows8OrGreater]: https://learn.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows8orgreater
func IsWindows8OrGreater() bool {
	ret, err := IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_WIN8))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_WIN8))),
		0)
	if err != nil {
		panic(err)
	}
	return ret
}

// [IsWindows8Point1OrGreater] function.
//
// Panics on error.
//
// [IsWindows8Point1OrGreater]: https://learn.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows8point1orgreater
func IsWindows8Point1OrGreater() bool {
	ret, err := IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_WINBLUE))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_WINBLUE))),
		0)
	if err != nil {
		panic(err)
	}
	return ret
}

// [IsWindowsVistaOrGreater] function.
//
// Panics on error.
//
// [IsWindowsVistaOrGreater]: https://learn.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsvistaorgreater
func IsWindowsVistaOrGreater() bool {
	ret, err := IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_VISTA))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_VISTA))),
		0)
	if err != nil {
		panic(err)
	}
	return ret
}

// [IsWindowsXpOrGreater] function.
//
// Panics on error.
//
// [IsWindowsXpOrGreater]: https://learn.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsxporgreater
func IsWindowsXpOrGreater() bool {
	ret, err := IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_WINXP))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_WINXP))),
		0)
	if err != nil {
		panic(err)
	}
	return ret
}

// [IsWindowsVersionOrGreater] function.
//
// [IsWindowsVersionOrGreater]: https://learn.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsversionorgreater
func IsWindowsVersionOrGreater(majorVersion, minorVersion uint32, servicePackMajor uint16) (bool, error) {
	ovi := OSVERSIONINFOEX{
		DwMajorVersion:    majorVersion,
		DwMinorVersion:    minorVersion,
		WServicePackMajor: servicePackMajor,
	}
	ovi.SetDwOsVersionInfoSize()

	conditionMask := VerSetConditionMask(
		VerSetConditionMask(
			VerSetConditionMask(0, co.VER_MAJORVERSION, co.VER_COND_GREATER_EQUAL),
			co.VER_MINORVERSION, co.VER_COND_GREATER_EQUAL),
		co.VER_SERVICEPACKMAJOR, co.VER_COND_GREATER_EQUAL)

	ret, err := VerifyVersionInfo(&ovi,
		co.VER_MAJORVERSION|co.VER_MINORVERSION|co.VER_SERVICEPACKMAJOR,
		conditionMask)
	if err != nil {
		return false, err
	}
	return ret, nil
}

// [HIBYTE] macro.
//
// [HIBYTE]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632656(v=vs.85)
func HIBYTE(val uint16) uint8 {
	_, hi := utl.Break16(val)
	return hi
}

// [HIWORD] macro.
//
// [HIWORD]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632657(v=vs.85)
func HIWORD(val uint32) uint16 {
	_, hi := utl.Break32(val)
	return hi
}

// [LOBYTE] macro.
//
// [LOBYTE]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632658(v=vs.85)
func LOBYTE(val uint16) uint8 {
	lo, _ := utl.Break16(val)
	return lo
}

// [LOWORD] macro.
//
// [LOWORD]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632659(v=vs.85)
func LOWORD(val uint32) uint16 {
	lo, _ := utl.Break32(val)
	return lo
}

// [MAKELONG] macro.
//
// [MAKELONG]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632660(v=vs.85)
func MAKELONG(lo, hi uint16) uint32 {
	return utl.Make32(lo, hi)
}

// [MAKEWORD] macro.
//
// [MAKEWORD]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632663(v=vs.85)
func MAKEWORD(lo, hi uint8) uint16 {
	return utl.Make16(lo, hi)
}

// [SetConsoleTitle] function.
//
// [SetConsoleTitle]: https://learn.microsoft.com/en-us/windows/console/setconsoletitle
func SetConsoleTitle(title string) error {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pTitle := wbuf.PtrEmptyIsNil(title)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_SetConsoleTitleW, "SetConsoleTitleW"),
		uintptr(pTitle))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetConsoleTitleW *syscall.Proc

// [SetCurrentDirectory] function.
//
// [SetCurrentDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setcurrentdirectory
func SetCurrentDirectory(pathName string) error {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pPathName := wbuf.PtrEmptyIsNil(pathName)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_SetCurrentDirectoryW, "SetCurrentDirectoryW"),
		uintptr(pPathName))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetCurrentDirectoryW *syscall.Proc

// [SystemTimeToFileTime] function.
//
// [SystemTimeToFileTime]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetofiletime
func SystemTimeToFileTime(st *SYSTEMTIME) (FILETIME, error) {
	var ft FILETIME
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_SystemTimeToFileTime, "SystemTimeToFileTime"),
		uintptr(unsafe.Pointer(st)),
		uintptr(unsafe.Pointer(&ft)))
	if ret == 0 {
		return FILETIME{}, co.ERROR(err)
	}
	return ft, nil
}

var _SystemTimeToFileTime *syscall.Proc

// [SystemTimeToTzSpecificLocalTime] function.
//
// [SystemTimeToTzSpecificLocalTime]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetotzspecificlocaltime
func SystemTimeToTzSpecificLocalTime(
	timeZoneInfo *TIME_ZONE_INFORMATION,
	inUniversalTime *SYSTEMTIME,
) (SYSTEMTIME, error) {
	var st SYSTEMTIME
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_SystemTimeToTzSpecificLocalTime, "SystemTimeToTzSpecificLocalTime"),
		uintptr(unsafe.Pointer(timeZoneInfo)),
		uintptr(unsafe.Pointer(inUniversalTime)),
		uintptr(unsafe.Pointer(&st)))
	if ret == 0 {
		return SYSTEMTIME{}, co.ERROR(err)
	}
	return st, nil
}

var _SystemTimeToTzSpecificLocalTime *syscall.Proc
