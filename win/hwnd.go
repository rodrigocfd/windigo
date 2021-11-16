package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a window.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hwnd
type HWND HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func CreateWindowEx(exStyle co.WS_EX, className ClassName, title StrOrNil,
	style co.WS, x, y, width, height int32,
	parent HWND, menu HMENU, instance HINSTANCE, param LPARAM) HWND {

	classNameVal, classNameBuf := variantClassName(className)
	ret, _, err := syscall.Syscall12(proc.CreateWindowEx.Addr(), 12,
		uintptr(exStyle), classNameVal, uintptr(variantStrOrNil(title)),
		uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(parent), uintptr(menu), uintptr(instance), uintptr(param))
	runtime.KeepAlive(classNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-findwindoww
func FindWindow(className ClassName, title StrOrNil) (HWND, bool) {
	classNameVal, classNameBuf := variantClassName(className)
	ret, _, _ := syscall.Syscall(proc.FindWindow.Addr(), 2,
		classNameVal, uintptr(variantStrOrNil(title)), 0)
	runtime.KeepAlive(classNameBuf)
	return HWND(ret), ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdesktopwindow
func GetDesktopWindow() HWND {
	ret, _, _ := syscall.Syscall(proc.GetDesktopWindow.Addr(), 0,
		0, 0, 0)
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getfocus
func GetFocus() HWND {
	ret, _, _ := syscall.Syscall(proc.GetFocus.Addr(), 0,
		0, 0, 0)
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
func GetForegroundWindow() HWND {
	ret, _, _ := syscall.Syscall(proc.GetForegroundWindow.Addr(), 0,
		0, 0, 0)
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getshellwindow
func GetShellWindow() HWND {
	ret, _, _ := syscall.Syscall(proc.GetShellWindow.Addr(), 0,
		0, 0, 0)
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-gettopwindow
func GetTopWindow(hwnd HWND) HWND {
	ret, _, _ := syscall.Syscall(proc.GetTopWindow.Addr(), 1,
		uintptr(hwnd), 0, 0)
	return HWND(ret)
}

// ⚠️ You must defer HWND.EndPaint().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-beginpaint
func (hWnd HWND) BeginPaint(ps *PAINTSTRUCT) HDC {
	ret, _, err := syscall.Syscall(proc.BeginPaint.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(ps)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HDC(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-childwindowfrompoint
func (hWnd HWND) ChildWindowFromPoint(pt POINT) (HWND, bool) {
	ret, _, _ := syscall.Syscall(proc.ChildWindowFromPoint.Addr(), 3,
		uintptr(hWnd), uintptr(pt.X), uintptr(pt.Y))
	if ret == 0 {
		return HWND(0), false
	}
	return HWND(ret), true
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-childwindowfrompointex
func (hWnd HWND) ChildWindowFromPointEx(pt POINT, flags co.CWP) (HWND, bool) {
	ret, _, _ := syscall.Syscall6(proc.ChildWindowFromPointEx.Addr(), 4,
		uintptr(hWnd), uintptr(pt.X), uintptr(pt.Y), uintptr(flags), 0, 0)
	if ret == 0 {
		return HWND(0), false
	}
	return HWND(ret), true
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenPt(pt *POINT) {
	ret, _, err := syscall.Syscall(proc.ClientToScreen.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(pt)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clienttoscreen
func (hWnd HWND) ClientToScreenRc(rc *RECT) {
	ret, _, err := syscall.Syscall(proc.ClientToScreen.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	ret, _, err = syscall.Syscall(proc.ClientToScreen.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(&rc.Right)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defdlgprocw
func (hWnd HWND) DefDlgProc(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.Syscall6(proc.DefDlgProc.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
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
	ret, _, err := syscall.Syscall(proc.DestroyWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragacceptfiles
func (hWnd HWND) DragAcceptFiles(accept bool) {
	syscall.Syscall(proc.DragAcceptFiles.Addr(), 2,
		uintptr(hWnd), util.BoolToUintptr(accept), 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawmenubar
func (hWnd HWND) DrawMenuBar() {
	ret, _, err := syscall.Syscall(proc.DrawMenuBar.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea
func (hWnd HWND) DwmExtendFrameIntoClientArea(marginsInset *MARGINS) {
	ret, _, _ := syscall.Syscall(proc.DwmExtendFrameIntoClientArea.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(marginsInset)), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmseticoniclivepreviewbitmap
func (hWnd HWND) DwmSetIconicLivePreviewBitmap(
	hBmp HBITMAP, ptClient POINT, sitFlags co.DWM_SIT) {

	ret, _, _ := syscall.Syscall6(proc.DwmSetIconicLivePreviewBitmap.Addr(), 4,
		uintptr(hWnd), uintptr(hBmp), uintptr(unsafe.Pointer(&ptClient)),
		uintptr(sitFlags), 0, 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmseticonicthumbnail
func (hWnd HWND) DwmSetIconicThumbnail(hBmp HBITMAP, sitFlags co.DWM_SIT) {
	ret, _, _ := syscall.Syscall(proc.DwmSetIconicThumbnail.Addr(), 3,
		uintptr(hWnd), uintptr(hBmp), uintptr(sitFlags))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enablewindow
func (hWnd HWND) EnableWindow(enable bool) bool {
	ret, _, _ := syscall.Syscall(proc.EnableWindow.Addr(), 2,
		uintptr(hWnd), util.BoolToUintptr(enable), 0)
	return ret != 0 // the window was previously disabled?
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddialog
func (hWnd HWND) EndDialog(result uintptr) {
	ret, _, err := syscall.Syscall(proc.EndDialog.Addr(), 2,
		uintptr(hWnd), result, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endpaint
func (hWnd HWND) EndPaint(ps *PAINTSTRUCT) {
	syscall.Syscall(proc.EndPaint.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(ps)), 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumchildwindows
func (hWnd HWND) EnumChildWindows(enumFunc func(hChild HWND) bool) {
	pPack := &_EnumChildPack{f: enumFunc}
	if _globalEnumChildFuncs == nil {
		_globalEnumChildFuncs = make(map[*_EnumChildPack]struct{}, 2)
	}
	_globalEnumChildFuncs[pPack] = struct{}{} // store pointer in the set

	syscall.Syscall(proc.EnumChildWindows.Addr(), 3,
		uintptr(hWnd), _globalEnumChildCallback,
		uintptr(unsafe.Pointer(pPack)))
}

type _EnumChildPack struct{ f func(hChild HWND) bool }

var (
	_globalEnumChildCallback uintptr = syscall.NewCallback(_EnumChildProc)
	_globalEnumChildFuncs    map[*_EnumChildPack]struct{}
)

func _EnumChildProc(hChild HWND, lParam LPARAM) uintptr {
	pPack := (*_EnumChildPack)(unsafe.Pointer(lParam))
	retVal := uintptr(0)
	if _, isStored := _globalEnumChildFuncs[pPack]; isStored {
		retVal = util.BoolToUintptr(pPack.f(hChild))
		if retVal == 0 {
			delete(_globalEnumChildFuncs, pPack) // remove from set
		}
	}
	return retVal
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getancestor
func (hWnd HWND) GetAncestor(gaFlags co.GA) HWND {
	ret, _, _ := syscall.Syscall(proc.GetAncestor.Addr(), 2,
		uintptr(hWnd), uintptr(gaFlags), 0)
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclasslongptrw
func (hWnd HWND) GetClassLongPtr(index co.GCL) uint32 {
	ret, _, err := syscall.Syscall(proc.GetClassLongPtr.Addr(), 2,
		uintptr(hWnd), uintptr(index), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassnamew
func (hWnd HWND) GetClassName() string {
	var buf [256 + 1]uint16
	ret, _, err := syscall.Syscall(proc.GetClassName.Addr(), 3,
		uintptr(hWnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return Str.FromNativeSlice(buf[:])
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclientrect
func (hWnd HWND) GetClientRect() RECT {
	var rc RECT
	ret, _, err := syscall.Syscall(proc.GetClientRect.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(&rc)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return rc
}

// ⚠️ You must defer HDC.ReleaseDC().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
func (hWnd HWND) GetDC() HDC {
	ret, _, err := syscall.Syscall(proc.GetDC.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HDC(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgctrlid
func (hWnd HWND) GetDlgCtrlID() int32 {
	ret, _, err := syscall.Syscall(proc.GetDlgCtrlID.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgitem
func (hWnd HWND) GetDlgItem(dlgId int32) HWND {
	ret, _, err := syscall.Syscall(proc.GetDlgItem.Addr(), 2,
		uintptr(hWnd), uintptr(dlgId), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenu
func (hWnd HWND) GetMenu() HMENU {
	ret, _, _ := syscall.Syscall(proc.GetMenu.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return HMENU(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlggroupitem
func (hWnd HWND) GetNextDlgGroupItem(hChild HWND, isPrevious bool) HWND {
	ret, _, err := syscall.Syscall(proc.GetNextDlgGroupItem.Addr(), 3,
		uintptr(hWnd), uintptr(hChild), util.BoolToUintptr(isPrevious))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getnextdlgtabitem
func (hWnd HWND) GetNextDlgTabItem(hChild HWND, isPrevious bool) HWND {
	ret, _, err := syscall.Syscall(proc.GetNextDlgTabItem.Addr(), 3,
		uintptr(hWnd), uintptr(hChild), util.BoolToUintptr(isPrevious))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getparent
func (hWnd HWND) GetParent() HWND {
	ret, _, err := syscall.Syscall(proc.GetParent.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getscrollinfo
func (hWnd HWND) GetScrollInfo(bar co.SB_TYPE, si *SCROLLINFO) {
	ret, _, err := syscall.Syscall(proc.GetScrollInfo.Addr(), 3,
		uintptr(hWnd), uintptr(bar), uintptr(unsafe.Pointer(si)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmenu
func (hWnd HWND) GetSystemMenu(revert bool) HMENU {
	ret, _, _ := syscall.Syscall(proc.GetSystemMenu.Addr(), 2,
		uintptr(hWnd), util.BoolToUintptr(revert), 0)
	return HMENU(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindow
func (hWnd HWND) GetWindow(cmd co.GW) HWND {
	ret, _, err := syscall.Syscall(proc.GetWindow.Addr(), 2,
		uintptr(hWnd), uintptr(cmd), 0)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return HWND(ret)
}

// ⚠️ You must defer HDC.ReleaseDC().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowdc
func (hWnd HWND) GetWindowDC() HDC {
	ret, _, err := syscall.Syscall(proc.GetWindowDC.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HDC(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
func (hWnd HWND) GetWindowLongPtr(index co.GWLP) uintptr {
	ret, _, err := syscall.Syscall(proc.GetWindowLongPtr.Addr(), 2,
		uintptr(hWnd), uintptr(index), 0)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowrect
func (hWnd HWND) GetWindowRect() RECT {
	var rc RECT
	ret, _, err := syscall.Syscall(proc.GetWindowRect.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(&rc)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return rc
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextw
func (hWnd HWND) GetWindowText() string {
	len := hWnd.GetWindowTextLength() + 1
	buf := make([]uint16, len)

	ret, _, err := syscall.Syscall(proc.GetWindowText.Addr(), 3,
		uintptr(hWnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(len))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return Str.FromNativeSlice(buf)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextlengthw
func (hWnd HWND) GetWindowTextLength() int32 {
	ret, _, _ := syscall.Syscall(proc.GetWindowTextLength.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-hidecaret
func (hWnd HWND) HideCaret() {
	ret, _, err := syscall.Syscall(proc.HideCaret.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-hilitemenuitem
func (hWnd HWND) HiliteMenuItem(hMenu HMENU, item MenuItem, hilite bool) bool {
	idPos, mf := variantMenuItem(item)
	flags := util.Iif(hilite, co.MFS_HILITE, co.MFS_UNHILITE).(co.MF) | mf

	ret, _, _ := syscall.Syscall6(proc.HiliteMenuItem.Addr(), 4,
		uintptr(hWnd), uintptr(hMenu), idPos, uintptr((flags)), 0, 0)
	return ret != 0
}

// Returns the window instance with GetWindowLongPtr().
func (hWnd HWND) Hinstance() HINSTANCE {
	return HINSTANCE(hWnd.GetWindowLongPtr(co.GWLP_HINSTANCE))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invalidaterect
func (hWnd HWND) InvalidateRect(rc *RECT, erase bool) {
	ret, _, err := syscall.Syscall(proc.InvalidateRect.Addr(), 3,
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)),
		util.BoolToUintptr(erase))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-ischild
func (hWnd HWND) IsChild(hChild HWND) bool {
	ret, _, _ := syscall.Syscall(proc.IsChild.Addr(), 2,
		uintptr(hWnd), uintptr(hChild), 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdlgbuttonchecked
func (hWnd HWND) IsDlgButtonChecked(idButton int32) co.BST {
	ret, _, _ := syscall.Syscall(proc.IsDlgButtonChecked.Addr(), 2,
		uintptr(hWnd), uintptr(idButton), 0)
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemedialogtextureenabled
func (hWnd HWND) IsThemeDialogTextureEnabled() bool {
	ret, _, _ := syscall.Syscall(proc.IsThemeDialogTextureEnabled.Addr(), 1,
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-iswindowvisible
func (hWnd HWND) IsWindowVisible() bool {
	ret, _, _ := syscall.Syscall(proc.IsWindowVisible.Addr(), 1,
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
func (hWnd HWND) KillTimer(id uintptr) {
	ret, _, err := syscall.Syscall(proc.KillTimer.Addr(), 2,
		uintptr(hWnd), id, 0)

	if id > 0xffff { // guess: Win32 pointers are greater than WORDs
		delete(_globalTimerFuncs, (*_TimerPack)(unsafe.Pointer(id))) // remove from set
		// At this moment, the callback pointer has no more references. If
		// KillTimer() is called from within the callback itself, it's unsure
		// whether the running function will be enough to keep its pointer from
		// being collected by the GC, but it's reasonable to think so.
	}

	if ret == 0 && errco.ERROR(err) != errco.SUCCESS {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-logicaltophysicalpoint
func (hWnd HWND) LogicalToPhysicalPoint(pt *POINT) {
	ret, _, err := syscall.Syscall(proc.LogicalToPhysicalPoint.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(pt)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapdialogrect
func (hWnd HWND) MapDialogRect(rc *RECT) {
	ret, _, err := syscall.Syscall(proc.MapDialogRect.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-menuitemfrompoint
func (hWnd HWND) MenuItemFromPoint(hMenu HMENU, pt POINT) (int, bool) {
	ret, _, _ := syscall.Syscall6(proc.MenuItemFromPoint.Addr(), 4,
		uintptr(hWnd), uintptr(hMenu), uintptr(pt.X), uintptr(pt.Y), 0, 0)
	return int(ret), int(ret) != -1
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
func (hWnd HWND) MessageBox(text, caption string, uType co.MB) co.ID {
	ret, _, err := syscall.Syscall6(proc.MessageBox.Addr(), 4,
		uintptr(hWnd), uintptr(unsafe.Pointer(Str.ToNativePtr(text))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(caption))), uintptr(uType),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.ID(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfromwindow
func (hWnd HWND) MonitorFromWindow(flags co.MONITOR) HMONITOR {
	ret, _, _ := syscall.Syscall(proc.MonitorFromWindow.Addr(), 2,
		uintptr(hWnd), uintptr(flags), 0)
	return HMONITOR(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-movewindow
func (hWnd HWND) MoveWindow(x, y, width, height int32, repaint bool) {
	ret, _, err := syscall.Syscall6(proc.MoveWindow.Addr(), 6,
		uintptr(hWnd), uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		util.BoolToUintptr(repaint))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HCLIPBOARD.CloseClipboard().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-openclipboard
func (hWnd HWND) OpenClipboard() HCLIPBOARD {
	ret, _, err := syscall.Syscall(proc.OpenClipboard.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HCLIPBOARD{}
}

// ⚠️ You must defer HTHEME.CloseThemeData().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-openthemedata
func (hWnd HWND) OpenThemeData(classNames string) (HTHEME, error) {
	ret, _, err := syscall.Syscall(proc.OpenThemeData.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(Str.ToNativePtr(classNames))),
		0)
	if ret == 0 {
		return HTHEME(0), errco.ERROR(err)
	}
	return HTHEME(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postmessagew
func (hWnd HWND) PostMessage(msg co.WM, wParam WPARAM, lParam LPARAM) {
	ret, _, err := syscall.Syscall6(proc.PostMessage.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-releasedc
func (hWnd HWND) ReleaseDC(hdc HDC) {
	ret, _, err := syscall.Syscall(proc.ReleaseDC.Addr(), 2,
		uintptr(hWnd), uintptr(hdc), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-removewindowsubclass
func (hWnd HWND) RemoveWindowSubclass(
	subclassProc uintptr, idSubclass uint32) {

	ret, _, err := syscall.Syscall(proc.RemoveWindowSubclass.Addr(), 3,
		uintptr(hWnd), subclassProc, uintptr(idSubclass))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientPt(pt *POINT) {
	ret, _, err := syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(pt)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-screentoclient
func (hWnd HWND) ScreenToClientRc(rc *RECT) {
	ret, _, err := syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(rc)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	ret, _, err = syscall.Syscall(proc.ScreenToClient.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(&rc.Right)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew
func (hWnd HWND) SendMessage(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.Syscall6(proc.SendMessage.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

// Returns a handle to the previously focused window.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setfocus
func (hWnd HWND) SetFocus() (HWND, error) {
	ret, _, err := syscall.Syscall(proc.SetFocus.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if hPrev, err := HWND(ret), errco.ERROR(err); hPrev == 0 && err != errco.S_OK {
		return hPrev, err
	} else {
		return hPrev, nil
	}
}

// Returns true if the window was brought to the foreground.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setforegroundwindow
func (hWnd HWND) SetForegroundWindow() bool {
	ret, _, _ := syscall.Syscall(proc.SetForegroundWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

// Returns the handle of the previous parent window.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setparent
func (hWnd HWND) SetParent(hwndNewParent HWND) HWND {
	ret, _, err := syscall.Syscall(proc.SetParent.Addr(), 2,
		uintptr(hWnd), uintptr(hwndNewParent), 0)
	if hPrev, wErr := HWND(ret), errco.ERROR(err); hPrev == 0 && wErr != errco.S_OK {
		panic(wErr)
	} else {
		return hPrev
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenu
func (hWnd HWND) SetMenu(hMenu HMENU) {
	ret, _, err := syscall.Syscall(proc.SetMenu.Addr(), 2,
		uintptr(hWnd), uintptr(hMenu), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// Returns the current position of the scroll box.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setscrollinfo
func (hWnd HWND) SetScrollInfo(
	bar co.SB_TYPE, si *SCROLLINFO, redraw bool) int32 {

	ret, _, _ := syscall.Syscall6(proc.SetScrollInfo.Addr(), 4,
		uintptr(hWnd), uintptr(bar), uintptr(unsafe.Pointer(si)),
		util.BoolToUintptr(redraw), 0, 0)
	return int32(ret)
}

// This method will create a timer that will post WM_TIMER messages, instead of
// running a callback.
//
// The method returns the timer ID.
//
// ⚠️ You must call HWND.KillTimer() to stop the timer.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-settimer
func (hWnd HWND) SetTimer(msElapse uint32, id uintptr) uintptr {
	ret, _, err := syscall.Syscall6(proc.SetTimer.Addr(), 4,
		uintptr(hWnd), id, uintptr(msElapse), 0, 0, 0)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return id
}

// Creates a timer with SetTimer(), which runs the given callback instead of
// posting WM_TIMER messages.
//
// The method returns the timer ID, which is also sent to the callback.
//
// ⚠️ You must call HWND.KillTimer() to stop the timer and free the allocated
// resources.
//
// Example:
//
//  var hWnd HWND // initialized somewhere
//
//  hWnd.SetTimerCallback(2000, func(id uintptr) {
//      println("Run once")
//      hWnd.KillTimer(id)
//  })
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-settimer
func (hWnd HWND) SetTimerCallback(
	msElapse uint32, timerFunc func(id uintptr)) uintptr {

	pPack := &_TimerPack{f: timerFunc}
	if _globalTimerFuncs == nil {
		_globalTimerFuncs = make(map[*_TimerPack]struct{}, 1)
	}
	_globalTimerFuncs[pPack] = struct{}{} // store pointer in the set

	id := uintptr(unsafe.Pointer(pPack)) // use the pointer as the timer ID

	ret, _, err := syscall.Syscall6(proc.SetTimer.Addr(), 4,
		uintptr(hWnd), id, uintptr(msElapse), _globalTimerCallback, 0, 0)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return id
}

type _TimerPack struct{ f func(id uintptr) }

var (
	_globalTimerCallback uintptr = syscall.NewCallback(_TimerProc)
	_globalTimerFuncs    map[*_TimerPack]struct{}
)

func _TimerProc(_ HWND, _ co.WM, id uintptr, _ uint32) uintptr {
	pPack := (*_TimerPack)(unsafe.Pointer(id))
	if _, isStored := _globalTimerFuncs[pPack]; isStored {
		pPack.f(id)
	}
	return 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowlongptrw
func (hWnd HWND) SetWindowLongPtr(index co.GWLP, newLong uintptr) uintptr {
	syscall.Syscall(proc.SetLastError.Addr(), 0,
		0, 0, 0)

	ret, _, err := syscall.Syscall(proc.SetWindowLongPtr.Addr(), 3,
		uintptr(hWnd), uintptr(index), newLong)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return ret
}

// You can pass HWND or HWND_IA in hwndInsertAfter argument.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
func (hWnd HWND) SetWindowPos(
	hwndInsertAfter HWND, x, y, cx, cy int32, flags co.SWP) {

	ret, _, err := syscall.Syscall9(proc.SetWindowPos.Addr(), 7,
		uintptr(hWnd), uintptr(hwndInsertAfter),
		uintptr(x), uintptr(y), uintptr(cx), uintptr(cy),
		uintptr(flags), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-setwindowsubclass
func (hWnd HWND) SetWindowSubclass(
	subclassProc uintptr, idSubclass uint32, refData unsafe.Pointer) {

	ret, _, err := syscall.Syscall6(proc.SetWindowSubclass.Addr(), 4,
		uintptr(hWnd), subclassProc, uintptr(idSubclass), uintptr(refData),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowtextw
func (hWnd HWND) SetWindowText(text string) {
	syscall.Syscall(proc.SetWindowText.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(Str.ToNativePtr(text))),
		0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showcaret
func (hWnd HWND) ShowCaret() {
	ret, _, err := syscall.Syscall(proc.ShowCaret.Addr(), 1,
		uintptr(hWnd), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
func (hWnd HWND) ShowWindow(cmdShow co.SW) bool {
	ret, _, _ := syscall.Syscall(proc.ShowWindow.Addr(), 1,
		uintptr(hWnd), uintptr(cmdShow), 0)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialog
func (hWnd HWND) TaskDialog(
	hInstance HINSTANCE,
	windowTitle, mainInstruction, content StrOrNil,
	commonButtons co.TDCBF, icon co.TD_ICON) co.ID {

	var pnButton int32
	ret, _, _ := syscall.Syscall9(proc.TaskDialog.Addr(), 8,
		uintptr(hWnd), uintptr(hInstance),
		uintptr(variantStrOrNil(windowTitle)),
		uintptr(variantStrOrNil(mainInstruction)),
		uintptr(variantStrOrNil(content)),
		uintptr(commonButtons), uintptr(icon),
		uintptr(unsafe.Pointer(&pnButton)), 0)
	if wErr := errco.ERROR(ret); wErr != errco.S_OK {
		panic(wErr)
	}
	return co.ID(pnButton)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translateacceleratorw
func (hWnd HWND) TranslateAccelerator(hAccel HACCEL, msg *MSG) error {
	ret, _, err := syscall.Syscall(proc.TranslateAccelerator.Addr(), 3,
		uintptr(hWnd), uintptr(hAccel), uintptr(unsafe.Pointer(msg)))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-updatewindow
func (hWnd HWND) UpdateWindow() bool {
	ret, _, _ := syscall.Syscall(proc.UpdateWindow.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}
