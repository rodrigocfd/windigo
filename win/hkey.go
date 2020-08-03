/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type HKEY HANDLE

func (hKey HKEY) RegCloseKey() {
	ret := hKey.regCloseKeyNoPanic()
	if ret != co.ERROR_SUCCESS {
		panic(ret.Format("RegCloseKey failed."))
	}
}

func (hKey HKEY) regCloseKeyNoPanic() co.ERROR {
	return freeNoPanic(HANDLE(hKey), proc.RegCloseKey)
}

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
	panic(lerr.Format("RegEnumValue failed."))
}

// Returns zero if key doesn't exist.
func RegOpenKeyEx(hKeyPredef co.HKEY, lpSubKey string, ulOptions co.REG_OPTION,
	samDesired co.KEY) HKEY {

	hKey := HKEY(0)
	ret, _, _ := syscall.Syscall6(proc.RegOpenKeyEx.Addr(), 5,
		uintptr(hKeyPredef), uintptr(unsafe.Pointer(StrToPtr(lpSubKey))),
		uintptr(ulOptions), uintptr(samDesired), uintptr(unsafe.Pointer(&hKey)),
		0)

	lerr := co.ERROR(ret)
	if lerr == co.ERROR_FILE_NOT_FOUND {
		return 0 // not found
	} else if lerr != co.ERROR_SUCCESS {
		panic(lerr.Format("RegOpenKeyEx failed."))
	}
	return hKey
}

func (hKey HKEY) RegQueryValueEx(lpValueName string, lpType *co.REG,
	lpData unsafe.Pointer, lpcbData *uint32) co.ERROR {

	ret, _, _ := syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(StrToPtr(lpValueName))), 0,
		uintptr(unsafe.Pointer(lpType)), uintptr(lpData),
		uintptr(unsafe.Pointer(lpcbData)))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_SUCCESS {
		hKey.regCloseKeyNoPanic() // free resource
		panic(lerr.Format("RegQueryValueEx failed."))
	}
	return lerr
}
