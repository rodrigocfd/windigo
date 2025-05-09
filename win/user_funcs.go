//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [AdjustWindowRectEx] function.
//
// [AdjustWindowRectEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-adjustwindowrectex
func AdjustWindowRectEx(rc *RECT, style co.WS, hasMenu bool, exStyle co.WS_EX) error {
	ret, _, err := syscall.SyscallN(_AdjustWindowRectEx.Addr(),
		uintptr(unsafe.Pointer(rc)), uintptr(style),
		wutil.BoolToUintptr(hasMenu), uintptr(exStyle))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _AdjustWindowRectEx = dll.User32.NewProc("AdjustWindowRectEx")

// [AllowSetForegroundWindow] function.
//
// [AllowSetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-allowsetforegroundwindow
func AllowSetForegroundWindow(processId uint32) error {
	ret, _, err := syscall.SyscallN(_AllowSetForegroundWindow.Addr(),
		uintptr(processId))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _AllowSetForegroundWindow = dll.User32.NewProc("AllowSetForegroundWindow")

// [AnyPopup] function.
//
// [AnyPopup]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-anypopup
func AnyPopup() bool {
	ret, _, _ := syscall.SyscallN(_AnyPopup.Addr())
	return ret != 0
}

var _AnyPopup = dll.User32.NewProc("AnyPopup")

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
	ret, _, err := syscall.SyscallN(_BroadcastSystemMessageW.Addr(),
		uintptr(flags), uintptr(unsafe.Pointer(&receivers)),
		uintptr(msg), uintptr(wParam), uintptr(lParam))

	broadcastSuccessful = int(ret) > 1
	if ret == 0 {
		wErr = co.ERROR(err)
	}
	return
}

var _BroadcastSystemMessageW = dll.User32.NewProc("BroadcastSystemMessageW")

// [CreateIconFromResourceEx] function for cursor.
//
// This function creates HCURSOR only. The HICON variation is
// CreateIconFromResourceEx().
//
// ⚠️ You must defer HCURSOR.DestroyCursor().
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
// This function creates HICON only. The HCURSOR variation is
// CreateCursorFromResourceEx().
//
// ⚠️ You must defer HICON.DestroyIcon().
//
// [CreateIconFromResourceEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconfromresourceex
func CreateIconFromResourceEx(
	resBits []byte,
	fmtVersion int,
	cxDesired, cyDesired uint,
	flags co.LR,
) (HICON, error) {
	ret, _, err := syscall.SyscallN(_CreateIconFromResourceEx.Addr(),
		uintptr(unsafe.Pointer(&resBits[0])), uintptr(len(resBits)),
		1, uintptr(fmtVersion), uintptr(cxDesired), uintptr(cyDesired),
		uintptr(flags))
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

var _CreateIconFromResourceEx = dll.User32.NewProc("CreateIconFromResourceEx")

// [DestroyCaret] function.
//
// [DestroyCaret]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycarret
func DestroyCaret() error {
	ret, _, err := syscall.SyscallN(_DestroyCaret.Addr())
	return wutil.ZeroAsGetLastError(ret, err)
}

var _DestroyCaret = dll.User32.NewProc("DestroyCaret")

// [DispatchMessage] function.
//
// [DispatchMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.SyscallN(_DispatchMessageW.Addr(),
		uintptr(unsafe.Pointer(msg)))
	return ret
}

var _DispatchMessageW = dll.User32.NewProc("DispatchMessageW")

// [EndMenu] function.
//
// [EndMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endmenu
func EndMenu() error {
	ret, _, err := syscall.SyscallN(_EndMenu.Addr())
	return wutil.ZeroAsGetLastError(ret, err)
}

var _EndMenu = dll.User32.NewProc("EndMenu")

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
	devices := make([]DISPLAY_DEVICE, 0)
	devNum := 0
	device16 := wstr.NewBufWith[wstr.Stack20](device, wstr.EMPTY_IS_NIL)

	var dide DISPLAY_DEVICE
	dide.SetCb()

	for {
		// Ignore errors: only fails with devNum out-of-bounds, which never happens here.
		ret, _, _ := syscall.SyscallN(_EnumDisplayDevicesW.Addr(),
			uintptr(device16.UnsafePtr()), uintptr(devNum),
			uintptr(unsafe.Pointer(&dide)), uintptr(flags))
		if ret == 0 {
			break
		}
		devices = append(devices, dide)
		devNum++
	}

	return devices
}

