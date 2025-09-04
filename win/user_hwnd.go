//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// Handle to a [window].
//
// [window]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hwnd
type HWND HANDLE

// [CreateWindowEx] function.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func CreateWindowEx(
	exStyle co.WS_EX,
	className ClassName,
	title string,
	style co.WS,
	x, y, width, height int,
	parent HWND,
	menu HMENU,
	instance HINSTANCE,
	param LPARAM,
) (HWND, error) {
	var wClassName, wTitle wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_CreateWindowExW, "CreateWindowExW"),
		uintptr(exStyle),
		className.raw(&wClassName),
		uintptr(wTitle.EmptyIsNil(title)),
		uintptr(style),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(int32(width)),
		uintptr(int32(height)),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(param))
	if ret == 0 {
		return HWND(0), co.ERROR(err)
	}
	return HWND(ret), nil
}

var _CreateWindowExW *syscall.Proc

// [FindWindow] function.
//
// [FindWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-findwindoww
func FindWindow(className ClassName, title string) (HWND, bool) {
	var wClassName, wTitle wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_FindWindowW, "FindWindowW"),
		className.raw(&wClassName),
		uintptr(wTitle.EmptyIsNil(title)))
	return HWND(ret), ret != 0
}

var _FindWindowW *syscall.Proc

// [GetClipboardOwner] function.
//
// [GetClipboardOwner]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardowner
func GetClipboardOwner() (HWND, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetClipboardOwner, "GetClipboardOwner"))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetClipboardOwner *syscall.Proc

// [GetDesktopWindow] function.
//
// [GetDesktopWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdesktopwindow
func GetDesktopWindow() HWND {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetDesktopWindow, "GetDesktopWindow"))
	return HWND(ret)
}

var _GetDesktopWindow *syscall.Proc

// [GetFocus] function.
//
// [GetFocus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getfocus
func GetFocus() HWND {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetFocus, "GetFocus"))
	return HWND(ret)
}

var _GetFocus *syscall.Proc

// [GetForegroundWindow] function.
//
// [GetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
func GetForegroundWindow() HWND {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetForegroundWindow, "GetForegroundWindow"))
	return HWND(ret)
}

var _GetForegroundWindow *syscall.Proc

// [GetOpenClipboardWindow] function.
//
// [GetOpenClipboardWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getopenclipboardwindow
func GetOpenClipboardWindow() (HWND, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetOpenClipboardWindow, "GetOpenClipboardWindow"))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetOpenClipboardWindow *syscall.Proc

// [GetShellWindow] function.
//
// [GetShellWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getshellwindow
func GetShellWindow() HWND {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetShellWindow, "GetShellWindow"))
	return HWND(ret)
}

var _GetShellWindow *syscall.Proc

// [WindowFromPhysicalPoint] function.
//
// [WindowFromPhysicalPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-windowfromphysicalpoint
func WindowFromPhysicalPoint() HWND {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_WindowFromPhysicalPoint, "WindowFromPhysicalPoint"))
	return HWND(ret)
}

var _WindowFromPhysicalPoint *syscall.Proc

// [WindowFromPoint] function.
//
// [WindowFromPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-windowfrompoint
func WindowFromPoint() HWND {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_WindowFromPoint, "WindowFromPoint"))
	return HWND(ret)
}

var _WindowFromPoint *syscall.Proc

// [AnimateWindow] function.
//
// Panics if time is negative.
//
// [AnimateWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-animatewindow
func (hWnd HWND) AnimateWindow(time int, flags co.AW) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_AnimateWindow, "AnimateWindow"),
		uintptr(hWnd),
		uintptr(uint32(time)),
		uintptr(flags))
	return utl.ZeroAsGetLastError(ret, err)
}

var _AnimateWindow *syscall.Proc

