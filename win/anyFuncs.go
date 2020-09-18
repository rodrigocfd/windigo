/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win/proc"
)

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateInstance(rclsid *GUID, pUnkOuter unsafe.Pointer,
	dwClsContext co.CLSCTX, riid *GUID) (unsafe.Pointer, co.ERROR) {

	var ppv unsafe.Pointer = nil
	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(rclsid)), uintptr(pUnkOuter),
		uintptr(dwClsContext), uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppv)), 0)
	return ppv, co.ERROR(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
//
// Must be freed with CoUninitialize().
func CoInitializeEx(dwCoInit co.COINIT) {
	hr, _, _ := syscall.Syscall(proc.CoInitializeEx.Addr(), 2,
		0, uintptr(dwCoInit), 0)
	hr2 := co.ERROR(hr)
	if hr2 != co.ERROR_S_OK && hr2 != co.ERROR_S_FALSE {
		panic(NewWinError(hr2, "CoInitializeEx").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func CoUninitialize() {
	syscall.Syscall(proc.CoUninitialize.Addr(), 0, 0, 0, 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycaret
func DestroyCaret() {
	ret, _, lerr := syscall.Syscall(proc.DestroyCaret.Addr(), 0, 0, 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "DestroyCaret").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.Syscall(proc.DispatchMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endmenu
func EndMenu() {
	ret, _, lerr := syscall.Syscall(proc.EndMenu.Addr(), 0, 0, 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "EndMenu").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumwindows
func EnumWindows(
	lpEnumFunc func(hwnd HWND, lParam LPARAM) bool,
	lParam LPARAM) {

	ret, _, lerr := syscall.Syscall(proc.EnumWindows.Addr(), 2,
		syscall.NewCallback(
			func(hwnd HWND, lParam LPARAM) int32 {
				return _Win.BoolToInt32(lpEnumFunc(hwnd, lParam))
			}),
		uintptr(lParam), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "EnumWindow").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-filetimetosystemtime
func FileTimeToSystemTime(inFileTime *FILETIME, outSystemTime *SYSTEMTIME) {
	ret, _, lerr := syscall.Syscall(proc.FileTimeToSystemTime.Addr(), 2,
		uintptr(unsafe.Pointer(inFileTime)),
		uintptr(unsafe.Pointer(outSystemTime)), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "FileTimeToSystemTime").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getasynckeystate
func GetAsyncKeyState(virtKeyCode co.VK) uint16 {
	ret, _, _ := syscall.Syscall(proc.GetAsyncKeyState.Addr(), 1,
		uintptr(virtKeyCode), 0, 0)
	return uint16(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcaretpos
func GetCaretPos() *RECT {
	rc := &RECT{}
	ret, _, lerr := syscall.Syscall(proc.GetCaretPos.Addr(), 1,
		uintptr(unsafe.Pointer(rc)), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "GetCaretPos").Error())
	}
	return rc
}

// https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthreadid
func GetCurrentThreadId() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetCurrentThreadId.Addr(), 0,
		0, 0, 0)
	return uint32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
func GetCursorPos() *POINT {
	pt := &POINT{}
	ret, _, lerr := syscall.Syscall(proc.GetCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(pt)), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "GetCursorPos").Error())
	}
	return pt
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdpiforsystem
//
// Available in Windows 10, version 1607.
func GetDpiForSystem() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetDpiForSystem.Addr(), 0,
		0, 0, 0)
	return uint32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagew
func GetMessage(msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax uint32) int32 {
	ret, _, lerr := syscall.Syscall6(proc.GetMessage.Addr(), 4,
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax),
		0, 0)
	if int(ret) == -1 {
		panic(NewWinError(co.ERROR(lerr), "GetMessage").Error())
	}
	return int32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-getopenfilenamew
func GetOpenFileName(ofn *OPENFILENAME) bool {
	ofn.LStructSize = uint32(unsafe.Sizeof(*ofn)) // safety
	ret, _, _ := syscall.Syscall(proc.GetOpenFileName.Addr(), 1,
		uintptr(unsafe.Pointer(ofn)), 0, 0)

	if ret == 0 {
		ret, _, _ := syscall.Syscall(proc.CommDlgExtendedError.Addr(), 0,
			0, 0, 0)
		if ret != 0 {
			panic(fmt.Sprintf("GetOpenFileName failed: %d.", ret))
		} else {
			return false // user cancelled
		}
	}
	return true // user clicked OK
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getphysicalcursorpos
func GetPhysicalCursorPos() *POINT {
	pt := &POINT{}
	ret, _, lerr := syscall.Syscall(proc.GetPhysicalCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(pt)), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "GetPhysicalCursorPos").Error())
	}
	return pt
}

// https://docs.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-getsavefilenamew
func GetSaveFileName(ofn *OPENFILENAME) bool {
	ofn.LStructSize = uint32(unsafe.Sizeof(*ofn)) // safety
	ret, _, _ := syscall.Syscall(proc.GetSaveFileName.Addr(), 1,
		uintptr(unsafe.Pointer(ofn)), 0, 0)

	if ret == 0 {
		ret, _, _ := syscall.Syscall(proc.CommDlgExtendedError.Addr(), 0,
			0, 0, 0)
		if ret != 0 {
			panic(fmt.Sprintf("GetSaveFileName failed: %d.", ret))
		} else {
			return false // user cancelled
		}
	}
	return true // user clicked OK
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
func GetSystemMetrics(index co.SM) int32 {
	ret, _, _ := syscall.Syscall(proc.GetSystemMetrics.Addr(), 1,
		uintptr(index), 0, 0)
	return int32(ret)
}

// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632656(v=vs.85)
func HiByte(value uint16) uint8 {
	return uint8(value >> 8 & 0xff)
}

// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632657(v=vs.85)
func HiWord(value uint32) uint16 {
	return uint16(value >> 16 & 0xffff)
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrols
func InitCommonControls() {
	syscall.Syscall(proc.InitCommonControls.Addr(), 0, 0, 0, 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isguithread
//
// Warning: passing true will force current thread to GUI, and it may deadlock.
func IsGUIThread(bConvertToGuiThread bool) bool {
	ret, _, _ := syscall.Syscall(proc.IsGUIThread.Addr(), 1,
		_Win.BoolToUintptr(bConvertToGuiThread), 0, 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsversionorgreater
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
			VerSetConditionMask(0, co.VER_MAJORVERSION, co.VER_COND_GREATER_EQUAL),
			co.VER_MINORVERSION, co.VER_COND_GREATER_EQUAL),
		co.VER_SERVICEPACKMAJOR, co.VER_COND_GREATER_EQUAL)

	ret, _ := VerifyVersionInfo(&ovi,
		co.VER_MAJORVERSION|co.VER_MINORVERSION|co.VER_SERVICEPACKMAJOR,
		conditionMask)
	return ret
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows10orgreater
func IsWindows10OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HiByte(uint16(co.WIN32_WINNT_WINTHRESHOLD))),
		uint32(LoByte(uint16(co.WIN32_WINNT_WINTHRESHOLD))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows7orgreater
func IsWindows7OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HiByte(uint16(co.WIN32_WINNT_WIN7))),
		uint32(LoByte(uint16(co.WIN32_WINNT_WIN7))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows8orgreater
func IsWindows8OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HiByte(uint16(co.WIN32_WINNT_WIN8))),
		uint32(LoByte(uint16(co.WIN32_WINNT_WIN8))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows8point1orgreater
func IsWindows8Point1OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HiByte(uint16(co.WIN32_WINNT_WINBLUE))),
		uint32(LoByte(uint16(co.WIN32_WINNT_WINBLUE))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsvistaorgreater
func IsWindowsVistaOrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HiByte(uint16(co.WIN32_WINNT_VISTA))),
		uint32(LoByte(uint16(co.WIN32_WINNT_VISTA))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsxporgreater
func IsWindowsXpOrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(HiByte(uint16(co.WIN32_WINNT_WINXP))),
		uint32(LoByte(uint16(co.WIN32_WINNT_WINXP))),
		0)
}

// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632658(v=vs.85)
func LoByte(value uint16) uint8 {
	return uint8(value & 0xFF)
}

// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632659(v=vs.85)
func LoWord(value uint32) uint16 {
	return uint16(value & 0xFFFF)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makelparam
func MakeLParam(lo, hi uint16) LPARAM {
	return LPARAM((uint32(lo) & 0xffff) | ((uint32(hi) & 0xffff) << 16))
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makewparam
func MakeWParam(lo, hi uint16) WPARAM {
	return WPARAM((uint32(lo) & 0xffff) | ((uint32(hi) & 0xffff) << 16))
}

// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-muldiv
func MulDiv(number, numerator, denominator int32) int32 {
	ret, _, _ := syscall.Syscall(proc.MulDiv.Addr(), 3,
		uintptr(number), uintptr(numerator), uintptr(denominator))
	return int32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int32) {
	syscall.Syscall(proc.PostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessage(
	idThread uint32, Msg co.WM, wParam WPARAM, lParam LPARAM) {

	ret, _, lerr := syscall.Syscall6(proc.PostThreadMessage.Addr(), 4,
		uintptr(idThread), uintptr(Msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "PostThreadMessage").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wcx *WNDCLASSEX) (ATOM, *WinError) {
	wcx.CbSize = uint32(unsafe.Sizeof(*wcx)) // safety
	ret, _, lerr := syscall.Syscall(proc.RegisterClassEx.Addr(), 1,
		uintptr(unsafe.Pointer(wcx)), 0, 0)
	if ret == 0 {
		return ATOM(0), NewWinError(co.ERROR(lerr), "RegisterClassEx")
	}
	return ATOM(ret), nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessage(lpString string) (uint32, *WinError) {
	ret, _, lerr := syscall.Syscall(proc.RegisterWindowMessage.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpString))), 0, 0)
	if ret == 0 {
		return 0, NewWinError(co.ERROR(lerr), "RegisterWindowMessage")
	}
	return uint32(ret), nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-replymessage
func ReplyMessage(lResult uintptr) bool {
	ret, _, _ := syscall.Syscall(proc.ReplyMessage.Addr(), 1,
		lResult, 0, 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiawarenesscontext
//
// Available in Windows 10, version 1703.
func SetProcessDpiAwarenessContext(value co.DPI_AWARE_CTX) {
	ret, _, lerr := syscall.Syscall(proc.SetProcessDpiAwarenessContext.Addr(), 1,
		uintptr(value), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SetProcessDpiAwarenessContext").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiaware
//
// Available in Windows Vista.
func SetProcessDPIAware() {
	ret, _, _ := syscall.Syscall(proc.SetProcessDPIAware.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic("SetProcessDPIAware failed.")
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowshookexw
func SetWindowsHookEx(idHook co.WH,
	lpfn func(code int32, wp WPARAM, lp LPARAM) uintptr,
	hmod HINSTANCE, dwThreadId uint32) HHOOK {

	ret, _, lerr := syscall.Syscall6(proc.SetWindowsHookEx.Addr(), 4,
		uintptr(idHook), syscall.NewCallback(lpfn),
		uintptr(hmod), uintptr(dwThreadId), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SetWindowsHookEx").Error())
	}
	return HHOOK(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
//
// Depends of CoInitializeEx().
func SHGetFileInfo(pszPath string, dwFileAttributes co.FILE_ATTRIBUTE,
	uFlags co.SHGFI) *SHFILEINFO {

	shfi := &SHFILEINFO{}
	ret, _, _ := syscall.Syscall6(proc.SHGetFileInfo.Addr(), 5,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(pszPath))),
		uintptr(dwFileAttributes), uintptr(unsafe.Pointer(shfi)),
		unsafe.Sizeof(*shfi), uintptr(uFlags), 0)

	if (uFlags&co.SHGFI_EXETYPE) == 0 || (uFlags&co.SHGFI_SYSICONINDEX) == 0 {
		if ret == 0 {
			panic(NewWinError(co.ERROR_E_UNEXPECTED, "SHGetFileInfo").Error())
		}
	}

	if (uFlags & co.SHGFI_EXETYPE) != 0 {
		if ret == 0 {
			panic("SHGetFileInfo failed.")
		}
	}

	return shfi
}

// https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-sleep
func Sleep(dwMilliseconds uint32) {
	syscall.Syscall(proc.Sleep.Addr(), 1,
		uintptr(dwMilliseconds), 0, 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
func SystemParametersInfo(uiAction co.SPI, uiParam uint32,
	pvParam unsafe.Pointer, fWinIni uint32) {

	ret, _, lerr := syscall.Syscall6(proc.SystemParametersInfo.Addr(), 4,
		uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni),
		0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SystemParametersInfo").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetofiletime
func SystemTimeToFileTime(inSystemTime *SYSTEMTIME, outFileTime *FILETIME) {
	ret, _, lerr := syscall.Syscall(proc.SystemTimeToFileTime.Addr(), 2,
		uintptr(unsafe.Pointer(inSystemTime)),
		uintptr(unsafe.Pointer(outFileTime)), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SystemTimeToFileTime").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetotzspecificlocaltime
func SystemTimeToTzSpecificLocalTime(
	lpTimeZoneInformation *TIME_ZONE_INFORMATION,
	inUniversalTime *SYSTEMTIME, outLocalTime *SYSTEMTIME) {

	ret, _, lerr := syscall.Syscall(proc.SystemTimeToTzSpecificLocalTime.Addr(), 3,
		uintptr(unsafe.Pointer(lpTimeZoneInformation)),
		uintptr(unsafe.Pointer(inUniversalTime)),
		uintptr(unsafe.Pointer(outLocalTime)))
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SystemTimeToTzSpecificLocalTime").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(proc.TranslateMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-tzspecificlocaltimetosystemtime
func TzSpecificLocalTimeToSystemTime(
	lpTimeZoneInformation *TIME_ZONE_INFORMATION,
	inLocalTime *SYSTEMTIME, outUniversalTime *SYSTEMTIME) {

	ret, _, lerr := syscall.Syscall(proc.TzSpecificLocalTimeToSystemTime.Addr(), 3,
		uintptr(unsafe.Pointer(lpTimeZoneInformation)),
		uintptr(unsafe.Pointer(inLocalTime)),
		uintptr(unsafe.Pointer(outUniversalTime)))
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "TzSpecificLocalTimeToSystemTime").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-verifyversioninfow
func VerifyVersionInfo(ovi *OSVERSIONINFOEX, typeMask co.VER,
	conditionMask uint64) (bool, co.ERROR) {

	ret, _, lerr := syscall.Syscall(proc.VerifyVersionInfo.Addr(), 3,
		uintptr(unsafe.Pointer(ovi)),
		uintptr(typeMask), uintptr(conditionMask))
	return ret != 0, co.ERROR(lerr)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-versetconditionmask
func VerSetConditionMask(conditionMask uint64, typeMask co.VER,
	condition co.VER_COND) uint64 {

	ret, _, _ := syscall.Syscall(proc.VerSetConditionMask.Addr(), 3,
		uintptr(conditionMask), uintptr(typeMask), uintptr(condition))
	return uint64(ret)
}
