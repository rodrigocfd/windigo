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

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenPt(point *POINT) {
	ret, _, _ := syscall.Syscall(proc.ClientToScreen.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(point)), 0)
	if ret == 0 {
		panic("ClientToScreen failed for POINT.")
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
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

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
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

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-defsubclassproc
func (hWnd HWND) DefSubclassProc(msg co.WM,
	wParam WPARAM, lParam LPARAM) uintptr {

	ret, _, _ := syscall.Syscall6(proc.DefSubclassProc.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defwindowprocw
func (hWnd HWND) DefWindowProc(msg co.WM,
	wParam WPARAM, lParam LPARAM) uintptr {

	ret, _, _ := syscall.Syscall6(proc.DefWindowProc.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroywindow
func (hWnd HWND) DestroyWindow() {
	ret, _, lerr := syscall.Syscall(proc.DestroyWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("DestroyWindow failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragacceptfiles
func (hWnd HWND) DragAcceptFiles(fAccept bool) {
	syscall.Syscall(proc.DragAcceptFiles.Addr(), 2,
		uintptr(hWnd), boolToUintptr(fAccept), 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawmenubar
func (hWnd HWND) DrawMenuBar() {
	ret, _, lerr := syscall.Syscall(proc.DrawMenuBar.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("DrawMenuBar failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enablewindow
func (hWnd HWND) EnableWindow(bEnable bool) bool {
	ret, _, _ := syscall.Syscall(proc.EnableWindow.Addr(), 2,
		uintptr(hWnd), boolToUintptr(bEnable), 0)
	return ret != 0 // the window was previously disabled?
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumchildwindows
func (hWnd HWND) EnumChildWindows(
	lpEnumFunc func(hChild HWND, lParam LPARAM) bool,
	lParam LPARAM) {

	syscall.Syscall(proc.EnumChildWindows.Addr(), 3,
		uintptr(hWnd),
		syscall.NewCallback(
			func(hChild HWND, lParam LPARAM) int32 {
				return boolToInt32(lpEnumFunc(hChild, lParam))
			}),
		uintptr(lParam))
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getancestor
func (hWnd HWND) GetAncestor(gaFlags co.GA) HWND {
	ret, _, _ := syscall.Syscall(proc.GetAncestor.Addr(), 2,
		uintptr(hWnd), uintptr(gaFlags), 0)
	return HWND(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclientrect
func (hWnd HWND) GetClientRect() *RECT {
	rc := &RECT{}
	ret, _, lerr := syscall.Syscall(proc.GetClientRect.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)), 0)

	if ret == 0 {
		panic(co.ERROR(lerr).Format("GetClientRect failed."))
	}
	return rc
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
func (hWnd HWND) GetDC() HDC {
	ret, _, _ := syscall.Syscall(proc.GetDC.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic("GetDC failed.")
	}
	return HDC(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgctrlid
func (hWnd HWND) GetDlgCtrlID() int32 {
	ret, _, lerr := syscall.Syscall(proc.GetDlgCtrlID.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("GetDlgCtrlID failed."))
	}
	return int32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgitem
func (hWnd HWND) GetDlgItem(nIDDlgItem int32) HWND {
	ret, _, lerr := syscall.Syscall(proc.GetDlgItem.Addr(), 2,
		uintptr(hWnd), uintptr(nIDDlgItem), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("GetDlgItem failed."))
	}
	return HWND(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdpiforwindow
//
// Available in Windows 10, version 1607.
func (hWnd HWND) GetDpiForWindow() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetDpiForWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return uint32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
//
// GetWindowLongPtr() with GWLP_EXSTYLE flag.
func (hWnd HWND) GetExStyle() co.WS_EX {
	return co.WS_EX(hWnd.GetWindowLongPtr(co.GWLP_EXSTYLE))
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getfocus
func GetFocus() HWND {
	ret, _, _ := syscall.Syscall(proc.GetFocus.Addr(), 0, 0, 0, 0)
	return HWND(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
func GetForegroundWindow() HWND {
	ret, _, _ := syscall.Syscall(proc.GetForegroundWindow.Addr(), 0,
		0, 0, 0)
	return HWND(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
//
// GetWindowLongPtr() with GWLP_HINSTANCE flag.
func (hWnd HWND) GetInstance() HINSTANCE {
	return HINSTANCE(hWnd.GetWindowLongPtr(co.GWLP_HINSTANCE))
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenu
func (hWnd HWND) GetMenu() HMENU {
	ret, _, _ := syscall.Syscall(proc.GetMenu.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return HMENU(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlgtabitem
func (hWnd HWND) GetNextDlgTabItem(hChild HWND, bPrevious bool) HWND {
	ret, _, lerr := syscall.Syscall(proc.GetNextDlgTabItem.Addr(), 3,
		uintptr(hWnd), uintptr(hChild), boolToUintptr(bPrevious))

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("GetNextDlgTabItem failed."))
	}
	return HWND(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getparent
func (hWnd HWND) GetParent() HWND {
	ret, _, lerr := syscall.Syscall(proc.GetParent.Addr(), 1,
		uintptr(hWnd), 0, 0)

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("GetParent failed."))
	}
	return HWND(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
//
// GetWindowLongPtr() with GWLP_STYLE flag.
func (hWnd HWND) GetStyle() co.WS {
	return co.WS(hWnd.GetWindowLongPtr(co.GWLP_STYLE))
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindow
func (hWnd HWND) GetWindow(uCmd co.GW) HWND {
	ret, _, lerr := syscall.Syscall(proc.GetWindow.Addr(), 2,
		uintptr(hWnd), uintptr(uCmd), 0)

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("GetWindow failed."))
	}
	return HWND(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowdc
func (hWnd HWND) GetWindowDC() HDC {
	ret, _, _ := syscall.Syscall(proc.GetWindowDC.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic("GetWindowDC failed.")
	}
	return HDC(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
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

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowrect
func (hWnd HWND) GetWindowRect() *RECT {
	rc := &RECT{}
	ret, _, lerr := syscall.Syscall(proc.GetWindowRect.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)), 0)

	if ret == 0 {
		panic(co.ERROR(lerr).Format("GetWindowRect failed."))
	}
	return rc
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextw
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

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextlengthw
func (hWnd HWND) GetWindowTextLength() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetWindowTextLength.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return uint32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-hidecaret
func (hWnd HWND) HideCaret() {
	ret, _, lerr := syscall.Syscall(proc.HideCaret.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("HideCaret failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invalidaterect
func (hWnd HWND) InvalidateRect(lpRect *RECT, bErase bool) {
	ret, _, _ := syscall.Syscall(proc.InvalidateRect.Addr(), 3,
		uintptr(hWnd), uintptr(unsafe.Pointer(lpRect)), boolToUintptr(bErase))
	if ret == 0 {
		panic("InvalidateRect failed.")
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-ischild
func (hWnd HWND) IsChild(hChild HWND) bool {
	ret, _, _ := syscall.Syscall(proc.IsChild.Addr(), 2,
		uintptr(hWnd), uintptr(hChild), 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdialogmessagew
func (hWnd HWND) IsDialogMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(proc.IsDialogMessage.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(msg)), 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdlgbuttonchecked
func (hWnd HWND) IsDlgButtonChecked(nIDButton int32) co.BST {
	ret, _, _ := syscall.Syscall(proc.IsDlgButtonChecked.Addr(), 2,
		uintptr(hWnd), uintptr(nIDButton), 0)
	return co.BST(ret)
}

// Allegedly undocumented Win32 function; implemented here.
// https://stackoverflow.com/a/16975012
func (hWnd HWND) IsTopLevelWindow() bool {
	return hWnd == hWnd.GetAncestor(co.GA_ROOT)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iswindow
func (hWnd HWND) IsWindow() bool {
	ret, _, _ := syscall.Syscall(proc.IsWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iswindowenabled
func (hWnd HWND) IsWindowEnabled() bool {
	ret, _, _ := syscall.Syscall(proc.IsWindowEnabled.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
func (hWnd HWND) MessageBox(message, caption string, flags co.MB) co.MBID {
	ret, _, _ := syscall.Syscall6(proc.MessageBox.Addr(), 4,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(StrToPtr(message))),
		uintptr(unsafe.Pointer(StrToPtr(caption))),
		uintptr(flags), 0, 0)
	return co.MBID(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-movewindow
func (hWnd HWND) MoveWindow(x, y int32, width, height uint32, bRepaint bool) {
	ret, _, lerr := syscall.Syscall6(proc.MoveWindow.Addr(), 6,
		uintptr(hWnd), uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		boolToUintptr(bRepaint))
	if ret == 0 {
		panic(co.ERROR(lerr).Format("MoveWindow failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-openthemedata
func (hWnd HWND) OpenThemeData(classNames string) HTHEME {
	ret, _, _ := syscall.Syscall(proc.OpenThemeData.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(StrToPtr(classNames))),
		0)
	return HTHEME(ret) // zero if no match, never fails
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postmessagew
func (hWnd HWND) PostMessage(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.Syscall6(proc.PostMessage.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-releasedc
func (hWnd HWND) ReleaseDC(hdc HDC) int32 {
	ret, _, _ := syscall.Syscall(proc.ReleaseDC.Addr(), 2,
		uintptr(hWnd), uintptr(hdc), 0)
	return int32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-removewindowsubclass
func (hWnd HWND) RemoveWindowSubclass(
	subclassProc uintptr, uIdSubclass uint32) {

	ret, _, _ := syscall.Syscall(proc.RemoveWindowSubclass.Addr(), 3,
		uintptr(hWnd), subclassProc, uintptr(uIdSubclass))
	if ret == 0 {
		panic("RemoveWindowSubclass failed.")
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientPt(point *POINT) {
	ret, _, _ := syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(point)), 0)
	if ret == 0 {
		panic("ScreenToClient failed for POINT.")
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
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

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew
func (hWnd HWND) SendMessage(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.Syscall6(proc.SendMessage.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
//
// SetWindowLongPtr() with GWLP_EXSTYLE flag.
func (hWnd HWND) SetExStyle(style co.WS) {
	hWnd.SetWindowLongPtr(co.GWLP_EXSTYLE, uintptr(style))
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setforegroundwindow
func (hWnd HWND) SetForegroundWindow() bool {
	ret, _, _ := syscall.Syscall(proc.SetForegroundWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setfocus
func (hWnd HWND) SetFocus() HWND {
	ret, _, lerr := syscall.Syscall(proc.SetFocus.Addr(), 1,
		uintptr(hWnd), 0, 0)

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("SetFocus failed."))
	}
	return HWND(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setparent
func (hWnd HWND) SetParent(hWndNewParent HWND) HWND {
	ret, _, lerr := syscall.Syscall(proc.SetParent.Addr(), 2,
		uintptr(hWnd), uintptr(hWndNewParent), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("SetParent failed."))
	}
	return HWND(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
//
// SetWindowLongPtr() with GWLP_STYLE flag.
func (hWnd HWND) SetStyle(style co.WS) {
	hWnd.SetWindowLongPtr(co.GWLP_STYLE, uintptr(style))
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
func (hWnd HWND) SetWindowLongPtr(index co.GWLP, newLong uintptr) uintptr {
	ret, _, lerr := syscall.Syscall(proc.SetWindowLongPtr.Addr(), 3,
		uintptr(hWnd), uintptr(index), newLong)

	lerr2 := co.ERROR(lerr)
	if ret == 0 && lerr2 != co.ERROR_SUCCESS {
		panic(lerr2.Format("SetWindowLongPtr failed."))
	}
	return ret
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
//
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

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-setwindowsubclass
//
// Use syscall.NewCallback() to convert the closure to uintptr, and keep this
// uintptr to pass to RemoveWindowSubclass().
func (hWnd HWND) SetWindowSubclass(subclassProc uintptr, uIdSubclass uint32,
	dwRefData unsafe.Pointer) {

	ret, _, _ := syscall.Syscall6(proc.SetWindowSubclass.Addr(), 4,
		uintptr(hWnd), subclassProc, uintptr(uIdSubclass), uintptr(dwRefData),
		0, 0)
	if ret == 0 {
		panic("SetWindowSubclass failed.")
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowtextw
func (hWnd HWND) SetWindowText(lpString string) {
	syscall.Syscall(proc.SetWindowText.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(StrToPtr(lpString))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showcaret
func (hWnd HWND) ShowCaret() {
	ret, _, lerr := syscall.Syscall(proc.ShowCaret.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("ShowCaret failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
func (hWnd HWND) ShowWindow(nCmdShow co.SW) bool {
	ret, _, _ := syscall.Syscall(proc.ShowWindow.Addr(), 1,
		uintptr(hWnd), uintptr(nCmdShow), 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translateacceleratorw
func (hWnd HWND) TranslateAccelerator(hAccel HACCEL, msg *MSG) bool {
	ret, _, _ := syscall.Syscall(proc.TranslateAccelerator.Addr(), 3,
		uintptr(hWnd), uintptr(hAccel), uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-updatewindow
func (hWnd HWND) UpdateWindow() bool {
	ret, _, _ := syscall.Syscall(proc.UpdateWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}
