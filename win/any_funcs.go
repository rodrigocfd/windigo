package win

import (
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-adjustwindowrectex
func AdjustWindowRectEx(
	lpRect *RECT, dwStyle co.WS, bMenu bool, dwExStyle co.WS_EX) {

	ret, _, err := syscall.Syscall6(proc.AdjustWindowRectEx.Addr(), 4,
		uintptr(unsafe.Pointer(lpRect)), uintptr(dwStyle),
		util.BoolToUintptr(bMenu), uintptr(dwExStyle), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-allowsetforegroundwindow
func AllowSetForegroundWindow(dwProcessId uint32) {
	ret, _, err := syscall.Syscall(proc.AllowSetForegroundWindow.Addr(), 1,
		uintptr(dwProcessId), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms646912(v=vs.85)
func ChooseColor(lpcc *CHOOSECOLOR) bool {
	ret, _, _ := syscall.Syscall(proc.ChooseColor.Addr(), 1,
		uintptr(unsafe.Pointer(lpcc)), 0, 0)
	if ret == 0 {
		dlgErr := CommDlgExtendedError()
		if dlgErr == errco.CDERR(0) {
			return false
		} else {
			panic(dlgErr)
		}
	}
	return true
}

// Loads the COM module. This needs to be done only once in your application.
// Typically uses COINIT_APARTMENTTHREADED.
//
// ⚠️ You must defer CoUninitialize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
func CoInitializeEx(dwCoInit co.COINIT) {
	ret, _, _ := syscall.Syscall(proc.CoInitializeEx.Addr(), 2,
		0, uintptr(dwCoInit), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-commdlgextendederror
func CommDlgExtendedError() errco.CDERR {
	ret, _, _ := syscall.Syscall(proc.CommDlgExtendedError.Addr(), 0,
		0, 0, 0)
	return errco.CDERR(ret)
}

// Typically used with GetCommandLine().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-commandlinetoargvw
func CommandLineToArgv(lpCmdLine string) []string {
	pNumArgs := int32(0)
	ret, _, err := syscall.Syscall(proc.CommandLineToArgv.Addr(), 2,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpCmdLine))),
		uintptr(unsafe.Pointer(&pNumArgs)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}

	lpPtrs := unsafe.Slice((**uint16)(unsafe.Pointer(ret)), pNumArgs) // []*uint16
	strs := make([]string, 0, pNumArgs)

	for _, lpPtr := range lpPtrs {
		strs = append(strs, Str.FromUint16Ptr(lpPtr))
	}
	return strs
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemfree
func CoTaskMemFree(pv unsafe.Pointer) {
	syscall.Syscall(proc.CoTaskMemFree.Addr(), 1,
		uintptr(pv), 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func CoUninitialize() {
	syscall.Syscall(proc.CoUninitialize.Addr(), 0, 0, 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createdirectoryw
func CreateDirectory(
	pathName string, securityAttributes *SECURITY_ATTRIBUTES) error {

	ret, _, err := syscall.Syscall(proc.CreateDirectory.Addr(), 2,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(pathName))),
		uintptr(unsafe.Pointer(securityAttributes)), 0)
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
	lpApplicationName, lpCommandLine string,
	lpProcessAttributes, lpThreadAttributes *SECURITY_ATTRIBUTES,
	bInheritHandles bool,
	dwCreationFlags co.CREATE,
	lpEnvironment uintptr,
	lpCurrentDirectory string,
	lpStartupInfo *STARTUPINFO,
	lpProcessInformation *PROCESS_INFORMATION) {

	ret, _, err := syscall.Syscall12(proc.CreateProcess.Addr(), 10,
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(lpApplicationName))),
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(lpCommandLine))),
		uintptr(unsafe.Pointer(lpProcessAttributes)),
		uintptr(unsafe.Pointer(lpThreadAttributes)),
		util.BoolToUintptr(bInheritHandles),
		uintptr(dwCreationFlags),
		lpEnvironment,
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(lpCurrentDirectory))),
		uintptr(unsafe.Pointer(lpStartupInfo)),
		uintptr(unsafe.Pointer(lpProcessInformation)),
		0, 0)

	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-deletefilew
func DeleteFile(fileName string) error {
	ret, _, err := syscall.Syscall(proc.DeleteFile.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(fileName))), 0, 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycaret
func DestroyCaret() {
	ret, _, err := syscall.Syscall(proc.DestroyCaret.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.Syscall(proc.DispatchMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmgetcolorizationcolor
func DwmGetColorizationColor() (color COLORREF, isOpaqueBlend bool) {
	ret, _, _ := syscall.Syscall(proc.DwmGetColorizationColor.Addr(), 2,
		uintptr(unsafe.Pointer(&color)), uintptr(unsafe.Pointer(&isOpaqueBlend)), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmiscompositionenabled
func DwmIsCompositionEnabled() bool {
	pfEnabled := BOOL(0)
	ret, _, _ := syscall.Syscall(proc.DwmIsCompositionEnabled.Addr(), 1,
		uintptr(unsafe.Pointer(&pfEnabled)), 0, 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return pfEnabled != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endmenu
func EndMenu() {
	ret, _, err := syscall.Syscall(proc.EndMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumwindows
func EnumWindows(lpEnumFunc func(hwnd HWND) bool) {
	ret, _, err := syscall.Syscall(proc.EnumWindows.Addr(), 2,
		syscall.NewCallback(
			func(hwnd HWND, lParam LPARAM) uintptr {
				return util.BoolToUintptr(lpEnumFunc(hwnd))
			}),
		0, 0) // no need to use LPARAM, Go automatically allocs closure contexts in the heap
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-exitprocess
func ExitProcess(uExitCode uint32) {
	syscall.Syscall(proc.ExitProcess.Addr(), 1,
		uintptr(uExitCode), 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-expandenvironmentstringsw
func ExpandEnvironmentStrings(lpSrc string) string {
	ret, _, _ := syscall.Syscall(proc.ExpandEnvironmentStrings.Addr(), 3,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpSrc))), 0, 0)

	buf := make([]uint16, ret)
	ret, _, err := syscall.Syscall(proc.ExpandEnvironmentStrings.Addr(), 3,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpSrc))),
		uintptr(unsafe.Pointer(&buf[0])), ret)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromUint16Slice(buf)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-filetimetosystemtime
func FileTimeToSystemTime(inFileTime *FILETIME, outSystemTime *SYSTEMTIME) {
	ret, _, err := syscall.Syscall(proc.FileTimeToSystemTime.Addr(), 2,
		uintptr(unsafe.Pointer(inFileTime)),
		uintptr(unsafe.Pointer(outSystemTime)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getasynckeystate
func GetAsyncKeyState(virtKeyCode co.VK) uint16 {
	ret, _, _ := syscall.Syscall(proc.GetAsyncKeyState.Addr(), 1,
		uintptr(virtKeyCode), 0, 0)
	return uint16(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcaretpos
func GetCaretPos() RECT {
	rc := RECT{}
	ret, _, err := syscall.Syscall(proc.GetCaretPos.Addr(), 1,
		uintptr(unsafe.Pointer(&rc)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return rc
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-getcommandlinew
func GetCommandLine() string {
	ret, _, _ := syscall.Syscall(proc.GetCommandLine.Addr(), 0,
		0, 0, 0)
	return Str.FromUint16Ptr((*uint16)(unsafe.Pointer(ret)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocessid
func GetCurrentProcessId() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetCurrentProcessId.Addr(), 0,
		0, 0, 0)
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthreadid
func GetCurrentThreadId() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetCurrentThreadId.Addr(), 0,
		0, 0, 0)
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
func GetCursorPos() POINT {
	pt := POINT{}
	ret, _, err := syscall.Syscall(proc.GetCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(&pt)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdialogbaseunits
func GetDialogBaseUnits() (horz, vert uint16) {
	ret, _, _ := syscall.Syscall(proc.GetDialogBaseUnits.Addr(), 0,
		0, 0, 0)
	horz = Bytes.Lo16(uint32(ret))
	vert = Bytes.Hi16(uint32(ret))
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-getdynamictimezoneinformation
func GetDynamicTimeZoneInformation(
	pTimeZoneInformation *DYNAMIC_TIME_ZONE_INFORMATION) co.TIME_ZONE_ID {

	ret, _, _ := syscall.Syscall(proc.GetDynamicTimeZoneInformation.Addr(), 1,
		uintptr(unsafe.Pointer(pTimeZoneInformation)), 0, 0)
	return co.TIME_ZONE_ID(ret)
}

// You don't need to call FreeEnvironmentStrings(), it's automatically called.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-getenvironmentstringsw
func GetEnvironmentStrings() map[string]string {
	ret, _, err := syscall.Syscall(proc.GetEnvironmentStrings.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	rawEntries := Str.FromUint16PtrMulti((*uint16)(unsafe.Pointer(ret)))

	ret, _, err = syscall.Syscall(proc.FreeEnvironmentStrings.Addr(), 1,
		ret, 0, 0)
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
func GetFileAttributes(lpFileName string) (co.FILE_ATTRIBUTE, error) {
	ret, _, err := syscall.Syscall(proc.GetFileAttributes.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpFileName))), 0, 0)

	if retAttr := co.FILE_ATTRIBUTE(ret); retAttr == co.FILE_ATTRIBUTE_INVALID {
		return retAttr, errco.ERROR(err)
	} else {
		return retAttr, nil
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfow
func GetFileVersionInfo(lptstrFilename string) []byte {
	visz := GetFileVersionInfoSize(lptstrFilename)
	buf := make([]byte, visz)

	ret, _, err := syscall.Syscall6(proc.GetFileVersionInfo.Addr(), 4,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lptstrFilename))),
		0, uintptr(visz), uintptr(unsafe.Pointer(&buf[0])), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return buf
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfosizew
func GetFileVersionInfoSize(lptstrFilename string) uint32 {
	lpdwHandle := uint32(0)
	ret, _, err := syscall.Syscall(proc.GetFileVersionInfoSize.Addr(), 2,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lptstrFilename))),
		uintptr(unsafe.Pointer(&lpdwHandle)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagew
func GetMessage(
	msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax uint32) (int32, error) {

	ret, _, err := syscall.Syscall6(proc.GetMessage.Addr(), 4,
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax),
		0, 0)
	if int(ret) == -1 {
		return 0, errco.ERROR(err)
	}
	return int32(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagepos
func GetMessagePos() POINT {
	ret, _, _ := syscall.Syscall(proc.GetMessagePos.Addr(), 0,
		0, 0, 0)
	return POINT{
		X: int32(Bytes.Lo16(uint32(ret))),
		Y: int32(Bytes.Hi16(uint32(ret))),
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagetime
func GetMessageTime() time.Duration {
	ret, _, _ := syscall.Syscall(proc.GetMessageTime.Addr(), 0,
		0, 0, 0)
	return time.Duration(ret * uintptr(time.Millisecond))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getphysicalcursorpos
func GetPhysicalCursorPos() POINT {
	pt := POINT{}
	ret, _, err := syscall.Syscall(proc.GetPhysicalCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(&pt)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getqueuestatus
func GetQueueStatus(flags co.QS) uint32 {
	ret, _, _ := syscall.Syscall(proc.GetQueueStatus.Addr(), 1,
		uintptr(flags), 0, 0)
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getstartupinfow
func GetStartupInfo(lpStartupInfo *STARTUPINFO) {
	syscall.Syscall(proc.GetStartupInfo.Addr(), 1,
		uintptr(unsafe.Pointer(lpStartupInfo)), 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
func GetSysColor(nIndex co.COLOR) COLORREF {
	ret, _, _ := syscall.Syscall(proc.GetSysColor.Addr(), 1,
		uintptr(nIndex), 0, 0)
	return COLORREF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getsysteminfo
func GetSystemInfo(lpSystemInfo *SYSTEM_INFO) {
	syscall.Syscall(proc.GetSystemInfo.Addr(), 1,
		uintptr(unsafe.Pointer(lpSystemInfo)), 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
func GetSystemMetrics(index co.SM) int32 {
	ret, _, _ := syscall.Syscall(proc.GetSystemMetrics.Addr(), 1,
		uintptr(index), 0, 0)
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getsystemtime
func GetSystemTime(lpSystemTime *SYSTEMTIME) {
	syscall.Syscall(proc.GetSystemTime.Addr(), 1,
		uintptr(unsafe.Pointer(lpSystemTime)), 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getsystemtimes
func GetSystemTimes(lpIdleTime, lpKernelTime, lpUserTime *FILETIME) {
	ret, _, err := syscall.Syscall(proc.GetSystemTimes.Addr(), 3,
		uintptr(unsafe.Pointer(lpIdleTime)), uintptr(unsafe.Pointer(lpKernelTime)),
		uintptr(unsafe.Pointer(lpUserTime)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getsystemtimeasfiletime
func GetSystemTimeAsFileTime() FILETIME {
	ft := FILETIME{}
	syscall.Syscall(proc.GetSystemTimeAsFileTime.Addr(), 1,
		uintptr(unsafe.Pointer(&ft)), 0, 0)
	return ft
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getsystemtimepreciseasfiletime
func GetSystemTimePreciseAsFileTime() FILETIME {
	ft := FILETIME{}
	syscall.Syscall(proc.GetSystemTimePreciseAsFileTime.Addr(), 1,
		uintptr(unsafe.Pointer(&ft)), 0, 0)
	return ft
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-gettickcount64
func GetTickCount64() uint64 {
	ret, _, _ := syscall.Syscall(proc.GetTickCount64.Addr(), 0,
		0, 0, 0)
	return uint64(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-gettimezoneinformation
func GetTimeZoneInformation(
	lpTimeZoneInformation *TIME_ZONE_INFORMATION) co.TIME_ZONE_ID {

	ret, _, _ := syscall.Syscall(proc.GetTimeZoneInformation.Addr(), 1,
		uintptr(unsafe.Pointer(lpTimeZoneInformation)), 0, 0)
	return co.TIME_ZONE_ID(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-gettimezoneinformationforyear
func GetTimeZoneInformationForYear(
	wYear uint16,
	pdtzi *DYNAMIC_TIME_ZONE_INFORMATION, ptzi *TIME_ZONE_INFORMATION) {

	ret, _, err := syscall.Syscall(proc.GetTimeZoneInformationForYear.Addr(), 3,
		uintptr(wYear),
		uintptr(unsafe.Pointer(pdtzi)), uintptr(unsafe.Pointer(ptzi)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrols
func InitCommonControls() {
	syscall.Syscall(proc.InitCommonControls.Addr(), 0, 0, 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isappthemed
func IsAppThemed() bool {
	ret, _, _ := syscall.Syscall(proc.IsAppThemed.Addr(), 0,
		0, 0, 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-iscompositionactive
func IsCompositionActive() bool {
	ret, _, _ := syscall.Syscall(proc.IsCompositionActive.Addr(), 0,
		0, 0, 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isguithread
func IsGUIThread(bConvertToGuiThread bool) (bool, error) {
	ret, _, _ := syscall.Syscall(proc.IsGUIThread.Addr(), 1,
		util.BoolToUintptr(bConvertToGuiThread), 0, 0)
	if bConvertToGuiThread && errco.ERROR(ret) == errco.NOT_ENOUGH_MEMORY {
		return false, errco.NOT_ENOUGH_MEMORY
	}
	return ret != 0, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemeactive
func IsThemeActive() bool {
	ret, _, _ := syscall.Syscall(proc.IsThemeActive.Addr(), 0,
		0, 0, 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows10orgreater
func IsWindows10OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(Bytes.Hi8(uint16(co.WIN32_WINNT_WINTHRESHOLD))),
		uint32(Bytes.Lo8(uint16(co.WIN32_WINNT_WINTHRESHOLD))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows7orgreater
func IsWindows7OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(Bytes.Hi8(uint16(co.WIN32_WINNT_WIN7))),
		uint32(Bytes.Lo8(uint16(co.WIN32_WINNT_WIN7))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows8orgreater
func IsWindows8OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(Bytes.Hi8(uint16(co.WIN32_WINNT_WIN8))),
		uint32(Bytes.Lo8(uint16(co.WIN32_WINNT_WIN8))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows8point1orgreater
func IsWindows8Point1OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(Bytes.Hi8(uint16(co.WIN32_WINNT_WINBLUE))),
		uint32(Bytes.Lo8(uint16(co.WIN32_WINNT_WINBLUE))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsvistaorgreater
func IsWindowsVistaOrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(Bytes.Hi8(uint16(co.WIN32_WINNT_VISTA))),
		uint32(Bytes.Lo8(uint16(co.WIN32_WINNT_VISTA))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsxporgreater
func IsWindowsXpOrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(Bytes.Hi8(uint16(co.WIN32_WINNT_WINXP))),
		uint32(Bytes.Lo8(uint16(co.WIN32_WINNT_WINXP))),
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

	ret, _ := VerifyVersionInfo(&ovi,
		co.VER_MAJORVERSION|co.VER_MINORVERSION|co.VER_SERVICEPACKMAJOR,
		conditionMask)
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-locksetforegroundwindow
func LockSetForegroundWindow(uLockCode co.LSFW) {
	ret, _, err := syscall.Syscall(proc.LockSetForegroundWindow.Addr(), 1,
		uintptr(uLockCode), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfrompoint
func MonitorFromPoint(pt POINT, dwFlags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.Syscall(proc.MonitorFromPoint.Addr(), 3,
		uintptr(pt.X), uintptr(pt.Y), uintptr(dwFlags))
	return HMONITOR(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-muldiv
func MulDiv(number, numerator, denominator int32) int32 {
	ret, _, _ := syscall.Syscall(proc.MulDiv.Addr(), 3,
		uintptr(number), uintptr(numerator), uintptr(denominator))
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-peekmessagew
func PeekMessage(
	lpMsg *MSG, hWnd HWND,
	wMsgFilterMin, wMsgFilterMax co.WM, wRemoveMsg co.PM) bool {

	ret, _, _ := syscall.Syscall6(proc.PeekMessageW.Addr(), 5,
		uintptr(unsafe.Pointer(lpMsg)), uintptr(hWnd),
		uintptr(wMsgFilterMin), uintptr(wMsgFilterMax), uintptr(wRemoveMsg), 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int32) {
	syscall.Syscall(proc.PostQuitMessage.Addr(), 1,
		uintptr(exitCode), 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessage(
	idThread uint32, Msg co.WM, wParam WPARAM, lParam LPARAM) {

	ret, _, err := syscall.Syscall6(proc.PostThreadMessage.Addr(), 4,
		uintptr(idThread), uintptr(Msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/profileapi/nf-profileapi-queryperformancecounter
func QueryPerformanceCounter() int64 {
	lpPerformanceCount := int64(0)
	ret, _, err := syscall.Syscall(proc.QueryPerformanceCounter.Addr(), 1,
		uintptr(unsafe.Pointer(&lpPerformanceCount)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return lpPerformanceCount
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/profileapi/nf-profileapi-queryperformancefrequency
func QueryPerformanceFrequency() int64 {
	lpFrequency := int64(0)
	ret, _, err := syscall.Syscall(proc.QueryPerformanceFrequency.Addr(), 1,
		uintptr(unsafe.Pointer(&lpFrequency)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return lpFrequency
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wcx *WNDCLASSEX) (ATOM, error) {
	wcx.SetCbSize() // safety
	ret, _, err := syscall.Syscall(proc.RegisterClassEx.Addr(), 1,
		uintptr(unsafe.Pointer(wcx)), 0, 0)
	if ret == 0 {
		return ATOM(0), errco.ERROR(err)
	}
	return ATOM(ret), nil
}

// Available in Windows Vista.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiaware
func SetProcessDPIAware() {
	ret, _, _ := syscall.Syscall(proc.SetProcessDPIAware.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic("SetProcessDPIAware failed.")
	}
}

// Available in Windows 10, version 1703.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiawarenesscontext
func SetProcessDpiAwarenessContext(value co.DPI_AWARE_CTX) {
	ret, _, err := syscall.Syscall(proc.SetProcessDpiAwarenessContext.Addr(), 1,
		uintptr(value), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw
func ShellNotifyIcon(dwMessage co.NIM, lpData *NOTIFYICONDATA) {
	ret, _, err := syscall.Syscall(proc.Shell_NotifyIcon.Addr(), 2,
		uintptr(dwMessage), uintptr(unsafe.Pointer(lpData)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// Depends of CoInitializeEx().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
func SHGetFileInfo(
	pszPath string, dwFileAttributes co.FILE_ATTRIBUTE,
	psfi *SHFILEINFO, uFlags co.SHGFI) {

	ret, _, err := syscall.Syscall6(proc.SHGetFileInfo.Addr(), 5,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(pszPath))),
		uintptr(dwFileAttributes), uintptr(unsafe.Pointer(psfi)),
		unsafe.Sizeof(*psfi), uintptr(uFlags), 0)

	if (uFlags&co.SHGFI_EXETYPE) == 0 || (uFlags&co.SHGFI_SYSICONINDEX) == 0 {
		if ret == 0 {
			panic(errco.ERROR(err))
		}
	}

	if (uFlags & co.SHGFI_EXETYPE) != 0 {
		if ret == 0 {
			panic(errco.ERROR(err))
		}
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-sleep
func Sleep(dwMilliseconds uint32) {
	syscall.Syscall(proc.Sleep.Addr(), 1,
		uintptr(dwMilliseconds), 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
func SystemParametersInfo(
	uiAction co.SPI, uiParam uint32, pvParam unsafe.Pointer, fWinIni co.SPIF) {

	ret, _, err := syscall.Syscall6(proc.SystemParametersInfo.Addr(), 4,
		uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetofiletime
func SystemTimeToFileTime(inSystemTime *SYSTEMTIME, outFileTime *FILETIME) {
	ret, _, err := syscall.Syscall(proc.SystemTimeToFileTime.Addr(), 2,
		uintptr(unsafe.Pointer(inSystemTime)),
		uintptr(unsafe.Pointer(outFileTime)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetotzspecificlocaltime
func SystemTimeToTzSpecificLocalTime(
	lpTimeZoneInformation *TIME_ZONE_INFORMATION,
	inUniversalTime *SYSTEMTIME, outLocalTime *SYSTEMTIME) {

	ret, _, err := syscall.Syscall(proc.SystemTimeToTzSpecificLocalTime.Addr(), 3,
		uintptr(unsafe.Pointer(lpTimeZoneInformation)),
		uintptr(unsafe.Pointer(inUniversalTime)),
		uintptr(unsafe.Pointer(outLocalTime)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialogindirect
func TaskDialogIndirect(pTaskConfig *TASKDIALOGCONFIG) co.ID {
	pnButton := co.ID(0)
	ret, _, _ := syscall.Syscall6(proc.TaskDialogIndirect.Addr(), 4,
		uintptr(unsafe.Pointer(pTaskConfig)), uintptr(unsafe.Pointer(&pnButton)),
		uintptr(0), uintptr(0), 0, 0)
	if wErr := errco.ERROR(ret); wErr != errco.S_OK {
		panic(wErr)
	}
	return co.ID(pnButton)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(proc.TranslateMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-tzspecificlocaltimetosystemtime
func TzSpecificLocalTimeToSystemTime(
	lpTimeZoneInformation *TIME_ZONE_INFORMATION,
	inLocalTime *SYSTEMTIME, outUniversalTime *SYSTEMTIME) {

	ret, _, err := syscall.Syscall(proc.TzSpecificLocalTimeToSystemTime.Addr(), 3,
		uintptr(unsafe.Pointer(lpTimeZoneInformation)),
		uintptr(unsafe.Pointer(inLocalTime)),
		uintptr(unsafe.Pointer(outUniversalTime)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winver/nf-winver-verqueryvaluew
func VerQueryValue(pBlock []byte, lpSubBlock string) ([]byte, bool) {
	lplpBuffer, puLen := uintptr(0), uint32(0)
	ret, _, _ := syscall.Syscall6(proc.VerQueryValue.Addr(), 4,
		uintptr(unsafe.Pointer(&pBlock[0])),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpSubBlock))),
		uintptr(unsafe.Pointer(&lplpBuffer)), uintptr(unsafe.Pointer(&puLen)),
		0, 0)
	if ret == 0 {
		return nil, false
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(lplpBuffer)), puLen), true
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-verifyversioninfow
func VerifyVersionInfo(
	ovi *OSVERSIONINFOEX,
	typeMask co.VER, conditionMask uint64) (bool, errco.ERROR) {

	ret, _, err := syscall.Syscall(proc.VerifyVersionInfo.Addr(), 3,
		uintptr(unsafe.Pointer(ovi)),
		uintptr(typeMask), uintptr(conditionMask))
	return ret != 0, errco.ERROR(err)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-versetconditionmask
func VerSetConditionMask(
	conditionMask uint64, typeMask co.VER, condition co.VER_COND) uint64 {

	ret, _, _ := syscall.Syscall(proc.VerSetConditionMask.Addr(), 3,
		uintptr(conditionMask), uintptr(typeMask), uintptr(condition))
	return uint64(ret)
}
