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
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_AdjustWindowRectEx),
		uintptr(unsafe.Pointer(rc)),
		uintptr(style),
		utl.BoolToUintptr(hasMenu),
		uintptr(exStyle))
	return utl.ZeroAsGetLastError(ret, err)
}

// [AllowSetForegroundWindow] function.
//
// [AllowSetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-allowsetforegroundwindow
func AllowSetForegroundWindow(processId uint32) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_AllowSetForegroundWindow),
		uintptr(processId))
	return utl.ZeroAsGetLastError(ret, err)
}

// [AnyPopup] function.
//
// [AnyPopup]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-anypopup
func AnyPopup() bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_AnyPopup))
	return ret != 0
}

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
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_BroadcastSystemMessageW),
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
	fmtVersion int,
	cxDesired, cyDesired uint,
	flags co.LR,
) (HCURSOR, error) {
	hIcon, err := CreateIconFromResourceEx(
		resBits, fmtVersion, cxDesired, cyDesired, flags)
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
	fmtVersion int,
	cxDesired, cyDesired uint,
	flags co.LR,
) (HICON, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_CreateIconFromResourceEx),
		uintptr(unsafe.Pointer(&resBits[0])),
		uintptr(len(resBits)),
		1,
		uintptr(fmtVersion),
		uintptr(cxDesired),
		uintptr(cyDesired),
		uintptr(flags))
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

// [DestroyCaret] function.
//
// [DestroyCaret]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycarret
func DestroyCaret() error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_DestroyCaret))
	return utl.ZeroAsGetLastError(ret, err)
}

// [DispatchMessage] function.
//
// [DispatchMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_DispatchMessageW),
		uintptr(unsafe.Pointer(msg)))
	return ret
}

// [EndMenu] function.
//
// [EndMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endmenu
func EndMenu() error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_EndMenu))
	return utl.ZeroAsGetLastError(ret, err)
}

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
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pDevice := wbuf.PtrEmptyIsNil(device)

	devices := make([]DISPLAY_DEVICE, 0) // to be returned
	devNum := 0

	var dide DISPLAY_DEVICE // buffer to receive each iteration
	dide.SetCb()

	for {
		// Ignore errors: only fails with devNum out-of-bounds, which never happens here.
		ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_EnumDisplayDevicesW),
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

// [EnumThreadWindows] function.
//
// [EnumThreadWindows]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumthreadwindows
func EnumThreadWindows(threadId uint32) []HWND {
	pPack := &_EnumThreadWindowsPack{
		arr: make([]HWND, 0),
	}

	syscall.SyscallN(dll.User(dll.PROC_EnumThreadWindows),
		enumThreadWindowsCallback(),
		uintptr(unsafe.Pointer(pPack)))
	return pPack.arr
}

type _EnumThreadWindowsPack struct{ arr []HWND }

var _enumThreadWindowsCallback uintptr

func enumThreadWindowsCallback() uintptr {
	if _enumThreadWindowsCallback == 0 {
		_enumThreadWindowsCallback = syscall.NewCallback(
			func(hWnd HWND, lParam LPARAM) uintptr {
				pPack := (*_EnumThreadWindowsPack)(unsafe.Pointer(lParam))
				pPack.arr = append(pPack.arr, hWnd)
				return 1
			},
		)
	}
	return _enumThreadWindowsCallback
}

// [EnumWindows] function.
//
// [EnumWindows]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumwindows
func EnumWindows() []HWND {
	pPack := &_EnumWindowsPack{
		arr: make([]HWND, 0),
	}

	syscall.SyscallN(dll.User(dll.PROC_EnumWindows),
		enumWindowsCallback(),
		uintptr(unsafe.Pointer(pPack)))
	return pPack.arr
}

type _EnumWindowsPack struct{ arr []HWND }

var _enumWindowsCallback uintptr

func enumWindowsCallback() uintptr {
	if _enumWindowsCallback == 0 {
		_enumWindowsCallback = syscall.NewCallback(
			func(hWnd HWND, lParam LPARAM) uintptr {
				pPack := (*_EnumWindowsPack)(unsafe.Pointer(lParam))
				pPack.arr = append(pPack.arr, hWnd)
				return 1
			},
		)
	}
	return _enumWindowsCallback
}

// [GetAsyncKeyState] function.
//
// [GetAsyncKeyState]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getasynckeystate
func GetAsyncKeyState(virtKeyCode co.VK) uint16 {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetAsyncKeyState),
		uintptr(virtKeyCode))
	return uint16(ret)
}

// [GetCaretPos] function.
//
// [GetCaretPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcaretpos
func GetCaretPos() (RECT, error) {
	var rc RECT
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetCaretPos),
		uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		return RECT{}, co.ERROR(err)
	}
	return rc, nil
}

// [GetCursorInfo] function.
//
// [GetCursorInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorinfo
func GetCursorInfo() (CURSORINFO, error) {
	var ci CURSORINFO
	ci.SetCbSize()

	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetCursorInfo),
		uintptr(unsafe.Pointer(&ci)))
	if ret == 0 {
		return CURSORINFO{}, co.ERROR(err)
	}
	return ci, nil
}

