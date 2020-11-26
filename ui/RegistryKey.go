/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Used in OpenRegistryKey().
//
// Behavior of the registry key opening.
type REGKEY_MODE uint8

const (
	REGKEY_MODE_R  REGKEY_MODE = iota // Open a registry key for read only.
	REGKEY_MODE_RW                    // Open a registry key for read and write.
)

// Information about a registry value.
type RegistryValueInfo struct {
	DataType co.REG
	Name     string
	Size     int // Size in bytes, note that wide strings are uint16 chars.
}

//------------------------------------------------------------------------------

// Manages a registry key resource.
type RegistryKey struct {
	hKey win.HKEY
}

// Constructor.
//
// You must defer Close().
func OpenRegistryKey(
	keyPredef co.HKEY, subKey string,
	behavior REGKEY_MODE) (*RegistryKey, error) {

	oMode := co.KEY_READ
	if behavior == REGKEY_MODE_RW {
		oMode |= co.KEY_WRITE
	}

	hKey, wErr := win.RegOpenKeyEx(keyPredef, subKey, co.REG_OPTION_NONE, oMode)
	if wErr != nil {
		return nil, wErr
	}

	return &RegistryKey{
		hKey: hKey,
	}, nil
}

// Calls RegCloseKey() to free the resource.
func (me *RegistryKey) Close() {
	if me.hKey != 0 {
		me.hKey.RegCloseKey()
		me.hKey = win.HKEY(0)
	}
}

// Enumerates all values in this registry key.
func (me *RegistryKey) EnumValues() ([]*RegistryValueInfo, error) {
	retVals := make([]*RegistryValueInfo, 0)
	index := uint32(0)
	nameBufSz := uint32(64) // arbitrary
	nameBuf := make([]uint16, nameBufSz)
	dataType := co.REG_NONE
	dataBufSz := uint32(0)

	for {
		status, wErr := me.hKey.RegEnumValue(index, nameBuf, &nameBufSz,
			&dataType, nil, &dataBufSz)

		if wErr != nil {
			return nil, wErr

		} else if status == co.ERROR_SUCCESS { // we got this one, but there's more
			retVals = append(retVals, &RegistryValueInfo{
				DataType: dataType,
				Name:     syscall.UTF16ToString(nameBuf),
				Size:     int(dataBufSz),
			})
			index++

		} else if status == co.ERROR_NO_MORE_ITEMS { // we're done
			break

		} else if status == co.ERROR_MORE_DATA { // increase buffer size
			nameBufSz += 8 // arbitrary
			nameBuf = make([]uint16, nameBufSz)
		}
	}

	sort.Slice(retVals, func(i, j int) bool {
		return strings.ToUpper(retVals[i].Name) < strings.ToUpper(retVals[j].Name)
	})
	return retVals, nil
}

// Reads an uint32.
func (me *RegistryKey) ReadDword(valueName string) (int, error) {
	info, err := me.ValueInfo(valueName)
	if err != nil {
		return 0, err
	}

	if info.DataType != co.REG_DWORD {
		return 0, errors.New(
			fmt.Sprintf("Registry value %s is not an uint32, type is %d.",
				valueName, info.DataType))
	}

	dataBuf := uint32(0) // 4 bytes
	dataBufSize32 := uint32(info.Size)
	wErr := me.hKey.RegQueryValueEx(valueName, &info.DataType,
		unsafe.Pointer(&dataBuf), &dataBufSize32)
	if wErr != nil {
		return 0, wErr
	}

	return int(dataBuf), nil
}

// Reads a string.
func (me *RegistryKey) ReadString(valueName string) (string, error) {
	info, err := me.ValueInfo(valueName)
	if err != nil {
		return "", err
	}

	if info.DataType != co.REG_SZ {
		return "", errors.New(
			fmt.Sprintf("Registry value %s is not a string, type is %d.",
				valueName, info.DataType))
	}

	dataBuf := make([]uint16, info.Size/2) // retrieved size is in bytes
	dataBufSize32 := uint32(info.Size)
	wErr := me.hKey.RegQueryValueEx(valueName, &info.DataType,
		unsafe.Pointer(&dataBuf[0]), &dataBufSize32)
	if wErr != nil {
		return "", wErr
	}

	return syscall.UTF16ToString(dataBuf), nil
}

// Checks if a value exists within the key.
func (me *RegistryKey) ValueExists(valueName string) bool {
	return me.hKey.RegQueryValueEx(valueName, nil, nil, nil) != nil
}

// Retrieves information about a specific value.
func (me *RegistryKey) ValueInfo(
	valueName string) (*RegistryValueInfo, error) {

	dataType := co.REG_NONE
	dataBufSize := uint32(0)
	wErr := me.hKey.RegQueryValueEx(valueName, &dataType, nil, &dataBufSize)
	if wErr != nil {
		return nil, wErr
	}

	return &RegistryValueInfo{
		DataType: dataType,
		Name:     valueName,
		Size:     int(dataBufSize),
	}, nil
}
