/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type HWND HANDLE

func (hWnd HWND) ClientToScreenPt(point *POINT) {
	ret, _, _ := syscall.Syscall(proc.ClientToScreen.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(point)), 0)
	if ret == 0 {
		panic("ClientToScreen failed for POINT.")
	}
}

func (hWnd HWND) ClientToScreenRc(rect *RECT) {
	ret, _, _ := syscall.Syscall(proc.ClientToScreen.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(rect)), 0)
	if ret == 0 {
		panic("ClientToScreen failed for RECT (1).")
	}
	ret, _, _ = syscall.Syscall(proc.ClientToScreen.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(&rect.Right)), 0)
	if ret == 0 {
		panic("ClientToScreen failed for RECT (2).")
	}
}

func CreateWindowEx(exStyle co.WS_EX, className, title string, style co.WS,
	x, y int32, width, height uint32, parent HWND, menu HMENU,
	instance HINSTANCE, param unsafe.Pointer) HWND {

	ret, _, lerr := syscall.Syscall12(proc.CreateWindowEx.Addr(), 12,
		uintptr(exStyle),
		uintptr(unsafe.Pointer(StrToPtrBlankIsNil(className))),
		uintptr(unsafe.Pointer(StrToPtrBlankIsNil(title))),
		uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(parent), uintptr(menu), uintptr(instance), uintptr(param))

	if ret == 0 {
		panic(co.ERROR(lerr).Format(
			fmt.Sprintf("CreateWindowEx failed %s.", className),
		))
	}

	return HWND(ret)
}

func (hWnd HWND) DefSubclassProc(msg co.WM,
	wParam WPARAM, lParam LPARAM) uintptr {

	ret, _, _ := syscall.Syscall6(proc.DefSubclassProc.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

func (hWnd HWND) DefWindowProc(msg co.WM,
	wParam WPARAM, lParam LPARAM) uintptr {

	ret, _, _ := syscall.Syscall6(proc.DefWindowProc.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

func (hWnd HWND) DestroyWindow() {
	ret, _, lerr := syscall.Syscall(proc.DestroyWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("DestroyWindow failed."))
	}
}

func (hWnd HWND) DragAcceptFiles(fAccept bool) {
	syscall.Syscall(proc.DragAcceptFiles.Addr(), 2,
		uintptr(hWnd), boolToUintptr(fAccept), 0)
}

func (hWnd HWND) DrawMenuBar() {
	ret, _, lerr := syscall.Syscall(proc.DrawMenuBar.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("DrawMenuBar failed."))
	}
}

func (hWnd HWND) EnableWindow(bEnable bool) bool {
	ret, _, _ := syscall.Syscall(proc.EnableWindow.Addr(), 2,
		uintptr(hWnd), boolToUintptr(bEnable), 0)
	return ret != 0 // the window was previously disabled?
}

func (hWnd HWND) EnumChildWindows() []HWND {
	hChildren := make([]HWND, 0)
	syscall.Syscall(proc.EnumChildWindows.Addr(), 3,
		uintptr(hWnd),
		syscall.NewCallback(
			func(hChild HWND, lp LPARAM) uintptr {
				hChildren = append(hChildren, hChild)
				return boolToUintptr(true)
			}), 0)
	return hChildren
}

func (hWnd HWND) GetAncestor(gaFlags co.GA) HWND {
	ret, _, _ := syscall.Syscall(proc.GetAncestor.Addr(), 2,
		uintptr(hWnd), uintptr(gaFlags), 0)
	return HWND(ret)
}

func (hWnd HWND) GetClientRect() *RECT {
	rc := &RECT{}
	ret, _, lerr := syscall.Syscall(proc.GetClientRect.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)), 0)

	if ret == 0 {
		panic(co.ERROR(lerr).Format("GetClientRect failed."))
	}
	return rc
}

func (hWnd HWND) GetDC() HDC {
	ret, _, _ := syscall.Syscall(proc.GetDC.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic("GetDC failed.")
	}
	return HDC(ret)
}

func (hWnd HWND) GetDlgItem(nIDDlgItem int32) HWND {
	ret, _, lerr := syscall.Syscall(proc.GetDlgItem.Addr(), 2,
		uintptr(hWnd), uintptr(nIDDlgItem), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("GetDlgItem failed."))
	}
	return HWND(ret)
}

// Available in Windows 10, version 1607.
func (hWnd HWND) GetDpiForWindow() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetDpiForWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return uint32(ret)
}

func (hWnd HWND) GetExStyle() co.WS_EX {
	return co.WS_EX(hWnd.GetWindowLongPtr(co.GWLP_EXSTYLE))
}

func GetFocus() HWND {
	ret, _, _ := syscall.Syscall(proc.GetFocus.Addr(), 0, 0, 0, 0)
	return HWND(ret)
}

