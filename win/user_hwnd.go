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
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pClassName := className.raw(&wbuf)
	pTitle := wbuf.PtrEmptyIsNil(title)

	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_CreateWindowExW),
		uintptr(exStyle),
		pClassName,
		uintptr(pTitle),
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

// [FindWindow] function.
//
// [FindWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-findwindoww
func FindWindow(className ClassName, title string) (HWND, bool) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pClassName := className.raw(&wbuf)
	pTitle := wbuf.PtrEmptyIsNil(title)

	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_FindWindowW),
		pClassName,
		uintptr(pTitle))
	return HWND(ret), ret != 0
}

// [GetClipboardOwner] function.
//
// [GetClipboardOwner]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardowner
func GetClipboardOwner() (HWND, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetClipboardOwner))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

// [GetDesktopWindow] function.
//
// [GetDesktopWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdesktopwindow
func GetDesktopWindow() HWND {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetDesktopWindow))
	return HWND(ret)
}

// [GetFocus] function.
//
// [GetFocus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getfocus
func GetFocus() HWND {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetFocus))
	return HWND(ret)
}

// [GetForegroundWindow] function.
//
// [GetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
func GetForegroundWindow() HWND {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetForegroundWindow))
	return HWND(ret)
}

// [GetOpenClipboardWindow] function.
//
// [GetOpenClipboardWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getopenclipboardwindow
func GetOpenClipboardWindow() (HWND, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetOpenClipboardWindow))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

// [GetShellWindow] function.
//
// [GetShellWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getshellwindow
func GetShellWindow() HWND {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetShellWindow))
	return HWND(ret)
}

// [AnimateWindow] function.
//
// [AnimateWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-animatewindow
func (hWnd HWND) AnimateWindow(time uint, flags co.AW) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_AnimateWindow),
		uintptr(hWnd),
		uintptr(time),
		uintptr(flags))
	return utl.ZeroAsGetLastError(ret, err)
}

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
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_BeginPaint),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(ps)))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

// [BringWindowToTop] function.
//
// [BringWindowToTop]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-bringwindowtotop
func (hWnd HWND) BringWindowToTop() error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_BringWindowToTop),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

// [ChildWindowFromPoint] function.
//
// [ChildWindowFromPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-childwindowfrompoint
func (hWnd HWND) ChildWindowFromPoint(pt POINT) (HWND, bool) {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_ChildWindowFromPoint),
		uintptr(hWnd),
		uintptr(pt.X),
		uintptr(pt.Y))
	if ret == 0 {
		return HWND(0), false
	}
	return HWND(ret), true
}

// [ChildWindowFromPointEx] function.
//
// [ChildWindowFromPointEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-childwindowfrompointex
func (hWnd HWND) ChildWindowFromPointEx(pt POINT, flags co.CWP) (HWND, bool) {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_ChildWindowFromPointEx),
		uintptr(hWnd),
		uintptr(pt.X),
		uintptr(pt.Y),
		uintptr(flags))
	if ret == 0 {
		return HWND(0), false
	}
	return HWND(ret), true
}

// [ClientToScreen] function for [POINT].
//
// [ClientToScreen]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenPt(pt *POINT) error {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_ClientToScreen),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pt)))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [ClientToScreen] function for [RECT].
//
// [ClientToScreen]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenRc(rc *RECT) error {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_ClientToScreen),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)))
	if ret == 0 {
		return co.ERROR_INVALID_PARAMETER
	}

	ret, _, _ = syscall.SyscallN(dll.User(dll.PROC_ClientToScreen),
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
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_CloseWindow),
		uintptr(hWnd))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [DefDlgProc] function.
//
// [DefDlgProc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defdlgprocw
func (hWnd HWND) DefDlgProc(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_DefDlgProcW),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return ret
}

// [DefWindowProc] function.
//
// [DefWindowProc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defwindowprocw
func (hWnd HWND) DefWindowProc(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_DefWindowProcW),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return ret
}

// [DestroyWindow] function.
//
// Note: don't call this function to close a window. The correct way to close a
// window is sending [co.WM_CLOSE].
//
// [DestroyWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroywindow
func (hWnd HWND) DestroyWindow() error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_DestroyWindow),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

// [DrawMenuBar] function.
//
// [DrawMenuBar]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawmenubar
func (hWnd HWND) DrawMenuBar() error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_DrawMenuBar),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