var _EnumDisplayDevicesW = dll.User32.NewProc("EnumDisplayDevicesW")

// [EnumThreadWindows] function.
//
// [EnumThreadWindows]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumthreadwindows
func EnumThreadWindows(threadId uint32) []HWND {
	pPack := &_EnumThreadWindowsPack{
		arr: make([]HWND, 0),
	}

	syscall.SyscallN(_EnumThreadWindows.Addr(),
		enumThreadWindowsCallback(), uintptr(unsafe.Pointer(pPack)))
	return pPack.arr
}

type _EnumThreadWindowsPack struct{ arr []HWND }

var (
	_EnumThreadWindows         = dll.User32.NewProc("EnumThreadWindows")
	_enumThreadWindowsCallback uintptr
)

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

	syscall.SyscallN(_EnumWindows.Addr(),
		enumWindowsCallback(), uintptr(unsafe.Pointer(pPack)))
	return pPack.arr
}

type _EnumWindowsPack struct{ arr []HWND }

var (
	_EnumWindows         = dll.User32.NewProc("EnumWindows")
	_enumWindowsCallback uintptr
)

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
	ret, _, _ := syscall.SyscallN(_GetAsyncKeyState.Addr(),
		uintptr(virtKeyCode))
	return uint16(ret)
}

var _GetAsyncKeyState = dll.User32.NewProc("GetAsyncKeyState")

// [GetCaretPos] function.
//
// [GetCaretPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcaretpos
func GetCaretPos() (RECT, error) {
	var rc RECT
	ret, _, err := syscall.SyscallN(_GetCaretPos.Addr(),
		uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		return RECT{}, co.ERROR(err)
	}
	return rc, nil
}

var _GetCaretPos = dll.User32.NewProc("GetCaretPos")

// [GetCursorInfo] function.
//
// [GetCursorInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorinfo
func GetCursorInfo() (CURSORINFO, error) {
	var ci CURSORINFO
	ci.SetCbSize()

	ret, _, err := syscall.SyscallN(_GetCursorInfo.Addr(),
		uintptr(unsafe.Pointer(&ci)))
	if ret == 0 {
		return CURSORINFO{}, co.ERROR(err)
	}
	return ci, nil
}

var _GetCursorInfo = dll.User32.NewProc("GetCursorInfo")

// [GetCursorPos] function.
//
// [GetCursorPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
func GetCursorPos() (POINT, error) {
	var pt POINT
	ret, _, err := syscall.SyscallN(_GetCursorPos.Addr(),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR(err)
	}
	return pt, nil
}

var _GetCursorPos = dll.User32.NewProc("GetCursorPos")

// [GetDialogBaseUnits] function.
//
// [GetDialogBaseUnits]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdialogbaseunits
func GetDialogBaseUnits() (horz, vert uint16) {
	ret, _, _ := syscall.SyscallN(_GetDialogBaseUnits.Addr())
	horz, vert = LOWORD(uint32(ret)), HIWORD(uint32(ret))
	return
}

var _GetDialogBaseUnits = dll.User32.NewProc("GetDialogBaseUnits")

// [GetGUIThreadInfo] function.
//
// [GetGUIThreadInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getguithreadinfo
func GetGUIThreadInfo(thread_id uint32) (GUITHREADINFO, error) {
	var info GUITHREADINFO
	info.SetCbSize()

	ret, _, err := syscall.SyscallN(_GetGUIThreadInfo.Addr(),
		uintptr(thread_id), uintptr(unsafe.Pointer(&info)))
	if ret == 0 {
		return GUITHREADINFO{}, co.ERROR(err)
	}
	return info, nil
}

var _GetGUIThreadInfo = dll.User32.NewProc("GetGUIThreadInfo")

// [GetInputState] function.
//
// [GetInputState]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getinputstate
func GetInputState() bool {
	ret, _, _ := syscall.SyscallN(_GetInputState.Addr())
	return ret != 0
}

var _GetInputState = dll.User32.NewProc("GetInputState")

// [GetMessage] function.
//
// [GetMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagew
func GetMessage(msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax uint32) (int32, error) {
	ret, _, err := syscall.SyscallN(_GetMessageW.Addr(),
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax))
	if int32(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return int32(ret), nil
}

