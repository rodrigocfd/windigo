package api

import (
	"syscall"
	"unsafe"
	c "winffi/consts"
	p "winffi/procs"
)

type HWND HANDLE

func (hwnd HWND) ClientToScreen(point *POINT) {
	syscall.Syscall(p.ClientToScreen.Addr(), 2,
		uintptr(hwnd), uintptr(unsafe.Pointer(point)), 0)
}

func CreateWindowEx(exStyle c.WS_EX, className, title string,
	style c.WS, x, y, width, height int32, parent HWND, menu HMENU,
	instance HINSTANCE, param unsafe.Pointer) (HWND, syscall.Errno) {

	ret, _, errno := syscall.Syscall12(p.CreateWindowEx.Addr(), 12,
		uintptr(exStyle), toUtf16BlankIsNilToUintptr(className),
		toUtf16BlankIsNilToUintptr(title), uintptr(style),
		uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(parent), uintptr(menu), uintptr(instance), uintptr(param))
	return HWND(ret), errno
}

func (hwnd HWND) DefWindowProc(msg c.WM, wParam WPARAM,
	lParam LPARAM) uintptr {

	ret, _, _ := syscall.Syscall6(p.DefWindowProc.Addr(), 4,
		uintptr(hwnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

func (hwnd HWND) DestroyWindow() syscall.Errno {
	_, _, errno := syscall.Syscall(p.DestroyWindow.Addr(), 1,
		uintptr(hwnd), 0, 0)
	return errno
}

func (hwnd HWND) EnumChildWindows() []HWND {
	hChildren := []HWND{}
	syscall.Syscall(p.EnumChildWindows.Addr(), 3,
		uintptr(hwnd),
		syscall.NewCallback(func(hChild HWND, lp LPARAM) uintptr {
			hChildren = append(hChildren, hChild)
			return uintptr(1)
		}), 0)
	return hChildren
}

func GetForegroundWindow() HWND {
	ret, _, _ := syscall.Syscall(p.GetForegroundWindow.Addr(), 0,
		0, 0, 0)
	return HWND(ret)
}

func (hwnd HWND) GetParent() HWND {
	ret, _, _ := syscall.Syscall(p.GetParent.Addr(), 1,
		uintptr(hwnd), 0, 0)
	return HWND(ret)
}

func (hwnd HWND) GetWindowDC() HDC {
	ret, _, _ := syscall.Syscall(p.GetWindowDC.Addr(), 1,
		uintptr(hwnd), 0, 0)
	return HDC(ret)
}

func (hwnd HWND) IsDialogMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(p.IsDialogMessage.Addr(), 2,
		uintptr(hwnd), uintptr(unsafe.Pointer(msg)), 0)
	return ret != 0
}

func (hwnd HWND) MessageBox(message, caption string, flags c.MB) int32 {
	ret, _, _ := syscall.Syscall6(p.MessageBox.Addr(), 4,
		uintptr(0), toUtf16ToUintptr(message), toUtf16ToUintptr(caption),
		uintptr(flags), 0, 0)
	return int32(ret)
}
