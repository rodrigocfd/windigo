//go:build windows

package win

import (
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// ⚠️ You must defer FreeConsole().
//
// 📑 https://docs.microsoft.com/en-us/windows/console/allocconsole
func AllocConsole() error {
	ret, _, err := syscall.SyscallN(proc.AllocConsole.Addr())
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// ⚠️ You must defer FreeConsole().
//
// 📑 https://docs.microsoft.com/en-us/windows/console/attachconsole
func AttachConsole(processId uint32) error {
	ret, _, err := syscall.SyscallN(proc.AttachConsole.Addr(),
		uintptr(processId))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-copyfilew
func CopyFile(existingFile, newFile string, failIfExists bool) error {
	ret, _, err := syscall.SyscallN(proc.CopyFile.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(existingFile))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(newFile))),
		util.BoolToUintptr(failIfExists))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createdirectoryw
func CreateDirectory(
	pathName string, securityAttributes *SECURITY_ATTRIBUTES) error {

	ret, _, err := syscall.SyscallN(proc.CreateDirectory.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(pathName))),
		uintptr(unsafe.Pointer(securityAttributes)))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// ⚠️ You must defer HPROCESS.CloseHandle() and HTHREAD.CloseHandle() on
// HProcess and HThread members of PROCESS_INFORMATION.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-createprocessw
func CreateProcess(
	applicationName, commandLine StrOpt,
	processAttributes, threadAttributes *SECURITY_ATTRIBUTES,
	inheritHandles bool,
	creationFlags co.CREATE,
	environment []struct {
		name string
		val  string
	},
	currentDirectory StrOpt,
	startupInfo *STARTUPINFO,
	processInformation *PROCESS_INFORMATION) {

	var envStrsPtr unsafe.Pointer
	if environment != nil {
		envStrs := make([]string, 0, len(environment))
		for _, pair := range environment {
			envStrs = append(envStrs, pair.name+"="+pair.val)
		}
		envStrsPtr = unsafe.Pointer(Str.ToNativePtrMulti(envStrs))
	}

	ret, _, err := syscall.SyscallN(proc.CreateProcess.Addr(),
		uintptr(applicationName.Raw()),
		uintptr(commandLine.Raw()),
		uintptr(unsafe.Pointer(processAttributes)),
		uintptr(unsafe.Pointer(threadAttributes)),
		util.BoolToUintptr(inheritHandles),
		uintptr(creationFlags),
		uintptr(envStrsPtr),
		uintptr(currentDirectory.Raw()),
		uintptr(unsafe.Pointer(startupInfo)),
		uintptr(unsafe.Pointer(processInformation)))

	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-deletefilew
func DeleteFile(fileName string) error {
	ret, _, err := syscall.SyscallN(proc.DeleteFile.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(fileName))))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-exitprocess
