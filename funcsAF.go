package winffi

import (
	"syscall"
	"unsafe"
	"winffi/consts"
	"winffi/procs"
)

func (hWnd HWND) ClientToScreen(point *POINT) {
	syscall.Syscall(procs.ClientToScreen.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(point)), 0)
}

func (hAccel HACCEL) CopyAcceleratorTable() []ACCEL {
	szRet, _, _ := syscall.Syscall(procs.CopyAcceleratorTable.Addr(), 3,
		uintptr(hAccel), 0, 0)
	if szRet == 0 {
		return []ACCEL{}
	}

	accelList := make([]ACCEL, uint32(szRet))
	syscall.Syscall(procs.CopyAcceleratorTable.Addr(), 3,
		uintptr(hAccel), uintptr(unsafe.Pointer(&accelList[0])), szRet)
	return accelList
}

func CreateAcceleratorTable(accelList []ACCEL) (HACCEL, syscall.Errno) {
	ret, _, errno := syscall.Syscall(procs.CreateAcceleratorTable.Addr(), 2,
		uintptr(unsafe.Pointer(&accelList[0])), uintptr(len(accelList)),
		0)
	return HACCEL(ret), errno
}

func (lf *LOGFONT) CreateFontIndirect() HFONT {
	ret, _, _ := syscall.Syscall(procs.CreateFontIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lf)), 0, 0)
	return HFONT(ret)
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

func (hWnd HWND) DefWindowProc(msg consts.WM, wParam WPARAM,
	lParam LPARAM) uintptr {

	ret, _, _ := syscall.Syscall6(procs.DefWindowProc.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

func (hBrush HBRUSH) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(procs.DeleteObject.Addr(), 1,
		uintptr(hBrush), 0, 0)
	return ret != 0
}

func (hFont HFONT) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(procs.DeleteObject.Addr(), 1,
		uintptr(hFont), 0, 0)
	return ret != 0
}

func (hObject HGDIOBJ) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(procs.DeleteObject.Addr(), 1,
		uintptr(hObject), 0, 0)
	return ret != 0
}

func (hPen HPEN) DeleteObject() bool {
	ret, _, _ := syscall.Syscall(procs.DeleteObject.Addr(), 1,
		uintptr(hPen), 0, 0)
	return ret != 0
}

func (hAccel HACCEL) DestroyAcceleratorTable() bool {
	ret, _, _ := syscall.Syscall(procs.DestroyAcceleratorTable.Addr(), 1,
		uintptr(hAccel), 0, 0)
	return ret != 0
}

func (hWnd HWND) DestroyWindow() syscall.Errno {
	_, _, errno := syscall.Syscall(procs.DestroyWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return errno
}

func (msg *MSG) DispatchMessage() uintptr {
	ret, _, _ := syscall.Syscall(procs.DispatchMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret
}

func (hWnd HWND) EnumChildWindows() []HWND {
	hChildren := []HWND{}
	syscall.Syscall(procs.EnumChildWindows.Addr(), 3,
		uintptr(hWnd),
		syscall.NewCallback(func(hChild HWND, lp LPARAM) uintptr {
			hChildren = append(hChildren, hChild)
			return uintptr(1)
		}), 0)
	return hChildren
}

func EnumDisplayMonitors(hdc HDC, rcClip *RECT) []HMONITOR {
	hMons := []HMONITOR{}
	syscall.Syscall6(procs.EnumDisplayMonitors.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(rcClip)),
		syscall.NewCallback(func(hMon HMONITOR, hdcMon HDC, rcMon *RECT,
			lp LPARAM) uintptr {
			hMons = append(hMons, hMon)
			return uintptr(1)
		}), 0, 0, 0)
	return hMons
}
