//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [AdjustWindowRectEx] function.
//
// [AdjustWindowRectEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-adjustwindowrectex
func AdjustWindowRectEx(rc *RECT, style co.WS, hasMenu bool, exStyle co.WS_EX) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_AdjustWindowRectEx, "AdjustWindowRectEx"),
		uintptr(unsafe.Pointer(rc)),
		uintptr(style),
		utl.BoolToUintptr(hasMenu),
		uintptr(exStyle))
	return utl.ZeroAsGetLastError(ret, err)
}

var _AdjustWindowRectEx *syscall.Proc

// [AllowSetForegroundWindow] function.
//
// [AllowSetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-allowsetforegroundwindow
func AllowSetForegroundWindow(processId uint32) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_AllowSetForegroundWindow, "AllowSetForegroundWindow"),
		uintptr(processId))
	return utl.ZeroAsGetLastError(ret, err)
}

var _AllowSetForegroundWindow *syscall.Proc

// [AnyPopup] function.
//
// [AnyPopup]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-anypopup
func AnyPopup() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_AnyPopup, "AnyPopup"))
	return ret != 0
}

var _AnyPopup *syscall.Proc

// [BroadcastSystemMessage] function.
//
// [BroadcastSystemMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-broadcastsystemmessagew
func BroadcastSystemMessage(
	flags co.BSF,
	recipients co.BSM,
	msg co.WM,
	wParam WPARAM,
	lParam LPARAM,
) (broadcastSuccessful bool, receivers co.BSM, wErr error) {
	receivers = recipients
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_BroadcastSystemMessageW, "BroadcastSystemMessageW"),
		uintptr(flags),
		uintptr(unsafe.Pointer(&receivers)),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))

	broadcastSuccessful = int(ret) > 1
	if ret == 0 {
		wErr = co.ERROR(err)
	}
	return
}

var _BroadcastSystemMessageW *syscall.Proc

// [CreateIconFromResourceEx] function for cursor.
//
// This function creates [HCURSOR] only. The [HICON] variation is
// [CreateIconFromResourceEx].
//
// ⚠️ You must defer [HCURSOR.DestroyCursor].
//
// [CreateIconFromResourceEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconfromresourceex
func CreateCursorFromResourceEx(
	resBits []byte,
	fmtVersion uint32,
	cxDesired, cyDesired uint,
	flags co.LR,
) (HCURSOR, error) {
	hIcon, err := CreateIconFromResourceEx(resBits, fmtVersion, cxDesired, cyDesired, flags)
	return HCURSOR(hIcon), err
}

// [CreateIconFromResourceEx] function.
//
// This function creates [HICON] only. The [HCURSOR] variation is
// [CreateCursorFromResourceEx].
//
// ⚠️ You must defer [HICON.DestroyIcon].
//
// [CreateIconFromResourceEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconfromresourceex
func CreateIconFromResourceEx(
	resBits []byte,
	fmtVersion uint32,
	cxDesired, cyDesired uint,
	flags co.LR,
) (HICON, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_CreateIconFromResourceEx, "CreateIconFromResourceEx"),
		uintptr(unsafe.Pointer(&resBits[0])),
		uintptr(uint32(len(resBits))),
		1,
		uintptr(fmtVersion),
		uintptr(int32(cxDesired)),
		uintptr(int32(cyDesired)),
		uintptr(flags))
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

var _CreateIconFromResourceEx *syscall.Proc

// [DestroyCaret] function.
//
// [DestroyCaret]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycarret
func DestroyCaret() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_DestroyCaret, "DestroyCaret"))
	return utl.ZeroAsGetLastError(ret, err)
}

var _DestroyCaret *syscall.Proc

// [DispatchMessage] function.
//
// [DispatchMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_DispatchMessageW, "DispatchMessageW"),
		uintptr(unsafe.Pointer(msg)))
	return ret
}

var _DispatchMessageW *syscall.Proc

// [EndMenu] function.
//
// [EndMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endmenu
func EndMenu() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_EndMenu, "EndMenu"))
	return utl.ZeroAsGetLastError(ret, err)
}

var _EndMenu *syscall.Proc