// [BeginPaint] function.
//
// ⚠️ You must defer [HWND.EndPaint].
//
// Example:
//
//	var hWnd win.HWND // initialized somewhere
//
//	var ps win.PAINTSTRUCT
//	hdc, _ := hWnd.BeginPaint(&ps)
//	defer hWnd.EndPaint(&ps)
//
// [BeginPaint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-beginpaint
func (hWnd HWND) BeginPaint(ps *PAINTSTRUCT) (HDC, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_BeginPaint, "BeginPaint"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(ps)))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _BeginPaint *syscall.Proc

// [BringWindowToTop] function.
//
// [BringWindowToTop]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-bringwindowtotop
func (hWnd HWND) BringWindowToTop() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_BringWindowToTop, "BringWindowToTop"),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

var _BringWindowToTop *syscall.Proc

// [ChildWindowFromPoint] function.
//
// [ChildWindowFromPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-childwindowfrompoint
func (hWnd HWND) ChildWindowFromPoint(pt POINT) (HWND, bool) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_ChildWindowFromPoint, "ChildWindowFromPoint"),
		uintptr(hWnd),
		uintptr(pt.X),
		uintptr(pt.Y))
	if ret == 0 {
		return HWND(0), false
	}
	return HWND(ret), true
}

var _ChildWindowFromPoint *syscall.Proc

// [ChildWindowFromPointEx] function.
//
// [ChildWindowFromPointEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-childwindowfrompointex
func (hWnd HWND) ChildWindowFromPointEx(pt POINT, flags co.CWP) (HWND, bool) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_ChildWindowFromPointEx, "ChildWindowFromPointEx"),
		uintptr(hWnd),
		uintptr(pt.X),
		uintptr(pt.Y),
		uintptr(flags))
	if ret == 0 {
		return HWND(0), false
	}
	return HWND(ret), true
}

var _ChildWindowFromPointEx *syscall.Proc

// [ClientToScreen] function for [POINT].
//
// [ClientToScreen]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenPt(pt *POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_ClientToScreen, "ClientToScreen"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pt)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ClientToScreen *syscall.Proc

// [ClientToScreen] function for [RECT].
//
// [ClientToScreen]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenRc(rc *RECT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_ClientToScreen, "ClientToScreen"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)))
	if ret == 0 {
		return co.ERROR_INVALID_PARAMETER
	}

	ret, _, _ = syscall.SyscallN(
		dll.Load(dll.USER32, &_ClientToScreen, "ClientToScreen"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&rc.Right)))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [CloseWindow] function.
//
// Note that this function will actually minimize the window, not destroy it.
//
// [CloseWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-closewindow
func (hWnd HWND) CloseWindow() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_CloseWindow, "CloseWindow"),
		uintptr(hWnd))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _CloseWindow *syscall.Proc

// [DefDlgProc] function.
//
// [DefDlgProc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defdlgprocw
func (hWnd HWND) DefDlgProc(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_DefDlgProcW, "DefDlgProcW"),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return ret
}

var _DefDlgProcW *syscall.Proc

// [DefWindowProc] function.
//
// [DefWindowProc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defwindowprocw
func (hWnd HWND) DefWindowProc(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_DefWindowProcW, "DefWindowProcW"),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return ret
}

var _DefWindowProcW *syscall.Proc

// [DestroyWindow] function.
//
// Note: don't call this function to close a window. The correct way to close a
// window is sending [co.WM_CLOSE].
//
// [DestroyWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroywindow
func (hWnd HWND) DestroyWindow() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_DestroyWindow, "DestroyWindow"),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

var _DestroyWindow *syscall.Proc

// [DrawMenuBar] function.
//
// [DrawMenuBar]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawmenubar
func (hWnd HWND) DrawMenuBar() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_DrawMenuBar, "DrawMenuBar"),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

var _DrawMenuBar *syscall.Proc

