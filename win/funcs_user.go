//go:build windows

package win

import (
	"runtime"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [AdjustWindowRectEx] function.
//
// [AdjustWindowRectEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-adjustwindowrectex
func AdjustWindowRectEx(rc *RECT, style co.WS, hasMenu bool, exStyle co.WS_EX) {
	ret, _, err := syscall.SyscallN(proc.AdjustWindowRectEx.Addr(),
		uintptr(unsafe.Pointer(rc)), uintptr(style),
		util.BoolToUintptr(hasMenu), uintptr(exStyle))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [AllowSetForegroundWindow] function.
//
// [AllowSetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-allowsetforegroundwindow
func AllowSetForegroundWindow(processId uint32) {
	ret, _, err := syscall.SyscallN(proc.AllowSetForegroundWindow.Addr(),
		uintptr(processId))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [BroadcastSystemMessage] function.
//
// [BroadcastSystemMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-broadcastsystemmessagew
func BroadcastSystemMessage(
	flags co.BSF,
	recipients co.BSM,
	msg co.WM,
	wParam WPARAM,
	lParam LPARAM) (broadcastSuccessful bool, receivers co.BSM, e error) {

	receivers = recipients

	ret, _, err := syscall.SyscallN(proc.BroadcastSystemMessage.Addr(),
		uintptr(flags), uintptr(unsafe.Pointer(&receivers)),
		uintptr(msg), uintptr(wParam), uintptr(lParam))

	broadcastSuccessful = int(ret) > 1
	if ret == 0 {
		e = errco.ERROR(err)
	}
	return
}

// [CreateCursorFromResourceEx] function.
//
// This function creates HCURSOR only. The HICON variation is
// CreateIconFromResourceEx().
//
// ⚠️ You must defer HCURSOR.DestroyCursor().
//
// [CreateCursorFromResourceEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconfromresourceex
func CreateCursorFromResourceEx(
	resBits []byte, fmtVersion int,
	cxDesired, cyDesired int,
	flags co.LR) (HCURSOR, error) {

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
	resBits []byte, fmtVersion int,
	cxDesired, cyDesired int,
	flags co.LR) (HICON, error) {

	ret, _, err := syscall.SyscallN(proc.CreateIconFromResourceEx.Addr(),
		uintptr(unsafe.Pointer(&resBits[0])), uintptr(len(resBits)),
		1, uintptr(fmtVersion), uintptr(cxDesired), uintptr(cyDesired),
		uintptr(flags))
	if ret == 0 {
		return HICON(0), errco.ERROR(err)
	}
	return HICON(ret), nil
}

// [DestroyCaret] function.
//
// [DestroyCaret]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycarret
func DestroyCaret() error {
	ret, _, err := syscall.SyscallN(proc.DestroyCaret.Addr())
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [DispatchMessage] function.
//
// [DispatchMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.SyscallN(proc.DispatchMessage.Addr(),
		uintptr(unsafe.Pointer(msg)))
	return ret
}

// [EndMenu] function.
//
// [EndMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endmenu
func EndMenu() {
	ret, _, err := syscall.SyscallN(proc.EndMenu.Addr())
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [EnumDisplayDevices] function.
//
// To continue enumeration, the callback function must return true; to stop
// enumeration, it must return false.
//
// [EnumDisplayDevices]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaydevicesw
func EnumDisplayDevices(
	device StrOpt, flags co.EDD,
	callback func(devNum int, info *DISPLAY_DEVICE) bool) {

	devicePtr := device.Raw()
	devNum := 0

	dide := DISPLAY_DEVICE{}
	dide.SetCb()

	for {
		ret, _, err := syscall.SyscallN(proc.EnumDisplayDevices.Addr(),
			uintptr(devicePtr), uintptr(devNum),
			uintptr(unsafe.Pointer(&dide)), uintptr(flags))

		if ret == 0 {
			if wErr := errco.ERROR(err); wErr != errco.SUCCESS {
				panic(wErr)
			} else {
				break
			}
		}

		if !callback(devNum, &dide) {
			break
		}
		devNum++
	}

	runtime.KeepAlive(devicePtr)
}

// [EnumWindows] function.
//
// To continue enumeration, the callback function must return true; to stop
// enumeration, it must return false.
//
// [EnumWindows]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumwindows
func EnumWindows(callback func(hWnd HWND) bool) {
	pPack := &_EnumWindowsPack{f: callback}
	_globalEnumWindowsMutex.Lock()
	if _globalEnumWindowsFuncs == nil { // the set was not initialized yet?
		_globalEnumWindowsFuncs = make(map[*_EnumWindowsPack]struct{}, 1)
	}
	_globalEnumWindowsFuncs[pPack] = struct{}{} // store pointer in the set
	_globalEnumWindowsMutex.Unlock()

	ret, _, err := syscall.SyscallN(proc.EnumWindows.Addr(),
		_globalEnumWindowsCallback, uintptr(unsafe.Pointer(pPack)))

	_globalEnumWindowsMutex.Lock()
	delete(_globalEnumWindowsFuncs, pPack) // remove from the set
	_globalEnumWindowsMutex.Unlock()

	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

type _EnumWindowsPack struct{ f func(hWnd HWND) bool }

var (
	_globalEnumWindowsFuncs    map[*_EnumWindowsPack]struct{} // keeps pointers from being collected by GC
	_globalEnumWindowsMutex    = sync.Mutex{}
	_globalEnumWindowsCallback = syscall.NewCallback(
		func(hWnd HWND, lParam LPARAM) uintptr {
			pPack := (*_EnumWindowsPack)(unsafe.Pointer(lParam))
			return util.BoolToUintptr(pPack.f(hWnd))
		})
)

// [GetAsyncKeyState] function.
//
// [GetAsyncKeyState]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getasynckeystate
func GetAsyncKeyState(virtKeyCode co.VK) uint16 {
	ret, _, _ := syscall.SyscallN(proc.GetAsyncKeyState.Addr(),
		uintptr(virtKeyCode))
	return uint16(ret)
}

// [GetCaretPos] function.
//
// [GetCaretPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcaretpos
func GetCaretPos() RECT {
	var rc RECT
	ret, _, err := syscall.SyscallN(proc.GetCaretPos.Addr(),
		uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return rc
}

// [GetCursorPos] function.
//
// [GetCursorPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
func GetCursorPos() POINT {
	var pt POINT
	ret, _, err := syscall.SyscallN(proc.GetCursorPos.Addr(),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// [GetDialogBaseUnits] function.
//
// [GetDialogBaseUnits]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdialogbaseunits
func GetDialogBaseUnits() (horz, vert uint16) {
	ret, _, _ := syscall.SyscallN(proc.GetDialogBaseUnits.Addr())
	horz, vert = LOWORD(uint32(ret)), HIWORD(uint32(ret))
	return
}

// [GetGUIThreadInfo] function.
//
// [GetGUIThreadInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getguithreadinfo
func GetGUIThreadInfo(thread_id uint32, info *GUITHREADINFO) {
	info.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(proc.GetGUIThreadInfo.Addr(),
		uintptr(thread_id), uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [GetInputState] function.
//
// [GetInputState]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getinputstate
func GetInputState() bool {
	ret, _, _ := syscall.SyscallN(proc.GetInputState.Addr())
	return ret != 0
}

// [GetMessage] function.
//
// [GetMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagew
func GetMessage(
	msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax uint32) (int32, error) {

	ret, _, err := syscall.SyscallN(proc.GetMessage.Addr(),
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax))
	if int(ret) == -1 {
		return 0, errco.ERROR(err)
	}
	return int32(ret), nil
}

// [GetMessageExtraInfo] function.
//
// [GetMessageExtraInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessageextrainfo
func GetMessageExtraInfo() LPARAM {
	ret, _, _ := syscall.SyscallN(proc.GetMessageExtraInfo.Addr())
	return LPARAM(ret)
}

// [GetMessagePos] function.
//
// [GetMessagePos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagepos
func GetMessagePos() POINT {
	ret, _, _ := syscall.SyscallN(proc.GetMessagePos.Addr())
	return POINT{
		X: int32(LOWORD(uint32(ret))),
		Y: int32(HIWORD(uint32(ret))),
	}
}

// [GetMessageTime] function.
//
// [GetMessageTime]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagetime
func GetMessageTime() time.Duration {
	ret, _, _ := syscall.SyscallN(proc.GetMessageTime.Addr())
	return time.Duration(ret * uintptr(time.Millisecond))
}

// [GetPhysicalCursorPos] function.
//
// [GetPhysicalCursorPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getphysicalcursorpos
func GetPhysicalCursorPos() POINT {
	var pt POINT
	ret, _, err := syscall.SyscallN(proc.GetPhysicalCursorPos.Addr(),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// [GetProcessDefaultLayout] function.
//
// [GetProcessDefaultLayout]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getprocessdefaultlayout
func GetProcessDefaultLayout() co.LAYOUT {
	var defaultLayout co.LAYOUT
	ret, _, err := syscall.SyscallN(proc.GetProcessDefaultLayout.Addr(),
		uintptr(unsafe.Pointer(&defaultLayout)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return defaultLayout
}

// [GetQueueStatus] function.
//
// [GetQueueStatus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getqueuestatus
func GetQueueStatus(flags co.QS) uint32 {
	ret, _, _ := syscall.SyscallN(proc.GetQueueStatus.Addr(),
		uintptr(flags))
	return uint32(ret)
}

// [GetSysColor] function.
//
// [GetSysColor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
func GetSysColor(index co.COLOR) COLORREF {
	ret, _, _ := syscall.SyscallN(proc.GetSysColor.Addr(),
		uintptr(index))
	return COLORREF(ret)
}

// [GetSystemMetrics] function.
//
// [GetSystemMetrics]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
func GetSystemMetrics(index co.SM) int32 {
	ret, _, _ := syscall.SyscallN(proc.GetSystemMetrics.Addr(),
		uintptr(index))
	return int32(ret)
}

// [GetSystemMetricsForDpi] function.
//
// Available in Windows 10, version 1607.
//
// [GetSystemMetricsForDpi]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetricsfordpi
func GetSystemMetricsForDpi(index co.SM, dpi uint32) int32 {
	ret, _, err := syscall.SyscallN(proc.GetSystemMetricsForDpi.Addr(),
		uintptr(index))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return int32(ret)
}

// [InSendMessage] function.
//
// [InSendMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessage
func InSendMessage() bool {
	ret, _, _ := syscall.SyscallN(proc.InSendMessage.Addr())
	return ret != 0
}

// [InSendMessageEx] function.
//
// [InSendMessageEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessageex
func InSendMessageEx() co.ISMEX {
	ret, _, _ := syscall.SyscallN(proc.InSendMessageEx.Addr())
	return co.ISMEX(ret)
}

// [IsGUIThread] function.
//
// [IsGUIThread]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isguithread
func IsGUIThread(convertToGuiThread bool) (bool, error) {
	ret, _, _ := syscall.SyscallN(proc.IsGUIThread.Addr(),
		util.BoolToUintptr(convertToGuiThread))
	if convertToGuiThread && errco.ERROR(ret) == errco.NOT_ENOUGH_MEMORY {
		return false, errco.NOT_ENOUGH_MEMORY
	}
	return ret != 0, nil
}

// [LockSetForegroundWindow] function.
//
// [LockSetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-locksetforegroundwindow
func LockSetForegroundWindow(lockCode co.LSFW) {
	ret, _, err := syscall.SyscallN(proc.LockSetForegroundWindow.Addr(),
		uintptr(lockCode))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [PeekMessage] function.
//
// [PeekMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-peekmessagew
func PeekMessage(
	msg *MSG, hWnd HWND,
	msgFilterMin, msgFilterMax co.WM, removeMsg co.PM) bool {

	ret, _, _ := syscall.SyscallN(proc.PeekMessage.Addr(),
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax), uintptr(removeMsg))
	return ret != 0
}

// [PostQuitMessage] function.
//
// [PostQuitMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int32) {
	syscall.SyscallN(proc.PostQuitMessage.Addr(),
		uintptr(exitCode))
}

// [PostThreadMessage] function.
//
// [PostThreadMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessage(
	idThread uint32, msg co.WM, wParam WPARAM, lParam LPARAM) error {

	ret, _, err := syscall.SyscallN(proc.PostThreadMessage.Addr(),
		uintptr(idThread), uintptr(msg), uintptr(wParam), uintptr(lParam))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [RegisterClassEx] function.
//
// [RegisterClassEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wcx *WNDCLASSEX) (ATOM, error) {
	wcx.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(proc.RegisterClassEx.Addr(),
		uintptr(unsafe.Pointer(wcx)))

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return ATOM(0), wErr
	} else {
		return ATOM(ret), nil
	}
}

// [RegisterWindowMessage] function.
//
// [RegisterWindowMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessage(message string) (co.WM, error) {
	ret, _, err := syscall.SyscallN(proc.RegisterWindowMessage.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(message))))

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return co.WM(0), wErr
	} else {
		return co.WM(ret), nil
	}
}

// [ReplyMessage] function.
//
// [ReplyMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-replymessage
func ReplyMessage(result uintptr) bool {
	ret, _, _ := syscall.SyscallN(proc.ReplyMessage.Addr(),
		result)
	return ret != 0
}

// [TranslateMessage] function.
//
// [TranslateMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.SyscallN(proc.TranslateMessage.Addr(),
		uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

// [SetProcessDpiAwarenessContext] function.
//
// Available in Windows 10, version 1703.
//
// [SetProcessDpiAwarenessContext]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiawarenesscontext
func SetProcessDpiAwarenessContext(value co.DPI_AWARE_CTX) error {
	ret, _, err := syscall.SyscallN(proc.SetProcessDpiAwarenessContext.Addr(),
		uintptr(value))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [SetProcessDPIAware] function.
//
// Available in Windows Vista.
//
// [SetProcessDPIAware]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiaware
func SetProcessDPIAware() {
	ret, _, _ := syscall.SyscallN(proc.SetProcessDPIAware.Addr())
	if ret == 0 {
		panic("SetProcessDPIAware() failed.")
	}
}

// [SetMessageExtraInfo] function.
//
// [SetMessageExtraInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmessageextrainfo
func SetMessageExtraInfo(lp LPARAM) LPARAM {
	ret, _, _ := syscall.SyscallN(proc.SetMessageExtraInfo.Addr(),
		uintptr(lp))
	return LPARAM(ret)
}

// [SetProcessDefaultLayout] function.
//
// [SetProcessDefaultLayout]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdefaultlayout
func SetProcessDefaultLayout(defaultLayout co.LAYOUT) {
	ret, _, err := syscall.SyscallN(proc.SetProcessDefaultLayout.Addr(),
		uintptr(defaultLayout))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [UnregisterClass] function.
//
// [UnregisterClass]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unregisterclassw
func UnregisterClass(className ClassName, hInst HINSTANCE) error {
	classNameVal, classNameBuf := className.raw()
	ret, _, err := syscall.SyscallN(proc.UnregisterClass.Addr(),
		classNameVal, uintptr(hInst))
	runtime.KeepAlive(classNameBuf)

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return wErr
	} else {
		return nil
	}
}