// [EnumDisplayDevices] function.
//
// # Example
//
//	devices := win.EnumDisplayDevices("", co.EDD_GET_DEVICE_INTERFACE_NAME)
//	for _, device := devices {
//		println(device.DeviceName(), device.DeviceString())
//	}
//
// [EnumDisplayDevices]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaydevicesw
func EnumDisplayDevices(device string, flags co.EDD) []DISPLAY_DEVICE {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pDevice := wbuf.PtrEmptyIsNil(device)

	devices := make([]DISPLAY_DEVICE, 0) // to be returned
	devNum := 0

	var dide DISPLAY_DEVICE // buffer to receive each iteration
	dide.SetCb()

	for {
		// Ignore errors: only fails with devNum out-of-bounds, which never happens here.
		ret, _, _ := syscall.SyscallN(
			dll.Load(dll.USER32, &_EnumDisplayDevicesW, "EnumDisplayDevicesW"),
			uintptr(pDevice),
			uintptr(devNum),
			uintptr(unsafe.Pointer(&dide)),
			uintptr(flags))
		if ret == 0 {
			break
		}
		devices = append(devices, dide)
		devNum++
	}

	return devices
}

var _EnumDisplayDevicesW *syscall.Proc

// [EnumThreadWindows] function.
//
// [EnumThreadWindows]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumthreadwindows
func EnumThreadWindows(threadId uint32) []HWND {
	pPack := &_EnumThreadWindowsPack{
		arr: make([]HWND, 0),
	}

	syscall.SyscallN(
		dll.Load(dll.USER32, &_EnumThreadWindows, "EnumThreadWindows"),
		enumThreadWindowsCallback(),
		uintptr(unsafe.Pointer(pPack)))
	return pPack.arr
}

var _EnumThreadWindows *syscall.Proc

type _EnumThreadWindowsPack struct{ arr []HWND }

func enumThreadWindowsCallback() uintptr {
	if _enumThreadWindowsCallback != 0 {
		return _enumThreadWindowsCallback
	}

	_enumThreadWindowsCallback = syscall.NewCallback(
		func(hWnd HWND, lParam LPARAM) uintptr {
			pPack := (*_EnumThreadWindowsPack)(unsafe.Pointer(lParam))
			pPack.arr = append(pPack.arr, hWnd)
			return 1
		},
	)
	return _enumThreadWindowsCallback
}

var _enumThreadWindowsCallback uintptr

// [EnumWindows] function.
//
// [EnumWindows]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumwindows
func EnumWindows() []HWND {
	pPack := &_EnumWindowsPack{
		arr: make([]HWND, 0),
	}

	syscall.SyscallN(
		dll.Load(dll.USER32, &_EnumWindows, "EnumWindows"),
		enumWindowsCallback(),
		uintptr(unsafe.Pointer(pPack)))
	return pPack.arr
}

var _EnumWindows *syscall.Proc

type _EnumWindowsPack struct{ arr []HWND }

var _enumWindowsCallback uintptr

func enumWindowsCallback() uintptr {
	if _enumWindowsCallback != 0 {
		return _enumWindowsCallback
	}

	_enumWindowsCallback = syscall.NewCallback(
		func(hWnd HWND, lParam LPARAM) uintptr {
			pPack := (*_EnumWindowsPack)(unsafe.Pointer(lParam))
			pPack.arr = append(pPack.arr, hWnd)
			return 1
		},
	)
	return _enumWindowsCallback
}

// [ExitWindowsEx] function.
//
// [ExitWindowsEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-exitwindowsex
func ExitWindowsEx(flags co.EXW, reason co.SHTDN) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_ExitWindowsEx, "ExitWindowsEx"),
		uintptr(flags),
		uintptr(reason))
	return utl.ZeroAsGetLastError(ret, err)
}

var _ExitWindowsEx *syscall.Proc

// [GetAsyncKeyState] function.
//
// [GetAsyncKeyState]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getasynckeystate
func GetAsyncKeyState(virtKeyCode co.VK) uint16 {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetAsyncKeyState, "GetAsyncKeyState"),
		uintptr(virtKeyCode))
	return uint16(ret)
}

var _GetAsyncKeyState *syscall.Proc

// [GetCaretPos] function.
//
// [GetCaretPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcaretpos
func GetCaretPos() (RECT, error) {
	var rc RECT
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetCaretPos, "GetCaretPos"),
		uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		return RECT{}, co.ERROR(err)
	}
	return rc, nil
}

var _GetCaretPos *syscall.Proc

// [GetCursorInfo] function.
//
// [GetCursorInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorinfo
func GetCursorInfo() (CURSORINFO, error) {
	var ci CURSORINFO
	ci.SetCbSize()

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetCursorInfo, "GetCursorInfo"),
		uintptr(unsafe.Pointer(&ci)))
	if ret == 0 {
		return CURSORINFO{}, co.ERROR(err)
	}
	return ci, nil
}

var _GetCursorInfo *syscall.Proc