func GetForegroundWindow() HWND {
	ret, _, _ := syscall.Syscall(proc.GetForegroundWindow.Addr(), 0,
		0, 0, 0)
	return HWND(ret)
}

func (hWnd HWND) GetInstance() HINSTANCE {
	return HINSTANCE(hWnd.GetWindowLongPtr(co.GWLP_HINSTANCE))
}

func (hWnd HWND) GetMenu() HMENU {
	ret, _, _ := syscall.Syscall(proc.GetMenu.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return HMENU(ret)
}

func (hWnd HWND) GetNextDlgTabItem(hChild HWND, bPrevious bool) HWND {
	ret, _, lerr := syscall.Syscall(proc.GetNextDlgTabItem.Addr(), 3,
		uintptr(hWnd), uintptr(hChild), boolToUintptr(bPrevious))

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("GetNextDlgTabItem failed."))
	}
	return HWND(ret)
}

func (hWnd HWND) GetParent() HWND {
	ret, _, lerr := syscall.Syscall(proc.GetParent.Addr(), 1,
		uintptr(hWnd), 0, 0)

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("GetParent failed."))
	}
	return HWND(ret)
}

func (hWnd HWND) GetStyle() co.WS {
	return co.WS(hWnd.GetWindowLongPtr(co.GWLP_STYLE))
}

func (hWnd HWND) GetWindow(uCmd co.GW) HWND {
	ret, _, lerr := syscall.Syscall(proc.GetWindow.Addr(), 2,
		uintptr(hWnd), uintptr(uCmd), 0)

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("GetWindow failed."))
	}
	return HWND(ret)
}

func (hWnd HWND) GetWindowDC() HDC {
	ret, _, _ := syscall.Syscall(proc.GetWindowDC.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic("GetWindowDC failed.")
	}
	return HDC(ret)
}

func (hWnd HWND) GetWindowLongPtr(index co.GWLP) uintptr {
	ret, _, lerr := syscall.Syscall(proc.GetWindowLongPtr.Addr(), 2,
		uintptr(hWnd), uintptr(index),
		0)

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("GetWindowLongPtr failed."))
	}
	return ret
}

func (hWnd HWND) GetWindowRect() *RECT {
	rc := &RECT{}
	ret, _, lerr := syscall.Syscall(proc.GetWindowRect.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)), 0)

	if ret == 0 {
		panic(co.ERROR(lerr).Format("GetWindowRect failed."))
	}
	return rc
}

func (hWnd HWND) GetWindowText() string {
	len := hWnd.GetWindowTextLength() + 1
	buf := make([]uint16, len)

	ret, _, lerr := syscall.Syscall(proc.GetWindowText.Addr(), 3,
		uintptr(hWnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(len))

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("GetWindowText failed."))
	}
	return syscall.UTF16ToString(buf)
}

func (hWnd HWND) GetWindowTextLength() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetWindowTextLength.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return uint32(ret)
}

func (hWnd HWND) InvalidateRect(lpRect *RECT, bErase bool) {
	ret, _, _ := syscall.Syscall(proc.InvalidateRect.Addr(), 3,
		uintptr(hWnd), uintptr(unsafe.Pointer(lpRect)), boolToUintptr(bErase))
	if ret == 0 {
		panic("InvalidateRect failed.")
	}
}

func (hWnd HWND) IsChild(hChild HWND) bool {
	ret, _, _ := syscall.Syscall(proc.IsChild.Addr(), 2,
		uintptr(hWnd), uintptr(hChild), 0)
	return ret != 0
}

func (hWnd HWND) IsDialogMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(proc.IsDialogMessage.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(msg)), 0)
	return ret != 0
}

func (hWnd HWND) IsDlgButtonChecked(nIDButton int32) co.BST {
	ret, _, _ := syscall.Syscall(proc.IsDlgButtonChecked.Addr(), 2,
		uintptr(hWnd), uintptr(nIDButton), 0)
	return co.BST(ret)
}

func (hWnd HWND) IsTopLevelWindow() bool {
	// Allegedly undocumented Win32 function; implemented here.
	// https://stackoverflow.com/a/16975012
	return hWnd == hWnd.GetAncestor(co.GA_ROOT)
}

