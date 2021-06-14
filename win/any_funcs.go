package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-adjustwindowrectex
func AdjustWindowRectEx(
	lpRect *RECT, dwStyle co.WS, bMenu bool, dwExStyle co.WS_EX) {

	ret, _, lerr := syscall.Syscall6(proc.AdjustWindowRectEx.Addr(), 4,
		uintptr(unsafe.Pointer(lpRect)), uintptr(dwStyle),
		util.BoolToUintptr(bMenu), uintptr(dwExStyle), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-allowsetforegroundwindow
func AllowSetForegroundWindow(dwProcessId uint32) {
	ret, _, lerr := syscall.Syscall(proc.AllowSetForegroundWindow.Addr(), 1,
		uintptr(dwProcessId), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-begindeferwindowpos
func BeginDeferWindowPos(numWindows int32) HDWP {
	ret, _, lerr := syscall.Syscall(proc.BeginDeferWindowPos.Addr(), 1,
		uintptr(numWindows), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HDWP(ret)
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
	hr := err.ERROR(ret)
	if hr != err.S_OK && hr != err.S_FALSE {
		panic(hr)
	}
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

// ⚠️ You must defer DestroyAcceleratorTable().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createacceleratortablew
func CreateAcceleratorTable(accelList []ACCEL) HACCEL {
	ret, _, lerr := syscall.Syscall(proc.CreateAcceleratorTable.Addr(), 2,
		uintptr(unsafe.Pointer(&accelList[0])), uintptr(len(accelList)),
		0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HACCEL(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createdirectoryw
func CreateDirectory(
	pathName string, securityAttributes *SECURITY_ATTRIBUTES) error {

	ret, _, lerr := syscall.Syscall(proc.CreateDirectory.Addr(), 2,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(pathName))),
		uintptr(unsafe.Pointer(securityAttributes)), 0)
	if ret == 0 {
		return err.ERROR(lerr)
	}
	return nil
}

// ⚠️ You must defer CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
func CreateFile(fileName string, desiredAccess co.GENERIC,
	shareMode co.FILE_SHARE, securityAttributes *SECURITY_ATTRIBUTES,
	creationDisposition co.DISPOSITION, attributes co.FILE_ATTRIBUTE,
	flags co.FILE_FLAG, security co.SECURITY,
	hTemplateFile HFILE) (HFILE, error) {

	ret, _, lerr := syscall.Syscall9(proc.CreateFile.Addr(), 7,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(fileName))),
		uintptr(desiredAccess), uintptr(shareMode),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(creationDisposition),
		uintptr(uint32(attributes)|uint32(flags)|uint32(security)),
		uintptr(hTemplateFile), 0, 0)

	if int(ret) == _INVALID_HANDLE_VALUE {
		return HFILE(0), err.ERROR(lerr)
	}
	return HFILE(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createfontindirectw
func CreateFontIndirect(lf *LOGFONT) HFONT {
	ret, _, lerr := syscall.Syscall(proc.CreateFontIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lf)), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HFONT(ret)
}

// ⚠️ You must defer DestroyMenu().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createmenu
func CreateMenu() HMENU {
	ret, _, lerr := syscall.Syscall(proc.CreateMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HMENU(ret)
}

// ⚠️ You must defer DestroyMenu().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createpopupmenu
func CreatePopupMenu() HMENU {
	ret, _, lerr := syscall.Syscall(proc.CreatePopupMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HMENU(ret)
}

// ⚠️ You must defer DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createrectrgnindirect
func CreateRectRgnIndirect(lprect *RECT) HRGN {
	ret, _, lerr := syscall.Syscall(proc.CreateRectRgnIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lprect)), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HRGN(ret)
}

// ⚠️ You must defer DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createroundrectrgn
func CreateRoundRectRgn(x1, y1, x2, y2, w, h int32) HRGN {
	ret, _, lerr := syscall.Syscall6(proc.CreateRoundRectRgn.Addr(), 6,
		uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2),
		uintptr(w), uintptr(h))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HRGN(ret)
}

// ⚠️ You must defer DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createsolidbrush
func CreateSolidBrush(color COLORREF) HBRUSH {
	ret, _, lerr := syscall.Syscall(proc.CreateSolidBrush.Addr(), 1,
		uintptr(color), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HBRUSH(ret)
}

// Not an actual Win32 function, just a tricky conversion to create a brush from
// a system color, particularly used when registering a window class.
func CreateSysColorBrush(sysColor co.COLOR) HBRUSH {
	return HBRUSH(sysColor + 1)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func CreateWindowEx(exStyle co.WS_EX, className, title string, style co.WS,
	x, y, width, height int32, parent HWND, menu HMENU,
	instance HINSTANCE, param LPARAM) HWND {

	ret, _, lerr := syscall.Syscall12(proc.CreateWindowEx.Addr(), 12,
		uintptr(exStyle),
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(className))),
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(title))),
		uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(parent), uintptr(menu), uintptr(instance), uintptr(param))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-deletefilew
func DeleteFile(fileName string) error {
	ret, _, lerr := syscall.Syscall(proc.DeleteFile.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(fileName))), 0, 0)
	if ret == 0 {
		return err.ERROR(lerr)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycaret
func DestroyCaret() {
	ret, _, lerr := syscall.Syscall(proc.DestroyCaret.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.Syscall(proc.DispatchMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-emptyclipboard
func EmptyClipboard() {
	ret, _, lerr := syscall.Syscall(proc.EmptyClipboard.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endmenu
func EndMenu() {
	ret, _, lerr := syscall.Syscall(proc.EndMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumwindows
func EnumWindows(
	lpEnumFunc func(hwnd HWND, lParam LPARAM) bool,
	lParam LPARAM) {

	ret, _, lerr := syscall.Syscall(proc.EnumWindows.Addr(), 2,
		syscall.NewCallback(
			func(hwnd HWND, lParam LPARAM) uintptr {
				return util.BoolToUintptr(lpEnumFunc(hwnd, lParam))
			}),
		uintptr(lParam), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// Extracts all icons: big and small.
//
// ⚠️ You must defer DestroyIcon() on each icon returned in both slices.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-extracticonexw
func ExtractIconEx(lpszFile string) ([]HICON, []HICON) {
	retrieveIdx := -1
	ret, _, lerr := syscall.Syscall6(proc.ExtractIconEx.Addr(), 5,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpszFile))),
		uintptr(retrieveIdx), 0, 0, 0, 0)
	if ret == _UINT_MAX {
		panic(err.ERROR(lerr))
	}

	numIcons := int(ret)
	largeIcons := make([]HICON, numIcons)
	smallIcons := make([]HICON, numIcons)

	ret, _, lerr = syscall.Syscall6(proc.ExtractIconEx.Addr(), 5,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpszFile))), 0,
		uintptr(unsafe.Pointer(&largeIcons[0])),
		uintptr(unsafe.Pointer(&smallIcons[0])), uintptr(numIcons), 0)
	if ret == _UINT_MAX {
		panic(err.ERROR(lerr))
	}

	return largeIcons, smallIcons
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-filetimetosystemtime
func FileTimeToSystemTime(inFileTime *FILETIME, outSystemTime *SYSTEMTIME) {
	ret, _, lerr := syscall.Syscall(proc.FileTimeToSystemTime.Addr(), 2,
		uintptr(unsafe.Pointer(inFileTime)),
		uintptr(unsafe.Pointer(outSystemTime)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// Returns true if a file was found.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
func FindFirstFile(lpFileName string,
	lpFindFileData *WIN32_FIND_DATA) (HFIND, bool, error) {

	ret, _, lerr := syscall.Syscall(proc.FindFirstFile.Addr(), 2,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpFileName))),
		uintptr(unsafe.Pointer(lpFindFileData)), 0)

	errCode := err.ERROR(lerr)
	if int(ret) == _INVALID_HANDLE_VALUE {
		if errCode == err.FILE_NOT_FOUND || errCode == err.PATH_NOT_FOUND { // no matching files, not an error
			return HFIND(0), false, nil
		} else {
			return HFIND(0), false, errCode
		}
	}
	return HFIND(ret), true, nil // a file was found
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
	ret, _, lerr := syscall.Syscall(proc.GetCaretPos.Addr(), 1,
		uintptr(unsafe.Pointer(&rc)), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return rc
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
	ret, _, lerr := syscall.Syscall(proc.GetCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(&pt)), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdesktopwindow
func GetDesktopWindow() HWND {
	ret, _, _ := syscall.Syscall(proc.GetDesktopWindow.Addr(), 0, 0, 0, 0)
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-getdynamictimezoneinformation
func GetDynamicTimeZoneInformation(
	pTimeZoneInformation *DYNAMIC_TIME_ZONE_INFORMATION) co.TIME_ZONE_ID {

	ret, _, _ := syscall.Syscall(proc.GetDynamicTimeZoneInformation.Addr(), 1,
		uintptr(unsafe.Pointer(pTimeZoneInformation)), 0, 0)
	return co.TIME_ZONE_ID(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfileattributesw
func GetFileAttributes(lpFileName string) (co.FILE_ATTRIBUTE, error) {
	ret, _, lerr := syscall.Syscall(proc.GetFileAttributes.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpFileName))), 0, 0)

	retAttr := co.FILE_ATTRIBUTE(ret)
	if retAttr == co.FILE_ATTRIBUTE_INVALID {
		return retAttr, err.ERROR(lerr)
	}
	return retAttr, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getfocus
func GetFocus() HWND {
	ret, _, _ := syscall.Syscall(proc.GetFocus.Addr(), 0, 0, 0, 0)
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
func GetForegroundWindow() HWND {
	ret, _, _ := syscall.Syscall(proc.GetForegroundWindow.Addr(), 0,
		0, 0, 0)
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagew
func GetMessage(
	msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax uint32) (int32, error) {

	ret, _, lerr := syscall.Syscall6(proc.GetMessage.Addr(), 4,
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax),
		0, 0)
	if int(ret) == -1 {
		return 0, err.ERROR(lerr)
	}
	return int32(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandle(moduleName string) HINSTANCE {
	ret, _, lerr := syscall.Syscall(proc.GetModuleHandle.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(moduleName))),
		0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HINSTANCE(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getphysicalcursorpos
func GetPhysicalCursorPos() POINT {
	pt := POINT{}
	ret, _, lerr := syscall.Syscall(proc.GetPhysicalCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(&pt)), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
func GetSysColor(nIndex co.COLOR) COLORREF {
	ret, _, _ := syscall.Syscall(proc.GetSysColor.Addr(), 1,
		uintptr(nIndex), 0, 0)
	return COLORREF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolorbrush
func GetSysColorBrush(nIndex co.COLOR) HBRUSH {
	ret, _, lerr := syscall.Syscall(proc.GetSysColorBrush.Addr(), 1,
		uintptr(nIndex), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HBRUSH(ret)
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

	ret, _, lerr := syscall.Syscall(proc.GetTimeZoneInformationForYear.Addr(), 3,
		uintptr(wYear),
		uintptr(unsafe.Pointer(pdtzi)), uintptr(unsafe.Pointer(ptzi)))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632656(v=vs.85)
func HIBYTE(value uint16) uint8 {
	return uint8(value >> 8 & 0xff)
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632657(v=vs.85)
func HIWORD(value uint32) uint16 {
	return uint16(value >> 16 & 0xffff)
}

// Usually flags is ILC_COLOR32.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_create
func ImageListCreate(cx, cy uint32, flags co.ILC,
	cInitial, cGrow uint32) HIMAGELIST {

	ret, _, lerr := syscall.Syscall6(proc.ImageList_Create.Addr(), 5,
		uintptr(cx), uintptr(cy), uintptr(flags),
		uintptr(cInitial), uintptr(cGrow), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HIMAGELIST(ret)
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isguithread
func IsGUIThread(bConvertToGuiThread bool) (bool, error) {
	ret, _, _ := syscall.Syscall(proc.IsGUIThread.Addr(), 1,
		util.BoolToUintptr(bConvertToGuiThread), 0, 0)
	if bConvertToGuiThread && err.ERROR(ret) == err.NOT_ENOUGH_MEMORY {
		return false, err.NOT_ENOUGH_MEMORY
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

// ⚠️ You must defer FreeLibrary().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadlibraryw
func LoadLibrary(lpLibFileName string) HINSTANCE {
	ret, _, lerr := syscall.Syscall(proc.LoadLibrary.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpLibFileName))),
		0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HINSTANCE(ret)
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632658(v=vs.85)
func LOBYTE(value uint16) uint8 {
	return uint8(value & 0xff)
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632659(v=vs.85)
func LOWORD(value uint32) uint16 {
	return uint16(value & 0xffff)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makelparam
func MAKELPARAM(lo, hi uint16) LPARAM {
	return LPARAM((uint32(lo) & 0xffff) | ((uint32(hi) & 0xffff) << 16))
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms632663(v=vs.85)
func MAKEWORD(lo, hi uint8) uint16 {
	return (uint16(lo) & 0xff) | ((uint16(hi) & 0xff) << 8)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makewparam
func MAKEWPARAM(lo, hi uint16) WPARAM {
	return WPARAM((uint32(lo) & 0xffff) | ((uint32(hi) & 0xffff) << 16))
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int32) {
	syscall.Syscall(proc.PostQuitMessage.Addr(), 1,
		uintptr(exitCode), 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessage(
	idThread uint32, Msg co.WM, wParam WPARAM, lParam LPARAM) {

	ret, _, lerr := syscall.Syscall6(proc.PostThreadMessage.Addr(), 4,
		uintptr(idThread), uintptr(Msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/profileapi/nf-profileapi-queryperformancecounter
func QueryPerformanceCounter() int64 {
	lpPerformanceCount := int64(0)
	ret, _, lerr := syscall.Syscall(proc.QueryPerformanceCounter.Addr(), 1,
		uintptr(unsafe.Pointer(&lpPerformanceCount)), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return lpPerformanceCount
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/profileapi/nf-profileapi-queryperformancefrequency
func QueryPerformanceFrequency() int64 {
	lpFrequency := int64(0)
	ret, _, lerr := syscall.Syscall(proc.QueryPerformanceFrequency.Addr(), 1,
		uintptr(unsafe.Pointer(&lpFrequency)), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return lpFrequency
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wcx *WNDCLASSEX) (ATOM, error) {
	wcx.CbSize = uint32(unsafe.Sizeof(*wcx)) // safety
	ret, _, lerr := syscall.Syscall(proc.RegisterClassEx.Addr(), 1,
		uintptr(unsafe.Pointer(wcx)), 0, 0)
	if ret == 0 {
		return ATOM(0), err.ERROR(lerr)
	}
	return ATOM(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rgb
func RGB(r, g, b uint8) COLORREF {
	return COLORREF(uint32(r) | (uint32(g) << 8) | (uint32(b) << 16))
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
	ret, _, lerr := syscall.Syscall(proc.SetProcessDpiAwarenessContext.Addr(), 1,
		uintptr(value), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowshookexw
func SetWindowsHookEx(idHook co.WH,
	lpfn func(code int32, wp WPARAM, lp LPARAM) uintptr,
	hmod HINSTANCE, dwThreadId uint32) HHOOK {

	ret, _, lerr := syscall.Syscall6(proc.SetWindowsHookEx.Addr(), 4,
		uintptr(idHook), syscall.NewCallback(lpfn),
		uintptr(hmod), uintptr(dwThreadId), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HHOOK(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw
func ShellNotifyIcon(dwMessage co.NIM, lpData *NOTIFYICONDATA) {
	ret, _, lerr := syscall.Syscall(proc.Shell_NotifyIcon.Addr(), 2,
		uintptr(dwMessage), uintptr(unsafe.Pointer(lpData)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// Depends of CoInitializeEx().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
func SHGetFileInfo(
	pszPath string, dwFileAttributes co.FILE_ATTRIBUTE,
	uFlags co.SHGFI) *SHFILEINFO {

	shfi := SHFILEINFO{}
	ret, _, lerr := syscall.Syscall6(proc.SHGetFileInfo.Addr(), 5,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(pszPath))),
		uintptr(dwFileAttributes), uintptr(unsafe.Pointer(&shfi)),
		unsafe.Sizeof(shfi), uintptr(uFlags), 0)

	if (uFlags&co.SHGFI_EXETYPE) == 0 || (uFlags&co.SHGFI_SYSICONINDEX) == 0 {
		if ret == 0 {
			panic(err.ERROR(lerr))
		}
	}

	if (uFlags & co.SHGFI_EXETYPE) != 0 {
		if ret == 0 {
			panic(err.ERROR(lerr))
		}
	}

	return &shfi
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-sleep
func Sleep(dwMilliseconds uint32) {
	syscall.Syscall(proc.Sleep.Addr(), 1,
		uintptr(dwMilliseconds), 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
func SystemParametersInfo(
	uiAction co.SPI, uiParam uint32, pvParam unsafe.Pointer, fWinIni co.SPIF) {

	ret, _, lerr := syscall.Syscall6(proc.SystemParametersInfo.Addr(), 4,
		uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni),
		0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetofiletime
func SystemTimeToFileTime(inSystemTime *SYSTEMTIME, outFileTime *FILETIME) {
	ret, _, lerr := syscall.Syscall(proc.SystemTimeToFileTime.Addr(), 2,
		uintptr(unsafe.Pointer(inSystemTime)),
		uintptr(unsafe.Pointer(outFileTime)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetotzspecificlocaltime
func SystemTimeToTzSpecificLocalTime(
	lpTimeZoneInformation *TIME_ZONE_INFORMATION,
	inUniversalTime *SYSTEMTIME, outLocalTime *SYSTEMTIME) {

	ret, _, lerr := syscall.Syscall(proc.SystemTimeToTzSpecificLocalTime.Addr(), 3,
		uintptr(unsafe.Pointer(lpTimeZoneInformation)),
		uintptr(unsafe.Pointer(inUniversalTime)),
		uintptr(unsafe.Pointer(outLocalTime)))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
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

	ret, _, lerr := syscall.Syscall(proc.TzSpecificLocalTimeToSystemTime.Addr(), 3,
		uintptr(unsafe.Pointer(lpTimeZoneInformation)),
		uintptr(unsafe.Pointer(inLocalTime)),
		uintptr(unsafe.Pointer(outUniversalTime)))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-verifyversioninfow
func VerifyVersionInfo(
	ovi *OSVERSIONINFOEX,
	typeMask co.VER, conditionMask uint64) (bool, err.ERROR) {

	ret, _, lerr := syscall.Syscall(proc.VerifyVersionInfo.Addr(), 3,
		uintptr(unsafe.Pointer(ovi)),
		uintptr(typeMask), uintptr(conditionMask))
	return ret != 0, err.ERROR(lerr)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-versetconditionmask
func VerSetConditionMask(
	conditionMask uint64, typeMask co.VER, condition co.VER_COND) uint64 {

	ret, _, _ := syscall.Syscall(proc.VerSetConditionMask.Addr(), 3,
		uintptr(conditionMask), uintptr(typeMask), uintptr(condition))
	return uint64(ret)
}
