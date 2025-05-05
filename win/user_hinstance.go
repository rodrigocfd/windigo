//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [CreateDialogParam] function.
//
// [CreateDialogParam]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createdialogparamw
func (hInst HINSTANCE) CreateDialogParam(
	templateName ResId,
	hwndParent HWND,
	dialogFunc uintptr,
	dwInitParam LPARAM,
) (HWND, error) {
	templateName16 := wstr.NewBuf[wstr.Stack20]()
	templateNameVal := templateName.raw(&templateName16)

	ret, _, err := syscall.SyscallN(_CreateDialogParamW.Addr(),
		uintptr(hInst), templateNameVal,
		uintptr(hwndParent), dialogFunc, uintptr(dwInitParam))
	if ret == 0 {
		return HWND(0), co.ERROR(err)
	}
	return HWND(ret), nil
}

var _CreateDialogParamW = dll.User32.NewProc("CreateDialogParamW")

// [DialogBoxIndirectParam] function.
//
// [DialogBoxIndirectParam]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dialogboxindirectparamw
func (hInst HINSTANCE) DialogBoxIndirectParam(
	template *DLGTEMPLATE,
	hwndParent HWND,
	dialogFunc uintptr,
	dwInitParam LPARAM,
) (uintptr, error) {
	ret, _, err := syscall.SyscallN(_DialogBoxIndirectParamW.Addr(),
		uintptr(hInst), uintptr(unsafe.Pointer(template)),
		uintptr(hwndParent), dialogFunc, uintptr(dwInitParam))
	if int(ret) == -1 && co.ERROR(err) != co.ERROR_SUCCESS {
		return 0, co.ERROR(err)
	}
	return ret, nil
}

var _DialogBoxIndirectParamW = dll.User32.NewProc("DialogBoxIndirectParamW")

// [DialogBoxParam] function.
//
// [DialogBoxParam]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dialogboxparamw
func (hInst HINSTANCE) DialogBoxParam(
	templateName ResId,
	hwndParent HWND,
	dialogFunc uintptr,
	dwInitParam LPARAM,
) (uintptr, error) {
	templateName16 := wstr.NewBuf[wstr.Stack20]()
	templateNameVal := templateName.raw(&templateName16)

	ret, _, err := syscall.SyscallN(_DialogBoxParamW.Addr(),
		uintptr(hInst), templateNameVal,
		uintptr(hwndParent), dialogFunc, uintptr(dwInitParam))
	if int(ret) == -1 && co.ERROR(err) != co.ERROR_SUCCESS {
		return 0, co.ERROR(err)
	}
	return ret, nil
}

var _DialogBoxParamW = dll.User32.NewProc("DialogBoxParamW")

// [GetClassInfoEx] function.
//
// [GetClassInfoEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassinfoexw
func (hInst HINSTANCE) GetClassInfoEx(
	className *uint16,
	destBuf *WNDCLASSEX,
) (ATOM, error) {
	ret, _, err := syscall.SyscallN(_GetClassInfoExW.Addr(),
		uintptr(hInst),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(destBuf)))
	if ret == 0 {
		return ATOM(0), co.ERROR(err)
	}
	return ATOM(ret), nil
}

var _GetClassInfoExW = dll.User32.NewProc("GetClassInfoExW")

// [LoadAccelerators] function.
//
// Accelerator tables loaded from resources are shared, and don't need to be
// deleted.
//
// [LoadAccelerators]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadacceleratorsw
func (hInst HINSTANCE) LoadAccelerators(tableName ResId) (HACCEL, error) {
	tableName16 := wstr.NewBuf[wstr.Stack20]()
	tableNameVal := tableName.raw(&tableName16)

	ret, _, err := syscall.SyscallN(_LoadAcceleratorsW.Addr(),
		uintptr(hInst), tableNameVal)
	if ret == 0 {
		return HACCEL(0), co.ERROR(err)
	}
	return HACCEL(ret), nil
}

var _LoadAcceleratorsW = dll.User32.NewProc("LoadAcceleratorsW")

