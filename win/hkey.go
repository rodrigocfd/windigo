/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hkey
type HKEY HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regclosekey
func (hKey HKEY) RegCloseKey() {
	ret := hKey.regCloseKeyNoPanic()
	if ret != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegCloseKey failed. %s", ret.Error()))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenumvaluew
func (hKey HKEY) RegEnumValue(dwIndex uint32,
	lpValueName []uint16, lpcchValueName *uint32, lpType *co.REG,
	lpData unsafe.Pointer, lpcbData *uint32) co.ERROR {

	ret, _, _ := syscall.Syscall9(proc.RegEnumValue.Addr(), 8,
		uintptr(hKey), uintptr(dwIndex),
		uintptr(unsafe.Pointer(&lpValueName[0])),
		uintptr(unsafe.Pointer(lpcchValueName)), 0,
		uintptr(unsafe.Pointer(lpType)),
		uintptr(lpData), uintptr(unsafe.Pointer(lpcbData)), 0)

	lerr := co.ERROR(ret)
	if lerr == co.ERROR_SUCCESS ||
		lerr == co.ERROR_NO_MORE_ITEMS ||
		lerr == co.ERROR_MORE_DATA {
		// These are not really errors.
		return lerr
	}

	hKey.regCloseKeyNoPanic() // free resource
	panic(fmt.Sprintf("RegEnumValue failed. %s", lerr.Error()))
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regopenkeyexw
func RegOpenKeyEx(hKeyPredef co.HKEY, lpSubKey string, ulOptions co.REG_OPTION,
	samDesired co.KEY) (HKEY, co.ERROR) {

	hKey := HKEY(0)
	ret, _, _ := syscall.Syscall6(proc.RegOpenKeyEx.Addr(), 5,
		uintptr(hKeyPredef), uintptr(unsafe.Pointer(StrToPtr(lpSubKey))),
		uintptr(ulOptions), uintptr(samDesired), uintptr(unsafe.Pointer(&hKey)),
		0)
	return hKey, co.ERROR(ret)
}

// https://www.google.com/search?client=firefox-b-d&q=RegQueryValueExW
func (hKey HKEY) RegQueryValueEx(lpValueName string, lpType *co.REG,
	lpData unsafe.Pointer, lpcbData *uint32) co.ERROR {

	ret, _, _ := syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(StrToPtr(lpValueName))), 0,
		uintptr(unsafe.Pointer(lpType)), uintptr(lpData),
		uintptr(unsafe.Pointer(lpcbData)))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_SUCCESS {
		hKey.regCloseKeyNoPanic() // free resource
		panic(fmt.Sprintf("RegQueryValueEx failed. %s", lerr.Error()))
	}
	return lerr
}

func (hKey HKEY) regCloseKeyNoPanic() co.ERROR {
	if hKey == 0 { // handle is null, do nothing
		return co.ERROR_SUCCESS
	}
	ret, _, lerr := syscall.Syscall(proc.RegCloseKey.Addr(), 1,
		uintptr(hKey), 0, 0)
	if ret == 0 { // an error occurred
		return co.ERROR(lerr)
	}
	return co.ERROR_SUCCESS
}