// [GetCursorPos] function.
//
// [GetCursorPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
func GetCursorPos() (POINT, error) {
	var pt POINT
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetCursorPos),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR(err)
	}
	return pt, nil
}

// [GetDialogBaseUnits] function.
//
// [GetDialogBaseUnits]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdialogbaseunits
func GetDialogBaseUnits() (horz, vert uint16) {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetDialogBaseUnits))
	horz, vert = LOWORD(uint32(ret)), HIWORD(uint32(ret))
	return
}

// [GetGUIThreadInfo] function.
//
// [GetGUIThreadInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getguithreadinfo
func GetGUIThreadInfo(thread_id uint32) (GUITHREADINFO, error) {
	var info GUITHREADINFO
	info.SetCbSize()

	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetGUIThreadInfo),
		uintptr(thread_id),
		uintptr(unsafe.Pointer(&info)))
	if ret == 0 {
		return GUITHREADINFO{}, co.ERROR(err)
	}
	return info, nil
}

// [GetInputState] function.
//
// [GetInputState]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getinputstate
func GetInputState() bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetInputState))
	return ret != 0
}

// [GetMessage] function.
//
// [GetMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagew
func GetMessage(msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax uint32) (int32, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetMessageW),
		uintptr(unsafe.Pointer(msg)),
		uintptr(hWnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))
	if int32(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return int32(ret), nil
}

// [GetMessageExtraInfo] function.
//
// [GetMessageExtraInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessageextrainfo
func GetMessageExtraInfo() LPARAM {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetMessageExtraInfo))
	return LPARAM(ret)
}

// [GetMessagePos] function.
//
// [GetMessagePos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagepos
func GetMessagePos() POINT {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetMessagePos))
	return POINT{
		X: int32(LOWORD(uint32(ret))),
		Y: int32(HIWORD(uint32(ret))),
	}
}

// [GetMessageTime] function.
//
// [GetMessageTime]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagetime
func GetMessageTime() time.Duration {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetMessageTime))
	return time.Duration(ret * uintptr(time.Millisecond))
}

// [GetPhysicalCursorPos] function.
//
// [GetPhysicalCursorPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getphysicalcursorpos
func GetPhysicalCursorPos() (POINT, error) {
	var pt POINT
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetPhysicalCursorPos),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR(err)
	}
	return pt, nil
}

// [GetProcessDefaultLayout] function.
//
// [GetProcessDefaultLayout]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getprocessdefaultlayout
func GetProcessDefaultLayout() (co.LAYOUT, error) {
	var defaultLayout co.LAYOUT
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetProcessDefaultLayout),
		uintptr(unsafe.Pointer(&defaultLayout)))
	if ret == 0 {
		return co.LAYOUT(0), co.ERROR(err)
	}
	return defaultLayout, nil
}

// [GetQueueStatus] function.
//
// [GetQueueStatus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getqueuestatus
func GetQueueStatus(flags co.QS) (currentlyInQueue, addedToQueue co.QS) {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetQueueStatus),
		uintptr(flags))
	currentlyInQueue = co.QS(HIWORD(uint32(ret)))
	addedToQueue = co.QS(LOWORD(uint32(ret)))
	return
}

// [GetSysColor] function.
//
// [GetSysColor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
func GetSysColor(index co.COLOR) COLORREF {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetSysColor),
		uintptr(index))
	return COLORREF(ret)
}

// [GetSystemMetrics] function.
//
// [GetSystemMetrics]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
func GetSystemMetrics(index co.SM) int32 {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetSystemMetrics),
		uintptr(index))
	return int32(ret)
}

// [InflateRect] function.
//
// [InflateRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-inflaterect
func InflateRect(rc *RECT, dx, dy int) error {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_InflateRect),
		uintptr(unsafe.Pointer(rc)),
		uintptr(dx),
		uintptr(dy))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [InSendMessage] function.
//
// [InSendMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessage
func InSendMessage() bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_InSendMessage))
	return ret != 0
}

// [InSendMessageEx] function.
//
// [InSendMessageEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessageex
func InSendMessageEx() co.ISMEX {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_InSendMessageEx))
	return co.ISMEX(ret)
}

// [IsGUIThread] function.
//
// [IsGUIThread]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isguithread
func IsGUIThread(convertToGuiThread bool) (bool, error) {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_IsGUIThread),
		utl.BoolToUintptr(convertToGuiThread))
	if convertToGuiThread && co.ERROR(ret) == co.ERROR_NOT_ENOUGH_MEMORY {
		return false, co.ERROR_NOT_ENOUGH_MEMORY
	}
	return ret != 0, nil
}

// [LockSetForegroundWindow] function.
//
// [LockSetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-locksetforegroundwindow
func LockSetForegroundWindow(lockCode co.LSFW) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_LockSetForegroundWindow),
		uintptr(lockCode))
	return utl.ZeroAsGetLastError(ret, err)
}

