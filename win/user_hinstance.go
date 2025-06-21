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
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pTemplateName := templateName.raw(&wbuf)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_CreateDialogParamW, "CreateDialogParamW"),
		uintptr(hInst),
		pTemplateName,
		uintptr(hwndParent),
		dialogFunc,
		uintptr(dwInitParam))
	if ret == 0 {
		return HWND(0), co.ERROR(err)
	}
	return HWND(ret), nil
}

var _CreateDialogParamW *syscall.Proc

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
		dll.Load(dll.USER32, &_DialogBoxIndirectParamW, "DialogBoxIndirectParamW"),
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

var _DialogBoxIndirectParamW *syscall.Proc

// [DialogBoxParam] function.
//
// [DialogBoxParam]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dialogboxparamw
func (hInst HINSTANCE) DialogBoxParam(
	templateName ResId,
	hwndParent HWND,
	dialogFunc uintptr,
	dwInitParam LPARAM,
) (uintptr, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pTemplateName := templateName.raw(&wbuf)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_DialogBoxParamW, "DialogBoxParamW"),
		uintptr(hInst),
		pTemplateName,
		uintptr(hwndParent),
		dialogFunc,
		uintptr(dwInitParam))
	if int32(ret) == -1 && co.ERROR(err) != co.ERROR_SUCCESS {
		return 0, co.ERROR(err)
	}
	return ret, nil
}

var _DialogBoxParamW *syscall.Proc

// [GetClassInfoEx] function.
//
// [GetClassInfoEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassinfoexw
func (hInst HINSTANCE) GetClassInfoEx(
	className *uint16,
	destBuf *WNDCLASSEX,
) (ATOM, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetClassInfoExW, "GetClassInfoExW"),
		uintptr(hInst),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(destBuf)))
	if ret == 0 {
		return ATOM(0), co.ERROR(err)
	}
	return ATOM(ret), nil
}

var _GetClassInfoExW *syscall.Proc

// [LoadAccelerators] function.
//
// Accelerator tables loaded from resources are shared, and don't need to be
// deleted.
//
// [LoadAccelerators]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadacceleratorsw
func (hInst HINSTANCE) LoadAccelerators(tableName ResId) (HACCEL, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pTableName := tableName.raw(&wbuf)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_LoadAcceleratorsW, "LoadAcceleratorsW"),
		uintptr(hInst),
		pTableName)
	if ret == 0 {
		return HACCEL(0), co.ERROR(err)
	}
	return HACCEL(ret), nil
}

var _LoadAcceleratorsW *syscall.Proc

// [LoadCursor] function.
//
// Icons loaded from resources are shared, and don't need to be deleted.
//
// [LoadCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
func (hInst HINSTANCE) LoadCursor(cursorName CursorRes) (HCURSOR, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pCursorName := cursorName.raw(&wbuf)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_LoadCursorW, "LoadCursorW"),
		uintptr(hInst),
		pCursorName)
	if ret == 0 {
		return HCURSOR(0), co.ERROR(err)
	}
	return HCURSOR(ret), nil
}

var _LoadCursorW *syscall.Proc

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
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pIconName := iconName.raw(&wbuf)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_LoadIconW, "LoadIconW"),
		uintptr(hInst),
		pIconName)
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

var _LoadIconW *syscall.Proc

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
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pName := name.raw(&wbuf)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_LoadImageW, "LoadImageW"),
		uintptr(hInst),
		pName,
		uintptr(imgType),
		uintptr(int32(cx)),
		uintptr(int32(cy)),
		uintptr(fuLoad))
	if ret == 0 {
		return HGDIOBJ(0), co.ERROR(err)
	}
	return HGDIOBJ(ret), nil
}

var _LoadImageW *syscall.Proc

// [LoadMenu] function.
//
// ⚠️ You must defer [HMENU.DestroyMenu].
//
// [LoadMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadmenuw
func (hInst HINSTANCE) LoadMenu(menuName ResId) (HMENU, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pMenuName := menuName.raw(&wbuf)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.USER32, &_LoadMenuW, "LoadMenuW"),
		uintptr(hInst),
		pMenuName)
	if ret == 0 {
		return HMENU(0), co.ERROR(err)
	}
	return HMENU(ret), nil
}

var _LoadMenuW *syscall.Proc
