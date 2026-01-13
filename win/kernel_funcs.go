//go:build windows

package win

import (
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [CopyFile] function.
//
// [CopyFile]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-copyfilew
func CopyFile(existingFile, newFile string, failIfExists bool) error {
	var wExistingFile, wNewFile wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_CopyFileW, "CopyFileW"),
		uintptr(wExistingFile.EmptyIsNil(existingFile)),
		uintptr(wNewFile.EmptyIsNil(newFile)),
		utl.BoolToUintptr(failIfExists))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_CopyFileW *syscall.Proc

// [CreateDirectory] function.
//
// [CreateDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createdirectoryw
func CreateDirectory(pathName string, securityAttributes *SECURITY_ATTRIBUTES) error {
	var wPathName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_CreateDirectoryW, "CreateDirectoryW"),
		uintptr(wPathName.EmptyIsNil(pathName)),
		uintptr(unsafe.Pointer(securityAttributes)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_CreateDirectoryW *syscall.Proc

// [CreateProcess] function.
//
// ⚠️ You must defer [HPROCESS.CloseHandle] and [HTHREAD.CloseHandle] on
// HProcess and HThread members returned in [PROCESS_INFORMATION].
//
// Example:
//
//	var si win.STARTUPINFO
//	si.SetCb()
//
//	pi, _ := win.CreateProcess(
//		"C:\\Windows\\notepad.exe",
//		"",
//		nil,
//		nil,
//		false,
//		co.CREATE_NONE,
//		map[string]string{
//			"FOO": "bar",
//			"BAZ": "44",
//		},
//		"",
//		&si,
//	)
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
	environment map[string]string,
	currentDirectory string,
	startupInfo *STARTUPINFO,
) (PROCESS_INFORMATION, error) {
	var wApplicationName, wCommandLine, wCurrentDirectory wstr.BufEncoder
	var pi PROCESS_INFORMATION

	var pEnvironment *uint16
	if len(environment) > 0 {
		envStrs := make([]string, 0, len(environment))
		for k, v := range environment {
			envStrs = append(envStrs, k+"="+v) // concat key=value for each env var
		}
		pEnvironment = wstr.EncodeArrToPtr(envStrs...)
	}

	if startupInfo != nil {
		startupInfo.SetCb() // safety
	}

	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_CreateProcessW, "CreateProcessW"),
		uintptr(wApplicationName.EmptyIsNil(applicationName)),
		uintptr(wCommandLine.EmptyIsNil(commandLine)),
		uintptr(unsafe.Pointer(processAttributes)),
		uintptr(unsafe.Pointer(threadAttributes)),
		utl.BoolToUintptr(inheritHandles),
		uintptr(creationFlags|co.CREATE_UNICODE_ENVIRONMENT), // env strings are always UTF-16
		uintptr(unsafe.Pointer(pEnvironment)),
		uintptr(wCurrentDirectory.EmptyIsNil(currentDirectory)),
		uintptr(unsafe.Pointer(startupInfo)),
		uintptr(unsafe.Pointer(&pi)))
	if ret == 0 {
		return PROCESS_INFORMATION{}, co.ERROR(err)
	}
	return pi, nil
}

var _kernel_CreateProcessW *syscall.Proc

// [DeactivateActCtx] function.
//
// Cookie is returned by [HACTCTX.ActivateActCtx].
//
// [DeactivateActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-deactivateactctx
func DeactivateActCtx(flags co.DEACTIVATE_ACTCTX, cookie int) error {
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_DeactivateActCtx, "DeactivateActCtx"),
		uintptr(flags),
		uintptr(cookie))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_DeactivateActCtx *syscall.Proc