// [GetCursorPos] function.
//
// [GetCursorPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
func GetCursorPos() (POINT, error) {
	var pt POINT
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetCursorPos, "GetCursorPos"),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR(err)
	}
	return pt, nil
}

var _GetCursorPos *syscall.Proc

// [GetDialogBaseUnits] function.
//
// [GetDialogBaseUnits]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdialogbaseunits
func GetDialogBaseUnits() (horz, vert uint16) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetDialogBaseUnits, "GetDialogBaseUnits"))
	horz, vert = LOWORD(uint32(ret)), HIWORD(uint32(ret))
	return
}

var _GetDialogBaseUnits *syscall.Proc

// [GetGUIThreadInfo] function.
//
// [GetGUIThreadInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getguithreadinfo
func GetGUIThreadInfo(thread_id uint32) (GUITHREADINFO, error) {
	var info GUITHREADINFO
	info.SetCbSize()

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetGUIThreadInfo, "GetGUIThreadInfo"),
		uintptr(thread_id),
		uintptr(unsafe.Pointer(&info)))
	if ret == 0 {
		return GUITHREADINFO{}, co.ERROR(err)
	}
	return info, nil
}

var _GetGUIThreadInfo *syscall.Proc

// [GetInputState] function.
//
// [GetInputState]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getinputstate
func GetInputState() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetInputState, "GetInputState"))
	return ret != 0
}

var _GetInputState *syscall.Proc

// [GetMessage] function.
//
// [GetMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagew
func GetMessage(msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax uint32) (int32, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetMessageW, "GetMessageW"),
		uintptr(unsafe.Pointer(msg)),
		uintptr(hWnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))
	if int32(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return int32(ret), nil
}

var _GetMessageW *syscall.Proc

// [GetMessageExtraInfo] function.
//
// [GetMessageExtraInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessageextrainfo
func GetMessageExtraInfo() LPARAM {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetMessageExtraInfo, "GetMessageExtraInfo"))
	return LPARAM(ret)
}

var _GetMessageExtraInfo *syscall.Proc

// [GetMessagePos] function.
//
// [GetMessagePos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagepos
func GetMessagePos() POINT {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetMessagePos, "GetMessagePos"))
	return POINT{
		X: int32(LOWORD(uint32(ret))),
		Y: int32(HIWORD(uint32(ret))),
	}
}

var _GetMessagePos *syscall.Proc

// [GetMessageTime] function.
//
// [GetMessageTime]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagetime
func GetMessageTime() time.Duration {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetMessageTime, "GetMessageTime"))
	return time.Duration(ret * uintptr(time.Millisecond))
}

var _GetMessageTime *syscall.Proc

// [GetPhysicalCursorPos] function.
//
// [GetPhysicalCursorPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getphysicalcursorpos
func GetPhysicalCursorPos() (POINT, error) {
	var pt POINT
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetPhysicalCursorPos, "GetPhysicalCursorPos"),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR(err)
	}
	return pt, nil
}

var _GetPhysicalCursorPos *syscall.Proc

// [GetProcessDefaultLayout] function.
//
// [GetProcessDefaultLayout]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getprocessdefaultlayout
func GetProcessDefaultLayout() (co.LAYOUT, error) {
	var defaultLayout co.LAYOUT
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetProcessDefaultLayout, "GetProcessDefaultLayout"),
		uintptr(unsafe.Pointer(&defaultLayout)))
	if ret == 0 {
		return co.LAYOUT(0), co.ERROR(err)
	}
	return defaultLayout, nil
}

var _GetProcessDefaultLayout *syscall.Proc

// [GetQueueStatus] function.
//
// [GetQueueStatus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getqueuestatus
func GetQueueStatus(flags co.QS) (currentlyInQueue, addedToQueue co.QS) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetQueueStatus, "GetQueueStatus"),
		uintptr(flags))
	currentlyInQueue = co.QS(HIWORD(uint32(ret)))
	addedToQueue = co.QS(LOWORD(uint32(ret)))
	return
}

var _GetQueueStatus *syscall.Proc

// [GetSysColor] function.
//
// [GetSysColor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
func GetSysColor(index co.COLOR) COLORREF {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetSysColor, "GetSysColor"),
		uintptr(index))
	return COLORREF(ret)
}

var _GetSysColor *syscall.Proc

// [GetSystemMetrics] function.
//
// [GetSystemMetrics]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
func GetSystemMetrics(index co.SM) int32 {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetSystemMetrics, "GetSystemMetrics"),
		uintptr(index))
	return int32(ret)
}

var _GetSystemMetrics *syscall.Proc

