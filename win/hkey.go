/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"encoding/binary"
	"sort"
	"syscall"
	"unsafe"
	"windigo/co"
	proc "windigo/win/internal"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hkey
type HKEY HANDLE

// You must defer RegCloseKey().
//
// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regopenkeyexw
func RegOpenKeyEx(
	hKeyPredef co.HKEY, lpSubKey string,
	ulOptions co.REG_OPTION, samDesired co.KEY) (HKEY, error) {

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

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regclosekey
func (hKey HKEY) RegCloseKey() {
	if hKey != 0 {
		syscall.Syscall(proc.RegCloseKey.Addr(), 1,
			uintptr(hKey), 0, 0)
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenumvaluew
func (hKey HKEY) RegEnumValue() ([]string, error) {
	valueNames := make([]string, 0)
	dwIndex := uint32(0)
	valueNameSz := uint32(64) // arbitrary
	valueNameBuf := make([]uint16, valueNameSz)

	for {
		ret, _, _ := syscall.Syscall9(proc.RegEnumValue.Addr(), 8,
			uintptr(hKey), uintptr(dwIndex),
			uintptr(unsafe.Pointer(&valueNameBuf[0])),
			uintptr(unsafe.Pointer(&valueNameSz)), 0, 0, 0, 0, 0)

		lerr := co.ERROR(ret)
		if lerr == co.ERROR_SUCCESS { // we got this one, but there's more
			valueNames = append(valueNames, Str.FromUint16Slice(valueNameBuf))
			dwIndex++
		} else if lerr == co.ERROR_NO_MORE_ITEMS { // we're done
			break
		} else if lerr == co.ERROR_MORE_DATA { // increase buffer size
			valueNameSz += 8 // arbitrary
			valueNameBuf = make([]uint16, valueNameSz)
		} else {
			return nil, NewWinError(lerr, "RegEnumValue")
		}
	}

	sort.Strings(valueNames)
	return valueNames, nil
}

// Supported return types: []byte, uint32, uint64, string, []string.
//
// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regqueryvalueexw
func (hKey HKEY) RegQueryValueEx(lpValueName string) (interface{}, error) {
	lpType := co.REG(0)
	lpcbData := uint32(0)

	ret, _, _ := syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValueName))), 0,
		uintptr(unsafe.Pointer(&lpType)), 0,
		uintptr(unsafe.Pointer(&lpcbData))) // query type and size
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		return nil, NewWinError(co.ERROR(ret), "RegQueryValueEx")
	}

	if lpType == co.REG_NONE {
		return nil, nil // no value to query
	}

	lpData := make([]byte, lpcbData) // buffer to receive data

	ret, _, _ = syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValueName))), 0,
		uintptr(unsafe.Pointer(&lpType)), uintptr(unsafe.Pointer(&lpData[0])),
		uintptr(unsafe.Pointer(&lpcbData))) // query value itself
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		return nil, NewWinError(co.ERROR(ret), "RegQueryValueEx")
	}

	switch lpType {
	case co.REG_BINARY:
		return lpData, nil
	case co.REG_DWORD:
		return binary.LittleEndian.Uint32(lpData), nil
	case co.REG_QWORD:
		return binary.LittleEndian.Uint64(lpData), nil
	case co.REG_SZ:
		return Str.FromUint16Ptr((*uint16)(unsafe.Pointer((&lpData[0])))), nil
	case co.REG_MULTI_SZ:
		return Str.FromUint16PtrMulti((*uint16)(unsafe.Pointer((&lpData[0])))), nil
	}

	panic("Unsupported RegQueryValueEx type.")
}
