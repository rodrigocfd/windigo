package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

// A handle to an instance. This is the base address of the module in memory.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hinstance
type HINSTANCE HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createdialogparamw
func (hInst HINSTANCE) CreateDialogParam(
	lpTemplateName int32, hWndParent HWND,
	lpDialogFunc uintptr, dwInitParam LPARAM) HWND {

	ret, _, lerr := syscall.Syscall6(proc.CreateDialogParam.Addr(), 5,
		uintptr(hInst), uintptr(lpTemplateName), uintptr(hWndParent),
		lpDialogFunc, uintptr(dwInitParam), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dialogboxparamw
func (hInst HINSTANCE) DialogBoxParam(
	lpTemplateName int32, hWndParent HWND,
	lpDialogFunc uintptr, dwInitParam LPARAM) uintptr {

	ret, _, lerr := syscall.Syscall6(proc.DialogBoxParam.Addr(), 5,
		uintptr(hInst), uintptr(lpTemplateName), uintptr(hWndParent),
		lpDialogFunc, uintptr(dwInitParam), 0)
	if int(ret) == -1 && err.ERROR(lerr) != err.SUCCESS {
		panic(err.ERROR(lerr))
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-duplicateicon
func (hInst HINSTANCE) DuplicateIcon(hIcon HICON) HICON {
	ret, _, lerr := syscall.Syscall(proc.DuplicateIcon.Addr(), 2,
		uintptr(hInst), uintptr(hIcon), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HICON(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-freelibrary
func (hInst HINSTANCE) FreeLibrary() {
	ret, _, lerr := syscall.Syscall(proc.FreeLibrary.Addr(), 1,
		uintptr(hInst), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassinfoexw
func (hInst HINSTANCE) GetClassInfoEx(
	className *uint16, destBuf *WNDCLASSEX) (ATOM, error) {

	ret, _, lerr := syscall.Syscall(proc.GetClassInfoEx.Addr(), 3,
		uintptr(hInst),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(destBuf)))
	if ret == 0 {
		return ATOM(0), err.ERROR(lerr)
	}
	return ATOM(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getprocaddress
func (hInst HINSTANCE) GetProcAddress(lpProcName string) uintptr {
	ascii := []byte(lpProcName)
	ascii = append(ascii, 0x00) // terminating null

	ret, _, lerr := syscall.Syscall(proc.GetProcAddress.Addr(), 2,
		uintptr(hInst), uintptr(unsafe.Pointer(&ascii[0])), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadacceleratorsw
func (hInst HINSTANCE) LoadAccelerators(lpTableName uintptr) HACCEL {
	ret, _, lerr := syscall.Syscall(proc.LoadAccelerators.Addr(), 2,
		uintptr(hInst), uintptr(lpTableName), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HACCEL(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
func (hInst HINSTANCE) LoadCursor(lpCursorName co.IDC) HCURSOR {
	ret, _, lerr := syscall.Syscall(proc.LoadCursor.Addr(), 2,
		uintptr(hInst), uintptr(lpCursorName), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HCURSOR(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
func (hInst HINSTANCE) LoadIcon(lpIconName co.IDI) HICON {
	ret, _, lerr := syscall.Syscall(proc.LoadIcon.Addr(), 2,
		uintptr(hInst), uintptr(lpIconName), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HICON(ret)
}

// Returned HANDLE must be cast into HBITMAP, HCURSOR or HICON.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
func (hInst HINSTANCE) LoadImage(
	name int32, imgType co.IMAGE, cx, cy int32, fuLoad co.LR) HANDLE {

	ret, _, lerr := syscall.Syscall6(proc.LoadImage.Addr(), 6,
		uintptr(hInst), uintptr(name), uintptr(imgType),
		uintptr(cx), uintptr(cy), uintptr(fuLoad))
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HANDLE(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadmenuw
func (hInst HINSTANCE) LoadMenu(lpMenuName int32) HMENU {
	ret, _, lerr := syscall.Syscall(proc.LoadMenu.Addr(), 2,
		uintptr(hInst), uintptr(lpMenuName), 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return HMENU(ret)
}