var _GetMessageW = dll.User32.NewProc("GetMessageW")

// [GetMessageExtraInfo] function.
//
// [GetMessageExtraInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessageextrainfo
func GetMessageExtraInfo() LPARAM {
	ret, _, _ := syscall.SyscallN(_GetMessageExtraInfo.Addr())
	return LPARAM(ret)
}

var _GetMessageExtraInfo = dll.User32.NewProc("GetMessageExtraInfo")

// [GetMessagePos] function.
//
// [GetMessagePos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagepos
func GetMessagePos() POINT {
	ret, _, _ := syscall.SyscallN(_GetMessagePos.Addr())
	return POINT{
		X: int32(LOWORD(uint32(ret))),
		Y: int32(HIWORD(uint32(ret))),
	}
}

var _GetMessagePos = dll.User32.NewProc("GetMessagePos")

// [GetMessageTime] function.
//
// [GetMessageTime]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagetime
func GetMessageTime() time.Duration {
	ret, _, _ := syscall.SyscallN(_GetMessageTime.Addr())
	return time.Duration(ret * uintptr(time.Millisecond))
}

var _GetMessageTime = dll.User32.NewProc("GetMessageTime")

// [GetPhysicalCursorPos] function.
//
// [GetPhysicalCursorPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getphysicalcursorpos
func GetPhysicalCursorPos() (POINT, error) {
	var pt POINT
	ret, _, err := syscall.SyscallN(_GetPhysicalCursorPos.Addr(),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, co.ERROR(err)
	}
	return pt, nil
}

var _GetPhysicalCursorPos = dll.User32.NewProc("GetPhysicalCursorPos")

// [GetProcessDefaultLayout] function.
//
// [GetProcessDefaultLayout]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getprocessdefaultlayout
func GetProcessDefaultLayout() (co.LAYOUT, error) {
	var defaultLayout co.LAYOUT
	ret, _, err := syscall.SyscallN(_GetProcessDefaultLayout.Addr(),
		uintptr(unsafe.Pointer(&defaultLayout)))
	if ret == 0 {
		return co.LAYOUT(0), co.ERROR(err)
	}
	return defaultLayout, nil
}

var _GetProcessDefaultLayout = dll.User32.NewProc("GetProcessDefaultLayout")

// [GetQueueStatus] function.
//
// [GetQueueStatus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getqueuestatus
func GetQueueStatus(flags co.QS) (currentlyInQueue, addedToQueue co.QS) {
	ret, _, _ := syscall.SyscallN(_GetQueueStatus.Addr(),
		uintptr(flags))
	currentlyInQueue = co.QS(HIWORD(uint32(ret)))
	addedToQueue = co.QS(LOWORD(uint32(ret)))
	return
}

var _GetQueueStatus = dll.User32.NewProc("GetQueueStatus")

// [GetSysColor] function.
//
// [GetSysColor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
func GetSysColor(index co.COLOR) COLORREF {
	ret, _, _ := syscall.SyscallN(_GetSysColor.Addr(),
		uintptr(index))
	return COLORREF(ret)
}

var _GetSysColor = dll.User32.NewProc("GetSysColor")

// [GetSystemMetrics] function.
//
// [GetSystemMetrics]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
func GetSystemMetrics(index co.SM) int32 {
	ret, _, _ := syscall.SyscallN(_GetSystemMetrics.Addr(),
		uintptr(index))
	return int32(ret)
}

var _GetSystemMetrics = dll.User32.NewProc("GetSystemMetrics")

// [InflateRect] function.
//
// [InflateRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-inflaterect
func InflateRect(rc *RECT, dx, dy int) error {
	ret, _, _ := syscall.SyscallN(_InflateRect.Addr(),
		uintptr(unsafe.Pointer(rc)), uintptr(dx), uintptr(dy))
	return wutil.ZeroAsSysInvalidParm(ret)
}

var _InflateRect = dll.User32.NewProc("InflateRect")

// [InSendMessage] function.
//
// [InSendMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessage
func InSendMessage() bool {
	ret, _, _ := syscall.SyscallN(_InSendMessage.Addr())
	return ret != 0
}

var _InSendMessage = dll.User32.NewProc("InSendMessage")

