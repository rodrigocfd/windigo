package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

// A handle to a window.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hwnd
type HWND HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func CreateWindowEx(exStyle co.WS_EX, className, title string, style co.WS,
	x, y, width, height int32, parent HWND, menu HMENU,
	instance HINSTANCE, param LPARAM) HWND {

	ret, _, lerr := syscall.Syscall12(proc.CreateWindowEx.Addr(), 12,
		uintptr(exStyle),
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(className))),
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(title))),
		uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(parent), uintptr(menu), uintptr(instance), uintptr(param))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdesktopwindow
func GetDesktopWindow() HWND {
	ret, _, _ := syscall.Syscall(proc.GetDesktopWindow.Addr(), 0, 0, 0, 0)
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getfocus
func GetFocus() HWND {
	ret, _, _ := syscall.Syscall(proc.GetFocus.Addr(), 0, 0, 0, 0)
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
func GetForegroundWindow() HWND {
	ret, _, _ := syscall.Syscall(proc.GetForegroundWindow.Addr(), 0,
		0, 0, 0)
	return HWND(ret)
}

// ⚠️ You must defer EndPaint().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-beginpaint
func (hWnd HWND) BeginPaint(lpPaint *PAINTSTRUCT) HDC {
	ret, _, lerr := syscall.Syscall(proc.BeginPaint.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(lpPaint)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HDC(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenPt(point *POINT) {
	ret, _, lerr := syscall.Syscall(proc.ClientToScreen.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(point)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenRc(rect *RECT) {
	ret, _, lerr := syscall.Syscall(proc.ClientToScreen.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(rect)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	ret, _, lerr = syscall.Syscall(proc.ClientToScreen.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(&rect.Right)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-defsubclassproc
func (hWnd HWND) DefSubclassProc(
	msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {

	ret, _, _ := syscall.Syscall6(proc.DefSubclassProc.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defwindowprocw
func (hWnd HWND) DefWindowProc(
	msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {

	ret, _, _ := syscall.Syscall6(proc.DefWindowProc.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroywindow
func (hWnd HWND) DestroyWindow() {
	ret, _, lerr := syscall.Syscall(proc.DestroyWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragacceptfiles
func (hWnd HWND) DragAcceptFiles(fAccept bool) {
	syscall.Syscall(proc.DragAcceptFiles.Addr(), 2,
		uintptr(hWnd), util.BoolToUintptr(fAccept), 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawmenubar
func (hWnd HWND) DrawMenuBar() {
	ret, _, lerr := syscall.Syscall(proc.DrawMenuBar.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enablewindow
func (hWnd HWND) EnableWindow(bEnable bool) bool {
	ret, _, _ := syscall.Syscall(proc.EnableWindow.Addr(), 2,
		uintptr(hWnd), util.BoolToUintptr(bEnable), 0)
	return ret != 0 // the window was previously disabled?
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddialog
func (hWnd HWND) EndDialog(nResult uintptr) {
	ret, _, lerr := syscall.Syscall(proc.EndDialog.Addr(), 2,
		uintptr(hWnd), nResult, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endpaint
func (hWnd HWND) EndPaint(lpPaint *PAINTSTRUCT) {
	syscall.Syscall(proc.EndPaint.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(lpPaint)), 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumchildwindows
func (hWnd HWND) EnumChildWindows(
	lpEnumFunc func(hChild HWND, lParam LPARAM) bool,
	lParam LPARAM) {

	syscall.Syscall(proc.EnumChildWindows.Addr(), 3,
		uintptr(hWnd),
		syscall.NewCallback(
			func(hChild HWND, lParam LPARAM) uintptr {
				return util.BoolToUintptr(lpEnumFunc(hChild, lParam))
			}),
		uintptr(lParam))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getancestor
func (hWnd HWND) GetAncestor(gaFlags co.GA) HWND {
	ret, _, _ := syscall.Syscall(proc.GetAncestor.Addr(), 2,
		uintptr(hWnd), uintptr(gaFlags), 0)
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclasslongptrw
func (hWnd HWND) GetClassLongPtr(nIndex co.GCL) uint32 {
	ret, _, lerr := syscall.Syscall(proc.GetClassLongPtr.Addr(), 2,
		uintptr(hWnd), uintptr(nIndex), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassnamew
func (hWnd HWND) GetClassName() string {
	var buf [256 + 1]uint16

	ret, _, lerr := syscall.Syscall(proc.GetClassName.Addr(), 3,
		uintptr(hWnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(256+1))
	if ret == 0 && err.ERROR(lerr) != err.SUCCESS {
		panic(err.ERROR(lerr))
	}
	return Str.FromUint16Slice(buf[:])
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclientrect
func (hWnd HWND) GetClientRect() RECT {
	rc := RECT{}
	ret, _, lerr := syscall.Syscall(proc.GetClientRect.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(&rc)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return rc
}

// ⚠️ You must defer ReleaseDC().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
func (hWnd HWND) GetDC() HDC {
	ret, _, lerr := syscall.Syscall(proc.GetDC.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HDC(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgctrlid
func (hWnd HWND) GetDlgCtrlID() int32 {
	ret, _, lerr := syscall.Syscall(proc.GetDlgCtrlID.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgitem
func (hWnd HWND) GetDlgItem(nIDDlgItem int32) HWND {
	ret, _, lerr := syscall.Syscall(proc.GetDlgItem.Addr(), 2,
		uintptr(hWnd), uintptr(nIDDlgItem), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenu
func (hWnd HWND) GetMenu() HMENU {
	ret, _, _ := syscall.Syscall(proc.GetMenu.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return HMENU(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlgtabitem
func (hWnd HWND) GetNextDlgTabItem(hChild HWND, bPrevious bool) HWND {
	ret, _, lerr := syscall.Syscall(proc.GetNextDlgTabItem.Addr(), 3,
		uintptr(hWnd), uintptr(hChild), util.BoolToUintptr(bPrevious))
	if ret == 0 && err.ERROR(lerr) != err.SUCCESS {
		panic(err.ERROR(lerr))
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getparent
func (hWnd HWND) GetParent() HWND {
	ret, _, lerr := syscall.Syscall(proc.GetParent.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 && err.ERROR(lerr) != err.SUCCESS {
		panic(err.ERROR(lerr))
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getscrollinfo
func (hWnd HWND) GetScrollInfo(nBar co.SB_TYPE, lpsi *SCROLLINFO) {
	ret, _, lerr := syscall.Syscall(proc.GetScrollInfo.Addr(), 3,
		uintptr(hWnd), uintptr(nBar), uintptr(unsafe.Pointer(lpsi)))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmenu
func (hWnd HWND) GetSystemMenu(bRevert bool) HMENU {
	ret, _, _ := syscall.Syscall(proc.GetSystemMenu.Addr(), 2,
		uintptr(hWnd), util.BoolToUintptr(bRevert), 0)
	return HMENU(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindow
func (hWnd HWND) GetWindow(uCmd co.GW) HWND {
	ret, _, lerr := syscall.Syscall(proc.GetWindow.Addr(), 2,
		uintptr(hWnd), uintptr(uCmd), 0)
	if ret == 0 && err.ERROR(lerr) != err.SUCCESS {
		panic(err.ERROR(lerr))
	}
	return HWND(ret)
}

// ⚠️ You must defer ReleaseDC().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowdc
func (hWnd HWND) GetWindowDC() HDC {
	ret, _, lerr := syscall.Syscall(proc.GetWindowDC.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HDC(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
func (hWnd HWND) GetWindowLongPtr(index co.GWLP) uintptr {
	ret, _, lerr := syscall.Syscall(proc.GetWindowLongPtr.Addr(), 2,
		uintptr(hWnd), uintptr(index), 0)
	if ret == 0 && err.ERROR(lerr) != err.SUCCESS {
		panic(err.ERROR(lerr))
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowrect
func (hWnd HWND) GetWindowRect() RECT {
	rc := RECT{}
	ret, _, lerr := syscall.Syscall(proc.GetWindowRect.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(&rc)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return rc
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextw
func (hWnd HWND) GetWindowText() string {
	len := hWnd.GetWindowTextLength() + 1
	buf := make([]uint16, len)

	ret, _, lerr := syscall.Syscall(proc.GetWindowText.Addr(), 3,
		uintptr(hWnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(len))
	if ret == 0 && err.ERROR(lerr) != err.SUCCESS {
		panic(err.ERROR(lerr))
	}
	return Str.FromUint16Slice(buf)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextlengthw
func (hWnd HWND) GetWindowTextLength() int32 {
	ret, _, _ := syscall.Syscall(proc.GetWindowTextLength.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-hidecaret
func (hWnd HWND) HideCaret() {
	ret, _, lerr := syscall.Syscall(proc.HideCaret.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// Returns the window instance with GetWindowLongPtr().
func (hWnd HWND) Hinstance() HINSTANCE {
	return HINSTANCE(hWnd.GetWindowLongPtr(co.GWLP_HINSTANCE))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invalidaterect
func (hWnd HWND) InvalidateRect(lpRect *RECT, bErase bool) {
	ret, _, lerr := syscall.Syscall(proc.InvalidateRect.Addr(), 3,
		uintptr(hWnd), uintptr(unsafe.Pointer(lpRect)),
		util.BoolToUintptr(bErase))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-ischild
func (hWnd HWND) IsChild(hChild HWND) bool {
	ret, _, _ := syscall.Syscall(proc.IsChild.Addr(), 2,
		uintptr(hWnd), uintptr(hChild), 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdlgbuttonchecked
func (hWnd HWND) IsDlgButtonChecked(nIDButton int32) co.BST {
	ret, _, _ := syscall.Syscall(proc.IsDlgButtonChecked.Addr(), 2,
		uintptr(hWnd), uintptr(nIDButton), 0)
	return co.BST(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdialogmessagew
func (hWnd HWND) IsDialogMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(proc.IsDialogMessage.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(msg)), 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isiconic
func (hWnd HWND) IsIconic() bool {
	ret, _, _ := syscall.Syscall(proc.IsIconic.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

// Allegedly undocumented Win32 function; implemented here.
//
// https://stackoverflow.com/a/16975012
func (hWnd HWND) IsTopLevelWindow() bool {
	return hWnd == hWnd.GetAncestor(co.GA_ROOT)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iswindow
func (hWnd HWND) IsWindow() bool {
	ret, _, _ := syscall.Syscall(proc.IsWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iswindowenabled
func (hWnd HWND) IsWindowEnabled() bool {
	ret, _, _ := syscall.Syscall(proc.IsWindowEnabled.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iszoomed
func (hWnd HWND) IsZoomed() bool {
	ret, _, _ := syscall.Syscall(proc.IsZoomed.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-killtimer
func (hWnd HWND) KillTimer(uIDEvent uintptr) {
	ret, _, lerr := syscall.Syscall(proc.KillTimer.Addr(), 2,
		uintptr(hWnd), uIDEvent, 0)
	if ret == 0 && err.ERROR(lerr) != err.SUCCESS {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
func (hWnd HWND) MessageBox(lpText, lpCaption string, uType co.MB) co.ID {
	ret, _, lerr := syscall.Syscall6(proc.MessageBox.Addr(), 4,
		uintptr(hWnd), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpText))),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpCaption))), uintptr(uType),
		0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return co.ID(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-movewindow
func (hWnd HWND) MoveWindow(x, y, width, height int32, bRepaint bool) {
	ret, _, lerr := syscall.Syscall6(proc.MoveWindow.Addr(), 6,
		uintptr(hWnd), uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		util.BoolToUintptr(bRepaint))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// ⚠️ You must defer CloseThemeData().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-openthemedata
func (hWnd HWND) OpenThemeData(classNames string) (HTHEME, error) {
	ret, _, lerr := syscall.Syscall(proc.OpenThemeData.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(Str.ToUint16Ptr(classNames))),
		0)
	if ret == 0 {
		return HTHEME(0), err.ERROR(lerr)
	}
	return HTHEME(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postmessagew
func (hWnd HWND) PostMessage(msg co.WM, wParam WPARAM, lParam LPARAM) {
	ret, _, lerr := syscall.Syscall6(proc.PostMessage.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-releasedc
func (hWnd HWND) ReleaseDC(hdc HDC) {
	ret, _, lerr := syscall.Syscall(proc.ReleaseDC.Addr(), 2,
		uintptr(hWnd), uintptr(hdc), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-removewindowsubclass
func (hWnd HWND) RemoveWindowSubclass(
	subclassProc uintptr, uIdSubclass uint32) {

	ret, _, lerr := syscall.Syscall(proc.RemoveWindowSubclass.Addr(), 3,
		uintptr(hWnd), subclassProc, uintptr(uIdSubclass))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientPt(point *POINT) {
	ret, _, lerr := syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(point)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientRc(rect *RECT) {
	ret, _, lerr := syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(rect)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	ret, _, lerr = syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(&rect.Right)), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew
func (hWnd HWND) SendMessage(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.Syscall6(proc.SendMessage.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setfocus
func (hWnd HWND) SetFocus() (HWND, error) {
	ret, _, lerr := syscall.Syscall(proc.SetFocus.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		return HWND(0), err.ERROR(lerr)
	}
	return HWND(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setforegroundwindow
func (hWnd HWND) SetForegroundWindow() {
	ret, _, lerr := syscall.Syscall(proc.SetForegroundWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setparent
func (hWnd HWND) SetParent(hWndNewParent HWND) HWND {
	ret, _, lerr := syscall.Syscall(proc.SetParent.Addr(), 2,
		uintptr(hWnd), uintptr(hWndNewParent), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setscrollinfo
func (hWnd HWND) SetScrollInfo(
	nBar co.SB_TYPE, lpsi *SCROLLINFO, redraw bool) int32 {

	ret, _, _ := syscall.Syscall6(proc.SetScrollInfo.Addr(), 4,
		uintptr(hWnd), uintptr(nBar), uintptr(unsafe.Pointer(lpsi)),
		util.BoolToUintptr(redraw), 0, 0)
	return int32(ret)
}

// ⚠️ You must defer KillTimer().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-settimer
func (hWnd HWND) SetTimer(
	nIDEvent, uElapse uintptr, lpTimerFunc func(msElapsed uint32)) {

	cbTimer := uintptr(0)
	if lpTimerFunc != nil {
		cbTimer = syscall.NewCallback(
			func(hWnd HWND, nIDEvent uintptr, wmTimer uint32, msElapsed uint32) uintptr {
				lpTimerFunc(msElapsed)
				return 0
			})
	}

	ret, _, lerr := syscall.Syscall6(proc.SetTimer.Addr(), 4,
		uintptr(hWnd), nIDEvent, uElapse, cbTimer, 0, 0)
	if ret == 0 && err.ERROR(lerr) != err.SUCCESS {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
func (hWnd HWND) SetWindowLongPtr(index co.GWLP, newLong uintptr) uintptr {
	ret, _, lerr := syscall.Syscall(proc.SetWindowLongPtr.Addr(), 3,
		uintptr(hWnd), uintptr(index), newLong)
	if ret == 0 && err.ERROR(lerr) != err.SUCCESS {
		panic(err.ERROR(lerr))
	}
	return ret
}

// You can pass HWND or HWND_IA in hwndInsertAfter argument.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
func (hWnd HWND) SetWindowPos(
	hwndInsertAfter HWND, x, y, cx, cy int32, uFlags co.SWP) {

	ret, _, lerr := syscall.Syscall9(proc.SetWindowPos.Addr(), 7,
		uintptr(hWnd), uintptr(hwndInsertAfter),
		uintptr(x), uintptr(y), uintptr(cx), uintptr(cy),
		uintptr(uFlags), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// Use syscall.NewCallback() to convert the closure to uintptr, and keep this
// uintptr to pass to RemoveWindowSubclass().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-setwindowsubclass
func (hWnd HWND) SetWindowSubclass(
	subclassProc uintptr, uIdSubclass uint32, dwRefData unsafe.Pointer) {

	ret, _, lerr := syscall.Syscall6(proc.SetWindowSubclass.Addr(), 4,
		uintptr(hWnd), subclassProc, uintptr(uIdSubclass), uintptr(dwRefData),
		0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowtextw
func (hWnd HWND) SetWindowText(lpString string) {
	syscall.Syscall(proc.SetWindowText.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpString))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showcaret
func (hWnd HWND) ShowCaret() {
	ret, _, lerr := syscall.Syscall(proc.ShowCaret.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
func (hWnd HWND) ShowWindow(nCmdShow co.SW) bool {
	ret, _, _ := syscall.Syscall(proc.ShowWindow.Addr(), 1,
		uintptr(hWnd), uintptr(nCmdShow), 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialog
func (hWnd HWND) TaskDialog(
	hInstance HINSTANCE,
	pszWindowTitle, pszMainInstruction, pszContent string,
	dwCommonButtons co.TDCBF, pszIcon co.TD_ICON) co.ID {

	pnButton := int32(0)
	ret, _, _ := syscall.Syscall9(proc.TaskDialog.Addr(), 8,
		uintptr(hWnd), uintptr(hInstance),
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(pszWindowTitle))),
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(pszMainInstruction))),
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(pszContent))),
		uintptr(dwCommonButtons), uintptr(pszIcon),
		uintptr(unsafe.Pointer(&pnButton)), 0)
	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return co.ID(pnButton)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translateacceleratorw
func (hWnd HWND) TranslateAccelerator(hAccel HACCEL, msg *MSG) error {
	ret, _, lerr := syscall.Syscall(proc.TranslateAccelerator.Addr(), 3,
		uintptr(hWnd), uintptr(hAccel), uintptr(unsafe.Pointer(msg)))
	if ret == 0 {
		return err.ERROR(lerr)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-updatewindow
func (hWnd HWND) UpdateWindow() bool {
	ret, _, _ := syscall.Syscall(proc.UpdateWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}