// [EnableWindow] function.
//
// [EnableWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enablewindow
func (hWnd HWND) EnableWindow(enable bool) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_EnableWindow, "EnableWindow"),
		uintptr(hWnd), utl.BoolToUintptr(enable))
	return ret != 0 // the window was previously disabled?
}

var _EnableWindow *syscall.Proc

// [EndDialog] function.
//
// [EndDialog]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddialog
func (hWnd HWND) EndDialog(result uintptr) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_EndDialog, "EndDialog"),
		uintptr(hWnd),
		result)
	return utl.ZeroAsGetLastError(ret, err)
}

var _EndDialog *syscall.Proc

// [EndPaint] function.
//
// Paired with [HWND.BeginPaint].
//
// [EndPaint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endpaint
func (hWnd HWND) EndPaint(ps *PAINTSTRUCT) {
	syscall.SyscallN(
		dll.Load(dll.USER32, &_EndPaint, "EndPaint"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(ps)))
}

var _EndPaint *syscall.Proc

// [EnumChildWindows] function.
//
// [EnumChildWindows]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumchildwindows
func (hWnd HWND) EnumChildWindows() []HWND {
	pPack := &_EnumChildWindowsPack{
		arr: make([]HWND, 0),
	}

	syscall.SyscallN(
		dll.Load(dll.USER32, &_EnumChildWindows, "EnumChildWindows"),
		uintptr(hWnd),
		enumChildWindowsCallback(),
		uintptr(unsafe.Pointer(pPack)))
	return pPack.arr
}

var _EnumChildWindows *syscall.Proc

type _EnumChildWindowsPack struct{ arr []HWND }

var _enumChildWindowsCallback uintptr

func enumChildWindowsCallback() uintptr {
	if _enumChildWindowsCallback != 0 {
		return _enumChildWindowsCallback
	}

	_enumChildWindowsCallback = syscall.NewCallback(
		func(hChild HWND, lParam LPARAM) uintptr {
			pPack := (*_EnumChildWindowsPack)(unsafe.Pointer(lParam))
			pPack.arr = append(pPack.arr, hChild)
			return 1
		},
	)
	return _enumChildWindowsCallback
}

// Calls [HWND.GetWindowLongPtr] to retrieve the window extended style.
func (hWnd HWND) ExStyle() (co.WS_EX, error) {
	exStyle, err := hWnd.GetWindowLongPtr(co.GWLP_EXSTYLE)
	return co.WS_EX(exStyle), err
}

// [GetAncestor] function.
//
// [GetAncestor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getancestor
func (hWnd HWND) GetAncestor(gaFlags co.GA) (HWND, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetAncestor, "GetAncestor"),
		uintptr(hWnd),
		uintptr(gaFlags))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetAncestor *syscall.Proc

// [GetClassName] function.
//
// [GetClassName]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassnamew
func (hWnd HWND) GetClassName() (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetClassNameW, "GetClassNameW"),
		uintptr(hWnd),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len())))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return "", wErr
	}
	return wBuf.String(), nil
}

var _GetClassNameW *syscall.Proc

// [GetClientRect] function.
//
// [GetClientRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclientrect
func (hWnd HWND) GetClientRect() (RECT, error) {
	var rc RECT
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetClientRect, "GetClientRect"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		return RECT{}, co.ERROR(err)
	}
	return rc, nil
}

var _GetClientRect *syscall.Proc

// [GetDC] function.
//
// ⚠️ You must defer [HWND.ReleaseDC].
//
// Example:
//
// Retrieving the DC of the entire screen:
//
//	hdcScreen, _ := win.HWND(0).GetDC()
//	defer win.HWND(0).ReleaseDC(hdcScreen)
//
// [GetDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
func (hWnd HWND) GetDC() (HDC, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetDC, "GetDC"),
		uintptr(hWnd))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _GetDC *syscall.Proc

