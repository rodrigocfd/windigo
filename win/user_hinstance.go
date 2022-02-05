package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createdialogparamw
func (hInst HINSTANCE) CreateDialogParam(
	templateName ResId, hwndParent HWND,
	dialogFunc uintptr, dwInitParam LPARAM) HWND {

	templateNameVal, templateNameBuf := variantResId(templateName)
	ret, _, err := syscall.Syscall6(proc.CreateDialogParam.Addr(), 5,
		uintptr(hInst), templateNameVal,
		uintptr(hwndParent), dialogFunc, uintptr(dwInitParam), 0)
	runtime.KeepAlive(templateNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dialogboxparamw
func (hInst HINSTANCE) DialogBoxParam(
	templateName ResId, hwndParent HWND,
	dialogFunc uintptr, dwInitParam LPARAM) uintptr {

	templateNameVal, templateNameBuf := variantResId(templateName)
	ret, _, err := syscall.Syscall6(proc.DialogBoxParam.Addr(), 5,
		uintptr(hInst), templateNameVal,
		uintptr(hwndParent), dialogFunc, uintptr(dwInitParam), 0)
	runtime.KeepAlive(templateNameBuf)
	if int(ret) == -1 && errco.ERROR(err) != errco.SUCCESS {
		panic(errco.ERROR(err))
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassinfoexw
func (hInst HINSTANCE) GetClassInfoEx(
	className *uint16, destBuf *WNDCLASSEX) (ATOM, error) {

	ret, _, err := syscall.Syscall(proc.GetClassInfoEx.Addr(), 3,
		uintptr(hInst),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(destBuf)))
	if ret == 0 {
		return ATOM(0), errco.ERROR(err)
	}
	return ATOM(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadacceleratorsw
func (hInst HINSTANCE) LoadAccelerators(tableName ResId) HACCEL {
	tableNameVal, tableNameBuf := variantResId(tableName)
	ret, _, err := syscall.Syscall(proc.LoadAccelerators.Addr(), 2,
		uintptr(hInst), tableNameVal, 0)
	runtime.KeepAlive(tableNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HACCEL(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
func (hInst HINSTANCE) LoadCursor(cursorName CursorRes) HCURSOR {
	cursorNameVal, cursorNameBuf := variantCursorResId(cursorName)
	ret, _, err := syscall.Syscall(proc.LoadCursor.Addr(), 2,
		uintptr(hInst), cursorNameVal, 0)
	runtime.KeepAlive(cursorNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HCURSOR(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
func (hInst HINSTANCE) LoadIcon(iconName IconRes) HICON {
	iconNameVal, iconNameBuf := variantIconResId(iconName)
	ret, _, err := syscall.Syscall(proc.LoadIcon.Addr(), 2,
		uintptr(hInst), iconNameVal, 0)
	runtime.KeepAlive(iconNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HICON(ret)
}

// Returned HGDIOBJ must be cast into HBITMAP, HCURSOR or HICON.
//
// ⚠️ If the object is not being loaded from the application resources, you must
// defer its respective DeleteObject().
//
// Example loading a 16x16 icon resource:
//
//  const MY_ICON_ID int = 101
//
//  hIcon := win.HICON(
//      win.GetModuleHandle(nil).LoadImage(
//          win.ResIdInt(MY_ICON_ID),
//          co.IMAGE_ICON,
//          16, 16,
//          co.LR_DEFAULTCOLOR,
//      ),
//  )
//
// Example loading a bitmap from file:
//
//  hBmp := win.HBITMAP(
//      win.HINSTANCE(0).LoadImage(
//          win.ResIdStr("C:\\Temp\\image.bmp"),
//          co.IMAGE_BITMAP,
//          0, 0,
//          co.LR_LOADFROMFILE,
//      ),
//  )
//  defer hBmp.DeleteObject()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
func (hInst HINSTANCE) LoadImage(
	name ResId, imgType co.IMAGE, cx, cy int32, fuLoad co.LR) HGDIOBJ {

	nameVal, nameBuf := variantResId(name)
	ret, _, err := syscall.Syscall6(proc.LoadImage.Addr(), 6,
		uintptr(hInst), nameVal, uintptr(imgType),
		uintptr(cx), uintptr(cy), uintptr(fuLoad))
	runtime.KeepAlive(nameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HGDIOBJ(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadmenuw
func (hInst HINSTANCE) LoadMenu(menuName ResId) HMENU {
	menuNameVal, menuNameBuf := variantResId(menuName)
	ret, _, err := syscall.Syscall(proc.LoadMenu.Addr(), 2,
		uintptr(hInst), menuNameVal, 0)
	runtime.KeepAlive(menuNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HMENU(ret)
}