// [InflateRect] function.
//
// [InflateRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-inflaterect
func InflateRect(rc *RECT, dx, dy int) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_InflateRect, "InflateRect"),
		uintptr(unsafe.Pointer(rc)),
		uintptr(dx),
		uintptr(dy))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _InflateRect *syscall.Proc

// [InSendMessage] function.
//
// [InSendMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessage
func InSendMessage() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_InSendMessage, "InSendMessage"))
	return ret != 0
}

var _InSendMessage *syscall.Proc

// [InSendMessageEx] function.
//
// [InSendMessageEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessageex
func InSendMessageEx() co.ISMEX {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_InSendMessageEx, "InSendMessageEx"))
	return co.ISMEX(ret)
}

var _InSendMessageEx *syscall.Proc

// [IsGUIThread] function.
//
// [IsGUIThread]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isguithread
func IsGUIThread(convertToGuiThread bool) (bool, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_IsGUIThread, "IsGUIThread"),
		utl.BoolToUintptr(convertToGuiThread))
	if convertToGuiThread && co.ERROR(ret) == co.ERROR_NOT_ENOUGH_MEMORY {
		return false, co.ERROR_NOT_ENOUGH_MEMORY
	}
	return ret != 0, nil
}

var _IsGUIThread *syscall.Proc

// [IsProcessDPIAware] function.
//
// [IsProcessDPIAware]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isprocessdpiaware
func IsProcessDPIAware() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_IsProcessDPIAware, "IsProcessDPIAware"))
	return ret != 0
}

var _IsProcessDPIAware *syscall.Proc

// [LockSetForegroundWindow] function.
//
// [LockSetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-locksetforegroundwindow
func LockSetForegroundWindow(lockCode co.LSFW) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_LockSetForegroundWindow, "LockSetForegroundWindow"),
		uintptr(lockCode))
	return utl.ZeroAsGetLastError(ret, err)
}

var _LockSetForegroundWindow *syscall.Proc

// [LockWorkStation] function.
//
// [LockWorkStation]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-lockworkstation
func LockWorkStation() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_LockWorkStation, "LockWorkStation"))
	return utl.ZeroAsGetLastError(ret, err)
}

var _LockWorkStation *syscall.Proc

// [MapVirtualKey] function.
//
// [MapVirtualKey]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapvirtualkeyw
func MapVirtualKey(code co.VK, mapType co.MAPVK) uint32 {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_MapVirtualKeyW, "MapVirtualKeyW"),
		uintptr(uint32(code)),
		uintptr(mapType))
	return uint32(ret)
}

var _MapVirtualKeyW *syscall.Proc

// [OffsetRect] function.
//
// [OffsetRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-offsetrect
func OffsetRect(rc *RECT, dx, dy int) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_OffsetRect, "OffsetRect"),
		uintptr(unsafe.Pointer(rc)),
		uintptr(dx),
		uintptr(dy))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _OffsetRect *syscall.Proc

// [PeekMessage] function.
//
// [PeekMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-peekmessagew
func PeekMessage(msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax co.WM, removeMsg co.PM) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_PeekMessageW, "PeekMessageW"),
		uintptr(unsafe.Pointer(msg)),
		uintptr(hWnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
		uintptr(removeMsg))
	return ret != 0
}

var _PeekMessageW *syscall.Proc

// [PostQuitMessage] function.
//
// [PostQuitMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int) {
	syscall.SyscallN(
		dll.Load(dll.USER32, &_PostQuitMessage, "PostQuitMessage"),
		uintptr(exitCode))
}

var _PostQuitMessage *syscall.Proc

// [PostThreadMessage] function.
//
// [PostThreadMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessage(idThread uint32, msg co.WM, wParam WPARAM, lParam LPARAM) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_PostThreadMessageW, "PostThreadMessageW"),
		uintptr(idThread),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return utl.ZeroAsGetLastError(ret, err)
}

var _PostThreadMessageW *syscall.Proc

// [RegisterClassEx] function.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wcx *WNDCLASSEX) (ATOM, error) {
	wcx.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_RegisterClassExW, "RegisterClassExW"),
		uintptr(unsafe.Pointer(wcx)))

	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return ATOM(0), wErr
	} else {
		return ATOM(ret), nil
	}
}

var _RegisterClassExW *syscall.Proc

// [RegisterClipboardFormat] function.
//
// [RegisterClipboardFormat]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclipboardformatw
func RegisterClipboardFormat(name string) (co.CF, error) {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pName := wbuf.PtrAllowEmpty(name)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_RegisterClipboardFormatW, "RegisterClipboardFormatW"),
		uintptr(pName))
	if ret == 0 {
		return co.CF(0), co.ERROR(err)
	}
	return co.CF(ret), nil
}