// [GetDCEx] function.
//
// ⚠️ You must defer [HWND.ReleaseDC].
//
// [GetDCEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
func (hWnd HWND) GetDCEx(hRgnClip HRGN, flags co.DCX) (HDC, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetDCEx, "GetDCEx"),
		uintptr(hWnd),
		uintptr(hRgnClip),
		uintptr(flags))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _GetDCEx *syscall.Proc

// [GetDlgCtrlID] function.
//
// [GetDlgCtrlID]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgctrlid
func (hWnd HWND) GetDlgCtrlID() (uint16, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetDlgCtrlID, "GetDlgCtrlID"),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint16(ret), nil
}

var _GetDlgCtrlID *syscall.Proc

// [GetDlgItem] function.
//
// [GetDlgItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgitem
func (hWnd HWND) GetDlgItem(dlgId uint16) (HWND, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetDlgItem, "GetDlgItem"),
		uintptr(hWnd),
		uintptr(int32(dlgId)))
	if ret == 0 {
		return HWND(0), co.ERROR(err)
	}
	return HWND(ret), nil
}

var _GetDlgItem *syscall.Proc

// [GetLastActivePopup] function.
//
// [GetLastActivePopup]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getlastactivepopup
func (hWnd HWND) GetLastActivePopup() HWND {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetLastActivePopup, "GetLastActivePopup"),
		uintptr(hWnd))
	return HWND(ret)
}

var _GetLastActivePopup *syscall.Proc

// [GetMenu] function.
//
// [GetMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenu
func (hWnd HWND) GetMenu() HMENU {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetMenu, "GetMenu"),
		uintptr(hWnd))
	return HMENU(ret)
}

var _GetMenu *syscall.Proc

// [GetNextDlgGroupItem] function.
//
// [GetNextDlgGroupItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlggroupitem
func (hWnd HWND) GetNextDlgGroupItem(hChild HWND, isPrevious bool) (HWND, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetNextDlgGroupItem, "GetNextDlgGroupItem"),
		uintptr(hWnd),
		uintptr(hChild),
		utl.BoolToUintptr(isPrevious))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetNextDlgGroupItem *syscall.Proc

// [GetNextDlgTabItem] function.
//
// [GetNextDlgTabItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlgtabitem
func (hWnd HWND) GetNextDlgTabItem(hChild HWND, isPrevious bool) (HWND, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetNextDlgTabItem, "GetNextDlgTabItem"),
		uintptr(hWnd),
		uintptr(hChild),
		utl.BoolToUintptr(isPrevious))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetNextDlgTabItem *syscall.Proc

// [GetParent] function.
//
// [GetParent]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getparent
func (hWnd HWND) GetParent() (HWND, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetParent, "GetParent"),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetParent *syscall.Proc

// [GetTitleBarInfo] function.
//
// [GetTitleBarInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-gettitlebarinfo
func (hWnd HWND) GetTitleBarInfo() (TITLEBARINFO, error) {
	var ti TITLEBARINFO
	ti.SetCbSize()

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetTitleBarInfo, "GetTitleBarInfo"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&ti)))
	if ret == 0 {
		return TITLEBARINFO{}, co.ERROR(err)
	}
	return ti, nil
}

var _GetTitleBarInfo *syscall.Proc

// [GetTopWindow] function.
//
// [GetTopWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-gettopwindow
func (hWnd HWND) GetTopWindow() (HWND, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetTopWindow, "GetTopWindow"),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetTopWindow *syscall.Proc

// [GetWindow] function.
//
// [GetWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindow
func (hWnd HWND) GetWindow(cmd co.GW) (HWND, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetWindow, "GetWindow"),
		uintptr(hWnd),
		uintptr(cmd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetWindow *syscall.Proc

// [GetWindowDC] function.
//
// ⚠️ You must defer [HWND.ReleaseDC].
//
// [GetWindowDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowdc
func (hWnd HWND) GetWindowDC() (HDC, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetWindowDC, "GetWindowDC"),
		uintptr(hWnd))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _GetWindowDC *syscall.Proc