func (hWnd HWND) IsWindow() bool {
	ret, _, _ := syscall.Syscall(proc.IsWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

func (hWnd HWND) IsWindowEnabled() bool {
	ret, _, _ := syscall.Syscall(proc.IsWindowEnabled.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

func (hWnd HWND) MessageBox(message, caption string, flags co.MB) co.MBID {
	ret, _, _ := syscall.Syscall6(proc.MessageBox.Addr(), 4,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(StrToPtr(message))),
		uintptr(unsafe.Pointer(StrToPtr(caption))),
		uintptr(flags), 0, 0)
	return co.MBID(ret)
}

func (hWnd HWND) OpenThemeData(classNames string) HTHEME {
	ret, _, _ := syscall.Syscall(proc.OpenThemeData.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(StrToPtr(classNames))),
		0)
	return HTHEME(ret) // zero if no match, never fails
}

func (hWnd HWND) PostMessage(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.Syscall6(proc.PostMessage.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

func (hWnd HWND) ReleaseDC(hdc HDC) int32 {
	ret, _, _ := syscall.Syscall(proc.ReleaseDC.Addr(), 2,
		uintptr(hWnd), uintptr(hdc), 0)
	return int32(ret)
}

func (hWnd HWND) RemoveWindowSubclass(
	subclassProc uintptr, uIdSubclass uint32) {

	ret, _, _ := syscall.Syscall(proc.RemoveWindowSubclass.Addr(), 3,
		uintptr(hWnd), subclassProc, uintptr(uIdSubclass))
	if ret == 0 {
		panic("RemoveWindowSubclass failed.")
	}
}

func (hWnd HWND) ScreenToClientPt(point *POINT) {
	ret, _, _ := syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(point)), 0)
	if ret == 0 {
		panic("ScreenToClient failed for POINT.")
	}
}

func (hWnd HWND) ScreenToClientRc(rect *RECT) {
	ret, _, _ := syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(rect)), 0)
	if ret == 0 {
		panic("ScreenToClient failed for RECT (1).")
	}
	ret, _, _ = syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(&rect.Right)), 0)
	if ret == 0 {
		panic("ScreenToClient failed for RECT (2).")
	}
}

func (hWnd HWND) SendMessage(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.Syscall6(proc.SendMessage.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

func (hWnd HWND) SetExStyle(style co.WS) {
	hWnd.SetWindowLongPtr(co.GWLP_EXSTYLE, uintptr(style))
}

// Returns true if window was brought to foreground.
func (hWnd HWND) SetForegroundWindow() bool {
	ret, _, _ := syscall.Syscall(proc.SetForegroundWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

// Returns a handle to the window that previously had the focus, if any.
func (hWnd HWND) SetFocus() HWND {
	ret, _, lerr := syscall.Syscall(proc.SetFocus.Addr(), 1,
		uintptr(hWnd), 0, 0)

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("SetFocus failed."))
	}
	return HWND(ret)
}

func (hWnd HWND) SetStyle(style co.WS) {
	hWnd.SetWindowLongPtr(co.GWLP_STYLE, uintptr(style))
}

func (hWnd HWND) SetWindowLongPtr(index co.GWLP, newLong uintptr) uintptr {
	ret, _, lerr := syscall.Syscall(proc.SetWindowLongPtr.Addr(), 3,
		uintptr(hWnd), uintptr(index), newLong)

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("SetWindowLongPtr failed."))
	}
	return ret
}

func (hWnd HWND) ShowWindow(nCmdShow co.SW) bool {
	ret, _, _ := syscall.Syscall(proc.ShowWindow.Addr(), 1,
		uintptr(hWnd), uintptr(nCmdShow), 0)
	return ret != 0
}

// You can pass a HWND handle or SWP_HWND constants in hwndInsertAfter argument.
func (hWnd HWND) SetWindowPos(hwndInsertAfter co.SWP_HWND, x, y int32,
	cx, cy uint32, uFlags co.SWP) {

	ret, _, lerr := syscall.Syscall9(proc.SetWindowPos.Addr(), 7,
		uintptr(hWnd), uintptr(hwndInsertAfter),
		uintptr(x), uintptr(y), uintptr(cx), uintptr(cy),
		uintptr(uFlags), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("SetWindowPos failed."))
	}
}

// Use syscall.NewCallback() to convert the closure to uintptr, and keep this
// uintptr to pass to RemoveWindowSubclass.
func (hWnd HWND) SetWindowSubclass(subclassProc uintptr, uIdSubclass uint32,
	dwRefData unsafe.Pointer) {

	ret, _, _ := syscall.Syscall6(proc.SetWindowSubclass.Addr(), 4,
		uintptr(hWnd), subclassProc, uintptr(uIdSubclass), uintptr(dwRefData),
		0, 0)
	if ret == 0 {
		panic("SetWindowSubclass failed.")
	}
}

func (hWnd HWND) SetWindowText(lpString string) {
	syscall.Syscall(proc.SetWindowText.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(StrToPtr(lpString))),
		0)
}

func (hWnd HWND) TranslateAccelerator(hAccel HACCEL, msg *MSG) bool {
	ret, _, _ := syscall.Syscall(proc.TranslateAccelerator.Addr(), 3,
		uintptr(hWnd), uintptr(hAccel), uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

func (hWnd HWND) UpdateWindow() bool {
	ret, _, _ := syscall.Syscall(proc.UpdateWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}