// [InSendMessageEx] function.
//
// [InSendMessageEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessageex
func InSendMessageEx() co.ISMEX {
	ret, _, _ := syscall.SyscallN(_InSendMessageEx.Addr())
	return co.ISMEX(ret)
}

var _InSendMessageEx = dll.User32.NewProc("InSendMessageEx")

// [IsGUIThread] function.
//
// [IsGUIThread]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isguithread
func IsGUIThread(convertToGuiThread bool) (bool, error) {
	ret, _, _ := syscall.SyscallN(_IsGUIThread.Addr(),
		wutil.BoolToUintptr(convertToGuiThread))
	if convertToGuiThread && co.ERROR(ret) == co.ERROR_NOT_ENOUGH_MEMORY {
		return false, co.ERROR_NOT_ENOUGH_MEMORY
	}
	return ret != 0, nil
}

var _IsGUIThread = dll.User32.NewProc("IsGUIThread")

// [LockSetForegroundWindow] function.
//
// [LockSetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-locksetforegroundwindow
func LockSetForegroundWindow(lockCode co.LSFW) error {
	ret, _, err := syscall.SyscallN(_LockSetForegroundWindow.Addr(),
		uintptr(lockCode))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _LockSetForegroundWindow = dll.User32.NewProc("LockSetForegroundWindow")

// [OffsetRect] function.
//
// [OffsetRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-offsetrect
func OffsetRect(rc *RECT, dx, dy int) error {
	ret, _, _ := syscall.SyscallN(_OffsetRect.Addr(),
		uintptr(unsafe.Pointer(rc)), uintptr(dx), uintptr(dy))
	return wutil.ZeroAsSysInvalidParm(ret)
}

var _OffsetRect = dll.User32.NewProc("OffsetRect")

// [PeekMessage] function.
//
// [PeekMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-peekmessagew
func PeekMessage(msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax co.WM, removeMsg co.PM) bool {
	ret, _, _ := syscall.SyscallN(_PeekMessageW.Addr(),
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax), uintptr(removeMsg))
	return ret != 0
}

var _PeekMessageW = dll.User32.NewProc("PeekMessageW")

// [PostQuitMessage] function.
//
// [PostQuitMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int) {
	syscall.SyscallN(_PostQuitMessage.Addr(),
		uintptr(exitCode))
}

var _PostQuitMessage = dll.User32.NewProc("PostQuitMessage")

// [PostThreadMessage] function.
//
// [PostThreadMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessage(idThread uint32, msg co.WM, wParam WPARAM, lParam LPARAM) error {
	ret, _, err := syscall.SyscallN(_PostThreadMessageW.Addr(),
		uintptr(idThread), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _PostThreadMessageW = dll.User32.NewProc("PostThreadMessageW")

// [RegisterClassEx] function.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wcx *WNDCLASSEX) (ATOM, error) {
	wcx.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(_RegisterClassExW.Addr(),
		uintptr(unsafe.Pointer(wcx)))

	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return ATOM(0), wErr
	} else {
		return ATOM(ret), nil
	}
}

var _RegisterClassExW = dll.User32.NewProc("RegisterClassExW")

// [RegisterClipboardFormat] function.
//
// [RegisterClipboardFormat]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclipboardformatw
func RegisterClipboardFormat(name string) (co.CF, error) {
	name16 := wstr.NewBufWith[wstr.Stack20](name, wstr.ALLOW_EMPTY)
	ret, _, err := syscall.SyscallN(_RegisterClipboardFormatW.Addr(),
		uintptr(name16.UnsafePtr()))
	if ret == 0 {
		return co.CF(0), co.ERROR(err)
	}
	return co.CF(ret), nil
}

var _RegisterClipboardFormatW = dll.User32.NewProc("RegisterClipboardFormatW")

// [RegisterWindowMessage] function.
//
// [RegisterWindowMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessage(message string) (co.WM, error) {
	message16 := wstr.NewBufWith[wstr.Stack20](message, wstr.EMPTY_IS_NIL)
	ret, _, err := syscall.SyscallN(_RegisterWindowMessageW.Addr(),
		uintptr(message16.UnsafePtr()))

	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return co.WM(0), wErr
	} else {
		return co.WM(ret), nil
	}
}