func ExitProcess(exitCode uint32) {
	syscall.SyscallN(proc.ExitProcess.Addr(),
		uintptr(exitCode))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-expandenvironmentstringsw
func ExpandEnvironmentStrings(src string) string {
	pSrc := Str.ToNativePtr(src)
	ret, _, _ := syscall.SyscallN(proc.ExpandEnvironmentStrings.Addr(),
		uintptr(unsafe.Pointer(pSrc)), 0, 0)

	buf := make([]uint16, ret)
	ret, _, err := syscall.SyscallN(proc.ExpandEnvironmentStrings.Addr(),
		uintptr(unsafe.Pointer(pSrc)),
		uintptr(unsafe.Pointer(&buf[0])), ret)
	runtime.KeepAlive(pSrc)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromNativeSlice(buf)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-filetimetosystemtime
func FileTimeToSystemTime(inFileTime *FILETIME, outSystemTime *SYSTEMTIME) {
	ret, _, err := syscall.SyscallN(proc.FileTimeToSystemTime.Addr(),
		uintptr(unsafe.Pointer(inFileTime)),
		uintptr(unsafe.Pointer(outSystemTime)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/console/freeconsole
func FreeConsole() error {
	ret, _, err := syscall.SyscallN(proc.FreeConsole.Addr())
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-getcommandlinew
func GetCommandLine() string {
	ret, _, _ := syscall.SyscallN(proc.GetCommandLine.Addr())
	return Str.FromNativePtr((*uint16)(unsafe.Pointer(ret)))
}

// 📑 https://docs.microsoft.com/en-us/windows/console/getconsolecp
func GetConsoleCP() (co.CP, error) {
	ret, _, err := syscall.SyscallN(proc.GetConsoleCP.Addr())
	if ret == 0 {
		return co.CP(0), errco.ERROR(err)
	}
	return co.CP(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/console/getconsoletitle
func GetConsoleTitle() (string, error) {
	const BUF_SZ = _MAX_PATH * 2
	buf := make([]uint16, BUF_SZ)

	ret, _, err := syscall.SyscallN(proc.GetConsoleTitle.Addr(),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(BUF_SZ))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return "", wErr
	}
	return Str.FromNativeSlice(buf), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/console/getconsolewindow
func GetConsoleWindow() HWND {
	ret, _, _ := syscall.SyscallN(proc.GetConsoleWindow.Addr())
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-getcurrentdirectory
func GetCurrentDirectory() string {
	var buf [_MAX_PATH + 1]uint16
	ret, _, err := syscall.SyscallN(proc.GetCurrentDirectory.Addr(),
		uintptr(len(buf)), uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromNativeSlice(buf[:])
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-getdynamictimezoneinformation
func GetDynamicTimeZoneInformation(
	timeZoneInfo *DYNAMIC_TIME_ZONE_INFORMATION) co.TIME_ZONE_ID {

	ret, _, _ := syscall.SyscallN(proc.GetDynamicTimeZoneInformation.Addr(),
		uintptr(unsafe.Pointer(timeZoneInfo)))
	return co.TIME_ZONE_ID(ret)
}

// You don't need to call FreeEnvironmentStrings(), it's automatically called.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-getenvironmentstringsw
func GetEnvironmentStrings() map[string]string {
	ret, _, err := syscall.SyscallN(proc.GetEnvironmentStrings.Addr())
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	rawEntries := Str.FromNativePtrMulti((*uint16)(unsafe.Pointer(ret)))

	ret, _, err = syscall.SyscallN(proc.FreeEnvironmentStrings.Addr(),
		ret)
	if ret == 0 {
		panic(errco.ERROR(err))
	}

	mapEntries := make(map[string]string, len(rawEntries))
	for _, entry := range rawEntries {
		keyVal := strings.SplitN(entry, "=", 2)
		mapEntries[keyVal[0]] = keyVal[1]
	}
	return mapEntries
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfileattributesw
func GetFileAttributes(fileName string) (co.FILE_ATTRIBUTE, error) {
	ret, _, err := syscall.SyscallN(proc.GetFileAttributes.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(fileName))))

	if retAttr := co.FILE_ATTRIBUTE(ret); retAttr == co.FILE_ATTRIBUTE_INVALID {
		return retAttr, errco.ERROR(err) // err is extended error information
	} else {
		return retAttr, nil
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getstartupinfow
func GetStartupInfo(startupInfo *STARTUPINFO) {
	syscall.SyscallN(proc.GetStartupInfo.Addr(),
		uintptr(unsafe.Pointer(startupInfo)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getsysteminfo
func GetSystemInfo(systemInfo *SYSTEM_INFO) {
	syscall.SyscallN(proc.GetSystemInfo.Addr(),
		uintptr(unsafe.Pointer(systemInfo)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getsystemtime
func GetSystemTime(systemTime *SYSTEMTIME) {
	syscall.SyscallN(proc.GetSystemTime.Addr(),
		uintptr(unsafe.Pointer(systemTime)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getsystemtimes
func GetSystemTimes(idleTime, kernelTime, userTime *FILETIME) {
	ret, _, err := syscall.SyscallN(proc.GetSystemTimes.Addr(),
		uintptr(unsafe.Pointer(idleTime)), uintptr(unsafe.Pointer(kernelTime)),
		uintptr(unsafe.Pointer(userTime)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getsystemtimeasfiletime
func GetSystemTimeAsFileTime() FILETIME {
	var ft FILETIME
	syscall.SyscallN(proc.GetSystemTimeAsFileTime.Addr(),
		uintptr(unsafe.Pointer(&ft)))
	return ft
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getsystemtimepreciseasfiletime
func GetSystemTimePreciseAsFileTime() FILETIME {
	var ft FILETIME
	syscall.SyscallN(proc.GetSystemTimePreciseAsFileTime.Addr(),
		uintptr(unsafe.Pointer(&ft)))
	return ft
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-gettimezoneinformation
func GetTimeZoneInformation(
	timeZoneInfo *TIME_ZONE_INFORMATION) co.TIME_ZONE_ID {

	ret, _, _ := syscall.SyscallN(proc.GetTimeZoneInformation.Addr(),
		uintptr(unsafe.Pointer(timeZoneInfo)))
	return co.TIME_ZONE_ID(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-gettimezoneinformationforyear
func GetTimeZoneInformationForYear(
	wYear uint16,
	dtzi *DYNAMIC_TIME_ZONE_INFORMATION, tzi *TIME_ZONE_INFORMATION) {

	ret, _, err := syscall.SyscallN(proc.GetTimeZoneInformationForYear.Addr(),
		uintptr(wYear),
		uintptr(unsafe.Pointer(dtzi)), uintptr(unsafe.Pointer(tzi)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

type _VolumeInfo struct {
	Name               string
	SerialNumber       uint32
	MaxComponentLength uint32
	FileSystemFlags    co.FILE_VOL
	FileSystemName     string
}

// Example:
//
//	nfo, err := win.GetVolumeInformation(win.StrOptSome("C:\\"))
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Printf("Name: %s\n", nfo.Name)
//	fmt.Printf("File system name: %s\n", nfo.FileSystemName)
//	fmt.Printf("Max component length: %d\n", nfo.MaxComponentLength)
//	fmt.Printf("Serial number: 0x08%x\n", nfo.SerialNumber)
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getvolumeinformationw
func GetVolumeInformation(rootPathName StrOpt) (_VolumeInfo, error) {
	var info _VolumeInfo
	var nameBuf [_MAX_PATH + 1]uint16
	var sysNameBuf [_MAX_PATH + 1]uint16

	ret, _, err := syscall.SyscallN(proc.GetVolumeInformation.Addr(),
		uintptr(rootPathName.Raw()),
		uintptr(unsafe.Pointer(&nameBuf[0])), _MAX_PATH+1,
		uintptr(unsafe.Pointer(&info.SerialNumber)),
		uintptr(unsafe.Pointer(&info.MaxComponentLength)),
		uintptr(unsafe.Pointer(&info.FileSystemFlags)),
		uintptr(unsafe.Pointer(&sysNameBuf[0])), _MAX_PATH+1)

	if ret == 0 {
		return _VolumeInfo{}, errco.ERROR(err)
	}

	info.Name = Str.FromNativeSlice(nameBuf[:])
	info.FileSystemName = Str.FromNativeSlice(sysNameBuf[:])
	return info, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getwindowsdirectoryw
func GetWindowsDirectory() string {
	var buf [_MAX_PATH + 1]uint16
	ret, _, err := syscall.SyscallN(proc.GetWindowsDirectory.Addr(),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromNativeSlice(buf[:])
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632656(v=vs.85)
func HIBYTE(val uint16) uint8 {
	_, hi := util.Break16(val)
	return hi
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632657(v=vs.85)
func HIWORD(val uint32) uint16 {
	_, hi := util.Break32(val)
	return hi
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows10orgreater
func IsWindows10OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_WINTHRESHOLD))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_WINTHRESHOLD))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows7orgreater
func IsWindows7OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_WIN7))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_WIN7))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows8orgreater
func IsWindows8OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_WIN8))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_WIN8))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows8point1orgreater
func IsWindows8Point1OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_WINBLUE))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_WINBLUE))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsvistaorgreater
func IsWindowsVistaOrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_VISTA))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_VISTA))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsxporgreater
func IsWindowsXpOrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HIBYTE(uint16(co.WIN32_WINNT_WINXP))),
		uint32(LOBYTE(uint16(co.WIN32_WINNT_WINXP))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsversionorgreater
func IsWindowsVersionOrGreater(
	majorVersion, minorVersion uint32, servicePackMajor uint16) bool {

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
		panic(err)
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632658(v=vs.85)
func LOBYTE(val uint16) uint8 {
	lo, _ := util.Break16(val)
	return lo
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632659(v=vs.85)
func LOWORD(val uint32) uint16 {
	lo, _ := util.Break32(val)
	return lo
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632660(v=vs.85)
func MAKELONG(lo, hi uint16) uint32 {
	return util.Make32(lo, hi)
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632663(v=vs.85)
func MAKEWORD(lo, hi uint8) uint16 {
	return util.Make16(lo, hi)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-movefilew
func MoveFile(existingFile, newFile string) error {
	ret, _, err := syscall.SyscallN(proc.MoveFile.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(existingFile))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(newFile))))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-movefileexw
func MoveFileEx(existingFile, newFile string, flags co.MOVEFILE) error {
	ret, _, err := syscall.SyscallN(proc.MoveFileEx.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(existingFile))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(newFile))),
		uintptr(flags))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// Note: You'll achieve a much better performance with ordinary Go code:
//
//	res := int32((int64(n) * int64(num)) / int64(den))
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-muldiv
func MulDiv(number, numerator, denominator int32) int32 {
	ret, _, _ := syscall.SyscallN(proc.MulDiv.Addr(),
		uintptr(number), uintptr(numerator), uintptr(denominator))
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/profileapi/nf-profileapi-queryperformancecounter
func QueryPerformanceCounter() int64 {
	var lpPerformanceCount int64
	ret, _, err := syscall.SyscallN(proc.QueryPerformanceCounter.Addr(),
		uintptr(unsafe.Pointer(&lpPerformanceCount)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return lpPerformanceCount
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/profileapi/nf-profileapi-queryperformancefrequency
func QueryPerformanceFrequency() int64 {
	var lpFrequency int64
	ret, _, err := syscall.SyscallN(proc.QueryPerformanceFrequency.Addr(),
		uintptr(unsafe.Pointer(&lpFrequency)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return lpFrequency
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-removedirectoryw
func RemoveDirectory(pathName string) error {
	ret, _, err := syscall.SyscallN(proc.RemoveDirectory.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(pathName))))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-replacefilew
func ReplaceFile(
	replaced, replacement string,
	backup StrOpt, replaceFlags co.REPLACEFILE) error {

	ret, _, err := syscall.SyscallN(proc.ReplaceFile.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(replaced))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(replacement))),
		uintptr(backup.Raw()), uintptr(replaceFlags), 0, 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/console/setconsoleoutputcp
func SetConsoleOutputCP(codePage co.CP) error {
	ret, _, err := syscall.SyscallN(proc.SetConsoleOutputCP.Addr(),
		uintptr(codePage))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/console/setconsoletitle
func SetConsoleTitle(title string) error {
	ret, _, err := syscall.SyscallN(proc.SetConsoleTitle.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(title))))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setcurrentdirectory
func SetCurrentDirectory(pathName string) error {
	ret, _, err := syscall.SyscallN(proc.SetCurrentDirectory.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(pathName))))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfileattributesw
func SetFileAttributes(fileName string, attrs co.FILE_ATTRIBUTE) error {
	ret, _, err := syscall.SyscallN(proc.SetFileAttributes.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(fileName))), uintptr(attrs))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-sleep
func Sleep(milliseconds uint32) {
	syscall.SyscallN(proc.Sleep.Addr(),
		uintptr(milliseconds))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
func SystemParametersInfo(
	uiAction co.SPI, uiParam uint32, pvParam unsafe.Pointer, fWinIni co.SPIF) {

	ret, _, err := syscall.SyscallN(proc.SystemParametersInfo.Addr(),
		uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetofiletime
func SystemTimeToFileTime(inSystemTime *SYSTEMTIME, outFileTime *FILETIME) {
	ret, _, err := syscall.SyscallN(proc.SystemTimeToFileTime.Addr(),
		uintptr(unsafe.Pointer(inSystemTime)),
		uintptr(unsafe.Pointer(outFileTime)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetotzspecificlocaltime
func SystemTimeToTzSpecificLocalTime(
	timeZoneInfo *TIME_ZONE_INFORMATION,
	inUniversalTime *SYSTEMTIME, outLocalTime *SYSTEMTIME) {

	ret, _, err := syscall.SyscallN(proc.SystemTimeToTzSpecificLocalTime.Addr(),
		uintptr(unsafe.Pointer(timeZoneInfo)),
		uintptr(unsafe.Pointer(inUniversalTime)),
		uintptr(unsafe.Pointer(outLocalTime)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-tzspecificlocaltimetosystemtime
func TzSpecificLocalTimeToSystemTime(
	timeZoneInfo *TIME_ZONE_INFORMATION,
	inLocalTime *SYSTEMTIME, outUniversalTime *SYSTEMTIME) {

	ret, _, err := syscall.SyscallN(proc.TzSpecificLocalTimeToSystemTime.Addr(),
		uintptr(unsafe.Pointer(timeZoneInfo)),
		uintptr(unsafe.Pointer(inLocalTime)),
		uintptr(unsafe.Pointer(outUniversalTime)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