// [EnableWindow] function.
//
// [EnableWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enablewindow
func (hWnd HWND) EnableWindow(enable bool) bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_EnableWindow),
		uintptr(hWnd), utl.BoolToUintptr(enable))
	return ret != 0 // the window was previously disabled?
}

// [EndDialog] function.
//
// [EndDialog]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddialog
func (hWnd HWND) EndDialog(result uintptr) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_EndDialog),
		uintptr(hWnd),
		result)
	return utl.ZeroAsGetLastError(ret, err)
}

// [EndPaint] function.
//
// Paired with [HWND.BeginPaint].
//
// [EndPaint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endpaint
func (hWnd HWND) EndPaint(ps *PAINTSTRUCT) {
	syscall.SyscallN(dll.User(dll.PROC_EndPaint),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(ps)))
}

// [EnumChildWindows] function.
//
// [EnumChildWindows]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumchildwindows
func (hWnd HWND) EnumChildWindows() []HWND {
	pPack := &_EnumChildWindowsPack{
		arr: make([]HWND, 0),
	}

	syscall.SyscallN(dll.User(dll.PROC_EnumChildWindows),
		uintptr(hWnd),
		enumChildWindowsCallback(),
		uintptr(unsafe.Pointer(pPack)))
	return pPack.arr
}

type _EnumChildWindowsPack struct{ arr []HWND }

var _enumChildWindowsCallback uintptr

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
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetAncestor),
		uintptr(hWnd),
		uintptr(gaFlags))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

// [GetClassName] function.
//
// [GetClassName]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassnamew
func (hWnd HWND) GetClassName() (string, error) {
	recvBuf := wstr.NewBufReceiver(wstr.BUF_MAX)
	defer recvBuf.Free()

	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetClassNameW),
		uintptr(hWnd),
		uintptr(recvBuf.UnsafePtr()),
		uintptr(int32(recvBuf.Len())))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return "", wErr
	}
	return recvBuf.String(), nil
}

// [GetClientRect] function.
//
// [GetClientRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclientrect
func (hWnd HWND) GetClientRect() (RECT, error) {
	var rc RECT
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetClientRect),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		return RECT{}, co.ERROR(err)
	}
	return rc, nil
}

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
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetDC),
		uintptr(hWnd))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

// [GetDCEx] function.
//
// ⚠️ You must defer [HWND.ReleaseDC].
//
// [GetDCEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
func (hWnd HWND) GetDCEx(hRgnClip HRGN, flags co.DCX) (HDC, error) {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetDCEx),
		uintptr(hWnd),
		uintptr(hRgnClip),
		uintptr(flags))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

// [GetDlgCtrlID] function.
//
// [GetDlgCtrlID]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgctrlid
func (hWnd HWND) GetDlgCtrlID() (uint16, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetDlgCtrlID),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint16(ret), nil
}

// [GetDlgItem] function.
//
// [GetDlgItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgitem
func (hWnd HWND) GetDlgItem(dlgId uint16) (HWND, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetDlgItem),
		uintptr(hWnd),
		uintptr(int32(dlgId)))
	if ret == 0 {
		return HWND(0), co.ERROR(err)
	}
	return HWND(ret), nil
}

// [GetLastActivePopup] function.
//
// [GetLastActivePopup]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getlastactivepopup
func (hWnd HWND) GetLastActivePopup() HWND {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetLastActivePopup),
		uintptr(hWnd))
	return HWND(ret)
}

// [GetMenu] function.
//
// [GetMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenu
func (hWnd HWND) GetMenu() HMENU {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetMenu),
		uintptr(hWnd))
	return HMENU(ret)
}

// [GetNextDlgGroupItem] function.
//
// [GetNextDlgGroupItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlggroupitem
func (hWnd HWND) GetNextDlgGroupItem(hChild HWND, isPrevious bool) (HWND, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetNextDlgGroupItem),
		uintptr(hWnd),
		uintptr(hChild),
		utl.BoolToUintptr(isPrevious))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

// [GetNextDlgTabItem] function.
//
// [GetNextDlgTabItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlgtabitem
func (hWnd HWND) GetNextDlgTabItem(hChild HWND, isPrevious bool) (HWND, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetNextDlgTabItem),
		uintptr(hWnd),
		uintptr(hChild),
		utl.BoolToUintptr(isPrevious))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