// [GetWindowDisplayAffinity] function.
//
// [GetWindowDisplayAffinity]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowdisplayaffinity
func (hWnd HWND) GetWindowDisplayAffinity() (co.WDA, error) {
	var da co.WDA
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetWindowDisplayAffinity, "GetWindowDisplayAffinity"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&da)))
	if ret == 0 {
		return co.WDA(0), co.ERROR(err)
	}
	return da, nil
}

var _GetWindowDisplayAffinity *syscall.Proc

// [GetWindowPlacement] function.
//
// [GetWindowPlacement]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowplacement
func (hWnd HWND) GetWindowPlacement() (WINDOWPLACEMENT, error) {
	var wp WINDOWPLACEMENT
	wp.SetLength()

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetWindowPlacement, "GetWindowPlacement"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&wp)))
	if ret == 0 {
		return WINDOWPLACEMENT{}, co.ERROR(err)
	}
	return wp, nil
}

var _GetWindowPlacement *syscall.Proc

// [GetWindowRect] function.
//
// [GetWindowRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowrect
func (hWnd HWND) GetWindowRect() (RECT, error) {
	var rc RECT
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetWindowRect, "GetWindowRect"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		return RECT{}, co.ERROR(err)
	}
	return rc, nil
}

var _GetWindowRect *syscall.Proc

// [GetWindowText] function.
//
// Calls [HWND.GetWindowTextLength] to allocate the memory block.
//
// [GetWindowText]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextw
func (hWnd HWND) GetWindowText() (string, error) {
	bufSz, errLen := hWnd.GetWindowTextLength()
	if errLen != nil {
		return "", errLen
	}
	bufSz += 1 // needed buffer, also counting terminating null

	var wBuf wstr.BufDecoder
	wBuf.Alloc(bufSz)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetWindowTextW, "GetWindowTextW"),
		uintptr(hWnd),
		uintptr(wBuf.Ptr()),
		uintptr(int32(bufSz)))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return "", wErr
	}
	return wBuf.String(), nil
}

var _GetWindowTextW *syscall.Proc

// [GetWindowTextLength] function.
//
// You usually don't need to call this function since [HWND.GetWindowText] already
// calls it to allocate the memory block.
//
// [GetWindowTextLength]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextlengthw
func (hWnd HWND) GetWindowTextLength() (int, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetWindowTextLengthW, "GetWindowTextLengthW"),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return int(int32(ret)), nil
}

var _GetWindowTextLengthW *syscall.Proc

// [GetWindowThreadProcessId] function.
//
// [GetWindowThreadProcessId]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowthreadprocessid
func (hWnd HWND) GetWindowThreadProcessId() (threadId, processId uint32, wErr error) {
	var pid uint32
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetWindowThreadProcessId, "GetWindowThreadProcessId"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&pid)))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, 0, wErr
	}
	return uint32(ret), pid, nil
}

var _GetWindowThreadProcessId *syscall.Proc

// Calls [HWND.GetWindowLongPtr] to retrieve the instance to which this window
// belongs.
func (hWnd HWND) HInstance() (HINSTANCE, error) {
	hInst, err := hWnd.GetWindowLongPtr(co.GWLP_HINSTANCE)
	return HINSTANCE(hInst), err
}

// [InvalidateRect] function.
//
// [InvalidateRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invalidaterect
func (hWnd HWND) InvalidateRect(rc *RECT, erase bool) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_InvalidateRect, "InvalidateRect"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)),
		utl.BoolToUintptr(erase))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _InvalidateRect *syscall.Proc

// [IsChild] function.
//
// [IsChild]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-ischild
func (hWnd HWND) IsChild(hChild HWND) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_IsChild, "IsChild"),
		uintptr(hWnd),
		uintptr(hChild))
	return ret != 0
}

var _IsChild *syscall.Proc

