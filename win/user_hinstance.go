//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/wstr"
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
	var wTemplateName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_CreateDialogParamW, "CreateDialogParamW"),
		uintptr(hInst),
		templateName.raw(&wTemplateName),
		uintptr(hwndParent),
		dialogFunc,
		uintptr(dwInitParam))
	if ret == 0 {
		return HWND(0), co.ERROR(err)
	}
	return HWND(ret), nil
}

var _user_CreateDialogParamW *syscall.Proc

// [DialogBoxIndirectParam] function.
//
// [DialogBoxIndirectParam]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dialogboxindirectparamw
func (hInst HINSTANCE) DialogBoxIndirectParam(
	template *DLGTEMPLATE,
	hwndParent HWND,
	dialogFunc uintptr,
	dwInitParam LPARAM,
) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_DialogBoxIndirectParamW, "DialogBoxIndirectParamW"),
		uintptr(hInst),
		uintptr(unsafe.Pointer(template)),
		uintptr(hwndParent),
		dialogFunc,
		uintptr(dwInitParam))
	if int32(ret) == -1 && co.ERROR(err) != co.ERROR_SUCCESS {
		return 0, co.ERROR(err)
	}
	return ret, nil
}

var _user_DialogBoxIndirectParamW *syscall.Proc

// [DialogBoxParam] function.
//
// [DialogBoxParam]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dialogboxparamw
func (hInst HINSTANCE) DialogBoxParam(
	templateName ResId,
	hwndParent HWND,
	dialogFunc uintptr,
	dwInitParam LPARAM,
) (uintptr, error) {
	var wTemplateName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_DialogBoxParamW, "DialogBoxParamW"),
		uintptr(hInst),
		templateName.raw(&wTemplateName),
		uintptr(hwndParent),
		dialogFunc,
		uintptr(dwInitParam))
	if int32(ret) == -1 && co.ERROR(err) != co.ERROR_SUCCESS {
		return 0, co.ERROR(err)
	}
	return ret, nil
}

var _user_DialogBoxParamW *syscall.Proc

// [GetClassInfoEx] function.
//
// [GetClassInfoEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassinfoexw
func (hInst HINSTANCE) GetClassInfoEx(
	className *uint16,
	destBuf *WNDCLASSEX,
) (ATOM, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_GetClassInfoExW, "GetClassInfoExW"),
		uintptr(hInst),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(destBuf)))
	if ret == 0 {
		return ATOM(0), co.ERROR(err)
	}
	return ATOM(ret), nil
}

var _user_GetClassInfoExW *syscall.Proc

// [LoadAccelerators] function.
//
// Accelerator tables loaded from resources are shared, and don't need to be
// deleted.
//
// [LoadAccelerators]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadacceleratorsw
func (hInst HINSTANCE) LoadAccelerators(tableName ResId) (HACCEL, error) {
	var wTableName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_LoadAcceleratorsW, "LoadAcceleratorsW"),
		uintptr(hInst),
		tableName.raw(&wTableName))
	if ret == 0 {
		return HACCEL(0), co.ERROR(err)
	}
	return HACCEL(ret), nil
}

var _user_LoadAcceleratorsW *syscall.Proc

// [LoadCursor] function.
//
// Cursors loaded from resources are shared, and don't need to be deleted.
//
// [LoadCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
func (hInst HINSTANCE) LoadCursor(cursorName CursorRes) (HCURSOR, error) {
	var wCursorName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_LoadCursorW, "LoadCursorW"),
		uintptr(hInst),
		cursorName.raw(&wCursorName))
	if ret == 0 {
		return HCURSOR(0), co.ERROR(err)
	}
	return HCURSOR(ret), nil
}

var _user_LoadCursorW *syscall.Proc

// [LoadIcon] function.
//
// Icons loaded from resources are shared, and don't need to be deleted.
//
// Example:
//
// Loading an icon from the resource:
//
//	hInst, _ := win.GetModuleHandle("")
//	hIco, _ := hInst.LoadIcon(win.IconResId(101))
//
// [LoadIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
func (hInst HINSTANCE) LoadIcon(iconName IconRes) (HICON, error) {
	var wIconName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_LoadIconW, "LoadIconW"),
		uintptr(hInst),
		iconName.raw(&wIconName))
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

var _user_LoadIconW *syscall.Proc

// [LoadImage] function.
//
// Returned [HGDIOBJ] must be cast into [HBITMAP], [HCURSOR] or [HICON].
//
// ⚠️ If the object is not being loaded from the application resources with
// [co.LR_SHARED], you must defer its respective free method:
//   - [HBITMAP.DeleteObject]
//   - [HCURSOR.DestroyCursor]
//   - [HICON.DestroyIcon]
//
// Example:
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
	cx, cy int,
	fuLoad co.LR,
) (HGDIOBJ, error) {
	var wName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_LoadImageW, "LoadImageW"),
		uintptr(hInst),
		name.raw(&wName),
		uintptr(imgType),
		uintptr(int32(cx)),
		uintptr(int32(cy)),
		uintptr(fuLoad))
	if ret == 0 {
		return HGDIOBJ(0), co.ERROR(err)
	}
	return HGDIOBJ(ret), nil
}

var _user_LoadImageW *syscall.Proc

// [LoadMenu] function.
//
// ⚠️ You must defer [HMENU.DestroyMenu].
//
// [LoadMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadmenuw
func (hInst HINSTANCE) LoadMenu(menuName ResId) (HMENU, error) {
	var wMenuName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_user_LoadMenuW, "LoadMenuW"),
		uintptr(hInst),
		menuName.raw(&wMenuName))
	if ret == 0 {
		return HMENU(0), co.ERROR(err)
	}
	return HMENU(ret), nil
}

var _user_LoadMenuW *syscall.Proc
