/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Manages a registry key resource.
type RegistryKey struct {
	hKey win.HKEY
}

// Calls RegCloseKey and sets the HKEY to zero.
func (me *RegistryKey) Close() {
	if me.hKey != 0 {
		me.hKey.RegCloseKey()
		me.hKey = win.HKEY(0)
	}
}

func (me *RegistryKey) EnumValues() []win.RegistryValueInfo {
	return me.hKey.RegEnumValue()
}

func (me *RegistryKey) OpenForRead(keyPredef co.HKEY, subKey string) {
	me.hKey = win.RegOpenKeyEx(keyPredef, subKey,
		co.REG_OPTION_NONE, co.KEY_READ)
	if me.hKey == win.HKEY(0) {
		panic("Key doesn't exist.")
	}
}

// Retrieves data type and size.
func (me *RegistryKey) ValueInfo(valueName string) (co.REG, uint32) {
	dataType := co.REG_NONE
	dataBufSize := uint32(0)
	me.hKey.RegQueryValueEx(valueName, &dataType, 0, &dataBufSize)
	return dataType, dataBufSize
}

func (me *RegistryKey) ReadString(valueName string) string {
	dataType, dataBufSize := me.ValueInfo(valueName)
	if dataType != co.REG_SZ {
		panic(fmt.Sprintf("Registry value isn't string, type is %d.", dataType))
	}

	dataBuf := make([]uint16, dataBufSize/2) // returned size is in bytes, we've got wide chars
	me.hKey.RegQueryValueEx(valueName, &dataType,
		uintptr(unsafe.Pointer(&dataBuf[0])), &dataBufSize)
	return syscall.UTF16ToString(dataBuf)
}

func (me *RegistryKey) ReadUint32(valueName string) uint32 {
	dataType, dataBufSize := me.ValueInfo(valueName)
	if dataType != co.REG_DWORD {
		panic(fmt.Sprintf("Registry value isn't uint32, type is %d.", dataType))
	}

	dataBuf := uint32(0) // 4 bytes
	me.hKey.RegQueryValueEx(valueName, &dataType,
		uintptr(unsafe.Pointer(&dataBuf)), &dataBufSize)
	return dataBuf
}

func (me *RegistryKey) ValueExists(valueName string) bool {
	return me.hKey.RegQueryValueEx(valueName, nil, 0, nil) == co.ERROR_SUCCESS
}