// Calls [HWND.GetClassLongPtr] and checks the [ATOM] against WC_DIALOG.
func (hWnd HWND) IsDialog() (bool, error) {
	uAtom, err := hWnd.GetClassLongPtr(co.GCL_ATOM)
	if err != nil {
		return false, err
	}
	return uint16(uAtom) == utl.WC_DIALOG, nil
}

// [IsDialogMessage] function.
//
// [IsDialogMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdialogmessagew
func (hWnd HWND) IsDialogMessage(msg *MSG) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_IsDialogMessageW, "IsDialogMessageW"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

var _IsDialogMessageW *syscall.Proc

// [IsIconic] function.
//
// [IsIconic]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isiconic
func (hWnd HWND) IsIconic() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_IsIconic, "IsIconic"),
		uintptr(hWnd))
	return ret != 0
}

var _IsIconic *syscall.Proc

// [IsWindow] function.
//
// [IsWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iswindow
func (hWnd HWND) IsWindow() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_IsWindow, "IsWindow"),
		uintptr(hWnd))
	return ret != 0
}

var _IsWindow *syscall.Proc

// [MapDialogRect] function.
//
// [MapDialogRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapdialogrect
func (hWnd HWND) MapDialogRect(rc *RECT) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_MapDialogRect, "MapDialogRect"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _MapDialogRect *syscall.Proc

// [MessageBox] function.
//
// [MessageBox]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
func (hWnd HWND) MessageBox(text, caption string, uType co.MB) (co.ID, error) {
	var wText, wCaption wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_MessageBoxW, "MessageBoxW"),
		uintptr(hWnd),
		uintptr(wText.EmptyIsNil(text)),
		uintptr(wCaption.EmptyIsNil(caption)),
		uintptr(uType))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return co.ID(0), wErr
	}
	return co.ID(ret), nil
}

var _MessageBoxW *syscall.Proc

// [MonitorFromWindow] function.
//
// [MonitorFromWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfromwindow
func (hWnd HWND) MonitorFromWindow(flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_MonitorFromWindow, "MonitorFromWindow"),
		uintptr(hWnd),
		uintptr(flags))
	return HMONITOR(ret)
}

var _MonitorFromWindow *syscall.Proc

// [PostMessage] function.
//
// [PostMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postmessagew
func (hWnd HWND) PostMessage(msg co.WM, wParam WPARAM, lParam LPARAM) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_PostMessageW, "PostMessageW"),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return utl.ZeroAsGetLastError(ret, err)
}

var _PostMessageW *syscall.Proc

// [RedrawWindow] function.
//
// [RedrawWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-redrawwindow
func (hWnd HWND) RedrawWindow(rcUpdate *RECT, hrgnUpdate HRGN, flags co.RDW) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_RedrawWindow, "RedrawWindow"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rcUpdate)),
		uintptr(hrgnUpdate),
		uintptr(flags))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _RedrawWindow *syscall.Proc

// [ReleaseDC] function.
//
// [ReleaseDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-releasedc
func (hWnd HWND) ReleaseDC(hdc HDC) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_ReleaseDC, "ReleaseDC"),
		uintptr(hWnd),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ReleaseDC *syscall.Proc

// [ScreenToClient] function for [POINT].
//
// [ScreenToClient]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientPt(pt *POINT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_ScreenToClient, "ScreenToClient"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pt)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ScreenToClient *syscall.Proc

// [ScreenToClient] function for [RECT].
//
// [ScreenToClient]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientRc(rc *RECT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_ScreenToClient, "ScreenToClient"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)))
	if ret == 0 {
		return co.ERROR_INVALID_PARAMETER
	}

	ret, _, _ = syscall.SyscallN(
		dll.Load(dll.USER32, &_ScreenToClient, "ScreenToClient"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&rc.Right)))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [SendMessage] function.
