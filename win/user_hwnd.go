//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
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
	x, y int,
	width, height uint,
	parent HWND,
	menu HMENU,
	instance HINSTANCE,
	param LPARAM,
) (HWND, error) {
	className16 := wstr.NewBuf[wstr.Stack20]()
	classNameVal := className.raw(&className16)

	title16 := wstr.NewBufWith[wstr.Stack20](title, wstr.EMPTY_IS_NIL)

	ret, _, err := syscall.SyscallN(_CreateWindowExW.Addr(),
		uintptr(exStyle),
		classNameVal,
		uintptr(title16.UnsafePtr()),
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

var _CreateWindowExW = dll.User32.NewProc("CreateWindowExW")

// [FindWindow] function.
//
// [FindWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-findwindoww
func FindWindow(className ClassName, title string) (HWND, bool) {
	className16 := wstr.NewBuf[wstr.Stack20]()
	classNameVal := className.raw(&className16)

	title16 := wstr.NewBufWith[wstr.Stack20](title, wstr.EMPTY_IS_NIL)

	ret, _, _ := syscall.SyscallN(_FindWindowW.Addr(),
		classNameVal,
		uintptr(title16.UnsafePtr()))
	return HWND(ret), ret != 0
}

var _FindWindowW = dll.User32.NewProc("FindWindowW")

// [GetClipboardOwner] function.
//
// [GetClipboardOwner]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardowner
func GetClipboardOwner() (HWND, error) {
	ret, _, err := syscall.SyscallN(_GetClipboardOwner.Addr())
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetClipboardOwner = dll.User32.NewProc("GetClipboardOwner")

// [GetDesktopWindow] function.
//
// [GetDesktopWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdesktopwindow
func GetDesktopWindow() HWND {
	ret, _, _ := syscall.SyscallN(_GetDesktopWindow.Addr())
	return HWND(ret)
}

var _GetDesktopWindow = dll.User32.NewProc("GetDesktopWindow")

// [GetFocus] function.
//
// [GetFocus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getfocus
func GetFocus() HWND {
	ret, _, _ := syscall.SyscallN(_GetFocus.Addr())
	return HWND(ret)
}

var _GetFocus = dll.User32.NewProc("GetFocus")

// [GetForegroundWindow] function.
//
// [GetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
func GetForegroundWindow() HWND {
	ret, _, _ := syscall.SyscallN(_GetForegroundWindow.Addr())
	return HWND(ret)
}

var _GetForegroundWindow = dll.User32.NewProc("GetForegroundWindow")

// [GetOpenClipboardWindow] function.
//
// [GetOpenClipboardWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getopenclipboardwindow
func GetOpenClipboardWindow() (HWND, error) {
	ret, _, err := syscall.SyscallN(_GetOpenClipboardWindow.Addr())
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetOpenClipboardWindow = dll.User32.NewProc("GetOpenClipboardWindow")

// [GetShellWindow] function.
//
// [GetShellWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getshellwindow
func GetShellWindow() HWND {
	ret, _, _ := syscall.SyscallN(_GetShellWindow.Addr())
	return HWND(ret)
}

var _GetShellWindow = dll.User32.NewProc("GetShellWindow")

// [AnimateWindow] function.
//
// [AnimateWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-animatewindow
func (hWnd HWND) AnimateWindow(time uint, flags co.AW) error {
	ret, _, err := syscall.SyscallN(_AnimateWindow.Addr(),
		uintptr(hWnd),
		uintptr(time),
		uintptr(flags))
	return utl.ZeroAsGetLastError(ret, err)
}

var _AnimateWindow = dll.User32.NewProc("AnimateWindow")

// [BeginPaint] function.
//
// ⚠️ You must defer [HWND.EndPaint].
//
// # Example
//
//	var hWnd win.HWND // initialized somewhere
//
//	var ps win.PAINTSTRUCT
//	hdc, _ := hWnd.BeginPaint(&ps)
//	defer hWnd.EndPaint(&ps)
//
// [BeginPaint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-beginpaint
func (hWnd HWND) BeginPaint(ps *PAINTSTRUCT) (HDC, error) {
	ret, _, _ := syscall.SyscallN(_BeginPaint.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(ps)))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _BeginPaint = dll.User32.NewProc("BeginPaint")

// [BringWindowToTop] function.
//
// [BringWindowToTop]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-bringwindowtotop
func (hWnd HWND) BringWindowToTop() error {
	ret, _, err := syscall.SyscallN(_BringWindowToTop.Addr(),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

var _BringWindowToTop = dll.User32.NewProc("BringWindowToTop")

// [ChildWindowFromPoint] function.
//
// [ChildWindowFromPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-childwindowfrompoint
func (hWnd HWND) ChildWindowFromPoint(pt POINT) (HWND, bool) {
	ret, _, _ := syscall.SyscallN(_ChildWindowFromPoint.Addr(),
		uintptr(hWnd),
		uintptr(pt.X),
		uintptr(pt.Y))
	if ret == 0 {
		return HWND(0), false
	}
	return HWND(ret), true
}

var _ChildWindowFromPoint = dll.User32.NewProc("ChildWindowFromPoint")

// [ChildWindowFromPointEx] function.
//
// [ChildWindowFromPointEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-childwindowfrompointex
func (hWnd HWND) ChildWindowFromPointEx(pt POINT, flags co.CWP) (HWND, bool) {
	ret, _, _ := syscall.SyscallN(_ChildWindowFromPointEx.Addr(),
		uintptr(hWnd),
		uintptr(pt.X),
		uintptr(pt.Y),
		uintptr(flags))
	if ret == 0 {
		return HWND(0), false
	}
	return HWND(ret), true
}

var _ChildWindowFromPointEx = dll.User32.NewProc("ChildWindowFromPointEx")

// [ClientToScreen] function for [POINT].
//
// [ClientToScreen]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenPt(pt *POINT) error {
	ret, _, _ := syscall.SyscallN(_ClientToScreen.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pt)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ClientToScreen = dll.User32.NewProc("ClientToScreen")

// [ClientToScreen] function for [RECT].
//
// [ClientToScreen]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenRc(rc *RECT) error {
	ret, _, _ := syscall.SyscallN(_ClientToScreen.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)))
	if ret == 0 {
		return co.ERROR_INVALID_PARAMETER
	}

	ret, _, _ = syscall.SyscallN(_ClientToScreen.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(&rc.Right)))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [CloseWindow] function.
//
// Note that this function will actually minimize the window, not destroy it.
//
// [CloseWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-closewindow
func (hWnd HWND) CloseWindow() error {
	ret, _, _ := syscall.SyscallN(_CloseWindow.Addr(),
		uintptr(hWnd))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _CloseWindow = dll.User32.NewProc("CloseWindow")

// [DefDlgProc] function.
//
// [DefDlgProc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defdlgprocw
func (hWnd HWND) DefDlgProc(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.SyscallN(_DefDlgProcW.Addr(),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return ret
}

var _DefDlgProcW = dll.User32.NewProc("DefDlgProcW")

// [DefWindowProc] function.
//
// [DefWindowProc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defwindowprocw
func (hWnd HWND) DefWindowProc(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.SyscallN(_DefWindowProcW.Addr(),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return ret
}

var _DefWindowProcW = dll.User32.NewProc("DefWindowProcW")

// [DestroyWindow] function.
//
// Note: don't call this function to close a window. The correct way to close a
// window is sending [co.WM_CLOSE].
//
// [DestroyWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroywindow
func (hWnd HWND) DestroyWindow() error {
	ret, _, err := syscall.SyscallN(_DestroyWindow.Addr(),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

var _DestroyWindow = dll.User32.NewProc("DestroyWindow")

// [DrawMenuBar] function.
//
// [DrawMenuBar]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawmenubar
func (hWnd HWND) DrawMenuBar() error {
	ret, _, err := syscall.SyscallN(_DrawMenuBar.Addr(),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

var _DrawMenuBar = dll.User32.NewProc("DrawMenuBar")

// [EnableWindow] function.
//
// [EnableWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enablewindow
func (hWnd HWND) EnableWindow(enable bool) bool {
	ret, _, _ := syscall.SyscallN(_EnableWindow.Addr(),
		uintptr(hWnd), utl.BoolToUintptr(enable))
	return ret != 0 // the window was previously disabled?
}

var _EnableWindow = dll.User32.NewProc("EnableWindow")

// [EndDialog] function.
//
// [EndDialog]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddialog
func (hWnd HWND) EndDialog(result uintptr) error {
	ret, _, err := syscall.SyscallN(_EndDialog.Addr(),
		uintptr(hWnd),
		result)
	return utl.ZeroAsGetLastError(ret, err)
}

var _EndDialog = dll.User32.NewProc("EndDialog")

// [EndPaint] function.
//
// Paired with [HWND.BeginPaint].
//
// [EndPaint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endpaint
func (hWnd HWND) EndPaint(ps *PAINTSTRUCT) {
	syscall.SyscallN(_EndPaint.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(ps)))
}

var _EndPaint = dll.User32.NewProc("EndPaint")

// [EnumChildWindows] function.
//
// [EnumChildWindows]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumchildwindows
func (hWnd HWND) EnumChildWindows() []HWND {
	pPack := &_EnumChildWindowsPack{
		arr: make([]HWND, 0),
	}

	syscall.SyscallN(_EnumChildWindows.Addr(),
		uintptr(hWnd),
		enumChildWindowsCallback(),
		uintptr(unsafe.Pointer(pPack)))
	return pPack.arr
}

type _EnumChildWindowsPack struct{ arr []HWND }

var (
	_EnumChildWindows         = dll.User32.NewProc("EnumChildWindows")
	_enumChildWindowsCallback uintptr
)

func enumChildWindowsCallback() uintptr {
	if _enumChildWindowsCallback == 0 {
		_enumChildWindowsCallback = syscall.NewCallback(
			func(hChild HWND, lParam LPARAM) uintptr {
				pPack := (*_EnumChildWindowsPack)(unsafe.Pointer(lParam))
				pPack.arr = append(pPack.arr, hChild)
				return 1
			},
		)
	}
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
	ret, _, err := syscall.SyscallN(_GetAncestor.Addr(),
		uintptr(hWnd),
		uintptr(gaFlags))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetAncestor = dll.User32.NewProc("GetAncestor")

// [GetClassName] function.
//
// [GetClassName]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassnamew
func (hWnd HWND) GetClassName() (string, error) {
	var buf [256 + 1]uint16
	ret, _, err := syscall.SyscallN(_GetClassNameW.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(int32(len(buf))))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return "", wErr
	}
	return wstr.WstrSliceToStr(buf[:]), nil
}

var _GetClassNameW = dll.User32.NewProc("GetClassNameW")

// [GetClientRect] function.
//
// [GetClientRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclientrect
func (hWnd HWND) GetClientRect() (RECT, error) {
	var rc RECT
	ret, _, err := syscall.SyscallN(_GetClientRect.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		return RECT{}, co.ERROR(err)
	}
	return rc, nil
}

var _GetClientRect = dll.User32.NewProc("GetClientRect")

// [GetDC] function.
//
// ⚠️ You must defer [HWND.ReleaseDC].
//
// # Example
//
// Retrieving the DC of the entire screen:
//
//	hdcScreen, _ := win.HWND(0).GetDC()
//	defer win.HWND(0).ReleaseDC(hdcScreen)
//
// [GetDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
func (hWnd HWND) GetDC() (HDC, error) {
	ret, _, _ := syscall.SyscallN(_GetDC.Addr(),
		uintptr(hWnd))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _GetDC = dll.User32.NewProc("GetDC")

// [GetDCEx] function.
//
// ⚠️ You must defer [HWND.ReleaseDC].
//
// [GetDCEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
func (hWnd HWND) GetDCEx(hRgnClip HRGN, flags co.DCX) (HDC, error) {
	ret, _, _ := syscall.SyscallN(_GetDCEx.Addr(),
		uintptr(hWnd),
		uintptr(hRgnClip),
		uintptr(flags))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _GetDCEx = dll.User32.NewProc("GetDCEx")

// [GetDlgCtrlID] function.
//
// [GetDlgCtrlID]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgctrlid
func (hWnd HWND) GetDlgCtrlID() (uint16, error) {
	ret, _, err := syscall.SyscallN(_GetDlgCtrlID.Addr(),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint16(ret), nil
}

var _GetDlgCtrlID = dll.User32.NewProc("GetDlgCtrlID")

// [GetDlgItem] function.
//
// [GetDlgItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgitem
func (hWnd HWND) GetDlgItem(dlgId uint16) (HWND, error) {
	ret, _, err := syscall.SyscallN(_GetDlgItem.Addr(),
		uintptr(hWnd),
		uintptr(int32(dlgId)))
	if ret == 0 {
		return HWND(0), co.ERROR(err)
	}
	return HWND(ret), nil
}

var _GetDlgItem = dll.User32.NewProc("GetDlgItem")

// [GetLastActivePopup] function.
//
// [GetLastActivePopup]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getlastactivepopup
func (hWnd HWND) GetLastActivePopup() HWND {
	ret, _, _ := syscall.SyscallN(_GetLastActivePopup.Addr(),
		uintptr(hWnd))
	return HWND(ret)
}

var _GetLastActivePopup = dll.User32.NewProc("GetLastActivePopup")

// [GetMenu] function.
//
// [GetMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenu
func (hWnd HWND) GetMenu() HMENU {
	ret, _, _ := syscall.SyscallN(_GetMenu.Addr(),
		uintptr(hWnd))
	return HMENU(ret)
}

var _GetMenu = dll.User32.NewProc("GetMenu")

// [GetNextDlgGroupItem] function.
//
// [GetNextDlgGroupItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlggroupitem
func (hWnd HWND) GetNextDlgGroupItem(hChild HWND, isPrevious bool) (HWND, error) {
	ret, _, err := syscall.SyscallN(_GetNextDlgGroupItem.Addr(),
		uintptr(hWnd),
		uintptr(hChild),
		utl.BoolToUintptr(isPrevious))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetNextDlgGroupItem = dll.User32.NewProc("GetNextDlgGroupItem")

// [GetNextDlgTabItem] function.
//
// [GetNextDlgTabItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlgtabitem
func (hWnd HWND) GetNextDlgTabItem(hChild HWND, isPrevious bool) (HWND, error) {
	ret, _, err := syscall.SyscallN(_GetNextDlgTabItem.Addr(),
		uintptr(hWnd),
		uintptr(hChild),
		utl.BoolToUintptr(isPrevious))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetNextDlgTabItem = dll.User32.NewProc("GetNextDlgTabItem")

// [GetParent] function.
//
// [GetParent]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getparent
func (hWnd HWND) GetParent() (HWND, error) {
	ret, _, err := syscall.SyscallN(_GetParent.Addr(),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetParent = dll.User32.NewProc("GetParent")

// [GetWindow] function.
//
// [GetWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindow
func (hWnd HWND) GetWindow(cmd co.GW) (HWND, error) {
	ret, _, err := syscall.SyscallN(_GetWindow.Addr(),
		uintptr(hWnd),
		uintptr(cmd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

var _GetWindow = dll.User32.NewProc("GetWindow")

// [GetWindowDC] function.
//
// ⚠️ You must defer [HWND.ReleaseDC].
//
// [GetWindowDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowdc
func (hWnd HWND) GetWindowDC() (HDC, error) {
	ret, _, _ := syscall.SyscallN(_GetWindowDC.Addr(),
		uintptr(hWnd))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

var _GetWindowDC = dll.User32.NewProc("GetWindowDC")

// [GetWindowRect] function.
//
// [GetWindowRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowrect
func (hWnd HWND) GetWindowRect() (RECT, error) {
	var rc RECT
	ret, _, err := syscall.SyscallN(_GetWindowRect.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		return RECT{}, co.ERROR(err)
	}
	return rc, nil
}

var _GetWindowRect = dll.User32.NewProc("GetWindowRect")

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

	buf := wstr.NewBufSized[wstr.Stack64](bufSz)

	ret, _, err := syscall.SyscallN(_GetWindowTextW.Addr(),
		uintptr(hWnd),
		uintptr(buf.UnsafePtr()),
		uintptr(int32(bufSz)))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return "", wErr
	}
	return wstr.WstrSliceToStr(buf.HotSlice()), nil
}

var _GetWindowTextW = dll.User32.NewProc("GetWindowTextW")

// [GetWindowTextLength] function.
//
// You usually don't need to call this function since [HWND.GetWindowText] already
// calls it to allocate the memory block.
//
// [GetWindowTextLength]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextlengthw
func (hWnd HWND) GetWindowTextLength() (uint, error) {
	ret, _, err := syscall.SyscallN(_GetWindowTextLengthW.Addr(),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint(ret), nil
}

var _GetWindowTextLengthW = dll.User32.NewProc("GetWindowTextLengthW")

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
	ret, _, _ := syscall.SyscallN(_InvalidateRect.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)),
		utl.BoolToUintptr(erase))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _InvalidateRect = dll.User32.NewProc("InvalidateRect")

// [IsChild] function.
//
// [IsChild]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-ischild
func (hWnd HWND) IsChild(hChild HWND) bool {
	ret, _, _ := syscall.SyscallN(_IsChild.Addr(),
		uintptr(hWnd),
		uintptr(hChild))
	return ret != 0
}

var _IsChild = dll.User32.NewProc("IsChild")

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
	ret, _, _ := syscall.SyscallN(_IsDialogMessageW.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

var _IsDialogMessageW = dll.User32.NewProc("IsDialogMessage")

// [IsWindow] function.
//
// [IsWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iswindow
func (hWnd HWND) IsWindow() bool {
	ret, _, _ := syscall.SyscallN(_IsWindow.Addr(),
		uintptr(hWnd))
	return ret != 0
}

var _IsWindow = dll.User32.NewProc("IsWindow")

// [MapDialogRect] function.
//
// [MapDialogRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapdialogrect
func (hWnd HWND) MapDialogRect(rc *RECT) error {
	ret, _, err := syscall.SyscallN(_MapDialogRect.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _MapDialogRect = dll.User32.NewProc("MapDialogRect")

// [MessageBox] function.
//
// [MessageBox]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
func (hWnd HWND) MessageBox(text, caption string, uType co.MB) (co.ID, error) {
	text16 := wstr.NewBufWith[wstr.Stack20](text, wstr.EMPTY_IS_NIL)
	caption16 := wstr.NewBufWith[wstr.Stack20](caption, wstr.EMPTY_IS_NIL)

	ret, _, err := syscall.SyscallN(_MessageBoxW.Addr(),
		uintptr(hWnd),
		uintptr(text16.UnsafePtr()),
		uintptr(caption16.UnsafePtr()),
		uintptr(uType))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return co.ID(0), wErr
	}
	return co.ID(ret), nil
}

var _MessageBoxW = dll.User32.NewProc("MessageBoxW")

// [MonitorFromWindow] function.
//
// [MonitorFromWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfromwindow
func (hWnd HWND) MonitorFromWindow(flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.SyscallN(_MonitorFromWindow.Addr(),
		uintptr(hWnd),
		uintptr(flags))
	return HMONITOR(ret)
}

var _MonitorFromWindow = dll.User32.NewProc("MonitorFromWindow")

// [PostMessage] function.
//
// [PostMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postmessagew
func (hWnd HWND) PostMessage(msg co.WM, wParam WPARAM, lParam LPARAM) error {
	ret, _, err := syscall.SyscallN(_PostMessageW.Addr(),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return utl.ZeroAsGetLastError(ret, err)
}

var _PostMessageW = dll.User32.NewProc("PostMessageW")

// [RedrawWindow] function.
//
// [RedrawWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-redrawwindow
func (hWnd HWND) RedrawWindow(rcUpdate *RECT, hrgnUpdate HRGN, flags co.RDW) error {
	ret, _, _ := syscall.SyscallN(_RedrawWindow.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rcUpdate)),
		uintptr(hrgnUpdate),
		uintptr(flags))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _RedrawWindow = dll.User32.NewProc("RedrawWindow")

// [ReleaseDC] function.
//
// [ReleaseDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-releasedc
func (hWnd HWND) ReleaseDC(hdc HDC) error {
	ret, _, _ := syscall.SyscallN(_ReleaseDC.Addr(),
		uintptr(hWnd),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ReleaseDC = dll.User32.NewProc("ReleaseDC")

// [ScreenToClient] function for [POINT].
//
// [ScreenToClient]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientPt(pt *POINT) error {
	ret, _, _ := syscall.SyscallN(_ScreenToClient.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pt)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ScreenToClient = dll.User32.NewProc("ScreenToClient")

// [ScreenToClient] function for [RECT].
//
// [ScreenToClient]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientRc(rc *RECT) error {
	ret, _, _ := syscall.SyscallN(_ScreenToClient.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)))
	if ret == 0 {
		return co.ERROR_INVALID_PARAMETER
	}

	ret, _, _ = syscall.SyscallN(_ScreenToClient.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&rc.Right)))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [SendMessage] function.
//
// [SendMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew
func (hWnd HWND) SendMessage(msg co.WM, wParam WPARAM, lParam LPARAM) (uintptr, error) {
	ret, _, err := syscall.SyscallN(_SendMessageW.Addr(),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	if wErr := co.ERROR(err); wErr == co.ERROR_ACCESS_DENIED {
		return 0, wErr
	}
	return ret, nil
}

var _SendMessageW = dll.User32.NewProc("SendMessageW")

// [SetFocus] function.
//
// Returns a handle to the previously focused window.
//
// [SetFocus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setfocus
func (hWnd HWND) SetFocus() (HWND, error) {
	ret, _, err := syscall.SyscallN(_SetFocus.Addr(),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	} else {
		return HWND(ret), nil
	}
}

var _SetFocus = dll.User32.NewProc("SetFocus")

// [SetForegroundWindow] function.
//
// Returns true if the window was brought to the foreground.
//
// [SetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setforegroundwindow
func (hWnd HWND) SetForegroundWindow() bool {
	ret, _, _ := syscall.SyscallN(_SetForegroundWindow.Addr(),
		uintptr(hWnd))
	return ret != 0
}

var _SetForegroundWindow = dll.User32.NewProc("SetForegroundWindow")

// [SetMenu] function.
//
// [SetMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenu
func (hWnd HWND) SetMenu(hMenu HMENU) error {
	ret, _, err := syscall.SyscallN(_SetMenu.Addr(),
		uintptr(hWnd),
		uintptr(hMenu))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetMenu = dll.User32.NewProc("SetMenu")

// [SetWindowPos] function.
//
// [SetWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
func (hWnd HWND) SetWindowPos(hwndInsertAfter HWND, x, y int, cx, cy uint, flags co.SWP) error {
	ret, _, err := syscall.SyscallN(_SetWindowPos.Addr(),
		uintptr(hWnd),
		uintptr(hwndInsertAfter),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(int32(cx)),
		uintptr(int32(cy)),
		uintptr(flags))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetWindowPos = dll.User32.NewProc("SetWindowPos")

// [SetWindowRgn] function.
//
// [SetWindowRgn]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowrgn
func (hWnd HWND) SetWindowRgn(hRgn HRGN, redraw bool) error {
	ret, _, _ := syscall.SyscallN(_SetWindowRgn.Addr(),
		uintptr(hWnd),
		uintptr(hRgn),
		utl.BoolToUintptr(redraw))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _SetWindowRgn = dll.User32.NewProc("SetWindowRgn")

// [SetWindowText] function.
//
// [SetWindowText]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowtextw
func (hWnd HWND) SetWindowText(text string) {
	text16 := wstr.NewBufWith[wstr.Stack20](text, wstr.EMPTY_IS_NIL)
	syscall.SyscallN(_SetWindowTextW.Addr(),
		uintptr(hWnd),
		uintptr(text16.UnsafePtr()))
}

var _SetWindowTextW = dll.User32.NewProc("SetWindowTextW")

// [ShowCaret] function.
//
// [ShowCaret]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showcaret
func (hWnd HWND) ShowCaret() error {
	ret, _, err := syscall.SyscallN(_ShowCaret.Addr(),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

var _ShowCaret = dll.User32.NewProc("ShowCaret")

// [ShowWindow] function.
//
// [ShowWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
func (hWnd HWND) ShowWindow(cmdShow co.SW) bool {
	ret, _, _ := syscall.SyscallN(_ShowWindow.Addr(),
		uintptr(hWnd),
		uintptr(cmdShow))
	return ret != 0
}

var _ShowWindow = dll.User32.NewProc("ShowWindow")

// Calls [HWND.GetWindowLongPtr] to retrieve the window style.
func (hWnd HWND) Style() (co.WS, error) {
	style, err := hWnd.GetWindowLongPtr(co.GWLP_STYLE)
	return co.WS(style), err
}

// [TranslateAccelerator] function.
//
// [TranslateAccelerator]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translateacceleratorw
func (hWnd HWND) TranslateAccelerator(hAccel HACCEL, msg *MSG) error {
	ret, _, err := syscall.SyscallN(_TranslateAcceleratorW.Addr(),
		uintptr(hWnd),
		uintptr(hAccel),
		uintptr(unsafe.Pointer(msg)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _TranslateAcceleratorW = dll.User32.NewProc("TranslateAcceleratorW")

// [UpdateWindow] function.
//
// [UpdateWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-updatewindow
func (hWnd HWND) UpdateWindow() bool {
	ret, _, _ := syscall.SyscallN(_UpdateWindow.Addr(),
		uintptr(hWnd))
	return ret != 0
}

var _UpdateWindow = dll.User32.NewProc("UpdateWindow")
