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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-adjustwindowrectex
func AdjustWindowRectEx(rc *RECT, style co.WS, hasMenu bool, exStyle co.WS_EX) {
	ret, _, err := syscall.Syscall6(proc.AdjustWindowRectEx.Addr(), 4,
		uintptr(unsafe.Pointer(rc)), uintptr(style),
		util.BoolToUintptr(hasMenu), uintptr(exStyle), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-allowsetforegroundwindow
func AllowSetForegroundWindow(processId uint32) {
	ret, _, err := syscall.Syscall(proc.AllowSetForegroundWindow.Addr(), 1,
		uintptr(processId), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-broadcastsystemmessagew
func BroadcastSystemMessage(
	flags co.BSF,
	recipients co.BSM,
	msg co.WM,
	wParam WPARAM,
	lParam LPARAM) (broadcastSuccessful bool, receivers co.BSM, e error) {

	receivers = recipients

	ret, _, err := syscall.Syscall6(proc.BroadcastSystemMessage.Addr(), 5,
		uintptr(flags), uintptr(unsafe.Pointer(&receivers)),
		uintptr(msg), uintptr(wParam), uintptr(lParam), 0)

	broadcastSuccessful = int(ret) > 1
	if ret == 0 {
		e = errco.ERROR(err)
	}
	return
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endmenu
func EndMenu() {
	ret, _, err := syscall.Syscall(proc.EndMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// To continue enumeration, the callback function must return true; to stop
// enumeration, it must return false.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaydevicesw
func EnumDisplayDevices(
	device StrOpt, flags co.EDD,
	callback func(devNum int, info *DISPLAY_DEVICE) bool) {

	devicePtr := device.Raw()
	devNum := 0

	dide := DISPLAY_DEVICE{}
	dide.SetCb()

	for {
		ret, _, err := syscall.Syscall6(proc.EnumDisplayDevices.Addr(), 4,
			uintptr(devicePtr), uintptr(devNum),
			uintptr(unsafe.Pointer(&dide)), uintptr(flags),
			0, 0)

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

// To continue enumeration, the callback function must return true; to stop
// enumeration, it must return false.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumwindows
func EnumWindows(callback func(hWnd HWND) bool) {
	pPack := &_EnumWindowsPack{f: callback}
	_globalEnumWindowsMutex.Lock()
	if _globalEnumWindowsFuncs == nil { // the set was not initialized yet?
		_globalEnumWindowsFuncs = make(map[*_EnumWindowsPack]struct{}, 1)
	}
	_globalEnumWindowsFuncs[pPack] = struct{}{} // store pointer in the set
	_globalEnumWindowsMutex.Unlock()

	ret, _, err := syscall.Syscall(proc.EnumWindows.Addr(), 2,
		_globalEnumWindowsCallback, uintptr(unsafe.Pointer(pPack)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

type _EnumWindowsPack struct{ f func(hWnd HWND) bool }

var (
	_globalEnumWindowsCallback uintptr = syscall.NewCallback(_EnumWindowsProc)
	_globalEnumWindowsFuncs    map[*_EnumWindowsPack]struct{}
	_globalEnumWindowsMutex    = sync.Mutex{}
)

func _EnumWindowsProc(hWnd HWND, lParam LPARAM) uintptr {
	pPack := (*_EnumWindowsPack)(unsafe.Pointer(lParam))
	retVal := uintptr(0)

	_globalEnumWindowsMutex.Lock()
	_, isStored := _globalEnumWindowsFuncs[pPack]
	_globalEnumWindowsMutex.Unlock()

	if isStored {
		retVal = util.BoolToUintptr(pPack.f(hWnd))
		if retVal == 0 {
			_globalEnumWindowsMutex.Lock()
			delete(_globalEnumWindowsFuncs, pPack) // remove from the set
			_globalEnumWindowsMutex.Unlock()
		}
	}
	return retVal
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getasynckeystate
func GetAsyncKeyState(virtKeyCode co.VK) uint16 {
	ret, _, _ := syscall.Syscall(proc.GetAsyncKeyState.Addr(), 1,
		uintptr(virtKeyCode), 0, 0)
	return uint16(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcaretpos
func GetCaretPos() RECT {
	var rc RECT
	ret, _, err := syscall.Syscall(proc.GetCaretPos.Addr(), 1,
		uintptr(unsafe.Pointer(&rc)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return rc
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
func GetCursorPos() POINT {
	var pt POINT
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
	horz, vert = LOWORD(uint32(ret)), HIWORD(uint32(ret))
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getguithreadinfo
func GetGUIThreadInfo(thread_id uint32, info *GUITHREADINFO) {
	info.SetCbSize() // safety
	ret, _, err := syscall.Syscall(proc.GetGUIThreadInfo.Addr(), 2,
		uintptr(thread_id), uintptr(unsafe.Pointer(info)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getinputstate
func GetInputState() bool {
	ret, _, _ := syscall.Syscall(proc.GetInputState.Addr(), 0,
		0, 0, 0)
	return ret != 0
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessageextrainfo
func GetMessageExtraInfo() LPARAM {
	ret, _, _ := syscall.Syscall(proc.GetMessageExtraInfo.Addr(), 0,
		0, 0, 0)
	return LPARAM(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagepos
func GetMessagePos() POINT {
	ret, _, _ := syscall.Syscall(proc.GetMessagePos.Addr(), 0,
		0, 0, 0)
	return POINT{
		X: int32(LOWORD(uint32(ret))),
		Y: int32(HIWORD(uint32(ret))),
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
	var pt POINT
	ret, _, err := syscall.Syscall(proc.GetPhysicalCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(&pt)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getprocessdefaultlayout
func GetProcessDefaultLayout() co.LAYOUT {
	var defaultLayout co.LAYOUT
	ret, _, err := syscall.Syscall(proc.GetProcessDefaultLayout.Addr(), 1,
		uintptr(unsafe.Pointer(&defaultLayout)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return defaultLayout
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getqueuestatus
func GetQueueStatus(flags co.QS) uint32 {
	ret, _, _ := syscall.Syscall(proc.GetQueueStatus.Addr(), 1,
		uintptr(flags), 0, 0)
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
func GetSysColor(index co.COLOR) COLORREF {
	ret, _, _ := syscall.Syscall(proc.GetSysColor.Addr(), 1,
		uintptr(index), 0, 0)
	return COLORREF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
func GetSystemMetrics(index co.SM) int32 {
	ret, _, _ := syscall.Syscall(proc.GetSystemMetrics.Addr(), 1,
		uintptr(index), 0, 0)
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessage
func InSendMessage() bool {
	ret, _, _ := syscall.Syscall(proc.InSendMessage.Addr(), 0,
		0, 0, 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessageex
func InSendMessageEx() co.ISMEX {
	ret, _, _ := syscall.Syscall(proc.InSendMessageEx.Addr(), 0,
		0, 0, 0)
	return co.ISMEX(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isguithread
func IsGUIThread(convertToGuiThread bool) (bool, error) {
	ret, _, _ := syscall.Syscall(proc.IsGUIThread.Addr(), 1,
		util.BoolToUintptr(convertToGuiThread), 0, 0)
	if convertToGuiThread && errco.ERROR(ret) == errco.NOT_ENOUGH_MEMORY {
		return false, errco.NOT_ENOUGH_MEMORY
	}
	return ret != 0, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-locksetforegroundwindow
func LockSetForegroundWindow(lockCode co.LSFW) {
	ret, _, err := syscall.Syscall(proc.LockSetForegroundWindow.Addr(), 1,
		uintptr(lockCode), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-peekmessagew
func PeekMessage(
	msg *MSG, hWnd HWND,
	msgFilterMin, msgFilterMax co.WM, removeMsg co.PM) bool {

	ret, _, _ := syscall.Syscall6(proc.PeekMessage.Addr(), 5,
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax), uintptr(removeMsg), 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int32) {
	syscall.Syscall(proc.PostQuitMessage.Addr(), 1,
		uintptr(exitCode), 0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessage(
	idThread uint32, msg co.WM, wParam WPARAM, lParam LPARAM) error {

	ret, _, err := syscall.Syscall6(proc.PostThreadMessage.Addr(), 4,
		uintptr(idThread), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wcx *WNDCLASSEX) (ATOM, error) {
	wcx.SetCbSize() // safety
	ret, _, err := syscall.Syscall(proc.RegisterClassEx.Addr(), 1,
		uintptr(unsafe.Pointer(wcx)), 0, 0)

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return ATOM(0), wErr
	} else {
		return ATOM(ret), nil
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessage(message string) (co.WM, error) {
	ret, _, err := syscall.Syscall(proc.RegisterWindowMessage.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToNativePtr(message))), 0, 0)

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return co.WM(0), wErr
	} else {
		return co.WM(ret), nil
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-replymessage
func ReplyMessage(result uintptr) bool {
	ret, _, _ := syscall.Syscall(proc.ReplyMessage.Addr(), 1,
		result, 0, 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(proc.TranslateMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret != 0
}

// Available in Windows 10, version 1703.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiawarenesscontext
func SetProcessDpiAwarenessContext(value co.DPI_AWARE_CTX) error {
	ret, _, err := syscall.Syscall(proc.SetProcessDpiAwarenessContext.Addr(), 1,
		uintptr(value), 0, 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// Available in Windows Vista.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiaware
func SetProcessDPIAware() {
	ret, _, _ := syscall.Syscall(proc.SetProcessDPIAware.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic("SetProcessDPIAware() failed.")
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmessageextrainfo
func SetMessageExtraInfo(lp LPARAM) LPARAM {
	ret, _, _ := syscall.Syscall(proc.SetMessageExtraInfo.Addr(), 1,
		uintptr(lp), 0, 0)
	return LPARAM(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdefaultlayout
func SetProcessDefaultLayout(defaultLayout co.LAYOUT) {
	ret, _, err := syscall.Syscall(proc.SetProcessDefaultLayout.Addr(), 1,
		uintptr(defaultLayout), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unregisterclassw
func UnregisterClass(className ClassName, hInst HINSTANCE) error {
	classNameVal, classNameBuf := className.raw()
	ret, _, err := syscall.Syscall(proc.UnregisterClass.Addr(), 2,
		classNameVal, uintptr(hInst), 0)
	runtime.KeepAlive(classNameBuf)

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return wErr
	} else {
		return nil
	}
}
