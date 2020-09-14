/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win/proc"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hkey
type HKEY HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regclosekey
func (hKey HKEY) RegCloseKey() {
	if hKey != 0 {
		syscall.Syscall(proc.RegCloseKey.Addr(), 1,
			uintptr(hKey), 0, 0)
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenumvaluew
//
// This function returns co.ERROR as status flag.
func (hKey HKEY) RegEnumValue(
	dwIndex uint32, lpValueName []uint16, lpcchValueName *uint32, lpType *co.REG,
	lpData unsafe.Pointer, lpcbData *uint32) (co.ERROR, *WinError) {

	ret, _, _ := syscall.Syscall9(proc.RegEnumValue.Addr(), 8,
		uintptr(hKey), uintptr(dwIndex),
		uintptr(unsafe.Pointer(&lpValueName[0])),
		uintptr(unsafe.Pointer(lpcchValueName)), 0,
		uintptr(unsafe.Pointer(lpType)),
		uintptr(lpData), uintptr(unsafe.Pointer(lpcbData)), 0)

	status := co.ERROR(ret)
	if status == co.ERROR_SUCCESS ||
		status == co.ERROR_NO_MORE_ITEMS ||
		status == co.ERROR_MORE_DATA {
		// These are not really errors.
		return status, nil
	}
	return status, NewWinError(status, "RegEnumValue") // any other status is an error
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regopenkeyexw
func RegOpenKeyEx(hKeyPredef co.HKEY, lpSubKey string, ulOptions co.REG_OPTION,
	samDesired co.KEY) (HKEY, *WinError) {

	hKey := HKEY(0)
	ret, _, _ := syscall.Syscall6(proc.RegOpenKeyEx.Addr(), 5,
		uintptr(hKeyPredef), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpSubKey))),
		uintptr(ulOptions), uintptr(samDesired), uintptr(unsafe.Pointer(&hKey)),
		0)
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		return HKEY(0), NewWinError(co.ERROR(ret), "RegOpenKeyEx")
	}
	return hKey, nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regqueryvalueexw
func (hKey HKEY) RegQueryValueEx(lpValueName string, lpType *co.REG,
	lpData unsafe.Pointer, lpcbData *uint32) *WinError {

	ret, _, _ := syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValueName))), 0,
		uintptr(unsafe.Pointer(lpType)), uintptr(lpData),
		uintptr(unsafe.Pointer(lpcbData)))
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		return NewWinError(co.ERROR(ret), "RegQueryValueEx")
	}
	return nil
}