// [GetParent] function.
//
// [GetParent]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getparent
func (hWnd HWND) GetParent() (HWND, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetParent),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

// [GetWindow] function.
//
// [GetWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindow
func (hWnd HWND) GetWindow(cmd co.GW) (HWND, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetWindow),
		uintptr(hWnd),
		uintptr(cmd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	}
	return HWND(ret), nil
}

// [GetWindowDC] function.
//
// ⚠️ You must defer [HWND.ReleaseDC].
//
// [GetWindowDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowdc
func (hWnd HWND) GetWindowDC() (HDC, error) {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_GetWindowDC),
		uintptr(hWnd))
	if ret == 0 {
		return HDC(0), co.ERROR_INVALID_PARAMETER
	}
	return HDC(ret), nil
}

// [GetWindowRect] function.
//
// [GetWindowRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowrect
func (hWnd HWND) GetWindowRect() (RECT, error) {
	var rc RECT
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetWindowRect),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		return RECT{}, co.ERROR(err)
	}
	return rc, nil
}

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

	recvBuf := wstr.NewBufReceiver(bufSz)
	defer recvBuf.Free()

	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetWindowTextW),
		uintptr(hWnd),
		uintptr(recvBuf.UnsafePtr()),
		uintptr(int32(bufSz)))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return "", wErr
	}
	return recvBuf.String(), nil
}

// [GetWindowTextLength] function.
//
// You usually don't need to call this function since [HWND.GetWindowText] already
// calls it to allocate the memory block.
//
// [GetWindowTextLength]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextlengthw
func (hWnd HWND) GetWindowTextLength() (uint, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetWindowTextLengthW),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint(ret), nil
}

// [GetWindowThreadProcessId] function.
//
// [GetWindowThreadProcessId]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowthreadprocessid
func (hWnd HWND) GetWindowThreadProcessId() (threadId, processId uint32, wErr error) {
	var pid uint32
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_GetWindowThreadProcessId),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&pid)))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, 0, wErr
	}
	return uint32(ret), pid, nil
}

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
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_InvalidateRect),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)),
		utl.BoolToUintptr(erase))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [IsChild] function.
//
// [IsChild]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-ischild
func (hWnd HWND) IsChild(hChild HWND) bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_IsChild),
		uintptr(hWnd),
		uintptr(hChild))
	return ret != 0
}

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
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_IsDialogMessageW),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

// [IsWindow] function.
//
// [IsWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iswindow
func (hWnd HWND) IsWindow() bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_IsWindow),
		uintptr(hWnd))
	return ret != 0
}

// [MapDialogRect] function.
//
// [MapDialogRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapdialogrect
func (hWnd HWND) MapDialogRect(rc *RECT) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_MapDialogRect),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)))
	return utl.ZeroAsGetLastError(ret, err)
}

// [MessageBox] function.
//
// [MessageBox]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
func (hWnd HWND) MessageBox(text, caption string, uType co.MB) (co.ID, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pText := wbuf.PtrEmptyIsNil(text)
	pCaption := wbuf.PtrEmptyIsNil(caption)

	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_MessageBoxW),
		uintptr(hWnd),
		uintptr(pText),
		uintptr(pCaption),
		uintptr(uType))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return co.ID(0), wErr
	}
	return co.ID(ret), nil
}

// [MonitorFromWindow] function.
//
// [MonitorFromWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfromwindow
func (hWnd HWND) MonitorFromWindow(flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_MonitorFromWindow),
		uintptr(hWnd),
		uintptr(flags))
	return HMONITOR(ret)
}

// [PostMessage] function.
//
// [PostMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postmessagew
func (hWnd HWND) PostMessage(msg co.WM, wParam WPARAM, lParam LPARAM) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_PostMessageW),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return utl.ZeroAsGetLastError(ret, err)
}

// [RedrawWindow] function.
//
// [RedrawWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-redrawwindow
func (hWnd HWND) RedrawWindow(rcUpdate *RECT, hrgnUpdate HRGN, flags co.RDW) error {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_RedrawWindow),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rcUpdate)),
		uintptr(hrgnUpdate),
		uintptr(flags))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [ReleaseDC] function.
//
// [ReleaseDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-releasedc
func (hWnd HWND) ReleaseDC(hdc HDC) error {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_ReleaseDC),
		uintptr(hWnd),
		uintptr(hdc))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [ScreenToClient] function for [POINT].
