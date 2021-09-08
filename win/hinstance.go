package win

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to an instance. This is the base address of the module in memory.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hinstance
type HINSTANCE HANDLE

// ⚠️ lpModuleName must be string or nil.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandle(moduleName interface{}) HINSTANCE {
	ret, _, err := syscall.Syscall(proc.GetModuleHandle.Addr(), 1,
		uintptr(util.VariantNilString(moduleName)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HINSTANCE(ret)
}

// ⚠️ You must defer HINSTANCE.FreeLibrary().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadlibraryw
func LoadLibrary(libFileName string) HINSTANCE {
	ret, _, err := syscall.Syscall(proc.LoadLibrary.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(libFileName))),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HINSTANCE(ret)
}

// ⚠️ lpTemplateName must be uint16 or string.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createdialogparamw
func (hInst HINSTANCE) CreateDialogParam(
	templateName interface{}, hwndParent HWND,
	dialogFunc uintptr, dwInitParam LPARAM) HWND {

	ret, _, err := syscall.Syscall6(proc.CreateDialogParam.Addr(), 5,
		uintptr(hInst), uintptr(util.VariantUint16String(templateName)),
		uintptr(hwndParent), dialogFunc, uintptr(dwInitParam), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HWND(ret)
}

// ⚠️ lpTemplateName must be uint16 or string.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dialogboxparamw
func (hInst HINSTANCE) DialogBoxParam(
	templateName interface{}, hwndParent HWND,
	dialogFunc uintptr, dwInitParam LPARAM) uintptr {

	ret, _, err := syscall.Syscall6(proc.DialogBoxParam.Addr(), 5,
		uintptr(hInst), uintptr(util.VariantUint16String(templateName)),
		uintptr(hwndParent), dialogFunc, uintptr(dwInitParam), 0)
	if int(ret) == -1 && errco.ERROR(err) != errco.SUCCESS {
		panic(errco.ERROR(err))
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-duplicateicon
func (hInst HINSTANCE) DuplicateIcon(hIcon HICON) HICON {
	ret, _, err := syscall.Syscall(proc.DuplicateIcon.Addr(), 2,
		uintptr(hInst), uintptr(hIcon), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HICON(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-freelibrary
func (hInst HINSTANCE) FreeLibrary() {
	ret, _, err := syscall.Syscall(proc.FreeLibrary.Addr(), 1,
		uintptr(hInst), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew
func (hInst HINSTANCE) GetModuleFileName() string {
	buf := [_MAX_PATH + 1]uint16{}
	ret, _, err := syscall.Syscall(proc.GetModuleFileName.Addr(), 3,
		uintptr(hInst), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromUint16Slice(buf[:])
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getprocaddress
func (hInst HINSTANCE) GetProcAddress(procName string) uintptr {
	ascii := []byte(procName)
	ascii = append(ascii, 0x00) // terminating null

	ret, _, err := syscall.Syscall(proc.GetProcAddress.Addr(), 2,
		uintptr(hInst), uintptr(unsafe.Pointer(&ascii[0])), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return ret
}

// ⚠️ lpTableName must be uint16 or string.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadacceleratorsw
func (hInst HINSTANCE) LoadAccelerators(tableName interface{}) HACCEL {
	ret, _, err := syscall.Syscall(proc.LoadAccelerators.Addr(), 2,
		uintptr(hInst), uintptr(util.VariantUint16String(tableName)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HACCEL(ret)
}

// ⚠️ lpCursorName must be uint16, co.IDC or string.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
func (hInst HINSTANCE) LoadCursor(cursorName interface{}) HCURSOR {
	var pName unsafe.Pointer
	switch v := cursorName.(type) {
	case uint16:
		pName = unsafe.Pointer(uintptr(v))
	case co.IDC:
		pName = unsafe.Pointer(uintptr(v))
	case string:
		pName = unsafe.Pointer(Str.ToUint16Ptr(v))
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(cursorName)))
	}

	ret, _, err := syscall.Syscall(proc.LoadCursor.Addr(), 2,
		uintptr(hInst), uintptr(pName), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HCURSOR(ret)
}

// ⚠️ lpIconName must be uint16, co.IDI or string.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
func (hInst HINSTANCE) LoadIcon(iconName interface{}) HICON {
	var pName unsafe.Pointer
	switch v := iconName.(type) {
	case uint16:
		pName = unsafe.Pointer(uintptr(v))
	case co.IDI:
		pName = unsafe.Pointer(uintptr(v))
	case string:
		pName = unsafe.Pointer(Str.ToUint16Ptr(v))
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(iconName)))
	}

	ret, _, err := syscall.Syscall(proc.LoadIcon.Addr(), 2,
		uintptr(hInst), uintptr(pName), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HICON(ret)
}

// Returned HANDLE can be cast into HBITMAP, HCURSOR or HICON.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
func (hInst HINSTANCE) LoadImage(
	name int32, imgType co.IMAGE, cx, cy int32, fuLoad co.LR) HANDLE {

	ret, _, err := syscall.Syscall6(proc.LoadImage.Addr(), 6,
		uintptr(hInst), uintptr(name), uintptr(imgType),
		uintptr(cx), uintptr(cy), uintptr(fuLoad))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HANDLE(ret)
}

// ⚠️ lpMenuName must be uint16 or string.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadmenuw
func (hInst HINSTANCE) LoadMenu(menuName interface{}) HMENU {
	ret, _, err := syscall.Syscall(proc.LoadMenu.Addr(), 2,
		uintptr(hInst), uintptr(util.VariantUint16String(menuName)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HMENU(ret)
}