// [DeleteFile] function.
//
// [DeleteFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-deletefilew
func DeleteFile(fileName string) error {
	var wFileName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_DeleteFileW, "DeleteFileW"),
		uintptr(wFileName.EmptyIsNil(fileName)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_DeleteFileW *syscall.Proc

// [ExitProcess] function.
//
// [ExitProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-exitprocess
func ExitProcess(exitCode uint32) {
	syscall.SyscallN(
		dll.Kernel.Load(&_kernel_ExitProcess, "ExitProcess"),
		uintptr(exitCode))
}

var _kernel_ExitProcess *syscall.Proc

// [ExpandEnvironmentStrings] function.
//
// [ExpandEnvironmentStrings]: https://learn.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-expandenvironmentstringsw
func ExpandEnvironmentStrings(s string) (string, error) {
	var wS wstr.BufEncoder
	pS := wS.AllowEmpty(s)

	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_ExpandEnvironmentStringsW, "ExpandEnvironmentStringsW"),
		uintptr(pS),
		0, 0) // 1st call to retrieve the required length
	if ret == 0 {
		return "", co.ERROR(err)
	}

	szBuf := int(ret) // includes terminating null
	var wBuf wstr.BufDecoder

	for {
		wBuf.AllocAndZero(szBuf)

		ret, _, err = syscall.SyscallN(
			dll.Kernel.Load(&_kernel_ExpandEnvironmentStringsW, "ExpandEnvironmentStringsW"),
			uintptr(pS),
			uintptr(wBuf.Ptr()),
			uintptr(uint32(szBuf)))
		if ret == 0 {
			return "", co.ERROR(err)
		}
		required := int(ret) // plus terminating null count

		if required <= szBuf {
			return wBuf.String(), nil
		}

		szBuf = required // set new buffer size to try again
	}
}

var _kernel_ExpandEnvironmentStringsW *syscall.Proc

// [FileTimeToSystemTime] function.
//
// [FileTimeToSystemTime]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-filetimetosystemtime
func FileTimeToSystemTime(ft *FILETIME) (SYSTEMTIME, error) {
	var st SYSTEMTIME
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_FileTimeToSystemTime, "FileTimeToSystemTime"),
		uintptr(unsafe.Pointer(ft)),
		uintptr(unsafe.Pointer(&st)))
	if ret == 0 {
		return SYSTEMTIME{}, co.ERROR(err)
	}
	return st, nil
}

var _kernel_FileTimeToSystemTime *syscall.Proc

// [FlushProcessWriteBuffers] function.
//
// [FlushProcessWriteBuffers]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-flushprocesswritebuffers
func FlushProcessWriteBuffers() {
	syscall.SyscallN(
		dll.Kernel.Load(&_kernel_FlushProcessWriteBuffers, "FlushProcessWriteBuffers"))
}

var _kernel_FlushProcessWriteBuffers *syscall.Proc

// [GetActiveProcessorCount] function.
//
// For ALL_PROCESSOR_GROUPS, pass 0xffff.
//
// Panics if groupNumber is negative.
//
// [GetActiveProcessorCount]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-getactiveprocessorcount
func GetActiveProcessorCount(groupNumber int) (int, error) {
	utl.PanicNeg(groupNumber)
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetActiveProcessorCount, "GetActiveProcessorCount"),
		uintptr(uint16(groupNumber)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return int(uint32(ret)), nil
}

var _kernel_GetActiveProcessorCount *syscall.Proc

// [GetActiveProcessorGroupCount] function.
//
// [GetActiveProcessorGroupCount]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-getactiveprocessorgroupcount
func GetActiveProcessorGroupCount() (int, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetActiveProcessorGroupCount, "GetActiveProcessorGroupCount"))
	if ret == 0 {
		return 0, co.ERROR_INVALID_PARAMETER
	}
	return int(uint16(ret)), nil
}

var _kernel_GetActiveProcessorGroupCount *syscall.Proc

// [GetCommandLine] function.
//
// [GetCommandLine]: https://learn.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-getcommandlinew
func GetCommandLine() string {
	ret, _, _ := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetCommandLineW, "GetCommandLineW"))
	return wstr.DecodePtr((*uint16)(unsafe.Pointer(ret)))
}

var _kernel_GetCommandLineW *syscall.Proc