//
// [ScreenToClient]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientPt(pt *POINT) error {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_ScreenToClient),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pt)))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [ScreenToClient] function for [RECT].
//
// [ScreenToClient]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientRc(rc *RECT) error {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_ScreenToClient),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rc)))
	if ret == 0 {
		return co.ERROR_INVALID_PARAMETER
	}

	ret, _, _ = syscall.SyscallN(dll.User(dll.PROC_ScreenToClient),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&rc.Right)))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [SendMessage] function.
//
// [SendMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew
func (hWnd HWND) SendMessage(msg co.WM, wParam WPARAM, lParam LPARAM) (uintptr, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_SendMessageW),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	if wErr := co.ERROR(err); wErr == co.ERROR_ACCESS_DENIED {
		return 0, wErr
	}
	return ret, nil
}

// [SetFocus] function.
//
// Returns a handle to the previously focused window.
//
// [SetFocus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setfocus
func (hWnd HWND) SetFocus() (HWND, error) {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_SetFocus),
		uintptr(hWnd))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return HWND(0), wErr
	} else {
		return HWND(ret), nil
	}
}

// [SetForegroundWindow] function.
//
// Returns true if the window was brought to the foreground.
//
// [SetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setforegroundwindow
func (hWnd HWND) SetForegroundWindow() bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_SetForegroundWindow),
		uintptr(hWnd))
	return ret != 0
}

// [SetMenu] function.
//
// [SetMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenu
func (hWnd HWND) SetMenu(hMenu HMENU) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_SetMenu),
		uintptr(hWnd),
		uintptr(hMenu))
	return utl.ZeroAsGetLastError(ret, err)
}

// [SetWindowPos] function.
//
// [SetWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
func (hWnd HWND) SetWindowPos(hwndInsertAfter HWND, x, y int, cx, cy uint, flags co.SWP) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_SetWindowPos),
		uintptr(hWnd),
		uintptr(hwndInsertAfter),
		uintptr(int32(x)),
		uintptr(int32(y)),
		uintptr(int32(cx)),
		uintptr(int32(cy)),
		uintptr(flags))
	return utl.ZeroAsGetLastError(ret, err)
}

// [SetWindowRgn] function.
//
// [SetWindowRgn]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowrgn
func (hWnd HWND) SetWindowRgn(hRgn HRGN, redraw bool) error {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_SetWindowRgn),
		uintptr(hWnd),
		uintptr(hRgn),
		utl.BoolToUintptr(redraw))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [SetWindowText] function.
//
// [SetWindowText]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowtextw
func (hWnd HWND) SetWindowText(text string) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pText := wbuf.PtrEmptyIsNil(text)

	syscall.SyscallN(dll.User(dll.PROC_SetWindowTextW),
		uintptr(hWnd),
		uintptr(pText))
}

// [ShowCaret] function.
//
// [ShowCaret]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showcaret
func (hWnd HWND) ShowCaret() error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_ShowCaret),
		uintptr(hWnd))
	return utl.ZeroAsGetLastError(ret, err)
}

// [ShowWindow] function.
//
// [ShowWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
func (hWnd HWND) ShowWindow(cmdShow co.SW) bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_ShowWindow),
		uintptr(hWnd),
		uintptr(cmdShow))
	return ret != 0
}

// Calls [HWND.GetWindowLongPtr] to retrieve the window style.
func (hWnd HWND) Style() (co.WS, error) {
	style, err := hWnd.GetWindowLongPtr(co.GWLP_STYLE)
	return co.WS(style), err
}

// [TranslateAccelerator] function.
//
// [TranslateAccelerator]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translateacceleratorw
func (hWnd HWND) TranslateAccelerator(hAccel HACCEL, msg *MSG) error {
	ret, _, err := syscall.SyscallN(dll.User(dll.PROC_TranslateAcceleratorW),
		uintptr(hWnd),
		uintptr(hAccel),
		uintptr(unsafe.Pointer(msg)))
	return utl.ZeroAsGetLastError(ret, err)
}

// [UpdateWindow] function.
//
// [UpdateWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-updatewindow
func (hWnd HWND) UpdateWindow() bool {
	ret, _, _ := syscall.SyscallN(dll.User(dll.PROC_UpdateWindow),
		uintptr(hWnd))
	return ret != 0
}