var _RegisterClipboardFormatW *syscall.Proc

// [RegisterWindowMessage] function.
//
// [RegisterWindowMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessage(message string) (co.WM, error) {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pMessage := wbuf.PtrEmptyIsNil(message)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_RegisterWindowMessageW, "RegisterWindowMessageW"),
		uintptr(pMessage))

	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return co.WM(0), wErr
	} else {
		return co.WM(ret), nil
	}
}

var _RegisterWindowMessageW *syscall.Proc

// [ReplyMessage] function.
//
// [ReplyMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-replymessage
func ReplyMessage(result uintptr) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_ReplyMessage, "ReplyMessage"),
		result)
	return ret != 0
}

var _ReplyMessage *syscall.Proc

// [SendInput] function.
//
// [SendInput]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendinput
func SendInput(inputs []INPUT) (uint, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SendInput, "SendInput"),
		uintptr(uint32(len(inputs))),
		uintptr(unsafe.Pointer(&inputs[0])),
		uintptr(uint32(unsafe.Sizeof(INPUT{}))))

	if wErr := co.ERROR(err); wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint(ret), nil
}

var _SendInput *syscall.Proc

// [SetCaretPos] function.
//
// [SetCaretPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setcaretpos
func SetCaretPos(x, y int) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetCaretPos, "SetCaretPos"),
		uintptr(x),
		uintptr(y))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetCaretPos *syscall.Proc

// [SetCursorPos] function.
//
// [SetCursorPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setcursorpos
func SetCursorPos(x, y int) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetCursorPos, "SetCursorPos"),
		uintptr(x),
		uintptr(y))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetCursorPos *syscall.Proc

// [SetMessageExtraInfo] function.
//
// [SetMessageExtraInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmessageextrainfo
func SetMessageExtraInfo(lp LPARAM) LPARAM {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetMessageExtraInfo, "SetMessageExtraInfo"),
		uintptr(lp))
	return LPARAM(ret)
}

var _SetMessageExtraInfo *syscall.Proc

// [SetProcessDefaultLayout] function.
//
// [SetProcessDefaultLayout]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdefaultlayout
func SetProcessDefaultLayout(defaultLayout co.LAYOUT) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetProcessDefaultLayout, "SetProcessDefaultLayout"),
		uintptr(defaultLayout))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetProcessDefaultLayout *syscall.Proc

// [SetProcessDPIAware] function.
//
// [SetProcessDPIAware]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiaware
func SetProcessDPIAware() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetProcessDPIAware, "SetProcessDPIAware"))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetProcessDPIAware *syscall.Proc

// [ShowCursor] function.
//
// [ShowCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showcursor
func ShowCursor(show bool) int {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_ShowCursor, "ShowCursor"),
		utl.BoolToUintptr(show))
	return int(ret)
}

var _ShowCursor *syscall.Proc

// [SoundSentry] function.
//
// [SoundSentry]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-soundsentry
func SoundSentry() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_SoundSentry, "SoundSentry"))
	return ret != 0
}

var _SoundSentry *syscall.Proc

// [SystemParametersInfo] function.
//
// [SystemParametersInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
func SystemParametersInfo(
	uiAction co.SPI,
	uiParam uint32,
	pvParam unsafe.Pointer,
	fWinIni co.SPIF,
) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SystemParametersInfoW, "SystemParametersInfoW"),
		uintptr(uiAction),
		uintptr(uiParam),
		uintptr(pvParam),
		uintptr(fWinIni))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SystemParametersInfoW *syscall.Proc

// [TranslateMessage] function.
//
// [TranslateMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_TranslateMessage, "TranslateMessage"),
		uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

var _TranslateMessage *syscall.Proc

// [UnregisterClass] function.
//
// Paired with [RegisterClassEx].
//
// [UnregisterClass]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unregisterclassw
func UnregisterClass(className ClassName, hInst HINSTANCE) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pClassName := className.raw(&wbuf)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_UnregisterClassW, "UnregisterClassW"),
		pClassName,
		uintptr(hInst))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return wErr
	} else {
		return nil
	}
}

var _UnregisterClassW *syscall.Proc

// [WaitMessage] function.
//
// [WaitMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-waitmessage
func WaitMessage() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_WaitMessage, "WaitMessage"))
	return utl.ZeroAsGetLastError(ret, err)
}

var _WaitMessage *syscall.Proc