var _RegisterWindowMessageW = dll.User32.NewProc("RegisterWindowMessageW")

// [ReplyMessage] function.
//
// [ReplyMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-replymessage
func ReplyMessage(result uintptr) bool {
	ret, _, _ := syscall.SyscallN(_ReplyMessage.Addr(),
		result)
	return ret != 0
}

var _ReplyMessage = dll.User32.NewProc("ReplyMessage")

// [SetCaretPos] function.
//
// [SetCaretPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setcaretpos
func SetCaretPos(x, y int) error {
	ret, _, err := syscall.SyscallN(_SetCaretPos.Addr(),
		uintptr(x), uintptr(y))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _SetCaretPos = dll.User32.NewProc("SetCaretPos")

// [SetCursorPos] function.
//
// [SetCursorPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setcursorpos
func SetCursorPos(x, y int) error {
	ret, _, err := syscall.SyscallN(_SetCursorPos.Addr(),
		uintptr(x), uintptr(y))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _SetCursorPos = dll.User32.NewProc("SetCursorPos")

// [SetMessageExtraInfo] function.
//
// [SetMessageExtraInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmessageextrainfo
func SetMessageExtraInfo(lp LPARAM) LPARAM {
	ret, _, _ := syscall.SyscallN(_SetMessageExtraInfo.Addr(),
		uintptr(lp))
	return LPARAM(ret)
}

var _SetMessageExtraInfo = dll.User32.NewProc("SetMessageExtraInfo")

// [SetProcessDefaultLayout] function.
//
// [SetProcessDefaultLayout]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdefaultlayout
func SetProcessDefaultLayout(defaultLayout co.LAYOUT) error {
	ret, _, err := syscall.SyscallN(_SetProcessDefaultLayout.Addr(),
		uintptr(defaultLayout))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _SetProcessDefaultLayout = dll.User32.NewProc("SetProcessDefaultLayout")

// [SetProcessDPIAware] function.
//
// [SetProcessDPIAware]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiaware
func SetProcessDPIAware() error {
	ret, _, err := syscall.SyscallN(_SetProcessDPIAware.Addr())
	return wutil.ZeroAsGetLastError(ret, err)
}

var _SetProcessDPIAware = dll.User32.NewProc("SetProcessDPIAware")

// [ShowCursor] function.
//
// [ShowCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showcursor
func ShowCursor(show bool) int {
	ret, _, _ := syscall.SyscallN(_ShowCursor.Addr(),
		wutil.BoolToUintptr(show))
	return int(ret)
}

var _ShowCursor = dll.User32.NewProc("ShowCursor")

// [SoundSentry] function.
//
// [SoundSentry]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-soundsentry
func SoundSentry() bool {
	ret, _, _ := syscall.SyscallN(_SoundSentry.Addr())
	return ret != 0
}

var _SoundSentry = dll.User32.NewProc("SoundSentry")

// [SystemParametersInfo] function.
//
// [SystemParametersInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
func SystemParametersInfo(
	uiAction co.SPI,
	uiParam uint32,
	pvParam unsafe.Pointer,
	fWinIni co.SPIF,
) error {
	ret, _, err := syscall.SyscallN(_SystemParametersInfoW.Addr(),
		uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _SystemParametersInfoW = dll.User32.NewProc("SystemParametersInfoW")

// [TranslateMessage] function.
//
// [TranslateMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.SyscallN(_TranslateMessage.Addr(),
		uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

var _TranslateMessage = dll.User32.NewProc("TranslateMessage")

// [UnregisterClass] function.
//
// [UnregisterClass]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unregisterclassw
func UnregisterClass(className ClassName, hInst HINSTANCE) error {
	className16 := wstr.NewBuf[wstr.Stack20]()
	classNameVal := className.raw(&className16)

	ret, _, err := syscall.SyscallN(_UnregisterClassW.Addr(),
		classNameVal, uintptr(hInst))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return wErr
	} else {
		return nil
	}
}

var _UnregisterClassW = dll.User32.NewProc("UnregisterClassW")

// [WaitMessage] function.
//
// [WaitMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-waitmessage
func WaitMessage() error {
	ret, _, err := syscall.SyscallN(_WaitMessage.Addr())
	return wutil.ZeroAsGetLastError(ret, err)
}

var _WaitMessage = dll.User32.NewProc("WaitMessage")