//
// [SendMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew
func (hWnd HWND) SendMessage(msg co.WM, wParam WPARAM, lParam LPARAM) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SendMessageW, "SendMessageW"),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	if wErr := co.ERROR(err); wErr == co.ERROR_ACCESS_DENIED {
		return 0, wErr
	}
	return ret, nil
}

var _SendMessageW *syscall.Proc

// [SetFocus] function.
//
// Returns a handle to the previously focused window.
//
// [SetFocus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setfocus
func (hWnd HWND) SetFocus() (HWND, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetFocus, "SetFocus"),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	} else {
		return HWND(ret), nil
	}
}

var _SetFocus *syscall.Proc

// [SetForegroundWindow] function.
//
// Returns true if the window was brought to the foreground.
//
// [SetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setforegroundwindow
func (hWnd HWND) SetForegroundWindow() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetForegroundWindow, "SetForegroundWindow"),
		uintptr(hWnd))
	return ret != 0
}

var _SetForegroundWindow *syscall.Proc

// [SetMenu] function.
//
// [SetMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenu
func (hWnd HWND) SetMenu(hMenu HMENU) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetMenu, "SetMenu"),
		uintptr(hWnd),
		uintptr(hMenu))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetMenu *syscall.Proc

// [SetParent] function.
//
// [SetParent]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setparent
func (hWnd HWND) SetParent(hNewParent HWND) (HWND, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetParent, "SetParent"),
		uintptr(hWnd),
		uintptr(hNewParent))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _SetParent *syscall.Proc

// [SetWindowDisplayAffinity] function.
//
// [SetWindowDisplayAffinity]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowdisplayaffinity
func (hWnd HWND) SetWindowDisplayAffinity(da co.WDA) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetWindowDisplayAffinity, "SetWindowDisplayAffinity"),
		uintptr(hWnd),
		uintptr(da))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetWindowDisplayAffinity *syscall.Proc

// [SetWindowPlacement] function.
//
// [SetWindowPlacement]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowplacement
func (hWnd HWND) SetWindowPlacement(wp *WINDOWPLACEMENT) error {
	wp.SetLength() // safety
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetWindowPlacement, "SetWindowPlacement"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(wp)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetWindowPlacement *syscall.Proc

// [SetWindowPos] function.
//
// [SetWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
func (hWnd HWND) SetWindowPos(hwndInsertAfter HWND, x, y, cx, cy int, flags co.SWP) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetWindowPos, "SetWindowPos"),
		uintptr(hWnd),
		uintptr(hwndInsertAfter),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(int32(cx)),
		uintptr(int32(cy)),
		uintptr(flags))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetWindowPos *syscall.Proc

// [SetWindowRgn] function.
//
// [SetWindowRgn]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowrgn
func (hWnd HWND) SetWindowRgn(hRgn HRGN, redraw bool) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetWindowRgn, "SetWindowRgn"),
		uintptr(hWnd),
		uintptr(hRgn),
		utl.BoolToUintptr(redraw))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _SetWindowRgn *syscall.Proc

// [SetWindowText] function.
//
// [SetWindowText]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowtextw
func (hWnd HWND) SetWindowText(text string) error {
	var wText wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_SetWindowTextW, "SetWindowTextW"),
		uintptr(hWnd),
		uintptr(wText.EmptyIsNil(text)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetWindowTextW *syscall.Proc

// [ShowCaret] function.
//
// [ShowCaret]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showcaret
func (hWnd HWND) ShowCaret() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_ShowCaret, "ShowCaret"),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

var _ShowCaret *syscall.Proc

// [ShowOwnedPopups] function.
//
// [ShowOwnedPopups]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showownedpopups
func (hWnd HWND) ShowOwnedPopups(show bool) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_ShowOwnedPopups, "ShowOwnedPopups"),
		uintptr(hWnd),
		utl.BoolToUintptr(show))
	return utl.ZeroAsGetLastError(ret, err)
}

var _ShowOwnedPopups *syscall.Proc

