package api

import (
	"syscall"
	"unsafe"
	"winffi/consts"
	"winffi/procs"
)

type HWND HANDLE

func (hwnd HWND) ClientToScreen(point *POINT) {
	syscall.Syscall(procs.ClientToScreen.Addr(), 2,
		uintptr(hwnd), uintptr(unsafe.Pointer(point)), 0)
}

func CreateWindowEx(exStyle consts.WS_EX, className, title string,
	style consts.WS, x, y, width, height int32, parent HWND, menu HMENU,
	instance HINSTANCE, param unsafe.Pointer) (HWND, syscall.Errno) {

	ret, _, errno := syscall.Syscall12(procs.CreateWindowEx.Addr(), 12,
		uintptr(exStyle), toUtf16BlankIsNilToUintptr(className),
		toUtf16BlankIsNilToUintptr(title), uintptr(style),
		uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(parent), uintptr(menu), uintptr(instance), uintptr(param))
	return HWND(ret), errno
}

func (hwnd HWND) DefWindowProc(msg consts.WM, wParam WPARAM,
	lParam LPARAM) uintptr {

	ret, _, _ := syscall.Syscall6(procs.DefWindowProc.Addr(), 4,
		uintptr(hwnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

func (hwnd HWND) DestroyWindow() syscall.Errno {
	_, _, errno := syscall.Syscall(procs.DestroyWindow.Addr(), 1,
		uintptr(hwnd), 0, 0)
	return errno
}

func (hwnd HWND) EnumChildWindows() []HWND {
	hChildren := []HWND{}
	syscall.Syscall(procs.EnumChildWindows.Addr(), 3,
		uintptr(hwnd),
		syscall.NewCallback(func(hChild HWND, lp LPARAM) uintptr {
			hChildren = append(hChildren, hChild)
			return uintptr(1)
		}), 0)
	return hChildren
}

func GetForegroundWindow() HWND {
	ret, _, _ := syscall.Syscall(procs.GetForegroundWindow.Addr(), 0,
		0, 0, 0)
	return HWND(ret)
}

func (hwnd HWND) GetParent() HWND {
	ret, _, _ := syscall.Syscall(procs.GetParent.Addr(), 1,
		uintptr(hwnd), 0, 0)
	return HWND(ret)
}

func (hwnd HWND) GetWindowDC() HDC {
	ret, _, _ := syscall.Syscall(procs.GetWindowDC.Addr(), 1,
		uintptr(hwnd), 0, 0)
	return HDC(ret)
}

func (hwnd HWND) IsDialogMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(procs.IsDialogMessage.Addr(), 2,
		uintptr(hwnd), uintptr(unsafe.Pointer(msg)), 0)
	return ret != 0
}

func (hwnd HWND) MessageBox(message, caption string, flags consts.MB) int32 {
	ret, _, _ := syscall.Syscall6(procs.MessageBox.Addr(), 4,
		uintptr(0), toUtf16ToUintptr(message), toUtf16ToUintptr(caption),
		uintptr(flags), 0, 0)
	return int32(ret)
}