// [GetCurrentProcessId] function.
//
// [GetCurrentProcessId]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocessid
func GetCurrentProcessId() uint32 {
	ret, _, _ := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetCurrentProcessId, "GetCurrentProcessId"))
	return uint32(ret)
}

var _kernel_GetCurrentProcessId *syscall.Proc

// [GetCurrentProcessorNumber] function.
//
// [GetCurrentProcessorNumber]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocessornumber
func GetCurrentProcessorNumber() int {
	ret, _, _ := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetCurrentProcessorNumber, "GetCurrentProcessorNumber"))
	return int(uint16(ret))
}

var _kernel_GetCurrentProcessorNumber *syscall.Proc

// [GetCurrentProcessorNumberEx] function.
//
// [GetCurrentProcessorNumberEx]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocessornumberex
func GetCurrentProcessorNumberEx() PROCESSOR_NUMBER {
	var pn PROCESSOR_NUMBER
	syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetCurrentProcessorNumberEx, "GetCurrentProcessorNumberEx"),
		uintptr(unsafe.Pointer(&pn)))
	return pn
}

var _kernel_GetCurrentProcessorNumberEx *syscall.Proc

// [GetCurrentThreadId] function.
//
// [GetCurrentThreadId]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthreadid
func GetCurrentThreadId() uint32 {
	ret, _, _ := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetCurrentThreadId, "GetCurrentThreadId"))
	return uint32(ret)
}

var _kernel_GetCurrentThreadId *syscall.Proc

// [GetEnvironmentStrings] function.
//
// [FreeEnvironmentStrings] is automatically called after the data retrieval.
//
// [GetEnvironmentStrings]: https://learn.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-getenvironmentstringsw
// [FreeEnvironmentStrings]: https://learn.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-freeenvironmentstringsw
func GetEnvironmentStrings() (map[string]string, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetEnvironmentStringsW, "GetEnvironmentStringsW"))
	if ret == 0 {
		return nil, co.ERROR_NOT_CAPABLE
	}
	rawEntries := wstr.DecodeArrPtr((*uint16)(unsafe.Pointer(ret)))

	ret, _, _ = syscall.SyscallN( // free right away
		dll.Kernel.Load(&_kernel_FreeEnvironmentStringsW, "FreeEnvironmentStringsW"),
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

var _kernel_GetEnvironmentStringsW *syscall.Proc
var _kernel_FreeEnvironmentStringsW *syscall.Proc

// [GetFileAttributes] function.
//
// [GetFileAttributes]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfileattributesw
func GetFileAttributes(fileName string) (co.FILE_ATTRIBUTE, error) {
	var wFileName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetFileAttributesW, "GetFileAttributesW"),
		uintptr(wFileName.EmptyIsNil(fileName)))

	if retAttr := co.FILE_ATTRIBUTE(ret); retAttr == co.FILE_ATTRIBUTE_INVALID {
		return retAttr, co.ERROR(err) // err is extended error information
	} else {
		return retAttr, nil
	}
}

var _kernel_GetFileAttributesW *syscall.Proc

// [GetLocalTime] function.
//
// [GetLocalTime]: https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getlocaltime
func GetLocalTime() SYSTEMTIME {
	var st SYSTEMTIME
	syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetLocalTime, "GetLocalTime"),
		uintptr(unsafe.Pointer(&st)))
	return st
}

var _kernel_GetLocalTime *syscall.Proc

// [GetNumberOfConsoleMouseButtons] function.
//
// [GetNumberOfConsoleMouseButtons]: https://learn.microsoft.com/en-us/windows/console/getnumberofconsolemousebuttons
func GetNumberOfConsoleMouseButtons() (int, error) {
	var numberOfMouseButtons uint32
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetNumberOfConsoleMouseButtons, "GetNumberOfConsoleMouseButtons"),
		uintptr(unsafe.Pointer(&numberOfMouseButtons)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return int(numberOfMouseButtons), nil
}

