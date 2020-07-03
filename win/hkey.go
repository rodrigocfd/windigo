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

type HKEY HANDLE

func (hKey HKEY) RegCloseKey() {
	ret, _, _ := syscall.Syscall(proc.RegCloseKey.Addr(), 1,
		uintptr(hKey), 0, 0)
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegCloseKey failed: %d %s",
			ret, syscall.Errno(ret).Error()))
	}
}

func (hKey HKEY) RegEnumValue(dwIndex uint32,
	lpValueName []uint16, lpcchValueName *uint32, lpType *co.REG,
	lpData uintptr, lpcbData *uint32) co.ERROR {

	ret, _, _ := syscall.Syscall9(proc.RegEnumValue.Addr(), 8,
		uintptr(hKey), uintptr(dwIndex),
		uintptr(unsafe.Pointer(&lpValueName[0])),
		uintptr(unsafe.Pointer(lpcchValueName)), 0,
		uintptr(unsafe.Pointer(lpType)),
		lpData, uintptr(unsafe.Pointer(lpcbData)), 0)

	switch co.ERROR(ret) {
	case co.ERROR_SUCCESS:
		fallthrough
	case co.ERROR_NO_MORE_ITEMS:
		fallthrough
	case co.ERROR_MORE_DATA:
		return co.ERROR(ret)
	}

	panic(fmt.Sprintf("RegEnumValue failed: %d %s",
		ret, syscall.Errno(ret).Error()))
}

// Returns zero if key doesn't exist.
func RegOpenKeyEx(hKeyPredef co.HKEY, lpSubKey string, ulOptions co.REG_OPTION,
	samDesired co.KEY) HKEY {

	hKey := HKEY(0)
	ret, _, _ := syscall.Syscall6(proc.RegOpenKeyEx.Addr(), 5,
		uintptr(hKeyPredef), uintptr(unsafe.Pointer(StrToPtr(lpSubKey))),
		uintptr(ulOptions), uintptr(samDesired), uintptr(unsafe.Pointer(&hKey)),
		0)
	if co.ERROR(ret) == co.ERROR_FILE_NOT_FOUND {
		return HKEY(0) // not found
	} else if co.ERROR(ret) != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegOpenKeyEx failed: %d %s",
			ret, syscall.Errno(ret).Error()))
	}
	return hKey
}

func (hKey HKEY) RegQueryValueEx(lpValueName string, lpType *co.REG,
	lpData uintptr, lpcbData *uint32) co.ERROR {

	ret, _, _ := syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(StrToPtr(lpValueName))), 0,
		uintptr(unsafe.Pointer(lpType)), lpData,
		uintptr(unsafe.Pointer(lpcbData)))
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegQueryValueEx failed: %d %s",
			ret, syscall.Errno(ret).Error()))
	}
	return co.ERROR(ret)
}
