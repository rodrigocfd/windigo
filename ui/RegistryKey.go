/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"sort"
	"strings"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Manages a registry key resource.
type RegistryKey struct {
	hKey win.HKEY
}

// Data returned by RegistryKey.EnumValues().
type RegistryValueInfo struct {
	DataType co.REG
	Name     string
	Size     uint // in uint8, notice wide strings are uint16 chars
}

// Calls RegCloseKey and sets the HKEY to zero.
func (me *RegistryKey) Close() {
	if me.hKey != 0 {
		me.hKey.RegCloseKey()
		me.hKey = win.HKEY(0)
	}
}

// Enumerates all values in this registry key.
func (me *RegistryKey) EnumValues() []RegistryValueInfo {
	retVals := make([]RegistryValueInfo, 0)
	index := uint32(0)
	nameBufSz := uint32(64) // arbitrary
	nameBuf := make([]uint16, nameBufSz)
	dataType := co.REG_NONE
	dataBufSz := uint32(0)

	for {
		status, err := me.hKey.RegEnumValue(index, nameBuf, &nameBufSz,
			&dataType, nil, &dataBufSz)

		if err != nil {
			panic(err.Error())
		} else if status == co.ERROR_SUCCESS { // we got this one, but there's more
			retVals = append(retVals, RegistryValueInfo{
				DataType: dataType,
				Name:     syscall.UTF16ToString(nameBuf),
				Size:     uint(dataBufSz),
			})
			index++
		} else if status == co.ERROR_NO_MORE_ITEMS { // we're done
			break
		} else if status == co.ERROR_MORE_DATA { // increase buffer size
			nameBufSz += 4 // arbitrary
			nameBuf = make([]uint16, nameBufSz)
		}
	}

	sort.Slice(retVals, func(i, j int) bool {
		return strings.ToUpper(retVals[i].Name) < strings.ToUpper(retVals[j].Name)
	})
	return retVals
}

// Opens a registry key for reading.
func (me *RegistryKey) OpenForRead(
	keyPredef co.HKEY, subKey string) *win.WinError {

	var err *win.WinError
	me.hKey, err = win.RegOpenKeyEx(keyPredef, subKey,
		co.REG_OPTION_NONE, co.KEY_READ)
	return err
}

// Retrieves data type and size.
func (me *RegistryKey) ValueInfo(valueName string) (co.REG, uint) {
	dataType := co.REG_NONE
	dataBufSize := uint32(0)
	err := me.hKey.RegQueryValueEx(valueName, &dataType, nil, &dataBufSize)
	if err != nil {
		panic(err.Error())
	}
	return dataType, uint(dataBufSize)
}

// Reads a string. Panics if data type is different.
func (me *RegistryKey) ReadString(valueName string) string {
	dataType, dataBufSize := me.ValueInfo(valueName)
	if dataType != co.REG_SZ {
		panic(fmt.Sprintf("Registry value isn't string, type is %d.", dataType))
	}

	dataBuf := make([]uint16, dataBufSize/2) // returned size is in bytes, we've got wide chars
	dataBufSize32 := uint32(dataBufSize)
	if err := me.hKey.RegQueryValueEx(valueName, &dataType,
		unsafe.Pointer(&dataBuf[0]), &dataBufSize32); err != nil {
		panic(err.Error())
	}
	return syscall.UTF16ToString(dataBuf)
}

// Reads an uint32. Panics if data type is different.
func (me *RegistryKey) ReadDword(valueName string) uint {
	dataType, dataBufSize := me.ValueInfo(valueName)
	if dataType != co.REG_DWORD {
		panic(fmt.Sprintf("Registry value isn't uint32, type is %d.", dataType))
	}

	dataBuf := uint32(0) // 4 bytes
	dataBufSize32 := uint32(dataBufSize)
	if err := me.hKey.RegQueryValueEx(valueName, &dataType,
		unsafe.Pointer(&dataBuf), &dataBufSize32); err != nil {
		panic(err.Error())
	}
	return uint(dataBuf)
}

// Checks if a value exists within the key.
func (me *RegistryKey) ValueExists(valueName string) bool {
	return me.hKey.RegQueryValueEx(valueName, nil, nil, nil) != nil
}
