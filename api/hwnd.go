package api

import (
	"log"
	"syscall"
	"unsafe"
	"winffi/api/proc"
	c "winffi/consts"
)

type HWND HANDLE

func (hwnd HWND) ClientToScreen(point *POINT) {
	ret, _, _ := syscall.Syscall(proc.ClientToScreen.Addr(), 2,
		uintptr(hwnd), uintptr(unsafe.Pointer(point)), 0)
	if ret == 0 {
		panic("ClientToScreen failed.")
	}
}

func CreateWindowEx(exStyle c.WS_EX, className, title string, style c.WS,
	x, y int32, width, height uint32, parent HWND, menu HMENU,
	instance HINSTANCE, param unsafe.Pointer) HWND {

	ret, _, errno := syscall.Syscall12(proc.CreateWindowEx.Addr(), 12,
		uintptr(exStyle),
		uintptr(unsafe.Pointer(toUtf16PtrBlankIsNil(className))),
		uintptr(unsafe.Pointer(toUtf16PtrBlankIsNil(title))),
		uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(parent), uintptr(menu), uintptr(instance), uintptr(param))
	if ret == 0 {
		log.Panicf("CreateWindowEx failed: %d %s", errno, errno.Error())
	}
	return HWND(ret)
}

func (hwnd HWND) DefWindowProc(msg c.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.Syscall6(proc.DefWindowProc.Addr(), 4,
		uintptr(hwnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

func (hwnd HWND) DestroyWindow() {
	ret, _, errno := syscall.Syscall(proc.DestroyWindow.Addr(), 1,
		uintptr(hwnd), 0, 0)
	if ret == 0 {
		log.Panicf("DestroyWindow failed: %d %s", errno, errno.Error())
	}
}

func (hwnd HWND) EnumChildWindows() []HWND {
	hChildren := make([]HWND, 0)
	syscall.Syscall(proc.EnumChildWindows.Addr(), 3,
		uintptr(hwnd),
		syscall.NewCallback(
			func(hChild HWND, lp LPARAM) uintptr {
				hChildren = append(hChildren, hChild)
				return boolToUintptr(true)
			}), 0)
	return hChildren
}

func (hwnd HWND) GetExStyle() c.WS_EX {
	return c.WS_EX(hwnd.GetWindowLongPtr(c.GWLP_EXSTYLE))
}

func GetForegroundWindow() HWND {
	ret, _, _ := syscall.Syscall(proc.GetForegroundWindow.Addr(), 0,
		0, 0, 0)
	return HWND(ret)
}

func (hwnd HWND) GetInstance() HINSTANCE {
	return HINSTANCE(hwnd.GetWindowLongPtr(c.GWLP_HINSTANCE))
}

func (hwnd HWND) GetParent() HWND {
	ret, _, errno := syscall.Syscall(proc.GetParent.Addr(), 1,
		uintptr(hwnd), 0, 0)
	if ret == 0 {
		log.Panicf("GetParent failed: %d %s", errno, errno.Error())
	}
	return HWND(ret)
}

func (hwnd HWND) GetStyle() c.WS {
	return c.WS(hwnd.GetWindowLongPtr(c.GWLP_STYLE))
}

func (hwnd HWND) GetWindowDC() HDC {
	ret, _, _ := syscall.Syscall(proc.GetWindowDC.Addr(), 1,
		uintptr(hwnd), 0, 0)
	if ret == 0 {
		panic("GetWindowDC failed.")
	}
	return HDC(ret)
}

func (hwnd HWND) GetWindowLongPtr(index c.GWLP) uintptr {
	ret, _, _ := syscall.Syscall(proc.GetWindowLongPtr.Addr(), 2,
		uintptr(hwnd), uintptr(index),
		0)
	return ret
}

func (hwnd HWND) GetWindowText() string {
	len := hwnd.GetWindowTextLength() + 1
	buf := make([]uint16, len)
	ret, _, errno := syscall.Syscall(proc.GetWindowText.Addr(), 3,
		uintptr(hwnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(len))
	if ret == 0 && errno != 0 {
		log.Panicf("GetWindowText failed: %d %s", errno, errno.Error())
	}
	return syscall.UTF16ToString(buf)
}

func (hwnd HWND) GetWindowTextLength() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetWindowTextLength.Addr(), 1,
		uintptr(hwnd), 0, 0)
	return uint32(ret)
}

func (hwnd HWND) IsDialogMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(proc.IsDialogMessage.Addr(), 2,
		uintptr(hwnd), uintptr(unsafe.Pointer(msg)), 0)
	return ret != 0
}

func (hwnd HWND) MessageBox(message, caption string, flags c.MB) c.ID {
	ret, _, _ := syscall.Syscall6(proc.MessageBox.Addr(), 4,
		uintptr(0),
		uintptr(unsafe.Pointer(toUtf16Ptr(message))),
		uintptr(unsafe.Pointer(toUtf16Ptr(caption))),
		uintptr(flags), 0, 0)
	return c.ID(ret)
}

func (hwnd HWND) ReleaseDC(hdc HDC) int32 {
	ret, _, _ := syscall.Syscall(proc.ReleaseDC.Addr(), 2,
		uintptr(hwnd), uintptr(hdc), 0)
	return int32(ret)
}

func (hwnd HWND) ScreenToClientPoint(point *POINT) {
	syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hwnd), uintptr(unsafe.Pointer(point)), 0)
}

func (hwnd HWND) ScreenToClientRect(rect *RECT) {
	syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hwnd), uintptr(unsafe.Pointer(rect)), 0)
	syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hwnd), uintptr(unsafe.Pointer(&rect.Right)), 0)
}

func (hwnd HWND) SendMessage(msg c.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.Syscall6(proc.SendMessage.Addr(), 4,
		uintptr(hwnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

func (hwnd HWND) SetWindowLongPtr(index c.GWLP, newLong uintptr) uintptr {
	ret, _, _ := syscall.Syscall(proc.SetWindowLongPtr.Addr(), 3,
		uintptr(hwnd), uintptr(index), newLong)
	return ret
}

func (hwnd HWND) ShowWindow(nCmdShow c.SW) bool {
	ret, _, _ := syscall.Syscall(proc.ShowWindow.Addr(), 1,
		uintptr(hwnd), uintptr(nCmdShow), 0)
	return ret != 0
}

func (hwnd HWND) TranslateAccelerator(hAccel HACCEL,
	msg *MSG) (int32, syscall.Errno) {

	ret, _, errno := syscall.Syscall(proc.TranslateAccelerator.Addr(), 3,
		uintptr(hwnd), uintptr(hAccel), uintptr(unsafe.Pointer(msg)))
	return int32(ret), errno
}

func (hwnd HWND) SetWindowText(text string) {
	syscall.Syscall(proc.SetWindowText.Addr(), 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(toUtf16Ptr(text))),
		0)
}

func (hwnd HWND) UpdateWindow() bool {
	ret, _, _ := syscall.Syscall(proc.UpdateWindow.Addr(), 1,
		uintptr(hwnd), 0, 0)
	return ret != 0
}