// [LoadCursor] function.
//
// Icons loaded from resources are shared, and don't need to be deleted.
//
// [LoadCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
func (hInst HINSTANCE) LoadCursor(cursorName CursorRes) (HCURSOR, error) {
	cursorName16 := wstr.NewBuf[wstr.Stack20]()
	cursorNameVal := cursorName.raw(&cursorName16)

	ret, _, err := syscall.SyscallN(_LoadCursorW.Addr(),
		uintptr(hInst), cursorNameVal)
	if ret == 0 {
		return HCURSOR(0), co.ERROR(err)
	}
	return HCURSOR(ret), nil
}

var _LoadCursorW = dll.User32.NewProc("LoadCursorW")

// [LoadIcon] function.
//
// Icons loaded from resources are shared, and don't need to be deleted.
//
// # Example
//
// Loading an icon from the resource:
//
//	hInst, _ := win.GetModuleHandle("")
//	hIco, _ := hInst.LoadIcon(win.IconResId(101))
//
// [LoadIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
func (hInst HINSTANCE) LoadIcon(iconName IconRes) (HICON, error) {
	iconName16 := wstr.NewBuf[wstr.Stack20]()
	iconNameVal := iconName.raw(&iconName16)

	ret, _, err := syscall.SyscallN(_LoadIconW.Addr(),
		uintptr(hInst), iconNameVal)
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

var _LoadIconW = dll.User32.NewProc("LoadIconW")

// [LoadImage] function.
//
// Returned HGDIOBJ must be cast into HBITMAP, HCURSOR or HICON.
//
// ⚠️ If the object is not being loaded from the application resources with
// co.LR_SHARED, you must defer its respective DeleteObject().
//
// # Examples
//
// Loading a 16x16 icon resource:
//
//	const MY_ICON_ID uint16 = 101
//
//	hInst, _ := win.GetModuleHandle("")
//	hGdi, _ := hInst.LoadImage(
//		win.ResIdInt(MY_ICON_ID),
//		co.IMAGE_ICON,
//		16, 16,
//		co.LR_DEFAULTCOLOR | co.LR_SHARED,
//	)
//	hIcon := win.HICON(hGdi)
//
// Loading a bitmap from a file:
//
//	hGdi, _ := win.HINSTANCE(0).LoadImage(
//		win.ResIdStr("C:\\Temp\\image.bmp"),
//		co.IMAGE_BITMAP,
//		0, 0,
//		co.LR_LOADFROMFILE,
//	)
//	hBmp := win.HBITMAP(hGdi)
//	defer hBmp.DeleteObject()
//
// [LoadImage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
func (hInst HINSTANCE) LoadImage(
	name ResId,
	imgType co.IMAGE,
	cx, cy uint,
	fuLoad co.LR,
) (HGDIOBJ, error) {
	name16 := wstr.NewBuf[wstr.Stack20]()
	nameVal := name.raw(&name16)

	ret, _, err := syscall.SyscallN(_LoadImageW.Addr(),
		uintptr(hInst), nameVal, uintptr(imgType),
		uintptr(cx), uintptr(cy), uintptr(fuLoad))
	if ret == 0 {
		return HGDIOBJ(0), co.ERROR(err)
	}
	return HGDIOBJ(ret), nil
}

var _LoadImageW = dll.User32.NewProc("LoadImageW")

// [LoadMenu] function.
//
// ⚠️ You must defer HMENU.DestroyMenu().
//
// [LoadMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadmenuw
func (hInst HINSTANCE) LoadMenu(menuName ResId) (HMENU, error) {
	menuName16 := wstr.NewBuf[wstr.Stack20]()
	menuNameVal := menuName.raw(&menuName16)

	ret, _, err := syscall.SyscallN(_LoadMenuW.Addr(),
		uintptr(hInst), menuNameVal)
	if ret == 0 {
		return HMENU(0), co.ERROR(err)
	}
	return HMENU(ret), nil
}

var _LoadMenuW = dll.User32.NewProc("LoadMenuW")