// [OffsetRect] function.
//
// [OffsetRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-offsetrect
func OffsetRect(rc *RECT, dx, dy int) error {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_OffsetRect),
		uintptr(unsafe.Pointer(rc)),
		uintptr(dx),
		uintptr(dy))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [PeekMessage] function.
//
// [PeekMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-peekmessagew
func PeekMessage(msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax co.WM, removeMsg co.PM) bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_PeekMessageW),
		uintptr(unsafe.Pointer(msg)),
		uintptr(hWnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
		uintptr(removeMsg))
	return ret != 0
}

// [PostQuitMessage] function.
//
// [PostQuitMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int) {
	syscall.SyscallN(dll.User(dll.PROC_PostQuitMessage),
		uintptr(exitCode))
}

// [PostThreadMessage] function.
//
// [PostThreadMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessage(idThread uint32, msg co.WM, wParam WPARAM, lParam LPARAM) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_PostThreadMessageW),
		uintptr(idThread),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return utl.ZeroAsGetLastError(ret, err)
}

// [RegisterClassEx] function.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wcx *WNDCLASSEX) (ATOM, error) {
	wcx.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_RegisterClassExW),
		uintptr(unsafe.Pointer(wcx)))

	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return ATOM(0), wErr
	} else {
		return ATOM(ret), nil
	}
}

// [RegisterClipboardFormat] function.
//
// [RegisterClipboardFormat]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclipboardformatw
func RegisterClipboardFormat(name string) (co.CF, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pName := wbuf.PtrAllowEmpty(name)

	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_RegisterClipboardFormatW),
		uintptr(pName))
	if ret == 0 {
		return co.CF(0), co.ERROR(err)
	}
	return co.CF(ret), nil
}

// [RegisterWindowMessage] function.
//
// [RegisterWindowMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessage(message string) (co.WM, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pMessage := wbuf.PtrEmptyIsNil(message)

	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_RegisterWindowMessageW),
		uintptr(pMessage))

	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return co.WM(0), wErr
	} else {
		return co.WM(ret), nil
	}
}

// [ReplyMessage] function.
//
// [ReplyMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-replymessage
func ReplyMessage(result uintptr) bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_ReplyMessage),
		result)
	return ret != 0
}

// [SetCaretPos] function.
//
// [SetCaretPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setcaretpos
func SetCaretPos(x, y int) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_SetCaretPos),
		uintptr(x),
		uintptr(y))
	return utl.ZeroAsGetLastError(ret, err)
}

// [SetCursorPos] function.
//
// [SetCursorPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setcursorpos
func SetCursorPos(x, y int) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_SetCursorPos),
		uintptr(x),
		uintptr(y))
	return utl.ZeroAsGetLastError(ret, err)
}

// [SetMessageExtraInfo] function.
//
// [SetMessageExtraInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmessageextrainfo
func SetMessageExtraInfo(lp LPARAM) LPARAM {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_SetMessageExtraInfo),
		uintptr(lp))
	return LPARAM(ret)
}

// [SetProcessDefaultLayout] function.
//
// [SetProcessDefaultLayout]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdefaultlayout
func SetProcessDefaultLayout(defaultLayout co.LAYOUT) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_SetProcessDefaultLayout),
		uintptr(defaultLayout))
	return utl.ZeroAsGetLastError(ret, err)
}

// [SetProcessDPIAware] function.
//
// [SetProcessDPIAware]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiaware
func SetProcessDPIAware() error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_SetProcessDPIAware))
	return utl.ZeroAsGetLastError(ret, err)
}

// [ShowCursor] function.
//
// [ShowCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showcursor
func ShowCursor(show bool) int {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_ShowCursor),
		utl.BoolToUintptr(show))
	return int(ret)
}

// [SoundSentry] function.
//
// [SoundSentry]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-soundsentry
func SoundSentry() bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_SoundSentry))
	return ret != 0
}

// [SystemParametersInfo] function.
//
// [SystemParametersInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
func SystemParametersInfo(
	uiAction co.SPI,
	uiParam uint32,
	pvParam unsafe.Pointer,
	fWinIni co.SPIF,
) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_SystemParametersInfoW),
		uintptr(uiAction),
		uintptr(uiParam),
		uintptr(pvParam),
		uintptr(fWinIni))
	return utl.ZeroAsGetLastError(ret, err)
}

// [TranslateMessage] function.
//
// [TranslateMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_TranslateMessage),
		uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

// [UnregisterClass] function.
//
// Paired with [RegisterClassEx].
//
// [UnregisterClass]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unregisterclassw
func UnregisterClass(className ClassName, hInst HINSTANCE) error {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pClassName := className.raw(&wbuf)

	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_UnregisterClassW),
		pClassName,
		uintptr(hInst))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return wErr
	} else {
		return nil
	}
}

// [WaitMessage] function.
//
// [WaitMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-waitmessage
func WaitMessage() error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_WaitMessage))
	return utl.ZeroAsGetLastError(ret, err)
}
