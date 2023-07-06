//go:build windows

package win

import (
	"runtime"
	"sync"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [window].
//
// [window]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hwnd
type HWND HANDLE

// [CreateWindowEx] function.
//
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func CreateWindowEx(
	exStyle co.WS_EX,
	className ClassName,
	title StrOpt,
	style co.WS,
	x, y, width, height int32,
	parent HWND,
	menu HMENU,
	instance HINSTANCE,
	param LPARAM) HWND {

	classNameVal, classNameBuf := className.raw()
	ret, _, err := syscall.SyscallN(proc.CreateWindowEx.Addr(),
		uintptr(exStyle), classNameVal, uintptr(title.Raw()),
		uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(parent), uintptr(menu), uintptr(instance), uintptr(param))
	runtime.KeepAlive(classNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HWND(ret)
}

// [FindWindow] function.
//
// [FindWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-findwindoww
func FindWindow(className ClassName, title StrOpt) (HWND, bool) {
	classNameVal, classNameBuf := className.raw()
	ret, _, _ := syscall.SyscallN(proc.FindWindow.Addr(),
		classNameVal, uintptr(title.Raw()))
	runtime.KeepAlive(classNameBuf)
	return HWND(ret), ret != 0
}

// [GetClipboardOwner] function.
//
// [GetClipboardOwner]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardowner
func GetClipboardOwner() HWND {
	ret, _, err := syscall.SyscallN(proc.GetClipboardOwner.Addr())
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return HWND(ret)
}

// [GetDesktopWindow] function.
//
// [GetDesktopWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdesktopwindow
func GetDesktopWindow() HWND {
	ret, _, _ := syscall.SyscallN(proc.GetDesktopWindow.Addr())
	return HWND(ret)
}

// [GetFocus] function.
//
// [GetFocus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getfocus
func GetFocus() HWND {
	ret, _, _ := syscall.SyscallN(proc.GetFocus.Addr())
	return HWND(ret)
}

// [GetForegroundWindow] function.
//
// [GetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
func GetForegroundWindow() HWND {
	ret, _, _ := syscall.SyscallN(proc.GetForegroundWindow.Addr())
	return HWND(ret)
}

// [GetOpenClipboardWindow] function.
//
// [GetOpenClipboardWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getopenclipboardwindow
func GetOpenClipboardWindow() HWND {
	ret, _, err := syscall.SyscallN(proc.GetOpenClipboardWindow.Addr())
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return HWND(ret)
}

// [GetShellWindow] function.
//
// [GetShellWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getshellwindow
func GetShellWindow() HWND {
	ret, _, _ := syscall.SyscallN(proc.GetShellWindow.Addr())
	return HWND(ret)
}

// [BeginPaint] function.
//
// ⚠️ You must defer HWND.EndPaint().
//
// [BeginPaint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-beginpaint
func (hWnd HWND) BeginPaint(ps *PAINTSTRUCT) HDC {
	ret, _, err := syscall.SyscallN(proc.BeginPaint.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(ps)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HDC(ret)
}

// [ChildWindowFromPoint] function.
//
// [ChildWindowFromPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-childwindowfrompoint
func (hWnd HWND) ChildWindowFromPoint(pt POINT) (HWND, bool) {
	ret, _, _ := syscall.SyscallN(proc.ChildWindowFromPoint.Addr(),
		uintptr(hWnd), uintptr(pt.X), uintptr(pt.Y))
	if ret == 0 {
		return HWND(0), false
	}
	return HWND(ret), true
}

// [ChildWindowFromPointEx] function.
//
// [ChildWindowFromPointEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-childwindowfrompointex
func (hWnd HWND) ChildWindowFromPointEx(pt POINT, flags co.CWP) (HWND, bool) {
	ret, _, _ := syscall.SyscallN(proc.ChildWindowFromPointEx.Addr(),
		uintptr(hWnd), uintptr(pt.X), uintptr(pt.Y), uintptr(flags))
	if ret == 0 {
		return HWND(0), false
	}
	return HWND(ret), true
}

// [ClientToScreenPt] function.
//
// [ClientToScreenPt]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenPt(pt *POINT) {
	ret, _, err := syscall.SyscallN(proc.ClientToScreen.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [ClientToScreenRc] function.
//
// [ClientToScreenRc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenRc(rc *RECT) {
	ret, _, err := syscall.SyscallN(proc.ClientToScreen.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	ret, _, err = syscall.SyscallN(proc.ClientToScreen.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(&rc.Right)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [DefDlgProc] function.
//
// [DefDlgProc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defdlgprocw
func (hWnd HWND) DefDlgProc(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.SyscallN(proc.DefDlgProc.Addr(),
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return ret
}

// [DefWindowProc] function.
//
// [DefWindowProc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defwindowprocw
func (hWnd HWND) DefWindowProc(
	msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {

	ret, _, _ := syscall.SyscallN(proc.DefWindowProc.Addr(),
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return ret
}

// [DestroyWindow] function.
//
// Note: don't call this function to close a window. The correct way to close a
// window is calling SendMessage with WM_CLOSE.
//
// [DestroyWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroywindow
func (hWnd HWND) DestroyWindow() error {
	ret, _, err := syscall.SyscallN(proc.DestroyWindow.Addr(),
		uintptr(hWnd))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [DrawMenuBar] function.
//
// [DrawMenuBar]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawmenubar
func (hWnd HWND) DrawMenuBar() {
	ret, _, err := syscall.SyscallN(proc.DrawMenuBar.Addr(),
		uintptr(hWnd))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [EnableWindow] function.
//
// [EnableWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enablewindow
func (hWnd HWND) EnableWindow(enable bool) bool {
	ret, _, _ := syscall.SyscallN(proc.EnableWindow.Addr(),
		uintptr(hWnd), util.BoolToUintptr(enable))
	return ret != 0 // the window was previously disabled?
}

// [EndDialog] function.
//
// [EndDialog]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddialog
func (hWnd HWND) EndDialog(result uintptr) error {
	ret, _, err := syscall.SyscallN(proc.EndDialog.Addr(),
		uintptr(hWnd), result)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [EndPaint] function.
//
// [EndPaint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endpaint
func (hWnd HWND) EndPaint(ps *PAINTSTRUCT) {
	syscall.SyscallN(proc.EndPaint.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(ps)))
}

// [EnumChildWindows] function.
//
// To continue enumeration, the callback function must return true; to stop
// enumeration, it must return false.
//
// [EnumChildWindows]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumchildwindows
func (hWnd HWND) EnumChildWindows(callback func(hChild HWND) bool) {
	pPack := &_EnumChildPack{f: callback}
	_globalEnumChildMutex.Lock()
	if _globalEnumChildFuncs == nil { // the set was not initialized yet?
		_globalEnumChildFuncs = make(map[*_EnumChildPack]struct{}, 1)
	}
	_globalEnumChildFuncs[pPack] = struct{}{} // store pointer in the set
	_globalEnumChildMutex.Unlock()

	syscall.SyscallN(proc.EnumChildWindows.Addr(),
		uintptr(hWnd), _globalEnumChildCallback,
		uintptr(unsafe.Pointer(pPack)))

	_globalEnumChildMutex.Lock()
	delete(_globalEnumChildFuncs, pPack) // remove from the set
	_globalEnumChildMutex.Unlock()
}

type _EnumChildPack struct{ f func(hChild HWND) bool }

var (
	_globalEnumChildFuncs    map[*_EnumChildPack]struct{} // keeps pointers from being collected by GC
	_globalEnumChildMutex    = sync.Mutex{}
	_globalEnumChildCallback = syscall.NewCallback(
		func(hChild HWND, lParam LPARAM) uintptr {
			pPack := (*_EnumChildPack)(unsafe.Pointer(lParam))
			return util.BoolToUintptr(pPack.f(hChild))
		})
)

// [GetAncestor] function.
//
// [GetAncestor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getancestor
func (hWnd HWND) GetAncestor(gaFlags co.GA) HWND {
	ret, _, _ := syscall.SyscallN(proc.GetAncestor.Addr(),
		uintptr(hWnd), uintptr(gaFlags))
	return HWND(ret)
}

// [GetClassLongPtr] function.
//
// [GetClassLongPtr]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclasslongptrw
func (hWnd HWND) GetClassLongPtr(index co.GCL) uint32 {
	ret, _, err := syscall.SyscallN(proc.GetClassLongPtr.Addr(),
		uintptr(hWnd), uintptr(index))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// [GetClassName] function.
//
// [GetClassName]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassnamew
func (hWnd HWND) GetClassName() string {
	var buf [256 + 1]uint16
	ret, _, err := syscall.SyscallN(proc.GetClassName.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return Str.FromNativeSlice(buf[:])
}

// [GetClientRect] function.
//
// [GetClientRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclientrect
func (hWnd HWND) GetClientRect() RECT {
	var rc RECT
	ret, _, err := syscall.SyscallN(proc.GetClientRect.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return rc
}

// [GetDC] function.
//
// Call HWND(0).GetDC() to retrieve the DC for the entire screen.
//
// ⚠️ You must defer HDC.ReleaseDC().
//
// [GetDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
func (hWnd HWND) GetDC() HDC {
	ret, _, err := syscall.SyscallN(proc.GetDC.Addr(),
		uintptr(hWnd))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HDC(ret)
}

// [GetDlgCtrlID] function.
//
// [GetDlgCtrlID]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgctrlid
func (hWnd HWND) GetDlgCtrlID() int32 {
	syscall.SyscallN(proc.SetLastError.Addr(), 0)

	ret, _, err := syscall.SyscallN(proc.GetDlgCtrlID.Addr(),
		uintptr(hWnd))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return int32(ret)
}

// [GetDlgItem] function.
//
// [GetDlgItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgitem
func (hWnd HWND) GetDlgItem(dlgId int32) HWND {
	ret, _, err := syscall.SyscallN(proc.GetDlgItem.Addr(),
		uintptr(hWnd), uintptr(dlgId))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HWND(ret)
}

// [GetLastActivePopup] function.
//
// [GetLastActivePopup]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getlastactivepopup
func (hWnd HWND) GetLastActivePopup() HWND {
	ret, _, _ := syscall.SyscallN(proc.GetLastActivePopup.Addr(),
		uintptr(hWnd))
	return HWND(ret)
}

// [GetMenu] function.
//
// [GetMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenu
func (hWnd HWND) GetMenu() HMENU {
	ret, _, _ := syscall.SyscallN(proc.GetMenu.Addr(),
		uintptr(hWnd))
	return HMENU(ret)
}

// [GetNextDlgGroupItem] function.
//
// [GetNextDlgGroupItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlggroupitem
func (hWnd HWND) GetNextDlgGroupItem(hChild HWND, isPrevious bool) HWND {
	ret, _, err := syscall.SyscallN(proc.GetNextDlgGroupItem.Addr(),
		uintptr(hWnd), uintptr(hChild), util.BoolToUintptr(isPrevious))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return HWND(ret)
}

// [GetNextDlgTabItem] function.
//
// [GetNextDlgTabItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlgtabitem
func (hWnd HWND) GetNextDlgTabItem(hChild HWND, isPrevious bool) HWND {
	ret, _, err := syscall.SyscallN(proc.GetNextDlgTabItem.Addr(),
		uintptr(hWnd), uintptr(hChild), util.BoolToUintptr(isPrevious))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return HWND(ret)
}

// [GetParent] function.
//
// [GetParent]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getparent
func (hWnd HWND) GetParent() HWND {
	ret, _, err := syscall.SyscallN(proc.GetParent.Addr(),
		uintptr(hWnd))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return HWND(ret)
}

// [GetScrollInfo] function.
//
// [GetScrollInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getscrollinfo
func (hWnd HWND) GetScrollInfo(bar co.SB_TYPE, si *SCROLLINFO) {
	ret, _, err := syscall.SyscallN(proc.GetScrollInfo.Addr(),
		uintptr(hWnd), uintptr(bar), uintptr(unsafe.Pointer(si)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [GetSystemMenu] function.
//
// [GetSystemMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmenu
func (hWnd HWND) GetSystemMenu(revert bool) HMENU {
	ret, _, _ := syscall.SyscallN(proc.GetSystemMenu.Addr(),
		uintptr(hWnd), util.BoolToUintptr(revert))
	return HMENU(ret)
}

// [GetTopWindow] function.
//
// [GetTopWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-gettopwindow
func (hWnd HWND) GetTopWindow() HWND {
	ret, _, err := syscall.SyscallN(proc.GetTopWindow.Addr(),
		uintptr(hWnd))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return HWND(ret)
}

// [GetWindow] function.
//
// [GetWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindow
func (hWnd HWND) GetWindow(cmd co.GW) HWND {
	ret, _, err := syscall.SyscallN(proc.GetWindow.Addr(),
		uintptr(hWnd), uintptr(cmd))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return HWND(ret)
}

// [GetWindowDC] function.
//
// ⚠️ You must defer HDC.ReleaseDC().
//
// [GetWindowDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowdc
func (hWnd HWND) GetWindowDC() HDC {
	ret, _, err := syscall.SyscallN(proc.GetWindowDC.Addr(),
		uintptr(hWnd))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HDC(ret)
}

// [GetWindowLongPtr] function.
//
// [GetWindowLongPtr]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
func (hWnd HWND) GetWindowLongPtr(index co.GWLP) uintptr {
	ret, _, err := syscall.SyscallN(proc.GetWindowLongPtr.Addr(),
		uintptr(hWnd), uintptr(index))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return ret
}

// [GetWindowRect] function.
//
// [GetWindowRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowrect
func (hWnd HWND) GetWindowRect() RECT {
	var rc RECT
	ret, _, err := syscall.SyscallN(proc.GetWindowRect.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return rc
}

// [GetWindowText] function.
//
// Calls [GetWindowTextLength] to allocate the memory block.
//
// [GetWindowText]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextw
// [GetWindowTextLength]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextlengthw
func (hWnd HWND) GetWindowText() string {
	len := hWnd.GetWindowTextLength() + 1
	buf := make([]uint16, len)

	ret, _, err := syscall.SyscallN(proc.GetWindowText.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(len))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return Str.FromNativeSlice(buf)
}

// [GetWindowTextLength] function.
//
// You usually don't need to call this function since [GetWindowText] already
// calls it to allocate the memory block.
//
// [GetWindowTextLength]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextlengthw
// [GetWindowText]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextw
func (hWnd HWND) GetWindowTextLength() int32 {
	ret, _, _ := syscall.SyscallN(proc.GetWindowTextLength.Addr(),
		uintptr(hWnd))
	return int32(ret)
}

// [GetWindowThreadProcessId] function.
//
// [GetWindowThreadProcessId]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowthreadprocessid
func (hWnd HWND) GetWindowThreadProcessId() (threadId, processId uint32) {
	ret, _, _ := syscall.SyscallN(proc.GetWindowThreadProcessId.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(&processId)))
	return uint32(ret), processId
}

// [HideCaret] function.
//
// [HideCaret]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-hidecaret
func (hWnd HWND) HideCaret() {
	ret, _, err := syscall.SyscallN(proc.HideCaret.Addr(),
		uintptr(hWnd))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [HiliteMenuItem] function.
//
// [HiliteMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-hilitemenuitem
func (hWnd HWND) HiliteMenuItem(hMenu HMENU, item MenuItem, hilite bool) bool {
	idPos, mf := item.raw()
	flags := util.Iif(hilite, co.MFS_HILITE, co.MFS_UNHILITE).(co.MF) | mf

	ret, _, _ := syscall.SyscallN(proc.HiliteMenuItem.Addr(),
		uintptr(hWnd), uintptr(hMenu), idPos, uintptr((flags)))
	return ret != 0
}

// [InvalidateRect] function.
//
// [InvalidateRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invalidaterect
func (hWnd HWND) InvalidateRect(rc *RECT, erase bool) {
	ret, _, err := syscall.SyscallN(proc.InvalidateRect.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)),
		util.BoolToUintptr(erase))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [IsChild] function.
//
// [IsChild]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-ischild
func (hWnd HWND) IsChild(hChild HWND) bool {
	ret, _, _ := syscall.SyscallN(proc.IsChild.Addr(),
		uintptr(hWnd), uintptr(hChild))
	return ret != 0
}

// [IsDlgButtonChecked] function.
//
// [IsDlgButtonChecked]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdlgbuttonchecked
func (hWnd HWND) IsDlgButtonChecked(idButton int32) co.BST {
	ret, _, _ := syscall.SyscallN(proc.IsDlgButtonChecked.Addr(),
		uintptr(hWnd), uintptr(idButton))
	return co.BST(ret)
}

// [IsDialogMessage] function.
//
// [IsDialogMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdialogmessagew
func (hWnd HWND) IsDialogMessage(msg *MSG) bool {
	ret, _, _ := syscall.SyscallN(proc.IsDialogMessage.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

// [IsIconic] function.
//
// [IsIconic]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isiconic
func (hWnd HWND) IsIconic() bool {
	ret, _, _ := syscall.SyscallN(proc.IsIconic.Addr(),
		uintptr(hWnd))
	return ret != 0
}

// [IsWindow] function.
//
// [IsWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iswindow
func (hWnd HWND) IsWindow() bool {
	ret, _, _ := syscall.SyscallN(proc.IsWindow.Addr(),
		uintptr(hWnd))
	return ret != 0
}

// [IsWindowEnabled] function.
//
// [IsWindowEnabled]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iswindowenabled
func (hWnd HWND) IsWindowEnabled() bool {
	ret, _, _ := syscall.SyscallN(proc.IsWindowEnabled.Addr(),
		uintptr(hWnd))
	return ret != 0
}

// [IsWindowVisible] function.
//
// [IsWindowVisible]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iswindowvisible
func (hWnd HWND) IsWindowVisible() bool {
	ret, _, _ := syscall.SyscallN(proc.IsWindowVisible.Addr(),
		uintptr(hWnd))
	return ret != 0
}

// [IsZoomed] function.
//
// [IsZoomed]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iszoomed
func (hWnd HWND) IsZoomed() bool {
	ret, _, _ := syscall.SyscallN(proc.IsZoomed.Addr(),
		uintptr(hWnd))
	return ret != 0
}

// [KillTimer] function.
//
// [KillTimer]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-killtimer
func (hWnd HWND) KillTimer(timerId uintptr) {
	ret, _, err := syscall.SyscallN(proc.KillTimer.Addr(),
		uintptr(hWnd), timerId)

	if timerId > 0xffff { // guess: Win32 pointers are greater than WORDs
		_globalTimerMutex.Lock()
		delete(_globalTimerFuncs, (*_TimerPack)(unsafe.Pointer(timerId))) // remove from set
		_globalTimerMutex.Unlock()
		// At this moment, the callback pointer has no more references. If
		// KillTimer() is called from within the callback itself, it's unsure
		// whether the running function will be enough to keep its pointer from
		// being collected by the GC, but it's reasonable to think so.
	}

	if ret == 0 && errco.ERROR(err) != errco.SUCCESS {
		panic(errco.ERROR(err))
	}
}

// [LockWindowUpdate] function.
//
// [LockWindowUpdate]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-lockwindowupdate
func (hWnd HWND) LockWindowUpdate() error {
	ret, _, err := syscall.SyscallN(proc.LockWindowUpdate.Addr(),
		uintptr(hWnd))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [LogicalToPhysicalPoint] function.
//
// [LogicalToPhysicalPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-logicaltophysicalpoint
func (hWnd HWND) LogicalToPhysicalPoint(pt *POINT) {
	ret, _, err := syscall.SyscallN(proc.LogicalToPhysicalPoint.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [MapDialogRect] function.
//
// [MapDialogRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapdialogrect
func (hWnd HWND) MapDialogRect(rc *RECT) {
	ret, _, err := syscall.SyscallN(proc.MapDialogRect.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [MapWindowPoints] function.
//
// Returns the number of pixels added horizontally and vertically to the passed
// points.
//
// [MapWindowPoints]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapwindowpoints
func (hWnd HWND) MapWindowPoints(hWndTo HWND, points []POINT) (int, int) {
	syscall.SyscallN(proc.SetLastError.Addr(), 0)

	ret, _, _ := syscall.SyscallN(proc.MapWindowPoints.Addr(),
		uintptr(hWnd), uintptr(hWndTo),
		uintptr(unsafe.Pointer(&points[0])), uintptr(len(points)))
	return int(LOWORD(uint32(ret))), int(HIWORD(uint32(ret)))
}

// [MenuItemFromPoint] function.
//
// [MenuItemFromPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-menuitemfrompoint
func (hWnd HWND) MenuItemFromPoint(hMenu HMENU, pt POINT) (int, bool) {
	ret, _, _ := syscall.SyscallN(proc.MenuItemFromPoint.Addr(),
		uintptr(hWnd), uintptr(hMenu), uintptr(pt.X), uintptr(pt.Y))
	return int(ret), int(ret) != -1
}

// [MessageBox] function.
//
// [MessageBox]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
func (hWnd HWND) MessageBox(text, caption string, uType co.MB) co.ID {
	ret, _, err := syscall.SyscallN(proc.MessageBox.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(Str.ToNativePtr(text))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(caption))), uintptr(uType))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.ID(ret)
}

// [MonitorFromWindow] function.
//
// [MonitorFromWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfromwindow
func (hWnd HWND) MonitorFromWindow(flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.SyscallN(proc.MonitorFromWindow.Addr(),
		uintptr(hWnd), uintptr(flags))
	return HMONITOR(ret)
}

// [MoveWindow] function.
//
// [MoveWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-movewindow
func (hWnd HWND) MoveWindow(x, y, width, height int32, repaint bool) {
	ret, _, err := syscall.SyscallN(proc.MoveWindow.Addr(),
		uintptr(hWnd), uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		util.BoolToUintptr(repaint))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [OpenClipboard] function.
//
// ⚠️ You must defer HCLIPBOARD.CloseClipboard().
//
// [OpenClipboard]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-openclipboard
func (hWnd HWND) OpenClipboard() HCLIPBOARD {
	ret, _, err := syscall.SyscallN(proc.OpenClipboard.Addr(),
		uintptr(hWnd))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HCLIPBOARD{}
}

// [PhysicalToLogicalPoint] function.
//
// [PhysicalToLogicalPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-physicaltologicalpoint
func (hWnd HWND) PhysicalToLogicalPoint(pt *POINT) {
	ret, _, err := syscall.SyscallN(proc.PhysicalToLogicalPoint.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [PostMessage] function.
//
// [PostMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postmessagew
func (hWnd HWND) PostMessage(msg co.WM, wParam WPARAM, lParam LPARAM) {
	ret, _, err := syscall.SyscallN(proc.PostMessage.Addr(),
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [RealChildWindowFromPoint] function.
//
// [RealChildWindowFromPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-realchildwindowfrompoint
func (hWnd HWND) RealChildWindowFromPoint(
	parentClientCoords POINT) (HWND, bool) {

	ret, _, _ := syscall.SyscallN(proc.RealChildWindowFromPoint.Addr(),
		uintptr(hWnd),
		uintptr(parentClientCoords.X), uintptr(parentClientCoords.Y))
	return HWND(ret), ret != 0
}

// [RealGetWindowClass] function.
//
// [RealGetWindowClass]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-realgetwindowclassw
func (hWnd HWND) RealGetWindowClass() string {
	var buf [256 + 1]uint16
	ret, _, err := syscall.SyscallN(proc.RealGetWindowClass.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return Str.FromNativeSlice(buf[:])
}

// [ReleaseDC] function.
//
// [ReleaseDC]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-releasedc
func (hWnd HWND) ReleaseDC(hdc HDC) {
	ret, _, err := syscall.SyscallN(proc.ReleaseDC.Addr(),
		uintptr(hWnd), uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [ScreenToClientPt] function.
//
// [ScreenToClientPt]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientPt(pt *POINT) {
	ret, _, err := syscall.SyscallN(proc.ScreenToClient.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [ScreenToClientRc] function.
//
// [ScreenToClientRc]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientRc(rc *RECT) {
	ret, _, err := syscall.SyscallN(proc.ScreenToClient.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	ret, _, err = syscall.SyscallN(proc.ScreenToClient.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(&rc.Right)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SendMessage] function.
//
// [SendMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew
func (hWnd HWND) SendMessage(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.SyscallN(proc.SendMessage.Addr(),
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return ret
}

// [SendMessageTimeout] function.
//
// [SendMessageTimeout]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagetimeoutw
func (hWnd HWND) SendMessageTimeout(
	msg co.WM, wParam WPARAM, lParam LPARAM,
	flags co.SMTO, msTimeout int) (uintptr, error) {

	var procRet uintptr
	ret, _, err := syscall.SyscallN(proc.SendMessageTimeout.Addr(),
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		uintptr(flags), uintptr(msTimeout), uintptr(unsafe.Pointer(&procRet)))

	if ret == 0 {
		return procRet, errco.ERROR(err)
	} else {
		return procRet, nil
	}
}

// [SetFocus] function.
//
// Returns a handle to the previously focused window.
//
// [SetFocus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setfocus
func (hWnd HWND) SetFocus() (HWND, error) {
	ret, _, err := syscall.SyscallN(proc.SetFocus.Addr(),
		uintptr(hWnd))
	if hPrev, err := HWND(ret), errco.ERROR(err); hPrev == 0 && err != errco.S_OK {
		return hPrev, err
	} else {
		return hPrev, nil
	}
}

// [SetForegroundWindow] function.
//
// Returns true if the window was brought to the foreground.
//
// [SetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setforegroundwindow
func (hWnd HWND) SetForegroundWindow() bool {
	ret, _, _ := syscall.SyscallN(proc.SetForegroundWindow.Addr(),
		uintptr(hWnd))
	return ret != 0
}

// [SetLayeredWindowAttributes] function.
//
// [SetLayeredWindowAttributes]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setlayeredwindowattributes
func (hWnd HWND) SetLayeredWindowAttributes(
	transparencyColorKey COLORREF, alpha uint8, flags co.LWA) {

	ret, _, err := syscall.SyscallN(proc.SetLayeredWindowAttributes.Addr(),
		uintptr(hWnd), uintptr(transparencyColorKey),
		uintptr(alpha), uintptr(flags))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SetMenu] function.
//
// [SetMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenu
func (hWnd HWND) SetMenu(hMenu HMENU) {
	ret, _, err := syscall.SyscallN(proc.SetMenu.Addr(),
		uintptr(hWnd), uintptr(hMenu))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SetParent] function.
//
// Returns the handle of the previous parent window.
//
// [SetParent]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setparent
func (hWnd HWND) SetParent(hwndNewParent HWND) HWND {
	ret, _, err := syscall.SyscallN(proc.SetParent.Addr(),
		uintptr(hWnd), uintptr(hwndNewParent))
	if hPrev, wErr := HWND(ret), errco.ERROR(err); hPrev == 0 && wErr != errco.S_OK {
		panic(wErr)
	} else {
		return hPrev
	}
}

// [SetScrollInfo] function.
//
// Returns the current position of the scroll box.
//
// [SetScrollInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setscrollinfo
func (hWnd HWND) SetScrollInfo(
	bar co.SB_TYPE, si *SCROLLINFO, redraw bool) int32 {

	ret, _, _ := syscall.SyscallN(proc.SetScrollInfo.Addr(),
		uintptr(hWnd), uintptr(bar), uintptr(unsafe.Pointer(si)),
		util.BoolToUintptr(redraw))
	return int32(ret)
}

// [SetScrollPos] function.
//
// [SetScrollPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setscrollpos
func (hWnd HWND) SetScrollPos(bar co.SB_TYPE, pos int32, redraw bool) int32 {
	ret, _, _ := syscall.SyscallN(proc.SetScrollPos.Addr(),
		uintptr(hWnd), uintptr(bar), uintptr(pos), util.BoolToUintptr(redraw))
	return int32(ret)
}

// [SetScrollRange] function.
//
// [SetScrollRange]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setscrollrange
func (hWnd HWND) SetScrollRange(
	bar co.SB_TYPE, minPos, maxPos int32, redraw bool) bool {

	ret, _, _ := syscall.SyscallN(proc.SetScrollRange.Addr(),
		uintptr(hWnd), uintptr(bar), uintptr(minPos), uintptr(maxPos),
		util.BoolToUintptr(redraw))
	return ret != 0
}

// [SetTimer] function.
//
// This method will create a timer that will post WM_TIMER messages, instead of
// running a callback.
//
// The method returns the timer ID.
//
// ⚠️ You must call HWND.KillTimer() to stop the timer.
//
// [SetTimer]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-settimer
func (hWnd HWND) SetTimer(msElapse int, timerId uintptr) uintptr {
	ret, _, err := syscall.SyscallN(proc.SetTimer.Addr(),
		uintptr(hWnd), timerId, uintptr(msElapse), 0)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return ret
}

// Creates a timer with [SetTimer], which runs the given callback instead of
// posting WM_TIMER messages.
//
// The method returns the timer ID, which is also sent to the callback.
//
// ⚠️ You must call HWND.KillTimer() to stop the timer and free the allocated
// resources.
//
// # Example
//
//	var hWnd HWND // initialized somewhere
//
//	hWnd.SetTimerCallback(2000, func(timerId uintptr) {
//		hWnd.KillTimer(timerId)
//		println("This callback will run once.")
//	})
//
// [SetTimer]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-settimer
func (hWnd HWND) SetTimerCallback(
	msElapse int, timerFunc func(timerId uintptr)) uintptr {

	pPack := &_TimerPack{f: timerFunc}
	_globalTimerMutex.Lock()
	if _globalTimerFuncs == nil { // the set was not initialized yet?
		_globalTimerFuncs = make(map[*_TimerPack]struct{}, 1)
	}
	_globalTimerFuncs[pPack] = struct{}{} // store pointer in the set
	_globalTimerMutex.Unlock()

	timerId := uintptr(unsafe.Pointer(pPack)) // use the pack pointer as the timer ID

	ret, _, err := syscall.SyscallN(proc.SetTimer.Addr(),
		uintptr(hWnd), timerId, uintptr(msElapse), _globalTimerCallback)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return ret
}

type _TimerPack struct{ f func(timerId uintptr) }

var (
	_globalTimerFuncs    map[*_TimerPack]struct{} // keeps pointers from being collected by GC
	_globalTimerMutex    = sync.Mutex{}
	_globalTimerCallback = syscall.NewCallback(
		func(_ HWND, _ co.WM, timerId uintptr, _ uint32) uintptr {
			pPack := (*_TimerPack)(unsafe.Pointer(timerId))

			_globalTimerMutex.Lock()
			_, isStored := _globalTimerFuncs[pPack]
			_globalTimerMutex.Unlock()

			if isStored {
				pPack.f(timerId) // invoke user callback
			}
			return 0
		})
)

// [SetWindowDisplayAffinity] function.
//
// [SetWindowDisplayAffinity]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowdisplayaffinity
func (hWnd HWND) SetWindowDisplayAffinity(affinity co.WDA) {
	ret, _, err := syscall.SyscallN(proc.SetWindowDisplayAffinity.Addr(),
		uintptr(hWnd), uintptr(affinity))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SetWindowLongPtr] function.
//
// [SetWindowLongPtr]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowlongptrw
func (hWnd HWND) SetWindowLongPtr(index co.GWLP, newLong uintptr) uintptr {
	syscall.SyscallN(proc.SetLastError.Addr(), 0)

	ret, _, err := syscall.SyscallN(proc.SetWindowLongPtr.Addr(),
		uintptr(hWnd), uintptr(index), newLong)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return ret
}

// [SetWindowPos] function.
//
// You can pass HWND or HWND_IA in hwndInsertAfter argument.
//
// [SetWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
func (hWnd HWND) SetWindowPos(
	hwndInsertAfter HWND, x, y, cx, cy int32, flags co.SWP) {

	ret, _, err := syscall.SyscallN(proc.SetWindowPos.Addr(),
		uintptr(hWnd), uintptr(hwndInsertAfter),
		uintptr(x), uintptr(y), uintptr(cx), uintptr(cy),
		uintptr(flags))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SetWindowRgn] function.
//
// [SetWindowRgn]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowrgn
func (hWnd HWND) SetWindowRgn(hRgn HRGN, redraw bool) {
	ret, _, err := syscall.SyscallN(proc.SetWindowRgn.Addr(),
		uintptr(hWnd), uintptr(hRgn), util.BoolToUintptr(redraw))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SetWindowText] function.
//
// [SetWindowText]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowtextw
func (hWnd HWND) SetWindowText(text string) {
	syscall.SyscallN(proc.SetWindowText.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(Str.ToNativePtr(text))))
}

// [ShowCaret] function.
//
// [ShowCaret]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showcaret
func (hWnd HWND) ShowCaret() {
	ret, _, err := syscall.SyscallN(proc.ShowCaret.Addr(),
		uintptr(hWnd))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [ShowWindow] function.
//
// [ShowWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
func (hWnd HWND) ShowWindow(cmdShow co.SW) bool {
	ret, _, _ := syscall.SyscallN(proc.ShowWindow.Addr(),
		uintptr(hWnd), uintptr(cmdShow))
	return ret != 0
}

// [TranslateAccelerator] function.
//
// [TranslateAccelerator]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translateacceleratorw
func (hWnd HWND) TranslateAccelerator(hAccel HACCEL, msg *MSG) error {
	ret, _, err := syscall.SyscallN(proc.TranslateAccelerator.Addr(),
		uintptr(hWnd), uintptr(hAccel), uintptr(unsafe.Pointer(msg)))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [UpdateWindow] function.
//
// [UpdateWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-updatewindow
func (hWnd HWND) UpdateWindow() bool {
	ret, _, _ := syscall.SyscallN(proc.UpdateWindow.Addr(),
		uintptr(hWnd))
	return ret != 0
}
