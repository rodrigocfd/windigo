/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"encoding/binary"
	"fmt"
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

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyvaluew
func (hKey HKEY) RegDeleteKeyValue(lpSubKey, lpValueName string) error {
	ret, _, _ := syscall.Syscall(proc.RegDeleteKeyValue.Addr(), 3,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpSubKey))),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValueName))))
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		return NewWinError(co.ERROR(ret), "RegDeleteKeyValue")
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletevaluew
func (hKey HKEY) RegDeleteValue(lpValueName string) error {
	ret, _, _ := syscall.Syscall(proc.RegDeleteValue.Addr(), 2,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValueName))), 0)
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		return NewWinError(co.ERROR(ret), "RegDeleteValue")
	}
	return nil
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

// Supported return types:
//
// - []byte - REG_BINARY
// - uint32 - REG_DWORD
// - uint64 - REG_QWORD
// - string - REG_SZ
// - string - REG_EXPAND_SZ
// - []string - REG_MULTI_SZ
//
// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-reggetvaluew
func (hKey HKEY) RegGetValue(
	lpSubKey, lpValue string) (interface{}, co.REG, error) {

	dwFlags := co.RRF_RT_ANY | co.RRF_NOEXPAND
	lpType := co.REG(0)
	lpcbData := uint32(0)

	ret, _, _ := syscall.Syscall9(proc.RegGetValue.Addr(), 7,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpSubKey))),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValue))),
		uintptr(dwFlags), uintptr(unsafe.Pointer(&lpType)), 0,
		uintptr(unsafe.Pointer(&lpcbData)), 0, 0) // query type and size
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		return nil, co.REG_NONE, NewWinError(co.ERROR(ret), "RegGetValue")
	}

	if lpType == co.REG_NONE {
		return nil, lpType, nil // no value to query
	}

	lpData := make([]byte, lpcbData) // buffer to receive data

	ret, _, _ = syscall.Syscall9(proc.RegGetValue.Addr(), 7,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpSubKey))),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValue))),
		uintptr(dwFlags), uintptr(unsafe.Pointer(&lpType)),
		uintptr(unsafe.Pointer(&lpData[0])),
		uintptr(unsafe.Pointer(&lpcbData)), 0, 0) // query type and size
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		return nil, co.REG_NONE, NewWinError(co.ERROR(ret), "RegGetValue")
	}

	return hKey.outputValue(lpData, lpType)
}

// Supported return types:
//
// - []byte - REG_BINARY
// - uint32 - REG_DWORD
// - uint64 - REG_QWORD
// - string - REG_SZ
// - string - REG_EXPAND_SZ
// - []string - REG_MULTI_SZ
//
// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regqueryvalueexw
func (hKey HKEY) RegQueryValueEx(
	lpValueName string) (interface{}, co.REG, error) {

	lpType := co.REG(0)
	lpcbData := uint32(0)

	ret, _, _ := syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValueName))), 0,
		uintptr(unsafe.Pointer(&lpType)), 0,
		uintptr(unsafe.Pointer(&lpcbData))) // query type and size
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		return nil, co.REG_NONE, NewWinError(co.ERROR(ret), "RegQueryValueEx")
	}

	if lpType == co.REG_NONE {
		return nil, lpType, nil // no value to query
	}

	lpData := make([]byte, lpcbData) // buffer to receive data

	ret, _, _ = syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValueName))), 0,
		uintptr(unsafe.Pointer(&lpType)), uintptr(unsafe.Pointer(&lpData[0])),
		uintptr(unsafe.Pointer(&lpcbData))) // query value itself
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		return nil, co.REG_NONE, NewWinError(co.ERROR(ret), "RegQueryValueEx")
	}

	return hKey.outputValue(lpData, lpType)
}

// Key will be create if it doesn't exist. If new type is different from current
// type, new type will prevail.
//
// Will panic on a wrong data type. Supported types:
//
// - []byte - REG_BINARY
// - uint32 - REG_DWORD
// - uint64 - REG_QWORD
// - string - REG_SZ
// - string - REG_EXPAND_SZ
// - []string - REG_MULTI_SZ
//
// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsetkeyvaluew
func (hKey HKEY) RegSetKeyValue(
	lpSubKey, lpValueName string, dwType co.REG, lpData interface{}) error {

	var dataPtr *byte
	var dataSz uint32

	switch dwType {
	case co.REG_BINARY:
		binSlice := lpData.([]byte)
		dataSz = uint32(len(binSlice))
		dataPtr = &binSlice[0]

	case co.REG_DWORD:
		dataSz = uint32(unsafe.Sizeof(uint32(0)))
		binSlice := make([]byte, dataSz)
		binary.LittleEndian.PutUint32(binSlice, lpData.(uint32))
		dataPtr = &binSlice[0]

	case co.REG_QWORD:
		dataSz = uint32(unsafe.Sizeof(uint64(0)))
		binSlice := make([]byte, dataSz)
		binary.LittleEndian.PutUint64(binSlice, lpData.(uint64))
		dataPtr = &binSlice[0]

	case co.REG_SZ, co.REG_EXPAND_SZ:
		buf16 := Str.ToUint16Slice(lpData.(string)) // null-terminated
		dataSz = uint32(len(buf16) * 2)             // wide chars counted in bytes
		dataPtr = (*byte)(unsafe.Pointer(&buf16[0]))

	case co.REG_MULTI_SZ:
		buf16 := Str.ToUint16SliceMulti(lpData.([]string)) // double null-terminated
		dataSz = uint32(len(buf16) * 2)                    // wide chars counted in bytes
		dataPtr = (*byte)(unsafe.Pointer(&buf16[0]))

	default:
		panic(fmt.Sprintf("Unsupported registry type: %d.", dwType))
	}

	ret, _, _ := syscall.Syscall6(proc.RegSetKeyValue.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpSubKey))),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValueName))),
		uintptr(dwType), uintptr(unsafe.Pointer(dataPtr)), uintptr(dataSz))
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		return NewWinError(co.ERROR(ret), "RegSetKeyValue")
	}

	return nil
}

// Returns the converted registry value.
func (hKey HKEY) outputValue(
	lpData []byte, lpType co.REG) (interface{}, co.REG, error) {

	switch lpType {
	case co.REG_BINARY:
		return lpData, lpType, nil
	case co.REG_DWORD:
		return binary.LittleEndian.Uint32(lpData), lpType, nil
	case co.REG_QWORD:
		return binary.LittleEndian.Uint64(lpData), lpType, nil
	case co.REG_SZ, co.REG_EXPAND_SZ:
		return Str.FromUint16Ptr((*uint16)(unsafe.Pointer((&lpData[0])))), lpType, nil
	case co.REG_MULTI_SZ:
		return Str.FromUint16PtrMulti((*uint16)(unsafe.Pointer((&lpData[0])))), lpType, nil
	}

	panic(fmt.Sprintf("Unsupported registry type: %d.", lpType))
}