var _kernel_GetNumberOfConsoleMouseButtons *syscall.Proc

// [GetStartupInfo] function.
//
// [GetStartupInfo]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getstartupinfow
func GetStartupInfo() STARTUPINFO {
	var si STARTUPINFO
	si.SetCb()

	syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetStartupInfoW, "GetStartupInfoW"),
		uintptr(unsafe.Pointer(&si)))
	return si
}

var _kernel_GetStartupInfoW *syscall.Proc

// [GetTimeZoneInformation] function.
//
// [GetTimeZoneInformation]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-gettimezoneinformation
func GetTimeZoneInformation() (TIME_ZONE_INFORMATION, co.TIME_ZONE_ID, error) {
	var tzi TIME_ZONE_INFORMATION
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetTimeZoneInformation, "GetTimeZoneInformation"),
		uintptr(unsafe.Pointer(&tzi)))
	if ret == utl.TIME_ZONE_INVALID {
		return TIME_ZONE_INFORMATION{}, co.TIME_ZONE_ID(0), co.ERROR(err)
	}
	return tzi, co.TIME_ZONE_ID(ret), nil
}

var _kernel_GetTimeZoneInformation *syscall.Proc

// [GetSystemInfo] function.
//
// [GetSystemInfo]: https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getsysteminfo
func GetSystemInfo() SYSTEM_INFO {
	var si SYSTEM_INFO
	syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetSystemInfo, "GetSystemInfo"),
		uintptr(unsafe.Pointer(&si)))
	return si
}

var _kernel_GetSystemInfo *syscall.Proc

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
	var wTitle wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_SetConsoleTitleW, "SetConsoleTitleW"),
		uintptr(wTitle.EmptyIsNil(title)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_SetConsoleTitleW *syscall.Proc

// [SetCurrentDirectory] function.
//
// [SetCurrentDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setcurrentdirectory
func SetCurrentDirectory(pathName string) error {
	var wPathName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_SetCurrentDirectoryW, "SetCurrentDirectoryW"),
		uintptr(wPathName.EmptyIsNil(pathName)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_SetCurrentDirectoryW *syscall.Proc

// [Sleep] function.
//
// Example:
//
//	win.Sleep(5 * time.Second)
//
// [Sleep]: https://learn.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-sleep
func Sleep(duration time.Duration) {
	syscall.SyscallN(
		dll.Kernel.Load(&_kernel_Sleep, "Sleep"),
		uintptr(uint32(duration.Milliseconds())))
}

var _kernel_Sleep *syscall.Proc

// [SystemTimeToFileTime] function.
//
// [SystemTimeToFileTime]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetofiletime
func SystemTimeToFileTime(st *SYSTEMTIME) (FILETIME, error) {
	var ft FILETIME
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_SystemTimeToFileTime, "SystemTimeToFileTime"),
		uintptr(unsafe.Pointer(st)),
		uintptr(unsafe.Pointer(&ft)))
	if ret == 0 {
		return FILETIME{}, co.ERROR(err)
	}
	return ft, nil
}

var _kernel_SystemTimeToFileTime *syscall.Proc

// [SystemTimeToTzSpecificLocalTime] function.
//
// [SystemTimeToTzSpecificLocalTime]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetotzspecificlocaltime
func SystemTimeToTzSpecificLocalTime(
	timeZoneInfo *TIME_ZONE_INFORMATION,
	inUniversalTime *SYSTEMTIME,
) (SYSTEMTIME, error) {
	var st SYSTEMTIME
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_SystemTimeToTzSpecificLocalTime, "SystemTimeToTzSpecificLocalTime"),
		uintptr(unsafe.Pointer(timeZoneInfo)),
		uintptr(unsafe.Pointer(inUniversalTime)),
		uintptr(unsafe.Pointer(&st)))
	if ret == 0 {
		return SYSTEMTIME{}, co.ERROR(err)
	}
	return st, nil
}

var _kernel_SystemTimeToTzSpecificLocalTime *syscall.Proc