// [ShowScrollBar] function.
//
// [ShowScrollBar]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showscrollbar
func (hWnd HWND) ShowScrollBar(bar co.SB_SHOW, show bool) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_ShowScrollBar, "ShowScrollBar"),
		uintptr(hWnd),
		uintptr(bar),
		utl.BoolToUintptr(show))
	return utl.ZeroAsGetLastError(ret, err)
}

var _ShowScrollBar *syscall.Proc

// [ShowWindow] function.
//
// [ShowWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
func (hWnd HWND) ShowWindow(cmdShow co.SW) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_ShowWindow, "ShowWindow"),
		uintptr(hWnd),
		uintptr(cmdShow))
	return ret != 0
}

var _ShowWindow *syscall.Proc

// [ShowWindowAsync] function.
//
// [ShowWindowAsync]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindowasync
func (hWnd HWND) ShowWindowAsync(cmdShow co.SW) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_ShowWindowAsync, "ShowWindowAsync"),
		uintptr(hWnd),
		uintptr(cmdShow))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ShowWindowAsync *syscall.Proc

// [ShutdownBlockReasonCreate] function.
//
// [ShutdownBlockReasonCreate]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-shutdownblockreasoncreate
func (hWnd HWND) ShutdownBlockReasonCreate(reason string) error {
	var wReason wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_ShutdownBlockReasonCreate, "ShutdownBlockReasonCreate"),
		uintptr(hWnd),
		uintptr(wReason.AllowEmpty(reason)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _ShutdownBlockReasonCreate *syscall.Proc

// [ShutdownBlockReasonDestroy] function.
//
// [ShutdownBlockReasonDestroy]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-shutdownblockreasondestroy
func (hWnd HWND) ShutdownBlockReasonDestroy() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_ShutdownBlockReasonDestroy, "ShutdownBlockReasonDestroy"),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

var _ShutdownBlockReasonDestroy *syscall.Proc

// [ShutdownBlockReasonQuery] function.
//
// [ShutdownBlockReasonQuery]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-shutdownblockreasonquery
func (hWnd HWND) ShutdownBlockReasonQuery() (string, error) {
	var bufSz uint32
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_ShutdownBlockReasonQuery, "ShutdownBlockReasonQuery"),
		uintptr(hWnd),
		0,
		uintptr(unsafe.Pointer(&bufSz)))
	if ret == 0 {
		return "", co.ERROR(err)
	}

	var wBuf wstr.BufDecoder
	wBuf.Alloc(int(bufSz))

	ret, _, err = syscall.SyscallN(
		dll.Load(dll.USER32, &_ShutdownBlockReasonQuery, "ShutdownBlockReasonQuery"),
		uintptr(hWnd),
		uintptr(wBuf.Ptr()),
		uintptr(unsafe.Pointer(&bufSz)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wBuf.String(), nil
}

var _ShutdownBlockReasonQuery *syscall.Proc

// Calls [HWND.GetWindowLongPtr] to retrieve the window style.
func (hWnd HWND) Style() (co.WS, error) {
	style, err := hWnd.GetWindowLongPtr(co.GWLP_STYLE)
	return co.WS(style), err
}

// [TranslateAccelerator] function.
//
// [TranslateAccelerator]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translateacceleratorw
func (hWnd HWND) TranslateAccelerator(hAccel HACCEL, msg *MSG) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_TranslateAcceleratorW, "TranslateAcceleratorW"),
		uintptr(hWnd),
		uintptr(hAccel),
		uintptr(unsafe.Pointer(msg)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _TranslateAcceleratorW *syscall.Proc

// [UpdateWindow] function.
//
// [UpdateWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-updatewindow
func (hWnd HWND) UpdateWindow() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_UpdateWindow, "UpdateWindow"),
		uintptr(hWnd))
	return ret != 0
}

var _UpdateWindow *syscall.Proc
